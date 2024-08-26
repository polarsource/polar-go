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

// AdvertisementService contains methods and other services that help with
// interacting with the polar API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdvertisementService] method instead.
type AdvertisementService struct {
	Options []option.RequestOption
}

// NewAdvertisementService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAdvertisementService(opts ...option.RequestOption) (r *AdvertisementService) {
	r = &AdvertisementService{}
	r.Options = opts
	return
}

// Get an advertisement campaign by ID.
func (r *AdvertisementService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *AdvertisementCampaign, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/advertisements/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List active advertisement campaigns for a benefit.
func (r *AdvertisementService) List(ctx context.Context, query AdvertisementListParams, opts ...option.RequestOption) (res *pagination.PolarPagination[AdvertisementCampaign], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/advertisements/"
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

// List active advertisement campaigns for a benefit.
func (r *AdvertisementService) ListAutoPaging(ctx context.Context, query AdvertisementListParams, opts ...option.RequestOption) *pagination.PolarPaginationAutoPager[AdvertisementCampaign] {
	return pagination.NewPolarPaginationAutoPager(r.List(ctx, query, opts...))
}

type AdvertisementCampaign struct {
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt    time.Time `json:"created_at,required" format:"date-time"`
	ImageURL     string    `json:"image_url,required" format:"uri"`
	ImageURLDark string    `json:"image_url_dark,required,nullable" format:"uri"`
	LinkURL      string    `json:"link_url,required" format:"uri"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time                 `json:"modified_at,required,nullable" format:"date-time"`
	Text       string                    `json:"text,required"`
	JSON       advertisementCampaignJSON `json:"-"`
}

// advertisementCampaignJSON contains the JSON metadata for the struct
// [AdvertisementCampaign]
type advertisementCampaignJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	ImageURL     apijson.Field
	ImageURLDark apijson.Field
	LinkURL      apijson.Field
	ModifiedAt   apijson.Field
	Text         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AdvertisementCampaign) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advertisementCampaignJSON) RawJSON() string {
	return r.raw
}

type AdvertisementCampaignListResource struct {
	// The dimensions (width, height) in pixels of the advertisement images.
	Dimensions []int64                                     `json:"dimensions,required"`
	Pagination AdvertisementCampaignListResourcePagination `json:"pagination,required"`
	Items      []AdvertisementCampaign                     `json:"items"`
	JSON       advertisementCampaignListResourceJSON       `json:"-"`
}

// advertisementCampaignListResourceJSON contains the JSON metadata for the struct
// [AdvertisementCampaignListResource]
type advertisementCampaignListResourceJSON struct {
	Dimensions  apijson.Field
	Pagination  apijson.Field
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvertisementCampaignListResource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advertisementCampaignListResourceJSON) RawJSON() string {
	return r.raw
}

type AdvertisementCampaignListResourcePagination struct {
	MaxPage    int64                                           `json:"max_page,required"`
	TotalCount int64                                           `json:"total_count,required"`
	JSON       advertisementCampaignListResourcePaginationJSON `json:"-"`
}

// advertisementCampaignListResourcePaginationJSON contains the JSON metadata for
// the struct [AdvertisementCampaignListResourcePagination]
type advertisementCampaignListResourcePaginationJSON struct {
	MaxPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvertisementCampaignListResourcePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advertisementCampaignListResourcePaginationJSON) RawJSON() string {
	return r.raw
}

type AdvertisementListParams struct {
	// The benefit ID.
	BenefitID param.Field[string] `query:"benefit_id,required" format:"uuid4"`
	// Size of a page, defaults to 10. Maximum is 100.
	Limit param.Field[int64] `query:"limit"`
	// Page number, defaults to 1.
	Page param.Field[int64] `query:"page"`
	// Sorting criterion. Several criteria can be used simultaneously and will be
	// applied in order. Add a minus sign `-` before the criteria name to sort by
	// descending order.
	Sorting param.Field[[]string] `query:"sorting"`
}

// URLQuery serializes [AdvertisementListParams]'s query parameters as
// `url.Values`.
func (r AdvertisementListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
