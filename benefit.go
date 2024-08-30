// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package polar

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/polarsource/polar-go/internal/apijson"
	"github.com/polarsource/polar-go/internal/apiquery"
	"github.com/polarsource/polar-go/internal/pagination"
	"github.com/polarsource/polar-go/internal/param"
	"github.com/polarsource/polar-go/internal/requestconfig"
	"github.com/polarsource/polar-go/option"
	"github.com/tidwall/gjson"
)

// BenefitService contains methods and other services that help with interacting
// with the polar API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBenefitService] method instead.
type BenefitService struct {
	Options []option.RequestOption
	Grants  *BenefitGrantService
}

// NewBenefitService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBenefitService(opts ...option.RequestOption) (r *BenefitService) {
	r = &BenefitService{}
	r.Options = opts
	r.Grants = NewBenefitGrantService(opts...)
	return
}

// Create a benefit.
func (r *BenefitService) New(ctx context.Context, body BenefitNewParams, opts ...option.RequestOption) (res *BenefitNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/benefits/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a benefit by ID.
func (r *BenefitService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *BenefitGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/benefits/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a benefit.
func (r *BenefitService) Update(ctx context.Context, id string, body BenefitUpdateParams, opts ...option.RequestOption) (res *BenefitUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/benefits/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List benefits.
func (r *BenefitService) List(ctx context.Context, query BenefitListParams, opts ...option.RequestOption) (res *pagination.PolarPagination[BenefitListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/benefits/"
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

// List benefits.
func (r *BenefitService) ListAutoPaging(ctx context.Context, query BenefitListParams, opts ...option.RequestOption) *pagination.PolarPaginationAutoPager[BenefitListResponse] {
	return pagination.NewPolarPaginationAutoPager(r.List(ctx, query, opts...))
}

// Delete a benefit.
//
// > [!WARNING] Every grants associated with the benefit will be revoked. Users
// > will lose access to the benefit.
func (r *BenefitService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/benefits/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// A benefit of type `articles`.
//
// Use it to grant access to posts.
type BenefitNewResponse struct {
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The ID of the benefit.
	ID   string                 `json:"id,required" format:"uuid4"`
	Type BenefitNewResponseType `json:"type,required"`
	// The description of the benefit.
	Description string `json:"description,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool `json:"selectable,required"`
	// Whether the benefit is deletable.
	Deletable bool `json:"deletable,required"`
	// The ID of the organization owning the benefit.
	OrganizationID string `json:"organization_id,required" format:"uuid4"`
	// This field can have the runtime type of
	// [BenefitNewResponseBenefitArticlesProperties],
	// [BenefitNewResponseBenefitAdsProperties],
	// [BenefitNewResponseBenefitCustomProperties],
	// [BenefitNewResponseBenefitDiscordOutputProperties],
	// [BenefitNewResponseBenefitGitHubRepositoryProperties],
	// [BenefitNewResponseBenefitDownloadablesProperties],
	// [BenefitNewResponseBenefitLicenseKeysOutputProperties].
	Properties interface{} `json:"properties"`
	// Whether the benefit is taxable.
	IsTaxApplicable bool                   `json:"is_tax_applicable"`
	JSON            benefitNewResponseJSON `json:"-"`
	union           BenefitNewResponseUnion
}

// benefitNewResponseJSON contains the JSON metadata for the struct
// [BenefitNewResponse]
type benefitNewResponseJSON struct {
	CreatedAt       apijson.Field
	ModifiedAt      apijson.Field
	ID              apijson.Field
	Type            apijson.Field
	Description     apijson.Field
	Selectable      apijson.Field
	Deletable       apijson.Field
	OrganizationID  apijson.Field
	Properties      apijson.Field
	IsTaxApplicable apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r benefitNewResponseJSON) RawJSON() string {
	return r.raw
}

func (r *BenefitNewResponse) UnmarshalJSON(data []byte) (err error) {
	*r = BenefitNewResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [BenefitNewResponseUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [BenefitNewResponseBenefitArticles],
// [BenefitNewResponseBenefitAds], [BenefitNewResponseBenefitCustom],
// [BenefitNewResponseBenefitDiscordOutput],
// [BenefitNewResponseBenefitGitHubRepository],
// [BenefitNewResponseBenefitDownloadables],
// [BenefitNewResponseBenefitLicenseKeysOutput].
func (r BenefitNewResponse) AsUnion() BenefitNewResponseUnion {
	return r.union
}

// A benefit of type `articles`.
//
// Use it to grant access to posts.
//
// Union satisfied by [BenefitNewResponseBenefitArticles],
// [BenefitNewResponseBenefitAds], [BenefitNewResponseBenefitCustom],
// [BenefitNewResponseBenefitDiscordOutput],
// [BenefitNewResponseBenefitGitHubRepository],
// [BenefitNewResponseBenefitDownloadables] or
// [BenefitNewResponseBenefitLicenseKeysOutput].
type BenefitNewResponseUnion interface {
	implementsBenefitNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BenefitNewResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BenefitNewResponseBenefitArticles{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BenefitNewResponseBenefitAds{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BenefitNewResponseBenefitCustom{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BenefitNewResponseBenefitDiscordOutput{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BenefitNewResponseBenefitGitHubRepository{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BenefitNewResponseBenefitDownloadables{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BenefitNewResponseBenefitLicenseKeysOutput{}),
		},
	)
}

// A benefit of type `articles`.
//
// Use it to grant access to posts.
type BenefitNewResponseBenefitArticles struct {
	// The ID of the benefit.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the benefit is deletable.
	Deletable bool `json:"deletable,required"`
	// The description of the benefit.
	Description string `json:"description,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The ID of the organization owning the benefit.
	OrganizationID string `json:"organization_id,required" format:"uuid4"`
	// Properties for a benefit of type `articles`.
	Properties BenefitNewResponseBenefitArticlesProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                  `json:"selectable,required"`
	Type       BenefitNewResponseBenefitArticlesType `json:"type,required"`
	JSON       benefitNewResponseBenefitArticlesJSON `json:"-"`
}

// benefitNewResponseBenefitArticlesJSON contains the JSON metadata for the struct
// [BenefitNewResponseBenefitArticles]
type benefitNewResponseBenefitArticlesJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Deletable      apijson.Field
	Description    apijson.Field
	ModifiedAt     apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	Selectable     apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BenefitNewResponseBenefitArticles) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitNewResponseBenefitArticlesJSON) RawJSON() string {
	return r.raw
}

func (r BenefitNewResponseBenefitArticles) implementsBenefitNewResponse() {}

// Properties for a benefit of type `articles`.
type BenefitNewResponseBenefitArticlesProperties struct {
	// Whether the user can access paid articles.
	PaidArticles bool                                            `json:"paid_articles,required"`
	JSON         benefitNewResponseBenefitArticlesPropertiesJSON `json:"-"`
}

// benefitNewResponseBenefitArticlesPropertiesJSON contains the JSON metadata for
// the struct [BenefitNewResponseBenefitArticlesProperties]
type benefitNewResponseBenefitArticlesPropertiesJSON struct {
	PaidArticles apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *BenefitNewResponseBenefitArticlesProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitNewResponseBenefitArticlesPropertiesJSON) RawJSON() string {
	return r.raw
}

type BenefitNewResponseBenefitArticlesType string

const (
	BenefitNewResponseBenefitArticlesTypeArticles BenefitNewResponseBenefitArticlesType = "articles"
)

func (r BenefitNewResponseBenefitArticlesType) IsKnown() bool {
	switch r {
	case BenefitNewResponseBenefitArticlesTypeArticles:
		return true
	}
	return false
}

// A benefit of type `ads`.
//
// Use it so your backers can display ads on your README, website, etc.
type BenefitNewResponseBenefitAds struct {
	// The ID of the benefit.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the benefit is deletable.
	Deletable bool `json:"deletable,required"`
	// The description of the benefit.
	Description string `json:"description,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The ID of the organization owning the benefit.
	OrganizationID string `json:"organization_id,required" format:"uuid4"`
	// Properties for a benefit of type `ads`.
	Properties BenefitNewResponseBenefitAdsProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                             `json:"selectable,required"`
	Type       BenefitNewResponseBenefitAdsType `json:"type,required"`
	JSON       benefitNewResponseBenefitAdsJSON `json:"-"`
}

// benefitNewResponseBenefitAdsJSON contains the JSON metadata for the struct
// [BenefitNewResponseBenefitAds]
type benefitNewResponseBenefitAdsJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Deletable      apijson.Field
	Description    apijson.Field
	ModifiedAt     apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	Selectable     apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BenefitNewResponseBenefitAds) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitNewResponseBenefitAdsJSON) RawJSON() string {
	return r.raw
}

func (r BenefitNewResponseBenefitAds) implementsBenefitNewResponse() {}

// Properties for a benefit of type `ads`.
type BenefitNewResponseBenefitAdsProperties struct {
	// The height of the displayed ad.
	ImageHeight int64 `json:"image_height"`
	// The width of the displayed ad.
	ImageWidth int64                                      `json:"image_width"`
	JSON       benefitNewResponseBenefitAdsPropertiesJSON `json:"-"`
}

// benefitNewResponseBenefitAdsPropertiesJSON contains the JSON metadata for the
// struct [BenefitNewResponseBenefitAdsProperties]
type benefitNewResponseBenefitAdsPropertiesJSON struct {
	ImageHeight apijson.Field
	ImageWidth  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BenefitNewResponseBenefitAdsProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitNewResponseBenefitAdsPropertiesJSON) RawJSON() string {
	return r.raw
}

type BenefitNewResponseBenefitAdsType string

const (
	BenefitNewResponseBenefitAdsTypeAds BenefitNewResponseBenefitAdsType = "ads"
)

func (r BenefitNewResponseBenefitAdsType) IsKnown() bool {
	switch r {
	case BenefitNewResponseBenefitAdsTypeAds:
		return true
	}
	return false
}

// A benefit of type `custom`.
//
// Use it to grant any kind of benefit that doesn't fit in the other types.
type BenefitNewResponseBenefitCustom struct {
	// The ID of the benefit.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the benefit is deletable.
	Deletable bool `json:"deletable,required"`
	// The description of the benefit.
	Description string `json:"description,required"`
	// Whether the benefit is taxable.
	IsTaxApplicable bool `json:"is_tax_applicable,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The ID of the organization owning the benefit.
	OrganizationID string `json:"organization_id,required" format:"uuid4"`
	// Properties for a benefit of type `custom`.
	Properties BenefitNewResponseBenefitCustomProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                `json:"selectable,required"`
	Type       BenefitNewResponseBenefitCustomType `json:"type,required"`
	JSON       benefitNewResponseBenefitCustomJSON `json:"-"`
}

// benefitNewResponseBenefitCustomJSON contains the JSON metadata for the struct
// [BenefitNewResponseBenefitCustom]
type benefitNewResponseBenefitCustomJSON struct {
	ID              apijson.Field
	CreatedAt       apijson.Field
	Deletable       apijson.Field
	Description     apijson.Field
	IsTaxApplicable apijson.Field
	ModifiedAt      apijson.Field
	OrganizationID  apijson.Field
	Properties      apijson.Field
	Selectable      apijson.Field
	Type            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *BenefitNewResponseBenefitCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitNewResponseBenefitCustomJSON) RawJSON() string {
	return r.raw
}

func (r BenefitNewResponseBenefitCustom) implementsBenefitNewResponse() {}

// Properties for a benefit of type `custom`.
type BenefitNewResponseBenefitCustomProperties struct {
	// Private note to be shared with users who have this benefit granted.
	Note string                                        `json:"note,required,nullable"`
	JSON benefitNewResponseBenefitCustomPropertiesJSON `json:"-"`
}

// benefitNewResponseBenefitCustomPropertiesJSON contains the JSON metadata for the
// struct [BenefitNewResponseBenefitCustomProperties]
type benefitNewResponseBenefitCustomPropertiesJSON struct {
	Note        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BenefitNewResponseBenefitCustomProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitNewResponseBenefitCustomPropertiesJSON) RawJSON() string {
	return r.raw
}

type BenefitNewResponseBenefitCustomType string

const (
	BenefitNewResponseBenefitCustomTypeCustom BenefitNewResponseBenefitCustomType = "custom"
)

func (r BenefitNewResponseBenefitCustomType) IsKnown() bool {
	switch r {
	case BenefitNewResponseBenefitCustomTypeCustom:
		return true
	}
	return false
}

// A benefit of type `discord`.
//
// Use it to automatically invite your backers to a Discord server.
type BenefitNewResponseBenefitDiscordOutput struct {
	// The ID of the benefit.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the benefit is deletable.
	Deletable bool `json:"deletable,required"`
	// The description of the benefit.
	Description string `json:"description,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The ID of the organization owning the benefit.
	OrganizationID string `json:"organization_id,required" format:"uuid4"`
	// Properties for a benefit of type `discord`.
	Properties BenefitNewResponseBenefitDiscordOutputProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                       `json:"selectable,required"`
	Type       BenefitNewResponseBenefitDiscordOutputType `json:"type,required"`
	JSON       benefitNewResponseBenefitDiscordOutputJSON `json:"-"`
}

// benefitNewResponseBenefitDiscordOutputJSON contains the JSON metadata for the
// struct [BenefitNewResponseBenefitDiscordOutput]
type benefitNewResponseBenefitDiscordOutputJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Deletable      apijson.Field
	Description    apijson.Field
	ModifiedAt     apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	Selectable     apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BenefitNewResponseBenefitDiscordOutput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitNewResponseBenefitDiscordOutputJSON) RawJSON() string {
	return r.raw
}

func (r BenefitNewResponseBenefitDiscordOutput) implementsBenefitNewResponse() {}

// Properties for a benefit of type `discord`.
type BenefitNewResponseBenefitDiscordOutputProperties struct {
	// The ID of the Discord server.
	GuildID    string `json:"guild_id,required"`
	GuildToken string `json:"guild_token,required"`
	// The ID of the Discord role to grant.
	RoleID string                                               `json:"role_id,required"`
	JSON   benefitNewResponseBenefitDiscordOutputPropertiesJSON `json:"-"`
}

// benefitNewResponseBenefitDiscordOutputPropertiesJSON contains the JSON metadata
// for the struct [BenefitNewResponseBenefitDiscordOutputProperties]
type benefitNewResponseBenefitDiscordOutputPropertiesJSON struct {
	GuildID     apijson.Field
	GuildToken  apijson.Field
	RoleID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BenefitNewResponseBenefitDiscordOutputProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitNewResponseBenefitDiscordOutputPropertiesJSON) RawJSON() string {
	return r.raw
}

type BenefitNewResponseBenefitDiscordOutputType string

const (
	BenefitNewResponseBenefitDiscordOutputTypeDiscord BenefitNewResponseBenefitDiscordOutputType = "discord"
)

func (r BenefitNewResponseBenefitDiscordOutputType) IsKnown() bool {
	switch r {
	case BenefitNewResponseBenefitDiscordOutputTypeDiscord:
		return true
	}
	return false
}

// A benefit of type `github_repository`.
//
// Use it to automatically invite your backers to a private GitHub repository.
type BenefitNewResponseBenefitGitHubRepository struct {
	// The ID of the benefit.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the benefit is deletable.
	Deletable bool `json:"deletable,required"`
	// The description of the benefit.
	Description string `json:"description,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The ID of the organization owning the benefit.
	OrganizationID string `json:"organization_id,required" format:"uuid4"`
	// Properties for a benefit of type `github_repository`.
	Properties BenefitNewResponseBenefitGitHubRepositoryProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                          `json:"selectable,required"`
	Type       BenefitNewResponseBenefitGitHubRepositoryType `json:"type,required"`
	JSON       benefitNewResponseBenefitGitHubRepositoryJSON `json:"-"`
}

// benefitNewResponseBenefitGitHubRepositoryJSON contains the JSON metadata for the
// struct [BenefitNewResponseBenefitGitHubRepository]
type benefitNewResponseBenefitGitHubRepositoryJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Deletable      apijson.Field
	Description    apijson.Field
	ModifiedAt     apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	Selectable     apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BenefitNewResponseBenefitGitHubRepository) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitNewResponseBenefitGitHubRepositoryJSON) RawJSON() string {
	return r.raw
}

func (r BenefitNewResponseBenefitGitHubRepository) implementsBenefitNewResponse() {}

// Properties for a benefit of type `github_repository`.
type BenefitNewResponseBenefitGitHubRepositoryProperties struct {
	// The permission level to grant. Read more about roles and their permissions on
	// [GitHub documentation](https://docs.github.com/en/organizations/managing-user-access-to-your-organizations-repositories/managing-repository-roles/repository-roles-for-an-organization#permissions-for-each-role).
	Permission   BenefitNewResponseBenefitGitHubRepositoryPropertiesPermission `json:"permission,required"`
	RepositoryID string                                                        `json:"repository_id,required,nullable" format:"uuid4"`
	// The name of the repository.
	RepositoryName string `json:"repository_name,required"`
	// The owner of the repository.
	RepositoryOwner string                                                  `json:"repository_owner,required"`
	JSON            benefitNewResponseBenefitGitHubRepositoryPropertiesJSON `json:"-"`
}

// benefitNewResponseBenefitGitHubRepositoryPropertiesJSON contains the JSON
// metadata for the struct [BenefitNewResponseBenefitGitHubRepositoryProperties]
type benefitNewResponseBenefitGitHubRepositoryPropertiesJSON struct {
	Permission      apijson.Field
	RepositoryID    apijson.Field
	RepositoryName  apijson.Field
	RepositoryOwner apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *BenefitNewResponseBenefitGitHubRepositoryProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitNewResponseBenefitGitHubRepositoryPropertiesJSON) RawJSON() string {
	return r.raw
}

// The permission level to grant. Read more about roles and their permissions on
// [GitHub documentation](https://docs.github.com/en/organizations/managing-user-access-to-your-organizations-repositories/managing-repository-roles/repository-roles-for-an-organization#permissions-for-each-role).
type BenefitNewResponseBenefitGitHubRepositoryPropertiesPermission string

const (
	BenefitNewResponseBenefitGitHubRepositoryPropertiesPermissionPull     BenefitNewResponseBenefitGitHubRepositoryPropertiesPermission = "pull"
	BenefitNewResponseBenefitGitHubRepositoryPropertiesPermissionTriage   BenefitNewResponseBenefitGitHubRepositoryPropertiesPermission = "triage"
	BenefitNewResponseBenefitGitHubRepositoryPropertiesPermissionPush     BenefitNewResponseBenefitGitHubRepositoryPropertiesPermission = "push"
	BenefitNewResponseBenefitGitHubRepositoryPropertiesPermissionMaintain BenefitNewResponseBenefitGitHubRepositoryPropertiesPermission = "maintain"
	BenefitNewResponseBenefitGitHubRepositoryPropertiesPermissionAdmin    BenefitNewResponseBenefitGitHubRepositoryPropertiesPermission = "admin"
)

func (r BenefitNewResponseBenefitGitHubRepositoryPropertiesPermission) IsKnown() bool {
	switch r {
	case BenefitNewResponseBenefitGitHubRepositoryPropertiesPermissionPull, BenefitNewResponseBenefitGitHubRepositoryPropertiesPermissionTriage, BenefitNewResponseBenefitGitHubRepositoryPropertiesPermissionPush, BenefitNewResponseBenefitGitHubRepositoryPropertiesPermissionMaintain, BenefitNewResponseBenefitGitHubRepositoryPropertiesPermissionAdmin:
		return true
	}
	return false
}

type BenefitNewResponseBenefitGitHubRepositoryType string

const (
	BenefitNewResponseBenefitGitHubRepositoryTypeGitHubRepository BenefitNewResponseBenefitGitHubRepositoryType = "github_repository"
)

func (r BenefitNewResponseBenefitGitHubRepositoryType) IsKnown() bool {
	switch r {
	case BenefitNewResponseBenefitGitHubRepositoryTypeGitHubRepository:
		return true
	}
	return false
}

type BenefitNewResponseBenefitDownloadables struct {
	// The ID of the benefit.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the benefit is deletable.
	Deletable bool `json:"deletable,required"`
	// The description of the benefit.
	Description string `json:"description,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The ID of the organization owning the benefit.
	OrganizationID string                                           `json:"organization_id,required" format:"uuid4"`
	Properties     BenefitNewResponseBenefitDownloadablesProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                       `json:"selectable,required"`
	Type       BenefitNewResponseBenefitDownloadablesType `json:"type,required"`
	JSON       benefitNewResponseBenefitDownloadablesJSON `json:"-"`
}

// benefitNewResponseBenefitDownloadablesJSON contains the JSON metadata for the
// struct [BenefitNewResponseBenefitDownloadables]
type benefitNewResponseBenefitDownloadablesJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Deletable      apijson.Field
	Description    apijson.Field
	ModifiedAt     apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	Selectable     apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BenefitNewResponseBenefitDownloadables) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitNewResponseBenefitDownloadablesJSON) RawJSON() string {
	return r.raw
}

func (r BenefitNewResponseBenefitDownloadables) implementsBenefitNewResponse() {}

type BenefitNewResponseBenefitDownloadablesProperties struct {
	Archived map[string]bool                                      `json:"archived,required"`
	Files    []string                                             `json:"files,required" format:"uuid4"`
	JSON     benefitNewResponseBenefitDownloadablesPropertiesJSON `json:"-"`
}

// benefitNewResponseBenefitDownloadablesPropertiesJSON contains the JSON metadata
// for the struct [BenefitNewResponseBenefitDownloadablesProperties]
type benefitNewResponseBenefitDownloadablesPropertiesJSON struct {
	Archived    apijson.Field
	Files       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BenefitNewResponseBenefitDownloadablesProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitNewResponseBenefitDownloadablesPropertiesJSON) RawJSON() string {
	return r.raw
}

type BenefitNewResponseBenefitDownloadablesType string

const (
	BenefitNewResponseBenefitDownloadablesTypeDownloadables BenefitNewResponseBenefitDownloadablesType = "downloadables"
)

func (r BenefitNewResponseBenefitDownloadablesType) IsKnown() bool {
	switch r {
	case BenefitNewResponseBenefitDownloadablesTypeDownloadables:
		return true
	}
	return false
}

type BenefitNewResponseBenefitLicenseKeysOutput struct {
	// The ID of the benefit.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the benefit is deletable.
	Deletable bool `json:"deletable,required"`
	// The description of the benefit.
	Description string `json:"description,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The ID of the organization owning the benefit.
	OrganizationID string                                               `json:"organization_id,required" format:"uuid4"`
	Properties     BenefitNewResponseBenefitLicenseKeysOutputProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                           `json:"selectable,required"`
	Type       BenefitNewResponseBenefitLicenseKeysOutputType `json:"type,required"`
	JSON       benefitNewResponseBenefitLicenseKeysOutputJSON `json:"-"`
}

// benefitNewResponseBenefitLicenseKeysOutputJSON contains the JSON metadata for
// the struct [BenefitNewResponseBenefitLicenseKeysOutput]
type benefitNewResponseBenefitLicenseKeysOutputJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Deletable      apijson.Field
	Description    apijson.Field
	ModifiedAt     apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	Selectable     apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BenefitNewResponseBenefitLicenseKeysOutput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitNewResponseBenefitLicenseKeysOutputJSON) RawJSON() string {
	return r.raw
}

func (r BenefitNewResponseBenefitLicenseKeysOutput) implementsBenefitNewResponse() {}

type BenefitNewResponseBenefitLicenseKeysOutputProperties struct {
	Activations BenefitNewResponseBenefitLicenseKeysOutputPropertiesActivations `json:"activations,required,nullable"`
	Expires     BenefitNewResponseBenefitLicenseKeysOutputPropertiesExpires     `json:"expires,required,nullable"`
	LimitUsage  int64                                                           `json:"limit_usage,required,nullable"`
	Prefix      string                                                          `json:"prefix,required,nullable"`
	JSON        benefitNewResponseBenefitLicenseKeysOutputPropertiesJSON        `json:"-"`
}

// benefitNewResponseBenefitLicenseKeysOutputPropertiesJSON contains the JSON
// metadata for the struct [BenefitNewResponseBenefitLicenseKeysOutputProperties]
type benefitNewResponseBenefitLicenseKeysOutputPropertiesJSON struct {
	Activations apijson.Field
	Expires     apijson.Field
	LimitUsage  apijson.Field
	Prefix      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BenefitNewResponseBenefitLicenseKeysOutputProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitNewResponseBenefitLicenseKeysOutputPropertiesJSON) RawJSON() string {
	return r.raw
}

type BenefitNewResponseBenefitLicenseKeysOutputPropertiesActivations struct {
	EnableUserAdmin bool                                                                `json:"enable_user_admin,required"`
	Limit           int64                                                               `json:"limit,required"`
	JSON            benefitNewResponseBenefitLicenseKeysOutputPropertiesActivationsJSON `json:"-"`
}

// benefitNewResponseBenefitLicenseKeysOutputPropertiesActivationsJSON contains the
// JSON metadata for the struct
// [BenefitNewResponseBenefitLicenseKeysOutputPropertiesActivations]
type benefitNewResponseBenefitLicenseKeysOutputPropertiesActivationsJSON struct {
	EnableUserAdmin apijson.Field
	Limit           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *BenefitNewResponseBenefitLicenseKeysOutputPropertiesActivations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitNewResponseBenefitLicenseKeysOutputPropertiesActivationsJSON) RawJSON() string {
	return r.raw
}

type BenefitNewResponseBenefitLicenseKeysOutputPropertiesExpires struct {
	Timeframe BenefitNewResponseBenefitLicenseKeysOutputPropertiesExpiresTimeframe `json:"timeframe,required"`
	Ttl       int64                                                                `json:"ttl,required"`
	JSON      benefitNewResponseBenefitLicenseKeysOutputPropertiesExpiresJSON      `json:"-"`
}

// benefitNewResponseBenefitLicenseKeysOutputPropertiesExpiresJSON contains the
// JSON metadata for the struct
// [BenefitNewResponseBenefitLicenseKeysOutputPropertiesExpires]
type benefitNewResponseBenefitLicenseKeysOutputPropertiesExpiresJSON struct {
	Timeframe   apijson.Field
	Ttl         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BenefitNewResponseBenefitLicenseKeysOutputPropertiesExpires) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitNewResponseBenefitLicenseKeysOutputPropertiesExpiresJSON) RawJSON() string {
	return r.raw
}

