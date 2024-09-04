// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package polar_test

import (
	"context"
	"os"
	"testing"

	"github.com/polarsource/polar-go"
	"github.com/polarsource/polar-go/internal/testutil"
	"github.com/polarsource/polar-go/option"
)

func TestUsage(t *testing.T) {
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
	checkout, err := client.Checkouts.New(context.TODO(), polar.CheckoutNewParams{
		ProductPriceID: polar.F("product_price_id"),
		SuccessURL:     polar.F("https://example.com"),
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", checkout.ID)
}
