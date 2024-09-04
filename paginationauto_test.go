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

func TestAutoPagination(t *testing.T) {
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
	iter := client.Products.ListAutoPaging(context.TODO(), polar.ProductListParams{})
	// Prism mock isn't going to give us real pagination
	for i := 0; i < 3 && iter.Next(); i++ {
		product := iter.Current()
		t.Logf("%+v\n", product.ID)
	}
	if err := iter.Err(); err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
