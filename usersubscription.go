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

// UserSubscriptionService contains methods and other services that help with
// interacting with the polar API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserSubscriptionService] method instead.
type UserSubscriptionService struct {
Options []option.RequestOption
}

// NewUserSubscriptionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewUserSubscriptionService(opts ...option.RequestOption) (r *UserSubscriptionService) {
  r = &UserSubscriptionService{}
  r.Options = opts
  return
}

// Create a subscription on a **free** tier.
//
// If you want to subscribe to a paid tier, you need to create a checkout session.
func (r *UserSubscriptionService) New(ctx context.Context, body UserSubscriptionNewParams, opts ...option.RequestOption) (res *UserSubscriptionNewResponse, err error) {
  opts = append(r.Options[:], opts...)
  path := "v1/users/subscriptions/"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
  return
}

// Get a subscription by ID.
func (r *UserSubscriptionService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *UserSubscriptionGetResponse, err error) {
  opts = append(r.Options[:], opts...)
  if id == "" {
    err = errors.New("missing required id parameter")
    return
  }
  path := fmt.Sprintf("v1/users/subscriptions/%s", id)
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
  return
}

// Update a subscription.
func (r *UserSubscriptionService) Update(ctx context.Context, id string, body UserSubscriptionUpdateParams, opts ...option.RequestOption) (res *UserSubscriptionUpdateResponse, err error) {
  opts = append(r.Options[:], opts...)
  if id == "" {
    err = errors.New("missing required id parameter")
    return
  }
  path := fmt.Sprintf("v1/users/subscriptions/%s", id)
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
  return
}

// List my subscriptions.
func (r *UserSubscriptionService) List(ctx context.Context, query UserSubscriptionListParams, opts ...option.RequestOption) (res *pagination.PolarPagination[UserSubscriptionListResponse], err error) {
  var raw *http.Response
  opts = append(r.Options[:], opts...)
  opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
  path := "v1/users/subscriptions/"
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

// List my subscriptions.
func (r *UserSubscriptionService) ListAutoPaging(ctx context.Context, query UserSubscriptionListParams, opts ...option.RequestOption) (*pagination.PolarPaginationAutoPager[UserSubscriptionListResponse]) {
  return pagination.NewPolarPaginationAutoPager(r.List(ctx, query, opts...))
}

// Cancel a subscription.
func (r *UserSubscriptionService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *UserSubscriptionDeleteResponse, err error) {
  opts = append(r.Options[:], opts...)
  if id == "" {
    err = errors.New("missing required id parameter")
    return
  }
  path := fmt.Sprintf("v1/users/subscriptions/%s", id)
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
  return
}

type UserSubscriptionNewResponse struct {
ID string `json:"id,required" format:"uuid4"`
CancelAtPeriodEnd bool `json:"cancel_at_period_end,required"`
// Creation timestamp of the object.
CreatedAt time.Time `json:"created_at,required" format:"date-time"`
CurrentPeriodEnd time.Time `json:"current_period_end,required,nullable" format:"date-time"`
CurrentPeriodStart time.Time `json:"current_period_start,required" format:"date-time"`
EndedAt time.Time `json:"ended_at,required,nullable" format:"date-time"`
// Last modification timestamp of the object.
ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
// A recurring price for a product, i.e. a subscription.
Price UserSubscriptionNewResponsePrice `json:"price,required,nullable"`
PriceID string `json:"price_id,required,nullable" format:"uuid4"`
Product UserSubscriptionNewResponseProduct `json:"product,required"`
ProductID string `json:"product_id,required" format:"uuid4"`
StartedAt time.Time `json:"started_at,required,nullable" format:"date-time"`
Status UserSubscriptionNewResponseStatus `json:"status,required"`
UserID string `json:"user_id,required" format:"uuid4"`
JSON userSubscriptionNewResponseJSON `json:"-"`
}

// userSubscriptionNewResponseJSON contains the JSON metadata for the struct
// [UserSubscriptionNewResponse]
type userSubscriptionNewResponseJSON struct {
ID apijson.Field
CancelAtPeriodEnd apijson.Field
CreatedAt apijson.Field
CurrentPeriodEnd apijson.Field
CurrentPeriodStart apijson.Field
EndedAt apijson.Field
ModifiedAt apijson.Field
Price apijson.Field
PriceID apijson.Field
Product apijson.Field
ProductID apijson.Field
StartedAt apijson.Field
Status apijson.Field
UserID apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionNewResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userSubscriptionNewResponseJSON) RawJSON() (string) {
  return r.raw
}

// A recurring price for a product, i.e. a subscription.
type UserSubscriptionNewResponsePrice struct {
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
Type UserSubscriptionNewResponsePriceType `json:"type,required"`
// The recurring interval of the price, if type is `recurring`.
RecurringInterval UserSubscriptionNewResponsePriceRecurringInterval `json:"recurring_interval,nullable"`
JSON userSubscriptionNewResponsePriceJSON `json:"-"`
union UserSubscriptionNewResponsePriceUnion
}

// userSubscriptionNewResponsePriceJSON contains the JSON metadata for the struct
// [UserSubscriptionNewResponsePrice]
type userSubscriptionNewResponsePriceJSON struct {
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

func (r userSubscriptionNewResponsePriceJSON) RawJSON() (string) {
  return r.raw
}

func (r *UserSubscriptionNewResponsePrice) UnmarshalJSON(data []byte) (err error) {
  *r = UserSubscriptionNewResponsePrice{}
  err = apijson.UnmarshalRoot(data, &r.union)
  if err != nil {
    return err
  }
  return apijson.Port(r.union, &r)
}

// AsUnion returns a [UserSubscriptionNewResponsePriceUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [UserSubscriptionNewResponsePriceProductPriceRecurring],
// [UserSubscriptionNewResponsePriceProductPriceOneTime].
func (r UserSubscriptionNewResponsePrice) AsUnion() (UserSubscriptionNewResponsePriceUnion) {
  return r.union
}

// A recurring price for a product, i.e. a subscription.
//
// Union satisfied by [UserSubscriptionNewResponsePriceProductPriceRecurring] or
// [UserSubscriptionNewResponsePriceProductPriceOneTime].
type UserSubscriptionNewResponsePriceUnion interface {
  implementsUserSubscriptionNewResponsePrice()
}

func init() {
  apijson.RegisterUnion(
    reflect.TypeOf((*UserSubscriptionNewResponsePriceUnion)(nil)).Elem(),
    "type",
    apijson.UnionVariant{
      TypeFilter: gjson.JSON,
      Type: reflect.TypeOf(UserSubscriptionNewResponsePriceProductPriceRecurring{}),
      DiscriminatorValue: "recurring",
    },
    apijson.UnionVariant{
      TypeFilter: gjson.JSON,
      Type: reflect.TypeOf(UserSubscriptionNewResponsePriceProductPriceOneTime{}),
      DiscriminatorValue: "one_time",
    },
  )
}

// A recurring price for a product, i.e. a subscription.
type UserSubscriptionNewResponsePriceProductPriceRecurring struct {
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
RecurringInterval UserSubscriptionNewResponsePriceProductPriceRecurringRecurringInterval `json:"recurring_interval,required,nullable"`
// The type of the price.
Type UserSubscriptionNewResponsePriceProductPriceRecurringType `json:"type,required"`
JSON userSubscriptionNewResponsePriceProductPriceRecurringJSON `json:"-"`
}

// userSubscriptionNewResponsePriceProductPriceRecurringJSON contains the JSON
// metadata for the struct [UserSubscriptionNewResponsePriceProductPriceRecurring]
type userSubscriptionNewResponsePriceProductPriceRecurringJSON struct {
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

func (r *UserSubscriptionNewResponsePriceProductPriceRecurring) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userSubscriptionNewResponsePriceProductPriceRecurringJSON) RawJSON() (string) {
  return r.raw
}

func (r UserSubscriptionNewResponsePriceProductPriceRecurring) implementsUserSubscriptionNewResponsePrice() {}

// The recurring interval of the price, if type is `recurring`.
type UserSubscriptionNewResponsePriceProductPriceRecurringRecurringInterval string

const (
  UserSubscriptionNewResponsePriceProductPriceRecurringRecurringIntervalMonth UserSubscriptionNewResponsePriceProductPriceRecurringRecurringInterval = "month"
  UserSubscriptionNewResponsePriceProductPriceRecurringRecurringIntervalYear UserSubscriptionNewResponsePriceProductPriceRecurringRecurringInterval = "year"
)

func (r UserSubscriptionNewResponsePriceProductPriceRecurringRecurringInterval) IsKnown() (bool) {
  switch r {
  case UserSubscriptionNewResponsePriceProductPriceRecurringRecurringIntervalMonth, UserSubscriptionNewResponsePriceProductPriceRecurringRecurringIntervalYear:
      return true
  }
  return false
}

// The type of the price.
type UserSubscriptionNewResponsePriceProductPriceRecurringType string

const (
  UserSubscriptionNewResponsePriceProductPriceRecurringTypeRecurring UserSubscriptionNewResponsePriceProductPriceRecurringType = "recurring"
)

func (r UserSubscriptionNewResponsePriceProductPriceRecurringType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionNewResponsePriceProductPriceRecurringTypeRecurring:
      return true
  }
  return false
}

// A one-time price for a product.
type UserSubscriptionNewResponsePriceProductPriceOneTime struct {
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
Type UserSubscriptionNewResponsePriceProductPriceOneTimeType `json:"type,required"`
JSON userSubscriptionNewResponsePriceProductPriceOneTimeJSON `json:"-"`
}

// userSubscriptionNewResponsePriceProductPriceOneTimeJSON contains the JSON
// metadata for the struct [UserSubscriptionNewResponsePriceProductPriceOneTime]
type userSubscriptionNewResponsePriceProductPriceOneTimeJSON struct {
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

func (r *UserSubscriptionNewResponsePriceProductPriceOneTime) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userSubscriptionNewResponsePriceProductPriceOneTimeJSON) RawJSON() (string) {
  return r.raw
}

func (r UserSubscriptionNewResponsePriceProductPriceOneTime) implementsUserSubscriptionNewResponsePrice() {}

// The type of the price.
type UserSubscriptionNewResponsePriceProductPriceOneTimeType string

const (
  UserSubscriptionNewResponsePriceProductPriceOneTimeTypeOneTime UserSubscriptionNewResponsePriceProductPriceOneTimeType = "one_time"
)

func (r UserSubscriptionNewResponsePriceProductPriceOneTimeType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionNewResponsePriceProductPriceOneTimeTypeOneTime:
      return true
  }
  return false
}

// The type of the price.
type UserSubscriptionNewResponsePriceType string

const (
  UserSubscriptionNewResponsePriceTypeRecurring UserSubscriptionNewResponsePriceType = "recurring"
  UserSubscriptionNewResponsePriceTypeOneTime UserSubscriptionNewResponsePriceType = "one_time"
)

func (r UserSubscriptionNewResponsePriceType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionNewResponsePriceTypeRecurring, UserSubscriptionNewResponsePriceTypeOneTime:
      return true
  }
  return false
}

// The recurring interval of the price, if type is `recurring`.
type UserSubscriptionNewResponsePriceRecurringInterval string

const (
  UserSubscriptionNewResponsePriceRecurringIntervalMonth UserSubscriptionNewResponsePriceRecurringInterval = "month"
  UserSubscriptionNewResponsePriceRecurringIntervalYear UserSubscriptionNewResponsePriceRecurringInterval = "year"
)

func (r UserSubscriptionNewResponsePriceRecurringInterval) IsKnown() (bool) {
  switch r {
  case UserSubscriptionNewResponsePriceRecurringIntervalMonth, UserSubscriptionNewResponsePriceRecurringIntervalYear:
      return true
  }
  return false
}

type UserSubscriptionNewResponseProduct struct {
// The ID of the product.
ID string `json:"id,required" format:"uuid4"`
// The benefits granted by the product.
Benefits []UserSubscriptionNewResponseProductBenefit `json:"benefits,required"`
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
Prices []UserSubscriptionNewResponseProductPrice `json:"prices,required"`
Type UserSubscriptionNewResponseProductType `json:"type,required,nullable"`
JSON userSubscriptionNewResponseProductJSON `json:"-"`
}

// userSubscriptionNewResponseProductJSON contains the JSON metadata for the struct
// [UserSubscriptionNewResponseProduct]
type userSubscriptionNewResponseProductJSON struct {
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

func (r *UserSubscriptionNewResponseProduct) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userSubscriptionNewResponseProductJSON) RawJSON() (string) {
  return r.raw
}

// A benefit of type `articles`.
//
// Use it to grant access to posts.
type UserSubscriptionNewResponseProductBenefit struct {
// Creation timestamp of the object.
CreatedAt time.Time `json:"created_at,required" format:"date-time"`
// Last modification timestamp of the object.
ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
// The ID of the benefit.
ID string `json:"id,required" format:"uuid4"`
// The type of the benefit.
Type UserSubscriptionNewResponseProductBenefitsType `json:"type,required"`
// The description of the benefit.
Description string `json:"description,required"`
// Whether the benefit is selectable when creating a product.
Selectable bool `json:"selectable,required"`
// Whether the benefit is deletable.
Deletable bool `json:"deletable,required"`
// The ID of the organization owning the benefit.
OrganizationID string `json:"organization_id,required" format:"uuid4"`
// This field can have the runtime type of
// [UserSubscriptionNewResponseProductBenefitsBenefitArticlesProperties].
Properties interface{} `json:"properties,required"`
JSON userSubscriptionNewResponseProductBenefitJSON `json:"-"`
union UserSubscriptionNewResponseProductBenefitsUnion
}

// userSubscriptionNewResponseProductBenefitJSON contains the JSON metadata for the
// struct [UserSubscriptionNewResponseProductBenefit]
type userSubscriptionNewResponseProductBenefitJSON struct {
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

func (r userSubscriptionNewResponseProductBenefitJSON) RawJSON() (string) {
  return r.raw
}

func (r *UserSubscriptionNewResponseProductBenefit) UnmarshalJSON(data []byte) (err error) {
  *r = UserSubscriptionNewResponseProductBenefit{}
  err = apijson.UnmarshalRoot(data, &r.union)
  if err != nil {
    return err
  }
  return apijson.Port(r.union, &r)
}

// AsUnion returns a [UserSubscriptionNewResponseProductBenefitsUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [UserSubscriptionNewResponseProductBenefitsBenefitBase],
// [UserSubscriptionNewResponseProductBenefitsBenefitArticles].
func (r UserSubscriptionNewResponseProductBenefit) AsUnion() (UserSubscriptionNewResponseProductBenefitsUnion) {
  return r.union
}

// A benefit of type `articles`.
//
// Use it to grant access to posts.
//
// Union satisfied by [UserSubscriptionNewResponseProductBenefitsBenefitBase] or
// [UserSubscriptionNewResponseProductBenefitsBenefitArticles].
type UserSubscriptionNewResponseProductBenefitsUnion interface {
  implementsUserSubscriptionNewResponseProductBenefit()
}

func init() {
  apijson.RegisterUnion(
    reflect.TypeOf((*UserSubscriptionNewResponseProductBenefitsUnion)(nil)).Elem(),
    "",
    apijson.UnionVariant{
      TypeFilter: gjson.JSON,
      Type: reflect.TypeOf(UserSubscriptionNewResponseProductBenefitsBenefitBase{}),
    },
    apijson.UnionVariant{
      TypeFilter: gjson.JSON,
      Type: reflect.TypeOf(UserSubscriptionNewResponseProductBenefitsBenefitArticles{}),
    },
  )
}

type UserSubscriptionNewResponseProductBenefitsBenefitBase struct {
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
Type UserSubscriptionNewResponseProductBenefitsBenefitBaseType `json:"type,required"`
JSON userSubscriptionNewResponseProductBenefitsBenefitBaseJSON `json:"-"`
}

// userSubscriptionNewResponseProductBenefitsBenefitBaseJSON contains the JSON
// metadata for the struct [UserSubscriptionNewResponseProductBenefitsBenefitBase]
type userSubscriptionNewResponseProductBenefitsBenefitBaseJSON struct {
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

func (r *UserSubscriptionNewResponseProductBenefitsBenefitBase) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userSubscriptionNewResponseProductBenefitsBenefitBaseJSON) RawJSON() (string) {
  return r.raw
}

func (r UserSubscriptionNewResponseProductBenefitsBenefitBase) implementsUserSubscriptionNewResponseProductBenefit() {}

// The type of the benefit.
type UserSubscriptionNewResponseProductBenefitsBenefitBaseType string

const (
  UserSubscriptionNewResponseProductBenefitsBenefitBaseTypeCustom UserSubscriptionNewResponseProductBenefitsBenefitBaseType = "custom"
  UserSubscriptionNewResponseProductBenefitsBenefitBaseTypeArticles UserSubscriptionNewResponseProductBenefitsBenefitBaseType = "articles"
  UserSubscriptionNewResponseProductBenefitsBenefitBaseTypeAds UserSubscriptionNewResponseProductBenefitsBenefitBaseType = "ads"
  UserSubscriptionNewResponseProductBenefitsBenefitBaseTypeDiscord UserSubscriptionNewResponseProductBenefitsBenefitBaseType = "discord"
  UserSubscriptionNewResponseProductBenefitsBenefitBaseTypeGitHubRepository UserSubscriptionNewResponseProductBenefitsBenefitBaseType = "github_repository"
  UserSubscriptionNewResponseProductBenefitsBenefitBaseTypeDownloadables UserSubscriptionNewResponseProductBenefitsBenefitBaseType = "downloadables"
)

func (r UserSubscriptionNewResponseProductBenefitsBenefitBaseType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionNewResponseProductBenefitsBenefitBaseTypeCustom, UserSubscriptionNewResponseProductBenefitsBenefitBaseTypeArticles, UserSubscriptionNewResponseProductBenefitsBenefitBaseTypeAds, UserSubscriptionNewResponseProductBenefitsBenefitBaseTypeDiscord, UserSubscriptionNewResponseProductBenefitsBenefitBaseTypeGitHubRepository, UserSubscriptionNewResponseProductBenefitsBenefitBaseTypeDownloadables:
      return true
  }
  return false
}

// A benefit of type `articles`.
//
// Use it to grant access to posts.
type UserSubscriptionNewResponseProductBenefitsBenefitArticles struct {
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
Properties UserSubscriptionNewResponseProductBenefitsBenefitArticlesProperties `json:"properties,required"`
// Whether the benefit is selectable when creating a product.
Selectable bool `json:"selectable,required"`
Type UserSubscriptionNewResponseProductBenefitsBenefitArticlesType `json:"type,required"`
JSON userSubscriptionNewResponseProductBenefitsBenefitArticlesJSON `json:"-"`
}

// userSubscriptionNewResponseProductBenefitsBenefitArticlesJSON contains the JSON
// metadata for the struct
// [UserSubscriptionNewResponseProductBenefitsBenefitArticles]
type userSubscriptionNewResponseProductBenefitsBenefitArticlesJSON struct {
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

func (r *UserSubscriptionNewResponseProductBenefitsBenefitArticles) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userSubscriptionNewResponseProductBenefitsBenefitArticlesJSON) RawJSON() (string) {
  return r.raw
}

func (r UserSubscriptionNewResponseProductBenefitsBenefitArticles) implementsUserSubscriptionNewResponseProductBenefit() {}

