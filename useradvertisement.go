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
func (r *UserAdvertisementService) List(ctx context.Context, query UserAdvertisementListParams, opts ...option.RequestOption) (res *UserAdvertisementListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/users/advertisements/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
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
	ID     string `json:"id,required" format:"uuid4"`
	Clicks int64  `json:"clicks,required"`
	// Creation timestamp of the object.
	CreatedAt    time.Time `json:"created_at,required" format:"date-time"`
	ImageURL     string    `json:"image_url,required" format:"uri"`
	LinkURL      string    `json:"link_url,required" format:"uri"`
	Text         string    `json:"text,required"`
	UserID       string    `json:"user_id,required" format:"uuid4"`
	Views        int64     `json:"views,required"`
	ImageURLDark string    `json:"image_url_dark,nullable" format:"uri"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time                        `json:"modified_at,nullable" format:"date-time"`
	JSON       userAdvertisementNewResponseJSON `json:"-"`
}

// userAdvertisementNewResponseJSON contains the JSON metadata for the struct
// [UserAdvertisementNewResponse]
type userAdvertisementNewResponseJSON struct {
	ID           apijson.Field
	Clicks       apijson.Field
	CreatedAt    apijson.Field
	ImageURL     apijson.Field
	LinkURL      apijson.Field
	Text         apijson.Field
	UserID       apijson.Field
	Views        apijson.Field
	ImageURLDark apijson.Field
	ModifiedAt   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *UserAdvertisementNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userAdvertisementNewResponseJSON) RawJSON() string {
	return r.raw
}

type UserAdvertisementGetResponse struct {
	ID     string `json:"id,required" format:"uuid4"`
	Clicks int64  `json:"clicks,required"`
	// Creation timestamp of the object.
	CreatedAt    time.Time `json:"created_at,required" format:"date-time"`
	ImageURL     string    `json:"image_url,required" format:"uri"`
	LinkURL      string    `json:"link_url,required" format:"uri"`
	Text         string    `json:"text,required"`
	UserID       string    `json:"user_id,required" format:"uuid4"`
	Views        int64     `json:"views,required"`
	ImageURLDark string    `json:"image_url_dark,nullable" format:"uri"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time                        `json:"modified_at,nullable" format:"date-time"`
	JSON       userAdvertisementGetResponseJSON `json:"-"`
}

// userAdvertisementGetResponseJSON contains the JSON metadata for the struct
// [UserAdvertisementGetResponse]
type userAdvertisementGetResponseJSON struct {
	ID           apijson.Field
	Clicks       apijson.Field
	CreatedAt    apijson.Field
	ImageURL     apijson.Field
	LinkURL      apijson.Field
	Text         apijson.Field
	UserID       apijson.Field
	Views        apijson.Field
	ImageURLDark apijson.Field
	ModifiedAt   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *UserAdvertisementGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userAdvertisementGetResponseJSON) RawJSON() string {
	return r.raw
}

