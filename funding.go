// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package polar

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/polar-go/internal/apijson"
	"github.com/stainless-sdks/polar-go/internal/apiquery"
	"github.com/stainless-sdks/polar-go/internal/param"
	"github.com/stainless-sdks/polar-go/internal/requestconfig"
	"github.com/stainless-sdks/polar-go/option"
)

// FundingService contains methods and other services that help with interacting
// with the polar API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFundingService] method instead.
type FundingService struct {
	Options []option.RequestOption
}

// NewFundingService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFundingService(opts ...option.RequestOption) (r *FundingService) {
	r = &FundingService{}
	r.Options = opts
	return
}

// Lookup
func (r *FundingService) Lookup(ctx context.Context, query FundingLookupParams, opts ...option.RequestOption) (res *FundingLookupResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/funding/lookup"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Search
func (r *FundingService) Search(ctx context.Context, query FundingSearchParams, opts ...option.RequestOption) (res *FundingSearchResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/funding/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type FundingLookupResponse struct {
	Issue            FundingLookupResponseIssue            `json:"issue,required"`
	PledgesSummaries FundingLookupResponsePledgesSummaries `json:"pledges_summaries,required"`
	Total            FundingLookupResponseTotal            `json:"total,required"`
	FundingGoal      FundingLookupResponseFundingGoal      `json:"funding_goal,nullable"`
	JSON             fundingLookupResponseJSON             `json:"-"`
}

// fundingLookupResponseJSON contains the JSON metadata for the struct
// [FundingLookupResponse]
type fundingLookupResponseJSON struct {
	Issue            apijson.Field
	PledgesSummaries apijson.Field
	Total            apijson.Field
	FundingGoal      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *FundingLookupResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingLookupResponseJSON) RawJSON() string {
	return r.raw
}

type FundingLookupResponseIssue struct {
	ID             string                            `json:"id,required" format:"uuid"`
	Funding        FundingLookupResponseIssueFunding `json:"funding,required"`
	IssueCreatedAt time.Time                         `json:"issue_created_at,required" format:"date-time"`
	// If a maintainer needs to mark this issue as solved
	NeedsConfirmationSolved bool `json:"needs_confirmation_solved,required"`
	// GitHub #number
	Number int64 `json:"number,required"`
	// Issue platform (currently always GitHub)
	Platform FundingLookupResponseIssuePlatform `json:"platform,required"`
	// If this issue currently has the Polar badge SVG embedded
	PledgeBadgeCurrentlyEmbedded bool `json:"pledge_badge_currently_embedded,required"`
	// The repository that the issue is in
	Repository FundingLookupResponseIssueRepository `json:"repository,required"`
	State      FundingLookupResponseIssueState      `json:"state,required"`
	// GitHub issue title
	Title string `json:"title,required"`
	// GitHub assignees
	Assignees []FundingLookupResponseIssueAssignee `json:"assignees,nullable"`
	// GitHub author
	Author FundingLookupResponseIssueAuthor `json:"author,nullable"`
	// Optional custom badge SVG promotional content
	BadgeCustomContent string `json:"badge_custom_content,nullable"`
	// GitHub issue body
	Body string `json:"body,nullable"`
	// Number of GitHub comments made on the issue
	Comments int64 `json:"comments,nullable"`
	// If this issue has been marked as confirmed solved through Polar
	ConfirmedSolvedAt time.Time                         `json:"confirmed_solved_at,nullable" format:"date-time"`
	IssueClosedAt     time.Time                         `json:"issue_closed_at,nullable" format:"date-time"`
	IssueModifiedAt   time.Time                         `json:"issue_modified_at,nullable" format:"date-time"`
	Labels            []FundingLookupResponseIssueLabel `json:"labels"`
	// GitHub reactions
	Reactions FundingLookupResponseIssueReactions `json:"reactions,nullable"`
	// Share of rewrads that will be rewarded to contributors of this issue. A number
	// between 0 and 100 (inclusive).
	UpfrontSplitToContributors int64                          `json:"upfront_split_to_contributors,nullable"`
	JSON                       fundingLookupResponseIssueJSON `json:"-"`
}

// fundingLookupResponseIssueJSON contains the JSON metadata for the struct
// [FundingLookupResponseIssue]
type fundingLookupResponseIssueJSON struct {
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

func (r *FundingLookupResponseIssue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingLookupResponseIssueJSON) RawJSON() string {
	return r.raw
}

type FundingLookupResponseIssueFunding struct {
	FundingGoal FundingLookupResponseIssueFundingFundingGoal `json:"funding_goal,nullable"`
	// Sum of pledges to this isuse (including currently open pledges and pledges that
	// have been paid out). Always in USD.
	PledgesSum FundingLookupResponseIssueFundingPledgesSum `json:"pledges_sum,nullable"`
	JSON       fundingLookupResponseIssueFundingJSON       `json:"-"`
}

// fundingLookupResponseIssueFundingJSON contains the JSON metadata for the struct
// [FundingLookupResponseIssueFunding]
type fundingLookupResponseIssueFundingJSON struct {
	FundingGoal apijson.Field
	PledgesSum  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FundingLookupResponseIssueFunding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingLookupResponseIssueFundingJSON) RawJSON() string {
	return r.raw
}

type FundingLookupResponseIssueFundingFundingGoal struct {
	// Amount in the currencies smallest unit (cents if currency is USD)
	Amount int64 `json:"amount,required"`
	// Three letter currency code (eg: USD)
	Currency string                                           `json:"currency,required"`
	JSON     fundingLookupResponseIssueFundingFundingGoalJSON `json:"-"`
}

// fundingLookupResponseIssueFundingFundingGoalJSON contains the JSON metadata for
// the struct [FundingLookupResponseIssueFundingFundingGoal]
type fundingLookupResponseIssueFundingFundingGoalJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FundingLookupResponseIssueFundingFundingGoal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingLookupResponseIssueFundingFundingGoalJSON) RawJSON() string {
	return r.raw
}

// Sum of pledges to this isuse (including currently open pledges and pledges that
// have been paid out). Always in USD.
type FundingLookupResponseIssueFundingPledgesSum struct {
	// Amount in the currencies smallest unit (cents if currency is USD)
	Amount int64 `json:"amount,required"`
	// Three letter currency code (eg: USD)
	Currency string                                          `json:"currency,required"`
	JSON     fundingLookupResponseIssueFundingPledgesSumJSON `json:"-"`
}

// fundingLookupResponseIssueFundingPledgesSumJSON contains the JSON metadata for
// the struct [FundingLookupResponseIssueFundingPledgesSum]
type fundingLookupResponseIssueFundingPledgesSumJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FundingLookupResponseIssueFundingPledgesSum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingLookupResponseIssueFundingPledgesSumJSON) RawJSON() string {
	return r.raw
}

// Issue platform (currently always GitHub)
type FundingLookupResponseIssuePlatform string

const (
	FundingLookupResponseIssuePlatformGitHub FundingLookupResponseIssuePlatform = "github"
)

func (r FundingLookupResponseIssuePlatform) IsKnown() bool {
	switch r {
	case FundingLookupResponseIssuePlatformGitHub:
		return true
	}
	return false
}

// The repository that the issue is in
type FundingLookupResponseIssueRepository struct {
	ID           string                                           `json:"id,required" format:"uuid"`
	IsPrivate    bool                                             `json:"is_private,required"`
	Name         string                                           `json:"name,required"`
	Organization FundingLookupResponseIssueRepositoryOrganization `json:"organization,required"`
	Platform     FundingLookupResponseIssueRepositoryPlatform     `json:"platform,required"`
	// Settings for the repository profile
	ProfileSettings FundingLookupResponseIssueRepositoryProfileSettings `json:"profile_settings,required,nullable"`
	Description     string                                              `json:"description,nullable"`
	Homepage        string                                              `json:"homepage,nullable"`
	License         string                                              `json:"license,nullable"`
	Stars           int64                                               `json:"stars,nullable"`
	JSON            fundingLookupResponseIssueRepositoryJSON            `json:"-"`
}

// fundingLookupResponseIssueRepositoryJSON contains the JSON metadata for the
// struct [FundingLookupResponseIssueRepository]
type fundingLookupResponseIssueRepositoryJSON struct {
	ID              apijson.Field
	IsPrivate       apijson.Field
	Name            apijson.Field
	Organization    apijson.Field
	Platform        apijson.Field
	ProfileSettings apijson.Field
	Description     apijson.Field
	Homepage        apijson.Field
	License         apijson.Field
	Stars           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *FundingLookupResponseIssueRepository) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingLookupResponseIssueRepositoryJSON) RawJSON() string {
	return r.raw
}