// Properties for a benefit of type `articles`.
type UserSubscriptionNewResponseProductBenefitsBenefitArticlesProperties struct {
// Whether the user can access paid articles.
PaidArticles bool `json:"paid_articles,required"`
JSON userSubscriptionNewResponseProductBenefitsBenefitArticlesPropertiesJSON `json:"-"`
}

// userSubscriptionNewResponseProductBenefitsBenefitArticlesPropertiesJSON contains
// the JSON metadata for the struct
// [UserSubscriptionNewResponseProductBenefitsBenefitArticlesProperties]
type userSubscriptionNewResponseProductBenefitsBenefitArticlesPropertiesJSON struct {
PaidArticles apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionNewResponseProductBenefitsBenefitArticlesProperties) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userSubscriptionNewResponseProductBenefitsBenefitArticlesPropertiesJSON) RawJSON() (string) {
  return r.raw
}

type UserSubscriptionNewResponseProductBenefitsBenefitArticlesType string

const (
  UserSubscriptionNewResponseProductBenefitsBenefitArticlesTypeArticles UserSubscriptionNewResponseProductBenefitsBenefitArticlesType = "articles"
)

func (r UserSubscriptionNewResponseProductBenefitsBenefitArticlesType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionNewResponseProductBenefitsBenefitArticlesTypeArticles:
      return true
  }
  return false
}

// The type of the benefit.
type UserSubscriptionNewResponseProductBenefitsType string

const (
  UserSubscriptionNewResponseProductBenefitsTypeCustom UserSubscriptionNewResponseProductBenefitsType = "custom"
  UserSubscriptionNewResponseProductBenefitsTypeArticles UserSubscriptionNewResponseProductBenefitsType = "articles"
  UserSubscriptionNewResponseProductBenefitsTypeAds UserSubscriptionNewResponseProductBenefitsType = "ads"
  UserSubscriptionNewResponseProductBenefitsTypeDiscord UserSubscriptionNewResponseProductBenefitsType = "discord"
  UserSubscriptionNewResponseProductBenefitsTypeGitHubRepository UserSubscriptionNewResponseProductBenefitsType = "github_repository"
  UserSubscriptionNewResponseProductBenefitsTypeDownloadables UserSubscriptionNewResponseProductBenefitsType = "downloadables"
)

func (r UserSubscriptionNewResponseProductBenefitsType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionNewResponseProductBenefitsTypeCustom, UserSubscriptionNewResponseProductBenefitsTypeArticles, UserSubscriptionNewResponseProductBenefitsTypeAds, UserSubscriptionNewResponseProductBenefitsTypeDiscord, UserSubscriptionNewResponseProductBenefitsTypeGitHubRepository, UserSubscriptionNewResponseProductBenefitsTypeDownloadables:
      return true
  }
  return false
}

// A recurring price for a product, i.e. a subscription.
type UserSubscriptionNewResponseProductPrice struct {
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
Type UserSubscriptionNewResponseProductPricesType `json:"type,required"`
// The recurring interval of the price, if type is `recurring`.
RecurringInterval UserSubscriptionNewResponseProductPricesRecurringInterval `json:"recurring_interval,nullable"`
JSON userSubscriptionNewResponseProductPriceJSON `json:"-"`
union UserSubscriptionNewResponseProductPricesUnion
}

// userSubscriptionNewResponseProductPriceJSON contains the JSON metadata for the
// struct [UserSubscriptionNewResponseProductPrice]
type userSubscriptionNewResponseProductPriceJSON struct {
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

func (r userSubscriptionNewResponseProductPriceJSON) RawJSON() (string) {
  return r.raw
}

func (r *UserSubscriptionNewResponseProductPrice) UnmarshalJSON(data []byte) (err error) {
  *r = UserSubscriptionNewResponseProductPrice{}
  err = apijson.UnmarshalRoot(data, &r.union)
  if err != nil {
    return err
  }
  return apijson.Port(r.union, &r)
}

// AsUnion returns a [UserSubscriptionNewResponseProductPricesUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [UserSubscriptionNewResponseProductPricesProductPriceRecurring],
// [UserSubscriptionNewResponseProductPricesProductPriceOneTime].
func (r UserSubscriptionNewResponseProductPrice) AsUnion() (UserSubscriptionNewResponseProductPricesUnion) {
  return r.union
}

// A recurring price for a product, i.e. a subscription.
//
// Union satisfied by
// [UserSubscriptionNewResponseProductPricesProductPriceRecurring] or
// [UserSubscriptionNewResponseProductPricesProductPriceOneTime].
type UserSubscriptionNewResponseProductPricesUnion interface {
  implementsUserSubscriptionNewResponseProductPrice()
}

func init() {
  apijson.RegisterUnion(
    reflect.TypeOf((*UserSubscriptionNewResponseProductPricesUnion)(nil)).Elem(),
    "type",
    apijson.UnionVariant{
      TypeFilter: gjson.JSON,
      Type: reflect.TypeOf(UserSubscriptionNewResponseProductPricesProductPriceRecurring{}),
      DiscriminatorValue: "recurring",
    },
    apijson.UnionVariant{
      TypeFilter: gjson.JSON,
      Type: reflect.TypeOf(UserSubscriptionNewResponseProductPricesProductPriceOneTime{}),
      DiscriminatorValue: "one_time",
    },
  )
}

// A recurring price for a product, i.e. a subscription.
type UserSubscriptionNewResponseProductPricesProductPriceRecurring struct {
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
RecurringInterval UserSubscriptionNewResponseProductPricesProductPriceRecurringRecurringInterval `json:"recurring_interval,required,nullable"`
// The type of the price.
Type UserSubscriptionNewResponseProductPricesProductPriceRecurringType `json:"type,required"`
JSON userSubscriptionNewResponseProductPricesProductPriceRecurringJSON `json:"-"`
}

// userSubscriptionNewResponseProductPricesProductPriceRecurringJSON contains the
// JSON metadata for the struct
// [UserSubscriptionNewResponseProductPricesProductPriceRecurring]
type userSubscriptionNewResponseProductPricesProductPriceRecurringJSON struct {
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

func (r *UserSubscriptionNewResponseProductPricesProductPriceRecurring) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userSubscriptionNewResponseProductPricesProductPriceRecurringJSON) RawJSON() (string) {
  return r.raw
}

func (r UserSubscriptionNewResponseProductPricesProductPriceRecurring) implementsUserSubscriptionNewResponseProductPrice() {}

// The recurring interval of the price, if type is `recurring`.
type UserSubscriptionNewResponseProductPricesProductPriceRecurringRecurringInterval string

const (
  UserSubscriptionNewResponseProductPricesProductPriceRecurringRecurringIntervalMonth UserSubscriptionNewResponseProductPricesProductPriceRecurringRecurringInterval = "month"
  UserSubscriptionNewResponseProductPricesProductPriceRecurringRecurringIntervalYear UserSubscriptionNewResponseProductPricesProductPriceRecurringRecurringInterval = "year"
)

func (r UserSubscriptionNewResponseProductPricesProductPriceRecurringRecurringInterval) IsKnown() (bool) {
  switch r {
  case UserSubscriptionNewResponseProductPricesProductPriceRecurringRecurringIntervalMonth, UserSubscriptionNewResponseProductPricesProductPriceRecurringRecurringIntervalYear:
      return true
  }
  return false
}

// The type of the price.
type UserSubscriptionNewResponseProductPricesProductPriceRecurringType string

const (
  UserSubscriptionNewResponseProductPricesProductPriceRecurringTypeRecurring UserSubscriptionNewResponseProductPricesProductPriceRecurringType = "recurring"
)

func (r UserSubscriptionNewResponseProductPricesProductPriceRecurringType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionNewResponseProductPricesProductPriceRecurringTypeRecurring:
      return true
  }
  return false
}

// A one-time price for a product.
type UserSubscriptionNewResponseProductPricesProductPriceOneTime struct {
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
Type UserSubscriptionNewResponseProductPricesProductPriceOneTimeType `json:"type,required"`
JSON userSubscriptionNewResponseProductPricesProductPriceOneTimeJSON `json:"-"`
}

// userSubscriptionNewResponseProductPricesProductPriceOneTimeJSON contains the
// JSON metadata for the struct
// [UserSubscriptionNewResponseProductPricesProductPriceOneTime]
type userSubscriptionNewResponseProductPricesProductPriceOneTimeJSON struct {
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

func (r *UserSubscriptionNewResponseProductPricesProductPriceOneTime) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userSubscriptionNewResponseProductPricesProductPriceOneTimeJSON) RawJSON() (string) {
  return r.raw
}

func (r UserSubscriptionNewResponseProductPricesProductPriceOneTime) implementsUserSubscriptionNewResponseProductPrice() {}

// The type of the price.
type UserSubscriptionNewResponseProductPricesProductPriceOneTimeType string

const (
  UserSubscriptionNewResponseProductPricesProductPriceOneTimeTypeOneTime UserSubscriptionNewResponseProductPricesProductPriceOneTimeType = "one_time"
)

func (r UserSubscriptionNewResponseProductPricesProductPriceOneTimeType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionNewResponseProductPricesProductPriceOneTimeTypeOneTime:
      return true
  }
  return false
}

// The type of the price.
type UserSubscriptionNewResponseProductPricesType string

const (
  UserSubscriptionNewResponseProductPricesTypeRecurring UserSubscriptionNewResponseProductPricesType = "recurring"
  UserSubscriptionNewResponseProductPricesTypeOneTime UserSubscriptionNewResponseProductPricesType = "one_time"
)

func (r UserSubscriptionNewResponseProductPricesType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionNewResponseProductPricesTypeRecurring, UserSubscriptionNewResponseProductPricesTypeOneTime:
      return true
  }
  return false
}

// The recurring interval of the price, if type is `recurring`.
type UserSubscriptionNewResponseProductPricesRecurringInterval string

const (
  UserSubscriptionNewResponseProductPricesRecurringIntervalMonth UserSubscriptionNewResponseProductPricesRecurringInterval = "month"
  UserSubscriptionNewResponseProductPricesRecurringIntervalYear UserSubscriptionNewResponseProductPricesRecurringInterval = "year"
)

func (r UserSubscriptionNewResponseProductPricesRecurringInterval) IsKnown() (bool) {
  switch r {
  case UserSubscriptionNewResponseProductPricesRecurringIntervalMonth, UserSubscriptionNewResponseProductPricesRecurringIntervalYear:
      return true
  }
  return false
}

type UserSubscriptionNewResponseProductType string

const (
  UserSubscriptionNewResponseProductTypeFree UserSubscriptionNewResponseProductType = "free"
  UserSubscriptionNewResponseProductTypeIndividual UserSubscriptionNewResponseProductType = "individual"
  UserSubscriptionNewResponseProductTypeBusiness UserSubscriptionNewResponseProductType = "business"
)

func (r UserSubscriptionNewResponseProductType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionNewResponseProductTypeFree, UserSubscriptionNewResponseProductTypeIndividual, UserSubscriptionNewResponseProductTypeBusiness:
      return true
  }
  return false
}

type UserSubscriptionNewResponseStatus string

const (
  UserSubscriptionNewResponseStatusIncomplete UserSubscriptionNewResponseStatus = "incomplete"
  UserSubscriptionNewResponseStatusIncompleteExpired UserSubscriptionNewResponseStatus = "incomplete_expired"
  UserSubscriptionNewResponseStatusTrialing UserSubscriptionNewResponseStatus = "trialing"
  UserSubscriptionNewResponseStatusActive UserSubscriptionNewResponseStatus = "active"
  UserSubscriptionNewResponseStatusPastDue UserSubscriptionNewResponseStatus = "past_due"
  UserSubscriptionNewResponseStatusCanceled UserSubscriptionNewResponseStatus = "canceled"
  UserSubscriptionNewResponseStatusUnpaid UserSubscriptionNewResponseStatus = "unpaid"
)

func (r UserSubscriptionNewResponseStatus) IsKnown() (bool) {
  switch r {
  case UserSubscriptionNewResponseStatusIncomplete, UserSubscriptionNewResponseStatusIncompleteExpired, UserSubscriptionNewResponseStatusTrialing, UserSubscriptionNewResponseStatusActive, UserSubscriptionNewResponseStatusPastDue, UserSubscriptionNewResponseStatusCanceled, UserSubscriptionNewResponseStatusUnpaid:
      return true
  }
  return false
}

type UserSubscriptionGetResponse struct {
ID string `json:"id,required" format:"uuid4"`
CancelAtPeriodEnd bool `json:"cancel_at_period_end,required"`
// Creation timestamp of the object.
CreatedAt time.Time `json:"created_at,required" format:"date-time"`
CurrentPeriodEnd time.Time `json:"current_period_end,required,nullable" format:"date-time"`
CurrentPeriodStart time.Time `json:"current_period_start,required" format:"date-time"`
EndedAt time.Time `json:"ended_at,required,nullable" format:"date-time"`
// Last modification timestamp of the object.
ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
// A recurring price for a product, i.e. a subscription.
Price UserSubscriptionGetResponsePrice `json:"price,required,nullable"`
PriceID string `json:"price_id,required,nullable" format:"uuid4"`
Product UserSubscriptionGetResponseProduct `json:"product,required"`
ProductID string `json:"product_id,required" format:"uuid4"`
StartedAt time.Time `json:"started_at,required,nullable" format:"date-time"`
Status UserSubscriptionGetResponseStatus `json:"status,required"`
UserID string `json:"user_id,required" format:"uuid4"`
JSON userSubscriptionGetResponseJSON `json:"-"`
}

// userSubscriptionGetResponseJSON contains the JSON metadata for the struct
// [UserSubscriptionGetResponse]
type userSubscriptionGetResponseJSON struct {
ID apijson.Field
CancelAtPeriodEnd apijson.Field
CreatedAt apijson.Field
CurrentPeriodEnd apijson.Field
CurrentPeriodStart apijson.Field
EndedAt apijson.Field
ModifiedAt apijson.Field
Price apijson.Field
PriceID apijson.Field
Product apijson.Field
ProductID apijson.Field
StartedAt apijson.Field
Status apijson.Field
UserID apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionGetResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userSubscriptionGetResponseJSON) RawJSON() (string) {
  return r.raw
}

// A recurring price for a product, i.e. a subscription.
type UserSubscriptionGetResponsePrice struct {
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
Type UserSubscriptionGetResponsePriceType `json:"type,required"`
// The recurring interval of the price, if type is `recurring`.
RecurringInterval UserSubscriptionGetResponsePriceRecurringInterval `json:"recurring_interval,nullable"`
JSON userSubscriptionGetResponsePriceJSON `json:"-"`
union UserSubscriptionGetResponsePriceUnion
}

// userSubscriptionGetResponsePriceJSON contains the JSON metadata for the struct
// [UserSubscriptionGetResponsePrice]
type userSubscriptionGetResponsePriceJSON struct {
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

func (r userSubscriptionGetResponsePriceJSON) RawJSON() (string) {
  return r.raw
}

func (r *UserSubscriptionGetResponsePrice) UnmarshalJSON(data []byte) (err error) {
  *r = UserSubscriptionGetResponsePrice{}
  err = apijson.UnmarshalRoot(data, &r.union)
  if err != nil {
    return err
  }
  return apijson.Port(r.union, &r)
}

// AsUnion returns a [UserSubscriptionGetResponsePriceUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [UserSubscriptionGetResponsePriceProductPriceRecurring],
// [UserSubscriptionGetResponsePriceProductPriceOneTime].
func (r UserSubscriptionGetResponsePrice) AsUnion() (UserSubscriptionGetResponsePriceUnion) {
  return r.union
}

// A recurring price for a product, i.e. a subscription.
//
// Union satisfied by [UserSubscriptionGetResponsePriceProductPriceRecurring] or
// [UserSubscriptionGetResponsePriceProductPriceOneTime].
type UserSubscriptionGetResponsePriceUnion interface {
  implementsUserSubscriptionGetResponsePrice()
}

func init() {
  apijson.RegisterUnion(
    reflect.TypeOf((*UserSubscriptionGetResponsePriceUnion)(nil)).Elem(),
    "type",
    apijson.UnionVariant{
      TypeFilter: gjson.JSON,
      Type: reflect.TypeOf(UserSubscriptionGetResponsePriceProductPriceRecurring{}),
      DiscriminatorValue: "recurring",
    },
    apijson.UnionVariant{
      TypeFilter: gjson.JSON,
      Type: reflect.TypeOf(UserSubscriptionGetResponsePriceProductPriceOneTime{}),
      DiscriminatorValue: "one_time",
    },
  )
}

// A recurring price for a product, i.e. a subscription.
type UserSubscriptionGetResponsePriceProductPriceRecurring struct {
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
RecurringInterval UserSubscriptionGetResponsePriceProductPriceRecurringRecurringInterval `json:"recurring_interval,required,nullable"`
// The type of the price.
Type UserSubscriptionGetResponsePriceProductPriceRecurringType `json:"type,required"`
JSON userSubscriptionGetResponsePriceProductPriceRecurringJSON `json:"-"`
}

// userSubscriptionGetResponsePriceProductPriceRecurringJSON contains the JSON
// metadata for the struct [UserSubscriptionGetResponsePriceProductPriceRecurring]
type userSubscriptionGetResponsePriceProductPriceRecurringJSON struct {
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

func (r *UserSubscriptionGetResponsePriceProductPriceRecurring) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userSubscriptionGetResponsePriceProductPriceRecurringJSON) RawJSON() (string) {
  return r.raw
}

func (r UserSubscriptionGetResponsePriceProductPriceRecurring) implementsUserSubscriptionGetResponsePrice() {}

// The recurring interval of the price, if type is `recurring`.
type UserSubscriptionGetResponsePriceProductPriceRecurringRecurringInterval string

const (
  UserSubscriptionGetResponsePriceProductPriceRecurringRecurringIntervalMonth UserSubscriptionGetResponsePriceProductPriceRecurringRecurringInterval = "month"
  UserSubscriptionGetResponsePriceProductPriceRecurringRecurringIntervalYear UserSubscriptionGetResponsePriceProductPriceRecurringRecurringInterval = "year"
)

func (r UserSubscriptionGetResponsePriceProductPriceRecurringRecurringInterval) IsKnown() (bool) {
  switch r {
  case UserSubscriptionGetResponsePriceProductPriceRecurringRecurringIntervalMonth, UserSubscriptionGetResponsePriceProductPriceRecurringRecurringIntervalYear:
      return true
  }
  return false
}

// The type of the price.
type UserSubscriptionGetResponsePriceProductPriceRecurringType string

const (
  UserSubscriptionGetResponsePriceProductPriceRecurringTypeRecurring UserSubscriptionGetResponsePriceProductPriceRecurringType = "recurring"
)

func (r UserSubscriptionGetResponsePriceProductPriceRecurringType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionGetResponsePriceProductPriceRecurringTypeRecurring:
      return true
  }
  return false
}

// A one-time price for a product.
type UserSubscriptionGetResponsePriceProductPriceOneTime struct {
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
Type UserSubscriptionGetResponsePriceProductPriceOneTimeType `json:"type,required"`
JSON userSubscriptionGetResponsePriceProductPriceOneTimeJSON `json:"-"`
}

