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

func TestUserBenefitGet(t *testing.T) {
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
	_, err := client.Users.Benefits.Get(context.TODO(), "id")
	if err != nil {
		var apierr *polar.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserBenefitListWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.Benefits.List(context.TODO(), polar.UserBenefitListParams{
		Limit:          polar.F(int64(1)),
		OrderID:        polar.F[polar.UserBenefitListParamsOrderIDUnion](shared.UnionString("string")),
		OrganizationID: polar.F[polar.UserBenefitListParamsOrganizationIDUnion](shared.UnionString("string")),
		Page:           polar.F(int64(1)),
		Sorting:        polar.F([]string{"string", "string", "string"}),
		SubscriptionID: polar.F[polar.UserBenefitListParamsSubscriptionIDUnion](shared.UnionString("string")),
		Type:           polar.F[polar.UserBenefitListParamsTypeUnion](polar.UserBenefitListParamsTypeBenefitType(polar.UserBenefitListParamsTypeBenefitTypeCustom)),
	})
	if err != nil {
		var apierr *polar.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
