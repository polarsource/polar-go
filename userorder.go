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
  "github.com/polarsource/polar-go/internal/pagination"
  "github.com/polarsource/polar-go/internal/param"
  "github.com/polarsource/polar-go/internal/requestconfig"
  "github.com/polarsource/polar-go/option"
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
func (r *UserOrderService) List(ctx context.Context, query UserOrderListParams, opts ...option.RequestOption) (res *pagination.PolarPagination[UserOrderListResponse], err error) {
  var raw *http.Response
  opts = append(r.Options[:], opts...)
  opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
  path := "v1/users/orders/"
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

// List my orders.
func (r *UserOrderService) ListAutoPaging(ctx context.Context, query UserOrderListParams, opts ...option.RequestOption) (*pagination.PolarPaginationAutoPager[UserOrderListResponse]) {
  return pagination.NewPolarPaginationAutoPager(r.List(ctx, query, opts...))
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
ID string `json:"id,required" format:"uuid4"`
Amount int64 `json:"amount,required"`
// Creation timestamp of the object.
CreatedAt time.Time `json:"created_at,required" format:"date-time"`
Currency string `json:"currency,required"`
// Last modification timestamp of the object.
ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
Product UserOrderGetResponseProduct `json:"product,required"`
ProductID string `json:"product_id,required" format:"uuid4"`
// A recurring price for a product, i.e. a subscription.
ProductPrice UserOrderGetResponseProductPrice `json:"product_price,required"`
ProductPriceID string `json:"product_price_id,required" format:"uuid4"`
Subscription UserOrderGetResponseSubscription `json:"subscription,required,nullable"`
SubscriptionID string `json:"subscription_id,required,nullable" format:"uuid4"`
TaxAmount int64 `json:"tax_amount,required"`
UserID string `json:"user_id,required" format:"uuid4"`
JSON userOrderGetResponseJSON `json:"-"`
}

// userOrderGetResponseJSON contains the JSON metadata for the struct
// [UserOrderGetResponse]
type userOrderGetResponseJSON struct {
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
UserID apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *UserOrderGetResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userOrderGetResponseJSON) RawJSON() (string) {
  return r.raw
}

type UserOrderGetResponseProduct struct {
// The ID of the product.
ID string `json:"id,required" format:"uuid4"`
// The benefits granted by the product.
Benefits []UserOrderGetResponseProductBenefit `json:"benefits,required"`
// Creation timestamp of the object.
CreatedAt time.Time `json:"created_at,required" format:"date-time"`
// The description of the product.
Description string `json:"description,required,nullable"`
// Whether the product is archived and no longer available.
IsArchived bool `json:"is_archived,required"`
IsHighlighted bool `json:"is_highlighted,required,nullable"`
// Whether the product is a subscription tier.
IsRecurring bool `json:"is_recurring,required"`
// The medias associated to the product.
Medias []ProductMediaFileReadOutput `json:"medias,required"`
// Last modification timestamp of the object.
ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
// The name of the product.
Name string `json:"name,required"`
// The ID of the organization owning the product.
OrganizationID string `json:"organization_id,required" format:"uuid4"`
// List of available prices for this product.
Prices []UserOrderGetResponseProductPrice `json:"prices,required"`
Type UserOrderGetResponseProductType `json:"type,required,nullable"`
JSON userOrderGetResponseProductJSON `json:"-"`
}

// userOrderGetResponseProductJSON contains the JSON metadata for the struct
// [UserOrderGetResponseProduct]
type userOrderGetResponseProductJSON struct {
ID apijson.Field
Benefits apijson.Field
CreatedAt apijson.Field
Description apijson.Field
IsArchived apijson.Field
IsHighlighted apijson.Field
IsRecurring apijson.Field
Medias apijson.Field
ModifiedAt apijson.Field
Name apijson.Field
OrganizationID apijson.Field
Prices apijson.Field
Type apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *UserOrderGetResponseProduct) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userOrderGetResponseProductJSON) RawJSON() (string) {
  return r.raw
}

// A benefit of type `articles`.
//
// Use it to grant access to posts.
type UserOrderGetResponseProductBenefit struct {
// Creation timestamp of the object.
CreatedAt time.Time `json:"created_at,required" format:"date-time"`
// Last modification timestamp of the object.
ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
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
Properties interface{} `json:"properties,required"`
JSON userOrderGetResponseProductBenefitJSON `json:"-"`
union UserOrderGetResponseProductBenefitsUnion
}

// userOrderGetResponseProductBenefitJSON contains the JSON metadata for the struct
// [UserOrderGetResponseProductBenefit]
type userOrderGetResponseProductBenefitJSON struct {
CreatedAt apijson.Field
ModifiedAt apijson.Field
ID apijson.Field
Type apijson.Field
Description apijson.Field
Selectable apijson.Field
Deletable apijson.Field
OrganizationID apijson.Field
Properties apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r userOrderGetResponseProductBenefitJSON) RawJSON() (string) {
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
func (r UserOrderGetResponseProductBenefit) AsUnion() (UserOrderGetResponseProductBenefitsUnion) {
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
      Type: reflect.TypeOf(UserOrderGetResponseProductBenefitsBenefitBase{}),
    },
    apijson.UnionVariant{
      TypeFilter: gjson.JSON,
      Type: reflect.TypeOf(UserOrderGetResponseProductBenefitsBenefitArticles{}),
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
// Last modification timestamp of the object.
ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
// The ID of the organization owning the benefit.
OrganizationID string `json:"organization_id,required" format:"uuid4"`
// Whether the benefit is selectable when creating a product.
Selectable bool `json:"selectable,required"`
// The type of the benefit.
Type UserOrderGetResponseProductBenefitsBenefitBaseType `json:"type,required"`
JSON userOrderGetResponseProductBenefitsBenefitBaseJSON `json:"-"`
}

// userOrderGetResponseProductBenefitsBenefitBaseJSON contains the JSON metadata
// for the struct [UserOrderGetResponseProductBenefitsBenefitBase]
type userOrderGetResponseProductBenefitsBenefitBaseJSON struct {
ID apijson.Field
CreatedAt apijson.Field
Deletable apijson.Field
Description apijson.Field
ModifiedAt apijson.Field
OrganizationID apijson.Field
Selectable apijson.Field
Type apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *UserOrderGetResponseProductBenefitsBenefitBase) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userOrderGetResponseProductBenefitsBenefitBaseJSON) RawJSON() (string) {
  return r.raw
}

