package hooks

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/polarsource/polar-go/models/components"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockServer struct {
	server       *httptest.Server
	mu           sync.Mutex
	requests     []components.EventsIngest
	requestCount int32
}

func newMockServer(t *testing.T) *mockServer {
	ms := &mockServer{
		requests: make([]components.EventsIngest, 0),
	}

	ms.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/events/ingest" {
			t.Logf("Unexpected path: %s", r.URL.Path)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		atomic.AddInt32(&ms.requestCount, 1)

		body, err := io.ReadAll(r.Body)
		require.NoError(t, err)

		var ingest components.EventsIngest
		err = json.Unmarshal(body, &ingest)
		require.NoError(t, err)

		ms.mu.Lock()
		ms.requests = append(ms.requests, ingest)
		ms.mu.Unlock()
		response := components.EventsIngestResponse{
			Inserted: int64(len(ingest.Events)),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(response)
	}))

	return ms
}

func (ms *mockServer) getRequests() []components.EventsIngest {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	return append([]components.EventsIngest{}, ms.requests...)
}

func (ms *mockServer) getRequestCount() int {
	return int(atomic.LoadInt32(&ms.requestCount))
}

func (ms *mockServer) close() {
	ms.server.Close()
}

func createTestEvent(name string) components.Events {
	timestamp := time.Now()
	return components.CreateEventsEventCreateCustomer(components.EventCreateCustomer{
		Name:       name,
		CustomerID: "test-customer-id",
		Timestamp:  &timestamp,
	})
}

func TestEventBatcher_BatchSizeTrigger(t *testing.T) {
	server := newMockServer(t)
	defer server.close()

	batcher := NewEventBatcher()
	batcher.client = server.server.Client()
	batcher.baseURL = server.server.URL
	batcher.security = func(ctx context.Context) (interface{}, error) { return nil, nil }

	hookCtx := BeforeRequestContext{
		HookContext: HookContext{
			OperationID: "events:ingest",
		},
	}

	afterSuccessCtx := AfterSuccessContext{
		HookContext: HookContext{
			OperationID: "events:ingest",
		},
	}

	for i := 0; i < 9; i++ {
		ingest := components.EventsIngest{
			Events: []components.Events{createTestEvent("event-" + string(rune(i+'0')))},
		}
		body, _ := json.Marshal(ingest)
		req, _ := http.NewRequest("POST", server.server.URL+"/v1/events/ingest", io.NopCloser(bytes.NewBuffer(body)))
		modifiedReq, err := batcher.BeforeRequest(hookCtx, req)
		assert.NoError(t, err)
		assert.NotNil(t, modifiedReq.Context().Value(batchedRequestMarker))

		dummyResp := &http.Response{
			StatusCode: http.StatusOK,
			Request:    modifiedReq,
			Body:       io.NopCloser(bytes.NewBuffer([]byte(`{"inserted":0}`))),
		}
		finalResp, err := batcher.AfterSuccess(afterSuccessCtx, dummyResp)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, finalResp.StatusCode)
		finalResp.Body.Close()
	}

	time.Sleep(100 * time.Millisecond)
	assert.Equal(t, 0, server.getRequestCount(), "Should not flush before reaching batch size")

	ingest := components.EventsIngest{
		Events: []components.Events{createTestEvent("event-9")},
	}
	body, _ := json.Marshal(ingest)
	req, _ := http.NewRequest("POST", server.server.URL+"/v1/events/ingest", io.NopCloser(bytes.NewBuffer(body)))
	modifiedReq, err := batcher.BeforeRequest(hookCtx, req)
	assert.NoError(t, err)

	dummyResp := &http.Response{
		StatusCode: http.StatusOK,
		Request:    modifiedReq,
		Body:       io.NopCloser(bytes.NewBuffer([]byte(`{"inserted":0}`))),
	}
	finalResp, err := batcher.AfterSuccess(afterSuccessCtx, dummyResp)
	assert.NoError(t, err)
	finalResp.Body.Close()

	time.Sleep(200 * time.Millisecond)

	assert.Equal(t, 1, server.getRequestCount(), "Should have flushed once")
	requests := server.getRequests()
	require.Len(t, requests, 1)
	assert.Len(t, requests[0].Events, 10, "Should have batched 10 events")
}