type FundingLookupResponseIssueRepositoryOrganization struct {
	ID         string                                                   `json:"id,required" format:"uuid"`
	AvatarURL  string                                                   `json:"avatar_url,required"`
	IsPersonal bool                                                     `json:"is_personal,required"`
	Name       string                                                   `json:"name,required"`
	Platform   FundingLookupResponseIssueRepositoryOrganizationPlatform `json:"platform,required"`
	Bio        string                                                   `json:"bio,nullable"`
	Blog       string                                                   `json:"blog,nullable"`
	Company    string                                                   `json:"company,nullable"`
	Email      string                                                   `json:"email,nullable"`
	Location   string                                                   `json:"location,nullable"`
	// The organization ID.
	OrganizationID  string                                               `json:"organization_id,nullable" format:"uuid4"`
	PrettyName      string                                               `json:"pretty_name,nullable"`
	TwitterUsername string                                               `json:"twitter_username,nullable"`
	JSON            fundingLookupResponseIssueRepositoryOrganizationJSON `json:"-"`
}

// fundingLookupResponseIssueRepositoryOrganizationJSON contains the JSON metadata
// for the struct [FundingLookupResponseIssueRepositoryOrganization]
type fundingLookupResponseIssueRepositoryOrganizationJSON struct {
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

func (r *FundingLookupResponseIssueRepositoryOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingLookupResponseIssueRepositoryOrganizationJSON) RawJSON() string {
	return r.raw
}

type FundingLookupResponseIssueRepositoryOrganizationPlatform string

const (
	FundingLookupResponseIssueRepositoryOrganizationPlatformGitHub FundingLookupResponseIssueRepositoryOrganizationPlatform = "github"
)

func (r FundingLookupResponseIssueRepositoryOrganizationPlatform) IsKnown() bool {
	switch r {
	case FundingLookupResponseIssueRepositoryOrganizationPlatformGitHub:
		return true
	}
	return false
}

type FundingLookupResponseIssueRepositoryPlatform string

const (
	FundingLookupResponseIssueRepositoryPlatformGitHub FundingLookupResponseIssueRepositoryPlatform = "github"
)

func (r FundingLookupResponseIssueRepositoryPlatform) IsKnown() bool {
	switch r {
	case FundingLookupResponseIssueRepositoryPlatformGitHub:
		return true
	}
	return false
}

// Settings for the repository profile
type FundingLookupResponseIssueRepositoryProfileSettings struct {
	// A URL to a cover image
	CoverImageURL string `json:"cover_image_url,nullable"`
	// A description of the repository
	Description string `json:"description,nullable"`
	// A list of featured organizations
	FeaturedOrganizations []string `json:"featured_organizations,nullable" format:"uuid4"`
	// A list of highlighted subscription tiers
	HighlightedSubscriptionTiers []string `json:"highlighted_subscription_tiers,nullable" format:"uuid4"`
	// A list of links related to the repository
	Links []string                                                `json:"links,nullable" format:"uri"`
	JSON  fundingLookupResponseIssueRepositoryProfileSettingsJSON `json:"-"`
}

// fundingLookupResponseIssueRepositoryProfileSettingsJSON contains the JSON
// metadata for the struct [FundingLookupResponseIssueRepositoryProfileSettings]
type fundingLookupResponseIssueRepositoryProfileSettingsJSON struct {
	CoverImageURL                apijson.Field
	Description                  apijson.Field
	FeaturedOrganizations        apijson.Field
	HighlightedSubscriptionTiers apijson.Field
	Links                        apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *FundingLookupResponseIssueRepositoryProfileSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingLookupResponseIssueRepositoryProfileSettingsJSON) RawJSON() string {
	return r.raw
}

type FundingLookupResponseIssueState string

const (
	FundingLookupResponseIssueStateOpen   FundingLookupResponseIssueState = "open"
	FundingLookupResponseIssueStateClosed FundingLookupResponseIssueState = "closed"
)

func (r FundingLookupResponseIssueState) IsKnown() bool {
	switch r {
	case FundingLookupResponseIssueStateOpen, FundingLookupResponseIssueStateClosed:
		return true
	}
	return false
}

type FundingLookupResponseIssueAssignee struct {
	ID        int64                                  `json:"id,required"`
	AvatarURL string                                 `json:"avatar_url,required" format:"uri"`
	HTMLURL   string                                 `json:"html_url,required" format:"uri"`
	Login     string                                 `json:"login,required"`
	JSON      fundingLookupResponseIssueAssigneeJSON `json:"-"`
}

// fundingLookupResponseIssueAssigneeJSON contains the JSON metadata for the struct
// [FundingLookupResponseIssueAssignee]
type fundingLookupResponseIssueAssigneeJSON struct {
	ID          apijson.Field
	AvatarURL   apijson.Field
	HTMLURL     apijson.Field
	Login       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FundingLookupResponseIssueAssignee) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingLookupResponseIssueAssigneeJSON) RawJSON() string {
	return r.raw
}

// GitHub author
type FundingLookupResponseIssueAuthor struct {
	ID        int64                                `json:"id,required"`
	AvatarURL string                               `json:"avatar_url,required" format:"uri"`
	HTMLURL   string                               `json:"html_url,required" format:"uri"`
	Login     string                               `json:"login,required"`
	JSON      fundingLookupResponseIssueAuthorJSON `json:"-"`
}

