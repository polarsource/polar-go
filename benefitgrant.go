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

// BenefitGrantService contains methods and other services that help with
// interacting with the polar API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBenefitGrantService] method instead.
type BenefitGrantService struct {
	Options []option.RequestOption
}

// NewBenefitGrantService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBenefitGrantService(opts ...option.RequestOption) (r *BenefitGrantService) {
	r = &BenefitGrantService{}
	r.Options = opts
	return
}

// List the individual grants for a benefit.
//
// It's especially useful to check if a user has been granted a benefit.
func (r *BenefitGrantService) List(ctx context.Context, id string, query BenefitGrantListParams, opts ...option.RequestOption) (res *pagination.PolarPagination[BenefitGrantListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/benefits/%s/grants", id)
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

// List the individual grants for a benefit.
//
// It's especially useful to check if a user has been granted a benefit.
func (r *BenefitGrantService) ListAutoPaging(ctx context.Context, id string, query BenefitGrantListParams, opts ...option.RequestOption) *pagination.PolarPaginationAutoPager[BenefitGrantListResponse] {
	return pagination.NewPolarPaginationAutoPager(r.List(ctx, id, query, opts...))
}

type ListResourceBenefitGrant struct {
	Pagination ListResourceBenefitGrantPagination `json:"pagination,required"`
	Items      []ListResourceBenefitGrantItem     `json:"items"`
	JSON       listResourceBenefitGrantJSON       `json:"-"`
}

// listResourceBenefitGrantJSON contains the JSON metadata for the struct
// [ListResourceBenefitGrant]
type listResourceBenefitGrantJSON struct {
	Pagination  apijson.Field
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListResourceBenefitGrant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResourceBenefitGrantJSON) RawJSON() string {
	return r.raw
}

type ListResourceBenefitGrantPagination struct {
	MaxPage    int64                                  `json:"max_page,required"`
	TotalCount int64                                  `json:"total_count,required"`
	JSON       listResourceBenefitGrantPaginationJSON `json:"-"`
}

// listResourceBenefitGrantPaginationJSON contains the JSON metadata for the struct
// [ListResourceBenefitGrantPagination]
type listResourceBenefitGrantPaginationJSON struct {
	MaxPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListResourceBenefitGrantPagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResourceBenefitGrantPaginationJSON) RawJSON() string {
	return r.raw
}

// A grant of a benefit to a user.
type ListResourceBenefitGrantItem struct {
	// The ID of the grant.
	ID string `json:"id,required" format:"uuid4"`
	// The ID of the benefit concerned by this grant.
	BenefitID string `json:"benefit_id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the benefit is granted.
	IsGranted bool `json:"is_granted,required"`
	// Whether the benefit is revoked.
	IsRevoked bool `json:"is_revoked,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The ID of the order that granted this benefit.
	OrderID string `json:"order_id,required,nullable" format:"uuid4"`
	// Properties for a benefit grant.
	Properties interface{} `json:"properties,required"`
	// The ID of the subscription that granted this benefit.
	SubscriptionID string `json:"subscription_id,required,nullable" format:"uuid4"`
	// The ID of the user concerned by this grant.
	UserID string `json:"user_id,required" format:"uuid4"`
	// The timestamp when the benefit was granted. If `None`, the benefit is not
	// granted.
	GrantedAt time.Time `json:"granted_at,nullable" format:"date-time"`
	// The timestamp when the benefit was revoked. If `None`, the benefit is not
	// revoked.
	RevokedAt time.Time                        `json:"revoked_at,nullable" format:"date-time"`
	JSON      listResourceBenefitGrantItemJSON `json:"-"`
}

// listResourceBenefitGrantItemJSON contains the JSON metadata for the struct
// [ListResourceBenefitGrantItem]
type listResourceBenefitGrantItemJSON struct {
	ID             apijson.Field
	BenefitID      apijson.Field
	CreatedAt      apijson.Field
	IsGranted      apijson.Field
	IsRevoked      apijson.Field
	ModifiedAt     apijson.Field
	OrderID        apijson.Field
	Properties     apijson.Field
	SubscriptionID apijson.Field
	UserID         apijson.Field
	GrantedAt      apijson.Field
	RevokedAt      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ListResourceBenefitGrantItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResourceBenefitGrantItemJSON) RawJSON() string {
	return r.raw
}

// A grant of a benefit to a user.
type BenefitGrantListResponse struct {
	// The ID of the grant.
	ID string `json:"id,required" format:"uuid4"`
	// The ID of the benefit concerned by this grant.
	BenefitID string `json:"benefit_id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the benefit is granted.
	IsGranted bool `json:"is_granted,required"`
	// Whether the benefit is revoked.
	IsRevoked bool `json:"is_revoked,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The ID of the order that granted this benefit.
	OrderID string `json:"order_id,required,nullable" format:"uuid4"`
	// Properties for a benefit grant.
	Properties interface{} `json:"properties,required"`
	// The ID of the subscription that granted this benefit.
	SubscriptionID string `json:"subscription_id,required,nullable" format:"uuid4"`
	// The ID of the user concerned by this grant.
	UserID string `json:"user_id,required" format:"uuid4"`
	// The timestamp when the benefit was granted. If `None`, the benefit is not
	// granted.
	GrantedAt time.Time `json:"granted_at,nullable" format:"date-time"`
	// The timestamp when the benefit was revoked. If `None`, the benefit is not
	// revoked.
	RevokedAt time.Time                    `json:"revoked_at,nullable" format:"date-time"`
	JSON      benefitGrantListResponseJSON `json:"-"`
}

// benefitGrantListResponseJSON contains the JSON metadata for the struct
// [BenefitGrantListResponse]
type benefitGrantListResponseJSON struct {
	ID             apijson.Field
	BenefitID      apijson.Field
	CreatedAt      apijson.Field
	IsGranted      apijson.Field
	IsRevoked      apijson.Field
	ModifiedAt     apijson.Field
	OrderID        apijson.Field
	Properties     apijson.Field
	SubscriptionID apijson.Field
	UserID         apijson.Field
	GrantedAt      apijson.Field
	RevokedAt      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BenefitGrantListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitGrantListResponseJSON) RawJSON() string {
	return r.raw
}

type BenefitGrantListParams struct {
	// Filter by GitHub user ID. Only available for users who have linked their GitHub
	// account on Polar.
	GitHubUserID param.Field[int64] `query:"github_user_id"`
	// Filter by granted status. If `true`, only granted benefits will be returned. If
	// `false`, only revoked benefits will be returned.
	IsGranted param.Field[bool] `query:"is_granted"`
	// Size of a page, defaults to 10. Maximum is 100.
	Limit param.Field[int64] `query:"limit"`
	// Page number, defaults to 1.
	Page param.Field[int64] `query:"page"`
	// Filter by user ID.
	UserID param.Field[string] `query:"user_id" format:"uuid4"`
}

// URLQuery serializes [BenefitGrantListParams]'s query parameters as `url.Values`.
func (r BenefitGrantListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