// userSubscriptionGetResponsePriceProductPriceOneTimeJSON contains the JSON
// metadata for the struct [UserSubscriptionGetResponsePriceProductPriceOneTime]
type userSubscriptionGetResponsePriceProductPriceOneTimeJSON struct {
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

func (r *UserSubscriptionGetResponsePriceProductPriceOneTime) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userSubscriptionGetResponsePriceProductPriceOneTimeJSON) RawJSON() (string) {
  return r.raw
}

func (r UserSubscriptionGetResponsePriceProductPriceOneTime) implementsUserSubscriptionGetResponsePrice() {}

// The type of the price.
type UserSubscriptionGetResponsePriceProductPriceOneTimeType string

const (
  UserSubscriptionGetResponsePriceProductPriceOneTimeTypeOneTime UserSubscriptionGetResponsePriceProductPriceOneTimeType = "one_time"
)

func (r UserSubscriptionGetResponsePriceProductPriceOneTimeType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionGetResponsePriceProductPriceOneTimeTypeOneTime:
      return true
  }
  return false
}

// The type of the price.
type UserSubscriptionGetResponsePriceType string

const (
  UserSubscriptionGetResponsePriceTypeRecurring UserSubscriptionGetResponsePriceType = "recurring"
  UserSubscriptionGetResponsePriceTypeOneTime UserSubscriptionGetResponsePriceType = "one_time"
)

func (r UserSubscriptionGetResponsePriceType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionGetResponsePriceTypeRecurring, UserSubscriptionGetResponsePriceTypeOneTime:
      return true
  }
  return false
}

// The recurring interval of the price, if type is `recurring`.
type UserSubscriptionGetResponsePriceRecurringInterval string

const (
  UserSubscriptionGetResponsePriceRecurringIntervalMonth UserSubscriptionGetResponsePriceRecurringInterval = "month"
  UserSubscriptionGetResponsePriceRecurringIntervalYear UserSubscriptionGetResponsePriceRecurringInterval = "year"
)

func (r UserSubscriptionGetResponsePriceRecurringInterval) IsKnown() (bool) {
  switch r {
  case UserSubscriptionGetResponsePriceRecurringIntervalMonth, UserSubscriptionGetResponsePriceRecurringIntervalYear:
      return true
  }
  return false
}

type UserSubscriptionGetResponseProduct struct {
// The ID of the product.
ID string `json:"id,required" format:"uuid4"`
// The benefits granted by the product.
Benefits []UserSubscriptionGetResponseProductBenefit `json:"benefits,required"`
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
Prices []UserSubscriptionGetResponseProductPrice `json:"prices,required"`
Type UserSubscriptionGetResponseProductType `json:"type,required,nullable"`
JSON userSubscriptionGetResponseProductJSON `json:"-"`
}

// userSubscriptionGetResponseProductJSON contains the JSON metadata for the struct
// [UserSubscriptionGetResponseProduct]
type userSubscriptionGetResponseProductJSON struct {
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

func (r *UserSubscriptionGetResponseProduct) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userSubscriptionGetResponseProductJSON) RawJSON() (string) {
  return r.raw
}

// A benefit of type `articles`.
//
// Use it to grant access to posts.
type UserSubscriptionGetResponseProductBenefit struct {
// Creation timestamp of the object.
CreatedAt time.Time `json:"created_at,required" format:"date-time"`
// Last modification timestamp of the object.
ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
// The ID of the benefit.
ID string `json:"id,required" format:"uuid4"`
// The type of the benefit.
Type UserSubscriptionGetResponseProductBenefitsType `json:"type,required"`
// The description of the benefit.
Description string `json:"description,required"`
// Whether the benefit is selectable when creating a product.
Selectable bool `json:"selectable,required"`
// Whether the benefit is deletable.
Deletable bool `json:"deletable,required"`
// The ID of the organization owning the benefit.
OrganizationID string `json:"organization_id,required" format:"uuid4"`
// This field can have the runtime type of
// [UserSubscriptionGetResponseProductBenefitsBenefitArticlesProperties].
Properties interface{} `json:"properties,required"`
JSON userSubscriptionGetResponseProductBenefitJSON `json:"-"`
union UserSubscriptionGetResponseProductBenefitsUnion
}

// userSubscriptionGetResponseProductBenefitJSON contains the JSON metadata for the
// struct [UserSubscriptionGetResponseProductBenefit]
type userSubscriptionGetResponseProductBenefitJSON struct {
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

func (r userSubscriptionGetResponseProductBenefitJSON) RawJSON() (string) {
  return r.raw
}

func (r *UserSubscriptionGetResponseProductBenefit) UnmarshalJSON(data []byte) (err error) {
  *r = UserSubscriptionGetResponseProductBenefit{}
  err = apijson.UnmarshalRoot(data, &r.union)
  if err != nil {
    return err
  }
  return apijson.Port(r.union, &r)
}

// AsUnion returns a [UserSubscriptionGetResponseProductBenefitsUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [UserSubscriptionGetResponseProductBenefitsBenefitBase],
// [UserSubscriptionGetResponseProductBenefitsBenefitArticles].
func (r UserSubscriptionGetResponseProductBenefit) AsUnion() (UserSubscriptionGetResponseProductBenefitsUnion) {
  return r.union
}

// A benefit of type `articles`.
//
// Use it to grant access to posts.
//
// Union satisfied by [UserSubscriptionGetResponseProductBenefitsBenefitBase] or
// [UserSubscriptionGetResponseProductBenefitsBenefitArticles].
type UserSubscriptionGetResponseProductBenefitsUnion interface {
  implementsUserSubscriptionGetResponseProductBenefit()
}

func init() {
  apijson.RegisterUnion(
    reflect.TypeOf((*UserSubscriptionGetResponseProductBenefitsUnion)(nil)).Elem(),
    "",
    apijson.UnionVariant{
      TypeFilter: gjson.JSON,
      Type: reflect.TypeOf(UserSubscriptionGetResponseProductBenefitsBenefitBase{}),
    },
    apijson.UnionVariant{
      TypeFilter: gjson.JSON,
      Type: reflect.TypeOf(UserSubscriptionGetResponseProductBenefitsBenefitArticles{}),
    },
  )
}

type UserSubscriptionGetResponseProductBenefitsBenefitBase struct {
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
Type UserSubscriptionGetResponseProductBenefitsBenefitBaseType `json:"type,required"`
JSON userSubscriptionGetResponseProductBenefitsBenefitBaseJSON `json:"-"`
}

// userSubscriptionGetResponseProductBenefitsBenefitBaseJSON contains the JSON
// metadata for the struct [UserSubscriptionGetResponseProductBenefitsBenefitBase]
type userSubscriptionGetResponseProductBenefitsBenefitBaseJSON struct {
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

func (r *UserSubscriptionGetResponseProductBenefitsBenefitBase) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userSubscriptionGetResponseProductBenefitsBenefitBaseJSON) RawJSON() (string) {
  return r.raw
}

func (r UserSubscriptionGetResponseProductBenefitsBenefitBase) implementsUserSubscriptionGetResponseProductBenefit() {}

// The type of the benefit.
type UserSubscriptionGetResponseProductBenefitsBenefitBaseType string

const (
  UserSubscriptionGetResponseProductBenefitsBenefitBaseTypeCustom UserSubscriptionGetResponseProductBenefitsBenefitBaseType = "custom"
  UserSubscriptionGetResponseProductBenefitsBenefitBaseTypeArticles UserSubscriptionGetResponseProductBenefitsBenefitBaseType = "articles"
  UserSubscriptionGetResponseProductBenefitsBenefitBaseTypeAds UserSubscriptionGetResponseProductBenefitsBenefitBaseType = "ads"
  UserSubscriptionGetResponseProductBenefitsBenefitBaseTypeDiscord UserSubscriptionGetResponseProductBenefitsBenefitBaseType = "discord"
  UserSubscriptionGetResponseProductBenefitsBenefitBaseTypeGitHubRepository UserSubscriptionGetResponseProductBenefitsBenefitBaseType = "github_repository"
  UserSubscriptionGetResponseProductBenefitsBenefitBaseTypeDownloadables UserSubscriptionGetResponseProductBenefitsBenefitBaseType = "downloadables"
)

func (r UserSubscriptionGetResponseProductBenefitsBenefitBaseType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionGetResponseProductBenefitsBenefitBaseTypeCustom, UserSubscriptionGetResponseProductBenefitsBenefitBaseTypeArticles, UserSubscriptionGetResponseProductBenefitsBenefitBaseTypeAds, UserSubscriptionGetResponseProductBenefitsBenefitBaseTypeDiscord, UserSubscriptionGetResponseProductBenefitsBenefitBaseTypeGitHubRepository, UserSubscriptionGetResponseProductBenefitsBenefitBaseTypeDownloadables:
      return true
  }
  return false
}

// A benefit of type `articles`.
//
// Use it to grant access to posts.
type UserSubscriptionGetResponseProductBenefitsBenefitArticles struct {
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
Properties UserSubscriptionGetResponseProductBenefitsBenefitArticlesProperties `json:"properties,required"`
// Whether the benefit is selectable when creating a product.
Selectable bool `json:"selectable,required"`
Type UserSubscriptionGetResponseProductBenefitsBenefitArticlesType `json:"type,required"`
JSON userSubscriptionGetResponseProductBenefitsBenefitArticlesJSON `json:"-"`
}

// userSubscriptionGetResponseProductBenefitsBenefitArticlesJSON contains the JSON
// metadata for the struct
// [UserSubscriptionGetResponseProductBenefitsBenefitArticles]
type userSubscriptionGetResponseProductBenefitsBenefitArticlesJSON struct {
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

func (r *UserSubscriptionGetResponseProductBenefitsBenefitArticles) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userSubscriptionGetResponseProductBenefitsBenefitArticlesJSON) RawJSON() (string) {
  return r.raw
}

func (r UserSubscriptionGetResponseProductBenefitsBenefitArticles) implementsUserSubscriptionGetResponseProductBenefit() {}

// Properties for a benefit of type `articles`.
type UserSubscriptionGetResponseProductBenefitsBenefitArticlesProperties struct {
// Whether the user can access paid articles.
PaidArticles bool `json:"paid_articles,required"`
JSON userSubscriptionGetResponseProductBenefitsBenefitArticlesPropertiesJSON `json:"-"`
}

// userSubscriptionGetResponseProductBenefitsBenefitArticlesPropertiesJSON contains
// the JSON metadata for the struct
// [UserSubscriptionGetResponseProductBenefitsBenefitArticlesProperties]
type userSubscriptionGetResponseProductBenefitsBenefitArticlesPropertiesJSON struct {
PaidArticles apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionGetResponseProductBenefitsBenefitArticlesProperties) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userSubscriptionGetResponseProductBenefitsBenefitArticlesPropertiesJSON) RawJSON() (string) {
  return r.raw
}

type UserSubscriptionGetResponseProductBenefitsBenefitArticlesType string

const (
  UserSubscriptionGetResponseProductBenefitsBenefitArticlesTypeArticles UserSubscriptionGetResponseProductBenefitsBenefitArticlesType = "articles"
)

func (r UserSubscriptionGetResponseProductBenefitsBenefitArticlesType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionGetResponseProductBenefitsBenefitArticlesTypeArticles:
      return true
  }
  return false
}

// The type of the benefit.
type UserSubscriptionGetResponseProductBenefitsType string

const (
  UserSubscriptionGetResponseProductBenefitsTypeCustom UserSubscriptionGetResponseProductBenefitsType = "custom"
  UserSubscriptionGetResponseProductBenefitsTypeArticles UserSubscriptionGetResponseProductBenefitsType = "articles"
  UserSubscriptionGetResponseProductBenefitsTypeAds UserSubscriptionGetResponseProductBenefitsType = "ads"
  UserSubscriptionGetResponseProductBenefitsTypeDiscord UserSubscriptionGetResponseProductBenefitsType = "discord"
  UserSubscriptionGetResponseProductBenefitsTypeGitHubRepository UserSubscriptionGetResponseProductBenefitsType = "github_repository"
  UserSubscriptionGetResponseProductBenefitsTypeDownloadables UserSubscriptionGetResponseProductBenefitsType = "downloadables"
)

func (r UserSubscriptionGetResponseProductBenefitsType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionGetResponseProductBenefitsTypeCustom, UserSubscriptionGetResponseProductBenefitsTypeArticles, UserSubscriptionGetResponseProductBenefitsTypeAds, UserSubscriptionGetResponseProductBenefitsTypeDiscord, UserSubscriptionGetResponseProductBenefitsTypeGitHubRepository, UserSubscriptionGetResponseProductBenefitsTypeDownloadables:
      return true
  }
  return false
}

// A recurring price for a product, i.e. a subscription.
type UserSubscriptionGetResponseProductPrice struct {
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
Type UserSubscriptionGetResponseProductPricesType `json:"type,required"`
// The recurring interval of the price, if type is `recurring`.
RecurringInterval UserSubscriptionGetResponseProductPricesRecurringInterval `json:"recurring_interval,nullable"`
JSON userSubscriptionGetResponseProductPriceJSON `json:"-"`
union UserSubscriptionGetResponseProductPricesUnion
}

// userSubscriptionGetResponseProductPriceJSON contains the JSON metadata for the
// struct [UserSubscriptionGetResponseProductPrice]
type userSubscriptionGetResponseProductPriceJSON struct {
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

func (r userSubscriptionGetResponseProductPriceJSON) RawJSON() (string) {
  return r.raw
}

func (r *UserSubscriptionGetResponseProductPrice) UnmarshalJSON(data []byte) (err error) {
  *r = UserSubscriptionGetResponseProductPrice{}
  err = apijson.UnmarshalRoot(data, &r.union)
  if err != nil {
    return err
  }
  return apijson.Port(r.union, &r)
}

// AsUnion returns a [UserSubscriptionGetResponseProductPricesUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [UserSubscriptionGetResponseProductPricesProductPriceRecurring],
// [UserSubscriptionGetResponseProductPricesProductPriceOneTime].
func (r UserSubscriptionGetResponseProductPrice) AsUnion() (UserSubscriptionGetResponseProductPricesUnion) {
  return r.union
}

// A recurring price for a product, i.e. a subscription.
//
// Union satisfied by
// [UserSubscriptionGetResponseProductPricesProductPriceRecurring] or
// [UserSubscriptionGetResponseProductPricesProductPriceOneTime].
type UserSubscriptionGetResponseProductPricesUnion interface {
  implementsUserSubscriptionGetResponseProductPrice()
}

func init() {
  apijson.RegisterUnion(
    reflect.TypeOf((*UserSubscriptionGetResponseProductPricesUnion)(nil)).Elem(),
    "type",
    apijson.UnionVariant{
      TypeFilter: gjson.JSON,
      Type: reflect.TypeOf(UserSubscriptionGetResponseProductPricesProductPriceRecurring{}),
      DiscriminatorValue: "recurring",
    },
    apijson.UnionVariant{
      TypeFilter: gjson.JSON,
      Type: reflect.TypeOf(UserSubscriptionGetResponseProductPricesProductPriceOneTime{}),
      DiscriminatorValue: "one_time",
    },
  )
}

// A recurring price for a product, i.e. a subscription.
type UserSubscriptionGetResponseProductPricesProductPriceRecurring struct {
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
RecurringInterval UserSubscriptionGetResponseProductPricesProductPriceRecurringRecurringInterval `json:"recurring_interval,required,nullable"`
// The type of the price.
Type UserSubscriptionGetResponseProductPricesProductPriceRecurringType `json:"type,required"`
JSON userSubscriptionGetResponseProductPricesProductPriceRecurringJSON `json:"-"`
}

// userSubscriptionGetResponseProductPricesProductPriceRecurringJSON contains the
// JSON metadata for the struct
// [UserSubscriptionGetResponseProductPricesProductPriceRecurring]
type userSubscriptionGetResponseProductPricesProductPriceRecurringJSON struct {
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

func (r *UserSubscriptionGetResponseProductPricesProductPriceRecurring) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userSubscriptionGetResponseProductPricesProductPriceRecurringJSON) RawJSON() (string) {
  return r.raw
}

func (r UserSubscriptionGetResponseProductPricesProductPriceRecurring) implementsUserSubscriptionGetResponseProductPrice() {}

// The recurring interval of the price, if type is `recurring`.
type UserSubscriptionGetResponseProductPricesProductPriceRecurringRecurringInterval string

const (
  UserSubscriptionGetResponseProductPricesProductPriceRecurringRecurringIntervalMonth UserSubscriptionGetResponseProductPricesProductPriceRecurringRecurringInterval = "month"
  UserSubscriptionGetResponseProductPricesProductPriceRecurringRecurringIntervalYear UserSubscriptionGetResponseProductPricesProductPriceRecurringRecurringInterval = "year"
)

func (r UserSubscriptionGetResponseProductPricesProductPriceRecurringRecurringInterval) IsKnown() (bool) {
  switch r {
  case UserSubscriptionGetResponseProductPricesProductPriceRecurringRecurringIntervalMonth, UserSubscriptionGetResponseProductPricesProductPriceRecurringRecurringIntervalYear:
      return true
  }
  return false
}

// The type of the price.
type UserSubscriptionGetResponseProductPricesProductPriceRecurringType string

const (
  UserSubscriptionGetResponseProductPricesProductPriceRecurringTypeRecurring UserSubscriptionGetResponseProductPricesProductPriceRecurringType = "recurring"
)

func (r UserSubscriptionGetResponseProductPricesProductPriceRecurringType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionGetResponseProductPricesProductPriceRecurringTypeRecurring:
      return true
  }
  return false
}

// A one-time price for a product.
type UserSubscriptionGetResponseProductPricesProductPriceOneTime struct {
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
Type UserSubscriptionGetResponseProductPricesProductPriceOneTimeType `json:"type,required"`
JSON userSubscriptionGetResponseProductPricesProductPriceOneTimeJSON `json:"-"`
}

// userSubscriptionGetResponseProductPricesProductPriceOneTimeJSON contains the
// JSON metadata for the struct
// [UserSubscriptionGetResponseProductPricesProductPriceOneTime]
type userSubscriptionGetResponseProductPricesProductPriceOneTimeJSON struct {
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

func (r *UserSubscriptionGetResponseProductPricesProductPriceOneTime) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userSubscriptionGetResponseProductPricesProductPriceOneTimeJSON) RawJSON() (string) {
  return r.raw
}

func (r UserSubscriptionGetResponseProductPricesProductPriceOneTime) implementsUserSubscriptionGetResponseProductPrice() {}

// The type of the price.
type UserSubscriptionGetResponseProductPricesProductPriceOneTimeType string

const (
  UserSubscriptionGetResponseProductPricesProductPriceOneTimeTypeOneTime UserSubscriptionGetResponseProductPricesProductPriceOneTimeType = "one_time"
)

func (r UserSubscriptionGetResponseProductPricesProductPriceOneTimeType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionGetResponseProductPricesProductPriceOneTimeTypeOneTime:
      return true
  }
  return false
}

// The type of the price.
type UserSubscriptionGetResponseProductPricesType string

const (
  UserSubscriptionGetResponseProductPricesTypeRecurring UserSubscriptionGetResponseProductPricesType = "recurring"
  UserSubscriptionGetResponseProductPricesTypeOneTime UserSubscriptionGetResponseProductPricesType = "one_time"
)

func (r UserSubscriptionGetResponseProductPricesType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionGetResponseProductPricesTypeRecurring, UserSubscriptionGetResponseProductPricesTypeOneTime:
      return true
  }
  return false
}