// fundingLookupResponseIssueAuthorJSON contains the JSON metadata for the struct
// [FundingLookupResponseIssueAuthor]
type fundingLookupResponseIssueAuthorJSON struct {
	ID          apijson.Field
	AvatarURL   apijson.Field
	HTMLURL     apijson.Field
	Login       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FundingLookupResponseIssueAuthor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingLookupResponseIssueAuthorJSON) RawJSON() string {
	return r.raw
}

type FundingLookupResponseIssueLabel struct {
	Color string                              `json:"color,required"`
	Name  string                              `json:"name,required"`
	JSON  fundingLookupResponseIssueLabelJSON `json:"-"`
}

// fundingLookupResponseIssueLabelJSON contains the JSON metadata for the struct
// [FundingLookupResponseIssueLabel]
type fundingLookupResponseIssueLabelJSON struct {
	Color       apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FundingLookupResponseIssueLabel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingLookupResponseIssueLabelJSON) RawJSON() string {
	return r.raw
}

// GitHub reactions
type FundingLookupResponseIssueReactions struct {
	Confused   int64                                   `json:"confused,required"`
	Eyes       int64                                   `json:"eyes,required"`
	Heart      int64                                   `json:"heart,required"`
	Hooray     int64                                   `json:"hooray,required"`
	Laugh      int64                                   `json:"laugh,required"`
	MinusOne   int64                                   `json:"minus_one,required"`
	PlusOne    int64                                   `json:"plus_one,required"`
	Rocket     int64                                   `json:"rocket,required"`
	TotalCount int64                                   `json:"total_count,required"`
	JSON       fundingLookupResponseIssueReactionsJSON `json:"-"`
}

// fundingLookupResponseIssueReactionsJSON contains the JSON metadata for the
// struct [FundingLookupResponseIssueReactions]
type fundingLookupResponseIssueReactionsJSON struct {
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

func (r *FundingLookupResponseIssueReactions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingLookupResponseIssueReactionsJSON) RawJSON() string {
	return r.raw
}

type FundingLookupResponsePledgesSummaries struct {
	PayDirectly     FundingLookupResponsePledgesSummariesPayDirectly     `json:"pay_directly,required"`
	PayOnCompletion FundingLookupResponsePledgesSummariesPayOnCompletion `json:"pay_on_completion,required"`
	PayUpfront      FundingLookupResponsePledgesSummariesPayUpfront      `json:"pay_upfront,required"`
	JSON            fundingLookupResponsePledgesSummariesJSON            `json:"-"`
}

// fundingLookupResponsePledgesSummariesJSON contains the JSON metadata for the
// struct [FundingLookupResponsePledgesSummaries]
type fundingLookupResponsePledgesSummariesJSON struct {
	PayDirectly     apijson.Field
	PayOnCompletion apijson.Field
	PayUpfront      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *FundingLookupResponsePledgesSummaries) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingLookupResponsePledgesSummariesJSON) RawJSON() string {
	return r.raw
}

type FundingLookupResponsePledgesSummariesPayDirectly struct {
	Pledgers []FundingLookupResponsePledgesSummariesPayDirectlyPledger `json:"pledgers,required"`
	Total    FundingLookupResponsePledgesSummariesPayDirectlyTotal     `json:"total,required"`
	JSON     fundingLookupResponsePledgesSummariesPayDirectlyJSON      `json:"-"`
}

// fundingLookupResponsePledgesSummariesPayDirectlyJSON contains the JSON metadata
// for the struct [FundingLookupResponsePledgesSummariesPayDirectly]
type fundingLookupResponsePledgesSummariesPayDirectlyJSON struct {
	Pledgers    apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FundingLookupResponsePledgesSummariesPayDirectly) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingLookupResponsePledgesSummariesPayDirectlyJSON) RawJSON() string {
	return r.raw
}

type FundingLookupResponsePledgesSummariesPayDirectlyPledger struct {
	Name           string                                                      `json:"name,required"`
	AvatarURL      string                                                      `json:"avatar_url,nullable"`
	GitHubUsername string                                                      `json:"github_username,nullable"`
	JSON           fundingLookupResponsePledgesSummariesPayDirectlyPledgerJSON `json:"-"`
}

// fundingLookupResponsePledgesSummariesPayDirectlyPledgerJSON contains the JSON
// metadata for the struct
// [FundingLookupResponsePledgesSummariesPayDirectlyPledger]
type fundingLookupResponsePledgesSummariesPayDirectlyPledgerJSON struct {
	Name           apijson.Field
	AvatarURL      apijson.Field
	GitHubUsername apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *FundingLookupResponsePledgesSummariesPayDirectlyPledger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingLookupResponsePledgesSummariesPayDirectlyPledgerJSON) RawJSON() string {
	return r.raw
}

type FundingLookupResponsePledgesSummariesPayDirectlyTotal struct {
	// Amount in the currencies smallest unit (cents if currency is USD)
	Amount int64 `json:"amount,required"`
	// Three letter currency code (eg: USD)
	Currency string                                                    `json:"currency,required"`
	JSON     fundingLookupResponsePledgesSummariesPayDirectlyTotalJSON `json:"-"`
}

// fundingLookupResponsePledgesSummariesPayDirectlyTotalJSON contains the JSON
// metadata for the struct [FundingLookupResponsePledgesSummariesPayDirectlyTotal]
type fundingLookupResponsePledgesSummariesPayDirectlyTotalJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FundingLookupResponsePledgesSummariesPayDirectlyTotal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingLookupResponsePledgesSummariesPayDirectlyTotalJSON) RawJSON() string {
	return r.raw
}

type FundingLookupResponsePledgesSummariesPayOnCompletion struct {
	Pledgers []FundingLookupResponsePledgesSummariesPayOnCompletionPledger `json:"pledgers,required"`
	Total    FundingLookupResponsePledgesSummariesPayOnCompletionTotal     `json:"total,required"`
	JSON     fundingLookupResponsePledgesSummariesPayOnCompletionJSON      `json:"-"`
}

// fundingLookupResponsePledgesSummariesPayOnCompletionJSON contains the JSON
// metadata for the struct [FundingLookupResponsePledgesSummariesPayOnCompletion]
type fundingLookupResponsePledgesSummariesPayOnCompletionJSON struct {
	Pledgers    apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FundingLookupResponsePledgesSummariesPayOnCompletion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingLookupResponsePledgesSummariesPayOnCompletionJSON) RawJSON() string {
	return r.raw
}

type FundingLookupResponsePledgesSummariesPayOnCompletionPledger struct {
	Name           string                                                          `json:"name,required"`
	AvatarURL      string                                                          `json:"avatar_url,nullable"`
	GitHubUsername string                                                          `json:"github_username,nullable"`
	JSON           fundingLookupResponsePledgesSummariesPayOnCompletionPledgerJSON `json:"-"`
}

