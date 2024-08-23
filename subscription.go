// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package polar

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/polarsource/polar-go/internal/apiform"
	"github.com/polarsource/polar-go/internal/apijson"
	"github.com/polarsource/polar-go/internal/apiquery"
	"github.com/polarsource/polar-go/internal/param"
	"github.com/polarsource/polar-go/internal/requestconfig"
	"github.com/polarsource/polar-go/option"
	"github.com/tidwall/gjson"
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
func (r *SubscriptionService) List(ctx context.Context, query SubscriptionListParams, opts ...option.RequestOption) (res *SubscriptionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/subscriptions/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
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
	ID                string `json:"id,required" format:"uuid4"`
	CancelAtPeriodEnd bool   `json:"cancel_at_period_end,required"`
	// Creation timestamp of the object.
	CreatedAt          time.Time `json:"created_at,required" format:"date-time"`
	CurrentPeriodEnd   time.Time `json:"current_period_end,required,nullable" format:"date-time"`
	CurrentPeriodStart time.Time `json:"current_period_start,required" format:"date-time"`
	EndedAt            time.Time `json:"ended_at,required,nullable" format:"date-time"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// A recurring price for a product, i.e. a subscription.
	Price   SubscriptionNewResponsePrice `json:"price,required,nullable"`
	PriceID string                       `json:"price_id,required,nullable" format:"uuid4"`
	// A product.
	Product   ProductOutput                 `json:"product,required"`
	ProductID string                        `json:"product_id,required" format:"uuid4"`
	StartedAt time.Time                     `json:"started_at,required,nullable" format:"date-time"`
	Status    SubscriptionNewResponseStatus `json:"status,required"`
	User      SubscriptionNewResponseUser   `json:"user,required"`
	UserID    string                        `json:"user_id,required" format:"uuid4"`
	JSON      subscriptionNewResponseJSON   `json:"-"`
}

// subscriptionNewResponseJSON contains the JSON metadata for the struct
// [SubscriptionNewResponse]
type subscriptionNewResponseJSON struct {
	ID                 apijson.Field
	CancelAtPeriodEnd  apijson.Field
	CreatedAt          apijson.Field
	CurrentPeriodEnd   apijson.Field
	CurrentPeriodStart apijson.Field
	EndedAt            apijson.Field
	ModifiedAt         apijson.Field
	Price              apijson.Field
	PriceID            apijson.Field
	Product            apijson.Field
	ProductID          apijson.Field
	StartedAt          apijson.Field
	Status             apijson.Field
	User               apijson.Field
	UserID             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SubscriptionNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionNewResponseJSON) RawJSON() string {
	return r.raw
}

// A recurring price for a product, i.e. a subscription.
type SubscriptionNewResponsePrice struct {
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The ID of the price.
	ID string `json:"id,required" format:"uuid4"`
	// The price in cents.
	PriceAmount int64 `json:"price_amount,required"`
	// The currency.
	PriceCurrency string `json:"price_currency,required"`
	// Whether the price is archived and no longer available.
	IsArchived bool `json:"is_archived,required"`
	// The type of the price.
	Type SubscriptionNewResponsePriceType `json:"type,required"`
	// The recurring interval of the price, if type is `recurring`.
	RecurringInterval SubscriptionNewResponsePriceRecurringInterval `json:"recurring_interval,nullable"`
	JSON              subscriptionNewResponsePriceJSON              `json:"-"`
	union             SubscriptionNewResponsePriceUnion
}

// subscriptionNewResponsePriceJSON contains the JSON metadata for the struct
// [SubscriptionNewResponsePrice]
type subscriptionNewResponsePriceJSON struct {
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

func (r subscriptionNewResponsePriceJSON) RawJSON() string {
	return r.raw
}

func (r *SubscriptionNewResponsePrice) UnmarshalJSON(data []byte) (err error) {
	*r = SubscriptionNewResponsePrice{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SubscriptionNewResponsePriceUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SubscriptionNewResponsePriceProductPriceRecurring],
// [SubscriptionNewResponsePriceProductPriceOneTime].
func (r SubscriptionNewResponsePrice) AsUnion() SubscriptionNewResponsePriceUnion {
	return r.union
}

// A recurring price for a product, i.e. a subscription.
//
// Union satisfied by [SubscriptionNewResponsePriceProductPriceRecurring] or
// [SubscriptionNewResponsePriceProductPriceOneTime].
type SubscriptionNewResponsePriceUnion interface {
	implementsSubscriptionNewResponsePrice()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SubscriptionNewResponsePriceUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SubscriptionNewResponsePriceProductPriceRecurring{}),
			DiscriminatorValue: "recurring",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SubscriptionNewResponsePriceProductPriceOneTime{}),
			DiscriminatorValue: "one_time",
		},
	)
}

// A recurring price for a product, i.e. a subscription.
type SubscriptionNewResponsePriceProductPriceRecurring struct {
	// The ID of the price.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the price is archived and no longer available.
	IsArchived bool `json:"is_archived,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The price in cents.
	PriceAmount int64 `json:"price_amount,required"`
	// The currency.
	PriceCurrency string `json:"price_currency,required"`
	// The recurring interval of the price, if type is `recurring`.
	RecurringInterval SubscriptionNewResponsePriceProductPriceRecurringRecurringInterval `json:"recurring_interval,required,nullable"`
	// The type of the price.
	Type SubscriptionNewResponsePriceProductPriceRecurringType `json:"type,required"`
	JSON subscriptionNewResponsePriceProductPriceRecurringJSON `json:"-"`
}

