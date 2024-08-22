// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package polar_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/polar-go"
	"github.com/stainless-sdks/polar-go/internal/testutil"
	"github.com/stainless-sdks/polar-go/option"
	"github.com/stainless-sdks/polar-go/shared"
)

func TestMetricListWithOptionalParams(t *testing.T) {
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
	_, err := client.Metrics.List(context.TODO(), polar.MetricListParams{
		EndDate:          polar.F(time.Now()),
		Interval:         polar.F(polar.MetricListParamsIntervalYear),
		StartDate:        polar.F(time.Now()),
		OrganizationID:   polar.F[polar.MetricListParamsOrganizationIDUnion](shared.UnionString("string")),
		ProductID:        polar.F[polar.MetricListParamsProductIDUnion](shared.UnionString("string")),
		ProductPriceType: polar.F[polar.MetricListParamsProductPriceTypeUnion](polar.MetricListParamsProductPriceTypeProductPriceType(polar.MetricListParamsProductPriceTypeProductPriceTypeOneTime)),
	})
	if err != nil {
		var apierr *polar.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
