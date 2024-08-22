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

// OrderService contains methods and other services that help with interacting with
// the polar API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrderService] method instead.
type OrderService struct {
	Options []option.RequestOption
	Invoice *OrderInvoiceService
}

// NewOrderService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewOrderService(opts ...option.RequestOption) (r *OrderService) {
	r = &OrderService{}
	r.Options = opts
	r.Invoice = NewOrderInvoiceService(opts...)
	return
}

// Get an order by ID.
func (r *OrderService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *OrderOutput, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/orders/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List orders.
func (r *OrderService) List(ctx context.Context, query OrderListParams, opts ...option.RequestOption) (res *ListResourceOrder, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/orders/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ListResourceOrder struct {
	Pagination ListResourceOrderPagination `json:"pagination,required"`
	Items      []OrderOutput               `json:"items"`
	JSON       listResourceOrderJSON       `json:"-"`
}

// listResourceOrderJSON contains the JSON metadata for the struct
// [ListResourceOrder]
type listResourceOrderJSON struct {
	Pagination  apijson.Field
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListResourceOrder) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResourceOrderJSON) RawJSON() string {
	return r.raw
}

type ListResourceOrderPagination struct {
	MaxPage    int64                           `json:"max_page,required"`
	TotalCount int64                           `json:"total_count,required"`
	JSON       listResourceOrderPaginationJSON `json:"-"`
}

// listResourceOrderPaginationJSON contains the JSON metadata for the struct
// [ListResourceOrderPagination]
type listResourceOrderPaginationJSON struct {
	MaxPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListResourceOrderPagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResourceOrderPaginationJSON) RawJSON() string {
	return r.raw
}

type OrderOutput struct {
	// The ID of the object.
	ID     string `json:"id,required" format:"uuid4"`
	Amount int64  `json:"amount,required"`
	// Creation timestamp of the object.
	CreatedAt time.Time          `json:"created_at,required" format:"date-time"`
	Currency  string             `json:"currency,required"`
	Product   OrderOutputProduct `json:"product,required"`
	ProductID string             `json:"product_id,required" format:"uuid4"`
	// A recurring price for a product, i.e. a subscription.
	ProductPrice   OrderOutputProductPrice `json:"product_price,required"`
	ProductPriceID string                  `json:"product_price_id,required" format:"uuid4"`
	TaxAmount      int64                   `json:"tax_amount,required"`
	User           OrderOutputUser         `json:"user,required"`
	UserID         string                  `json:"user_id,required" format:"uuid4"`
	// Last modification timestamp of the object.
	ModifiedAt     time.Time               `json:"modified_at,nullable" format:"date-time"`
	Subscription   OrderOutputSubscription `json:"subscription,nullable"`
	SubscriptionID string                  `json:"subscription_id,nullable" format:"uuid4"`
	JSON           orderOutputJSON         `json:"-"`
}

// orderOutputJSON contains the JSON metadata for the struct [OrderOutput]
type orderOutputJSON struct {
	ID             apijson.Field
	Amount         apijson.Field
	CreatedAt      apijson.Field
	Currency       apijson.Field
	Product        apijson.Field
	ProductID      apijson.Field
	ProductPrice   apijson.Field
	ProductPriceID apijson.Field
	TaxAmount      apijson.Field
	User           apijson.Field
	UserID         apijson.Field
	ModifiedAt     apijson.Field
	Subscription   apijson.Field
	SubscriptionID apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *OrderOutput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orderOutputJSON) RawJSON() string {
	return r.raw
}

type OrderOutputProduct struct {
	// The ID of the product.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the product is archived and no longer available.
	IsArchived bool `json:"is_archived,required"`
	// Whether the product is a subscription tier.
	IsRecurring bool `json:"is_recurring,required"`
	// The name of the product.
	Name string `json:"name,required"`
	// The ID of the organization owning the product.
	OrganizationID string `json:"organization_id,required" format:"uuid4"`
	// The description of the product.
	Description   string `json:"description,nullable"`
	IsHighlighted bool   `json:"is_highlighted,nullable"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time              `json:"modified_at,nullable" format:"date-time"`
	Type       OrderOutputProductType `json:"type,nullable"`
	JSON       orderOutputProductJSON `json:"-"`
}

// orderOutputProductJSON contains the JSON metadata for the struct
// [OrderOutputProduct]
type orderOutputProductJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	IsArchived     apijson.Field
	IsRecurring    apijson.Field
	Name           apijson.Field
	OrganizationID apijson.Field
	Description    apijson.Field
	IsHighlighted  apijson.Field
	ModifiedAt     apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *OrderOutputProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orderOutputProductJSON) RawJSON() string {
	return r.raw
}

type OrderOutputProductType string

const (
	OrderOutputProductTypeFree       OrderOutputProductType = "free"
	OrderOutputProductTypeIndividual OrderOutputProductType = "individual"
	OrderOutputProductTypeBusiness   OrderOutputProductType = "business"
)

func (r OrderOutputProductType) IsKnown() bool {
	switch r {
	case OrderOutputProductTypeFree, OrderOutputProductTypeIndividual, OrderOutputProductTypeBusiness:
		return true
	}
	return false
}

// A recurring price for a product, i.e. a subscription.
type OrderOutputProductPrice struct {
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
	Type OrderOutputProductPriceType `json:"type,required"`
	// The recurring interval of the price, if type is `recurring`.
	RecurringInterval OrderOutputProductPriceRecurringInterval `json:"recurring_interval,nullable"`
	JSON              orderOutputProductPriceJSON              `json:"-"`
	union             OrderOutputProductPriceUnion
}

// orderOutputProductPriceJSON contains the JSON metadata for the struct
// [OrderOutputProductPrice]
type orderOutputProductPriceJSON struct {
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

func (r orderOutputProductPriceJSON) RawJSON() string {
	return r.raw
}

func (r *OrderOutputProductPrice) UnmarshalJSON(data []byte) (err error) {
	*r = OrderOutputProductPrice{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [OrderOutputProductPriceUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are
// [OrderOutputProductPriceProductPriceRecurring],
// [OrderOutputProductPriceProductPriceOneTime].
func (r OrderOutputProductPrice) AsUnion() OrderOutputProductPriceUnion {
	return r.union
}

// A recurring price for a product, i.e. a subscription.
//
// Union satisfied by [OrderOutputProductPriceProductPriceRecurring] or
// [OrderOutputProductPriceProductPriceOneTime].
type OrderOutputProductPriceUnion interface {
	implementsOrderOutputProductPrice()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*OrderOutputProductPriceUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(OrderOutputProductPriceProductPriceRecurring{}),
			DiscriminatorValue: "recurring",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(OrderOutputProductPriceProductPriceOneTime{}),
			DiscriminatorValue: "one_time",
		},
	)
}

// A recurring price for a product, i.e. a subscription.
type OrderOutputProductPriceProductPriceRecurring struct {
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
	Type OrderOutputProductPriceProductPriceRecurringType `json:"type,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,nullable" format:"date-time"`
	// The recurring interval of the price, if type is `recurring`.
	RecurringInterval OrderOutputProductPriceProductPriceRecurringRecurringInterval `json:"recurring_interval,nullable"`
	JSON              orderOutputProductPriceProductPriceRecurringJSON              `json:"-"`
}

// orderOutputProductPriceProductPriceRecurringJSON contains the JSON metadata for
// the struct [OrderOutputProductPriceProductPriceRecurring]
type orderOutputProductPriceProductPriceRecurringJSON struct {
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

func (r *OrderOutputProductPriceProductPriceRecurring) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orderOutputProductPriceProductPriceRecurringJSON) RawJSON() string {
	return r.raw
}

func (r OrderOutputProductPriceProductPriceRecurring) implementsOrderOutputProductPrice() {}

// The type of the price.
type OrderOutputProductPriceProductPriceRecurringType string

const (
	OrderOutputProductPriceProductPriceRecurringTypeRecurring OrderOutputProductPriceProductPriceRecurringType = "recurring"
)

func (r OrderOutputProductPriceProductPriceRecurringType) IsKnown() bool {
	switch r {
	case OrderOutputProductPriceProductPriceRecurringTypeRecurring:
		return true
	}
	return false
}

// The recurring interval of the price, if type is `recurring`.
type OrderOutputProductPriceProductPriceRecurringRecurringInterval string

const (
	OrderOutputProductPriceProductPriceRecurringRecurringIntervalMonth OrderOutputProductPriceProductPriceRecurringRecurringInterval = "month"
	OrderOutputProductPriceProductPriceRecurringRecurringIntervalYear  OrderOutputProductPriceProductPriceRecurringRecurringInterval = "year"
)

func (r OrderOutputProductPriceProductPriceRecurringRecurringInterval) IsKnown() bool {
	switch r {
	case OrderOutputProductPriceProductPriceRecurringRecurringIntervalMonth, OrderOutputProductPriceProductPriceRecurringRecurringIntervalYear:
		return true
	}
	return false
}

// A one-time price for a product.
type OrderOutputProductPriceProductPriceOneTime struct {
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
	Type OrderOutputProductPriceProductPriceOneTimeType `json:"type,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time                                      `json:"modified_at,nullable" format:"date-time"`
	JSON       orderOutputProductPriceProductPriceOneTimeJSON `json:"-"`
}

// orderOutputProductPriceProductPriceOneTimeJSON contains the JSON metadata for
// the struct [OrderOutputProductPriceProductPriceOneTime]
type orderOutputProductPriceProductPriceOneTimeJSON struct {
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

func (r *OrderOutputProductPriceProductPriceOneTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orderOutputProductPriceProductPriceOneTimeJSON) RawJSON() string {
	return r.raw
}

func (r OrderOutputProductPriceProductPriceOneTime) implementsOrderOutputProductPrice() {}

// The type of the price.
type OrderOutputProductPriceProductPriceOneTimeType string

const (
	OrderOutputProductPriceProductPriceOneTimeTypeOneTime OrderOutputProductPriceProductPriceOneTimeType = "one_time"
)

func (r OrderOutputProductPriceProductPriceOneTimeType) IsKnown() bool {
	switch r {
	case OrderOutputProductPriceProductPriceOneTimeTypeOneTime:
		return true
	}
	return false
}

// The type of the price.
type OrderOutputProductPriceType string

const (
	OrderOutputProductPriceTypeRecurring OrderOutputProductPriceType = "recurring"
	OrderOutputProductPriceTypeOneTime   OrderOutputProductPriceType = "one_time"
)

func (r OrderOutputProductPriceType) IsKnown() bool {
	switch r {
	case OrderOutputProductPriceTypeRecurring, OrderOutputProductPriceTypeOneTime:
		return true
	}
	return false
}

// The recurring interval of the price, if type is `recurring`.
type OrderOutputProductPriceRecurringInterval string

const (
	OrderOutputProductPriceRecurringIntervalMonth OrderOutputProductPriceRecurringInterval = "month"
	OrderOutputProductPriceRecurringIntervalYear  OrderOutputProductPriceRecurringInterval = "year"
)

func (r OrderOutputProductPriceRecurringInterval) IsKnown() bool {
	switch r {
	case OrderOutputProductPriceRecurringIntervalMonth, OrderOutputProductPriceRecurringIntervalYear:
		return true
	}
	return false
}

type OrderOutputUser struct {
	ID             string              `json:"id,required" format:"uuid4"`
	Email          string              `json:"email,required"`
	PublicName     string              `json:"public_name,required"`
	AvatarURL      string              `json:"avatar_url,nullable"`
	GitHubUsername string              `json:"github_username,nullable"`
	JSON           orderOutputUserJSON `json:"-"`
}

// orderOutputUserJSON contains the JSON metadata for the struct [OrderOutputUser]
type orderOutputUserJSON struct {
	ID             apijson.Field
	Email          apijson.Field
	PublicName     apijson.Field
	AvatarURL      apijson.Field
	GitHubUsername apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *OrderOutputUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orderOutputUserJSON) RawJSON() string {
	return r.raw
}

type OrderOutputSubscription struct {
	// The ID of the object.
	ID                string `json:"id,required" format:"uuid4"`
	CancelAtPeriodEnd bool   `json:"cancel_at_period_end,required"`
	// Creation timestamp of the object.
	CreatedAt          time.Time                     `json:"created_at,required" format:"date-time"`
	CurrentPeriodStart time.Time                     `json:"current_period_start,required" format:"date-time"`
	ProductID          string                        `json:"product_id,required" format:"uuid4"`
	Status             OrderOutputSubscriptionStatus `json:"status,required"`
	UserID             string                        `json:"user_id,required" format:"uuid4"`
	CurrentPeriodEnd   time.Time                     `json:"current_period_end,nullable" format:"date-time"`
	EndedAt            time.Time                     `json:"ended_at,nullable" format:"date-time"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time                   `json:"modified_at,nullable" format:"date-time"`
	PriceID    string                      `json:"price_id,nullable" format:"uuid4"`
	StartedAt  time.Time                   `json:"started_at,nullable" format:"date-time"`
	JSON       orderOutputSubscriptionJSON `json:"-"`
}

// orderOutputSubscriptionJSON contains the JSON metadata for the struct
// [OrderOutputSubscription]
type orderOutputSubscriptionJSON struct {
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

func (r *OrderOutputSubscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orderOutputSubscriptionJSON) RawJSON() string {
	return r.raw
}

type OrderOutputSubscriptionStatus string

const (
	OrderOutputSubscriptionStatusIncomplete        OrderOutputSubscriptionStatus = "incomplete"
	OrderOutputSubscriptionStatusIncompleteExpired OrderOutputSubscriptionStatus = "incomplete_expired"
	OrderOutputSubscriptionStatusTrialing          OrderOutputSubscriptionStatus = "trialing"
	OrderOutputSubscriptionStatusActive            OrderOutputSubscriptionStatus = "active"
	OrderOutputSubscriptionStatusPastDue           OrderOutputSubscriptionStatus = "past_due"
	OrderOutputSubscriptionStatusCanceled          OrderOutputSubscriptionStatus = "canceled"
	OrderOutputSubscriptionStatusUnpaid            OrderOutputSubscriptionStatus = "unpaid"
)

func (r OrderOutputSubscriptionStatus) IsKnown() bool {
	switch r {
	case OrderOutputSubscriptionStatusIncomplete, OrderOutputSubscriptionStatusIncompleteExpired, OrderOutputSubscriptionStatusTrialing, OrderOutputSubscriptionStatusActive, OrderOutputSubscriptionStatusPastDue, OrderOutputSubscriptionStatusCanceled, OrderOutputSubscriptionStatusUnpaid:
		return true
	}
	return false
}

type OrderListParams struct {
	// Size of a page, defaults to 10. Maximum is 100.
	Limit param.Field[int64] `query:"limit"`
	// Filter by organization ID.
	OrganizationID param.Field[OrderListParamsOrganizationIDUnion] `query:"organization_id" format:"uuid4"`
	// Page number, defaults to 1.
	Page param.Field[int64] `query:"page"`
	// Filter by product ID.
	ProductID param.Field[OrderListParamsProductIDUnion] `query:"product_id" format:"uuid4"`
	// Filter by product price type. `recurring` will return orders corresponding to
	// subscriptions creations or renewals. `one_time` will return orders corresponding
	// to one-time purchases.
	ProductPriceType param.Field[OrderListParamsProductPriceTypeUnion] `query:"product_price_type"`
	// Sorting criterion. Several criteria can be used simultaneously and will be
	// applied in order. Add a minus sign `-` before the criteria name to sort by
	// descending order.
	Sorting param.Field[[]string] `query:"sorting"`
	// Filter by customer's user ID.
	UserID param.Field[OrderListParamsUserIDUnion] `query:"user_id" format:"uuid4"`
}

// URLQuery serializes [OrderListParams]'s query parameters as `url.Values`.
func (r OrderListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by organization ID.
//
// Satisfied by [shared.UnionString], [OrderListParamsOrganizationIDArray].
type OrderListParamsOrganizationIDUnion interface {
	ImplementsOrderListParamsOrganizationIDUnion()
}

type OrderListParamsOrganizationIDArray []string

func (r OrderListParamsOrganizationIDArray) ImplementsOrderListParamsOrganizationIDUnion() {}

// Filter by product ID.
//
// Satisfied by [shared.UnionString], [OrderListParamsProductIDArray].
type OrderListParamsProductIDUnion interface {
	ImplementsOrderListParamsProductIDUnion()
}

type OrderListParamsProductIDArray []string

func (r OrderListParamsProductIDArray) ImplementsOrderListParamsProductIDUnion() {}

// Filter by product price type. `recurring` will return orders corresponding to
// subscriptions creations or renewals. `one_time` will return orders corresponding
// to one-time purchases.
//
// Satisfied by [OrderListParamsProductPriceTypeProductPriceType],
// [OrderListParamsProductPriceTypeArray].
type OrderListParamsProductPriceTypeUnion interface {
	implementsOrderListParamsProductPriceTypeUnion()
}

type OrderListParamsProductPriceTypeProductPriceType string

const (
	OrderListParamsProductPriceTypeProductPriceTypeOneTime   OrderListParamsProductPriceTypeProductPriceType = "one_time"
	OrderListParamsProductPriceTypeProductPriceTypeRecurring OrderListParamsProductPriceTypeProductPriceType = "recurring"
)

func (r OrderListParamsProductPriceTypeProductPriceType) IsKnown() bool {
	switch r {
	case OrderListParamsProductPriceTypeProductPriceTypeOneTime, OrderListParamsProductPriceTypeProductPriceTypeRecurring:
		return true
	}
	return false
}

func (r OrderListParamsProductPriceTypeProductPriceType) implementsOrderListParamsProductPriceTypeUnion() {
}

type OrderListParamsProductPriceTypeArray []OrderListParamsProductPriceTypeArray

func (r OrderListParamsProductPriceTypeArray) implementsOrderListParamsProductPriceTypeUnion() {}

// Filter by customer's user ID.
//
// Satisfied by [shared.UnionString], [OrderListParamsUserIDArray].
type OrderListParamsUserIDUnion interface {
	ImplementsOrderListParamsUserIDUnion()
}

type OrderListParamsUserIDArray []string

func (r OrderListParamsUserIDArray) ImplementsOrderListParamsUserIDUnion() {}
