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
	"github.com/polarsource/polar-go/internal/param"
	"github.com/polarsource/polar-go/internal/requestconfig"
	"github.com/polarsource/polar-go/option"
	"github.com/tidwall/gjson"
)

// UserBenefitService contains methods and other services that help with
// interacting with the polar API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserBenefitService] method instead.
type UserBenefitService struct {
	Options []option.RequestOption
}

// NewUserBenefitService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserBenefitService(opts ...option.RequestOption) (r *UserBenefitService) {
	r = &UserBenefitService{}
	r.Options = opts
	return
}

// Get a granted benefit by ID.
func (r *UserBenefitService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *UserBenefitGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/users/benefits/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List my granted benefits.
func (r *UserBenefitService) List(ctx context.Context, query UserBenefitListParams, opts ...option.RequestOption) (res *UserBenefitListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/users/benefits/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type UserBenefitGetResponse struct {
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The ID of the benefit.
	ID   string                     `json:"id,required" format:"uuid4"`
	Type UserBenefitGetResponseType `json:"type,required"`
	// The description of the benefit.
	Description string `json:"description,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool `json:"selectable,required"`
	// Whether the benefit is deletable.
	Deletable bool `json:"deletable,required"`
	// The ID of the organization owning the benefit.
	OrganizationID string `json:"organization_id,required" format:"uuid4"`
	// This field can have the runtime type of
	// [UserBenefitGetResponseBenefitArticlesSubscriberProperties],
	// [UserBenefitGetResponseBenefitAdsSubscriberProperties],
	// [UserBenefitGetResponseBenefitDiscordSubscriberProperties],
	// [UserBenefitGetResponseBenefitCustomSubscriberProperties],
	// [UserBenefitGetResponseBenefitGitHubRepositorySubscriberProperties],
	// [UserBenefitGetResponseBenefitDownloadablesSubscriberProperties].
	Properties interface{} `json:"properties"`
	// This field can have the runtime type of
	// [[]UserBenefitGetResponseBenefitAdsSubscriberGrant],
	// [[]UserBenefitGetResponseBenefitCustomSubscriberGrant].
	Grants interface{}                `json:"grants,required"`
	JSON   userBenefitGetResponseJSON `json:"-"`
	union  UserBenefitGetResponseUnion
}

// userBenefitGetResponseJSON contains the JSON metadata for the struct
// [UserBenefitGetResponse]
type userBenefitGetResponseJSON struct {
	CreatedAt      apijson.Field
	ModifiedAt     apijson.Field
	ID             apijson.Field
	Type           apijson.Field
	Description    apijson.Field
	Selectable     apijson.Field
	Deletable      apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	Grants         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r userBenefitGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *UserBenefitGetResponse) UnmarshalJSON(data []byte) (err error) {
	*r = UserBenefitGetResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [UserBenefitGetResponseUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are
// [UserBenefitGetResponseBenefitArticlesSubscriber],
// [UserBenefitGetResponseBenefitAdsSubscriber],
// [UserBenefitGetResponseBenefitDiscordSubscriber],
// [UserBenefitGetResponseBenefitCustomSubscriber],
// [UserBenefitGetResponseBenefitGitHubRepositorySubscriber],
// [UserBenefitGetResponseBenefitDownloadablesSubscriber].
func (r UserBenefitGetResponse) AsUnion() UserBenefitGetResponseUnion {
	return r.union
}

// Union satisfied by [UserBenefitGetResponseBenefitArticlesSubscriber],
// [UserBenefitGetResponseBenefitAdsSubscriber],
// [UserBenefitGetResponseBenefitDiscordSubscriber],
// [UserBenefitGetResponseBenefitCustomSubscriber],
// [UserBenefitGetResponseBenefitGitHubRepositorySubscriber] or
// [UserBenefitGetResponseBenefitDownloadablesSubscriber].
type UserBenefitGetResponseUnion interface {
	implementsUserBenefitGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserBenefitGetResponseUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(UserBenefitGetResponseBenefitArticlesSubscriber{}),
			DiscriminatorValue: "articles",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(UserBenefitGetResponseBenefitAdsSubscriber{}),
			DiscriminatorValue: "ads",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(UserBenefitGetResponseBenefitDiscordSubscriber{}),
			DiscriminatorValue: "discord",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(UserBenefitGetResponseBenefitCustomSubscriber{}),
			DiscriminatorValue: "custom",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(UserBenefitGetResponseBenefitGitHubRepositorySubscriber{}),
			DiscriminatorValue: "github_repository",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(UserBenefitGetResponseBenefitDownloadablesSubscriber{}),
			DiscriminatorValue: "downloadables",
		},
	)
}

type UserBenefitGetResponseBenefitArticlesSubscriber struct {
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
	// Properties available to subscribers for a benefit of type `articles`.
	Properties UserBenefitGetResponseBenefitArticlesSubscriberProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                                `json:"selectable,required"`
	Type       UserBenefitGetResponseBenefitArticlesSubscriberType `json:"type,required"`
	JSON       userBenefitGetResponseBenefitArticlesSubscriberJSON `json:"-"`
}

// userBenefitGetResponseBenefitArticlesSubscriberJSON contains the JSON metadata
// for the struct [UserBenefitGetResponseBenefitArticlesSubscriber]
type userBenefitGetResponseBenefitArticlesSubscriberJSON struct {
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

func (r *UserBenefitGetResponseBenefitArticlesSubscriber) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userBenefitGetResponseBenefitArticlesSubscriberJSON) RawJSON() string {
	return r.raw
}

func (r UserBenefitGetResponseBenefitArticlesSubscriber) implementsUserBenefitGetResponse() {}

// Properties available to subscribers for a benefit of type `articles`.
type UserBenefitGetResponseBenefitArticlesSubscriberProperties struct {
	// Whether the user can access paid articles.
	PaidArticles bool                                                          `json:"paid_articles,required"`
	JSON         userBenefitGetResponseBenefitArticlesSubscriberPropertiesJSON `json:"-"`
}

// userBenefitGetResponseBenefitArticlesSubscriberPropertiesJSON contains the JSON
// metadata for the struct
// [UserBenefitGetResponseBenefitArticlesSubscriberProperties]
type userBenefitGetResponseBenefitArticlesSubscriberPropertiesJSON struct {
	PaidArticles apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *UserBenefitGetResponseBenefitArticlesSubscriberProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userBenefitGetResponseBenefitArticlesSubscriberPropertiesJSON) RawJSON() string {
	return r.raw
}

type UserBenefitGetResponseBenefitArticlesSubscriberType string

const (
	UserBenefitGetResponseBenefitArticlesSubscriberTypeArticles UserBenefitGetResponseBenefitArticlesSubscriberType = "articles"
)

func (r UserBenefitGetResponseBenefitArticlesSubscriberType) IsKnown() bool {
	switch r {
	case UserBenefitGetResponseBenefitArticlesSubscriberTypeArticles:
		return true
	}
	return false
}

type UserBenefitGetResponseBenefitAdsSubscriber struct {
	// The ID of the benefit.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the benefit is deletable.
	Deletable bool `json:"deletable,required"`
	// The description of the benefit.
	Description string                                            `json:"description,required"`
	Grants      []UserBenefitGetResponseBenefitAdsSubscriberGrant `json:"grants,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The ID of the organization owning the benefit.
	OrganizationID string `json:"organization_id,required" format:"uuid4"`
	// Properties for a benefit of type `ads`.
	Properties UserBenefitGetResponseBenefitAdsSubscriberProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                           `json:"selectable,required"`
	Type       UserBenefitGetResponseBenefitAdsSubscriberType `json:"type,required"`
	JSON       userBenefitGetResponseBenefitAdsSubscriberJSON `json:"-"`
}

// userBenefitGetResponseBenefitAdsSubscriberJSON contains the JSON metadata for
// the struct [UserBenefitGetResponseBenefitAdsSubscriber]
type userBenefitGetResponseBenefitAdsSubscriberJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Deletable      apijson.Field
	Description    apijson.Field
	Grants         apijson.Field
	ModifiedAt     apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	Selectable     apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *UserBenefitGetResponseBenefitAdsSubscriber) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userBenefitGetResponseBenefitAdsSubscriberJSON) RawJSON() string {
	return r.raw
}

