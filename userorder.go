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

	"github.com/stainless-sdks/polar-go/internal/apijson"
	"github.com/stainless-sdks/polar-go/internal/apiquery"
	"github.com/stainless-sdks/polar-go/internal/param"
	"github.com/stainless-sdks/polar-go/internal/requestconfig"
	"github.com/stainless-sdks/polar-go/option"
	"github.com/tidwall/gjson"
)

// UserOrderService contains methods and other services that help with interacting
// with the polar API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserOrderService] method instead.
type UserOrderService struct {
	Options []option.RequestOption
}

// NewUserOrderService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserOrderService(opts ...option.RequestOption) (r *UserOrderService) {
	r = &UserOrderService{}
	r.Options = opts
	return
}

// Get an order by ID.
func (r *UserOrderService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *UserOrderGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/users/orders/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List my orders.
func (r *UserOrderService) List(ctx context.Context, query UserOrderListParams, opts ...option.RequestOption) (res *UserOrderListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/users/orders/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get an order's invoice data.
func (r *UserOrderService) Invoice(ctx context.Context, id string, opts ...option.RequestOption) (res *UserOrderInvoiceResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/users/orders/%s/invoice", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type UserOrderGetResponse struct {
	ID     string `json:"id,required" format:"uuid4"`
	Amount int64  `json:"amount,required"`
	// Creation timestamp of the object.
	CreatedAt time.Time                   `json:"created_at,required" format:"date-time"`
	Currency  string                      `json:"currency,required"`
	Product   UserOrderGetResponseProduct `json:"product,required"`
	ProductID string                      `json:"product_id,required" format:"uuid4"`
	// A recurring price for a product, i.e. a subscription.
	ProductPrice   UserOrderGetResponseProductPrice `json:"product_price,required"`
	ProductPriceID string                           `json:"product_price_id,required" format:"uuid4"`
	TaxAmount      int64                            `json:"tax_amount,required"`
	UserID         string                           `json:"user_id,required" format:"uuid4"`
	// Last modification timestamp of the object.
	ModifiedAt     time.Time                        `json:"modified_at,nullable" format:"date-time"`
	Subscription   UserOrderGetResponseSubscription `json:"subscription,nullable"`
	SubscriptionID string                           `json:"subscription_id,nullable" format:"uuid4"`
	JSON           userOrderGetResponseJSON         `json:"-"`
}

// userOrderGetResponseJSON contains the JSON metadata for the struct
// [UserOrderGetResponse]
type userOrderGetResponseJSON struct {
	ID             apijson.Field
	Amount         apijson.Field
	CreatedAt      apijson.Field
	Currency       apijson.Field
	Product        apijson.Field
	ProductID      apijson.Field
	ProductPrice   apijson.Field
	ProductPriceID apijson.Field
	TaxAmount      apijson.Field
	UserID         apijson.Field
	ModifiedAt     apijson.Field
	Subscription   apijson.Field
	SubscriptionID apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *UserOrderGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userOrderGetResponseJSON) RawJSON() string {
	return r.raw
}

type UserOrderGetResponseProduct struct {
	// The ID of the product.
	ID string `json:"id,required" format:"uuid4"`
	// The benefits granted by the product.
	Benefits []UserOrderGetResponseProductBenefit `json:"benefits,required"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the product is archived and no longer available.
	IsArchived bool `json:"is_archived,required"`
	// Whether the product is a subscription tier.
	IsRecurring bool `json:"is_recurring,required"`
	// The medias associated to the product.
	Medias []ProductMediaFileReadOutput `json:"medias,required"`
	// The name of the product.
	Name string `json:"name,required"`
	// The ID of the organization owning the product.
	OrganizationID string `json:"organization_id,required" format:"uuid4"`
	// List of available prices for this product.
	Prices []UserOrderGetResponseProductPrice `json:"prices,required"`
	// The description of the product.
	Description   string `json:"description,nullable"`
	IsHighlighted bool   `json:"is_highlighted,nullable"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time                       `json:"modified_at,nullable" format:"date-time"`
	Type       UserOrderGetResponseProductType `json:"type,nullable"`
	JSON       userOrderGetResponseProductJSON `json:"-"`
}

// userOrderGetResponseProductJSON contains the JSON metadata for the struct
// [UserOrderGetResponseProduct]
type userOrderGetResponseProductJSON struct {
	ID             apijson.Field
	Benefits       apijson.Field
	CreatedAt      apijson.Field
	IsArchived     apijson.Field
	IsRecurring    apijson.Field
	Medias         apijson.Field
	Name           apijson.Field
	OrganizationID apijson.Field
	Prices         apijson.Field
	Description    apijson.Field
	IsHighlighted  apijson.Field
	ModifiedAt     apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *UserOrderGetResponseProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userOrderGetResponseProductJSON) RawJSON() string {
	return r.raw
}

// A benefit of type `articles`.
//
// Use it to grant access to posts.
type UserOrderGetResponseProductBenefit struct {
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,nullable" format:"date-time"`
	// The ID of the benefit.
	ID string `json:"id,required" format:"uuid4"`
	// The type of the benefit.
	Type UserOrderGetResponseProductBenefitsType `json:"type,required"`
	// The description of the benefit.
	Description string `json:"description,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool `json:"selectable,required"`
	// Whether the benefit is deletable.
	Deletable bool `json:"deletable,required"`
	// The ID of the organization owning the benefit.
	OrganizationID string `json:"organization_id,required" format:"uuid4"`
	// This field can have the runtime type of
	// [UserOrderGetResponseProductBenefitsBenefitArticlesProperties].
	Properties interface{}                            `json:"properties,required"`
	JSON       userOrderGetResponseProductBenefitJSON `json:"-"`
	union      UserOrderGetResponseProductBenefitsUnion
}

// userOrderGetResponseProductBenefitJSON contains the JSON metadata for the struct
// [UserOrderGetResponseProductBenefit]
type userOrderGetResponseProductBenefitJSON struct {
	CreatedAt      apijson.Field
	ModifiedAt     apijson.Field
	ID             apijson.Field
	Type           apijson.Field
	Description    apijson.Field
	Selectable     apijson.Field
	Deletable      apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r userOrderGetResponseProductBenefitJSON) RawJSON() string {
	return r.raw
}

func (r *UserOrderGetResponseProductBenefit) UnmarshalJSON(data []byte) (err error) {
	*r = UserOrderGetResponseProductBenefit{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [UserOrderGetResponseProductBenefitsUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [UserOrderGetResponseProductBenefitsBenefitBase],
// [UserOrderGetResponseProductBenefitsBenefitArticles].
func (r UserOrderGetResponseProductBenefit) AsUnion() UserOrderGetResponseProductBenefitsUnion {
	return r.union
}

// A benefit of type `articles`.
//
// Use it to grant access to posts.
//
// Union satisfied by [UserOrderGetResponseProductBenefitsBenefitBase] or
// [UserOrderGetResponseProductBenefitsBenefitArticles].
type UserOrderGetResponseProductBenefitsUnion interface {
	implementsUserOrderGetResponseProductBenefit()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserOrderGetResponseProductBenefitsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserOrderGetResponseProductBenefitsBenefitBase{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserOrderGetResponseProductBenefitsBenefitArticles{}),
		},
	)
}

type UserOrderGetResponseProductBenefitsBenefitBase struct {
	// The ID of the benefit.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the benefit is deletable.
	Deletable bool `json:"deletable,required"`
	// The description of the benefit.
	Description string `json:"description,required"`
	// The ID of the organization owning the benefit.
	OrganizationID string `json:"organization_id,required" format:"uuid4"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool `json:"selectable,required"`
	// The type of the benefit.
	Type UserOrderGetResponseProductBenefitsBenefitBaseType `json:"type,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time                                          `json:"modified_at,nullable" format:"date-time"`
	JSON       userOrderGetResponseProductBenefitsBenefitBaseJSON `json:"-"`
}

// userOrderGetResponseProductBenefitsBenefitBaseJSON contains the JSON metadata
// for the struct [UserOrderGetResponseProductBenefitsBenefitBase]
type userOrderGetResponseProductBenefitsBenefitBaseJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Deletable      apijson.Field
	Description    apijson.Field
	OrganizationID apijson.Field
	Selectable     apijson.Field
	Type           apijson.Field
	ModifiedAt     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *UserOrderGetResponseProductBenefitsBenefitBase) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userOrderGetResponseProductBenefitsBenefitBaseJSON) RawJSON() string {
	return r.raw
}

func (r UserOrderGetResponseProductBenefitsBenefitBase) implementsUserOrderGetResponseProductBenefit() {
}

// The type of the benefit.
type UserOrderGetResponseProductBenefitsBenefitBaseType string

const (
	UserOrderGetResponseProductBenefitsBenefitBaseTypeCustom           UserOrderGetResponseProductBenefitsBenefitBaseType = "custom"
	UserOrderGetResponseProductBenefitsBenefitBaseTypeArticles         UserOrderGetResponseProductBenefitsBenefitBaseType = "articles"
	UserOrderGetResponseProductBenefitsBenefitBaseTypeAds              UserOrderGetResponseProductBenefitsBenefitBaseType = "ads"
	UserOrderGetResponseProductBenefitsBenefitBaseTypeDiscord          UserOrderGetResponseProductBenefitsBenefitBaseType = "discord"
	UserOrderGetResponseProductBenefitsBenefitBaseTypeGitHubRepository UserOrderGetResponseProductBenefitsBenefitBaseType = "github_repository"
	UserOrderGetResponseProductBenefitsBenefitBaseTypeDownloadables    UserOrderGetResponseProductBenefitsBenefitBaseType = "downloadables"
)

func (r UserOrderGetResponseProductBenefitsBenefitBaseType) IsKnown() bool {
	switch r {
	case UserOrderGetResponseProductBenefitsBenefitBaseTypeCustom, UserOrderGetResponseProductBenefitsBenefitBaseTypeArticles, UserOrderGetResponseProductBenefitsBenefitBaseTypeAds, UserOrderGetResponseProductBenefitsBenefitBaseTypeDiscord, UserOrderGetResponseProductBenefitsBenefitBaseTypeGitHubRepository, UserOrderGetResponseProductBenefitsBenefitBaseTypeDownloadables:
		return true
	}
	return false
}

// A benefit of type `articles`.
//
// Use it to grant access to posts.
type UserOrderGetResponseProductBenefitsBenefitArticles struct {
	// The ID of the benefit.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the benefit is deletable.
	Deletable bool `json:"deletable,required"`
	// The description of the benefit.
	Description string `json:"description,required"`
	// The ID of the organization owning the benefit.
	OrganizationID string `json:"organization_id,required" format:"uuid4"`
	// Properties for a benefit of type `articles`.
	Properties UserOrderGetResponseProductBenefitsBenefitArticlesProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                                   `json:"selectable,required"`
	Type       UserOrderGetResponseProductBenefitsBenefitArticlesType `json:"type,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time                                              `json:"modified_at,nullable" format:"date-time"`
	JSON       userOrderGetResponseProductBenefitsBenefitArticlesJSON `json:"-"`
}

// userOrderGetResponseProductBenefitsBenefitArticlesJSON contains the JSON
// metadata for the struct [UserOrderGetResponseProductBenefitsBenefitArticles]
type userOrderGetResponseProductBenefitsBenefitArticlesJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Deletable      apijson.Field
	Description    apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	Selectable     apijson.Field
	Type           apijson.Field
	ModifiedAt     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *UserOrderGetResponseProductBenefitsBenefitArticles) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userOrderGetResponseProductBenefitsBenefitArticlesJSON) RawJSON() string {
	return r.raw
}

