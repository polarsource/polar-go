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
Price ProductPrice `json:"price,required,nullable"`
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
Prices []ProductPrice `json:"prices,required"`
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
  UserSubscriptionNewResponseProductBenefitsBenefitBaseTypeLicenseKeys UserSubscriptionNewResponseProductBenefitsBenefitBaseType = "license_keys"
)

func (r UserSubscriptionNewResponseProductBenefitsBenefitBaseType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionNewResponseProductBenefitsBenefitBaseTypeCustom, UserSubscriptionNewResponseProductBenefitsBenefitBaseTypeArticles, UserSubscriptionNewResponseProductBenefitsBenefitBaseTypeAds, UserSubscriptionNewResponseProductBenefitsBenefitBaseTypeDiscord, UserSubscriptionNewResponseProductBenefitsBenefitBaseTypeGitHubRepository, UserSubscriptionNewResponseProductBenefitsBenefitBaseTypeDownloadables, UserSubscriptionNewResponseProductBenefitsBenefitBaseTypeLicenseKeys:
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
  UserSubscriptionNewResponseProductBenefitsTypeLicenseKeys UserSubscriptionNewResponseProductBenefitsType = "license_keys"
)

func (r UserSubscriptionNewResponseProductBenefitsType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionNewResponseProductBenefitsTypeCustom, UserSubscriptionNewResponseProductBenefitsTypeArticles, UserSubscriptionNewResponseProductBenefitsTypeAds, UserSubscriptionNewResponseProductBenefitsTypeDiscord, UserSubscriptionNewResponseProductBenefitsTypeGitHubRepository, UserSubscriptionNewResponseProductBenefitsTypeDownloadables, UserSubscriptionNewResponseProductBenefitsTypeLicenseKeys:
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
Price ProductPrice `json:"price,required,nullable"`
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
Prices []ProductPrice `json:"prices,required"`
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
  UserSubscriptionGetResponseProductBenefitsBenefitBaseTypeLicenseKeys UserSubscriptionGetResponseProductBenefitsBenefitBaseType = "license_keys"
)

func (r UserSubscriptionGetResponseProductBenefitsBenefitBaseType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionGetResponseProductBenefitsBenefitBaseTypeCustom, UserSubscriptionGetResponseProductBenefitsBenefitBaseTypeArticles, UserSubscriptionGetResponseProductBenefitsBenefitBaseTypeAds, UserSubscriptionGetResponseProductBenefitsBenefitBaseTypeDiscord, UserSubscriptionGetResponseProductBenefitsBenefitBaseTypeGitHubRepository, UserSubscriptionGetResponseProductBenefitsBenefitBaseTypeDownloadables, UserSubscriptionGetResponseProductBenefitsBenefitBaseTypeLicenseKeys:
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
  UserSubscriptionGetResponseProductBenefitsTypeLicenseKeys UserSubscriptionGetResponseProductBenefitsType = "license_keys"
)

func (r UserSubscriptionGetResponseProductBenefitsType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionGetResponseProductBenefitsTypeCustom, UserSubscriptionGetResponseProductBenefitsTypeArticles, UserSubscriptionGetResponseProductBenefitsTypeAds, UserSubscriptionGetResponseProductBenefitsTypeDiscord, UserSubscriptionGetResponseProductBenefitsTypeGitHubRepository, UserSubscriptionGetResponseProductBenefitsTypeDownloadables, UserSubscriptionGetResponseProductBenefitsTypeLicenseKeys:
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
Price ProductPrice `json:"price,required,nullable"`
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
Prices []ProductPrice `json:"prices,required"`
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
  UserSubscriptionUpdateResponseProductBenefitsBenefitBaseTypeLicenseKeys UserSubscriptionUpdateResponseProductBenefitsBenefitBaseType = "license_keys"
)

func (r UserSubscriptionUpdateResponseProductBenefitsBenefitBaseType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionUpdateResponseProductBenefitsBenefitBaseTypeCustom, UserSubscriptionUpdateResponseProductBenefitsBenefitBaseTypeArticles, UserSubscriptionUpdateResponseProductBenefitsBenefitBaseTypeAds, UserSubscriptionUpdateResponseProductBenefitsBenefitBaseTypeDiscord, UserSubscriptionUpdateResponseProductBenefitsBenefitBaseTypeGitHubRepository, UserSubscriptionUpdateResponseProductBenefitsBenefitBaseTypeDownloadables, UserSubscriptionUpdateResponseProductBenefitsBenefitBaseTypeLicenseKeys:
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
  UserSubscriptionUpdateResponseProductBenefitsTypeLicenseKeys UserSubscriptionUpdateResponseProductBenefitsType = "license_keys"
)

