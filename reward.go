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

// RewardService contains methods and other services that help with interacting
// with the polar API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRewardService] method instead.
type RewardService struct {
	Options []option.RequestOption
}

// NewRewardService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRewardService(opts ...option.RequestOption) (r *RewardService) {
	r = &RewardService{}
	r.Options = opts
	return
}

// Search rewards.
func (r *RewardService) Search(ctx context.Context, query RewardSearchParams, opts ...option.RequestOption) (res *RewardSearchResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/rewards/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get summary of rewards for resource.
func (r *RewardService) Summary(ctx context.Context, query RewardSummaryParams, opts ...option.RequestOption) (res *RewardSummaryResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/rewards/summary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RewardSearchResponse struct {
	Pagination RewardSearchResponsePagination `json:"pagination,required"`
	Items      []RewardSearchResponseItem     `json:"items"`
	JSON       rewardSearchResponseJSON       `json:"-"`
}

// rewardSearchResponseJSON contains the JSON metadata for the struct
// [RewardSearchResponse]
type rewardSearchResponseJSON struct {
	Pagination  apijson.Field
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RewardSearchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rewardSearchResponseJSON) RawJSON() string {
	return r.raw
}

type RewardSearchResponsePagination struct {
	MaxPage    int64                              `json:"max_page,required"`
	TotalCount int64                              `json:"total_count,required"`
	JSON       rewardSearchResponsePaginationJSON `json:"-"`
}

// rewardSearchResponsePaginationJSON contains the JSON metadata for the struct
// [RewardSearchResponsePagination]
type rewardSearchResponsePaginationJSON struct {
	MaxPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RewardSearchResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rewardSearchResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type RewardSearchResponseItem struct {
	Amount RewardSearchResponseItemsAmount `json:"amount,required"`
	// The pledge that the reward was split from
	Pledge RewardSearchResponseItemsPledge `json:"pledge,required"`
	State  RewardSearchResponseItemsState  `json:"state,required"`
	// The organization that received the reward (if any)
	Organization RewardSearchResponseItemsOrganization `json:"organization,nullable"`
	// If and when the reward was paid out.
	PaidAt time.Time `json:"paid_at,nullable" format:"date-time"`
	// The user that received the reward (if any)
	User RewardSearchResponseItemsUser `json:"user,nullable"`
	JSON rewardSearchResponseItemJSON  `json:"-"`
}

// rewardSearchResponseItemJSON contains the JSON metadata for the struct
// [RewardSearchResponseItem]
type rewardSearchResponseItemJSON struct {
	Amount       apijson.Field
	Pledge       apijson.Field
	State        apijson.Field
	Organization apijson.Field
	PaidAt       apijson.Field
	User         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RewardSearchResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rewardSearchResponseItemJSON) RawJSON() string {
	return r.raw
}

type RewardSearchResponseItemsAmount struct {
	// Amount in the currencies smallest unit (cents if currency is USD)
	Amount int64 `json:"amount,required"`
	// Three letter currency code (eg: USD)
	Currency string                              `json:"currency,required"`
	JSON     rewardSearchResponseItemsAmountJSON `json:"-"`
}

// rewardSearchResponseItemsAmountJSON contains the JSON metadata for the struct
// [RewardSearchResponseItemsAmount]
type rewardSearchResponseItemsAmountJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RewardSearchResponseItemsAmount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rewardSearchResponseItemsAmountJSON) RawJSON() string {
	return r.raw
}

// The pledge that the reward was split from
type RewardSearchResponseItemsPledge struct {
	// The ID of the object.
	ID string `json:"id,required" format:"uuid4"`
	// Amount pledged towards the issue
	Amount int64 `json:"amount,required"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	Currency  string    `json:"currency,required"`
	// The issue that the pledge was made towards
	Issue RewardSearchResponseItemsPledgeIssue `json:"issue,required"`
	// Current state of the pledge
	State RewardSearchResponseItemsPledgeState `json:"state,required"`
	// Type of pledge
	Type RewardSearchResponseItemsPledgeType `json:"type,required"`
	// If the currently authenticated subject can perform admin actions on behalf of
	// the receiver of the peldge
	AuthedCanAdminReceived bool `json:"authed_can_admin_received"`
	// If the currently authenticated subject can perform admin actions on behalf of
	// the maker of the peldge
	AuthedCanAdminSender bool `json:"authed_can_admin_sender"`
	// For pledges made by an organization, or on behalf of an organization. This is
	// the user that made the pledge. Only visible for members of said organization.
	CreatedBy RewardSearchResponseItemsPledgeCreatedBy `json:"created_by,nullable"`
	// URL of invoice for this pledge
	HostedInvoiceURL string `json:"hosted_invoice_url,nullable"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,nullable" format:"date-time"`
	// The user or organization that made this pledge
	Pledger RewardSearchResponseItemsPledgePledger `json:"pledger,nullable"`
	// If and when the pledge was refunded to the pledger
	RefundedAt time.Time `json:"refunded_at,nullable" format:"date-time"`
	// When the payout is scheduled to be made to the maintainers behind the issue.
	// Disputes must be made before this date.
	ScheduledPayoutAt time.Time                           `json:"scheduled_payout_at,nullable" format:"date-time"`
	JSON              rewardSearchResponseItemsPledgeJSON `json:"-"`
}

// rewardSearchResponseItemsPledgeJSON contains the JSON metadata for the struct
// [RewardSearchResponseItemsPledge]
type rewardSearchResponseItemsPledgeJSON struct {
	ID                     apijson.Field
	Amount                 apijson.Field
	CreatedAt              apijson.Field
	Currency               apijson.Field
	Issue                  apijson.Field
	State                  apijson.Field
	Type                   apijson.Field
	AuthedCanAdminReceived apijson.Field
	AuthedCanAdminSender   apijson.Field
	CreatedBy              apijson.Field
	HostedInvoiceURL       apijson.Field
	ModifiedAt             apijson.Field
	Pledger                apijson.Field
	RefundedAt             apijson.Field
	ScheduledPayoutAt      apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *RewardSearchResponseItemsPledge) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rewardSearchResponseItemsPledgeJSON) RawJSON() string {
	return r.raw
}

// The issue that the pledge was made towards
type RewardSearchResponseItemsPledgeIssue struct {
	ID             string                                      `json:"id,required" format:"uuid"`
	Funding        RewardSearchResponseItemsPledgeIssueFunding `json:"funding,required"`
	IssueCreatedAt time.Time                                   `json:"issue_created_at,required" format:"date-time"`
	// If a maintainer needs to mark this issue as solved
	NeedsConfirmationSolved bool `json:"needs_confirmation_solved,required"`
	// GitHub #number
	Number int64 `json:"number,required"`
	// Issue platform (currently always GitHub)
	Platform RewardSearchResponseItemsPledgeIssuePlatform `json:"platform,required"`
	// If this issue currently has the Polar badge SVG embedded
	PledgeBadgeCurrentlyEmbedded bool `json:"pledge_badge_currently_embedded,required"`
	// The repository that the issue is in
	Repository RewardSearchResponseItemsPledgeIssueRepository `json:"repository,required"`
	State      RewardSearchResponseItemsPledgeIssueState      `json:"state,required"`
	// GitHub issue title
	Title string `json:"title,required"`
	// GitHub assignees
	Assignees []RewardSearchResponseItemsPledgeIssueAssignee `json:"assignees,nullable"`
	// GitHub author
	Author RewardSearchResponseItemsPledgeIssueAuthor `json:"author,nullable"`
	// Optional custom badge SVG promotional content
	BadgeCustomContent string `json:"badge_custom_content,nullable"`
	// GitHub issue body
	Body string `json:"body,nullable"`
	// Number of GitHub comments made on the issue
	Comments int64 `json:"comments,nullable"`
	// If this issue has been marked as confirmed solved through Polar
	ConfirmedSolvedAt time.Time                                   `json:"confirmed_solved_at,nullable" format:"date-time"`
	IssueClosedAt     time.Time                                   `json:"issue_closed_at,nullable" format:"date-time"`
	IssueModifiedAt   time.Time                                   `json:"issue_modified_at,nullable" format:"date-time"`
	Labels            []RewardSearchResponseItemsPledgeIssueLabel `json:"labels"`
	// GitHub reactions
	Reactions RewardSearchResponseItemsPledgeIssueReactions `json:"reactions,nullable"`
	// Share of rewrads that will be rewarded to contributors of this issue. A number
	// between 0 and 100 (inclusive).
	UpfrontSplitToContributors int64                                    `json:"upfront_split_to_contributors,nullable"`
	JSON                       rewardSearchResponseItemsPledgeIssueJSON `json:"-"`
}

// rewardSearchResponseItemsPledgeIssueJSON contains the JSON metadata for the
// struct [RewardSearchResponseItemsPledgeIssue]
type rewardSearchResponseItemsPledgeIssueJSON struct {
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

func (r *RewardSearchResponseItemsPledgeIssue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rewardSearchResponseItemsPledgeIssueJSON) RawJSON() string {
	return r.raw
}

type RewardSearchResponseItemsPledgeIssueFunding struct {
	FundingGoal RewardSearchResponseItemsPledgeIssueFundingFundingGoal `json:"funding_goal,nullable"`
	// Sum of pledges to this isuse (including currently open pledges and pledges that
	// have been paid out). Always in USD.
	PledgesSum RewardSearchResponseItemsPledgeIssueFundingPledgesSum `json:"pledges_sum,nullable"`
	JSON       rewardSearchResponseItemsPledgeIssueFundingJSON       `json:"-"`
}

// rewardSearchResponseItemsPledgeIssueFundingJSON contains the JSON metadata for
// the struct [RewardSearchResponseItemsPledgeIssueFunding]
type rewardSearchResponseItemsPledgeIssueFundingJSON struct {
	FundingGoal apijson.Field
	PledgesSum  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RewardSearchResponseItemsPledgeIssueFunding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rewardSearchResponseItemsPledgeIssueFundingJSON) RawJSON() string {
	return r.raw
}

type RewardSearchResponseItemsPledgeIssueFundingFundingGoal struct {
	// Amount in the currencies smallest unit (cents if currency is USD)
	Amount int64 `json:"amount,required"`
	// Three letter currency code (eg: USD)
	Currency string                                                     `json:"currency,required"`
	JSON     rewardSearchResponseItemsPledgeIssueFundingFundingGoalJSON `json:"-"`
}

// rewardSearchResponseItemsPledgeIssueFundingFundingGoalJSON contains the JSON
// metadata for the struct [RewardSearchResponseItemsPledgeIssueFundingFundingGoal]
type rewardSearchResponseItemsPledgeIssueFundingFundingGoalJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RewardSearchResponseItemsPledgeIssueFundingFundingGoal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rewardSearchResponseItemsPledgeIssueFundingFundingGoalJSON) RawJSON() string {
	return r.raw
}

// Sum of pledges to this isuse (including currently open pledges and pledges that
// have been paid out). Always in USD.
type RewardSearchResponseItemsPledgeIssueFundingPledgesSum struct {
	// Amount in the currencies smallest unit (cents if currency is USD)
	Amount int64 `json:"amount,required"`
	// Three letter currency code (eg: USD)
	Currency string                                                    `json:"currency,required"`
	JSON     rewardSearchResponseItemsPledgeIssueFundingPledgesSumJSON `json:"-"`
}

// rewardSearchResponseItemsPledgeIssueFundingPledgesSumJSON contains the JSON
// metadata for the struct [RewardSearchResponseItemsPledgeIssueFundingPledgesSum]
type rewardSearchResponseItemsPledgeIssueFundingPledgesSumJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RewardSearchResponseItemsPledgeIssueFundingPledgesSum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rewardSearchResponseItemsPledgeIssueFundingPledgesSumJSON) RawJSON() string {
	return r.raw
}