// subscriptionNewResponsePriceProductPriceRecurringJSON contains the JSON metadata
// for the struct [SubscriptionNewResponsePriceProductPriceRecurring]
type subscriptionNewResponsePriceProductPriceRecurringJSON struct {
	ID                apijson.Field
	CreatedAt         apijson.Field
	IsArchived        apijson.Field
	ModifiedAt        apijson.Field
	PriceAmount       apijson.Field
	PriceCurrency     apijson.Field
	RecurringInterval apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SubscriptionNewResponsePriceProductPriceRecurring) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionNewResponsePriceProductPriceRecurringJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionNewResponsePriceProductPriceRecurring) implementsSubscriptionNewResponsePrice() {}

// The recurring interval of the price, if type is `recurring`.
type SubscriptionNewResponsePriceProductPriceRecurringRecurringInterval string

const (
	SubscriptionNewResponsePriceProductPriceRecurringRecurringIntervalMonth SubscriptionNewResponsePriceProductPriceRecurringRecurringInterval = "month"
	SubscriptionNewResponsePriceProductPriceRecurringRecurringIntervalYear  SubscriptionNewResponsePriceProductPriceRecurringRecurringInterval = "year"
)

func (r SubscriptionNewResponsePriceProductPriceRecurringRecurringInterval) IsKnown() bool {
	switch r {
	case SubscriptionNewResponsePriceProductPriceRecurringRecurringIntervalMonth, SubscriptionNewResponsePriceProductPriceRecurringRecurringIntervalYear:
		return true
	}
	return false
}

// The type of the price.
type SubscriptionNewResponsePriceProductPriceRecurringType string

const (
	SubscriptionNewResponsePriceProductPriceRecurringTypeRecurring SubscriptionNewResponsePriceProductPriceRecurringType = "recurring"
)

func (r SubscriptionNewResponsePriceProductPriceRecurringType) IsKnown() bool {
	switch r {
	case SubscriptionNewResponsePriceProductPriceRecurringTypeRecurring:
		return true
	}
	return false
}