// The recurring interval of the price, if type is `recurring`.
type UserSubscriptionGetResponseProductPricesRecurringInterval string

const (
  UserSubscriptionGetResponseProductPricesRecurringIntervalMonth UserSubscriptionGetResponseProductPricesRecurringInterval = "month"
  UserSubscriptionGetResponseProductPricesRecurringIntervalYear UserSubscriptionGetResponseProductPricesRecurringInterval = "year"
)

func (r UserSubscriptionGetResponseProductPricesRecurringInterval) IsKnown() (bool) {
  switch r {
  case UserSubscriptionGetResponseProductPricesRecurringIntervalMonth, UserSubscriptionGetResponseProductPricesRecurringIntervalYear:
      return true
  }
  return false
}

type UserSubscriptionGetResponseProductType string

const (
  UserSubscriptionGetResponseProductTypeFree UserSubscriptionGetResponseProductType = "free"
  UserSubscriptionGetResponseProductTypeIndividual UserSubscriptionGetResponseProductType = "individual"
  UserSubscriptionGetResponseProductTypeBusiness UserSubscriptionGetResponseProductType = "business"
)

func (r UserSubscriptionGetResponseProductType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionGetResponseProductTypeFree, UserSubscriptionGetResponseProductTypeIndividual, UserSubscriptionGetResponseProductTypeBusiness:
      return true
  }
  return false
}

type UserSubscriptionGetResponseStatus string

const (
  UserSubscriptionGetResponseStatusIncomplete UserSubscriptionGetResponseStatus = "incomplete"
  UserSubscriptionGetResponseStatusIncompleteExpired UserSubscriptionGetResponseStatus = "incomplete_expired"
  UserSubscriptionGetResponseStatusTrialing UserSubscriptionGetResponseStatus = "trialing"
  UserSubscriptionGetResponseStatusActive UserSubscriptionGetResponseStatus = "active"
  UserSubscriptionGetResponseStatusPastDue UserSubscriptionGetResponseStatus = "past_due"
  UserSubscriptionGetResponseStatusCanceled UserSubscriptionGetResponseStatus = "canceled"
  UserSubscriptionGetResponseStatusUnpaid UserSubscriptionGetResponseStatus = "unpaid"
)

func (r UserSubscriptionGetResponseStatus) IsKnown() (bool) {
  switch r {
  case UserSubscriptionGetResponseStatusIncomplete, UserSubscriptionGetResponseStatusIncompleteExpired, UserSubscriptionGetResponseStatusTrialing, UserSubscriptionGetResponseStatusActive, UserSubscriptionGetResponseStatusPastDue, UserSubscriptionGetResponseStatusCanceled, UserSubscriptionGetResponseStatusUnpaid:
      return true
  }
  return false
}

type UserSubscriptionUpdateResponse struct {
ID string `json:"id,required" format:"uuid4"`
CancelAtPeriodEnd bool `json:"cancel_at_period_end,required"`
// Creation timestamp of the object.
CreatedAt time.Time `json:"created_at,required" format:"date-time"`
CurrentPeriodEnd time.Time `json:"current_period_end,required,nullable" format:"date-time"`
CurrentPeriodStart time.Time `json:"current_period_start,required" format:"date-time"`
EndedAt time.Time `json:"ended_at,required,nullable" format:"date-time"`
// Last modification timestamp of the object.
ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
// A recurring price for a product, i.e. a subscription.
Price UserSubscriptionUpdateResponsePrice `json:"price,required,nullable"`
PriceID string `json:"price_id,required,nullable" format:"uuid4"`
Product UserSubscriptionUpdateResponseProduct `json:"product,required"`
ProductID string `json:"product_id,required" format:"uuid4"`
StartedAt time.Time `json:"started_at,required,nullable" format:"date-time"`
Status UserSubscriptionUpdateResponseStatus `json:"status,required"`
UserID string `json:"user_id,required" format:"uuid4"`
JSON userSubscriptionUpdateResponseJSON `json:"-"`
}

// userSubscriptionUpdateResponseJSON contains the JSON metadata for the struct
// [UserSubscriptionUpdateResponse]
type userSubscriptionUpdateResponseJSON struct {
ID apijson.Field
CancelAtPeriodEnd apijson.Field
CreatedAt apijson.Field
CurrentPeriodEnd apijson.Field
CurrentPeriodStart apijson.Field
EndedAt apijson.Field
ModifiedAt apijson.Field
Price apijson.Field
PriceID apijson.Field
Product apijson.Field
ProductID apijson.Field
StartedAt apijson.Field
Status apijson.Field
UserID apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionUpdateResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userSubscriptionUpdateResponseJSON) RawJSON() (string) {
  return r.raw
}

// A recurring price for a product, i.e. a subscription.
type UserSubscriptionUpdateResponsePrice struct {
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
Type UserSubscriptionUpdateResponsePriceType `json:"type,required"`
// The recurring interval of the price, if type is `recurring`.
RecurringInterval UserSubscriptionUpdateResponsePriceRecurringInterval `json:"recurring_interval,nullable"`
JSON userSubscriptionUpdateResponsePriceJSON `json:"-"`
union UserSubscriptionUpdateResponsePriceUnion
}

// userSubscriptionUpdateResponsePriceJSON contains the JSON metadata for the
// struct [UserSubscriptionUpdateResponsePrice]
type userSubscriptionUpdateResponsePriceJSON struct {
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

func (r userSubscriptionUpdateResponsePriceJSON) RawJSON() (string) {
  return r.raw
}

func (r *UserSubscriptionUpdateResponsePrice) UnmarshalJSON(data []byte) (err error) {
  *r = UserSubscriptionUpdateResponsePrice{}
  err = apijson.UnmarshalRoot(data, &r.union)
  if err != nil {
    return err
  }
  return apijson.Port(r.union, &r)
}

// AsUnion returns a [UserSubscriptionUpdateResponsePriceUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [UserSubscriptionUpdateResponsePriceProductPriceRecurring],
// [UserSubscriptionUpdateResponsePriceProductPriceOneTime].
func (r UserSubscriptionUpdateResponsePrice) AsUnion() (UserSubscriptionUpdateResponsePriceUnion) {
  return r.union
}

// A recurring price for a product, i.e. a subscription.
//
// Union satisfied by [UserSubscriptionUpdateResponsePriceProductPriceRecurring] or
// [UserSubscriptionUpdateResponsePriceProductPriceOneTime].
type UserSubscriptionUpdateResponsePriceUnion interface {
  implementsUserSubscriptionUpdateResponsePrice()
}

func init() {
  apijson.RegisterUnion(
    reflect.TypeOf((*UserSubscriptionUpdateResponsePriceUnion)(nil)).Elem(),
    "type",
    apijson.UnionVariant{
      TypeFilter: gjson.JSON,
      Type: reflect.TypeOf(UserSubscriptionUpdateResponsePriceProductPriceRecurring{}),
      DiscriminatorValue: "recurring",
    },
    apijson.UnionVariant{
      TypeFilter: gjson.JSON,
      Type: reflect.TypeOf(UserSubscriptionUpdateResponsePriceProductPriceOneTime{}),
      DiscriminatorValue: "one_time",
    },
  )
}

// A recurring price for a product, i.e. a subscription.
type UserSubscriptionUpdateResponsePriceProductPriceRecurring struct {
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
RecurringInterval UserSubscriptionUpdateResponsePriceProductPriceRecurringRecurringInterval `json:"recurring_interval,required,nullable"`
// The type of the price.
Type UserSubscriptionUpdateResponsePriceProductPriceRecurringType `json:"type,required"`
JSON userSubscriptionUpdateResponsePriceProductPriceRecurringJSON `json:"-"`
}

// userSubscriptionUpdateResponsePriceProductPriceRecurringJSON contains the JSON
// metadata for the struct
// [UserSubscriptionUpdateResponsePriceProductPriceRecurring]
type userSubscriptionUpdateResponsePriceProductPriceRecurringJSON struct {
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

func (r *UserSubscriptionUpdateResponsePriceProductPriceRecurring) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userSubscriptionUpdateResponsePriceProductPriceRecurringJSON) RawJSON() (string) {
  return r.raw
}

func (r UserSubscriptionUpdateResponsePriceProductPriceRecurring) implementsUserSubscriptionUpdateResponsePrice() {}

// The recurring interval of the price, if type is `recurring`.
type UserSubscriptionUpdateResponsePriceProductPriceRecurringRecurringInterval string

const (
  UserSubscriptionUpdateResponsePriceProductPriceRecurringRecurringIntervalMonth UserSubscriptionUpdateResponsePriceProductPriceRecurringRecurringInterval = "month"
  UserSubscriptionUpdateResponsePriceProductPriceRecurringRecurringIntervalYear UserSubscriptionUpdateResponsePriceProductPriceRecurringRecurringInterval = "year"
)

func (r UserSubscriptionUpdateResponsePriceProductPriceRecurringRecurringInterval) IsKnown() (bool) {
  switch r {
  case UserSubscriptionUpdateResponsePriceProductPriceRecurringRecurringIntervalMonth, UserSubscriptionUpdateResponsePriceProductPriceRecurringRecurringIntervalYear:
      return true
  }
  return false
}

// The type of the price.
type UserSubscriptionUpdateResponsePriceProductPriceRecurringType string

const (
  UserSubscriptionUpdateResponsePriceProductPriceRecurringTypeRecurring UserSubscriptionUpdateResponsePriceProductPriceRecurringType = "recurring"
)

func (r UserSubscriptionUpdateResponsePriceProductPriceRecurringType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionUpdateResponsePriceProductPriceRecurringTypeRecurring:
      return true
  }
  return false
}

// A one-time price for a product.
type UserSubscriptionUpdateResponsePriceProductPriceOneTime struct {
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
Type UserSubscriptionUpdateResponsePriceProductPriceOneTimeType `json:"type,required"`
JSON userSubscriptionUpdateResponsePriceProductPriceOneTimeJSON `json:"-"`
}

// userSubscriptionUpdateResponsePriceProductPriceOneTimeJSON contains the JSON
// metadata for the struct [UserSubscriptionUpdateResponsePriceProductPriceOneTime]
type userSubscriptionUpdateResponsePriceProductPriceOneTimeJSON struct {
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

func (r *UserSubscriptionUpdateResponsePriceProductPriceOneTime) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userSubscriptionUpdateResponsePriceProductPriceOneTimeJSON) RawJSON() (string) {
  return r.raw
}

func (r UserSubscriptionUpdateResponsePriceProductPriceOneTime) implementsUserSubscriptionUpdateResponsePrice() {}

// The type of the price.
type UserSubscriptionUpdateResponsePriceProductPriceOneTimeType string

const (
  UserSubscriptionUpdateResponsePriceProductPriceOneTimeTypeOneTime UserSubscriptionUpdateResponsePriceProductPriceOneTimeType = "one_time"
)

func (r UserSubscriptionUpdateResponsePriceProductPriceOneTimeType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionUpdateResponsePriceProductPriceOneTimeTypeOneTime:
      return true
  }
  return false
}

// The type of the price.
type UserSubscriptionUpdateResponsePriceType string

const (
  UserSubscriptionUpdateResponsePriceTypeRecurring UserSubscriptionUpdateResponsePriceType = "recurring"
  UserSubscriptionUpdateResponsePriceTypeOneTime UserSubscriptionUpdateResponsePriceType = "one_time"
)

func (r UserSubscriptionUpdateResponsePriceType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionUpdateResponsePriceTypeRecurring, UserSubscriptionUpdateResponsePriceTypeOneTime:
      return true
  }
  return false
}

// The recurring interval of the price, if type is `recurring`.
type UserSubscriptionUpdateResponsePriceRecurringInterval string

const (
  UserSubscriptionUpdateResponsePriceRecurringIntervalMonth UserSubscriptionUpdateResponsePriceRecurringInterval = "month"
  UserSubscriptionUpdateResponsePriceRecurringIntervalYear UserSubscriptionUpdateResponsePriceRecurringInterval = "year"
)

func (r UserSubscriptionUpdateResponsePriceRecurringInterval) IsKnown() (bool) {
  switch r {
  case UserSubscriptionUpdateResponsePriceRecurringIntervalMonth, UserSubscriptionUpdateResponsePriceRecurringIntervalYear:
      return true
  }
  return false
}

type UserSubscriptionUpdateResponseProduct struct {
// The ID of the product.
ID string `json:"id,required" format:"uuid4"`
// The benefits granted by the product.
Benefits []UserSubscriptionUpdateResponseProductBenefit `json:"benefits,required"`
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
Prices []UserSubscriptionUpdateResponseProductPrice `json:"prices,required"`
Type UserSubscriptionUpdateResponseProductType `json:"type,required,nullable"`
JSON userSubscriptionUpdateResponseProductJSON `json:"-"`
}

// userSubscriptionUpdateResponseProductJSON contains the JSON metadata for the
// struct [UserSubscriptionUpdateResponseProduct]
type userSubscriptionUpdateResponseProductJSON struct {
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

func (r *UserSubscriptionUpdateResponseProduct) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userSubscriptionUpdateResponseProductJSON) RawJSON() (string) {
  return r.raw
}

// A benefit of type `articles`.
//
// Use it to grant access to posts.
type UserSubscriptionUpdateResponseProductBenefit struct {
// Creation timestamp of the object.
CreatedAt time.Time `json:"created_at,required" format:"date-time"`
// Last modification timestamp of the object.
ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
// The ID of the benefit.
ID string `json:"id,required" format:"uuid4"`
// The type of the benefit.
Type UserSubscriptionUpdateResponseProductBenefitsType `json:"type,required"`
// The description of the benefit.
Description string `json:"description,required"`
// Whether the benefit is selectable when creating a product.
Selectable bool `json:"selectable,required"`
// Whether the benefit is deletable.
Deletable bool `json:"deletable,required"`
// The ID of the organization owning the benefit.
OrganizationID string `json:"organization_id,required" format:"uuid4"`
// This field can have the runtime type of
// [UserSubscriptionUpdateResponseProductBenefitsBenefitArticlesProperties].
Properties interface{} `json:"properties,required"`
JSON userSubscriptionUpdateResponseProductBenefitJSON `json:"-"`
union UserSubscriptionUpdateResponseProductBenefitsUnion
}

// userSubscriptionUpdateResponseProductBenefitJSON contains the JSON metadata for
// the struct [UserSubscriptionUpdateResponseProductBenefit]
type userSubscriptionUpdateResponseProductBenefitJSON struct {
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

func (r userSubscriptionUpdateResponseProductBenefitJSON) RawJSON() (string) {
  return r.raw
}

func (r *UserSubscriptionUpdateResponseProductBenefit) UnmarshalJSON(data []byte) (err error) {
  *r = UserSubscriptionUpdateResponseProductBenefit{}
  err = apijson.UnmarshalRoot(data, &r.union)
  if err != nil {
    return err
  }
  return apijson.Port(r.union, &r)
}

// AsUnion returns a [UserSubscriptionUpdateResponseProductBenefitsUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [UserSubscriptionUpdateResponseProductBenefitsBenefitBase],
// [UserSubscriptionUpdateResponseProductBenefitsBenefitArticles].
func (r UserSubscriptionUpdateResponseProductBenefit) AsUnion() (UserSubscriptionUpdateResponseProductBenefitsUnion) {
  return r.union
}

// A benefit of type `articles`.
//
// Use it to grant access to posts.
//
// Union satisfied by [UserSubscriptionUpdateResponseProductBenefitsBenefitBase] or
// [UserSubscriptionUpdateResponseProductBenefitsBenefitArticles].
type UserSubscriptionUpdateResponseProductBenefitsUnion interface {
  implementsUserSubscriptionUpdateResponseProductBenefit()
}

func init() {
  apijson.RegisterUnion(
    reflect.TypeOf((*UserSubscriptionUpdateResponseProductBenefitsUnion)(nil)).Elem(),
    "",
    apijson.UnionVariant{
      TypeFilter: gjson.JSON,
      Type: reflect.TypeOf(UserSubscriptionUpdateResponseProductBenefitsBenefitBase{}),
    },
    apijson.UnionVariant{
      TypeFilter: gjson.JSON,
      Type: reflect.TypeOf(UserSubscriptionUpdateResponseProductBenefitsBenefitArticles{}),
    },
  )
}

type UserSubscriptionUpdateResponseProductBenefitsBenefitBase struct {
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
Type UserSubscriptionUpdateResponseProductBenefitsBenefitBaseType `json:"type,required"`
JSON userSubscriptionUpdateResponseProductBenefitsBenefitBaseJSON `json:"-"`
}

// userSubscriptionUpdateResponseProductBenefitsBenefitBaseJSON contains the JSON
// metadata for the struct
// [UserSubscriptionUpdateResponseProductBenefitsBenefitBase]
type userSubscriptionUpdateResponseProductBenefitsBenefitBaseJSON struct {
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

func (r *UserSubscriptionUpdateResponseProductBenefitsBenefitBase) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userSubscriptionUpdateResponseProductBenefitsBenefitBaseJSON) RawJSON() (string) {
  return r.raw
}

func (r UserSubscriptionUpdateResponseProductBenefitsBenefitBase) implementsUserSubscriptionUpdateResponseProductBenefit() {}

// The type of the benefit.
type UserSubscriptionUpdateResponseProductBenefitsBenefitBaseType string

const (
  UserSubscriptionUpdateResponseProductBenefitsBenefitBaseTypeCustom UserSubscriptionUpdateResponseProductBenefitsBenefitBaseType = "custom"
  UserSubscriptionUpdateResponseProductBenefitsBenefitBaseTypeArticles UserSubscriptionUpdateResponseProductBenefitsBenefitBaseType = "articles"
  UserSubscriptionUpdateResponseProductBenefitsBenefitBaseTypeAds UserSubscriptionUpdateResponseProductBenefitsBenefitBaseType = "ads"
  UserSubscriptionUpdateResponseProductBenefitsBenefitBaseTypeDiscord UserSubscriptionUpdateResponseProductBenefitsBenefitBaseType = "discord"
  UserSubscriptionUpdateResponseProductBenefitsBenefitBaseTypeGitHubRepository UserSubscriptionUpdateResponseProductBenefitsBenefitBaseType = "github_repository"
  UserSubscriptionUpdateResponseProductBenefitsBenefitBaseTypeDownloadables UserSubscriptionUpdateResponseProductBenefitsBenefitBaseType = "downloadables"
)

func (r UserSubscriptionUpdateResponseProductBenefitsBenefitBaseType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionUpdateResponseProductBenefitsBenefitBaseTypeCustom, UserSubscriptionUpdateResponseProductBenefitsBenefitBaseTypeArticles, UserSubscriptionUpdateResponseProductBenefitsBenefitBaseTypeAds, UserSubscriptionUpdateResponseProductBenefitsBenefitBaseTypeDiscord, UserSubscriptionUpdateResponseProductBenefitsBenefitBaseTypeGitHubRepository, UserSubscriptionUpdateResponseProductBenefitsBenefitBaseTypeDownloadables:
      return true
  }
  return false
}

