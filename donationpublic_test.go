// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package polar_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/polarsource/polar-go"
	"github.com/polarsource/polar-go/internal/testutil"
	"github.com/polarsource/polar-go/option"
)

func TestDonationPublicSearchWithOptionalParams(t *testing.T) {
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
	_, err := client.Donations.Public.Search(context.TODO(), polar.DonationPublicSearchParams{
		OrganizationID: polar.F("organization_id"),
		Limit:          polar.F(int64(1)),
		Page:           polar.F(int64(1)),
		Sorting:        polar.F([]string{"string", "string", "string"}),
	})
	if err != nil {
		var apierr *polar.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