// A one-time price for a product.
type SubscriptionNewResponsePriceProductPriceOneTime struct {
	// The ID of the price.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the price is archived and no longer available.
	IsArchived bool `json:"is_archived,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The price in cents.
	PriceAmount int64 `json:"price_amount,required"`
	// The currency.
	PriceCurrency string `json:"price_currency,required"`
	// The type of the price.
	Type SubscriptionNewResponsePriceProductPriceOneTimeType `json:"type,required"`
	JSON subscriptionNewResponsePriceProductPriceOneTimeJSON `json:"-"`
}

// subscriptionNewResponsePriceProductPriceOneTimeJSON contains the JSON metadata
// for the struct [SubscriptionNewResponsePriceProductPriceOneTime]
type subscriptionNewResponsePriceProductPriceOneTimeJSON struct {
	ID            apijson.Field
	CreatedAt     apijson.Field
	IsArchived    apijson.Field
	ModifiedAt    apijson.Field
	PriceAmount   apijson.Field
	PriceCurrency apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SubscriptionNewResponsePriceProductPriceOneTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionNewResponsePriceProductPriceOneTimeJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionNewResponsePriceProductPriceOneTime) implementsSubscriptionNewResponsePrice() {}

// The type of the price.
type SubscriptionNewResponsePriceProductPriceOneTimeType string

const (
	SubscriptionNewResponsePriceProductPriceOneTimeTypeOneTime SubscriptionNewResponsePriceProductPriceOneTimeType = "one_time"
)

func (r SubscriptionNewResponsePriceProductPriceOneTimeType) IsKnown() bool {
	switch r {
	case SubscriptionNewResponsePriceProductPriceOneTimeTypeOneTime:
		return true
	}
	return false
}

// The type of the price.
type SubscriptionNewResponsePriceType string

const (
	SubscriptionNewResponsePriceTypeRecurring SubscriptionNewResponsePriceType = "recurring"
	SubscriptionNewResponsePriceTypeOneTime   SubscriptionNewResponsePriceType = "one_time"
)

func (r SubscriptionNewResponsePriceType) IsKnown() bool {
	switch r {
	case SubscriptionNewResponsePriceTypeRecurring, SubscriptionNewResponsePriceTypeOneTime:
		return true
	}
	return false
}

// The recurring interval of the price, if type is `recurring`.
type SubscriptionNewResponsePriceRecurringInterval string

const (
	SubscriptionNewResponsePriceRecurringIntervalMonth SubscriptionNewResponsePriceRecurringInterval = "month"
	SubscriptionNewResponsePriceRecurringIntervalYear  SubscriptionNewResponsePriceRecurringInterval = "year"
)

func (r SubscriptionNewResponsePriceRecurringInterval) IsKnown() bool {
	switch r {
	case SubscriptionNewResponsePriceRecurringIntervalMonth, SubscriptionNewResponsePriceRecurringIntervalYear:
		return true
	}
	return false
}

type SubscriptionNewResponseStatus string

const (
	SubscriptionNewResponseStatusIncomplete        SubscriptionNewResponseStatus = "incomplete"
	SubscriptionNewResponseStatusIncompleteExpired SubscriptionNewResponseStatus = "incomplete_expired"
	SubscriptionNewResponseStatusTrialing          SubscriptionNewResponseStatus = "trialing"
	SubscriptionNewResponseStatusActive            SubscriptionNewResponseStatus = "active"
	SubscriptionNewResponseStatusPastDue           SubscriptionNewResponseStatus = "past_due"
	SubscriptionNewResponseStatusCanceled          SubscriptionNewResponseStatus = "canceled"
	SubscriptionNewResponseStatusUnpaid            SubscriptionNewResponseStatus = "unpaid"
)

func (r SubscriptionNewResponseStatus) IsKnown() bool {
	switch r {
	case SubscriptionNewResponseStatusIncomplete, SubscriptionNewResponseStatusIncompleteExpired, SubscriptionNewResponseStatusTrialing, SubscriptionNewResponseStatusActive, SubscriptionNewResponseStatusPastDue, SubscriptionNewResponseStatusCanceled, SubscriptionNewResponseStatusUnpaid:
		return true
	}
	return false
}

type SubscriptionNewResponseUser struct {
	AvatarURL      string                          `json:"avatar_url,required,nullable"`
	Email          string                          `json:"email,required"`
	GitHubUsername string                          `json:"github_username,required,nullable"`
	PublicName     string                          `json:"public_name,required"`
	JSON           subscriptionNewResponseUserJSON `json:"-"`
}

// subscriptionNewResponseUserJSON contains the JSON metadata for the struct
// [SubscriptionNewResponseUser]
type subscriptionNewResponseUserJSON struct {
	AvatarURL      apijson.Field
	Email          apijson.Field
	GitHubUsername apijson.Field
	PublicName     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SubscriptionNewResponseUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionNewResponseUserJSON) RawJSON() string {
	return r.raw
}

type SubscriptionListResponse struct {
	Pagination SubscriptionListResponsePagination `json:"pagination,required"`
	Items      []SubscriptionListResponseItem     `json:"items"`
	JSON       subscriptionListResponseJSON       `json:"-"`
}

// subscriptionListResponseJSON contains the JSON metadata for the struct
// [SubscriptionListResponse]
type subscriptionListResponseJSON struct {
	Pagination  apijson.Field
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionListResponseJSON) RawJSON() string {
	return r.raw
}

type SubscriptionListResponsePagination struct {
	MaxPage    int64                                  `json:"max_page,required"`
	TotalCount int64                                  `json:"total_count,required"`
	JSON       subscriptionListResponsePaginationJSON `json:"-"`
}

// subscriptionListResponsePaginationJSON contains the JSON metadata for the struct
// [SubscriptionListResponsePagination]
type subscriptionListResponsePaginationJSON struct {
	MaxPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionListResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionListResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type SubscriptionListResponseItem struct {
	// The ID of the object.
	ID                string `json:"id,required" format:"uuid4"`
	CancelAtPeriodEnd bool   `json:"cancel_at_period_end,required"`
	// Creation timestamp of the object.
	CreatedAt          time.Time `json:"created_at,required" format:"date-time"`
	CurrentPeriodEnd   time.Time `json:"current_period_end,required,nullable" format:"date-time"`
	CurrentPeriodStart time.Time `json:"current_period_start,required" format:"date-time"`
	EndedAt            time.Time `json:"ended_at,required,nullable" format:"date-time"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// A recurring price for a product, i.e. a subscription.
	Price   SubscriptionListResponseItemsPrice `json:"price,required,nullable"`
	PriceID string                             `json:"price_id,required,nullable" format:"uuid4"`
	// A product.
	Product   ProductOutput                       `json:"product,required"`
	ProductID string                              `json:"product_id,required" format:"uuid4"`
	StartedAt time.Time                           `json:"started_at,required,nullable" format:"date-time"`
	Status    SubscriptionListResponseItemsStatus `json:"status,required"`
	User      SubscriptionListResponseItemsUser   `json:"user,required"`
	UserID    string                              `json:"user_id,required" format:"uuid4"`
	JSON      subscriptionListResponseItemJSON    `json:"-"`
}

