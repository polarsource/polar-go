package hooks

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/polarsource/polar-go/internal/utils"
	"github.com/polarsource/polar-go/models/components"
)

const (
	BatchSize     = 10
	FlushInterval = 5 * time.Second
)

type batchedRequestMarkerType struct{}

var batchedRequestMarker = batchedRequestMarkerType{}

type EventBatcher struct {
	mu            sync.Mutex
	events        []components.Events
	timer         *time.Timer
	client        HTTPClient
	baseURL       string
	security      func(context.Context) (interface{}, error)
	flushCallback func([]components.Events) error
}

var _ beforeRequestHook = (*EventBatcher)(nil)
var _ afterSuccessHook = (*EventBatcher)(nil)

func NewEventBatcher() *EventBatcher {
	eb := &EventBatcher{
		events: make([]components.Events, 0, BatchSize),
	}
	eb.resetTimer()
	return eb
}

func (eb *EventBatcher) BeforeRequest(hookCtx BeforeRequestContext, req *http.Request) (*http.Request, error) {
	if hookCtx.OperationID != "events:ingest" {
		return req, nil
	}

	eb.mu.Lock()
	if eb.client == nil {
		eb.client = hookCtx.SDKConfiguration.Client
		eb.baseURL = hookCtx.BaseURL
	}
	if eb.security == nil {
		eb.security = hookCtx.SecuritySource
	}
	eb.mu.Unlock()

	if req.Body == nil || req.Body == http.NoBody {
		return req, nil
	}

	bodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		return req, err
	}
	req.Body.Close()

	var ingestRequest components.EventsIngest
	if err := json.Unmarshal(bodyBytes, &ingestRequest); err != nil {
		req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		return req, nil
	}

	eb.mu.Lock()
	eb.events = append(eb.events, ingestRequest.Events...)
	eventCount := len(ingestRequest.Events)
	shouldFlush := len(eb.events) >= BatchSize
	eb.mu.Unlock()

	if shouldFlush {
		if err := eb.Flush(); err != nil {
			req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			return req, err
		}
	}

	ctx := context.WithValue(req.Context(), batchedRequestMarker, eventCount)
	req = req.WithContext(ctx)

	return req, nil
}

func (eb *EventBatcher) AfterSuccess(hookCtx AfterSuccessContext, res *http.Response) (*http.Response, error) {
	if res.Request == nil {
		return res, nil
	}

	if eventCount, ok := res.Request.Context().Value(batchedRequestMarker).(int); ok {
		syntheticResponse := &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(bytes.NewBuffer([]byte(fmt.Sprintf(`{"inserted":%d}`, eventCount)))),
			Request:    res.Request,
		}
		syntheticResponse.Header.Set("Content-Type", "application/json")

		res.Body.Close()
		return syntheticResponse, nil
	}

	return res, nil
}

func (eb *EventBatcher) resetTimer() {
	if eb.timer != nil {
		eb.timer.Stop()
	}
	eb.timer = time.AfterFunc(FlushInterval, func() {
		_ = eb.Flush()
	})
}

func (eb *EventBatcher) Flush() error {
	eb.mu.Lock()
	defer eb.mu.Unlock()

	if len(eb.events) == 0 {
		return nil
	}

	if eb.timer != nil {
		eb.timer.Stop()
	}

	eventsToSend := make([]components.Events, len(eb.events))
	copy(eventsToSend, eb.events)
	eb.events = eb.events[:0]

	eb.resetTimer()

	if eb.flushCallback != nil {
		return eb.flushCallback(eventsToSend)
	}

	return eb.sendEvents(eventsToSend)
}

func (eb *EventBatcher) sendEvents(events []components.Events) error {
	if eb.client == nil {
		return fmt.Errorf("client not initialized")
	}

	ingestRequest := components.EventsIngest{
		Events: events,
	}

	bodyBytes, err := json.Marshal(ingestRequest)
	if err != nil {
		return fmt.Errorf("failed to marshal events: %w", err)
	}

	url := eb.baseURL + "/v1/events/ingest"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	if eb.security != nil {
		ctx := context.Background()
		sec, err := eb.security(ctx)
		if err == nil && sec != nil {
			if err := utils.PopulateSecurity(ctx, req, eb.security); err != nil {
				return fmt.Errorf("failed to apply security: %w", err)
			}
		}
	}

	resp, err := eb.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send events: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API returned status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	return nil
}
