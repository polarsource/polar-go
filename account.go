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

// AccountService contains methods and other services that help with interacting
// with the polar API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountService] method instead.
type AccountService struct {
	Options []option.RequestOption
}

// NewAccountService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAccountService(opts ...option.RequestOption) (r *AccountService) {
	r = &AccountService{}
	r.Options = opts
	return
}

// Create
func (r *AccountService) New(ctx context.Context, body AccountNewParams, opts ...option.RequestOption) (res *AccountNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get
func (r *AccountService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *AccountGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/accounts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Dashboard Link
func (r *AccountService) DashboardLink(ctx context.Context, id string, opts ...option.RequestOption) (res *AccountDashboardLinkResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/accounts/%s/dashboard_link", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Onboarding Link
func (r *AccountService) OnboardingLink(ctx context.Context, id string, body AccountOnboardingLinkParams, opts ...option.RequestOption) (res *AccountOnboardingLinkResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/accounts/%s/onboarding_link", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Search
func (r *AccountService) Search(ctx context.Context, query AccountSearchParams, opts ...option.RequestOption) (res *AccountSearchResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/accounts/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountNewResponse struct {
	ID                 string                           `json:"id,required" format:"uuid"`
	AccountType        AccountNewResponseAccountType    `json:"account_type,required"`
	Country            string                           `json:"country,required"`
	IsChargesEnabled   bool                             `json:"is_charges_enabled,required"`
	IsDetailsSubmitted bool                             `json:"is_details_submitted,required"`
	IsPayoutsEnabled   bool                             `json:"is_payouts_enabled,required"`
	OpenCollectiveSlug string                           `json:"open_collective_slug,required,nullable"`
	Organizations      []AccountNewResponseOrganization `json:"organizations,required"`
	Status             AccountNewResponseStatus         `json:"status,required"`
	StripeID           string                           `json:"stripe_id,required,nullable"`
	Users              []AccountNewResponseUser         `json:"users,required"`
	JSON               accountNewResponseJSON           `json:"-"`
}

// accountNewResponseJSON contains the JSON metadata for the struct
// [AccountNewResponse]
type accountNewResponseJSON struct {
	ID                 apijson.Field
	AccountType        apijson.Field
	Country            apijson.Field
	IsChargesEnabled   apijson.Field
	IsDetailsSubmitted apijson.Field
	IsPayoutsEnabled   apijson.Field
	OpenCollectiveSlug apijson.Field
	Organizations      apijson.Field
	Status             apijson.Field
	StripeID           apijson.Field
	Users              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountNewResponseAccountType string

const (
	AccountNewResponseAccountTypeStripe         AccountNewResponseAccountType = "stripe"
	AccountNewResponseAccountTypeOpenCollective AccountNewResponseAccountType = "open_collective"
)

func (r AccountNewResponseAccountType) IsKnown() bool {
	switch r {
	case AccountNewResponseAccountTypeStripe, AccountNewResponseAccountTypeOpenCollective:
		return true
	}
	return false
}

type AccountNewResponseOrganization struct {
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
	FeatureSettings AccountNewResponseOrganizationsFeatureSettings `json:"feature_settings,required,nullable"`
	Location        string                                         `json:"location,required,nullable"`
	// Last modification timestamp of the object.
	ModifiedAt            time.Time `json:"modified_at,required,nullable" format:"date-time"`
	Name                  string    `json:"name,required"`
	PledgeBadgeShowAmount bool      `json:"pledge_badge_show_amount,required"`
	PledgeMinimumAmount   int64     `json:"pledge_minimum_amount,required"`
	// Settings for the organization profile
	ProfileSettings AccountNewResponseOrganizationsProfileSettings `json:"profile_settings,required,nullable"`
	Slug            string                                         `json:"slug,required"`
	TwitterUsername string                                         `json:"twitter_username,required,nullable"`
	JSON            accountNewResponseOrganizationJSON             `json:"-"`
}

// accountNewResponseOrganizationJSON contains the JSON metadata for the struct
// [AccountNewResponseOrganization]
type accountNewResponseOrganizationJSON struct {
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

func (r *AccountNewResponseOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountNewResponseOrganizationJSON) RawJSON() string {
	return r.raw
}

// Settings for the organization features
type AccountNewResponseOrganizationsFeatureSettings struct {
	// If this organization has articles enabled
	ArticlesEnabled bool `json:"articles_enabled"`
	// If this organization has issue funding enabled
	IssueFundingEnabled bool `json:"issue_funding_enabled"`
	// If this organization has subscriptions enabled
	SubscriptionsEnabled bool                                               `json:"subscriptions_enabled"`
	JSON                 accountNewResponseOrganizationsFeatureSettingsJSON `json:"-"`
}

// accountNewResponseOrganizationsFeatureSettingsJSON contains the JSON metadata
// for the struct [AccountNewResponseOrganizationsFeatureSettings]
type accountNewResponseOrganizationsFeatureSettingsJSON struct {
	ArticlesEnabled      apijson.Field
	IssueFundingEnabled  apijson.Field
	SubscriptionsEnabled apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountNewResponseOrganizationsFeatureSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountNewResponseOrganizationsFeatureSettingsJSON) RawJSON() string {
	return r.raw
}

// Settings for the organization profile
type AccountNewResponseOrganizationsProfileSettings struct {
	// A description of the organization
	Description string `json:"description,nullable"`
	// A list of featured organizations
	FeaturedOrganizations []string `json:"featured_organizations,nullable" format:"uuid4"`
	// A list of featured projects
	FeaturedProjects []string `json:"featured_projects,nullable" format:"uuid4"`
	// A list of links associated with the organization
	Links []string `json:"links,nullable" format:"uri"`
	// Subscription promotion settings
	Subscribe AccountNewResponseOrganizationsProfileSettingsSubscribe `json:"subscribe,nullable"`
	JSON      accountNewResponseOrganizationsProfileSettingsJSON      `json:"-"`
}

// accountNewResponseOrganizationsProfileSettingsJSON contains the JSON metadata
// for the struct [AccountNewResponseOrganizationsProfileSettings]
type accountNewResponseOrganizationsProfileSettingsJSON struct {
	Description           apijson.Field
	FeaturedOrganizations apijson.Field
	FeaturedProjects      apijson.Field
	Links                 apijson.Field
	Subscribe             apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AccountNewResponseOrganizationsProfileSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountNewResponseOrganizationsProfileSettingsJSON) RawJSON() string {
	return r.raw
}

// Subscription promotion settings
type AccountNewResponseOrganizationsProfileSettingsSubscribe struct {
	// Include free subscribers in total count
	CountFree bool `json:"count_free"`
	// Promote email subscription (free)
	Promote bool `json:"promote"`
	// Show subscription count publicly
	ShowCount bool                                                        `json:"show_count"`
	JSON      accountNewResponseOrganizationsProfileSettingsSubscribeJSON `json:"-"`
}

// accountNewResponseOrganizationsProfileSettingsSubscribeJSON contains the JSON
// metadata for the struct
// [AccountNewResponseOrganizationsProfileSettingsSubscribe]
type accountNewResponseOrganizationsProfileSettingsSubscribeJSON struct {
	CountFree   apijson.Field
	Promote     apijson.Field
	ShowCount   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountNewResponseOrganizationsProfileSettingsSubscribe) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountNewResponseOrganizationsProfileSettingsSubscribeJSON) RawJSON() string {
	return r.raw
}

type AccountNewResponseStatus string

const (
	AccountNewResponseStatusCreated           AccountNewResponseStatus = "created"
	AccountNewResponseStatusOnboardingStarted AccountNewResponseStatus = "onboarding_started"
	AccountNewResponseStatusUnderReview       AccountNewResponseStatus = "under_review"
	AccountNewResponseStatusActive            AccountNewResponseStatus = "active"
)

func (r AccountNewResponseStatus) IsKnown() bool {
	switch r {
	case AccountNewResponseStatusCreated, AccountNewResponseStatusOnboardingStarted, AccountNewResponseStatusUnderReview, AccountNewResponseStatusActive:
		return true
	}
	return false
}

type AccountNewResponseUser struct {
	AccountID string                     `json:"account_id,required,nullable" format:"uuid4"`
	AvatarURL string                     `json:"avatar_url,required,nullable"`
	Email     string                     `json:"email,required" format:"email"`
	Username  string                     `json:"username,required"`
	JSON      accountNewResponseUserJSON `json:"-"`
}

// accountNewResponseUserJSON contains the JSON metadata for the struct
// [AccountNewResponseUser]
type accountNewResponseUserJSON struct {
	AccountID   apijson.Field
	AvatarURL   apijson.Field
	Email       apijson.Field
	Username    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountNewResponseUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountNewResponseUserJSON) RawJSON() string {
	return r.raw
}

type AccountGetResponse struct {
	ID                 string                           `json:"id,required" format:"uuid"`
	AccountType        AccountGetResponseAccountType    `json:"account_type,required"`
	Country            string                           `json:"country,required"`
	IsChargesEnabled   bool                             `json:"is_charges_enabled,required"`
	IsDetailsSubmitted bool                             `json:"is_details_submitted,required"`
	IsPayoutsEnabled   bool                             `json:"is_payouts_enabled,required"`
	OpenCollectiveSlug string                           `json:"open_collective_slug,required,nullable"`
	Organizations      []AccountGetResponseOrganization `json:"organizations,required"`
	Status             AccountGetResponseStatus         `json:"status,required"`
	StripeID           string                           `json:"stripe_id,required,nullable"`
	Users              []AccountGetResponseUser         `json:"users,required"`
	JSON               accountGetResponseJSON           `json:"-"`
}

// accountGetResponseJSON contains the JSON metadata for the struct
// [AccountGetResponse]
type accountGetResponseJSON struct {
	ID                 apijson.Field
	AccountType        apijson.Field
	Country            apijson.Field
	IsChargesEnabled   apijson.Field
	IsDetailsSubmitted apijson.Field
	IsPayoutsEnabled   apijson.Field
	OpenCollectiveSlug apijson.Field
	Organizations      apijson.Field
	Status             apijson.Field
	StripeID           apijson.Field
	Users              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountGetResponseAccountType string

const (
	AccountGetResponseAccountTypeStripe         AccountGetResponseAccountType = "stripe"
	AccountGetResponseAccountTypeOpenCollective AccountGetResponseAccountType = "open_collective"
)

func (r AccountGetResponseAccountType) IsKnown() bool {
	switch r {
	case AccountGetResponseAccountTypeStripe, AccountGetResponseAccountTypeOpenCollective:
		return true
	}
	return false
}

type AccountGetResponseOrganization struct {
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
	FeatureSettings AccountGetResponseOrganizationsFeatureSettings `json:"feature_settings,required,nullable"`
	Location        string                                         `json:"location,required,nullable"`
	// Last modification timestamp of the object.
	ModifiedAt            time.Time `json:"modified_at,required,nullable" format:"date-time"`
	Name                  string    `json:"name,required"`
	PledgeBadgeShowAmount bool      `json:"pledge_badge_show_amount,required"`
	PledgeMinimumAmount   int64     `json:"pledge_minimum_amount,required"`
	// Settings for the organization profile
	ProfileSettings AccountGetResponseOrganizationsProfileSettings `json:"profile_settings,required,nullable"`
	Slug            string                                         `json:"slug,required"`
	TwitterUsername string                                         `json:"twitter_username,required,nullable"`
	JSON            accountGetResponseOrganizationJSON             `json:"-"`
}

// accountGetResponseOrganizationJSON contains the JSON metadata for the struct
// [AccountGetResponseOrganization]
type accountGetResponseOrganizationJSON struct {
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

func (r *AccountGetResponseOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGetResponseOrganizationJSON) RawJSON() string {
	return r.raw
}

// Settings for the organization features
type AccountGetResponseOrganizationsFeatureSettings struct {
	// If this organization has articles enabled
	ArticlesEnabled bool `json:"articles_enabled"`
	// If this organization has issue funding enabled
	IssueFundingEnabled bool `json:"issue_funding_enabled"`
	// If this organization has subscriptions enabled
	SubscriptionsEnabled bool                                               `json:"subscriptions_enabled"`
	JSON                 accountGetResponseOrganizationsFeatureSettingsJSON `json:"-"`
}

// accountGetResponseOrganizationsFeatureSettingsJSON contains the JSON metadata
// for the struct [AccountGetResponseOrganizationsFeatureSettings]
type accountGetResponseOrganizationsFeatureSettingsJSON struct {
	ArticlesEnabled      apijson.Field
	IssueFundingEnabled  apijson.Field
	SubscriptionsEnabled apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountGetResponseOrganizationsFeatureSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGetResponseOrganizationsFeatureSettingsJSON) RawJSON() string {
	return r.raw
}

// Settings for the organization profile
type AccountGetResponseOrganizationsProfileSettings struct {
	// A description of the organization
	Description string `json:"description,nullable"`
	// A list of featured organizations
	FeaturedOrganizations []string `json:"featured_organizations,nullable" format:"uuid4"`
	// A list of featured projects
	FeaturedProjects []string `json:"featured_projects,nullable" format:"uuid4"`
	// A list of links associated with the organization
	Links []string `json:"links,nullable" format:"uri"`
	// Subscription promotion settings
	Subscribe AccountGetResponseOrganizationsProfileSettingsSubscribe `json:"subscribe,nullable"`
	JSON      accountGetResponseOrganizationsProfileSettingsJSON      `json:"-"`
}

// accountGetResponseOrganizationsProfileSettingsJSON contains the JSON metadata
// for the struct [AccountGetResponseOrganizationsProfileSettings]
type accountGetResponseOrganizationsProfileSettingsJSON struct {
	Description           apijson.Field
	FeaturedOrganizations apijson.Field
	FeaturedProjects      apijson.Field
	Links                 apijson.Field
	Subscribe             apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AccountGetResponseOrganizationsProfileSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGetResponseOrganizationsProfileSettingsJSON) RawJSON() string {
	return r.raw
}

// Subscription promotion settings
type AccountGetResponseOrganizationsProfileSettingsSubscribe struct {
	// Include free subscribers in total count
	CountFree bool `json:"count_free"`
	// Promote email subscription (free)
	Promote bool `json:"promote"`
	// Show subscription count publicly
	ShowCount bool                                                        `json:"show_count"`
	JSON      accountGetResponseOrganizationsProfileSettingsSubscribeJSON `json:"-"`
}

// accountGetResponseOrganizationsProfileSettingsSubscribeJSON contains the JSON
// metadata for the struct
// [AccountGetResponseOrganizationsProfileSettingsSubscribe]
type accountGetResponseOrganizationsProfileSettingsSubscribeJSON struct {
	CountFree   apijson.Field
	Promote     apijson.Field
	ShowCount   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGetResponseOrganizationsProfileSettingsSubscribe) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGetResponseOrganizationsProfileSettingsSubscribeJSON) RawJSON() string {
	return r.raw
}

type AccountGetResponseStatus string

const (
	AccountGetResponseStatusCreated           AccountGetResponseStatus = "created"
	AccountGetResponseStatusOnboardingStarted AccountGetResponseStatus = "onboarding_started"
	AccountGetResponseStatusUnderReview       AccountGetResponseStatus = "under_review"
	AccountGetResponseStatusActive            AccountGetResponseStatus = "active"
)

func (r AccountGetResponseStatus) IsKnown() bool {
	switch r {
	case AccountGetResponseStatusCreated, AccountGetResponseStatusOnboardingStarted, AccountGetResponseStatusUnderReview, AccountGetResponseStatusActive:
		return true
	}
	return false
}

type AccountGetResponseUser struct {
	AccountID string                     `json:"account_id,required,nullable" format:"uuid4"`
	AvatarURL string                     `json:"avatar_url,required,nullable"`
	Email     string                     `json:"email,required" format:"email"`
	Username  string                     `json:"username,required"`
	JSON      accountGetResponseUserJSON `json:"-"`
}

// accountGetResponseUserJSON contains the JSON metadata for the struct
// [AccountGetResponseUser]
type accountGetResponseUserJSON struct {
	AccountID   apijson.Field
	AvatarURL   apijson.Field
	Email       apijson.Field
	Username    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGetResponseUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGetResponseUserJSON) RawJSON() string {
	return r.raw
}

type AccountDashboardLinkResponse struct {
	URL  string                           `json:"url,required"`
	JSON accountDashboardLinkResponseJSON `json:"-"`
}

// accountDashboardLinkResponseJSON contains the JSON metadata for the struct
// [AccountDashboardLinkResponse]
type accountDashboardLinkResponseJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDashboardLinkResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDashboardLinkResponseJSON) RawJSON() string {
	return r.raw
}

type AccountOnboardingLinkResponse struct {
	URL  string                            `json:"url,required"`
	JSON accountOnboardingLinkResponseJSON `json:"-"`
}

// accountOnboardingLinkResponseJSON contains the JSON metadata for the struct
// [AccountOnboardingLinkResponse]
type accountOnboardingLinkResponseJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountOnboardingLinkResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountOnboardingLinkResponseJSON) RawJSON() string {
	return r.raw
}

type AccountSearchResponse struct {
	Pagination AccountSearchResponsePagination `json:"pagination,required"`
	Items      []AccountSearchResponseItem     `json:"items"`
	JSON       accountSearchResponseJSON       `json:"-"`
}

// accountSearchResponseJSON contains the JSON metadata for the struct
// [AccountSearchResponse]
type accountSearchResponseJSON struct {
	Pagination  apijson.Field
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSearchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSearchResponseJSON) RawJSON() string {
	return r.raw
}

type AccountSearchResponsePagination struct {
	MaxPage    int64                               `json:"max_page,required"`
	TotalCount int64                               `json:"total_count,required"`
	JSON       accountSearchResponsePaginationJSON `json:"-"`
}

// accountSearchResponsePaginationJSON contains the JSON metadata for the struct
// [AccountSearchResponsePagination]
type accountSearchResponsePaginationJSON struct {
	MaxPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSearchResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSearchResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type AccountSearchResponseItem struct {
	ID                 string                                   `json:"id,required" format:"uuid"`
	AccountType        AccountSearchResponseItemsAccountType    `json:"account_type,required"`
	Country            string                                   `json:"country,required"`
	IsChargesEnabled   bool                                     `json:"is_charges_enabled,required"`
	IsDetailsSubmitted bool                                     `json:"is_details_submitted,required"`
	IsPayoutsEnabled   bool                                     `json:"is_payouts_enabled,required"`
	OpenCollectiveSlug string                                   `json:"open_collective_slug,required,nullable"`
	Organizations      []AccountSearchResponseItemsOrganization `json:"organizations,required"`
	Status             AccountSearchResponseItemsStatus         `json:"status,required"`
	StripeID           string                                   `json:"stripe_id,required,nullable"`
	Users              []AccountSearchResponseItemsUser         `json:"users,required"`
	JSON               accountSearchResponseItemJSON            `json:"-"`
}

// accountSearchResponseItemJSON contains the JSON metadata for the struct
// [AccountSearchResponseItem]
type accountSearchResponseItemJSON struct {
	ID                 apijson.Field
	AccountType        apijson.Field
	Country            apijson.Field
	IsChargesEnabled   apijson.Field
	IsDetailsSubmitted apijson.Field
	IsPayoutsEnabled   apijson.Field
	OpenCollectiveSlug apijson.Field
	Organizations      apijson.Field
	Status             apijson.Field
	StripeID           apijson.Field
	Users              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountSearchResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSearchResponseItemJSON) RawJSON() string {
	return r.raw
}

type AccountSearchResponseItemsAccountType string

const (
	AccountSearchResponseItemsAccountTypeStripe         AccountSearchResponseItemsAccountType = "stripe"
	AccountSearchResponseItemsAccountTypeOpenCollective AccountSearchResponseItemsAccountType = "open_collective"
)

func (r AccountSearchResponseItemsAccountType) IsKnown() bool {
	switch r {
	case AccountSearchResponseItemsAccountTypeStripe, AccountSearchResponseItemsAccountTypeOpenCollective:
		return true
	}
	return false
}

type AccountSearchResponseItemsOrganization struct {
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
	FeatureSettings AccountSearchResponseItemsOrganizationsFeatureSettings `json:"feature_settings,required,nullable"`
	Location        string                                                 `json:"location,required,nullable"`
	// Last modification timestamp of the object.
	ModifiedAt            time.Time `json:"modified_at,required,nullable" format:"date-time"`
	Name                  string    `json:"name,required"`
	PledgeBadgeShowAmount bool      `json:"pledge_badge_show_amount,required"`
	PledgeMinimumAmount   int64     `json:"pledge_minimum_amount,required"`
	// Settings for the organization profile
	ProfileSettings AccountSearchResponseItemsOrganizationsProfileSettings `json:"profile_settings,required,nullable"`
	Slug            string                                                 `json:"slug,required"`
	TwitterUsername string                                                 `json:"twitter_username,required,nullable"`
	JSON            accountSearchResponseItemsOrganizationJSON             `json:"-"`
}

// accountSearchResponseItemsOrganizationJSON contains the JSON metadata for the
// struct [AccountSearchResponseItemsOrganization]
type accountSearchResponseItemsOrganizationJSON struct {
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

func (r *AccountSearchResponseItemsOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSearchResponseItemsOrganizationJSON) RawJSON() string {
	return r.raw
}

// Settings for the organization features
type AccountSearchResponseItemsOrganizationsFeatureSettings struct {
	// If this organization has articles enabled
	ArticlesEnabled bool `json:"articles_enabled"`
	// If this organization has issue funding enabled
	IssueFundingEnabled bool `json:"issue_funding_enabled"`
	// If this organization has subscriptions enabled
	SubscriptionsEnabled bool                                                       `json:"subscriptions_enabled"`
	JSON                 accountSearchResponseItemsOrganizationsFeatureSettingsJSON `json:"-"`
}

// accountSearchResponseItemsOrganizationsFeatureSettingsJSON contains the JSON
// metadata for the struct [AccountSearchResponseItemsOrganizationsFeatureSettings]
type accountSearchResponseItemsOrganizationsFeatureSettingsJSON struct {
	ArticlesEnabled      apijson.Field
	IssueFundingEnabled  apijson.Field
	SubscriptionsEnabled apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountSearchResponseItemsOrganizationsFeatureSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSearchResponseItemsOrganizationsFeatureSettingsJSON) RawJSON() string {
	return r.raw
}

// Settings for the organization profile
type AccountSearchResponseItemsOrganizationsProfileSettings struct {
	// A description of the organization
	Description string `json:"description,nullable"`
	// A list of featured organizations
	FeaturedOrganizations []string `json:"featured_organizations,nullable" format:"uuid4"`
	// A list of featured projects
	FeaturedProjects []string `json:"featured_projects,nullable" format:"uuid4"`
	// A list of links associated with the organization
	Links []string `json:"links,nullable" format:"uri"`
	// Subscription promotion settings
	Subscribe AccountSearchResponseItemsOrganizationsProfileSettingsSubscribe `json:"subscribe,nullable"`
	JSON      accountSearchResponseItemsOrganizationsProfileSettingsJSON      `json:"-"`
}

// accountSearchResponseItemsOrganizationsProfileSettingsJSON contains the JSON
// metadata for the struct [AccountSearchResponseItemsOrganizationsProfileSettings]
type accountSearchResponseItemsOrganizationsProfileSettingsJSON struct {
	Description           apijson.Field
	FeaturedOrganizations apijson.Field
	FeaturedProjects      apijson.Field
	Links                 apijson.Field
	Subscribe             apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AccountSearchResponseItemsOrganizationsProfileSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSearchResponseItemsOrganizationsProfileSettingsJSON) RawJSON() string {
	return r.raw
}

// Subscription promotion settings
type AccountSearchResponseItemsOrganizationsProfileSettingsSubscribe struct {
	// Include free subscribers in total count
	CountFree bool `json:"count_free"`
	// Promote email subscription (free)
	Promote bool `json:"promote"`
	// Show subscription count publicly
	ShowCount bool                                                                `json:"show_count"`
	JSON      accountSearchResponseItemsOrganizationsProfileSettingsSubscribeJSON `json:"-"`
}

// accountSearchResponseItemsOrganizationsProfileSettingsSubscribeJSON contains the
// JSON metadata for the struct
// [AccountSearchResponseItemsOrganizationsProfileSettingsSubscribe]
type accountSearchResponseItemsOrganizationsProfileSettingsSubscribeJSON struct {
	CountFree   apijson.Field
	Promote     apijson.Field
	ShowCount   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSearchResponseItemsOrganizationsProfileSettingsSubscribe) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSearchResponseItemsOrganizationsProfileSettingsSubscribeJSON) RawJSON() string {
	return r.raw
}

type AccountSearchResponseItemsStatus string

const (
	AccountSearchResponseItemsStatusCreated           AccountSearchResponseItemsStatus = "created"
	AccountSearchResponseItemsStatusOnboardingStarted AccountSearchResponseItemsStatus = "onboarding_started"
	AccountSearchResponseItemsStatusUnderReview       AccountSearchResponseItemsStatus = "under_review"
	AccountSearchResponseItemsStatusActive            AccountSearchResponseItemsStatus = "active"
)

func (r AccountSearchResponseItemsStatus) IsKnown() bool {
	switch r {
	case AccountSearchResponseItemsStatusCreated, AccountSearchResponseItemsStatusOnboardingStarted, AccountSearchResponseItemsStatusUnderReview, AccountSearchResponseItemsStatusActive:
		return true
	}
	return false
}

type AccountSearchResponseItemsUser struct {
	AccountID string                             `json:"account_id,required,nullable" format:"uuid4"`
	AvatarURL string                             `json:"avatar_url,required,nullable"`
	Email     string                             `json:"email,required" format:"email"`
	Username  string                             `json:"username,required"`
	JSON      accountSearchResponseItemsUserJSON `json:"-"`
}

// accountSearchResponseItemsUserJSON contains the JSON metadata for the struct
// [AccountSearchResponseItemsUser]
type accountSearchResponseItemsUserJSON struct {
	AccountID   apijson.Field
	AvatarURL   apijson.Field
	Email       apijson.Field
	Username    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSearchResponseItemsUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSearchResponseItemsUserJSON) RawJSON() string {
	return r.raw
}

type AccountNewParams struct {
	AccountType param.Field[AccountNewParamsAccountType] `json:"account_type,required"`
	// Two letter uppercase country code
	Country            param.Field[string] `json:"country,required"`
	OpenCollectiveSlug param.Field[string] `json:"open_collective_slug"`
}

func (r AccountNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountNewParamsAccountType string

const (
	AccountNewParamsAccountTypeStripe         AccountNewParamsAccountType = "stripe"
	AccountNewParamsAccountTypeOpenCollective AccountNewParamsAccountType = "open_collective"
)

func (r AccountNewParamsAccountType) IsKnown() bool {
	switch r {
	case AccountNewParamsAccountTypeStripe, AccountNewParamsAccountTypeOpenCollective:
		return true
	}
	return false
}

type AccountOnboardingLinkParams struct {
	ReturnPath param.Field[string] `query:"return_path,required"`
}

// URLQuery serializes [AccountOnboardingLinkParams]'s query parameters as
// `url.Values`.
func (r AccountOnboardingLinkParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountSearchParams struct {
	// Size of a page, defaults to 10. Maximum is 100.
	Limit param.Field[int64] `query:"limit"`
	// Page number, defaults to 1.
	Page param.Field[int64] `query:"page"`
}

// URLQuery serializes [AccountSearchParams]'s query parameters as `url.Values`.
func (r AccountSearchParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