// subscriptionListResponseItemJSON contains the JSON metadata for the struct
// [SubscriptionListResponseItem]
type subscriptionListResponseItemJSON struct {
	ID                 apijson.Field
	CancelAtPeriodEnd  apijson.Field
	CreatedAt          apijson.Field
	CurrentPeriodEnd   apijson.Field
	CurrentPeriodStart apijson.Field
	EndedAt            apijson.Field
	ModifiedAt         apijson.Field
	Price              apijson.Field
	PriceID            apijson.Field
	Product            apijson.Field
	ProductID          apijson.Field
	StartedAt          apijson.Field
	Status             apijson.Field
	User               apijson.Field
	UserID             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SubscriptionListResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionListResponseItemJSON) RawJSON() string {
	return r.raw
}

// A recurring price for a product, i.e. a subscription.
type SubscriptionListResponseItemsPrice struct {
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The ID of the price.
	ID string `json:"id,required" format:"uuid4"`
	// The price in cents.
	PriceAmount int64 `json:"price_amount,required"`
	// The currency.
	PriceCurrency string `json:"price_currency,required"`
	// Whether the price is archived and no longer available.
	IsArchived bool `json:"is_archived,required"`
	// The type of the price.
	Type SubscriptionListResponseItemsPriceType `json:"type,required"`
	// The recurring interval of the price, if type is `recurring`.
	RecurringInterval SubscriptionListResponseItemsPriceRecurringInterval `json:"recurring_interval,nullable"`
	JSON              subscriptionListResponseItemsPriceJSON              `json:"-"`
	union             SubscriptionListResponseItemsPriceUnion
}