type BenefitNewResponseBenefitLicenseKeysOutputPropertiesExpiresTimeframe string

const (
	BenefitNewResponseBenefitLicenseKeysOutputPropertiesExpiresTimeframeYear  BenefitNewResponseBenefitLicenseKeysOutputPropertiesExpiresTimeframe = "year"
	BenefitNewResponseBenefitLicenseKeysOutputPropertiesExpiresTimeframeMonth BenefitNewResponseBenefitLicenseKeysOutputPropertiesExpiresTimeframe = "month"
	BenefitNewResponseBenefitLicenseKeysOutputPropertiesExpiresTimeframeDay   BenefitNewResponseBenefitLicenseKeysOutputPropertiesExpiresTimeframe = "day"
)

func (r BenefitNewResponseBenefitLicenseKeysOutputPropertiesExpiresTimeframe) IsKnown() bool {
	switch r {
	case BenefitNewResponseBenefitLicenseKeysOutputPropertiesExpiresTimeframeYear, BenefitNewResponseBenefitLicenseKeysOutputPropertiesExpiresTimeframeMonth, BenefitNewResponseBenefitLicenseKeysOutputPropertiesExpiresTimeframeDay:
		return true
	}
	return false
}

type BenefitNewResponseBenefitLicenseKeysOutputType string

const (
	BenefitNewResponseBenefitLicenseKeysOutputTypeLicenseKeys BenefitNewResponseBenefitLicenseKeysOutputType = "license_keys"
)

func (r BenefitNewResponseBenefitLicenseKeysOutputType) IsKnown() bool {
	switch r {
	case BenefitNewResponseBenefitLicenseKeysOutputTypeLicenseKeys:
		return true
	}
	return false
}

type BenefitNewResponseType string

const (
	BenefitNewResponseTypeArticles         BenefitNewResponseType = "articles"
	BenefitNewResponseTypeAds              BenefitNewResponseType = "ads"
	BenefitNewResponseTypeCustom           BenefitNewResponseType = "custom"
	BenefitNewResponseTypeDiscord          BenefitNewResponseType = "discord"
	BenefitNewResponseTypeGitHubRepository BenefitNewResponseType = "github_repository"
	BenefitNewResponseTypeDownloadables    BenefitNewResponseType = "downloadables"
	BenefitNewResponseTypeLicenseKeys      BenefitNewResponseType = "license_keys"
)

func (r BenefitNewResponseType) IsKnown() bool {
	switch r {
	case BenefitNewResponseTypeArticles, BenefitNewResponseTypeAds, BenefitNewResponseTypeCustom, BenefitNewResponseTypeDiscord, BenefitNewResponseTypeGitHubRepository, BenefitNewResponseTypeDownloadables, BenefitNewResponseTypeLicenseKeys:
		return true
	}
	return false
}

// A benefit of type `articles`.
//
// Use it to grant access to posts.
type BenefitGetResponse struct {
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The ID of the benefit.
	ID   string                 `json:"id,required" format:"uuid4"`
	Type BenefitGetResponseType `json:"type,required"`
	// The description of the benefit.
	Description string `json:"description,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool `json:"selectable,required"`
	// Whether the benefit is deletable.
	Deletable bool `json:"deletable,required"`
	// The ID of the organization owning the benefit.
	OrganizationID string `json:"organization_id,required" format:"uuid4"`
	// This field can have the runtime type of
	// [BenefitGetResponseBenefitArticlesProperties],
	// [BenefitGetResponseBenefitAdsProperties],
	// [BenefitGetResponseBenefitCustomProperties],
	// [BenefitGetResponseBenefitDiscordOutputProperties],
	// [BenefitGetResponseBenefitGitHubRepositoryProperties],
	// [BenefitGetResponseBenefitDownloadablesProperties],
	// [BenefitGetResponseBenefitLicenseKeysOutputProperties].
	Properties interface{} `json:"properties"`
	// Whether the benefit is taxable.
	IsTaxApplicable bool                   `json:"is_tax_applicable"`
	JSON            benefitGetResponseJSON `json:"-"`
	union           BenefitGetResponseUnion
}

// benefitGetResponseJSON contains the JSON metadata for the struct
// [BenefitGetResponse]
type benefitGetResponseJSON struct {
	CreatedAt       apijson.Field
	ModifiedAt      apijson.Field
	ID              apijson.Field
	Type            apijson.Field
	Description     apijson.Field
	Selectable      apijson.Field
	Deletable       apijson.Field
	OrganizationID  apijson.Field
	Properties      apijson.Field
	IsTaxApplicable apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r benefitGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *BenefitGetResponse) UnmarshalJSON(data []byte) (err error) {
	*r = BenefitGetResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [BenefitGetResponseUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [BenefitGetResponseBenefitArticles],
// [BenefitGetResponseBenefitAds], [BenefitGetResponseBenefitCustom],
// [BenefitGetResponseBenefitDiscordOutput],
// [BenefitGetResponseBenefitGitHubRepository],
// [BenefitGetResponseBenefitDownloadables],
// [BenefitGetResponseBenefitLicenseKeysOutput].
func (r BenefitGetResponse) AsUnion() BenefitGetResponseUnion {
	return r.union
}

// A benefit of type `articles`.
//
// Use it to grant access to posts.
//
// Union satisfied by [BenefitGetResponseBenefitArticles],
// [BenefitGetResponseBenefitAds], [BenefitGetResponseBenefitCustom],
// [BenefitGetResponseBenefitDiscordOutput],
// [BenefitGetResponseBenefitGitHubRepository],
// [BenefitGetResponseBenefitDownloadables] or
// [BenefitGetResponseBenefitLicenseKeysOutput].
type BenefitGetResponseUnion interface {
	implementsBenefitGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BenefitGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BenefitGetResponseBenefitArticles{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BenefitGetResponseBenefitAds{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BenefitGetResponseBenefitCustom{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BenefitGetResponseBenefitDiscordOutput{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BenefitGetResponseBenefitGitHubRepository{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BenefitGetResponseBenefitDownloadables{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BenefitGetResponseBenefitLicenseKeysOutput{}),
		},
	)
}

// A benefit of type `articles`.
//
// Use it to grant access to posts.
type BenefitGetResponseBenefitArticles struct {
	// The ID of the benefit.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the benefit is deletable.
	Deletable bool `json:"deletable,required"`
	// The description of the benefit.
	Description string `json:"description,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The ID of the organization owning the benefit.
	OrganizationID string `json:"organization_id,required" format:"uuid4"`
	// Properties for a benefit of type `articles`.
	Properties BenefitGetResponseBenefitArticlesProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                  `json:"selectable,required"`
	Type       BenefitGetResponseBenefitArticlesType `json:"type,required"`
	JSON       benefitGetResponseBenefitArticlesJSON `json:"-"`
}

// benefitGetResponseBenefitArticlesJSON contains the JSON metadata for the struct
// [BenefitGetResponseBenefitArticles]
type benefitGetResponseBenefitArticlesJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Deletable      apijson.Field
	Description    apijson.Field
	ModifiedAt     apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	Selectable     apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BenefitGetResponseBenefitArticles) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitGetResponseBenefitArticlesJSON) RawJSON() string {
	return r.raw
}

func (r BenefitGetResponseBenefitArticles) implementsBenefitGetResponse() {}

// Properties for a benefit of type `articles`.
type BenefitGetResponseBenefitArticlesProperties struct {
	// Whether the user can access paid articles.
	PaidArticles bool                                            `json:"paid_articles,required"`
	JSON         benefitGetResponseBenefitArticlesPropertiesJSON `json:"-"`
}

// benefitGetResponseBenefitArticlesPropertiesJSON contains the JSON metadata for
// the struct [BenefitGetResponseBenefitArticlesProperties]
type benefitGetResponseBenefitArticlesPropertiesJSON struct {
	PaidArticles apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *BenefitGetResponseBenefitArticlesProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitGetResponseBenefitArticlesPropertiesJSON) RawJSON() string {
	return r.raw
}

type BenefitGetResponseBenefitArticlesType string

const (
	BenefitGetResponseBenefitArticlesTypeArticles BenefitGetResponseBenefitArticlesType = "articles"
)

func (r BenefitGetResponseBenefitArticlesType) IsKnown() bool {
	switch r {
	case BenefitGetResponseBenefitArticlesTypeArticles:
		return true
	}
	return false
}

// A benefit of type `ads`.
//
// Use it so your backers can display ads on your README, website, etc.
type BenefitGetResponseBenefitAds struct {
	// The ID of the benefit.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the benefit is deletable.
	Deletable bool `json:"deletable,required"`
	// The description of the benefit.
	Description string `json:"description,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The ID of the organization owning the benefit.
	OrganizationID string `json:"organization_id,required" format:"uuid4"`
	// Properties for a benefit of type `ads`.
	Properties BenefitGetResponseBenefitAdsProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                             `json:"selectable,required"`
	Type       BenefitGetResponseBenefitAdsType `json:"type,required"`
	JSON       benefitGetResponseBenefitAdsJSON `json:"-"`
}

// benefitGetResponseBenefitAdsJSON contains the JSON metadata for the struct
// [BenefitGetResponseBenefitAds]
type benefitGetResponseBenefitAdsJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Deletable      apijson.Field
	Description    apijson.Field
	ModifiedAt     apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	Selectable     apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BenefitGetResponseBenefitAds) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitGetResponseBenefitAdsJSON) RawJSON() string {
	return r.raw
}

func (r BenefitGetResponseBenefitAds) implementsBenefitGetResponse() {}

// Properties for a benefit of type `ads`.
type BenefitGetResponseBenefitAdsProperties struct {
	// The height of the displayed ad.
	ImageHeight int64 `json:"image_height"`
	// The width of the displayed ad.
	ImageWidth int64                                      `json:"image_width"`
	JSON       benefitGetResponseBenefitAdsPropertiesJSON `json:"-"`
}

// benefitGetResponseBenefitAdsPropertiesJSON contains the JSON metadata for the
// struct [BenefitGetResponseBenefitAdsProperties]
type benefitGetResponseBenefitAdsPropertiesJSON struct {
	ImageHeight apijson.Field
	ImageWidth  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BenefitGetResponseBenefitAdsProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitGetResponseBenefitAdsPropertiesJSON) RawJSON() string {
	return r.raw
}

type BenefitGetResponseBenefitAdsType string

const (
	BenefitGetResponseBenefitAdsTypeAds BenefitGetResponseBenefitAdsType = "ads"
)

func (r BenefitGetResponseBenefitAdsType) IsKnown() bool {
	switch r {
	case BenefitGetResponseBenefitAdsTypeAds:
		return true
	}
	return false
}

// A benefit of type `custom`.
//
// Use it to grant any kind of benefit that doesn't fit in the other types.
type BenefitGetResponseBenefitCustom struct {
	// The ID of the benefit.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the benefit is deletable.
	Deletable bool `json:"deletable,required"`
	// The description of the benefit.
	Description string `json:"description,required"`
	// Whether the benefit is taxable.
	IsTaxApplicable bool `json:"is_tax_applicable,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The ID of the organization owning the benefit.
	OrganizationID string `json:"organization_id,required" format:"uuid4"`
	// Properties for a benefit of type `custom`.
	Properties BenefitGetResponseBenefitCustomProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                `json:"selectable,required"`
	Type       BenefitGetResponseBenefitCustomType `json:"type,required"`
	JSON       benefitGetResponseBenefitCustomJSON `json:"-"`
}

