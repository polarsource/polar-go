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

func TestOauth2RegisterNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Oauth2.Register.New(context.TODO(), polar.Oauth2RegisterNewParams{
		ClientName:              polar.F("client_name"),
		RedirectUris:            polar.F([]string{"https://example.com", "https://example.com", "https://example.com"}),
		ClientUri:               polar.F("client_uri"),
		GrantTypes:              polar.F([]polar.Oauth2RegisterNewParamsGrantType{polar.Oauth2RegisterNewParamsGrantTypeAuthorizationCode, polar.Oauth2RegisterNewParamsGrantTypeRefreshToken}),
		LogoUri:                 polar.F("https://example.com"),
		PolicyUri:               polar.F("https://example.com"),
		ResponseTypes:           polar.F([]polar.Oauth2RegisterNewParamsResponseType{polar.Oauth2RegisterNewParamsResponseTypeCode}),
		Scope:                   polar.F("scope"),
		TokenEndpointAuthMethod: polar.F(polar.Oauth2RegisterNewParamsTokenEndpointAuthMethodClientSecretBasic),
		TosUri:                  polar.F("https://example.com"),
	})
	if err != nil {
		var apierr *polar.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOauth2RegisterGet(t *testing.T) {
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
	_, err := client.Oauth2.Register.Get(context.TODO(), "client_id")
	if err != nil {
		var apierr *polar.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOauth2RegisterUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Oauth2.Register.Update(context.TODO(), polar.Oauth2RegisterUpdateParams{
		PathClientID:            polar.F("client_id"),
		BodyClientID:            polar.F("client_id"),
		ClientName:              polar.F("client_name"),
		RedirectUris:            polar.F([]string{"https://example.com", "https://example.com", "https://example.com"}),
		ClientUri:               polar.F("client_uri"),
		GrantTypes:              polar.F([]polar.Oauth2RegisterUpdateParamsGrantType{polar.Oauth2RegisterUpdateParamsGrantTypeAuthorizationCode, polar.Oauth2RegisterUpdateParamsGrantTypeRefreshToken}),
		LogoUri:                 polar.F("https://example.com"),
		PolicyUri:               polar.F("https://example.com"),
		ResponseTypes:           polar.F([]polar.Oauth2RegisterUpdateParamsResponseType{polar.Oauth2RegisterUpdateParamsResponseTypeCode}),
		Scope:                   polar.F("scope"),
		TokenEndpointAuthMethod: polar.F(polar.Oauth2RegisterUpdateParamsTokenEndpointAuthMethodClientSecretBasic),
		TosUri:                  polar.F("https://example.com"),
	})
	if err != nil {
		var apierr *polar.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOauth2RegisterDelete(t *testing.T) {
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
	_, err := client.Oauth2.Register.Delete(context.TODO(), "client_id")
	if err != nil {
		var apierr *polar.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
