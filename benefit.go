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
	"github.com/polarsource/polar-go/internal/param"
	"github.com/polarsource/polar-go/internal/requestconfig"
	"github.com/polarsource/polar-go/option"
	"github.com/tidwall/gjson"
)

// BenefitService contains methods and other services that help with interacting
// with the polar API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBenefitService] method instead.
type BenefitService struct {
	Options []option.RequestOption
	Grants  *BenefitGrantService
}

// NewBenefitService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBenefitService(opts ...option.RequestOption) (r *BenefitService) {
	r = &BenefitService{}
	r.Options = opts
	r.Grants = NewBenefitGrantService(opts...)
	return
}

// Create a benefit.
func (r *BenefitService) New(ctx context.Context, body BenefitNewParams, opts ...option.RequestOption) (res *BenefitNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/benefits/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a benefit by ID.
func (r *BenefitService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *BenefitGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/benefits/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a benefit.
func (r *BenefitService) Update(ctx context.Context, id string, body BenefitUpdateParams, opts ...option.RequestOption) (res *BenefitUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/benefits/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List benefits.
func (r *BenefitService) List(ctx context.Context, query BenefitListParams, opts ...option.RequestOption) (res *ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadables, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/benefits/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a benefit.
//
// > [!WARNING] Every grants associated with the benefit will be revoked. Users
// > will lose access to the benefit.
func (r *BenefitService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/benefits/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadables struct {
	Pagination ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesPagination `json:"pagination,required"`
	Items      []ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItem     `json:"items"`
	JSON       listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesJSON       `json:"-"`
}

// listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesJSON
// contains the JSON metadata for the struct
// [ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadables]
type listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesJSON struct {
	Pagination  apijson.Field
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadables) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesJSON) RawJSON() string {
	return r.raw
}

type ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesPagination struct {
	MaxPage    int64                                                                                                                          `json:"max_page,required"`
	TotalCount int64                                                                                                                          `json:"total_count,required"`
	JSON       listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesPaginationJSON `json:"-"`
}

// listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesPaginationJSON
// contains the JSON metadata for the struct
// [ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesPagination]
type listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesPaginationJSON struct {
	MaxPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesPagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesPaginationJSON) RawJSON() string {
	return r.raw
}

// A benefit of type `articles`.
//
// Use it to grant access to posts.
type ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItem struct {
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The ID of the benefit.
	ID   string                                                                                                                    `json:"id,required" format:"uuid4"`
	Type ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsType `json:"type,required"`
	// The description of the benefit.
	Description string `json:"description,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool `json:"selectable,required"`
	// Whether the benefit is deletable.
	Deletable bool `json:"deletable,required"`
	// The ID of the organization owning the benefit.
	OrganizationID string `json:"organization_id,required" format:"uuid4"`
	// This field can have the runtime type of
	// [ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitArticlesProperties],
	// [ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitAdsProperties],
	// [ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitCustomProperties],
	// [ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitDiscordOutputProperties],
	// [ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitGitHubRepositoryProperties],
	// [ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitDownloadablesProperties].
	Properties interface{} `json:"properties"`
	// Whether the benefit is taxable.
	IsTaxApplicable bool                                                                                                                     `json:"is_tax_applicable"`
	JSON            listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemJSON `json:"-"`
	union           ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsUnion
}

// listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemJSON
// contains the JSON metadata for the struct
// [ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItem]
type listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemJSON struct {
	CreatedAt       apijson.Field
	ModifiedAt      apijson.Field
	ID              apijson.Field
	Type            apijson.Field
	Description     apijson.Field
	Selectable      apijson.Field
	Deletable       apijson.Field
	OrganizationID  apijson.Field
	Properties      apijson.Field
	IsTaxApplicable apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemJSON) RawJSON() string {
	return r.raw
}

func (r *ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItem) UnmarshalJSON(data []byte) (err error) {
	*r = ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItem{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitArticles],
// [ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitAds],
// [ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitCustom],
// [ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitDiscordOutput],
// [ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitGitHubRepository],
// [ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitDownloadables].
func (r ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItem) AsUnion() ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsUnion {
	return r.union
}

// A benefit of type `articles`.
//
// Use it to grant access to posts.
//
// Union satisfied by
// [ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitArticles],
// [ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitAds],
// [ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitCustom],
// [ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitDiscordOutput],
// [ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitGitHubRepository]
// or
// [ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitDownloadables].
type ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsUnion interface {
	implementsListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItem()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitArticles{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitAds{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitCustom{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitDiscordOutput{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitGitHubRepository{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitDownloadables{}),
		},
	)
}

// A benefit of type `articles`.
//
// Use it to grant access to posts.
type ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitArticles struct {
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
	Properties ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitArticlesProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                                                                                                                     `json:"selectable,required"`
	Type       ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitArticlesType `json:"type,required"`
	JSON       listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitArticlesJSON `json:"-"`
}

// listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitArticlesJSON
// contains the JSON metadata for the struct
// [ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitArticles]
type listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitArticlesJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Deletable      apijson.Field
	Description    apijson.Field
	ModifiedAt     apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	Selectable     apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitArticles) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitArticlesJSON) RawJSON() string {
	return r.raw
}

func (r ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitArticles) implementsListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItem() {
}

// Properties for a benefit of type `articles`.
type ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitArticlesProperties struct {
	// Whether the user can access paid articles.
	PaidArticles bool                                                                                                                                               `json:"paid_articles,required"`
	JSON         listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitArticlesPropertiesJSON `json:"-"`
}

// listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitArticlesPropertiesJSON
// contains the JSON metadata for the struct
// [ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitArticlesProperties]
type listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitArticlesPropertiesJSON struct {
	PaidArticles apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitArticlesProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitArticlesPropertiesJSON) RawJSON() string {
	return r.raw
}

type ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitArticlesType string

const (
	ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitArticlesTypeArticles ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitArticlesType = "articles"
)

func (r ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitArticlesType) IsKnown() bool {
	switch r {
	case ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitArticlesTypeArticles:
		return true
	}
	return false
}

// A benefit of type `ads`.
//
// Use it so your backers can display ads on your README, website, etc.
type ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitAds struct {
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
	// Properties for a benefit of type `ads`.
	Properties ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitAdsProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                                                                                                                `json:"selectable,required"`
	Type       ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitAdsType `json:"type,required"`
	JSON       listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitAdsJSON `json:"-"`
}

// listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitAdsJSON
// contains the JSON metadata for the struct
// [ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitAds]
type listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitAdsJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Deletable      apijson.Field
	Description    apijson.Field
	ModifiedAt     apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	Selectable     apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitAds) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitAdsJSON) RawJSON() string {
	return r.raw
}

func (r ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitAds) implementsListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItem() {
}

// Properties for a benefit of type `ads`.
type ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitAdsProperties struct {
	// The height of the displayed ad.
	ImageHeight int64 `json:"image_height"`
	// The width of the displayed ad.
	ImageWidth int64                                                                                                                                         `json:"image_width"`
	JSON       listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitAdsPropertiesJSON `json:"-"`
}

// listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitAdsPropertiesJSON
// contains the JSON metadata for the struct
// [ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitAdsProperties]
type listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitAdsPropertiesJSON struct {
	ImageHeight apijson.Field
	ImageWidth  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitAdsProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitAdsPropertiesJSON) RawJSON() string {
	return r.raw
}

type ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitAdsType string

const (
	ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitAdsTypeAds ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitAdsType = "ads"
)

func (r ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitAdsType) IsKnown() bool {
	switch r {
	case ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitAdsTypeAds:
		return true
	}
	return false
}

// A benefit of type `custom`.
//
// Use it to grant any kind of benefit that doesn't fit in the other types.
type ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitCustom struct {
	// The ID of the benefit.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the benefit is deletable.
	Deletable bool `json:"deletable,required"`
	// The description of the benefit.
	Description string `json:"description,required"`
	// Whether the benefit is taxable.
	IsTaxApplicable bool `json:"is_tax_applicable,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The ID of the organization owning the benefit.
	OrganizationID string `json:"organization_id,required" format:"uuid4"`
	// Properties for a benefit of type `custom`.
	Properties ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitCustomProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                                                                                                                   `json:"selectable,required"`
	Type       ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitCustomType `json:"type,required"`
	JSON       listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitCustomJSON `json:"-"`
}

// listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitCustomJSON
// contains the JSON metadata for the struct
// [ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitCustom]
type listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitCustomJSON struct {
	ID              apijson.Field
	CreatedAt       apijson.Field
	Deletable       apijson.Field
	Description     apijson.Field
	IsTaxApplicable apijson.Field
	ModifiedAt      apijson.Field
	OrganizationID  apijson.Field
	Properties      apijson.Field
	Selectable      apijson.Field
	Type            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitCustomJSON) RawJSON() string {
	return r.raw
}

func (r ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitCustom) implementsListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItem() {
}

// Properties for a benefit of type `custom`.
type ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitCustomProperties struct {
	// Private note to be shared with users who have this benefit granted.
	Note string                                                                                                                                           `json:"note,required,nullable"`
	JSON listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitCustomPropertiesJSON `json:"-"`
}

// listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitCustomPropertiesJSON
// contains the JSON metadata for the struct
// [ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitCustomProperties]
type listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitCustomPropertiesJSON struct {
	Note        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitCustomProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitCustomPropertiesJSON) RawJSON() string {
	return r.raw
}

type ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitCustomType string

const (
	ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitCustomTypeCustom ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitCustomType = "custom"
)

func (r ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitCustomType) IsKnown() bool {
	switch r {
	case ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitCustomTypeCustom:
		return true
	}
	return false
}

// A benefit of type `discord`.
//
// Use it to automatically invite your backers to a Discord server.
type ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitDiscordOutput struct {
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
	// Properties for a benefit of type `discord`.
	Properties ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitDiscordOutputProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                                                                                                                          `json:"selectable,required"`
	Type       ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitDiscordOutputType `json:"type,required"`
	JSON       listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitDiscordOutputJSON `json:"-"`
}

// listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitDiscordOutputJSON
// contains the JSON metadata for the struct
// [ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitDiscordOutput]
type listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitDiscordOutputJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Deletable      apijson.Field
	Description    apijson.Field
	ModifiedAt     apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	Selectable     apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitDiscordOutput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitDiscordOutputJSON) RawJSON() string {
	return r.raw
}

func (r ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitDiscordOutput) implementsListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItem() {
}

// Properties for a benefit of type `discord`.
type ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitDiscordOutputProperties struct {
	// The ID of the Discord server.
	GuildID    string `json:"guild_id,required"`
	GuildToken string `json:"guild_token,required"`
	// The ID of the Discord role to grant.
	RoleID string                                                                                                                                                  `json:"role_id,required"`
	JSON   listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitDiscordOutputPropertiesJSON `json:"-"`
}

// listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitDiscordOutputPropertiesJSON
// contains the JSON metadata for the struct
// [ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitDiscordOutputProperties]
type listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitDiscordOutputPropertiesJSON struct {
	GuildID     apijson.Field
	GuildToken  apijson.Field
	RoleID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitDiscordOutputProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitDiscordOutputPropertiesJSON) RawJSON() string {
	return r.raw
}

type ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitDiscordOutputType string

const (
	ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitDiscordOutputTypeDiscord ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitDiscordOutputType = "discord"
)

func (r ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitDiscordOutputType) IsKnown() bool {
	switch r {
	case ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitDiscordOutputTypeDiscord:
		return true
	}
	return false
}

// A benefit of type `github_repository`.
//
// Use it to automatically invite your backers to a private GitHub repository.
type ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitGitHubRepository struct {
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
	// Properties for a benefit of type `github_repository`.
	Properties ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitGitHubRepositoryProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                                                                                                                             `json:"selectable,required"`
	Type       ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitGitHubRepositoryType `json:"type,required"`
	JSON       listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitGitHubRepositoryJSON `json:"-"`
}

// listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitGitHubRepositoryJSON
// contains the JSON metadata for the struct
// [ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitGitHubRepository]
type listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitGitHubRepositoryJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Deletable      apijson.Field
	Description    apijson.Field
	ModifiedAt     apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	Selectable     apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitGitHubRepository) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitGitHubRepositoryJSON) RawJSON() string {
	return r.raw
}

func (r ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitGitHubRepository) implementsListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItem() {
}

// Properties for a benefit of type `github_repository`.
type ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitGitHubRepositoryProperties struct {
	// The permission level to grant. Read more about roles and their permissions on
	// [GitHub documentation](https://docs.github.com/en/organizations/managing-user-access-to-your-organizations-repositories/managing-repository-roles/repository-roles-for-an-organization#permissions-for-each-role).
	Permission   ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitGitHubRepositoryPropertiesPermission `json:"permission,required"`
	RepositoryID string                                                                                                                                                           `json:"repository_id,required,nullable" format:"uuid4"`
	// The name of the repository.
	RepositoryName string `json:"repository_name,required"`
	// The owner of the repository.
	RepositoryOwner string                                                                                                                                                     `json:"repository_owner,required"`
	JSON            listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitGitHubRepositoryPropertiesJSON `json:"-"`
}

// listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitGitHubRepositoryPropertiesJSON
// contains the JSON metadata for the struct
// [ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitGitHubRepositoryProperties]
type listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitGitHubRepositoryPropertiesJSON struct {
	Permission      apijson.Field
	RepositoryID    apijson.Field
	RepositoryName  apijson.Field
	RepositoryOwner apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitGitHubRepositoryProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitGitHubRepositoryPropertiesJSON) RawJSON() string {
	return r.raw
}

// The permission level to grant. Read more about roles and their permissions on
// [GitHub documentation](https://docs.github.com/en/organizations/managing-user-access-to-your-organizations-repositories/managing-repository-roles/repository-roles-for-an-organization#permissions-for-each-role).
type ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitGitHubRepositoryPropertiesPermission string

const (
	ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitGitHubRepositoryPropertiesPermissionPull     ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitGitHubRepositoryPropertiesPermission = "pull"
	ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitGitHubRepositoryPropertiesPermissionTriage   ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitGitHubRepositoryPropertiesPermission = "triage"
	ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitGitHubRepositoryPropertiesPermissionPush     ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitGitHubRepositoryPropertiesPermission = "push"
	ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitGitHubRepositoryPropertiesPermissionMaintain ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitGitHubRepositoryPropertiesPermission = "maintain"
	ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitGitHubRepositoryPropertiesPermissionAdmin    ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitGitHubRepositoryPropertiesPermission = "admin"
)

func (r ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitGitHubRepositoryPropertiesPermission) IsKnown() bool {
	switch r {
	case ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitGitHubRepositoryPropertiesPermissionPull, ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitGitHubRepositoryPropertiesPermissionTriage, ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitGitHubRepositoryPropertiesPermissionPush, ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitGitHubRepositoryPropertiesPermissionMaintain, ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitGitHubRepositoryPropertiesPermissionAdmin:
		return true
	}
	return false
}

type ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitGitHubRepositoryType string

const (
	ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitGitHubRepositoryTypeGitHubRepository ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitGitHubRepositoryType = "github_repository"
)

func (r ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitGitHubRepositoryType) IsKnown() bool {
	switch r {
	case ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitGitHubRepositoryTypeGitHubRepository:
		return true
	}
	return false
}

type ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitDownloadables struct {
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
	OrganizationID string                                                                                                                                              `json:"organization_id,required" format:"uuid4"`
	Properties     ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitDownloadablesProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                                                                                                                          `json:"selectable,required"`
	Type       ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitDownloadablesType `json:"type,required"`
	JSON       listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitDownloadablesJSON `json:"-"`
}

// listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitDownloadablesJSON
// contains the JSON metadata for the struct
// [ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitDownloadables]
type listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitDownloadablesJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Deletable      apijson.Field
	Description    apijson.Field
	ModifiedAt     apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	Selectable     apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitDownloadables) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitDownloadablesJSON) RawJSON() string {
	return r.raw
}

