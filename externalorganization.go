// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package polar

import (
  "context"
  "net/http"
  "net/url"

  "github.com/polarsource/polar-go/internal/apijson"
  "github.com/polarsource/polar-go/internal/apiquery"
  "github.com/polarsource/polar-go/internal/pagination"
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
func (r *ExternalOrganizationService) List(ctx context.Context, query ExternalOrganizationListParams, opts ...option.RequestOption) (res *pagination.PolarPagination[ExternalOrganizationListResponse], err error) {
  var raw *http.Response
  opts = append(r.Options[:], opts...)
  opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
  path := "v1/external_organizations/"
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

// List external organizations.
func (r *ExternalOrganizationService) ListAutoPaging(ctx context.Context, query ExternalOrganizationListParams, opts ...option.RequestOption) (*pagination.PolarPaginationAutoPager[ExternalOrganizationListResponse]) {
  return pagination.NewPolarPaginationAutoPager(r.List(ctx, query, opts...))
}

type ExternalOrganizationListResponse struct {
ID string `json:"id,required" format:"uuid"`
AvatarURL string `json:"avatar_url,required"`
Bio string `json:"bio,required,nullable"`
Blog string `json:"blog,required,nullable"`
Company string `json:"company,required,nullable"`
Email string `json:"email,required,nullable"`
IsPersonal bool `json:"is_personal,required"`
Location string `json:"location,required,nullable"`
Name string `json:"name,required"`
// The organization ID.
OrganizationID string `json:"organization_id,required,nullable" format:"uuid4"`
Platform ExternalOrganizationListResponsePlatform `json:"platform,required"`
PrettyName string `json:"pretty_name,required,nullable"`
TwitterUsername string `json:"twitter_username,required,nullable"`
JSON externalOrganizationListResponseJSON `json:"-"`
}

// externalOrganizationListResponseJSON contains the JSON metadata for the struct
// [ExternalOrganizationListResponse]
type externalOrganizationListResponseJSON struct {
ID apijson.Field
AvatarURL apijson.Field
Bio apijson.Field
Blog apijson.Field
Company apijson.Field
Email apijson.Field
IsPersonal apijson.Field
Location apijson.Field
Name apijson.Field
OrganizationID apijson.Field
Platform apijson.Field
PrettyName apijson.Field
TwitterUsername apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ExternalOrganizationListResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r externalOrganizationListResponseJSON) RawJSON() (string) {
  return r.raw
}

type ExternalOrganizationListResponsePlatform string

const (
  ExternalOrganizationListResponsePlatformGitHub ExternalOrganizationListResponsePlatform = "github"
)

func (r ExternalOrganizationListResponsePlatform) IsKnown() (bool) {
  switch r {
  case ExternalOrganizationListResponsePlatformGitHub:
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
Sorting param.Field[[]ExternalOrganizationListParamsSorting] `query:"sorting"`
}

// URLQuery serializes [ExternalOrganizationListParams]'s query parameters as
// `url.Values`.
func (r ExternalOrganizationListParams) URLQuery() (v url.Values) {
  return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
    ArrayFormat: apiquery.ArrayQueryFormatComma,
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

func (r ExternalOrganizationListParamsNameArray) ImplementsExternalOrganizationListParamsNameUnion() {}

// Filter by organization ID.
//
// Satisfied by [shared.UnionString],
// [ExternalOrganizationListParamsOrganizationIDArray].
type ExternalOrganizationListParamsOrganizationIDUnion interface {
  ImplementsExternalOrganizationListParamsOrganizationIDUnion()
}

type ExternalOrganizationListParamsOrganizationIDArray []string

func (r ExternalOrganizationListParamsOrganizationIDArray) ImplementsExternalOrganizationListParamsOrganizationIDUnion() {}

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

func (r ExternalOrganizationListParamsPlatformPlatforms) IsKnown() (bool) {
  switch r {
  case ExternalOrganizationListParamsPlatformPlatformsGitHub:
      return true
  }
  return false
}

func (r ExternalOrganizationListParamsPlatformPlatforms) implementsExternalOrganizationListParamsPlatformUnion() {}

type ExternalOrganizationListParamsPlatformArray []ExternalOrganizationListParamsPlatformArray

func (r ExternalOrganizationListParamsPlatformArray) implementsExternalOrganizationListParamsPlatformUnion() {}

type ExternalOrganizationListParamsSorting string

const (
  ExternalOrganizationListParamsSortingCreatedAt ExternalOrganizationListParamsSorting = "created_at"
  ExternalOrganizationListParamsSorting-CreatedAt ExternalOrganizationListParamsSorting = "-created_at"
  ExternalOrganizationListParamsSortingName ExternalOrganizationListParamsSorting = "name"
  ExternalOrganizationListParamsSorting-Name ExternalOrganizationListParamsSorting = "-name"
)

func (r ExternalOrganizationListParamsSorting) IsKnown() (bool) {
  switch r {
  case ExternalOrganizationListParamsSortingCreatedAt, ExternalOrganizationListParamsSorting-CreatedAt, ExternalOrganizationListParamsSortingName, ExternalOrganizationListParamsSorting-Name:
      return true
  }
  return false
}