// fundingLookupResponsePledgesSummariesPayOnCompletionPledgerJSON contains the
// JSON metadata for the struct
// [FundingLookupResponsePledgesSummariesPayOnCompletionPledger]
type fundingLookupResponsePledgesSummariesPayOnCompletionPledgerJSON struct {
	Name           apijson.Field
	AvatarURL      apijson.Field
	GitHubUsername apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *FundingLookupResponsePledgesSummariesPayOnCompletionPledger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingLookupResponsePledgesSummariesPayOnCompletionPledgerJSON) RawJSON() string {
	return r.raw
}

type FundingLookupResponsePledgesSummariesPayOnCompletionTotal struct {
	// Amount in the currencies smallest unit (cents if currency is USD)
	Amount int64 `json:"amount,required"`
	// Three letter currency code (eg: USD)
	Currency string                                                        `json:"currency,required"`
	JSON     fundingLookupResponsePledgesSummariesPayOnCompletionTotalJSON `json:"-"`
}

// fundingLookupResponsePledgesSummariesPayOnCompletionTotalJSON contains the JSON
// metadata for the struct
// [FundingLookupResponsePledgesSummariesPayOnCompletionTotal]
type fundingLookupResponsePledgesSummariesPayOnCompletionTotalJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FundingLookupResponsePledgesSummariesPayOnCompletionTotal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingLookupResponsePledgesSummariesPayOnCompletionTotalJSON) RawJSON() string {
	return r.raw
}

type FundingLookupResponsePledgesSummariesPayUpfront struct {
	Pledgers []FundingLookupResponsePledgesSummariesPayUpfrontPledger `json:"pledgers,required"`
	Total    FundingLookupResponsePledgesSummariesPayUpfrontTotal     `json:"total,required"`
	JSON     fundingLookupResponsePledgesSummariesPayUpfrontJSON      `json:"-"`
}

// fundingLookupResponsePledgesSummariesPayUpfrontJSON contains the JSON metadata
// for the struct [FundingLookupResponsePledgesSummariesPayUpfront]
type fundingLookupResponsePledgesSummariesPayUpfrontJSON struct {
	Pledgers    apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FundingLookupResponsePledgesSummariesPayUpfront) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingLookupResponsePledgesSummariesPayUpfrontJSON) RawJSON() string {
	return r.raw
}

type FundingLookupResponsePledgesSummariesPayUpfrontPledger struct {
	Name           string                                                     `json:"name,required"`
	AvatarURL      string                                                     `json:"avatar_url,nullable"`
	GitHubUsername string                                                     `json:"github_username,nullable"`
	JSON           fundingLookupResponsePledgesSummariesPayUpfrontPledgerJSON `json:"-"`
}

// fundingLookupResponsePledgesSummariesPayUpfrontPledgerJSON contains the JSON
// metadata for the struct [FundingLookupResponsePledgesSummariesPayUpfrontPledger]
type fundingLookupResponsePledgesSummariesPayUpfrontPledgerJSON struct {
	Name           apijson.Field
	AvatarURL      apijson.Field
	GitHubUsername apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *FundingLookupResponsePledgesSummariesPayUpfrontPledger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingLookupResponsePledgesSummariesPayUpfrontPledgerJSON) RawJSON() string {
	return r.raw
}

type FundingLookupResponsePledgesSummariesPayUpfrontTotal struct {
	// Amount in the currencies smallest unit (cents if currency is USD)
	Amount int64 `json:"amount,required"`
	// Three letter currency code (eg: USD)
	Currency string                                                   `json:"currency,required"`
	JSON     fundingLookupResponsePledgesSummariesPayUpfrontTotalJSON `json:"-"`
}

// fundingLookupResponsePledgesSummariesPayUpfrontTotalJSON contains the JSON
// metadata for the struct [FundingLookupResponsePledgesSummariesPayUpfrontTotal]
type fundingLookupResponsePledgesSummariesPayUpfrontTotalJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FundingLookupResponsePledgesSummariesPayUpfrontTotal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingLookupResponsePledgesSummariesPayUpfrontTotalJSON) RawJSON() string {
	return r.raw
}

type FundingLookupResponseTotal struct {
	// Amount in the currencies smallest unit (cents if currency is USD)
	Amount int64 `json:"amount,required"`
	// Three letter currency code (eg: USD)
	Currency string                         `json:"currency,required"`
	JSON     fundingLookupResponseTotalJSON `json:"-"`
}

// fundingLookupResponseTotalJSON contains the JSON metadata for the struct
// [FundingLookupResponseTotal]
type fundingLookupResponseTotalJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FundingLookupResponseTotal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingLookupResponseTotalJSON) RawJSON() string {
	return r.raw
}

type FundingLookupResponseFundingGoal struct {
	// Amount in the currencies smallest unit (cents if currency is USD)
	Amount int64 `json:"amount,required"`
	// Three letter currency code (eg: USD)
	Currency string                               `json:"currency,required"`
	JSON     fundingLookupResponseFundingGoalJSON `json:"-"`
}

// fundingLookupResponseFundingGoalJSON contains the JSON metadata for the struct
// [FundingLookupResponseFundingGoal]
type fundingLookupResponseFundingGoalJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FundingLookupResponseFundingGoal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingLookupResponseFundingGoalJSON) RawJSON() string {
	return r.raw
}

type FundingSearchResponse struct {
	Pagination FundingSearchResponsePagination `json:"pagination,required"`
	Items      []FundingSearchResponseItem     `json:"items"`
	JSON       fundingSearchResponseJSON       `json:"-"`
}

// fundingSearchResponseJSON contains the JSON metadata for the struct
// [FundingSearchResponse]
type fundingSearchResponseJSON struct {
	Pagination  apijson.Field
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FundingSearchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingSearchResponseJSON) RawJSON() string {
	return r.raw
}

type FundingSearchResponsePagination struct {
	MaxPage    int64                               `json:"max_page,required"`
	TotalCount int64                               `json:"total_count,required"`
	JSON       fundingSearchResponsePaginationJSON `json:"-"`
}