func (r ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitDownloadables) implementsListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItem() {
}

type ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitDownloadablesProperties struct {
	Archived map[string]bool                                                                                                                                         `json:"archived,required"`
	Files    []string                                                                                                                                                `json:"files,required" format:"uuid4"`
	JSON     listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitDownloadablesPropertiesJSON `json:"-"`
}

// listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitDownloadablesPropertiesJSON
// contains the JSON metadata for the struct
// [ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitDownloadablesProperties]
type listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitDownloadablesPropertiesJSON struct {
	Archived    apijson.Field
	Files       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitDownloadablesProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitDownloadablesPropertiesJSON) RawJSON() string {
	return r.raw
}

type ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitDownloadablesType string

const (
	ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitDownloadablesTypeDownloadables ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitDownloadablesType = "downloadables"
)

func (r ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitDownloadablesType) IsKnown() bool {
	switch r {
	case ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsBenefitDownloadablesTypeDownloadables:
		return true
	}
	return false
}

type ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsType string

const (
	ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsTypeArticles         ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsType = "articles"
	ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsTypeAds              ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsType = "ads"
	ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsTypeCustom           ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsType = "custom"
	ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsTypeDiscord          ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsType = "discord"
	ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsTypeGitHubRepository ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsType = "github_repository"
	ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsTypeDownloadables    ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsType = "downloadables"
)

func (r ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsType) IsKnown() bool {
	switch r {
	case ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsTypeArticles, ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsTypeAds, ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsTypeCustom, ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsTypeDiscord, ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsTypeGitHubRepository, ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadablesItemsTypeDownloadables:
		return true
	}
	return false
}

// A benefit of type `articles`.
//
// Use it to grant access to posts.
type BenefitNewResponse struct {
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The ID of the benefit.
	ID   string                 `json:"id,required" format:"uuid4"`
	Type BenefitNewResponseType `json:"type,required"`
	// The description of the benefit.
	Description string `json:"description,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool `json:"selectable,required"`
	// Whether the benefit is deletable.
	Deletable bool `json:"deletable,required"`
	// The ID of the organization owning the benefit.
	OrganizationID string `json:"organization_id,required" format:"uuid4"`
	// This field can have the runtime type of
	// [BenefitNewResponseBenefitArticlesProperties],
	// [BenefitNewResponseBenefitAdsProperties],
	// [BenefitNewResponseBenefitCustomProperties],
	// [BenefitNewResponseBenefitDiscordOutputProperties],
	// [BenefitNewResponseBenefitGitHubRepositoryProperties],
	// [BenefitNewResponseBenefitDownloadablesProperties].
	Properties interface{} `json:"properties"`
	// Whether the benefit is taxable.
	IsTaxApplicable bool                   `json:"is_tax_applicable"`
	JSON            benefitNewResponseJSON `json:"-"`
	union           BenefitNewResponseUnion
}

// benefitNewResponseJSON contains the JSON metadata for the struct
// [BenefitNewResponse]
type benefitNewResponseJSON struct {
	CreatedAt       apijson.Field
	ModifiedAt      apijson.Field
	ID              apijson.Field
	Type            apijson.Field
	Description     apijson.Field
	Selectable      apijson.Field
	Deletable       apijson.Field
	OrganizationID  apijson.Field
	Properties      apijson.Field
	IsTaxApplicable apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r benefitNewResponseJSON) RawJSON() string {
	return r.raw
}

func (r *BenefitNewResponse) UnmarshalJSON(data []byte) (err error) {
	*r = BenefitNewResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [BenefitNewResponseUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [BenefitNewResponseBenefitArticles],
// [BenefitNewResponseBenefitAds], [BenefitNewResponseBenefitCustom],
// [BenefitNewResponseBenefitDiscordOutput],
// [BenefitNewResponseBenefitGitHubRepository],
// [BenefitNewResponseBenefitDownloadables].
func (r BenefitNewResponse) AsUnion() BenefitNewResponseUnion {
	return r.union
}

// A benefit of type `articles`.
//
// Use it to grant access to posts.
//
// Union satisfied by [BenefitNewResponseBenefitArticles],
// [BenefitNewResponseBenefitAds], [BenefitNewResponseBenefitCustom],
// [BenefitNewResponseBenefitDiscordOutput],
// [BenefitNewResponseBenefitGitHubRepository] or
// [BenefitNewResponseBenefitDownloadables].
type BenefitNewResponseUnion interface {
	implementsBenefitNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BenefitNewResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BenefitNewResponseBenefitArticles{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BenefitNewResponseBenefitAds{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BenefitNewResponseBenefitCustom{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BenefitNewResponseBenefitDiscordOutput{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BenefitNewResponseBenefitGitHubRepository{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BenefitNewResponseBenefitDownloadables{}),
		},
	)
}

// A benefit of type `articles`.
//
// Use it to grant access to posts.
type BenefitNewResponseBenefitArticles struct {
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
	Properties BenefitNewResponseBenefitArticlesProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                  `json:"selectable,required"`
	Type       BenefitNewResponseBenefitArticlesType `json:"type,required"`
	JSON       benefitNewResponseBenefitArticlesJSON `json:"-"`
}

// benefitNewResponseBenefitArticlesJSON contains the JSON metadata for the struct
// [BenefitNewResponseBenefitArticles]
type benefitNewResponseBenefitArticlesJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Deletable      apijson.Field
	Description    apijson.Field
	ModifiedAt     apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	Selectable     apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BenefitNewResponseBenefitArticles) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitNewResponseBenefitArticlesJSON) RawJSON() string {
	return r.raw
}

func (r BenefitNewResponseBenefitArticles) implementsBenefitNewResponse() {}

// Properties for a benefit of type `articles`.
type BenefitNewResponseBenefitArticlesProperties struct {
	// Whether the user can access paid articles.
	PaidArticles bool                                            `json:"paid_articles,required"`
	JSON         benefitNewResponseBenefitArticlesPropertiesJSON `json:"-"`
}

// benefitNewResponseBenefitArticlesPropertiesJSON contains the JSON metadata for
// the struct [BenefitNewResponseBenefitArticlesProperties]
type benefitNewResponseBenefitArticlesPropertiesJSON struct {
	PaidArticles apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *BenefitNewResponseBenefitArticlesProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitNewResponseBenefitArticlesPropertiesJSON) RawJSON() string {
	return r.raw
}

type BenefitNewResponseBenefitArticlesType string

const (
	BenefitNewResponseBenefitArticlesTypeArticles BenefitNewResponseBenefitArticlesType = "articles"
)

func (r BenefitNewResponseBenefitArticlesType) IsKnown() bool {
	switch r {
	case BenefitNewResponseBenefitArticlesTypeArticles:
		return true
	}
	return false
}

// A benefit of type `ads`.
//
// Use it so your backers can display ads on your README, website, etc.
type BenefitNewResponseBenefitAds struct {
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
	// Properties for a benefit of type `ads`.
	Properties BenefitNewResponseBenefitAdsProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                             `json:"selectable,required"`
	Type       BenefitNewResponseBenefitAdsType `json:"type,required"`
	JSON       benefitNewResponseBenefitAdsJSON `json:"-"`
}

// benefitNewResponseBenefitAdsJSON contains the JSON metadata for the struct
// [BenefitNewResponseBenefitAds]
type benefitNewResponseBenefitAdsJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Deletable      apijson.Field
	Description    apijson.Field
	ModifiedAt     apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	Selectable     apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BenefitNewResponseBenefitAds) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitNewResponseBenefitAdsJSON) RawJSON() string {
	return r.raw
}

func (r BenefitNewResponseBenefitAds) implementsBenefitNewResponse() {}

// Properties for a benefit of type `ads`.
type BenefitNewResponseBenefitAdsProperties struct {
	// The height of the displayed ad.
	ImageHeight int64 `json:"image_height"`
	// The width of the displayed ad.
	ImageWidth int64                                      `json:"image_width"`
	JSON       benefitNewResponseBenefitAdsPropertiesJSON `json:"-"`
}

// benefitNewResponseBenefitAdsPropertiesJSON contains the JSON metadata for the
// struct [BenefitNewResponseBenefitAdsProperties]
type benefitNewResponseBenefitAdsPropertiesJSON struct {
	ImageHeight apijson.Field
	ImageWidth  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BenefitNewResponseBenefitAdsProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitNewResponseBenefitAdsPropertiesJSON) RawJSON() string {
	return r.raw
}

type BenefitNewResponseBenefitAdsType string

const (
	BenefitNewResponseBenefitAdsTypeAds BenefitNewResponseBenefitAdsType = "ads"
)

func (r BenefitNewResponseBenefitAdsType) IsKnown() bool {
	switch r {
	case BenefitNewResponseBenefitAdsTypeAds:
		return true
	}
	return false
}

// A benefit of type `custom`.
//
// Use it to grant any kind of benefit that doesn't fit in the other types.
type BenefitNewResponseBenefitCustom struct {
	// The ID of the benefit.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the benefit is deletable.
	Deletable bool `json:"deletable,required"`
	// The description of the benefit.
	Description string `json:"description,required"`
	// Whether the benefit is taxable.
	IsTaxApplicable bool `json:"is_tax_applicable,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The ID of the organization owning the benefit.
	OrganizationID string `json:"organization_id,required" format:"uuid4"`
	// Properties for a benefit of type `custom`.
	Properties BenefitNewResponseBenefitCustomProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                `json:"selectable,required"`
	Type       BenefitNewResponseBenefitCustomType `json:"type,required"`
	JSON       benefitNewResponseBenefitCustomJSON `json:"-"`
}

// benefitNewResponseBenefitCustomJSON contains the JSON metadata for the struct
// [BenefitNewResponseBenefitCustom]
type benefitNewResponseBenefitCustomJSON struct {
	ID              apijson.Field
	CreatedAt       apijson.Field
	Deletable       apijson.Field
	Description     apijson.Field
	IsTaxApplicable apijson.Field
	ModifiedAt      apijson.Field
	OrganizationID  apijson.Field
	Properties      apijson.Field
	Selectable      apijson.Field
	Type            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *BenefitNewResponseBenefitCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitNewResponseBenefitCustomJSON) RawJSON() string {
	return r.raw
}

func (r BenefitNewResponseBenefitCustom) implementsBenefitNewResponse() {}

// Properties for a benefit of type `custom`.
type BenefitNewResponseBenefitCustomProperties struct {
	// Private note to be shared with users who have this benefit granted.
	Note string                                        `json:"note,required,nullable"`
	JSON benefitNewResponseBenefitCustomPropertiesJSON `json:"-"`
}

