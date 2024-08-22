// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package polar

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/polar-go/internal/apijson"
	"github.com/stainless-sdks/polar-go/internal/requestconfig"
	"github.com/stainless-sdks/polar-go/option"
)

// OrderInvoiceService contains methods and other services that help with
// interacting with the polar API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrderInvoiceService] method instead.
type OrderInvoiceService struct {
	Options []option.RequestOption
}

// NewOrderInvoiceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOrderInvoiceService(opts ...option.RequestOption) (r *OrderInvoiceService) {
	r = &OrderInvoiceService{}
	r.Options = opts
	return
}

// Get an order's invoice data.
func (r *OrderInvoiceService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *OrderInvoice, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/orders/%s/invoice", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Order's invoice data.
type OrderInvoice struct {
	// The URL to the invoice.
	URL  string           `json:"url,required"`
	JSON orderInvoiceJSON `json:"-"`
}

// orderInvoiceJSON contains the JSON metadata for the struct [OrderInvoice]
type orderInvoiceJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrderInvoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orderInvoiceJSON) RawJSON() string {
	return r.raw
}