func (r UserOrderGetResponseProductBenefitsBenefitBase) implementsUserOrderGetResponseProductBenefit() {}

// The type of the benefit.
type UserOrderGetResponseProductBenefitsBenefitBaseType string

const (
  UserOrderGetResponseProductBenefitsBenefitBaseTypeCustom UserOrderGetResponseProductBenefitsBenefitBaseType = "custom"
  UserOrderGetResponseProductBenefitsBenefitBaseTypeArticles UserOrderGetResponseProductBenefitsBenefitBaseType = "articles"
  UserOrderGetResponseProductBenefitsBenefitBaseTypeAds UserOrderGetResponseProductBenefitsBenefitBaseType = "ads"
  UserOrderGetResponseProductBenefitsBenefitBaseTypeDiscord UserOrderGetResponseProductBenefitsBenefitBaseType = "discord"
  UserOrderGetResponseProductBenefitsBenefitBaseTypeGitHubRepository UserOrderGetResponseProductBenefitsBenefitBaseType = "github_repository"
  UserOrderGetResponseProductBenefitsBenefitBaseTypeDownloadables UserOrderGetResponseProductBenefitsBenefitBaseType = "downloadables"
)

func (r UserOrderGetResponseProductBenefitsBenefitBaseType) IsKnown() (bool) {
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
// Last modification timestamp of the object.
ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
// The ID of the organization owning the benefit.
OrganizationID string `json:"organization_id,required" format:"uuid4"`
// Properties for a benefit of type `articles`.
Properties UserOrderGetResponseProductBenefitsBenefitArticlesProperties `json:"properties,required"`
// Whether the benefit is selectable when creating a product.
Selectable bool `json:"selectable,required"`
Type UserOrderGetResponseProductBenefitsBenefitArticlesType `json:"type,required"`
JSON userOrderGetResponseProductBenefitsBenefitArticlesJSON `json:"-"`
}

// userOrderGetResponseProductBenefitsBenefitArticlesJSON contains the JSON
// metadata for the struct [UserOrderGetResponseProductBenefitsBenefitArticles]
type userOrderGetResponseProductBenefitsBenefitArticlesJSON struct {
ID apijson.Field
CreatedAt apijson.Field
Deletable apijson.Field
Description apijson.Field
ModifiedAt apijson.Field
OrganizationID apijson.Field
Properties apijson.Field
Selectable apijson.Field
Type apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *UserOrderGetResponseProductBenefitsBenefitArticles) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userOrderGetResponseProductBenefitsBenefitArticlesJSON) RawJSON() (string) {
  return r.raw
}

func (r UserOrderGetResponseProductBenefitsBenefitArticles) implementsUserOrderGetResponseProductBenefit() {}

// Properties for a benefit of type `articles`.
type UserOrderGetResponseProductBenefitsBenefitArticlesProperties struct {
// Whether the user can access paid articles.
PaidArticles bool `json:"paid_articles,required"`
JSON userOrderGetResponseProductBenefitsBenefitArticlesPropertiesJSON `json:"-"`
}

// userOrderGetResponseProductBenefitsBenefitArticlesPropertiesJSON contains the
// JSON metadata for the struct
// [UserOrderGetResponseProductBenefitsBenefitArticlesProperties]
type userOrderGetResponseProductBenefitsBenefitArticlesPropertiesJSON struct {
PaidArticles apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *UserOrderGetResponseProductBenefitsBenefitArticlesProperties) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userOrderGetResponseProductBenefitsBenefitArticlesPropertiesJSON) RawJSON() (string) {
  return r.raw
}

type UserOrderGetResponseProductBenefitsBenefitArticlesType string

const (
  UserOrderGetResponseProductBenefitsBenefitArticlesTypeArticles UserOrderGetResponseProductBenefitsBenefitArticlesType = "articles"
)

func (r UserOrderGetResponseProductBenefitsBenefitArticlesType) IsKnown() (bool) {
  switch r {
  case UserOrderGetResponseProductBenefitsBenefitArticlesTypeArticles:
      return true
  }
  return false
}

// The type of the benefit.
type UserOrderGetResponseProductBenefitsType string

const (
  UserOrderGetResponseProductBenefitsTypeCustom UserOrderGetResponseProductBenefitsType = "custom"
  UserOrderGetResponseProductBenefitsTypeArticles UserOrderGetResponseProductBenefitsType = "articles"
  UserOrderGetResponseProductBenefitsTypeAds UserOrderGetResponseProductBenefitsType = "ads"
  UserOrderGetResponseProductBenefitsTypeDiscord UserOrderGetResponseProductBenefitsType = "discord"
  UserOrderGetResponseProductBenefitsTypeGitHubRepository UserOrderGetResponseProductBenefitsType = "github_repository"
  UserOrderGetResponseProductBenefitsTypeDownloadables UserOrderGetResponseProductBenefitsType = "downloadables"
)

func (r UserOrderGetResponseProductBenefitsType) IsKnown() (bool) {
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
ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
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
JSON userOrderGetResponseProductPriceJSON `json:"-"`
union UserOrderGetResponseProductPricesUnion
}