func (r UserBenefitGetResponseBenefitAdsSubscriber) implementsUserBenefitGetResponse() {}

type UserBenefitGetResponseBenefitAdsSubscriberGrant struct {
	// The ID of the grant.
	ID string `json:"id,required" format:"uuid4"`
	// The ID of the benefit concerned by this grant.
	BenefitID string `json:"benefit_id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the benefit is granted.
	IsGranted bool `json:"is_granted,required"`
	// Whether the benefit is revoked.
	IsRevoked bool `json:"is_revoked,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The ID of the order that granted this benefit.
	OrderID    string                                                     `json:"order_id,required,nullable" format:"uuid4"`
	Properties UserBenefitGetResponseBenefitAdsSubscriberGrantsProperties `json:"properties,required"`
	// The ID of the subscription that granted this benefit.
	SubscriptionID string `json:"subscription_id,required,nullable" format:"uuid4"`
	// The ID of the user concerned by this grant.
	UserID string `json:"user_id,required" format:"uuid4"`
	// The timestamp when the benefit was granted. If `None`, the benefit is not
	// granted.
	GrantedAt time.Time `json:"granted_at,nullable" format:"date-time"`
	// The timestamp when the benefit was revoked. If `None`, the benefit is not
	// revoked.
	RevokedAt time.Time                                           `json:"revoked_at,nullable" format:"date-time"`
	JSON      userBenefitGetResponseBenefitAdsSubscriberGrantJSON `json:"-"`
}

// userBenefitGetResponseBenefitAdsSubscriberGrantJSON contains the JSON metadata
// for the struct [UserBenefitGetResponseBenefitAdsSubscriberGrant]
type userBenefitGetResponseBenefitAdsSubscriberGrantJSON struct {
	ID             apijson.Field
	BenefitID      apijson.Field
	CreatedAt      apijson.Field
	IsGranted      apijson.Field
	IsRevoked      apijson.Field
	ModifiedAt     apijson.Field
	OrderID        apijson.Field
	Properties     apijson.Field
	SubscriptionID apijson.Field
	UserID         apijson.Field
	GrantedAt      apijson.Field
	RevokedAt      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *UserBenefitGetResponseBenefitAdsSubscriberGrant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userBenefitGetResponseBenefitAdsSubscriberGrantJSON) RawJSON() string {
	return r.raw
}

type UserBenefitGetResponseBenefitAdsSubscriberGrantsProperties struct {
	// The ID of the enabled advertisement campaign for this benefit grant.
	AdvertisementCampaignID string                                                         `json:"advertisement_campaign_id,nullable" format:"uuid4"`
	JSON                    userBenefitGetResponseBenefitAdsSubscriberGrantsPropertiesJSON `json:"-"`
}

// userBenefitGetResponseBenefitAdsSubscriberGrantsPropertiesJSON contains the JSON
// metadata for the struct
// [UserBenefitGetResponseBenefitAdsSubscriberGrantsProperties]
type userBenefitGetResponseBenefitAdsSubscriberGrantsPropertiesJSON struct {
	AdvertisementCampaignID apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *UserBenefitGetResponseBenefitAdsSubscriberGrantsProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userBenefitGetResponseBenefitAdsSubscriberGrantsPropertiesJSON) RawJSON() string {
	return r.raw
}

// Properties for a benefit of type `ads`.
type UserBenefitGetResponseBenefitAdsSubscriberProperties struct {
	// The height of the displayed ad.
	ImageHeight int64 `json:"image_height"`
	// The width of the displayed ad.
	ImageWidth int64                                                    `json:"image_width"`
	JSON       userBenefitGetResponseBenefitAdsSubscriberPropertiesJSON `json:"-"`
}

// userBenefitGetResponseBenefitAdsSubscriberPropertiesJSON contains the JSON
// metadata for the struct [UserBenefitGetResponseBenefitAdsSubscriberProperties]
type userBenefitGetResponseBenefitAdsSubscriberPropertiesJSON struct {
	ImageHeight apijson.Field
	ImageWidth  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserBenefitGetResponseBenefitAdsSubscriberProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userBenefitGetResponseBenefitAdsSubscriberPropertiesJSON) RawJSON() string {
	return r.raw
}

type UserBenefitGetResponseBenefitAdsSubscriberType string

const (
	UserBenefitGetResponseBenefitAdsSubscriberTypeAds UserBenefitGetResponseBenefitAdsSubscriberType = "ads"
)

func (r UserBenefitGetResponseBenefitAdsSubscriberType) IsKnown() bool {
	switch r {
	case UserBenefitGetResponseBenefitAdsSubscriberTypeAds:
		return true
	}
	return false
}

type UserBenefitGetResponseBenefitDiscordSubscriber struct {
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
	// Properties available to subscribers for a benefit of type `discord`.
	Properties UserBenefitGetResponseBenefitDiscordSubscriberProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                               `json:"selectable,required"`
	Type       UserBenefitGetResponseBenefitDiscordSubscriberType `json:"type,required"`
	JSON       userBenefitGetResponseBenefitDiscordSubscriberJSON `json:"-"`
}

// userBenefitGetResponseBenefitDiscordSubscriberJSON contains the JSON metadata
// for the struct [UserBenefitGetResponseBenefitDiscordSubscriber]
type userBenefitGetResponseBenefitDiscordSubscriberJSON struct {
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

func (r *UserBenefitGetResponseBenefitDiscordSubscriber) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userBenefitGetResponseBenefitDiscordSubscriberJSON) RawJSON() string {
	return r.raw
}

func (r UserBenefitGetResponseBenefitDiscordSubscriber) implementsUserBenefitGetResponse() {}

// Properties available to subscribers for a benefit of type `discord`.
type UserBenefitGetResponseBenefitDiscordSubscriberProperties struct {
	// The ID of the Discord server.
	GuildID string                                                       `json:"guild_id,required"`
	JSON    userBenefitGetResponseBenefitDiscordSubscriberPropertiesJSON `json:"-"`
}

// userBenefitGetResponseBenefitDiscordSubscriberPropertiesJSON contains the JSON
// metadata for the struct
// [UserBenefitGetResponseBenefitDiscordSubscriberProperties]
type userBenefitGetResponseBenefitDiscordSubscriberPropertiesJSON struct {
	GuildID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserBenefitGetResponseBenefitDiscordSubscriberProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userBenefitGetResponseBenefitDiscordSubscriberPropertiesJSON) RawJSON() string {
	return r.raw
}

