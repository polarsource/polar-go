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

// UserAdvertisementService contains methods and other services that help with
// interacting with the polar API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserAdvertisementService] method instead.
type UserAdvertisementService struct {
Options []option.RequestOption
}

// NewUserAdvertisementService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewUserAdvertisementService(opts ...option.RequestOption) (r *UserAdvertisementService) {
  r = &UserAdvertisementService{}
  r.Options = opts
  return
}

// Create an advertisement campaign.
func (r *UserAdvertisementService) New(ctx context.Context, body UserAdvertisementNewParams, opts ...option.RequestOption) (res *UserAdvertisementNewResponse, err error) {
  opts = append(r.Options[:], opts...)
  path := "v1/users/advertisements/"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
  return
}

// Get an advertisement campaign by ID.
func (r *UserAdvertisementService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *UserAdvertisementGetResponse, err error) {
  opts = append(r.Options[:], opts...)
  if id == "" {
    err = errors.New("missing required id parameter")
    return
  }
  path := fmt.Sprintf("v1/users/advertisements/%s", id)
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
  return
}

// Update an advertisement campaign.
func (r *UserAdvertisementService) Update(ctx context.Context, id string, body UserAdvertisementUpdateParams, opts ...option.RequestOption) (res *UserAdvertisementUpdateResponse, err error) {
  opts = append(r.Options[:], opts...)
  if id == "" {
    err = errors.New("missing required id parameter")
    return
  }
  path := fmt.Sprintf("v1/users/advertisements/%s", id)
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
  return
}