// subscriptionListResponseItemsPriceJSON contains the JSON metadata for the struct
// [SubscriptionListResponseItemsPrice]
type subscriptionListResponseItemsPriceJSON struct {
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

func (r subscriptionListResponseItemsPriceJSON) RawJSON() string {
	return r.raw
}

func (r *SubscriptionListResponseItemsPrice) UnmarshalJSON(data []byte) (err error) {
	*r = SubscriptionListResponseItemsPrice{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SubscriptionListResponseItemsPriceUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SubscriptionListResponseItemsPriceProductPriceRecurring],
// [SubscriptionListResponseItemsPriceProductPriceOneTime].
func (r SubscriptionListResponseItemsPrice) AsUnion() SubscriptionListResponseItemsPriceUnion {
	return r.union
}

// A recurring price for a product, i.e. a subscription.
//
// Union satisfied by [SubscriptionListResponseItemsPriceProductPriceRecurring] or
// [SubscriptionListResponseItemsPriceProductPriceOneTime].
type SubscriptionListResponseItemsPriceUnion interface {
	implementsSubscriptionListResponseItemsPrice()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SubscriptionListResponseItemsPriceUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SubscriptionListResponseItemsPriceProductPriceRecurring{}),
			DiscriminatorValue: "recurring",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SubscriptionListResponseItemsPriceProductPriceOneTime{}),
			DiscriminatorValue: "one_time",
		},
	)
}

// A recurring price for a product, i.e. a subscription.
type SubscriptionListResponseItemsPriceProductPriceRecurring struct {
	// The ID of the price.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the price is archived and no longer available.
	IsArchived bool `json:"is_archived,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The price in cents.
	PriceAmount int64 `json:"price_amount,required"`
	// The currency.
	PriceCurrency string `json:"price_currency,required"`
	// The recurring interval of the price, if type is `recurring`.
	RecurringInterval SubscriptionListResponseItemsPriceProductPriceRecurringRecurringInterval `json:"recurring_interval,required,nullable"`
	// The type of the price.
	Type SubscriptionListResponseItemsPriceProductPriceRecurringType `json:"type,required"`
	JSON subscriptionListResponseItemsPriceProductPriceRecurringJSON `json:"-"`
}

// subscriptionListResponseItemsPriceProductPriceRecurringJSON contains the JSON
// metadata for the struct
// [SubscriptionListResponseItemsPriceProductPriceRecurring]
type subscriptionListResponseItemsPriceProductPriceRecurringJSON struct {
	ID                apijson.Field
	CreatedAt         apijson.Field
	IsArchived        apijson.Field
	ModifiedAt        apijson.Field
	PriceAmount       apijson.Field
	PriceCurrency     apijson.Field
	RecurringInterval apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SubscriptionListResponseItemsPriceProductPriceRecurring) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionListResponseItemsPriceProductPriceRecurringJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionListResponseItemsPriceProductPriceRecurring) implementsSubscriptionListResponseItemsPrice() {
}

// The recurring interval of the price, if type is `recurring`.
type SubscriptionListResponseItemsPriceProductPriceRecurringRecurringInterval string

const (
	SubscriptionListResponseItemsPriceProductPriceRecurringRecurringIntervalMonth SubscriptionListResponseItemsPriceProductPriceRecurringRecurringInterval = "month"
	SubscriptionListResponseItemsPriceProductPriceRecurringRecurringIntervalYear  SubscriptionListResponseItemsPriceProductPriceRecurringRecurringInterval = "year"
)

func (r SubscriptionListResponseItemsPriceProductPriceRecurringRecurringInterval) IsKnown() bool {
	switch r {
	case SubscriptionListResponseItemsPriceProductPriceRecurringRecurringIntervalMonth, SubscriptionListResponseItemsPriceProductPriceRecurringRecurringIntervalYear:
		return true
	}
	return false
}

// The type of the price.
type SubscriptionListResponseItemsPriceProductPriceRecurringType string

const (
	SubscriptionListResponseItemsPriceProductPriceRecurringTypeRecurring SubscriptionListResponseItemsPriceProductPriceRecurringType = "recurring"
)

func (r SubscriptionListResponseItemsPriceProductPriceRecurringType) IsKnown() bool {
	switch r {
	case SubscriptionListResponseItemsPriceProductPriceRecurringTypeRecurring:
		return true
	}
	return false
}

// A one-time price for a product.
type SubscriptionListResponseItemsPriceProductPriceOneTime struct {
	// The ID of the price.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the price is archived and no longer available.
	IsArchived bool `json:"is_archived,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The price in cents.
	PriceAmount int64 `json:"price_amount,required"`
	// The currency.
	PriceCurrency string `json:"price_currency,required"`
	// The type of the price.
	Type SubscriptionListResponseItemsPriceProductPriceOneTimeType `json:"type,required"`
	JSON subscriptionListResponseItemsPriceProductPriceOneTimeJSON `json:"-"`
}

