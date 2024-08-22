// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package polar

import (
	"github.com/polarsource/polar-go/option"
)

// UserService contains methods and other services that help with interacting with
// the polar API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserService] method instead.
type UserService struct {
	Options        []option.RequestOption
	Benefits       *UserBenefitService
	Orders         *UserOrderService
	Subscriptions  *UserSubscriptionService
	Advertisements *UserAdvertisementService
	Downloadables  *UserDownloadableService
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r *UserService) {
	r = &UserService{}
	r.Options = opts
	r.Benefits = NewUserBenefitService(opts...)
	r.Orders = NewUserOrderService(opts...)
	r.Subscriptions = NewUserSubscriptionService(opts...)
	r.Advertisements = NewUserAdvertisementService(opts...)
	r.Downloadables = NewUserDownloadableService(opts...)
	return
}