// benefitNewResponseBenefitCustomPropertiesJSON contains the JSON metadata for the
// struct [BenefitNewResponseBenefitCustomProperties]
type benefitNewResponseBenefitCustomPropertiesJSON struct {
	Note        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BenefitNewResponseBenefitCustomProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitNewResponseBenefitCustomPropertiesJSON) RawJSON() string {
	return r.raw
}

type BenefitNewResponseBenefitCustomType string

const (
	BenefitNewResponseBenefitCustomTypeCustom BenefitNewResponseBenefitCustomType = "custom"
)

func (r BenefitNewResponseBenefitCustomType) IsKnown() bool {
	switch r {
	case BenefitNewResponseBenefitCustomTypeCustom:
		return true
	}
	return false
}

// A benefit of type `discord`.
//
// Use it to automatically invite your backers to a Discord server.
type BenefitNewResponseBenefitDiscordOutput struct {
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
	// Properties for a benefit of type `discord`.
	Properties BenefitNewResponseBenefitDiscordOutputProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                       `json:"selectable,required"`
	Type       BenefitNewResponseBenefitDiscordOutputType `json:"type,required"`
	JSON       benefitNewResponseBenefitDiscordOutputJSON `json:"-"`
}

// benefitNewResponseBenefitDiscordOutputJSON contains the JSON metadata for the
// struct [BenefitNewResponseBenefitDiscordOutput]
type benefitNewResponseBenefitDiscordOutputJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Deletable      apijson.Field
	Description    apijson.Field
	ModifiedAt     apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	Selectable     apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BenefitNewResponseBenefitDiscordOutput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitNewResponseBenefitDiscordOutputJSON) RawJSON() string {
	return r.raw
}

func (r BenefitNewResponseBenefitDiscordOutput) implementsBenefitNewResponse() {}

// Properties for a benefit of type `discord`.
type BenefitNewResponseBenefitDiscordOutputProperties struct {
	// The ID of the Discord server.
	GuildID    string `json:"guild_id,required"`
	GuildToken string `json:"guild_token,required"`
	// The ID of the Discord role to grant.
	RoleID string                                               `json:"role_id,required"`
	JSON   benefitNewResponseBenefitDiscordOutputPropertiesJSON `json:"-"`
}

// benefitNewResponseBenefitDiscordOutputPropertiesJSON contains the JSON metadata
// for the struct [BenefitNewResponseBenefitDiscordOutputProperties]
type benefitNewResponseBenefitDiscordOutputPropertiesJSON struct {
	GuildID     apijson.Field
	GuildToken  apijson.Field
	RoleID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BenefitNewResponseBenefitDiscordOutputProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitNewResponseBenefitDiscordOutputPropertiesJSON) RawJSON() string {
	return r.raw
}

type BenefitNewResponseBenefitDiscordOutputType string

const (
	BenefitNewResponseBenefitDiscordOutputTypeDiscord BenefitNewResponseBenefitDiscordOutputType = "discord"
)

func (r BenefitNewResponseBenefitDiscordOutputType) IsKnown() bool {
	switch r {
	case BenefitNewResponseBenefitDiscordOutputTypeDiscord:
		return true
	}
	return false
}

// A benefit of type `github_repository`.
//
// Use it to automatically invite your backers to a private GitHub repository.
type BenefitNewResponseBenefitGitHubRepository struct {
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
	// Properties for a benefit of type `github_repository`.
	Properties BenefitNewResponseBenefitGitHubRepositoryProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                          `json:"selectable,required"`
	Type       BenefitNewResponseBenefitGitHubRepositoryType `json:"type,required"`
	JSON       benefitNewResponseBenefitGitHubRepositoryJSON `json:"-"`
}

// benefitNewResponseBenefitGitHubRepositoryJSON contains the JSON metadata for the
// struct [BenefitNewResponseBenefitGitHubRepository]
type benefitNewResponseBenefitGitHubRepositoryJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Deletable      apijson.Field
	Description    apijson.Field
	ModifiedAt     apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	Selectable     apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BenefitNewResponseBenefitGitHubRepository) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitNewResponseBenefitGitHubRepositoryJSON) RawJSON() string {
	return r.raw
}

func (r BenefitNewResponseBenefitGitHubRepository) implementsBenefitNewResponse() {}

// Properties for a benefit of type `github_repository`.
type BenefitNewResponseBenefitGitHubRepositoryProperties struct {
	// The permission level to grant. Read more about roles and their permissions on
	// [GitHub documentation](https://docs.github.com/en/organizations/managing-user-access-to-your-organizations-repositories/managing-repository-roles/repository-roles-for-an-organization#permissions-for-each-role).
	Permission   BenefitNewResponseBenefitGitHubRepositoryPropertiesPermission `json:"permission,required"`
	RepositoryID string                                                        `json:"repository_id,required,nullable" format:"uuid4"`
	// The name of the repository.
	RepositoryName string `json:"repository_name,required"`
	// The owner of the repository.
	RepositoryOwner string                                                  `json:"repository_owner,required"`
	JSON            benefitNewResponseBenefitGitHubRepositoryPropertiesJSON `json:"-"`
}

// benefitNewResponseBenefitGitHubRepositoryPropertiesJSON contains the JSON
// metadata for the struct [BenefitNewResponseBenefitGitHubRepositoryProperties]
type benefitNewResponseBenefitGitHubRepositoryPropertiesJSON struct {
	Permission      apijson.Field
	RepositoryID    apijson.Field
	RepositoryName  apijson.Field
	RepositoryOwner apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *BenefitNewResponseBenefitGitHubRepositoryProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitNewResponseBenefitGitHubRepositoryPropertiesJSON) RawJSON() string {
	return r.raw
}

// The permission level to grant. Read more about roles and their permissions on
// [GitHub documentation](https://docs.github.com/en/organizations/managing-user-access-to-your-organizations-repositories/managing-repository-roles/repository-roles-for-an-organization#permissions-for-each-role).
type BenefitNewResponseBenefitGitHubRepositoryPropertiesPermission string

const (
	BenefitNewResponseBenefitGitHubRepositoryPropertiesPermissionPull     BenefitNewResponseBenefitGitHubRepositoryPropertiesPermission = "pull"
	BenefitNewResponseBenefitGitHubRepositoryPropertiesPermissionTriage   BenefitNewResponseBenefitGitHubRepositoryPropertiesPermission = "triage"
	BenefitNewResponseBenefitGitHubRepositoryPropertiesPermissionPush     BenefitNewResponseBenefitGitHubRepositoryPropertiesPermission = "push"
	BenefitNewResponseBenefitGitHubRepositoryPropertiesPermissionMaintain BenefitNewResponseBenefitGitHubRepositoryPropertiesPermission = "maintain"
	BenefitNewResponseBenefitGitHubRepositoryPropertiesPermissionAdmin    BenefitNewResponseBenefitGitHubRepositoryPropertiesPermission = "admin"
)

func (r BenefitNewResponseBenefitGitHubRepositoryPropertiesPermission) IsKnown() bool {
	switch r {
	case BenefitNewResponseBenefitGitHubRepositoryPropertiesPermissionPull, BenefitNewResponseBenefitGitHubRepositoryPropertiesPermissionTriage, BenefitNewResponseBenefitGitHubRepositoryPropertiesPermissionPush, BenefitNewResponseBenefitGitHubRepositoryPropertiesPermissionMaintain, BenefitNewResponseBenefitGitHubRepositoryPropertiesPermissionAdmin:
		return true
	}
	return false
}

type BenefitNewResponseBenefitGitHubRepositoryType string

const (
	BenefitNewResponseBenefitGitHubRepositoryTypeGitHubRepository BenefitNewResponseBenefitGitHubRepositoryType = "github_repository"
)

func (r BenefitNewResponseBenefitGitHubRepositoryType) IsKnown() bool {
	switch r {
	case BenefitNewResponseBenefitGitHubRepositoryTypeGitHubRepository:
		return true
	}
	return false
}

type BenefitNewResponseBenefitDownloadables struct {
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
	OrganizationID string                                           `json:"organization_id,required" format:"uuid4"`
	Properties     BenefitNewResponseBenefitDownloadablesProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                       `json:"selectable,required"`
	Type       BenefitNewResponseBenefitDownloadablesType `json:"type,required"`
	JSON       benefitNewResponseBenefitDownloadablesJSON `json:"-"`
}

// benefitNewResponseBenefitDownloadablesJSON contains the JSON metadata for the
// struct [BenefitNewResponseBenefitDownloadables]
type benefitNewResponseBenefitDownloadablesJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Deletable      apijson.Field
	Description    apijson.Field
	ModifiedAt     apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	Selectable     apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BenefitNewResponseBenefitDownloadables) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitNewResponseBenefitDownloadablesJSON) RawJSON() string {
	return r.raw
}

func (r BenefitNewResponseBenefitDownloadables) implementsBenefitNewResponse() {}

type BenefitNewResponseBenefitDownloadablesProperties struct {
	Archived map[string]bool                                      `json:"archived,required"`
	Files    []string                                             `json:"files,required" format:"uuid4"`
	JSON     benefitNewResponseBenefitDownloadablesPropertiesJSON `json:"-"`
}

// benefitNewResponseBenefitDownloadablesPropertiesJSON contains the JSON metadata
// for the struct [BenefitNewResponseBenefitDownloadablesProperties]
type benefitNewResponseBenefitDownloadablesPropertiesJSON struct {
	Archived    apijson.Field
	Files       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BenefitNewResponseBenefitDownloadablesProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitNewResponseBenefitDownloadablesPropertiesJSON) RawJSON() string {
	return r.raw
}

type BenefitNewResponseBenefitDownloadablesType string

const (
	BenefitNewResponseBenefitDownloadablesTypeDownloadables BenefitNewResponseBenefitDownloadablesType = "downloadables"
)

func (r BenefitNewResponseBenefitDownloadablesType) IsKnown() bool {
	switch r {
	case BenefitNewResponseBenefitDownloadablesTypeDownloadables:
		return true
	}
	return false
}

type BenefitNewResponseType string

const (
	BenefitNewResponseTypeArticles         BenefitNewResponseType = "articles"
	BenefitNewResponseTypeAds              BenefitNewResponseType = "ads"
	BenefitNewResponseTypeCustom           BenefitNewResponseType = "custom"
	BenefitNewResponseTypeDiscord          BenefitNewResponseType = "discord"
	BenefitNewResponseTypeGitHubRepository BenefitNewResponseType = "github_repository"
	BenefitNewResponseTypeDownloadables    BenefitNewResponseType = "downloadables"
)

func (r BenefitNewResponseType) IsKnown() bool {
	switch r {
	case BenefitNewResponseTypeArticles, BenefitNewResponseTypeAds, BenefitNewResponseTypeCustom, BenefitNewResponseTypeDiscord, BenefitNewResponseTypeGitHubRepository, BenefitNewResponseTypeDownloadables:
		return true
	}
	return false
}

