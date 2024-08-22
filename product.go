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

// ProductService contains methods and other services that help with interacting
// with the polar API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProductService] method instead.
type ProductService struct {
	Options  []option.RequestOption
	Benefits *ProductBenefitService
}

// NewProductService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProductService(opts ...option.RequestOption) (r *ProductService) {
	r = &ProductService{}
	r.Options = opts
	r.Benefits = NewProductBenefitService(opts...)
	return
}

// Create a product.
func (r *ProductService) New(ctx context.Context, body ProductNewParams, opts ...option.RequestOption) (res *ProductOutput, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/products/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a product by ID.
func (r *ProductService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *ProductOutput, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/products/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a product.
func (r *ProductService) Update(ctx context.Context, id string, body ProductUpdateParams, opts ...option.RequestOption) (res *ProductOutput, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/products/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List products.
func (r *ProductService) List(ctx context.Context, query ProductListParams, opts ...option.RequestOption) (res *ListResourceProduct, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/products/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ListResourceProduct struct {
	Pagination ListResourceProductPagination `json:"pagination,required"`
	Items      []ProductOutput               `json:"items"`
	JSON       listResourceProductJSON       `json:"-"`
}

// listResourceProductJSON contains the JSON metadata for the struct
// [ListResourceProduct]
type listResourceProductJSON struct {
	Pagination  apijson.Field
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListResourceProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResourceProductJSON) RawJSON() string {
	return r.raw
}

type ListResourceProductPagination struct {
	MaxPage    int64                             `json:"max_page,required"`
	TotalCount int64                             `json:"total_count,required"`
	JSON       listResourceProductPaginationJSON `json:"-"`
}

// listResourceProductPaginationJSON contains the JSON metadata for the struct
// [ListResourceProductPagination]
type listResourceProductPaginationJSON struct {
	MaxPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListResourceProductPagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResourceProductPaginationJSON) RawJSON() string {
	return r.raw
}

// A product.
type ProductOutput struct {
	// The ID of the product.
	ID string `json:"id,required" format:"uuid4"`
	// The benefits granted by the product.
	Benefits []ProductOutputBenefit `json:"benefits,required"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the product is archived and no longer available.
	IsArchived bool `json:"is_archived,required"`
	// Whether the product is a subscription tier.
	IsRecurring bool `json:"is_recurring,required"`
	// The medias associated to the product.
	Medias []ProductMediaFileReadOutput `json:"medias,required"`
	// The name of the product.
	Name string `json:"name,required"`
	// The ID of the organization owning the product.
	OrganizationID string `json:"organization_id,required" format:"uuid4"`
	// List of available prices for this product.
	Prices []ProductOutputPrice `json:"prices,required"`
	// The description of the product.
	Description   string `json:"description,nullable"`
	IsHighlighted bool   `json:"is_highlighted,nullable"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time         `json:"modified_at,nullable" format:"date-time"`
	Type       ProductOutputType `json:"type,nullable"`
	JSON       productOutputJSON `json:"-"`
}

// productOutputJSON contains the JSON metadata for the struct [ProductOutput]
type productOutputJSON struct {
	ID             apijson.Field
	Benefits       apijson.Field
	CreatedAt      apijson.Field
	IsArchived     apijson.Field
	IsRecurring    apijson.Field
	Medias         apijson.Field
	Name           apijson.Field
	OrganizationID apijson.Field
	Prices         apijson.Field
	Description    apijson.Field
	IsHighlighted  apijson.Field
	ModifiedAt     apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ProductOutput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productOutputJSON) RawJSON() string {
	return r.raw
}

// A benefit of type `articles`.
//
// Use it to grant access to posts.
type ProductOutputBenefit struct {
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,nullable" format:"date-time"`
	// The ID of the benefit.
	ID string `json:"id,required" format:"uuid4"`
	// The type of the benefit.
	Type ProductOutputBenefitsType `json:"type,required"`
	// The description of the benefit.
	Description string `json:"description,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool `json:"selectable,required"`
	// Whether the benefit is deletable.
	Deletable bool `json:"deletable,required"`
	// The ID of the organization owning the benefit.
	OrganizationID string `json:"organization_id,required" format:"uuid4"`
	// This field can have the runtime type of
	// [ProductOutputBenefitsBenefitArticlesProperties].
	Properties interface{}              `json:"properties,required"`
	JSON       productOutputBenefitJSON `json:"-"`
	union      ProductOutputBenefitsUnion
}

// productOutputBenefitJSON contains the JSON metadata for the struct
// [ProductOutputBenefit]
type productOutputBenefitJSON struct {
	CreatedAt      apijson.Field
	ModifiedAt     apijson.Field
	ID             apijson.Field
	Type           apijson.Field
	Description    apijson.Field
	Selectable     apijson.Field
	Deletable      apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r productOutputBenefitJSON) RawJSON() string {
	return r.raw
}

func (r *ProductOutputBenefit) UnmarshalJSON(data []byte) (err error) {
	*r = ProductOutputBenefit{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProductOutputBenefitsUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [ProductOutputBenefitsBenefitBase],
// [ProductOutputBenefitsBenefitArticles].
func (r ProductOutputBenefit) AsUnion() ProductOutputBenefitsUnion {
	return r.union
}

// A benefit of type `articles`.
//
// Use it to grant access to posts.
//
// Union satisfied by [ProductOutputBenefitsBenefitBase] or
// [ProductOutputBenefitsBenefitArticles].
type ProductOutputBenefitsUnion interface {
	implementsProductOutputBenefit()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProductOutputBenefitsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProductOutputBenefitsBenefitBase{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProductOutputBenefitsBenefitArticles{}),
		},
	)
}

type ProductOutputBenefitsBenefitBase struct {
	// The ID of the benefit.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the benefit is deletable.
	Deletable bool `json:"deletable,required"`
	// The description of the benefit.
	Description string `json:"description,required"`
	// The ID of the organization owning the benefit.
	OrganizationID string `json:"organization_id,required" format:"uuid4"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool `json:"selectable,required"`
	// The type of the benefit.
	Type ProductOutputBenefitsBenefitBaseType `json:"type,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time                            `json:"modified_at,nullable" format:"date-time"`
	JSON       productOutputBenefitsBenefitBaseJSON `json:"-"`
}

// productOutputBenefitsBenefitBaseJSON contains the JSON metadata for the struct
// [ProductOutputBenefitsBenefitBase]
type productOutputBenefitsBenefitBaseJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Deletable      apijson.Field
	Description    apijson.Field
	OrganizationID apijson.Field
	Selectable     apijson.Field
	Type           apijson.Field
	ModifiedAt     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ProductOutputBenefitsBenefitBase) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productOutputBenefitsBenefitBaseJSON) RawJSON() string {
	return r.raw
}

func (r ProductOutputBenefitsBenefitBase) implementsProductOutputBenefit() {}

// The type of the benefit.
type ProductOutputBenefitsBenefitBaseType string

const (
	ProductOutputBenefitsBenefitBaseTypeCustom           ProductOutputBenefitsBenefitBaseType = "custom"
	ProductOutputBenefitsBenefitBaseTypeArticles         ProductOutputBenefitsBenefitBaseType = "articles"
	ProductOutputBenefitsBenefitBaseTypeAds              ProductOutputBenefitsBenefitBaseType = "ads"
	ProductOutputBenefitsBenefitBaseTypeDiscord          ProductOutputBenefitsBenefitBaseType = "discord"
	ProductOutputBenefitsBenefitBaseTypeGitHubRepository ProductOutputBenefitsBenefitBaseType = "github_repository"
	ProductOutputBenefitsBenefitBaseTypeDownloadables    ProductOutputBenefitsBenefitBaseType = "downloadables"
)

func (r ProductOutputBenefitsBenefitBaseType) IsKnown() bool {
	switch r {
	case ProductOutputBenefitsBenefitBaseTypeCustom, ProductOutputBenefitsBenefitBaseTypeArticles, ProductOutputBenefitsBenefitBaseTypeAds, ProductOutputBenefitsBenefitBaseTypeDiscord, ProductOutputBenefitsBenefitBaseTypeGitHubRepository, ProductOutputBenefitsBenefitBaseTypeDownloadables:
		return true
	}
	return false
}

// A benefit of type `articles`.
//
// Use it to grant access to posts.
type ProductOutputBenefitsBenefitArticles struct {
	// The ID of the benefit.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the benefit is deletable.
	Deletable bool `json:"deletable,required"`
	// The description of the benefit.
	Description string `json:"description,required"`
	// The ID of the organization owning the benefit.
	OrganizationID string `json:"organization_id,required" format:"uuid4"`
	// Properties for a benefit of type `articles`.
	Properties ProductOutputBenefitsBenefitArticlesProperties `json:"properties,required"`
	// Whether the benefit is selectable when creating a product.
	Selectable bool                                     `json:"selectable,required"`
	Type       ProductOutputBenefitsBenefitArticlesType `json:"type,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time                                `json:"modified_at,nullable" format:"date-time"`
	JSON       productOutputBenefitsBenefitArticlesJSON `json:"-"`
}

// productOutputBenefitsBenefitArticlesJSON contains the JSON metadata for the
// struct [ProductOutputBenefitsBenefitArticles]
type productOutputBenefitsBenefitArticlesJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Deletable      apijson.Field
	Description    apijson.Field
	OrganizationID apijson.Field
	Properties     apijson.Field
	Selectable     apijson.Field
	Type           apijson.Field
	ModifiedAt     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ProductOutputBenefitsBenefitArticles) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productOutputBenefitsBenefitArticlesJSON) RawJSON() string {
	return r.raw
}

func (r ProductOutputBenefitsBenefitArticles) implementsProductOutputBenefit() {}

// Properties for a benefit of type `articles`.
type ProductOutputBenefitsBenefitArticlesProperties struct {
	// Whether the user can access paid articles.
	PaidArticles bool                                               `json:"paid_articles,required"`
	JSON         productOutputBenefitsBenefitArticlesPropertiesJSON `json:"-"`
}

// productOutputBenefitsBenefitArticlesPropertiesJSON contains the JSON metadata
// for the struct [ProductOutputBenefitsBenefitArticlesProperties]
type productOutputBenefitsBenefitArticlesPropertiesJSON struct {
	PaidArticles apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ProductOutputBenefitsBenefitArticlesProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productOutputBenefitsBenefitArticlesPropertiesJSON) RawJSON() string {
	return r.raw
}

type ProductOutputBenefitsBenefitArticlesType string

const (
	ProductOutputBenefitsBenefitArticlesTypeArticles ProductOutputBenefitsBenefitArticlesType = "articles"
)

func (r ProductOutputBenefitsBenefitArticlesType) IsKnown() bool {
	switch r {
	case ProductOutputBenefitsBenefitArticlesTypeArticles:
		return true
	}
	return false
}

// The type of the benefit.
type ProductOutputBenefitsType string

const (
	ProductOutputBenefitsTypeCustom           ProductOutputBenefitsType = "custom"
	ProductOutputBenefitsTypeArticles         ProductOutputBenefitsType = "articles"
	ProductOutputBenefitsTypeAds              ProductOutputBenefitsType = "ads"
	ProductOutputBenefitsTypeDiscord          ProductOutputBenefitsType = "discord"
	ProductOutputBenefitsTypeGitHubRepository ProductOutputBenefitsType = "github_repository"
	ProductOutputBenefitsTypeDownloadables    ProductOutputBenefitsType = "downloadables"
)

func (r ProductOutputBenefitsType) IsKnown() bool {
	switch r {
	case ProductOutputBenefitsTypeCustom, ProductOutputBenefitsTypeArticles, ProductOutputBenefitsTypeAds, ProductOutputBenefitsTypeDiscord, ProductOutputBenefitsTypeGitHubRepository, ProductOutputBenefitsTypeDownloadables:
		return true
	}
	return false
}

// A recurring price for a product, i.e. a subscription.
type ProductOutputPrice struct {
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,nullable" format:"date-time"`
	// The ID of the price.
	ID string `json:"id,required" format:"uuid4"`
	// The price in cents.
	PriceAmount int64 `json:"price_amount,required"`
	// The currency.
	PriceCurrency string `json:"price_currency,required"`
	// Whether the price is archived and no longer available.
	IsArchived bool `json:"is_archived,required"`
	// The type of the price.
	Type ProductOutputPricesType `json:"type,required"`
	// The recurring interval of the price, if type is `recurring`.
	RecurringInterval ProductOutputPricesRecurringInterval `json:"recurring_interval,nullable"`
	JSON              productOutputPriceJSON               `json:"-"`
	union             ProductOutputPricesUnion
}

// productOutputPriceJSON contains the JSON metadata for the struct
// [ProductOutputPrice]
type productOutputPriceJSON struct {
	CreatedAt         apijson.Field
	ModifiedAt        apijson.Field
	ID                apijson.Field
	PriceAmount       apijson.Field
	PriceCurrency     apijson.Field
	IsArchived        apijson.Field
	Type              apijson.Field
	RecurringInterval apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r productOutputPriceJSON) RawJSON() string {
	return r.raw
}

func (r *ProductOutputPrice) UnmarshalJSON(data []byte) (err error) {
	*r = ProductOutputPrice{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProductOutputPricesUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are
// [ProductOutputPricesProductPriceRecurring],
// [ProductOutputPricesProductPriceOneTime].
func (r ProductOutputPrice) AsUnion() ProductOutputPricesUnion {
	return r.union
}

// A recurring price for a product, i.e. a subscription.
//
// Union satisfied by [ProductOutputPricesProductPriceRecurring] or
// [ProductOutputPricesProductPriceOneTime].
type ProductOutputPricesUnion interface {
	implementsProductOutputPrice()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProductOutputPricesUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProductOutputPricesProductPriceRecurring{}),
			DiscriminatorValue: "recurring",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProductOutputPricesProductPriceOneTime{}),
			DiscriminatorValue: "one_time",
		},
	)
}