// benefitGetResponseBenefitCustomJSON contains the JSON metadata for the struct
// [BenefitGetResponseBenefitCustom]
type benefitGetResponseBenefitCustomJSON struct {
	ID              apijson.Field
	CreatedAt       apijson.Field
	Deletable       apijson.Field
	Description     apijson.Field
	IsTaxApplicable apijson.Field
	ModifiedAt      apijson.Field
	OrganizationID  apijson.Field
	Properties      apijson.Field
	Selectable      apijson.Field
	Type            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *BenefitGetResponseBenefitCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitGetResponseBenefitCustomJSON) RawJSON() string {
	return r.raw
}

func (r BenefitGetResponseBenefitCustom) implementsBenefitGetResponse() {}

// Properties for a benefit of type `custom`.
type BenefitGetResponseBenefitCustomProperties struct {
	// Private note to be shared with users who have this benefit granted.
	Note string                                        `json:"note,required,nullable"`
	JSON benefitGetResponseBenefitCustomPropertiesJSON `json:"-"`
}

// benefitGetResponseBenefitCustomPropertiesJSON contains the JSON metadata for the
// struct [BenefitGetResponseBenefitCustomProperties]
type benefitGetResponseBenefitCustomPropertiesJSON struct {
	Note        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BenefitGetResponseBenefitCustomProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitGetResponseBenefitCustomPropertiesJSON) RawJSON() string {
	return r.raw
}

type BenefitGetResponseBenefitCustomType string

const (
	BenefitGetResponseBenefitCustomTypeCustom BenefitGetResponseBenefitCustomType = "custom"
)

func (r BenefitGetResponseBenefitCustomType) IsKnown() bool {
	switch r {
	case BenefitGetResponseBenefitCustomTypeCustom:
		return true
	}
	return false
}

// A benefit of type `discord`.
//
// Use it to automatically invite your backers to a Discord server.
type BenefitGetResponseBenefitDiscordOutput struct {
	// The ID of the benefit.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the benefit is deletable.
	Deletable bool `json:"deletable,required"`
	// The description of the benefit.
	Description string `json:"description,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The ID of the organization owning the benefit.
	OrganizationID string `json:"organization_id,required" format:"uuid4"`
	// Properties for a benefit of type `discord`.
	Properties BenefitGetResponseBenefitDiscordOutputProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                       `json:"selectable,required"`
	Type       BenefitGetResponseBenefitDiscordOutputType `json:"type,required"`
	JSON       benefitGetResponseBenefitDiscordOutputJSON `json:"-"`
}

// benefitGetResponseBenefitDiscordOutputJSON contains the JSON metadata for the
// struct [BenefitGetResponseBenefitDiscordOutput]
type benefitGetResponseBenefitDiscordOutputJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Deletable      apijson.Field
	Description    apijson.Field
	ModifiedAt     apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	Selectable     apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BenefitGetResponseBenefitDiscordOutput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitGetResponseBenefitDiscordOutputJSON) RawJSON() string {
	return r.raw
}

func (r BenefitGetResponseBenefitDiscordOutput) implementsBenefitGetResponse() {}

// Properties for a benefit of type `discord`.
type BenefitGetResponseBenefitDiscordOutputProperties struct {
	// The ID of the Discord server.
	GuildID    string `json:"guild_id,required"`
	GuildToken string `json:"guild_token,required"`
	// The ID of the Discord role to grant.
	RoleID string                                               `json:"role_id,required"`
	JSON   benefitGetResponseBenefitDiscordOutputPropertiesJSON `json:"-"`
}

// benefitGetResponseBenefitDiscordOutputPropertiesJSON contains the JSON metadata
// for the struct [BenefitGetResponseBenefitDiscordOutputProperties]
type benefitGetResponseBenefitDiscordOutputPropertiesJSON struct {
	GuildID     apijson.Field
	GuildToken  apijson.Field
	RoleID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BenefitGetResponseBenefitDiscordOutputProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitGetResponseBenefitDiscordOutputPropertiesJSON) RawJSON() string {
	return r.raw
}

type BenefitGetResponseBenefitDiscordOutputType string

const (
	BenefitGetResponseBenefitDiscordOutputTypeDiscord BenefitGetResponseBenefitDiscordOutputType = "discord"
)

func (r BenefitGetResponseBenefitDiscordOutputType) IsKnown() bool {
	switch r {
	case BenefitGetResponseBenefitDiscordOutputTypeDiscord:
		return true
	}
	return false
}

// A benefit of type `github_repository`.
//
// Use it to automatically invite your backers to a private GitHub repository.
type BenefitGetResponseBenefitGitHubRepository struct {
	// The ID of the benefit.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the benefit is deletable.
	Deletable bool `json:"deletable,required"`
	// The description of the benefit.
	Description string `json:"description,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The ID of the organization owning the benefit.
	OrganizationID string `json:"organization_id,required" format:"uuid4"`
	// Properties for a benefit of type `github_repository`.
	Properties BenefitGetResponseBenefitGitHubRepositoryProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                          `json:"selectable,required"`
	Type       BenefitGetResponseBenefitGitHubRepositoryType `json:"type,required"`
	JSON       benefitGetResponseBenefitGitHubRepositoryJSON `json:"-"`
}

// benefitGetResponseBenefitGitHubRepositoryJSON contains the JSON metadata for the
// struct [BenefitGetResponseBenefitGitHubRepository]
type benefitGetResponseBenefitGitHubRepositoryJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Deletable      apijson.Field
	Description    apijson.Field
	ModifiedAt     apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	Selectable     apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BenefitGetResponseBenefitGitHubRepository) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitGetResponseBenefitGitHubRepositoryJSON) RawJSON() string {
	return r.raw
}

func (r BenefitGetResponseBenefitGitHubRepository) implementsBenefitGetResponse() {}

// Properties for a benefit of type `github_repository`.
type BenefitGetResponseBenefitGitHubRepositoryProperties struct {
	// The permission level to grant. Read more about roles and their permissions on
	// [GitHub documentation](https://docs.github.com/en/organizations/managing-user-access-to-your-organizations-repositories/managing-repository-roles/repository-roles-for-an-organization#permissions-for-each-role).
	Permission   BenefitGetResponseBenefitGitHubRepositoryPropertiesPermission `json:"permission,required"`
	RepositoryID string                                                        `json:"repository_id,required,nullable" format:"uuid4"`
	// The name of the repository.
	RepositoryName string `json:"repository_name,required"`
	// The owner of the repository.
	RepositoryOwner string                                                  `json:"repository_owner,required"`
	JSON            benefitGetResponseBenefitGitHubRepositoryPropertiesJSON `json:"-"`
}

// benefitGetResponseBenefitGitHubRepositoryPropertiesJSON contains the JSON
// metadata for the struct [BenefitGetResponseBenefitGitHubRepositoryProperties]
type benefitGetResponseBenefitGitHubRepositoryPropertiesJSON struct {
	Permission      apijson.Field
	RepositoryID    apijson.Field
	RepositoryName  apijson.Field
	RepositoryOwner apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *BenefitGetResponseBenefitGitHubRepositoryProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitGetResponseBenefitGitHubRepositoryPropertiesJSON) RawJSON() string {
	return r.raw
}

// The permission level to grant. Read more about roles and their permissions on
// [GitHub documentation](https://docs.github.com/en/organizations/managing-user-access-to-your-organizations-repositories/managing-repository-roles/repository-roles-for-an-organization#permissions-for-each-role).
type BenefitGetResponseBenefitGitHubRepositoryPropertiesPermission string

const (
	BenefitGetResponseBenefitGitHubRepositoryPropertiesPermissionPull     BenefitGetResponseBenefitGitHubRepositoryPropertiesPermission = "pull"
	BenefitGetResponseBenefitGitHubRepositoryPropertiesPermissionTriage   BenefitGetResponseBenefitGitHubRepositoryPropertiesPermission = "triage"
	BenefitGetResponseBenefitGitHubRepositoryPropertiesPermissionPush     BenefitGetResponseBenefitGitHubRepositoryPropertiesPermission = "push"
	BenefitGetResponseBenefitGitHubRepositoryPropertiesPermissionMaintain BenefitGetResponseBenefitGitHubRepositoryPropertiesPermission = "maintain"
	BenefitGetResponseBenefitGitHubRepositoryPropertiesPermissionAdmin    BenefitGetResponseBenefitGitHubRepositoryPropertiesPermission = "admin"
)

func (r BenefitGetResponseBenefitGitHubRepositoryPropertiesPermission) IsKnown() bool {
	switch r {
	case BenefitGetResponseBenefitGitHubRepositoryPropertiesPermissionPull, BenefitGetResponseBenefitGitHubRepositoryPropertiesPermissionTriage, BenefitGetResponseBenefitGitHubRepositoryPropertiesPermissionPush, BenefitGetResponseBenefitGitHubRepositoryPropertiesPermissionMaintain, BenefitGetResponseBenefitGitHubRepositoryPropertiesPermissionAdmin:
		return true
	}
	return false
}

type BenefitGetResponseBenefitGitHubRepositoryType string

const (
	BenefitGetResponseBenefitGitHubRepositoryTypeGitHubRepository BenefitGetResponseBenefitGitHubRepositoryType = "github_repository"
)

func (r BenefitGetResponseBenefitGitHubRepositoryType) IsKnown() bool {
	switch r {
	case BenefitGetResponseBenefitGitHubRepositoryTypeGitHubRepository:
		return true
	}
	return false
}

type BenefitGetResponseBenefitDownloadables struct {
	// The ID of the benefit.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the benefit is deletable.
	Deletable bool `json:"deletable,required"`
	// The description of the benefit.
	Description string `json:"description,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The ID of the organization owning the benefit.
	OrganizationID string                                           `json:"organization_id,required" format:"uuid4"`
	Properties     BenefitGetResponseBenefitDownloadablesProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                       `json:"selectable,required"`
	Type       BenefitGetResponseBenefitDownloadablesType `json:"type,required"`
	JSON       benefitGetResponseBenefitDownloadablesJSON `json:"-"`
}

// benefitGetResponseBenefitDownloadablesJSON contains the JSON metadata for the
// struct [BenefitGetResponseBenefitDownloadables]
type benefitGetResponseBenefitDownloadablesJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Deletable      apijson.Field
	Description    apijson.Field
	ModifiedAt     apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	Selectable     apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BenefitGetResponseBenefitDownloadables) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitGetResponseBenefitDownloadablesJSON) RawJSON() string {
	return r.raw
}

func (r BenefitGetResponseBenefitDownloadables) implementsBenefitGetResponse() {}

type BenefitGetResponseBenefitDownloadablesProperties struct {
	Archived map[string]bool                                      `json:"archived,required"`
	Files    []string                                             `json:"files,required" format:"uuid4"`
	JSON     benefitGetResponseBenefitDownloadablesPropertiesJSON `json:"-"`
}

// benefitGetResponseBenefitDownloadablesPropertiesJSON contains the JSON metadata
// for the struct [BenefitGetResponseBenefitDownloadablesProperties]
type benefitGetResponseBenefitDownloadablesPropertiesJSON struct {
	Archived    apijson.Field
	Files       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BenefitGetResponseBenefitDownloadablesProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitGetResponseBenefitDownloadablesPropertiesJSON) RawJSON() string {
	return r.raw
}

type BenefitGetResponseBenefitDownloadablesType string

const (
	BenefitGetResponseBenefitDownloadablesTypeDownloadables BenefitGetResponseBenefitDownloadablesType = "downloadables"
)

func (r BenefitGetResponseBenefitDownloadablesType) IsKnown() bool {
	switch r {
	case BenefitGetResponseBenefitDownloadablesTypeDownloadables:
		return true
	}
	return false
}

type BenefitGetResponseBenefitLicenseKeysOutput struct {
	// The ID of the benefit.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the benefit is deletable.
	Deletable bool `json:"deletable,required"`
	// The description of the benefit.
	Description string `json:"description,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The ID of the organization owning the benefit.
	OrganizationID string                                               `json:"organization_id,required" format:"uuid4"`
	Properties     BenefitGetResponseBenefitLicenseKeysOutputProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                           `json:"selectable,required"`
	Type       BenefitGetResponseBenefitLicenseKeysOutputType `json:"type,required"`
	JSON       benefitGetResponseBenefitLicenseKeysOutputJSON `json:"-"`
}

// benefitGetResponseBenefitLicenseKeysOutputJSON contains the JSON metadata for
// the struct [BenefitGetResponseBenefitLicenseKeysOutput]
type benefitGetResponseBenefitLicenseKeysOutputJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Deletable      apijson.Field
	Description    apijson.Field
	ModifiedAt     apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	Selectable     apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BenefitGetResponseBenefitLicenseKeysOutput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitGetResponseBenefitLicenseKeysOutputJSON) RawJSON() string {
	return r.raw
}

func (r BenefitGetResponseBenefitLicenseKeysOutput) implementsBenefitGetResponse() {}

type BenefitGetResponseBenefitLicenseKeysOutputProperties struct {
	Activations BenefitGetResponseBenefitLicenseKeysOutputPropertiesActivations `json:"activations,required,nullable"`
	Expires     BenefitGetResponseBenefitLicenseKeysOutputPropertiesExpires     `json:"expires,required,nullable"`
	LimitUsage  int64                                                           `json:"limit_usage,required,nullable"`
	Prefix      string                                                          `json:"prefix,required,nullable"`
	JSON        benefitGetResponseBenefitLicenseKeysOutputPropertiesJSON        `json:"-"`
}

// benefitGetResponseBenefitLicenseKeysOutputPropertiesJSON contains the JSON
// metadata for the struct [BenefitGetResponseBenefitLicenseKeysOutputProperties]
type benefitGetResponseBenefitLicenseKeysOutputPropertiesJSON struct {
	Activations apijson.Field
	Expires     apijson.Field
	LimitUsage  apijson.Field
	Prefix      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BenefitGetResponseBenefitLicenseKeysOutputProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitGetResponseBenefitLicenseKeysOutputPropertiesJSON) RawJSON() string {
	return r.raw
}

type BenefitGetResponseBenefitLicenseKeysOutputPropertiesActivations struct {
	EnableUserAdmin bool                                                                `json:"enable_user_admin,required"`
	Limit           int64                                                               `json:"limit,required"`
	JSON            benefitGetResponseBenefitLicenseKeysOutputPropertiesActivationsJSON `json:"-"`
}

// benefitGetResponseBenefitLicenseKeysOutputPropertiesActivationsJSON contains the
// JSON metadata for the struct
// [BenefitGetResponseBenefitLicenseKeysOutputPropertiesActivations]
type benefitGetResponseBenefitLicenseKeysOutputPropertiesActivationsJSON struct {
	EnableUserAdmin apijson.Field
	Limit           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *BenefitGetResponseBenefitLicenseKeysOutputPropertiesActivations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitGetResponseBenefitLicenseKeysOutputPropertiesActivationsJSON) RawJSON() string {
	return r.raw
}

type BenefitGetResponseBenefitLicenseKeysOutputPropertiesExpires struct {
	Timeframe BenefitGetResponseBenefitLicenseKeysOutputPropertiesExpiresTimeframe `json:"timeframe,required"`
	Ttl       int64                                                                `json:"ttl,required"`
	JSON      benefitGetResponseBenefitLicenseKeysOutputPropertiesExpiresJSON      `json:"-"`
}

// benefitGetResponseBenefitLicenseKeysOutputPropertiesExpiresJSON contains the
// JSON metadata for the struct
// [BenefitGetResponseBenefitLicenseKeysOutputPropertiesExpires]
type benefitGetResponseBenefitLicenseKeysOutputPropertiesExpiresJSON struct {
	Timeframe   apijson.Field
	Ttl         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BenefitGetResponseBenefitLicenseKeysOutputPropertiesExpires) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitGetResponseBenefitLicenseKeysOutputPropertiesExpiresJSON) RawJSON() string {
	return r.raw
}

type BenefitGetResponseBenefitLicenseKeysOutputPropertiesExpiresTimeframe string

const (
	BenefitGetResponseBenefitLicenseKeysOutputPropertiesExpiresTimeframeYear  BenefitGetResponseBenefitLicenseKeysOutputPropertiesExpiresTimeframe = "year"
	BenefitGetResponseBenefitLicenseKeysOutputPropertiesExpiresTimeframeMonth BenefitGetResponseBenefitLicenseKeysOutputPropertiesExpiresTimeframe = "month"
	BenefitGetResponseBenefitLicenseKeysOutputPropertiesExpiresTimeframeDay   BenefitGetResponseBenefitLicenseKeysOutputPropertiesExpiresTimeframe = "day"
)

func (r BenefitGetResponseBenefitLicenseKeysOutputPropertiesExpiresTimeframe) IsKnown() bool {
	switch r {
	case BenefitGetResponseBenefitLicenseKeysOutputPropertiesExpiresTimeframeYear, BenefitGetResponseBenefitLicenseKeysOutputPropertiesExpiresTimeframeMonth, BenefitGetResponseBenefitLicenseKeysOutputPropertiesExpiresTimeframeDay:
		return true
	}
	return false
}

type BenefitGetResponseBenefitLicenseKeysOutputType string

const (
	BenefitGetResponseBenefitLicenseKeysOutputTypeLicenseKeys BenefitGetResponseBenefitLicenseKeysOutputType = "license_keys"
)