func TestEventBatcher_TimeoutTrigger(t *testing.T) {
	server := newMockServer(t)
	defer server.close()

	batcher := NewEventBatcher()
	batcher.client = server.server.Client()
	batcher.baseURL = server.server.URL
	batcher.security = func(ctx context.Context) (interface{}, error) { return nil, nil }

	hookCtx := BeforeRequestContext{
		HookContext: HookContext{
			OperationID: "events:ingest",
		},
	}

	afterSuccessCtx := AfterSuccessContext{
		HookContext: HookContext{
			OperationID: "events:ingest",
		},
	}

	// Send 5 events (less than batch size)
	for i := 0; i < 5; i++ {
		ingest := components.EventsIngest{
			Events: []components.Events{createTestEvent("event-" + string(rune(i+'0')))},
		}
		body, _ := json.Marshal(ingest)
		req, _ := http.NewRequest("POST", server.server.URL+"/v1/events/ingest", io.NopCloser(bytes.NewBuffer(body)))
		modifiedReq, _ := batcher.BeforeRequest(hookCtx, req)
		dummyResp := &http.Response{StatusCode: http.StatusOK, Request: modifiedReq, Body: io.NopCloser(bytes.NewBuffer([]byte(`{"inserted":0}`)))}
		finalResp, _ := batcher.AfterSuccess(afterSuccessCtx, dummyResp)
		finalResp.Body.Close()
	}

	time.Sleep(100 * time.Millisecond)
	assert.Equal(t, 0, server.getRequestCount(), "Should not flush immediately")

	time.Sleep(5*time.Second + 500*time.Millisecond)
	assert.Equal(t, 1, server.getRequestCount(), "Should have flushed after timeout")
	requests := server.getRequests()
	require.Len(t, requests, 1)
	assert.Len(t, requests[0].Events, 5, "Should have batched 5 events")
}

func TestEventBatcher_ManualFlush(t *testing.T) {
	server := newMockServer(t)
	defer server.close()

	batcher := NewEventBatcher()
	batcher.client = server.server.Client()
	batcher.baseURL = server.server.URL
	batcher.security = func(ctx context.Context) (interface{}, error) { return nil, nil }

	hookCtx := BeforeRequestContext{
		HookContext: HookContext{
			OperationID: "events:ingest",
		},
	}

	afterSuccessCtx := AfterSuccessContext{
		HookContext: HookContext{
			OperationID: "events:ingest",
		},
	}

	// Send 3 events
	for i := 0; i < 3; i++ {
		ingest := components.EventsIngest{
			Events: []components.Events{createTestEvent("event-" + string(rune(i+'0')))},
		}
		body, _ := json.Marshal(ingest)
		req, _ := http.NewRequest("POST", server.server.URL+"/v1/events/ingest", io.NopCloser(bytes.NewBuffer(body)))
		modifiedReq, _ := batcher.BeforeRequest(hookCtx, req)
		dummyResp := &http.Response{StatusCode: http.StatusOK, Request: modifiedReq, Body: io.NopCloser(bytes.NewBuffer([]byte(`{"inserted":0}`)))}
		finalResp, _ := batcher.AfterSuccess(afterSuccessCtx, dummyResp)
		finalResp.Body.Close()
	}

	err := batcher.Flush()
	assert.NoError(t, err)

	time.Sleep(100 * time.Millisecond)
	assert.Equal(t, 1, server.getRequestCount(), "Should have flushed")
	requests := server.getRequests()
	require.Len(t, requests, 1)
	assert.Len(t, requests[0].Events, 3, "Should have batched 3 events")
	err = batcher.Flush()
	assert.NoError(t, err)
	assert.Equal(t, 1, server.getRequestCount(), "Should not flush again")
}

