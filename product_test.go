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

func TestProductNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Products.New(context.TODO(), polar.ProductNewParams{
		Body: polar.ProductNewParamsBodyProductRecurringCreate{
			Name: polar.F("xxx"),
			Prices: polar.F([]polar.ProductNewParamsBodyProductRecurringCreatePrice{{
				PriceAmount:       polar.F(int64(1)),
				RecurringInterval: polar.F(polar.ProductNewParamsBodyProductRecurringCreatePricesRecurringIntervalMonth),
				Type:              polar.F(polar.ProductNewParamsBodyProductRecurringCreatePricesTypeRecurring),
				PriceCurrency:     polar.F("price_currency"),
			}}),
			Type:           polar.F(polar.ProductNewParamsBodyProductRecurringCreateTypeIndividual),
			Description:    polar.F("description"),
			IsHighlighted:  polar.F(true),
			Medias:         polar.F([]string{"string", "string", "string"}),
			OrganizationID: polar.F("organization_id"),
		},
	})
	if err != nil {
		var apierr *polar.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProductGet(t *testing.T) {
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
	_, err := client.Products.Get(context.TODO(), "id")
	if err != nil {
		var apierr *polar.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProductUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Products.Update(
		context.TODO(),
		"id",
		polar.ProductUpdateParams{
			Description:   polar.F("description"),
			IsArchived:    polar.F(true),
			IsHighlighted: polar.F(true),
			Medias:        polar.F([]string{"string", "string", "string"}),
			Name:          polar.F("xxx"),
			Prices: polar.F([]polar.ProductUpdateParamsPriceUnion{polar.ProductUpdateParamsPricesExistingProductPrice{
				ID: polar.F("id"),
			}, polar.ProductUpdateParamsPricesExistingProductPrice{
				ID: polar.F("id"),
			}, polar.ProductUpdateParamsPricesExistingProductPrice{
				ID: polar.F("id"),
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

func TestProductListWithOptionalParams(t *testing.T) {
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
	_, err := client.Products.List(context.TODO(), polar.ProductListParams{
		BenefitID:      polar.F[polar.ProductListParamsBenefitIDUnion](shared.UnionString("string")),
		IsArchived:     polar.F(true),
		IsRecurring:    polar.F(true),
		Limit:          polar.F(int64(1)),
		OrganizationID: polar.F[polar.ProductListParamsOrganizationIDUnion](shared.UnionString("string")),
		Page:           polar.F(int64(1)),
		Type:           polar.F[polar.ProductListParamsTypeUnion](polar.ProductListParamsTypeSubscriptionTierType(polar.ProductListParamsTypeSubscriptionTierTypeFree)),
	})
	if err != nil {
		var apierr *polar.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