// fundingSearchResponsePaginationJSON contains the JSON metadata for the struct
// [FundingSearchResponsePagination]
type fundingSearchResponsePaginationJSON struct {
	MaxPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FundingSearchResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingSearchResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type FundingSearchResponseItem struct {
	Issue            FundingSearchResponseItemsIssue            `json:"issue,required"`
	PledgesSummaries FundingSearchResponseItemsPledgesSummaries `json:"pledges_summaries,required"`
	Total            FundingSearchResponseItemsTotal            `json:"total,required"`
	FundingGoal      FundingSearchResponseItemsFundingGoal      `json:"funding_goal,nullable"`
	JSON             fundingSearchResponseItemJSON              `json:"-"`
}

// fundingSearchResponseItemJSON contains the JSON metadata for the struct
// [FundingSearchResponseItem]
type fundingSearchResponseItemJSON struct {
	Issue            apijson.Field
	PledgesSummaries apijson.Field
	Total            apijson.Field
	FundingGoal      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *FundingSearchResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingSearchResponseItemJSON) RawJSON() string {
	return r.raw
}

type FundingSearchResponseItemsIssue struct {
	ID             string                                 `json:"id,required" format:"uuid"`
	Funding        FundingSearchResponseItemsIssueFunding `json:"funding,required"`
	IssueCreatedAt time.Time                              `json:"issue_created_at,required" format:"date-time"`
	// If a maintainer needs to mark this issue as solved
	NeedsConfirmationSolved bool `json:"needs_confirmation_solved,required"`
	// GitHub #number
	Number int64 `json:"number,required"`
	// Issue platform (currently always GitHub)
	Platform FundingSearchResponseItemsIssuePlatform `json:"platform,required"`
	// If this issue currently has the Polar badge SVG embedded
	PledgeBadgeCurrentlyEmbedded bool `json:"pledge_badge_currently_embedded,required"`
	// The repository that the issue is in
	Repository FundingSearchResponseItemsIssueRepository `json:"repository,required"`
	State      FundingSearchResponseItemsIssueState      `json:"state,required"`
	// GitHub issue title
	Title string `json:"title,required"`
	// GitHub assignees
	Assignees []FundingSearchResponseItemsIssueAssignee `json:"assignees,nullable"`
	// GitHub author
	Author FundingSearchResponseItemsIssueAuthor `json:"author,nullable"`
	// Optional custom badge SVG promotional content
	BadgeCustomContent string `json:"badge_custom_content,nullable"`
	// GitHub issue body
	Body string `json:"body,nullable"`
	// Number of GitHub comments made on the issue
	Comments int64 `json:"comments,nullable"`
	// If this issue has been marked as confirmed solved through Polar
	ConfirmedSolvedAt time.Time                              `json:"confirmed_solved_at,nullable" format:"date-time"`
	IssueClosedAt     time.Time                              `json:"issue_closed_at,nullable" format:"date-time"`
	IssueModifiedAt   time.Time                              `json:"issue_modified_at,nullable" format:"date-time"`
	Labels            []FundingSearchResponseItemsIssueLabel `json:"labels"`
	// GitHub reactions
	Reactions FundingSearchResponseItemsIssueReactions `json:"reactions,nullable"`
	// Share of rewrads that will be rewarded to contributors of this issue. A number
	// between 0 and 100 (inclusive).
	UpfrontSplitToContributors int64                               `json:"upfront_split_to_contributors,nullable"`
	JSON                       fundingSearchResponseItemsIssueJSON `json:"-"`
}

// fundingSearchResponseItemsIssueJSON contains the JSON metadata for the struct
// [FundingSearchResponseItemsIssue]
type fundingSearchResponseItemsIssueJSON struct {
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

func (r *FundingSearchResponseItemsIssue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingSearchResponseItemsIssueJSON) RawJSON() string {
	return r.raw
}

type FundingSearchResponseItemsIssueFunding struct {
	FundingGoal FundingSearchResponseItemsIssueFundingFundingGoal `json:"funding_goal,nullable"`
	// Sum of pledges to this isuse (including currently open pledges and pledges that
	// have been paid out). Always in USD.
	PledgesSum FundingSearchResponseItemsIssueFundingPledgesSum `json:"pledges_sum,nullable"`
	JSON       fundingSearchResponseItemsIssueFundingJSON       `json:"-"`
}

// fundingSearchResponseItemsIssueFundingJSON contains the JSON metadata for the
// struct [FundingSearchResponseItemsIssueFunding]
type fundingSearchResponseItemsIssueFundingJSON struct {
	FundingGoal apijson.Field
	PledgesSum  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FundingSearchResponseItemsIssueFunding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingSearchResponseItemsIssueFundingJSON) RawJSON() string {
	return r.raw
}

type FundingSearchResponseItemsIssueFundingFundingGoal struct {
	// Amount in the currencies smallest unit (cents if currency is USD)
	Amount int64 `json:"amount,required"`
	// Three letter currency code (eg: USD)
	Currency string                                                `json:"currency,required"`
	JSON     fundingSearchResponseItemsIssueFundingFundingGoalJSON `json:"-"`
}

// fundingSearchResponseItemsIssueFundingFundingGoalJSON contains the JSON metadata
// for the struct [FundingSearchResponseItemsIssueFundingFundingGoal]
type fundingSearchResponseItemsIssueFundingFundingGoalJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FundingSearchResponseItemsIssueFundingFundingGoal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingSearchResponseItemsIssueFundingFundingGoalJSON) RawJSON() string {
	return r.raw
}

// Sum of pledges to this isuse (including currently open pledges and pledges that
// have been paid out). Always in USD.
type FundingSearchResponseItemsIssueFundingPledgesSum struct {
	// Amount in the currencies smallest unit (cents if currency is USD)
	Amount int64 `json:"amount,required"`
	// Three letter currency code (eg: USD)
	Currency string                                               `json:"currency,required"`
	JSON     fundingSearchResponseItemsIssueFundingPledgesSumJSON `json:"-"`
}

// fundingSearchResponseItemsIssueFundingPledgesSumJSON contains the JSON metadata
// for the struct [FundingSearchResponseItemsIssueFundingPledgesSum]
type fundingSearchResponseItemsIssueFundingPledgesSumJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FundingSearchResponseItemsIssueFundingPledgesSum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingSearchResponseItemsIssueFundingPledgesSumJSON) RawJSON() string {
	return r.raw
}

// Issue platform (currently always GitHub)
type FundingSearchResponseItemsIssuePlatform string

const (
	FundingSearchResponseItemsIssuePlatformGitHub FundingSearchResponseItemsIssuePlatform = "github"
)

func (r FundingSearchResponseItemsIssuePlatform) IsKnown() bool {
	switch r {
	case FundingSearchResponseItemsIssuePlatformGitHub:
		return true
	}
	return false
}

// The repository that the issue is in
type FundingSearchResponseItemsIssueRepository struct {
	ID           string                                                `json:"id,required" format:"uuid"`
	IsPrivate    bool                                                  `json:"is_private,required"`
	Name         string                                                `json:"name,required"`
	Organization FundingSearchResponseItemsIssueRepositoryOrganization `json:"organization,required"`
	Platform     FundingSearchResponseItemsIssueRepositoryPlatform     `json:"platform,required"`
	// Settings for the repository profile
	ProfileSettings FundingSearchResponseItemsIssueRepositoryProfileSettings `json:"profile_settings,required,nullable"`
	Description     string                                                   `json:"description,nullable"`
	Homepage        string                                                   `json:"homepage,nullable"`
	License         string                                                   `json:"license,nullable"`
	Stars           int64                                                    `json:"stars,nullable"`
	JSON            fundingSearchResponseItemsIssueRepositoryJSON            `json:"-"`
}

// fundingSearchResponseItemsIssueRepositoryJSON contains the JSON metadata for the
// struct [FundingSearchResponseItemsIssueRepository]
type fundingSearchResponseItemsIssueRepositoryJSON struct {
	ID              apijson.Field
	IsPrivate       apijson.Field
	Name            apijson.Field
	Organization    apijson.Field
	Platform        apijson.Field
	ProfileSettings apijson.Field
	Description     apijson.Field
	Homepage        apijson.Field
	License         apijson.Field
	Stars           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *FundingSearchResponseItemsIssueRepository) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingSearchResponseItemsIssueRepositoryJSON) RawJSON() string {
	return r.raw
}