func (r BenefitGetResponseBenefitLicenseKeysOutputType) IsKnown() bool {
	switch r {
	case BenefitGetResponseBenefitLicenseKeysOutputTypeLicenseKeys:
		return true
	}
	return false
}

type BenefitGetResponseType string

const (
	BenefitGetResponseTypeArticles         BenefitGetResponseType = "articles"
	BenefitGetResponseTypeAds              BenefitGetResponseType = "ads"
	BenefitGetResponseTypeCustom           BenefitGetResponseType = "custom"
	BenefitGetResponseTypeDiscord          BenefitGetResponseType = "discord"
	BenefitGetResponseTypeGitHubRepository BenefitGetResponseType = "github_repository"
	BenefitGetResponseTypeDownloadables    BenefitGetResponseType = "downloadables"
	BenefitGetResponseTypeLicenseKeys      BenefitGetResponseType = "license_keys"
)

func (r BenefitGetResponseType) IsKnown() bool {
	switch r {
	case BenefitGetResponseTypeArticles, BenefitGetResponseTypeAds, BenefitGetResponseTypeCustom, BenefitGetResponseTypeDiscord, BenefitGetResponseTypeGitHubRepository, BenefitGetResponseTypeDownloadables, BenefitGetResponseTypeLicenseKeys:
		return true
	}
	return false
}

// A benefit of type `articles`.
//
// Use it to grant access to posts.
type BenefitUpdateResponse struct {
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The ID of the benefit.
	ID   string                    `json:"id,required" format:"uuid4"`
	Type BenefitUpdateResponseType `json:"type,required"`
	// The description of the benefit.
	Description string `json:"description,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool `json:"selectable,required"`
	// Whether the benefit is deletable.
	Deletable bool `json:"deletable,required"`
	// The ID of the organization owning the benefit.
	OrganizationID string `json:"organization_id,required" format:"uuid4"`
	// This field can have the runtime type of
	// [BenefitUpdateResponseBenefitArticlesProperties],
	// [BenefitUpdateResponseBenefitAdsProperties],
	// [BenefitUpdateResponseBenefitCustomProperties],
	// [BenefitUpdateResponseBenefitDiscordOutputProperties],
	// [BenefitUpdateResponseBenefitGitHubRepositoryProperties],
	// [BenefitUpdateResponseBenefitDownloadablesProperties],
	// [BenefitUpdateResponseBenefitLicenseKeysOutputProperties].
	Properties interface{} `json:"properties"`
	// Whether the benefit is taxable.
	IsTaxApplicable bool                      `json:"is_tax_applicable"`
	JSON            benefitUpdateResponseJSON `json:"-"`
	union           BenefitUpdateResponseUnion
}

// benefitUpdateResponseJSON contains the JSON metadata for the struct
// [BenefitUpdateResponse]
type benefitUpdateResponseJSON struct {
	CreatedAt       apijson.Field
	ModifiedAt      apijson.Field
	ID              apijson.Field
	Type            apijson.Field
	Description     apijson.Field
	Selectable      apijson.Field
	Deletable       apijson.Field
	OrganizationID  apijson.Field
	Properties      apijson.Field
	IsTaxApplicable apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r benefitUpdateResponseJSON) RawJSON() string {
	return r.raw
}

func (r *BenefitUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	*r = BenefitUpdateResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [BenefitUpdateResponseUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [BenefitUpdateResponseBenefitArticles],
// [BenefitUpdateResponseBenefitAds], [BenefitUpdateResponseBenefitCustom],
// [BenefitUpdateResponseBenefitDiscordOutput],
// [BenefitUpdateResponseBenefitGitHubRepository],
// [BenefitUpdateResponseBenefitDownloadables],
// [BenefitUpdateResponseBenefitLicenseKeysOutput].
func (r BenefitUpdateResponse) AsUnion() BenefitUpdateResponseUnion {
	return r.union
}

// A benefit of type `articles`.
//
// Use it to grant access to posts.
//
// Union satisfied by [BenefitUpdateResponseBenefitArticles],
// [BenefitUpdateResponseBenefitAds], [BenefitUpdateResponseBenefitCustom],
// [BenefitUpdateResponseBenefitDiscordOutput],
// [BenefitUpdateResponseBenefitGitHubRepository],
// [BenefitUpdateResponseBenefitDownloadables] or
// [BenefitUpdateResponseBenefitLicenseKeysOutput].
type BenefitUpdateResponseUnion interface {
	implementsBenefitUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BenefitUpdateResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BenefitUpdateResponseBenefitArticles{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BenefitUpdateResponseBenefitAds{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BenefitUpdateResponseBenefitCustom{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BenefitUpdateResponseBenefitDiscordOutput{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BenefitUpdateResponseBenefitGitHubRepository{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BenefitUpdateResponseBenefitDownloadables{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BenefitUpdateResponseBenefitLicenseKeysOutput{}),
		},
	)
}

// A benefit of type `articles`.
//
// Use it to grant access to posts.
type BenefitUpdateResponseBenefitArticles struct {
	// The ID of the benefit.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the benefit is deletable.
	Deletable bool `json:"deletable,required"`
	// The description of the benefit.
	Description string `json:"description,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The ID of the organization owning the benefit.
	OrganizationID string `json:"organization_id,required" format:"uuid4"`
	// Properties for a benefit of type `articles`.
	Properties BenefitUpdateResponseBenefitArticlesProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                     `json:"selectable,required"`
	Type       BenefitUpdateResponseBenefitArticlesType `json:"type,required"`
	JSON       benefitUpdateResponseBenefitArticlesJSON `json:"-"`
}

// benefitUpdateResponseBenefitArticlesJSON contains the JSON metadata for the
// struct [BenefitUpdateResponseBenefitArticles]
type benefitUpdateResponseBenefitArticlesJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Deletable      apijson.Field
	Description    apijson.Field
	ModifiedAt     apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	Selectable     apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BenefitUpdateResponseBenefitArticles) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitUpdateResponseBenefitArticlesJSON) RawJSON() string {
	return r.raw
}

func (r BenefitUpdateResponseBenefitArticles) implementsBenefitUpdateResponse() {}

// Properties for a benefit of type `articles`.
type BenefitUpdateResponseBenefitArticlesProperties struct {
	// Whether the user can access paid articles.
	PaidArticles bool                                               `json:"paid_articles,required"`
	JSON         benefitUpdateResponseBenefitArticlesPropertiesJSON `json:"-"`
}

// benefitUpdateResponseBenefitArticlesPropertiesJSON contains the JSON metadata
// for the struct [BenefitUpdateResponseBenefitArticlesProperties]
type benefitUpdateResponseBenefitArticlesPropertiesJSON struct {
	PaidArticles apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *BenefitUpdateResponseBenefitArticlesProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitUpdateResponseBenefitArticlesPropertiesJSON) RawJSON() string {
	return r.raw
}

type BenefitUpdateResponseBenefitArticlesType string

const (
	BenefitUpdateResponseBenefitArticlesTypeArticles BenefitUpdateResponseBenefitArticlesType = "articles"
)

func (r BenefitUpdateResponseBenefitArticlesType) IsKnown() bool {
	switch r {
	case BenefitUpdateResponseBenefitArticlesTypeArticles:
		return true
	}
	return false
}

// A benefit of type `ads`.
//
// Use it so your backers can display ads on your README, website, etc.
type BenefitUpdateResponseBenefitAds struct {
	// The ID of the benefit.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the benefit is deletable.
	Deletable bool `json:"deletable,required"`
	// The description of the benefit.
	Description string `json:"description,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The ID of the organization owning the benefit.
	OrganizationID string `json:"organization_id,required" format:"uuid4"`
	// Properties for a benefit of type `ads`.
	Properties BenefitUpdateResponseBenefitAdsProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                `json:"selectable,required"`
	Type       BenefitUpdateResponseBenefitAdsType `json:"type,required"`
	JSON       benefitUpdateResponseBenefitAdsJSON `json:"-"`
}

// benefitUpdateResponseBenefitAdsJSON contains the JSON metadata for the struct
// [BenefitUpdateResponseBenefitAds]
type benefitUpdateResponseBenefitAdsJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Deletable      apijson.Field
	Description    apijson.Field
	ModifiedAt     apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	Selectable     apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BenefitUpdateResponseBenefitAds) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitUpdateResponseBenefitAdsJSON) RawJSON() string {
	return r.raw
}

func (r BenefitUpdateResponseBenefitAds) implementsBenefitUpdateResponse() {}

// Properties for a benefit of type `ads`.
type BenefitUpdateResponseBenefitAdsProperties struct {
	// The height of the displayed ad.
	ImageHeight int64 `json:"image_height"`
	// The width of the displayed ad.
	ImageWidth int64                                         `json:"image_width"`
	JSON       benefitUpdateResponseBenefitAdsPropertiesJSON `json:"-"`
}

// benefitUpdateResponseBenefitAdsPropertiesJSON contains the JSON metadata for the
// struct [BenefitUpdateResponseBenefitAdsProperties]
type benefitUpdateResponseBenefitAdsPropertiesJSON struct {
	ImageHeight apijson.Field
	ImageWidth  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BenefitUpdateResponseBenefitAdsProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitUpdateResponseBenefitAdsPropertiesJSON) RawJSON() string {
	return r.raw
}

type BenefitUpdateResponseBenefitAdsType string

const (
	BenefitUpdateResponseBenefitAdsTypeAds BenefitUpdateResponseBenefitAdsType = "ads"
)

func (r BenefitUpdateResponseBenefitAdsType) IsKnown() bool {
	switch r {
	case BenefitUpdateResponseBenefitAdsTypeAds:
		return true
	}
	return false
}

// A benefit of type `custom`.
//
// Use it to grant any kind of benefit that doesn't fit in the other types.
type BenefitUpdateResponseBenefitCustom struct {
	// The ID of the benefit.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the benefit is deletable.
	Deletable bool `json:"deletable,required"`
	// The description of the benefit.
	Description string `json:"description,required"`
	// Whether the benefit is taxable.
	IsTaxApplicable bool `json:"is_tax_applicable,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The ID of the organization owning the benefit.
	OrganizationID string `json:"organization_id,required" format:"uuid4"`
	// Properties for a benefit of type `custom`.
	Properties BenefitUpdateResponseBenefitCustomProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                   `json:"selectable,required"`
	Type       BenefitUpdateResponseBenefitCustomType `json:"type,required"`
	JSON       benefitUpdateResponseBenefitCustomJSON `json:"-"`
}

// benefitUpdateResponseBenefitCustomJSON contains the JSON metadata for the struct
// [BenefitUpdateResponseBenefitCustom]
type benefitUpdateResponseBenefitCustomJSON struct {
	ID              apijson.Field
	CreatedAt       apijson.Field
	Deletable       apijson.Field
	Description     apijson.Field
	IsTaxApplicable apijson.Field
	ModifiedAt      apijson.Field
	OrganizationID  apijson.Field
	Properties      apijson.Field
	Selectable      apijson.Field
	Type            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *BenefitUpdateResponseBenefitCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitUpdateResponseBenefitCustomJSON) RawJSON() string {
	return r.raw
}

func (r BenefitUpdateResponseBenefitCustom) implementsBenefitUpdateResponse() {}

// Properties for a benefit of type `custom`.
type BenefitUpdateResponseBenefitCustomProperties struct {
	// Private note to be shared with users who have this benefit granted.
	Note string                                           `json:"note,required,nullable"`
	JSON benefitUpdateResponseBenefitCustomPropertiesJSON `json:"-"`
}

// benefitUpdateResponseBenefitCustomPropertiesJSON contains the JSON metadata for
// the struct [BenefitUpdateResponseBenefitCustomProperties]
type benefitUpdateResponseBenefitCustomPropertiesJSON struct {
	Note        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BenefitUpdateResponseBenefitCustomProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitUpdateResponseBenefitCustomPropertiesJSON) RawJSON() string {
	return r.raw
}

type BenefitUpdateResponseBenefitCustomType string

const (
	BenefitUpdateResponseBenefitCustomTypeCustom BenefitUpdateResponseBenefitCustomType = "custom"
)

func (r BenefitUpdateResponseBenefitCustomType) IsKnown() bool {
	switch r {
	case BenefitUpdateResponseBenefitCustomTypeCustom:
		return true
	}
	return false
}

// A benefit of type `discord`.
//
// Use it to automatically invite your backers to a Discord server.
type BenefitUpdateResponseBenefitDiscordOutput struct {
	// The ID of the benefit.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the benefit is deletable.
	Deletable bool `json:"deletable,required"`
	// The description of the benefit.
	Description string `json:"description,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The ID of the organization owning the benefit.
	OrganizationID string `json:"organization_id,required" format:"uuid4"`
	// Properties for a benefit of type `discord`.
	Properties BenefitUpdateResponseBenefitDiscordOutputProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                          `json:"selectable,required"`
	Type       BenefitUpdateResponseBenefitDiscordOutputType `json:"type,required"`
	JSON       benefitUpdateResponseBenefitDiscordOutputJSON `json:"-"`
}

// benefitUpdateResponseBenefitDiscordOutputJSON contains the JSON metadata for the
// struct [BenefitUpdateResponseBenefitDiscordOutput]
type benefitUpdateResponseBenefitDiscordOutputJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Deletable      apijson.Field
	Description    apijson.Field
	ModifiedAt     apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	Selectable     apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BenefitUpdateResponseBenefitDiscordOutput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitUpdateResponseBenefitDiscordOutputJSON) RawJSON() string {
	return r.raw
}

func (r BenefitUpdateResponseBenefitDiscordOutput) implementsBenefitUpdateResponse() {}

// Properties for a benefit of type `discord`.
type BenefitUpdateResponseBenefitDiscordOutputProperties struct {
	// The ID of the Discord server.
	GuildID    string `json:"guild_id,required"`
	GuildToken string `json:"guild_token,required"`
	// The ID of the Discord role to grant.
	RoleID string                                                  `json:"role_id,required"`
	JSON   benefitUpdateResponseBenefitDiscordOutputPropertiesJSON `json:"-"`
}

// benefitUpdateResponseBenefitDiscordOutputPropertiesJSON contains the JSON
// metadata for the struct [BenefitUpdateResponseBenefitDiscordOutputProperties]
type benefitUpdateResponseBenefitDiscordOutputPropertiesJSON struct {
	GuildID     apijson.Field
	GuildToken  apijson.Field
	RoleID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BenefitUpdateResponseBenefitDiscordOutputProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitUpdateResponseBenefitDiscordOutputPropertiesJSON) RawJSON() string {
	return r.raw
}

type BenefitUpdateResponseBenefitDiscordOutputType string

const (
	BenefitUpdateResponseBenefitDiscordOutputTypeDiscord BenefitUpdateResponseBenefitDiscordOutputType = "discord"
)

func (r BenefitUpdateResponseBenefitDiscordOutputType) IsKnown() bool {
	switch r {
	case BenefitUpdateResponseBenefitDiscordOutputTypeDiscord:
		return true
	}
	return false
}

// A benefit of type `github_repository`.
//
// Use it to automatically invite your backers to a private GitHub repository.
type BenefitUpdateResponseBenefitGitHubRepository struct {
	// The ID of the benefit.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the benefit is deletable.
	Deletable bool `json:"deletable,required"`
	// The description of the benefit.
	Description string `json:"description,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The ID of the organization owning the benefit.
	OrganizationID string `json:"organization_id,required" format:"uuid4"`
	// Properties for a benefit of type `github_repository`.
	Properties BenefitUpdateResponseBenefitGitHubRepositoryProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                             `json:"selectable,required"`
	Type       BenefitUpdateResponseBenefitGitHubRepositoryType `json:"type,required"`
	JSON       benefitUpdateResponseBenefitGitHubRepositoryJSON `json:"-"`
}

// benefitUpdateResponseBenefitGitHubRepositoryJSON contains the JSON metadata for
// the struct [BenefitUpdateResponseBenefitGitHubRepository]
type benefitUpdateResponseBenefitGitHubRepositoryJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Deletable      apijson.Field
	Description    apijson.Field
	ModifiedAt     apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	Selectable     apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BenefitUpdateResponseBenefitGitHubRepository) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitUpdateResponseBenefitGitHubRepositoryJSON) RawJSON() string {
	return r.raw
}

func (r BenefitUpdateResponseBenefitGitHubRepository) implementsBenefitUpdateResponse() {}

// Properties for a benefit of type `github_repository`.
type BenefitUpdateResponseBenefitGitHubRepositoryProperties struct {
	// The permission level to grant. Read more about roles and their permissions on
	// [GitHub documentation](https://docs.github.com/en/organizations/managing-user-access-to-your-organizations-repositories/managing-repository-roles/repository-roles-for-an-organization#permissions-for-each-role).
	Permission   BenefitUpdateResponseBenefitGitHubRepositoryPropertiesPermission `json:"permission,required"`
	RepositoryID string                                                           `json:"repository_id,required,nullable" format:"uuid4"`
	// The name of the repository.
	RepositoryName string `json:"repository_name,required"`
	// The owner of the repository.
	RepositoryOwner string                                                     `json:"repository_owner,required"`
	JSON            benefitUpdateResponseBenefitGitHubRepositoryPropertiesJSON `json:"-"`
}

// benefitUpdateResponseBenefitGitHubRepositoryPropertiesJSON contains the JSON
// metadata for the struct [BenefitUpdateResponseBenefitGitHubRepositoryProperties]
type benefitUpdateResponseBenefitGitHubRepositoryPropertiesJSON struct {
	Permission      apijson.Field
	RepositoryID    apijson.Field
	RepositoryName  apijson.Field
	RepositoryOwner apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *BenefitUpdateResponseBenefitGitHubRepositoryProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitUpdateResponseBenefitGitHubRepositoryPropertiesJSON) RawJSON() string {
	return r.raw
}

// The permission level to grant. Read more about roles and their permissions on
// [GitHub documentation](https://docs.github.com/en/organizations/managing-user-access-to-your-organizations-repositories/managing-repository-roles/repository-roles-for-an-organization#permissions-for-each-role).
type BenefitUpdateResponseBenefitGitHubRepositoryPropertiesPermission string

