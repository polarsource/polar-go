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
	"github.com/polarsource/polar-go/shared"
)

func TestRepositoryGet(t *testing.T) {
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
	_, err := client.Repositories.Get(context.TODO(), "id")
	if err != nil {
		var apierr *polar.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRepositoryUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Repositories.Update(
		context.TODO(),
		"id",
		polar.RepositoryUpdateParams{
			ProfileSettings: polar.F(polar.RepositoryUpdateParamsProfileSettings{
				CoverImageURL:                polar.F("cover_image_url"),
				Description:                  polar.F("description"),
				FeaturedOrganizations:        polar.F([]string{"string", "string", "string"}),
				HighlightedSubscriptionTiers: polar.F([]string{"string", "string", "string"}),
				Links:                        polar.F([]string{"https://example.com", "https://example.com", "https://example.com"}),
				SetCoverImageURL:             polar.F(true),
				SetDescription:               polar.F(true),
			}),
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

func TestRepositoryListWithOptionalParams(t *testing.T) {
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
	_, err := client.Repositories.List(context.TODO(), polar.RepositoryListParams{
		ExternalOrganizationName: polar.F[polar.RepositoryListParamsExternalOrganizationNameUnion](shared.UnionString("string")),
		IsPrivate:                polar.F(true),
		Limit:                    polar.F(int64(1)),
		Name:                     polar.F[polar.RepositoryListParamsNameUnion](shared.UnionString("string")),
		OrganizationID:           polar.F[polar.RepositoryListParamsOrganizationIDUnion](shared.UnionString("string")),
		Page:                     polar.F(int64(1)),
		Platform:                 polar.F[polar.RepositoryListParamsPlatformUnion](polar.RepositoryListParamsPlatformPlatforms(polar.RepositoryListParamsPlatformPlatformsGitHub)),
		Sorting:                  polar.F([]polar.RepositoryListParamsSorting{polar.RepositoryListParamsSortingCreatedAt, polar.RepositoryListParamsSorting - CreatedAt, polar.RepositoryListParamsSortingName}),
	})
	if err != nil {
		var apierr *polar.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
