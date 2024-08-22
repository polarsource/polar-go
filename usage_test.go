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
	accountNewResponse, err := client.Accounts.New(context.TODO(), polar.AccountNewParams{
		AccountType: polar.F(polar.AccountNewParamsAccountTypeStripe),
		Country:     polar.F("xx"),
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", accountNewResponse.ID)
}
