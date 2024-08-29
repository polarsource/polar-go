// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package polar

import (
  "bytes"
  "context"
  "io"
  "mime/multipart"
  "net/http"
  "net/url"
  "time"

  "github.com/polarsource/polar-go/internal/apiform"
  "github.com/polarsource/polar-go/internal/apijson"
  "github.com/polarsource/polar-go/internal/apiquery"
  "github.com/polarsource/polar-go/internal/pagination"
  "github.com/polarsource/polar-go/internal/param"
  "github.com/polarsource/polar-go/internal/requestconfig"
  "github.com/polarsource/polar-go/option"
)

// SubscriptionService contains methods and other services that help with
// interacting with the polar API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSubscriptionService] method instead.
type SubscriptionService struct {
Options []option.RequestOption
}

// NewSubscriptionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSubscriptionService(opts ...option.RequestOption) (r *SubscriptionService) {
  r = &SubscriptionService{}
  r.Options = opts
  return
}

// Create a subscription on the free tier for a given email.
func (r *SubscriptionService) New(ctx context.Context, body SubscriptionNewParams, opts ...option.RequestOption) (res *SubscriptionNewResponse, err error) {
  opts = append(r.Options[:], opts...)
  path := "v1/subscriptions/"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
  return
}

// List subscriptions.
func (r *SubscriptionService) List(ctx context.Context, query SubscriptionListParams, opts ...option.RequestOption) (res *pagination.PolarPagination[SubscriptionListResponse], err error) {
  var raw *http.Response
  opts = append(r.Options[:], opts...)
  opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
  path := "v1/subscriptions/"
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

// List subscriptions.
func (r *SubscriptionService) ListAutoPaging(ctx context.Context, query SubscriptionListParams, opts ...option.RequestOption) (*pagination.PolarPaginationAutoPager[SubscriptionListResponse]) {
  return pagination.NewPolarPaginationAutoPager(r.List(ctx, query, opts...))
}

// Export subscriptions as a CSV file.
func (r *SubscriptionService) Export(ctx context.Context, query SubscriptionExportParams, opts ...option.RequestOption) (res *SubscriptionExportResponse, err error) {
  opts = append(r.Options[:], opts...)
  path := "v1/subscriptions/export"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
  return
}

// Import subscriptions from a CSV file.
func (r *SubscriptionService) Import(ctx context.Context, body SubscriptionImportParams, opts ...option.RequestOption) (res *SubscriptionImportResponse, err error) {
  opts = append(r.Options[:], opts...)
  path := "v1/subscriptions/import"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
  return
}

type SubscriptionNewResponse struct {
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
// A recurring price for a product, i.e. a subscription.
Price ProductPrice `json:"price,required,nullable"`
PriceID string `json:"price_id,required,nullable" format:"uuid4"`
// A product.
Product Product `json:"product,required"`
ProductID string `json:"product_id,required" format:"uuid4"`
StartedAt time.Time `json:"started_at,required,nullable" format:"date-time"`
Status SubscriptionNewResponseStatus `json:"status,required"`
User SubscriptionNewResponseUser `json:"user,required"`
UserID string `json:"user_id,required" format:"uuid4"`
JSON subscriptionNewResponseJSON `json:"-"`
}

// subscriptionNewResponseJSON contains the JSON metadata for the struct
// [SubscriptionNewResponse]
type subscriptionNewResponseJSON struct {
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
User apijson.Field
UserID apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *SubscriptionNewResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionNewResponseJSON) RawJSON() (string) {
  return r.raw
}

type SubscriptionNewResponseStatus string

const (
  SubscriptionNewResponseStatusIncomplete SubscriptionNewResponseStatus = "incomplete"
  SubscriptionNewResponseStatusIncompleteExpired SubscriptionNewResponseStatus = "incomplete_expired"
  SubscriptionNewResponseStatusTrialing SubscriptionNewResponseStatus = "trialing"
  SubscriptionNewResponseStatusActive SubscriptionNewResponseStatus = "active"
  SubscriptionNewResponseStatusPastDue SubscriptionNewResponseStatus = "past_due"
  SubscriptionNewResponseStatusCanceled SubscriptionNewResponseStatus = "canceled"
  SubscriptionNewResponseStatusUnpaid SubscriptionNewResponseStatus = "unpaid"
)

func (r SubscriptionNewResponseStatus) IsKnown() (bool) {
  switch r {
  case SubscriptionNewResponseStatusIncomplete, SubscriptionNewResponseStatusIncompleteExpired, SubscriptionNewResponseStatusTrialing, SubscriptionNewResponseStatusActive, SubscriptionNewResponseStatusPastDue, SubscriptionNewResponseStatusCanceled, SubscriptionNewResponseStatusUnpaid:
      return true
  }
  return false
}

type SubscriptionNewResponseUser struct {
AvatarURL string `json:"avatar_url,required,nullable"`
Email string `json:"email,required"`
GitHubUsername string `json:"github_username,required,nullable"`
PublicName string `json:"public_name,required"`
JSON subscriptionNewResponseUserJSON `json:"-"`
}

// subscriptionNewResponseUserJSON contains the JSON metadata for the struct
// [SubscriptionNewResponseUser]
type subscriptionNewResponseUserJSON struct {
AvatarURL apijson.Field
Email apijson.Field
GitHubUsername apijson.Field
PublicName apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *SubscriptionNewResponseUser) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionNewResponseUserJSON) RawJSON() (string) {
  return r.raw
}