// userOrderGetResponseProductPriceJSON contains the JSON metadata for the struct
// [UserOrderGetResponseProductPrice]
type userOrderGetResponseProductPriceJSON struct {
CreatedAt apijson.Field
ModifiedAt apijson.Field
ID apijson.Field
PriceAmount apijson.Field
PriceCurrency apijson.Field
IsArchived apijson.Field
Type apijson.Field
RecurringInterval apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r userOrderGetResponseProductPriceJSON) RawJSON() (string) {
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
func (r UserOrderGetResponseProductPrice) AsUnion() (UserOrderGetResponseProductPricesUnion) {
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
      TypeFilter: gjson.JSON,
      Type: reflect.TypeOf(UserOrderGetResponseProductPricesProductPriceRecurring{}),
      DiscriminatorValue: "recurring",
    },
    apijson.UnionVariant{
      TypeFilter: gjson.JSON,
      Type: reflect.TypeOf(UserOrderGetResponseProductPricesProductPriceOneTime{}),
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
// Last modification timestamp of the object.
ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
// The price in cents.
PriceAmount int64 `json:"price_amount,required"`
// The currency.
PriceCurrency string `json:"price_currency,required"`
// The recurring interval of the price, if type is `recurring`.
RecurringInterval UserOrderGetResponseProductPricesProductPriceRecurringRecurringInterval `json:"recurring_interval,required,nullable"`
// The type of the price.
Type UserOrderGetResponseProductPricesProductPriceRecurringType `json:"type,required"`
JSON userOrderGetResponseProductPricesProductPriceRecurringJSON `json:"-"`
}

// userOrderGetResponseProductPricesProductPriceRecurringJSON contains the JSON
// metadata for the struct [UserOrderGetResponseProductPricesProductPriceRecurring]
type userOrderGetResponseProductPricesProductPriceRecurringJSON struct {
ID apijson.Field
CreatedAt apijson.Field
IsArchived apijson.Field
ModifiedAt apijson.Field
PriceAmount apijson.Field
PriceCurrency apijson.Field
RecurringInterval apijson.Field
Type apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *UserOrderGetResponseProductPricesProductPriceRecurring) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userOrderGetResponseProductPricesProductPriceRecurringJSON) RawJSON() (string) {
  return r.raw
}

func (r UserOrderGetResponseProductPricesProductPriceRecurring) implementsUserOrderGetResponseProductPrice() {}

// The recurring interval of the price, if type is `recurring`.
type UserOrderGetResponseProductPricesProductPriceRecurringRecurringInterval string

const (
  UserOrderGetResponseProductPricesProductPriceRecurringRecurringIntervalMonth UserOrderGetResponseProductPricesProductPriceRecurringRecurringInterval = "month"
  UserOrderGetResponseProductPricesProductPriceRecurringRecurringIntervalYear UserOrderGetResponseProductPricesProductPriceRecurringRecurringInterval = "year"
)

func (r UserOrderGetResponseProductPricesProductPriceRecurringRecurringInterval) IsKnown() (bool) {
  switch r {
  case UserOrderGetResponseProductPricesProductPriceRecurringRecurringIntervalMonth, UserOrderGetResponseProductPricesProductPriceRecurringRecurringIntervalYear:
      return true
  }
  return false
}

// The type of the price.
type UserOrderGetResponseProductPricesProductPriceRecurringType string

const (
  UserOrderGetResponseProductPricesProductPriceRecurringTypeRecurring UserOrderGetResponseProductPricesProductPriceRecurringType = "recurring"
)

func (r UserOrderGetResponseProductPricesProductPriceRecurringType) IsKnown() (bool) {
  switch r {
  case UserOrderGetResponseProductPricesProductPriceRecurringTypeRecurring:
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
// Last modification timestamp of the object.
ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
// The price in cents.
PriceAmount int64 `json:"price_amount,required"`
// The currency.
PriceCurrency string `json:"price_currency,required"`
// The type of the price.
Type UserOrderGetResponseProductPricesProductPriceOneTimeType `json:"type,required"`
JSON userOrderGetResponseProductPricesProductPriceOneTimeJSON `json:"-"`
}

// userOrderGetResponseProductPricesProductPriceOneTimeJSON contains the JSON
// metadata for the struct [UserOrderGetResponseProductPricesProductPriceOneTime]
type userOrderGetResponseProductPricesProductPriceOneTimeJSON struct {
ID apijson.Field
CreatedAt apijson.Field
IsArchived apijson.Field
ModifiedAt apijson.Field
PriceAmount apijson.Field
PriceCurrency apijson.Field
Type apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *UserOrderGetResponseProductPricesProductPriceOneTime) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userOrderGetResponseProductPricesProductPriceOneTimeJSON) RawJSON() (string) {
  return r.raw
}

func (r UserOrderGetResponseProductPricesProductPriceOneTime) implementsUserOrderGetResponseProductPrice() {}

// The type of the price.
type UserOrderGetResponseProductPricesProductPriceOneTimeType string

const (
  UserOrderGetResponseProductPricesProductPriceOneTimeTypeOneTime UserOrderGetResponseProductPricesProductPriceOneTimeType = "one_time"
)

func (r UserOrderGetResponseProductPricesProductPriceOneTimeType) IsKnown() (bool) {
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
  UserOrderGetResponseProductPricesTypeOneTime UserOrderGetResponseProductPricesType = "one_time"
)

func (r UserOrderGetResponseProductPricesType) IsKnown() (bool) {
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
  UserOrderGetResponseProductPricesRecurringIntervalYear UserOrderGetResponseProductPricesRecurringInterval = "year"
)

func (r UserOrderGetResponseProductPricesRecurringInterval) IsKnown() (bool) {
  switch r {
  case UserOrderGetResponseProductPricesRecurringIntervalMonth, UserOrderGetResponseProductPricesRecurringIntervalYear:
      return true
  }
  return false
}

type UserOrderGetResponseProductType string

const (
  UserOrderGetResponseProductTypeFree UserOrderGetResponseProductType = "free"
  UserOrderGetResponseProductTypeIndividual UserOrderGetResponseProductType = "individual"
  UserOrderGetResponseProductTypeBusiness UserOrderGetResponseProductType = "business"
)

func (r UserOrderGetResponseProductType) IsKnown() (bool) {
  switch r {
  case UserOrderGetResponseProductTypeFree, UserOrderGetResponseProductTypeIndividual, UserOrderGetResponseProductTypeBusiness:
      return true
  }
  return false
}

type UserOrderGetResponseSubscription struct {
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
Status UserOrderGetResponseSubscriptionStatus `json:"status,required"`
UserID string `json:"user_id,required" format:"uuid4"`
JSON userOrderGetResponseSubscriptionJSON `json:"-"`
}