type UserBenefitGetResponseBenefitDiscordSubscriberType string

const (
	UserBenefitGetResponseBenefitDiscordSubscriberTypeDiscord UserBenefitGetResponseBenefitDiscordSubscriberType = "discord"
)

func (r UserBenefitGetResponseBenefitDiscordSubscriberType) IsKnown() bool {
	switch r {
	case UserBenefitGetResponseBenefitDiscordSubscriberTypeDiscord:
		return true
	}
	return false
}

type UserBenefitGetResponseBenefitCustomSubscriber struct {
	// The ID of the benefit.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the benefit is deletable.
	Deletable bool `json:"deletable,required"`
	// The description of the benefit.
	Description string                                               `json:"description,required"`
	Grants      []UserBenefitGetResponseBenefitCustomSubscriberGrant `json:"grants,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The ID of the organization owning the benefit.
	OrganizationID string `json:"organization_id,required" format:"uuid4"`
	// Properties available to subscribers for a benefit of type `custom`.
	Properties UserBenefitGetResponseBenefitCustomSubscriberProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                              `json:"selectable,required"`
	Type       UserBenefitGetResponseBenefitCustomSubscriberType `json:"type,required"`
	JSON       userBenefitGetResponseBenefitCustomSubscriberJSON `json:"-"`
}

// userBenefitGetResponseBenefitCustomSubscriberJSON contains the JSON metadata for
// the struct [UserBenefitGetResponseBenefitCustomSubscriber]
type userBenefitGetResponseBenefitCustomSubscriberJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Deletable      apijson.Field
	Description    apijson.Field
	Grants         apijson.Field
	ModifiedAt     apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	Selectable     apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *UserBenefitGetResponseBenefitCustomSubscriber) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userBenefitGetResponseBenefitCustomSubscriberJSON) RawJSON() string {
	return r.raw
}

func (r UserBenefitGetResponseBenefitCustomSubscriber) implementsUserBenefitGetResponse() {}

// A grant of a benefit to a user.
type UserBenefitGetResponseBenefitCustomSubscriberGrant struct {
	// The ID of the grant.
	ID string `json:"id,required" format:"uuid4"`
	// The ID of the benefit concerned by this grant.
	BenefitID string `json:"benefit_id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the benefit is granted.
	IsGranted bool `json:"is_granted,required"`
	// Whether the benefit is revoked.
	IsRevoked bool `json:"is_revoked,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The ID of the order that granted this benefit.
	OrderID string `json:"order_id,required,nullable" format:"uuid4"`
	// Properties for a benefit grant.
	Properties interface{} `json:"properties,required"`
	// The ID of the subscription that granted this benefit.
	SubscriptionID string `json:"subscription_id,required,nullable" format:"uuid4"`
	// The ID of the user concerned by this grant.
	UserID string `json:"user_id,required" format:"uuid4"`
	// The timestamp when the benefit was granted. If `None`, the benefit is not
	// granted.
	GrantedAt time.Time `json:"granted_at,nullable" format:"date-time"`
	// The timestamp when the benefit was revoked. If `None`, the benefit is not
	// revoked.
	RevokedAt time.Time                                              `json:"revoked_at,nullable" format:"date-time"`
	JSON      userBenefitGetResponseBenefitCustomSubscriberGrantJSON `json:"-"`
}

// userBenefitGetResponseBenefitCustomSubscriberGrantJSON contains the JSON
// metadata for the struct [UserBenefitGetResponseBenefitCustomSubscriberGrant]
type userBenefitGetResponseBenefitCustomSubscriberGrantJSON struct {
	ID             apijson.Field
	BenefitID      apijson.Field
	CreatedAt      apijson.Field
	IsGranted      apijson.Field
	IsRevoked      apijson.Field
	ModifiedAt     apijson.Field
	OrderID        apijson.Field
	Properties     apijson.Field
	SubscriptionID apijson.Field
	UserID         apijson.Field
	GrantedAt      apijson.Field
	RevokedAt      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *UserBenefitGetResponseBenefitCustomSubscriberGrant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userBenefitGetResponseBenefitCustomSubscriberGrantJSON) RawJSON() string {
	return r.raw
}

// Properties available to subscribers for a benefit of type `custom`.
type UserBenefitGetResponseBenefitCustomSubscriberProperties struct {
	// Private note to be shared with users who have this benefit granted.
	Note string                                                      `json:"note,required,nullable"`
	JSON userBenefitGetResponseBenefitCustomSubscriberPropertiesJSON `json:"-"`
}

// userBenefitGetResponseBenefitCustomSubscriberPropertiesJSON contains the JSON
// metadata for the struct
// [UserBenefitGetResponseBenefitCustomSubscriberProperties]
type userBenefitGetResponseBenefitCustomSubscriberPropertiesJSON struct {
	Note        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserBenefitGetResponseBenefitCustomSubscriberProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userBenefitGetResponseBenefitCustomSubscriberPropertiesJSON) RawJSON() string {
	return r.raw
}

type UserBenefitGetResponseBenefitCustomSubscriberType string

const (
	UserBenefitGetResponseBenefitCustomSubscriberTypeCustom UserBenefitGetResponseBenefitCustomSubscriberType = "custom"
)

func (r UserBenefitGetResponseBenefitCustomSubscriberType) IsKnown() bool {
	switch r {
	case UserBenefitGetResponseBenefitCustomSubscriberTypeCustom:
		return true
	}
	return false
}

type UserBenefitGetResponseBenefitGitHubRepositorySubscriber struct {
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
	// Properties available to subscribers for a benefit of type `github_repository`.
	Properties UserBenefitGetResponseBenefitGitHubRepositorySubscriberProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                                        `json:"selectable,required"`
	Type       UserBenefitGetResponseBenefitGitHubRepositorySubscriberType `json:"type,required"`
	JSON       userBenefitGetResponseBenefitGitHubRepositorySubscriberJSON `json:"-"`
}

// userBenefitGetResponseBenefitGitHubRepositorySubscriberJSON contains the JSON
// metadata for the struct
// [UserBenefitGetResponseBenefitGitHubRepositorySubscriber]
type userBenefitGetResponseBenefitGitHubRepositorySubscriberJSON struct {
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

func (r *UserBenefitGetResponseBenefitGitHubRepositorySubscriber) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userBenefitGetResponseBenefitGitHubRepositorySubscriberJSON) RawJSON() string {
	return r.raw
}

func (r UserBenefitGetResponseBenefitGitHubRepositorySubscriber) implementsUserBenefitGetResponse() {}

// Properties available to subscribers for a benefit of type `github_repository`.
type UserBenefitGetResponseBenefitGitHubRepositorySubscriberProperties struct {
	// The name of the repository.
	RepositoryName string `json:"repository_name,required"`
	// The owner of the repository.
	RepositoryOwner string                                                                `json:"repository_owner,required"`
	JSON            userBenefitGetResponseBenefitGitHubRepositorySubscriberPropertiesJSON `json:"-"`
}

// userBenefitGetResponseBenefitGitHubRepositorySubscriberPropertiesJSON contains
// the JSON metadata for the struct
// [UserBenefitGetResponseBenefitGitHubRepositorySubscriberProperties]
type userBenefitGetResponseBenefitGitHubRepositorySubscriberPropertiesJSON struct {
	RepositoryName  apijson.Field
	RepositoryOwner apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *UserBenefitGetResponseBenefitGitHubRepositorySubscriberProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userBenefitGetResponseBenefitGitHubRepositorySubscriberPropertiesJSON) RawJSON() string {
	return r.raw
}

type UserBenefitGetResponseBenefitGitHubRepositorySubscriberType string

const (
	UserBenefitGetResponseBenefitGitHubRepositorySubscriberTypeGitHubRepository UserBenefitGetResponseBenefitGitHubRepositorySubscriberType = "github_repository"
)

func (r UserBenefitGetResponseBenefitGitHubRepositorySubscriberType) IsKnown() bool {
	switch r {
	case UserBenefitGetResponseBenefitGitHubRepositorySubscriberTypeGitHubRepository:
		return true
	}
	return false
}

type UserBenefitGetResponseBenefitDownloadablesSubscriber struct {
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
	OrganizationID string                                                         `json:"organization_id,required" format:"uuid4"`
	Properties     UserBenefitGetResponseBenefitDownloadablesSubscriberProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                                     `json:"selectable,required"`
	Type       UserBenefitGetResponseBenefitDownloadablesSubscriberType `json:"type,required"`
	JSON       userBenefitGetResponseBenefitDownloadablesSubscriberJSON `json:"-"`
}

// userBenefitGetResponseBenefitDownloadablesSubscriberJSON contains the JSON
// metadata for the struct [UserBenefitGetResponseBenefitDownloadablesSubscriber]
type userBenefitGetResponseBenefitDownloadablesSubscriberJSON struct {
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

func (r *UserBenefitGetResponseBenefitDownloadablesSubscriber) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userBenefitGetResponseBenefitDownloadablesSubscriberJSON) RawJSON() string {
	return r.raw
}

func (r UserBenefitGetResponseBenefitDownloadablesSubscriber) implementsUserBenefitGetResponse() {}

type UserBenefitGetResponseBenefitDownloadablesSubscriberProperties struct {
	ActiveFiles []string                                                           `json:"active_files,required" format:"uuid4"`
	JSON        userBenefitGetResponseBenefitDownloadablesSubscriberPropertiesJSON `json:"-"`
}

// userBenefitGetResponseBenefitDownloadablesSubscriberPropertiesJSON contains the
// JSON metadata for the struct
// [UserBenefitGetResponseBenefitDownloadablesSubscriberProperties]
type userBenefitGetResponseBenefitDownloadablesSubscriberPropertiesJSON struct {
	ActiveFiles apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserBenefitGetResponseBenefitDownloadablesSubscriberProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userBenefitGetResponseBenefitDownloadablesSubscriberPropertiesJSON) RawJSON() string {
	return r.raw
}

type UserBenefitGetResponseBenefitDownloadablesSubscriberType string

const (
	UserBenefitGetResponseBenefitDownloadablesSubscriberTypeDownloadables UserBenefitGetResponseBenefitDownloadablesSubscriberType = "downloadables"
)

func (r UserBenefitGetResponseBenefitDownloadablesSubscriberType) IsKnown() bool {
	switch r {
	case UserBenefitGetResponseBenefitDownloadablesSubscriberTypeDownloadables:
		return true
	}
	return false
}

type UserBenefitGetResponseType string

const (
	UserBenefitGetResponseTypeArticles         UserBenefitGetResponseType = "articles"
	UserBenefitGetResponseTypeAds              UserBenefitGetResponseType = "ads"
	UserBenefitGetResponseTypeDiscord          UserBenefitGetResponseType = "discord"
	UserBenefitGetResponseTypeCustom           UserBenefitGetResponseType = "custom"
	UserBenefitGetResponseTypeGitHubRepository UserBenefitGetResponseType = "github_repository"
	UserBenefitGetResponseTypeDownloadables    UserBenefitGetResponseType = "downloadables"
)

func (r UserBenefitGetResponseType) IsKnown() bool {
	switch r {
	case UserBenefitGetResponseTypeArticles, UserBenefitGetResponseTypeAds, UserBenefitGetResponseTypeDiscord, UserBenefitGetResponseTypeCustom, UserBenefitGetResponseTypeGitHubRepository, UserBenefitGetResponseTypeDownloadables:
		return true
	}
	return false
}

type UserBenefitListResponse struct {
	Pagination UserBenefitListResponsePagination `json:"pagination,required"`
	Items      []UserBenefitListResponseItem     `json:"items"`
	JSON       userBenefitListResponseJSON       `json:"-"`
}

// userBenefitListResponseJSON contains the JSON metadata for the struct
// [UserBenefitListResponse]
type userBenefitListResponseJSON struct {
	Pagination  apijson.Field
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserBenefitListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userBenefitListResponseJSON) RawJSON() string {
	return r.raw
}

type UserBenefitListResponsePagination struct {
	MaxPage    int64                                 `json:"max_page,required"`
	TotalCount int64                                 `json:"total_count,required"`
	JSON       userBenefitListResponsePaginationJSON `json:"-"`
}

// userBenefitListResponsePaginationJSON contains the JSON metadata for the struct
// [UserBenefitListResponsePagination]
type userBenefitListResponsePaginationJSON struct {
	MaxPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserBenefitListResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userBenefitListResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type UserBenefitListResponseItem struct {
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The ID of the benefit.
	ID   string                           `json:"id,required" format:"uuid4"`
	Type UserBenefitListResponseItemsType `json:"type,required"`
	// The description of the benefit.
	Description string `json:"description,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool `json:"selectable,required"`
	// Whether the benefit is deletable.
	Deletable bool `json:"deletable,required"`
	// The ID of the organization owning the benefit.
	OrganizationID string `json:"organization_id,required" format:"uuid4"`
	// This field can have the runtime type of
	// [UserBenefitListResponseItemsBenefitArticlesSubscriberProperties],
	// [UserBenefitListResponseItemsBenefitAdsSubscriberProperties],
	// [UserBenefitListResponseItemsBenefitDiscordSubscriberProperties],
	// [UserBenefitListResponseItemsBenefitCustomSubscriberProperties],
	// [UserBenefitListResponseItemsBenefitGitHubRepositorySubscriberProperties],
	// [UserBenefitListResponseItemsBenefitDownloadablesSubscriberProperties].
	Properties interface{} `json:"properties"`
	// This field can have the runtime type of
	// [[]UserBenefitListResponseItemsBenefitAdsSubscriberGrant],
	// [[]UserBenefitListResponseItemsBenefitCustomSubscriberGrant].
	Grants interface{}                     `json:"grants,required"`
	JSON   userBenefitListResponseItemJSON `json:"-"`
	union  UserBenefitListResponseItemsUnion
}