func (r UserOrderGetResponseProductBenefitsBenefitArticles) implementsUserOrderGetResponseProductBenefit() {
}

// Properties for a benefit of type `articles`.
type UserOrderGetResponseProductBenefitsBenefitArticlesProperties struct {
	// Whether the user can access paid articles.
	PaidArticles bool                                                             `json:"paid_articles,required"`
	JSON         userOrderGetResponseProductBenefitsBenefitArticlesPropertiesJSON `json:"-"`
}

// userOrderGetResponseProductBenefitsBenefitArticlesPropertiesJSON contains the
// JSON metadata for the struct
// [UserOrderGetResponseProductBenefitsBenefitArticlesProperties]
type userOrderGetResponseProductBenefitsBenefitArticlesPropertiesJSON struct {
	PaidArticles apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *UserOrderGetResponseProductBenefitsBenefitArticlesProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userOrderGetResponseProductBenefitsBenefitArticlesPropertiesJSON) RawJSON() string {
	return r.raw
}

type UserOrderGetResponseProductBenefitsBenefitArticlesType string

const (
	UserOrderGetResponseProductBenefitsBenefitArticlesTypeArticles UserOrderGetResponseProductBenefitsBenefitArticlesType = "articles"
)

func (r UserOrderGetResponseProductBenefitsBenefitArticlesType) IsKnown() bool {
	switch r {
	case UserOrderGetResponseProductBenefitsBenefitArticlesTypeArticles:
		return true
	}
	return false
}

// The type of the benefit.
type UserOrderGetResponseProductBenefitsType string

const (
	UserOrderGetResponseProductBenefitsTypeCustom           UserOrderGetResponseProductBenefitsType = "custom"
	UserOrderGetResponseProductBenefitsTypeArticles         UserOrderGetResponseProductBenefitsType = "articles"
	UserOrderGetResponseProductBenefitsTypeAds              UserOrderGetResponseProductBenefitsType = "ads"
	UserOrderGetResponseProductBenefitsTypeDiscord          UserOrderGetResponseProductBenefitsType = "discord"
	UserOrderGetResponseProductBenefitsTypeGitHubRepository UserOrderGetResponseProductBenefitsType = "github_repository"
	UserOrderGetResponseProductBenefitsTypeDownloadables    UserOrderGetResponseProductBenefitsType = "downloadables"
)

func (r UserOrderGetResponseProductBenefitsType) IsKnown() bool {
	switch r {
	case UserOrderGetResponseProductBenefitsTypeCustom, UserOrderGetResponseProductBenefitsTypeArticles, UserOrderGetResponseProductBenefitsTypeAds, UserOrderGetResponseProductBenefitsTypeDiscord, UserOrderGetResponseProductBenefitsTypeGitHubRepository, UserOrderGetResponseProductBenefitsTypeDownloadables:
		return true
	}
	return false
}

// A recurring price for a product, i.e. a subscription.
type UserOrderGetResponseProductPrice struct {
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,nullable" format:"date-time"`
	// The ID of the price.
	ID string `json:"id,required" format:"uuid4"`
	// The price in cents.
	PriceAmount int64 `json:"price_amount,required"`
	// The currency.
	PriceCurrency string `json:"price_currency,required"`
	// Whether the price is archived and no longer available.
	IsArchived bool `json:"is_archived,required"`
	// The type of the price.
	Type UserOrderGetResponseProductPricesType `json:"type,required"`
	// The recurring interval of the price, if type is `recurring`.
	RecurringInterval UserOrderGetResponseProductPricesRecurringInterval `json:"recurring_interval,nullable"`
	JSON              userOrderGetResponseProductPriceJSON               `json:"-"`
	union             UserOrderGetResponseProductPricesUnion
}