// A benefit of type `articles`.
//
// Use it to grant access to posts.
type UserSubscriptionUpdateResponseProductBenefitsBenefitArticles struct {
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
Properties UserSubscriptionUpdateResponseProductBenefitsBenefitArticlesProperties `json:"properties,required"`
// Whether the benefit is selectable when creating a product.
Selectable bool `json:"selectable,required"`
Type UserSubscriptionUpdateResponseProductBenefitsBenefitArticlesType `json:"type,required"`
JSON userSubscriptionUpdateResponseProductBenefitsBenefitArticlesJSON `json:"-"`
}

// userSubscriptionUpdateResponseProductBenefitsBenefitArticlesJSON contains the
// JSON metadata for the struct
// [UserSubscriptionUpdateResponseProductBenefitsBenefitArticles]
type userSubscriptionUpdateResponseProductBenefitsBenefitArticlesJSON struct {
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

func (r *UserSubscriptionUpdateResponseProductBenefitsBenefitArticles) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userSubscriptionUpdateResponseProductBenefitsBenefitArticlesJSON) RawJSON() (string) {
  return r.raw
}

func (r UserSubscriptionUpdateResponseProductBenefitsBenefitArticles) implementsUserSubscriptionUpdateResponseProductBenefit() {}

// Properties for a benefit of type `articles`.
type UserSubscriptionUpdateResponseProductBenefitsBenefitArticlesProperties struct {
// Whether the user can access paid articles.
PaidArticles bool `json:"paid_articles,required"`
JSON userSubscriptionUpdateResponseProductBenefitsBenefitArticlesPropertiesJSON `json:"-"`
}

// userSubscriptionUpdateResponseProductBenefitsBenefitArticlesPropertiesJSON
// contains the JSON metadata for the struct
// [UserSubscriptionUpdateResponseProductBenefitsBenefitArticlesProperties]
type userSubscriptionUpdateResponseProductBenefitsBenefitArticlesPropertiesJSON struct {
PaidArticles apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionUpdateResponseProductBenefitsBenefitArticlesProperties) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userSubscriptionUpdateResponseProductBenefitsBenefitArticlesPropertiesJSON) RawJSON() (string) {
  return r.raw
}

type UserSubscriptionUpdateResponseProductBenefitsBenefitArticlesType string

const (
  UserSubscriptionUpdateResponseProductBenefitsBenefitArticlesTypeArticles UserSubscriptionUpdateResponseProductBenefitsBenefitArticlesType = "articles"
)

func (r UserSubscriptionUpdateResponseProductBenefitsBenefitArticlesType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionUpdateResponseProductBenefitsBenefitArticlesTypeArticles:
      return true
  }
  return false
}

// The type of the benefit.
type UserSubscriptionUpdateResponseProductBenefitsType string

const (
  UserSubscriptionUpdateResponseProductBenefitsTypeCustom UserSubscriptionUpdateResponseProductBenefitsType = "custom"
  UserSubscriptionUpdateResponseProductBenefitsTypeArticles UserSubscriptionUpdateResponseProductBenefitsType = "articles"
  UserSubscriptionUpdateResponseProductBenefitsTypeAds UserSubscriptionUpdateResponseProductBenefitsType = "ads"
  UserSubscriptionUpdateResponseProductBenefitsTypeDiscord UserSubscriptionUpdateResponseProductBenefitsType = "discord"
  UserSubscriptionUpdateResponseProductBenefitsTypeGitHubRepository UserSubscriptionUpdateResponseProductBenefitsType = "github_repository"
  UserSubscriptionUpdateResponseProductBenefitsTypeDownloadables UserSubscriptionUpdateResponseProductBenefitsType = "downloadables"
)

func (r UserSubscriptionUpdateResponseProductBenefitsType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionUpdateResponseProductBenefitsTypeCustom, UserSubscriptionUpdateResponseProductBenefitsTypeArticles, UserSubscriptionUpdateResponseProductBenefitsTypeAds, UserSubscriptionUpdateResponseProductBenefitsTypeDiscord, UserSubscriptionUpdateResponseProductBenefitsTypeGitHubRepository, UserSubscriptionUpdateResponseProductBenefitsTypeDownloadables:
      return true
  }
  return false
}

// A recurring price for a product, i.e. a subscription.
type UserSubscriptionUpdateResponseProductPrice struct {
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
Type UserSubscriptionUpdateResponseProductPricesType `json:"type,required"`
// The recurring interval of the price, if type is `recurring`.
RecurringInterval UserSubscriptionUpdateResponseProductPricesRecurringInterval `json:"recurring_interval,nullable"`
JSON userSubscriptionUpdateResponseProductPriceJSON `json:"-"`
union UserSubscriptionUpdateResponseProductPricesUnion
}

// userSubscriptionUpdateResponseProductPriceJSON contains the JSON metadata for
// the struct [UserSubscriptionUpdateResponseProductPrice]
type userSubscriptionUpdateResponseProductPriceJSON struct {
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

func (r userSubscriptionUpdateResponseProductPriceJSON) RawJSON() (string) {
  return r.raw
}

func (r *UserSubscriptionUpdateResponseProductPrice) UnmarshalJSON(data []byte) (err error) {
  *r = UserSubscriptionUpdateResponseProductPrice{}
  err = apijson.UnmarshalRoot(data, &r.union)
  if err != nil {
    return err
  }
  return apijson.Port(r.union, &r)
}

// AsUnion returns a [UserSubscriptionUpdateResponseProductPricesUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [UserSubscriptionUpdateResponseProductPricesProductPriceRecurring],
// [UserSubscriptionUpdateResponseProductPricesProductPriceOneTime].
func (r UserSubscriptionUpdateResponseProductPrice) AsUnion() (UserSubscriptionUpdateResponseProductPricesUnion) {
  return r.union
}

// A recurring price for a product, i.e. a subscription.
//
// Union satisfied by
// [UserSubscriptionUpdateResponseProductPricesProductPriceRecurring] or
// [UserSubscriptionUpdateResponseProductPricesProductPriceOneTime].
type UserSubscriptionUpdateResponseProductPricesUnion interface {
  implementsUserSubscriptionUpdateResponseProductPrice()
}

func init() {
  apijson.RegisterUnion(
    reflect.TypeOf((*UserSubscriptionUpdateResponseProductPricesUnion)(nil)).Elem(),
    "type",
    apijson.UnionVariant{
      TypeFilter: gjson.JSON,
      Type: reflect.TypeOf(UserSubscriptionUpdateResponseProductPricesProductPriceRecurring{}),
      DiscriminatorValue: "recurring",
    },
    apijson.UnionVariant{
      TypeFilter: gjson.JSON,
      Type: reflect.TypeOf(UserSubscriptionUpdateResponseProductPricesProductPriceOneTime{}),
      DiscriminatorValue: "one_time",
    },
  )
}

// A recurring price for a product, i.e. a subscription.
type UserSubscriptionUpdateResponseProductPricesProductPriceRecurring struct {
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
RecurringInterval UserSubscriptionUpdateResponseProductPricesProductPriceRecurringRecurringInterval `json:"recurring_interval,required,nullable"`
// The type of the price.
Type UserSubscriptionUpdateResponseProductPricesProductPriceRecurringType `json:"type,required"`
JSON userSubscriptionUpdateResponseProductPricesProductPriceRecurringJSON `json:"-"`
}

// userSubscriptionUpdateResponseProductPricesProductPriceRecurringJSON contains
// the JSON metadata for the struct
// [UserSubscriptionUpdateResponseProductPricesProductPriceRecurring]
type userSubscriptionUpdateResponseProductPricesProductPriceRecurringJSON struct {
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

func (r *UserSubscriptionUpdateResponseProductPricesProductPriceRecurring) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userSubscriptionUpdateResponseProductPricesProductPriceRecurringJSON) RawJSON() (string) {
  return r.raw
}

func (r UserSubscriptionUpdateResponseProductPricesProductPriceRecurring) implementsUserSubscriptionUpdateResponseProductPrice() {}

// The recurring interval of the price, if type is `recurring`.
type UserSubscriptionUpdateResponseProductPricesProductPriceRecurringRecurringInterval string

const (
  UserSubscriptionUpdateResponseProductPricesProductPriceRecurringRecurringIntervalMonth UserSubscriptionUpdateResponseProductPricesProductPriceRecurringRecurringInterval = "month"
  UserSubscriptionUpdateResponseProductPricesProductPriceRecurringRecurringIntervalYear UserSubscriptionUpdateResponseProductPricesProductPriceRecurringRecurringInterval = "year"
)

func (r UserSubscriptionUpdateResponseProductPricesProductPriceRecurringRecurringInterval) IsKnown() (bool) {
  switch r {
  case UserSubscriptionUpdateResponseProductPricesProductPriceRecurringRecurringIntervalMonth, UserSubscriptionUpdateResponseProductPricesProductPriceRecurringRecurringIntervalYear:
      return true
  }
  return false
}

// The type of the price.
type UserSubscriptionUpdateResponseProductPricesProductPriceRecurringType string

const (
  UserSubscriptionUpdateResponseProductPricesProductPriceRecurringTypeRecurring UserSubscriptionUpdateResponseProductPricesProductPriceRecurringType = "recurring"
)

func (r UserSubscriptionUpdateResponseProductPricesProductPriceRecurringType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionUpdateResponseProductPricesProductPriceRecurringTypeRecurring:
      return true
  }
  return false
}

// A one-time price for a product.
type UserSubscriptionUpdateResponseProductPricesProductPriceOneTime struct {
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
Type UserSubscriptionUpdateResponseProductPricesProductPriceOneTimeType `json:"type,required"`
JSON userSubscriptionUpdateResponseProductPricesProductPriceOneTimeJSON `json:"-"`
}

// userSubscriptionUpdateResponseProductPricesProductPriceOneTimeJSON contains the
// JSON metadata for the struct
// [UserSubscriptionUpdateResponseProductPricesProductPriceOneTime]
type userSubscriptionUpdateResponseProductPricesProductPriceOneTimeJSON struct {
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

func (r *UserSubscriptionUpdateResponseProductPricesProductPriceOneTime) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userSubscriptionUpdateResponseProductPricesProductPriceOneTimeJSON) RawJSON() (string) {
  return r.raw
}

func (r UserSubscriptionUpdateResponseProductPricesProductPriceOneTime) implementsUserSubscriptionUpdateResponseProductPrice() {}

// The type of the price.
type UserSubscriptionUpdateResponseProductPricesProductPriceOneTimeType string

const (
  UserSubscriptionUpdateResponseProductPricesProductPriceOneTimeTypeOneTime UserSubscriptionUpdateResponseProductPricesProductPriceOneTimeType = "one_time"
)

func (r UserSubscriptionUpdateResponseProductPricesProductPriceOneTimeType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionUpdateResponseProductPricesProductPriceOneTimeTypeOneTime:
      return true
  }
  return false
}

// The type of the price.
type UserSubscriptionUpdateResponseProductPricesType string

const (
  UserSubscriptionUpdateResponseProductPricesTypeRecurring UserSubscriptionUpdateResponseProductPricesType = "recurring"
  UserSubscriptionUpdateResponseProductPricesTypeOneTime UserSubscriptionUpdateResponseProductPricesType = "one_time"
)

func (r UserSubscriptionUpdateResponseProductPricesType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionUpdateResponseProductPricesTypeRecurring, UserSubscriptionUpdateResponseProductPricesTypeOneTime:
      return true
  }
  return false
}

// The recurring interval of the price, if type is `recurring`.
type UserSubscriptionUpdateResponseProductPricesRecurringInterval string

const (
  UserSubscriptionUpdateResponseProductPricesRecurringIntervalMonth UserSubscriptionUpdateResponseProductPricesRecurringInterval = "month"
  UserSubscriptionUpdateResponseProductPricesRecurringIntervalYear UserSubscriptionUpdateResponseProductPricesRecurringInterval = "year"
)

func (r UserSubscriptionUpdateResponseProductPricesRecurringInterval) IsKnown() (bool) {
  switch r {
  case UserSubscriptionUpdateResponseProductPricesRecurringIntervalMonth, UserSubscriptionUpdateResponseProductPricesRecurringIntervalYear:
      return true
  }
  return false
}

type UserSubscriptionUpdateResponseProductType string

const (
  UserSubscriptionUpdateResponseProductTypeFree UserSubscriptionUpdateResponseProductType = "free"
  UserSubscriptionUpdateResponseProductTypeIndividual UserSubscriptionUpdateResponseProductType = "individual"
  UserSubscriptionUpdateResponseProductTypeBusiness UserSubscriptionUpdateResponseProductType = "business"
)

func (r UserSubscriptionUpdateResponseProductType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionUpdateResponseProductTypeFree, UserSubscriptionUpdateResponseProductTypeIndividual, UserSubscriptionUpdateResponseProductTypeBusiness:
      return true
  }
  return false
}

type UserSubscriptionUpdateResponseStatus string

const (
  UserSubscriptionUpdateResponseStatusIncomplete UserSubscriptionUpdateResponseStatus = "incomplete"
  UserSubscriptionUpdateResponseStatusIncompleteExpired UserSubscriptionUpdateResponseStatus = "incomplete_expired"
  UserSubscriptionUpdateResponseStatusTrialing UserSubscriptionUpdateResponseStatus = "trialing"
  UserSubscriptionUpdateResponseStatusActive UserSubscriptionUpdateResponseStatus = "active"
  UserSubscriptionUpdateResponseStatusPastDue UserSubscriptionUpdateResponseStatus = "past_due"
  UserSubscriptionUpdateResponseStatusCanceled UserSubscriptionUpdateResponseStatus = "canceled"
  UserSubscriptionUpdateResponseStatusUnpaid UserSubscriptionUpdateResponseStatus = "unpaid"
)

func (r UserSubscriptionUpdateResponseStatus) IsKnown() (bool) {
  switch r {
  case UserSubscriptionUpdateResponseStatusIncomplete, UserSubscriptionUpdateResponseStatusIncompleteExpired, UserSubscriptionUpdateResponseStatusTrialing, UserSubscriptionUpdateResponseStatusActive, UserSubscriptionUpdateResponseStatusPastDue, UserSubscriptionUpdateResponseStatusCanceled, UserSubscriptionUpdateResponseStatusUnpaid:
      return true
  }
  return false
}

type UserSubscriptionListResponse struct {
ID string `json:"id,required" format:"uuid4"`
CancelAtPeriodEnd bool `json:"cancel_at_period_end,required"`
// Creation timestamp of the object.
CreatedAt time.Time `json:"created_at,required" format:"date-time"`
CurrentPeriodEnd time.Time `json:"current_period_end,required,nullable" format:"date-time"`
CurrentPeriodStart time.Time `json:"current_period_start,required" format:"date-time"`
EndedAt time.Time `json:"ended_at,required,nullable" format:"date-time"`
// Last modification timestamp of the object.
ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
// A recurring price for a product, i.e. a subscription.
Price UserSubscriptionListResponsePrice `json:"price,required,nullable"`
PriceID string `json:"price_id,required,nullable" format:"uuid4"`
Product UserSubscriptionListResponseProduct `json:"product,required"`
ProductID string `json:"product_id,required" format:"uuid4"`
StartedAt time.Time `json:"started_at,required,nullable" format:"date-time"`
Status UserSubscriptionListResponseStatus `json:"status,required"`
UserID string `json:"user_id,required" format:"uuid4"`
JSON userSubscriptionListResponseJSON `json:"-"`
}

// userSubscriptionListResponseJSON contains the JSON metadata for the struct
// [UserSubscriptionListResponse]
type userSubscriptionListResponseJSON struct {
ID apijson.Field
CancelAtPeriodEnd apijson.Field
CreatedAt apijson.Field
CurrentPeriodEnd apijson.Field
CurrentPeriodStart apijson.Field
EndedAt apijson.Field
ModifiedAt apijson.Field
Price apijson.Field
PriceID apijson.Field
Product apijson.Field
ProductID apijson.Field
StartedAt apijson.Field
Status apijson.Field
UserID apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionListResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userSubscriptionListResponseJSON) RawJSON() (string) {
  return r.raw
}

// A recurring price for a product, i.e. a subscription.
type UserSubscriptionListResponsePrice struct {
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
Type UserSubscriptionListResponsePriceType `json:"type,required"`
// The recurring interval of the price, if type is `recurring`.
RecurringInterval UserSubscriptionListResponsePriceRecurringInterval `json:"recurring_interval,nullable"`
JSON userSubscriptionListResponsePriceJSON `json:"-"`
union UserSubscriptionListResponsePriceUnion
}

// userSubscriptionListResponsePriceJSON contains the JSON metadata for the struct
// [UserSubscriptionListResponsePrice]
type userSubscriptionListResponsePriceJSON struct {
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

func (r userSubscriptionListResponsePriceJSON) RawJSON() (string) {
  return r.raw
}

func (r *UserSubscriptionListResponsePrice) UnmarshalJSON(data []byte) (err error) {
  *r = UserSubscriptionListResponsePrice{}
  err = apijson.UnmarshalRoot(data, &r.union)
  if err != nil {
    return err
  }
  return apijson.Port(r.union, &r)
}

// AsUnion returns a [UserSubscriptionListResponsePriceUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [UserSubscriptionListResponsePriceProductPriceRecurring],
// [UserSubscriptionListResponsePriceProductPriceOneTime].
func (r UserSubscriptionListResponsePrice) AsUnion() (UserSubscriptionListResponsePriceUnion) {
  return r.union
}

// A recurring price for a product, i.e. a subscription.
//
// Union satisfied by [UserSubscriptionListResponsePriceProductPriceRecurring] or
// [UserSubscriptionListResponsePriceProductPriceOneTime].
type UserSubscriptionListResponsePriceUnion interface {
  implementsUserSubscriptionListResponsePrice()
}

func init() {
  apijson.RegisterUnion(
    reflect.TypeOf((*UserSubscriptionListResponsePriceUnion)(nil)).Elem(),
    "type",
    apijson.UnionVariant{
      TypeFilter: gjson.JSON,
      Type: reflect.TypeOf(UserSubscriptionListResponsePriceProductPriceRecurring{}),
      DiscriminatorValue: "recurring",
    },
    apijson.UnionVariant{
      TypeFilter: gjson.JSON,
      Type: reflect.TypeOf(UserSubscriptionListResponsePriceProductPriceOneTime{}),
      DiscriminatorValue: "one_time",
    },
  )
}

// A recurring price for a product, i.e. a subscription.
type UserSubscriptionListResponsePriceProductPriceRecurring struct {
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
RecurringInterval UserSubscriptionListResponsePriceProductPriceRecurringRecurringInterval `json:"recurring_interval,required,nullable"`
// The type of the price.
Type UserSubscriptionListResponsePriceProductPriceRecurringType `json:"type,required"`
JSON userSubscriptionListResponsePriceProductPriceRecurringJSON `json:"-"`
}

// userSubscriptionListResponsePriceProductPriceRecurringJSON contains the JSON
// metadata for the struct [UserSubscriptionListResponsePriceProductPriceRecurring]
type userSubscriptionListResponsePriceProductPriceRecurringJSON struct {
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

func (r *UserSubscriptionListResponsePriceProductPriceRecurring) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userSubscriptionListResponsePriceProductPriceRecurringJSON) RawJSON() (string) {
  return r.raw
}

func (r UserSubscriptionListResponsePriceProductPriceRecurring) implementsUserSubscriptionListResponsePrice() {}

// The recurring interval of the price, if type is `recurring`.
type UserSubscriptionListResponsePriceProductPriceRecurringRecurringInterval string

const (
  UserSubscriptionListResponsePriceProductPriceRecurringRecurringIntervalMonth UserSubscriptionListResponsePriceProductPriceRecurringRecurringInterval = "month"
  UserSubscriptionListResponsePriceProductPriceRecurringRecurringIntervalYear UserSubscriptionListResponsePriceProductPriceRecurringRecurringInterval = "year"
)

func (r UserSubscriptionListResponsePriceProductPriceRecurringRecurringInterval) IsKnown() (bool) {
  switch r {
  case UserSubscriptionListResponsePriceProductPriceRecurringRecurringIntervalMonth, UserSubscriptionListResponsePriceProductPriceRecurringRecurringIntervalYear:
      return true
  }
  return false
}

// The type of the price.
type UserSubscriptionListResponsePriceProductPriceRecurringType string

const (
  UserSubscriptionListResponsePriceProductPriceRecurringTypeRecurring UserSubscriptionListResponsePriceProductPriceRecurringType = "recurring"
)

func (r UserSubscriptionListResponsePriceProductPriceRecurringType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionListResponsePriceProductPriceRecurringTypeRecurring:
      return true
  }
  return false
}

