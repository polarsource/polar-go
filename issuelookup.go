// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package polar

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/polarsource/polar-go/internal/apijson"
	"github.com/polarsource/polar-go/internal/apiquery"
	"github.com/polarsource/polar-go/internal/param"
	"github.com/polarsource/polar-go/internal/requestconfig"
	"github.com/polarsource/polar-go/option"
)

// IssueLookupService contains methods and other services that help with
// interacting with the polar API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIssueLookupService] method instead.
type IssueLookupService struct {
	Options []option.RequestOption
}

// NewIssueLookupService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewIssueLookupService(opts ...option.RequestOption) (r *IssueLookupService) {
	r = &IssueLookupService{}
	r.Options = opts
	return
}

// Lookup
func (r *IssueLookupService) Get(ctx context.Context, query IssueLookupGetParams, opts ...option.RequestOption) (res *IssueLookupGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/issues/lookup"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type IssueLookupGetResponse struct {
	ID             string                        `json:"id,required" format:"uuid"`
	Funding        IssueLookupGetResponseFunding `json:"funding,required"`
	IssueCreatedAt time.Time                     `json:"issue_created_at,required" format:"date-time"`
	// If a maintainer needs to mark this issue as solved
	NeedsConfirmationSolved bool `json:"needs_confirmation_solved,required"`
	// GitHub #number
	Number int64 `json:"number,required"`
	// Issue platform (currently always GitHub)
	Platform IssueLookupGetResponsePlatform `json:"platform,required"`
	// If this issue currently has the Polar badge SVG embedded
	PledgeBadgeCurrentlyEmbedded bool `json:"pledge_badge_currently_embedded,required"`
	// The repository that the issue is in
	Repository IssueLookupGetResponseRepository `json:"repository,required"`
	State      IssueLookupGetResponseState      `json:"state,required"`
	// GitHub issue title
	Title string `json:"title,required"`
	// GitHub assignees
	Assignees []IssueLookupGetResponseAssignee `json:"assignees,nullable"`
	// GitHub author
	Author IssueLookupGetResponseAuthor `json:"author,nullable"`
	// Optional custom badge SVG promotional content
	BadgeCustomContent string `json:"badge_custom_content,nullable"`
	// GitHub issue body
	Body string `json:"body,nullable"`
	// Number of GitHub comments made on the issue
	Comments int64 `json:"comments,nullable"`
	// If this issue has been marked as confirmed solved through Polar
	ConfirmedSolvedAt time.Time                     `json:"confirmed_solved_at,nullable" format:"date-time"`
	IssueClosedAt     time.Time                     `json:"issue_closed_at,nullable" format:"date-time"`
	IssueModifiedAt   time.Time                     `json:"issue_modified_at,nullable" format:"date-time"`
	Labels            []IssueLookupGetResponseLabel `json:"labels"`
	// GitHub reactions
	Reactions IssueLookupGetResponseReactions `json:"reactions,nullable"`
	// Share of rewrads that will be rewarded to contributors of this issue. A number
	// between 0 and 100 (inclusive).
	UpfrontSplitToContributors int64                      `json:"upfront_split_to_contributors,nullable"`
	JSON                       issueLookupGetResponseJSON `json:"-"`
}

// issueLookupGetResponseJSON contains the JSON metadata for the struct
// [IssueLookupGetResponse]
type issueLookupGetResponseJSON struct {
	ID                           apijson.Field
	Funding                      apijson.Field
	IssueCreatedAt               apijson.Field
	NeedsConfirmationSolved      apijson.Field
	Number                       apijson.Field
	Platform                     apijson.Field
	PledgeBadgeCurrentlyEmbedded apijson.Field
	Repository                   apijson.Field
	State                        apijson.Field
	Title                        apijson.Field
	Assignees                    apijson.Field
	Author                       apijson.Field
	BadgeCustomContent           apijson.Field
	Body                         apijson.Field
	Comments                     apijson.Field
	ConfirmedSolvedAt            apijson.Field
	IssueClosedAt                apijson.Field
	IssueModifiedAt              apijson.Field
	Labels                       apijson.Field
	Reactions                    apijson.Field
	UpfrontSplitToContributors   apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *IssueLookupGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueLookupGetResponseJSON) RawJSON() string {
	return r.raw
}

type IssueLookupGetResponseFunding struct {
	FundingGoal IssueLookupGetResponseFundingFundingGoal `json:"funding_goal,nullable"`
	// Sum of pledges to this isuse (including currently open pledges and pledges that
	// have been paid out). Always in USD.
	PledgesSum IssueLookupGetResponseFundingPledgesSum `json:"pledges_sum,nullable"`
	JSON       issueLookupGetResponseFundingJSON       `json:"-"`
}

// issueLookupGetResponseFundingJSON contains the JSON metadata for the struct
// [IssueLookupGetResponseFunding]
type issueLookupGetResponseFundingJSON struct {
	FundingGoal apijson.Field
	PledgesSum  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IssueLookupGetResponseFunding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueLookupGetResponseFundingJSON) RawJSON() string {
	return r.raw
}