const (
	BenefitUpdateResponseBenefitGitHubRepositoryPropertiesPermissionPull     BenefitUpdateResponseBenefitGitHubRepositoryPropertiesPermission = "pull"
	BenefitUpdateResponseBenefitGitHubRepositoryPropertiesPermissionTriage   BenefitUpdateResponseBenefitGitHubRepositoryPropertiesPermission = "triage"
	BenefitUpdateResponseBenefitGitHubRepositoryPropertiesPermissionPush     BenefitUpdateResponseBenefitGitHubRepositoryPropertiesPermission = "push"
	BenefitUpdateResponseBenefitGitHubRepositoryPropertiesPermissionMaintain BenefitUpdateResponseBenefitGitHubRepositoryPropertiesPermission = "maintain"
	BenefitUpdateResponseBenefitGitHubRepositoryPropertiesPermissionAdmin    BenefitUpdateResponseBenefitGitHubRepositoryPropertiesPermission = "admin"
)

func (r BenefitUpdateResponseBenefitGitHubRepositoryPropertiesPermission) IsKnown() bool {
	switch r {
	case BenefitUpdateResponseBenefitGitHubRepositoryPropertiesPermissionPull, BenefitUpdateResponseBenefitGitHubRepositoryPropertiesPermissionTriage, BenefitUpdateResponseBenefitGitHubRepositoryPropertiesPermissionPush, BenefitUpdateResponseBenefitGitHubRepositoryPropertiesPermissionMaintain, BenefitUpdateResponseBenefitGitHubRepositoryPropertiesPermissionAdmin:
		return true
	}
	return false
}

type BenefitUpdateResponseBenefitGitHubRepositoryType string

const (
	BenefitUpdateResponseBenefitGitHubRepositoryTypeGitHubRepository BenefitUpdateResponseBenefitGitHubRepositoryType = "github_repository"
)

func (r BenefitUpdateResponseBenefitGitHubRepositoryType) IsKnown() bool {
	switch r {
	case BenefitUpdateResponseBenefitGitHubRepositoryTypeGitHubRepository:
		return true
	}
	return false
}

type BenefitUpdateResponseBenefitDownloadables struct {
	// The ID of the benefit.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the benefit is deletable.
	Deletable bool `json:"deletable,required"`
	// The description of the benefit.
	Description string `json:"description,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The ID of the organization owning the benefit.
	OrganizationID string                                              `json:"organization_id,required" format:"uuid4"`
	Properties     BenefitUpdateResponseBenefitDownloadablesProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                          `json:"selectable,required"`
	Type       BenefitUpdateResponseBenefitDownloadablesType `json:"type,required"`
	JSON       benefitUpdateResponseBenefitDownloadablesJSON `json:"-"`
}

// benefitUpdateResponseBenefitDownloadablesJSON contains the JSON metadata for the
// struct [BenefitUpdateResponseBenefitDownloadables]
type benefitUpdateResponseBenefitDownloadablesJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Deletable      apijson.Field
	Description    apijson.Field
	ModifiedAt     apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	Selectable     apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BenefitUpdateResponseBenefitDownloadables) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitUpdateResponseBenefitDownloadablesJSON) RawJSON() string {
	return r.raw
}

func (r BenefitUpdateResponseBenefitDownloadables) implementsBenefitUpdateResponse() {}

type BenefitUpdateResponseBenefitDownloadablesProperties struct {
	Archived map[string]bool                                         `json:"archived,required"`
	Files    []string                                                `json:"files,required" format:"uuid4"`
	JSON     benefitUpdateResponseBenefitDownloadablesPropertiesJSON `json:"-"`
}

// benefitUpdateResponseBenefitDownloadablesPropertiesJSON contains the JSON
// metadata for the struct [BenefitUpdateResponseBenefitDownloadablesProperties]
type benefitUpdateResponseBenefitDownloadablesPropertiesJSON struct {
	Archived    apijson.Field
	Files       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BenefitUpdateResponseBenefitDownloadablesProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitUpdateResponseBenefitDownloadablesPropertiesJSON) RawJSON() string {
	return r.raw
}

type BenefitUpdateResponseBenefitDownloadablesType string

const (
	BenefitUpdateResponseBenefitDownloadablesTypeDownloadables BenefitUpdateResponseBenefitDownloadablesType = "downloadables"
)

func (r BenefitUpdateResponseBenefitDownloadablesType) IsKnown() bool {
	switch r {
	case BenefitUpdateResponseBenefitDownloadablesTypeDownloadables:
		return true
	}
	return false
}

type BenefitUpdateResponseBenefitLicenseKeysOutput struct {
	// The ID of the benefit.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the benefit is deletable.
	Deletable bool `json:"deletable,required"`
	// The description of the benefit.
	Description string `json:"description,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The ID of the organization owning the benefit.
	OrganizationID string                                                  `json:"organization_id,required" format:"uuid4"`
	Properties     BenefitUpdateResponseBenefitLicenseKeysOutputProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                              `json:"selectable,required"`
	Type       BenefitUpdateResponseBenefitLicenseKeysOutputType `json:"type,required"`
	JSON       benefitUpdateResponseBenefitLicenseKeysOutputJSON `json:"-"`
}

// benefitUpdateResponseBenefitLicenseKeysOutputJSON contains the JSON metadata for
// the struct [BenefitUpdateResponseBenefitLicenseKeysOutput]
type benefitUpdateResponseBenefitLicenseKeysOutputJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Deletable      apijson.Field
	Description    apijson.Field
	ModifiedAt     apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	Selectable     apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BenefitUpdateResponseBenefitLicenseKeysOutput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitUpdateResponseBenefitLicenseKeysOutputJSON) RawJSON() string {
	return r.raw
}

func (r BenefitUpdateResponseBenefitLicenseKeysOutput) implementsBenefitUpdateResponse() {}

type BenefitUpdateResponseBenefitLicenseKeysOutputProperties struct {
	Activations BenefitUpdateResponseBenefitLicenseKeysOutputPropertiesActivations `json:"activations,required,nullable"`
	Expires     BenefitUpdateResponseBenefitLicenseKeysOutputPropertiesExpires     `json:"expires,required,nullable"`
	LimitUsage  int64                                                              `json:"limit_usage,required,nullable"`
	Prefix      string                                                             `json:"prefix,required,nullable"`
	JSON        benefitUpdateResponseBenefitLicenseKeysOutputPropertiesJSON        `json:"-"`
}

// benefitUpdateResponseBenefitLicenseKeysOutputPropertiesJSON contains the JSON
// metadata for the struct
// [BenefitUpdateResponseBenefitLicenseKeysOutputProperties]
type benefitUpdateResponseBenefitLicenseKeysOutputPropertiesJSON struct {
	Activations apijson.Field
	Expires     apijson.Field
	LimitUsage  apijson.Field
	Prefix      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BenefitUpdateResponseBenefitLicenseKeysOutputProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitUpdateResponseBenefitLicenseKeysOutputPropertiesJSON) RawJSON() string {
	return r.raw
}

type BenefitUpdateResponseBenefitLicenseKeysOutputPropertiesActivations struct {
	EnableUserAdmin bool                                                                   `json:"enable_user_admin,required"`
	Limit           int64                                                                  `json:"limit,required"`
	JSON            benefitUpdateResponseBenefitLicenseKeysOutputPropertiesActivationsJSON `json:"-"`
}

// benefitUpdateResponseBenefitLicenseKeysOutputPropertiesActivationsJSON contains
// the JSON metadata for the struct
// [BenefitUpdateResponseBenefitLicenseKeysOutputPropertiesActivations]
type benefitUpdateResponseBenefitLicenseKeysOutputPropertiesActivationsJSON struct {
	EnableUserAdmin apijson.Field
	Limit           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *BenefitUpdateResponseBenefitLicenseKeysOutputPropertiesActivations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitUpdateResponseBenefitLicenseKeysOutputPropertiesActivationsJSON) RawJSON() string {
	return r.raw
}

type BenefitUpdateResponseBenefitLicenseKeysOutputPropertiesExpires struct {
	Timeframe BenefitUpdateResponseBenefitLicenseKeysOutputPropertiesExpiresTimeframe `json:"timeframe,required"`
	Ttl       int64                                                                   `json:"ttl,required"`
	JSON      benefitUpdateResponseBenefitLicenseKeysOutputPropertiesExpiresJSON      `json:"-"`
}

// benefitUpdateResponseBenefitLicenseKeysOutputPropertiesExpiresJSON contains the
// JSON metadata for the struct
// [BenefitUpdateResponseBenefitLicenseKeysOutputPropertiesExpires]
type benefitUpdateResponseBenefitLicenseKeysOutputPropertiesExpiresJSON struct {
	Timeframe   apijson.Field
	Ttl         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BenefitUpdateResponseBenefitLicenseKeysOutputPropertiesExpires) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitUpdateResponseBenefitLicenseKeysOutputPropertiesExpiresJSON) RawJSON() string {
	return r.raw
}

type BenefitUpdateResponseBenefitLicenseKeysOutputPropertiesExpiresTimeframe string

const (
	BenefitUpdateResponseBenefitLicenseKeysOutputPropertiesExpiresTimeframeYear  BenefitUpdateResponseBenefitLicenseKeysOutputPropertiesExpiresTimeframe = "year"
	BenefitUpdateResponseBenefitLicenseKeysOutputPropertiesExpiresTimeframeMonth BenefitUpdateResponseBenefitLicenseKeysOutputPropertiesExpiresTimeframe = "month"
	BenefitUpdateResponseBenefitLicenseKeysOutputPropertiesExpiresTimeframeDay   BenefitUpdateResponseBenefitLicenseKeysOutputPropertiesExpiresTimeframe = "day"
)

func (r BenefitUpdateResponseBenefitLicenseKeysOutputPropertiesExpiresTimeframe) IsKnown() bool {
	switch r {
	case BenefitUpdateResponseBenefitLicenseKeysOutputPropertiesExpiresTimeframeYear, BenefitUpdateResponseBenefitLicenseKeysOutputPropertiesExpiresTimeframeMonth, BenefitUpdateResponseBenefitLicenseKeysOutputPropertiesExpiresTimeframeDay:
		return true
	}
	return false
}

type BenefitUpdateResponseBenefitLicenseKeysOutputType string

const (
	BenefitUpdateResponseBenefitLicenseKeysOutputTypeLicenseKeys BenefitUpdateResponseBenefitLicenseKeysOutputType = "license_keys"
)

func (r BenefitUpdateResponseBenefitLicenseKeysOutputType) IsKnown() bool {
	switch r {
	case BenefitUpdateResponseBenefitLicenseKeysOutputTypeLicenseKeys:
		return true
	}
	return false
}

type BenefitUpdateResponseType string

const (
	BenefitUpdateResponseTypeArticles         BenefitUpdateResponseType = "articles"
	BenefitUpdateResponseTypeAds              BenefitUpdateResponseType = "ads"
	BenefitUpdateResponseTypeCustom           BenefitUpdateResponseType = "custom"
	BenefitUpdateResponseTypeDiscord          BenefitUpdateResponseType = "discord"
	BenefitUpdateResponseTypeGitHubRepository BenefitUpdateResponseType = "github_repository"
	BenefitUpdateResponseTypeDownloadables    BenefitUpdateResponseType = "downloadables"
	BenefitUpdateResponseTypeLicenseKeys      BenefitUpdateResponseType = "license_keys"
)

func (r BenefitUpdateResponseType) IsKnown() bool {
	switch r {
	case BenefitUpdateResponseTypeArticles, BenefitUpdateResponseTypeAds, BenefitUpdateResponseTypeCustom, BenefitUpdateResponseTypeDiscord, BenefitUpdateResponseTypeGitHubRepository, BenefitUpdateResponseTypeDownloadables, BenefitUpdateResponseTypeLicenseKeys:
		return true
	}
	return false
}

// A benefit of type `articles`.
//
// Use it to grant access to posts.
type BenefitListResponse struct {
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The ID of the benefit.
	ID   string                  `json:"id,required" format:"uuid4"`
	Type BenefitListResponseType `json:"type,required"`
	// The description of the benefit.
	Description string `json:"description,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool `json:"selectable,required"`
	// Whether the benefit is deletable.
	Deletable bool `json:"deletable,required"`
	// The ID of the organization owning the benefit.
	OrganizationID string `json:"organization_id,required" format:"uuid4"`
	// This field can have the runtime type of
	// [BenefitListResponseBenefitArticlesProperties],
	// [BenefitListResponseBenefitAdsProperties],
	// [BenefitListResponseBenefitCustomProperties],
	// [BenefitListResponseBenefitDiscordOutputProperties],
	// [BenefitListResponseBenefitGitHubRepositoryProperties],
	// [BenefitListResponseBenefitDownloadablesProperties],
	// [BenefitListResponseBenefitLicenseKeysOutputProperties].
	Properties interface{} `json:"properties"`
	// Whether the benefit is taxable.
	IsTaxApplicable bool                    `json:"is_tax_applicable"`
	JSON            benefitListResponseJSON `json:"-"`
	union           BenefitListResponseUnion
}

// benefitListResponseJSON contains the JSON metadata for the struct
// [BenefitListResponse]
type benefitListResponseJSON struct {
	CreatedAt       apijson.Field
	ModifiedAt      apijson.Field
	ID              apijson.Field
	Type            apijson.Field
	Description     apijson.Field
	Selectable      apijson.Field
	Deletable       apijson.Field
	OrganizationID  apijson.Field
	Properties      apijson.Field
	IsTaxApplicable apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r benefitListResponseJSON) RawJSON() string {
	return r.raw
}

func (r *BenefitListResponse) UnmarshalJSON(data []byte) (err error) {
	*r = BenefitListResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [BenefitListResponseUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [BenefitListResponseBenefitArticles],
// [BenefitListResponseBenefitAds], [BenefitListResponseBenefitCustom],
// [BenefitListResponseBenefitDiscordOutput],
// [BenefitListResponseBenefitGitHubRepository],
// [BenefitListResponseBenefitDownloadables],
// [BenefitListResponseBenefitLicenseKeysOutput].
func (r BenefitListResponse) AsUnion() BenefitListResponseUnion {
	return r.union
}

// A benefit of type `articles`.
//
// Use it to grant access to posts.
//
// Union satisfied by [BenefitListResponseBenefitArticles],
// [BenefitListResponseBenefitAds], [BenefitListResponseBenefitCustom],
// [BenefitListResponseBenefitDiscordOutput],
// [BenefitListResponseBenefitGitHubRepository],
// [BenefitListResponseBenefitDownloadables] or
// [BenefitListResponseBenefitLicenseKeysOutput].
type BenefitListResponseUnion interface {
	implementsBenefitListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BenefitListResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BenefitListResponseBenefitArticles{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BenefitListResponseBenefitAds{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BenefitListResponseBenefitCustom{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BenefitListResponseBenefitDiscordOutput{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BenefitListResponseBenefitGitHubRepository{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BenefitListResponseBenefitDownloadables{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BenefitListResponseBenefitLicenseKeysOutput{}),
		},
	)
}

// A benefit of type `articles`.
//
// Use it to grant access to posts.
type BenefitListResponseBenefitArticles struct {
	// The ID of the benefit.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the benefit is deletable.
	Deletable bool `json:"deletable,required"`
	// The description of the benefit.
	Description string `json:"description,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The ID of the organization owning the benefit.
	OrganizationID string `json:"organization_id,required" format:"uuid4"`
	// Properties for a benefit of type `articles`.
	Properties BenefitListResponseBenefitArticlesProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                   `json:"selectable,required"`
	Type       BenefitListResponseBenefitArticlesType `json:"type,required"`
	JSON       benefitListResponseBenefitArticlesJSON `json:"-"`
}

// benefitListResponseBenefitArticlesJSON contains the JSON metadata for the struct
// [BenefitListResponseBenefitArticles]
type benefitListResponseBenefitArticlesJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Deletable      apijson.Field
	Description    apijson.Field
	ModifiedAt     apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	Selectable     apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BenefitListResponseBenefitArticles) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitListResponseBenefitArticlesJSON) RawJSON() string {
	return r.raw
}

func (r BenefitListResponseBenefitArticles) implementsBenefitListResponse() {}

// Properties for a benefit of type `articles`.
type BenefitListResponseBenefitArticlesProperties struct {
	// Whether the user can access paid articles.
	PaidArticles bool                                             `json:"paid_articles,required"`
	JSON         benefitListResponseBenefitArticlesPropertiesJSON `json:"-"`
}

// benefitListResponseBenefitArticlesPropertiesJSON contains the JSON metadata for
// the struct [BenefitListResponseBenefitArticlesProperties]
type benefitListResponseBenefitArticlesPropertiesJSON struct {
	PaidArticles apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *BenefitListResponseBenefitArticlesProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitListResponseBenefitArticlesPropertiesJSON) RawJSON() string {
	return r.raw
}

type BenefitListResponseBenefitArticlesType string

const (
	BenefitListResponseBenefitArticlesTypeArticles BenefitListResponseBenefitArticlesType = "articles"
)

func (r BenefitListResponseBenefitArticlesType) IsKnown() bool {
	switch r {
	case BenefitListResponseBenefitArticlesTypeArticles:
		return true
	}
	return false
}