// userBenefitListResponseItemJSON contains the JSON metadata for the struct
// [UserBenefitListResponseItem]
type userBenefitListResponseItemJSON struct {
	CreatedAt      apijson.Field
	ModifiedAt     apijson.Field
	ID             apijson.Field
	Type           apijson.Field
	Description    apijson.Field
	Selectable     apijson.Field
	Deletable      apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	Grants         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r userBenefitListResponseItemJSON) RawJSON() string {
	return r.raw
}

func (r *UserBenefitListResponseItem) UnmarshalJSON(data []byte) (err error) {
	*r = UserBenefitListResponseItem{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [UserBenefitListResponseItemsUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [UserBenefitListResponseItemsBenefitArticlesSubscriber],
// [UserBenefitListResponseItemsBenefitAdsSubscriber],
// [UserBenefitListResponseItemsBenefitDiscordSubscriber],
// [UserBenefitListResponseItemsBenefitCustomSubscriber],
// [UserBenefitListResponseItemsBenefitGitHubRepositorySubscriber],
// [UserBenefitListResponseItemsBenefitDownloadablesSubscriber].
func (r UserBenefitListResponseItem) AsUnion() UserBenefitListResponseItemsUnion {
	return r.union
}

// Union satisfied by [UserBenefitListResponseItemsBenefitArticlesSubscriber],
// [UserBenefitListResponseItemsBenefitAdsSubscriber],
// [UserBenefitListResponseItemsBenefitDiscordSubscriber],
// [UserBenefitListResponseItemsBenefitCustomSubscriber],
// [UserBenefitListResponseItemsBenefitGitHubRepositorySubscriber] or
// [UserBenefitListResponseItemsBenefitDownloadablesSubscriber].
type UserBenefitListResponseItemsUnion interface {
	implementsUserBenefitListResponseItem()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserBenefitListResponseItemsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(UserBenefitListResponseItemsBenefitArticlesSubscriber{}),
			DiscriminatorValue: "articles",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(UserBenefitListResponseItemsBenefitAdsSubscriber{}),
			DiscriminatorValue: "ads",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(UserBenefitListResponseItemsBenefitDiscordSubscriber{}),
			DiscriminatorValue: "discord",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(UserBenefitListResponseItemsBenefitCustomSubscriber{}),
			DiscriminatorValue: "custom",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(UserBenefitListResponseItemsBenefitGitHubRepositorySubscriber{}),
			DiscriminatorValue: "github_repository",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(UserBenefitListResponseItemsBenefitDownloadablesSubscriber{}),
			DiscriminatorValue: "downloadables",
		},
	)
}

type UserBenefitListResponseItemsBenefitArticlesSubscriber struct {
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
	// Properties available to subscribers for a benefit of type `articles`.
	Properties UserBenefitListResponseItemsBenefitArticlesSubscriberProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                                      `json:"selectable,required"`
	Type       UserBenefitListResponseItemsBenefitArticlesSubscriberType `json:"type,required"`
	JSON       userBenefitListResponseItemsBenefitArticlesSubscriberJSON `json:"-"`
}

// userBenefitListResponseItemsBenefitArticlesSubscriberJSON contains the JSON
// metadata for the struct [UserBenefitListResponseItemsBenefitArticlesSubscriber]
type userBenefitListResponseItemsBenefitArticlesSubscriberJSON struct {
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

func (r *UserBenefitListResponseItemsBenefitArticlesSubscriber) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userBenefitListResponseItemsBenefitArticlesSubscriberJSON) RawJSON() string {
	return r.raw
}

func (r UserBenefitListResponseItemsBenefitArticlesSubscriber) implementsUserBenefitListResponseItem() {
}

// Properties available to subscribers for a benefit of type `articles`.
type UserBenefitListResponseItemsBenefitArticlesSubscriberProperties struct {
	// Whether the user can access paid articles.
	PaidArticles bool                                                                `json:"paid_articles,required"`
	JSON         userBenefitListResponseItemsBenefitArticlesSubscriberPropertiesJSON `json:"-"`
}

// userBenefitListResponseItemsBenefitArticlesSubscriberPropertiesJSON contains the
// JSON metadata for the struct
// [UserBenefitListResponseItemsBenefitArticlesSubscriberProperties]
type userBenefitListResponseItemsBenefitArticlesSubscriberPropertiesJSON struct {
	PaidArticles apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *UserBenefitListResponseItemsBenefitArticlesSubscriberProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userBenefitListResponseItemsBenefitArticlesSubscriberPropertiesJSON) RawJSON() string {
	return r.raw
}

type UserBenefitListResponseItemsBenefitArticlesSubscriberType string

const (
	UserBenefitListResponseItemsBenefitArticlesSubscriberTypeArticles UserBenefitListResponseItemsBenefitArticlesSubscriberType = "articles"
)

func (r UserBenefitListResponseItemsBenefitArticlesSubscriberType) IsKnown() bool {
	switch r {
	case UserBenefitListResponseItemsBenefitArticlesSubscriberTypeArticles:
		return true
	}
	return false
}

type UserBenefitListResponseItemsBenefitAdsSubscriber struct {
	// The ID of the benefit.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the benefit is deletable.
	Deletable bool `json:"deletable,required"`
	// The description of the benefit.
	Description string                                                  `json:"description,required"`
	Grants      []UserBenefitListResponseItemsBenefitAdsSubscriberGrant `json:"grants,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The ID of the organization owning the benefit.
	OrganizationID string `json:"organization_id,required" format:"uuid4"`
	// Properties for a benefit of type `ads`.
	Properties UserBenefitListResponseItemsBenefitAdsSubscriberProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                                 `json:"selectable,required"`
	Type       UserBenefitListResponseItemsBenefitAdsSubscriberType `json:"type,required"`
	JSON       userBenefitListResponseItemsBenefitAdsSubscriberJSON `json:"-"`
}

// userBenefitListResponseItemsBenefitAdsSubscriberJSON contains the JSON metadata
// for the struct [UserBenefitListResponseItemsBenefitAdsSubscriber]
type userBenefitListResponseItemsBenefitAdsSubscriberJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Deletable      apijson.Field
	Description    apijson.Field
	Grants         apijson.Field
	ModifiedAt     apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	Selectable     apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *UserBenefitListResponseItemsBenefitAdsSubscriber) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userBenefitListResponseItemsBenefitAdsSubscriberJSON) RawJSON() string {
	return r.raw
}

func (r UserBenefitListResponseItemsBenefitAdsSubscriber) implementsUserBenefitListResponseItem() {}

type UserBenefitListResponseItemsBenefitAdsSubscriberGrant struct {
	// The ID of the grant.
	ID string `json:"id,required" format:"uuid4"`
	// The ID of the benefit concerned by this grant.
	BenefitID string `json:"benefit_id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the benefit is granted.
	IsGranted bool `json:"is_granted,required"`
	// Whether the benefit is revoked.
	IsRevoked bool `json:"is_revoked,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The ID of the order that granted this benefit.
	OrderID    string                                                           `json:"order_id,required,nullable" format:"uuid4"`
	Properties UserBenefitListResponseItemsBenefitAdsSubscriberGrantsProperties `json:"properties,required"`
	// The ID of the subscription that granted this benefit.
	SubscriptionID string `json:"subscription_id,required,nullable" format:"uuid4"`
	// The ID of the user concerned by this grant.
	UserID string `json:"user_id,required" format:"uuid4"`
	// The timestamp when the benefit was granted. If `None`, the benefit is not
	// granted.
	GrantedAt time.Time `json:"granted_at,nullable" format:"date-time"`
	// The timestamp when the benefit was revoked. If `None`, the benefit is not
	// revoked.
	RevokedAt time.Time                                                 `json:"revoked_at,nullable" format:"date-time"`
	JSON      userBenefitListResponseItemsBenefitAdsSubscriberGrantJSON `json:"-"`
}

// userBenefitListResponseItemsBenefitAdsSubscriberGrantJSON contains the JSON
// metadata for the struct [UserBenefitListResponseItemsBenefitAdsSubscriberGrant]
type userBenefitListResponseItemsBenefitAdsSubscriberGrantJSON struct {
	ID             apijson.Field
	BenefitID      apijson.Field
	CreatedAt      apijson.Field
	IsGranted      apijson.Field
	IsRevoked      apijson.Field
	ModifiedAt     apijson.Field
	OrderID        apijson.Field
	Properties     apijson.Field
	SubscriptionID apijson.Field
	UserID         apijson.Field
	GrantedAt      apijson.Field
	RevokedAt      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *UserBenefitListResponseItemsBenefitAdsSubscriberGrant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userBenefitListResponseItemsBenefitAdsSubscriberGrantJSON) RawJSON() string {
	return r.raw
}

type UserBenefitListResponseItemsBenefitAdsSubscriberGrantsProperties struct {
	// The ID of the enabled advertisement campaign for this benefit grant.
	AdvertisementCampaignID string                                                               `json:"advertisement_campaign_id,nullable" format:"uuid4"`
	JSON                    userBenefitListResponseItemsBenefitAdsSubscriberGrantsPropertiesJSON `json:"-"`
}

// userBenefitListResponseItemsBenefitAdsSubscriberGrantsPropertiesJSON contains
// the JSON metadata for the struct
// [UserBenefitListResponseItemsBenefitAdsSubscriberGrantsProperties]
type userBenefitListResponseItemsBenefitAdsSubscriberGrantsPropertiesJSON struct {
	AdvertisementCampaignID apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *UserBenefitListResponseItemsBenefitAdsSubscriberGrantsProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userBenefitListResponseItemsBenefitAdsSubscriberGrantsPropertiesJSON) RawJSON() string {
	return r.raw
}

// Properties for a benefit of type `ads`.
type UserBenefitListResponseItemsBenefitAdsSubscriberProperties struct {
	// The height of the displayed ad.
	ImageHeight int64 `json:"image_height"`
	// The width of the displayed ad.
	ImageWidth int64                                                          `json:"image_width"`
	JSON       userBenefitListResponseItemsBenefitAdsSubscriberPropertiesJSON `json:"-"`
}

// userBenefitListResponseItemsBenefitAdsSubscriberPropertiesJSON contains the JSON
// metadata for the struct
// [UserBenefitListResponseItemsBenefitAdsSubscriberProperties]
type userBenefitListResponseItemsBenefitAdsSubscriberPropertiesJSON struct {
	ImageHeight apijson.Field
	ImageWidth  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserBenefitListResponseItemsBenefitAdsSubscriberProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userBenefitListResponseItemsBenefitAdsSubscriberPropertiesJSON) RawJSON() string {
	return r.raw
}

type UserBenefitListResponseItemsBenefitAdsSubscriberType string

const (
	UserBenefitListResponseItemsBenefitAdsSubscriberTypeAds UserBenefitListResponseItemsBenefitAdsSubscriberType = "ads"
)

func (r UserBenefitListResponseItemsBenefitAdsSubscriberType) IsKnown() bool {
	switch r {
	case UserBenefitListResponseItemsBenefitAdsSubscriberTypeAds:
		return true
	}
	return false
}

type UserBenefitListResponseItemsBenefitDiscordSubscriber struct {
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
	// Properties available to subscribers for a benefit of type `discord`.
	Properties UserBenefitListResponseItemsBenefitDiscordSubscriberProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                                     `json:"selectable,required"`
	Type       UserBenefitListResponseItemsBenefitDiscordSubscriberType `json:"type,required"`
	JSON       userBenefitListResponseItemsBenefitDiscordSubscriberJSON `json:"-"`
}

// userBenefitListResponseItemsBenefitDiscordSubscriberJSON contains the JSON
// metadata for the struct [UserBenefitListResponseItemsBenefitDiscordSubscriber]
type userBenefitListResponseItemsBenefitDiscordSubscriberJSON struct {
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

func (r *UserBenefitListResponseItemsBenefitDiscordSubscriber) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userBenefitListResponseItemsBenefitDiscordSubscriberJSON) RawJSON() string {
	return r.raw
}

func (r UserBenefitListResponseItemsBenefitDiscordSubscriber) implementsUserBenefitListResponseItem() {
}

// Properties available to subscribers for a benefit of type `discord`.
type UserBenefitListResponseItemsBenefitDiscordSubscriberProperties struct {
	// The ID of the Discord server.
	GuildID string                                                             `json:"guild_id,required"`
	JSON    userBenefitListResponseItemsBenefitDiscordSubscriberPropertiesJSON `json:"-"`
}

// userBenefitListResponseItemsBenefitDiscordSubscriberPropertiesJSON contains the
// JSON metadata for the struct
// [UserBenefitListResponseItemsBenefitDiscordSubscriberProperties]
type userBenefitListResponseItemsBenefitDiscordSubscriberPropertiesJSON struct {
	GuildID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserBenefitListResponseItemsBenefitDiscordSubscriberProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userBenefitListResponseItemsBenefitDiscordSubscriberPropertiesJSON) RawJSON() string {
	return r.raw
}

type UserBenefitListResponseItemsBenefitDiscordSubscriberType string

const (
	UserBenefitListResponseItemsBenefitDiscordSubscriberTypeDiscord UserBenefitListResponseItemsBenefitDiscordSubscriberType = "discord"
)

func (r UserBenefitListResponseItemsBenefitDiscordSubscriberType) IsKnown() bool {
	switch r {
	case UserBenefitListResponseItemsBenefitDiscordSubscriberTypeDiscord:
		return true
	}
	return false
}

type UserBenefitListResponseItemsBenefitCustomSubscriber struct {
	// The ID of the benefit.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the benefit is deletable.
	Deletable bool `json:"deletable,required"`
	// The description of the benefit.
	Description string                                                     `json:"description,required"`
	Grants      []UserBenefitListResponseItemsBenefitCustomSubscriberGrant `json:"grants,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The ID of the organization owning the benefit.
	OrganizationID string `json:"organization_id,required" format:"uuid4"`
	// Properties available to subscribers for a benefit of type `custom`.
	Properties UserBenefitListResponseItemsBenefitCustomSubscriberProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                                    `json:"selectable,required"`
	Type       UserBenefitListResponseItemsBenefitCustomSubscriberType `json:"type,required"`
	JSON       userBenefitListResponseItemsBenefitCustomSubscriberJSON `json:"-"`
}

// userBenefitListResponseItemsBenefitCustomSubscriberJSON contains the JSON
// metadata for the struct [UserBenefitListResponseItemsBenefitCustomSubscriber]
type userBenefitListResponseItemsBenefitCustomSubscriberJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Deletable      apijson.Field
	Description    apijson.Field
	Grants         apijson.Field
	ModifiedAt     apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	Selectable     apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *UserBenefitListResponseItemsBenefitCustomSubscriber) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userBenefitListResponseItemsBenefitCustomSubscriberJSON) RawJSON() string {
	return r.raw
}