type IssueLookupGetResponseFundingFundingGoal struct {
	// Amount in the currencies smallest unit (cents if currency is USD)
	Amount int64 `json:"amount,required"`
	// Three letter currency code (eg: USD)
	Currency string                                       `json:"currency,required"`
	JSON     issueLookupGetResponseFundingFundingGoalJSON `json:"-"`
}

// issueLookupGetResponseFundingFundingGoalJSON contains the JSON metadata for the
// struct [IssueLookupGetResponseFundingFundingGoal]
type issueLookupGetResponseFundingFundingGoalJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IssueLookupGetResponseFundingFundingGoal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueLookupGetResponseFundingFundingGoalJSON) RawJSON() string {
	return r.raw
}

// Sum of pledges to this isuse (including currently open pledges and pledges that
// have been paid out). Always in USD.
type IssueLookupGetResponseFundingPledgesSum struct {
	// Amount in the currencies smallest unit (cents if currency is USD)
	Amount int64 `json:"amount,required"`
	// Three letter currency code (eg: USD)
	Currency string                                      `json:"currency,required"`
	JSON     issueLookupGetResponseFundingPledgesSumJSON `json:"-"`
}

// issueLookupGetResponseFundingPledgesSumJSON contains the JSON metadata for the
// struct [IssueLookupGetResponseFundingPledgesSum]
type issueLookupGetResponseFundingPledgesSumJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IssueLookupGetResponseFundingPledgesSum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueLookupGetResponseFundingPledgesSumJSON) RawJSON() string {
	return r.raw
}

// Issue platform (currently always GitHub)
type IssueLookupGetResponsePlatform string

const (
	IssueLookupGetResponsePlatformGitHub IssueLookupGetResponsePlatform = "github"
)

func (r IssueLookupGetResponsePlatform) IsKnown() bool {
	switch r {
	case IssueLookupGetResponsePlatformGitHub:
		return true
	}
	return false
}

// The repository that the issue is in
type IssueLookupGetResponseRepository struct {
	ID           string                                       `json:"id,required" format:"uuid"`
	Description  string                                       `json:"description,required,nullable"`
	Homepage     string                                       `json:"homepage,required,nullable"`
	IsPrivate    bool                                         `json:"is_private,required"`
	License      string                                       `json:"license,required,nullable"`
	Name         string                                       `json:"name,required"`
	Organization IssueLookupGetResponseRepositoryOrganization `json:"organization,required"`
	Platform     IssueLookupGetResponseRepositoryPlatform     `json:"platform,required"`
	// Settings for the repository profile
	ProfileSettings IssueLookupGetResponseRepositoryProfileSettings `json:"profile_settings,required,nullable"`
	Stars           int64                                           `json:"stars,required,nullable"`
	JSON            issueLookupGetResponseRepositoryJSON            `json:"-"`
}

// issueLookupGetResponseRepositoryJSON contains the JSON metadata for the struct
// [IssueLookupGetResponseRepository]
type issueLookupGetResponseRepositoryJSON struct {
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

func (r *IssueLookupGetResponseRepository) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueLookupGetResponseRepositoryJSON) RawJSON() string {
	return r.raw
}

type IssueLookupGetResponseRepositoryOrganization struct {
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
	OrganizationID  string                                               `json:"organization_id,required,nullable" format:"uuid4"`
	Platform        IssueLookupGetResponseRepositoryOrganizationPlatform `json:"platform,required"`
	PrettyName      string                                               `json:"pretty_name,required,nullable"`
	TwitterUsername string                                               `json:"twitter_username,required,nullable"`
	JSON            issueLookupGetResponseRepositoryOrganizationJSON     `json:"-"`
}

