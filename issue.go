// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package polar

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/polar-go/internal/apijson"
	"github.com/stainless-sdks/polar-go/internal/apiquery"
	"github.com/stainless-sdks/polar-go/internal/param"
	"github.com/stainless-sdks/polar-go/internal/requestconfig"
	"github.com/stainless-sdks/polar-go/option"
)

// IssueService contains methods and other services that help with interacting with
// the polar API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIssueService] method instead.
type IssueService struct {
	Options []option.RequestOption
	Lookup  *IssueLookupService
	Body    *IssueBodyService
}

// NewIssueService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewIssueService(opts ...option.RequestOption) (r *IssueService) {
	r = &IssueService{}
	r.Options = opts
	r.Lookup = NewIssueLookupService(opts...)
	r.Body = NewIssueBodyService(opts...)
	return
}

// Get issue
func (r *IssueService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *IssueGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/issues/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update issue. Requires authentication.
func (r *IssueService) Update(ctx context.Context, id string, body IssueUpdateParams, opts ...option.RequestOption) (res *IssueUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/issues/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List issues.
func (r *IssueService) List(ctx context.Context, query IssueListParams, opts ...option.RequestOption) (res *IssueListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/issues/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Mark an issue as confirmed solved, and configure issue reward splits. Enables
// payouts of pledges. Can only be done once per issue. Requires authentication.
func (r *IssueService) ConfirmSolved(ctx context.Context, id string, body IssueConfirmSolvedParams, opts ...option.RequestOption) (res *IssueConfirmSolvedResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/issues/%s/confirm_solved", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type IssueGetResponse struct {
	ID             string                  `json:"id,required" format:"uuid"`
	Funding        IssueGetResponseFunding `json:"funding,required"`
	IssueCreatedAt time.Time               `json:"issue_created_at,required" format:"date-time"`
	// If a maintainer needs to mark this issue as solved
	NeedsConfirmationSolved bool `json:"needs_confirmation_solved,required"`
	// GitHub #number
	Number int64 `json:"number,required"`
	// Issue platform (currently always GitHub)
	Platform IssueGetResponsePlatform `json:"platform,required"`
	// If this issue currently has the Polar badge SVG embedded
	PledgeBadgeCurrentlyEmbedded bool `json:"pledge_badge_currently_embedded,required"`
	// The repository that the issue is in
	Repository IssueGetResponseRepository `json:"repository,required"`
	State      IssueGetResponseState      `json:"state,required"`
	// GitHub issue title
	Title string `json:"title,required"`
	// GitHub assignees
	Assignees []IssueGetResponseAssignee `json:"assignees,nullable"`
	// GitHub author
	Author IssueGetResponseAuthor `json:"author,nullable"`
	// Optional custom badge SVG promotional content
	BadgeCustomContent string `json:"badge_custom_content,nullable"`
	// GitHub issue body
	Body string `json:"body,nullable"`
	// Number of GitHub comments made on the issue
	Comments int64 `json:"comments,nullable"`
	// If this issue has been marked as confirmed solved through Polar
	ConfirmedSolvedAt time.Time               `json:"confirmed_solved_at,nullable" format:"date-time"`
	IssueClosedAt     time.Time               `json:"issue_closed_at,nullable" format:"date-time"`
	IssueModifiedAt   time.Time               `json:"issue_modified_at,nullable" format:"date-time"`
	Labels            []IssueGetResponseLabel `json:"labels"`
	// GitHub reactions
	Reactions IssueGetResponseReactions `json:"reactions,nullable"`
	// Share of rewrads that will be rewarded to contributors of this issue. A number
	// between 0 and 100 (inclusive).
	UpfrontSplitToContributors int64                `json:"upfront_split_to_contributors,nullable"`
	JSON                       issueGetResponseJSON `json:"-"`
}

// issueGetResponseJSON contains the JSON metadata for the struct
// [IssueGetResponse]
type issueGetResponseJSON struct {
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

func (r *IssueGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueGetResponseJSON) RawJSON() string {
	return r.raw
}

type IssueGetResponseFunding struct {
	FundingGoal IssueGetResponseFundingFundingGoal `json:"funding_goal,nullable"`
	// Sum of pledges to this isuse (including currently open pledges and pledges that
	// have been paid out). Always in USD.
	PledgesSum IssueGetResponseFundingPledgesSum `json:"pledges_sum,nullable"`
	JSON       issueGetResponseFundingJSON       `json:"-"`
}

// issueGetResponseFundingJSON contains the JSON metadata for the struct
// [IssueGetResponseFunding]
type issueGetResponseFundingJSON struct {
	FundingGoal apijson.Field
	PledgesSum  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IssueGetResponseFunding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueGetResponseFundingJSON) RawJSON() string {
	return r.raw
}

type IssueGetResponseFundingFundingGoal struct {
	// Amount in the currencies smallest unit (cents if currency is USD)
	Amount int64 `json:"amount,required"`
	// Three letter currency code (eg: USD)
	Currency string                                 `json:"currency,required"`
	JSON     issueGetResponseFundingFundingGoalJSON `json:"-"`
}

// issueGetResponseFundingFundingGoalJSON contains the JSON metadata for the struct
// [IssueGetResponseFundingFundingGoal]
type issueGetResponseFundingFundingGoalJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IssueGetResponseFundingFundingGoal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueGetResponseFundingFundingGoalJSON) RawJSON() string {
	return r.raw
}

// Sum of pledges to this isuse (including currently open pledges and pledges that
// have been paid out). Always in USD.
type IssueGetResponseFundingPledgesSum struct {
	// Amount in the currencies smallest unit (cents if currency is USD)
	Amount int64 `json:"amount,required"`
	// Three letter currency code (eg: USD)
	Currency string                                `json:"currency,required"`
	JSON     issueGetResponseFundingPledgesSumJSON `json:"-"`
}

// issueGetResponseFundingPledgesSumJSON contains the JSON metadata for the struct
// [IssueGetResponseFundingPledgesSum]
type issueGetResponseFundingPledgesSumJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IssueGetResponseFundingPledgesSum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueGetResponseFundingPledgesSumJSON) RawJSON() string {
	return r.raw
}

// Issue platform (currently always GitHub)
type IssueGetResponsePlatform string

const (
	IssueGetResponsePlatformGitHub IssueGetResponsePlatform = "github"
)

func (r IssueGetResponsePlatform) IsKnown() bool {
	switch r {
	case IssueGetResponsePlatformGitHub:
		return true
	}
	return false
}

// The repository that the issue is in
type IssueGetResponseRepository struct {
	ID           string                                 `json:"id,required" format:"uuid"`
	IsPrivate    bool                                   `json:"is_private,required"`
	Name         string                                 `json:"name,required"`
	Organization IssueGetResponseRepositoryOrganization `json:"organization,required"`
	Platform     IssueGetResponseRepositoryPlatform     `json:"platform,required"`
	// Settings for the repository profile
	ProfileSettings IssueGetResponseRepositoryProfileSettings `json:"profile_settings,required,nullable"`
	Description     string                                    `json:"description,nullable"`
	Homepage        string                                    `json:"homepage,nullable"`
	License         string                                    `json:"license,nullable"`
	Stars           int64                                     `json:"stars,nullable"`
	JSON            issueGetResponseRepositoryJSON            `json:"-"`
}

// issueGetResponseRepositoryJSON contains the JSON metadata for the struct
// [IssueGetResponseRepository]
type issueGetResponseRepositoryJSON struct {
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

func (r *IssueGetResponseRepository) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueGetResponseRepositoryJSON) RawJSON() string {
	return r.raw
}

type IssueGetResponseRepositoryOrganization struct {
	ID         string                                         `json:"id,required" format:"uuid"`
	AvatarURL  string                                         `json:"avatar_url,required"`
	IsPersonal bool                                           `json:"is_personal,required"`
	Name       string                                         `json:"name,required"`
	Platform   IssueGetResponseRepositoryOrganizationPlatform `json:"platform,required"`
	Bio        string                                         `json:"bio,nullable"`
	Blog       string                                         `json:"blog,nullable"`
	Company    string                                         `json:"company,nullable"`
	Email      string                                         `json:"email,nullable"`
	Location   string                                         `json:"location,nullable"`
	// The organization ID.
	OrganizationID  string                                     `json:"organization_id,nullable" format:"uuid4"`
	PrettyName      string                                     `json:"pretty_name,nullable"`
	TwitterUsername string                                     `json:"twitter_username,nullable"`
	JSON            issueGetResponseRepositoryOrganizationJSON `json:"-"`
}

// issueGetResponseRepositoryOrganizationJSON contains the JSON metadata for the
// struct [IssueGetResponseRepositoryOrganization]
type issueGetResponseRepositoryOrganizationJSON struct {
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

func (r *IssueGetResponseRepositoryOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueGetResponseRepositoryOrganizationJSON) RawJSON() string {
	return r.raw
}

type IssueGetResponseRepositoryOrganizationPlatform string

const (
	IssueGetResponseRepositoryOrganizationPlatformGitHub IssueGetResponseRepositoryOrganizationPlatform = "github"
)

func (r IssueGetResponseRepositoryOrganizationPlatform) IsKnown() bool {
	switch r {
	case IssueGetResponseRepositoryOrganizationPlatformGitHub:
		return true
	}
	return false
}

type IssueGetResponseRepositoryPlatform string

const (
	IssueGetResponseRepositoryPlatformGitHub IssueGetResponseRepositoryPlatform = "github"
)

func (r IssueGetResponseRepositoryPlatform) IsKnown() bool {
	switch r {
	case IssueGetResponseRepositoryPlatformGitHub:
		return true
	}
	return false
}

// Settings for the repository profile
type IssueGetResponseRepositoryProfileSettings struct {
	// A URL to a cover image
	CoverImageURL string `json:"cover_image_url,nullable"`
	// A description of the repository
	Description string `json:"description,nullable"`
	// A list of featured organizations
	FeaturedOrganizations []string `json:"featured_organizations,nullable" format:"uuid4"`
	// A list of highlighted subscription tiers
	HighlightedSubscriptionTiers []string `json:"highlighted_subscription_tiers,nullable" format:"uuid4"`
	// A list of links related to the repository
	Links []string                                      `json:"links,nullable" format:"uri"`
	JSON  issueGetResponseRepositoryProfileSettingsJSON `json:"-"`
}

// issueGetResponseRepositoryProfileSettingsJSON contains the JSON metadata for the
// struct [IssueGetResponseRepositoryProfileSettings]
type issueGetResponseRepositoryProfileSettingsJSON struct {
	CoverImageURL                apijson.Field
	Description                  apijson.Field
	FeaturedOrganizations        apijson.Field
	HighlightedSubscriptionTiers apijson.Field
	Links                        apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *IssueGetResponseRepositoryProfileSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueGetResponseRepositoryProfileSettingsJSON) RawJSON() string {
	return r.raw
}

type IssueGetResponseState string

const (
	IssueGetResponseStateOpen   IssueGetResponseState = "open"
	IssueGetResponseStateClosed IssueGetResponseState = "closed"
)

func (r IssueGetResponseState) IsKnown() bool {
	switch r {
	case IssueGetResponseStateOpen, IssueGetResponseStateClosed:
		return true
	}
	return false
}

type IssueGetResponseAssignee struct {
	ID        int64                        `json:"id,required"`
	AvatarURL string                       `json:"avatar_url,required" format:"uri"`
	HTMLURL   string                       `json:"html_url,required" format:"uri"`
	Login     string                       `json:"login,required"`
	JSON      issueGetResponseAssigneeJSON `json:"-"`
}

// issueGetResponseAssigneeJSON contains the JSON metadata for the struct
// [IssueGetResponseAssignee]
type issueGetResponseAssigneeJSON struct {
	ID          apijson.Field
	AvatarURL   apijson.Field
	HTMLURL     apijson.Field
	Login       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IssueGetResponseAssignee) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueGetResponseAssigneeJSON) RawJSON() string {
	return r.raw
}

// GitHub author
type IssueGetResponseAuthor struct {
	ID        int64                      `json:"id,required"`
	AvatarURL string                     `json:"avatar_url,required" format:"uri"`
	HTMLURL   string                     `json:"html_url,required" format:"uri"`
	Login     string                     `json:"login,required"`
	JSON      issueGetResponseAuthorJSON `json:"-"`
}

// issueGetResponseAuthorJSON contains the JSON metadata for the struct
// [IssueGetResponseAuthor]
type issueGetResponseAuthorJSON struct {
	ID          apijson.Field
	AvatarURL   apijson.Field
	HTMLURL     apijson.Field
	Login       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IssueGetResponseAuthor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueGetResponseAuthorJSON) RawJSON() string {
	return r.raw
}

type IssueGetResponseLabel struct {
	Color string                    `json:"color,required"`
	Name  string                    `json:"name,required"`
	JSON  issueGetResponseLabelJSON `json:"-"`
}

// issueGetResponseLabelJSON contains the JSON metadata for the struct
// [IssueGetResponseLabel]
type issueGetResponseLabelJSON struct {
	Color       apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IssueGetResponseLabel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueGetResponseLabelJSON) RawJSON() string {
	return r.raw
}

// GitHub reactions
type IssueGetResponseReactions struct {
	Confused   int64                         `json:"confused,required"`
	Eyes       int64                         `json:"eyes,required"`
	Heart      int64                         `json:"heart,required"`
	Hooray     int64                         `json:"hooray,required"`
	Laugh      int64                         `json:"laugh,required"`
	MinusOne   int64                         `json:"minus_one,required"`
	PlusOne    int64                         `json:"plus_one,required"`
	Rocket     int64                         `json:"rocket,required"`
	TotalCount int64                         `json:"total_count,required"`
	JSON       issueGetResponseReactionsJSON `json:"-"`
}

// issueGetResponseReactionsJSON contains the JSON metadata for the struct
// [IssueGetResponseReactions]
type issueGetResponseReactionsJSON struct {
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

func (r *IssueGetResponseReactions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueGetResponseReactionsJSON) RawJSON() string {
	return r.raw
}

type IssueUpdateResponse struct {
	ID             string                     `json:"id,required" format:"uuid"`
	Funding        IssueUpdateResponseFunding `json:"funding,required"`
	IssueCreatedAt time.Time                  `json:"issue_created_at,required" format:"date-time"`
	// If a maintainer needs to mark this issue as solved
	NeedsConfirmationSolved bool `json:"needs_confirmation_solved,required"`
	// GitHub #number
	Number int64 `json:"number,required"`
	// Issue platform (currently always GitHub)
	Platform IssueUpdateResponsePlatform `json:"platform,required"`
	// If this issue currently has the Polar badge SVG embedded
	PledgeBadgeCurrentlyEmbedded bool `json:"pledge_badge_currently_embedded,required"`
	// The repository that the issue is in
	Repository IssueUpdateResponseRepository `json:"repository,required"`
	State      IssueUpdateResponseState      `json:"state,required"`
	// GitHub issue title
	Title string `json:"title,required"`
	// GitHub assignees
	Assignees []IssueUpdateResponseAssignee `json:"assignees,nullable"`
	// GitHub author
	Author IssueUpdateResponseAuthor `json:"author,nullable"`
	// Optional custom badge SVG promotional content
	BadgeCustomContent string `json:"badge_custom_content,nullable"`
	// GitHub issue body
	Body string `json:"body,nullable"`
	// Number of GitHub comments made on the issue
	Comments int64 `json:"comments,nullable"`
	// If this issue has been marked as confirmed solved through Polar
	ConfirmedSolvedAt time.Time                  `json:"confirmed_solved_at,nullable" format:"date-time"`
	IssueClosedAt     time.Time                  `json:"issue_closed_at,nullable" format:"date-time"`
	IssueModifiedAt   time.Time                  `json:"issue_modified_at,nullable" format:"date-time"`
	Labels            []IssueUpdateResponseLabel `json:"labels"`
	// GitHub reactions
	Reactions IssueUpdateResponseReactions `json:"reactions,nullable"`
	// Share of rewrads that will be rewarded to contributors of this issue. A number
	// between 0 and 100 (inclusive).
	UpfrontSplitToContributors int64                   `json:"upfront_split_to_contributors,nullable"`
	JSON                       issueUpdateResponseJSON `json:"-"`
}

// issueUpdateResponseJSON contains the JSON metadata for the struct
// [IssueUpdateResponse]
type issueUpdateResponseJSON struct {
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

func (r *IssueUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type IssueUpdateResponseFunding struct {
	FundingGoal IssueUpdateResponseFundingFundingGoal `json:"funding_goal,nullable"`
	// Sum of pledges to this isuse (including currently open pledges and pledges that
	// have been paid out). Always in USD.
	PledgesSum IssueUpdateResponseFundingPledgesSum `json:"pledges_sum,nullable"`
	JSON       issueUpdateResponseFundingJSON       `json:"-"`
}

// issueUpdateResponseFundingJSON contains the JSON metadata for the struct
// [IssueUpdateResponseFunding]
type issueUpdateResponseFundingJSON struct {
	FundingGoal apijson.Field
	PledgesSum  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IssueUpdateResponseFunding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueUpdateResponseFundingJSON) RawJSON() string {
	return r.raw
}

type IssueUpdateResponseFundingFundingGoal struct {
	// Amount in the currencies smallest unit (cents if currency is USD)
	Amount int64 `json:"amount,required"`
	// Three letter currency code (eg: USD)
	Currency string                                    `json:"currency,required"`
	JSON     issueUpdateResponseFundingFundingGoalJSON `json:"-"`
}

// issueUpdateResponseFundingFundingGoalJSON contains the JSON metadata for the
// struct [IssueUpdateResponseFundingFundingGoal]
type issueUpdateResponseFundingFundingGoalJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IssueUpdateResponseFundingFundingGoal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueUpdateResponseFundingFundingGoalJSON) RawJSON() string {
	return r.raw
}

// Sum of pledges to this isuse (including currently open pledges and pledges that
// have been paid out). Always in USD.
type IssueUpdateResponseFundingPledgesSum struct {
	// Amount in the currencies smallest unit (cents if currency is USD)
	Amount int64 `json:"amount,required"`
	// Three letter currency code (eg: USD)
	Currency string                                   `json:"currency,required"`
	JSON     issueUpdateResponseFundingPledgesSumJSON `json:"-"`
}

// issueUpdateResponseFundingPledgesSumJSON contains the JSON metadata for the
// struct [IssueUpdateResponseFundingPledgesSum]
type issueUpdateResponseFundingPledgesSumJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IssueUpdateResponseFundingPledgesSum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueUpdateResponseFundingPledgesSumJSON) RawJSON() string {
	return r.raw
}

// Issue platform (currently always GitHub)
type IssueUpdateResponsePlatform string

const (
	IssueUpdateResponsePlatformGitHub IssueUpdateResponsePlatform = "github"
)

func (r IssueUpdateResponsePlatform) IsKnown() bool {
	switch r {
	case IssueUpdateResponsePlatformGitHub:
		return true
	}
	return false
}

// The repository that the issue is in
type IssueUpdateResponseRepository struct {
	ID           string                                    `json:"id,required" format:"uuid"`
	IsPrivate    bool                                      `json:"is_private,required"`
	Name         string                                    `json:"name,required"`
	Organization IssueUpdateResponseRepositoryOrganization `json:"organization,required"`
	Platform     IssueUpdateResponseRepositoryPlatform     `json:"platform,required"`
	// Settings for the repository profile
	ProfileSettings IssueUpdateResponseRepositoryProfileSettings `json:"profile_settings,required,nullable"`
	Description     string                                       `json:"description,nullable"`
	Homepage        string                                       `json:"homepage,nullable"`
	License         string                                       `json:"license,nullable"`
	Stars           int64                                        `json:"stars,nullable"`
	JSON            issueUpdateResponseRepositoryJSON            `json:"-"`
}

// issueUpdateResponseRepositoryJSON contains the JSON metadata for the struct
// [IssueUpdateResponseRepository]
type issueUpdateResponseRepositoryJSON struct {
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

func (r *IssueUpdateResponseRepository) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueUpdateResponseRepositoryJSON) RawJSON() string {
	return r.raw
}

type IssueUpdateResponseRepositoryOrganization struct {
	ID         string                                            `json:"id,required" format:"uuid"`
	AvatarURL  string                                            `json:"avatar_url,required"`
	IsPersonal bool                                              `json:"is_personal,required"`
	Name       string                                            `json:"name,required"`
	Platform   IssueUpdateResponseRepositoryOrganizationPlatform `json:"platform,required"`
	Bio        string                                            `json:"bio,nullable"`
	Blog       string                                            `json:"blog,nullable"`
	Company    string                                            `json:"company,nullable"`
	Email      string                                            `json:"email,nullable"`
	Location   string                                            `json:"location,nullable"`
	// The organization ID.
	OrganizationID  string                                        `json:"organization_id,nullable" format:"uuid4"`
	PrettyName      string                                        `json:"pretty_name,nullable"`
	TwitterUsername string                                        `json:"twitter_username,nullable"`
	JSON            issueUpdateResponseRepositoryOrganizationJSON `json:"-"`
}

// issueUpdateResponseRepositoryOrganizationJSON contains the JSON metadata for the
// struct [IssueUpdateResponseRepositoryOrganization]
type issueUpdateResponseRepositoryOrganizationJSON struct {
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

func (r *IssueUpdateResponseRepositoryOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueUpdateResponseRepositoryOrganizationJSON) RawJSON() string {
	return r.raw
}

type IssueUpdateResponseRepositoryOrganizationPlatform string

const (
	IssueUpdateResponseRepositoryOrganizationPlatformGitHub IssueUpdateResponseRepositoryOrganizationPlatform = "github"
)

func (r IssueUpdateResponseRepositoryOrganizationPlatform) IsKnown() bool {
	switch r {
	case IssueUpdateResponseRepositoryOrganizationPlatformGitHub:
		return true
	}
	return false
}

type IssueUpdateResponseRepositoryPlatform string

const (
	IssueUpdateResponseRepositoryPlatformGitHub IssueUpdateResponseRepositoryPlatform = "github"
)

func (r IssueUpdateResponseRepositoryPlatform) IsKnown() bool {
	switch r {
	case IssueUpdateResponseRepositoryPlatformGitHub:
		return true
	}
	return false
}

// Settings for the repository profile
type IssueUpdateResponseRepositoryProfileSettings struct {
	// A URL to a cover image
	CoverImageURL string `json:"cover_image_url,nullable"`
	// A description of the repository
	Description string `json:"description,nullable"`
	// A list of featured organizations
	FeaturedOrganizations []string `json:"featured_organizations,nullable" format:"uuid4"`
	// A list of highlighted subscription tiers
	HighlightedSubscriptionTiers []string `json:"highlighted_subscription_tiers,nullable" format:"uuid4"`
	// A list of links related to the repository
	Links []string                                         `json:"links,nullable" format:"uri"`
	JSON  issueUpdateResponseRepositoryProfileSettingsJSON `json:"-"`
}

// issueUpdateResponseRepositoryProfileSettingsJSON contains the JSON metadata for
// the struct [IssueUpdateResponseRepositoryProfileSettings]
type issueUpdateResponseRepositoryProfileSettingsJSON struct {
	CoverImageURL                apijson.Field
	Description                  apijson.Field
	FeaturedOrganizations        apijson.Field
	HighlightedSubscriptionTiers apijson.Field
	Links                        apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *IssueUpdateResponseRepositoryProfileSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueUpdateResponseRepositoryProfileSettingsJSON) RawJSON() string {
	return r.raw
}

type IssueUpdateResponseState string

const (
	IssueUpdateResponseStateOpen   IssueUpdateResponseState = "open"
	IssueUpdateResponseStateClosed IssueUpdateResponseState = "closed"
)

func (r IssueUpdateResponseState) IsKnown() bool {
	switch r {
	case IssueUpdateResponseStateOpen, IssueUpdateResponseStateClosed:
		return true
	}
	return false
}

type IssueUpdateResponseAssignee struct {
	ID        int64                           `json:"id,required"`
	AvatarURL string                          `json:"avatar_url,required" format:"uri"`
	HTMLURL   string                          `json:"html_url,required" format:"uri"`
	Login     string                          `json:"login,required"`
	JSON      issueUpdateResponseAssigneeJSON `json:"-"`
}

// issueUpdateResponseAssigneeJSON contains the JSON metadata for the struct
// [IssueUpdateResponseAssignee]
type issueUpdateResponseAssigneeJSON struct {
	ID          apijson.Field
	AvatarURL   apijson.Field
	HTMLURL     apijson.Field
	Login       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IssueUpdateResponseAssignee) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueUpdateResponseAssigneeJSON) RawJSON() string {
	return r.raw
}

// GitHub author
type IssueUpdateResponseAuthor struct {
	ID        int64                         `json:"id,required"`
	AvatarURL string                        `json:"avatar_url,required" format:"uri"`
	HTMLURL   string                        `json:"html_url,required" format:"uri"`
	Login     string                        `json:"login,required"`
	JSON      issueUpdateResponseAuthorJSON `json:"-"`
}

// issueUpdateResponseAuthorJSON contains the JSON metadata for the struct
// [IssueUpdateResponseAuthor]
type issueUpdateResponseAuthorJSON struct {
	ID          apijson.Field
	AvatarURL   apijson.Field
	HTMLURL     apijson.Field
	Login       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IssueUpdateResponseAuthor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueUpdateResponseAuthorJSON) RawJSON() string {
	return r.raw
}

type IssueUpdateResponseLabel struct {
	Color string                       `json:"color,required"`
	Name  string                       `json:"name,required"`
	JSON  issueUpdateResponseLabelJSON `json:"-"`
}

// issueUpdateResponseLabelJSON contains the JSON metadata for the struct
// [IssueUpdateResponseLabel]
type issueUpdateResponseLabelJSON struct {
	Color       apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IssueUpdateResponseLabel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueUpdateResponseLabelJSON) RawJSON() string {
	return r.raw
}

// GitHub reactions
type IssueUpdateResponseReactions struct {
	Confused   int64                            `json:"confused,required"`
	Eyes       int64                            `json:"eyes,required"`
	Heart      int64                            `json:"heart,required"`
	Hooray     int64                            `json:"hooray,required"`
	Laugh      int64                            `json:"laugh,required"`
	MinusOne   int64                            `json:"minus_one,required"`
	PlusOne    int64                            `json:"plus_one,required"`
	Rocket     int64                            `json:"rocket,required"`
	TotalCount int64                            `json:"total_count,required"`
	JSON       issueUpdateResponseReactionsJSON `json:"-"`
}

// issueUpdateResponseReactionsJSON contains the JSON metadata for the struct
// [IssueUpdateResponseReactions]
type issueUpdateResponseReactionsJSON struct {
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

func (r *IssueUpdateResponseReactions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueUpdateResponseReactionsJSON) RawJSON() string {
	return r.raw
}

type IssueListResponse struct {
	Pagination IssueListResponsePagination `json:"pagination,required"`
	Items      []IssueListResponseItem     `json:"items"`
	JSON       issueListResponseJSON       `json:"-"`
}

// issueListResponseJSON contains the JSON metadata for the struct
// [IssueListResponse]
type issueListResponseJSON struct {
	Pagination  apijson.Field
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IssueListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueListResponseJSON) RawJSON() string {
	return r.raw
}

type IssueListResponsePagination struct {
	MaxPage    int64                           `json:"max_page,required"`
	TotalCount int64                           `json:"total_count,required"`
	JSON       issueListResponsePaginationJSON `json:"-"`
}

// issueListResponsePaginationJSON contains the JSON metadata for the struct
// [IssueListResponsePagination]
type issueListResponsePaginationJSON struct {
	MaxPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IssueListResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueListResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type IssueListResponseItem struct {
	ID             string                        `json:"id,required" format:"uuid"`
	Funding        IssueListResponseItemsFunding `json:"funding,required"`
	IssueCreatedAt time.Time                     `json:"issue_created_at,required" format:"date-time"`
	// If a maintainer needs to mark this issue as solved
	NeedsConfirmationSolved bool `json:"needs_confirmation_solved,required"`
	// GitHub #number
	Number int64 `json:"number,required"`
	// Issue platform (currently always GitHub)
	Platform IssueListResponseItemsPlatform `json:"platform,required"`
	// If this issue currently has the Polar badge SVG embedded
	PledgeBadgeCurrentlyEmbedded bool `json:"pledge_badge_currently_embedded,required"`
	// The repository that the issue is in
	Repository IssueListResponseItemsRepository `json:"repository,required"`
	State      IssueListResponseItemsState      `json:"state,required"`
	// GitHub issue title
	Title string `json:"title,required"`
	// GitHub assignees
	Assignees []IssueListResponseItemsAssignee `json:"assignees,nullable"`
	// GitHub author
	Author IssueListResponseItemsAuthor `json:"author,nullable"`
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
	Labels            []IssueListResponseItemsLabel `json:"labels"`
	// GitHub reactions
	Reactions IssueListResponseItemsReactions `json:"reactions,nullable"`
	// Share of rewrads that will be rewarded to contributors of this issue. A number
	// between 0 and 100 (inclusive).
	UpfrontSplitToContributors int64                     `json:"upfront_split_to_contributors,nullable"`
	JSON                       issueListResponseItemJSON `json:"-"`
}

// issueListResponseItemJSON contains the JSON metadata for the struct
// [IssueListResponseItem]
type issueListResponseItemJSON struct {
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

func (r *IssueListResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueListResponseItemJSON) RawJSON() string {
	return r.raw
}

type IssueListResponseItemsFunding struct {
	FundingGoal IssueListResponseItemsFundingFundingGoal `json:"funding_goal,nullable"`
	// Sum of pledges to this isuse (including currently open pledges and pledges that
	// have been paid out). Always in USD.
	PledgesSum IssueListResponseItemsFundingPledgesSum `json:"pledges_sum,nullable"`
	JSON       issueListResponseItemsFundingJSON       `json:"-"`
}

// issueListResponseItemsFundingJSON contains the JSON metadata for the struct
// [IssueListResponseItemsFunding]
type issueListResponseItemsFundingJSON struct {
	FundingGoal apijson.Field
	PledgesSum  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IssueListResponseItemsFunding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueListResponseItemsFundingJSON) RawJSON() string {
	return r.raw
}

type IssueListResponseItemsFundingFundingGoal struct {
	// Amount in the currencies smallest unit (cents if currency is USD)
	Amount int64 `json:"amount,required"`
	// Three letter currency code (eg: USD)
	Currency string                                       `json:"currency,required"`
	JSON     issueListResponseItemsFundingFundingGoalJSON `json:"-"`
}

// issueListResponseItemsFundingFundingGoalJSON contains the JSON metadata for the
// struct [IssueListResponseItemsFundingFundingGoal]
type issueListResponseItemsFundingFundingGoalJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IssueListResponseItemsFundingFundingGoal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueListResponseItemsFundingFundingGoalJSON) RawJSON() string {
	return r.raw
}

// Sum of pledges to this isuse (including currently open pledges and pledges that
// have been paid out). Always in USD.
type IssueListResponseItemsFundingPledgesSum struct {
	// Amount in the currencies smallest unit (cents if currency is USD)
	Amount int64 `json:"amount,required"`
	// Three letter currency code (eg: USD)
	Currency string                                      `json:"currency,required"`
	JSON     issueListResponseItemsFundingPledgesSumJSON `json:"-"`
}

// issueListResponseItemsFundingPledgesSumJSON contains the JSON metadata for the
// struct [IssueListResponseItemsFundingPledgesSum]
type issueListResponseItemsFundingPledgesSumJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IssueListResponseItemsFundingPledgesSum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueListResponseItemsFundingPledgesSumJSON) RawJSON() string {
	return r.raw
}

// Issue platform (currently always GitHub)
type IssueListResponseItemsPlatform string

const (
	IssueListResponseItemsPlatformGitHub IssueListResponseItemsPlatform = "github"
)

func (r IssueListResponseItemsPlatform) IsKnown() bool {
	switch r {
	case IssueListResponseItemsPlatformGitHub:
		return true
	}
	return false
}

// The repository that the issue is in
type IssueListResponseItemsRepository struct {
	ID           string                                       `json:"id,required" format:"uuid"`
	IsPrivate    bool                                         `json:"is_private,required"`
	Name         string                                       `json:"name,required"`
	Organization IssueListResponseItemsRepositoryOrganization `json:"organization,required"`
	Platform     IssueListResponseItemsRepositoryPlatform     `json:"platform,required"`
	// Settings for the repository profile
	ProfileSettings IssueListResponseItemsRepositoryProfileSettings `json:"profile_settings,required,nullable"`
	Description     string                                          `json:"description,nullable"`
	Homepage        string                                          `json:"homepage,nullable"`
	License         string                                          `json:"license,nullable"`
	Stars           int64                                           `json:"stars,nullable"`
	JSON            issueListResponseItemsRepositoryJSON            `json:"-"`
}

// issueListResponseItemsRepositoryJSON contains the JSON metadata for the struct
// [IssueListResponseItemsRepository]
type issueListResponseItemsRepositoryJSON struct {
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

func (r *IssueListResponseItemsRepository) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueListResponseItemsRepositoryJSON) RawJSON() string {
	return r.raw
}

type IssueListResponseItemsRepositoryOrganization struct {
	ID         string                                               `json:"id,required" format:"uuid"`
	AvatarURL  string                                               `json:"avatar_url,required"`
	IsPersonal bool                                                 `json:"is_personal,required"`
	Name       string                                               `json:"name,required"`
	Platform   IssueListResponseItemsRepositoryOrganizationPlatform `json:"platform,required"`
	Bio        string                                               `json:"bio,nullable"`
	Blog       string                                               `json:"blog,nullable"`
	Company    string                                               `json:"company,nullable"`
	Email      string                                               `json:"email,nullable"`
	Location   string                                               `json:"location,nullable"`
	// The organization ID.
	OrganizationID  string                                           `json:"organization_id,nullable" format:"uuid4"`
	PrettyName      string                                           `json:"pretty_name,nullable"`
	TwitterUsername string                                           `json:"twitter_username,nullable"`
	JSON            issueListResponseItemsRepositoryOrganizationJSON `json:"-"`
}

// issueListResponseItemsRepositoryOrganizationJSON contains the JSON metadata for
// the struct [IssueListResponseItemsRepositoryOrganization]
type issueListResponseItemsRepositoryOrganizationJSON struct {
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

func (r *IssueListResponseItemsRepositoryOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueListResponseItemsRepositoryOrganizationJSON) RawJSON() string {
	return r.raw
}

type IssueListResponseItemsRepositoryOrganizationPlatform string

const (
	IssueListResponseItemsRepositoryOrganizationPlatformGitHub IssueListResponseItemsRepositoryOrganizationPlatform = "github"
)

func (r IssueListResponseItemsRepositoryOrganizationPlatform) IsKnown() bool {
	switch r {
	case IssueListResponseItemsRepositoryOrganizationPlatformGitHub:
		return true
	}
	return false
}

type IssueListResponseItemsRepositoryPlatform string

const (
	IssueListResponseItemsRepositoryPlatformGitHub IssueListResponseItemsRepositoryPlatform = "github"
)

func (r IssueListResponseItemsRepositoryPlatform) IsKnown() bool {
	switch r {
	case IssueListResponseItemsRepositoryPlatformGitHub:
		return true
	}
	return false
}

// Settings for the repository profile
type IssueListResponseItemsRepositoryProfileSettings struct {
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
	JSON  issueListResponseItemsRepositoryProfileSettingsJSON `json:"-"`
}

// issueListResponseItemsRepositoryProfileSettingsJSON contains the JSON metadata
// for the struct [IssueListResponseItemsRepositoryProfileSettings]
type issueListResponseItemsRepositoryProfileSettingsJSON struct {
	CoverImageURL                apijson.Field
	Description                  apijson.Field
	FeaturedOrganizations        apijson.Field
	HighlightedSubscriptionTiers apijson.Field
	Links                        apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *IssueListResponseItemsRepositoryProfileSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueListResponseItemsRepositoryProfileSettingsJSON) RawJSON() string {
	return r.raw
}

type IssueListResponseItemsState string

const (
	IssueListResponseItemsStateOpen   IssueListResponseItemsState = "open"
	IssueListResponseItemsStateClosed IssueListResponseItemsState = "closed"
)

func (r IssueListResponseItemsState) IsKnown() bool {
	switch r {
	case IssueListResponseItemsStateOpen, IssueListResponseItemsStateClosed:
		return true
	}
	return false
}

type IssueListResponseItemsAssignee struct {
	ID        int64                              `json:"id,required"`
	AvatarURL string                             `json:"avatar_url,required" format:"uri"`
	HTMLURL   string                             `json:"html_url,required" format:"uri"`
	Login     string                             `json:"login,required"`
	JSON      issueListResponseItemsAssigneeJSON `json:"-"`
}

// issueListResponseItemsAssigneeJSON contains the JSON metadata for the struct
// [IssueListResponseItemsAssignee]
type issueListResponseItemsAssigneeJSON struct {
	ID          apijson.Field
	AvatarURL   apijson.Field
	HTMLURL     apijson.Field
	Login       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IssueListResponseItemsAssignee) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueListResponseItemsAssigneeJSON) RawJSON() string {
	return r.raw
}

// GitHub author
type IssueListResponseItemsAuthor struct {
	ID        int64                            `json:"id,required"`
	AvatarURL string                           `json:"avatar_url,required" format:"uri"`
	HTMLURL   string                           `json:"html_url,required" format:"uri"`
	Login     string                           `json:"login,required"`
	JSON      issueListResponseItemsAuthorJSON `json:"-"`
}

// issueListResponseItemsAuthorJSON contains the JSON metadata for the struct
// [IssueListResponseItemsAuthor]
type issueListResponseItemsAuthorJSON struct {
	ID          apijson.Field
	AvatarURL   apijson.Field
	HTMLURL     apijson.Field
	Login       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IssueListResponseItemsAuthor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueListResponseItemsAuthorJSON) RawJSON() string {
	return r.raw
}

type IssueListResponseItemsLabel struct {
	Color string                          `json:"color,required"`
	Name  string                          `json:"name,required"`
	JSON  issueListResponseItemsLabelJSON `json:"-"`
}

// issueListResponseItemsLabelJSON contains the JSON metadata for the struct
// [IssueListResponseItemsLabel]
type issueListResponseItemsLabelJSON struct {
	Color       apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IssueListResponseItemsLabel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueListResponseItemsLabelJSON) RawJSON() string {
	return r.raw
}

// GitHub reactions
type IssueListResponseItemsReactions struct {
	Confused   int64                               `json:"confused,required"`
	Eyes       int64                               `json:"eyes,required"`
	Heart      int64                               `json:"heart,required"`
	Hooray     int64                               `json:"hooray,required"`
	Laugh      int64                               `json:"laugh,required"`
	MinusOne   int64                               `json:"minus_one,required"`
	PlusOne    int64                               `json:"plus_one,required"`
	Rocket     int64                               `json:"rocket,required"`
	TotalCount int64                               `json:"total_count,required"`
	JSON       issueListResponseItemsReactionsJSON `json:"-"`
}

// issueListResponseItemsReactionsJSON contains the JSON metadata for the struct
// [IssueListResponseItemsReactions]
type issueListResponseItemsReactionsJSON struct {
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

func (r *IssueListResponseItemsReactions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueListResponseItemsReactionsJSON) RawJSON() string {
	return r.raw
}

type IssueConfirmSolvedResponse struct {
	ID             string                            `json:"id,required" format:"uuid"`
	Funding        IssueConfirmSolvedResponseFunding `json:"funding,required"`
	IssueCreatedAt time.Time                         `json:"issue_created_at,required" format:"date-time"`
	// If a maintainer needs to mark this issue as solved
	NeedsConfirmationSolved bool `json:"needs_confirmation_solved,required"`
	// GitHub #number
	Number int64 `json:"number,required"`
	// Issue platform (currently always GitHub)
	Platform IssueConfirmSolvedResponsePlatform `json:"platform,required"`
	// If this issue currently has the Polar badge SVG embedded
	PledgeBadgeCurrentlyEmbedded bool `json:"pledge_badge_currently_embedded,required"`
	// The repository that the issue is in
	Repository IssueConfirmSolvedResponseRepository `json:"repository,required"`
	State      IssueConfirmSolvedResponseState      `json:"state,required"`
	// GitHub issue title
	Title string `json:"title,required"`
	// GitHub assignees
	Assignees []IssueConfirmSolvedResponseAssignee `json:"assignees,nullable"`
	// GitHub author
	Author IssueConfirmSolvedResponseAuthor `json:"author,nullable"`
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
	Labels            []IssueConfirmSolvedResponseLabel `json:"labels"`
	// GitHub reactions
	Reactions IssueConfirmSolvedResponseReactions `json:"reactions,nullable"`
	// Share of rewrads that will be rewarded to contributors of this issue. A number
	// between 0 and 100 (inclusive).
	UpfrontSplitToContributors int64                          `json:"upfront_split_to_contributors,nullable"`
	JSON                       issueConfirmSolvedResponseJSON `json:"-"`
}

// issueConfirmSolvedResponseJSON contains the JSON metadata for the struct
// [IssueConfirmSolvedResponse]
type issueConfirmSolvedResponseJSON struct {
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

func (r *IssueConfirmSolvedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueConfirmSolvedResponseJSON) RawJSON() string {
	return r.raw
}

type IssueConfirmSolvedResponseFunding struct {
	FundingGoal IssueConfirmSolvedResponseFundingFundingGoal `json:"funding_goal,nullable"`
	// Sum of pledges to this isuse (including currently open pledges and pledges that
	// have been paid out). Always in USD.
	PledgesSum IssueConfirmSolvedResponseFundingPledgesSum `json:"pledges_sum,nullable"`
	JSON       issueConfirmSolvedResponseFundingJSON       `json:"-"`
}

// issueConfirmSolvedResponseFundingJSON contains the JSON metadata for the struct
// [IssueConfirmSolvedResponseFunding]
type issueConfirmSolvedResponseFundingJSON struct {
	FundingGoal apijson.Field
	PledgesSum  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IssueConfirmSolvedResponseFunding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueConfirmSolvedResponseFundingJSON) RawJSON() string {
	return r.raw
}

type IssueConfirmSolvedResponseFundingFundingGoal struct {
	// Amount in the currencies smallest unit (cents if currency is USD)
	Amount int64 `json:"amount,required"`
	// Three letter currency code (eg: USD)
	Currency string                                           `json:"currency,required"`
	JSON     issueConfirmSolvedResponseFundingFundingGoalJSON `json:"-"`
}

// issueConfirmSolvedResponseFundingFundingGoalJSON contains the JSON metadata for
// the struct [IssueConfirmSolvedResponseFundingFundingGoal]
type issueConfirmSolvedResponseFundingFundingGoalJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IssueConfirmSolvedResponseFundingFundingGoal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueConfirmSolvedResponseFundingFundingGoalJSON) RawJSON() string {
	return r.raw
}

// Sum of pledges to this isuse (including currently open pledges and pledges that
// have been paid out). Always in USD.
type IssueConfirmSolvedResponseFundingPledgesSum struct {
	// Amount in the currencies smallest unit (cents if currency is USD)
	Amount int64 `json:"amount,required"`
	// Three letter currency code (eg: USD)
	Currency string                                          `json:"currency,required"`
	JSON     issueConfirmSolvedResponseFundingPledgesSumJSON `json:"-"`
}

// issueConfirmSolvedResponseFundingPledgesSumJSON contains the JSON metadata for
// the struct [IssueConfirmSolvedResponseFundingPledgesSum]
type issueConfirmSolvedResponseFundingPledgesSumJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IssueConfirmSolvedResponseFundingPledgesSum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueConfirmSolvedResponseFundingPledgesSumJSON) RawJSON() string {
	return r.raw
}

// Issue platform (currently always GitHub)
type IssueConfirmSolvedResponsePlatform string

const (
	IssueConfirmSolvedResponsePlatformGitHub IssueConfirmSolvedResponsePlatform = "github"
)

func (r IssueConfirmSolvedResponsePlatform) IsKnown() bool {
	switch r {
	case IssueConfirmSolvedResponsePlatformGitHub:
		return true
	}
	return false
}

// The repository that the issue is in
type IssueConfirmSolvedResponseRepository struct {
	ID           string                                           `json:"id,required" format:"uuid"`
	IsPrivate    bool                                             `json:"is_private,required"`
	Name         string                                           `json:"name,required"`
	Organization IssueConfirmSolvedResponseRepositoryOrganization `json:"organization,required"`
	Platform     IssueConfirmSolvedResponseRepositoryPlatform     `json:"platform,required"`
	// Settings for the repository profile
	ProfileSettings IssueConfirmSolvedResponseRepositoryProfileSettings `json:"profile_settings,required,nullable"`
	Description     string                                              `json:"description,nullable"`
	Homepage        string                                              `json:"homepage,nullable"`
	License         string                                              `json:"license,nullable"`
	Stars           int64                                               `json:"stars,nullable"`
	JSON            issueConfirmSolvedResponseRepositoryJSON            `json:"-"`
}

// issueConfirmSolvedResponseRepositoryJSON contains the JSON metadata for the
// struct [IssueConfirmSolvedResponseRepository]
type issueConfirmSolvedResponseRepositoryJSON struct {
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

func (r *IssueConfirmSolvedResponseRepository) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueConfirmSolvedResponseRepositoryJSON) RawJSON() string {
	return r.raw
}

type IssueConfirmSolvedResponseRepositoryOrganization struct {
	ID         string                                                   `json:"id,required" format:"uuid"`
	AvatarURL  string                                                   `json:"avatar_url,required"`
	IsPersonal bool                                                     `json:"is_personal,required"`
	Name       string                                                   `json:"name,required"`
	Platform   IssueConfirmSolvedResponseRepositoryOrganizationPlatform `json:"platform,required"`
	Bio        string                                                   `json:"bio,nullable"`
	Blog       string                                                   `json:"blog,nullable"`
	Company    string                                                   `json:"company,nullable"`
	Email      string                                                   `json:"email,nullable"`
	Location   string                                                   `json:"location,nullable"`
	// The organization ID.
	OrganizationID  string                                               `json:"organization_id,nullable" format:"uuid4"`
	PrettyName      string                                               `json:"pretty_name,nullable"`
	TwitterUsername string                                               `json:"twitter_username,nullable"`
	JSON            issueConfirmSolvedResponseRepositoryOrganizationJSON `json:"-"`
}

// issueConfirmSolvedResponseRepositoryOrganizationJSON contains the JSON metadata
// for the struct [IssueConfirmSolvedResponseRepositoryOrganization]
type issueConfirmSolvedResponseRepositoryOrganizationJSON struct {
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

func (r *IssueConfirmSolvedResponseRepositoryOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueConfirmSolvedResponseRepositoryOrganizationJSON) RawJSON() string {
	return r.raw
}

type IssueConfirmSolvedResponseRepositoryOrganizationPlatform string

const (
	IssueConfirmSolvedResponseRepositoryOrganizationPlatformGitHub IssueConfirmSolvedResponseRepositoryOrganizationPlatform = "github"
)

func (r IssueConfirmSolvedResponseRepositoryOrganizationPlatform) IsKnown() bool {
	switch r {
	case IssueConfirmSolvedResponseRepositoryOrganizationPlatformGitHub:
		return true
	}
	return false
}

type IssueConfirmSolvedResponseRepositoryPlatform string

const (
	IssueConfirmSolvedResponseRepositoryPlatformGitHub IssueConfirmSolvedResponseRepositoryPlatform = "github"
)

func (r IssueConfirmSolvedResponseRepositoryPlatform) IsKnown() bool {
	switch r {
	case IssueConfirmSolvedResponseRepositoryPlatformGitHub:
		return true
	}
	return false
}

// Settings for the repository profile
type IssueConfirmSolvedResponseRepositoryProfileSettings struct {
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
	JSON  issueConfirmSolvedResponseRepositoryProfileSettingsJSON `json:"-"`
}

// issueConfirmSolvedResponseRepositoryProfileSettingsJSON contains the JSON
// metadata for the struct [IssueConfirmSolvedResponseRepositoryProfileSettings]
type issueConfirmSolvedResponseRepositoryProfileSettingsJSON struct {
	CoverImageURL                apijson.Field
	Description                  apijson.Field
	FeaturedOrganizations        apijson.Field
	HighlightedSubscriptionTiers apijson.Field
	Links                        apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *IssueConfirmSolvedResponseRepositoryProfileSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueConfirmSolvedResponseRepositoryProfileSettingsJSON) RawJSON() string {
	return r.raw
}

type IssueConfirmSolvedResponseState string

const (
	IssueConfirmSolvedResponseStateOpen   IssueConfirmSolvedResponseState = "open"
	IssueConfirmSolvedResponseStateClosed IssueConfirmSolvedResponseState = "closed"
)

func (r IssueConfirmSolvedResponseState) IsKnown() bool {
	switch r {
	case IssueConfirmSolvedResponseStateOpen, IssueConfirmSolvedResponseStateClosed:
		return true
	}
	return false
}

type IssueConfirmSolvedResponseAssignee struct {
	ID        int64                                  `json:"id,required"`
	AvatarURL string                                 `json:"avatar_url,required" format:"uri"`
	HTMLURL   string                                 `json:"html_url,required" format:"uri"`
	Login     string                                 `json:"login,required"`
	JSON      issueConfirmSolvedResponseAssigneeJSON `json:"-"`
}

// issueConfirmSolvedResponseAssigneeJSON contains the JSON metadata for the struct
// [IssueConfirmSolvedResponseAssignee]
type issueConfirmSolvedResponseAssigneeJSON struct {
	ID          apijson.Field
	AvatarURL   apijson.Field
	HTMLURL     apijson.Field
	Login       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IssueConfirmSolvedResponseAssignee) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueConfirmSolvedResponseAssigneeJSON) RawJSON() string {
	return r.raw
}

// GitHub author
type IssueConfirmSolvedResponseAuthor struct {
	ID        int64                                `json:"id,required"`
	AvatarURL string                               `json:"avatar_url,required" format:"uri"`
	HTMLURL   string                               `json:"html_url,required" format:"uri"`
	Login     string                               `json:"login,required"`
	JSON      issueConfirmSolvedResponseAuthorJSON `json:"-"`
}

// issueConfirmSolvedResponseAuthorJSON contains the JSON metadata for the struct
// [IssueConfirmSolvedResponseAuthor]
type issueConfirmSolvedResponseAuthorJSON struct {
	ID          apijson.Field
	AvatarURL   apijson.Field
	HTMLURL     apijson.Field
	Login       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IssueConfirmSolvedResponseAuthor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueConfirmSolvedResponseAuthorJSON) RawJSON() string {
	return r.raw
}

type IssueConfirmSolvedResponseLabel struct {
	Color string                              `json:"color,required"`
	Name  string                              `json:"name,required"`
	JSON  issueConfirmSolvedResponseLabelJSON `json:"-"`
}

// issueConfirmSolvedResponseLabelJSON contains the JSON metadata for the struct
// [IssueConfirmSolvedResponseLabel]
type issueConfirmSolvedResponseLabelJSON struct {
	Color       apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IssueConfirmSolvedResponseLabel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueConfirmSolvedResponseLabelJSON) RawJSON() string {
	return r.raw
}

// GitHub reactions
type IssueConfirmSolvedResponseReactions struct {
	Confused   int64                                   `json:"confused,required"`
	Eyes       int64                                   `json:"eyes,required"`
	Heart      int64                                   `json:"heart,required"`
	Hooray     int64                                   `json:"hooray,required"`
	Laugh      int64                                   `json:"laugh,required"`
	MinusOne   int64                                   `json:"minus_one,required"`
	PlusOne    int64                                   `json:"plus_one,required"`
	Rocket     int64                                   `json:"rocket,required"`
	TotalCount int64                                   `json:"total_count,required"`
	JSON       issueConfirmSolvedResponseReactionsJSON `json:"-"`
}

// issueConfirmSolvedResponseReactionsJSON contains the JSON metadata for the
// struct [IssueConfirmSolvedResponseReactions]
type issueConfirmSolvedResponseReactionsJSON struct {
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

func (r *IssueConfirmSolvedResponseReactions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueConfirmSolvedResponseReactionsJSON) RawJSON() string {
	return r.raw
}

type IssueUpdateParams struct {
	FundingGoal                   param.Field[IssueUpdateParamsFundingGoal] `json:"funding_goal"`
	SetUpfrontSplitToContributors param.Field[bool]                         `json:"set_upfront_split_to_contributors"`
	UpfrontSplitToContributors    param.Field[int64]                        `json:"upfront_split_to_contributors"`
}

func (r IssueUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IssueUpdateParamsFundingGoal struct {
	// Amount in the currencies smallest unit (cents if currency is USD)
	Amount param.Field[int64] `json:"amount,required"`
	// Three letter currency code (eg: USD)
	Currency param.Field[string] `json:"currency,required"`
}

func (r IssueUpdateParamsFundingGoal) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IssueListParams struct {
	// Filter by external organization name.
	ExternalOrganizationName param.Field[IssueListParamsExternalOrganizationNameUnion] `query:"external_organization_name"`
	// Filter by badged status.
	IsBadged param.Field[bool] `query:"is_badged"`
	// Size of a page, defaults to 10. Maximum is 100.
	Limit param.Field[int64] `query:"limit"`
	// Filter by issue number.
	Number param.Field[IssueListParamsNumberUnion] `query:"number"`
	// Filter by organization ID.
	OrganizationID param.Field[IssueListParamsOrganizationIDUnion] `query:"organization_id" format:"uuid4"`
	// Page number, defaults to 1.
	Page param.Field[int64] `query:"page"`
	// Filter by platform.
	Platform param.Field[IssueListParamsPlatformUnion] `query:"platform"`
	// Filter by repository name.
	RepositoryName param.Field[IssueListParamsRepositoryNameUnion] `query:"repository_name"`
	// Sorting criterion. Several criteria can be used simultaneously and will be
	// applied in order. Add a minus sign `-` before the criteria name to sort by
	// descending order.
	Sorting param.Field[[]string] `query:"sorting"`
}

// URLQuery serializes [IssueListParams]'s query parameters as `url.Values`.
func (r IssueListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by external organization name.
//
// Satisfied by [shared.UnionString],
// [IssueListParamsExternalOrganizationNameArray].
type IssueListParamsExternalOrganizationNameUnion interface {
	ImplementsIssueListParamsExternalOrganizationNameUnion()
}

type IssueListParamsExternalOrganizationNameArray []string

func (r IssueListParamsExternalOrganizationNameArray) ImplementsIssueListParamsExternalOrganizationNameUnion() {
}

// Filter by issue number.
//
// Satisfied by [shared.UnionInt], [IssueListParamsNumberArray].
type IssueListParamsNumberUnion interface {
	ImplementsIssueListParamsNumberUnion()
}

type IssueListParamsNumberArray []int64

func (r IssueListParamsNumberArray) ImplementsIssueListParamsNumberUnion() {}

// Filter by organization ID.
//
// Satisfied by [shared.UnionString], [IssueListParamsOrganizationIDArray].
type IssueListParamsOrganizationIDUnion interface {
	ImplementsIssueListParamsOrganizationIDUnion()
}

type IssueListParamsOrganizationIDArray []string

func (r IssueListParamsOrganizationIDArray) ImplementsIssueListParamsOrganizationIDUnion() {}

// Filter by platform.
//
// Satisfied by [IssueListParamsPlatformPlatforms], [IssueListParamsPlatformArray].
type IssueListParamsPlatformUnion interface {
	implementsIssueListParamsPlatformUnion()
}

type IssueListParamsPlatformPlatforms string

const (
	IssueListParamsPlatformPlatformsGitHub IssueListParamsPlatformPlatforms = "github"
)

func (r IssueListParamsPlatformPlatforms) IsKnown() bool {
	switch r {
	case IssueListParamsPlatformPlatformsGitHub:
		return true
	}
	return false
}

func (r IssueListParamsPlatformPlatforms) implementsIssueListParamsPlatformUnion() {}

type IssueListParamsPlatformArray []IssueListParamsPlatformArray

func (r IssueListParamsPlatformArray) implementsIssueListParamsPlatformUnion() {}

// Filter by repository name.
//
// Satisfied by [shared.UnionString], [IssueListParamsRepositoryNameArray].
type IssueListParamsRepositoryNameUnion interface {
	ImplementsIssueListParamsRepositoryNameUnion()
}

type IssueListParamsRepositoryNameArray []string

func (r IssueListParamsRepositoryNameArray) ImplementsIssueListParamsRepositoryNameUnion() {}

type IssueConfirmSolvedParams struct {
	Splits param.Field[[]IssueConfirmSolvedParamsSplit] `json:"splits,required"`
}

func (r IssueConfirmSolvedParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IssueConfirmSolvedParamsSplit struct {
	ShareThousands param.Field[int64]  `json:"share_thousands,required"`
	GitHubUsername param.Field[string] `json:"github_username"`
	OrganizationID param.Field[string] `json:"organization_id" format:"uuid"`
}

func (r IssueConfirmSolvedParamsSplit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
