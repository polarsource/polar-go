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

// ProductBenefitService contains methods and other services that help with
// interacting with the polar API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProductBenefitService] method instead.
type ProductBenefitService struct {
	Options []option.RequestOption
}

// NewProductBenefitService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewProductBenefitService(opts ...option.RequestOption) (r *ProductBenefitService) {
	r = &ProductBenefitService{}
	r.Options = opts
	return
}

// Update benefits granted by a product.
func (r *ProductBenefitService) Update(ctx context.Context, id string, body ProductBenefitUpdateParams, opts ...option.RequestOption) (res *Product, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/products/%s/benefits", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ProductBenefitUpdateParams struct {
	// List of benefit IDs. Each one must be on the same organization as the product.
	Benefits param.Field[[]string] `json:"benefits,required" format:"uuid4"`
}

func (r ProductBenefitUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
