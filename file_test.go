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

func TestFileNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Files.New(context.TODO(), polar.FileNewParams{
		Body: polar.FileNewParamsBodyDownloadableFileCreate{
			MimeType: polar.F("mime_type"),
			Name:     polar.F("name"),
			Service:  polar.F(polar.FileNewParamsBodyDownloadableFileCreateServiceDownloadable),
			Size:     polar.F(int64(0)),
			Upload: polar.F(polar.FileNewParamsBodyDownloadableFileCreateUpload{
				Parts: polar.F([]polar.FileNewParamsBodyDownloadableFileCreateUploadPart{{
					ChunkEnd:             polar.F(int64(0)),
					ChunkStart:           polar.F(int64(0)),
					Number:               polar.F(int64(0)),
					ChecksumSha256Base64: polar.F("checksum_sha256_base64"),
				}, {
					ChunkEnd:             polar.F(int64(0)),
					ChunkStart:           polar.F(int64(0)),
					Number:               polar.F(int64(0)),
					ChecksumSha256Base64: polar.F("checksum_sha256_base64"),
				}, {
					ChunkEnd:             polar.F(int64(0)),
					ChunkStart:           polar.F(int64(0)),
					Number:               polar.F(int64(0)),
					ChecksumSha256Base64: polar.F("checksum_sha256_base64"),
				}}),
			}),
			ChecksumSha256Base64: polar.F("checksum_sha256_base64"),
			OrganizationID:       polar.F("organization_id"),
			Version:              polar.F("version"),
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

func TestFileUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Files.Update(
		context.TODO(),
		"id",
		polar.FileUpdateParams{
			Name:    polar.F("name"),
			Version: polar.F("version"),
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

func TestFileListWithOptionalParams(t *testing.T) {
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
	_, err := client.Files.List(context.TODO(), polar.FileListParams{
		IDs:            polar.F([]string{"string", "string", "string"}),
		Limit:          polar.F(int64(1)),
		OrganizationID: polar.F("organization_id"),
		Page:           polar.F(int64(1)),
	})
	if err != nil {
		var apierr *polar.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFileDelete(t *testing.T) {
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
	err := client.Files.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *polar.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFileUploaded(t *testing.T) {
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
	_, err := client.Files.Uploaded(context.TODO(), polar.FileUploadedParams{
		PathID: polar.F("id"),
		BodyID: polar.F("id"),
		Parts: polar.F([]polar.FileUploadedParamsPart{{
			ChecksumEtag:         polar.F("checksum_etag"),
			ChecksumSha256Base64: polar.F("checksum_sha256_base64"),
			Number:               polar.F(int64(0)),
		}, {
			ChecksumEtag:         polar.F("checksum_etag"),
			ChecksumSha256Base64: polar.F("checksum_sha256_base64"),
			Number:               polar.F(int64(0)),
		}, {
			ChecksumEtag:         polar.F("checksum_etag"),
			ChecksumSha256Base64: polar.F("checksum_sha256_base64"),
			Number:               polar.F(int64(0)),
		}}),
		Path: polar.F("path"),
	})
	if err != nil {
		var apierr *polar.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
