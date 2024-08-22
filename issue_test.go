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

func TestIssueGet(t *testing.T) {
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
	_, err := client.Issues.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *polar.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestIssueUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Issues.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		polar.IssueUpdateParams{
			FundingGoal: polar.F(polar.IssueUpdateParamsFundingGoal{
				Amount:   polar.F(int64(0)),
				Currency: polar.F("currency"),
			}),
			SetUpfrontSplitToContributors: polar.F(true),
			UpfrontSplitToContributors:    polar.F(int64(0)),
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

func TestIssueListWithOptionalParams(t *testing.T) {
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
	_, err := client.Issues.List(context.TODO(), polar.IssueListParams{
		ExternalOrganizationName: polar.F[polar.IssueListParamsExternalOrganizationNameUnion](shared.UnionString("string")),
		IsBadged:                 polar.F(true),
		Limit:                    polar.F(int64(1)),
		Number:                   polar.F[polar.IssueListParamsNumberUnion](shared.UnionInt(int64(0))),
		OrganizationID:           polar.F[polar.IssueListParamsOrganizationIDUnion](shared.UnionString("string")),
		Page:                     polar.F(int64(1)),
		Platform:                 polar.F[polar.IssueListParamsPlatformUnion](polar.IssueListParamsPlatformPlatforms(polar.IssueListParamsPlatformPlatformsGitHub)),
		RepositoryName:           polar.F[polar.IssueListParamsRepositoryNameUnion](shared.UnionString("string")),
		Sorting:                  polar.F([]string{"string", "string", "string"}),
	})
	if err != nil {
		var apierr *polar.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestIssueConfirmSolved(t *testing.T) {
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
	_, err := client.Issues.ConfirmSolved(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		polar.IssueConfirmSolvedParams{
			Splits: polar.F([]polar.IssueConfirmSolvedParamsSplit{{
				ShareThousands: polar.F(int64(0)),
				GitHubUsername: polar.F("github_username"),
				OrganizationID: polar.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			}, {
				ShareThousands: polar.F(int64(0)),
				GitHubUsername: polar.F("github_username"),
				OrganizationID: polar.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			}, {
				ShareThousands: polar.F(int64(0)),
				GitHubUsername: polar.F("github_username"),
				OrganizationID: polar.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			}}),
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