// userOrderGetResponseSubscriptionJSON contains the JSON metadata for the struct
// [UserOrderGetResponseSubscription]
type userOrderGetResponseSubscriptionJSON struct {
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

func (r *UserOrderGetResponseSubscription) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userOrderGetResponseSubscriptionJSON) RawJSON() (string) {
  return r.raw
}

type UserOrderGetResponseSubscriptionStatus string

const (
  UserOrderGetResponseSubscriptionStatusIncomplete UserOrderGetResponseSubscriptionStatus = "incomplete"
  UserOrderGetResponseSubscriptionStatusIncompleteExpired UserOrderGetResponseSubscriptionStatus = "incomplete_expired"
  UserOrderGetResponseSubscriptionStatusTrialing UserOrderGetResponseSubscriptionStatus = "trialing"
  UserOrderGetResponseSubscriptionStatusActive UserOrderGetResponseSubscriptionStatus = "active"
  UserOrderGetResponseSubscriptionStatusPastDue UserOrderGetResponseSubscriptionStatus = "past_due"
  UserOrderGetResponseSubscriptionStatusCanceled UserOrderGetResponseSubscriptionStatus = "canceled"
  UserOrderGetResponseSubscriptionStatusUnpaid UserOrderGetResponseSubscriptionStatus = "unpaid"
)

func (r UserOrderGetResponseSubscriptionStatus) IsKnown() (bool) {
  switch r {
  case UserOrderGetResponseSubscriptionStatusIncomplete, UserOrderGetResponseSubscriptionStatusIncompleteExpired, UserOrderGetResponseSubscriptionStatusTrialing, UserOrderGetResponseSubscriptionStatusActive, UserOrderGetResponseSubscriptionStatusPastDue, UserOrderGetResponseSubscriptionStatusCanceled, UserOrderGetResponseSubscriptionStatusUnpaid:
      return true
  }
  return false
}

type UserOrderListResponse struct {
ID string `json:"id,required" format:"uuid4"`
Amount int64 `json:"amount,required"`
// Creation timestamp of the object.
CreatedAt time.Time `json:"created_at,required" format:"date-time"`
Currency string `json:"currency,required"`
// Last modification timestamp of the object.
ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
Product UserOrderListResponseProduct `json:"product,required"`
ProductID string `json:"product_id,required" format:"uuid4"`
// A recurring price for a product, i.e. a subscription.
ProductPrice UserOrderListResponseProductPrice `json:"product_price,required"`
ProductPriceID string `json:"product_price_id,required" format:"uuid4"`
Subscription UserOrderListResponseSubscription `json:"subscription,required,nullable"`
SubscriptionID string `json:"subscription_id,required,nullable" format:"uuid4"`
TaxAmount int64 `json:"tax_amount,required"`
UserID string `json:"user_id,required" format:"uuid4"`
JSON userOrderListResponseJSON `json:"-"`
}

// userOrderListResponseJSON contains the JSON metadata for the struct
// [UserOrderListResponse]
type userOrderListResponseJSON struct {
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
UserID apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *UserOrderListResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userOrderListResponseJSON) RawJSON() (string) {
  return r.raw
}

type UserOrderListResponseProduct struct {
// The ID of the product.
ID string `json:"id,required" format:"uuid4"`
// The benefits granted by the product.
Benefits []UserOrderListResponseProductBenefit `json:"benefits,required"`
// Creation timestamp of the object.
CreatedAt time.Time `json:"created_at,required" format:"date-time"`
// The description of the product.
Description string `json:"description,required,nullable"`
// Whether the product is archived and no longer available.
IsArchived bool `json:"is_archived,required"`
IsHighlighted bool `json:"is_highlighted,required,nullable"`
// Whether the product is a subscription tier.
IsRecurring bool `json:"is_recurring,required"`
// The medias associated to the product.
Medias []ProductMediaFileReadOutput `json:"medias,required"`
// Last modification timestamp of the object.
ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
// The name of the product.
Name string `json:"name,required"`
// The ID of the organization owning the product.
OrganizationID string `json:"organization_id,required" format:"uuid4"`
// List of available prices for this product.
Prices []UserOrderListResponseProductPrice `json:"prices,required"`
Type UserOrderListResponseProductType `json:"type,required,nullable"`
JSON userOrderListResponseProductJSON `json:"-"`
}

// userOrderListResponseProductJSON contains the JSON metadata for the struct
// [UserOrderListResponseProduct]
type userOrderListResponseProductJSON struct {
ID apijson.Field
Benefits apijson.Field
CreatedAt apijson.Field
Description apijson.Field
IsArchived apijson.Field
IsHighlighted apijson.Field
IsRecurring apijson.Field
Medias apijson.Field
ModifiedAt apijson.Field
Name apijson.Field
OrganizationID apijson.Field
Prices apijson.Field
Type apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *UserOrderListResponseProduct) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userOrderListResponseProductJSON) RawJSON() (string) {
  return r.raw
}

// A benefit of type `articles`.
//
// Use it to grant access to posts.
type UserOrderListResponseProductBenefit struct {
// Creation timestamp of the object.
CreatedAt time.Time `json:"created_at,required" format:"date-time"`
// Last modification timestamp of the object.
ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
// The ID of the benefit.
ID string `json:"id,required" format:"uuid4"`
// The type of the benefit.
Type UserOrderListResponseProductBenefitsType `json:"type,required"`
// The description of the benefit.
Description string `json:"description,required"`
// Whether the benefit is selectable when creating a product.
Selectable bool `json:"selectable,required"`
// Whether the benefit is deletable.
Deletable bool `json:"deletable,required"`
// The ID of the organization owning the benefit.
OrganizationID string `json:"organization_id,required" format:"uuid4"`
// This field can have the runtime type of
// [UserOrderListResponseProductBenefitsBenefitArticlesProperties].
Properties interface{} `json:"properties,required"`
JSON userOrderListResponseProductBenefitJSON `json:"-"`
union UserOrderListResponseProductBenefitsUnion
}