// A benefit of type `articles`.
//
// Use it to grant access to posts.
type BenefitGetResponse struct {
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The ID of the benefit.
	ID   string                 `json:"id,required" format:"uuid4"`
	Type BenefitGetResponseType `json:"type,required"`
	// The description of the benefit.
	Description string `json:"description,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool `json:"selectable,required"`
	// Whether the benefit is deletable.
	Deletable bool `json:"deletable,required"`
	// The ID of the organization owning the benefit.
	OrganizationID string `json:"organization_id,required" format:"uuid4"`
	// This field can have the runtime type of
	// [BenefitGetResponseBenefitArticlesProperties],
	// [BenefitGetResponseBenefitAdsProperties],
	// [BenefitGetResponseBenefitCustomProperties],
	// [BenefitGetResponseBenefitDiscordOutputProperties],
	// [BenefitGetResponseBenefitGitHubRepositoryProperties],
	// [BenefitGetResponseBenefitDownloadablesProperties].
	Properties interface{} `json:"properties"`
	// Whether the benefit is taxable.
	IsTaxApplicable bool                   `json:"is_tax_applicable"`
	JSON            benefitGetResponseJSON `json:"-"`
	union           BenefitGetResponseUnion
}

// benefitGetResponseJSON contains the JSON metadata for the struct
// [BenefitGetResponse]
type benefitGetResponseJSON struct {
	CreatedAt       apijson.Field
	ModifiedAt      apijson.Field
	ID              apijson.Field
	Type            apijson.Field
	Description     apijson.Field
	Selectable      apijson.Field
	Deletable       apijson.Field
	OrganizationID  apijson.Field
	Properties      apijson.Field
	IsTaxApplicable apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r benefitGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *BenefitGetResponse) UnmarshalJSON(data []byte) (err error) {
	*r = BenefitGetResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [BenefitGetResponseUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [BenefitGetResponseBenefitArticles],
// [BenefitGetResponseBenefitAds], [BenefitGetResponseBenefitCustom],
// [BenefitGetResponseBenefitDiscordOutput],
// [BenefitGetResponseBenefitGitHubRepository],
// [BenefitGetResponseBenefitDownloadables].
func (r BenefitGetResponse) AsUnion() BenefitGetResponseUnion {
	return r.union
}

// A benefit of type `articles`.
//
// Use it to grant access to posts.
//
// Union satisfied by [BenefitGetResponseBenefitArticles],
// [BenefitGetResponseBenefitAds], [BenefitGetResponseBenefitCustom],
// [BenefitGetResponseBenefitDiscordOutput],
// [BenefitGetResponseBenefitGitHubRepository] or
// [BenefitGetResponseBenefitDownloadables].
type BenefitGetResponseUnion interface {
	implementsBenefitGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BenefitGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BenefitGetResponseBenefitArticles{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BenefitGetResponseBenefitAds{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BenefitGetResponseBenefitCustom{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BenefitGetResponseBenefitDiscordOutput{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BenefitGetResponseBenefitGitHubRepository{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BenefitGetResponseBenefitDownloadables{}),
		},
	)
}

// A benefit of type `articles`.
//
// Use it to grant access to posts.
type BenefitGetResponseBenefitArticles struct {
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
	Properties BenefitGetResponseBenefitArticlesProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                  `json:"selectable,required"`
	Type       BenefitGetResponseBenefitArticlesType `json:"type,required"`
	JSON       benefitGetResponseBenefitArticlesJSON `json:"-"`
}

// benefitGetResponseBenefitArticlesJSON contains the JSON metadata for the struct
// [BenefitGetResponseBenefitArticles]
type benefitGetResponseBenefitArticlesJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Deletable      apijson.Field
	Description    apijson.Field
	ModifiedAt     apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	Selectable     apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BenefitGetResponseBenefitArticles) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitGetResponseBenefitArticlesJSON) RawJSON() string {
	return r.raw
}

func (r BenefitGetResponseBenefitArticles) implementsBenefitGetResponse() {}

// Properties for a benefit of type `articles`.
type BenefitGetResponseBenefitArticlesProperties struct {
	// Whether the user can access paid articles.
	PaidArticles bool                                            `json:"paid_articles,required"`
	JSON         benefitGetResponseBenefitArticlesPropertiesJSON `json:"-"`
}

// benefitGetResponseBenefitArticlesPropertiesJSON contains the JSON metadata for
// the struct [BenefitGetResponseBenefitArticlesProperties]
type benefitGetResponseBenefitArticlesPropertiesJSON struct {
	PaidArticles apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *BenefitGetResponseBenefitArticlesProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitGetResponseBenefitArticlesPropertiesJSON) RawJSON() string {
	return r.raw
}

type BenefitGetResponseBenefitArticlesType string

const (
	BenefitGetResponseBenefitArticlesTypeArticles BenefitGetResponseBenefitArticlesType = "articles"
)

func (r BenefitGetResponseBenefitArticlesType) IsKnown() bool {
	switch r {
	case BenefitGetResponseBenefitArticlesTypeArticles:
		return true
	}
	return false
}

// A benefit of type `ads`.
//
// Use it so your backers can display ads on your README, website, etc.
type BenefitGetResponseBenefitAds struct {
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
	// Properties for a benefit of type `ads`.
	Properties BenefitGetResponseBenefitAdsProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                             `json:"selectable,required"`
	Type       BenefitGetResponseBenefitAdsType `json:"type,required"`
	JSON       benefitGetResponseBenefitAdsJSON `json:"-"`
}

// benefitGetResponseBenefitAdsJSON contains the JSON metadata for the struct
// [BenefitGetResponseBenefitAds]
type benefitGetResponseBenefitAdsJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Deletable      apijson.Field
	Description    apijson.Field
	ModifiedAt     apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	Selectable     apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BenefitGetResponseBenefitAds) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitGetResponseBenefitAdsJSON) RawJSON() string {
	return r.raw
}

func (r BenefitGetResponseBenefitAds) implementsBenefitGetResponse() {}

// Properties for a benefit of type `ads`.
type BenefitGetResponseBenefitAdsProperties struct {
	// The height of the displayed ad.
	ImageHeight int64 `json:"image_height"`
	// The width of the displayed ad.
	ImageWidth int64                                      `json:"image_width"`
	JSON       benefitGetResponseBenefitAdsPropertiesJSON `json:"-"`
}

// benefitGetResponseBenefitAdsPropertiesJSON contains the JSON metadata for the
// struct [BenefitGetResponseBenefitAdsProperties]
type benefitGetResponseBenefitAdsPropertiesJSON struct {
	ImageHeight apijson.Field
	ImageWidth  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BenefitGetResponseBenefitAdsProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitGetResponseBenefitAdsPropertiesJSON) RawJSON() string {
	return r.raw
}

type BenefitGetResponseBenefitAdsType string

const (
	BenefitGetResponseBenefitAdsTypeAds BenefitGetResponseBenefitAdsType = "ads"
)

func (r BenefitGetResponseBenefitAdsType) IsKnown() bool {
	switch r {
	case BenefitGetResponseBenefitAdsTypeAds:
		return true
	}
	return false
}

// A benefit of type `custom`.
//
// Use it to grant any kind of benefit that doesn't fit in the other types.
type BenefitGetResponseBenefitCustom struct {
	// The ID of the benefit.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the benefit is deletable.
	Deletable bool `json:"deletable,required"`
	// The description of the benefit.
	Description string `json:"description,required"`
	// Whether the benefit is taxable.
	IsTaxApplicable bool `json:"is_tax_applicable,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The ID of the organization owning the benefit.
	OrganizationID string `json:"organization_id,required" format:"uuid4"`
	// Properties for a benefit of type `custom`.
	Properties BenefitGetResponseBenefitCustomProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                `json:"selectable,required"`
	Type       BenefitGetResponseBenefitCustomType `json:"type,required"`
	JSON       benefitGetResponseBenefitCustomJSON `json:"-"`
}

// benefitGetResponseBenefitCustomJSON contains the JSON metadata for the struct
// [BenefitGetResponseBenefitCustom]
type benefitGetResponseBenefitCustomJSON struct {
	ID              apijson.Field
	CreatedAt       apijson.Field
	Deletable       apijson.Field
	Description     apijson.Field
	IsTaxApplicable apijson.Field
	ModifiedAt      apijson.Field
	OrganizationID  apijson.Field
	Properties      apijson.Field
	Selectable      apijson.Field
	Type            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *BenefitGetResponseBenefitCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitGetResponseBenefitCustomJSON) RawJSON() string {
	return r.raw
}

func (r BenefitGetResponseBenefitCustom) implementsBenefitGetResponse() {}

// Properties for a benefit of type `custom`.
type BenefitGetResponseBenefitCustomProperties struct {
	// Private note to be shared with users who have this benefit granted.
	Note string                                        `json:"note,required,nullable"`
	JSON benefitGetResponseBenefitCustomPropertiesJSON `json:"-"`
}

// benefitGetResponseBenefitCustomPropertiesJSON contains the JSON metadata for the
// struct [BenefitGetResponseBenefitCustomProperties]
type benefitGetResponseBenefitCustomPropertiesJSON struct {
	Note        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BenefitGetResponseBenefitCustomProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitGetResponseBenefitCustomPropertiesJSON) RawJSON() string {
	return r.raw
}

type BenefitGetResponseBenefitCustomType string

const (
	BenefitGetResponseBenefitCustomTypeCustom BenefitGetResponseBenefitCustomType = "custom"
)

func (r BenefitGetResponseBenefitCustomType) IsKnown() bool {
	switch r {
	case BenefitGetResponseBenefitCustomTypeCustom:
		return true
	}
	return false
}

// A benefit of type `discord`.
//
// Use it to automatically invite your backers to a Discord server.
type BenefitGetResponseBenefitDiscordOutput struct {
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
	// Properties for a benefit of type `discord`.
	Properties BenefitGetResponseBenefitDiscordOutputProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                       `json:"selectable,required"`
	Type       BenefitGetResponseBenefitDiscordOutputType `json:"type,required"`
	JSON       benefitGetResponseBenefitDiscordOutputJSON `json:"-"`
}

// benefitGetResponseBenefitDiscordOutputJSON contains the JSON metadata for the
// struct [BenefitGetResponseBenefitDiscordOutput]
type benefitGetResponseBenefitDiscordOutputJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Deletable      apijson.Field
	Description    apijson.Field
	ModifiedAt     apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	Selectable     apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BenefitGetResponseBenefitDiscordOutput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitGetResponseBenefitDiscordOutputJSON) RawJSON() string {
	return r.raw
}

func (r BenefitGetResponseBenefitDiscordOutput) implementsBenefitGetResponse() {}

// Properties for a benefit of type `discord`.
type BenefitGetResponseBenefitDiscordOutputProperties struct {
	// The ID of the Discord server.
	GuildID    string `json:"guild_id,required"`
	GuildToken string `json:"guild_token,required"`
	// The ID of the Discord role to grant.
	RoleID string                                               `json:"role_id,required"`
	JSON   benefitGetResponseBenefitDiscordOutputPropertiesJSON `json:"-"`
}

// benefitGetResponseBenefitDiscordOutputPropertiesJSON contains the JSON metadata
// for the struct [BenefitGetResponseBenefitDiscordOutputProperties]
type benefitGetResponseBenefitDiscordOutputPropertiesJSON struct {
	GuildID     apijson.Field
	GuildToken  apijson.Field
	RoleID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BenefitGetResponseBenefitDiscordOutputProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitGetResponseBenefitDiscordOutputPropertiesJSON) RawJSON() string {
	return r.raw
}

type BenefitGetResponseBenefitDiscordOutputType string

const (
	BenefitGetResponseBenefitDiscordOutputTypeDiscord BenefitGetResponseBenefitDiscordOutputType = "discord"
)

func (r BenefitGetResponseBenefitDiscordOutputType) IsKnown() bool {
	switch r {
	case BenefitGetResponseBenefitDiscordOutputTypeDiscord:
		return true
	}
	return false
}

// A benefit of type `github_repository`.
//
// Use it to automatically invite your backers to a private GitHub repository.
type BenefitGetResponseBenefitGitHubRepository struct {
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
	// Properties for a benefit of type `github_repository`.
	Properties BenefitGetResponseBenefitGitHubRepositoryProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                          `json:"selectable,required"`
	Type       BenefitGetResponseBenefitGitHubRepositoryType `json:"type,required"`
	JSON       benefitGetResponseBenefitGitHubRepositoryJSON `json:"-"`
}

// benefitGetResponseBenefitGitHubRepositoryJSON contains the JSON metadata for the
// struct [BenefitGetResponseBenefitGitHubRepository]
type benefitGetResponseBenefitGitHubRepositoryJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Deletable      apijson.Field
	Description    apijson.Field
	ModifiedAt     apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	Selectable     apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BenefitGetResponseBenefitGitHubRepository) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitGetResponseBenefitGitHubRepositoryJSON) RawJSON() string {
	return r.raw
}

func (r BenefitGetResponseBenefitGitHubRepository) implementsBenefitGetResponse() {}

// Properties for a benefit of type `github_repository`.
type BenefitGetResponseBenefitGitHubRepositoryProperties struct {
	// The permission level to grant. Read more about roles and their permissions on
	// [GitHub documentation](https://docs.github.com/en/organizations/managing-user-access-to-your-organizations-repositories/managing-repository-roles/repository-roles-for-an-organization#permissions-for-each-role).
	Permission   BenefitGetResponseBenefitGitHubRepositoryPropertiesPermission `json:"permission,required"`
	RepositoryID string                                                        `json:"repository_id,required,nullable" format:"uuid4"`
	// The name of the repository.
	RepositoryName string `json:"repository_name,required"`
	// The owner of the repository.
	RepositoryOwner string                                                  `json:"repository_owner,required"`
	JSON            benefitGetResponseBenefitGitHubRepositoryPropertiesJSON `json:"-"`
}

// benefitGetResponseBenefitGitHubRepositoryPropertiesJSON contains the JSON
// metadata for the struct [BenefitGetResponseBenefitGitHubRepositoryProperties]
type benefitGetResponseBenefitGitHubRepositoryPropertiesJSON struct {
	Permission      apijson.Field
	RepositoryID    apijson.Field
	RepositoryName  apijson.Field
	RepositoryOwner apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *BenefitGetResponseBenefitGitHubRepositoryProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitGetResponseBenefitGitHubRepositoryPropertiesJSON) RawJSON() string {
	return r.raw
}

// The permission level to grant. Read more about roles and their permissions on
// [GitHub documentation](https://docs.github.com/en/organizations/managing-user-access-to-your-organizations-repositories/managing-repository-roles/repository-roles-for-an-organization#permissions-for-each-role).
type BenefitGetResponseBenefitGitHubRepositoryPropertiesPermission string

const (
	BenefitGetResponseBenefitGitHubRepositoryPropertiesPermissionPull     BenefitGetResponseBenefitGitHubRepositoryPropertiesPermission = "pull"
	BenefitGetResponseBenefitGitHubRepositoryPropertiesPermissionTriage   BenefitGetResponseBenefitGitHubRepositoryPropertiesPermission = "triage"
	BenefitGetResponseBenefitGitHubRepositoryPropertiesPermissionPush     BenefitGetResponseBenefitGitHubRepositoryPropertiesPermission = "push"
	BenefitGetResponseBenefitGitHubRepositoryPropertiesPermissionMaintain BenefitGetResponseBenefitGitHubRepositoryPropertiesPermission = "maintain"
	BenefitGetResponseBenefitGitHubRepositoryPropertiesPermissionAdmin    BenefitGetResponseBenefitGitHubRepositoryPropertiesPermission = "admin"
)

