// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package polar

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/polarsource/polar-go/internal/apijson"
	"github.com/polarsource/polar-go/internal/param"
	"github.com/polarsource/polar-go/internal/requestconfig"
	"github.com/polarsource/polar-go/option"
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
	ID            string `json:"id,required"`
	CustomerEmail string `json:"customer_email,required,nullable"`
	CustomerName  string `json:"customer_name,required,nullable"`
	// A product.
	Product Product `json:"product,required"`
	// A recurring price for a product, i.e. a subscription.
	ProductPrice ProductPrice `json:"product_price,required"`
	// URL the customer should be redirected to complete the purchase.
	URL  string       `json:"url,nullable"`
	JSON checkoutJSON `json:"-"`
}

// checkoutJSON contains the JSON metadata for the struct [Checkout]
type checkoutJSON struct {
	ID            apijson.Field
	CustomerEmail apijson.Field
	CustomerName  apijson.Field
	Product       apijson.Field
	ProductPrice  apijson.Field
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