func (r UserSubscriptionUpdateResponseProductBenefitsType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionUpdateResponseProductBenefitsTypeCustom, UserSubscriptionUpdateResponseProductBenefitsTypeArticles, UserSubscriptionUpdateResponseProductBenefitsTypeAds, UserSubscriptionUpdateResponseProductBenefitsTypeDiscord, UserSubscriptionUpdateResponseProductBenefitsTypeGitHubRepository, UserSubscriptionUpdateResponseProductBenefitsTypeDownloadables, UserSubscriptionUpdateResponseProductBenefitsTypeLicenseKeys:
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
Price ProductPrice `json:"price,required,nullable"`
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
Prices []ProductPrice `json:"prices,required"`
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
  UserSubscriptionListResponseProductBenefitsBenefitBaseTypeLicenseKeys UserSubscriptionListResponseProductBenefitsBenefitBaseType = "license_keys"
)

func (r UserSubscriptionListResponseProductBenefitsBenefitBaseType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionListResponseProductBenefitsBenefitBaseTypeCustom, UserSubscriptionListResponseProductBenefitsBenefitBaseTypeArticles, UserSubscriptionListResponseProductBenefitsBenefitBaseTypeAds, UserSubscriptionListResponseProductBenefitsBenefitBaseTypeDiscord, UserSubscriptionListResponseProductBenefitsBenefitBaseTypeGitHubRepository, UserSubscriptionListResponseProductBenefitsBenefitBaseTypeDownloadables, UserSubscriptionListResponseProductBenefitsBenefitBaseTypeLicenseKeys:
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
  UserSubscriptionListResponseProductBenefitsTypeLicenseKeys UserSubscriptionListResponseProductBenefitsType = "license_keys"
)

func (r UserSubscriptionListResponseProductBenefitsType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionListResponseProductBenefitsTypeCustom, UserSubscriptionListResponseProductBenefitsTypeArticles, UserSubscriptionListResponseProductBenefitsTypeAds, UserSubscriptionListResponseProductBenefitsTypeDiscord, UserSubscriptionListResponseProductBenefitsTypeGitHubRepository, UserSubscriptionListResponseProductBenefitsTypeDownloadables, UserSubscriptionListResponseProductBenefitsTypeLicenseKeys:
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
Price ProductPrice `json:"price,required,nullable"`
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
Prices []ProductPrice `json:"prices,required"`
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
  UserSubscriptionDeleteResponseProductBenefitsBenefitBaseTypeLicenseKeys UserSubscriptionDeleteResponseProductBenefitsBenefitBaseType = "license_keys"
)

func (r UserSubscriptionDeleteResponseProductBenefitsBenefitBaseType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionDeleteResponseProductBenefitsBenefitBaseTypeCustom, UserSubscriptionDeleteResponseProductBenefitsBenefitBaseTypeArticles, UserSubscriptionDeleteResponseProductBenefitsBenefitBaseTypeAds, UserSubscriptionDeleteResponseProductBenefitsBenefitBaseTypeDiscord, UserSubscriptionDeleteResponseProductBenefitsBenefitBaseTypeGitHubRepository, UserSubscriptionDeleteResponseProductBenefitsBenefitBaseTypeDownloadables, UserSubscriptionDeleteResponseProductBenefitsBenefitBaseTypeLicenseKeys:
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
  UserSubscriptionDeleteResponseProductBenefitsTypeLicenseKeys UserSubscriptionDeleteResponseProductBenefitsType = "license_keys"
)

func (r UserSubscriptionDeleteResponseProductBenefitsType) IsKnown() (bool) {
  switch r {
  case UserSubscriptionDeleteResponseProductBenefitsTypeCustom, UserSubscriptionDeleteResponseProductBenefitsTypeArticles, UserSubscriptionDeleteResponseProductBenefitsTypeAds, UserSubscriptionDeleteResponseProductBenefitsTypeDiscord, UserSubscriptionDeleteResponseProductBenefitsTypeGitHubRepository, UserSubscriptionDeleteResponseProductBenefitsTypeDownloadables, UserSubscriptionDeleteResponseProductBenefitsTypeLicenseKeys:
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