type SubscriptionListResponse struct {
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
// A recurring price for a product, i.e. a subscription.
Price ProductPrice `json:"price,required,nullable"`
PriceID string `json:"price_id,required,nullable" format:"uuid4"`
// A product.
Product Product `json:"product,required"`
ProductID string `json:"product_id,required" format:"uuid4"`
StartedAt time.Time `json:"started_at,required,nullable" format:"date-time"`
Status SubscriptionListResponseStatus `json:"status,required"`
User SubscriptionListResponseUser `json:"user,required"`
UserID string `json:"user_id,required" format:"uuid4"`
JSON subscriptionListResponseJSON `json:"-"`
}

// subscriptionListResponseJSON contains the JSON metadata for the struct
// [SubscriptionListResponse]
type subscriptionListResponseJSON struct {
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
User apijson.Field
UserID apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *SubscriptionListResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionListResponseJSON) RawJSON() (string) {
  return r.raw
}

type SubscriptionListResponseStatus string

const (
  SubscriptionListResponseStatusIncomplete SubscriptionListResponseStatus = "incomplete"
  SubscriptionListResponseStatusIncompleteExpired SubscriptionListResponseStatus = "incomplete_expired"
  SubscriptionListResponseStatusTrialing SubscriptionListResponseStatus = "trialing"
  SubscriptionListResponseStatusActive SubscriptionListResponseStatus = "active"
  SubscriptionListResponseStatusPastDue SubscriptionListResponseStatus = "past_due"
  SubscriptionListResponseStatusCanceled SubscriptionListResponseStatus = "canceled"
  SubscriptionListResponseStatusUnpaid SubscriptionListResponseStatus = "unpaid"
)

func (r SubscriptionListResponseStatus) IsKnown() (bool) {
  switch r {
  case SubscriptionListResponseStatusIncomplete, SubscriptionListResponseStatusIncompleteExpired, SubscriptionListResponseStatusTrialing, SubscriptionListResponseStatusActive, SubscriptionListResponseStatusPastDue, SubscriptionListResponseStatusCanceled, SubscriptionListResponseStatusUnpaid:
      return true
  }
  return false
}

type SubscriptionListResponseUser struct {
AvatarURL string `json:"avatar_url,required,nullable"`
Email string `json:"email,required"`
GitHubUsername string `json:"github_username,required,nullable"`
PublicName string `json:"public_name,required"`
JSON subscriptionListResponseUserJSON `json:"-"`
}

// subscriptionListResponseUserJSON contains the JSON metadata for the struct
// [SubscriptionListResponseUser]
type subscriptionListResponseUserJSON struct {
AvatarURL apijson.Field
Email apijson.Field
GitHubUsername apijson.Field
PublicName apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *SubscriptionListResponseUser) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionListResponseUserJSON) RawJSON() (string) {
  return r.raw
}

type SubscriptionExportResponse = interface{}

