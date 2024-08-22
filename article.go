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

// ArticleService contains methods and other services that help with interacting
// with the polar API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewArticleService] method instead.
type ArticleService struct {
	Options   []option.RequestOption
	Receivers *ArticleReceiverService
}

// NewArticleService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewArticleService(opts ...option.RequestOption) (r *ArticleService) {
	r = &ArticleService{}
	r.Options = opts
	r.Receivers = NewArticleReceiverService(opts...)
	return
}

// Create an article.
func (r *ArticleService) New(ctx context.Context, body ArticleNewParams, opts ...option.RequestOption) (res *Article, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/articles/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get an article by ID.
func (r *ArticleService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Article, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/articles/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an article.
func (r *ArticleService) Update(ctx context.Context, id string, body ArticleUpdateParams, opts ...option.RequestOption) (res *Article, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/articles/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List articles.
func (r *ArticleService) List(ctx context.Context, query ArticleListParams, opts ...option.RequestOption) (res *ListResourceArticle, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/articles/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete an article.
func (r *ArticleService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/articles/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Send an article preview by email.
func (r *ArticleService) Preview(ctx context.Context, id string, body ArticlePreviewParams, opts ...option.RequestOption) (res *ArticlePreviewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/articles/%s/preview", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Send an article by email to all subscribers.
func (r *ArticleService) Send(ctx context.Context, id string, opts ...option.RequestOption) (res *ArticleSendResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/articles/%s/send", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type Article struct {
	ID                        string              `json:"id,required" format:"uuid4"`
	Body                      string              `json:"body,required"`
	Byline                    ArticleByline       `json:"byline,required"`
	IsPinned                  bool                `json:"is_pinned,required"`
	IsPreview                 bool                `json:"is_preview,required"`
	Organization              ArticleOrganization `json:"organization,required"`
	OrganizationID            string              `json:"organization_id,required" format:"uuid4"`
	Slug                      string              `json:"slug,required"`
	Title                     string              `json:"title,required"`
	Visibility                ArticleVisibility   `json:"visibility,required"`
	EmailSentToCount          int64               `json:"email_sent_to_count,nullable"`
	NotificationsSentAt       time.Time           `json:"notifications_sent_at,nullable" format:"date-time"`
	NotifySubscribers         bool                `json:"notify_subscribers,nullable"`
	OgDescription             string              `json:"og_description,nullable"`
	OgImageURL                string              `json:"og_image_url,nullable"`
	PaidSubscribersOnly       bool                `json:"paid_subscribers_only,nullable"`
	PaidSubscribersOnlyEndsAt time.Time           `json:"paid_subscribers_only_ends_at,nullable" format:"date-time"`
	PublishedAt               time.Time           `json:"published_at,nullable" format:"date-time"`
	UserID                    string              `json:"user_id,nullable" format:"uuid4"`
	JSON                      articleJSON         `json:"-"`
}

// articleJSON contains the JSON metadata for the struct [Article]
type articleJSON struct {
	ID                        apijson.Field
	Body                      apijson.Field
	Byline                    apijson.Field
	IsPinned                  apijson.Field
	IsPreview                 apijson.Field
	Organization              apijson.Field
	OrganizationID            apijson.Field
	Slug                      apijson.Field
	Title                     apijson.Field
	Visibility                apijson.Field
	EmailSentToCount          apijson.Field
	NotificationsSentAt       apijson.Field
	NotifySubscribers         apijson.Field
	OgDescription             apijson.Field
	OgImageURL                apijson.Field
	PaidSubscribersOnly       apijson.Field
	PaidSubscribersOnlyEndsAt apijson.Field
	PublishedAt               apijson.Field
	UserID                    apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *Article) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r articleJSON) RawJSON() string {
	return r.raw
}

type ArticleByline struct {
	Name      string            `json:"name,required"`
	AvatarURL string            `json:"avatar_url,nullable"`
	JSON      articleBylineJSON `json:"-"`
}

// articleBylineJSON contains the JSON metadata for the struct [ArticleByline]
type articleBylineJSON struct {
	Name        apijson.Field
	AvatarURL   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArticleByline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r articleBylineJSON) RawJSON() string {
	return r.raw
}

type ArticleOrganization struct {
	// The organization ID.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// If this organizations accepts donations
	DonationsEnabled bool `json:"donations_enabled,required"`
	// Settings for the organization features
	FeatureSettings       ArticleOrganizationFeatureSettings `json:"feature_settings,required,nullable"`
	Name                  string                             `json:"name,required"`
	PledgeBadgeShowAmount bool                               `json:"pledge_badge_show_amount,required"`
	PledgeMinimumAmount   int64                              `json:"pledge_minimum_amount,required"`
	// Settings for the organization profile
	ProfileSettings                   ArticleOrganizationProfileSettings `json:"profile_settings,required,nullable"`
	Slug                              string                             `json:"slug,required"`
	AvatarURL                         string                             `json:"avatar_url,nullable"`
	Bio                               string                             `json:"bio,nullable"`
	Blog                              string                             `json:"blog,nullable"`
	Company                           string                             `json:"company,nullable"`
	DefaultUpfrontSplitToContributors int64                              `json:"default_upfront_split_to_contributors,nullable"`
	Email                             string                             `json:"email,nullable"`
	Location                          string                             `json:"location,nullable"`
	// Last modification timestamp of the object.
	ModifiedAt      time.Time               `json:"modified_at,nullable" format:"date-time"`
	TwitterUsername string                  `json:"twitter_username,nullable"`
	JSON            articleOrganizationJSON `json:"-"`
}

// articleOrganizationJSON contains the JSON metadata for the struct
// [ArticleOrganization]
type articleOrganizationJSON struct {
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

func (r *ArticleOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r articleOrganizationJSON) RawJSON() string {
	return r.raw
}

// Settings for the organization features
type ArticleOrganizationFeatureSettings struct {
	// If this organization has articles enabled
	ArticlesEnabled bool `json:"articles_enabled"`
	// If this organization has issue funding enabled
	IssueFundingEnabled bool `json:"issue_funding_enabled"`
	// If this organization has subscriptions enabled
	SubscriptionsEnabled bool                                   `json:"subscriptions_enabled"`
	JSON                 articleOrganizationFeatureSettingsJSON `json:"-"`
}

// articleOrganizationFeatureSettingsJSON contains the JSON metadata for the struct
// [ArticleOrganizationFeatureSettings]
type articleOrganizationFeatureSettingsJSON struct {
	ArticlesEnabled      apijson.Field
	IssueFundingEnabled  apijson.Field
	SubscriptionsEnabled apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ArticleOrganizationFeatureSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r articleOrganizationFeatureSettingsJSON) RawJSON() string {
	return r.raw
}

// Settings for the organization profile
type ArticleOrganizationProfileSettings struct {
	// A description of the organization
	Description string `json:"description,nullable"`
	// A list of featured organizations
	FeaturedOrganizations []string `json:"featured_organizations,nullable" format:"uuid4"`
	// A list of featured projects
	FeaturedProjects []string `json:"featured_projects,nullable" format:"uuid4"`
	// A list of links associated with the organization
	Links []string `json:"links,nullable" format:"uri"`
	// Subscription promotion settings
	Subscribe ArticleOrganizationProfileSettingsSubscribe `json:"subscribe,nullable"`
	JSON      articleOrganizationProfileSettingsJSON      `json:"-"`
}

// articleOrganizationProfileSettingsJSON contains the JSON metadata for the struct
// [ArticleOrganizationProfileSettings]
type articleOrganizationProfileSettingsJSON struct {
	Description           apijson.Field
	FeaturedOrganizations apijson.Field
	FeaturedProjects      apijson.Field
	Links                 apijson.Field
	Subscribe             apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ArticleOrganizationProfileSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r articleOrganizationProfileSettingsJSON) RawJSON() string {
	return r.raw
}

// Subscription promotion settings
type ArticleOrganizationProfileSettingsSubscribe struct {
	// Include free subscribers in total count
	CountFree bool `json:"count_free"`
	// Promote email subscription (free)
	Promote bool `json:"promote"`
	// Show subscription count publicly
	ShowCount bool                                            `json:"show_count"`
	JSON      articleOrganizationProfileSettingsSubscribeJSON `json:"-"`
}

// articleOrganizationProfileSettingsSubscribeJSON contains the JSON metadata for
// the struct [ArticleOrganizationProfileSettingsSubscribe]
type articleOrganizationProfileSettingsSubscribeJSON struct {
	CountFree   apijson.Field
	Promote     apijson.Field
	ShowCount   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArticleOrganizationProfileSettingsSubscribe) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r articleOrganizationProfileSettingsSubscribeJSON) RawJSON() string {
	return r.raw
}

type ArticleVisibility string

const (
	ArticleVisibilityPublic  ArticleVisibility = "public"
	ArticleVisibilityHidden  ArticleVisibility = "hidden"
	ArticleVisibilityPrivate ArticleVisibility = "private"
)

func (r ArticleVisibility) IsKnown() bool {
	switch r {
	case ArticleVisibilityPublic, ArticleVisibilityHidden, ArticleVisibilityPrivate:
		return true
	}
	return false
}

type ArticleReceivers struct {
	FreeSubscribers     int64                `json:"free_subscribers,required"`
	OrganizationMembers int64                `json:"organization_members,required"`
	PremiumSubscribers  int64                `json:"premium_subscribers,required"`
	JSON                articleReceiversJSON `json:"-"`
}

// articleReceiversJSON contains the JSON metadata for the struct
// [ArticleReceivers]
type articleReceiversJSON struct {
	FreeSubscribers     apijson.Field
	OrganizationMembers apijson.Field
	PremiumSubscribers  apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ArticleReceivers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r articleReceiversJSON) RawJSON() string {
	return r.raw
}

type ListResourceArticle struct {
	Pagination ListResourceArticlePagination `json:"pagination,required"`
	Items      []Article                     `json:"items"`
	JSON       listResourceArticleJSON       `json:"-"`
}

// listResourceArticleJSON contains the JSON metadata for the struct
// [ListResourceArticle]
type listResourceArticleJSON struct {
	Pagination  apijson.Field
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListResourceArticle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResourceArticleJSON) RawJSON() string {
	return r.raw
}

type ListResourceArticlePagination struct {
	MaxPage    int64                             `json:"max_page,required"`
	TotalCount int64                             `json:"total_count,required"`
	JSON       listResourceArticlePaginationJSON `json:"-"`
}

// listResourceArticlePaginationJSON contains the JSON metadata for the struct
// [ListResourceArticlePagination]
type listResourceArticlePaginationJSON struct {
	MaxPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListResourceArticlePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResourceArticlePaginationJSON) RawJSON() string {
	return r.raw
}

type ArticlePreviewResponse = interface{}

type ArticleSendResponse = interface{}

type ArticleNewParams struct {
	// Title of the article.
	Title param.Field[string] `json:"title,required"`
	// Body in string format. Either one of body or body_base64 is required.
	Body param.Field[string] `json:"body"`
	// Body in base64-encoded format. Can be helpful to bypass Web Application
	// Firewalls (WAF). Either one of body or body_base64 is required.
	BodyBase64 param.Field[string] `json:"body_base64"`
	// If the user or organization should be credited in the byline.
	Byline param.Field[ArticleNewParamsByline] `json:"byline"`
	// If the article should be pinned
	IsPinned param.Field[bool] `json:"is_pinned"`
	// Set to true to deliver this article via email and/or notifications to
	// subscribers.
	NotifySubscribers param.Field[bool] `json:"notify_subscribers"`
	// Custom og:description value
	OgDescription param.Field[string] `json:"og_description"`
	// Custom og:image URL value
	OgImageURL param.Field[string] `json:"og_image_url" format:"uri"`
	// The organization ID.
	OrganizationID param.Field[string] `json:"organization_id" format:"uuid4"`
	// Set to true to only make this article available for subscribers to a paid
	// subscription tier in the organization.
	PaidSubscribersOnly param.Field[bool] `json:"paid_subscribers_only"`
	// If specified, time at which the article should no longer be restricted to paid
	// subscribers. Only relevant if `paid_subscribers_only` is true.
	PaidSubscribersOnlyEndsAt param.Field[time.Time] `json:"paid_subscribers_only_ends_at" format:"date-time"`
	// Time of publishing. If this date is in the future, the post will be scheduled to
	// publish at this time. If visibility is 'public', published_at will default to
	// the current time.
	PublishedAt param.Field[time.Time] `json:"published_at" format:"date-time"`
	// Slug of the article to be used in URLs. If no slug is provided one will be
	// generated from the title.
	Slug       param.Field[string]                     `json:"slug"`
	Visibility param.Field[ArticleNewParamsVisibility] `json:"visibility"`
}

func (r ArticleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If the user or organization should be credited in the byline.
type ArticleNewParamsByline string

const (
	ArticleNewParamsBylineUser         ArticleNewParamsByline = "user"
	ArticleNewParamsBylineOrganization ArticleNewParamsByline = "organization"
)

func (r ArticleNewParamsByline) IsKnown() bool {
	switch r {
	case ArticleNewParamsBylineUser, ArticleNewParamsBylineOrganization:
		return true
	}
	return false
}

type ArticleNewParamsVisibility string

const (
	ArticleNewParamsVisibilityPublic  ArticleNewParamsVisibility = "public"
	ArticleNewParamsVisibilityHidden  ArticleNewParamsVisibility = "hidden"
	ArticleNewParamsVisibilityPrivate ArticleNewParamsVisibility = "private"
)

func (r ArticleNewParamsVisibility) IsKnown() bool {
	switch r {
	case ArticleNewParamsVisibilityPublic, ArticleNewParamsVisibilityHidden, ArticleNewParamsVisibilityPrivate:
		return true
	}
	return false
}

type ArticleUpdateParams struct {
	// Body in string format. body and body_base64 are mutually exclusive.
	Body param.Field[string] `json:"body"`
	// Body in base64-encoded format. Can be helpful to bypass Web Application
	// Firewalls (WAF). body and body_base64 are mutually exclusive.
	BodyBase64 param.Field[string] `json:"body_base64"`
	// If the user or organization should be credited in the byline.
	Byline param.Field[ArticleUpdateParamsByline] `json:"byline"`
	// If the article should be pinned
	IsPinned param.Field[bool] `json:"is_pinned"`
	// Set to true to deliver this article via email and/or notifications to
	// subscribers.
	NotifySubscribers param.Field[bool] `json:"notify_subscribers"`
	// Custom og:description value
	OgDescription param.Field[string] `json:"og_description"`
	// Custom og:image URL value
	OgImageURL param.Field[string] `json:"og_image_url" format:"uri"`
	// Set to true to only make this article available for subscribers to a paid
	// subscription tier in the organization.
	PaidSubscribersOnly param.Field[bool] `json:"paid_subscribers_only"`
	// If specified, time at which the article should no longer be restricted to paid
	// subscribers. Only relevant if `paid_subscribers_only` is true.
	PaidSubscribersOnlyEndsAt param.Field[time.Time] `json:"paid_subscribers_only_ends_at" format:"date-time"`
	// Time of publishing. If this date is in the future, the post will be scheduled to
	// publish at this time.
	PublishedAt param.Field[time.Time]                     `json:"published_at" format:"date-time"`
	Slug        param.Field[string]                        `json:"slug"`
	Title       param.Field[string]                        `json:"title"`
	Visibility  param.Field[ArticleUpdateParamsVisibility] `json:"visibility"`
}

func (r ArticleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If the user or organization should be credited in the byline.
type ArticleUpdateParamsByline string

const (
	ArticleUpdateParamsBylineUser         ArticleUpdateParamsByline = "user"
	ArticleUpdateParamsBylineOrganization ArticleUpdateParamsByline = "organization"
)

func (r ArticleUpdateParamsByline) IsKnown() bool {
	switch r {
	case ArticleUpdateParamsBylineUser, ArticleUpdateParamsBylineOrganization:
		return true
	}
	return false
}

type ArticleUpdateParamsVisibility string

const (
	ArticleUpdateParamsVisibilityPublic  ArticleUpdateParamsVisibility = "public"
	ArticleUpdateParamsVisibilityHidden  ArticleUpdateParamsVisibility = "hidden"
	ArticleUpdateParamsVisibilityPrivate ArticleUpdateParamsVisibility = "private"
)

func (r ArticleUpdateParamsVisibility) IsKnown() bool {
	switch r {
	case ArticleUpdateParamsVisibilityPublic, ArticleUpdateParamsVisibilityHidden, ArticleUpdateParamsVisibilityPrivate:
		return true
	}
	return false
}

type ArticleListParams struct {
	// Filter by pinned status.
	IsPinned param.Field[bool] `query:"is_pinned"`
	// Filter by published status.
	IsPublished param.Field[bool] `query:"is_published"`
	// Filter by subscription status.
	IsSubscribed param.Field[bool] `query:"is_subscribed"`
	// Size of a page, defaults to 10. Maximum is 100.
	Limit param.Field[int64] `query:"limit"`
	// Filter by organization ID.
	OrganizationID param.Field[ArticleListParamsOrganizationIDUnion] `query:"organization_id" format:"uuid4"`
	// Page number, defaults to 1.
	Page param.Field[int64] `query:"page"`
	// Filter by slug.
	Slug param.Field[string] `query:"slug"`
	// Filter by visibility.
	Visibility param.Field[ArticleListParamsVisibilityUnion] `query:"visibility"`
}

// URLQuery serializes [ArticleListParams]'s query parameters as `url.Values`.
func (r ArticleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by organization ID.
//
// Satisfied by [shared.UnionString], [ArticleListParamsOrganizationIDArray].
type ArticleListParamsOrganizationIDUnion interface {
	ImplementsArticleListParamsOrganizationIDUnion()
}

type ArticleListParamsOrganizationIDArray []string

func (r ArticleListParamsOrganizationIDArray) ImplementsArticleListParamsOrganizationIDUnion() {}

// Filter by visibility.
//
// Satisfied by [ArticleListParamsVisibilityArticleVisibility],
// [ArticleListParamsVisibilityArray].
type ArticleListParamsVisibilityUnion interface {
	implementsArticleListParamsVisibilityUnion()
}

type ArticleListParamsVisibilityArticleVisibility string

const (
	ArticleListParamsVisibilityArticleVisibilityPublic  ArticleListParamsVisibilityArticleVisibility = "public"
	ArticleListParamsVisibilityArticleVisibilityHidden  ArticleListParamsVisibilityArticleVisibility = "hidden"
	ArticleListParamsVisibilityArticleVisibilityPrivate ArticleListParamsVisibilityArticleVisibility = "private"
)

func (r ArticleListParamsVisibilityArticleVisibility) IsKnown() bool {
	switch r {
	case ArticleListParamsVisibilityArticleVisibilityPublic, ArticleListParamsVisibilityArticleVisibilityHidden, ArticleListParamsVisibilityArticleVisibilityPrivate:
		return true
	}
	return false
}

func (r ArticleListParamsVisibilityArticleVisibility) implementsArticleListParamsVisibilityUnion() {}

type ArticleListParamsVisibilityArray []ArticleListParamsVisibilityArray

func (r ArticleListParamsVisibilityArray) implementsArticleListParamsVisibilityUnion() {}

type ArticlePreviewParams struct {
	// Email address to send the preview to. The user must be registered on Polar.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r ArticlePreviewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