// Issue platform (currently always GitHub)
type RewardSearchResponseItemsPledgeIssuePlatform string

const (
	RewardSearchResponseItemsPledgeIssuePlatformGitHub RewardSearchResponseItemsPledgeIssuePlatform = "github"
)

func (r RewardSearchResponseItemsPledgeIssuePlatform) IsKnown() bool {
	switch r {
	case RewardSearchResponseItemsPledgeIssuePlatformGitHub:
		return true
	}
	return false
}

// The repository that the issue is in
type RewardSearchResponseItemsPledgeIssueRepository struct {
	ID           string                                                     `json:"id,required" format:"uuid"`
	IsPrivate    bool                                                       `json:"is_private,required"`
	Name         string                                                     `json:"name,required"`
	Organization RewardSearchResponseItemsPledgeIssueRepositoryOrganization `json:"organization,required"`
	Platform     RewardSearchResponseItemsPledgeIssueRepositoryPlatform     `json:"platform,required"`
	// Settings for the repository profile
	ProfileSettings RewardSearchResponseItemsPledgeIssueRepositoryProfileSettings `json:"profile_settings,required,nullable"`
	Description     string                                                        `json:"description,nullable"`
	Homepage        string                                                        `json:"homepage,nullable"`
	License         string                                                        `json:"license,nullable"`
	Stars           int64                                                         `json:"stars,nullable"`
	JSON            rewardSearchResponseItemsPledgeIssueRepositoryJSON            `json:"-"`
}

