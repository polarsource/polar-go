// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package polar_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/polar-go"
	"github.com/stainless-sdks/polar-go/internal/testutil"
	"github.com/stainless-sdks/polar-go/option"
)

func TestWebhookEndpointNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := polar.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Webhooks.Endpoints.New(context.TODO(), polar.WebhookEndpointNewParams{
		Events:         polar.F([]polar.WebhookEndpointNewParamsEvent{polar.WebhookEndpointNewParamsEventOrderCreated, polar.WebhookEndpointNewParamsEventSubscriptionCreated, polar.WebhookEndpointNewParamsEventSubscriptionUpdated}),
		Format:         polar.F(polar.WebhookEndpointNewParamsFormatRaw),
		Secret:         polar.F("f_z6mfSpxkjogyw3FkA2aH2gYE5huxruNf34MpdWMcA"),
		URL:            polar.F("https://webhook.site/cb791d80-f26e-4f8c-be88-6e56054192b0"),
		OrganizationID: polar.F("organization_id"),
	})
	if err != nil {
		var apierr *polar.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWebhookEndpointGet(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := polar.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Webhooks.Endpoints.Get(context.TODO(), "id")
	if err != nil {
		var apierr *polar.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWebhookEndpointUpdateWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := polar.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Webhooks.Endpoints.Update(
		context.TODO(),
		"id",
		polar.WebhookEndpointUpdateParams{
			Events: polar.F([]polar.WebhookEndpointUpdateParamsEvent{polar.WebhookEndpointUpdateParamsEventOrderCreated, polar.WebhookEndpointUpdateParamsEventSubscriptionCreated, polar.WebhookEndpointUpdateParamsEventSubscriptionUpdated}),
			Format: polar.F(polar.WebhookEndpointUpdateParamsFormatRaw),
			Secret: polar.F("f_z6mfSpxkjogyw3FkA2aH2gYE5huxruNf34MpdWMcA"),
			URL:    polar.F("https://webhook.site/cb791d80-f26e-4f8c-be88-6e56054192b0"),
		},
	)
	if err != nil {
		var apierr *polar.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWebhookEndpointListWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := polar.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Webhooks.Endpoints.List(context.TODO(), polar.WebhookEndpointListParams{
		Limit:          polar.F(int64(1)),
		OrganizationID: polar.F("organization_id"),
		Page:           polar.F(int64(1)),
		UserID:         polar.F("user_id"),
	})
	if err != nil {
		var apierr *polar.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWebhookEndpointDelete(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := polar.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	err := client.Webhooks.Endpoints.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *polar.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
