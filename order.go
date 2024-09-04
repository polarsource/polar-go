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
  "github.com/polarsource/polar-go/internal/pagination"
  "github.com/polarsource/polar-go/internal/param"
  "github.com/polarsource/polar-go/internal/requestconfig"
  "github.com/polarsource/polar-go/option"
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
func (r *OrderService) List(ctx context.Context, query OrderListParams, opts ...option.RequestOption) (res *pagination.PolarPagination[OrderOutput], err error) {
  var raw *http.Response
  opts = append(r.Options[:], opts...)
  opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
  path := "v1/orders/"
  cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
  if err != nil {
    return nil, err
  }
  err = cfg.Execute()
  if err != nil {
    return nil, err
  }
  res.SetPageConfig(cfg, raw)
  return res, nil
}

// List orders.
func (r *OrderService) ListAutoPaging(ctx context.Context, query OrderListParams, opts ...option.RequestOption) (*pagination.PolarPaginationAutoPager[OrderOutput]) {
  return pagination.NewPolarPaginationAutoPager(r.List(ctx, query, opts...))
}

type ListResourceOrder struct {
Items []OrderOutput `json:"items,required"`
Pagination ListResourceOrderPagination `json:"pagination,required"`
JSON listResourceOrderJSON `json:"-"`
}

// listResourceOrderJSON contains the JSON metadata for the struct
// [ListResourceOrder]
type listResourceOrderJSON struct {
Items apijson.Field
Pagination apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ListResourceOrder) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r listResourceOrderJSON) RawJSON() (string) {
  return r.raw
}

type ListResourceOrderPagination struct {
MaxPage int64 `json:"max_page,required"`
TotalCount int64 `json:"total_count,required"`
JSON listResourceOrderPaginationJSON `json:"-"`
}

// listResourceOrderPaginationJSON contains the JSON metadata for the struct
// [ListResourceOrderPagination]
type listResourceOrderPaginationJSON struct {
MaxPage apijson.Field
TotalCount apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ListResourceOrderPagination) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r listResourceOrderPaginationJSON) RawJSON() (string) {
  return r.raw
}

type OrderOutput struct {
// The ID of the object.
ID string `json:"id,required" format:"uuid4"`
Amount int64 `json:"amount,required"`
// Creation timestamp of the object.
CreatedAt time.Time `json:"created_at,required" format:"date-time"`
Currency string `json:"currency,required"`
// Last modification timestamp of the object.
ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
Product OrderOutputProduct `json:"product,required"`
ProductID string `json:"product_id,required" format:"uuid4"`
// A recurring price for a product, i.e. a subscription.
ProductPrice ProductPrice `json:"product_price,required"`
ProductPriceID string `json:"product_price_id,required" format:"uuid4"`
Subscription OrderOutputSubscription `json:"subscription,required,nullable"`
SubscriptionID string `json:"subscription_id,required,nullable" format:"uuid4"`
TaxAmount int64 `json:"tax_amount,required"`
User OrderOutputUser `json:"user,required"`
UserID string `json:"user_id,required" format:"uuid4"`
JSON orderOutputJSON `json:"-"`
}

// orderOutputJSON contains the JSON metadata for the struct [OrderOutput]
type orderOutputJSON struct {
ID apijson.Field
Amount apijson.Field
CreatedAt apijson.Field
Currency apijson.Field
ModifiedAt apijson.Field
Product apijson.Field
ProductID apijson.Field
ProductPrice apijson.Field
ProductPriceID apijson.Field
Subscription apijson.Field
SubscriptionID apijson.Field
TaxAmount apijson.Field
User apijson.Field
UserID apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *OrderOutput) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r orderOutputJSON) RawJSON() (string) {
  return r.raw
}

