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

// PledgeService contains methods and other services that help with interacting
// with the polar API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPledgeService] method instead.
type PledgeService struct {
	Options []option.RequestOption
	Summary *PledgeSummaryService
}

// NewPledgeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPledgeService(opts ...option.RequestOption) (r *PledgeService) {
	r = &PledgeService{}
	r.Options = opts
	r.Summary = NewPledgeSummaryService(opts...)
	return
}

// Get a pledge. Requires authentication.
func (r *PledgeService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *PledgeGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/pledges/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Search pledges. Requires authentication. The user can only read pledges that
// they have made (personally or via an organization) or received (to organizations
// that they are a member of).
func (r *PledgeService) Search(ctx context.Context, query PledgeSearchParams, opts ...option.RequestOption) (res *PledgeSearchResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/pledges/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get current user spending in the current period. Used together with spending
// limits.
func (r *PledgeService) Spending(ctx context.Context, query PledgeSpendingParams, opts ...option.RequestOption) (res *PledgeSpendingResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/pledges/spending"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type PledgeGetResponse struct {
	// The ID of the object.
	ID string `json:"id,required" format:"uuid4"`
	// Amount pledged towards the issue
	Amount int64 `json:"amount,required"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	Currency  string    `json:"currency,required"`
	// The issue that the pledge was made towards
	Issue PledgeGetResponseIssue `json:"issue,required"`
	// Current state of the pledge
	State PledgeGetResponseState `json:"state,required"`
	// Type of pledge
	Type PledgeGetResponseType `json:"type,required"`
	// If the currently authenticated subject can perform admin actions on behalf of
	// the receiver of the peldge
	AuthedCanAdminReceived bool `json:"authed_can_admin_received"`
	// If the currently authenticated subject can perform admin actions on behalf of
	// the maker of the peldge
	AuthedCanAdminSender bool `json:"authed_can_admin_sender"`
	// For pledges made by an organization, or on behalf of an organization. This is
	// the user that made the pledge. Only visible for members of said organization.
	CreatedBy PledgeGetResponseCreatedBy `json:"created_by,nullable"`
	// URL of invoice for this pledge
	HostedInvoiceURL string `json:"hosted_invoice_url,nullable"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,nullable" format:"date-time"`
	// The user or organization that made this pledge
	Pledger PledgeGetResponsePledger `json:"pledger,nullable"`
	// If and when the pledge was refunded to the pledger
	RefundedAt time.Time `json:"refunded_at,nullable" format:"date-time"`
	// When the payout is scheduled to be made to the maintainers behind the issue.
	// Disputes must be made before this date.
	ScheduledPayoutAt time.Time             `json:"scheduled_payout_at,nullable" format:"date-time"`
	JSON              pledgeGetResponseJSON `json:"-"`
}

// pledgeGetResponseJSON contains the JSON metadata for the struct
// [PledgeGetResponse]
type pledgeGetResponseJSON struct {
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

func (r *PledgeGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pledgeGetResponseJSON) RawJSON() string {
	return r.raw
}

// The issue that the pledge was made towards
type PledgeGetResponseIssue struct {
	ID             string                        `json:"id,required" format:"uuid"`
	Funding        PledgeGetResponseIssueFunding `json:"funding,required"`
	IssueCreatedAt time.Time                     `json:"issue_created_at,required" format:"date-time"`
	// If a maintainer needs to mark this issue as solved
	NeedsConfirmationSolved bool `json:"needs_confirmation_solved,required"`
	// GitHub #number
	Number int64 `json:"number,required"`
	// Issue platform (currently always GitHub)
	Platform PledgeGetResponseIssuePlatform `json:"platform,required"`
	// If this issue currently has the Polar badge SVG embedded
	PledgeBadgeCurrentlyEmbedded bool `json:"pledge_badge_currently_embedded,required"`
	// The repository that the issue is in
	Repository PledgeGetResponseIssueRepository `json:"repository,required"`
	State      PledgeGetResponseIssueState      `json:"state,required"`
	// GitHub issue title
	Title string `json:"title,required"`
	// GitHub assignees
	Assignees []PledgeGetResponseIssueAssignee `json:"assignees,nullable"`
	// GitHub author
	Author PledgeGetResponseIssueAuthor `json:"author,nullable"`
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
	Labels            []PledgeGetResponseIssueLabel `json:"labels"`
	// GitHub reactions
	Reactions PledgeGetResponseIssueReactions `json:"reactions,nullable"`
	// Share of rewrads that will be rewarded to contributors of this issue. A number
	// between 0 and 100 (inclusive).
	UpfrontSplitToContributors int64                      `json:"upfront_split_to_contributors,nullable"`
	JSON                       pledgeGetResponseIssueJSON `json:"-"`
}

// pledgeGetResponseIssueJSON contains the JSON metadata for the struct
// [PledgeGetResponseIssue]
type pledgeGetResponseIssueJSON struct {
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

func (r *PledgeGetResponseIssue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pledgeGetResponseIssueJSON) RawJSON() string {
	return r.raw
}

type PledgeGetResponseIssueFunding struct {
	FundingGoal PledgeGetResponseIssueFundingFundingGoal `json:"funding_goal,nullable"`
	// Sum of pledges to this isuse (including currently open pledges and pledges that
	// have been paid out). Always in USD.
	PledgesSum PledgeGetResponseIssueFundingPledgesSum `json:"pledges_sum,nullable"`
	JSON       pledgeGetResponseIssueFundingJSON       `json:"-"`
}

// pledgeGetResponseIssueFundingJSON contains the JSON metadata for the struct
// [PledgeGetResponseIssueFunding]
type pledgeGetResponseIssueFundingJSON struct {
	FundingGoal apijson.Field
	PledgesSum  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PledgeGetResponseIssueFunding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pledgeGetResponseIssueFundingJSON) RawJSON() string {
	return r.raw
}

type PledgeGetResponseIssueFundingFundingGoal struct {
	// Amount in the currencies smallest unit (cents if currency is USD)
	Amount int64 `json:"amount,required"`
	// Three letter currency code (eg: USD)
	Currency string                                       `json:"currency,required"`
	JSON     pledgeGetResponseIssueFundingFundingGoalJSON `json:"-"`
}

// pledgeGetResponseIssueFundingFundingGoalJSON contains the JSON metadata for the
// struct [PledgeGetResponseIssueFundingFundingGoal]
type pledgeGetResponseIssueFundingFundingGoalJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PledgeGetResponseIssueFundingFundingGoal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pledgeGetResponseIssueFundingFundingGoalJSON) RawJSON() string {
	return r.raw
}

// Sum of pledges to this isuse (including currently open pledges and pledges that
// have been paid out). Always in USD.
type PledgeGetResponseIssueFundingPledgesSum struct {
	// Amount in the currencies smallest unit (cents if currency is USD)
	Amount int64 `json:"amount,required"`
	// Three letter currency code (eg: USD)
	Currency string                                      `json:"currency,required"`
	JSON     pledgeGetResponseIssueFundingPledgesSumJSON `json:"-"`
}

// pledgeGetResponseIssueFundingPledgesSumJSON contains the JSON metadata for the
// struct [PledgeGetResponseIssueFundingPledgesSum]
type pledgeGetResponseIssueFundingPledgesSumJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PledgeGetResponseIssueFundingPledgesSum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pledgeGetResponseIssueFundingPledgesSumJSON) RawJSON() string {
	return r.raw
}

// Issue platform (currently always GitHub)
type PledgeGetResponseIssuePlatform string

const (
	PledgeGetResponseIssuePlatformGitHub PledgeGetResponseIssuePlatform = "github"
)

func (r PledgeGetResponseIssuePlatform) IsKnown() bool {
	switch r {
	case PledgeGetResponseIssuePlatformGitHub:
		return true
	}
	return false
}

// The repository that the issue is in
type PledgeGetResponseIssueRepository struct {
	ID           string                                       `json:"id,required" format:"uuid"`
	IsPrivate    bool                                         `json:"is_private,required"`
	Name         string                                       `json:"name,required"`
	Organization PledgeGetResponseIssueRepositoryOrganization `json:"organization,required"`
	Platform     PledgeGetResponseIssueRepositoryPlatform     `json:"platform,required"`
	// Settings for the repository profile
	ProfileSettings PledgeGetResponseIssueRepositoryProfileSettings `json:"profile_settings,required,nullable"`
	Description     string                                          `json:"description,nullable"`
	Homepage        string                                          `json:"homepage,nullable"`
	License         string                                          `json:"license,nullable"`
	Stars           int64                                           `json:"stars,nullable"`
	JSON            pledgeGetResponseIssueRepositoryJSON            `json:"-"`
}

// pledgeGetResponseIssueRepositoryJSON contains the JSON metadata for the struct
// [PledgeGetResponseIssueRepository]
type pledgeGetResponseIssueRepositoryJSON struct {
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

func (r *PledgeGetResponseIssueRepository) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pledgeGetResponseIssueRepositoryJSON) RawJSON() string {
	return r.raw
}

type PledgeGetResponseIssueRepositoryOrganization struct {
	ID         string                                               `json:"id,required" format:"uuid"`
	AvatarURL  string                                               `json:"avatar_url,required"`
	IsPersonal bool                                                 `json:"is_personal,required"`
	Name       string                                               `json:"name,required"`
	Platform   PledgeGetResponseIssueRepositoryOrganizationPlatform `json:"platform,required"`
	Bio        string                                               `json:"bio,nullable"`
	Blog       string                                               `json:"blog,nullable"`
	Company    string                                               `json:"company,nullable"`
	Email      string                                               `json:"email,nullable"`
	Location   string                                               `json:"location,nullable"`
	// The organization ID.
	OrganizationID  string                                           `json:"organization_id,nullable" format:"uuid4"`
	PrettyName      string                                           `json:"pretty_name,nullable"`
	TwitterUsername string                                           `json:"twitter_username,nullable"`
	JSON            pledgeGetResponseIssueRepositoryOrganizationJSON `json:"-"`
}

// pledgeGetResponseIssueRepositoryOrganizationJSON contains the JSON metadata for
// the struct [PledgeGetResponseIssueRepositoryOrganization]
type pledgeGetResponseIssueRepositoryOrganizationJSON struct {
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

func (r *PledgeGetResponseIssueRepositoryOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pledgeGetResponseIssueRepositoryOrganizationJSON) RawJSON() string {
	return r.raw
}

type PledgeGetResponseIssueRepositoryOrganizationPlatform string

const (
	PledgeGetResponseIssueRepositoryOrganizationPlatformGitHub PledgeGetResponseIssueRepositoryOrganizationPlatform = "github"
)

func (r PledgeGetResponseIssueRepositoryOrganizationPlatform) IsKnown() bool {
	switch r {
	case PledgeGetResponseIssueRepositoryOrganizationPlatformGitHub:
		return true
	}
	return false
}

type PledgeGetResponseIssueRepositoryPlatform string

const (
	PledgeGetResponseIssueRepositoryPlatformGitHub PledgeGetResponseIssueRepositoryPlatform = "github"
)

func (r PledgeGetResponseIssueRepositoryPlatform) IsKnown() bool {
	switch r {
	case PledgeGetResponseIssueRepositoryPlatformGitHub:
		return true
	}
	return false
}

// Settings for the repository profile
type PledgeGetResponseIssueRepositoryProfileSettings struct {
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
	JSON  pledgeGetResponseIssueRepositoryProfileSettingsJSON `json:"-"`
}

// pledgeGetResponseIssueRepositoryProfileSettingsJSON contains the JSON metadata
// for the struct [PledgeGetResponseIssueRepositoryProfileSettings]
type pledgeGetResponseIssueRepositoryProfileSettingsJSON struct {
	CoverImageURL                apijson.Field
	Description                  apijson.Field
	FeaturedOrganizations        apijson.Field
	HighlightedSubscriptionTiers apijson.Field
	Links                        apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *PledgeGetResponseIssueRepositoryProfileSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pledgeGetResponseIssueRepositoryProfileSettingsJSON) RawJSON() string {
	return r.raw
}

type PledgeGetResponseIssueState string

const (
	PledgeGetResponseIssueStateOpen   PledgeGetResponseIssueState = "open"
	PledgeGetResponseIssueStateClosed PledgeGetResponseIssueState = "closed"
)

func (r PledgeGetResponseIssueState) IsKnown() bool {
	switch r {
	case PledgeGetResponseIssueStateOpen, PledgeGetResponseIssueStateClosed:
		return true
	}
	return false
}

type PledgeGetResponseIssueAssignee struct {
	ID        int64                              `json:"id,required"`
	AvatarURL string                             `json:"avatar_url,required" format:"uri"`
	HTMLURL   string                             `json:"html_url,required" format:"uri"`
	Login     string                             `json:"login,required"`
	JSON      pledgeGetResponseIssueAssigneeJSON `json:"-"`
}

// pledgeGetResponseIssueAssigneeJSON contains the JSON metadata for the struct
// [PledgeGetResponseIssueAssignee]
type pledgeGetResponseIssueAssigneeJSON struct {
	ID          apijson.Field
	AvatarURL   apijson.Field
	HTMLURL     apijson.Field
	Login       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PledgeGetResponseIssueAssignee) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pledgeGetResponseIssueAssigneeJSON) RawJSON() string {
	return r.raw
}

// GitHub author
type PledgeGetResponseIssueAuthor struct {
	ID        int64                            `json:"id,required"`
	AvatarURL string                           `json:"avatar_url,required" format:"uri"`
	HTMLURL   string                           `json:"html_url,required" format:"uri"`
	Login     string                           `json:"login,required"`
	JSON      pledgeGetResponseIssueAuthorJSON `json:"-"`
}

// pledgeGetResponseIssueAuthorJSON contains the JSON metadata for the struct
// [PledgeGetResponseIssueAuthor]
type pledgeGetResponseIssueAuthorJSON struct {
	ID          apijson.Field
	AvatarURL   apijson.Field
	HTMLURL     apijson.Field
	Login       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PledgeGetResponseIssueAuthor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pledgeGetResponseIssueAuthorJSON) RawJSON() string {
	return r.raw
}

type PledgeGetResponseIssueLabel struct {
	Color string                          `json:"color,required"`
	Name  string                          `json:"name,required"`
	JSON  pledgeGetResponseIssueLabelJSON `json:"-"`
}

// pledgeGetResponseIssueLabelJSON contains the JSON metadata for the struct
// [PledgeGetResponseIssueLabel]
type pledgeGetResponseIssueLabelJSON struct {
	Color       apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PledgeGetResponseIssueLabel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pledgeGetResponseIssueLabelJSON) RawJSON() string {
	return r.raw
}

// GitHub reactions
type PledgeGetResponseIssueReactions struct {
	Confused   int64                               `json:"confused,required"`
	Eyes       int64                               `json:"eyes,required"`
	Heart      int64                               `json:"heart,required"`
	Hooray     int64                               `json:"hooray,required"`
	Laugh      int64                               `json:"laugh,required"`
	MinusOne   int64                               `json:"minus_one,required"`
	PlusOne    int64                               `json:"plus_one,required"`
	Rocket     int64                               `json:"rocket,required"`
	TotalCount int64                               `json:"total_count,required"`
	JSON       pledgeGetResponseIssueReactionsJSON `json:"-"`
}

// pledgeGetResponseIssueReactionsJSON contains the JSON metadata for the struct
// [PledgeGetResponseIssueReactions]
type pledgeGetResponseIssueReactionsJSON struct {
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

func (r *PledgeGetResponseIssueReactions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pledgeGetResponseIssueReactionsJSON) RawJSON() string {
	return r.raw
}

// Current state of the pledge
type PledgeGetResponseState string

const (
	PledgeGetResponseStateInitiated      PledgeGetResponseState = "initiated"
	PledgeGetResponseStateCreated        PledgeGetResponseState = "created"
	PledgeGetResponseStatePending        PledgeGetResponseState = "pending"
	PledgeGetResponseStateRefunded       PledgeGetResponseState = "refunded"
	PledgeGetResponseStateDisputed       PledgeGetResponseState = "disputed"
	PledgeGetResponseStateChargeDisputed PledgeGetResponseState = "charge_disputed"
	PledgeGetResponseStateCancelled      PledgeGetResponseState = "cancelled"
)

func (r PledgeGetResponseState) IsKnown() bool {
	switch r {
	case PledgeGetResponseStateInitiated, PledgeGetResponseStateCreated, PledgeGetResponseStatePending, PledgeGetResponseStateRefunded, PledgeGetResponseStateDisputed, PledgeGetResponseStateChargeDisputed, PledgeGetResponseStateCancelled:
		return true
	}
	return false
}

// Type of pledge
type PledgeGetResponseType string

const (
	PledgeGetResponseTypePayUpfront      PledgeGetResponseType = "pay_upfront"
	PledgeGetResponseTypePayOnCompletion PledgeGetResponseType = "pay_on_completion"
	PledgeGetResponseTypePayDirectly     PledgeGetResponseType = "pay_directly"
)

func (r PledgeGetResponseType) IsKnown() bool {
	switch r {
	case PledgeGetResponseTypePayUpfront, PledgeGetResponseTypePayOnCompletion, PledgeGetResponseTypePayDirectly:
		return true
	}
	return false
}

// For pledges made by an organization, or on behalf of an organization. This is
// the user that made the pledge. Only visible for members of said organization.
type PledgeGetResponseCreatedBy struct {
	Name           string                         `json:"name,required"`
	AvatarURL      string                         `json:"avatar_url,nullable"`
	GitHubUsername string                         `json:"github_username,nullable"`
	JSON           pledgeGetResponseCreatedByJSON `json:"-"`
}

// pledgeGetResponseCreatedByJSON contains the JSON metadata for the struct
// [PledgeGetResponseCreatedBy]
type pledgeGetResponseCreatedByJSON struct {
	Name           apijson.Field
	AvatarURL      apijson.Field
	GitHubUsername apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PledgeGetResponseCreatedBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pledgeGetResponseCreatedByJSON) RawJSON() string {
	return r.raw
}

// The user or organization that made this pledge
type PledgeGetResponsePledger struct {
	Name           string                       `json:"name,required"`
	AvatarURL      string                       `json:"avatar_url,nullable"`
	GitHubUsername string                       `json:"github_username,nullable"`
	JSON           pledgeGetResponsePledgerJSON `json:"-"`
}

// pledgeGetResponsePledgerJSON contains the JSON metadata for the struct
// [PledgeGetResponsePledger]
type pledgeGetResponsePledgerJSON struct {
	Name           apijson.Field
	AvatarURL      apijson.Field
	GitHubUsername apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PledgeGetResponsePledger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pledgeGetResponsePledgerJSON) RawJSON() string {
	return r.raw
}

type PledgeSearchResponse struct {
	Pagination PledgeSearchResponsePagination `json:"pagination,required"`
	Items      []PledgeSearchResponseItem     `json:"items"`
	JSON       pledgeSearchResponseJSON       `json:"-"`
}

// pledgeSearchResponseJSON contains the JSON metadata for the struct
// [PledgeSearchResponse]
type pledgeSearchResponseJSON struct {
	Pagination  apijson.Field
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PledgeSearchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pledgeSearchResponseJSON) RawJSON() string {
	return r.raw
}

type PledgeSearchResponsePagination struct {
	MaxPage    int64                              `json:"max_page,required"`
	TotalCount int64                              `json:"total_count,required"`
	JSON       pledgeSearchResponsePaginationJSON `json:"-"`
}

// pledgeSearchResponsePaginationJSON contains the JSON metadata for the struct
// [PledgeSearchResponsePagination]
type pledgeSearchResponsePaginationJSON struct {
	MaxPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PledgeSearchResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pledgeSearchResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type PledgeSearchResponseItem struct {
	// The ID of the object.
	ID string `json:"id,required" format:"uuid4"`
	// Amount pledged towards the issue
	Amount int64 `json:"amount,required"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	Currency  string    `json:"currency,required"`
	// The issue that the pledge was made towards
	Issue PledgeSearchResponseItemsIssue `json:"issue,required"`
	// Current state of the pledge
	State PledgeSearchResponseItemsState `json:"state,required"`
	// Type of pledge
	Type PledgeSearchResponseItemsType `json:"type,required"`
	// If the currently authenticated subject can perform admin actions on behalf of
	// the receiver of the peldge
	AuthedCanAdminReceived bool `json:"authed_can_admin_received"`
	// If the currently authenticated subject can perform admin actions on behalf of
	// the maker of the peldge
	AuthedCanAdminSender bool `json:"authed_can_admin_sender"`
	// For pledges made by an organization, or on behalf of an organization. This is
	// the user that made the pledge. Only visible for members of said organization.
	CreatedBy PledgeSearchResponseItemsCreatedBy `json:"created_by,nullable"`
	// URL of invoice for this pledge
	HostedInvoiceURL string `json:"hosted_invoice_url,nullable"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,nullable" format:"date-time"`
	// The user or organization that made this pledge
	Pledger PledgeSearchResponseItemsPledger `json:"pledger,nullable"`
	// If and when the pledge was refunded to the pledger
	RefundedAt time.Time `json:"refunded_at,nullable" format:"date-time"`
	// When the payout is scheduled to be made to the maintainers behind the issue.
	// Disputes must be made before this date.
	ScheduledPayoutAt time.Time                    `json:"scheduled_payout_at,nullable" format:"date-time"`
	JSON              pledgeSearchResponseItemJSON `json:"-"`
}

// pledgeSearchResponseItemJSON contains the JSON metadata for the struct
// [PledgeSearchResponseItem]
type pledgeSearchResponseItemJSON struct {
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

func (r *PledgeSearchResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pledgeSearchResponseItemJSON) RawJSON() string {
	return r.raw
}

// The issue that the pledge was made towards
type PledgeSearchResponseItemsIssue struct {
	ID             string                                `json:"id,required" format:"uuid"`
	Funding        PledgeSearchResponseItemsIssueFunding `json:"funding,required"`
	IssueCreatedAt time.Time                             `json:"issue_created_at,required" format:"date-time"`
	// If a maintainer needs to mark this issue as solved
	NeedsConfirmationSolved bool `json:"needs_confirmation_solved,required"`
	// GitHub #number
	Number int64 `json:"number,required"`
	// Issue platform (currently always GitHub)
	Platform PledgeSearchResponseItemsIssuePlatform `json:"platform,required"`
	// If this issue currently has the Polar badge SVG embedded
	PledgeBadgeCurrentlyEmbedded bool `json:"pledge_badge_currently_embedded,required"`
	// The repository that the issue is in
	Repository PledgeSearchResponseItemsIssueRepository `json:"repository,required"`
	State      PledgeSearchResponseItemsIssueState      `json:"state,required"`
	// GitHub issue title
	Title string `json:"title,required"`
	// GitHub assignees
	Assignees []PledgeSearchResponseItemsIssueAssignee `json:"assignees,nullable"`
	// GitHub author
	Author PledgeSearchResponseItemsIssueAuthor `json:"author,nullable"`
	// Optional custom badge SVG promotional content
	BadgeCustomContent string `json:"badge_custom_content,nullable"`
	// GitHub issue body
	Body string `json:"body,nullable"`
	// Number of GitHub comments made on the issue
	Comments int64 `json:"comments,nullable"`
	// If this issue has been marked as confirmed solved through Polar
	ConfirmedSolvedAt time.Time                             `json:"confirmed_solved_at,nullable" format:"date-time"`
	IssueClosedAt     time.Time                             `json:"issue_closed_at,nullable" format:"date-time"`
	IssueModifiedAt   time.Time                             `json:"issue_modified_at,nullable" format:"date-time"`
	Labels            []PledgeSearchResponseItemsIssueLabel `json:"labels"`
	// GitHub reactions
	Reactions PledgeSearchResponseItemsIssueReactions `json:"reactions,nullable"`
	// Share of rewrads that will be rewarded to contributors of this issue. A number
	// between 0 and 100 (inclusive).
	UpfrontSplitToContributors int64                              `json:"upfront_split_to_contributors,nullable"`
	JSON                       pledgeSearchResponseItemsIssueJSON `json:"-"`
}

// pledgeSearchResponseItemsIssueJSON contains the JSON metadata for the struct
// [PledgeSearchResponseItemsIssue]
type pledgeSearchResponseItemsIssueJSON struct {
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

func (r *PledgeSearchResponseItemsIssue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pledgeSearchResponseItemsIssueJSON) RawJSON() string {
	return r.raw
}

type PledgeSearchResponseItemsIssueFunding struct {
	FundingGoal PledgeSearchResponseItemsIssueFundingFundingGoal `json:"funding_goal,nullable"`
	// Sum of pledges to this isuse (including currently open pledges and pledges that
	// have been paid out). Always in USD.
	PledgesSum PledgeSearchResponseItemsIssueFundingPledgesSum `json:"pledges_sum,nullable"`
	JSON       pledgeSearchResponseItemsIssueFundingJSON       `json:"-"`
}

// pledgeSearchResponseItemsIssueFundingJSON contains the JSON metadata for the
// struct [PledgeSearchResponseItemsIssueFunding]
type pledgeSearchResponseItemsIssueFundingJSON struct {
	FundingGoal apijson.Field
	PledgesSum  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PledgeSearchResponseItemsIssueFunding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pledgeSearchResponseItemsIssueFundingJSON) RawJSON() string {
	return r.raw
}

type PledgeSearchResponseItemsIssueFundingFundingGoal struct {
	// Amount in the currencies smallest unit (cents if currency is USD)
	Amount int64 `json:"amount,required"`
	// Three letter currency code (eg: USD)
	Currency string                                               `json:"currency,required"`
	JSON     pledgeSearchResponseItemsIssueFundingFundingGoalJSON `json:"-"`
}

// pledgeSearchResponseItemsIssueFundingFundingGoalJSON contains the JSON metadata
// for the struct [PledgeSearchResponseItemsIssueFundingFundingGoal]
type pledgeSearchResponseItemsIssueFundingFundingGoalJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PledgeSearchResponseItemsIssueFundingFundingGoal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pledgeSearchResponseItemsIssueFundingFundingGoalJSON) RawJSON() string {
	return r.raw
}

// Sum of pledges to this isuse (including currently open pledges and pledges that
// have been paid out). Always in USD.
type PledgeSearchResponseItemsIssueFundingPledgesSum struct {
	// Amount in the currencies smallest unit (cents if currency is USD)
	Amount int64 `json:"amount,required"`
	// Three letter currency code (eg: USD)
	Currency string                                              `json:"currency,required"`
	JSON     pledgeSearchResponseItemsIssueFundingPledgesSumJSON `json:"-"`
}

// pledgeSearchResponseItemsIssueFundingPledgesSumJSON contains the JSON metadata
// for the struct [PledgeSearchResponseItemsIssueFundingPledgesSum]
type pledgeSearchResponseItemsIssueFundingPledgesSumJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PledgeSearchResponseItemsIssueFundingPledgesSum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pledgeSearchResponseItemsIssueFundingPledgesSumJSON) RawJSON() string {
	return r.raw
}

// Issue platform (currently always GitHub)
type PledgeSearchResponseItemsIssuePlatform string

const (
	PledgeSearchResponseItemsIssuePlatformGitHub PledgeSearchResponseItemsIssuePlatform = "github"
)

func (r PledgeSearchResponseItemsIssuePlatform) IsKnown() bool {
	switch r {
	case PledgeSearchResponseItemsIssuePlatformGitHub:
		return true
	}
	return false
}

// The repository that the issue is in
type PledgeSearchResponseItemsIssueRepository struct {
	ID           string                                               `json:"id,required" format:"uuid"`
	IsPrivate    bool                                                 `json:"is_private,required"`
	Name         string                                               `json:"name,required"`
	Organization PledgeSearchResponseItemsIssueRepositoryOrganization `json:"organization,required"`
	Platform     PledgeSearchResponseItemsIssueRepositoryPlatform     `json:"platform,required"`
	// Settings for the repository profile
	ProfileSettings PledgeSearchResponseItemsIssueRepositoryProfileSettings `json:"profile_settings,required,nullable"`
	Description     string                                                  `json:"description,nullable"`
	Homepage        string                                                  `json:"homepage,nullable"`
	License         string                                                  `json:"license,nullable"`
	Stars           int64                                                   `json:"stars,nullable"`
	JSON            pledgeSearchResponseItemsIssueRepositoryJSON            `json:"-"`
}

// pledgeSearchResponseItemsIssueRepositoryJSON contains the JSON metadata for the
// struct [PledgeSearchResponseItemsIssueRepository]
type pledgeSearchResponseItemsIssueRepositoryJSON struct {
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

func (r *PledgeSearchResponseItemsIssueRepository) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pledgeSearchResponseItemsIssueRepositoryJSON) RawJSON() string {
	return r.raw
}

type PledgeSearchResponseItemsIssueRepositoryOrganization struct {
	ID         string                                                       `json:"id,required" format:"uuid"`
	AvatarURL  string                                                       `json:"avatar_url,required"`
	IsPersonal bool                                                         `json:"is_personal,required"`
	Name       string                                                       `json:"name,required"`
	Platform   PledgeSearchResponseItemsIssueRepositoryOrganizationPlatform `json:"platform,required"`
	Bio        string                                                       `json:"bio,nullable"`
	Blog       string                                                       `json:"blog,nullable"`
	Company    string                                                       `json:"company,nullable"`
	Email      string                                                       `json:"email,nullable"`
	Location   string                                                       `json:"location,nullable"`
	// The organization ID.
	OrganizationID  string                                                   `json:"organization_id,nullable" format:"uuid4"`
	PrettyName      string                                                   `json:"pretty_name,nullable"`
	TwitterUsername string                                                   `json:"twitter_username,nullable"`
	JSON            pledgeSearchResponseItemsIssueRepositoryOrganizationJSON `json:"-"`
}

// pledgeSearchResponseItemsIssueRepositoryOrganizationJSON contains the JSON
// metadata for the struct [PledgeSearchResponseItemsIssueRepositoryOrganization]
type pledgeSearchResponseItemsIssueRepositoryOrganizationJSON struct {
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

func (r *PledgeSearchResponseItemsIssueRepositoryOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pledgeSearchResponseItemsIssueRepositoryOrganizationJSON) RawJSON() string {
	return r.raw
}

type PledgeSearchResponseItemsIssueRepositoryOrganizationPlatform string

const (
	PledgeSearchResponseItemsIssueRepositoryOrganizationPlatformGitHub PledgeSearchResponseItemsIssueRepositoryOrganizationPlatform = "github"
)

func (r PledgeSearchResponseItemsIssueRepositoryOrganizationPlatform) IsKnown() bool {
	switch r {
	case PledgeSearchResponseItemsIssueRepositoryOrganizationPlatformGitHub:
		return true
	}
	return false
}

type PledgeSearchResponseItemsIssueRepositoryPlatform string

const (
	PledgeSearchResponseItemsIssueRepositoryPlatformGitHub PledgeSearchResponseItemsIssueRepositoryPlatform = "github"
)

func (r PledgeSearchResponseItemsIssueRepositoryPlatform) IsKnown() bool {
	switch r {
	case PledgeSearchResponseItemsIssueRepositoryPlatformGitHub:
		return true
	}
	return false
}

// Settings for the repository profile
type PledgeSearchResponseItemsIssueRepositoryProfileSettings struct {
	// A URL to a cover image
	CoverImageURL string `json:"cover_image_url,nullable"`
	// A description of the repository
	Description string `json:"description,nullable"`
	// A list of featured organizations
	FeaturedOrganizations []string `json:"featured_organizations,nullable" format:"uuid4"`
	// A list of highlighted subscription tiers
	HighlightedSubscriptionTiers []string `json:"highlighted_subscription_tiers,nullable" format:"uuid4"`
	// A list of links related to the repository
	Links []string                                                    `json:"links,nullable" format:"uri"`
	JSON  pledgeSearchResponseItemsIssueRepositoryProfileSettingsJSON `json:"-"`
}

// pledgeSearchResponseItemsIssueRepositoryProfileSettingsJSON contains the JSON
// metadata for the struct
// [PledgeSearchResponseItemsIssueRepositoryProfileSettings]
type pledgeSearchResponseItemsIssueRepositoryProfileSettingsJSON struct {
	CoverImageURL                apijson.Field
	Description                  apijson.Field
	FeaturedOrganizations        apijson.Field
	HighlightedSubscriptionTiers apijson.Field
	Links                        apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *PledgeSearchResponseItemsIssueRepositoryProfileSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pledgeSearchResponseItemsIssueRepositoryProfileSettingsJSON) RawJSON() string {
	return r.raw
}

type PledgeSearchResponseItemsIssueState string

const (
	PledgeSearchResponseItemsIssueStateOpen   PledgeSearchResponseItemsIssueState = "open"
	PledgeSearchResponseItemsIssueStateClosed PledgeSearchResponseItemsIssueState = "closed"
)

func (r PledgeSearchResponseItemsIssueState) IsKnown() bool {
	switch r {
	case PledgeSearchResponseItemsIssueStateOpen, PledgeSearchResponseItemsIssueStateClosed:
		return true
	}
	return false
}

type PledgeSearchResponseItemsIssueAssignee struct {
	ID        int64                                      `json:"id,required"`
	AvatarURL string                                     `json:"avatar_url,required" format:"uri"`
	HTMLURL   string                                     `json:"html_url,required" format:"uri"`
	Login     string                                     `json:"login,required"`
	JSON      pledgeSearchResponseItemsIssueAssigneeJSON `json:"-"`
}

// pledgeSearchResponseItemsIssueAssigneeJSON contains the JSON metadata for the
// struct [PledgeSearchResponseItemsIssueAssignee]
type pledgeSearchResponseItemsIssueAssigneeJSON struct {
	ID          apijson.Field
	AvatarURL   apijson.Field
	HTMLURL     apijson.Field
	Login       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PledgeSearchResponseItemsIssueAssignee) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pledgeSearchResponseItemsIssueAssigneeJSON) RawJSON() string {
	return r.raw
}

// GitHub author
type PledgeSearchResponseItemsIssueAuthor struct {
	ID        int64                                    `json:"id,required"`
	AvatarURL string                                   `json:"avatar_url,required" format:"uri"`
	HTMLURL   string                                   `json:"html_url,required" format:"uri"`
	Login     string                                   `json:"login,required"`
	JSON      pledgeSearchResponseItemsIssueAuthorJSON `json:"-"`
}

// pledgeSearchResponseItemsIssueAuthorJSON contains the JSON metadata for the
// struct [PledgeSearchResponseItemsIssueAuthor]
type pledgeSearchResponseItemsIssueAuthorJSON struct {
	ID          apijson.Field
	AvatarURL   apijson.Field
	HTMLURL     apijson.Field
	Login       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PledgeSearchResponseItemsIssueAuthor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pledgeSearchResponseItemsIssueAuthorJSON) RawJSON() string {
	return r.raw
}

type PledgeSearchResponseItemsIssueLabel struct {
	Color string                                  `json:"color,required"`
	Name  string                                  `json:"name,required"`
	JSON  pledgeSearchResponseItemsIssueLabelJSON `json:"-"`
}

// pledgeSearchResponseItemsIssueLabelJSON contains the JSON metadata for the
// struct [PledgeSearchResponseItemsIssueLabel]
type pledgeSearchResponseItemsIssueLabelJSON struct {
	Color       apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PledgeSearchResponseItemsIssueLabel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pledgeSearchResponseItemsIssueLabelJSON) RawJSON() string {
	return r.raw
}

// GitHub reactions
type PledgeSearchResponseItemsIssueReactions struct {
	Confused   int64                                       `json:"confused,required"`
	Eyes       int64                                       `json:"eyes,required"`
	Heart      int64                                       `json:"heart,required"`
	Hooray     int64                                       `json:"hooray,required"`
	Laugh      int64                                       `json:"laugh,required"`
	MinusOne   int64                                       `json:"minus_one,required"`
	PlusOne    int64                                       `json:"plus_one,required"`
	Rocket     int64                                       `json:"rocket,required"`
	TotalCount int64                                       `json:"total_count,required"`
	JSON       pledgeSearchResponseItemsIssueReactionsJSON `json:"-"`
}

// pledgeSearchResponseItemsIssueReactionsJSON contains the JSON metadata for the
// struct [PledgeSearchResponseItemsIssueReactions]
type pledgeSearchResponseItemsIssueReactionsJSON struct {
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

func (r *PledgeSearchResponseItemsIssueReactions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pledgeSearchResponseItemsIssueReactionsJSON) RawJSON() string {
	return r.raw
}

// Current state of the pledge
type PledgeSearchResponseItemsState string

const (
	PledgeSearchResponseItemsStateInitiated      PledgeSearchResponseItemsState = "initiated"
	PledgeSearchResponseItemsStateCreated        PledgeSearchResponseItemsState = "created"
	PledgeSearchResponseItemsStatePending        PledgeSearchResponseItemsState = "pending"
	PledgeSearchResponseItemsStateRefunded       PledgeSearchResponseItemsState = "refunded"
	PledgeSearchResponseItemsStateDisputed       PledgeSearchResponseItemsState = "disputed"
	PledgeSearchResponseItemsStateChargeDisputed PledgeSearchResponseItemsState = "charge_disputed"
	PledgeSearchResponseItemsStateCancelled      PledgeSearchResponseItemsState = "cancelled"
)

func (r PledgeSearchResponseItemsState) IsKnown() bool {
	switch r {
	case PledgeSearchResponseItemsStateInitiated, PledgeSearchResponseItemsStateCreated, PledgeSearchResponseItemsStatePending, PledgeSearchResponseItemsStateRefunded, PledgeSearchResponseItemsStateDisputed, PledgeSearchResponseItemsStateChargeDisputed, PledgeSearchResponseItemsStateCancelled:
		return true
	}
	return false
}

// Type of pledge
type PledgeSearchResponseItemsType string

const (
	PledgeSearchResponseItemsTypePayUpfront      PledgeSearchResponseItemsType = "pay_upfront"
	PledgeSearchResponseItemsTypePayOnCompletion PledgeSearchResponseItemsType = "pay_on_completion"
	PledgeSearchResponseItemsTypePayDirectly     PledgeSearchResponseItemsType = "pay_directly"
)

func (r PledgeSearchResponseItemsType) IsKnown() bool {
	switch r {
	case PledgeSearchResponseItemsTypePayUpfront, PledgeSearchResponseItemsTypePayOnCompletion, PledgeSearchResponseItemsTypePayDirectly:
		return true
	}
	return false
}

// For pledges made by an organization, or on behalf of an organization. This is
// the user that made the pledge. Only visible for members of said organization.
type PledgeSearchResponseItemsCreatedBy struct {
	Name           string                                 `json:"name,required"`
	AvatarURL      string                                 `json:"avatar_url,nullable"`
	GitHubUsername string                                 `json:"github_username,nullable"`
	JSON           pledgeSearchResponseItemsCreatedByJSON `json:"-"`
}

// pledgeSearchResponseItemsCreatedByJSON contains the JSON metadata for the struct
// [PledgeSearchResponseItemsCreatedBy]
type pledgeSearchResponseItemsCreatedByJSON struct {
	Name           apijson.Field
	AvatarURL      apijson.Field
	GitHubUsername apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PledgeSearchResponseItemsCreatedBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pledgeSearchResponseItemsCreatedByJSON) RawJSON() string {
	return r.raw
}

// The user or organization that made this pledge
type PledgeSearchResponseItemsPledger struct {
	Name           string                               `json:"name,required"`
	AvatarURL      string                               `json:"avatar_url,nullable"`
	GitHubUsername string                               `json:"github_username,nullable"`
	JSON           pledgeSearchResponseItemsPledgerJSON `json:"-"`
}

// pledgeSearchResponseItemsPledgerJSON contains the JSON metadata for the struct
// [PledgeSearchResponseItemsPledger]
type pledgeSearchResponseItemsPledgerJSON struct {
	Name           apijson.Field
	AvatarURL      apijson.Field
	GitHubUsername apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PledgeSearchResponseItemsPledger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pledgeSearchResponseItemsPledgerJSON) RawJSON() string {
	return r.raw
}

type PledgeSpendingResponse struct {
	Amount   int64                      `json:"amount,required"`
	Currency string                     `json:"currency,required"`
	JSON     pledgeSpendingResponseJSON `json:"-"`
}

// pledgeSpendingResponseJSON contains the JSON metadata for the struct
// [PledgeSpendingResponse]
type pledgeSpendingResponseJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PledgeSpendingResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pledgeSpendingResponseJSON) RawJSON() string {
	return r.raw
}

type PledgeSearchParams struct {
	// Search pledges made by this organization.
	ByOrganizationID param.Field[string] `query:"by_organization_id" format:"uuid"`
	// Search pledges made by this user.
	ByUserID param.Field[string] `query:"by_user_id" format:"uuid"`
	// Search pledges to this issue
	IssueID param.Field[string] `query:"issue_id" format:"uuid"`
	// The organization ID.
	OrganizationID param.Field[string] `query:"organization_id" format:"uuid4"`
	// Search pledges in the repository with this name. Can only be used if
	// organization_name is set.
	RepositoryName param.Field[string] `query:"repository_name"`
}

// URLQuery serializes [PledgeSearchParams]'s query parameters as `url.Values`.
func (r PledgeSearchParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PledgeSpendingParams struct {
	// Spending in this organization. Required.
	OrganizationID param.Field[string] `query:"organization_id,required" format:"uuid"`
}

// URLQuery serializes [PledgeSpendingParams]'s query parameters as `url.Values`.
func (r PledgeSpendingParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