// A one-time price for a product.
type UserSubscriptionListResponsePriceProductPriceOneTime struct {
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
Type UserSubscriptionListResponsePriceProductPriceOneTimeType `json:"type,required"`
JSON userSubscriptionListResponsePriceProductPriceOneTimeJSON `json:"-"`
}

// userSubscriptionListResponsePriceProductPriceOneTimeJSON contains the JSON
// metadata for the struct [UserSubscriptionListResponsePriceProductPriceOneTime]
type userSubscriptionListResponsePriceProductPriceOneTimeJSON struct {
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

func (r *UserSubscriptionListResponsePriceProductPriceOneTime) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userSubscriptionListResponsePriceProductPriceOneTimeJSON) RawJSON() (string) {
  return r.raw
}

func (r UserSubscriptionListResponsePriceProductPriceOneTime) implementsUserSubscriptionListResponsePrice() {}

// The type of the price.
type UserSubscriptionListResponsePriceProductPriceOneTimeType string

const (
  UserSubscriptionListResponsePriceProductPriceOneTimeTypeOneTime UserSubscriptionListResponsePriceProductPriceOneTimeType = "one_time"
)

func (r UserSubscriptionListResponsePriceProductPriceOneTimeType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionListResponsePriceProductPriceOneTimeTypeOneTime:
      return true
  }
  return false
}

// The type of the price.
type UserSubscriptionListResponsePriceType string

const (
  UserSubscriptionListResponsePriceTypeRecurring UserSubscriptionListResponsePriceType = "recurring"
  UserSubscriptionListResponsePriceTypeOneTime UserSubscriptionListResponsePriceType = "one_time"
)

func (r UserSubscriptionListResponsePriceType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionListResponsePriceTypeRecurring, UserSubscriptionListResponsePriceTypeOneTime:
      return true
  }
  return false
}

// The recurring interval of the price, if type is `recurring`.
type UserSubscriptionListResponsePriceRecurringInterval string

const (
  UserSubscriptionListResponsePriceRecurringIntervalMonth UserSubscriptionListResponsePriceRecurringInterval = "month"
  UserSubscriptionListResponsePriceRecurringIntervalYear UserSubscriptionListResponsePriceRecurringInterval = "year"
)

func (r UserSubscriptionListResponsePriceRecurringInterval) IsKnown() (bool) {
  switch r {
  case UserSubscriptionListResponsePriceRecurringIntervalMonth, UserSubscriptionListResponsePriceRecurringIntervalYear:
      return true
  }
  return false
}

type UserSubscriptionListResponseProduct struct {
// The ID of the product.
ID string `json:"id,required" format:"uuid4"`
// The benefits granted by the product.
Benefits []UserSubscriptionListResponseProductBenefit `json:"benefits,required"`
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
Prices []UserSubscriptionListResponseProductPrice `json:"prices,required"`
Type UserSubscriptionListResponseProductType `json:"type,required,nullable"`
JSON userSubscriptionListResponseProductJSON `json:"-"`
}

// userSubscriptionListResponseProductJSON contains the JSON metadata for the
// struct [UserSubscriptionListResponseProduct]
type userSubscriptionListResponseProductJSON struct {
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

func (r *UserSubscriptionListResponseProduct) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userSubscriptionListResponseProductJSON) RawJSON() (string) {
  return r.raw
}

// A benefit of type `articles`.
//
// Use it to grant access to posts.
type UserSubscriptionListResponseProductBenefit struct {
// Creation timestamp of the object.
CreatedAt time.Time `json:"created_at,required" format:"date-time"`
// Last modification timestamp of the object.
ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
// The ID of the benefit.
ID string `json:"id,required" format:"uuid4"`
// The type of the benefit.
Type UserSubscriptionListResponseProductBenefitsType `json:"type,required"`
// The description of the benefit.
Description string `json:"description,required"`
// Whether the benefit is selectable when creating a product.
Selectable bool `json:"selectable,required"`
// Whether the benefit is deletable.
Deletable bool `json:"deletable,required"`
// The ID of the organization owning the benefit.
OrganizationID string `json:"organization_id,required" format:"uuid4"`
// This field can have the runtime type of
// [UserSubscriptionListResponseProductBenefitsBenefitArticlesProperties].
Properties interface{} `json:"properties,required"`
JSON userSubscriptionListResponseProductBenefitJSON `json:"-"`
union UserSubscriptionListResponseProductBenefitsUnion
}

// userSubscriptionListResponseProductBenefitJSON contains the JSON metadata for
// the struct [UserSubscriptionListResponseProductBenefit]
type userSubscriptionListResponseProductBenefitJSON struct {
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

func (r userSubscriptionListResponseProductBenefitJSON) RawJSON() (string) {
  return r.raw
}

func (r *UserSubscriptionListResponseProductBenefit) UnmarshalJSON(data []byte) (err error) {
  *r = UserSubscriptionListResponseProductBenefit{}
  err = apijson.UnmarshalRoot(data, &r.union)
  if err != nil {
    return err
  }
  return apijson.Port(r.union, &r)
}

// AsUnion returns a [UserSubscriptionListResponseProductBenefitsUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [UserSubscriptionListResponseProductBenefitsBenefitBase],
// [UserSubscriptionListResponseProductBenefitsBenefitArticles].
func (r UserSubscriptionListResponseProductBenefit) AsUnion() (UserSubscriptionListResponseProductBenefitsUnion) {
  return r.union
}

// A benefit of type `articles`.
//
// Use it to grant access to posts.
//
// Union satisfied by [UserSubscriptionListResponseProductBenefitsBenefitBase] or
// [UserSubscriptionListResponseProductBenefitsBenefitArticles].
type UserSubscriptionListResponseProductBenefitsUnion interface {
  implementsUserSubscriptionListResponseProductBenefit()
}

func init() {
  apijson.RegisterUnion(
    reflect.TypeOf((*UserSubscriptionListResponseProductBenefitsUnion)(nil)).Elem(),
    "",
    apijson.UnionVariant{
      TypeFilter: gjson.JSON,
      Type: reflect.TypeOf(UserSubscriptionListResponseProductBenefitsBenefitBase{}),
    },
    apijson.UnionVariant{
      TypeFilter: gjson.JSON,
      Type: reflect.TypeOf(UserSubscriptionListResponseProductBenefitsBenefitArticles{}),
    },
  )
}

type UserSubscriptionListResponseProductBenefitsBenefitBase struct {
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
Type UserSubscriptionListResponseProductBenefitsBenefitBaseType `json:"type,required"`
JSON userSubscriptionListResponseProductBenefitsBenefitBaseJSON `json:"-"`
}

// userSubscriptionListResponseProductBenefitsBenefitBaseJSON contains the JSON
// metadata for the struct [UserSubscriptionListResponseProductBenefitsBenefitBase]
type userSubscriptionListResponseProductBenefitsBenefitBaseJSON struct {
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

func (r *UserSubscriptionListResponseProductBenefitsBenefitBase) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userSubscriptionListResponseProductBenefitsBenefitBaseJSON) RawJSON() (string) {
  return r.raw
}

func (r UserSubscriptionListResponseProductBenefitsBenefitBase) implementsUserSubscriptionListResponseProductBenefit() {}

// The type of the benefit.
type UserSubscriptionListResponseProductBenefitsBenefitBaseType string

const (
  UserSubscriptionListResponseProductBenefitsBenefitBaseTypeCustom UserSubscriptionListResponseProductBenefitsBenefitBaseType = "custom"
  UserSubscriptionListResponseProductBenefitsBenefitBaseTypeArticles UserSubscriptionListResponseProductBenefitsBenefitBaseType = "articles"
  UserSubscriptionListResponseProductBenefitsBenefitBaseTypeAds UserSubscriptionListResponseProductBenefitsBenefitBaseType = "ads"
  UserSubscriptionListResponseProductBenefitsBenefitBaseTypeDiscord UserSubscriptionListResponseProductBenefitsBenefitBaseType = "discord"
  UserSubscriptionListResponseProductBenefitsBenefitBaseTypeGitHubRepository UserSubscriptionListResponseProductBenefitsBenefitBaseType = "github_repository"
  UserSubscriptionListResponseProductBenefitsBenefitBaseTypeDownloadables UserSubscriptionListResponseProductBenefitsBenefitBaseType = "downloadables"
)

func (r UserSubscriptionListResponseProductBenefitsBenefitBaseType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionListResponseProductBenefitsBenefitBaseTypeCustom, UserSubscriptionListResponseProductBenefitsBenefitBaseTypeArticles, UserSubscriptionListResponseProductBenefitsBenefitBaseTypeAds, UserSubscriptionListResponseProductBenefitsBenefitBaseTypeDiscord, UserSubscriptionListResponseProductBenefitsBenefitBaseTypeGitHubRepository, UserSubscriptionListResponseProductBenefitsBenefitBaseTypeDownloadables:
      return true
  }
  return false
}

// A benefit of type `articles`.
//
// Use it to grant access to posts.
type UserSubscriptionListResponseProductBenefitsBenefitArticles struct {
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
Properties UserSubscriptionListResponseProductBenefitsBenefitArticlesProperties `json:"properties,required"`
// Whether the benefit is selectable when creating a product.
Selectable bool `json:"selectable,required"`
Type UserSubscriptionListResponseProductBenefitsBenefitArticlesType `json:"type,required"`
JSON userSubscriptionListResponseProductBenefitsBenefitArticlesJSON `json:"-"`
}

// userSubscriptionListResponseProductBenefitsBenefitArticlesJSON contains the JSON
// metadata for the struct
// [UserSubscriptionListResponseProductBenefitsBenefitArticles]
type userSubscriptionListResponseProductBenefitsBenefitArticlesJSON struct {
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

func (r *UserSubscriptionListResponseProductBenefitsBenefitArticles) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userSubscriptionListResponseProductBenefitsBenefitArticlesJSON) RawJSON() (string) {
  return r.raw
}

func (r UserSubscriptionListResponseProductBenefitsBenefitArticles) implementsUserSubscriptionListResponseProductBenefit() {}

// Properties for a benefit of type `articles`.
type UserSubscriptionListResponseProductBenefitsBenefitArticlesProperties struct {
// Whether the user can access paid articles.
PaidArticles bool `json:"paid_articles,required"`
JSON userSubscriptionListResponseProductBenefitsBenefitArticlesPropertiesJSON `json:"-"`
}

// userSubscriptionListResponseProductBenefitsBenefitArticlesPropertiesJSON
// contains the JSON metadata for the struct
// [UserSubscriptionListResponseProductBenefitsBenefitArticlesProperties]
type userSubscriptionListResponseProductBenefitsBenefitArticlesPropertiesJSON struct {
PaidArticles apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionListResponseProductBenefitsBenefitArticlesProperties) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userSubscriptionListResponseProductBenefitsBenefitArticlesPropertiesJSON) RawJSON() (string) {
  return r.raw
}

type UserSubscriptionListResponseProductBenefitsBenefitArticlesType string

const (
  UserSubscriptionListResponseProductBenefitsBenefitArticlesTypeArticles UserSubscriptionListResponseProductBenefitsBenefitArticlesType = "articles"
)

func (r UserSubscriptionListResponseProductBenefitsBenefitArticlesType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionListResponseProductBenefitsBenefitArticlesTypeArticles:
      return true
  }
  return false
}

// The type of the benefit.
type UserSubscriptionListResponseProductBenefitsType string

const (
  UserSubscriptionListResponseProductBenefitsTypeCustom UserSubscriptionListResponseProductBenefitsType = "custom"
  UserSubscriptionListResponseProductBenefitsTypeArticles UserSubscriptionListResponseProductBenefitsType = "articles"
  UserSubscriptionListResponseProductBenefitsTypeAds UserSubscriptionListResponseProductBenefitsType = "ads"
  UserSubscriptionListResponseProductBenefitsTypeDiscord UserSubscriptionListResponseProductBenefitsType = "discord"
  UserSubscriptionListResponseProductBenefitsTypeGitHubRepository UserSubscriptionListResponseProductBenefitsType = "github_repository"
  UserSubscriptionListResponseProductBenefitsTypeDownloadables UserSubscriptionListResponseProductBenefitsType = "downloadables"
)

func (r UserSubscriptionListResponseProductBenefitsType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionListResponseProductBenefitsTypeCustom, UserSubscriptionListResponseProductBenefitsTypeArticles, UserSubscriptionListResponseProductBenefitsTypeAds, UserSubscriptionListResponseProductBenefitsTypeDiscord, UserSubscriptionListResponseProductBenefitsTypeGitHubRepository, UserSubscriptionListResponseProductBenefitsTypeDownloadables:
      return true
  }
  return false
}

// A recurring price for a product, i.e. a subscription.
type UserSubscriptionListResponseProductPrice struct {
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
Type UserSubscriptionListResponseProductPricesType `json:"type,required"`
// The recurring interval of the price, if type is `recurring`.
RecurringInterval UserSubscriptionListResponseProductPricesRecurringInterval `json:"recurring_interval,nullable"`
JSON userSubscriptionListResponseProductPriceJSON `json:"-"`
union UserSubscriptionListResponseProductPricesUnion
}

// userSubscriptionListResponseProductPriceJSON contains the JSON metadata for the
// struct [UserSubscriptionListResponseProductPrice]
type userSubscriptionListResponseProductPriceJSON struct {
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

func (r userSubscriptionListResponseProductPriceJSON) RawJSON() (string) {
  return r.raw
}

func (r *UserSubscriptionListResponseProductPrice) UnmarshalJSON(data []byte) (err error) {
  *r = UserSubscriptionListResponseProductPrice{}
  err = apijson.UnmarshalRoot(data, &r.union)
  if err != nil {
    return err
  }
  return apijson.Port(r.union, &r)
}

// AsUnion returns a [UserSubscriptionListResponseProductPricesUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [UserSubscriptionListResponseProductPricesProductPriceRecurring],
// [UserSubscriptionListResponseProductPricesProductPriceOneTime].
func (r UserSubscriptionListResponseProductPrice) AsUnion() (UserSubscriptionListResponseProductPricesUnion) {
  return r.union
}

// A recurring price for a product, i.e. a subscription.
//
// Union satisfied by
// [UserSubscriptionListResponseProductPricesProductPriceRecurring] or
// [UserSubscriptionListResponseProductPricesProductPriceOneTime].
type UserSubscriptionListResponseProductPricesUnion interface {
  implementsUserSubscriptionListResponseProductPrice()
}

func init() {
  apijson.RegisterUnion(
    reflect.TypeOf((*UserSubscriptionListResponseProductPricesUnion)(nil)).Elem(),
    "type",
    apijson.UnionVariant{
      TypeFilter: gjson.JSON,
      Type: reflect.TypeOf(UserSubscriptionListResponseProductPricesProductPriceRecurring{}),
      DiscriminatorValue: "recurring",
    },
    apijson.UnionVariant{
      TypeFilter: gjson.JSON,
      Type: reflect.TypeOf(UserSubscriptionListResponseProductPricesProductPriceOneTime{}),
      DiscriminatorValue: "one_time",
    },
  )
}

// A recurring price for a product, i.e. a subscription.
type UserSubscriptionListResponseProductPricesProductPriceRecurring struct {
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
RecurringInterval UserSubscriptionListResponseProductPricesProductPriceRecurringRecurringInterval `json:"recurring_interval,required,nullable"`
// The type of the price.
Type UserSubscriptionListResponseProductPricesProductPriceRecurringType `json:"type,required"`
JSON userSubscriptionListResponseProductPricesProductPriceRecurringJSON `json:"-"`
}

// userSubscriptionListResponseProductPricesProductPriceRecurringJSON contains the
// JSON metadata for the struct
// [UserSubscriptionListResponseProductPricesProductPriceRecurring]
type userSubscriptionListResponseProductPricesProductPriceRecurringJSON struct {
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

func (r *UserSubscriptionListResponseProductPricesProductPriceRecurring) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userSubscriptionListResponseProductPricesProductPriceRecurringJSON) RawJSON() (string) {
  return r.raw
}

func (r UserSubscriptionListResponseProductPricesProductPriceRecurring) implementsUserSubscriptionListResponseProductPrice() {}

// The recurring interval of the price, if type is `recurring`.
type UserSubscriptionListResponseProductPricesProductPriceRecurringRecurringInterval string

const (
  UserSubscriptionListResponseProductPricesProductPriceRecurringRecurringIntervalMonth UserSubscriptionListResponseProductPricesProductPriceRecurringRecurringInterval = "month"
  UserSubscriptionListResponseProductPricesProductPriceRecurringRecurringIntervalYear UserSubscriptionListResponseProductPricesProductPriceRecurringRecurringInterval = "year"
)

func (r UserSubscriptionListResponseProductPricesProductPriceRecurringRecurringInterval) IsKnown() (bool) {
  switch r {
  case UserSubscriptionListResponseProductPricesProductPriceRecurringRecurringIntervalMonth, UserSubscriptionListResponseProductPricesProductPriceRecurringRecurringIntervalYear:
      return true
  }
  return false
}

// The type of the price.
type UserSubscriptionListResponseProductPricesProductPriceRecurringType string

const (
  UserSubscriptionListResponseProductPricesProductPriceRecurringTypeRecurring UserSubscriptionListResponseProductPricesProductPriceRecurringType = "recurring"
)

func (r UserSubscriptionListResponseProductPricesProductPriceRecurringType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionListResponseProductPricesProductPriceRecurringTypeRecurring:
      return true
  }
  return false
}

// A one-time price for a product.
type UserSubscriptionListResponseProductPricesProductPriceOneTime struct {
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
Type UserSubscriptionListResponseProductPricesProductPriceOneTimeType `json:"type,required"`
JSON userSubscriptionListResponseProductPricesProductPriceOneTimeJSON `json:"-"`
}

// userSubscriptionListResponseProductPricesProductPriceOneTimeJSON contains the
// JSON metadata for the struct
// [UserSubscriptionListResponseProductPricesProductPriceOneTime]
type userSubscriptionListResponseProductPricesProductPriceOneTimeJSON struct {
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

func (r *UserSubscriptionListResponseProductPricesProductPriceOneTime) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userSubscriptionListResponseProductPricesProductPriceOneTimeJSON) RawJSON() (string) {
  return r.raw
}

