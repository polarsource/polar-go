// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package polar_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/polar-go"
	"github.com/stainless-sdks/polar-go/internal/testutil"
	"github.com/stainless-sdks/polar-go/option"
	"github.com/stainless-sdks/polar-go/shared"
)

func TestArticleNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Articles.New(context.TODO(), polar.ArticleNewParams{
		Title:                     polar.F("x"),
		Body:                      polar.F("body"),
		BodyBase64:                polar.F("body_base64"),
		Byline:                    polar.F(polar.ArticleNewParamsBylineUser),
		IsPinned:                  polar.F(true),
		NotifySubscribers:         polar.F(true),
		OgDescription:             polar.F("og_description"),
		OgImageURL:                polar.F("https://example.com"),
		OrganizationID:            polar.F("organization_id"),
		PaidSubscribersOnly:       polar.F(true),
		PaidSubscribersOnlyEndsAt: polar.F(time.Now()),
		PublishedAt:               polar.F(time.Now()),
		Slug:                      polar.F("slug"),
		Visibility:                polar.F(polar.ArticleNewParamsVisibilityPublic),
	})
	if err != nil {
		var apierr *polar.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestArticleGet(t *testing.T) {
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
	_, err := client.Articles.Get(context.TODO(), "id")
	if err != nil {
		var apierr *polar.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestArticleUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Articles.Update(
		context.TODO(),
		"id",
		polar.ArticleUpdateParams{
			Body:                      polar.F("body"),
			BodyBase64:                polar.F("body_base64"),
			Byline:                    polar.F(polar.ArticleUpdateParamsBylineUser),
			IsPinned:                  polar.F(true),
			NotifySubscribers:         polar.F(true),
			OgDescription:             polar.F("og_description"),
			OgImageURL:                polar.F("https://example.com"),
			PaidSubscribersOnly:       polar.F(true),
			PaidSubscribersOnlyEndsAt: polar.F(time.Now()),
			PublishedAt:               polar.F(time.Now()),
			Slug:                      polar.F("slug"),
			Title:                     polar.F("title"),
			Visibility:                polar.F(polar.ArticleUpdateParamsVisibilityPublic),
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

func TestArticleListWithOptionalParams(t *testing.T) {
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
	_, err := client.Articles.List(context.TODO(), polar.ArticleListParams{
		IsPinned:       polar.F(true),
		IsPublished:    polar.F(true),
		IsSubscribed:   polar.F(true),
		Limit:          polar.F(int64(1)),
		OrganizationID: polar.F[polar.ArticleListParamsOrganizationIDUnion](shared.UnionString("string")),
		Page:           polar.F(int64(1)),
		Slug:           polar.F("slug"),
		Visibility:     polar.F[polar.ArticleListParamsVisibilityUnion](polar.ArticleListParamsVisibilityArticleVisibility(polar.ArticleListParamsVisibilityArticleVisibilityPublic)),
	})
	if err != nil {
		var apierr *polar.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestArticleDelete(t *testing.T) {
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
	err := client.Articles.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *polar.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestArticlePreview(t *testing.T) {
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
	_, err := client.Articles.Preview(
		context.TODO(),
		"id",
		polar.ArticlePreviewParams{
			Email: polar.F("dev@stainlessapi.com"),
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

func TestArticleSend(t *testing.T) {
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
	_, err := client.Articles.Send(context.TODO(), "id")
	if err != nil {
		var apierr *polar.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