func (r UserBenefitListResponseItemsBenefitCustomSubscriber) implementsUserBenefitListResponseItem() {
}

// A grant of a benefit to a user.
type UserBenefitListResponseItemsBenefitCustomSubscriberGrant struct {
	// The ID of the grant.
	ID string `json:"id,required" format:"uuid4"`
	// The ID of the benefit concerned by this grant.
	BenefitID string `json:"benefit_id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the benefit is granted.
	IsGranted bool `json:"is_granted,required"`
	// Whether the benefit is revoked.
	IsRevoked bool `json:"is_revoked,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The ID of the order that granted this benefit.
	OrderID string `json:"order_id,required,nullable" format:"uuid4"`
	// Properties for a benefit grant.
	Properties interface{} `json:"properties,required"`
	// The ID of the subscription that granted this benefit.
	SubscriptionID string `json:"subscription_id,required,nullable" format:"uuid4"`
	// The ID of the user concerned by this grant.
	UserID string `json:"user_id,required" format:"uuid4"`
	// The timestamp when the benefit was granted. If `None`, the benefit is not
	// granted.
	GrantedAt time.Time `json:"granted_at,nullable" format:"date-time"`
	// The timestamp when the benefit was revoked. If `None`, the benefit is not
	// revoked.
	RevokedAt time.Time                                                    `json:"revoked_at,nullable" format:"date-time"`
	JSON      userBenefitListResponseItemsBenefitCustomSubscriberGrantJSON `json:"-"`
}