func (r UserSubscriptionListResponseProductPricesProductPriceOneTime) implementsUserSubscriptionListResponseProductPrice() {}

// The type of the price.
type UserSubscriptionListResponseProductPricesProductPriceOneTimeType string

const (
  UserSubscriptionListResponseProductPricesProductPriceOneTimeTypeOneTime UserSubscriptionListResponseProductPricesProductPriceOneTimeType = "one_time"
)

func (r UserSubscriptionListResponseProductPricesProductPriceOneTimeType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionListResponseProductPricesProductPriceOneTimeTypeOneTime:
      return true
  }
  return false
}

// The type of the price.
type UserSubscriptionListResponseProductPricesType string

const (
  UserSubscriptionListResponseProductPricesTypeRecurring UserSubscriptionListResponseProductPricesType = "recurring"
  UserSubscriptionListResponseProductPricesTypeOneTime UserSubscriptionListResponseProductPricesType = "one_time"
)

func (r UserSubscriptionListResponseProductPricesType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionListResponseProductPricesTypeRecurring, UserSubscriptionListResponseProductPricesTypeOneTime:
      return true
  }
  return false
}

// The recurring interval of the price, if type is `recurring`.
type UserSubscriptionListResponseProductPricesRecurringInterval string

const (
  UserSubscriptionListResponseProductPricesRecurringIntervalMonth UserSubscriptionListResponseProductPricesRecurringInterval = "month"
  UserSubscriptionListResponseProductPricesRecurringIntervalYear UserSubscriptionListResponseProductPricesRecurringInterval = "year"
)

func (r UserSubscriptionListResponseProductPricesRecurringInterval) IsKnown() (bool) {
  switch r {
  case UserSubscriptionListResponseProductPricesRecurringIntervalMonth, UserSubscriptionListResponseProductPricesRecurringIntervalYear:
      return true
  }
  return false
}

type UserSubscriptionListResponseProductType string

const (
  UserSubscriptionListResponseProductTypeFree UserSubscriptionListResponseProductType = "free"
  UserSubscriptionListResponseProductTypeIndividual UserSubscriptionListResponseProductType = "individual"
  UserSubscriptionListResponseProductTypeBusiness UserSubscriptionListResponseProductType = "business"
)

func (r UserSubscriptionListResponseProductType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionListResponseProductTypeFree, UserSubscriptionListResponseProductTypeIndividual, UserSubscriptionListResponseProductTypeBusiness:
      return true
  }
  return false
}

type UserSubscriptionListResponseStatus string

const (
  UserSubscriptionListResponseStatusIncomplete UserSubscriptionListResponseStatus = "incomplete"
  UserSubscriptionListResponseStatusIncompleteExpired UserSubscriptionListResponseStatus = "incomplete_expired"
  UserSubscriptionListResponseStatusTrialing UserSubscriptionListResponseStatus = "trialing"
  UserSubscriptionListResponseStatusActive UserSubscriptionListResponseStatus = "active"
  UserSubscriptionListResponseStatusPastDue UserSubscriptionListResponseStatus = "past_due"
  UserSubscriptionListResponseStatusCanceled UserSubscriptionListResponseStatus = "canceled"
  UserSubscriptionListResponseStatusUnpaid UserSubscriptionListResponseStatus = "unpaid"
)

func (r UserSubscriptionListResponseStatus) IsKnown() (bool) {
  switch r {
  case UserSubscriptionListResponseStatusIncomplete, UserSubscriptionListResponseStatusIncompleteExpired, UserSubscriptionListResponseStatusTrialing, UserSubscriptionListResponseStatusActive, UserSubscriptionListResponseStatusPastDue, UserSubscriptionListResponseStatusCanceled, UserSubscriptionListResponseStatusUnpaid:
      return true
  }
  return false
}

type UserSubscriptionDeleteResponse struct {
ID string `json:"id,required" format:"uuid4"`
CancelAtPeriodEnd bool `json:"cancel_at_period_end,required"`
// Creation timestamp of the object.
CreatedAt time.Time `json:"created_at,required" format:"date-time"`
CurrentPeriodEnd time.Time `json:"current_period_end,required,nullable" format:"date-time"`
CurrentPeriodStart time.Time `json:"current_period_start,required" format:"date-time"`
EndedAt time.Time `json:"ended_at,required,nullable" format:"date-time"`
// Last modification timestamp of the object.
ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
// A recurring price for a product, i.e. a subscription.
Price UserSubscriptionDeleteResponsePrice `json:"price,required,nullable"`
PriceID string `json:"price_id,required,nullable" format:"uuid4"`
Product UserSubscriptionDeleteResponseProduct `json:"product,required"`
ProductID string `json:"product_id,required" format:"uuid4"`
StartedAt time.Time `json:"started_at,required,nullable" format:"date-time"`
Status UserSubscriptionDeleteResponseStatus `json:"status,required"`
UserID string `json:"user_id,required" format:"uuid4"`
JSON userSubscriptionDeleteResponseJSON `json:"-"`
}

// userSubscriptionDeleteResponseJSON contains the JSON metadata for the struct
// [UserSubscriptionDeleteResponse]
type userSubscriptionDeleteResponseJSON struct {
ID apijson.Field
CancelAtPeriodEnd apijson.Field
CreatedAt apijson.Field
CurrentPeriodEnd apijson.Field
CurrentPeriodStart apijson.Field
EndedAt apijson.Field
ModifiedAt apijson.Field
Price apijson.Field
PriceID apijson.Field
Product apijson.Field
ProductID apijson.Field
StartedAt apijson.Field
Status apijson.Field
UserID apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionDeleteResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userSubscriptionDeleteResponseJSON) RawJSON() (string) {
  return r.raw
}

// A recurring price for a product, i.e. a subscription.
type UserSubscriptionDeleteResponsePrice struct {
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
Type UserSubscriptionDeleteResponsePriceType `json:"type,required"`
// The recurring interval of the price, if type is `recurring`.
RecurringInterval UserSubscriptionDeleteResponsePriceRecurringInterval `json:"recurring_interval,nullable"`
JSON userSubscriptionDeleteResponsePriceJSON `json:"-"`
union UserSubscriptionDeleteResponsePriceUnion
}

// userSubscriptionDeleteResponsePriceJSON contains the JSON metadata for the
// struct [UserSubscriptionDeleteResponsePrice]
type userSubscriptionDeleteResponsePriceJSON struct {
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

func (r userSubscriptionDeleteResponsePriceJSON) RawJSON() (string) {
  return r.raw
}

func (r *UserSubscriptionDeleteResponsePrice) UnmarshalJSON(data []byte) (err error) {
  *r = UserSubscriptionDeleteResponsePrice{}
  err = apijson.UnmarshalRoot(data, &r.union)
  if err != nil {
    return err
  }
  return apijson.Port(r.union, &r)
}

// AsUnion returns a [UserSubscriptionDeleteResponsePriceUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [UserSubscriptionDeleteResponsePriceProductPriceRecurring],
// [UserSubscriptionDeleteResponsePriceProductPriceOneTime].
func (r UserSubscriptionDeleteResponsePrice) AsUnion() (UserSubscriptionDeleteResponsePriceUnion) {
  return r.union
}

// A recurring price for a product, i.e. a subscription.
//
// Union satisfied by [UserSubscriptionDeleteResponsePriceProductPriceRecurring] or
// [UserSubscriptionDeleteResponsePriceProductPriceOneTime].
type UserSubscriptionDeleteResponsePriceUnion interface {
  implementsUserSubscriptionDeleteResponsePrice()
}

func init() {
  apijson.RegisterUnion(
    reflect.TypeOf((*UserSubscriptionDeleteResponsePriceUnion)(nil)).Elem(),
    "type",
    apijson.UnionVariant{
      TypeFilter: gjson.JSON,
      Type: reflect.TypeOf(UserSubscriptionDeleteResponsePriceProductPriceRecurring{}),
      DiscriminatorValue: "recurring",
    },
    apijson.UnionVariant{
      TypeFilter: gjson.JSON,
      Type: reflect.TypeOf(UserSubscriptionDeleteResponsePriceProductPriceOneTime{}),
      DiscriminatorValue: "one_time",
    },
  )
}

// A recurring price for a product, i.e. a subscription.
type UserSubscriptionDeleteResponsePriceProductPriceRecurring struct {
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
RecurringInterval UserSubscriptionDeleteResponsePriceProductPriceRecurringRecurringInterval `json:"recurring_interval,required,nullable"`
// The type of the price.
Type UserSubscriptionDeleteResponsePriceProductPriceRecurringType `json:"type,required"`
JSON userSubscriptionDeleteResponsePriceProductPriceRecurringJSON `json:"-"`
}

// userSubscriptionDeleteResponsePriceProductPriceRecurringJSON contains the JSON
// metadata for the struct
// [UserSubscriptionDeleteResponsePriceProductPriceRecurring]
type userSubscriptionDeleteResponsePriceProductPriceRecurringJSON struct {
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

func (r *UserSubscriptionDeleteResponsePriceProductPriceRecurring) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userSubscriptionDeleteResponsePriceProductPriceRecurringJSON) RawJSON() (string) {
  return r.raw
}

func (r UserSubscriptionDeleteResponsePriceProductPriceRecurring) implementsUserSubscriptionDeleteResponsePrice() {}

// The recurring interval of the price, if type is `recurring`.
type UserSubscriptionDeleteResponsePriceProductPriceRecurringRecurringInterval string

const (
  UserSubscriptionDeleteResponsePriceProductPriceRecurringRecurringIntervalMonth UserSubscriptionDeleteResponsePriceProductPriceRecurringRecurringInterval = "month"
  UserSubscriptionDeleteResponsePriceProductPriceRecurringRecurringIntervalYear UserSubscriptionDeleteResponsePriceProductPriceRecurringRecurringInterval = "year"
)

func (r UserSubscriptionDeleteResponsePriceProductPriceRecurringRecurringInterval) IsKnown() (bool) {
  switch r {
  case UserSubscriptionDeleteResponsePriceProductPriceRecurringRecurringIntervalMonth, UserSubscriptionDeleteResponsePriceProductPriceRecurringRecurringIntervalYear:
      return true
  }
  return false
}

// The type of the price.
type UserSubscriptionDeleteResponsePriceProductPriceRecurringType string

const (
  UserSubscriptionDeleteResponsePriceProductPriceRecurringTypeRecurring UserSubscriptionDeleteResponsePriceProductPriceRecurringType = "recurring"
)

func (r UserSubscriptionDeleteResponsePriceProductPriceRecurringType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionDeleteResponsePriceProductPriceRecurringTypeRecurring:
      return true
  }
  return false
}

// A one-time price for a product.
type UserSubscriptionDeleteResponsePriceProductPriceOneTime struct {
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
Type UserSubscriptionDeleteResponsePriceProductPriceOneTimeType `json:"type,required"`
JSON userSubscriptionDeleteResponsePriceProductPriceOneTimeJSON `json:"-"`
}

// userSubscriptionDeleteResponsePriceProductPriceOneTimeJSON contains the JSON
// metadata for the struct [UserSubscriptionDeleteResponsePriceProductPriceOneTime]
type userSubscriptionDeleteResponsePriceProductPriceOneTimeJSON struct {
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

func (r *UserSubscriptionDeleteResponsePriceProductPriceOneTime) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userSubscriptionDeleteResponsePriceProductPriceOneTimeJSON) RawJSON() (string) {
  return r.raw
}

func (r UserSubscriptionDeleteResponsePriceProductPriceOneTime) implementsUserSubscriptionDeleteResponsePrice() {}

// The type of the price.
type UserSubscriptionDeleteResponsePriceProductPriceOneTimeType string

const (
  UserSubscriptionDeleteResponsePriceProductPriceOneTimeTypeOneTime UserSubscriptionDeleteResponsePriceProductPriceOneTimeType = "one_time"
)

func (r UserSubscriptionDeleteResponsePriceProductPriceOneTimeType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionDeleteResponsePriceProductPriceOneTimeTypeOneTime:
      return true
  }
  return false
}

// The type of the price.
type UserSubscriptionDeleteResponsePriceType string

const (
  UserSubscriptionDeleteResponsePriceTypeRecurring UserSubscriptionDeleteResponsePriceType = "recurring"
  UserSubscriptionDeleteResponsePriceTypeOneTime UserSubscriptionDeleteResponsePriceType = "one_time"
)

func (r UserSubscriptionDeleteResponsePriceType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionDeleteResponsePriceTypeRecurring, UserSubscriptionDeleteResponsePriceTypeOneTime:
      return true
  }
  return false
}

// The recurring interval of the price, if type is `recurring`.
type UserSubscriptionDeleteResponsePriceRecurringInterval string

const (
  UserSubscriptionDeleteResponsePriceRecurringIntervalMonth UserSubscriptionDeleteResponsePriceRecurringInterval = "month"
  UserSubscriptionDeleteResponsePriceRecurringIntervalYear UserSubscriptionDeleteResponsePriceRecurringInterval = "year"
)

func (r UserSubscriptionDeleteResponsePriceRecurringInterval) IsKnown() (bool) {
  switch r {
  case UserSubscriptionDeleteResponsePriceRecurringIntervalMonth, UserSubscriptionDeleteResponsePriceRecurringIntervalYear:
      return true
  }
  return false
}

type UserSubscriptionDeleteResponseProduct struct {
// The ID of the product.
ID string `json:"id,required" format:"uuid4"`
// The benefits granted by the product.
Benefits []UserSubscriptionDeleteResponseProductBenefit `json:"benefits,required"`
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
Prices []UserSubscriptionDeleteResponseProductPrice `json:"prices,required"`
Type UserSubscriptionDeleteResponseProductType `json:"type,required,nullable"`
JSON userSubscriptionDeleteResponseProductJSON `json:"-"`
}

// userSubscriptionDeleteResponseProductJSON contains the JSON metadata for the
// struct [UserSubscriptionDeleteResponseProduct]
type userSubscriptionDeleteResponseProductJSON struct {
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

func (r *UserSubscriptionDeleteResponseProduct) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userSubscriptionDeleteResponseProductJSON) RawJSON() (string) {
  return r.raw
}

// A benefit of type `articles`.
//
// Use it to grant access to posts.
type UserSubscriptionDeleteResponseProductBenefit struct {
// Creation timestamp of the object.
CreatedAt time.Time `json:"created_at,required" format:"date-time"`
// Last modification timestamp of the object.
ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
// The ID of the benefit.
ID string `json:"id,required" format:"uuid4"`
// The type of the benefit.
Type UserSubscriptionDeleteResponseProductBenefitsType `json:"type,required"`
// The description of the benefit.
Description string `json:"description,required"`
// Whether the benefit is selectable when creating a product.
Selectable bool `json:"selectable,required"`
// Whether the benefit is deletable.
Deletable bool `json:"deletable,required"`
// The ID of the organization owning the benefit.
OrganizationID string `json:"organization_id,required" format:"uuid4"`
// This field can have the runtime type of
// [UserSubscriptionDeleteResponseProductBenefitsBenefitArticlesProperties].
Properties interface{} `json:"properties,required"`
JSON userSubscriptionDeleteResponseProductBenefitJSON `json:"-"`
union UserSubscriptionDeleteResponseProductBenefitsUnion
}

// userSubscriptionDeleteResponseProductBenefitJSON contains the JSON metadata for
// the struct [UserSubscriptionDeleteResponseProductBenefit]
type userSubscriptionDeleteResponseProductBenefitJSON struct {
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

func (r userSubscriptionDeleteResponseProductBenefitJSON) RawJSON() (string) {
  return r.raw
}

func (r *UserSubscriptionDeleteResponseProductBenefit) UnmarshalJSON(data []byte) (err error) {
  *r = UserSubscriptionDeleteResponseProductBenefit{}
  err = apijson.UnmarshalRoot(data, &r.union)
  if err != nil {
    return err
  }
  return apijson.Port(r.union, &r)
}

// AsUnion returns a [UserSubscriptionDeleteResponseProductBenefitsUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [UserSubscriptionDeleteResponseProductBenefitsBenefitBase],
// [UserSubscriptionDeleteResponseProductBenefitsBenefitArticles].
func (r UserSubscriptionDeleteResponseProductBenefit) AsUnion() (UserSubscriptionDeleteResponseProductBenefitsUnion) {
  return r.union
}

// A benefit of type `articles`.
//
// Use it to grant access to posts.
//
// Union satisfied by [UserSubscriptionDeleteResponseProductBenefitsBenefitBase] or
// [UserSubscriptionDeleteResponseProductBenefitsBenefitArticles].
type UserSubscriptionDeleteResponseProductBenefitsUnion interface {
  implementsUserSubscriptionDeleteResponseProductBenefit()
}

func init() {
  apijson.RegisterUnion(
    reflect.TypeOf((*UserSubscriptionDeleteResponseProductBenefitsUnion)(nil)).Elem(),
    "",
    apijson.UnionVariant{
      TypeFilter: gjson.JSON,
      Type: reflect.TypeOf(UserSubscriptionDeleteResponseProductBenefitsBenefitBase{}),
    },
    apijson.UnionVariant{
      TypeFilter: gjson.JSON,
      Type: reflect.TypeOf(UserSubscriptionDeleteResponseProductBenefitsBenefitArticles{}),
    },
  )
}

type UserSubscriptionDeleteResponseProductBenefitsBenefitBase struct {
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
Type UserSubscriptionDeleteResponseProductBenefitsBenefitBaseType `json:"type,required"`
JSON userSubscriptionDeleteResponseProductBenefitsBenefitBaseJSON `json:"-"`
}

// userSubscriptionDeleteResponseProductBenefitsBenefitBaseJSON contains the JSON
// metadata for the struct
// [UserSubscriptionDeleteResponseProductBenefitsBenefitBase]
type userSubscriptionDeleteResponseProductBenefitsBenefitBaseJSON struct {
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

func (r *UserSubscriptionDeleteResponseProductBenefitsBenefitBase) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userSubscriptionDeleteResponseProductBenefitsBenefitBaseJSON) RawJSON() (string) {
  return r.raw
}

func (r UserSubscriptionDeleteResponseProductBenefitsBenefitBase) implementsUserSubscriptionDeleteResponseProductBenefit() {}

// The type of the benefit.
type UserSubscriptionDeleteResponseProductBenefitsBenefitBaseType string

const (
  UserSubscriptionDeleteResponseProductBenefitsBenefitBaseTypeCustom UserSubscriptionDeleteResponseProductBenefitsBenefitBaseType = "custom"
  UserSubscriptionDeleteResponseProductBenefitsBenefitBaseTypeArticles UserSubscriptionDeleteResponseProductBenefitsBenefitBaseType = "articles"
  UserSubscriptionDeleteResponseProductBenefitsBenefitBaseTypeAds UserSubscriptionDeleteResponseProductBenefitsBenefitBaseType = "ads"
  UserSubscriptionDeleteResponseProductBenefitsBenefitBaseTypeDiscord UserSubscriptionDeleteResponseProductBenefitsBenefitBaseType = "discord"
  UserSubscriptionDeleteResponseProductBenefitsBenefitBaseTypeGitHubRepository UserSubscriptionDeleteResponseProductBenefitsBenefitBaseType = "github_repository"
  UserSubscriptionDeleteResponseProductBenefitsBenefitBaseTypeDownloadables UserSubscriptionDeleteResponseProductBenefitsBenefitBaseType = "downloadables"
)

func (r UserSubscriptionDeleteResponseProductBenefitsBenefitBaseType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionDeleteResponseProductBenefitsBenefitBaseTypeCustom, UserSubscriptionDeleteResponseProductBenefitsBenefitBaseTypeArticles, UserSubscriptionDeleteResponseProductBenefitsBenefitBaseTypeAds, UserSubscriptionDeleteResponseProductBenefitsBenefitBaseTypeDiscord, UserSubscriptionDeleteResponseProductBenefitsBenefitBaseTypeGitHubRepository, UserSubscriptionDeleteResponseProductBenefitsBenefitBaseTypeDownloadables:
      return true
  }
  return false
}