// userOrderListResponseProductBenefitJSON contains the JSON metadata for the
// struct [UserOrderListResponseProductBenefit]
type userOrderListResponseProductBenefitJSON struct {
CreatedAt apijson.Field
ModifiedAt apijson.Field
ID apijson.Field
Type apijson.Field
Description apijson.Field
Selectable apijson.Field
Deletable apijson.Field
OrganizationID apijson.Field
Properties apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r userOrderListResponseProductBenefitJSON) RawJSON() (string) {
  return r.raw
}

func (r *UserOrderListResponseProductBenefit) UnmarshalJSON(data []byte) (err error) {
  *r = UserOrderListResponseProductBenefit{}
  err = apijson.UnmarshalRoot(data, &r.union)
  if err != nil {
    return err
  }
  return apijson.Port(r.union, &r)
}

// AsUnion returns a [UserOrderListResponseProductBenefitsUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [UserOrderListResponseProductBenefitsBenefitBase],
// [UserOrderListResponseProductBenefitsBenefitArticles].
func (r UserOrderListResponseProductBenefit) AsUnion() (UserOrderListResponseProductBenefitsUnion) {
  return r.union
}

// A benefit of type `articles`.
//
// Use it to grant access to posts.
//
// Union satisfied by [UserOrderListResponseProductBenefitsBenefitBase] or
// [UserOrderListResponseProductBenefitsBenefitArticles].
type UserOrderListResponseProductBenefitsUnion interface {
  implementsUserOrderListResponseProductBenefit()
}

func init() {
  apijson.RegisterUnion(
    reflect.TypeOf((*UserOrderListResponseProductBenefitsUnion)(nil)).Elem(),
    "",
    apijson.UnionVariant{
      TypeFilter: gjson.JSON,
      Type: reflect.TypeOf(UserOrderListResponseProductBenefitsBenefitBase{}),
    },
    apijson.UnionVariant{
      TypeFilter: gjson.JSON,
      Type: reflect.TypeOf(UserOrderListResponseProductBenefitsBenefitArticles{}),
    },
  )
}

type UserOrderListResponseProductBenefitsBenefitBase struct {
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
// Whether the benefit is selectable when creating a product.
Selectable bool `json:"selectable,required"`
// The type of the benefit.
Type UserOrderListResponseProductBenefitsBenefitBaseType `json:"type,required"`
JSON userOrderListResponseProductBenefitsBenefitBaseJSON `json:"-"`
}

// userOrderListResponseProductBenefitsBenefitBaseJSON contains the JSON metadata
// for the struct [UserOrderListResponseProductBenefitsBenefitBase]
type userOrderListResponseProductBenefitsBenefitBaseJSON struct {
ID apijson.Field
CreatedAt apijson.Field
Deletable apijson.Field
Description apijson.Field
ModifiedAt apijson.Field
OrganizationID apijson.Field
Selectable apijson.Field
Type apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *UserOrderListResponseProductBenefitsBenefitBase) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userOrderListResponseProductBenefitsBenefitBaseJSON) RawJSON() (string) {
  return r.raw
}

func (r UserOrderListResponseProductBenefitsBenefitBase) implementsUserOrderListResponseProductBenefit() {}

// The type of the benefit.
type UserOrderListResponseProductBenefitsBenefitBaseType string

const (
  UserOrderListResponseProductBenefitsBenefitBaseTypeCustom UserOrderListResponseProductBenefitsBenefitBaseType = "custom"
  UserOrderListResponseProductBenefitsBenefitBaseTypeArticles UserOrderListResponseProductBenefitsBenefitBaseType = "articles"
  UserOrderListResponseProductBenefitsBenefitBaseTypeAds UserOrderListResponseProductBenefitsBenefitBaseType = "ads"
  UserOrderListResponseProductBenefitsBenefitBaseTypeDiscord UserOrderListResponseProductBenefitsBenefitBaseType = "discord"
  UserOrderListResponseProductBenefitsBenefitBaseTypeGitHubRepository UserOrderListResponseProductBenefitsBenefitBaseType = "github_repository"
  UserOrderListResponseProductBenefitsBenefitBaseTypeDownloadables UserOrderListResponseProductBenefitsBenefitBaseType = "downloadables"
)

func (r UserOrderListResponseProductBenefitsBenefitBaseType) IsKnown() (bool) {
  switch r {
  case UserOrderListResponseProductBenefitsBenefitBaseTypeCustom, UserOrderListResponseProductBenefitsBenefitBaseTypeArticles, UserOrderListResponseProductBenefitsBenefitBaseTypeAds, UserOrderListResponseProductBenefitsBenefitBaseTypeDiscord, UserOrderListResponseProductBenefitsBenefitBaseTypeGitHubRepository, UserOrderListResponseProductBenefitsBenefitBaseTypeDownloadables:
      return true
  }
  return false
}

// A benefit of type `articles`.
//
// Use it to grant access to posts.
type UserOrderListResponseProductBenefitsBenefitArticles struct {
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
// Properties for a benefit of type `articles`.
Properties UserOrderListResponseProductBenefitsBenefitArticlesProperties `json:"properties,required"`
// Whether the benefit is selectable when creating a product.
Selectable bool `json:"selectable,required"`
Type UserOrderListResponseProductBenefitsBenefitArticlesType `json:"type,required"`
JSON userOrderListResponseProductBenefitsBenefitArticlesJSON `json:"-"`
}

// userOrderListResponseProductBenefitsBenefitArticlesJSON contains the JSON
// metadata for the struct [UserOrderListResponseProductBenefitsBenefitArticles]
type userOrderListResponseProductBenefitsBenefitArticlesJSON struct {
ID apijson.Field
CreatedAt apijson.Field
Deletable apijson.Field
Description apijson.Field
ModifiedAt apijson.Field
OrganizationID apijson.Field
Properties apijson.Field
Selectable apijson.Field
Type apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *UserOrderListResponseProductBenefitsBenefitArticles) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userOrderListResponseProductBenefitsBenefitArticlesJSON) RawJSON() (string) {
  return r.raw
}

func (r UserOrderListResponseProductBenefitsBenefitArticles) implementsUserOrderListResponseProductBenefit() {}