type UserAdvertisementUpdateResponse struct {
	ID     string `json:"id,required" format:"uuid4"`
	Clicks int64  `json:"clicks,required"`
	// Creation timestamp of the object.
	CreatedAt    time.Time `json:"created_at,required" format:"date-time"`
	ImageURL     string    `json:"image_url,required" format:"uri"`
	LinkURL      string    `json:"link_url,required" format:"uri"`
	Text         string    `json:"text,required"`
	UserID       string    `json:"user_id,required" format:"uuid4"`
	Views        int64     `json:"views,required"`
	ImageURLDark string    `json:"image_url_dark,nullable" format:"uri"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time                           `json:"modified_at,nullable" format:"date-time"`
	JSON       userAdvertisementUpdateResponseJSON `json:"-"`
}

// userAdvertisementUpdateResponseJSON contains the JSON metadata for the struct
// [UserAdvertisementUpdateResponse]
type userAdvertisementUpdateResponseJSON struct {
	ID           apijson.Field
	Clicks       apijson.Field
	CreatedAt    apijson.Field
	ImageURL     apijson.Field
	LinkURL      apijson.Field
	Text         apijson.Field
	UserID       apijson.Field
	Views        apijson.Field
	ImageURLDark apijson.Field
	ModifiedAt   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *UserAdvertisementUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userAdvertisementUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type UserAdvertisementListResponse struct {
	Pagination UserAdvertisementListResponsePagination `json:"pagination,required"`
	Items      []UserAdvertisementListResponseItem     `json:"items"`
	JSON       userAdvertisementListResponseJSON       `json:"-"`
}

// userAdvertisementListResponseJSON contains the JSON metadata for the struct
// [UserAdvertisementListResponse]
type userAdvertisementListResponseJSON struct {
	Pagination  apijson.Field
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserAdvertisementListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userAdvertisementListResponseJSON) RawJSON() string {
	return r.raw
}

type UserAdvertisementListResponsePagination struct {
	MaxPage    int64                                       `json:"max_page,required"`
	TotalCount int64                                       `json:"total_count,required"`
	JSON       userAdvertisementListResponsePaginationJSON `json:"-"`
}

// userAdvertisementListResponsePaginationJSON contains the JSON metadata for the
// struct [UserAdvertisementListResponsePagination]
type userAdvertisementListResponsePaginationJSON struct {
	MaxPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserAdvertisementListResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userAdvertisementListResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type UserAdvertisementListResponseItem struct {
	ID     string `json:"id,required" format:"uuid4"`
	Clicks int64  `json:"clicks,required"`
	// Creation timestamp of the object.
	CreatedAt    time.Time `json:"created_at,required" format:"date-time"`
	ImageURL     string    `json:"image_url,required" format:"uri"`
	LinkURL      string    `json:"link_url,required" format:"uri"`
	Text         string    `json:"text,required"`
	UserID       string    `json:"user_id,required" format:"uuid4"`
	Views        int64     `json:"views,required"`
	ImageURLDark string    `json:"image_url_dark,nullable" format:"uri"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time                             `json:"modified_at,nullable" format:"date-time"`
	JSON       userAdvertisementListResponseItemJSON `json:"-"`
}

// userAdvertisementListResponseItemJSON contains the JSON metadata for the struct
// [UserAdvertisementListResponseItem]
type userAdvertisementListResponseItemJSON struct {
	ID           apijson.Field
	Clicks       apijson.Field
	CreatedAt    apijson.Field
	ImageURL     apijson.Field
	LinkURL      apijson.Field
	Text         apijson.Field
	UserID       apijson.Field
	Views        apijson.Field
	ImageURLDark apijson.Field
	ModifiedAt   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *UserAdvertisementListResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userAdvertisementListResponseItemJSON) RawJSON() string {
	return r.raw
}

type UserAdvertisementDeleteResponse = interface{}

type UserAdvertisementNewParams struct {
	ImageURL     param.Field[string] `json:"image_url,required" format:"uri"`
	LinkURL      param.Field[string] `json:"link_url,required" format:"uri"`
	Text         param.Field[string] `json:"text,required"`
	ImageURLDark param.Field[string] `json:"image_url_dark" format:"uri"`
}

func (r UserAdvertisementNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserAdvertisementUpdateParams struct {
	ImageURL     param.Field[string] `json:"image_url" format:"uri"`
	ImageURLDark param.Field[string] `json:"image_url_dark" format:"uri"`
	LinkURL      param.Field[string] `json:"link_url" format:"uri"`
	Text         param.Field[string] `json:"text"`
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
	Sorting param.Field[[]string] `query:"sorting"`
}

// URLQuery serializes [UserAdvertisementListParams]'s query parameters as
// `url.Values`.
func (r UserAdvertisementListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UserAdvertisementEnableParams struct {
	// The benefit ID.
	BenefitID param.Field[string] `json:"benefit_id,required" format:"uuid4"`
}

func (r UserAdvertisementEnableParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