type FundingSearchResponseItemsIssueRepositoryOrganization struct {
	ID         string                                                        `json:"id,required" format:"uuid"`
	AvatarURL  string                                                        `json:"avatar_url,required"`
	IsPersonal bool                                                          `json:"is_personal,required"`
	Name       string                                                        `json:"name,required"`
	Platform   FundingSearchResponseItemsIssueRepositoryOrganizationPlatform `json:"platform,required"`
	Bio        string                                                        `json:"bio,nullable"`
	Blog       string                                                        `json:"blog,nullable"`
	Company    string                                                        `json:"company,nullable"`
	Email      string                                                        `json:"email,nullable"`
	Location   string                                                        `json:"location,nullable"`
	// The organization ID.
	OrganizationID  string                                                    `json:"organization_id,nullable" format:"uuid4"`
	PrettyName      string                                                    `json:"pretty_name,nullable"`
	TwitterUsername string                                                    `json:"twitter_username,nullable"`
	JSON            fundingSearchResponseItemsIssueRepositoryOrganizationJSON `json:"-"`
}

// fundingSearchResponseItemsIssueRepositoryOrganizationJSON contains the JSON
// metadata for the struct [FundingSearchResponseItemsIssueRepositoryOrganization]
type fundingSearchResponseItemsIssueRepositoryOrganizationJSON struct {
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

func (r *FundingSearchResponseItemsIssueRepositoryOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingSearchResponseItemsIssueRepositoryOrganizationJSON) RawJSON() string {
	return r.raw
}

type FundingSearchResponseItemsIssueRepositoryOrganizationPlatform string

const (
	FundingSearchResponseItemsIssueRepositoryOrganizationPlatformGitHub FundingSearchResponseItemsIssueRepositoryOrganizationPlatform = "github"
)

func (r FundingSearchResponseItemsIssueRepositoryOrganizationPlatform) IsKnown() bool {
	switch r {
	case FundingSearchResponseItemsIssueRepositoryOrganizationPlatformGitHub:
		return true
	}
	return false
}

type FundingSearchResponseItemsIssueRepositoryPlatform string

const (
	FundingSearchResponseItemsIssueRepositoryPlatformGitHub FundingSearchResponseItemsIssueRepositoryPlatform = "github"
)

func (r FundingSearchResponseItemsIssueRepositoryPlatform) IsKnown() bool {
	switch r {
	case FundingSearchResponseItemsIssueRepositoryPlatformGitHub:
		return true
	}
	return false
}

// Settings for the repository profile
type FundingSearchResponseItemsIssueRepositoryProfileSettings struct {
	// A URL to a cover image
	CoverImageURL string `json:"cover_image_url,nullable"`
	// A description of the repository
	Description string `json:"description,nullable"`
	// A list of featured organizations
	FeaturedOrganizations []string `json:"featured_organizations,nullable" format:"uuid4"`
	// A list of highlighted subscription tiers
	HighlightedSubscriptionTiers []string `json:"highlighted_subscription_tiers,nullable" format:"uuid4"`
	// A list of links related to the repository
	Links []string                                                     `json:"links,nullable" format:"uri"`
	JSON  fundingSearchResponseItemsIssueRepositoryProfileSettingsJSON `json:"-"`
}

// fundingSearchResponseItemsIssueRepositoryProfileSettingsJSON contains the JSON
// metadata for the struct
// [FundingSearchResponseItemsIssueRepositoryProfileSettings]
type fundingSearchResponseItemsIssueRepositoryProfileSettingsJSON struct {
	CoverImageURL                apijson.Field
	Description                  apijson.Field
	FeaturedOrganizations        apijson.Field
	HighlightedSubscriptionTiers apijson.Field
	Links                        apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *FundingSearchResponseItemsIssueRepositoryProfileSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingSearchResponseItemsIssueRepositoryProfileSettingsJSON) RawJSON() string {
	return r.raw
}

type FundingSearchResponseItemsIssueState string

const (
	FundingSearchResponseItemsIssueStateOpen   FundingSearchResponseItemsIssueState = "open"
	FundingSearchResponseItemsIssueStateClosed FundingSearchResponseItemsIssueState = "closed"
)

func (r FundingSearchResponseItemsIssueState) IsKnown() bool {
	switch r {
	case FundingSearchResponseItemsIssueStateOpen, FundingSearchResponseItemsIssueStateClosed:
		return true
	}
	return false
}

type FundingSearchResponseItemsIssueAssignee struct {
	ID        int64                                       `json:"id,required"`
	AvatarURL string                                      `json:"avatar_url,required" format:"uri"`
	HTMLURL   string                                      `json:"html_url,required" format:"uri"`
	Login     string                                      `json:"login,required"`
	JSON      fundingSearchResponseItemsIssueAssigneeJSON `json:"-"`
}

// fundingSearchResponseItemsIssueAssigneeJSON contains the JSON metadata for the
// struct [FundingSearchResponseItemsIssueAssignee]
type fundingSearchResponseItemsIssueAssigneeJSON struct {
	ID          apijson.Field
	AvatarURL   apijson.Field
	HTMLURL     apijson.Field
	Login       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FundingSearchResponseItemsIssueAssignee) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingSearchResponseItemsIssueAssigneeJSON) RawJSON() string {
	return r.raw
}

// GitHub author
type FundingSearchResponseItemsIssueAuthor struct {
	ID        int64                                     `json:"id,required"`
	AvatarURL string                                    `json:"avatar_url,required" format:"uri"`
	HTMLURL   string                                    `json:"html_url,required" format:"uri"`
	Login     string                                    `json:"login,required"`
	JSON      fundingSearchResponseItemsIssueAuthorJSON `json:"-"`
}

// fundingSearchResponseItemsIssueAuthorJSON contains the JSON metadata for the
// struct [FundingSearchResponseItemsIssueAuthor]
type fundingSearchResponseItemsIssueAuthorJSON struct {
	ID          apijson.Field
	AvatarURL   apijson.Field
	HTMLURL     apijson.Field
	Login       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FundingSearchResponseItemsIssueAuthor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingSearchResponseItemsIssueAuthorJSON) RawJSON() string {
	return r.raw
}

type FundingSearchResponseItemsIssueLabel struct {
	Color string                                   `json:"color,required"`
	Name  string                                   `json:"name,required"`
	JSON  fundingSearchResponseItemsIssueLabelJSON `json:"-"`
}

// fundingSearchResponseItemsIssueLabelJSON contains the JSON metadata for the
// struct [FundingSearchResponseItemsIssueLabel]
type fundingSearchResponseItemsIssueLabelJSON struct {
	Color       apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FundingSearchResponseItemsIssueLabel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingSearchResponseItemsIssueLabelJSON) RawJSON() string {
	return r.raw
}