// rewardSearchResponseItemsPledgeIssueRepositoryJSON contains the JSON metadata
// for the struct [RewardSearchResponseItemsPledgeIssueRepository]
type rewardSearchResponseItemsPledgeIssueRepositoryJSON struct {
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

func (r *RewardSearchResponseItemsPledgeIssueRepository) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rewardSearchResponseItemsPledgeIssueRepositoryJSON) RawJSON() string {
	return r.raw
}

type RewardSearchResponseItemsPledgeIssueRepositoryOrganization struct {
	ID         string                                                             `json:"id,required" format:"uuid"`
	AvatarURL  string                                                             `json:"avatar_url,required"`
	IsPersonal bool                                                               `json:"is_personal,required"`
	Name       string                                                             `json:"name,required"`
	Platform   RewardSearchResponseItemsPledgeIssueRepositoryOrganizationPlatform `json:"platform,required"`
	Bio        string                                                             `json:"bio,nullable"`
	Blog       string                                                             `json:"blog,nullable"`
	Company    string                                                             `json:"company,nullable"`
	Email      string                                                             `json:"email,nullable"`
	Location   string                                                             `json:"location,nullable"`
	// The organization ID.
	OrganizationID  string                                                         `json:"organization_id,nullable" format:"uuid4"`
	PrettyName      string                                                         `json:"pretty_name,nullable"`
	TwitterUsername string                                                         `json:"twitter_username,nullable"`
	JSON            rewardSearchResponseItemsPledgeIssueRepositoryOrganizationJSON `json:"-"`
}