// userOrderGetResponseProductPriceJSON contains the JSON metadata for the struct
// [UserOrderGetResponseProductPrice]
type userOrderGetResponseProductPriceJSON struct {
	CreatedAt         apijson.Field
	ModifiedAt        apijson.Field
	ID                apijson.Field
	PriceAmount       apijson.Field
	PriceCurrency     apijson.Field
	IsArchived        apijson.Field
	Type              apijson.Field
	RecurringInterval apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r userOrderGetResponseProductPriceJSON) RawJSON() string {
	return r.raw
}

func (r *UserOrderGetResponseProductPrice) UnmarshalJSON(data []byte) (err error) {
	*r = UserOrderGetResponseProductPrice{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [UserOrderGetResponseProductPricesUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [UserOrderGetResponseProductPricesProductPriceRecurring],
// [UserOrderGetResponseProductPricesProductPriceOneTime].
func (r UserOrderGetResponseProductPrice) AsUnion() UserOrderGetResponseProductPricesUnion {
	return r.union
}

// A recurring price for a product, i.e. a subscription.
//
// Union satisfied by [UserOrderGetResponseProductPricesProductPriceRecurring] or
// [UserOrderGetResponseProductPricesProductPriceOneTime].
type UserOrderGetResponseProductPricesUnion interface {
	implementsUserOrderGetResponseProductPrice()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserOrderGetResponseProductPricesUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(UserOrderGetResponseProductPricesProductPriceRecurring{}),
			DiscriminatorValue: "recurring",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(UserOrderGetResponseProductPricesProductPriceOneTime{}),
			DiscriminatorValue: "one_time",
		},
	)
}

// A recurring price for a product, i.e. a subscription.
type UserOrderGetResponseProductPricesProductPriceRecurring struct {
	// The ID of the price.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the price is archived and no longer available.
	IsArchived bool `json:"is_archived,required"`
	// The price in cents.
	PriceAmount int64 `json:"price_amount,required"`
	// The currency.
	PriceCurrency string `json:"price_currency,required"`
	// The type of the price.
	Type UserOrderGetResponseProductPricesProductPriceRecurringType `json:"type,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,nullable" format:"date-time"`
	// The recurring interval of the price, if type is `recurring`.
	RecurringInterval UserOrderGetResponseProductPricesProductPriceRecurringRecurringInterval `json:"recurring_interval,nullable"`
	JSON              userOrderGetResponseProductPricesProductPriceRecurringJSON              `json:"-"`
}

// userOrderGetResponseProductPricesProductPriceRecurringJSON contains the JSON
// metadata for the struct [UserOrderGetResponseProductPricesProductPriceRecurring]
type userOrderGetResponseProductPricesProductPriceRecurringJSON struct {
	ID                apijson.Field
	CreatedAt         apijson.Field
	IsArchived        apijson.Field
	PriceAmount       apijson.Field
	PriceCurrency     apijson.Field
	Type              apijson.Field
	ModifiedAt        apijson.Field
	RecurringInterval apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *UserOrderGetResponseProductPricesProductPriceRecurring) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userOrderGetResponseProductPricesProductPriceRecurringJSON) RawJSON() string {
	return r.raw
}

func (r UserOrderGetResponseProductPricesProductPriceRecurring) implementsUserOrderGetResponseProductPrice() {
}

// The type of the price.
type UserOrderGetResponseProductPricesProductPriceRecurringType string

const (
	UserOrderGetResponseProductPricesProductPriceRecurringTypeRecurring UserOrderGetResponseProductPricesProductPriceRecurringType = "recurring"
)

func (r UserOrderGetResponseProductPricesProductPriceRecurringType) IsKnown() bool {
	switch r {
	case UserOrderGetResponseProductPricesProductPriceRecurringTypeRecurring:
		return true
	}
	return false
}

// The recurring interval of the price, if type is `recurring`.
type UserOrderGetResponseProductPricesProductPriceRecurringRecurringInterval string

const (
	UserOrderGetResponseProductPricesProductPriceRecurringRecurringIntervalMonth UserOrderGetResponseProductPricesProductPriceRecurringRecurringInterval = "month"
	UserOrderGetResponseProductPricesProductPriceRecurringRecurringIntervalYear  UserOrderGetResponseProductPricesProductPriceRecurringRecurringInterval = "year"
)

func (r UserOrderGetResponseProductPricesProductPriceRecurringRecurringInterval) IsKnown() bool {
	switch r {
	case UserOrderGetResponseProductPricesProductPriceRecurringRecurringIntervalMonth, UserOrderGetResponseProductPricesProductPriceRecurringRecurringIntervalYear:
		return true
	}
	return false
}

// A one-time price for a product.
type UserOrderGetResponseProductPricesProductPriceOneTime struct {
	// The ID of the price.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the price is archived and no longer available.
	IsArchived bool `json:"is_archived,required"`
	// The price in cents.
	PriceAmount int64 `json:"price_amount,required"`
	// The currency.
	PriceCurrency string `json:"price_currency,required"`
	// The type of the price.
	Type UserOrderGetResponseProductPricesProductPriceOneTimeType `json:"type,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time                                                `json:"modified_at,nullable" format:"date-time"`
	JSON       userOrderGetResponseProductPricesProductPriceOneTimeJSON `json:"-"`
}

// userOrderGetResponseProductPricesProductPriceOneTimeJSON contains the JSON
// metadata for the struct [UserOrderGetResponseProductPricesProductPriceOneTime]
type userOrderGetResponseProductPricesProductPriceOneTimeJSON struct {
	ID            apijson.Field
	CreatedAt     apijson.Field
	IsArchived    apijson.Field
	PriceAmount   apijson.Field
	PriceCurrency apijson.Field
	Type          apijson.Field
	ModifiedAt    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *UserOrderGetResponseProductPricesProductPriceOneTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userOrderGetResponseProductPricesProductPriceOneTimeJSON) RawJSON() string {
	return r.raw
}

func (r UserOrderGetResponseProductPricesProductPriceOneTime) implementsUserOrderGetResponseProductPrice() {
}

// The type of the price.
type UserOrderGetResponseProductPricesProductPriceOneTimeType string

const (
	UserOrderGetResponseProductPricesProductPriceOneTimeTypeOneTime UserOrderGetResponseProductPricesProductPriceOneTimeType = "one_time"
)

func (r UserOrderGetResponseProductPricesProductPriceOneTimeType) IsKnown() bool {
	switch r {
	case UserOrderGetResponseProductPricesProductPriceOneTimeTypeOneTime:
		return true
	}
	return false
}

// The type of the price.
type UserOrderGetResponseProductPricesType string

const (
	UserOrderGetResponseProductPricesTypeRecurring UserOrderGetResponseProductPricesType = "recurring"
	UserOrderGetResponseProductPricesTypeOneTime   UserOrderGetResponseProductPricesType = "one_time"
)

func (r UserOrderGetResponseProductPricesType) IsKnown() bool {
	switch r {
	case UserOrderGetResponseProductPricesTypeRecurring, UserOrderGetResponseProductPricesTypeOneTime:
		return true
	}
	return false
}

// The recurring interval of the price, if type is `recurring`.
type UserOrderGetResponseProductPricesRecurringInterval string

const (
	UserOrderGetResponseProductPricesRecurringIntervalMonth UserOrderGetResponseProductPricesRecurringInterval = "month"
	UserOrderGetResponseProductPricesRecurringIntervalYear  UserOrderGetResponseProductPricesRecurringInterval = "year"
)

func (r UserOrderGetResponseProductPricesRecurringInterval) IsKnown() bool {
	switch r {
	case UserOrderGetResponseProductPricesRecurringIntervalMonth, UserOrderGetResponseProductPricesRecurringIntervalYear:
		return true
	}
	return false
}

type UserOrderGetResponseProductType string

const (
	UserOrderGetResponseProductTypeFree       UserOrderGetResponseProductType = "free"
	UserOrderGetResponseProductTypeIndividual UserOrderGetResponseProductType = "individual"
	UserOrderGetResponseProductTypeBusiness   UserOrderGetResponseProductType = "business"
)

func (r UserOrderGetResponseProductType) IsKnown() bool {
	switch r {
	case UserOrderGetResponseProductTypeFree, UserOrderGetResponseProductTypeIndividual, UserOrderGetResponseProductTypeBusiness:
		return true
	}
	return false
}

type UserOrderGetResponseSubscription struct {
	// The ID of the object.
	ID                string `json:"id,required" format:"uuid4"`
	CancelAtPeriodEnd bool   `json:"cancel_at_period_end,required"`
	// Creation timestamp of the object.
	CreatedAt          time.Time                              `json:"created_at,required" format:"date-time"`
	CurrentPeriodStart time.Time                              `json:"current_period_start,required" format:"date-time"`
	ProductID          string                                 `json:"product_id,required" format:"uuid4"`
	Status             UserOrderGetResponseSubscriptionStatus `json:"status,required"`
	UserID             string                                 `json:"user_id,required" format:"uuid4"`
	CurrentPeriodEnd   time.Time                              `json:"current_period_end,nullable" format:"date-time"`
	EndedAt            time.Time                              `json:"ended_at,nullable" format:"date-time"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time                            `json:"modified_at,nullable" format:"date-time"`
	PriceID    string                               `json:"price_id,nullable" format:"uuid4"`
	StartedAt  time.Time                            `json:"started_at,nullable" format:"date-time"`
	JSON       userOrderGetResponseSubscriptionJSON `json:"-"`
}