func (r BenefitGetResponseBenefitGitHubRepositoryPropertiesPermission) IsKnown() bool {
	switch r {
	case BenefitGetResponseBenefitGitHubRepositoryPropertiesPermissionPull, BenefitGetResponseBenefitGitHubRepositoryPropertiesPermissionTriage, BenefitGetResponseBenefitGitHubRepositoryPropertiesPermissionPush, BenefitGetResponseBenefitGitHubRepositoryPropertiesPermissionMaintain, BenefitGetResponseBenefitGitHubRepositoryPropertiesPermissionAdmin:
		return true
	}
	return false
}

type BenefitGetResponseBenefitGitHubRepositoryType string

const (
	BenefitGetResponseBenefitGitHubRepositoryTypeGitHubRepository BenefitGetResponseBenefitGitHubRepositoryType = "github_repository"
)

func (r BenefitGetResponseBenefitGitHubRepositoryType) IsKnown() bool {
	switch r {
	case BenefitGetResponseBenefitGitHubRepositoryTypeGitHubRepository:
		return true
	}
	return false
}

type BenefitGetResponseBenefitDownloadables struct {
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
	OrganizationID string                                           `json:"organization_id,required" format:"uuid4"`
	Properties     BenefitGetResponseBenefitDownloadablesProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                       `json:"selectable,required"`
	Type       BenefitGetResponseBenefitDownloadablesType `json:"type,required"`
	JSON       benefitGetResponseBenefitDownloadablesJSON `json:"-"`
}

// benefitGetResponseBenefitDownloadablesJSON contains the JSON metadata for the
// struct [BenefitGetResponseBenefitDownloadables]
type benefitGetResponseBenefitDownloadablesJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Deletable      apijson.Field
	Description    apijson.Field
	ModifiedAt     apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	Selectable     apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BenefitGetResponseBenefitDownloadables) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitGetResponseBenefitDownloadablesJSON) RawJSON() string {
	return r.raw
}

func (r BenefitGetResponseBenefitDownloadables) implementsBenefitGetResponse() {}

type BenefitGetResponseBenefitDownloadablesProperties struct {
	Archived map[string]bool                                      `json:"archived,required"`
	Files    []string                                             `json:"files,required" format:"uuid4"`
	JSON     benefitGetResponseBenefitDownloadablesPropertiesJSON `json:"-"`
}

// benefitGetResponseBenefitDownloadablesPropertiesJSON contains the JSON metadata
// for the struct [BenefitGetResponseBenefitDownloadablesProperties]
type benefitGetResponseBenefitDownloadablesPropertiesJSON struct {
	Archived    apijson.Field
	Files       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BenefitGetResponseBenefitDownloadablesProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitGetResponseBenefitDownloadablesPropertiesJSON) RawJSON() string {
	return r.raw
}

type BenefitGetResponseBenefitDownloadablesType string

const (
	BenefitGetResponseBenefitDownloadablesTypeDownloadables BenefitGetResponseBenefitDownloadablesType = "downloadables"
)

func (r BenefitGetResponseBenefitDownloadablesType) IsKnown() bool {
	switch r {
	case BenefitGetResponseBenefitDownloadablesTypeDownloadables:
		return true
	}
	return false
}

type BenefitGetResponseType string

const (
	BenefitGetResponseTypeArticles         BenefitGetResponseType = "articles"
	BenefitGetResponseTypeAds              BenefitGetResponseType = "ads"
	BenefitGetResponseTypeCustom           BenefitGetResponseType = "custom"
	BenefitGetResponseTypeDiscord          BenefitGetResponseType = "discord"
	BenefitGetResponseTypeGitHubRepository BenefitGetResponseType = "github_repository"
	BenefitGetResponseTypeDownloadables    BenefitGetResponseType = "downloadables"
)

func (r BenefitGetResponseType) IsKnown() bool {
	switch r {
	case BenefitGetResponseTypeArticles, BenefitGetResponseTypeAds, BenefitGetResponseTypeCustom, BenefitGetResponseTypeDiscord, BenefitGetResponseTypeGitHubRepository, BenefitGetResponseTypeDownloadables:
		return true
	}
	return false
}

// A benefit of type `articles`.
//
// Use it to grant access to posts.
type BenefitUpdateResponse struct {
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The ID of the benefit.
	ID   string                    `json:"id,required" format:"uuid4"`
	Type BenefitUpdateResponseType `json:"type,required"`
	// The description of the benefit.
	Description string `json:"description,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool `json:"selectable,required"`
	// Whether the benefit is deletable.
	Deletable bool `json:"deletable,required"`
	// The ID of the organization owning the benefit.
	OrganizationID string `json:"organization_id,required" format:"uuid4"`
	// This field can have the runtime type of
	// [BenefitUpdateResponseBenefitArticlesProperties],
	// [BenefitUpdateResponseBenefitAdsProperties],
	// [BenefitUpdateResponseBenefitCustomProperties],
	// [BenefitUpdateResponseBenefitDiscordOutputProperties],
	// [BenefitUpdateResponseBenefitGitHubRepositoryProperties],
	// [BenefitUpdateResponseBenefitDownloadablesProperties].
	Properties interface{} `json:"properties"`
	// Whether the benefit is taxable.
	IsTaxApplicable bool                      `json:"is_tax_applicable"`
	JSON            benefitUpdateResponseJSON `json:"-"`
	union           BenefitUpdateResponseUnion
}

// benefitUpdateResponseJSON contains the JSON metadata for the struct
// [BenefitUpdateResponse]
type benefitUpdateResponseJSON struct {
	CreatedAt       apijson.Field
	ModifiedAt      apijson.Field
	ID              apijson.Field
	Type            apijson.Field
	Description     apijson.Field
	Selectable      apijson.Field
	Deletable       apijson.Field
	OrganizationID  apijson.Field
	Properties      apijson.Field
	IsTaxApplicable apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r benefitUpdateResponseJSON) RawJSON() string {
	return r.raw
}

func (r *BenefitUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	*r = BenefitUpdateResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [BenefitUpdateResponseUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [BenefitUpdateResponseBenefitArticles],
// [BenefitUpdateResponseBenefitAds], [BenefitUpdateResponseBenefitCustom],
// [BenefitUpdateResponseBenefitDiscordOutput],
// [BenefitUpdateResponseBenefitGitHubRepository],
// [BenefitUpdateResponseBenefitDownloadables].
func (r BenefitUpdateResponse) AsUnion() BenefitUpdateResponseUnion {
	return r.union
}

// A benefit of type `articles`.
//
// Use it to grant access to posts.
//
// Union satisfied by [BenefitUpdateResponseBenefitArticles],
// [BenefitUpdateResponseBenefitAds], [BenefitUpdateResponseBenefitCustom],
// [BenefitUpdateResponseBenefitDiscordOutput],
// [BenefitUpdateResponseBenefitGitHubRepository] or
// [BenefitUpdateResponseBenefitDownloadables].
type BenefitUpdateResponseUnion interface {
	implementsBenefitUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BenefitUpdateResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BenefitUpdateResponseBenefitArticles{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BenefitUpdateResponseBenefitAds{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BenefitUpdateResponseBenefitCustom{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BenefitUpdateResponseBenefitDiscordOutput{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BenefitUpdateResponseBenefitGitHubRepository{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BenefitUpdateResponseBenefitDownloadables{}),
		},
	)
}

// A benefit of type `articles`.
//
// Use it to grant access to posts.
type BenefitUpdateResponseBenefitArticles struct {
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
	Properties BenefitUpdateResponseBenefitArticlesProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                     `json:"selectable,required"`
	Type       BenefitUpdateResponseBenefitArticlesType `json:"type,required"`
	JSON       benefitUpdateResponseBenefitArticlesJSON `json:"-"`
}

// benefitUpdateResponseBenefitArticlesJSON contains the JSON metadata for the
// struct [BenefitUpdateResponseBenefitArticles]
type benefitUpdateResponseBenefitArticlesJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Deletable      apijson.Field
	Description    apijson.Field
	ModifiedAt     apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	Selectable     apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BenefitUpdateResponseBenefitArticles) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitUpdateResponseBenefitArticlesJSON) RawJSON() string {
	return r.raw
}

func (r BenefitUpdateResponseBenefitArticles) implementsBenefitUpdateResponse() {}

// Properties for a benefit of type `articles`.
type BenefitUpdateResponseBenefitArticlesProperties struct {
	// Whether the user can access paid articles.
	PaidArticles bool                                               `json:"paid_articles,required"`
	JSON         benefitUpdateResponseBenefitArticlesPropertiesJSON `json:"-"`
}

// benefitUpdateResponseBenefitArticlesPropertiesJSON contains the JSON metadata
// for the struct [BenefitUpdateResponseBenefitArticlesProperties]
type benefitUpdateResponseBenefitArticlesPropertiesJSON struct {
	PaidArticles apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *BenefitUpdateResponseBenefitArticlesProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitUpdateResponseBenefitArticlesPropertiesJSON) RawJSON() string {
	return r.raw
}

type BenefitUpdateResponseBenefitArticlesType string

const (
	BenefitUpdateResponseBenefitArticlesTypeArticles BenefitUpdateResponseBenefitArticlesType = "articles"
)

func (r BenefitUpdateResponseBenefitArticlesType) IsKnown() bool {
	switch r {
	case BenefitUpdateResponseBenefitArticlesTypeArticles:
		return true
	}
	return false
}

// A benefit of type `ads`.
//
// Use it so your backers can display ads on your README, website, etc.
type BenefitUpdateResponseBenefitAds struct {
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
	// Properties for a benefit of type `ads`.
	Properties BenefitUpdateResponseBenefitAdsProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                `json:"selectable,required"`
	Type       BenefitUpdateResponseBenefitAdsType `json:"type,required"`
	JSON       benefitUpdateResponseBenefitAdsJSON `json:"-"`
}

// benefitUpdateResponseBenefitAdsJSON contains the JSON metadata for the struct
// [BenefitUpdateResponseBenefitAds]
type benefitUpdateResponseBenefitAdsJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Deletable      apijson.Field
	Description    apijson.Field
	ModifiedAt     apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	Selectable     apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BenefitUpdateResponseBenefitAds) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitUpdateResponseBenefitAdsJSON) RawJSON() string {
	return r.raw
}

func (r BenefitUpdateResponseBenefitAds) implementsBenefitUpdateResponse() {}

// Properties for a benefit of type `ads`.
type BenefitUpdateResponseBenefitAdsProperties struct {
	// The height of the displayed ad.
	ImageHeight int64 `json:"image_height"`
	// The width of the displayed ad.
	ImageWidth int64                                         `json:"image_width"`
	JSON       benefitUpdateResponseBenefitAdsPropertiesJSON `json:"-"`
}

// benefitUpdateResponseBenefitAdsPropertiesJSON contains the JSON metadata for the
// struct [BenefitUpdateResponseBenefitAdsProperties]
type benefitUpdateResponseBenefitAdsPropertiesJSON struct {
	ImageHeight apijson.Field
	ImageWidth  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BenefitUpdateResponseBenefitAdsProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitUpdateResponseBenefitAdsPropertiesJSON) RawJSON() string {
	return r.raw
}

type BenefitUpdateResponseBenefitAdsType string

const (
	BenefitUpdateResponseBenefitAdsTypeAds BenefitUpdateResponseBenefitAdsType = "ads"
)

func (r BenefitUpdateResponseBenefitAdsType) IsKnown() bool {
	switch r {
	case BenefitUpdateResponseBenefitAdsTypeAds:
		return true
	}
	return false
}

// A benefit of type `custom`.
//
// Use it to grant any kind of benefit that doesn't fit in the other types.
type BenefitUpdateResponseBenefitCustom struct {
	// The ID of the benefit.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the benefit is deletable.
	Deletable bool `json:"deletable,required"`
	// The description of the benefit.
	Description string `json:"description,required"`
	// Whether the benefit is taxable.
	IsTaxApplicable bool `json:"is_tax_applicable,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The ID of the organization owning the benefit.
	OrganizationID string `json:"organization_id,required" format:"uuid4"`
	// Properties for a benefit of type `custom`.
	Properties BenefitUpdateResponseBenefitCustomProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                   `json:"selectable,required"`
	Type       BenefitUpdateResponseBenefitCustomType `json:"type,required"`
	JSON       benefitUpdateResponseBenefitCustomJSON `json:"-"`
}