// Result of a subscription import operation.
type SubscriptionImportResponse struct {
Count int64 `json:"count,required"`
JSON subscriptionImportResponseJSON `json:"-"`
}

// subscriptionImportResponseJSON contains the JSON metadata for the struct
// [SubscriptionImportResponse]
type subscriptionImportResponseJSON struct {
Count apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *SubscriptionImportResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionImportResponseJSON) RawJSON() (string) {
  return r.raw
}

type SubscriptionNewParams struct {
// The email address of the user.
Email param.Field[string] `json:"email,required" format:"email"`
// The ID of the product. **Must be the free subscription tier**.
ProductID param.Field[string] `json:"product_id,required" format:"uuid4"`
}

func (r SubscriptionNewParams) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type SubscriptionListParams struct {
// Filter by active or inactive subscription.
Active param.Field[bool] `query:"active"`
// Size of a page, defaults to 10. Maximum is 100.
Limit param.Field[int64] `query:"limit"`
// Filter by organization ID.
OrganizationID param.Field[SubscriptionListParamsOrganizationIDUnion] `query:"organization_id" format:"uuid4"`
// Page number, defaults to 1.
Page param.Field[int64] `query:"page"`
// Filter by product ID.
ProductID param.Field[SubscriptionListParamsProductIDUnion] `query:"product_id" format:"uuid4"`
// Sorting criterion. Several criteria can be used simultaneously and will be
// applied in order. Add a minus sign `-` before the criteria name to sort by
// descending order.
Sorting param.Field[[]SubscriptionListParamsSorting] `query:"sorting"`
// Filter by subscription tier type.
Type param.Field[SubscriptionListParamsTypeUnion] `query:"type"`
}

// URLQuery serializes [SubscriptionListParams]'s query parameters as `url.Values`.
func (r SubscriptionListParams) URLQuery() (v url.Values) {
  return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
    ArrayFormat: apiquery.ArrayQueryFormatComma,
    NestedFormat: apiquery.NestedQueryFormatBrackets,
  })
}

// Filter by organization ID.
//
// Satisfied by [shared.UnionString], [SubscriptionListParamsOrganizationIDArray].
type SubscriptionListParamsOrganizationIDUnion interface {
  ImplementsSubscriptionListParamsOrganizationIDUnion()
}

type SubscriptionListParamsOrganizationIDArray []string

func (r SubscriptionListParamsOrganizationIDArray) ImplementsSubscriptionListParamsOrganizationIDUnion() {}

// Filter by product ID.
//
// Satisfied by [shared.UnionString], [SubscriptionListParamsProductIDArray].
type SubscriptionListParamsProductIDUnion interface {
  ImplementsSubscriptionListParamsProductIDUnion()
}

type SubscriptionListParamsProductIDArray []string

func (r SubscriptionListParamsProductIDArray) ImplementsSubscriptionListParamsProductIDUnion() {}

type SubscriptionListParamsSorting string

const (
  SubscriptionListParamsSortingUser SubscriptionListParamsSorting = "user"
  SubscriptionListParamsSorting-User SubscriptionListParamsSorting = "-user"
  SubscriptionListParamsSortingStatus SubscriptionListParamsSorting = "status"
  SubscriptionListParamsSorting-Status SubscriptionListParamsSorting = "-status"
  SubscriptionListParamsSortingStartedAt SubscriptionListParamsSorting = "started_at"
  SubscriptionListParamsSorting-StartedAt SubscriptionListParamsSorting = "-started_at"
  SubscriptionListParamsSortingCurrentPeriodEnd SubscriptionListParamsSorting = "current_period_end"
  SubscriptionListParamsSorting-CurrentPeriodEnd SubscriptionListParamsSorting = "-current_period_end"
  SubscriptionListParamsSortingPriceAmount SubscriptionListParamsSorting = "price_amount"
  SubscriptionListParamsSorting-PriceAmount SubscriptionListParamsSorting = "-price_amount"
  SubscriptionListParamsSortingSubscriptionTierType SubscriptionListParamsSorting = "subscription_tier_type"
  SubscriptionListParamsSorting-SubscriptionTierType SubscriptionListParamsSorting = "-subscription_tier_type"
  SubscriptionListParamsSortingProduct SubscriptionListParamsSorting = "product"
  SubscriptionListParamsSorting-Product SubscriptionListParamsSorting = "-product"
)

