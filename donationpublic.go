// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package polar

import (
	"context"
	"net/http"
	"net/url"
	"reflect"

	"github.com/stainless-sdks/polar-go/internal/apijson"
	"github.com/stainless-sdks/polar-go/internal/apiquery"
	"github.com/stainless-sdks/polar-go/internal/param"
	"github.com/stainless-sdks/polar-go/internal/requestconfig"
	"github.com/stainless-sdks/polar-go/option"
	"github.com/tidwall/gjson"
)

// DonationPublicService contains methods and other services that help with
// interacting with the polar API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDonationPublicService] method instead.
type DonationPublicService struct {
	Options []option.RequestOption
}

// NewDonationPublicService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDonationPublicService(opts ...option.RequestOption) (r *DonationPublicService) {
	r = &DonationPublicService{}
	r.Options = opts
	return
}

// Donations Public Search
func (r *DonationPublicService) Search(ctx context.Context, query DonationPublicSearchParams, opts ...option.RequestOption) (res *ListResourcePublicDonation, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/donations/public/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ListResourcePublicDonation struct {
	Pagination ListResourcePublicDonationPagination `json:"pagination,required"`
	Items      []ListResourcePublicDonationItem     `json:"items"`
	JSON       listResourcePublicDonationJSON       `json:"-"`
}

// listResourcePublicDonationJSON contains the JSON metadata for the struct
// [ListResourcePublicDonation]
type listResourcePublicDonationJSON struct {
	Pagination  apijson.Field
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListResourcePublicDonation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResourcePublicDonationJSON) RawJSON() string {
	return r.raw
}

type ListResourcePublicDonationPagination struct {
	MaxPage    int64                                    `json:"max_page,required"`
	TotalCount int64                                    `json:"total_count,required"`
	JSON       listResourcePublicDonationPaginationJSON `json:"-"`
}

// listResourcePublicDonationPaginationJSON contains the JSON metadata for the
// struct [ListResourcePublicDonationPagination]
type listResourcePublicDonationPaginationJSON struct {
	MaxPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListResourcePublicDonationPagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResourcePublicDonationPaginationJSON) RawJSON() string {
	return r.raw
}

type ListResourcePublicDonationItem struct {
	ID       string                               `json:"id,required" format:"uuid4"`
	Amount   int64                                `json:"amount,required"`
	Currency string                               `json:"currency,required"`
	Donor    ListResourcePublicDonationItemsDonor `json:"donor,required,nullable"`
	Message  string                               `json:"message,required,nullable"`
	JSON     listResourcePublicDonationItemJSON   `json:"-"`
}

// listResourcePublicDonationItemJSON contains the JSON metadata for the struct
// [ListResourcePublicDonationItem]
type listResourcePublicDonationItemJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	Currency    apijson.Field
	Donor       apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListResourcePublicDonationItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResourcePublicDonationItemJSON) RawJSON() string {
	return r.raw
}

type ListResourcePublicDonationItemsDonor struct {
	ID         string                                       `json:"id,required" format:"uuid4"`
	Platform   ListResourcePublicDonationItemsDonorPlatform `json:"platform"`
	Name       string                                       `json:"name"`
	AvatarURL  string                                       `json:"avatar_url,required"`
	IsPersonal bool                                         `json:"is_personal"`
	PublicName string                                       `json:"public_name"`
	JSON       listResourcePublicDonationItemsDonorJSON     `json:"-"`
	union      ListResourcePublicDonationItemsDonorUnion
}

// listResourcePublicDonationItemsDonorJSON contains the JSON metadata for the
// struct [ListResourcePublicDonationItemsDonor]
type listResourcePublicDonationItemsDonorJSON struct {
	ID          apijson.Field
	Platform    apijson.Field
	Name        apijson.Field
	AvatarURL   apijson.Field
	IsPersonal  apijson.Field
	PublicName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r listResourcePublicDonationItemsDonorJSON) RawJSON() string {
	return r.raw
}