// A benefit of type `ads`.
//
// Use it so your backers can display ads on your README, website, etc.
type BenefitListResponseBenefitAds struct {
	// The ID of the benefit.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the benefit is deletable.
	Deletable bool `json:"deletable,required"`
	// The description of the benefit.
	Description string `json:"description,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The ID of the organization owning the benefit.
	OrganizationID string `json:"organization_id,required" format:"uuid4"`
	// Properties for a benefit of type `ads`.
	Properties BenefitListResponseBenefitAdsProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                              `json:"selectable,required"`
	Type       BenefitListResponseBenefitAdsType `json:"type,required"`
	JSON       benefitListResponseBenefitAdsJSON `json:"-"`
}

// benefitListResponseBenefitAdsJSON contains the JSON metadata for the struct
// [BenefitListResponseBenefitAds]
type benefitListResponseBenefitAdsJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Deletable      apijson.Field
	Description    apijson.Field
	ModifiedAt     apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	Selectable     apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BenefitListResponseBenefitAds) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitListResponseBenefitAdsJSON) RawJSON() string {
	return r.raw
}

func (r BenefitListResponseBenefitAds) implementsBenefitListResponse() {}

// Properties for a benefit of type `ads`.
type BenefitListResponseBenefitAdsProperties struct {
	// The height of the displayed ad.
	ImageHeight int64 `json:"image_height"`
	// The width of the displayed ad.
	ImageWidth int64                                       `json:"image_width"`
	JSON       benefitListResponseBenefitAdsPropertiesJSON `json:"-"`
}

// benefitListResponseBenefitAdsPropertiesJSON contains the JSON metadata for the
// struct [BenefitListResponseBenefitAdsProperties]
type benefitListResponseBenefitAdsPropertiesJSON struct {
	ImageHeight apijson.Field
	ImageWidth  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BenefitListResponseBenefitAdsProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitListResponseBenefitAdsPropertiesJSON) RawJSON() string {
	return r.raw
}

type BenefitListResponseBenefitAdsType string

const (
	BenefitListResponseBenefitAdsTypeAds BenefitListResponseBenefitAdsType = "ads"
)

func (r BenefitListResponseBenefitAdsType) IsKnown() bool {
	switch r {
	case BenefitListResponseBenefitAdsTypeAds:
		return true
	}
	return false
}

// A benefit of type `custom`.
//
// Use it to grant any kind of benefit that doesn't fit in the other types.
type BenefitListResponseBenefitCustom struct {
	// The ID of the benefit.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the benefit is deletable.
	Deletable bool `json:"deletable,required"`
	// The description of the benefit.
	Description string `json:"description,required"`
	// Whether the benefit is taxable.
	IsTaxApplicable bool `json:"is_tax_applicable,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The ID of the organization owning the benefit.
	OrganizationID string `json:"organization_id,required" format:"uuid4"`
	// Properties for a benefit of type `custom`.
	Properties BenefitListResponseBenefitCustomProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                 `json:"selectable,required"`
	Type       BenefitListResponseBenefitCustomType `json:"type,required"`
	JSON       benefitListResponseBenefitCustomJSON `json:"-"`
}

// benefitListResponseBenefitCustomJSON contains the JSON metadata for the struct
// [BenefitListResponseBenefitCustom]
type benefitListResponseBenefitCustomJSON struct {
	ID              apijson.Field
	CreatedAt       apijson.Field
	Deletable       apijson.Field
	Description     apijson.Field
	IsTaxApplicable apijson.Field
	ModifiedAt      apijson.Field
	OrganizationID  apijson.Field
	Properties      apijson.Field
	Selectable      apijson.Field
	Type            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *BenefitListResponseBenefitCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitListResponseBenefitCustomJSON) RawJSON() string {
	return r.raw
}

func (r BenefitListResponseBenefitCustom) implementsBenefitListResponse() {}

// Properties for a benefit of type `custom`.
type BenefitListResponseBenefitCustomProperties struct {
	// Private note to be shared with users who have this benefit granted.
	Note string                                         `json:"note,required,nullable"`
	JSON benefitListResponseBenefitCustomPropertiesJSON `json:"-"`
}

// benefitListResponseBenefitCustomPropertiesJSON contains the JSON metadata for
// the struct [BenefitListResponseBenefitCustomProperties]
type benefitListResponseBenefitCustomPropertiesJSON struct {
	Note        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BenefitListResponseBenefitCustomProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitListResponseBenefitCustomPropertiesJSON) RawJSON() string {
	return r.raw
}

type BenefitListResponseBenefitCustomType string

const (
	BenefitListResponseBenefitCustomTypeCustom BenefitListResponseBenefitCustomType = "custom"
)

func (r BenefitListResponseBenefitCustomType) IsKnown() bool {
	switch r {
	case BenefitListResponseBenefitCustomTypeCustom:
		return true
	}
	return false
}

// A benefit of type `discord`.
//
// Use it to automatically invite your backers to a Discord server.
type BenefitListResponseBenefitDiscordOutput struct {
	// The ID of the benefit.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the benefit is deletable.
	Deletable bool `json:"deletable,required"`
	// The description of the benefit.
	Description string `json:"description,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The ID of the organization owning the benefit.
	OrganizationID string `json:"organization_id,required" format:"uuid4"`
	// Properties for a benefit of type `discord`.
	Properties BenefitListResponseBenefitDiscordOutputProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                        `json:"selectable,required"`
	Type       BenefitListResponseBenefitDiscordOutputType `json:"type,required"`
	JSON       benefitListResponseBenefitDiscordOutputJSON `json:"-"`
}

// benefitListResponseBenefitDiscordOutputJSON contains the JSON metadata for the
// struct [BenefitListResponseBenefitDiscordOutput]
type benefitListResponseBenefitDiscordOutputJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Deletable      apijson.Field
	Description    apijson.Field
	ModifiedAt     apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	Selectable     apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BenefitListResponseBenefitDiscordOutput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitListResponseBenefitDiscordOutputJSON) RawJSON() string {
	return r.raw
}

func (r BenefitListResponseBenefitDiscordOutput) implementsBenefitListResponse() {}

// Properties for a benefit of type `discord`.
type BenefitListResponseBenefitDiscordOutputProperties struct {
	// The ID of the Discord server.
	GuildID    string `json:"guild_id,required"`
	GuildToken string `json:"guild_token,required"`
	// The ID of the Discord role to grant.
	RoleID string                                                `json:"role_id,required"`
	JSON   benefitListResponseBenefitDiscordOutputPropertiesJSON `json:"-"`
}

// benefitListResponseBenefitDiscordOutputPropertiesJSON contains the JSON metadata
// for the struct [BenefitListResponseBenefitDiscordOutputProperties]
type benefitListResponseBenefitDiscordOutputPropertiesJSON struct {
	GuildID     apijson.Field
	GuildToken  apijson.Field
	RoleID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BenefitListResponseBenefitDiscordOutputProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitListResponseBenefitDiscordOutputPropertiesJSON) RawJSON() string {
	return r.raw
}

type BenefitListResponseBenefitDiscordOutputType string

const (
	BenefitListResponseBenefitDiscordOutputTypeDiscord BenefitListResponseBenefitDiscordOutputType = "discord"
)

func (r BenefitListResponseBenefitDiscordOutputType) IsKnown() bool {
	switch r {
	case BenefitListResponseBenefitDiscordOutputTypeDiscord:
		return true
	}
	return false
}

// A benefit of type `github_repository`.
//
// Use it to automatically invite your backers to a private GitHub repository.
type BenefitListResponseBenefitGitHubRepository struct {
	// The ID of the benefit.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the benefit is deletable.
	Deletable bool `json:"deletable,required"`
	// The description of the benefit.
	Description string `json:"description,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The ID of the organization owning the benefit.
	OrganizationID string `json:"organization_id,required" format:"uuid4"`
	// Properties for a benefit of type `github_repository`.
	Properties BenefitListResponseBenefitGitHubRepositoryProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                           `json:"selectable,required"`
	Type       BenefitListResponseBenefitGitHubRepositoryType `json:"type,required"`
	JSON       benefitListResponseBenefitGitHubRepositoryJSON `json:"-"`
}

// benefitListResponseBenefitGitHubRepositoryJSON contains the JSON metadata for
// the struct [BenefitListResponseBenefitGitHubRepository]
type benefitListResponseBenefitGitHubRepositoryJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Deletable      apijson.Field
	Description    apijson.Field
	ModifiedAt     apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	Selectable     apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BenefitListResponseBenefitGitHubRepository) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitListResponseBenefitGitHubRepositoryJSON) RawJSON() string {
	return r.raw
}

func (r BenefitListResponseBenefitGitHubRepository) implementsBenefitListResponse() {}

// Properties for a benefit of type `github_repository`.
type BenefitListResponseBenefitGitHubRepositoryProperties struct {
	// The permission level to grant. Read more about roles and their permissions on
	// [GitHub documentation](https://docs.github.com/en/organizations/managing-user-access-to-your-organizations-repositories/managing-repository-roles/repository-roles-for-an-organization#permissions-for-each-role).
	Permission   BenefitListResponseBenefitGitHubRepositoryPropertiesPermission `json:"permission,required"`
	RepositoryID string                                                         `json:"repository_id,required,nullable" format:"uuid4"`
	// The name of the repository.
	RepositoryName string `json:"repository_name,required"`
	// The owner of the repository.
	RepositoryOwner string                                                   `json:"repository_owner,required"`
	JSON            benefitListResponseBenefitGitHubRepositoryPropertiesJSON `json:"-"`
}

// benefitListResponseBenefitGitHubRepositoryPropertiesJSON contains the JSON
// metadata for the struct [BenefitListResponseBenefitGitHubRepositoryProperties]
type benefitListResponseBenefitGitHubRepositoryPropertiesJSON struct {
	Permission      apijson.Field
	RepositoryID    apijson.Field
	RepositoryName  apijson.Field
	RepositoryOwner apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *BenefitListResponseBenefitGitHubRepositoryProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitListResponseBenefitGitHubRepositoryPropertiesJSON) RawJSON() string {
	return r.raw
}

// The permission level to grant. Read more about roles and their permissions on
// [GitHub documentation](https://docs.github.com/en/organizations/managing-user-access-to-your-organizations-repositories/managing-repository-roles/repository-roles-for-an-organization#permissions-for-each-role).
type BenefitListResponseBenefitGitHubRepositoryPropertiesPermission string

const (
	BenefitListResponseBenefitGitHubRepositoryPropertiesPermissionPull     BenefitListResponseBenefitGitHubRepositoryPropertiesPermission = "pull"
	BenefitListResponseBenefitGitHubRepositoryPropertiesPermissionTriage   BenefitListResponseBenefitGitHubRepositoryPropertiesPermission = "triage"
	BenefitListResponseBenefitGitHubRepositoryPropertiesPermissionPush     BenefitListResponseBenefitGitHubRepositoryPropertiesPermission = "push"
	BenefitListResponseBenefitGitHubRepositoryPropertiesPermissionMaintain BenefitListResponseBenefitGitHubRepositoryPropertiesPermission = "maintain"
	BenefitListResponseBenefitGitHubRepositoryPropertiesPermissionAdmin    BenefitListResponseBenefitGitHubRepositoryPropertiesPermission = "admin"
)

func (r BenefitListResponseBenefitGitHubRepositoryPropertiesPermission) IsKnown() bool {
	switch r {
	case BenefitListResponseBenefitGitHubRepositoryPropertiesPermissionPull, BenefitListResponseBenefitGitHubRepositoryPropertiesPermissionTriage, BenefitListResponseBenefitGitHubRepositoryPropertiesPermissionPush, BenefitListResponseBenefitGitHubRepositoryPropertiesPermissionMaintain, BenefitListResponseBenefitGitHubRepositoryPropertiesPermissionAdmin:
		return true
	}
	return false
}

type BenefitListResponseBenefitGitHubRepositoryType string

const (
	BenefitListResponseBenefitGitHubRepositoryTypeGitHubRepository BenefitListResponseBenefitGitHubRepositoryType = "github_repository"
)

func (r BenefitListResponseBenefitGitHubRepositoryType) IsKnown() bool {
	switch r {
	case BenefitListResponseBenefitGitHubRepositoryTypeGitHubRepository:
		return true
	}
	return false
}

type BenefitListResponseBenefitDownloadables struct {
	// The ID of the benefit.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the benefit is deletable.
	Deletable bool `json:"deletable,required"`
	// The description of the benefit.
	Description string `json:"description,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The ID of the organization owning the benefit.
	OrganizationID string                                            `json:"organization_id,required" format:"uuid4"`
	Properties     BenefitListResponseBenefitDownloadablesProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                        `json:"selectable,required"`
	Type       BenefitListResponseBenefitDownloadablesType `json:"type,required"`
	JSON       benefitListResponseBenefitDownloadablesJSON `json:"-"`
}

// benefitListResponseBenefitDownloadablesJSON contains the JSON metadata for the
// struct [BenefitListResponseBenefitDownloadables]
type benefitListResponseBenefitDownloadablesJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Deletable      apijson.Field
	Description    apijson.Field
	ModifiedAt     apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	Selectable     apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BenefitListResponseBenefitDownloadables) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitListResponseBenefitDownloadablesJSON) RawJSON() string {
	return r.raw
}

func (r BenefitListResponseBenefitDownloadables) implementsBenefitListResponse() {}

type BenefitListResponseBenefitDownloadablesProperties struct {
	Archived map[string]bool                                       `json:"archived,required"`
	Files    []string                                              `json:"files,required" format:"uuid4"`
	JSON     benefitListResponseBenefitDownloadablesPropertiesJSON `json:"-"`
}

// benefitListResponseBenefitDownloadablesPropertiesJSON contains the JSON metadata
// for the struct [BenefitListResponseBenefitDownloadablesProperties]
type benefitListResponseBenefitDownloadablesPropertiesJSON struct {
	Archived    apijson.Field
	Files       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BenefitListResponseBenefitDownloadablesProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitListResponseBenefitDownloadablesPropertiesJSON) RawJSON() string {
	return r.raw
}

type BenefitListResponseBenefitDownloadablesType string

const (
	BenefitListResponseBenefitDownloadablesTypeDownloadables BenefitListResponseBenefitDownloadablesType = "downloadables"
)

func (r BenefitListResponseBenefitDownloadablesType) IsKnown() bool {
	switch r {
	case BenefitListResponseBenefitDownloadablesTypeDownloadables:
		return true
	}
	return false
}

type BenefitListResponseBenefitLicenseKeysOutput struct {
	// The ID of the benefit.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the benefit is deletable.
	Deletable bool `json:"deletable,required"`
	// The description of the benefit.
	Description string `json:"description,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The ID of the organization owning the benefit.
	OrganizationID string                                                `json:"organization_id,required" format:"uuid4"`
	Properties     BenefitListResponseBenefitLicenseKeysOutputProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                            `json:"selectable,required"`
	Type       BenefitListResponseBenefitLicenseKeysOutputType `json:"type,required"`
	JSON       benefitListResponseBenefitLicenseKeysOutputJSON `json:"-"`
}

// benefitListResponseBenefitLicenseKeysOutputJSON contains the JSON metadata for
// the struct [BenefitListResponseBenefitLicenseKeysOutput]
type benefitListResponseBenefitLicenseKeysOutputJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Deletable      apijson.Field
	Description    apijson.Field
	ModifiedAt     apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	Selectable     apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BenefitListResponseBenefitLicenseKeysOutput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitListResponseBenefitLicenseKeysOutputJSON) RawJSON() string {
	return r.raw
}

func (r BenefitListResponseBenefitLicenseKeysOutput) implementsBenefitListResponse() {}

type BenefitListResponseBenefitLicenseKeysOutputProperties struct {
	Activations BenefitListResponseBenefitLicenseKeysOutputPropertiesActivations `json:"activations,required,nullable"`
	Expires     BenefitListResponseBenefitLicenseKeysOutputPropertiesExpires     `json:"expires,required,nullable"`
	LimitUsage  int64                                                            `json:"limit_usage,required,nullable"`
	Prefix      string                                                           `json:"prefix,required,nullable"`
	JSON        benefitListResponseBenefitLicenseKeysOutputPropertiesJSON        `json:"-"`
}

// benefitListResponseBenefitLicenseKeysOutputPropertiesJSON contains the JSON
// metadata for the struct [BenefitListResponseBenefitLicenseKeysOutputProperties]
type benefitListResponseBenefitLicenseKeysOutputPropertiesJSON struct {
	Activations apijson.Field
	Expires     apijson.Field
	LimitUsage  apijson.Field
	Prefix      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BenefitListResponseBenefitLicenseKeysOutputProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitListResponseBenefitLicenseKeysOutputPropertiesJSON) RawJSON() string {
	return r.raw
}

type BenefitListResponseBenefitLicenseKeysOutputPropertiesActivations struct {
	EnableUserAdmin bool                                                                 `json:"enable_user_admin,required"`
	Limit           int64                                                                `json:"limit,required"`
	JSON            benefitListResponseBenefitLicenseKeysOutputPropertiesActivationsJSON `json:"-"`
}

// benefitListResponseBenefitLicenseKeysOutputPropertiesActivationsJSON contains
// the JSON metadata for the struct
// [BenefitListResponseBenefitLicenseKeysOutputPropertiesActivations]
type benefitListResponseBenefitLicenseKeysOutputPropertiesActivationsJSON struct {
	EnableUserAdmin apijson.Field
	Limit           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *BenefitListResponseBenefitLicenseKeysOutputPropertiesActivations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitListResponseBenefitLicenseKeysOutputPropertiesActivationsJSON) RawJSON() string {
	return r.raw
}

type BenefitListResponseBenefitLicenseKeysOutputPropertiesExpires struct {
	Timeframe BenefitListResponseBenefitLicenseKeysOutputPropertiesExpiresTimeframe `json:"timeframe,required"`
	Ttl       int64                                                                 `json:"ttl,required"`
	JSON      benefitListResponseBenefitLicenseKeysOutputPropertiesExpiresJSON      `json:"-"`
}

// benefitListResponseBenefitLicenseKeysOutputPropertiesExpiresJSON contains the
// JSON metadata for the struct
// [BenefitListResponseBenefitLicenseKeysOutputPropertiesExpires]
type benefitListResponseBenefitLicenseKeysOutputPropertiesExpiresJSON struct {
	Timeframe   apijson.Field
	Ttl         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BenefitListResponseBenefitLicenseKeysOutputPropertiesExpires) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitListResponseBenefitLicenseKeysOutputPropertiesExpiresJSON) RawJSON() string {
	return r.raw
}

type BenefitListResponseBenefitLicenseKeysOutputPropertiesExpiresTimeframe string

const (
	BenefitListResponseBenefitLicenseKeysOutputPropertiesExpiresTimeframeYear  BenefitListResponseBenefitLicenseKeysOutputPropertiesExpiresTimeframe = "year"
	BenefitListResponseBenefitLicenseKeysOutputPropertiesExpiresTimeframeMonth BenefitListResponseBenefitLicenseKeysOutputPropertiesExpiresTimeframe = "month"
	BenefitListResponseBenefitLicenseKeysOutputPropertiesExpiresTimeframeDay   BenefitListResponseBenefitLicenseKeysOutputPropertiesExpiresTimeframe = "day"
)

func (r BenefitListResponseBenefitLicenseKeysOutputPropertiesExpiresTimeframe) IsKnown() bool {
	switch r {
	case BenefitListResponseBenefitLicenseKeysOutputPropertiesExpiresTimeframeYear, BenefitListResponseBenefitLicenseKeysOutputPropertiesExpiresTimeframeMonth, BenefitListResponseBenefitLicenseKeysOutputPropertiesExpiresTimeframeDay:
		return true
	}
	return false
}

type BenefitListResponseBenefitLicenseKeysOutputType string

