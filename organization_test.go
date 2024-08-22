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
)

func TestOrganizationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Organizations.New(context.TODO(), polar.OrganizationNewParams{
		Name:             polar.F("xxx"),
		Slug:             polar.F("xxx"),
		AvatarURL:        polar.F("https://example.com"),
		DonationsEnabled: polar.F(true),
		FeatureSettings: polar.F(polar.OrganizationNewParamsFeatureSettings{
			ArticlesEnabled:      polar.F(true),
			IssueFundingEnabled:  polar.F(true),
			SubscriptionsEnabled: polar.F(true),
		}),
	})
	if err != nil {
		var apierr *polar.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOrganizationGet(t *testing.T) {
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
	_, err := client.Organizations.Get(context.TODO(), "id")
	if err != nil {
		var apierr *polar.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOrganizationUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Organizations.Update(
		context.TODO(),
		"id",
		polar.OrganizationUpdateParams{
			AvatarURL:                         polar.F("https://example.com"),
			BillingEmail:                      polar.F("billing_email"),
			DefaultBadgeCustomContent:         polar.F("default_badge_custom_content"),
			DefaultUpfrontSplitToContributors: polar.F(int64(0)),
			DonationsEnabled:                  polar.F(true),
			FeatureSettings: polar.F(polar.OrganizationUpdateParamsFeatureSettings{
				ArticlesEnabled:      polar.F(true),
				IssueFundingEnabled:  polar.F(true),
				SubscriptionsEnabled: polar.F(true),
			}),
			Name:                        polar.F("xxx"),
			PerUserMonthlySpendingLimit: polar.F(int64(0)),
			PledgeBadgeShowAmount:       polar.F(true),
			PledgeMinimumAmount:         polar.F(int64(0)),
			ProfileSettings: polar.F(polar.OrganizationUpdateParamsProfileSettings{
				Description:           polar.F("description"),
				FeaturedOrganizations: polar.F([]string{"string", "string", "string"}),
				FeaturedProjects:      polar.F([]string{"string", "string", "string"}),
				Links:                 polar.F([]string{"https://example.com", "https://example.com", "https://example.com"}),
				Subscribe: polar.F(polar.OrganizationUpdateParamsProfileSettingsSubscribe{
					CountFree: polar.F(true),
					Promote:   polar.F(true),
					ShowCount: polar.F(true),
				}),
			}),
			TotalMonthlySpendingLimit: polar.F(int64(0)),
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

func TestOrganizationListWithOptionalParams(t *testing.T) {
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
	_, err := client.Organizations.List(context.TODO(), polar.OrganizationListParams{
		IsMember: polar.F(true),
		Limit:    polar.F(int64(1)),
		Page:     polar.F(int64(1)),
		Slug:     polar.F("slug"),
		Sorting:  polar.F([]string{"string", "string", "string"}),
	})
	if err != nil {
		var apierr *polar.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