type OrderOutputProduct struct {
// The ID of the product.
ID string `json:"id,required" format:"uuid4"`
// Creation timestamp of the object.
CreatedAt time.Time `json:"created_at,required" format:"date-time"`
// The description of the product.
Description string `json:"description,required,nullable"`
// Whether the product is archived and no longer available.
IsArchived bool `json:"is_archived,required"`
IsHighlighted bool `json:"is_highlighted,required,nullable"`
// Whether the product is a subscription tier.
IsRecurring bool `json:"is_recurring,required"`
// Last modification timestamp of the object.
ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
// The name of the product.
Name string `json:"name,required"`
// The ID of the organization owning the product.
OrganizationID string `json:"organization_id,required" format:"uuid4"`
Type OrderOutputProductType `json:"type,required,nullable"`
JSON orderOutputProductJSON `json:"-"`
}

// orderOutputProductJSON contains the JSON metadata for the struct
// [OrderOutputProduct]
type orderOutputProductJSON struct {
ID apijson.Field
CreatedAt apijson.Field
Description apijson.Field
IsArchived apijson.Field
IsHighlighted apijson.Field
IsRecurring apijson.Field
ModifiedAt apijson.Field
Name apijson.Field
OrganizationID apijson.Field
Type apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *OrderOutputProduct) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r orderOutputProductJSON) RawJSON() (string) {
  return r.raw
}

type OrderOutputProductType string

const (
  OrderOutputProductTypeFree OrderOutputProductType = "free"
  OrderOutputProductTypeIndividual OrderOutputProductType = "individual"
  OrderOutputProductTypeBusiness OrderOutputProductType = "business"
)

func (r OrderOutputProductType) IsKnown() (bool) {
  switch r {
  case OrderOutputProductTypeFree, OrderOutputProductTypeIndividual, OrderOutputProductTypeBusiness:
      return true
  }
  return false
}

type OrderOutputSubscription struct {
// The ID of the object.
ID string `json:"id,required" format:"uuid4"`
CancelAtPeriodEnd bool `json:"cancel_at_period_end,required"`
// Creation timestamp of the object.
CreatedAt time.Time `json:"created_at,required" format:"date-time"`
CurrentPeriodEnd time.Time `json:"current_period_end,required,nullable" format:"date-time"`
CurrentPeriodStart time.Time `json:"current_period_start,required" format:"date-time"`
EndedAt time.Time `json:"ended_at,required,nullable" format:"date-time"`
// Last modification timestamp of the object.
ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
PriceID string `json:"price_id,required,nullable" format:"uuid4"`
ProductID string `json:"product_id,required" format:"uuid4"`
StartedAt time.Time `json:"started_at,required,nullable" format:"date-time"`
Status OrderOutputSubscriptionStatus `json:"status,required"`
UserID string `json:"user_id,required" format:"uuid4"`
JSON orderOutputSubscriptionJSON `json:"-"`
}

// orderOutputSubscriptionJSON contains the JSON metadata for the struct
// [OrderOutputSubscription]
type orderOutputSubscriptionJSON struct {
ID apijson.Field
CancelAtPeriodEnd apijson.Field
CreatedAt apijson.Field
CurrentPeriodEnd apijson.Field
CurrentPeriodStart apijson.Field
EndedAt apijson.Field
ModifiedAt apijson.Field
PriceID apijson.Field
ProductID apijson.Field
StartedAt apijson.Field
Status apijson.Field
UserID apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *OrderOutputSubscription) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r orderOutputSubscriptionJSON) RawJSON() (string) {
  return r.raw
}

type OrderOutputSubscriptionStatus string

const (
  OrderOutputSubscriptionStatusIncomplete OrderOutputSubscriptionStatus = "incomplete"
  OrderOutputSubscriptionStatusIncompleteExpired OrderOutputSubscriptionStatus = "incomplete_expired"
  OrderOutputSubscriptionStatusTrialing OrderOutputSubscriptionStatus = "trialing"
  OrderOutputSubscriptionStatusActive OrderOutputSubscriptionStatus = "active"
  OrderOutputSubscriptionStatusPastDue OrderOutputSubscriptionStatus = "past_due"
  OrderOutputSubscriptionStatusCanceled OrderOutputSubscriptionStatus = "canceled"
  OrderOutputSubscriptionStatusUnpaid OrderOutputSubscriptionStatus = "unpaid"
)