// subscriptionListResponseItemsPriceProductPriceOneTimeJSON contains the JSON
// metadata for the struct [SubscriptionListResponseItemsPriceProductPriceOneTime]
type subscriptionListResponseItemsPriceProductPriceOneTimeJSON struct {
	ID            apijson.Field
	CreatedAt     apijson.Field
	IsArchived    apijson.Field
	ModifiedAt    apijson.Field
	PriceAmount   apijson.Field
	PriceCurrency apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SubscriptionListResponseItemsPriceProductPriceOneTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionListResponseItemsPriceProductPriceOneTimeJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionListResponseItemsPriceProductPriceOneTime) implementsSubscriptionListResponseItemsPrice() {
}

// The type of the price.
type SubscriptionListResponseItemsPriceProductPriceOneTimeType string

const (
	SubscriptionListResponseItemsPriceProductPriceOneTimeTypeOneTime SubscriptionListResponseItemsPriceProductPriceOneTimeType = "one_time"
)

func (r SubscriptionListResponseItemsPriceProductPriceOneTimeType) IsKnown() bool {
	switch r {
	case SubscriptionListResponseItemsPriceProductPriceOneTimeTypeOneTime:
		return true
	}
	return false
}

// The type of the price.
type SubscriptionListResponseItemsPriceType string

const (
	SubscriptionListResponseItemsPriceTypeRecurring SubscriptionListResponseItemsPriceType = "recurring"
	SubscriptionListResponseItemsPriceTypeOneTime   SubscriptionListResponseItemsPriceType = "one_time"
)

func (r SubscriptionListResponseItemsPriceType) IsKnown() bool {
	switch r {
	case SubscriptionListResponseItemsPriceTypeRecurring, SubscriptionListResponseItemsPriceTypeOneTime:
		return true
	}
	return false
}

// The recurring interval of the price, if type is `recurring`.
type SubscriptionListResponseItemsPriceRecurringInterval string

const (
	SubscriptionListResponseItemsPriceRecurringIntervalMonth SubscriptionListResponseItemsPriceRecurringInterval = "month"
	SubscriptionListResponseItemsPriceRecurringIntervalYear  SubscriptionListResponseItemsPriceRecurringInterval = "year"
)

func (r SubscriptionListResponseItemsPriceRecurringInterval) IsKnown() bool {
	switch r {
	case SubscriptionListResponseItemsPriceRecurringIntervalMonth, SubscriptionListResponseItemsPriceRecurringIntervalYear:
		return true
	}
	return false
}

type SubscriptionListResponseItemsStatus string