// Properties for a benefit of type `articles`.
type UserOrderListResponseProductBenefitsBenefitArticlesProperties struct {
// Whether the user can access paid articles.
PaidArticles bool `json:"paid_articles,required"`
JSON userOrderListResponseProductBenefitsBenefitArticlesPropertiesJSON `json:"-"`
}

// userOrderListResponseProductBenefitsBenefitArticlesPropertiesJSON contains the
// JSON metadata for the struct
// [UserOrderListResponseProductBenefitsBenefitArticlesProperties]
type userOrderListResponseProductBenefitsBenefitArticlesPropertiesJSON struct {
PaidArticles apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *UserOrderListResponseProductBenefitsBenefitArticlesProperties) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userOrderListResponseProductBenefitsBenefitArticlesPropertiesJSON) RawJSON() (string) {
  return r.raw
}

type UserOrderListResponseProductBenefitsBenefitArticlesType string

const (
  UserOrderListResponseProductBenefitsBenefitArticlesTypeArticles UserOrderListResponseProductBenefitsBenefitArticlesType = "articles"
)

func (r UserOrderListResponseProductBenefitsBenefitArticlesType) IsKnown() (bool) {
  switch r {
  case UserOrderListResponseProductBenefitsBenefitArticlesTypeArticles:
      return true
  }
  return false
}

// The type of the benefit.
type UserOrderListResponseProductBenefitsType string

const (
  UserOrderListResponseProductBenefitsTypeCustom UserOrderListResponseProductBenefitsType = "custom"
  UserOrderListResponseProductBenefitsTypeArticles UserOrderListResponseProductBenefitsType = "articles"
  UserOrderListResponseProductBenefitsTypeAds UserOrderListResponseProductBenefitsType = "ads"
  UserOrderListResponseProductBenefitsTypeDiscord UserOrderListResponseProductBenefitsType = "discord"
  UserOrderListResponseProductBenefitsTypeGitHubRepository UserOrderListResponseProductBenefitsType = "github_repository"
  UserOrderListResponseProductBenefitsTypeDownloadables UserOrderListResponseProductBenefitsType = "downloadables"
)

func (r UserOrderListResponseProductBenefitsType) IsKnown() (bool) {
  switch r {
  case UserOrderListResponseProductBenefitsTypeCustom, UserOrderListResponseProductBenefitsTypeArticles, UserOrderListResponseProductBenefitsTypeAds, UserOrderListResponseProductBenefitsTypeDiscord, UserOrderListResponseProductBenefitsTypeGitHubRepository, UserOrderListResponseProductBenefitsTypeDownloadables:
      return true
  }
  return false
}

// A recurring price for a product, i.e. a subscription.
type UserOrderListResponseProductPrice struct {
// Creation timestamp of the object.
CreatedAt time.Time `json:"created_at,required" format:"date-time"`
// Last modification timestamp of the object.
ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
// The ID of the price.
ID string `json:"id,required" format:"uuid4"`
// The price in cents.
PriceAmount int64 `json:"price_amount,required"`
// The currency.
PriceCurrency string `json:"price_currency,required"`
// Whether the price is archived and no longer available.
IsArchived bool `json:"is_archived,required"`
// The type of the price.
Type UserOrderListResponseProductPricesType `json:"type,required"`
// The recurring interval of the price, if type is `recurring`.
RecurringInterval UserOrderListResponseProductPricesRecurringInterval `json:"recurring_interval,nullable"`
JSON userOrderListResponseProductPriceJSON `json:"-"`
union UserOrderListResponseProductPricesUnion
}

// userOrderListResponseProductPriceJSON contains the JSON metadata for the struct
// [UserOrderListResponseProductPrice]
type userOrderListResponseProductPriceJSON struct {
CreatedAt apijson.Field
ModifiedAt apijson.Field
ID apijson.Field
PriceAmount apijson.Field
PriceCurrency apijson.Field
IsArchived apijson.Field
Type apijson.Field
RecurringInterval apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r userOrderListResponseProductPriceJSON) RawJSON() (string) {
  return r.raw
}

func (r *UserOrderListResponseProductPrice) UnmarshalJSON(data []byte) (err error) {
  *r = UserOrderListResponseProductPrice{}
  err = apijson.UnmarshalRoot(data, &r.union)
  if err != nil {
    return err
  }
  return apijson.Port(r.union, &r)
}

// AsUnion returns a [UserOrderListResponseProductPricesUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [UserOrderListResponseProductPricesProductPriceRecurring],
// [UserOrderListResponseProductPricesProductPriceOneTime].
func (r UserOrderListResponseProductPrice) AsUnion() (UserOrderListResponseProductPricesUnion) {
  return r.union
}

// A recurring price for a product, i.e. a subscription.
//
// Union satisfied by [UserOrderListResponseProductPricesProductPriceRecurring] or
// [UserOrderListResponseProductPricesProductPriceOneTime].
type UserOrderListResponseProductPricesUnion interface {
  implementsUserOrderListResponseProductPrice()
}

func init() {
  apijson.RegisterUnion(
    reflect.TypeOf((*UserOrderListResponseProductPricesUnion)(nil)).Elem(),
    "type",
    apijson.UnionVariant{
      TypeFilter: gjson.JSON,
      Type: reflect.TypeOf(UserOrderListResponseProductPricesProductPriceRecurring{}),
      DiscriminatorValue: "recurring",
    },
    apijson.UnionVariant{
      TypeFilter: gjson.JSON,
      Type: reflect.TypeOf(UserOrderListResponseProductPricesProductPriceOneTime{}),
      DiscriminatorValue: "one_time",
    },
  )
}

// A recurring price for a product, i.e. a subscription.
type UserOrderListResponseProductPricesProductPriceRecurring struct {
// The ID of the price.
ID string `json:"id,required" format:"uuid4"`
// Creation timestamp of the object.
CreatedAt time.Time `json:"created_at,required" format:"date-time"`
// Whether the price is archived and no longer available.
IsArchived bool `json:"is_archived,required"`
// Last modification timestamp of the object.
ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
// The price in cents.
PriceAmount int64 `json:"price_amount,required"`
// The currency.
PriceCurrency string `json:"price_currency,required"`
// The recurring interval of the price, if type is `recurring`.
RecurringInterval UserOrderListResponseProductPricesProductPriceRecurringRecurringInterval `json:"recurring_interval,required,nullable"`
// The type of the price.
Type UserOrderListResponseProductPricesProductPriceRecurringType `json:"type,required"`
JSON userOrderListResponseProductPricesProductPriceRecurringJSON `json:"-"`
}

