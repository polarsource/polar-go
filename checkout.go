// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package polar

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/stainless-sdks/polar-go/internal/apijson"
	"github.com/stainless-sdks/polar-go/internal/param"
	"github.com/stainless-sdks/polar-go/internal/requestconfig"
	"github.com/stainless-sdks/polar-go/option"
	"github.com/tidwall/gjson"
)

// CheckoutService contains methods and other services that help with interacting
// with the polar API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCheckoutService] method instead.
type CheckoutService struct {
	Options []option.RequestOption
}

// NewCheckoutService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCheckoutService(opts ...option.RequestOption) (r *CheckoutService) {
	r = &CheckoutService{}
	r.Options = opts
	return
}

// Create a checkout session.
func (r *CheckoutService) New(ctx context.Context, body CheckoutNewParams, opts ...option.RequestOption) (res *Checkout, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/checkouts/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get an active checkout session by ID.
func (r *CheckoutService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Checkout, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/checkouts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// A checkout session.
type Checkout struct {
	// The ID of the checkout.
	ID string `json:"id,required"`
	// A product.
	Product ProductOutput `json:"product,required"`
	// A recurring price for a product, i.e. a subscription.
	ProductPrice  CheckoutProductPrice `json:"product_price,required"`
	CustomerEmail string               `json:"customer_email,nullable"`
	CustomerName  string               `json:"customer_name,nullable"`
	// URL the customer should be redirected to complete the purchase.
	URL  string       `json:"url,nullable"`
	JSON checkoutJSON `json:"-"`
}

// checkoutJSON contains the JSON metadata for the struct [Checkout]
type checkoutJSON struct {
	ID            apijson.Field
	Product       apijson.Field
	ProductPrice  apijson.Field
	CustomerEmail apijson.Field
	CustomerName  apijson.Field
	URL           apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *Checkout) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkoutJSON) RawJSON() string {
	return r.raw
}

// A recurring price for a product, i.e. a subscription.
type CheckoutProductPrice struct {
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
	Type CheckoutProductPriceType `json:"type,required"`
	// The recurring interval of the price, if type is `recurring`.
	RecurringInterval CheckoutProductPriceRecurringInterval `json:"recurring_interval,nullable"`
	JSON              checkoutProductPriceJSON              `json:"-"`
	union             CheckoutProductPriceUnion
}

// checkoutProductPriceJSON contains the JSON metadata for the struct
// [CheckoutProductPrice]
type checkoutProductPriceJSON struct {
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

func (r checkoutProductPriceJSON) RawJSON() string {
	return r.raw
}

func (r *CheckoutProductPrice) UnmarshalJSON(data []byte) (err error) {
	*r = CheckoutProductPrice{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [CheckoutProductPriceUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are
// [CheckoutProductPriceProductPriceRecurring],
// [CheckoutProductPriceProductPriceOneTime].
func (r CheckoutProductPrice) AsUnion() CheckoutProductPriceUnion {
	return r.union
}

// A recurring price for a product, i.e. a subscription.
//
// Union satisfied by [CheckoutProductPriceProductPriceRecurring] or
// [CheckoutProductPriceProductPriceOneTime].
type CheckoutProductPriceUnion interface {
	implementsCheckoutProductPrice()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CheckoutProductPriceUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CheckoutProductPriceProductPriceRecurring{}),
			DiscriminatorValue: "recurring",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CheckoutProductPriceProductPriceOneTime{}),
			DiscriminatorValue: "one_time",
		},
	)
}

// A recurring price for a product, i.e. a subscription.
type CheckoutProductPriceProductPriceRecurring struct {
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
	Type CheckoutProductPriceProductPriceRecurringType `json:"type,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,nullable" format:"date-time"`
	// The recurring interval of the price, if type is `recurring`.
	RecurringInterval CheckoutProductPriceProductPriceRecurringRecurringInterval `json:"recurring_interval,nullable"`
	JSON              checkoutProductPriceProductPriceRecurringJSON              `json:"-"`
}

// checkoutProductPriceProductPriceRecurringJSON contains the JSON metadata for the
// struct [CheckoutProductPriceProductPriceRecurring]
type checkoutProductPriceProductPriceRecurringJSON struct {
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

func (r *CheckoutProductPriceProductPriceRecurring) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkoutProductPriceProductPriceRecurringJSON) RawJSON() string {
	return r.raw
}

func (r CheckoutProductPriceProductPriceRecurring) implementsCheckoutProductPrice() {}

// The type of the price.
type CheckoutProductPriceProductPriceRecurringType string

const (
	CheckoutProductPriceProductPriceRecurringTypeRecurring CheckoutProductPriceProductPriceRecurringType = "recurring"
)

func (r CheckoutProductPriceProductPriceRecurringType) IsKnown() bool {
	switch r {
	case CheckoutProductPriceProductPriceRecurringTypeRecurring:
		return true
	}
	return false
}

// The recurring interval of the price, if type is `recurring`.
type CheckoutProductPriceProductPriceRecurringRecurringInterval string

const (
	CheckoutProductPriceProductPriceRecurringRecurringIntervalMonth CheckoutProductPriceProductPriceRecurringRecurringInterval = "month"
	CheckoutProductPriceProductPriceRecurringRecurringIntervalYear  CheckoutProductPriceProductPriceRecurringRecurringInterval = "year"
)

func (r CheckoutProductPriceProductPriceRecurringRecurringInterval) IsKnown() bool {
	switch r {
	case CheckoutProductPriceProductPriceRecurringRecurringIntervalMonth, CheckoutProductPriceProductPriceRecurringRecurringIntervalYear:
		return true
	}
	return false
}

// A one-time price for a product.
type CheckoutProductPriceProductPriceOneTime struct {
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
	Type CheckoutProductPriceProductPriceOneTimeType `json:"type,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time                                   `json:"modified_at,nullable" format:"date-time"`
	JSON       checkoutProductPriceProductPriceOneTimeJSON `json:"-"`
}

// checkoutProductPriceProductPriceOneTimeJSON contains the JSON metadata for the
// struct [CheckoutProductPriceProductPriceOneTime]
type checkoutProductPriceProductPriceOneTimeJSON struct {
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

func (r *CheckoutProductPriceProductPriceOneTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkoutProductPriceProductPriceOneTimeJSON) RawJSON() string {
	return r.raw
}

func (r CheckoutProductPriceProductPriceOneTime) implementsCheckoutProductPrice() {}

// The type of the price.
type CheckoutProductPriceProductPriceOneTimeType string

const (
	CheckoutProductPriceProductPriceOneTimeTypeOneTime CheckoutProductPriceProductPriceOneTimeType = "one_time"
)

func (r CheckoutProductPriceProductPriceOneTimeType) IsKnown() bool {
	switch r {
	case CheckoutProductPriceProductPriceOneTimeTypeOneTime:
		return true
	}
	return false
}

// The type of the price.
type CheckoutProductPriceType string

const (
	CheckoutProductPriceTypeRecurring CheckoutProductPriceType = "recurring"
	CheckoutProductPriceTypeOneTime   CheckoutProductPriceType = "one_time"
)

func (r CheckoutProductPriceType) IsKnown() bool {
	switch r {
	case CheckoutProductPriceTypeRecurring, CheckoutProductPriceTypeOneTime:
		return true
	}
	return false
}

// The recurring interval of the price, if type is `recurring`.
type CheckoutProductPriceRecurringInterval string

const (
	CheckoutProductPriceRecurringIntervalMonth CheckoutProductPriceRecurringInterval = "month"
	CheckoutProductPriceRecurringIntervalYear  CheckoutProductPriceRecurringInterval = "year"
)

func (r CheckoutProductPriceRecurringInterval) IsKnown() bool {
	switch r {
	case CheckoutProductPriceRecurringIntervalMonth, CheckoutProductPriceRecurringIntervalYear:
		return true
	}
	return false
}

type CheckoutNewParams struct {
	// ID of the product price to subscribe to.
	ProductPriceID param.Field[string] `json:"product_price_id,required" format:"uuid4"`
	// URL where the customer will be redirected after a successful subscription. You
	// can add the `session_id={CHECKOUT_SESSION_ID}` query parameter to retrieve the
	// checkout session id.
	SuccessURL param.Field[string] `json:"success_url,required" format:"uri"`
	// If you already know the email of your customer, you can set it. It'll be
	// pre-filled on the checkout page.
	CustomerEmail param.Field[string] `json:"customer_email" format:"email"`
}

func (r CheckoutNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