// userOrderGetResponseSubscriptionJSON contains the JSON metadata for the struct
// [UserOrderGetResponseSubscription]
type userOrderGetResponseSubscriptionJSON struct {
	ID                 apijson.Field
	CancelAtPeriodEnd  apijson.Field
	CreatedAt          apijson.Field
	CurrentPeriodStart apijson.Field
	ProductID          apijson.Field
	Status             apijson.Field
	UserID             apijson.Field
	CurrentPeriodEnd   apijson.Field
	EndedAt            apijson.Field
	ModifiedAt         apijson.Field
	PriceID            apijson.Field
	StartedAt          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *UserOrderGetResponseSubscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userOrderGetResponseSubscriptionJSON) RawJSON() string {
	return r.raw
}

type UserOrderGetResponseSubscriptionStatus string

const (
	UserOrderGetResponseSubscriptionStatusIncomplete        UserOrderGetResponseSubscriptionStatus = "incomplete"
	UserOrderGetResponseSubscriptionStatusIncompleteExpired UserOrderGetResponseSubscriptionStatus = "incomplete_expired"
	UserOrderGetResponseSubscriptionStatusTrialing          UserOrderGetResponseSubscriptionStatus = "trialing"
	UserOrderGetResponseSubscriptionStatusActive            UserOrderGetResponseSubscriptionStatus = "active"
	UserOrderGetResponseSubscriptionStatusPastDue           UserOrderGetResponseSubscriptionStatus = "past_due"
	UserOrderGetResponseSubscriptionStatusCanceled          UserOrderGetResponseSubscriptionStatus = "canceled"
	UserOrderGetResponseSubscriptionStatusUnpaid            UserOrderGetResponseSubscriptionStatus = "unpaid"
)

func (r UserOrderGetResponseSubscriptionStatus) IsKnown() bool {
	switch r {
	case UserOrderGetResponseSubscriptionStatusIncomplete, UserOrderGetResponseSubscriptionStatusIncompleteExpired, UserOrderGetResponseSubscriptionStatusTrialing, UserOrderGetResponseSubscriptionStatusActive, UserOrderGetResponseSubscriptionStatusPastDue, UserOrderGetResponseSubscriptionStatusCanceled, UserOrderGetResponseSubscriptionStatusUnpaid:
		return true
	}
	return false
}

type UserOrderListResponse struct {
	Pagination UserOrderListResponsePagination `json:"pagination,required"`
	Items      []UserOrderListResponseItem     `json:"items"`
	JSON       userOrderListResponseJSON       `json:"-"`
}

// userOrderListResponseJSON contains the JSON metadata for the struct
// [UserOrderListResponse]
type userOrderListResponseJSON struct {
	Pagination  apijson.Field
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserOrderListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userOrderListResponseJSON) RawJSON() string {
	return r.raw
}

type UserOrderListResponsePagination struct {
	MaxPage    int64                               `json:"max_page,required"`
	TotalCount int64                               `json:"total_count,required"`
	JSON       userOrderListResponsePaginationJSON `json:"-"`
}