// GitHub reactions
type FundingSearchResponseItemsIssueReactions struct {
	Confused   int64                                        `json:"confused,required"`
	Eyes       int64                                        `json:"eyes,required"`
	Heart      int64                                        `json:"heart,required"`
	Hooray     int64                                        `json:"hooray,required"`
	Laugh      int64                                        `json:"laugh,required"`
	MinusOne   int64                                        `json:"minus_one,required"`
	PlusOne    int64                                        `json:"plus_one,required"`
	Rocket     int64                                        `json:"rocket,required"`
	TotalCount int64                                        `json:"total_count,required"`
	JSON       fundingSearchResponseItemsIssueReactionsJSON `json:"-"`
}

// fundingSearchResponseItemsIssueReactionsJSON contains the JSON metadata for the
// struct [FundingSearchResponseItemsIssueReactions]
type fundingSearchResponseItemsIssueReactionsJSON struct {
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

func (r *FundingSearchResponseItemsIssueReactions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingSearchResponseItemsIssueReactionsJSON) RawJSON() string {
	return r.raw
}

type FundingSearchResponseItemsPledgesSummaries struct {
	PayDirectly     FundingSearchResponseItemsPledgesSummariesPayDirectly     `json:"pay_directly,required"`
	PayOnCompletion FundingSearchResponseItemsPledgesSummariesPayOnCompletion `json:"pay_on_completion,required"`
	PayUpfront      FundingSearchResponseItemsPledgesSummariesPayUpfront      `json:"pay_upfront,required"`
	JSON            fundingSearchResponseItemsPledgesSummariesJSON            `json:"-"`
}

// fundingSearchResponseItemsPledgesSummariesJSON contains the JSON metadata for
// the struct [FundingSearchResponseItemsPledgesSummaries]
type fundingSearchResponseItemsPledgesSummariesJSON struct {
	PayDirectly     apijson.Field
	PayOnCompletion apijson.Field
	PayUpfront      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *FundingSearchResponseItemsPledgesSummaries) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingSearchResponseItemsPledgesSummariesJSON) RawJSON() string {
	return r.raw
}

type FundingSearchResponseItemsPledgesSummariesPayDirectly struct {
	Pledgers []FundingSearchResponseItemsPledgesSummariesPayDirectlyPledger `json:"pledgers,required"`
	Total    FundingSearchResponseItemsPledgesSummariesPayDirectlyTotal     `json:"total,required"`
	JSON     fundingSearchResponseItemsPledgesSummariesPayDirectlyJSON      `json:"-"`
}

// fundingSearchResponseItemsPledgesSummariesPayDirectlyJSON contains the JSON
// metadata for the struct [FundingSearchResponseItemsPledgesSummariesPayDirectly]
type fundingSearchResponseItemsPledgesSummariesPayDirectlyJSON struct {
	Pledgers    apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FundingSearchResponseItemsPledgesSummariesPayDirectly) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingSearchResponseItemsPledgesSummariesPayDirectlyJSON) RawJSON() string {
	return r.raw
}

type FundingSearchResponseItemsPledgesSummariesPayDirectlyPledger struct {
	Name           string                                                           `json:"name,required"`
	AvatarURL      string                                                           `json:"avatar_url,nullable"`
	GitHubUsername string                                                           `json:"github_username,nullable"`
	JSON           fundingSearchResponseItemsPledgesSummariesPayDirectlyPledgerJSON `json:"-"`
}

// fundingSearchResponseItemsPledgesSummariesPayDirectlyPledgerJSON contains the
// JSON metadata for the struct
// [FundingSearchResponseItemsPledgesSummariesPayDirectlyPledger]
type fundingSearchResponseItemsPledgesSummariesPayDirectlyPledgerJSON struct {
	Name           apijson.Field
	AvatarURL      apijson.Field
	GitHubUsername apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *FundingSearchResponseItemsPledgesSummariesPayDirectlyPledger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingSearchResponseItemsPledgesSummariesPayDirectlyPledgerJSON) RawJSON() string {
	return r.raw
}

type FundingSearchResponseItemsPledgesSummariesPayDirectlyTotal struct {
	// Amount in the currencies smallest unit (cents if currency is USD)
	Amount int64 `json:"amount,required"`
	// Three letter currency code (eg: USD)
	Currency string                                                         `json:"currency,required"`
	JSON     fundingSearchResponseItemsPledgesSummariesPayDirectlyTotalJSON `json:"-"`
}

// fundingSearchResponseItemsPledgesSummariesPayDirectlyTotalJSON contains the JSON
// metadata for the struct
// [FundingSearchResponseItemsPledgesSummariesPayDirectlyTotal]
type fundingSearchResponseItemsPledgesSummariesPayDirectlyTotalJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FundingSearchResponseItemsPledgesSummariesPayDirectlyTotal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingSearchResponseItemsPledgesSummariesPayDirectlyTotalJSON) RawJSON() string {
	return r.raw
}

type FundingSearchResponseItemsPledgesSummariesPayOnCompletion struct {
	Pledgers []FundingSearchResponseItemsPledgesSummariesPayOnCompletionPledger `json:"pledgers,required"`
	Total    FundingSearchResponseItemsPledgesSummariesPayOnCompletionTotal     `json:"total,required"`
	JSON     fundingSearchResponseItemsPledgesSummariesPayOnCompletionJSON      `json:"-"`
}

// fundingSearchResponseItemsPledgesSummariesPayOnCompletionJSON contains the JSON
// metadata for the struct
// [FundingSearchResponseItemsPledgesSummariesPayOnCompletion]
type fundingSearchResponseItemsPledgesSummariesPayOnCompletionJSON struct {
	Pledgers    apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FundingSearchResponseItemsPledgesSummariesPayOnCompletion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingSearchResponseItemsPledgesSummariesPayOnCompletionJSON) RawJSON() string {
	return r.raw
}

type FundingSearchResponseItemsPledgesSummariesPayOnCompletionPledger struct {
	Name           string                                                               `json:"name,required"`
	AvatarURL      string                                                               `json:"avatar_url,nullable"`
	GitHubUsername string                                                               `json:"github_username,nullable"`
	JSON           fundingSearchResponseItemsPledgesSummariesPayOnCompletionPledgerJSON `json:"-"`
}

// fundingSearchResponseItemsPledgesSummariesPayOnCompletionPledgerJSON contains
// the JSON metadata for the struct
// [FundingSearchResponseItemsPledgesSummariesPayOnCompletionPledger]
type fundingSearchResponseItemsPledgesSummariesPayOnCompletionPledgerJSON struct {
	Name           apijson.Field
	AvatarURL      apijson.Field
	GitHubUsername apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *FundingSearchResponseItemsPledgesSummariesPayOnCompletionPledger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingSearchResponseItemsPledgesSummariesPayOnCompletionPledgerJSON) RawJSON() string {
	return r.raw
}