// rewardSearchResponseItemsPledgeIssueRepositoryOrganizationJSON contains the JSON
// metadata for the struct
// [RewardSearchResponseItemsPledgeIssueRepositoryOrganization]
type rewardSearchResponseItemsPledgeIssueRepositoryOrganizationJSON struct {
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

func (r *RewardSearchResponseItemsPledgeIssueRepositoryOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rewardSearchResponseItemsPledgeIssueRepositoryOrganizationJSON) RawJSON() string {
	return r.raw
}

type RewardSearchResponseItemsPledgeIssueRepositoryOrganizationPlatform string

const (
	RewardSearchResponseItemsPledgeIssueRepositoryOrganizationPlatformGitHub RewardSearchResponseItemsPledgeIssueRepositoryOrganizationPlatform = "github"
)

func (r RewardSearchResponseItemsPledgeIssueRepositoryOrganizationPlatform) IsKnown() bool {
	switch r {
	case RewardSearchResponseItemsPledgeIssueRepositoryOrganizationPlatformGitHub:
		return true
	}
	return false
}

type RewardSearchResponseItemsPledgeIssueRepositoryPlatform string

const (
	RewardSearchResponseItemsPledgeIssueRepositoryPlatformGitHub RewardSearchResponseItemsPledgeIssueRepositoryPlatform = "github"
)

func (r RewardSearchResponseItemsPledgeIssueRepositoryPlatform) IsKnown() bool {
	switch r {
	case RewardSearchResponseItemsPledgeIssueRepositoryPlatformGitHub:
		return true
	}
	return false
}

// Settings for the repository profile
type RewardSearchResponseItemsPledgeIssueRepositoryProfileSettings struct {
	// A URL to a cover image
	CoverImageURL string `json:"cover_image_url,nullable"`
	// A description of the repository
	Description string `json:"description,nullable"`
	// A list of featured organizations
	FeaturedOrganizations []string `json:"featured_organizations,nullable" format:"uuid4"`
	// A list of highlighted subscription tiers
	HighlightedSubscriptionTiers []string `json:"highlighted_subscription_tiers,nullable" format:"uuid4"`
	// A list of links related to the repository
	Links []string                                                          `json:"links,nullable" format:"uri"`
	JSON  rewardSearchResponseItemsPledgeIssueRepositoryProfileSettingsJSON `json:"-"`
}

// rewardSearchResponseItemsPledgeIssueRepositoryProfileSettingsJSON contains the
// JSON metadata for the struct
// [RewardSearchResponseItemsPledgeIssueRepositoryProfileSettings]
type rewardSearchResponseItemsPledgeIssueRepositoryProfileSettingsJSON struct {
	CoverImageURL                apijson.Field
	Description                  apijson.Field
	FeaturedOrganizations        apijson.Field
	HighlightedSubscriptionTiers apijson.Field
	Links                        apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *RewardSearchResponseItemsPledgeIssueRepositoryProfileSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rewardSearchResponseItemsPledgeIssueRepositoryProfileSettingsJSON) RawJSON() string {
	return r.raw
}

type RewardSearchResponseItemsPledgeIssueState string

const (
	RewardSearchResponseItemsPledgeIssueStateOpen   RewardSearchResponseItemsPledgeIssueState = "open"
	RewardSearchResponseItemsPledgeIssueStateClosed RewardSearchResponseItemsPledgeIssueState = "closed"
)

func (r RewardSearchResponseItemsPledgeIssueState) IsKnown() bool {
	switch r {
	case RewardSearchResponseItemsPledgeIssueStateOpen, RewardSearchResponseItemsPledgeIssueStateClosed:
		return true
	}
	return false
}

type RewardSearchResponseItemsPledgeIssueAssignee struct {
	ID        int64                                            `json:"id,required"`
	AvatarURL string                                           `json:"avatar_url,required" format:"uri"`
	HTMLURL   string                                           `json:"html_url,required" format:"uri"`
	Login     string                                           `json:"login,required"`
	JSON      rewardSearchResponseItemsPledgeIssueAssigneeJSON `json:"-"`
}

// rewardSearchResponseItemsPledgeIssueAssigneeJSON contains the JSON metadata for
// the struct [RewardSearchResponseItemsPledgeIssueAssignee]
type rewardSearchResponseItemsPledgeIssueAssigneeJSON struct {
	ID          apijson.Field
	AvatarURL   apijson.Field
	HTMLURL     apijson.Field
	Login       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RewardSearchResponseItemsPledgeIssueAssignee) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rewardSearchResponseItemsPledgeIssueAssigneeJSON) RawJSON() string {
	return r.raw
}

