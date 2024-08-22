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

func TestExternalOrganizationListWithOptionalParams(t *testing.T) {
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
	_, err := client.ExternalOrganizations.List(context.TODO(), polar.ExternalOrganizationListParams{
		Limit:          polar.F(int64(1)),
		Name:           polar.F[polar.ExternalOrganizationListParamsNameUnion](shared.UnionString("string")),
		OrganizationID: polar.F[polar.ExternalOrganizationListParamsOrganizationIDUnion](shared.UnionString("string")),
		Page:           polar.F(int64(1)),
		Platform:       polar.F[polar.ExternalOrganizationListParamsPlatformUnion](polar.ExternalOrganizationListParamsPlatformPlatforms(polar.ExternalOrganizationListParamsPlatformPlatformsGitHub)),
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