func (r OrderOutputSubscriptionStatus) IsKnown() (bool) {
  switch r {
  case OrderOutputSubscriptionStatusIncomplete, OrderOutputSubscriptionStatusIncompleteExpired, OrderOutputSubscriptionStatusTrialing, OrderOutputSubscriptionStatusActive, OrderOutputSubscriptionStatusPastDue, OrderOutputSubscriptionStatusCanceled, OrderOutputSubscriptionStatusUnpaid:
      return true
  }
  return false
}

type OrderOutputUser struct {
ID string `json:"id,required" format:"uuid4"`
AvatarURL string `json:"avatar_url,required,nullable"`
Email string `json:"email,required"`
GitHubUsername string `json:"github_username,required,nullable"`
PublicName string `json:"public_name,required"`
JSON orderOutputUserJSON `json:"-"`
}

// orderOutputUserJSON contains the JSON metadata for the struct [OrderOutputUser]
type orderOutputUserJSON struct {
ID apijson.Field
AvatarURL apijson.Field
Email apijson.Field
GitHubUsername apijson.Field
PublicName apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *OrderOutputUser) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r orderOutputUserJSON) RawJSON() (string) {
  return r.raw
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
Sorting param.Field[[]OrderListParamsSorting] `query:"sorting"`
// Filter by customer's user ID.
UserID param.Field[OrderListParamsUserIDUnion] `query:"user_id" format:"uuid4"`
}

// URLQuery serializes [OrderListParams]'s query parameters as `url.Values`.
func (r OrderListParams) URLQuery() (v url.Values) {
  return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
    ArrayFormat: apiquery.ArrayQueryFormatComma,
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
  OrderListParamsProductPriceTypeProductPriceTypeOneTime OrderListParamsProductPriceTypeProductPriceType = "one_time"
  OrderListParamsProductPriceTypeProductPriceTypeRecurring OrderListParamsProductPriceTypeProductPriceType = "recurring"
)

func (r OrderListParamsProductPriceTypeProductPriceType) IsKnown() (bool) {
  switch r {
  case OrderListParamsProductPriceTypeProductPriceTypeOneTime, OrderListParamsProductPriceTypeProductPriceTypeRecurring:
      return true
  }
  return false
}

func (r OrderListParamsProductPriceTypeProductPriceType) implementsOrderListParamsProductPriceTypeUnion() {}

type OrderListParamsProductPriceTypeArray []OrderListParamsProductPriceTypeArray

func (r OrderListParamsProductPriceTypeArray) implementsOrderListParamsProductPriceTypeUnion() {}

type OrderListParamsSorting string

const (
  OrderListParamsSortingCreatedAt OrderListParamsSorting = "created_at"
  OrderListParamsSorting-CreatedAt OrderListParamsSorting = "-created_at"
  OrderListParamsSortingAmount OrderListParamsSorting = "amount"
  OrderListParamsSorting-Amount OrderListParamsSorting = "-amount"
  OrderListParamsSortingUser OrderListParamsSorting = "user"
  OrderListParamsSorting-User OrderListParamsSorting = "-user"
  OrderListParamsSortingProduct OrderListParamsSorting = "product"
  OrderListParamsSorting-Product OrderListParamsSorting = "-product"
  OrderListParamsSortingSubscription OrderListParamsSorting = "subscription"
  OrderListParamsSorting-Subscription OrderListParamsSorting = "-subscription"
)

func (r OrderListParamsSorting) IsKnown() (bool) {
  switch r {
  case OrderListParamsSortingCreatedAt, OrderListParamsSorting-CreatedAt, OrderListParamsSortingAmount, OrderListParamsSorting-Amount, OrderListParamsSortingUser, OrderListParamsSorting-User, OrderListParamsSortingProduct, OrderListParamsSorting-Product, OrderListParamsSortingSubscription, OrderListParamsSorting-Subscription:
      return true
  }
  return false
}

// Filter by customer's user ID.
//
// Satisfied by [shared.UnionString], [OrderListParamsUserIDArray].
type OrderListParamsUserIDUnion interface {
  ImplementsOrderListParamsUserIDUnion()
}

type OrderListParamsUserIDArray []string

func (r OrderListParamsUserIDArray) ImplementsOrderListParamsUserIDUnion() {}