// userOrderListResponseProductPricesProductPriceRecurringJSON contains the JSON
// metadata for the struct
// [UserOrderListResponseProductPricesProductPriceRecurring]
type userOrderListResponseProductPricesProductPriceRecurringJSON struct {
ID apijson.Field
CreatedAt apijson.Field
IsArchived apijson.Field
ModifiedAt apijson.Field
PriceAmount apijson.Field
PriceCurrency apijson.Field
RecurringInterval apijson.Field
Type apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *UserOrderListResponseProductPricesProductPriceRecurring) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userOrderListResponseProductPricesProductPriceRecurringJSON) RawJSON() (string) {
  return r.raw
}

func (r UserOrderListResponseProductPricesProductPriceRecurring) implementsUserOrderListResponseProductPrice() {}

// The recurring interval of the price, if type is `recurring`.
type UserOrderListResponseProductPricesProductPriceRecurringRecurringInterval string

const (
  UserOrderListResponseProductPricesProductPriceRecurringRecurringIntervalMonth UserOrderListResponseProductPricesProductPriceRecurringRecurringInterval = "month"
  UserOrderListResponseProductPricesProductPriceRecurringRecurringIntervalYear UserOrderListResponseProductPricesProductPriceRecurringRecurringInterval = "year"
)

func (r UserOrderListResponseProductPricesProductPriceRecurringRecurringInterval) IsKnown() (bool) {
  switch r {
  case UserOrderListResponseProductPricesProductPriceRecurringRecurringIntervalMonth, UserOrderListResponseProductPricesProductPriceRecurringRecurringIntervalYear:
      return true
  }
  return false
}

// The type of the price.
type UserOrderListResponseProductPricesProductPriceRecurringType string

const (
  UserOrderListResponseProductPricesProductPriceRecurringTypeRecurring UserOrderListResponseProductPricesProductPriceRecurringType = "recurring"
)

func (r UserOrderListResponseProductPricesProductPriceRecurringType) IsKnown() (bool) {
  switch r {
  case UserOrderListResponseProductPricesProductPriceRecurringTypeRecurring:
      return true
  }
  return false
}

// A one-time price for a product.
type UserOrderListResponseProductPricesProductPriceOneTime struct {
// The ID of the price.
ID string `json:"id,required" format:"uuid4"`
// Creation timestamp of the object.
CreatedAt time.Time `json:"created_at,required" format:"date-time"`
// Whether the price is archived and no longer available.
IsArchived bool `json:"is_archived,required"`
// Last modification timestamp of the object.
ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
// The price in cents.
PriceAmount int64 `json:"price_amount,required"`
// The currency.
PriceCurrency string `json:"price_currency,required"`
// The type of the price.
Type UserOrderListResponseProductPricesProductPriceOneTimeType `json:"type,required"`
JSON userOrderListResponseProductPricesProductPriceOneTimeJSON `json:"-"`
}

// userOrderListResponseProductPricesProductPriceOneTimeJSON contains the JSON
// metadata for the struct [UserOrderListResponseProductPricesProductPriceOneTime]
type userOrderListResponseProductPricesProductPriceOneTimeJSON struct {
ID apijson.Field
CreatedAt apijson.Field
IsArchived apijson.Field
ModifiedAt apijson.Field
PriceAmount apijson.Field
PriceCurrency apijson.Field
Type apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *UserOrderListResponseProductPricesProductPriceOneTime) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userOrderListResponseProductPricesProductPriceOneTimeJSON) RawJSON() (string) {
  return r.raw
}

func (r UserOrderListResponseProductPricesProductPriceOneTime) implementsUserOrderListResponseProductPrice() {}

// The type of the price.
type UserOrderListResponseProductPricesProductPriceOneTimeType string

const (
  UserOrderListResponseProductPricesProductPriceOneTimeTypeOneTime UserOrderListResponseProductPricesProductPriceOneTimeType = "one_time"
)

func (r UserOrderListResponseProductPricesProductPriceOneTimeType) IsKnown() (bool) {
  switch r {
  case UserOrderListResponseProductPricesProductPriceOneTimeTypeOneTime:
      return true
  }
  return false
}

// The type of the price.
type UserOrderListResponseProductPricesType string

const (
  UserOrderListResponseProductPricesTypeRecurring UserOrderListResponseProductPricesType = "recurring"
  UserOrderListResponseProductPricesTypeOneTime UserOrderListResponseProductPricesType = "one_time"
)

func (r UserOrderListResponseProductPricesType) IsKnown() (bool) {
  switch r {
  case UserOrderListResponseProductPricesTypeRecurring, UserOrderListResponseProductPricesTypeOneTime:
      return true
  }
  return false
}

// The recurring interval of the price, if type is `recurring`.
type UserOrderListResponseProductPricesRecurringInterval string

const (
  UserOrderListResponseProductPricesRecurringIntervalMonth UserOrderListResponseProductPricesRecurringInterval = "month"
  UserOrderListResponseProductPricesRecurringIntervalYear UserOrderListResponseProductPricesRecurringInterval = "year"
)

func (r UserOrderListResponseProductPricesRecurringInterval) IsKnown() (bool) {
  switch r {
  case UserOrderListResponseProductPricesRecurringIntervalMonth, UserOrderListResponseProductPricesRecurringIntervalYear:
      return true
  }
  return false
}

type UserOrderListResponseProductType string

const (
  UserOrderListResponseProductTypeFree UserOrderListResponseProductType = "free"
  UserOrderListResponseProductTypeIndividual UserOrderListResponseProductType = "individual"
  UserOrderListResponseProductTypeBusiness UserOrderListResponseProductType = "business"
)

