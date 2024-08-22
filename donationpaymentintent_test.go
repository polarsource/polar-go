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

func TestDonationPaymentIntentNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Donations.PaymentIntent.New(context.TODO(), polar.DonationPaymentIntentNewParams{
		Amount:                   polar.F(int64(1)),
		Email:                    polar.F("dev@stainlessapi.com"),
		ToOrganizationID:         polar.F("to_organization_id"),
		Currency:                 polar.F("currency"),
		IssueID:                  polar.F("issue_id"),
		Message:                  polar.F("message"),
		OnBehalfOfOrganizationID: polar.F("on_behalf_of_organization_id"),
		SetupFutureUsage:         polar.F(polar.DonationPaymentIntentNewParamsSetupFutureUsageOnSession),
	})
	if err != nil {
		var apierr *polar.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDonationPaymentIntentUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Donations.PaymentIntent.Update(
		context.TODO(),
		"id",
		polar.DonationPaymentIntentUpdateParams{
			Amount:                   polar.F(int64(1)),
			Email:                    polar.F("email"),
			Currency:                 polar.F("currency"),
			IssueID:                  polar.F("issue_id"),
			Message:                  polar.F("message"),
			OnBehalfOfOrganizationID: polar.F("on_behalf_of_organization_id"),
			SetupFutureUsage:         polar.F(polar.DonationPaymentIntentUpdateParamsSetupFutureUsageOnSession),
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