// benefitUpdateResponseBenefitCustomJSON contains the JSON metadata for the struct
// [BenefitUpdateResponseBenefitCustom]
type benefitUpdateResponseBenefitCustomJSON struct {
	ID              apijson.Field
	CreatedAt       apijson.Field
	Deletable       apijson.Field
	Description     apijson.Field
	IsTaxApplicable apijson.Field
	ModifiedAt      apijson.Field
	OrganizationID  apijson.Field
	Properties      apijson.Field
	Selectable      apijson.Field
	Type            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *BenefitUpdateResponseBenefitCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitUpdateResponseBenefitCustomJSON) RawJSON() string {
	return r.raw
}

func (r BenefitUpdateResponseBenefitCustom) implementsBenefitUpdateResponse() {}

// Properties for a benefit of type `custom`.
type BenefitUpdateResponseBenefitCustomProperties struct {
	// Private note to be shared with users who have this benefit granted.
	Note string                                           `json:"note,required,nullable"`
	JSON benefitUpdateResponseBenefitCustomPropertiesJSON `json:"-"`
}

// benefitUpdateResponseBenefitCustomPropertiesJSON contains the JSON metadata for
// the struct [BenefitUpdateResponseBenefitCustomProperties]
type benefitUpdateResponseBenefitCustomPropertiesJSON struct {
	Note        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BenefitUpdateResponseBenefitCustomProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitUpdateResponseBenefitCustomPropertiesJSON) RawJSON() string {
	return r.raw
}

type BenefitUpdateResponseBenefitCustomType string

const (
	BenefitUpdateResponseBenefitCustomTypeCustom BenefitUpdateResponseBenefitCustomType = "custom"
)

func (r BenefitUpdateResponseBenefitCustomType) IsKnown() bool {
	switch r {
	case BenefitUpdateResponseBenefitCustomTypeCustom:
		return true
	}
	return false
}

// A benefit of type `discord`.
//
// Use it to automatically invite your backers to a Discord server.
type BenefitUpdateResponseBenefitDiscordOutput struct {
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
	// Properties for a benefit of type `discord`.
	Properties BenefitUpdateResponseBenefitDiscordOutputProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                          `json:"selectable,required"`
	Type       BenefitUpdateResponseBenefitDiscordOutputType `json:"type,required"`
	JSON       benefitUpdateResponseBenefitDiscordOutputJSON `json:"-"`
}

// benefitUpdateResponseBenefitDiscordOutputJSON contains the JSON metadata for the
// struct [BenefitUpdateResponseBenefitDiscordOutput]
type benefitUpdateResponseBenefitDiscordOutputJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Deletable      apijson.Field
	Description    apijson.Field
	ModifiedAt     apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	Selectable     apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BenefitUpdateResponseBenefitDiscordOutput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitUpdateResponseBenefitDiscordOutputJSON) RawJSON() string {
	return r.raw
}

func (r BenefitUpdateResponseBenefitDiscordOutput) implementsBenefitUpdateResponse() {}

// Properties for a benefit of type `discord`.
type BenefitUpdateResponseBenefitDiscordOutputProperties struct {
	// The ID of the Discord server.
	GuildID    string `json:"guild_id,required"`
	GuildToken string `json:"guild_token,required"`
	// The ID of the Discord role to grant.
	RoleID string                                                  `json:"role_id,required"`
	JSON   benefitUpdateResponseBenefitDiscordOutputPropertiesJSON `json:"-"`
}

// benefitUpdateResponseBenefitDiscordOutputPropertiesJSON contains the JSON
// metadata for the struct [BenefitUpdateResponseBenefitDiscordOutputProperties]
type benefitUpdateResponseBenefitDiscordOutputPropertiesJSON struct {
	GuildID     apijson.Field
	GuildToken  apijson.Field
	RoleID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BenefitUpdateResponseBenefitDiscordOutputProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitUpdateResponseBenefitDiscordOutputPropertiesJSON) RawJSON() string {
	return r.raw
}

type BenefitUpdateResponseBenefitDiscordOutputType string

const (
	BenefitUpdateResponseBenefitDiscordOutputTypeDiscord BenefitUpdateResponseBenefitDiscordOutputType = "discord"
)

func (r BenefitUpdateResponseBenefitDiscordOutputType) IsKnown() bool {
	switch r {
	case BenefitUpdateResponseBenefitDiscordOutputTypeDiscord:
		return true
	}
	return false
}

// A benefit of type `github_repository`.
//
// Use it to automatically invite your backers to a private GitHub repository.
type BenefitUpdateResponseBenefitGitHubRepository struct {
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
	// Properties for a benefit of type `github_repository`.
	Properties BenefitUpdateResponseBenefitGitHubRepositoryProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                             `json:"selectable,required"`
	Type       BenefitUpdateResponseBenefitGitHubRepositoryType `json:"type,required"`
	JSON       benefitUpdateResponseBenefitGitHubRepositoryJSON `json:"-"`
}

// benefitUpdateResponseBenefitGitHubRepositoryJSON contains the JSON metadata for
// the struct [BenefitUpdateResponseBenefitGitHubRepository]
type benefitUpdateResponseBenefitGitHubRepositoryJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Deletable      apijson.Field
	Description    apijson.Field
	ModifiedAt     apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	Selectable     apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BenefitUpdateResponseBenefitGitHubRepository) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitUpdateResponseBenefitGitHubRepositoryJSON) RawJSON() string {
	return r.raw
}

func (r BenefitUpdateResponseBenefitGitHubRepository) implementsBenefitUpdateResponse() {}

// Properties for a benefit of type `github_repository`.
type BenefitUpdateResponseBenefitGitHubRepositoryProperties struct {
	// The permission level to grant. Read more about roles and their permissions on
	// [GitHub documentation](https://docs.github.com/en/organizations/managing-user-access-to-your-organizations-repositories/managing-repository-roles/repository-roles-for-an-organization#permissions-for-each-role).
	Permission   BenefitUpdateResponseBenefitGitHubRepositoryPropertiesPermission `json:"permission,required"`
	RepositoryID string                                                           `json:"repository_id,required,nullable" format:"uuid4"`
	// The name of the repository.
	RepositoryName string `json:"repository_name,required"`
	// The owner of the repository.
	RepositoryOwner string                                                     `json:"repository_owner,required"`
	JSON            benefitUpdateResponseBenefitGitHubRepositoryPropertiesJSON `json:"-"`
}

// benefitUpdateResponseBenefitGitHubRepositoryPropertiesJSON contains the JSON
// metadata for the struct [BenefitUpdateResponseBenefitGitHubRepositoryProperties]
type benefitUpdateResponseBenefitGitHubRepositoryPropertiesJSON struct {
	Permission      apijson.Field
	RepositoryID    apijson.Field
	RepositoryName  apijson.Field
	RepositoryOwner apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *BenefitUpdateResponseBenefitGitHubRepositoryProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitUpdateResponseBenefitGitHubRepositoryPropertiesJSON) RawJSON() string {
	return r.raw
}

// The permission level to grant. Read more about roles and their permissions on
// [GitHub documentation](https://docs.github.com/en/organizations/managing-user-access-to-your-organizations-repositories/managing-repository-roles/repository-roles-for-an-organization#permissions-for-each-role).
type BenefitUpdateResponseBenefitGitHubRepositoryPropertiesPermission string

const (
	BenefitUpdateResponseBenefitGitHubRepositoryPropertiesPermissionPull     BenefitUpdateResponseBenefitGitHubRepositoryPropertiesPermission = "pull"
	BenefitUpdateResponseBenefitGitHubRepositoryPropertiesPermissionTriage   BenefitUpdateResponseBenefitGitHubRepositoryPropertiesPermission = "triage"
	BenefitUpdateResponseBenefitGitHubRepositoryPropertiesPermissionPush     BenefitUpdateResponseBenefitGitHubRepositoryPropertiesPermission = "push"
	BenefitUpdateResponseBenefitGitHubRepositoryPropertiesPermissionMaintain BenefitUpdateResponseBenefitGitHubRepositoryPropertiesPermission = "maintain"
	BenefitUpdateResponseBenefitGitHubRepositoryPropertiesPermissionAdmin    BenefitUpdateResponseBenefitGitHubRepositoryPropertiesPermission = "admin"
)

func (r BenefitUpdateResponseBenefitGitHubRepositoryPropertiesPermission) IsKnown() bool {
	switch r {
	case BenefitUpdateResponseBenefitGitHubRepositoryPropertiesPermissionPull, BenefitUpdateResponseBenefitGitHubRepositoryPropertiesPermissionTriage, BenefitUpdateResponseBenefitGitHubRepositoryPropertiesPermissionPush, BenefitUpdateResponseBenefitGitHubRepositoryPropertiesPermissionMaintain, BenefitUpdateResponseBenefitGitHubRepositoryPropertiesPermissionAdmin:
		return true
	}
	return false
}

type BenefitUpdateResponseBenefitGitHubRepositoryType string

const (
	BenefitUpdateResponseBenefitGitHubRepositoryTypeGitHubRepository BenefitUpdateResponseBenefitGitHubRepositoryType = "github_repository"
)

func (r BenefitUpdateResponseBenefitGitHubRepositoryType) IsKnown() bool {
	switch r {
	case BenefitUpdateResponseBenefitGitHubRepositoryTypeGitHubRepository:
		return true
	}
	return false
}

type BenefitUpdateResponseBenefitDownloadables struct {
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
	OrganizationID string                                              `json:"organization_id,required" format:"uuid4"`
	Properties     BenefitUpdateResponseBenefitDownloadablesProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                          `json:"selectable,required"`
	Type       BenefitUpdateResponseBenefitDownloadablesType `json:"type,required"`
	JSON       benefitUpdateResponseBenefitDownloadablesJSON `json:"-"`
}

// benefitUpdateResponseBenefitDownloadablesJSON contains the JSON metadata for the
// struct [BenefitUpdateResponseBenefitDownloadables]
type benefitUpdateResponseBenefitDownloadablesJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Deletable      apijson.Field
	Description    apijson.Field
	ModifiedAt     apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	Selectable     apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BenefitUpdateResponseBenefitDownloadables) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitUpdateResponseBenefitDownloadablesJSON) RawJSON() string {
	return r.raw
}

func (r BenefitUpdateResponseBenefitDownloadables) implementsBenefitUpdateResponse() {}

type BenefitUpdateResponseBenefitDownloadablesProperties struct {
	Archived map[string]bool                                         `json:"archived,required"`
	Files    []string                                                `json:"files,required" format:"uuid4"`
	JSON     benefitUpdateResponseBenefitDownloadablesPropertiesJSON `json:"-"`
}

// benefitUpdateResponseBenefitDownloadablesPropertiesJSON contains the JSON
// metadata for the struct [BenefitUpdateResponseBenefitDownloadablesProperties]
type benefitUpdateResponseBenefitDownloadablesPropertiesJSON struct {
	Archived    apijson.Field
	Files       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BenefitUpdateResponseBenefitDownloadablesProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitUpdateResponseBenefitDownloadablesPropertiesJSON) RawJSON() string {
	return r.raw
}

type BenefitUpdateResponseBenefitDownloadablesType string

const (
	BenefitUpdateResponseBenefitDownloadablesTypeDownloadables BenefitUpdateResponseBenefitDownloadablesType = "downloadables"
)

func (r BenefitUpdateResponseBenefitDownloadablesType) IsKnown() bool {
	switch r {
	case BenefitUpdateResponseBenefitDownloadablesTypeDownloadables:
		return true
	}
	return false
}

type BenefitUpdateResponseType string

const (
	BenefitUpdateResponseTypeArticles         BenefitUpdateResponseType = "articles"
	BenefitUpdateResponseTypeAds              BenefitUpdateResponseType = "ads"
	BenefitUpdateResponseTypeCustom           BenefitUpdateResponseType = "custom"
	BenefitUpdateResponseTypeDiscord          BenefitUpdateResponseType = "discord"
	BenefitUpdateResponseTypeGitHubRepository BenefitUpdateResponseType = "github_repository"
	BenefitUpdateResponseTypeDownloadables    BenefitUpdateResponseType = "downloadables"
)

func (r BenefitUpdateResponseType) IsKnown() bool {
	switch r {
	case BenefitUpdateResponseTypeArticles, BenefitUpdateResponseTypeAds, BenefitUpdateResponseTypeCustom, BenefitUpdateResponseTypeDiscord, BenefitUpdateResponseTypeGitHubRepository, BenefitUpdateResponseTypeDownloadables:
		return true
	}
	return false
}

type BenefitNewParams struct {
	// Schema to create a benefit of type `custom`.
	Body BenefitNewParamsBodyUnion `json:"body,required"`
}