// GitHub author
type RewardSearchResponseItemsPledgeIssueAuthor struct {
	ID        int64                                          `json:"id,required"`
	AvatarURL string                                         `json:"avatar_url,required" format:"uri"`
	HTMLURL   string                                         `json:"html_url,required" format:"uri"`
	Login     string                                         `json:"login,required"`
	JSON      rewardSearchResponseItemsPledgeIssueAuthorJSON `json:"-"`
}

// rewardSearchResponseItemsPledgeIssueAuthorJSON contains the JSON metadata for
// the struct [RewardSearchResponseItemsPledgeIssueAuthor]
type rewardSearchResponseItemsPledgeIssueAuthorJSON struct {
	ID          apijson.Field
	AvatarURL   apijson.Field
	HTMLURL     apijson.Field
	Login       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RewardSearchResponseItemsPledgeIssueAuthor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rewardSearchResponseItemsPledgeIssueAuthorJSON) RawJSON() string {
	return r.raw
}

type RewardSearchResponseItemsPledgeIssueLabel struct {
	Color string                                        `json:"color,required"`
	Name  string                                        `json:"name,required"`
	JSON  rewardSearchResponseItemsPledgeIssueLabelJSON `json:"-"`
}

// rewardSearchResponseItemsPledgeIssueLabelJSON contains the JSON metadata for the
// struct [RewardSearchResponseItemsPledgeIssueLabel]
type rewardSearchResponseItemsPledgeIssueLabelJSON struct {
	Color       apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RewardSearchResponseItemsPledgeIssueLabel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rewardSearchResponseItemsPledgeIssueLabelJSON) RawJSON() string {
	return r.raw
}

// GitHub reactions
type RewardSearchResponseItemsPledgeIssueReactions struct {
	Confused   int64                                             `json:"confused,required"`
	Eyes       int64                                             `json:"eyes,required"`
	Heart      int64                                             `json:"heart,required"`
	Hooray     int64                                             `json:"hooray,required"`
	Laugh      int64                                             `json:"laugh,required"`
	MinusOne   int64                                             `json:"minus_one,required"`
	PlusOne    int64                                             `json:"plus_one,required"`
	Rocket     int64                                             `json:"rocket,required"`
	TotalCount int64                                             `json:"total_count,required"`
	JSON       rewardSearchResponseItemsPledgeIssueReactionsJSON `json:"-"`
}