// A recurring price for a product, i.e. a subscription.
type ProductOutputPricesProductPriceRecurring struct {
	// The ID of the price.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the price is archived and no longer available.
	IsArchived bool `json:"is_archived,required"`
	// The price in cents.
	PriceAmount int64 `json:"price_amount,required"`
	// The currency.
	PriceCurrency string `json:"price_currency,required"`
	// The type of the price.
	Type ProductOutputPricesProductPriceRecurringType `json:"type,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,nullable" format:"date-time"`
	// The recurring interval of the price, if type is `recurring`.
	RecurringInterval ProductOutputPricesProductPriceRecurringRecurringInterval `json:"recurring_interval,nullable"`
	JSON              productOutputPricesProductPriceRecurringJSON              `json:"-"`
}

// productOutputPricesProductPriceRecurringJSON contains the JSON metadata for the
// struct [ProductOutputPricesProductPriceRecurring]
type productOutputPricesProductPriceRecurringJSON struct {
	ID                apijson.Field
	CreatedAt         apijson.Field
	IsArchived        apijson.Field
	PriceAmount       apijson.Field
	PriceCurrency     apijson.Field
	Type              apijson.Field
	ModifiedAt        apijson.Field
	RecurringInterval apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProductOutputPricesProductPriceRecurring) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productOutputPricesProductPriceRecurringJSON) RawJSON() string {
	return r.raw
}

