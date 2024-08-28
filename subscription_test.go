// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package polar_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/polarsource/polar-go"
	"github.com/polarsource/polar-go/internal/testutil"
	"github.com/polarsource/polar-go/option"
	"github.com/polarsource/polar-go/shared"
)

func TestSubscriptionNew(t *testing.T) {
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
	_, err := client.Subscriptions.New(context.TODO(), polar.SubscriptionNewParams{
		Email:     polar.F("dev@stainlessapi.com"),
		ProductID: polar.F("product_id"),
	})
	if err != nil {
		var apierr *polar.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSubscriptionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Subscriptions.List(context.TODO(), polar.SubscriptionListParams{
		Active:         polar.F(true),
		Limit:          polar.F(int64(1)),
		OrganizationID: polar.F[polar.SubscriptionListParamsOrganizationIDUnion](shared.UnionString("string")),
		Page:           polar.F(int64(1)),
		ProductID:      polar.F[polar.SubscriptionListParamsProductIDUnion](shared.UnionString("string")),
		Sorting:        polar.F([]polar.SubscriptionListParamsSorting{polar.SubscriptionListParamsSortingUser, polar.SubscriptionListParamsSorting - User, polar.SubscriptionListParamsSortingStatus}),
		Type:           polar.F[polar.SubscriptionListParamsTypeUnion](polar.SubscriptionListParamsTypeSubscriptionTierType(polar.SubscriptionListParamsTypeSubscriptionTierTypeFree)),
	})
	if err != nil {
		var apierr *polar.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSubscriptionExportWithOptionalParams(t *testing.T) {
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
	_, err := client.Subscriptions.Export(context.TODO(), polar.SubscriptionExportParams{
		OrganizationID: polar.F[polar.SubscriptionExportParamsOrganizationIDUnion](shared.UnionString("string")),
	})
	if err != nil {
		var apierr *polar.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSubscriptionImport(t *testing.T) {
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
	_, err := client.Subscriptions.Import(context.TODO(), polar.SubscriptionImportParams{
		File:           polar.F(io.Reader(bytes.NewBuffer([]byte("some file contents")))),
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