const (
	BenefitListResponseBenefitLicenseKeysOutputTypeLicenseKeys BenefitListResponseBenefitLicenseKeysOutputType = "license_keys"
)

func (r BenefitListResponseBenefitLicenseKeysOutputType) IsKnown() bool {
	switch r {
	case BenefitListResponseBenefitLicenseKeysOutputTypeLicenseKeys:
		return true
	}
	return false
}

type BenefitListResponseType string

const (
	BenefitListResponseTypeArticles         BenefitListResponseType = "articles"
	BenefitListResponseTypeAds              BenefitListResponseType = "ads"
	BenefitListResponseTypeCustom           BenefitListResponseType = "custom"
	BenefitListResponseTypeDiscord          BenefitListResponseType = "discord"
	BenefitListResponseTypeGitHubRepository BenefitListResponseType = "github_repository"
	BenefitListResponseTypeDownloadables    BenefitListResponseType = "downloadables"
	BenefitListResponseTypeLicenseKeys      BenefitListResponseType = "license_keys"
)

func (r BenefitListResponseType) IsKnown() bool {
	switch r {
	case BenefitListResponseTypeArticles, BenefitListResponseTypeAds, BenefitListResponseTypeCustom, BenefitListResponseTypeDiscord, BenefitListResponseTypeGitHubRepository, BenefitListResponseTypeDownloadables, BenefitListResponseTypeLicenseKeys:
		return true
	}
	return false
}

type BenefitNewParams struct {
	// Schema to create a benefit of type `custom`.
	Body BenefitNewParamsBodyUnion `json:"body,required"`
}

