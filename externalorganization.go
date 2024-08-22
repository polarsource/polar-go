// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package polar

import (
	"context"
	"net/http"
	"net/url"

	"github.com/polarsource/polar-go/internal/apijson"
	"github.com/polarsource/polar-go/internal/apiquery"
	"github.com/polarsource/polar-go/internal/param"
	"github.com/polarsource/polar-go/internal/requestconfig"
	"github.com/polarsource/polar-go/option"
)

// ExternalOrganizationService contains methods and other services that help with
// interacting with the polar API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExternalOrganizationService] method instead.
type ExternalOrganizationService struct {
	Options []option.RequestOption
}

// NewExternalOrganizationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewExternalOrganizationService(opts ...option.RequestOption) (r *ExternalOrganizationService) {
	r = &ExternalOrganizationService{}
	r.Options = opts
	return
}

// List external organizations.
func (r *ExternalOrganizationService) List(ctx context.Context, query ExternalOrganizationListParams, opts ...option.RequestOption) (res *ExternalOrganizationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/external_organizations/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ExternalOrganizationListResponse struct {
	Pagination ExternalOrganizationListResponsePagination `json:"pagination,required"`
	Items      []ExternalOrganizationListResponseItem     `json:"items"`
	JSON       externalOrganizationListResponseJSON       `json:"-"`
}

// externalOrganizationListResponseJSON contains the JSON metadata for the struct
// [ExternalOrganizationListResponse]
type externalOrganizationListResponseJSON struct {
	Pagination  apijson.Field
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExternalOrganizationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r externalOrganizationListResponseJSON) RawJSON() string {
	return r.raw
}

type ExternalOrganizationListResponsePagination struct {
	MaxPage    int64                                          `json:"max_page,required"`
	TotalCount int64                                          `json:"total_count,required"`
	JSON       externalOrganizationListResponsePaginationJSON `json:"-"`
}

// externalOrganizationListResponsePaginationJSON contains the JSON metadata for
// the struct [ExternalOrganizationListResponsePagination]
type externalOrganizationListResponsePaginationJSON struct {
	MaxPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExternalOrganizationListResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r externalOrganizationListResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type ExternalOrganizationListResponseItem struct {
	ID         string                                        `json:"id,required" format:"uuid"`
	AvatarURL  string                                        `json:"avatar_url,required"`
	IsPersonal bool                                          `json:"is_personal,required"`
	Name       string                                        `json:"name,required"`
	Platform   ExternalOrganizationListResponseItemsPlatform `json:"platform,required"`
	Bio        string                                        `json:"bio,nullable"`
	Blog       string                                        `json:"blog,nullable"`
	Company    string                                        `json:"company,nullable"`
	Email      string                                        `json:"email,nullable"`
	Location   string                                        `json:"location,nullable"`
	// The organization ID.
	OrganizationID  string                                   `json:"organization_id,nullable" format:"uuid4"`
	PrettyName      string                                   `json:"pretty_name,nullable"`
	TwitterUsername string                                   `json:"twitter_username,nullable"`
	JSON            externalOrganizationListResponseItemJSON `json:"-"`
}

// externalOrganizationListResponseItemJSON contains the JSON metadata for the
// struct [ExternalOrganizationListResponseItem]
type externalOrganizationListResponseItemJSON struct {
	ID              apijson.Field
	AvatarURL       apijson.Field
	IsPersonal      apijson.Field
	Name            apijson.Field
	Platform        apijson.Field
	Bio             apijson.Field
	Blog            apijson.Field
	Company         apijson.Field
	Email           apijson.Field
	Location        apijson.Field
	OrganizationID  apijson.Field
	PrettyName      apijson.Field
	TwitterUsername apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ExternalOrganizationListResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r externalOrganizationListResponseItemJSON) RawJSON() string {
	return r.raw
}

type ExternalOrganizationListResponseItemsPlatform string

const (
	ExternalOrganizationListResponseItemsPlatformGitHub ExternalOrganizationListResponseItemsPlatform = "github"
)

func (r ExternalOrganizationListResponseItemsPlatform) IsKnown() bool {
	switch r {
	case ExternalOrganizationListResponseItemsPlatformGitHub:
		return true
	}
	return false
}

type ExternalOrganizationListParams struct {
	// Size of a page, defaults to 10. Maximum is 100.
	Limit param.Field[int64] `query:"limit"`
	// Filter by name.
	Name param.Field[ExternalOrganizationListParamsNameUnion] `query:"name"`
	// Filter by organization ID.
	OrganizationID param.Field[ExternalOrganizationListParamsOrganizationIDUnion] `query:"organization_id" format:"uuid4"`
	// Page number, defaults to 1.
	Page param.Field[int64] `query:"page"`
	// Filter by platform.
	Platform param.Field[ExternalOrganizationListParamsPlatformUnion] `query:"platform"`
	// Sorting criterion. Several criteria can be used simultaneously and will be
	// applied in order. Add a minus sign `-` before the criteria name to sort by
	// descending order.
	Sorting param.Field[[]string] `query:"sorting"`
}

// URLQuery serializes [ExternalOrganizationListParams]'s query parameters as
// `url.Values`.
func (r ExternalOrganizationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by name.
//
// Satisfied by [shared.UnionString], [ExternalOrganizationListParamsNameArray].
type ExternalOrganizationListParamsNameUnion interface {
	ImplementsExternalOrganizationListParamsNameUnion()
}

type ExternalOrganizationListParamsNameArray []string

func (r ExternalOrganizationListParamsNameArray) ImplementsExternalOrganizationListParamsNameUnion() {
}

// Filter by organization ID.
//
// Satisfied by [shared.UnionString],
// [ExternalOrganizationListParamsOrganizationIDArray].
type ExternalOrganizationListParamsOrganizationIDUnion interface {
	ImplementsExternalOrganizationListParamsOrganizationIDUnion()
}

type ExternalOrganizationListParamsOrganizationIDArray []string

func (r ExternalOrganizationListParamsOrganizationIDArray) ImplementsExternalOrganizationListParamsOrganizationIDUnion() {
}

// Filter by platform.
//
// Satisfied by [ExternalOrganizationListParamsPlatformPlatforms],
// [ExternalOrganizationListParamsPlatformArray].
type ExternalOrganizationListParamsPlatformUnion interface {
	implementsExternalOrganizationListParamsPlatformUnion()
}

type ExternalOrganizationListParamsPlatformPlatforms string

const (
	ExternalOrganizationListParamsPlatformPlatformsGitHub ExternalOrganizationListParamsPlatformPlatforms = "github"
)

func (r ExternalOrganizationListParamsPlatformPlatforms) IsKnown() bool {
	switch r {
	case ExternalOrganizationListParamsPlatformPlatformsGitHub:
		return true
	}
	return false
}

func (r ExternalOrganizationListParamsPlatformPlatforms) implementsExternalOrganizationListParamsPlatformUnion() {
}

type ExternalOrganizationListParamsPlatformArray []ExternalOrganizationListParamsPlatformArray

func (r ExternalOrganizationListParamsPlatformArray) implementsExternalOrganizationListParamsPlatformUnion() {
}