func (r SubscriptionListParamsSorting) IsKnown() (bool) {
  switch r {
  case SubscriptionListParamsSortingUser, SubscriptionListParamsSorting-User, SubscriptionListParamsSortingStatus, SubscriptionListParamsSorting-Status, SubscriptionListParamsSortingStartedAt, SubscriptionListParamsSorting-StartedAt, SubscriptionListParamsSortingCurrentPeriodEnd, SubscriptionListParamsSorting-CurrentPeriodEnd, SubscriptionListParamsSortingPriceAmount, SubscriptionListParamsSorting-PriceAmount, SubscriptionListParamsSortingSubscriptionTierType, SubscriptionListParamsSorting-SubscriptionTierType, SubscriptionListParamsSortingProduct, SubscriptionListParamsSorting-Product:
      return true
  }
  return false
}

// Filter by subscription tier type.
//
// Satisfied by [SubscriptionListParamsTypeSubscriptionTierType],
// [SubscriptionListParamsTypeArray].
type SubscriptionListParamsTypeUnion interface {
  implementsSubscriptionListParamsTypeUnion()
}

type SubscriptionListParamsTypeSubscriptionTierType string

const (
  SubscriptionListParamsTypeSubscriptionTierTypeFree SubscriptionListParamsTypeSubscriptionTierType = "free"
  SubscriptionListParamsTypeSubscriptionTierTypeIndividual SubscriptionListParamsTypeSubscriptionTierType = "individual"
  SubscriptionListParamsTypeSubscriptionTierTypeBusiness SubscriptionListParamsTypeSubscriptionTierType = "business"
)

func (r SubscriptionListParamsTypeSubscriptionTierType) IsKnown() (bool) {
  switch r {
  case SubscriptionListParamsTypeSubscriptionTierTypeFree, SubscriptionListParamsTypeSubscriptionTierTypeIndividual, SubscriptionListParamsTypeSubscriptionTierTypeBusiness:
      return true
  }
  return false
}

func (r SubscriptionListParamsTypeSubscriptionTierType) implementsSubscriptionListParamsTypeUnion() {}

type SubscriptionListParamsTypeArray []SubscriptionListParamsTypeArray

func (r SubscriptionListParamsTypeArray) implementsSubscriptionListParamsTypeUnion() {}

type SubscriptionExportParams struct {
// Filter by organization ID.
OrganizationID param.Field[SubscriptionExportParamsOrganizationIDUnion] `query:"organization_id" format:"uuid4"`
}

// URLQuery serializes [SubscriptionExportParams]'s query parameters as
// `url.Values`.
func (r SubscriptionExportParams) URLQuery() (v url.Values) {
  return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
    ArrayFormat: apiquery.ArrayQueryFormatComma,
    NestedFormat: apiquery.NestedQueryFormatBrackets,
  })
}

// Filter by organization ID.
//
// Satisfied by [shared.UnionString],
// [SubscriptionExportParamsOrganizationIDArray].
type SubscriptionExportParamsOrganizationIDUnion interface {
  ImplementsSubscriptionExportParamsOrganizationIDUnion()
}

type SubscriptionExportParamsOrganizationIDArray []string

func (r SubscriptionExportParamsOrganizationIDArray) ImplementsSubscriptionExportParamsOrganizationIDUnion() {}

type SubscriptionImportParams struct {
// CSV file with emails.
File param.Field[io.Reader] `json:"file,required" format:"binary"`
// The organization ID on which to import the subscriptions.
OrganizationID param.Field[string] `json:"organization_id,required" format:"uuid4"`
}

func (r SubscriptionImportParams) MarshalMultipart() (data []byte, contentType string, err error) {
  buf := bytes.NewBuffer(nil)
  writer := multipart.NewWriter(buf)
  err = apiform.MarshalRoot(r, writer)
  if err != nil {
    writer.Close()
    return nil, "", err
  }
  err = writer.Close()
  if err != nil {
    return nil, "", err
  }
  return buf.Bytes(), writer.FormDataContentType(), nil
}
