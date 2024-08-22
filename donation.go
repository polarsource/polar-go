// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package polar

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/stainless-sdks/polar-go/internal/apijson"
	"github.com/stainless-sdks/polar-go/internal/apiquery"
	"github.com/stainless-sdks/polar-go/internal/param"
	"github.com/stainless-sdks/polar-go/internal/requestconfig"
	"github.com/stainless-sdks/polar-go/option"
	"github.com/tidwall/gjson"
)

// DonationService contains methods and other services that help with interacting
// with the polar API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDonationService] method instead.
type DonationService struct {
	Options       []option.RequestOption
	PaymentIntent *DonationPaymentIntentService
	Public        *DonationPublicService
}

// NewDonationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDonationService(opts ...option.RequestOption) (r *DonationService) {
	r = &DonationService{}
	r.Options = opts
	r.PaymentIntent = NewDonationPaymentIntentService(opts...)
	r.Public = NewDonationPublicService(opts...)
	return
}

// Search Donations
func (r *DonationService) Search(ctx context.Context, query DonationSearchParams, opts ...option.RequestOption) (res *ListResourceDonation, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/donations/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Statistics
func (r *DonationService) Statistics(ctx context.Context, query DonationStatisticsParams, opts ...option.RequestOption) (res *DonationStatistics, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/donations/statistics"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type DonationStatistics struct {
	Periods []DonationStatisticsPeriod `json:"periods,required"`
	JSON    donationStatisticsJSON     `json:"-"`
}

// donationStatisticsJSON contains the JSON metadata for the struct
// [DonationStatistics]
type donationStatisticsJSON struct {
	Periods     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DonationStatistics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r donationStatisticsJSON) RawJSON() string {
	return r.raw
}

type DonationStatisticsPeriod struct {
	EndDate   time.Time                    `json:"end_date,required" format:"date"`
	StartDate time.Time                    `json:"start_date,required" format:"date"`
	Sum       int64                        `json:"sum,required"`
	JSON      donationStatisticsPeriodJSON `json:"-"`
}

// donationStatisticsPeriodJSON contains the JSON metadata for the struct
// [DonationStatisticsPeriod]
type donationStatisticsPeriodJSON struct {
	EndDate     apijson.Field
	StartDate   apijson.Field
	Sum         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DonationStatisticsPeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r donationStatisticsPeriodJSON) RawJSON() string {
	return r.raw
}

type ListResourceDonation struct {
	Pagination ListResourceDonationPagination `json:"pagination,required"`
	Items      []ListResourceDonationItem     `json:"items"`
	JSON       listResourceDonationJSON       `json:"-"`
}

// listResourceDonationJSON contains the JSON metadata for the struct
// [ListResourceDonation]
type listResourceDonationJSON struct {
	Pagination  apijson.Field
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListResourceDonation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResourceDonationJSON) RawJSON() string {
	return r.raw
}

type ListResourceDonationPagination struct {
	MaxPage    int64                              `json:"max_page,required"`
	TotalCount int64                              `json:"total_count,required"`
	JSON       listResourceDonationPaginationJSON `json:"-"`
}

// listResourceDonationPaginationJSON contains the JSON metadata for the struct
// [ListResourceDonationPagination]
type listResourceDonationPaginationJSON struct {
	MaxPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListResourceDonationPagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResourceDonationPaginationJSON) RawJSON() string {
	return r.raw
}

type ListResourceDonationItem struct {
	// The ID of the object.
	ID     string `json:"id,required" format:"uuid4"`
	Amount int64  `json:"amount,required"`
	// Creation timestamp of the object.
	CreatedAt time.Time                      `json:"created_at,required" format:"date-time"`
	Currency  string                         `json:"currency,required"`
	Email     string                         `json:"email,required"`
	Message   string                         `json:"message,required,nullable"`
	Donor     ListResourceDonationItemsDonor `json:"donor,nullable"`
	Issue     ListResourceDonationItemsIssue `json:"issue,nullable"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time                    `json:"modified_at,nullable" format:"date-time"`
	JSON       listResourceDonationItemJSON `json:"-"`
}

// listResourceDonationItemJSON contains the JSON metadata for the struct
// [ListResourceDonationItem]
type listResourceDonationItemJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	CreatedAt   apijson.Field
	Currency    apijson.Field
	Email       apijson.Field
	Message     apijson.Field
	Donor       apijson.Field
	Issue       apijson.Field
	ModifiedAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListResourceDonationItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResourceDonationItemJSON) RawJSON() string {
	return r.raw
}

type ListResourceDonationItemsDonor struct {
	ID         string                                 `json:"id,required" format:"uuid4"`
	Platform   ListResourceDonationItemsDonorPlatform `json:"platform"`
	Name       string                                 `json:"name"`
	AvatarURL  string                                 `json:"avatar_url,required"`
	IsPersonal bool                                   `json:"is_personal"`
	PublicName string                                 `json:"public_name"`
	JSON       listResourceDonationItemsDonorJSON     `json:"-"`
	union      ListResourceDonationItemsDonorUnion
}

// listResourceDonationItemsDonorJSON contains the JSON metadata for the struct
// [ListResourceDonationItemsDonor]
type listResourceDonationItemsDonorJSON struct {
	ID          apijson.Field
	Platform    apijson.Field
	Name        apijson.Field
	AvatarURL   apijson.Field
	IsPersonal  apijson.Field
	PublicName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r listResourceDonationItemsDonorJSON) RawJSON() string {
	return r.raw
}

func (r *ListResourceDonationItemsDonor) UnmarshalJSON(data []byte) (err error) {
	*r = ListResourceDonationItemsDonor{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ListResourceDonationItemsDonorUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ListResourceDonationItemsDonorDonationOrganization],
// [ListResourceDonationItemsDonorDonationUser].
func (r ListResourceDonationItemsDonor) AsUnion() ListResourceDonationItemsDonorUnion {
	return r.union
}

// Union satisfied by [ListResourceDonationItemsDonorDonationOrganization] or
// [ListResourceDonationItemsDonorDonationUser].
type ListResourceDonationItemsDonorUnion interface {
	implementsListResourceDonationItemsDonor()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ListResourceDonationItemsDonorUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ListResourceDonationItemsDonorDonationOrganization{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ListResourceDonationItemsDonorDonationUser{}),
		},
	)
}

type ListResourceDonationItemsDonorDonationOrganization struct {
	ID         string                                                     `json:"id,required" format:"uuid4"`
	AvatarURL  string                                                     `json:"avatar_url,required"`
	IsPersonal bool                                                       `json:"is_personal,required"`
	Name       string                                                     `json:"name,required"`
	Platform   ListResourceDonationItemsDonorDonationOrganizationPlatform `json:"platform,required"`
	JSON       listResourceDonationItemsDonorDonationOrganizationJSON     `json:"-"`
}

// listResourceDonationItemsDonorDonationOrganizationJSON contains the JSON
// metadata for the struct [ListResourceDonationItemsDonorDonationOrganization]
type listResourceDonationItemsDonorDonationOrganizationJSON struct {
	ID          apijson.Field
	AvatarURL   apijson.Field
	IsPersonal  apijson.Field
	Name        apijson.Field
	Platform    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListResourceDonationItemsDonorDonationOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResourceDonationItemsDonorDonationOrganizationJSON) RawJSON() string {
	return r.raw
}

func (r ListResourceDonationItemsDonorDonationOrganization) implementsListResourceDonationItemsDonor() {
}

type ListResourceDonationItemsDonorDonationOrganizationPlatform string

const (
	ListResourceDonationItemsDonorDonationOrganizationPlatformGitHub ListResourceDonationItemsDonorDonationOrganizationPlatform = "github"
)

func (r ListResourceDonationItemsDonorDonationOrganizationPlatform) IsKnown() bool {
	switch r {
	case ListResourceDonationItemsDonorDonationOrganizationPlatformGitHub:
		return true
	}
	return false
}

type ListResourceDonationItemsDonorDonationUser struct {
	ID         string                                         `json:"id,required" format:"uuid4"`
	AvatarURL  string                                         `json:"avatar_url,required,nullable"`
	PublicName string                                         `json:"public_name,required"`
	JSON       listResourceDonationItemsDonorDonationUserJSON `json:"-"`
}

// listResourceDonationItemsDonorDonationUserJSON contains the JSON metadata for
// the struct [ListResourceDonationItemsDonorDonationUser]
type listResourceDonationItemsDonorDonationUserJSON struct {
	ID          apijson.Field
	AvatarURL   apijson.Field
	PublicName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListResourceDonationItemsDonorDonationUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResourceDonationItemsDonorDonationUserJSON) RawJSON() string {
	return r.raw
}

func (r ListResourceDonationItemsDonorDonationUser) implementsListResourceDonationItemsDonor() {}

type ListResourceDonationItemsDonorPlatform string

const (
	ListResourceDonationItemsDonorPlatformGitHub ListResourceDonationItemsDonorPlatform = "github"
)

func (r ListResourceDonationItemsDonorPlatform) IsKnown() bool {
	switch r {
	case ListResourceDonationItemsDonorPlatformGitHub:
		return true
	}
	return false
}

type ListResourceDonationItemsIssue struct {
	ID             string                                `json:"id,required" format:"uuid"`
	Funding        ListResourceDonationItemsIssueFunding `json:"funding,required"`
	IssueCreatedAt time.Time                             `json:"issue_created_at,required" format:"date-time"`
	// If a maintainer needs to mark this issue as solved
	NeedsConfirmationSolved bool `json:"needs_confirmation_solved,required"`
	// GitHub #number
	Number int64 `json:"number,required"`
	// Issue platform (currently always GitHub)
	Platform ListResourceDonationItemsIssuePlatform `json:"platform,required"`
	// If this issue currently has the Polar badge SVG embedded
	PledgeBadgeCurrentlyEmbedded bool `json:"pledge_badge_currently_embedded,required"`
	// The repository that the issue is in
	Repository ListResourceDonationItemsIssueRepository `json:"repository,required"`
	State      ListResourceDonationItemsIssueState      `json:"state,required"`
	// GitHub issue title
	Title string `json:"title,required"`
	// GitHub assignees
	Assignees []ListResourceDonationItemsIssueAssignee `json:"assignees,nullable"`
	// GitHub author
	Author ListResourceDonationItemsIssueAuthor `json:"author,nullable"`
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
	Labels            []ListResourceDonationItemsIssueLabel `json:"labels"`
	// GitHub reactions
	Reactions ListResourceDonationItemsIssueReactions `json:"reactions,nullable"`
	// Share of rewrads that will be rewarded to contributors of this issue. A number
	// between 0 and 100 (inclusive).
	UpfrontSplitToContributors int64                              `json:"upfront_split_to_contributors,nullable"`
	JSON                       listResourceDonationItemsIssueJSON `json:"-"`
}

// listResourceDonationItemsIssueJSON contains the JSON metadata for the struct
// [ListResourceDonationItemsIssue]
type listResourceDonationItemsIssueJSON struct {
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

func (r *ListResourceDonationItemsIssue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResourceDonationItemsIssueJSON) RawJSON() string {
	return r.raw
}

type ListResourceDonationItemsIssueFunding struct {
	FundingGoal ListResourceDonationItemsIssueFundingFundingGoal `json:"funding_goal,nullable"`
	// Sum of pledges to this isuse (including currently open pledges and pledges that
	// have been paid out). Always in USD.
	PledgesSum ListResourceDonationItemsIssueFundingPledgesSum `json:"pledges_sum,nullable"`
	JSON       listResourceDonationItemsIssueFundingJSON       `json:"-"`
}

// listResourceDonationItemsIssueFundingJSON contains the JSON metadata for the
// struct [ListResourceDonationItemsIssueFunding]
type listResourceDonationItemsIssueFundingJSON struct {
	FundingGoal apijson.Field
	PledgesSum  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListResourceDonationItemsIssueFunding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResourceDonationItemsIssueFundingJSON) RawJSON() string {
	return r.raw
}

type ListResourceDonationItemsIssueFundingFundingGoal struct {
	// Amount in the currencies smallest unit (cents if currency is USD)
	Amount int64 `json:"amount,required"`
	// Three letter currency code (eg: USD)
	Currency string                                               `json:"currency,required"`
	JSON     listResourceDonationItemsIssueFundingFundingGoalJSON `json:"-"`
}

// listResourceDonationItemsIssueFundingFundingGoalJSON contains the JSON metadata
// for the struct [ListResourceDonationItemsIssueFundingFundingGoal]
type listResourceDonationItemsIssueFundingFundingGoalJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListResourceDonationItemsIssueFundingFundingGoal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResourceDonationItemsIssueFundingFundingGoalJSON) RawJSON() string {
	return r.raw
}

// Sum of pledges to this isuse (including currently open pledges and pledges that
// have been paid out). Always in USD.
type ListResourceDonationItemsIssueFundingPledgesSum struct {
	// Amount in the currencies smallest unit (cents if currency is USD)
	Amount int64 `json:"amount,required"`
	// Three letter currency code (eg: USD)
	Currency string                                              `json:"currency,required"`
	JSON     listResourceDonationItemsIssueFundingPledgesSumJSON `json:"-"`
}

// listResourceDonationItemsIssueFundingPledgesSumJSON contains the JSON metadata
// for the struct [ListResourceDonationItemsIssueFundingPledgesSum]
type listResourceDonationItemsIssueFundingPledgesSumJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListResourceDonationItemsIssueFundingPledgesSum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResourceDonationItemsIssueFundingPledgesSumJSON) RawJSON() string {
	return r.raw
}

// Issue platform (currently always GitHub)
type ListResourceDonationItemsIssuePlatform string

const (
	ListResourceDonationItemsIssuePlatformGitHub ListResourceDonationItemsIssuePlatform = "github"
)

func (r ListResourceDonationItemsIssuePlatform) IsKnown() bool {
	switch r {
	case ListResourceDonationItemsIssuePlatformGitHub:
		return true
	}
	return false
}

// The repository that the issue is in
type ListResourceDonationItemsIssueRepository struct {
	ID           string                                               `json:"id,required" format:"uuid"`
	IsPrivate    bool                                                 `json:"is_private,required"`
	Name         string                                               `json:"name,required"`
	Organization ListResourceDonationItemsIssueRepositoryOrganization `json:"organization,required"`
	Platform     ListResourceDonationItemsIssueRepositoryPlatform     `json:"platform,required"`
	// Settings for the repository profile
	ProfileSettings ListResourceDonationItemsIssueRepositoryProfileSettings `json:"profile_settings,required,nullable"`
	Description     string                                                  `json:"description,nullable"`
	Homepage        string                                                  `json:"homepage,nullable"`
	License         string                                                  `json:"license,nullable"`
	Stars           int64                                                   `json:"stars,nullable"`
	JSON            listResourceDonationItemsIssueRepositoryJSON            `json:"-"`
}

// listResourceDonationItemsIssueRepositoryJSON contains the JSON metadata for the
// struct [ListResourceDonationItemsIssueRepository]
type listResourceDonationItemsIssueRepositoryJSON struct {
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

func (r *ListResourceDonationItemsIssueRepository) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResourceDonationItemsIssueRepositoryJSON) RawJSON() string {
	return r.raw
}

type ListResourceDonationItemsIssueRepositoryOrganization struct {
	ID         string                                                       `json:"id,required" format:"uuid"`
	AvatarURL  string                                                       `json:"avatar_url,required"`
	IsPersonal bool                                                         `json:"is_personal,required"`
	Name       string                                                       `json:"name,required"`
	Platform   ListResourceDonationItemsIssueRepositoryOrganizationPlatform `json:"platform,required"`
	Bio        string                                                       `json:"bio,nullable"`
	Blog       string                                                       `json:"blog,nullable"`
	Company    string                                                       `json:"company,nullable"`
	Email      string                                                       `json:"email,nullable"`
	Location   string                                                       `json:"location,nullable"`
	// The organization ID.
	OrganizationID  string                                                   `json:"organization_id,nullable" format:"uuid4"`
	PrettyName      string                                                   `json:"pretty_name,nullable"`
	TwitterUsername string                                                   `json:"twitter_username,nullable"`
	JSON            listResourceDonationItemsIssueRepositoryOrganizationJSON `json:"-"`
}

// listResourceDonationItemsIssueRepositoryOrganizationJSON contains the JSON
// metadata for the struct [ListResourceDonationItemsIssueRepositoryOrganization]
type listResourceDonationItemsIssueRepositoryOrganizationJSON struct {
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

func (r *ListResourceDonationItemsIssueRepositoryOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResourceDonationItemsIssueRepositoryOrganizationJSON) RawJSON() string {
	return r.raw
}

type ListResourceDonationItemsIssueRepositoryOrganizationPlatform string

const (
	ListResourceDonationItemsIssueRepositoryOrganizationPlatformGitHub ListResourceDonationItemsIssueRepositoryOrganizationPlatform = "github"
)

func (r ListResourceDonationItemsIssueRepositoryOrganizationPlatform) IsKnown() bool {
	switch r {
	case ListResourceDonationItemsIssueRepositoryOrganizationPlatformGitHub:
		return true
	}
	return false
}

type ListResourceDonationItemsIssueRepositoryPlatform string

const (
	ListResourceDonationItemsIssueRepositoryPlatformGitHub ListResourceDonationItemsIssueRepositoryPlatform = "github"
)

func (r ListResourceDonationItemsIssueRepositoryPlatform) IsKnown() bool {
	switch r {
	case ListResourceDonationItemsIssueRepositoryPlatformGitHub:
		return true
	}
	return false
}

// Settings for the repository profile
type ListResourceDonationItemsIssueRepositoryProfileSettings struct {
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
	JSON  listResourceDonationItemsIssueRepositoryProfileSettingsJSON `json:"-"`
}

// listResourceDonationItemsIssueRepositoryProfileSettingsJSON contains the JSON
// metadata for the struct
// [ListResourceDonationItemsIssueRepositoryProfileSettings]
type listResourceDonationItemsIssueRepositoryProfileSettingsJSON struct {
	CoverImageURL                apijson.Field
	Description                  apijson.Field
	FeaturedOrganizations        apijson.Field
	HighlightedSubscriptionTiers apijson.Field
	Links                        apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *ListResourceDonationItemsIssueRepositoryProfileSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResourceDonationItemsIssueRepositoryProfileSettingsJSON) RawJSON() string {
	return r.raw
}

type ListResourceDonationItemsIssueState string

const (
	ListResourceDonationItemsIssueStateOpen   ListResourceDonationItemsIssueState = "open"
	ListResourceDonationItemsIssueStateClosed ListResourceDonationItemsIssueState = "closed"
)

func (r ListResourceDonationItemsIssueState) IsKnown() bool {
	switch r {
	case ListResourceDonationItemsIssueStateOpen, ListResourceDonationItemsIssueStateClosed:
		return true
	}
	return false
}

type ListResourceDonationItemsIssueAssignee struct {
	ID        int64                                      `json:"id,required"`
	AvatarURL string                                     `json:"avatar_url,required" format:"uri"`
	HTMLURL   string                                     `json:"html_url,required" format:"uri"`
	Login     string                                     `json:"login,required"`
	JSON      listResourceDonationItemsIssueAssigneeJSON `json:"-"`
}

// listResourceDonationItemsIssueAssigneeJSON contains the JSON metadata for the
// struct [ListResourceDonationItemsIssueAssignee]
type listResourceDonationItemsIssueAssigneeJSON struct {
	ID          apijson.Field
	AvatarURL   apijson.Field
	HTMLURL     apijson.Field
	Login       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListResourceDonationItemsIssueAssignee) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResourceDonationItemsIssueAssigneeJSON) RawJSON() string {
	return r.raw
}

// GitHub author
type ListResourceDonationItemsIssueAuthor struct {
	ID        int64                                    `json:"id,required"`
	AvatarURL string                                   `json:"avatar_url,required" format:"uri"`
	HTMLURL   string                                   `json:"html_url,required" format:"uri"`
	Login     string                                   `json:"login,required"`
	JSON      listResourceDonationItemsIssueAuthorJSON `json:"-"`
}

// listResourceDonationItemsIssueAuthorJSON contains the JSON metadata for the
// struct [ListResourceDonationItemsIssueAuthor]
type listResourceDonationItemsIssueAuthorJSON struct {
	ID          apijson.Field
	AvatarURL   apijson.Field
	HTMLURL     apijson.Field
	Login       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListResourceDonationItemsIssueAuthor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResourceDonationItemsIssueAuthorJSON) RawJSON() string {
	return r.raw
}

type ListResourceDonationItemsIssueLabel struct {
	Color string                                  `json:"color,required"`
	Name  string                                  `json:"name,required"`
	JSON  listResourceDonationItemsIssueLabelJSON `json:"-"`
}

// listResourceDonationItemsIssueLabelJSON contains the JSON metadata for the
// struct [ListResourceDonationItemsIssueLabel]
type listResourceDonationItemsIssueLabelJSON struct {
	Color       apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListResourceDonationItemsIssueLabel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResourceDonationItemsIssueLabelJSON) RawJSON() string {
	return r.raw
}

// GitHub reactions
type ListResourceDonationItemsIssueReactions struct {
	Confused   int64                                       `json:"confused,required"`
	Eyes       int64                                       `json:"eyes,required"`
	Heart      int64                                       `json:"heart,required"`
	Hooray     int64                                       `json:"hooray,required"`
	Laugh      int64                                       `json:"laugh,required"`
	MinusOne   int64                                       `json:"minus_one,required"`
	PlusOne    int64                                       `json:"plus_one,required"`
	Rocket     int64                                       `json:"rocket,required"`
	TotalCount int64                                       `json:"total_count,required"`
	JSON       listResourceDonationItemsIssueReactionsJSON `json:"-"`
}

// listResourceDonationItemsIssueReactionsJSON contains the JSON metadata for the
// struct [ListResourceDonationItemsIssueReactions]
type listResourceDonationItemsIssueReactionsJSON struct {
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

func (r *ListResourceDonationItemsIssueReactions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResourceDonationItemsIssueReactionsJSON) RawJSON() string {
	return r.raw
}

type DonationSearchParams struct {
	ToOrganizationID param.Field[string] `query:"to_organization_id,required" format:"uuid4"`
	// Size of a page, defaults to 10. Maximum is 100.
	Limit param.Field[int64] `query:"limit"`
	// Page number, defaults to 1.
	Page param.Field[int64] `query:"page"`
	// Sorting criterion. Several criteria can be used simultaneously and will be
	// applied in order. Add a minus sign `-` before the criteria name to sort by
	// descending order.
	Sorting param.Field[[]string] `query:"sorting"`
}

// URLQuery serializes [DonationSearchParams]'s query parameters as `url.Values`.
func (r DonationSearchParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DonationStatisticsParams struct {
	DonationsInterval param.Field[DonationStatisticsParamsDonationsInterval] `query:"donationsInterval,required"`
	EndDate           param.Field[time.Time]                                 `query:"end_date,required" format:"date"`
	StartDate         param.Field[time.Time]                                 `query:"start_date,required" format:"date"`
	ToOrganizationID  param.Field[string]                                    `query:"to_organization_id,required" format:"uuid4"`
}

// URLQuery serializes [DonationStatisticsParams]'s query parameters as
// `url.Values`.
func (r DonationStatisticsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DonationStatisticsParamsDonationsInterval string

const (
	DonationStatisticsParamsDonationsIntervalMonth DonationStatisticsParamsDonationsInterval = "month"
	DonationStatisticsParamsDonationsIntervalWeek  DonationStatisticsParamsDonationsInterval = "week"
	DonationStatisticsParamsDonationsIntervalDay   DonationStatisticsParamsDonationsInterval = "day"
)

func (r DonationStatisticsParamsDonationsInterval) IsKnown() bool {
	switch r {
	case DonationStatisticsParamsDonationsIntervalMonth, DonationStatisticsParamsDonationsIntervalWeek, DonationStatisticsParamsDonationsIntervalDay:
		return true
	}
	return false
}