func TestEventBatcher_ConcurrentAccess(t *testing.T) {
	server := newMockServer(t)
	defer server.close()

	batcher := NewEventBatcher()
	batcher.client = server.server.Client()
	batcher.baseURL = server.server.URL
	batcher.security = func(ctx context.Context) (interface{}, error) { return nil, nil }

	hookCtx := BeforeRequestContext{
		HookContext: HookContext{
			OperationID: "events:ingest",
		},
	}

	afterSuccessCtx := AfterSuccessContext{
		HookContext: HookContext{
			OperationID: "events:ingest",
		},
	}

	// Send 20 events concurrently from multiple goroutines
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			ingest := components.EventsIngest{
				Events: []components.Events{createTestEvent("event-" + string(rune(idx+'0')))},
			}
			body, _ := json.Marshal(ingest)
			req, _ := http.NewRequest("POST", server.server.URL+"/v1/events/ingest", io.NopCloser(bytes.NewBuffer(body)))
			modifiedReq, _ := batcher.BeforeRequest(hookCtx, req)
			dummyResp := &http.Response{StatusCode: http.StatusOK, Request: modifiedReq, Body: io.NopCloser(bytes.NewBuffer([]byte(`{"inserted":0}`)))}
			finalResp, _ := batcher.AfterSuccess(afterSuccessCtx, dummyResp)
			if finalResp != nil {
				finalResp.Body.Close()
			}
		}(i)
	}

	wg.Wait()

	time.Sleep(100 * time.Millisecond)
	err := batcher.Flush()
	assert.NoError(t, err)

	time.Sleep(200 * time.Millisecond)
	requests := server.getRequests()
	totalEvents := 0
	for _, req := range requests {
		totalEvents += len(req.Events)
	}
	assert.Equal(t, 20, totalEvents, "All events should have been sent")
}

func TestEventBatcher_EmptyFlush(t *testing.T) {
	server := newMockServer(t)
	defer server.close()

	batcher := NewEventBatcher()
	batcher.client = server.server.Client()
	batcher.baseURL = server.server.URL
	batcher.security = func(ctx context.Context) (interface{}, error) { return nil, nil }
	err := batcher.Flush()
	assert.NoError(t, err)

	time.Sleep(100 * time.Millisecond)
	assert.Equal(t, 0, server.getRequestCount(), "Should not send request when no events")
}

func TestEventBatcher_NonIngestOperations(t *testing.T) {
	server := newMockServer(t)
	defer server.close()

	batcher := NewEventBatcher()
	batcher.client = server.server.Client()
	batcher.baseURL = server.server.URL
	batcher.security = func(ctx context.Context) (interface{}, error) { return nil, nil }
	hookCtx := BeforeRequestContext{
		HookContext: HookContext{
			OperationID: "events:list",
		},
	}

	req, _ := http.NewRequest("GET", server.server.URL+"/v1/events/", nil)
	modifiedReq, err := batcher.BeforeRequest(hookCtx, req)
	assert.NoError(t, err)
	assert.Equal(t, req, modifiedReq, "Request should not be modified for non-ingest operations")
}

func TestEventBatcher_DisabledByDefault(t *testing.T) {
	originalGlobalBatcher := globalEventBatcher
	defer func() {
		globalEventBatcher = originalGlobalBatcher
	}()

	globalEventBatcher = nil

	err := FlushEvents()
	assert.NoError(t, err, "FlushEvents should be a safe no-op when batching is disabled")

	assert.Nil(t, globalEventBatcher, "globalEventBatcher should remain nil when batching is disabled")
}

func TestEventBatcher_EnvironmentVariableCheck(t *testing.T) {
	tests := []struct {
		name         string
		envValue     string
		shouldEnable bool
	}{
		{
			name:         "Batching enabled with true",
			envValue:     "true",
			shouldEnable: true,
		},
		{
			name:         "Batching disabled with empty",
			envValue:     "",
			shouldEnable: false,
		},
		{
			name:         "Batching disabled with false",
			envValue:     "false",
			shouldEnable: false,
		},
		{
			name:         "Batching disabled with random value",
			envValue:     "yes",
			shouldEnable: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			originalBatcher := globalEventBatcher
			defer func() {
				globalEventBatcher = originalBatcher
				os.Unsetenv("POLAR_ENABLE_EVENT_BATCHING")
			}()

			if tt.envValue != "" {
				os.Setenv("POLAR_ENABLE_EVENT_BATCHING", tt.envValue)
			} else {
				os.Unsetenv("POLAR_ENABLE_EVENT_BATCHING")
			}

			hooks := New()
			initHooks(hooks)

			if tt.shouldEnable {
				assert.NotNil(t, globalEventBatcher, "globalEventBatcher should be initialized when env var is 'true'")
			} else {
				assert.Nil(t, globalEventBatcher, "globalEventBatcher should be nil when env var is not 'true'")
			}

			err := FlushEvents()
			assert.NoError(t, err, "FlushEvents should always be safe to call")

			globalEventBatcher = nil
		})
	}
}