func (r ProductOutputPricesProductPriceRecurring) implementsProductOutputPrice() {}

// The type of the price.
type ProductOutputPricesProductPriceRecurringType string

const (
	ProductOutputPricesProductPriceRecurringTypeRecurring ProductOutputPricesProductPriceRecurringType = "recurring"
)

func (r ProductOutputPricesProductPriceRecurringType) IsKnown() bool {
	switch r {
	case ProductOutputPricesProductPriceRecurringTypeRecurring:
		return true
	}
	return false
}

// The recurring interval of the price, if type is `recurring`.
type ProductOutputPricesProductPriceRecurringRecurringInterval string

const (
	ProductOutputPricesProductPriceRecurringRecurringIntervalMonth ProductOutputPricesProductPriceRecurringRecurringInterval = "month"
	ProductOutputPricesProductPriceRecurringRecurringIntervalYear  ProductOutputPricesProductPriceRecurringRecurringInterval = "year"
)

func (r ProductOutputPricesProductPriceRecurringRecurringInterval) IsKnown() bool {
	switch r {
	case ProductOutputPricesProductPriceRecurringRecurringIntervalMonth, ProductOutputPricesProductPriceRecurringRecurringIntervalYear:
		return true
	}
	return false
}

// A one-time price for a product.
type ProductOutputPricesProductPriceOneTime struct {
	// The ID of the price.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the price is archived and no longer available.
	IsArchived bool `json:"is_archived,required"`
	// The price in cents.
	PriceAmount int64 `json:"price_amount,required"`
	// The currency.
	PriceCurrency string `json:"price_currency,required"`
	// The type of the price.
	Type ProductOutputPricesProductPriceOneTimeType `json:"type,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time                                  `json:"modified_at,nullable" format:"date-time"`
	JSON       productOutputPricesProductPriceOneTimeJSON `json:"-"`
}

// productOutputPricesProductPriceOneTimeJSON contains the JSON metadata for the
// struct [ProductOutputPricesProductPriceOneTime]
type productOutputPricesProductPriceOneTimeJSON struct {
	ID            apijson.Field
	CreatedAt     apijson.Field
	IsArchived    apijson.Field
	PriceAmount   apijson.Field
	PriceCurrency apijson.Field
	Type          apijson.Field
	ModifiedAt    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProductOutputPricesProductPriceOneTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productOutputPricesProductPriceOneTimeJSON) RawJSON() string {
	return r.raw
}

func (r ProductOutputPricesProductPriceOneTime) implementsProductOutputPrice() {}

// The type of the price.
type ProductOutputPricesProductPriceOneTimeType string

const (
	ProductOutputPricesProductPriceOneTimeTypeOneTime ProductOutputPricesProductPriceOneTimeType = "one_time"
)

func (r ProductOutputPricesProductPriceOneTimeType) IsKnown() bool {
	switch r {
	case ProductOutputPricesProductPriceOneTimeTypeOneTime:
		return true
	}
	return false
}

// The type of the price.
type ProductOutputPricesType string

const (
	ProductOutputPricesTypeRecurring ProductOutputPricesType = "recurring"
	ProductOutputPricesTypeOneTime   ProductOutputPricesType = "one_time"
)

func (r ProductOutputPricesType) IsKnown() bool {
	switch r {
	case ProductOutputPricesTypeRecurring, ProductOutputPricesTypeOneTime:
		return true
	}
	return false
}

// The recurring interval of the price, if type is `recurring`.
type ProductOutputPricesRecurringInterval string

const (
	ProductOutputPricesRecurringIntervalMonth ProductOutputPricesRecurringInterval = "month"
	ProductOutputPricesRecurringIntervalYear  ProductOutputPricesRecurringInterval = "year"
)

func (r ProductOutputPricesRecurringInterval) IsKnown() bool {
	switch r {
	case ProductOutputPricesRecurringIntervalMonth, ProductOutputPricesRecurringIntervalYear:
		return true
	}
	return false
}

type ProductOutputType string

const (
	ProductOutputTypeFree       ProductOutputType = "free"
	ProductOutputTypeIndividual ProductOutputType = "individual"
	ProductOutputTypeBusiness   ProductOutputType = "business"
)

func (r ProductOutputType) IsKnown() bool {
	switch r {
	case ProductOutputTypeFree, ProductOutputTypeIndividual, ProductOutputTypeBusiness:
		return true
	}
	return false
}

type ProductNewParams struct {
	// Schema to create a recurring product, i.e. a subscription.
	Body ProductNewParamsBodyUnion `json:"body,required"`
}

func (r ProductNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// Schema to create a recurring product, i.e. a subscription.
type ProductNewParamsBody struct {
	// The name of the product.
	Name param.Field[string] `json:"name,required"`
	// The description of the product.
	Description param.Field[string]      `json:"description"`
	Prices      param.Field[interface{}] `json:"prices"`
	Medias      param.Field[interface{}] `json:"medias,required"`
	// The organization ID.
	OrganizationID param.Field[string]                   `json:"organization_id" format:"uuid4"`
	Type           param.Field[ProductNewParamsBodyType] `json:"type"`
	IsHighlighted  param.Field[bool]                     `json:"is_highlighted"`
}

func (r ProductNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProductNewParamsBody) implementsProductNewParamsBodyUnion() {}

// Schema to create a recurring product, i.e. a subscription.
//
// Satisfied by [ProductNewParamsBodyProductRecurringCreate],
// [ProductNewParamsBodyProductOneTimeCreate], [ProductNewParamsBody].
type ProductNewParamsBodyUnion interface {
	implementsProductNewParamsBodyUnion()
}

// Schema to create a recurring product, i.e. a subscription.
type ProductNewParamsBodyProductRecurringCreate struct {
	// The name of the product.
	Name param.Field[string] `json:"name,required"`
	// List of available prices for this product.
	Prices param.Field[[]ProductNewParamsBodyProductRecurringCreatePrice] `json:"prices,required"`
	Type   param.Field[ProductNewParamsBodyProductRecurringCreateType]    `json:"type,required"`
	// The description of the product.
	Description   param.Field[string] `json:"description"`
	IsHighlighted param.Field[bool]   `json:"is_highlighted"`
	// List of file IDs. Each one must be on the same organization as the product, of
	// type `product_media` and correctly uploaded.
	Medias param.Field[[]string] `json:"medias" format:"uuid4"`
	// The organization ID.
	OrganizationID param.Field[string] `json:"organization_id" format:"uuid4"`
}

func (r ProductNewParamsBodyProductRecurringCreate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProductNewParamsBodyProductRecurringCreate) implementsProductNewParamsBodyUnion() {}

// Schema to create a recurring product price, i.e. a subscription.
type ProductNewParamsBodyProductRecurringCreatePrice struct {
	// The price in cents.
	PriceAmount param.Field[int64] `json:"price_amount,required"`
	// The recurring interval of the price.
	RecurringInterval param.Field[ProductNewParamsBodyProductRecurringCreatePricesRecurringInterval] `json:"recurring_interval,required"`
	Type              param.Field[ProductNewParamsBodyProductRecurringCreatePricesType]              `json:"type,required"`
	// The currency. Currently, only `usd` is supported.
	PriceCurrency param.Field[string] `json:"price_currency"`
}

func (r ProductNewParamsBodyProductRecurringCreatePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The recurring interval of the price.
type ProductNewParamsBodyProductRecurringCreatePricesRecurringInterval string

const (
	ProductNewParamsBodyProductRecurringCreatePricesRecurringIntervalMonth ProductNewParamsBodyProductRecurringCreatePricesRecurringInterval = "month"
	ProductNewParamsBodyProductRecurringCreatePricesRecurringIntervalYear  ProductNewParamsBodyProductRecurringCreatePricesRecurringInterval = "year"
)

func (r ProductNewParamsBodyProductRecurringCreatePricesRecurringInterval) IsKnown() bool {
	switch r {
	case ProductNewParamsBodyProductRecurringCreatePricesRecurringIntervalMonth, ProductNewParamsBodyProductRecurringCreatePricesRecurringIntervalYear:
		return true
	}
	return false
}

type ProductNewParamsBodyProductRecurringCreatePricesType string

const (
	ProductNewParamsBodyProductRecurringCreatePricesTypeRecurring ProductNewParamsBodyProductRecurringCreatePricesType = "recurring"
)

func (r ProductNewParamsBodyProductRecurringCreatePricesType) IsKnown() bool {
	switch r {
	case ProductNewParamsBodyProductRecurringCreatePricesTypeRecurring:
		return true
	}
	return false
}

type ProductNewParamsBodyProductRecurringCreateType string

const (
	ProductNewParamsBodyProductRecurringCreateTypeIndividual ProductNewParamsBodyProductRecurringCreateType = "individual"
	ProductNewParamsBodyProductRecurringCreateTypeBusiness   ProductNewParamsBodyProductRecurringCreateType = "business"
)

func (r ProductNewParamsBodyProductRecurringCreateType) IsKnown() bool {
	switch r {
	case ProductNewParamsBodyProductRecurringCreateTypeIndividual, ProductNewParamsBodyProductRecurringCreateTypeBusiness:
		return true
	}
	return false
}

// Schema to create a one-time product.
type ProductNewParamsBodyProductOneTimeCreate struct {
	// The name of the product.
	Name param.Field[string] `json:"name,required"`
	// List of available prices for this product.
	Prices param.Field[[]ProductNewParamsBodyProductOneTimeCreatePrice] `json:"prices,required"`
	// The description of the product.
	Description param.Field[string] `json:"description"`
	// List of file IDs. Each one must be on the same organization as the product, of
	// type `product_media` and correctly uploaded.
	Medias param.Field[[]string] `json:"medias" format:"uuid4"`
	// The organization ID.
	OrganizationID param.Field[string] `json:"organization_id" format:"uuid4"`
}

func (r ProductNewParamsBodyProductOneTimeCreate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProductNewParamsBodyProductOneTimeCreate) implementsProductNewParamsBodyUnion() {}

// Schema to create a one-time product price.
type ProductNewParamsBodyProductOneTimeCreatePrice struct {
	// The price in cents.
	PriceAmount param.Field[int64]                                              `json:"price_amount,required"`
	Type        param.Field[ProductNewParamsBodyProductOneTimeCreatePricesType] `json:"type,required"`
	// The currency. Currently, only `usd` is supported.
	PriceCurrency param.Field[string] `json:"price_currency"`
}

func (r ProductNewParamsBodyProductOneTimeCreatePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProductNewParamsBodyProductOneTimeCreatePricesType string

const (
	ProductNewParamsBodyProductOneTimeCreatePricesTypeOneTime ProductNewParamsBodyProductOneTimeCreatePricesType = "one_time"
)

func (r ProductNewParamsBodyProductOneTimeCreatePricesType) IsKnown() bool {
	switch r {
	case ProductNewParamsBodyProductOneTimeCreatePricesTypeOneTime:
		return true
	}
	return false
}

type ProductNewParamsBodyType string

const (
	ProductNewParamsBodyTypeIndividual ProductNewParamsBodyType = "individual"
	ProductNewParamsBodyTypeBusiness   ProductNewParamsBodyType = "business"
)

func (r ProductNewParamsBodyType) IsKnown() bool {
	switch r {
	case ProductNewParamsBodyTypeIndividual, ProductNewParamsBodyTypeBusiness:
		return true
	}
	return false
}

type ProductUpdateParams struct {
	// The description of the product.
	Description param.Field[string] `json:"description"`
	// Whether the product is archived. If `true`, the product won't be available for
	// purchase anymore. Existing customers will still have access to their benefits,
	// and subscriptions will continue normally.
	IsArchived    param.Field[bool] `json:"is_archived"`
	IsHighlighted param.Field[bool] `json:"is_highlighted"`
	// List of file IDs. Each one must be on the same organization as the product, of
	// type `product_media` and correctly uploaded.
	Medias param.Field[[]string] `json:"medias" format:"uuid4"`
	// The name of the product.
	Name param.Field[string] `json:"name"`
	// List of available prices for this product. If you want to keep existing prices,
	// include them in the list as an `ExistingProductPrice` object.
	Prices param.Field[[]ProductUpdateParamsPriceUnion] `json:"prices"`
}

func (r ProductUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A price that already exists for this product.
//
// Useful when updating a product if you want to keep an existing price.
type ProductUpdateParamsPrice struct {
	ID   param.Field[string]                        `json:"id" format:"uuid4"`
	Type param.Field[ProductUpdateParamsPricesType] `json:"type"`
	// The recurring interval of the price.
	RecurringInterval param.Field[ProductUpdateParamsPricesRecurringInterval] `json:"recurring_interval"`
	// The price in cents.
	PriceAmount param.Field[int64] `json:"price_amount"`
	// The currency. Currently, only `usd` is supported.
	PriceCurrency param.Field[string] `json:"price_currency"`
}

func (r ProductUpdateParamsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProductUpdateParamsPrice) implementsProductUpdateParamsPriceUnion() {}

// A price that already exists for this product.
//
// Useful when updating a product if you want to keep an existing price.
//
// Satisfied by [ProductUpdateParamsPricesExistingProductPrice],
// [ProductUpdateParamsPricesProductPriceRecurringCreate],
// [ProductUpdateParamsPricesProductPriceOneTimeCreate],
// [ProductUpdateParamsPrice].
type ProductUpdateParamsPriceUnion interface {
	implementsProductUpdateParamsPriceUnion()
}

// A price that already exists for this product.
//
// Useful when updating a product if you want to keep an existing price.
type ProductUpdateParamsPricesExistingProductPrice struct {
	ID param.Field[string] `json:"id,required" format:"uuid4"`
}

func (r ProductUpdateParamsPricesExistingProductPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProductUpdateParamsPricesExistingProductPrice) implementsProductUpdateParamsPriceUnion() {}

// Schema to create a recurring product price, i.e. a subscription.
type ProductUpdateParamsPricesProductPriceRecurringCreate struct {
	// The price in cents.
	PriceAmount param.Field[int64] `json:"price_amount,required"`
	// The recurring interval of the price.
	RecurringInterval param.Field[ProductUpdateParamsPricesProductPriceRecurringCreateRecurringInterval] `json:"recurring_interval,required"`
	Type              param.Field[ProductUpdateParamsPricesProductPriceRecurringCreateType]              `json:"type,required"`
	// The currency. Currently, only `usd` is supported.
	PriceCurrency param.Field[string] `json:"price_currency"`
}

func (r ProductUpdateParamsPricesProductPriceRecurringCreate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProductUpdateParamsPricesProductPriceRecurringCreate) implementsProductUpdateParamsPriceUnion() {
}

// The recurring interval of the price.
type ProductUpdateParamsPricesProductPriceRecurringCreateRecurringInterval string

const (
	ProductUpdateParamsPricesProductPriceRecurringCreateRecurringIntervalMonth ProductUpdateParamsPricesProductPriceRecurringCreateRecurringInterval = "month"
	ProductUpdateParamsPricesProductPriceRecurringCreateRecurringIntervalYear  ProductUpdateParamsPricesProductPriceRecurringCreateRecurringInterval = "year"
)

func (r ProductUpdateParamsPricesProductPriceRecurringCreateRecurringInterval) IsKnown() bool {
	switch r {
	case ProductUpdateParamsPricesProductPriceRecurringCreateRecurringIntervalMonth, ProductUpdateParamsPricesProductPriceRecurringCreateRecurringIntervalYear:
		return true
	}
	return false
}

type ProductUpdateParamsPricesProductPriceRecurringCreateType string

const (
	ProductUpdateParamsPricesProductPriceRecurringCreateTypeRecurring ProductUpdateParamsPricesProductPriceRecurringCreateType = "recurring"
)

func (r ProductUpdateParamsPricesProductPriceRecurringCreateType) IsKnown() bool {
	switch r {
	case ProductUpdateParamsPricesProductPriceRecurringCreateTypeRecurring:
		return true
	}
	return false
}

// Schema to create a one-time product price.
type ProductUpdateParamsPricesProductPriceOneTimeCreate struct {
	// The price in cents.
	PriceAmount param.Field[int64]                                                  `json:"price_amount,required"`
	Type        param.Field[ProductUpdateParamsPricesProductPriceOneTimeCreateType] `json:"type,required"`
	// The currency. Currently, only `usd` is supported.
	PriceCurrency param.Field[string] `json:"price_currency"`
}

func (r ProductUpdateParamsPricesProductPriceOneTimeCreate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProductUpdateParamsPricesProductPriceOneTimeCreate) implementsProductUpdateParamsPriceUnion() {
}

type ProductUpdateParamsPricesProductPriceOneTimeCreateType string

const (
	ProductUpdateParamsPricesProductPriceOneTimeCreateTypeOneTime ProductUpdateParamsPricesProductPriceOneTimeCreateType = "one_time"
)

func (r ProductUpdateParamsPricesProductPriceOneTimeCreateType) IsKnown() bool {
	switch r {
	case ProductUpdateParamsPricesProductPriceOneTimeCreateTypeOneTime:
		return true
	}
	return false
}

type ProductUpdateParamsPricesType string

const (
	ProductUpdateParamsPricesTypeRecurring ProductUpdateParamsPricesType = "recurring"
	ProductUpdateParamsPricesTypeOneTime   ProductUpdateParamsPricesType = "one_time"
)

func (r ProductUpdateParamsPricesType) IsKnown() bool {
	switch r {
	case ProductUpdateParamsPricesTypeRecurring, ProductUpdateParamsPricesTypeOneTime:
		return true
	}
	return false
}

// The recurring interval of the price.
type ProductUpdateParamsPricesRecurringInterval string

const (
	ProductUpdateParamsPricesRecurringIntervalMonth ProductUpdateParamsPricesRecurringInterval = "month"
	ProductUpdateParamsPricesRecurringIntervalYear  ProductUpdateParamsPricesRecurringInterval = "year"
)

func (r ProductUpdateParamsPricesRecurringInterval) IsKnown() bool {
	switch r {
	case ProductUpdateParamsPricesRecurringIntervalMonth, ProductUpdateParamsPricesRecurringIntervalYear:
		return true
	}
	return false
}

type ProductListParams struct {
	// Filter products granting specific benefit.
	BenefitID param.Field[ProductListParamsBenefitIDUnion] `query:"benefit_id" format:"uuid4"`
	// Filter on archived products.
	IsArchived param.Field[bool] `query:"is_archived"`
	// Filter on recurring products. If `true`, only subscriptions tiers are returned.
	// If `false`, only one-time purchase products are returned.
	IsRecurring param.Field[bool] `query:"is_recurring"`
	// Size of a page, defaults to 10. Maximum is 100.
	Limit param.Field[int64] `query:"limit"`
	// Filter by organization ID.
	OrganizationID param.Field[ProductListParamsOrganizationIDUnion] `query:"organization_id" format:"uuid4"`
	// Page number, defaults to 1.
	Page param.Field[int64] `query:"page"`
	// Filter by subscription tier type.
	Type param.Field[ProductListParamsTypeUnion] `query:"type"`
}

// URLQuery serializes [ProductListParams]'s query parameters as `url.Values`.
func (r ProductListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter products granting specific benefit.
//
// Satisfied by [shared.UnionString], [ProductListParamsBenefitIDArray].
type ProductListParamsBenefitIDUnion interface {
	ImplementsProductListParamsBenefitIDUnion()
}

type ProductListParamsBenefitIDArray []string

func (r ProductListParamsBenefitIDArray) ImplementsProductListParamsBenefitIDUnion() {}

// Filter by organization ID.
//
// Satisfied by [shared.UnionString], [ProductListParamsOrganizationIDArray].
type ProductListParamsOrganizationIDUnion interface {
	ImplementsProductListParamsOrganizationIDUnion()
}

type ProductListParamsOrganizationIDArray []string

func (r ProductListParamsOrganizationIDArray) ImplementsProductListParamsOrganizationIDUnion() {}

// Filter by subscription tier type.
//
// Satisfied by [ProductListParamsTypeSubscriptionTierType],
// [ProductListParamsTypeArray].
type ProductListParamsTypeUnion interface {
	implementsProductListParamsTypeUnion()
}

type ProductListParamsTypeSubscriptionTierType string

const (
	ProductListParamsTypeSubscriptionTierTypeFree       ProductListParamsTypeSubscriptionTierType = "free"
	ProductListParamsTypeSubscriptionTierTypeIndividual ProductListParamsTypeSubscriptionTierType = "individual"
	ProductListParamsTypeSubscriptionTierTypeBusiness   ProductListParamsTypeSubscriptionTierType = "business"
)

func (r ProductListParamsTypeSubscriptionTierType) IsKnown() bool {
	switch r {
	case ProductListParamsTypeSubscriptionTierTypeFree, ProductListParamsTypeSubscriptionTierTypeIndividual, ProductListParamsTypeSubscriptionTierTypeBusiness:
		return true
	}
	return false
}

func (r ProductListParamsTypeSubscriptionTierType) implementsProductListParamsTypeUnion() {}

type ProductListParamsTypeArray []ProductListParamsTypeArray

func (r ProductListParamsTypeArray) implementsProductListParamsTypeUnion() {}
