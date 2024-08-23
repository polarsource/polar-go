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
	"github.com/polarsource/polar-go/internal/param"
	"github.com/polarsource/polar-go/internal/requestconfig"
	"github.com/polarsource/polar-go/option"
)

// OrganizationService contains methods and other services that help with
// interacting with the polar API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationService] method instead.
type OrganizationService struct {
	Options   []option.RequestOption
	Customers *OrganizationCustomerService
}

// NewOrganizationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOrganizationService(opts ...option.RequestOption) (r *OrganizationService) {
	r = &OrganizationService{}
	r.Options = opts
	r.Customers = NewOrganizationCustomerService(opts...)
	return
}

// Create an organization.
func (r *OrganizationService) New(ctx context.Context, body OrganizationNewParams, opts ...option.RequestOption) (res *OrganizationNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/organizations/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get an organization by ID.
func (r *OrganizationService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *OrganizationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/organizations/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an organization.
func (r *OrganizationService) Update(ctx context.Context, id string, body OrganizationUpdateParams, opts ...option.RequestOption) (res *OrganizationUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/organizations/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List organizations.
func (r *OrganizationService) List(ctx context.Context, query OrganizationListParams, opts ...option.RequestOption) (res *OrganizationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/organizations/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type OrganizationNewResponse struct {
	// The organization ID.
	ID        string `json:"id,required" format:"uuid4"`
	AvatarURL string `json:"avatar_url,required,nullable"`
	Bio       string `json:"bio,required,nullable"`
	Blog      string `json:"blog,required,nullable"`
	Company   string `json:"company,required,nullable"`
	// Creation timestamp of the object.
	CreatedAt                         time.Time `json:"created_at,required" format:"date-time"`
	DefaultUpfrontSplitToContributors int64     `json:"default_upfront_split_to_contributors,required,nullable"`
	// If this organizations accepts donations
	DonationsEnabled bool   `json:"donations_enabled,required"`
	Email            string `json:"email,required,nullable"`
	// Settings for the organization features
	FeatureSettings OrganizationNewResponseFeatureSettings `json:"feature_settings,required,nullable"`
	Location        string                                 `json:"location,required,nullable"`
	// Last modification timestamp of the object.
	ModifiedAt            time.Time `json:"modified_at,required,nullable" format:"date-time"`
	Name                  string    `json:"name,required"`
	PledgeBadgeShowAmount bool      `json:"pledge_badge_show_amount,required"`
	PledgeMinimumAmount   int64     `json:"pledge_minimum_amount,required"`
	// Settings for the organization profile
	ProfileSettings OrganizationNewResponseProfileSettings `json:"profile_settings,required,nullable"`
	Slug            string                                 `json:"slug,required"`
	TwitterUsername string                                 `json:"twitter_username,required,nullable"`
	JSON            organizationNewResponseJSON            `json:"-"`
}

// organizationNewResponseJSON contains the JSON metadata for the struct
// [OrganizationNewResponse]
type organizationNewResponseJSON struct {
	ID                                apijson.Field
	AvatarURL                         apijson.Field
	Bio                               apijson.Field
	Blog                              apijson.Field
	Company                           apijson.Field
	CreatedAt                         apijson.Field
	DefaultUpfrontSplitToContributors apijson.Field
	DonationsEnabled                  apijson.Field
	Email                             apijson.Field
	FeatureSettings                   apijson.Field
	Location                          apijson.Field
	ModifiedAt                        apijson.Field
	Name                              apijson.Field
	PledgeBadgeShowAmount             apijson.Field
	PledgeMinimumAmount               apijson.Field
	ProfileSettings                   apijson.Field
	Slug                              apijson.Field
	TwitterUsername                   apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
}

func (r *OrganizationNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationNewResponseJSON) RawJSON() string {
	return r.raw
}

// Settings for the organization features
type OrganizationNewResponseFeatureSettings struct {
	// If this organization has articles enabled
	ArticlesEnabled bool `json:"articles_enabled"`
	// If this organization has issue funding enabled
	IssueFundingEnabled bool `json:"issue_funding_enabled"`
	// If this organization has subscriptions enabled
	SubscriptionsEnabled bool                                       `json:"subscriptions_enabled"`
	JSON                 organizationNewResponseFeatureSettingsJSON `json:"-"`
}

// organizationNewResponseFeatureSettingsJSON contains the JSON metadata for the
// struct [OrganizationNewResponseFeatureSettings]
type organizationNewResponseFeatureSettingsJSON struct {
	ArticlesEnabled      apijson.Field
	IssueFundingEnabled  apijson.Field
	SubscriptionsEnabled apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *OrganizationNewResponseFeatureSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationNewResponseFeatureSettingsJSON) RawJSON() string {
	return r.raw
}

// Settings for the organization profile
type OrganizationNewResponseProfileSettings struct {
	// A description of the organization
	Description string `json:"description,nullable"`
	// A list of featured organizations
	FeaturedOrganizations []string `json:"featured_organizations,nullable" format:"uuid4"`
	// A list of featured projects
	FeaturedProjects []string `json:"featured_projects,nullable" format:"uuid4"`
	// A list of links associated with the organization
	Links []string `json:"links,nullable" format:"uri"`
	// Subscription promotion settings
	Subscribe OrganizationNewResponseProfileSettingsSubscribe `json:"subscribe,nullable"`
	JSON      organizationNewResponseProfileSettingsJSON      `json:"-"`
}

// organizationNewResponseProfileSettingsJSON contains the JSON metadata for the
// struct [OrganizationNewResponseProfileSettings]
type organizationNewResponseProfileSettingsJSON struct {
	Description           apijson.Field
	FeaturedOrganizations apijson.Field
	FeaturedProjects      apijson.Field
	Links                 apijson.Field
	Subscribe             apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *OrganizationNewResponseProfileSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationNewResponseProfileSettingsJSON) RawJSON() string {
	return r.raw
}

// Subscription promotion settings
type OrganizationNewResponseProfileSettingsSubscribe struct {
	// Include free subscribers in total count
	CountFree bool `json:"count_free"`
	// Promote email subscription (free)
	Promote bool `json:"promote"`
	// Show subscription count publicly
	ShowCount bool                                                `json:"show_count"`
	JSON      organizationNewResponseProfileSettingsSubscribeJSON `json:"-"`
}

// organizationNewResponseProfileSettingsSubscribeJSON contains the JSON metadata
// for the struct [OrganizationNewResponseProfileSettingsSubscribe]
type organizationNewResponseProfileSettingsSubscribeJSON struct {
	CountFree   apijson.Field
	Promote     apijson.Field
	ShowCount   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationNewResponseProfileSettingsSubscribe) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationNewResponseProfileSettingsSubscribeJSON) RawJSON() string {
	return r.raw
}

type OrganizationGetResponse struct {
	// The organization ID.
	ID        string `json:"id,required" format:"uuid4"`
	AvatarURL string `json:"avatar_url,required,nullable"`
	Bio       string `json:"bio,required,nullable"`
	Blog      string `json:"blog,required,nullable"`
	Company   string `json:"company,required,nullable"`
	// Creation timestamp of the object.
	CreatedAt                         time.Time `json:"created_at,required" format:"date-time"`
	DefaultUpfrontSplitToContributors int64     `json:"default_upfront_split_to_contributors,required,nullable"`
	// If this organizations accepts donations
	DonationsEnabled bool   `json:"donations_enabled,required"`
	Email            string `json:"email,required,nullable"`
	// Settings for the organization features
	FeatureSettings OrganizationGetResponseFeatureSettings `json:"feature_settings,required,nullable"`
	Location        string                                 `json:"location,required,nullable"`
	// Last modification timestamp of the object.
	ModifiedAt            time.Time `json:"modified_at,required,nullable" format:"date-time"`
	Name                  string    `json:"name,required"`
	PledgeBadgeShowAmount bool      `json:"pledge_badge_show_amount,required"`
	PledgeMinimumAmount   int64     `json:"pledge_minimum_amount,required"`
	// Settings for the organization profile
	ProfileSettings OrganizationGetResponseProfileSettings `json:"profile_settings,required,nullable"`
	Slug            string                                 `json:"slug,required"`
	TwitterUsername string                                 `json:"twitter_username,required,nullable"`
	JSON            organizationGetResponseJSON            `json:"-"`
}

// organizationGetResponseJSON contains the JSON metadata for the struct
// [OrganizationGetResponse]
type organizationGetResponseJSON struct {
	ID                                apijson.Field
	AvatarURL                         apijson.Field
	Bio                               apijson.Field
	Blog                              apijson.Field
	Company                           apijson.Field
	CreatedAt                         apijson.Field
	DefaultUpfrontSplitToContributors apijson.Field
	DonationsEnabled                  apijson.Field
	Email                             apijson.Field
	FeatureSettings                   apijson.Field
	Location                          apijson.Field
	ModifiedAt                        apijson.Field
	Name                              apijson.Field
	PledgeBadgeShowAmount             apijson.Field
	PledgeMinimumAmount               apijson.Field
	ProfileSettings                   apijson.Field
	Slug                              apijson.Field
	TwitterUsername                   apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
}

func (r *OrganizationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationGetResponseJSON) RawJSON() string {
	return r.raw
}

// Settings for the organization features
type OrganizationGetResponseFeatureSettings struct {
	// If this organization has articles enabled
	ArticlesEnabled bool `json:"articles_enabled"`
	// If this organization has issue funding enabled
	IssueFundingEnabled bool `json:"issue_funding_enabled"`
	// If this organization has subscriptions enabled
	SubscriptionsEnabled bool                                       `json:"subscriptions_enabled"`
	JSON                 organizationGetResponseFeatureSettingsJSON `json:"-"`
}

// organizationGetResponseFeatureSettingsJSON contains the JSON metadata for the
// struct [OrganizationGetResponseFeatureSettings]
type organizationGetResponseFeatureSettingsJSON struct {
	ArticlesEnabled      apijson.Field
	IssueFundingEnabled  apijson.Field
	SubscriptionsEnabled apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *OrganizationGetResponseFeatureSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationGetResponseFeatureSettingsJSON) RawJSON() string {
	return r.raw
}

// Settings for the organization profile
type OrganizationGetResponseProfileSettings struct {
	// A description of the organization
	Description string `json:"description,nullable"`
	// A list of featured organizations
	FeaturedOrganizations []string `json:"featured_organizations,nullable" format:"uuid4"`
	// A list of featured projects
	FeaturedProjects []string `json:"featured_projects,nullable" format:"uuid4"`
	// A list of links associated with the organization
	Links []string `json:"links,nullable" format:"uri"`
	// Subscription promotion settings
	Subscribe OrganizationGetResponseProfileSettingsSubscribe `json:"subscribe,nullable"`
	JSON      organizationGetResponseProfileSettingsJSON      `json:"-"`
}

// organizationGetResponseProfileSettingsJSON contains the JSON metadata for the
// struct [OrganizationGetResponseProfileSettings]
type organizationGetResponseProfileSettingsJSON struct {
	Description           apijson.Field
	FeaturedOrganizations apijson.Field
	FeaturedProjects      apijson.Field
	Links                 apijson.Field
	Subscribe             apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *OrganizationGetResponseProfileSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationGetResponseProfileSettingsJSON) RawJSON() string {
	return r.raw
}

// Subscription promotion settings
type OrganizationGetResponseProfileSettingsSubscribe struct {
	// Include free subscribers in total count
	CountFree bool `json:"count_free"`
	// Promote email subscription (free)
	Promote bool `json:"promote"`
	// Show subscription count publicly
	ShowCount bool                                                `json:"show_count"`
	JSON      organizationGetResponseProfileSettingsSubscribeJSON `json:"-"`
}

// organizationGetResponseProfileSettingsSubscribeJSON contains the JSON metadata
// for the struct [OrganizationGetResponseProfileSettingsSubscribe]
type organizationGetResponseProfileSettingsSubscribeJSON struct {
	CountFree   apijson.Field
	Promote     apijson.Field
	ShowCount   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationGetResponseProfileSettingsSubscribe) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationGetResponseProfileSettingsSubscribeJSON) RawJSON() string {
	return r.raw
}

type OrganizationUpdateResponse struct {
	// The organization ID.
	ID        string `json:"id,required" format:"uuid4"`
	AvatarURL string `json:"avatar_url,required,nullable"`
	Bio       string `json:"bio,required,nullable"`
	Blog      string `json:"blog,required,nullable"`
	Company   string `json:"company,required,nullable"`
	// Creation timestamp of the object.
	CreatedAt                         time.Time `json:"created_at,required" format:"date-time"`
	DefaultUpfrontSplitToContributors int64     `json:"default_upfront_split_to_contributors,required,nullable"`
	// If this organizations accepts donations
	DonationsEnabled bool   `json:"donations_enabled,required"`
	Email            string `json:"email,required,nullable"`
	// Settings for the organization features
	FeatureSettings OrganizationUpdateResponseFeatureSettings `json:"feature_settings,required,nullable"`
	Location        string                                    `json:"location,required,nullable"`
	// Last modification timestamp of the object.
	ModifiedAt            time.Time `json:"modified_at,required,nullable" format:"date-time"`
	Name                  string    `json:"name,required"`
	PledgeBadgeShowAmount bool      `json:"pledge_badge_show_amount,required"`
	PledgeMinimumAmount   int64     `json:"pledge_minimum_amount,required"`
	// Settings for the organization profile
	ProfileSettings OrganizationUpdateResponseProfileSettings `json:"profile_settings,required,nullable"`
	Slug            string                                    `json:"slug,required"`
	TwitterUsername string                                    `json:"twitter_username,required,nullable"`
	JSON            organizationUpdateResponseJSON            `json:"-"`
}

// organizationUpdateResponseJSON contains the JSON metadata for the struct
// [OrganizationUpdateResponse]
type organizationUpdateResponseJSON struct {
	ID                                apijson.Field
	AvatarURL                         apijson.Field
	Bio                               apijson.Field
	Blog                              apijson.Field
	Company                           apijson.Field
	CreatedAt                         apijson.Field
	DefaultUpfrontSplitToContributors apijson.Field
	DonationsEnabled                  apijson.Field
	Email                             apijson.Field
	FeatureSettings                   apijson.Field
	Location                          apijson.Field
	ModifiedAt                        apijson.Field
	Name                              apijson.Field
	PledgeBadgeShowAmount             apijson.Field
	PledgeMinimumAmount               apijson.Field
	ProfileSettings                   apijson.Field
	Slug                              apijson.Field
	TwitterUsername                   apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
}

func (r *OrganizationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Settings for the organization features
type OrganizationUpdateResponseFeatureSettings struct {
	// If this organization has articles enabled
	ArticlesEnabled bool `json:"articles_enabled"`
	// If this organization has issue funding enabled
	IssueFundingEnabled bool `json:"issue_funding_enabled"`
	// If this organization has subscriptions enabled
	SubscriptionsEnabled bool                                          `json:"subscriptions_enabled"`
	JSON                 organizationUpdateResponseFeatureSettingsJSON `json:"-"`
}

// organizationUpdateResponseFeatureSettingsJSON contains the JSON metadata for the
// struct [OrganizationUpdateResponseFeatureSettings]
type organizationUpdateResponseFeatureSettingsJSON struct {
	ArticlesEnabled      apijson.Field
	IssueFundingEnabled  apijson.Field
	SubscriptionsEnabled apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *OrganizationUpdateResponseFeatureSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationUpdateResponseFeatureSettingsJSON) RawJSON() string {
	return r.raw
}

// Settings for the organization profile
type OrganizationUpdateResponseProfileSettings struct {
	// A description of the organization
	Description string `json:"description,nullable"`
	// A list of featured organizations
	FeaturedOrganizations []string `json:"featured_organizations,nullable" format:"uuid4"`
	// A list of featured projects
	FeaturedProjects []string `json:"featured_projects,nullable" format:"uuid4"`
	// A list of links associated with the organization
	Links []string `json:"links,nullable" format:"uri"`
	// Subscription promotion settings
	Subscribe OrganizationUpdateResponseProfileSettingsSubscribe `json:"subscribe,nullable"`
	JSON      organizationUpdateResponseProfileSettingsJSON      `json:"-"`
}

// organizationUpdateResponseProfileSettingsJSON contains the JSON metadata for the
// struct [OrganizationUpdateResponseProfileSettings]
type organizationUpdateResponseProfileSettingsJSON struct {
	Description           apijson.Field
	FeaturedOrganizations apijson.Field
	FeaturedProjects      apijson.Field
	Links                 apijson.Field
	Subscribe             apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *OrganizationUpdateResponseProfileSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationUpdateResponseProfileSettingsJSON) RawJSON() string {
	return r.raw
}

// Subscription promotion settings
type OrganizationUpdateResponseProfileSettingsSubscribe struct {
	// Include free subscribers in total count
	CountFree bool `json:"count_free"`
	// Promote email subscription (free)
	Promote bool `json:"promote"`
	// Show subscription count publicly
	ShowCount bool                                                   `json:"show_count"`
	JSON      organizationUpdateResponseProfileSettingsSubscribeJSON `json:"-"`
}

// organizationUpdateResponseProfileSettingsSubscribeJSON contains the JSON
// metadata for the struct [OrganizationUpdateResponseProfileSettingsSubscribe]
type organizationUpdateResponseProfileSettingsSubscribeJSON struct {
	CountFree   apijson.Field
	Promote     apijson.Field
	ShowCount   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationUpdateResponseProfileSettingsSubscribe) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationUpdateResponseProfileSettingsSubscribeJSON) RawJSON() string {
	return r.raw
}

type OrganizationListResponse struct {
	Pagination OrganizationListResponsePagination `json:"pagination,required"`
	Items      []OrganizationListResponseItem     `json:"items"`
	JSON       organizationListResponseJSON       `json:"-"`
}

// organizationListResponseJSON contains the JSON metadata for the struct
// [OrganizationListResponse]
type organizationListResponseJSON struct {
	Pagination  apijson.Field
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationListResponseJSON) RawJSON() string {
	return r.raw
}

type OrganizationListResponsePagination struct {
	MaxPage    int64                                  `json:"max_page,required"`
	TotalCount int64                                  `json:"total_count,required"`
	JSON       organizationListResponsePaginationJSON `json:"-"`
}

// organizationListResponsePaginationJSON contains the JSON metadata for the struct
// [OrganizationListResponsePagination]
type organizationListResponsePaginationJSON struct {
	MaxPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationListResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationListResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type OrganizationListResponseItem struct {
	// The organization ID.
	ID        string `json:"id,required" format:"uuid4"`
	AvatarURL string `json:"avatar_url,required,nullable"`
	Bio       string `json:"bio,required,nullable"`
	Blog      string `json:"blog,required,nullable"`
	Company   string `json:"company,required,nullable"`
	// Creation timestamp of the object.
	CreatedAt                         time.Time `json:"created_at,required" format:"date-time"`
	DefaultUpfrontSplitToContributors int64     `json:"default_upfront_split_to_contributors,required,nullable"`
	// If this organizations accepts donations
	DonationsEnabled bool   `json:"donations_enabled,required"`
	Email            string `json:"email,required,nullable"`
	// Settings for the organization features
	FeatureSettings OrganizationListResponseItemsFeatureSettings `json:"feature_settings,required,nullable"`
	Location        string                                       `json:"location,required,nullable"`
	// Last modification timestamp of the object.
	ModifiedAt            time.Time `json:"modified_at,required,nullable" format:"date-time"`
	Name                  string    `json:"name,required"`
	PledgeBadgeShowAmount bool      `json:"pledge_badge_show_amount,required"`
	PledgeMinimumAmount   int64     `json:"pledge_minimum_amount,required"`
	// Settings for the organization profile
	ProfileSettings OrganizationListResponseItemsProfileSettings `json:"profile_settings,required,nullable"`
	Slug            string                                       `json:"slug,required"`
	TwitterUsername string                                       `json:"twitter_username,required,nullable"`
	JSON            organizationListResponseItemJSON             `json:"-"`
}

// organizationListResponseItemJSON contains the JSON metadata for the struct
// [OrganizationListResponseItem]
type organizationListResponseItemJSON struct {
	ID                                apijson.Field
	AvatarURL                         apijson.Field
	Bio                               apijson.Field
	Blog                              apijson.Field
	Company                           apijson.Field
	CreatedAt                         apijson.Field
	DefaultUpfrontSplitToContributors apijson.Field
	DonationsEnabled                  apijson.Field
	Email                             apijson.Field
	FeatureSettings                   apijson.Field
	Location                          apijson.Field
	ModifiedAt                        apijson.Field
	Name                              apijson.Field
	PledgeBadgeShowAmount             apijson.Field
	PledgeMinimumAmount               apijson.Field
	ProfileSettings                   apijson.Field
	Slug                              apijson.Field
	TwitterUsername                   apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
}

func (r *OrganizationListResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationListResponseItemJSON) RawJSON() string {
	return r.raw
}

// Settings for the organization features
type OrganizationListResponseItemsFeatureSettings struct {
	// If this organization has articles enabled
	ArticlesEnabled bool `json:"articles_enabled"`
	// If this organization has issue funding enabled
	IssueFundingEnabled bool `json:"issue_funding_enabled"`
	// If this organization has subscriptions enabled
	SubscriptionsEnabled bool                                             `json:"subscriptions_enabled"`
	JSON                 organizationListResponseItemsFeatureSettingsJSON `json:"-"`
}

// organizationListResponseItemsFeatureSettingsJSON contains the JSON metadata for
// the struct [OrganizationListResponseItemsFeatureSettings]
type organizationListResponseItemsFeatureSettingsJSON struct {
	ArticlesEnabled      apijson.Field
	IssueFundingEnabled  apijson.Field
	SubscriptionsEnabled apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *OrganizationListResponseItemsFeatureSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationListResponseItemsFeatureSettingsJSON) RawJSON() string {
	return r.raw
}

// Settings for the organization profile
type OrganizationListResponseItemsProfileSettings struct {
	// A description of the organization
	Description string `json:"description,nullable"`
	// A list of featured organizations
	FeaturedOrganizations []string `json:"featured_organizations,nullable" format:"uuid4"`
	// A list of featured projects
	FeaturedProjects []string `json:"featured_projects,nullable" format:"uuid4"`
	// A list of links associated with the organization
	Links []string `json:"links,nullable" format:"uri"`
	// Subscription promotion settings
	Subscribe OrganizationListResponseItemsProfileSettingsSubscribe `json:"subscribe,nullable"`
	JSON      organizationListResponseItemsProfileSettingsJSON      `json:"-"`
}

// organizationListResponseItemsProfileSettingsJSON contains the JSON metadata for
// the struct [OrganizationListResponseItemsProfileSettings]
type organizationListResponseItemsProfileSettingsJSON struct {
	Description           apijson.Field
	FeaturedOrganizations apijson.Field
	FeaturedProjects      apijson.Field
	Links                 apijson.Field
	Subscribe             apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *OrganizationListResponseItemsProfileSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationListResponseItemsProfileSettingsJSON) RawJSON() string {
	return r.raw
}

// Subscription promotion settings
type OrganizationListResponseItemsProfileSettingsSubscribe struct {
	// Include free subscribers in total count
	CountFree bool `json:"count_free"`
	// Promote email subscription (free)
	Promote bool `json:"promote"`
	// Show subscription count publicly
	ShowCount bool                                                      `json:"show_count"`
	JSON      organizationListResponseItemsProfileSettingsSubscribeJSON `json:"-"`
}

// organizationListResponseItemsProfileSettingsSubscribeJSON contains the JSON
// metadata for the struct [OrganizationListResponseItemsProfileSettingsSubscribe]
type organizationListResponseItemsProfileSettingsSubscribeJSON struct {
	CountFree   apijson.Field
	Promote     apijson.Field
	ShowCount   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationListResponseItemsProfileSettingsSubscribe) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationListResponseItemsProfileSettingsSubscribeJSON) RawJSON() string {
	return r.raw
}

type OrganizationNewParams struct {
	Name             param.Field[string]                               `json:"name,required"`
	Slug             param.Field[string]                               `json:"slug,required"`
	AvatarURL        param.Field[string]                               `json:"avatar_url" format:"uri"`
	DonationsEnabled param.Field[bool]                                 `json:"donations_enabled"`
	FeatureSettings  param.Field[OrganizationNewParamsFeatureSettings] `json:"feature_settings"`
}

func (r OrganizationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationNewParamsFeatureSettings struct {
	// If this organization has articles enabled
	ArticlesEnabled param.Field[bool] `json:"articles_enabled"`
	// If this organization has issue funding enabled
	IssueFundingEnabled param.Field[bool] `json:"issue_funding_enabled"`
	// If this organization has subscriptions enabled
	SubscriptionsEnabled param.Field[bool] `json:"subscriptions_enabled"`
}

func (r OrganizationNewParamsFeatureSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationUpdateParams struct {
	AvatarURL                         param.Field[string]                                  `json:"avatar_url" format:"uri"`
	BillingEmail                      param.Field[string]                                  `json:"billing_email"`
	DefaultBadgeCustomContent         param.Field[string]                                  `json:"default_badge_custom_content"`
	DefaultUpfrontSplitToContributors param.Field[int64]                                   `json:"default_upfront_split_to_contributors"`
	DonationsEnabled                  param.Field[bool]                                    `json:"donations_enabled"`
	FeatureSettings                   param.Field[OrganizationUpdateParamsFeatureSettings] `json:"feature_settings"`
	Name                              param.Field[string]                                  `json:"name"`
	PerUserMonthlySpendingLimit       param.Field[int64]                                   `json:"per_user_monthly_spending_limit"`
	PledgeBadgeShowAmount             param.Field[bool]                                    `json:"pledge_badge_show_amount"`
	PledgeMinimumAmount               param.Field[int64]                                   `json:"pledge_minimum_amount"`
	ProfileSettings                   param.Field[OrganizationUpdateParamsProfileSettings] `json:"profile_settings"`
	TotalMonthlySpendingLimit         param.Field[int64]                                   `json:"total_monthly_spending_limit"`
}

func (r OrganizationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationUpdateParamsFeatureSettings struct {
	// If this organization has articles enabled
	ArticlesEnabled param.Field[bool] `json:"articles_enabled"`
	// If this organization has issue funding enabled
	IssueFundingEnabled param.Field[bool] `json:"issue_funding_enabled"`
	// If this organization has subscriptions enabled
	SubscriptionsEnabled param.Field[bool] `json:"subscriptions_enabled"`
}

func (r OrganizationUpdateParamsFeatureSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationUpdateParamsProfileSettings struct {
	// A description of the organization
	Description param.Field[string] `json:"description"`
	// A list of featured organizations
	FeaturedOrganizations param.Field[[]string] `json:"featured_organizations" format:"uuid4"`
	// A list of featured projects
	FeaturedProjects param.Field[[]string] `json:"featured_projects" format:"uuid4"`
	// A list of links associated with the organization
	Links param.Field[[]string] `json:"links" format:"uri"`
	// Subscription promotion settings
	Subscribe param.Field[OrganizationUpdateParamsProfileSettingsSubscribe] `json:"subscribe"`
}

func (r OrganizationUpdateParamsProfileSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Subscription promotion settings
type OrganizationUpdateParamsProfileSettingsSubscribe struct {
	// Include free subscribers in total count
	CountFree param.Field[bool] `json:"count_free"`
	// Promote email subscription (free)
	Promote param.Field[bool] `json:"promote"`
	// Show subscription count publicly
	ShowCount param.Field[bool] `json:"show_count"`
}

func (r OrganizationUpdateParamsProfileSettingsSubscribe) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationListParams struct {
	// Filter by membership. If `true`, only organizations the user is a member of are
	// returned. If `false`, only organizations the user is not a member of are
	// returned.
	IsMember param.Field[bool] `query:"is_member"`
	// Size of a page, defaults to 10. Maximum is 100.
	Limit param.Field[int64] `query:"limit"`
	// Page number, defaults to 1.
	Page param.Field[int64] `query:"page"`
	// Filter by slug.
	Slug param.Field[string] `query:"slug"`
	// Sorting criterion. Several criteria can be used simultaneously and will be
	// applied in order. Add a minus sign `-` before the criteria name to sort by
	// descending order.
	Sorting param.Field[[]string] `query:"sorting"`
}

// URLQuery serializes [OrganizationListParams]'s query parameters as `url.Values`.
func (r OrganizationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