// rewardSearchResponseItemsPledgeIssueReactionsJSON contains the JSON metadata for
// the struct [RewardSearchResponseItemsPledgeIssueReactions]
type rewardSearchResponseItemsPledgeIssueReactionsJSON struct {
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

func (r *RewardSearchResponseItemsPledgeIssueReactions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rewardSearchResponseItemsPledgeIssueReactionsJSON) RawJSON() string {
	return r.raw
}

// Current state of the pledge
type RewardSearchResponseItemsPledgeState string

const (
	RewardSearchResponseItemsPledgeStateInitiated      RewardSearchResponseItemsPledgeState = "initiated"
	RewardSearchResponseItemsPledgeStateCreated        RewardSearchResponseItemsPledgeState = "created"
	RewardSearchResponseItemsPledgeStatePending        RewardSearchResponseItemsPledgeState = "pending"
	RewardSearchResponseItemsPledgeStateRefunded       RewardSearchResponseItemsPledgeState = "refunded"
	RewardSearchResponseItemsPledgeStateDisputed       RewardSearchResponseItemsPledgeState = "disputed"
	RewardSearchResponseItemsPledgeStateChargeDisputed RewardSearchResponseItemsPledgeState = "charge_disputed"
	RewardSearchResponseItemsPledgeStateCancelled      RewardSearchResponseItemsPledgeState = "cancelled"
)

func (r RewardSearchResponseItemsPledgeState) IsKnown() bool {
	switch r {
	case RewardSearchResponseItemsPledgeStateInitiated, RewardSearchResponseItemsPledgeStateCreated, RewardSearchResponseItemsPledgeStatePending, RewardSearchResponseItemsPledgeStateRefunded, RewardSearchResponseItemsPledgeStateDisputed, RewardSearchResponseItemsPledgeStateChargeDisputed, RewardSearchResponseItemsPledgeStateCancelled:
		return true
	}
	return false
}

// Type of pledge
type RewardSearchResponseItemsPledgeType string

const (
	RewardSearchResponseItemsPledgeTypePayUpfront      RewardSearchResponseItemsPledgeType = "pay_upfront"
	RewardSearchResponseItemsPledgeTypePayOnCompletion RewardSearchResponseItemsPledgeType = "pay_on_completion"
	RewardSearchResponseItemsPledgeTypePayDirectly     RewardSearchResponseItemsPledgeType = "pay_directly"
)

func (r RewardSearchResponseItemsPledgeType) IsKnown() bool {
	switch r {
	case RewardSearchResponseItemsPledgeTypePayUpfront, RewardSearchResponseItemsPledgeTypePayOnCompletion, RewardSearchResponseItemsPledgeTypePayDirectly:
		return true
	}
	return false
}

// For pledges made by an organization, or on behalf of an organization. This is
// the user that made the pledge. Only visible for members of said organization.
type RewardSearchResponseItemsPledgeCreatedBy struct {
	Name           string                                       `json:"name,required"`
	AvatarURL      string                                       `json:"avatar_url,nullable"`
	GitHubUsername string                                       `json:"github_username,nullable"`
	JSON           rewardSearchResponseItemsPledgeCreatedByJSON `json:"-"`
}

// rewardSearchResponseItemsPledgeCreatedByJSON contains the JSON metadata for the
// struct [RewardSearchResponseItemsPledgeCreatedBy]
type rewardSearchResponseItemsPledgeCreatedByJSON struct {
	Name           apijson.Field
	AvatarURL      apijson.Field
	GitHubUsername apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RewardSearchResponseItemsPledgeCreatedBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rewardSearchResponseItemsPledgeCreatedByJSON) RawJSON() string {
	return r.raw
}

// The user or organization that made this pledge
type RewardSearchResponseItemsPledgePledger struct {
	Name           string                                     `json:"name,required"`
	AvatarURL      string                                     `json:"avatar_url,nullable"`
	GitHubUsername string                                     `json:"github_username,nullable"`
	JSON           rewardSearchResponseItemsPledgePledgerJSON `json:"-"`
}

// rewardSearchResponseItemsPledgePledgerJSON contains the JSON metadata for the
// struct [RewardSearchResponseItemsPledgePledger]
type rewardSearchResponseItemsPledgePledgerJSON struct {
	Name           apijson.Field
	AvatarURL      apijson.Field
	GitHubUsername apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RewardSearchResponseItemsPledgePledger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rewardSearchResponseItemsPledgePledgerJSON) RawJSON() string {
	return r.raw
}

type RewardSearchResponseItemsState string

const (
	RewardSearchResponseItemsStatePending RewardSearchResponseItemsState = "pending"
	RewardSearchResponseItemsStatePaid    RewardSearchResponseItemsState = "paid"
)

func (r RewardSearchResponseItemsState) IsKnown() bool {
	switch r {
	case RewardSearchResponseItemsStatePending, RewardSearchResponseItemsStatePaid:
		return true
	}
	return false
}

// The organization that received the reward (if any)
type RewardSearchResponseItemsOrganization struct {
	// The organization ID.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// If this organizations accepts donations
	DonationsEnabled bool `json:"donations_enabled,required"`
	// Settings for the organization features
	FeatureSettings       RewardSearchResponseItemsOrganizationFeatureSettings `json:"feature_settings,required,nullable"`
	Name                  string                                               `json:"name,required"`
	PledgeBadgeShowAmount bool                                                 `json:"pledge_badge_show_amount,required"`
	PledgeMinimumAmount   int64                                                `json:"pledge_minimum_amount,required"`
	// Settings for the organization profile
	ProfileSettings                   RewardSearchResponseItemsOrganizationProfileSettings `json:"profile_settings,required,nullable"`
	Slug                              string                                               `json:"slug,required"`
	AvatarURL                         string                                               `json:"avatar_url,nullable"`
	Bio                               string                                               `json:"bio,nullable"`
	Blog                              string                                               `json:"blog,nullable"`
	Company                           string                                               `json:"company,nullable"`
	DefaultUpfrontSplitToContributors int64                                                `json:"default_upfront_split_to_contributors,nullable"`
	Email                             string                                               `json:"email,nullable"`
	Location                          string                                               `json:"location,nullable"`
	// Last modification timestamp of the object.
	ModifiedAt      time.Time                                 `json:"modified_at,nullable" format:"date-time"`
	TwitterUsername string                                    `json:"twitter_username,nullable"`
	JSON            rewardSearchResponseItemsOrganizationJSON `json:"-"`
}

// rewardSearchResponseItemsOrganizationJSON contains the JSON metadata for the
// struct [RewardSearchResponseItemsOrganization]
type rewardSearchResponseItemsOrganizationJSON struct {
	ID                                apijson.Field
	CreatedAt                         apijson.Field
	DonationsEnabled                  apijson.Field
	FeatureSettings                   apijson.Field
	Name                              apijson.Field
	PledgeBadgeShowAmount             apijson.Field
	PledgeMinimumAmount               apijson.Field
	ProfileSettings                   apijson.Field
	Slug                              apijson.Field
	AvatarURL                         apijson.Field
	Bio                               apijson.Field
	Blog                              apijson.Field
	Company                           apijson.Field
	DefaultUpfrontSplitToContributors apijson.Field
	Email                             apijson.Field
	Location                          apijson.Field
	ModifiedAt                        apijson.Field
	TwitterUsername                   apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
}

func (r *RewardSearchResponseItemsOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rewardSearchResponseItemsOrganizationJSON) RawJSON() string {
	return r.raw
}

// Settings for the organization features
type RewardSearchResponseItemsOrganizationFeatureSettings struct {
	// If this organization has articles enabled
	ArticlesEnabled bool `json:"articles_enabled"`
	// If this organization has issue funding enabled
	IssueFundingEnabled bool `json:"issue_funding_enabled"`
	// If this organization has subscriptions enabled
	SubscriptionsEnabled bool                                                     `json:"subscriptions_enabled"`
	JSON                 rewardSearchResponseItemsOrganizationFeatureSettingsJSON `json:"-"`
}

// rewardSearchResponseItemsOrganizationFeatureSettingsJSON contains the JSON
// metadata for the struct [RewardSearchResponseItemsOrganizationFeatureSettings]
type rewardSearchResponseItemsOrganizationFeatureSettingsJSON struct {
	ArticlesEnabled      apijson.Field
	IssueFundingEnabled  apijson.Field
	SubscriptionsEnabled apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *RewardSearchResponseItemsOrganizationFeatureSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rewardSearchResponseItemsOrganizationFeatureSettingsJSON) RawJSON() string {
	return r.raw
}

// Settings for the organization profile
type RewardSearchResponseItemsOrganizationProfileSettings struct {
	// A description of the organization
	Description string `json:"description,nullable"`
	// A list of featured organizations
	FeaturedOrganizations []string `json:"featured_organizations,nullable" format:"uuid4"`
	// A list of featured projects
	FeaturedProjects []string `json:"featured_projects,nullable" format:"uuid4"`
	// A list of links associated with the organization
	Links []string `json:"links,nullable" format:"uri"`
	// Subscription promotion settings
	Subscribe RewardSearchResponseItemsOrganizationProfileSettingsSubscribe `json:"subscribe,nullable"`
	JSON      rewardSearchResponseItemsOrganizationProfileSettingsJSON      `json:"-"`
}

// rewardSearchResponseItemsOrganizationProfileSettingsJSON contains the JSON
// metadata for the struct [RewardSearchResponseItemsOrganizationProfileSettings]
type rewardSearchResponseItemsOrganizationProfileSettingsJSON struct {
	Description           apijson.Field
	FeaturedOrganizations apijson.Field
	FeaturedProjects      apijson.Field
	Links                 apijson.Field
	Subscribe             apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *RewardSearchResponseItemsOrganizationProfileSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rewardSearchResponseItemsOrganizationProfileSettingsJSON) RawJSON() string {
	return r.raw
}

// Subscription promotion settings
type RewardSearchResponseItemsOrganizationProfileSettingsSubscribe struct {
	// Include free subscribers in total count
	CountFree bool `json:"count_free"`
	// Promote email subscription (free)
	Promote bool `json:"promote"`
	// Show subscription count publicly
	ShowCount bool                                                              `json:"show_count"`
	JSON      rewardSearchResponseItemsOrganizationProfileSettingsSubscribeJSON `json:"-"`
}

// rewardSearchResponseItemsOrganizationProfileSettingsSubscribeJSON contains the
// JSON metadata for the struct
// [RewardSearchResponseItemsOrganizationProfileSettingsSubscribe]
type rewardSearchResponseItemsOrganizationProfileSettingsSubscribeJSON struct {
	CountFree   apijson.Field
	Promote     apijson.Field
	ShowCount   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RewardSearchResponseItemsOrganizationProfileSettingsSubscribe) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rewardSearchResponseItemsOrganizationProfileSettingsSubscribeJSON) RawJSON() string {
	return r.raw
}