// userBenefitListResponseItemsBenefitCustomSubscriberGrantJSON contains the JSON
// metadata for the struct
// [UserBenefitListResponseItemsBenefitCustomSubscriberGrant]
type userBenefitListResponseItemsBenefitCustomSubscriberGrantJSON struct {
	ID             apijson.Field
	BenefitID      apijson.Field
	CreatedAt      apijson.Field
	IsGranted      apijson.Field
	IsRevoked      apijson.Field
	ModifiedAt     apijson.Field
	OrderID        apijson.Field
	Properties     apijson.Field
	SubscriptionID apijson.Field
	UserID         apijson.Field
	GrantedAt      apijson.Field
	RevokedAt      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *UserBenefitListResponseItemsBenefitCustomSubscriberGrant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userBenefitListResponseItemsBenefitCustomSubscriberGrantJSON) RawJSON() string {
	return r.raw
}

// Properties available to subscribers for a benefit of type `custom`.
type UserBenefitListResponseItemsBenefitCustomSubscriberProperties struct {
	// Private note to be shared with users who have this benefit granted.
	Note string                                                            `json:"note,required,nullable"`
	JSON userBenefitListResponseItemsBenefitCustomSubscriberPropertiesJSON `json:"-"`
}

// userBenefitListResponseItemsBenefitCustomSubscriberPropertiesJSON contains the
// JSON metadata for the struct
// [UserBenefitListResponseItemsBenefitCustomSubscriberProperties]
type userBenefitListResponseItemsBenefitCustomSubscriberPropertiesJSON struct {
	Note        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserBenefitListResponseItemsBenefitCustomSubscriberProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userBenefitListResponseItemsBenefitCustomSubscriberPropertiesJSON) RawJSON() string {
	return r.raw
}