// issueLookupGetResponseRepositoryOrganizationJSON contains the JSON metadata for
// the struct [IssueLookupGetResponseRepositoryOrganization]
type issueLookupGetResponseRepositoryOrganizationJSON struct {
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

func (r *IssueLookupGetResponseRepositoryOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueLookupGetResponseRepositoryOrganizationJSON) RawJSON() string {
	return r.raw
}

type IssueLookupGetResponseRepositoryOrganizationPlatform string

const (
	IssueLookupGetResponseRepositoryOrganizationPlatformGitHub IssueLookupGetResponseRepositoryOrganizationPlatform = "github"
)

func (r IssueLookupGetResponseRepositoryOrganizationPlatform) IsKnown() bool {
	switch r {
	case IssueLookupGetResponseRepositoryOrganizationPlatformGitHub:
		return true
	}
	return false
}

type IssueLookupGetResponseRepositoryPlatform string

const (
	IssueLookupGetResponseRepositoryPlatformGitHub IssueLookupGetResponseRepositoryPlatform = "github"
)

func (r IssueLookupGetResponseRepositoryPlatform) IsKnown() bool {
	switch r {
	case IssueLookupGetResponseRepositoryPlatformGitHub:
		return true
	}
	return false
}

// Settings for the repository profile
type IssueLookupGetResponseRepositoryProfileSettings struct {
	// A URL to a cover image
	CoverImageURL string `json:"cover_image_url,nullable"`
	// A description of the repository
	Description string `json:"description,nullable"`
	// A list of featured organizations
	FeaturedOrganizations []string `json:"featured_organizations,nullable" format:"uuid4"`
	// A list of highlighted subscription tiers
	HighlightedSubscriptionTiers []string `json:"highlighted_subscription_tiers,nullable" format:"uuid4"`
	// A list of links related to the repository
	Links []string                                            `json:"links,nullable" format:"uri"`
	JSON  issueLookupGetResponseRepositoryProfileSettingsJSON `json:"-"`
}

// issueLookupGetResponseRepositoryProfileSettingsJSON contains the JSON metadata
// for the struct [IssueLookupGetResponseRepositoryProfileSettings]
type issueLookupGetResponseRepositoryProfileSettingsJSON struct {
	CoverImageURL                apijson.Field
	Description                  apijson.Field
	FeaturedOrganizations        apijson.Field
	HighlightedSubscriptionTiers apijson.Field
	Links                        apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *IssueLookupGetResponseRepositoryProfileSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueLookupGetResponseRepositoryProfileSettingsJSON) RawJSON() string {
	return r.raw
}

type IssueLookupGetResponseState string

const (
	IssueLookupGetResponseStateOpen   IssueLookupGetResponseState = "open"
	IssueLookupGetResponseStateClosed IssueLookupGetResponseState = "closed"
)

func (r IssueLookupGetResponseState) IsKnown() bool {
	switch r {
	case IssueLookupGetResponseStateOpen, IssueLookupGetResponseStateClosed:
		return true
	}
	return false
}

type IssueLookupGetResponseAssignee struct {
	ID        int64                              `json:"id,required"`
	AvatarURL string                             `json:"avatar_url,required" format:"uri"`
	HTMLURL   string                             `json:"html_url,required" format:"uri"`
	Login     string                             `json:"login,required"`
	JSON      issueLookupGetResponseAssigneeJSON `json:"-"`
}

// issueLookupGetResponseAssigneeJSON contains the JSON metadata for the struct
// [IssueLookupGetResponseAssignee]
type issueLookupGetResponseAssigneeJSON struct {
	ID          apijson.Field
	AvatarURL   apijson.Field
	HTMLURL     apijson.Field
	Login       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IssueLookupGetResponseAssignee) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueLookupGetResponseAssigneeJSON) RawJSON() string {
	return r.raw
}

// GitHub author
type IssueLookupGetResponseAuthor struct {
	ID        int64                            `json:"id,required"`
	AvatarURL string                           `json:"avatar_url,required" format:"uri"`
	HTMLURL   string                           `json:"html_url,required" format:"uri"`
	Login     string                           `json:"login,required"`
	JSON      issueLookupGetResponseAuthorJSON `json:"-"`
}

// issueLookupGetResponseAuthorJSON contains the JSON metadata for the struct
// [IssueLookupGetResponseAuthor]
type issueLookupGetResponseAuthorJSON struct {
	ID          apijson.Field
	AvatarURL   apijson.Field
	HTMLURL     apijson.Field
	Login       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IssueLookupGetResponseAuthor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueLookupGetResponseAuthorJSON) RawJSON() string {
	return r.raw
}

type IssueLookupGetResponseLabel struct {
	Color string                          `json:"color,required"`
	Name  string                          `json:"name,required"`
	JSON  issueLookupGetResponseLabelJSON `json:"-"`
}

// issueLookupGetResponseLabelJSON contains the JSON metadata for the struct
// [IssueLookupGetResponseLabel]
type issueLookupGetResponseLabelJSON struct {
	Color       apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IssueLookupGetResponseLabel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueLookupGetResponseLabelJSON) RawJSON() string {
	return r.raw
}

// GitHub reactions
type IssueLookupGetResponseReactions struct {
	Confused   int64                               `json:"confused,required"`
	Eyes       int64                               `json:"eyes,required"`
	Heart      int64                               `json:"heart,required"`
	Hooray     int64                               `json:"hooray,required"`
	Laugh      int64                               `json:"laugh,required"`
	MinusOne   int64                               `json:"minus_one,required"`
	PlusOne    int64                               `json:"plus_one,required"`
	Rocket     int64                               `json:"rocket,required"`
	TotalCount int64                               `json:"total_count,required"`
	JSON       issueLookupGetResponseReactionsJSON `json:"-"`
}

// issueLookupGetResponseReactionsJSON contains the JSON metadata for the struct
// [IssueLookupGetResponseReactions]
type issueLookupGetResponseReactionsJSON struct {
	Confused    apijson.Field
	Eyes        apijson.Field
	Heart       apijson.Field
	Hooray      apijson.Field
	Laugh       apijson.Field
	MinusOne    apijson.Field
	PlusOne     apijson.Field
	Rocket      apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IssueLookupGetResponseReactions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueLookupGetResponseReactionsJSON) RawJSON() string {
	return r.raw
}

type IssueLookupGetParams struct {
	// URL to issue on external source
	ExternalURL param.Field[string] `query:"external_url"`
}

// URLQuery serializes [IssueLookupGetParams]'s query parameters as `url.Values`.
func (r IssueLookupGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