// userOrderListResponsePaginationJSON contains the JSON metadata for the struct
// [UserOrderListResponsePagination]
type userOrderListResponsePaginationJSON struct {
	MaxPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserOrderListResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userOrderListResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type UserOrderListResponseItem struct {
	ID     string `json:"id,required" format:"uuid4"`
	Amount int64  `json:"amount,required"`
	// Creation timestamp of the object.
	CreatedAt time.Time                         `json:"created_at,required" format:"date-time"`
	Currency  string                            `json:"currency,required"`
	Product   UserOrderListResponseItemsProduct `json:"product,required"`
	ProductID string                            `json:"product_id,required" format:"uuid4"`
	// A recurring price for a product, i.e. a subscription.
	ProductPrice   UserOrderListResponseItemsProductPrice `json:"product_price,required"`
	ProductPriceID string                                 `json:"product_price_id,required" format:"uuid4"`
	TaxAmount      int64                                  `json:"tax_amount,required"`
	UserID         string                                 `json:"user_id,required" format:"uuid4"`
	// Last modification timestamp of the object.
	ModifiedAt     time.Time                              `json:"modified_at,nullable" format:"date-time"`
	Subscription   UserOrderListResponseItemsSubscription `json:"subscription,nullable"`
	SubscriptionID string                                 `json:"subscription_id,nullable" format:"uuid4"`
	JSON           userOrderListResponseItemJSON          `json:"-"`
}

// userOrderListResponseItemJSON contains the JSON metadata for the struct
// [UserOrderListResponseItem]
type userOrderListResponseItemJSON struct {
	ID             apijson.Field
	Amount         apijson.Field
	CreatedAt      apijson.Field
	Currency       apijson.Field
	Product        apijson.Field
	ProductID      apijson.Field
	ProductPrice   apijson.Field
	ProductPriceID apijson.Field
	TaxAmount      apijson.Field
	UserID         apijson.Field
	ModifiedAt     apijson.Field
	Subscription   apijson.Field
	SubscriptionID apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *UserOrderListResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userOrderListResponseItemJSON) RawJSON() string {
	return r.raw
}

type UserOrderListResponseItemsProduct struct {
	// The ID of the product.
	ID string `json:"id,required" format:"uuid4"`
	// The benefits granted by the product.
	Benefits []UserOrderListResponseItemsProductBenefit `json:"benefits,required"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the product is archived and no longer available.
	IsArchived bool `json:"is_archived,required"`
	// Whether the product is a subscription tier.
	IsRecurring bool `json:"is_recurring,required"`
	// The medias associated to the product.
	Medias []ProductMediaFileReadOutput `json:"medias,required"`
	// The name of the product.
	Name string `json:"name,required"`
	// The ID of the organization owning the product.
	OrganizationID string `json:"organization_id,required" format:"uuid4"`
	// List of available prices for this product.
	Prices []UserOrderListResponseItemsProductPrice `json:"prices,required"`
	// The description of the product.
	Description   string `json:"description,nullable"`
	IsHighlighted bool   `json:"is_highlighted,nullable"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time                             `json:"modified_at,nullable" format:"date-time"`
	Type       UserOrderListResponseItemsProductType `json:"type,nullable"`
	JSON       userOrderListResponseItemsProductJSON `json:"-"`
}

// userOrderListResponseItemsProductJSON contains the JSON metadata for the struct
// [UserOrderListResponseItemsProduct]
type userOrderListResponseItemsProductJSON struct {
	ID             apijson.Field
	Benefits       apijson.Field
	CreatedAt      apijson.Field
	IsArchived     apijson.Field
	IsRecurring    apijson.Field
	Medias         apijson.Field
	Name           apijson.Field
	OrganizationID apijson.Field
	Prices         apijson.Field
	Description    apijson.Field
	IsHighlighted  apijson.Field
	ModifiedAt     apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *UserOrderListResponseItemsProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userOrderListResponseItemsProductJSON) RawJSON() string {
	return r.raw
}

// A benefit of type `articles`.
//
// Use it to grant access to posts.
type UserOrderListResponseItemsProductBenefit struct {
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,nullable" format:"date-time"`
	// The ID of the benefit.
	ID string `json:"id,required" format:"uuid4"`
	// The type of the benefit.
	Type UserOrderListResponseItemsProductBenefitsType `json:"type,required"`
	// The description of the benefit.
	Description string `json:"description,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool `json:"selectable,required"`
	// Whether the benefit is deletable.
	Deletable bool `json:"deletable,required"`
	// The ID of the organization owning the benefit.
	OrganizationID string `json:"organization_id,required" format:"uuid4"`
	// This field can have the runtime type of
	// [UserOrderListResponseItemsProductBenefitsBenefitArticlesProperties].
	Properties interface{}                                  `json:"properties,required"`
	JSON       userOrderListResponseItemsProductBenefitJSON `json:"-"`
	union      UserOrderListResponseItemsProductBenefitsUnion
}

// userOrderListResponseItemsProductBenefitJSON contains the JSON metadata for the
// struct [UserOrderListResponseItemsProductBenefit]
type userOrderListResponseItemsProductBenefitJSON struct {
	CreatedAt      apijson.Field
	ModifiedAt     apijson.Field
	ID             apijson.Field
	Type           apijson.Field
	Description    apijson.Field
	Selectable     apijson.Field
	Deletable      apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r userOrderListResponseItemsProductBenefitJSON) RawJSON() string {
	return r.raw
}

func (r *UserOrderListResponseItemsProductBenefit) UnmarshalJSON(data []byte) (err error) {
	*r = UserOrderListResponseItemsProductBenefit{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [UserOrderListResponseItemsProductBenefitsUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [UserOrderListResponseItemsProductBenefitsBenefitBase],
// [UserOrderListResponseItemsProductBenefitsBenefitArticles].
func (r UserOrderListResponseItemsProductBenefit) AsUnion() UserOrderListResponseItemsProductBenefitsUnion {
	return r.union
}

// A benefit of type `articles`.
//
// Use it to grant access to posts.
//
// Union satisfied by [UserOrderListResponseItemsProductBenefitsBenefitBase] or
// [UserOrderListResponseItemsProductBenefitsBenefitArticles].
type UserOrderListResponseItemsProductBenefitsUnion interface {
	implementsUserOrderListResponseItemsProductBenefit()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserOrderListResponseItemsProductBenefitsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserOrderListResponseItemsProductBenefitsBenefitBase{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserOrderListResponseItemsProductBenefitsBenefitArticles{}),
		},
	)
}

type UserOrderListResponseItemsProductBenefitsBenefitBase struct {
	// The ID of the benefit.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the benefit is deletable.
	Deletable bool `json:"deletable,required"`
	// The description of the benefit.
	Description string `json:"description,required"`
	// The ID of the organization owning the benefit.
	OrganizationID string `json:"organization_id,required" format:"uuid4"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool `json:"selectable,required"`
	// The type of the benefit.
	Type UserOrderListResponseItemsProductBenefitsBenefitBaseType `json:"type,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time                                                `json:"modified_at,nullable" format:"date-time"`
	JSON       userOrderListResponseItemsProductBenefitsBenefitBaseJSON `json:"-"`
}

// userOrderListResponseItemsProductBenefitsBenefitBaseJSON contains the JSON
// metadata for the struct [UserOrderListResponseItemsProductBenefitsBenefitBase]
type userOrderListResponseItemsProductBenefitsBenefitBaseJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Deletable      apijson.Field
	Description    apijson.Field
	OrganizationID apijson.Field
	Selectable     apijson.Field
	Type           apijson.Field
	ModifiedAt     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *UserOrderListResponseItemsProductBenefitsBenefitBase) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userOrderListResponseItemsProductBenefitsBenefitBaseJSON) RawJSON() string {
	return r.raw
}

func (r UserOrderListResponseItemsProductBenefitsBenefitBase) implementsUserOrderListResponseItemsProductBenefit() {
}

// The type of the benefit.
type UserOrderListResponseItemsProductBenefitsBenefitBaseType string

const (
	UserOrderListResponseItemsProductBenefitsBenefitBaseTypeCustom           UserOrderListResponseItemsProductBenefitsBenefitBaseType = "custom"
	UserOrderListResponseItemsProductBenefitsBenefitBaseTypeArticles         UserOrderListResponseItemsProductBenefitsBenefitBaseType = "articles"
	UserOrderListResponseItemsProductBenefitsBenefitBaseTypeAds              UserOrderListResponseItemsProductBenefitsBenefitBaseType = "ads"
	UserOrderListResponseItemsProductBenefitsBenefitBaseTypeDiscord          UserOrderListResponseItemsProductBenefitsBenefitBaseType = "discord"
	UserOrderListResponseItemsProductBenefitsBenefitBaseTypeGitHubRepository UserOrderListResponseItemsProductBenefitsBenefitBaseType = "github_repository"
	UserOrderListResponseItemsProductBenefitsBenefitBaseTypeDownloadables    UserOrderListResponseItemsProductBenefitsBenefitBaseType = "downloadables"
)

func (r UserOrderListResponseItemsProductBenefitsBenefitBaseType) IsKnown() bool {
	switch r {
	case UserOrderListResponseItemsProductBenefitsBenefitBaseTypeCustom, UserOrderListResponseItemsProductBenefitsBenefitBaseTypeArticles, UserOrderListResponseItemsProductBenefitsBenefitBaseTypeAds, UserOrderListResponseItemsProductBenefitsBenefitBaseTypeDiscord, UserOrderListResponseItemsProductBenefitsBenefitBaseTypeGitHubRepository, UserOrderListResponseItemsProductBenefitsBenefitBaseTypeDownloadables:
		return true
	}
	return false
}

// A benefit of type `articles`.
//
// Use it to grant access to posts.
type UserOrderListResponseItemsProductBenefitsBenefitArticles struct {
	// The ID of the benefit.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the benefit is deletable.
	Deletable bool `json:"deletable,required"`
	// The description of the benefit.
	Description string `json:"description,required"`
	// The ID of the organization owning the benefit.
	OrganizationID string `json:"organization_id,required" format:"uuid4"`
	// Properties for a benefit of type `articles`.
	Properties UserOrderListResponseItemsProductBenefitsBenefitArticlesProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                                         `json:"selectable,required"`
	Type       UserOrderListResponseItemsProductBenefitsBenefitArticlesType `json:"type,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time                                                    `json:"modified_at,nullable" format:"date-time"`
	JSON       userOrderListResponseItemsProductBenefitsBenefitArticlesJSON `json:"-"`
}

// userOrderListResponseItemsProductBenefitsBenefitArticlesJSON contains the JSON
// metadata for the struct
// [UserOrderListResponseItemsProductBenefitsBenefitArticles]
type userOrderListResponseItemsProductBenefitsBenefitArticlesJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Deletable      apijson.Field
	Description    apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	Selectable     apijson.Field
	Type           apijson.Field
	ModifiedAt     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *UserOrderListResponseItemsProductBenefitsBenefitArticles) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userOrderListResponseItemsProductBenefitsBenefitArticlesJSON) RawJSON() string {
	return r.raw
}