func (r BenefitNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// Schema to create a benefit of type `custom`.
type BenefitNewParamsBody struct {
	Type param.Field[BenefitNewParamsBodyType] `json:"type,required"`
	// The description of the benefit. Will be displayed on products having this
	// benefit.
	Description param.Field[string] `json:"description,required"`
	// The organization ID.
	OrganizationID param.Field[string] `json:"organization_id" format:"uuid4"`
	// Whether the benefit is taxable.
	IsTaxApplicable param.Field[bool]        `json:"is_tax_applicable"`
	Properties      param.Field[interface{}] `json:"properties"`
}

func (r BenefitNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BenefitNewParamsBody) implementsBenefitNewParamsBodyUnion() {}

// Schema to create a benefit of type `custom`.
//
// Satisfied by [BenefitNewParamsBodyBenefitCustomCreate],
// [BenefitNewParamsBodyBenefitAdsCreate],
// [BenefitNewParamsBodyBenefitDiscordCreate],
// [BenefitNewParamsBodyBenefitGitHubRepositoryCreate],
// [BenefitNewParamsBodyBenefitDownloadablesCreate],
// [BenefitNewParamsBodyBenefitLicenseKeysCreate], [BenefitNewParamsBody].
type BenefitNewParamsBodyUnion interface {
	implementsBenefitNewParamsBodyUnion()
}

// Schema to create a benefit of type `custom`.
type BenefitNewParamsBodyBenefitCustomCreate struct {
	// The description of the benefit. Will be displayed on products having this
	// benefit.
	Description param.Field[string] `json:"description,required"`
	// Whether the benefit is taxable.
	IsTaxApplicable param.Field[bool] `json:"is_tax_applicable,required"`
	// Properties for a benefit of type `custom`.
	Properties param.Field[BenefitNewParamsBodyBenefitCustomCreateProperties] `json:"properties,required"`
	Type       param.Field[BenefitNewParamsBodyBenefitCustomCreateType]       `json:"type,required"`
	// The organization ID.
	OrganizationID param.Field[string] `json:"organization_id" format:"uuid4"`
}

func (r BenefitNewParamsBodyBenefitCustomCreate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BenefitNewParamsBodyBenefitCustomCreate) implementsBenefitNewParamsBodyUnion() {}

// Properties for a benefit of type `custom`.
type BenefitNewParamsBodyBenefitCustomCreateProperties struct {
	// Private note to be shared with users who have this benefit granted.
	Note param.Field[string] `json:"note,required"`
}

func (r BenefitNewParamsBodyBenefitCustomCreateProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BenefitNewParamsBodyBenefitCustomCreateType string

const (
	BenefitNewParamsBodyBenefitCustomCreateTypeCustom BenefitNewParamsBodyBenefitCustomCreateType = "custom"
)

func (r BenefitNewParamsBodyBenefitCustomCreateType) IsKnown() bool {
	switch r {
	case BenefitNewParamsBodyBenefitCustomCreateTypeCustom:
		return true
	}
	return false
}

type BenefitNewParamsBodyBenefitAdsCreate struct {
	// The description of the benefit. Will be displayed on products having this
	// benefit.
	Description param.Field[string] `json:"description,required"`
	// Properties for a benefit of type `ads`.
	Properties param.Field[BenefitNewParamsBodyBenefitAdsCreateProperties] `json:"properties,required"`
	Type       param.Field[BenefitNewParamsBodyBenefitAdsCreateType]       `json:"type,required"`
	// The organization ID.
	OrganizationID param.Field[string] `json:"organization_id" format:"uuid4"`
}

func (r BenefitNewParamsBodyBenefitAdsCreate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BenefitNewParamsBodyBenefitAdsCreate) implementsBenefitNewParamsBodyUnion() {}

// Properties for a benefit of type `ads`.
type BenefitNewParamsBodyBenefitAdsCreateProperties struct {
	// The height of the displayed ad.
	ImageHeight param.Field[int64] `json:"image_height"`
	// The width of the displayed ad.
	ImageWidth param.Field[int64] `json:"image_width"`
}

func (r BenefitNewParamsBodyBenefitAdsCreateProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BenefitNewParamsBodyBenefitAdsCreateType string

const (
	BenefitNewParamsBodyBenefitAdsCreateTypeAds BenefitNewParamsBodyBenefitAdsCreateType = "ads"
)

func (r BenefitNewParamsBodyBenefitAdsCreateType) IsKnown() bool {
	switch r {
	case BenefitNewParamsBodyBenefitAdsCreateTypeAds:
		return true
	}
	return false
}

type BenefitNewParamsBodyBenefitDiscordCreate struct {
	// The description of the benefit. Will be displayed on products having this
	// benefit.
	Description param.Field[string] `json:"description,required"`
	// Properties to create a benefit of type `discord`.
	Properties param.Field[BenefitNewParamsBodyBenefitDiscordCreateProperties] `json:"properties,required"`
	Type       param.Field[BenefitNewParamsBodyBenefitDiscordCreateType]       `json:"type,required"`
	// The organization ID.
	OrganizationID param.Field[string] `json:"organization_id" format:"uuid4"`
}

func (r BenefitNewParamsBodyBenefitDiscordCreate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BenefitNewParamsBodyBenefitDiscordCreate) implementsBenefitNewParamsBodyUnion() {}

// Properties to create a benefit of type `discord`.
type BenefitNewParamsBodyBenefitDiscordCreateProperties struct {
	GuildToken param.Field[string] `json:"guild_token,required"`
	// The ID of the Discord role to grant.
	RoleID param.Field[string] `json:"role_id,required"`
}

func (r BenefitNewParamsBodyBenefitDiscordCreateProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BenefitNewParamsBodyBenefitDiscordCreateType string

const (
	BenefitNewParamsBodyBenefitDiscordCreateTypeDiscord BenefitNewParamsBodyBenefitDiscordCreateType = "discord"
)

func (r BenefitNewParamsBodyBenefitDiscordCreateType) IsKnown() bool {
	switch r {
	case BenefitNewParamsBodyBenefitDiscordCreateTypeDiscord:
		return true
	}
	return false
}

type BenefitNewParamsBodyBenefitGitHubRepositoryCreate struct {
	// The description of the benefit. Will be displayed on products having this
	// benefit.
	Description param.Field[string] `json:"description,required"`
	// Properties to create a benefit of type `github_repository`.
	Properties param.Field[BenefitNewParamsBodyBenefitGitHubRepositoryCreateProperties] `json:"properties,required"`
	Type       param.Field[BenefitNewParamsBodyBenefitGitHubRepositoryCreateType]       `json:"type,required"`
	// The organization ID.
	OrganizationID param.Field[string] `json:"organization_id" format:"uuid4"`
}

func (r BenefitNewParamsBodyBenefitGitHubRepositoryCreate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BenefitNewParamsBodyBenefitGitHubRepositoryCreate) implementsBenefitNewParamsBodyUnion() {}

// Properties to create a benefit of type `github_repository`.
type BenefitNewParamsBodyBenefitGitHubRepositoryCreateProperties struct {
	// The permission level to grant. Read more about roles and their permissions on
	// [GitHub documentation](https://docs.github.com/en/organizations/managing-user-access-to-your-organizations-repositories/managing-repository-roles/repository-roles-for-an-organization#permissions-for-each-role).
	Permission   param.Field[BenefitNewParamsBodyBenefitGitHubRepositoryCreatePropertiesPermission] `json:"permission,required"`
	RepositoryID param.Field[string]                                                                `json:"repository_id" format:"uuid4"`
	// The name of the repository.
	RepositoryName param.Field[string] `json:"repository_name"`
	// The owner of the repository.
	RepositoryOwner param.Field[string] `json:"repository_owner"`
}

func (r BenefitNewParamsBodyBenefitGitHubRepositoryCreateProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The permission level to grant. Read more about roles and their permissions on
// [GitHub documentation](https://docs.github.com/en/organizations/managing-user-access-to-your-organizations-repositories/managing-repository-roles/repository-roles-for-an-organization#permissions-for-each-role).
type BenefitNewParamsBodyBenefitGitHubRepositoryCreatePropertiesPermission string

const (
	BenefitNewParamsBodyBenefitGitHubRepositoryCreatePropertiesPermissionPull     BenefitNewParamsBodyBenefitGitHubRepositoryCreatePropertiesPermission = "pull"
	BenefitNewParamsBodyBenefitGitHubRepositoryCreatePropertiesPermissionTriage   BenefitNewParamsBodyBenefitGitHubRepositoryCreatePropertiesPermission = "triage"
	BenefitNewParamsBodyBenefitGitHubRepositoryCreatePropertiesPermissionPush     BenefitNewParamsBodyBenefitGitHubRepositoryCreatePropertiesPermission = "push"
	BenefitNewParamsBodyBenefitGitHubRepositoryCreatePropertiesPermissionMaintain BenefitNewParamsBodyBenefitGitHubRepositoryCreatePropertiesPermission = "maintain"
	BenefitNewParamsBodyBenefitGitHubRepositoryCreatePropertiesPermissionAdmin    BenefitNewParamsBodyBenefitGitHubRepositoryCreatePropertiesPermission = "admin"
)

func (r BenefitNewParamsBodyBenefitGitHubRepositoryCreatePropertiesPermission) IsKnown() bool {
	switch r {
	case BenefitNewParamsBodyBenefitGitHubRepositoryCreatePropertiesPermissionPull, BenefitNewParamsBodyBenefitGitHubRepositoryCreatePropertiesPermissionTriage, BenefitNewParamsBodyBenefitGitHubRepositoryCreatePropertiesPermissionPush, BenefitNewParamsBodyBenefitGitHubRepositoryCreatePropertiesPermissionMaintain, BenefitNewParamsBodyBenefitGitHubRepositoryCreatePropertiesPermissionAdmin:
		return true
	}
	return false
}

type BenefitNewParamsBodyBenefitGitHubRepositoryCreateType string

const (
	BenefitNewParamsBodyBenefitGitHubRepositoryCreateTypeGitHubRepository BenefitNewParamsBodyBenefitGitHubRepositoryCreateType = "github_repository"
)

func (r BenefitNewParamsBodyBenefitGitHubRepositoryCreateType) IsKnown() bool {
	switch r {
	case BenefitNewParamsBodyBenefitGitHubRepositoryCreateTypeGitHubRepository:
		return true
	}
	return false
}

type BenefitNewParamsBodyBenefitDownloadablesCreate struct {
	// The description of the benefit. Will be displayed on products having this
	// benefit.
	Description param.Field[string]                                                   `json:"description,required"`
	Properties  param.Field[BenefitNewParamsBodyBenefitDownloadablesCreateProperties] `json:"properties,required"`
	Type        param.Field[BenefitNewParamsBodyBenefitDownloadablesCreateType]       `json:"type,required"`
	// The organization ID.
	OrganizationID param.Field[string] `json:"organization_id" format:"uuid4"`
}

func (r BenefitNewParamsBodyBenefitDownloadablesCreate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BenefitNewParamsBodyBenefitDownloadablesCreate) implementsBenefitNewParamsBodyUnion() {}

type BenefitNewParamsBodyBenefitDownloadablesCreateProperties struct {
	Files    param.Field[[]string]        `json:"files,required" format:"uuid4"`
	Archived param.Field[map[string]bool] `json:"archived"`
}

func (r BenefitNewParamsBodyBenefitDownloadablesCreateProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BenefitNewParamsBodyBenefitDownloadablesCreateType string

const (
	BenefitNewParamsBodyBenefitDownloadablesCreateTypeDownloadables BenefitNewParamsBodyBenefitDownloadablesCreateType = "downloadables"
)

func (r BenefitNewParamsBodyBenefitDownloadablesCreateType) IsKnown() bool {
	switch r {
	case BenefitNewParamsBodyBenefitDownloadablesCreateTypeDownloadables:
		return true
	}
	return false
}

type BenefitNewParamsBodyBenefitLicenseKeysCreate struct {
	// The description of the benefit. Will be displayed on products having this
	// benefit.
	Description param.Field[string]                                                 `json:"description,required"`
	Properties  param.Field[BenefitNewParamsBodyBenefitLicenseKeysCreateProperties] `json:"properties,required"`
	Type        param.Field[BenefitNewParamsBodyBenefitLicenseKeysCreateType]       `json:"type,required"`
	// The organization ID.
	OrganizationID param.Field[string] `json:"organization_id" format:"uuid4"`
}

func (r BenefitNewParamsBodyBenefitLicenseKeysCreate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BenefitNewParamsBodyBenefitLicenseKeysCreate) implementsBenefitNewParamsBodyUnion() {}

type BenefitNewParamsBodyBenefitLicenseKeysCreateProperties struct {
	Activations param.Field[BenefitNewParamsBodyBenefitLicenseKeysCreatePropertiesActivations] `json:"activations"`
	Expires     param.Field[BenefitNewParamsBodyBenefitLicenseKeysCreatePropertiesExpires]     `json:"expires"`
	LimitUsage  param.Field[int64]                                                             `json:"limit_usage"`
	Prefix      param.Field[string]                                                            `json:"prefix"`
}

func (r BenefitNewParamsBodyBenefitLicenseKeysCreateProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BenefitNewParamsBodyBenefitLicenseKeysCreatePropertiesActivations struct {
	EnableUserAdmin param.Field[bool]  `json:"enable_user_admin,required"`
	Limit           param.Field[int64] `json:"limit,required"`
}

func (r BenefitNewParamsBodyBenefitLicenseKeysCreatePropertiesActivations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BenefitNewParamsBodyBenefitLicenseKeysCreatePropertiesExpires struct {
	Timeframe param.Field[BenefitNewParamsBodyBenefitLicenseKeysCreatePropertiesExpiresTimeframe] `json:"timeframe,required"`
	Ttl       param.Field[int64]                                                                  `json:"ttl,required"`
}

func (r BenefitNewParamsBodyBenefitLicenseKeysCreatePropertiesExpires) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BenefitNewParamsBodyBenefitLicenseKeysCreatePropertiesExpiresTimeframe string

const (
	BenefitNewParamsBodyBenefitLicenseKeysCreatePropertiesExpiresTimeframeYear  BenefitNewParamsBodyBenefitLicenseKeysCreatePropertiesExpiresTimeframe = "year"
	BenefitNewParamsBodyBenefitLicenseKeysCreatePropertiesExpiresTimeframeMonth BenefitNewParamsBodyBenefitLicenseKeysCreatePropertiesExpiresTimeframe = "month"
	BenefitNewParamsBodyBenefitLicenseKeysCreatePropertiesExpiresTimeframeDay   BenefitNewParamsBodyBenefitLicenseKeysCreatePropertiesExpiresTimeframe = "day"
)

func (r BenefitNewParamsBodyBenefitLicenseKeysCreatePropertiesExpiresTimeframe) IsKnown() bool {
	switch r {
	case BenefitNewParamsBodyBenefitLicenseKeysCreatePropertiesExpiresTimeframeYear, BenefitNewParamsBodyBenefitLicenseKeysCreatePropertiesExpiresTimeframeMonth, BenefitNewParamsBodyBenefitLicenseKeysCreatePropertiesExpiresTimeframeDay:
		return true
	}
	return false
}

type BenefitNewParamsBodyBenefitLicenseKeysCreateType string

const (
	BenefitNewParamsBodyBenefitLicenseKeysCreateTypeLicenseKeys BenefitNewParamsBodyBenefitLicenseKeysCreateType = "license_keys"
)

func (r BenefitNewParamsBodyBenefitLicenseKeysCreateType) IsKnown() bool {
	switch r {
	case BenefitNewParamsBodyBenefitLicenseKeysCreateTypeLicenseKeys:
		return true
	}
	return false
}

type BenefitNewParamsBodyType string

const (
	BenefitNewParamsBodyTypeCustom           BenefitNewParamsBodyType = "custom"
	BenefitNewParamsBodyTypeAds              BenefitNewParamsBodyType = "ads"
	BenefitNewParamsBodyTypeDiscord          BenefitNewParamsBodyType = "discord"
	BenefitNewParamsBodyTypeGitHubRepository BenefitNewParamsBodyType = "github_repository"
	BenefitNewParamsBodyTypeDownloadables    BenefitNewParamsBodyType = "downloadables"
	BenefitNewParamsBodyTypeLicenseKeys      BenefitNewParamsBodyType = "license_keys"
)

func (r BenefitNewParamsBodyType) IsKnown() bool {
	switch r {
	case BenefitNewParamsBodyTypeCustom, BenefitNewParamsBodyTypeAds, BenefitNewParamsBodyTypeDiscord, BenefitNewParamsBodyTypeGitHubRepository, BenefitNewParamsBodyTypeDownloadables, BenefitNewParamsBodyTypeLicenseKeys:
		return true
	}
	return false
}

type BenefitUpdateParams struct {
	Body BenefitUpdateParamsBodyUnion `json:"body,required"`
}

func (r BenefitUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type BenefitUpdateParamsBody struct {
	// The description of the benefit. Will be displayed on products having this
	// benefit.
	Description param.Field[string]                      `json:"description"`
	Type        param.Field[BenefitUpdateParamsBodyType] `json:"type,required"`
	Properties  param.Field[interface{}]                 `json:"properties,required"`
}

func (r BenefitUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BenefitUpdateParamsBody) implementsBenefitUpdateParamsBodyUnion() {}

// Satisfied by [BenefitUpdateParamsBodyBenefitArticlesUpdate],
// [BenefitUpdateParamsBodyBenefitAdsUpdate],
// [BenefitUpdateParamsBodyBenefitCustomUpdate],
// [BenefitUpdateParamsBodyBenefitDiscordUpdate],
// [BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdate],
// [BenefitUpdateParamsBodyBenefitDownloadablesUpdate],
// [BenefitUpdateParamsBodyBenefitLicenseKeysUpdate], [BenefitUpdateParamsBody].
type BenefitUpdateParamsBodyUnion interface {
	implementsBenefitUpdateParamsBodyUnion()
}

type BenefitUpdateParamsBodyBenefitArticlesUpdate struct {
	Type param.Field[BenefitUpdateParamsBodyBenefitArticlesUpdateType] `json:"type,required"`
	// The description of the benefit. Will be displayed on products having this
	// benefit.
	Description param.Field[string] `json:"description"`
}

func (r BenefitUpdateParamsBodyBenefitArticlesUpdate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BenefitUpdateParamsBodyBenefitArticlesUpdate) implementsBenefitUpdateParamsBodyUnion() {}

type BenefitUpdateParamsBodyBenefitArticlesUpdateType string

const (
	BenefitUpdateParamsBodyBenefitArticlesUpdateTypeArticles BenefitUpdateParamsBodyBenefitArticlesUpdateType = "articles"
)

func (r BenefitUpdateParamsBodyBenefitArticlesUpdateType) IsKnown() bool {
	switch r {
	case BenefitUpdateParamsBodyBenefitArticlesUpdateTypeArticles:
		return true
	}
	return false
}

type BenefitUpdateParamsBodyBenefitAdsUpdate struct {
	Type param.Field[BenefitUpdateParamsBodyBenefitAdsUpdateType] `json:"type,required"`
	// The description of the benefit. Will be displayed on products having this
	// benefit.
	Description param.Field[string] `json:"description"`
	// Properties for a benefit of type `ads`.
	Properties param.Field[BenefitUpdateParamsBodyBenefitAdsUpdateProperties] `json:"properties"`
}

func (r BenefitUpdateParamsBodyBenefitAdsUpdate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BenefitUpdateParamsBodyBenefitAdsUpdate) implementsBenefitUpdateParamsBodyUnion() {}

type BenefitUpdateParamsBodyBenefitAdsUpdateType string

const (
	BenefitUpdateParamsBodyBenefitAdsUpdateTypeAds BenefitUpdateParamsBodyBenefitAdsUpdateType = "ads"
)

func (r BenefitUpdateParamsBodyBenefitAdsUpdateType) IsKnown() bool {
	switch r {
	case BenefitUpdateParamsBodyBenefitAdsUpdateTypeAds:
		return true
	}
	return false
}

// Properties for a benefit of type `ads`.
type BenefitUpdateParamsBodyBenefitAdsUpdateProperties struct {
	// The height of the displayed ad.
	ImageHeight param.Field[int64] `json:"image_height"`
	// The width of the displayed ad.
	ImageWidth param.Field[int64] `json:"image_width"`
}

func (r BenefitUpdateParamsBodyBenefitAdsUpdateProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BenefitUpdateParamsBodyBenefitCustomUpdate struct {
	Type param.Field[BenefitUpdateParamsBodyBenefitCustomUpdateType] `json:"type,required"`
	// The description of the benefit. Will be displayed on products having this
	// benefit.
	Description param.Field[string] `json:"description"`
	// Properties for a benefit of type `custom`.
	Properties param.Field[BenefitUpdateParamsBodyBenefitCustomUpdateProperties] `json:"properties"`
}

func (r BenefitUpdateParamsBodyBenefitCustomUpdate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BenefitUpdateParamsBodyBenefitCustomUpdate) implementsBenefitUpdateParamsBodyUnion() {}

type BenefitUpdateParamsBodyBenefitCustomUpdateType string

const (
	BenefitUpdateParamsBodyBenefitCustomUpdateTypeCustom BenefitUpdateParamsBodyBenefitCustomUpdateType = "custom"
)

func (r BenefitUpdateParamsBodyBenefitCustomUpdateType) IsKnown() bool {
	switch r {
	case BenefitUpdateParamsBodyBenefitCustomUpdateTypeCustom:
		return true
	}
	return false
}

// Properties for a benefit of type `custom`.
type BenefitUpdateParamsBodyBenefitCustomUpdateProperties struct {
	// Private note to be shared with users who have this benefit granted.
	Note param.Field[string] `json:"note,required"`
}

func (r BenefitUpdateParamsBodyBenefitCustomUpdateProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BenefitUpdateParamsBodyBenefitDiscordUpdate struct {
	Type param.Field[BenefitUpdateParamsBodyBenefitDiscordUpdateType] `json:"type,required"`
	// The description of the benefit. Will be displayed on products having this
	// benefit.
	Description param.Field[string] `json:"description"`
	// Properties to create a benefit of type `discord`.
	Properties param.Field[BenefitUpdateParamsBodyBenefitDiscordUpdateProperties] `json:"properties"`
}

func (r BenefitUpdateParamsBodyBenefitDiscordUpdate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BenefitUpdateParamsBodyBenefitDiscordUpdate) implementsBenefitUpdateParamsBodyUnion() {}

type BenefitUpdateParamsBodyBenefitDiscordUpdateType string

const (
	BenefitUpdateParamsBodyBenefitDiscordUpdateTypeDiscord BenefitUpdateParamsBodyBenefitDiscordUpdateType = "discord"
)

func (r BenefitUpdateParamsBodyBenefitDiscordUpdateType) IsKnown() bool {
	switch r {
	case BenefitUpdateParamsBodyBenefitDiscordUpdateTypeDiscord:
		return true
	}
	return false
}

// Properties to create a benefit of type `discord`.
type BenefitUpdateParamsBodyBenefitDiscordUpdateProperties struct {
	GuildToken param.Field[string] `json:"guild_token,required"`
	// The ID of the Discord role to grant.
	RoleID param.Field[string] `json:"role_id,required"`
}

func (r BenefitUpdateParamsBodyBenefitDiscordUpdateProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdate struct {
	Type param.Field[BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdateType] `json:"type,required"`
	// The description of the benefit. Will be displayed on products having this
	// benefit.
	Description param.Field[string] `json:"description"`
	// Properties to create a benefit of type `github_repository`.
	Properties param.Field[BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdateProperties] `json:"properties"`
}

func (r BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdate) implementsBenefitUpdateParamsBodyUnion() {
}

type BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdateType string

const (
	BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdateTypeGitHubRepository BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdateType = "github_repository"
)

func (r BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdateType) IsKnown() bool {
	switch r {
	case BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdateTypeGitHubRepository:
		return true
	}
	return false
}

// Properties to create a benefit of type `github_repository`.
type BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdateProperties struct {
	// The permission level to grant. Read more about roles and their permissions on
	// [GitHub documentation](https://docs.github.com/en/organizations/managing-user-access-to-your-organizations-repositories/managing-repository-roles/repository-roles-for-an-organization#permissions-for-each-role).
	Permission   param.Field[BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdatePropertiesPermission] `json:"permission,required"`
	RepositoryID param.Field[string]                                                                   `json:"repository_id" format:"uuid4"`
	// The name of the repository.
	RepositoryName param.Field[string] `json:"repository_name"`
	// The owner of the repository.
	RepositoryOwner param.Field[string] `json:"repository_owner"`
}

func (r BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdateProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The permission level to grant. Read more about roles and their permissions on
// [GitHub documentation](https://docs.github.com/en/organizations/managing-user-access-to-your-organizations-repositories/managing-repository-roles/repository-roles-for-an-organization#permissions-for-each-role).
type BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdatePropertiesPermission string

const (
	BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdatePropertiesPermissionPull     BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdatePropertiesPermission = "pull"
	BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdatePropertiesPermissionTriage   BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdatePropertiesPermission = "triage"
	BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdatePropertiesPermissionPush     BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdatePropertiesPermission = "push"
	BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdatePropertiesPermissionMaintain BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdatePropertiesPermission = "maintain"
	BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdatePropertiesPermissionAdmin    BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdatePropertiesPermission = "admin"
)

func (r BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdatePropertiesPermission) IsKnown() bool {
	switch r {
	case BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdatePropertiesPermissionPull, BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdatePropertiesPermissionTriage, BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdatePropertiesPermissionPush, BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdatePropertiesPermissionMaintain, BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdatePropertiesPermissionAdmin:
		return true
	}
	return false
}

type BenefitUpdateParamsBodyBenefitDownloadablesUpdate struct {
	Type param.Field[BenefitUpdateParamsBodyBenefitDownloadablesUpdateType] `json:"type,required"`
	// The description of the benefit. Will be displayed on products having this
	// benefit.
	Description param.Field[string]                                                      `json:"description"`
	Properties  param.Field[BenefitUpdateParamsBodyBenefitDownloadablesUpdateProperties] `json:"properties"`
}

func (r BenefitUpdateParamsBodyBenefitDownloadablesUpdate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BenefitUpdateParamsBodyBenefitDownloadablesUpdate) implementsBenefitUpdateParamsBodyUnion() {}

type BenefitUpdateParamsBodyBenefitDownloadablesUpdateType string

const (
	BenefitUpdateParamsBodyBenefitDownloadablesUpdateTypeDownloadables BenefitUpdateParamsBodyBenefitDownloadablesUpdateType = "downloadables"
)

func (r BenefitUpdateParamsBodyBenefitDownloadablesUpdateType) IsKnown() bool {
	switch r {
	case BenefitUpdateParamsBodyBenefitDownloadablesUpdateTypeDownloadables:
		return true
	}
	return false
}

type BenefitUpdateParamsBodyBenefitDownloadablesUpdateProperties struct {
	Files    param.Field[[]string]        `json:"files,required" format:"uuid4"`
	Archived param.Field[map[string]bool] `json:"archived"`
}

func (r BenefitUpdateParamsBodyBenefitDownloadablesUpdateProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BenefitUpdateParamsBodyBenefitLicenseKeysUpdate struct {
	Type param.Field[BenefitUpdateParamsBodyBenefitLicenseKeysUpdateType] `json:"type,required"`
	// The description of the benefit. Will be displayed on products having this
	// benefit.
	Description param.Field[string]                                                    `json:"description"`
	Properties  param.Field[BenefitUpdateParamsBodyBenefitLicenseKeysUpdateProperties] `json:"properties"`
}

func (r BenefitUpdateParamsBodyBenefitLicenseKeysUpdate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BenefitUpdateParamsBodyBenefitLicenseKeysUpdate) implementsBenefitUpdateParamsBodyUnion() {}

type BenefitUpdateParamsBodyBenefitLicenseKeysUpdateType string

const (
	BenefitUpdateParamsBodyBenefitLicenseKeysUpdateTypeLicenseKeys BenefitUpdateParamsBodyBenefitLicenseKeysUpdateType = "license_keys"
)

func (r BenefitUpdateParamsBodyBenefitLicenseKeysUpdateType) IsKnown() bool {
	switch r {
	case BenefitUpdateParamsBodyBenefitLicenseKeysUpdateTypeLicenseKeys:
		return true
	}
	return false
}

type BenefitUpdateParamsBodyBenefitLicenseKeysUpdateProperties struct {
	Activations param.Field[BenefitUpdateParamsBodyBenefitLicenseKeysUpdatePropertiesActivations] `json:"activations"`
	Expires     param.Field[BenefitUpdateParamsBodyBenefitLicenseKeysUpdatePropertiesExpires]     `json:"expires"`
	LimitUsage  param.Field[int64]                                                                `json:"limit_usage"`
	Prefix      param.Field[string]                                                               `json:"prefix"`
}

func (r BenefitUpdateParamsBodyBenefitLicenseKeysUpdateProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BenefitUpdateParamsBodyBenefitLicenseKeysUpdatePropertiesActivations struct {
	EnableUserAdmin param.Field[bool]  `json:"enable_user_admin,required"`
	Limit           param.Field[int64] `json:"limit,required"`
}

func (r BenefitUpdateParamsBodyBenefitLicenseKeysUpdatePropertiesActivations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BenefitUpdateParamsBodyBenefitLicenseKeysUpdatePropertiesExpires struct {
	Timeframe param.Field[BenefitUpdateParamsBodyBenefitLicenseKeysUpdatePropertiesExpiresTimeframe] `json:"timeframe,required"`
	Ttl       param.Field[int64]                                                                     `json:"ttl,required"`
}

func (r BenefitUpdateParamsBodyBenefitLicenseKeysUpdatePropertiesExpires) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BenefitUpdateParamsBodyBenefitLicenseKeysUpdatePropertiesExpiresTimeframe string

const (
	BenefitUpdateParamsBodyBenefitLicenseKeysUpdatePropertiesExpiresTimeframeYear  BenefitUpdateParamsBodyBenefitLicenseKeysUpdatePropertiesExpiresTimeframe = "year"
	BenefitUpdateParamsBodyBenefitLicenseKeysUpdatePropertiesExpiresTimeframeMonth BenefitUpdateParamsBodyBenefitLicenseKeysUpdatePropertiesExpiresTimeframe = "month"
	BenefitUpdateParamsBodyBenefitLicenseKeysUpdatePropertiesExpiresTimeframeDay   BenefitUpdateParamsBodyBenefitLicenseKeysUpdatePropertiesExpiresTimeframe = "day"
)

func (r BenefitUpdateParamsBodyBenefitLicenseKeysUpdatePropertiesExpiresTimeframe) IsKnown() bool {
	switch r {
	case BenefitUpdateParamsBodyBenefitLicenseKeysUpdatePropertiesExpiresTimeframeYear, BenefitUpdateParamsBodyBenefitLicenseKeysUpdatePropertiesExpiresTimeframeMonth, BenefitUpdateParamsBodyBenefitLicenseKeysUpdatePropertiesExpiresTimeframeDay:
		return true
	}
	return false
}

type BenefitUpdateParamsBodyType string

const (
	BenefitUpdateParamsBodyTypeArticles         BenefitUpdateParamsBodyType = "articles"
	BenefitUpdateParamsBodyTypeAds              BenefitUpdateParamsBodyType = "ads"
	BenefitUpdateParamsBodyTypeCustom           BenefitUpdateParamsBodyType = "custom"
	BenefitUpdateParamsBodyTypeDiscord          BenefitUpdateParamsBodyType = "discord"
	BenefitUpdateParamsBodyTypeGitHubRepository BenefitUpdateParamsBodyType = "github_repository"
	BenefitUpdateParamsBodyTypeDownloadables    BenefitUpdateParamsBodyType = "downloadables"
	BenefitUpdateParamsBodyTypeLicenseKeys      BenefitUpdateParamsBodyType = "license_keys"
)

func (r BenefitUpdateParamsBodyType) IsKnown() bool {
	switch r {
	case BenefitUpdateParamsBodyTypeArticles, BenefitUpdateParamsBodyTypeAds, BenefitUpdateParamsBodyTypeCustom, BenefitUpdateParamsBodyTypeDiscord, BenefitUpdateParamsBodyTypeGitHubRepository, BenefitUpdateParamsBodyTypeDownloadables, BenefitUpdateParamsBodyTypeLicenseKeys:
		return true
	}
	return false
}

type BenefitListParams struct {
	// Size of a page, defaults to 10. Maximum is 100.
	Limit param.Field[int64] `query:"limit"`
	// Filter by organization ID.
	OrganizationID param.Field[BenefitListParamsOrganizationIDUnion] `query:"organization_id" format:"uuid4"`
	// Page number, defaults to 1.
	Page param.Field[int64] `query:"page"`
	// Filter by benefit type.
	Type param.Field[BenefitListParamsTypeUnion] `query:"type"`
}

// URLQuery serializes [BenefitListParams]'s query parameters as `url.Values`.
func (r BenefitListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by organization ID.
//
// Satisfied by [shared.UnionString], [BenefitListParamsOrganizationIDArray].
type BenefitListParamsOrganizationIDUnion interface {
	ImplementsBenefitListParamsOrganizationIDUnion()
}

type BenefitListParamsOrganizationIDArray []string

func (r BenefitListParamsOrganizationIDArray) ImplementsBenefitListParamsOrganizationIDUnion() {}

// Filter by benefit type.
//
// Satisfied by [BenefitListParamsTypeBenefitType], [BenefitListParamsTypeArray].
type BenefitListParamsTypeUnion interface {
	implementsBenefitListParamsTypeUnion()
}

type BenefitListParamsTypeBenefitType string

const (
	BenefitListParamsTypeBenefitTypeCustom           BenefitListParamsTypeBenefitType = "custom"
	BenefitListParamsTypeBenefitTypeArticles         BenefitListParamsTypeBenefitType = "articles"
	BenefitListParamsTypeBenefitTypeAds              BenefitListParamsTypeBenefitType = "ads"
	BenefitListParamsTypeBenefitTypeDiscord          BenefitListParamsTypeBenefitType = "discord"
	BenefitListParamsTypeBenefitTypeGitHubRepository BenefitListParamsTypeBenefitType = "github_repository"
	BenefitListParamsTypeBenefitTypeDownloadables    BenefitListParamsTypeBenefitType = "downloadables"
	BenefitListParamsTypeBenefitTypeLicenseKeys      BenefitListParamsTypeBenefitType = "license_keys"
)

func (r BenefitListParamsTypeBenefitType) IsKnown() bool {
	switch r {
	case BenefitListParamsTypeBenefitTypeCustom, BenefitListParamsTypeBenefitTypeArticles, BenefitListParamsTypeBenefitTypeAds, BenefitListParamsTypeBenefitTypeDiscord, BenefitListParamsTypeBenefitTypeGitHubRepository, BenefitListParamsTypeBenefitTypeDownloadables, BenefitListParamsTypeBenefitTypeLicenseKeys:
		return true
	}
	return false
}

func (r BenefitListParamsTypeBenefitType) implementsBenefitListParamsTypeUnion() {}

type BenefitListParamsTypeArray []BenefitListParamsTypeArray

func (r BenefitListParamsTypeArray) implementsBenefitListParamsTypeUnion() {}
