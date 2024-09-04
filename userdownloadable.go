// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package polar

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/polarsource/polar-go/internal/apijson"
	"github.com/polarsource/polar-go/internal/apiquery"
	"github.com/polarsource/polar-go/internal/pagination"
	"github.com/polarsource/polar-go/internal/param"
	"github.com/polarsource/polar-go/internal/requestconfig"
	"github.com/polarsource/polar-go/option"
)

// UserDownloadableService contains methods and other services that help with
// interacting with the polar API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserDownloadableService] method instead.
type UserDownloadableService struct {
	Options []option.RequestOption
}

// NewUserDownloadableService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewUserDownloadableService(opts ...option.RequestOption) (r *UserDownloadableService) {
	r = &UserDownloadableService{}
	r.Options = opts
	return
}

// Get Downloadable
func (r *UserDownloadableService) Get(ctx context.Context, token string, opts ...option.RequestOption) (res *UserDownloadableGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if token == "" {
		err = errors.New("missing required token parameter")
		return
	}
	path := fmt.Sprintf("v1/users/downloadables/%s", token)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Downloadables
func (r *UserDownloadableService) List(ctx context.Context, query UserDownloadableListParams, opts ...option.RequestOption) (res *pagination.PolarPagination[UserDownloadableListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/users/downloadables/"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List Downloadables
func (r *UserDownloadableService) ListAutoPaging(ctx context.Context, query UserDownloadableListParams, opts ...option.RequestOption) *pagination.PolarPaginationAutoPager[UserDownloadableListResponse] {
	return pagination.NewPolarPaginationAutoPager(r.List(ctx, query, opts...))
}

type UserDownloadableGetResponse = interface{}

type UserDownloadableListResponse struct {
	ID        string                           `json:"id,required" format:"uuid4"`
	BenefitID string                           `json:"benefit_id,required" format:"uuid4"`
	File      UserDownloadableListResponseFile `json:"file,required"`
	JSON      userDownloadableListResponseJSON `json:"-"`
}

// userDownloadableListResponseJSON contains the JSON metadata for the struct
// [UserDownloadableListResponse]
type userDownloadableListResponseJSON struct {
	ID          apijson.Field
	BenefitID   apijson.Field
	File        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserDownloadableListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userDownloadableListResponseJSON) RawJSON() string {
	return r.raw
}

type UserDownloadableListResponseFile struct {
	ID                   string                                   `json:"id,required" format:"uuid4"`
	ChecksumEtag         string                                   `json:"checksum_etag,required,nullable"`
	ChecksumSha256Base64 string                                   `json:"checksum_sha256_base64,required,nullable"`
	ChecksumSha256Hex    string                                   `json:"checksum_sha256_hex,required,nullable"`
	Download             UserDownloadableListResponseFileDownload `json:"download,required"`
	IsUploaded           bool                                     `json:"is_uploaded,required"`
	LastModifiedAt       time.Time                                `json:"last_modified_at,required,nullable" format:"date-time"`
	MimeType             string                                   `json:"mime_type,required"`
	Name                 string                                   `json:"name,required"`
	OrganizationID       string                                   `json:"organization_id,required" format:"uuid4"`
	Path                 string                                   `json:"path,required"`
	Service              UserDownloadableListResponseFileService  `json:"service,required"`
	Size                 int64                                    `json:"size,required"`
	SizeReadable         string                                   `json:"size_readable,required"`
	StorageVersion       string                                   `json:"storage_version,required,nullable"`
	Version              string                                   `json:"version,required,nullable"`
	JSON                 userDownloadableListResponseFileJSON     `json:"-"`
}

// userDownloadableListResponseFileJSON contains the JSON metadata for the struct
// [UserDownloadableListResponseFile]
type userDownloadableListResponseFileJSON struct {
	ID                   apijson.Field
	ChecksumEtag         apijson.Field
	ChecksumSha256Base64 apijson.Field
	ChecksumSha256Hex    apijson.Field
	Download             apijson.Field
	IsUploaded           apijson.Field
	LastModifiedAt       apijson.Field
	MimeType             apijson.Field
	Name                 apijson.Field
	OrganizationID       apijson.Field
	Path                 apijson.Field
	Service              apijson.Field
	Size                 apijson.Field
	SizeReadable         apijson.Field
	StorageVersion       apijson.Field
	Version              apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *UserDownloadableListResponseFile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userDownloadableListResponseFileJSON) RawJSON() string {
	return r.raw
}

type UserDownloadableListResponseFileDownload struct {
	ExpiresAt time.Time                                    `json:"expires_at,required" format:"date-time"`
	URL       string                                       `json:"url,required"`
	Headers   map[string]string                            `json:"headers"`
	JSON      userDownloadableListResponseFileDownloadJSON `json:"-"`
}

// userDownloadableListResponseFileDownloadJSON contains the JSON metadata for the
// struct [UserDownloadableListResponseFileDownload]
type userDownloadableListResponseFileDownloadJSON struct {
	ExpiresAt   apijson.Field
	URL         apijson.Field
	Headers     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserDownloadableListResponseFileDownload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userDownloadableListResponseFileDownloadJSON) RawJSON() string {
	return r.raw
}

type UserDownloadableListResponseFileService string

const (
	UserDownloadableListResponseFileServiceDownloadable       UserDownloadableListResponseFileService = "downloadable"
	UserDownloadableListResponseFileServiceProductMedia       UserDownloadableListResponseFileService = "product_media"
	UserDownloadableListResponseFileServiceOrganizationAvatar UserDownloadableListResponseFileService = "organization_avatar"
)

func (r UserDownloadableListResponseFileService) IsKnown() bool {
	switch r {
	case UserDownloadableListResponseFileServiceDownloadable, UserDownloadableListResponseFileServiceProductMedia, UserDownloadableListResponseFileServiceOrganizationAvatar:
		return true
	}
	return false
}

type UserDownloadableListParams struct {
	// Filter by given benefit ID.
	BenefitID param.Field[UserDownloadableListParamsBenefitIDUnion] `query:"benefit_id" format:"uuid4"`
	// Size of a page, defaults to 10. Maximum is 100.
	Limit param.Field[int64] `query:"limit"`
	// Filter by organization ID.
	OrganizationID param.Field[UserDownloadableListParamsOrganizationIDUnion] `query:"organization_id" format:"uuid4"`
	// Page number, defaults to 1.
	Page param.Field[int64] `query:"page"`
}

// URLQuery serializes [UserDownloadableListParams]'s query parameters as
// `url.Values`.
func (r UserDownloadableListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by given benefit ID.
//
// Satisfied by [shared.UnionString], [UserDownloadableListParamsBenefitIDArray].
type UserDownloadableListParamsBenefitIDUnion interface {
	ImplementsUserDownloadableListParamsBenefitIDUnion()
}

type UserDownloadableListParamsBenefitIDArray []string

func (r UserDownloadableListParamsBenefitIDArray) ImplementsUserDownloadableListParamsBenefitIDUnion() {
}

// Filter by organization ID.
//
// Satisfied by [shared.UnionString],
// [UserDownloadableListParamsOrganizationIDArray].
type UserDownloadableListParamsOrganizationIDUnion interface {
	ImplementsUserDownloadableListParamsOrganizationIDUnion()
}

type UserDownloadableListParamsOrganizationIDArray []string

func (r UserDownloadableListParamsOrganizationIDArray) ImplementsUserDownloadableListParamsOrganizationIDUnion() {
}
