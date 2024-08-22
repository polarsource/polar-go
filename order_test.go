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
	"github.com/stainless-sdks/polar-go/shared"
)

func TestOrderGet(t *testing.T) {
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
	_, err := client.Orders.Get(context.TODO(), "id")
	if err != nil {
		var apierr *polar.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOrderListWithOptionalParams(t *testing.T) {
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
	_, err := client.Orders.List(context.TODO(), polar.OrderListParams{
		Limit:            polar.F(int64(1)),
		OrganizationID:   polar.F[polar.OrderListParamsOrganizationIDUnion](shared.UnionString("string")),
		Page:             polar.F(int64(1)),
		ProductID:        polar.F[polar.OrderListParamsProductIDUnion](shared.UnionString("string")),
		ProductPriceType: polar.F[polar.OrderListParamsProductPriceTypeUnion](polar.OrderListParamsProductPriceTypeProductPriceType(polar.OrderListParamsProductPriceTypeProductPriceTypeOneTime)),
		Sorting:          polar.F([]string{"string", "string", "string"}),
		UserID:           polar.F[polar.OrderListParamsUserIDUnion](shared.UnionString("string")),
	})
	if err != nil {
		var apierr *polar.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
