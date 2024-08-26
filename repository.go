// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package polar

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/polarsource/polar-go/internal/apijson"
	"github.com/polarsource/polar-go/internal/apiquery"
	"github.com/polarsource/polar-go/internal/pagination"
	"github.com/polarsource/polar-go/internal/param"
	"github.com/polarsource/polar-go/internal/requestconfig"
	"github.com/polarsource/polar-go/option"
)

// RepositoryService contains methods and other services that help with interacting
// with the polar API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRepositoryService] method instead.
type RepositoryService struct {
	Options []option.RequestOption
}

// NewRepositoryService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRepositoryService(opts ...option.RequestOption) (r *RepositoryService) {
	r = &RepositoryService{}
	r.Options = opts
	return
}

// Get a repository by ID.
func (r *RepositoryService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *RepositoryGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/repositories/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a repository.
func (r *RepositoryService) Update(ctx context.Context, id string, body RepositoryUpdateParams, opts ...option.RequestOption) (res *RepositoryUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/repositories/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List repositories.
func (r *RepositoryService) List(ctx context.Context, query RepositoryListParams, opts ...option.RequestOption) (res *pagination.PolarPagination[RepositoryListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/repositories/"
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

// List repositories.
func (r *RepositoryService) ListAutoPaging(ctx context.Context, query RepositoryListParams, opts ...option.RequestOption) *pagination.PolarPaginationAutoPager[RepositoryListResponse] {
	return pagination.NewPolarPaginationAutoPager(r.List(ctx, query, opts...))
}

type RepositoryGetResponse struct {
	ID           string                            `json:"id,required" format:"uuid"`
	Description  string                            `json:"description,required,nullable"`
	Homepage     string                            `json:"homepage,required,nullable"`
	IsPrivate    bool                              `json:"is_private,required"`
	License      string                            `json:"license,required,nullable"`
	Name         string                            `json:"name,required"`
	Organization RepositoryGetResponseOrganization `json:"organization,required"`
	Platform     RepositoryGetResponsePlatform     `json:"platform,required"`
	// Settings for the repository profile
	ProfileSettings RepositoryGetResponseProfileSettings `json:"profile_settings,required,nullable"`
	Stars           int64                                `json:"stars,required,nullable"`
	JSON            repositoryGetResponseJSON            `json:"-"`
}

// repositoryGetResponseJSON contains the JSON metadata for the struct
// [RepositoryGetResponse]
type repositoryGetResponseJSON struct {
	ID              apijson.Field
	Description     apijson.Field
	Homepage        apijson.Field
	IsPrivate       apijson.Field
	License         apijson.Field
	Name            apijson.Field
	Organization    apijson.Field
	Platform        apijson.Field
	ProfileSettings apijson.Field
	Stars           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RepositoryGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r repositoryGetResponseJSON) RawJSON() string {
	return r.raw
}

type RepositoryGetResponseOrganization struct {
	ID         string `json:"id,required" format:"uuid"`
	AvatarURL  string `json:"avatar_url,required"`
	Bio        string `json:"bio,required,nullable"`
	Blog       string `json:"blog,required,nullable"`
	Company    string `json:"company,required,nullable"`
	Email      string `json:"email,required,nullable"`
	IsPersonal bool   `json:"is_personal,required"`
	Location   string `json:"location,required,nullable"`
	Name       string `json:"name,required"`
	// The organization ID.
	OrganizationID  string                                    `json:"organization_id,required,nullable" format:"uuid4"`
	Platform        RepositoryGetResponseOrganizationPlatform `json:"platform,required"`
	PrettyName      string                                    `json:"pretty_name,required,nullable"`
	TwitterUsername string                                    `json:"twitter_username,required,nullable"`
	JSON            repositoryGetResponseOrganizationJSON     `json:"-"`
}

// repositoryGetResponseOrganizationJSON contains the JSON metadata for the struct
// [RepositoryGetResponseOrganization]
type repositoryGetResponseOrganizationJSON struct {
	ID              apijson.Field
	AvatarURL       apijson.Field
	Bio             apijson.Field
	Blog            apijson.Field
	Company         apijson.Field
	Email           apijson.Field
	IsPersonal      apijson.Field
	Location        apijson.Field
	Name            apijson.Field
	OrganizationID  apijson.Field
	Platform        apijson.Field
	PrettyName      apijson.Field
	TwitterUsername apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RepositoryGetResponseOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r repositoryGetResponseOrganizationJSON) RawJSON() string {
	return r.raw
}

type RepositoryGetResponseOrganizationPlatform string

const (
	RepositoryGetResponseOrganizationPlatformGitHub RepositoryGetResponseOrganizationPlatform = "github"
)

func (r RepositoryGetResponseOrganizationPlatform) IsKnown() bool {
	switch r {
	case RepositoryGetResponseOrganizationPlatformGitHub:
		return true
	}
	return false
}

type RepositoryGetResponsePlatform string

const (
	RepositoryGetResponsePlatformGitHub RepositoryGetResponsePlatform = "github"
)

func (r RepositoryGetResponsePlatform) IsKnown() bool {
	switch r {
	case RepositoryGetResponsePlatformGitHub:
		return true
	}
	return false
}

// Settings for the repository profile
type RepositoryGetResponseProfileSettings struct {
	// A URL to a cover image
	CoverImageURL string `json:"cover_image_url,nullable"`
	// A description of the repository
	Description string `json:"description,nullable"`
	// A list of featured organizations
	FeaturedOrganizations []string `json:"featured_organizations,nullable" format:"uuid4"`
	// A list of highlighted subscription tiers
	HighlightedSubscriptionTiers []string `json:"highlighted_subscription_tiers,nullable" format:"uuid4"`
	// A list of links related to the repository
	Links []string                                 `json:"links,nullable" format:"uri"`
	JSON  repositoryGetResponseProfileSettingsJSON `json:"-"`
}

// repositoryGetResponseProfileSettingsJSON contains the JSON metadata for the
// struct [RepositoryGetResponseProfileSettings]
type repositoryGetResponseProfileSettingsJSON struct {
	CoverImageURL                apijson.Field
	Description                  apijson.Field
	FeaturedOrganizations        apijson.Field
	HighlightedSubscriptionTiers apijson.Field
	Links                        apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *RepositoryGetResponseProfileSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r repositoryGetResponseProfileSettingsJSON) RawJSON() string {
	return r.raw
}

type RepositoryUpdateResponse struct {
	ID           string                               `json:"id,required" format:"uuid"`
	Description  string                               `json:"description,required,nullable"`
	Homepage     string                               `json:"homepage,required,nullable"`
	IsPrivate    bool                                 `json:"is_private,required"`
	License      string                               `json:"license,required,nullable"`
	Name         string                               `json:"name,required"`
	Organization RepositoryUpdateResponseOrganization `json:"organization,required"`
	Platform     RepositoryUpdateResponsePlatform     `json:"platform,required"`
	// Settings for the repository profile
	ProfileSettings RepositoryUpdateResponseProfileSettings `json:"profile_settings,required,nullable"`
	Stars           int64                                   `json:"stars,required,nullable"`
	JSON            repositoryUpdateResponseJSON            `json:"-"`
}

// repositoryUpdateResponseJSON contains the JSON metadata for the struct
// [RepositoryUpdateResponse]
type repositoryUpdateResponseJSON struct {
	ID              apijson.Field
	Description     apijson.Field
	Homepage        apijson.Field
	IsPrivate       apijson.Field
	License         apijson.Field
	Name            apijson.Field
	Organization    apijson.Field
	Platform        apijson.Field
	ProfileSettings apijson.Field
	Stars           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RepositoryUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r repositoryUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type RepositoryUpdateResponseOrganization struct {
	ID         string `json:"id,required" format:"uuid"`
	AvatarURL  string `json:"avatar_url,required"`
	Bio        string `json:"bio,required,nullable"`
	Blog       string `json:"blog,required,nullable"`
	Company    string `json:"company,required,nullable"`
	Email      string `json:"email,required,nullable"`
	IsPersonal bool   `json:"is_personal,required"`
	Location   string `json:"location,required,nullable"`
	Name       string `json:"name,required"`
	// The organization ID.
	OrganizationID  string                                       `json:"organization_id,required,nullable" format:"uuid4"`
	Platform        RepositoryUpdateResponseOrganizationPlatform `json:"platform,required"`
	PrettyName      string                                       `json:"pretty_name,required,nullable"`
	TwitterUsername string                                       `json:"twitter_username,required,nullable"`
	JSON            repositoryUpdateResponseOrganizationJSON     `json:"-"`
}

// repositoryUpdateResponseOrganizationJSON contains the JSON metadata for the
// struct [RepositoryUpdateResponseOrganization]
type repositoryUpdateResponseOrganizationJSON struct {
	ID              apijson.Field
	AvatarURL       apijson.Field
	Bio             apijson.Field
	Blog            apijson.Field
	Company         apijson.Field
	Email           apijson.Field
	IsPersonal      apijson.Field
	Location        apijson.Field
	Name            apijson.Field
	OrganizationID  apijson.Field
	Platform        apijson.Field
	PrettyName      apijson.Field
	TwitterUsername apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RepositoryUpdateResponseOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r repositoryUpdateResponseOrganizationJSON) RawJSON() string {
	return r.raw
}

type RepositoryUpdateResponseOrganizationPlatform string

const (
	RepositoryUpdateResponseOrganizationPlatformGitHub RepositoryUpdateResponseOrganizationPlatform = "github"
)

func (r RepositoryUpdateResponseOrganizationPlatform) IsKnown() bool {
	switch r {
	case RepositoryUpdateResponseOrganizationPlatformGitHub:
		return true
	}
	return false
}

type RepositoryUpdateResponsePlatform string

const (
	RepositoryUpdateResponsePlatformGitHub RepositoryUpdateResponsePlatform = "github"
)

func (r RepositoryUpdateResponsePlatform) IsKnown() bool {
	switch r {
	case RepositoryUpdateResponsePlatformGitHub:
		return true
	}
	return false
}

// Settings for the repository profile
type RepositoryUpdateResponseProfileSettings struct {
	// A URL to a cover image
	CoverImageURL string `json:"cover_image_url,nullable"`
	// A description of the repository
	Description string `json:"description,nullable"`
	// A list of featured organizations
	FeaturedOrganizations []string `json:"featured_organizations,nullable" format:"uuid4"`
	// A list of highlighted subscription tiers
	HighlightedSubscriptionTiers []string `json:"highlighted_subscription_tiers,nullable" format:"uuid4"`
	// A list of links related to the repository
	Links []string                                    `json:"links,nullable" format:"uri"`
	JSON  repositoryUpdateResponseProfileSettingsJSON `json:"-"`
}

// repositoryUpdateResponseProfileSettingsJSON contains the JSON metadata for the
// struct [RepositoryUpdateResponseProfileSettings]
type repositoryUpdateResponseProfileSettingsJSON struct {
	CoverImageURL                apijson.Field
	Description                  apijson.Field
	FeaturedOrganizations        apijson.Field
	HighlightedSubscriptionTiers apijson.Field
	Links                        apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *RepositoryUpdateResponseProfileSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r repositoryUpdateResponseProfileSettingsJSON) RawJSON() string {
	return r.raw
}

type RepositoryListResponse struct {
	ID           string                             `json:"id,required" format:"uuid"`
	Description  string                             `json:"description,required,nullable"`
	Homepage     string                             `json:"homepage,required,nullable"`
	IsPrivate    bool                               `json:"is_private,required"`
	License      string                             `json:"license,required,nullable"`
	Name         string                             `json:"name,required"`
	Organization RepositoryListResponseOrganization `json:"organization,required"`
	Platform     RepositoryListResponsePlatform     `json:"platform,required"`
	// Settings for the repository profile
	ProfileSettings RepositoryListResponseProfileSettings `json:"profile_settings,required,nullable"`
	Stars           int64                                 `json:"stars,required,nullable"`
	JSON            repositoryListResponseJSON            `json:"-"`
}

// repositoryListResponseJSON contains the JSON metadata for the struct
// [RepositoryListResponse]
type repositoryListResponseJSON struct {
	ID              apijson.Field
	Description     apijson.Field
	Homepage        apijson.Field
	IsPrivate       apijson.Field
	License         apijson.Field
	Name            apijson.Field
	Organization    apijson.Field
	Platform        apijson.Field
	ProfileSettings apijson.Field
	Stars           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RepositoryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r repositoryListResponseJSON) RawJSON() string {
	return r.raw
}

type RepositoryListResponseOrganization struct {
	ID         string `json:"id,required" format:"uuid"`
	AvatarURL  string `json:"avatar_url,required"`
	Bio        string `json:"bio,required,nullable"`
	Blog       string `json:"blog,required,nullable"`
	Company    string `json:"company,required,nullable"`
	Email      string `json:"email,required,nullable"`
	IsPersonal bool   `json:"is_personal,required"`
	Location   string `json:"location,required,nullable"`
	Name       string `json:"name,required"`
	// The organization ID.
	OrganizationID  string                                     `json:"organization_id,required,nullable" format:"uuid4"`
	Platform        RepositoryListResponseOrganizationPlatform `json:"platform,required"`
	PrettyName      string                                     `json:"pretty_name,required,nullable"`
	TwitterUsername string                                     `json:"twitter_username,required,nullable"`
	JSON            repositoryListResponseOrganizationJSON     `json:"-"`
}

// repositoryListResponseOrganizationJSON contains the JSON metadata for the struct
// [RepositoryListResponseOrganization]
type repositoryListResponseOrganizationJSON struct {
	ID              apijson.Field
	AvatarURL       apijson.Field
	Bio             apijson.Field
	Blog            apijson.Field
	Company         apijson.Field
	Email           apijson.Field
	IsPersonal      apijson.Field
	Location        apijson.Field
	Name            apijson.Field
	OrganizationID  apijson.Field
	Platform        apijson.Field
	PrettyName      apijson.Field
	TwitterUsername apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RepositoryListResponseOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r repositoryListResponseOrganizationJSON) RawJSON() string {
	return r.raw
}

type RepositoryListResponseOrganizationPlatform string

const (
	RepositoryListResponseOrganizationPlatformGitHub RepositoryListResponseOrganizationPlatform = "github"
)

func (r RepositoryListResponseOrganizationPlatform) IsKnown() bool {
	switch r {
	case RepositoryListResponseOrganizationPlatformGitHub:
		return true
	}
	return false
}

type RepositoryListResponsePlatform string

const (
	RepositoryListResponsePlatformGitHub RepositoryListResponsePlatform = "github"
)

func (r RepositoryListResponsePlatform) IsKnown() bool {
	switch r {
	case RepositoryListResponsePlatformGitHub:
		return true
	}
	return false
}

// Settings for the repository profile
type RepositoryListResponseProfileSettings struct {
	// A URL to a cover image
	CoverImageURL string `json:"cover_image_url,nullable"`
	// A description of the repository
	Description string `json:"description,nullable"`
	// A list of featured organizations
	FeaturedOrganizations []string `json:"featured_organizations,nullable" format:"uuid4"`
	// A list of highlighted subscription tiers
	HighlightedSubscriptionTiers []string `json:"highlighted_subscription_tiers,nullable" format:"uuid4"`
	// A list of links related to the repository
	Links []string                                  `json:"links,nullable" format:"uri"`
	JSON  repositoryListResponseProfileSettingsJSON `json:"-"`
}

// repositoryListResponseProfileSettingsJSON contains the JSON metadata for the
// struct [RepositoryListResponseProfileSettings]
type repositoryListResponseProfileSettingsJSON struct {
	CoverImageURL                apijson.Field
	Description                  apijson.Field
	FeaturedOrganizations        apijson.Field
	HighlightedSubscriptionTiers apijson.Field
	Links                        apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *RepositoryListResponseProfileSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r repositoryListResponseProfileSettingsJSON) RawJSON() string {
	return r.raw
}

type RepositoryUpdateParams struct {
	ProfileSettings param.Field[RepositoryUpdateParamsProfileSettings] `json:"profile_settings"`
}

func (r RepositoryUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RepositoryUpdateParamsProfileSettings struct {
	CoverImageURL                param.Field[string]   `json:"cover_image_url"`
	Description                  param.Field[string]   `json:"description"`
	FeaturedOrganizations        param.Field[[]string] `json:"featured_organizations" format:"uuid4"`
	HighlightedSubscriptionTiers param.Field[[]string] `json:"highlighted_subscription_tiers" format:"uuid4"`
	Links                        param.Field[[]string] `json:"links" format:"uri"`
	SetCoverImageURL             param.Field[bool]     `json:"set_cover_image_url"`
	SetDescription               param.Field[bool]     `json:"set_description"`
}

func (r RepositoryUpdateParamsProfileSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RepositoryListParams struct {
	// Filter by external organization name.
	ExternalOrganizationName param.Field[RepositoryListParamsExternalOrganizationNameUnion] `query:"external_organization_name"`
	// Filter by private status.
	IsPrivate param.Field[bool] `query:"is_private"`
	// Size of a page, defaults to 10. Maximum is 100.
	Limit param.Field[int64] `query:"limit"`
	// Filter by name.
	Name param.Field[RepositoryListParamsNameUnion] `query:"name"`
	// Filter by organization ID.
	OrganizationID param.Field[RepositoryListParamsOrganizationIDUnion] `query:"organization_id" format:"uuid4"`
	// Page number, defaults to 1.
	Page param.Field[int64] `query:"page"`
	// Filter by platform.
	Platform param.Field[RepositoryListParamsPlatformUnion] `query:"platform"`
	// Sorting criterion. Several criteria can be used simultaneously and will be
	// applied in order. Add a minus sign `-` before the criteria name to sort by
	// descending order.
	Sorting param.Field[[]string] `query:"sorting"`
}

// URLQuery serializes [RepositoryListParams]'s query parameters as `url.Values`.
func (r RepositoryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by external organization name.
//
// Satisfied by [shared.UnionString],
// [RepositoryListParamsExternalOrganizationNameArray].
type RepositoryListParamsExternalOrganizationNameUnion interface {
	ImplementsRepositoryListParamsExternalOrganizationNameUnion()
}

type RepositoryListParamsExternalOrganizationNameArray []string

func (r RepositoryListParamsExternalOrganizationNameArray) ImplementsRepositoryListParamsExternalOrganizationNameUnion() {
}

// Filter by name.
//
// Satisfied by [shared.UnionString], [RepositoryListParamsNameArray].
type RepositoryListParamsNameUnion interface {
	ImplementsRepositoryListParamsNameUnion()
}

type RepositoryListParamsNameArray []string

func (r RepositoryListParamsNameArray) ImplementsRepositoryListParamsNameUnion() {}

// Filter by organization ID.
//
// Satisfied by [shared.UnionString], [RepositoryListParamsOrganizationIDArray].
type RepositoryListParamsOrganizationIDUnion interface {
	ImplementsRepositoryListParamsOrganizationIDUnion()
}

type RepositoryListParamsOrganizationIDArray []string

func (r RepositoryListParamsOrganizationIDArray) ImplementsRepositoryListParamsOrganizationIDUnion() {
}

// Filter by platform.
//
// Satisfied by [RepositoryListParamsPlatformPlatforms],
// [RepositoryListParamsPlatformArray].
type RepositoryListParamsPlatformUnion interface {
	implementsRepositoryListParamsPlatformUnion()
}

type RepositoryListParamsPlatformPlatforms string

const (
	RepositoryListParamsPlatformPlatformsGitHub RepositoryListParamsPlatformPlatforms = "github"
)

func (r RepositoryListParamsPlatformPlatforms) IsKnown() bool {
	switch r {
	case RepositoryListParamsPlatformPlatformsGitHub:
		return true
	}
	return false
}

func (r RepositoryListParamsPlatformPlatforms) implementsRepositoryListParamsPlatformUnion() {}

type RepositoryListParamsPlatformArray []RepositoryListParamsPlatformArray

func (r RepositoryListParamsPlatformArray) implementsRepositoryListParamsPlatformUnion() {}