// The user that received the reward (if any)
type RewardSearchResponseItemsUser struct {
	AvatarURL string                            `json:"avatar_url,required"`
	Username  string                            `json:"username,required"`
	JSON      rewardSearchResponseItemsUserJSON `json:"-"`
}

// rewardSearchResponseItemsUserJSON contains the JSON metadata for the struct
// [RewardSearchResponseItemsUser]
type rewardSearchResponseItemsUserJSON struct {
	AvatarURL   apijson.Field
	Username    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RewardSearchResponseItemsUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rewardSearchResponseItemsUserJSON) RawJSON() string {
	return r.raw
}

type RewardSummaryResponse struct {
	Receivers []RewardSummaryResponseReceiver `json:"receivers,required"`
	JSON      rewardSummaryResponseJSON       `json:"-"`
}

// rewardSummaryResponseJSON contains the JSON metadata for the struct
// [RewardSummaryResponse]
type rewardSummaryResponseJSON struct {
	Receivers   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RewardSummaryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rewardSummaryResponseJSON) RawJSON() string {
	return r.raw
}

type RewardSummaryResponseReceiver struct {
	Name      string                            `json:"name,required"`
	AvatarURL string                            `json:"avatar_url,nullable"`
	JSON      rewardSummaryResponseReceiverJSON `json:"-"`
}

// rewardSummaryResponseReceiverJSON contains the JSON metadata for the struct
// [RewardSummaryResponseReceiver]
type rewardSummaryResponseReceiverJSON struct {
	Name        apijson.Field
	AvatarURL   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RewardSummaryResponseReceiver) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rewardSummaryResponseReceiverJSON) RawJSON() string {
	return r.raw
}

type RewardSearchParams struct {
	// Search rewards for pledges in this organization.
	PledgesToOrganization param.Field[string] `query:"pledges_to_organization" format:"uuid"`
	// Search rewards to organization.
	RewardsToOrg param.Field[string] `query:"rewards_to_org" format:"uuid"`
	// Search rewards to user.
	RewardsToUser param.Field[string] `query:"rewards_to_user" format:"uuid"`
}

// URLQuery serializes [RewardSearchParams]'s query parameters as `url.Values`.
func (r RewardSearchParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RewardSummaryParams struct {
	IssueID param.Field[string] `query:"issue_id,required" format:"uuid"`
}

// URLQuery serializes [RewardSummaryParams]'s query parameters as `url.Values`.
func (r RewardSummaryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