func (r *ListResourcePublicDonationItemsDonor) UnmarshalJSON(data []byte) (err error) {
	*r = ListResourcePublicDonationItemsDonor{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ListResourcePublicDonationItemsDonorUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ListResourcePublicDonationItemsDonorDonationOrganization],
// [ListResourcePublicDonationItemsDonorDonationUser].
func (r ListResourcePublicDonationItemsDonor) AsUnion() ListResourcePublicDonationItemsDonorUnion {
	return r.union
}

// Union satisfied by [ListResourcePublicDonationItemsDonorDonationOrganization] or
// [ListResourcePublicDonationItemsDonorDonationUser].
type ListResourcePublicDonationItemsDonorUnion interface {
	implementsListResourcePublicDonationItemsDonor()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ListResourcePublicDonationItemsDonorUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ListResourcePublicDonationItemsDonorDonationOrganization{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ListResourcePublicDonationItemsDonorDonationUser{}),
		},
	)
}

type ListResourcePublicDonationItemsDonorDonationOrganization struct {
	ID         string                                                           `json:"id,required" format:"uuid4"`
	AvatarURL  string                                                           `json:"avatar_url,required"`
	IsPersonal bool                                                             `json:"is_personal,required"`
	Name       string                                                           `json:"name,required"`
	Platform   ListResourcePublicDonationItemsDonorDonationOrganizationPlatform `json:"platform,required"`
	JSON       listResourcePublicDonationItemsDonorDonationOrganizationJSON     `json:"-"`
}

// listResourcePublicDonationItemsDonorDonationOrganizationJSON contains the JSON
// metadata for the struct
// [ListResourcePublicDonationItemsDonorDonationOrganization]
type listResourcePublicDonationItemsDonorDonationOrganizationJSON struct {
	ID          apijson.Field
	AvatarURL   apijson.Field
	IsPersonal  apijson.Field
	Name        apijson.Field
	Platform    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListResourcePublicDonationItemsDonorDonationOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResourcePublicDonationItemsDonorDonationOrganizationJSON) RawJSON() string {
	return r.raw
}

func (r ListResourcePublicDonationItemsDonorDonationOrganization) implementsListResourcePublicDonationItemsDonor() {
}

type ListResourcePublicDonationItemsDonorDonationOrganizationPlatform string

const (
	ListResourcePublicDonationItemsDonorDonationOrganizationPlatformGitHub ListResourcePublicDonationItemsDonorDonationOrganizationPlatform = "github"
)

func (r ListResourcePublicDonationItemsDonorDonationOrganizationPlatform) IsKnown() bool {
	switch r {
	case ListResourcePublicDonationItemsDonorDonationOrganizationPlatformGitHub:
		return true
	}
	return false
}

type ListResourcePublicDonationItemsDonorDonationUser struct {
	ID         string                                               `json:"id,required" format:"uuid4"`
	AvatarURL  string                                               `json:"avatar_url,required,nullable"`
	PublicName string                                               `json:"public_name,required"`
	JSON       listResourcePublicDonationItemsDonorDonationUserJSON `json:"-"`
}

// listResourcePublicDonationItemsDonorDonationUserJSON contains the JSON metadata
// for the struct [ListResourcePublicDonationItemsDonorDonationUser]
type listResourcePublicDonationItemsDonorDonationUserJSON struct {
	ID          apijson.Field
	AvatarURL   apijson.Field
	PublicName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListResourcePublicDonationItemsDonorDonationUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResourcePublicDonationItemsDonorDonationUserJSON) RawJSON() string {
	return r.raw
}

func (r ListResourcePublicDonationItemsDonorDonationUser) implementsListResourcePublicDonationItemsDonor() {
}

type ListResourcePublicDonationItemsDonorPlatform string

const (
	ListResourcePublicDonationItemsDonorPlatformGitHub ListResourcePublicDonationItemsDonorPlatform = "github"
)

func (r ListResourcePublicDonationItemsDonorPlatform) IsKnown() bool {
	switch r {
	case ListResourcePublicDonationItemsDonorPlatformGitHub:
		return true
	}
	return false
}

type DonationPublicSearchParams struct {
	// The organization ID.
	OrganizationID param.Field[string] `query:"organization_id,required" format:"uuid4"`
	// Size of a page, defaults to 10. Maximum is 100.
	Limit param.Field[int64] `query:"limit"`
	// Page number, defaults to 1.
	Page param.Field[int64] `query:"page"`
	// Sorting criterion. Several criteria can be used simultaneously and will be
	// applied in order. Add a minus sign `-` before the criteria name to sort by
	// descending order.
	Sorting param.Field[[]string] `query:"sorting"`
}

// URLQuery serializes [DonationPublicSearchParams]'s query parameters as
// `url.Values`.
func (r DonationPublicSearchParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