type FundingSearchResponseItemsPledgesSummariesPayOnCompletionTotal struct {
	// Amount in the currencies smallest unit (cents if currency is USD)
	Amount int64 `json:"amount,required"`
	// Three letter currency code (eg: USD)
	Currency string                                                             `json:"currency,required"`
	JSON     fundingSearchResponseItemsPledgesSummariesPayOnCompletionTotalJSON `json:"-"`
}

// fundingSearchResponseItemsPledgesSummariesPayOnCompletionTotalJSON contains the
// JSON metadata for the struct
// [FundingSearchResponseItemsPledgesSummariesPayOnCompletionTotal]
type fundingSearchResponseItemsPledgesSummariesPayOnCompletionTotalJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FundingSearchResponseItemsPledgesSummariesPayOnCompletionTotal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingSearchResponseItemsPledgesSummariesPayOnCompletionTotalJSON) RawJSON() string {
	return r.raw
}

type FundingSearchResponseItemsPledgesSummariesPayUpfront struct {
	Pledgers []FundingSearchResponseItemsPledgesSummariesPayUpfrontPledger `json:"pledgers,required"`
	Total    FundingSearchResponseItemsPledgesSummariesPayUpfrontTotal     `json:"total,required"`
	JSON     fundingSearchResponseItemsPledgesSummariesPayUpfrontJSON      `json:"-"`
}

// fundingSearchResponseItemsPledgesSummariesPayUpfrontJSON contains the JSON
// metadata for the struct [FundingSearchResponseItemsPledgesSummariesPayUpfront]
type fundingSearchResponseItemsPledgesSummariesPayUpfrontJSON struct {
	Pledgers    apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FundingSearchResponseItemsPledgesSummariesPayUpfront) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingSearchResponseItemsPledgesSummariesPayUpfrontJSON) RawJSON() string {
	return r.raw
}

type FundingSearchResponseItemsPledgesSummariesPayUpfrontPledger struct {
	Name           string                                                          `json:"name,required"`
	AvatarURL      string                                                          `json:"avatar_url,nullable"`
	GitHubUsername string                                                          `json:"github_username,nullable"`
	JSON           fundingSearchResponseItemsPledgesSummariesPayUpfrontPledgerJSON `json:"-"`
}

// fundingSearchResponseItemsPledgesSummariesPayUpfrontPledgerJSON contains the
// JSON metadata for the struct
// [FundingSearchResponseItemsPledgesSummariesPayUpfrontPledger]
type fundingSearchResponseItemsPledgesSummariesPayUpfrontPledgerJSON struct {
	Name           apijson.Field
	AvatarURL      apijson.Field
	GitHubUsername apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *FundingSearchResponseItemsPledgesSummariesPayUpfrontPledger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingSearchResponseItemsPledgesSummariesPayUpfrontPledgerJSON) RawJSON() string {
	return r.raw
}

type FundingSearchResponseItemsPledgesSummariesPayUpfrontTotal struct {
	// Amount in the currencies smallest unit (cents if currency is USD)
	Amount int64 `json:"amount,required"`
	// Three letter currency code (eg: USD)
	Currency string                                                        `json:"currency,required"`
	JSON     fundingSearchResponseItemsPledgesSummariesPayUpfrontTotalJSON `json:"-"`
}

// fundingSearchResponseItemsPledgesSummariesPayUpfrontTotalJSON contains the JSON
// metadata for the struct
// [FundingSearchResponseItemsPledgesSummariesPayUpfrontTotal]
type fundingSearchResponseItemsPledgesSummariesPayUpfrontTotalJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FundingSearchResponseItemsPledgesSummariesPayUpfrontTotal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingSearchResponseItemsPledgesSummariesPayUpfrontTotalJSON) RawJSON() string {
	return r.raw
}

type FundingSearchResponseItemsTotal struct {
	// Amount in the currencies smallest unit (cents if currency is USD)
	Amount int64 `json:"amount,required"`
	// Three letter currency code (eg: USD)
	Currency string                              `json:"currency,required"`
	JSON     fundingSearchResponseItemsTotalJSON `json:"-"`
}

// fundingSearchResponseItemsTotalJSON contains the JSON metadata for the struct
// [FundingSearchResponseItemsTotal]
type fundingSearchResponseItemsTotalJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FundingSearchResponseItemsTotal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingSearchResponseItemsTotalJSON) RawJSON() string {
	return r.raw
}

type FundingSearchResponseItemsFundingGoal struct {
	// Amount in the currencies smallest unit (cents if currency is USD)
	Amount int64 `json:"amount,required"`
	// Three letter currency code (eg: USD)
	Currency string                                    `json:"currency,required"`
	JSON     fundingSearchResponseItemsFundingGoalJSON `json:"-"`
}

// fundingSearchResponseItemsFundingGoalJSON contains the JSON metadata for the
// struct [FundingSearchResponseItemsFundingGoal]
type fundingSearchResponseItemsFundingGoalJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FundingSearchResponseItemsFundingGoal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingSearchResponseItemsFundingGoalJSON) RawJSON() string {
	return r.raw
}

type FundingLookupParams struct {
	IssueID param.Field[string] `query:"issue_id,required" format:"uuid"`
}

// URLQuery serializes [FundingLookupParams]'s query parameters as `url.Values`.
func (r FundingLookupParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FundingSearchParams struct {
	// The organization ID.
	OrganizationID param.Field[string] `query:"organization_id,required" format:"uuid4"`
	Badged         param.Field[bool]   `query:"badged"`
	Closed         param.Field[bool]   `query:"closed"`
	// Size of a page, defaults to 10. Maximum is 100.
	Limit param.Field[int64] `query:"limit"`
	// Page number, defaults to 1.
	Page  param.Field[int64]  `query:"page"`
	Query param.Field[string] `query:"query"`
	// Filter by repository name.
	RepositoryName param.Field[string] `query:"repository_name"`
	// Sorting criterion. Several criteria can be used simultaneously and will be
	// applied in order.
	Sorting param.Field[[]FundingSearchParamsSorting] `query:"sorting"`
}

// URLQuery serializes [FundingSearchParams]'s query parameters as `url.Values`.
func (r FundingSearchParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FundingSearchParamsSorting string

const (
	FundingSearchParamsSortingOldest             FundingSearchParamsSorting = "oldest"
	FundingSearchParamsSortingNewest             FundingSearchParamsSorting = "newest"
	FundingSearchParamsSortingMostFunded         FundingSearchParamsSorting = "most_funded"
	FundingSearchParamsSortingMostRecentlyFunded FundingSearchParamsSorting = "most_recently_funded"
	FundingSearchParamsSortingMostEngagement     FundingSearchParamsSorting = "most_engagement"
)

func (r FundingSearchParamsSorting) IsKnown() bool {
	switch r {
	case FundingSearchParamsSortingOldest, FundingSearchParamsSortingNewest, FundingSearchParamsSortingMostFunded, FundingSearchParamsSortingMostRecentlyFunded, FundingSearchParamsSortingMostEngagement:
		return true
	}
	return false
}