type UserBenefitListResponseItemsBenefitCustomSubscriberType string

const (
	UserBenefitListResponseItemsBenefitCustomSubscriberTypeCustom UserBenefitListResponseItemsBenefitCustomSubscriberType = "custom"
)

func (r UserBenefitListResponseItemsBenefitCustomSubscriberType) IsKnown() bool {
	switch r {
	case UserBenefitListResponseItemsBenefitCustomSubscriberTypeCustom:
		return true
	}
	return false
}

type UserBenefitListResponseItemsBenefitGitHubRepositorySubscriber struct {
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
	// Properties available to subscribers for a benefit of type `github_repository`.
	Properties UserBenefitListResponseItemsBenefitGitHubRepositorySubscriberProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                                              `json:"selectable,required"`
	Type       UserBenefitListResponseItemsBenefitGitHubRepositorySubscriberType `json:"type,required"`
	JSON       userBenefitListResponseItemsBenefitGitHubRepositorySubscriberJSON `json:"-"`
}

// userBenefitListResponseItemsBenefitGitHubRepositorySubscriberJSON contains the
// JSON metadata for the struct
// [UserBenefitListResponseItemsBenefitGitHubRepositorySubscriber]
type userBenefitListResponseItemsBenefitGitHubRepositorySubscriberJSON struct {
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

func (r *UserBenefitListResponseItemsBenefitGitHubRepositorySubscriber) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userBenefitListResponseItemsBenefitGitHubRepositorySubscriberJSON) RawJSON() string {
	return r.raw
}

func (r UserBenefitListResponseItemsBenefitGitHubRepositorySubscriber) implementsUserBenefitListResponseItem() {
}

// Properties available to subscribers for a benefit of type `github_repository`.
type UserBenefitListResponseItemsBenefitGitHubRepositorySubscriberProperties struct {
	// The name of the repository.
	RepositoryName string `json:"repository_name,required"`
	// The owner of the repository.
	RepositoryOwner string                                                                      `json:"repository_owner,required"`
	JSON            userBenefitListResponseItemsBenefitGitHubRepositorySubscriberPropertiesJSON `json:"-"`
}

// userBenefitListResponseItemsBenefitGitHubRepositorySubscriberPropertiesJSON
// contains the JSON metadata for the struct
// [UserBenefitListResponseItemsBenefitGitHubRepositorySubscriberProperties]
type userBenefitListResponseItemsBenefitGitHubRepositorySubscriberPropertiesJSON struct {
	RepositoryName  apijson.Field
	RepositoryOwner apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *UserBenefitListResponseItemsBenefitGitHubRepositorySubscriberProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userBenefitListResponseItemsBenefitGitHubRepositorySubscriberPropertiesJSON) RawJSON() string {
	return r.raw
}

type UserBenefitListResponseItemsBenefitGitHubRepositorySubscriberType string

const (
	UserBenefitListResponseItemsBenefitGitHubRepositorySubscriberTypeGitHubRepository UserBenefitListResponseItemsBenefitGitHubRepositorySubscriberType = "github_repository"
)

func (r UserBenefitListResponseItemsBenefitGitHubRepositorySubscriberType) IsKnown() bool {
	switch r {
	case UserBenefitListResponseItemsBenefitGitHubRepositorySubscriberTypeGitHubRepository:
		return true
	}
	return false
}

type UserBenefitListResponseItemsBenefitDownloadablesSubscriber struct {
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
	OrganizationID string                                                               `json:"organization_id,required" format:"uuid4"`
	Properties     UserBenefitListResponseItemsBenefitDownloadablesSubscriberProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                                           `json:"selectable,required"`
	Type       UserBenefitListResponseItemsBenefitDownloadablesSubscriberType `json:"type,required"`
	JSON       userBenefitListResponseItemsBenefitDownloadablesSubscriberJSON `json:"-"`
}