// List advertisement campaigns.
func (r *UserAdvertisementService) List(ctx context.Context, query UserAdvertisementListParams, opts ...option.RequestOption) (res *pagination.PolarPagination[UserAdvertisementListResponse], err error) {
  var raw *http.Response
  opts = append(r.Options[:], opts...)
  opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
  path := "v1/users/advertisements/"
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

// List advertisement campaigns.
func (r *UserAdvertisementService) ListAutoPaging(ctx context.Context, query UserAdvertisementListParams, opts ...option.RequestOption) (*pagination.PolarPaginationAutoPager[UserAdvertisementListResponse]) {
  return pagination.NewPolarPaginationAutoPager(r.List(ctx, query, opts...))
}

// Delete an advertisement campaign.
//
// It'll be automatically disabled on all granted benefits.
func (r *UserAdvertisementService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *UserAdvertisementDeleteResponse, err error) {
  opts = append(r.Options[:], opts...)
  if id == "" {
    err = errors.New("missing required id parameter")
    return
  }
  path := fmt.Sprintf("v1/users/advertisements/%s", id)
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
  return
}

// Enable an advertisement campaign on a granted benefit.
func (r *UserAdvertisementService) Enable(ctx context.Context, id string, body UserAdvertisementEnableParams, opts ...option.RequestOption) (err error) {
  opts = append(r.Options[:], opts...)
  opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
  if id == "" {
    err = errors.New("missing required id parameter")
    return
  }
  path := fmt.Sprintf("v1/users/advertisements/%s/enable", id)
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
  return
}

type UserAdvertisementNewResponse struct {
ID string `json:"id,required" format:"uuid4"`
Clicks int64 `json:"clicks,required"`
// Creation timestamp of the object.
CreatedAt time.Time `json:"created_at,required" format:"date-time"`
ImageURL string `json:"image_url,required" format:"uri"`
ImageURLDark string `json:"image_url_dark,required,nullable" format:"uri"`
LinkURL string `json:"link_url,required" format:"uri"`
// Last modification timestamp of the object.
ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
Text string `json:"text,required"`
UserID string `json:"user_id,required" format:"uuid4"`
Views int64 `json:"views,required"`
JSON userAdvertisementNewResponseJSON `json:"-"`
}

// userAdvertisementNewResponseJSON contains the JSON metadata for the struct
// [UserAdvertisementNewResponse]
type userAdvertisementNewResponseJSON struct {
ID apijson.Field
Clicks apijson.Field
CreatedAt apijson.Field
ImageURL apijson.Field
ImageURLDark apijson.Field
LinkURL apijson.Field
ModifiedAt apijson.Field
Text apijson.Field
UserID apijson.Field
Views apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *UserAdvertisementNewResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userAdvertisementNewResponseJSON) RawJSON() (string) {
  return r.raw
}

type UserAdvertisementGetResponse struct {
ID string `json:"id,required" format:"uuid4"`
Clicks int64 `json:"clicks,required"`
// Creation timestamp of the object.
CreatedAt time.Time `json:"created_at,required" format:"date-time"`
ImageURL string `json:"image_url,required" format:"uri"`
ImageURLDark string `json:"image_url_dark,required,nullable" format:"uri"`
LinkURL string `json:"link_url,required" format:"uri"`
// Last modification timestamp of the object.
ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
Text string `json:"text,required"`
UserID string `json:"user_id,required" format:"uuid4"`
Views int64 `json:"views,required"`
JSON userAdvertisementGetResponseJSON `json:"-"`
}

// userAdvertisementGetResponseJSON contains the JSON metadata for the struct
// [UserAdvertisementGetResponse]
type userAdvertisementGetResponseJSON struct {
ID apijson.Field
Clicks apijson.Field
CreatedAt apijson.Field
ImageURL apijson.Field
ImageURLDark apijson.Field
LinkURL apijson.Field
ModifiedAt apijson.Field
Text apijson.Field
UserID apijson.Field
Views apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *UserAdvertisementGetResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userAdvertisementGetResponseJSON) RawJSON() (string) {
  return r.raw
}

type UserAdvertisementUpdateResponse struct {
ID string `json:"id,required" format:"uuid4"`
Clicks int64 `json:"clicks,required"`
// Creation timestamp of the object.
CreatedAt time.Time `json:"created_at,required" format:"date-time"`
ImageURL string `json:"image_url,required" format:"uri"`
ImageURLDark string `json:"image_url_dark,required,nullable" format:"uri"`
LinkURL string `json:"link_url,required" format:"uri"`
// Last modification timestamp of the object.
ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
Text string `json:"text,required"`
UserID string `json:"user_id,required" format:"uuid4"`
Views int64 `json:"views,required"`
JSON userAdvertisementUpdateResponseJSON `json:"-"`
}

// userAdvertisementUpdateResponseJSON contains the JSON metadata for the struct
// [UserAdvertisementUpdateResponse]
type userAdvertisementUpdateResponseJSON struct {
ID apijson.Field
Clicks apijson.Field
CreatedAt apijson.Field
ImageURL apijson.Field
ImageURLDark apijson.Field
LinkURL apijson.Field
ModifiedAt apijson.Field
Text apijson.Field
UserID apijson.Field
Views apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *UserAdvertisementUpdateResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userAdvertisementUpdateResponseJSON) RawJSON() (string) {
  return r.raw
}

type UserAdvertisementListResponse struct {
ID string `json:"id,required" format:"uuid4"`
Clicks int64 `json:"clicks,required"`
// Creation timestamp of the object.
CreatedAt time.Time `json:"created_at,required" format:"date-time"`
ImageURL string `json:"image_url,required" format:"uri"`
ImageURLDark string `json:"image_url_dark,required,nullable" format:"uri"`
LinkURL string `json:"link_url,required" format:"uri"`
// Last modification timestamp of the object.
ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
Text string `json:"text,required"`
UserID string `json:"user_id,required" format:"uuid4"`
Views int64 `json:"views,required"`
JSON userAdvertisementListResponseJSON `json:"-"`
}

// userAdvertisementListResponseJSON contains the JSON metadata for the struct
// [UserAdvertisementListResponse]
type userAdvertisementListResponseJSON struct {
ID apijson.Field
Clicks apijson.Field
CreatedAt apijson.Field
ImageURL apijson.Field
ImageURLDark apijson.Field
LinkURL apijson.Field
ModifiedAt apijson.Field
Text apijson.Field
UserID apijson.Field
Views apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *UserAdvertisementListResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r userAdvertisementListResponseJSON) RawJSON() (string) {
  return r.raw
}

type UserAdvertisementDeleteResponse = interface{}

type UserAdvertisementNewParams struct {
ImageURL param.Field[string] `json:"image_url,required" format:"uri"`
LinkURL param.Field[string] `json:"link_url,required" format:"uri"`
Text param.Field[string] `json:"text,required"`
ImageURLDark param.Field[string] `json:"image_url_dark" format:"uri"`
}

func (r UserAdvertisementNewParams) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type UserAdvertisementUpdateParams struct {
ImageURL param.Field[string] `json:"image_url" format:"uri"`
ImageURLDark param.Field[string] `json:"image_url_dark" format:"uri"`
LinkURL param.Field[string] `json:"link_url" format:"uri"`
Text param.Field[string] `json:"text"`
}

func (r UserAdvertisementUpdateParams) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type UserAdvertisementListParams struct {
// Size of a page, defaults to 10. Maximum is 100.
Limit param.Field[int64] `query:"limit"`
// Page number, defaults to 1.
Page param.Field[int64] `query:"page"`
// Sorting criterion. Several criteria can be used simultaneously and will be
// applied in order. Add a minus sign `-` before the criteria name to sort by
// descending order.
Sorting param.Field[[]UserAdvertisementListParamsSorting] `query:"sorting"`
}

// URLQuery serializes [UserAdvertisementListParams]'s query parameters as
// `url.Values`.
func (r UserAdvertisementListParams) URLQuery() (v url.Values) {
  return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
    ArrayFormat: apiquery.ArrayQueryFormatComma,
    NestedFormat: apiquery.NestedQueryFormatBrackets,
  })
}

type UserAdvertisementListParamsSorting string

const (
  UserAdvertisementListParamsSortingCreatedAt UserAdvertisementListParamsSorting = "created_at"
  UserAdvertisementListParamsSorting-CreatedAt UserAdvertisementListParamsSorting = "-created_at"
  UserAdvertisementListParamsSortingViews UserAdvertisementListParamsSorting = "views"
  UserAdvertisementListParamsSorting-Views UserAdvertisementListParamsSorting = "-views"
  UserAdvertisementListParamsSortingClicks UserAdvertisementListParamsSorting = "clicks"
  UserAdvertisementListParamsSorting-Clicks UserAdvertisementListParamsSorting = "-clicks"
)

func (r UserAdvertisementListParamsSorting) IsKnown() (bool) {
  switch r {
  case UserAdvertisementListParamsSortingCreatedAt, UserAdvertisementListParamsSorting-CreatedAt, UserAdvertisementListParamsSortingViews, UserAdvertisementListParamsSorting-Views, UserAdvertisementListParamsSortingClicks, UserAdvertisementListParamsSorting-Clicks:
      return true
  }
  return false
}

type UserAdvertisementEnableParams struct {
// The benefit ID.
BenefitID param.Field[string] `json:"benefit_id,required" format:"uuid4"`
}

func (r UserAdvertisementEnableParams) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}
