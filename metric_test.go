// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package polar_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/polarsource/polar-go"
	"github.com/polarsource/polar-go/internal/testutil"
	"github.com/polarsource/polar-go/option"
	"github.com/polarsource/polar-go/shared"
)

func TestMetricGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Metrics.Get(context.TODO(), polar.MetricGetParams{
		EndDate:          polar.F(time.Now()),
		Interval:         polar.F(polar.MetricGetParamsIntervalYear),
		StartDate:        polar.F(time.Now()),
		OrganizationID:   polar.F[polar.MetricGetParamsOrganizationIDUnion](shared.UnionString("string")),
		ProductID:        polar.F[polar.MetricGetParamsProductIDUnion](shared.UnionString("string")),
		ProductPriceType: polar.F[polar.MetricGetParamsProductPriceTypeUnion](polar.MetricGetParamsProductPriceTypeProductPriceType(polar.MetricGetParamsProductPriceTypeProductPriceTypeOneTime)),
	})
	if err != nil {
		var apierr *polar.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