func (r BenefitNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// Schema to create a benefit of type `custom`.
type BenefitNewParamsBody struct {
	Type param.Field[BenefitNewParamsBodyType] `json:"type,required"`
	// The description of the benefit. Will be displayed on products having this
	// benefit.
	Description param.Field[string] `json:"description,required"`
	// The organization ID.
	OrganizationID param.Field[string] `json:"organization_id" format:"uuid4"`
	// Whether the benefit is taxable.
	IsTaxApplicable param.Field[bool]        `json:"is_tax_applicable"`
	Properties      param.Field[interface{}] `json:"properties"`
}

func (r BenefitNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BenefitNewParamsBody) implementsBenefitNewParamsBodyUnion() {}

// Schema to create a benefit of type `custom`.
//
// Satisfied by [BenefitNewParamsBodyBenefitCustomCreate],
// [BenefitNewParamsBodyBenefitAdsCreate],
// [BenefitNewParamsBodyBenefitDiscordCreate],
// [BenefitNewParamsBodyBenefitGitHubRepositoryCreate],
// [BenefitNewParamsBodyBenefitDownloadablesCreate], [BenefitNewParamsBody].
type BenefitNewParamsBodyUnion interface {
	implementsBenefitNewParamsBodyUnion()
}

// Schema to create a benefit of type `custom`.
type BenefitNewParamsBodyBenefitCustomCreate struct {
	// The description of the benefit. Will be displayed on products having this
	// benefit.
	Description param.Field[string] `json:"description,required"`
	// Whether the benefit is taxable.
	IsTaxApplicable param.Field[bool] `json:"is_tax_applicable,required"`
	// Properties for a benefit of type `custom`.
	Properties param.Field[BenefitNewParamsBodyBenefitCustomCreateProperties] `json:"properties,required"`
	Type       param.Field[BenefitNewParamsBodyBenefitCustomCreateType]       `json:"type,required"`
	// The organization ID.
	OrganizationID param.Field[string] `json:"organization_id" format:"uuid4"`
}

func (r BenefitNewParamsBodyBenefitCustomCreate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BenefitNewParamsBodyBenefitCustomCreate) implementsBenefitNewParamsBodyUnion() {}

// Properties for a benefit of type `custom`.
type BenefitNewParamsBodyBenefitCustomCreateProperties struct {
	// Private note to be shared with users who have this benefit granted.
	Note param.Field[string] `json:"note,required"`
}

func (r BenefitNewParamsBodyBenefitCustomCreateProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BenefitNewParamsBodyBenefitCustomCreateType string

const (
	BenefitNewParamsBodyBenefitCustomCreateTypeCustom BenefitNewParamsBodyBenefitCustomCreateType = "custom"
)

func (r BenefitNewParamsBodyBenefitCustomCreateType) IsKnown() bool {
	switch r {
	case BenefitNewParamsBodyBenefitCustomCreateTypeCustom:
		return true
	}
	return false
}

type BenefitNewParamsBodyBenefitAdsCreate struct {
	// The description of the benefit. Will be displayed on products having this
	// benefit.
	Description param.Field[string] `json:"description,required"`
	// Properties for a benefit of type `ads`.
	Properties param.Field[BenefitNewParamsBodyBenefitAdsCreateProperties] `json:"properties,required"`
	Type       param.Field[BenefitNewParamsBodyBenefitAdsCreateType]       `json:"type,required"`
	// The organization ID.
	OrganizationID param.Field[string] `json:"organization_id" format:"uuid4"`
}

func (r BenefitNewParamsBodyBenefitAdsCreate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BenefitNewParamsBodyBenefitAdsCreate) implementsBenefitNewParamsBodyUnion() {}

// Properties for a benefit of type `ads`.
type BenefitNewParamsBodyBenefitAdsCreateProperties struct {
	// The height of the displayed ad.
	ImageHeight param.Field[int64] `json:"image_height"`
	// The width of the displayed ad.
	ImageWidth param.Field[int64] `json:"image_width"`
}

func (r BenefitNewParamsBodyBenefitAdsCreateProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BenefitNewParamsBodyBenefitAdsCreateType string

const (
	BenefitNewParamsBodyBenefitAdsCreateTypeAds BenefitNewParamsBodyBenefitAdsCreateType = "ads"
)

func (r BenefitNewParamsBodyBenefitAdsCreateType) IsKnown() bool {
	switch r {
	case BenefitNewParamsBodyBenefitAdsCreateTypeAds:
		return true
	}
	return false
}

type BenefitNewParamsBodyBenefitDiscordCreate struct {
	// The description of the benefit. Will be displayed on products having this
	// benefit.
	Description param.Field[string] `json:"description,required"`
	// Properties to create a benefit of type `discord`.
	Properties param.Field[BenefitNewParamsBodyBenefitDiscordCreateProperties] `json:"properties,required"`
	Type       param.Field[BenefitNewParamsBodyBenefitDiscordCreateType]       `json:"type,required"`
	// The organization ID.
	OrganizationID param.Field[string] `json:"organization_id" format:"uuid4"`
}

func (r BenefitNewParamsBodyBenefitDiscordCreate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BenefitNewParamsBodyBenefitDiscordCreate) implementsBenefitNewParamsBodyUnion() {}

// Properties to create a benefit of type `discord`.
type BenefitNewParamsBodyBenefitDiscordCreateProperties struct {
	GuildToken param.Field[string] `json:"guild_token,required"`
	// The ID of the Discord role to grant.
	RoleID param.Field[string] `json:"role_id,required"`
}

func (r BenefitNewParamsBodyBenefitDiscordCreateProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BenefitNewParamsBodyBenefitDiscordCreateType string

const (
	BenefitNewParamsBodyBenefitDiscordCreateTypeDiscord BenefitNewParamsBodyBenefitDiscordCreateType = "discord"
)

func (r BenefitNewParamsBodyBenefitDiscordCreateType) IsKnown() bool {
	switch r {
	case BenefitNewParamsBodyBenefitDiscordCreateTypeDiscord:
		return true
	}
	return false
}

type BenefitNewParamsBodyBenefitGitHubRepositoryCreate struct {
	// The description of the benefit. Will be displayed on products having this
	// benefit.
	Description param.Field[string] `json:"description,required"`
	// Properties to create a benefit of type `github_repository`.
	Properties param.Field[BenefitNewParamsBodyBenefitGitHubRepositoryCreateProperties] `json:"properties,required"`
	Type       param.Field[BenefitNewParamsBodyBenefitGitHubRepositoryCreateType]       `json:"type,required"`
	// The organization ID.
	OrganizationID param.Field[string] `json:"organization_id" format:"uuid4"`
}

func (r BenefitNewParamsBodyBenefitGitHubRepositoryCreate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BenefitNewParamsBodyBenefitGitHubRepositoryCreate) implementsBenefitNewParamsBodyUnion() {}

// Properties to create a benefit of type `github_repository`.
type BenefitNewParamsBodyBenefitGitHubRepositoryCreateProperties struct {
	// The permission level to grant. Read more about roles and their permissions on
	// [GitHub documentation](https://docs.github.com/en/organizations/managing-user-access-to-your-organizations-repositories/managing-repository-roles/repository-roles-for-an-organization#permissions-for-each-role).
	Permission   param.Field[BenefitNewParamsBodyBenefitGitHubRepositoryCreatePropertiesPermission] `json:"permission,required"`
	RepositoryID param.Field[string]                                                                `json:"repository_id" format:"uuid4"`
	// The name of the repository.
	RepositoryName param.Field[string] `json:"repository_name"`
	// The owner of the repository.
	RepositoryOwner param.Field[string] `json:"repository_owner"`
}

func (r BenefitNewParamsBodyBenefitGitHubRepositoryCreateProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The permission level to grant. Read more about roles and their permissions on
// [GitHub documentation](https://docs.github.com/en/organizations/managing-user-access-to-your-organizations-repositories/managing-repository-roles/repository-roles-for-an-organization#permissions-for-each-role).
type BenefitNewParamsBodyBenefitGitHubRepositoryCreatePropertiesPermission string

const (
	BenefitNewParamsBodyBenefitGitHubRepositoryCreatePropertiesPermissionPull     BenefitNewParamsBodyBenefitGitHubRepositoryCreatePropertiesPermission = "pull"
	BenefitNewParamsBodyBenefitGitHubRepositoryCreatePropertiesPermissionTriage   BenefitNewParamsBodyBenefitGitHubRepositoryCreatePropertiesPermission = "triage"
	BenefitNewParamsBodyBenefitGitHubRepositoryCreatePropertiesPermissionPush     BenefitNewParamsBodyBenefitGitHubRepositoryCreatePropertiesPermission = "push"
	BenefitNewParamsBodyBenefitGitHubRepositoryCreatePropertiesPermissionMaintain BenefitNewParamsBodyBenefitGitHubRepositoryCreatePropertiesPermission = "maintain"
	BenefitNewParamsBodyBenefitGitHubRepositoryCreatePropertiesPermissionAdmin    BenefitNewParamsBodyBenefitGitHubRepositoryCreatePropertiesPermission = "admin"
)

func (r BenefitNewParamsBodyBenefitGitHubRepositoryCreatePropertiesPermission) IsKnown() bool {
	switch r {
	case BenefitNewParamsBodyBenefitGitHubRepositoryCreatePropertiesPermissionPull, BenefitNewParamsBodyBenefitGitHubRepositoryCreatePropertiesPermissionTriage, BenefitNewParamsBodyBenefitGitHubRepositoryCreatePropertiesPermissionPush, BenefitNewParamsBodyBenefitGitHubRepositoryCreatePropertiesPermissionMaintain, BenefitNewParamsBodyBenefitGitHubRepositoryCreatePropertiesPermissionAdmin:
		return true
	}
	return false
}

type BenefitNewParamsBodyBenefitGitHubRepositoryCreateType string

const (
	BenefitNewParamsBodyBenefitGitHubRepositoryCreateTypeGitHubRepository BenefitNewParamsBodyBenefitGitHubRepositoryCreateType = "github_repository"
)

func (r BenefitNewParamsBodyBenefitGitHubRepositoryCreateType) IsKnown() bool {
	switch r {
	case BenefitNewParamsBodyBenefitGitHubRepositoryCreateTypeGitHubRepository:
		return true
	}
	return false
}

type BenefitNewParamsBodyBenefitDownloadablesCreate struct {
	// The description of the benefit. Will be displayed on products having this
	// benefit.
	Description param.Field[string]                                                   `json:"description,required"`
	Properties  param.Field[BenefitNewParamsBodyBenefitDownloadablesCreateProperties] `json:"properties,required"`
	Type        param.Field[BenefitNewParamsBodyBenefitDownloadablesCreateType]       `json:"type,required"`
	// The organization ID.
	OrganizationID param.Field[string] `json:"organization_id" format:"uuid4"`
}

func (r BenefitNewParamsBodyBenefitDownloadablesCreate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BenefitNewParamsBodyBenefitDownloadablesCreate) implementsBenefitNewParamsBodyUnion() {}

type BenefitNewParamsBodyBenefitDownloadablesCreateProperties struct {
	Files    param.Field[[]string]        `json:"files,required" format:"uuid4"`
	Archived param.Field[map[string]bool] `json:"archived"`
}

func (r BenefitNewParamsBodyBenefitDownloadablesCreateProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BenefitNewParamsBodyBenefitDownloadablesCreateType string

const (
	BenefitNewParamsBodyBenefitDownloadablesCreateTypeDownloadables BenefitNewParamsBodyBenefitDownloadablesCreateType = "downloadables"
)

func (r BenefitNewParamsBodyBenefitDownloadablesCreateType) IsKnown() bool {
	switch r {
	case BenefitNewParamsBodyBenefitDownloadablesCreateTypeDownloadables:
		return true
	}
	return false
}

type BenefitNewParamsBodyType string

const (
	BenefitNewParamsBodyTypeCustom           BenefitNewParamsBodyType = "custom"
	BenefitNewParamsBodyTypeAds              BenefitNewParamsBodyType = "ads"
	BenefitNewParamsBodyTypeDiscord          BenefitNewParamsBodyType = "discord"
	BenefitNewParamsBodyTypeGitHubRepository BenefitNewParamsBodyType = "github_repository"
	BenefitNewParamsBodyTypeDownloadables    BenefitNewParamsBodyType = "downloadables"
)

func (r BenefitNewParamsBodyType) IsKnown() bool {
	switch r {
	case BenefitNewParamsBodyTypeCustom, BenefitNewParamsBodyTypeAds, BenefitNewParamsBodyTypeDiscord, BenefitNewParamsBodyTypeGitHubRepository, BenefitNewParamsBodyTypeDownloadables:
		return true
	}
	return false
}

type BenefitUpdateParams struct {
	Body BenefitUpdateParamsBodyUnion `json:"body,required"`
}

func (r BenefitUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type BenefitUpdateParamsBody struct {
	// The description of the benefit. Will be displayed on products having this
	// benefit.
	Description param.Field[string]                      `json:"description"`
	Type        param.Field[BenefitUpdateParamsBodyType] `json:"type,required"`
	Properties  param.Field[interface{}]                 `json:"properties,required"`
}

func (r BenefitUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BenefitUpdateParamsBody) implementsBenefitUpdateParamsBodyUnion() {}

// Satisfied by [BenefitUpdateParamsBodyBenefitArticlesUpdate],
// [BenefitUpdateParamsBodyBenefitAdsUpdate],
// [BenefitUpdateParamsBodyBenefitCustomUpdate],
// [BenefitUpdateParamsBodyBenefitDiscordUpdate],
// [BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdate],
// [BenefitUpdateParamsBodyBenefitDownloadablesUpdate], [BenefitUpdateParamsBody].
type BenefitUpdateParamsBodyUnion interface {
	implementsBenefitUpdateParamsBodyUnion()
}

type BenefitUpdateParamsBodyBenefitArticlesUpdate struct {
	Type param.Field[BenefitUpdateParamsBodyBenefitArticlesUpdateType] `json:"type,required"`
	// The description of the benefit. Will be displayed on products having this
	// benefit.
	Description param.Field[string] `json:"description"`
}

func (r BenefitUpdateParamsBodyBenefitArticlesUpdate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BenefitUpdateParamsBodyBenefitArticlesUpdate) implementsBenefitUpdateParamsBodyUnion() {}

type BenefitUpdateParamsBodyBenefitArticlesUpdateType string

const (
	BenefitUpdateParamsBodyBenefitArticlesUpdateTypeArticles BenefitUpdateParamsBodyBenefitArticlesUpdateType = "articles"
)

func (r BenefitUpdateParamsBodyBenefitArticlesUpdateType) IsKnown() bool {
	switch r {
	case BenefitUpdateParamsBodyBenefitArticlesUpdateTypeArticles:
		return true
	}
	return false
}

type BenefitUpdateParamsBodyBenefitAdsUpdate struct {
	Type param.Field[BenefitUpdateParamsBodyBenefitAdsUpdateType] `json:"type,required"`
	// The description of the benefit. Will be displayed on products having this
	// benefit.
	Description param.Field[string] `json:"description"`
	// Properties for a benefit of type `ads`.
	Properties param.Field[BenefitUpdateParamsBodyBenefitAdsUpdateProperties] `json:"properties"`
}

func (r BenefitUpdateParamsBodyBenefitAdsUpdate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BenefitUpdateParamsBodyBenefitAdsUpdate) implementsBenefitUpdateParamsBodyUnion() {}

type BenefitUpdateParamsBodyBenefitAdsUpdateType string

const (
	BenefitUpdateParamsBodyBenefitAdsUpdateTypeAds BenefitUpdateParamsBodyBenefitAdsUpdateType = "ads"
)

func (r BenefitUpdateParamsBodyBenefitAdsUpdateType) IsKnown() bool {
	switch r {
	case BenefitUpdateParamsBodyBenefitAdsUpdateTypeAds:
		return true
	}
	return false
}

// Properties for a benefit of type `ads`.
type BenefitUpdateParamsBodyBenefitAdsUpdateProperties struct {
	// The height of the displayed ad.
	ImageHeight param.Field[int64] `json:"image_height"`
	// The width of the displayed ad.
	ImageWidth param.Field[int64] `json:"image_width"`
}

func (r BenefitUpdateParamsBodyBenefitAdsUpdateProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BenefitUpdateParamsBodyBenefitCustomUpdate struct {
	Type param.Field[BenefitUpdateParamsBodyBenefitCustomUpdateType] `json:"type,required"`
	// The description of the benefit. Will be displayed on products having this
	// benefit.
	Description param.Field[string] `json:"description"`
	// Properties for a benefit of type `custom`.
	Properties param.Field[BenefitUpdateParamsBodyBenefitCustomUpdateProperties] `json:"properties"`
}

func (r BenefitUpdateParamsBodyBenefitCustomUpdate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BenefitUpdateParamsBodyBenefitCustomUpdate) implementsBenefitUpdateParamsBodyUnion() {}

type BenefitUpdateParamsBodyBenefitCustomUpdateType string

const (
	BenefitUpdateParamsBodyBenefitCustomUpdateTypeCustom BenefitUpdateParamsBodyBenefitCustomUpdateType = "custom"
)

func (r BenefitUpdateParamsBodyBenefitCustomUpdateType) IsKnown() bool {
	switch r {
	case BenefitUpdateParamsBodyBenefitCustomUpdateTypeCustom:
		return true
	}
	return false
}

// Properties for a benefit of type `custom`.
type BenefitUpdateParamsBodyBenefitCustomUpdateProperties struct {
	// Private note to be shared with users who have this benefit granted.
	Note param.Field[string] `json:"note,required"`
}

func (r BenefitUpdateParamsBodyBenefitCustomUpdateProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BenefitUpdateParamsBodyBenefitDiscordUpdate struct {
	Type param.Field[BenefitUpdateParamsBodyBenefitDiscordUpdateType] `json:"type,required"`
	// The description of the benefit. Will be displayed on products having this
	// benefit.
	Description param.Field[string] `json:"description"`
	// Properties to create a benefit of type `discord`.
	Properties param.Field[BenefitUpdateParamsBodyBenefitDiscordUpdateProperties] `json:"properties"`
}

func (r BenefitUpdateParamsBodyBenefitDiscordUpdate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BenefitUpdateParamsBodyBenefitDiscordUpdate) implementsBenefitUpdateParamsBodyUnion() {}

type BenefitUpdateParamsBodyBenefitDiscordUpdateType string

const (
	BenefitUpdateParamsBodyBenefitDiscordUpdateTypeDiscord BenefitUpdateParamsBodyBenefitDiscordUpdateType = "discord"
)

func (r BenefitUpdateParamsBodyBenefitDiscordUpdateType) IsKnown() bool {
	switch r {
	case BenefitUpdateParamsBodyBenefitDiscordUpdateTypeDiscord:
		return true
	}
	return false
}

// Properties to create a benefit of type `discord`.
type BenefitUpdateParamsBodyBenefitDiscordUpdateProperties struct {
	GuildToken param.Field[string] `json:"guild_token,required"`
	// The ID of the Discord role to grant.
	RoleID param.Field[string] `json:"role_id,required"`
}

func (r BenefitUpdateParamsBodyBenefitDiscordUpdateProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdate struct {
	Type param.Field[BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdateType] `json:"type,required"`
	// The description of the benefit. Will be displayed on products having this
	// benefit.
	Description param.Field[string] `json:"description"`
	// Properties to create a benefit of type `github_repository`.
	Properties param.Field[BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdateProperties] `json:"properties"`
}

func (r BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdate) implementsBenefitUpdateParamsBodyUnion() {
}

type BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdateType string

const (
	BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdateTypeGitHubRepository BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdateType = "github_repository"
)

func (r BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdateType) IsKnown() bool {
	switch r {
	case BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdateTypeGitHubRepository:
		return true
	}
	return false
}

// Properties to create a benefit of type `github_repository`.
type BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdateProperties struct {
	// The permission level to grant. Read more about roles and their permissions on
	// [GitHub documentation](https://docs.github.com/en/organizations/managing-user-access-to-your-organizations-repositories/managing-repository-roles/repository-roles-for-an-organization#permissions-for-each-role).
	Permission   param.Field[BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdatePropertiesPermission] `json:"permission,required"`
	RepositoryID param.Field[string]                                                                   `json:"repository_id" format:"uuid4"`
	// The name of the repository.
	RepositoryName param.Field[string] `json:"repository_name"`
	// The owner of the repository.
	RepositoryOwner param.Field[string] `json:"repository_owner"`
}

func (r BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdateProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The permission level to grant. Read more about roles and their permissions on
// [GitHub documentation](https://docs.github.com/en/organizations/managing-user-access-to-your-organizations-repositories/managing-repository-roles/repository-roles-for-an-organization#permissions-for-each-role).
type BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdatePropertiesPermission string

const (
	BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdatePropertiesPermissionPull     BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdatePropertiesPermission = "pull"
	BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdatePropertiesPermissionTriage   BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdatePropertiesPermission = "triage"
	BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdatePropertiesPermissionPush     BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdatePropertiesPermission = "push"
	BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdatePropertiesPermissionMaintain BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdatePropertiesPermission = "maintain"
	BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdatePropertiesPermissionAdmin    BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdatePropertiesPermission = "admin"
)

func (r BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdatePropertiesPermission) IsKnown() bool {
	switch r {
	case BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdatePropertiesPermissionPull, BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdatePropertiesPermissionTriage, BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdatePropertiesPermissionPush, BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdatePropertiesPermissionMaintain, BenefitUpdateParamsBodyBenefitGitHubRepositoryUpdatePropertiesPermissionAdmin:
		return true
	}
	return false
}

type BenefitUpdateParamsBodyBenefitDownloadablesUpdate struct {
	Type param.Field[BenefitUpdateParamsBodyBenefitDownloadablesUpdateType] `json:"type,required"`
	// The description of the benefit. Will be displayed on products having this
	// benefit.
	Description param.Field[string]                                                      `json:"description"`
	Properties  param.Field[BenefitUpdateParamsBodyBenefitDownloadablesUpdateProperties] `json:"properties"`
}

func (r BenefitUpdateParamsBodyBenefitDownloadablesUpdate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BenefitUpdateParamsBodyBenefitDownloadablesUpdate) implementsBenefitUpdateParamsBodyUnion() {}

type BenefitUpdateParamsBodyBenefitDownloadablesUpdateType string

const (
	BenefitUpdateParamsBodyBenefitDownloadablesUpdateTypeDownloadables BenefitUpdateParamsBodyBenefitDownloadablesUpdateType = "downloadables"
)

func (r BenefitUpdateParamsBodyBenefitDownloadablesUpdateType) IsKnown() bool {
	switch r {
	case BenefitUpdateParamsBodyBenefitDownloadablesUpdateTypeDownloadables:
		return true
	}
	return false
}

type BenefitUpdateParamsBodyBenefitDownloadablesUpdateProperties struct {
	Files    param.Field[[]string]        `json:"files,required" format:"uuid4"`
	Archived param.Field[map[string]bool] `json:"archived"`
}

func (r BenefitUpdateParamsBodyBenefitDownloadablesUpdateProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BenefitUpdateParamsBodyType string

const (
	BenefitUpdateParamsBodyTypeArticles         BenefitUpdateParamsBodyType = "articles"
	BenefitUpdateParamsBodyTypeAds              BenefitUpdateParamsBodyType = "ads"
	BenefitUpdateParamsBodyTypeCustom           BenefitUpdateParamsBodyType = "custom"
	BenefitUpdateParamsBodyTypeDiscord          BenefitUpdateParamsBodyType = "discord"
	BenefitUpdateParamsBodyTypeGitHubRepository BenefitUpdateParamsBodyType = "github_repository"
	BenefitUpdateParamsBodyTypeDownloadables    BenefitUpdateParamsBodyType = "downloadables"
)

func (r BenefitUpdateParamsBodyType) IsKnown() bool {
	switch r {
	case BenefitUpdateParamsBodyTypeArticles, BenefitUpdateParamsBodyTypeAds, BenefitUpdateParamsBodyTypeCustom, BenefitUpdateParamsBodyTypeDiscord, BenefitUpdateParamsBodyTypeGitHubRepository, BenefitUpdateParamsBodyTypeDownloadables:
		return true
	}
	return false
}

type BenefitListParams struct {
	// Size of a page, defaults to 10. Maximum is 100.
	Limit param.Field[int64] `query:"limit"`
	// Filter by organization ID.
	OrganizationID param.Field[BenefitListParamsOrganizationIDUnion] `query:"organization_id" format:"uuid4"`
	// Page number, defaults to 1.
	Page param.Field[int64] `query:"page"`
	// Filter by benefit type.
	Type param.Field[BenefitListParamsTypeUnion] `query:"type"`
}

// URLQuery serializes [BenefitListParams]'s query parameters as `url.Values`.
func (r BenefitListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by organization ID.
//
// Satisfied by [shared.UnionString], [BenefitListParamsOrganizationIDArray].
type BenefitListParamsOrganizationIDUnion interface {
	ImplementsBenefitListParamsOrganizationIDUnion()
}

type BenefitListParamsOrganizationIDArray []string

func (r BenefitListParamsOrganizationIDArray) ImplementsBenefitListParamsOrganizationIDUnion() {}

// Filter by benefit type.
//
// Satisfied by [BenefitListParamsTypeBenefitType], [BenefitListParamsTypeArray].
type BenefitListParamsTypeUnion interface {
	implementsBenefitListParamsTypeUnion()
}

type BenefitListParamsTypeBenefitType string

const (
	BenefitListParamsTypeBenefitTypeCustom           BenefitListParamsTypeBenefitType = "custom"
	BenefitListParamsTypeBenefitTypeArticles         BenefitListParamsTypeBenefitType = "articles"
	BenefitListParamsTypeBenefitTypeAds              BenefitListParamsTypeBenefitType = "ads"
	BenefitListParamsTypeBenefitTypeDiscord          BenefitListParamsTypeBenefitType = "discord"
	BenefitListParamsTypeBenefitTypeGitHubRepository BenefitListParamsTypeBenefitType = "github_repository"
	BenefitListParamsTypeBenefitTypeDownloadables    BenefitListParamsTypeBenefitType = "downloadables"
)

func (r BenefitListParamsTypeBenefitType) IsKnown() bool {
	switch r {
	case BenefitListParamsTypeBenefitTypeCustom, BenefitListParamsTypeBenefitTypeArticles, BenefitListParamsTypeBenefitTypeAds, BenefitListParamsTypeBenefitTypeDiscord, BenefitListParamsTypeBenefitTypeGitHubRepository, BenefitListParamsTypeBenefitTypeDownloadables:
		return true
	}
	return false
}

func (r BenefitListParamsTypeBenefitType) implementsBenefitListParamsTypeUnion() {}

type BenefitListParamsTypeArray []BenefitListParamsTypeArray

func (r BenefitListParamsTypeArray) implementsBenefitListParamsTypeUnion() {}