// A benefit of type `articles`.
//
// Use it to grant access to posts.
type UserSubscriptionDeleteResponseProductBenefitsBenefitArticles struct {
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
Properties UserSubscriptionDeleteResponseProductBenefitsBenefitArticlesProperties `json:"properties,required"`
// Whether the benefit is selectable when creating a product.
Selectable bool `json:"selectable,required"`
Type UserSubscriptionDeleteResponseProductBenefitsBenefitArticlesType `json:"type,required"`
JSON userSubscriptionDeleteResponseProductBenefitsBenefitArticlesJSON `json:"-"`
}

// userSubscriptionDeleteResponseProductBenefitsBenefitArticlesJSON contains the
// JSON metadata for the struct
// [UserSubscriptionDeleteResponseProductBenefitsBenefitArticles]
type userSubscriptionDeleteResponseProductBenefitsBenefitArticlesJSON struct {
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

func (r *UserSubscriptionDeleteResponseProductBenefitsBenefitArticles) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userSubscriptionDeleteResponseProductBenefitsBenefitArticlesJSON) RawJSON() (string) {
  return r.raw
}

func (r UserSubscriptionDeleteResponseProductBenefitsBenefitArticles) implementsUserSubscriptionDeleteResponseProductBenefit() {}

// Properties for a benefit of type `articles`.
type UserSubscriptionDeleteResponseProductBenefitsBenefitArticlesProperties struct {
// Whether the user can access paid articles.
PaidArticles bool `json:"paid_articles,required"`
JSON userSubscriptionDeleteResponseProductBenefitsBenefitArticlesPropertiesJSON `json:"-"`
}

// userSubscriptionDeleteResponseProductBenefitsBenefitArticlesPropertiesJSON
// contains the JSON metadata for the struct
// [UserSubscriptionDeleteResponseProductBenefitsBenefitArticlesProperties]
type userSubscriptionDeleteResponseProductBenefitsBenefitArticlesPropertiesJSON struct {
PaidArticles apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionDeleteResponseProductBenefitsBenefitArticlesProperties) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userSubscriptionDeleteResponseProductBenefitsBenefitArticlesPropertiesJSON) RawJSON() (string) {
  return r.raw
}

type UserSubscriptionDeleteResponseProductBenefitsBenefitArticlesType string

const (
  UserSubscriptionDeleteResponseProductBenefitsBenefitArticlesTypeArticles UserSubscriptionDeleteResponseProductBenefitsBenefitArticlesType = "articles"
)

func (r UserSubscriptionDeleteResponseProductBenefitsBenefitArticlesType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionDeleteResponseProductBenefitsBenefitArticlesTypeArticles:
      return true
  }
  return false
}

// The type of the benefit.
type UserSubscriptionDeleteResponseProductBenefitsType string

const (
  UserSubscriptionDeleteResponseProductBenefitsTypeCustom UserSubscriptionDeleteResponseProductBenefitsType = "custom"
  UserSubscriptionDeleteResponseProductBenefitsTypeArticles UserSubscriptionDeleteResponseProductBenefitsType = "articles"
  UserSubscriptionDeleteResponseProductBenefitsTypeAds UserSubscriptionDeleteResponseProductBenefitsType = "ads"
  UserSubscriptionDeleteResponseProductBenefitsTypeDiscord UserSubscriptionDeleteResponseProductBenefitsType = "discord"
  UserSubscriptionDeleteResponseProductBenefitsTypeGitHubRepository UserSubscriptionDeleteResponseProductBenefitsType = "github_repository"
  UserSubscriptionDeleteResponseProductBenefitsTypeDownloadables UserSubscriptionDeleteResponseProductBenefitsType = "downloadables"
)

func (r UserSubscriptionDeleteResponseProductBenefitsType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionDeleteResponseProductBenefitsTypeCustom, UserSubscriptionDeleteResponseProductBenefitsTypeArticles, UserSubscriptionDeleteResponseProductBenefitsTypeAds, UserSubscriptionDeleteResponseProductBenefitsTypeDiscord, UserSubscriptionDeleteResponseProductBenefitsTypeGitHubRepository, UserSubscriptionDeleteResponseProductBenefitsTypeDownloadables:
      return true
  }
  return false
}

// A recurring price for a product, i.e. a subscription.
type UserSubscriptionDeleteResponseProductPrice struct {
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
Type UserSubscriptionDeleteResponseProductPricesType `json:"type,required"`
// The recurring interval of the price, if type is `recurring`.
RecurringInterval UserSubscriptionDeleteResponseProductPricesRecurringInterval `json:"recurring_interval,nullable"`
JSON userSubscriptionDeleteResponseProductPriceJSON `json:"-"`
union UserSubscriptionDeleteResponseProductPricesUnion
}

// userSubscriptionDeleteResponseProductPriceJSON contains the JSON metadata for
// the struct [UserSubscriptionDeleteResponseProductPrice]
type userSubscriptionDeleteResponseProductPriceJSON struct {
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

func (r userSubscriptionDeleteResponseProductPriceJSON) RawJSON() (string) {
  return r.raw
}

func (r *UserSubscriptionDeleteResponseProductPrice) UnmarshalJSON(data []byte) (err error) {
  *r = UserSubscriptionDeleteResponseProductPrice{}
  err = apijson.UnmarshalRoot(data, &r.union)
  if err != nil {
    return err
  }
  return apijson.Port(r.union, &r)
}

// AsUnion returns a [UserSubscriptionDeleteResponseProductPricesUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [UserSubscriptionDeleteResponseProductPricesProductPriceRecurring],
// [UserSubscriptionDeleteResponseProductPricesProductPriceOneTime].
func (r UserSubscriptionDeleteResponseProductPrice) AsUnion() (UserSubscriptionDeleteResponseProductPricesUnion) {
  return r.union
}

// A recurring price for a product, i.e. a subscription.
//
// Union satisfied by
// [UserSubscriptionDeleteResponseProductPricesProductPriceRecurring] or
// [UserSubscriptionDeleteResponseProductPricesProductPriceOneTime].
type UserSubscriptionDeleteResponseProductPricesUnion interface {
  implementsUserSubscriptionDeleteResponseProductPrice()
}

func init() {
  apijson.RegisterUnion(
    reflect.TypeOf((*UserSubscriptionDeleteResponseProductPricesUnion)(nil)).Elem(),
    "type",
    apijson.UnionVariant{
      TypeFilter: gjson.JSON,
      Type: reflect.TypeOf(UserSubscriptionDeleteResponseProductPricesProductPriceRecurring{}),
      DiscriminatorValue: "recurring",
    },
    apijson.UnionVariant{
      TypeFilter: gjson.JSON,
      Type: reflect.TypeOf(UserSubscriptionDeleteResponseProductPricesProductPriceOneTime{}),
      DiscriminatorValue: "one_time",
    },
  )
}

// A recurring price for a product, i.e. a subscription.
type UserSubscriptionDeleteResponseProductPricesProductPriceRecurring struct {
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
RecurringInterval UserSubscriptionDeleteResponseProductPricesProductPriceRecurringRecurringInterval `json:"recurring_interval,required,nullable"`
// The type of the price.
Type UserSubscriptionDeleteResponseProductPricesProductPriceRecurringType `json:"type,required"`
JSON userSubscriptionDeleteResponseProductPricesProductPriceRecurringJSON `json:"-"`
}

// userSubscriptionDeleteResponseProductPricesProductPriceRecurringJSON contains
// the JSON metadata for the struct
// [UserSubscriptionDeleteResponseProductPricesProductPriceRecurring]
type userSubscriptionDeleteResponseProductPricesProductPriceRecurringJSON struct {
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

func (r *UserSubscriptionDeleteResponseProductPricesProductPriceRecurring) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userSubscriptionDeleteResponseProductPricesProductPriceRecurringJSON) RawJSON() (string) {
  return r.raw
}

func (r UserSubscriptionDeleteResponseProductPricesProductPriceRecurring) implementsUserSubscriptionDeleteResponseProductPrice() {}

// The recurring interval of the price, if type is `recurring`.
type UserSubscriptionDeleteResponseProductPricesProductPriceRecurringRecurringInterval string

const (
  UserSubscriptionDeleteResponseProductPricesProductPriceRecurringRecurringIntervalMonth UserSubscriptionDeleteResponseProductPricesProductPriceRecurringRecurringInterval = "month"
  UserSubscriptionDeleteResponseProductPricesProductPriceRecurringRecurringIntervalYear UserSubscriptionDeleteResponseProductPricesProductPriceRecurringRecurringInterval = "year"
)

func (r UserSubscriptionDeleteResponseProductPricesProductPriceRecurringRecurringInterval) IsKnown() (bool) {
  switch r {
  case UserSubscriptionDeleteResponseProductPricesProductPriceRecurringRecurringIntervalMonth, UserSubscriptionDeleteResponseProductPricesProductPriceRecurringRecurringIntervalYear:
      return true
  }
  return false
}

// The type of the price.
type UserSubscriptionDeleteResponseProductPricesProductPriceRecurringType string

const (
  UserSubscriptionDeleteResponseProductPricesProductPriceRecurringTypeRecurring UserSubscriptionDeleteResponseProductPricesProductPriceRecurringType = "recurring"
)

func (r UserSubscriptionDeleteResponseProductPricesProductPriceRecurringType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionDeleteResponseProductPricesProductPriceRecurringTypeRecurring:
      return true
  }
  return false
}

// A one-time price for a product.
type UserSubscriptionDeleteResponseProductPricesProductPriceOneTime struct {
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
Type UserSubscriptionDeleteResponseProductPricesProductPriceOneTimeType `json:"type,required"`
JSON userSubscriptionDeleteResponseProductPricesProductPriceOneTimeJSON `json:"-"`
}

// userSubscriptionDeleteResponseProductPricesProductPriceOneTimeJSON contains the
// JSON metadata for the struct
// [UserSubscriptionDeleteResponseProductPricesProductPriceOneTime]
type userSubscriptionDeleteResponseProductPricesProductPriceOneTimeJSON struct {
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

func (r *UserSubscriptionDeleteResponseProductPricesProductPriceOneTime) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userSubscriptionDeleteResponseProductPricesProductPriceOneTimeJSON) RawJSON() (string) {
  return r.raw
}

func (r UserSubscriptionDeleteResponseProductPricesProductPriceOneTime) implementsUserSubscriptionDeleteResponseProductPrice() {}

// The type of the price.
type UserSubscriptionDeleteResponseProductPricesProductPriceOneTimeType string

const (
  UserSubscriptionDeleteResponseProductPricesProductPriceOneTimeTypeOneTime UserSubscriptionDeleteResponseProductPricesProductPriceOneTimeType = "one_time"
)

func (r UserSubscriptionDeleteResponseProductPricesProductPriceOneTimeType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionDeleteResponseProductPricesProductPriceOneTimeTypeOneTime:
      return true
  }
  return false
}

// The type of the price.
type UserSubscriptionDeleteResponseProductPricesType string

const (
  UserSubscriptionDeleteResponseProductPricesTypeRecurring UserSubscriptionDeleteResponseProductPricesType = "recurring"
  UserSubscriptionDeleteResponseProductPricesTypeOneTime UserSubscriptionDeleteResponseProductPricesType = "one_time"
)

func (r UserSubscriptionDeleteResponseProductPricesType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionDeleteResponseProductPricesTypeRecurring, UserSubscriptionDeleteResponseProductPricesTypeOneTime:
      return true
  }
  return false
}

// The recurring interval of the price, if type is `recurring`.
type UserSubscriptionDeleteResponseProductPricesRecurringInterval string

const (
  UserSubscriptionDeleteResponseProductPricesRecurringIntervalMonth UserSubscriptionDeleteResponseProductPricesRecurringInterval = "month"
  UserSubscriptionDeleteResponseProductPricesRecurringIntervalYear UserSubscriptionDeleteResponseProductPricesRecurringInterval = "year"
)

func (r UserSubscriptionDeleteResponseProductPricesRecurringInterval) IsKnown() (bool) {
  switch r {
  case UserSubscriptionDeleteResponseProductPricesRecurringIntervalMonth, UserSubscriptionDeleteResponseProductPricesRecurringIntervalYear:
      return true
  }
  return false
}

type UserSubscriptionDeleteResponseProductType string

const (
  UserSubscriptionDeleteResponseProductTypeFree UserSubscriptionDeleteResponseProductType = "free"
  UserSubscriptionDeleteResponseProductTypeIndividual UserSubscriptionDeleteResponseProductType = "individual"
  UserSubscriptionDeleteResponseProductTypeBusiness UserSubscriptionDeleteResponseProductType = "business"
)

func (r UserSubscriptionDeleteResponseProductType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionDeleteResponseProductTypeFree, UserSubscriptionDeleteResponseProductTypeIndividual, UserSubscriptionDeleteResponseProductTypeBusiness:
      return true
  }
  return false
}

type UserSubscriptionDeleteResponseStatus string

const (
  UserSubscriptionDeleteResponseStatusIncomplete UserSubscriptionDeleteResponseStatus = "incomplete"
  UserSubscriptionDeleteResponseStatusIncompleteExpired UserSubscriptionDeleteResponseStatus = "incomplete_expired"
  UserSubscriptionDeleteResponseStatusTrialing UserSubscriptionDeleteResponseStatus = "trialing"
  UserSubscriptionDeleteResponseStatusActive UserSubscriptionDeleteResponseStatus = "active"
  UserSubscriptionDeleteResponseStatusPastDue UserSubscriptionDeleteResponseStatus = "past_due"
  UserSubscriptionDeleteResponseStatusCanceled UserSubscriptionDeleteResponseStatus = "canceled"
  UserSubscriptionDeleteResponseStatusUnpaid UserSubscriptionDeleteResponseStatus = "unpaid"
)

func (r UserSubscriptionDeleteResponseStatus) IsKnown() (bool) {
  switch r {
  case UserSubscriptionDeleteResponseStatusIncomplete, UserSubscriptionDeleteResponseStatusIncompleteExpired, UserSubscriptionDeleteResponseStatusTrialing, UserSubscriptionDeleteResponseStatusActive, UserSubscriptionDeleteResponseStatusPastDue, UserSubscriptionDeleteResponseStatusCanceled, UserSubscriptionDeleteResponseStatusUnpaid:
      return true
  }
  return false
}

type UserSubscriptionNewParams struct {
// ID of the free tier to subscribe to.
ProductID param.Field[string] `json:"product_id,required" format:"uuid4"`
// Email of the customer. This field is required if the API is called outside the
// Polar app.
CustomerEmail param.Field[string] `json:"customer_email" format:"email"`
}

func (r UserSubscriptionNewParams) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type UserSubscriptionUpdateParams struct {
ProductPriceID param.Field[string] `json:"product_price_id,required" format:"uuid4"`
}

func (r UserSubscriptionUpdateParams) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type UserSubscriptionListParams struct {
// Filter by active or cancelled subscription.
Active param.Field[bool] `query:"active"`
// Size of a page, defaults to 10. Maximum is 100.
Limit param.Field[int64] `query:"limit"`
// Filter by organization ID.
OrganizationID param.Field[UserSubscriptionListParamsOrganizationIDUnion] `query:"organization_id" format:"uuid4"`
// Page number, defaults to 1.
Page param.Field[int64] `query:"page"`
// Filter by product ID.
ProductID param.Field[UserSubscriptionListParamsProductIDUnion] `query:"product_id" format:"uuid4"`
// Search by product or organization name.
Query param.Field[string] `query:"query"`
// Sorting criterion. Several criteria can be used simultaneously and will be
// applied in order. Add a minus sign `-` before the criteria name to sort by
// descending order.
Sorting param.Field[[]UserSubscriptionListParamsSorting] `query:"sorting"`
}

// URLQuery serializes [UserSubscriptionListParams]'s query parameters as
// `url.Values`.
func (r UserSubscriptionListParams) URLQuery() (v url.Values) {
  return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
    ArrayFormat: apiquery.ArrayQueryFormatComma,
    NestedFormat: apiquery.NestedQueryFormatBrackets,
  })
}

// Filter by organization ID.
//
// Satisfied by [shared.UnionString],
// [UserSubscriptionListParamsOrganizationIDArray].
type UserSubscriptionListParamsOrganizationIDUnion interface {
  ImplementsUserSubscriptionListParamsOrganizationIDUnion()
}

type UserSubscriptionListParamsOrganizationIDArray []string

func (r UserSubscriptionListParamsOrganizationIDArray) ImplementsUserSubscriptionListParamsOrganizationIDUnion() {}

// Filter by product ID.
//
// Satisfied by [shared.UnionString], [UserSubscriptionListParamsProductIDArray].
type UserSubscriptionListParamsProductIDUnion interface {
  ImplementsUserSubscriptionListParamsProductIDUnion()
}

type UserSubscriptionListParamsProductIDArray []string

func (r UserSubscriptionListParamsProductIDArray) ImplementsUserSubscriptionListParamsProductIDUnion() {}

type UserSubscriptionListParamsSorting string

const (
  UserSubscriptionListParamsSortingStartedAt UserSubscriptionListParamsSorting = "started_at"
  UserSubscriptionListParamsSorting-StartedAt UserSubscriptionListParamsSorting = "-started_at"
  UserSubscriptionListParamsSortingPriceAmount UserSubscriptionListParamsSorting = "price_amount"
  UserSubscriptionListParamsSorting-PriceAmount UserSubscriptionListParamsSorting = "-price_amount"
  UserSubscriptionListParamsSortingStatus UserSubscriptionListParamsSorting = "status"
  UserSubscriptionListParamsSorting-Status UserSubscriptionListParamsSorting = "-status"
  UserSubscriptionListParamsSortingOrganization UserSubscriptionListParamsSorting = "organization"
  UserSubscriptionListParamsSorting-Organization UserSubscriptionListParamsSorting = "-organization"
  UserSubscriptionListParamsSortingProduct UserSubscriptionListParamsSorting = "product"
  UserSubscriptionListParamsSorting-Product UserSubscriptionListParamsSorting = "-product"
)

func (r UserSubscriptionListParamsSorting) IsKnown() (bool) {
  switch r {
  case UserSubscriptionListParamsSortingStartedAt, UserSubscriptionListParamsSorting-StartedAt, UserSubscriptionListParamsSortingPriceAmount, UserSubscriptionListParamsSorting-PriceAmount, UserSubscriptionListParamsSortingStatus, UserSubscriptionListParamsSorting-Status, UserSubscriptionListParamsSortingOrganization, UserSubscriptionListParamsSorting-Organization, UserSubscriptionListParamsSortingProduct, UserSubscriptionListParamsSorting-Product:
      return true
  }
  return false
}