// userBenefitListResponseItemsBenefitDownloadablesSubscriberJSON contains the JSON
// metadata for the struct
// [UserBenefitListResponseItemsBenefitDownloadablesSubscriber]
type userBenefitListResponseItemsBenefitDownloadablesSubscriberJSON struct {
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

func (r *UserBenefitListResponseItemsBenefitDownloadablesSubscriber) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userBenefitListResponseItemsBenefitDownloadablesSubscriberJSON) RawJSON() string {
	return r.raw
}

func (r UserBenefitListResponseItemsBenefitDownloadablesSubscriber) implementsUserBenefitListResponseItem() {
}

type UserBenefitListResponseItemsBenefitDownloadablesSubscriberProperties struct {
	ActiveFiles []string                                                                 `json:"active_files,required" format:"uuid4"`
	JSON        userBenefitListResponseItemsBenefitDownloadablesSubscriberPropertiesJSON `json:"-"`
}

// userBenefitListResponseItemsBenefitDownloadablesSubscriberPropertiesJSON
// contains the JSON metadata for the struct
// [UserBenefitListResponseItemsBenefitDownloadablesSubscriberProperties]
type userBenefitListResponseItemsBenefitDownloadablesSubscriberPropertiesJSON struct {
	ActiveFiles apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserBenefitListResponseItemsBenefitDownloadablesSubscriberProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userBenefitListResponseItemsBenefitDownloadablesSubscriberPropertiesJSON) RawJSON() string {
	return r.raw
}

type UserBenefitListResponseItemsBenefitDownloadablesSubscriberType string

const (
	UserBenefitListResponseItemsBenefitDownloadablesSubscriberTypeDownloadables UserBenefitListResponseItemsBenefitDownloadablesSubscriberType = "downloadables"
)

func (r UserBenefitListResponseItemsBenefitDownloadablesSubscriberType) IsKnown() bool {
	switch r {
	case UserBenefitListResponseItemsBenefitDownloadablesSubscriberTypeDownloadables:
		return true
	}
	return false
}

type UserBenefitListResponseItemsType string

const (
	UserBenefitListResponseItemsTypeArticles         UserBenefitListResponseItemsType = "articles"
	UserBenefitListResponseItemsTypeAds              UserBenefitListResponseItemsType = "ads"
	UserBenefitListResponseItemsTypeDiscord          UserBenefitListResponseItemsType = "discord"
	UserBenefitListResponseItemsTypeCustom           UserBenefitListResponseItemsType = "custom"
	UserBenefitListResponseItemsTypeGitHubRepository UserBenefitListResponseItemsType = "github_repository"
	UserBenefitListResponseItemsTypeDownloadables    UserBenefitListResponseItemsType = "downloadables"
)

func (r UserBenefitListResponseItemsType) IsKnown() bool {
	switch r {
	case UserBenefitListResponseItemsTypeArticles, UserBenefitListResponseItemsTypeAds, UserBenefitListResponseItemsTypeDiscord, UserBenefitListResponseItemsTypeCustom, UserBenefitListResponseItemsTypeGitHubRepository, UserBenefitListResponseItemsTypeDownloadables:
		return true
	}
	return false
}

type UserBenefitListParams struct {
	// Size of a page, defaults to 10. Maximum is 100.
	Limit param.Field[int64] `query:"limit"`
	// Filter by order ID.
	OrderID param.Field[UserBenefitListParamsOrderIDUnion] `query:"order_id" format:"uuid4"`
	// Filter by organization ID.
	OrganizationID param.Field[UserBenefitListParamsOrganizationIDUnion] `query:"organization_id" format:"uuid4"`
	// Page number, defaults to 1.
	Page param.Field[int64] `query:"page"`
	// Sorting criterion. Several criteria can be used simultaneously and will be
	// applied in order. Add a minus sign `-` before the criteria name to sort by
	// descending order.
	Sorting param.Field[[]string] `query:"sorting"`
	// Filter by subscription ID.
	SubscriptionID param.Field[UserBenefitListParamsSubscriptionIDUnion] `query:"subscription_id" format:"uuid4"`
	// Filter by benefit type.
	Type param.Field[UserBenefitListParamsTypeUnion] `query:"type"`
}

// URLQuery serializes [UserBenefitListParams]'s query parameters as `url.Values`.
func (r UserBenefitListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by order ID.
//
// Satisfied by [shared.UnionString], [UserBenefitListParamsOrderIDArray].
type UserBenefitListParamsOrderIDUnion interface {
	ImplementsUserBenefitListParamsOrderIDUnion()
}

type UserBenefitListParamsOrderIDArray []string

func (r UserBenefitListParamsOrderIDArray) ImplementsUserBenefitListParamsOrderIDUnion() {}

// Filter by organization ID.
//
// Satisfied by [shared.UnionString], [UserBenefitListParamsOrganizationIDArray].
type UserBenefitListParamsOrganizationIDUnion interface {
	ImplementsUserBenefitListParamsOrganizationIDUnion()
}

type UserBenefitListParamsOrganizationIDArray []string

func (r UserBenefitListParamsOrganizationIDArray) ImplementsUserBenefitListParamsOrganizationIDUnion() {
}

// Filter by subscription ID.
//
// Satisfied by [shared.UnionString], [UserBenefitListParamsSubscriptionIDArray].
type UserBenefitListParamsSubscriptionIDUnion interface {
	ImplementsUserBenefitListParamsSubscriptionIDUnion()
}

type UserBenefitListParamsSubscriptionIDArray []string

func (r UserBenefitListParamsSubscriptionIDArray) ImplementsUserBenefitListParamsSubscriptionIDUnion() {
}

// Filter by benefit type.
//
// Satisfied by [UserBenefitListParamsTypeBenefitType],
// [UserBenefitListParamsTypeArray].
type UserBenefitListParamsTypeUnion interface {
	implementsUserBenefitListParamsTypeUnion()
}

type UserBenefitListParamsTypeBenefitType string

const (
	UserBenefitListParamsTypeBenefitTypeCustom           UserBenefitListParamsTypeBenefitType = "custom"
	UserBenefitListParamsTypeBenefitTypeArticles         UserBenefitListParamsTypeBenefitType = "articles"
	UserBenefitListParamsTypeBenefitTypeAds              UserBenefitListParamsTypeBenefitType = "ads"
	UserBenefitListParamsTypeBenefitTypeDiscord          UserBenefitListParamsTypeBenefitType = "discord"
	UserBenefitListParamsTypeBenefitTypeGitHubRepository UserBenefitListParamsTypeBenefitType = "github_repository"
	UserBenefitListParamsTypeBenefitTypeDownloadables    UserBenefitListParamsTypeBenefitType = "downloadables"
)

func (r UserBenefitListParamsTypeBenefitType) IsKnown() bool {
	switch r {
	case UserBenefitListParamsTypeBenefitTypeCustom, UserBenefitListParamsTypeBenefitTypeArticles, UserBenefitListParamsTypeBenefitTypeAds, UserBenefitListParamsTypeBenefitTypeDiscord, UserBenefitListParamsTypeBenefitTypeGitHubRepository, UserBenefitListParamsTypeBenefitTypeDownloadables:
		return true
	}
	return false
}

func (r UserBenefitListParamsTypeBenefitType) implementsUserBenefitListParamsTypeUnion() {}

type UserBenefitListParamsTypeArray []UserBenefitListParamsTypeArray

func (r UserBenefitListParamsTypeArray) implementsUserBenefitListParamsTypeUnion() {}
