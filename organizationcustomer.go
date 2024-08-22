// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package polar

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/polarsource/polar-go/internal/apijson"
	"github.com/polarsource/polar-go/internal/apiquery"
	"github.com/polarsource/polar-go/internal/param"
	"github.com/polarsource/polar-go/internal/requestconfig"
	"github.com/polarsource/polar-go/option"
)

// OrganizationCustomerService contains methods and other services that help with
// interacting with the polar API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationCustomerService] method instead.
type OrganizationCustomerService struct {
	Options []option.RequestOption
}

// NewOrganizationCustomerService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOrganizationCustomerService(opts ...option.RequestOption) (r *OrganizationCustomerService) {
	r = &OrganizationCustomerService{}
	r.Options = opts
	return
}

// List organization customers.
func (r *OrganizationCustomerService) List(ctx context.Context, id string, query OrganizationCustomerListParams, opts ...option.RequestOption) (res *OrganizationCustomerListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/organizations/%s/customers", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type OrganizationCustomerListResponse struct {
	Pagination OrganizationCustomerListResponsePagination `json:"pagination,required"`
	Items      []OrganizationCustomerListResponseItem     `json:"items"`
	JSON       organizationCustomerListResponseJSON       `json:"-"`
}

// organizationCustomerListResponseJSON contains the JSON metadata for the struct
// [OrganizationCustomerListResponse]
type organizationCustomerListResponseJSON struct {
	Pagination  apijson.Field
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationCustomerListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationCustomerListResponseJSON) RawJSON() string {
	return r.raw
}

type OrganizationCustomerListResponsePagination struct {
	MaxPage    int64                                          `json:"max_page,required"`
	TotalCount int64                                          `json:"total_count,required"`
	JSON       organizationCustomerListResponsePaginationJSON `json:"-"`
}

// organizationCustomerListResponsePaginationJSON contains the JSON metadata for
// the struct [OrganizationCustomerListResponsePagination]
type organizationCustomerListResponsePaginationJSON struct {
	MaxPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationCustomerListResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationCustomerListResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type OrganizationCustomerListResponseItem struct {
	PublicName     string                                   `json:"public_name,required"`
	AvatarURL      string                                   `json:"avatar_url,nullable"`
	GitHubUsername string                                   `json:"github_username,nullable"`
	JSON           organizationCustomerListResponseItemJSON `json:"-"`
}

// organizationCustomerListResponseItemJSON contains the JSON metadata for the
// struct [OrganizationCustomerListResponseItem]
type organizationCustomerListResponseItemJSON struct {
	PublicName     apijson.Field
	AvatarURL      apijson.Field
	GitHubUsername apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *OrganizationCustomerListResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationCustomerListResponseItemJSON) RawJSON() string {
	return r.raw
}

type OrganizationCustomerListParams struct {
	// Filter by the type of purchase the customer made.
	CustomerTypes param.Field[[]OrganizationCustomerListParamsCustomerType] `query:"customer_types"`
	// Size of a page, defaults to 10. Maximum is 100.
	Limit param.Field[int64] `query:"limit"`
	// Page number, defaults to 1.
	Page param.Field[int64] `query:"page"`
}

// URLQuery serializes [OrganizationCustomerListParams]'s query parameters as
// `url.Values`.
func (r OrganizationCustomerListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrganizationCustomerListParamsCustomerType string

const (
	OrganizationCustomerListParamsCustomerTypeFreeSubscription OrganizationCustomerListParamsCustomerType = "free_subscription"
	OrganizationCustomerListParamsCustomerTypePaidSubscription OrganizationCustomerListParamsCustomerType = "paid_subscription"
	OrganizationCustomerListParamsCustomerTypeOrder            OrganizationCustomerListParamsCustomerType = "order"
	OrganizationCustomerListParamsCustomerTypeDonation         OrganizationCustomerListParamsCustomerType = "donation"
)

func (r OrganizationCustomerListParamsCustomerType) IsKnown() bool {
	switch r {
	case OrganizationCustomerListParamsCustomerTypeFreeSubscription, OrganizationCustomerListParamsCustomerTypePaidSubscription, OrganizationCustomerListParamsCustomerTypeOrder, OrganizationCustomerListParamsCustomerTypeDonation:
		return true
	}
	return false
}