func (r UserOrderListResponseItemsProductBenefitsBenefitArticles) implementsUserOrderListResponseItemsProductBenefit() {
}

// Properties for a benefit of type `articles`.
type UserOrderListResponseItemsProductBenefitsBenefitArticlesProperties struct {
	// Whether the user can access paid articles.
	PaidArticles bool                                                                   `json:"paid_articles,required"`
	JSON         userOrderListResponseItemsProductBenefitsBenefitArticlesPropertiesJSON `json:"-"`
}

// userOrderListResponseItemsProductBenefitsBenefitArticlesPropertiesJSON contains
// the JSON metadata for the struct
// [UserOrderListResponseItemsProductBenefitsBenefitArticlesProperties]
type userOrderListResponseItemsProductBenefitsBenefitArticlesPropertiesJSON struct {
	PaidArticles apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *UserOrderListResponseItemsProductBenefitsBenefitArticlesProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userOrderListResponseItemsProductBenefitsBenefitArticlesPropertiesJSON) RawJSON() string {
	return r.raw
}

type UserOrderListResponseItemsProductBenefitsBenefitArticlesType string

const (
	UserOrderListResponseItemsProductBenefitsBenefitArticlesTypeArticles UserOrderListResponseItemsProductBenefitsBenefitArticlesType = "articles"
)

func (r UserOrderListResponseItemsProductBenefitsBenefitArticlesType) IsKnown() bool {
	switch r {
	case UserOrderListResponseItemsProductBenefitsBenefitArticlesTypeArticles:
		return true
	}
	return false
}

// The type of the benefit.
type UserOrderListResponseItemsProductBenefitsType string

const (
	UserOrderListResponseItemsProductBenefitsTypeCustom           UserOrderListResponseItemsProductBenefitsType = "custom"
	UserOrderListResponseItemsProductBenefitsTypeArticles         UserOrderListResponseItemsProductBenefitsType = "articles"
	UserOrderListResponseItemsProductBenefitsTypeAds              UserOrderListResponseItemsProductBenefitsType = "ads"
	UserOrderListResponseItemsProductBenefitsTypeDiscord          UserOrderListResponseItemsProductBenefitsType = "discord"
	UserOrderListResponseItemsProductBenefitsTypeGitHubRepository UserOrderListResponseItemsProductBenefitsType = "github_repository"
	UserOrderListResponseItemsProductBenefitsTypeDownloadables    UserOrderListResponseItemsProductBenefitsType = "downloadables"
)

func (r UserOrderListResponseItemsProductBenefitsType) IsKnown() bool {
	switch r {
	case UserOrderListResponseItemsProductBenefitsTypeCustom, UserOrderListResponseItemsProductBenefitsTypeArticles, UserOrderListResponseItemsProductBenefitsTypeAds, UserOrderListResponseItemsProductBenefitsTypeDiscord, UserOrderListResponseItemsProductBenefitsTypeGitHubRepository, UserOrderListResponseItemsProductBenefitsTypeDownloadables:
		return true
	}
	return false
}

// A recurring price for a product, i.e. a subscription.
type UserOrderListResponseItemsProductPrice struct {
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,nullable" format:"date-time"`
	// The ID of the price.
	ID string `json:"id,required" format:"uuid4"`
	// The price in cents.
	PriceAmount int64 `json:"price_amount,required"`
	// The currency.
	PriceCurrency string `json:"price_currency,required"`
	// Whether the price is archived and no longer available.
	IsArchived bool `json:"is_archived,required"`
	// The type of the price.
	Type UserOrderListResponseItemsProductPricesType `json:"type,required"`
	// The recurring interval of the price, if type is `recurring`.
	RecurringInterval UserOrderListResponseItemsProductPricesRecurringInterval `json:"recurring_interval,nullable"`
	JSON              userOrderListResponseItemsProductPriceJSON               `json:"-"`
	union             UserOrderListResponseItemsProductPricesUnion
}

// userOrderListResponseItemsProductPriceJSON contains the JSON metadata for the
// struct [UserOrderListResponseItemsProductPrice]
type userOrderListResponseItemsProductPriceJSON struct {
	CreatedAt         apijson.Field
	ModifiedAt        apijson.Field
	ID                apijson.Field
	PriceAmount       apijson.Field
	PriceCurrency     apijson.Field
	IsArchived        apijson.Field
	Type              apijson.Field
	RecurringInterval apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r userOrderListResponseItemsProductPriceJSON) RawJSON() string {
	return r.raw
}

func (r *UserOrderListResponseItemsProductPrice) UnmarshalJSON(data []byte) (err error) {
	*r = UserOrderListResponseItemsProductPrice{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [UserOrderListResponseItemsProductPricesUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [UserOrderListResponseItemsProductPricesProductPriceRecurring],
// [UserOrderListResponseItemsProductPricesProductPriceOneTime].
func (r UserOrderListResponseItemsProductPrice) AsUnion() UserOrderListResponseItemsProductPricesUnion {
	return r.union
}

// A recurring price for a product, i.e. a subscription.
//
// Union satisfied by
// [UserOrderListResponseItemsProductPricesProductPriceRecurring] or
// [UserOrderListResponseItemsProductPricesProductPriceOneTime].
type UserOrderListResponseItemsProductPricesUnion interface {
	implementsUserOrderListResponseItemsProductPrice()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserOrderListResponseItemsProductPricesUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(UserOrderListResponseItemsProductPricesProductPriceRecurring{}),
			DiscriminatorValue: "recurring",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(UserOrderListResponseItemsProductPricesProductPriceOneTime{}),
			DiscriminatorValue: "one_time",
		},
	)
}

// A recurring price for a product, i.e. a subscription.
type UserOrderListResponseItemsProductPricesProductPriceRecurring struct {
	// The ID of the price.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the price is archived and no longer available.
	IsArchived bool `json:"is_archived,required"`
	// The price in cents.
	PriceAmount int64 `json:"price_amount,required"`
	// The currency.
	PriceCurrency string `json:"price_currency,required"`
	// The type of the price.
	Type UserOrderListResponseItemsProductPricesProductPriceRecurringType `json:"type,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,nullable" format:"date-time"`
	// The recurring interval of the price, if type is `recurring`.
	RecurringInterval UserOrderListResponseItemsProductPricesProductPriceRecurringRecurringInterval `json:"recurring_interval,nullable"`
	JSON              userOrderListResponseItemsProductPricesProductPriceRecurringJSON              `json:"-"`
}

// userOrderListResponseItemsProductPricesProductPriceRecurringJSON contains the
// JSON metadata for the struct
// [UserOrderListResponseItemsProductPricesProductPriceRecurring]
type userOrderListResponseItemsProductPricesProductPriceRecurringJSON struct {
	ID                apijson.Field
	CreatedAt         apijson.Field
	IsArchived        apijson.Field
	PriceAmount       apijson.Field
	PriceCurrency     apijson.Field
	Type              apijson.Field
	ModifiedAt        apijson.Field
	RecurringInterval apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *UserOrderListResponseItemsProductPricesProductPriceRecurring) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userOrderListResponseItemsProductPricesProductPriceRecurringJSON) RawJSON() string {
	return r.raw
}