func (r UserOrderListResponseProductType) IsKnown() (bool) {
  switch r {
  case UserOrderListResponseProductTypeFree, UserOrderListResponseProductTypeIndividual, UserOrderListResponseProductTypeBusiness:
      return true
  }
  return false
}

type UserOrderListResponseSubscription struct {
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
Status UserOrderListResponseSubscriptionStatus `json:"status,required"`
UserID string `json:"user_id,required" format:"uuid4"`
JSON userOrderListResponseSubscriptionJSON `json:"-"`
}

// userOrderListResponseSubscriptionJSON contains the JSON metadata for the struct
// [UserOrderListResponseSubscription]
type userOrderListResponseSubscriptionJSON struct {
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

func (r *UserOrderListResponseSubscription) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userOrderListResponseSubscriptionJSON) RawJSON() (string) {
  return r.raw
}

type UserOrderListResponseSubscriptionStatus string

const (
  UserOrderListResponseSubscriptionStatusIncomplete UserOrderListResponseSubscriptionStatus = "incomplete"
  UserOrderListResponseSubscriptionStatusIncompleteExpired UserOrderListResponseSubscriptionStatus = "incomplete_expired"
  UserOrderListResponseSubscriptionStatusTrialing UserOrderListResponseSubscriptionStatus = "trialing"
  UserOrderListResponseSubscriptionStatusActive UserOrderListResponseSubscriptionStatus = "active"
  UserOrderListResponseSubscriptionStatusPastDue UserOrderListResponseSubscriptionStatus = "past_due"
  UserOrderListResponseSubscriptionStatusCanceled UserOrderListResponseSubscriptionStatus = "canceled"
  UserOrderListResponseSubscriptionStatusUnpaid UserOrderListResponseSubscriptionStatus = "unpaid"
)

func (r UserOrderListResponseSubscriptionStatus) IsKnown() (bool) {
  switch r {
  case UserOrderListResponseSubscriptionStatusIncomplete, UserOrderListResponseSubscriptionStatusIncompleteExpired, UserOrderListResponseSubscriptionStatusTrialing, UserOrderListResponseSubscriptionStatusActive, UserOrderListResponseSubscriptionStatusPastDue, UserOrderListResponseSubscriptionStatusCanceled, UserOrderListResponseSubscriptionStatusUnpaid:
      return true
  }
  return false
}

// Order's invoice data.
type UserOrderInvoiceResponse struct {
// The URL to the invoice.
URL string `json:"url,required"`
JSON userOrderInvoiceResponseJSON `json:"-"`
}

// userOrderInvoiceResponseJSON contains the JSON metadata for the struct
// [UserOrderInvoiceResponse]
type userOrderInvoiceResponseJSON struct {
URL apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *UserOrderInvoiceResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userOrderInvoiceResponseJSON) RawJSON() (string) {
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
Sorting param.Field[[]UserOrderListParamsSorting] `query:"sorting"`
// Filter by subscription ID.
SubscriptionID param.Field[UserOrderListParamsSubscriptionIDUnion] `query:"subscription_id" format:"uuid4"`
}

// URLQuery serializes [UserOrderListParams]'s query parameters as `url.Values`.
func (r UserOrderListParams) URLQuery() (v url.Values) {
  return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
    ArrayFormat: apiquery.ArrayQueryFormatComma,
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
  UserOrderListParamsProductPriceTypeProductPriceTypeOneTime UserOrderListParamsProductPriceTypeProductPriceType = "one_time"
  UserOrderListParamsProductPriceTypeProductPriceTypeRecurring UserOrderListParamsProductPriceTypeProductPriceType = "recurring"
)

func (r UserOrderListParamsProductPriceTypeProductPriceType) IsKnown() (bool) {
  switch r {
  case UserOrderListParamsProductPriceTypeProductPriceTypeOneTime, UserOrderListParamsProductPriceTypeProductPriceTypeRecurring:
      return true
  }
  return false
}

func (r UserOrderListParamsProductPriceTypeProductPriceType) implementsUserOrderListParamsProductPriceTypeUnion() {}

type UserOrderListParamsProductPriceTypeArray []UserOrderListParamsProductPriceTypeArray

func (r UserOrderListParamsProductPriceTypeArray) implementsUserOrderListParamsProductPriceTypeUnion() {}

type UserOrderListParamsSorting string

const (
  UserOrderListParamsSortingCreatedAt UserOrderListParamsSorting = "created_at"
  UserOrderListParamsSorting-CreatedAt UserOrderListParamsSorting = "-created_at"
  UserOrderListParamsSortingAmount UserOrderListParamsSorting = "amount"
  UserOrderListParamsSorting-Amount UserOrderListParamsSorting = "-amount"
  UserOrderListParamsSortingOrganization UserOrderListParamsSorting = "organization"
  UserOrderListParamsSorting-Organization UserOrderListParamsSorting = "-organization"
  UserOrderListParamsSortingProduct UserOrderListParamsSorting = "product"
  UserOrderListParamsSorting-Product UserOrderListParamsSorting = "-product"
  UserOrderListParamsSortingSubscription UserOrderListParamsSorting = "subscription"
  UserOrderListParamsSorting-Subscription UserOrderListParamsSorting = "-subscription"
)

func (r UserOrderListParamsSorting) IsKnown() (bool) {
  switch r {
  case UserOrderListParamsSortingCreatedAt, UserOrderListParamsSorting-CreatedAt, UserOrderListParamsSortingAmount, UserOrderListParamsSorting-Amount, UserOrderListParamsSortingOrganization, UserOrderListParamsSorting-Organization, UserOrderListParamsSortingProduct, UserOrderListParamsSorting-Product, UserOrderListParamsSortingSubscription, UserOrderListParamsSorting-Subscription:
      return true
  }
  return false
}

// Filter by subscription ID.
//
// Satisfied by [shared.UnionString], [UserOrderListParamsSubscriptionIDArray].
type UserOrderListParamsSubscriptionIDUnion interface {
  ImplementsUserOrderListParamsSubscriptionIDUnion()
}

type UserOrderListParamsSubscriptionIDArray []string

func (r UserOrderListParamsSubscriptionIDArray) ImplementsUserOrderListParamsSubscriptionIDUnion() {}
