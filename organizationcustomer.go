// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package polar

import (
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