func (r UserOrderListResponseItemsProductPricesProductPriceRecurring) implementsUserOrderListResponseItemsProductPrice() {
}

// The type of the price.
type UserOrderListResponseItemsProductPricesProductPriceRecurringType string

const (
	UserOrderListResponseItemsProductPricesProductPriceRecurringTypeRecurring UserOrderListResponseItemsProductPricesProductPriceRecurringType = "recurring"
)

func (r UserOrderListResponseItemsProductPricesProductPriceRecurringType) IsKnown() bool {
	switch r {
	case UserOrderListResponseItemsProductPricesProductPriceRecurringTypeRecurring:
		return true
	}
	return false
}

// The recurring interval of the price, if type is `recurring`.
type UserOrderListResponseItemsProductPricesProductPriceRecurringRecurringInterval string

const (
	UserOrderListResponseItemsProductPricesProductPriceRecurringRecurringIntervalMonth UserOrderListResponseItemsProductPricesProductPriceRecurringRecurringInterval = "month"
	UserOrderListResponseItemsProductPricesProductPriceRecurringRecurringIntervalYear  UserOrderListResponseItemsProductPricesProductPriceRecurringRecurringInterval = "year"
)

func (r UserOrderListResponseItemsProductPricesProductPriceRecurringRecurringInterval) IsKnown() bool {
	switch r {
	case UserOrderListResponseItemsProductPricesProductPriceRecurringRecurringIntervalMonth, UserOrderListResponseItemsProductPricesProductPriceRecurringRecurringIntervalYear:
		return true
	}
	return false
}

// A one-time price for a product.
type UserOrderListResponseItemsProductPricesProductPriceOneTime struct {
	// The ID of the price.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the price is archived and no longer available.
	IsArchived bool `json:"is_archived,required"`
	// The price in cents.
	PriceAmount int64 `json:"price_amount,required"`
	// The currency.
	PriceCurrency string `json:"price_currency,required"`
	// The type of the price.
	Type UserOrderListResponseItemsProductPricesProductPriceOneTimeType `json:"type,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time                                                      `json:"modified_at,nullable" format:"date-time"`
	JSON       userOrderListResponseItemsProductPricesProductPriceOneTimeJSON `json:"-"`
}

// userOrderListResponseItemsProductPricesProductPriceOneTimeJSON contains the JSON
// metadata for the struct
// [UserOrderListResponseItemsProductPricesProductPriceOneTime]
type userOrderListResponseItemsProductPricesProductPriceOneTimeJSON struct {
	ID            apijson.Field
	CreatedAt     apijson.Field
	IsArchived    apijson.Field
	PriceAmount   apijson.Field
	PriceCurrency apijson.Field
	Type          apijson.Field
	ModifiedAt    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *UserOrderListResponseItemsProductPricesProductPriceOneTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userOrderListResponseItemsProductPricesProductPriceOneTimeJSON) RawJSON() string {
	return r.raw
}

func (r UserOrderListResponseItemsProductPricesProductPriceOneTime) implementsUserOrderListResponseItemsProductPrice() {
}

// The type of the price.
type UserOrderListResponseItemsProductPricesProductPriceOneTimeType string

const (
	UserOrderListResponseItemsProductPricesProductPriceOneTimeTypeOneTime UserOrderListResponseItemsProductPricesProductPriceOneTimeType = "one_time"
)

func (r UserOrderListResponseItemsProductPricesProductPriceOneTimeType) IsKnown() bool {
	switch r {
	case UserOrderListResponseItemsProductPricesProductPriceOneTimeTypeOneTime:
		return true
	}
	return false
}

// The type of the price.
type UserOrderListResponseItemsProductPricesType string

const (
	UserOrderListResponseItemsProductPricesTypeRecurring UserOrderListResponseItemsProductPricesType = "recurring"
	UserOrderListResponseItemsProductPricesTypeOneTime   UserOrderListResponseItemsProductPricesType = "one_time"
)

func (r UserOrderListResponseItemsProductPricesType) IsKnown() bool {
	switch r {
	case UserOrderListResponseItemsProductPricesTypeRecurring, UserOrderListResponseItemsProductPricesTypeOneTime:
		return true
	}
	return false
}

// The recurring interval of the price, if type is `recurring`.
type UserOrderListResponseItemsProductPricesRecurringInterval string

const (
	UserOrderListResponseItemsProductPricesRecurringIntervalMonth UserOrderListResponseItemsProductPricesRecurringInterval = "month"
	UserOrderListResponseItemsProductPricesRecurringIntervalYear  UserOrderListResponseItemsProductPricesRecurringInterval = "year"
)

func (r UserOrderListResponseItemsProductPricesRecurringInterval) IsKnown() bool {
	switch r {
	case UserOrderListResponseItemsProductPricesRecurringIntervalMonth, UserOrderListResponseItemsProductPricesRecurringIntervalYear:
		return true
	}
	return false
}

type UserOrderListResponseItemsProductType string

const (
	UserOrderListResponseItemsProductTypeFree       UserOrderListResponseItemsProductType = "free"
	UserOrderListResponseItemsProductTypeIndividual UserOrderListResponseItemsProductType = "individual"
	UserOrderListResponseItemsProductTypeBusiness   UserOrderListResponseItemsProductType = "business"
)

func (r UserOrderListResponseItemsProductType) IsKnown() bool {
	switch r {
	case UserOrderListResponseItemsProductTypeFree, UserOrderListResponseItemsProductTypeIndividual, UserOrderListResponseItemsProductTypeBusiness:
		return true
	}
	return false
}

type UserOrderListResponseItemsSubscription struct {
	// The ID of the object.
	ID                string `json:"id,required" format:"uuid4"`
	CancelAtPeriodEnd bool   `json:"cancel_at_period_end,required"`
	// Creation timestamp of the object.
	CreatedAt          time.Time                                    `json:"created_at,required" format:"date-time"`
	CurrentPeriodStart time.Time                                    `json:"current_period_start,required" format:"date-time"`
	ProductID          string                                       `json:"product_id,required" format:"uuid4"`
	Status             UserOrderListResponseItemsSubscriptionStatus `json:"status,required"`
	UserID             string                                       `json:"user_id,required" format:"uuid4"`
	CurrentPeriodEnd   time.Time                                    `json:"current_period_end,nullable" format:"date-time"`
	EndedAt            time.Time                                    `json:"ended_at,nullable" format:"date-time"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time                                  `json:"modified_at,nullable" format:"date-time"`
	PriceID    string                                     `json:"price_id,nullable" format:"uuid4"`
	StartedAt  time.Time                                  `json:"started_at,nullable" format:"date-time"`
	JSON       userOrderListResponseItemsSubscriptionJSON `json:"-"`
}