const (
	SubscriptionListResponseItemsStatusIncomplete        SubscriptionListResponseItemsStatus = "incomplete"
	SubscriptionListResponseItemsStatusIncompleteExpired SubscriptionListResponseItemsStatus = "incomplete_expired"
	SubscriptionListResponseItemsStatusTrialing          SubscriptionListResponseItemsStatus = "trialing"
	SubscriptionListResponseItemsStatusActive            SubscriptionListResponseItemsStatus = "active"
	SubscriptionListResponseItemsStatusPastDue           SubscriptionListResponseItemsStatus = "past_due"
	SubscriptionListResponseItemsStatusCanceled          SubscriptionListResponseItemsStatus = "canceled"
	SubscriptionListResponseItemsStatusUnpaid            SubscriptionListResponseItemsStatus = "unpaid"
)

func (r SubscriptionListResponseItemsStatus) IsKnown() bool {
	switch r {
	case SubscriptionListResponseItemsStatusIncomplete, SubscriptionListResponseItemsStatusIncompleteExpired, SubscriptionListResponseItemsStatusTrialing, SubscriptionListResponseItemsStatusActive, SubscriptionListResponseItemsStatusPastDue, SubscriptionListResponseItemsStatusCanceled, SubscriptionListResponseItemsStatusUnpaid:
		return true
	}
	return false
}

type SubscriptionListResponseItemsUser struct {
	AvatarURL      string                                `json:"avatar_url,required,nullable"`
	Email          string                                `json:"email,required"`
	GitHubUsername string                                `json:"github_username,required,nullable"`
	PublicName     string                                `json:"public_name,required"`
	JSON           subscriptionListResponseItemsUserJSON `json:"-"`
}

// subscriptionListResponseItemsUserJSON contains the JSON metadata for the struct
// [SubscriptionListResponseItemsUser]
type subscriptionListResponseItemsUserJSON struct {
	AvatarURL      apijson.Field
	Email          apijson.Field
	GitHubUsername apijson.Field
	PublicName     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SubscriptionListResponseItemsUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionListResponseItemsUserJSON) RawJSON() string {
	return r.raw
}

type SubscriptionExportResponse = interface{}

// Result of a subscription import operation.
type SubscriptionImportResponse struct {
	Count int64                          `json:"count,required"`
	JSON  subscriptionImportResponseJSON `json:"-"`
}

// subscriptionImportResponseJSON contains the JSON metadata for the struct
// [SubscriptionImportResponse]
type subscriptionImportResponseJSON struct {
	Count       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionImportResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionImportResponseJSON) RawJSON() string {
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
	Sorting param.Field[[]string] `query:"sorting"`
	// Filter by subscription tier type.
	Type param.Field[SubscriptionListParamsTypeUnion] `query:"type"`
}

// URLQuery serializes [SubscriptionListParams]'s query parameters as `url.Values`.
func (r SubscriptionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
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

func (r SubscriptionListParamsOrganizationIDArray) ImplementsSubscriptionListParamsOrganizationIDUnion() {
}

// Filter by product ID.
//
// Satisfied by [shared.UnionString], [SubscriptionListParamsProductIDArray].
type SubscriptionListParamsProductIDUnion interface {
	ImplementsSubscriptionListParamsProductIDUnion()
}

type SubscriptionListParamsProductIDArray []string

func (r SubscriptionListParamsProductIDArray) ImplementsSubscriptionListParamsProductIDUnion() {}

// Filter by subscription tier type.
//
// Satisfied by [SubscriptionListParamsTypeSubscriptionTierType],
// [SubscriptionListParamsTypeArray].
type SubscriptionListParamsTypeUnion interface {
	implementsSubscriptionListParamsTypeUnion()
}

type SubscriptionListParamsTypeSubscriptionTierType string

const (
	SubscriptionListParamsTypeSubscriptionTierTypeFree       SubscriptionListParamsTypeSubscriptionTierType = "free"
	SubscriptionListParamsTypeSubscriptionTierTypeIndividual SubscriptionListParamsTypeSubscriptionTierType = "individual"
	SubscriptionListParamsTypeSubscriptionTierTypeBusiness   SubscriptionListParamsTypeSubscriptionTierType = "business"
)

func (r SubscriptionListParamsTypeSubscriptionTierType) IsKnown() bool {
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
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
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

func (r SubscriptionExportParamsOrganizationIDArray) ImplementsSubscriptionExportParamsOrganizationIDUnion() {
}

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