// userOrderListResponseItemsSubscriptionJSON contains the JSON metadata for the
// struct [UserOrderListResponseItemsSubscription]
type userOrderListResponseItemsSubscriptionJSON struct {
	ID                 apijson.Field
	CancelAtPeriodEnd  apijson.Field
	CreatedAt          apijson.Field
	CurrentPeriodStart apijson.Field
	ProductID          apijson.Field
	Status             apijson.Field
	UserID             apijson.Field
	CurrentPeriodEnd   apijson.Field
	EndedAt            apijson.Field
	ModifiedAt         apijson.Field
	PriceID            apijson.Field
	StartedAt          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *UserOrderListResponseItemsSubscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userOrderListResponseItemsSubscriptionJSON) RawJSON() string {
	return r.raw
}

type UserOrderListResponseItemsSubscriptionStatus string

const (
	UserOrderListResponseItemsSubscriptionStatusIncomplete        UserOrderListResponseItemsSubscriptionStatus = "incomplete"
	UserOrderListResponseItemsSubscriptionStatusIncompleteExpired UserOrderListResponseItemsSubscriptionStatus = "incomplete_expired"
	UserOrderListResponseItemsSubscriptionStatusTrialing          UserOrderListResponseItemsSubscriptionStatus = "trialing"
	UserOrderListResponseItemsSubscriptionStatusActive            UserOrderListResponseItemsSubscriptionStatus = "active"
	UserOrderListResponseItemsSubscriptionStatusPastDue           UserOrderListResponseItemsSubscriptionStatus = "past_due"
	UserOrderListResponseItemsSubscriptionStatusCanceled          UserOrderListResponseItemsSubscriptionStatus = "canceled"
	UserOrderListResponseItemsSubscriptionStatusUnpaid            UserOrderListResponseItemsSubscriptionStatus = "unpaid"
)

func (r UserOrderListResponseItemsSubscriptionStatus) IsKnown() bool {
	switch r {
	case UserOrderListResponseItemsSubscriptionStatusIncomplete, UserOrderListResponseItemsSubscriptionStatusIncompleteExpired, UserOrderListResponseItemsSubscriptionStatusTrialing, UserOrderListResponseItemsSubscriptionStatusActive, UserOrderListResponseItemsSubscriptionStatusPastDue, UserOrderListResponseItemsSubscriptionStatusCanceled, UserOrderListResponseItemsSubscriptionStatusUnpaid:
		return true
	}
	return false
}

// Order's invoice data.
type UserOrderInvoiceResponse struct {
	// The URL to the invoice.
	URL  string                       `json:"url,required"`
	JSON userOrderInvoiceResponseJSON `json:"-"`
}

// userOrderInvoiceResponseJSON contains the JSON metadata for the struct
// [UserOrderInvoiceResponse]
type userOrderInvoiceResponseJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserOrderInvoiceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userOrderInvoiceResponseJSON) RawJSON() string {
	return r.raw
}

type UserOrderListParams struct {
	// Size of a page, defaults to 10. Maximum is 100.
	Limit param.Field[int64] `query:"limit"`
	// Filter by organization ID.
	OrganizationID param.Field[UserOrderListParamsOrganizationIDUnion] `query:"organization_id" format:"uuid4"`
	// Page number, defaults to 1.
	Page param.Field[int64] `query:"page"`
	// Filter by product ID.
	ProductID param.Field[UserOrderListParamsProductIDUnion] `query:"product_id" format:"uuid4"`
	// Filter by product price type. `recurring` will return orders corresponding to
	// subscriptions creations or renewals. `one_time` will return orders corresponding
	// to one-time purchases.
	ProductPriceType param.Field[UserOrderListParamsProductPriceTypeUnion] `query:"product_price_type"`
	// Search by product or organization name.
	Query param.Field[string] `query:"query"`
	// Sorting criterion. Several criteria can be used simultaneously and will be
	// applied in order. Add a minus sign `-` before the criteria name to sort by
	// descending order.
	Sorting param.Field[[]string] `query:"sorting"`
	// Filter by subscription ID.
	SubscriptionID param.Field[UserOrderListParamsSubscriptionIDUnion] `query:"subscription_id" format:"uuid4"`
}

// URLQuery serializes [UserOrderListParams]'s query parameters as `url.Values`.
func (r UserOrderListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by organization ID.
//
// Satisfied by [shared.UnionString], [UserOrderListParamsOrganizationIDArray].
type UserOrderListParamsOrganizationIDUnion interface {
	ImplementsUserOrderListParamsOrganizationIDUnion()
}

type UserOrderListParamsOrganizationIDArray []string

func (r UserOrderListParamsOrganizationIDArray) ImplementsUserOrderListParamsOrganizationIDUnion() {}

// Filter by product ID.
//
// Satisfied by [shared.UnionString], [UserOrderListParamsProductIDArray].
type UserOrderListParamsProductIDUnion interface {
	ImplementsUserOrderListParamsProductIDUnion()
}

type UserOrderListParamsProductIDArray []string

func (r UserOrderListParamsProductIDArray) ImplementsUserOrderListParamsProductIDUnion() {}

// Filter by product price type. `recurring` will return orders corresponding to
// subscriptions creations or renewals. `one_time` will return orders corresponding
// to one-time purchases.
//
// Satisfied by [UserOrderListParamsProductPriceTypeProductPriceType],
// [UserOrderListParamsProductPriceTypeArray].
type UserOrderListParamsProductPriceTypeUnion interface {
	implementsUserOrderListParamsProductPriceTypeUnion()
}

type UserOrderListParamsProductPriceTypeProductPriceType string

const (
	UserOrderListParamsProductPriceTypeProductPriceTypeOneTime   UserOrderListParamsProductPriceTypeProductPriceType = "one_time"
	UserOrderListParamsProductPriceTypeProductPriceTypeRecurring UserOrderListParamsProductPriceTypeProductPriceType = "recurring"
)

func (r UserOrderListParamsProductPriceTypeProductPriceType) IsKnown() bool {
	switch r {
	case UserOrderListParamsProductPriceTypeProductPriceTypeOneTime, UserOrderListParamsProductPriceTypeProductPriceTypeRecurring:
		return true
	}
	return false
}

func (r UserOrderListParamsProductPriceTypeProductPriceType) implementsUserOrderListParamsProductPriceTypeUnion() {
}

type UserOrderListParamsProductPriceTypeArray []UserOrderListParamsProductPriceTypeArray

func (r UserOrderListParamsProductPriceTypeArray) implementsUserOrderListParamsProductPriceTypeUnion() {
}

// Filter by subscription ID.
//
// Satisfied by [shared.UnionString], [UserOrderListParamsSubscriptionIDArray].
type UserOrderListParamsSubscriptionIDUnion interface {
	ImplementsUserOrderListParamsSubscriptionIDUnion()
}

type UserOrderListParamsSubscriptionIDArray []string

func (r UserOrderListParamsSubscriptionIDArray) ImplementsUserOrderListParamsSubscriptionIDUnion() {}
