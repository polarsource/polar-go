// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package polar

import (
	"github.com/stainless-sdks/polar-go/option"
)

// WebhookService contains methods and other services that help with interacting
// with the polar API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookService] method instead.
type WebhookService struct {
	Options    []option.RequestOption
	Endpoints  *WebhookEndpointService
	Deliveries *WebhookDeliveryService
	Events     *WebhookEventService
}

// NewWebhookService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWebhookService(opts ...option.RequestOption) (r *WebhookService) {
	r = &WebhookService{}
	r.Options = opts
	r.Endpoints = NewWebhookEndpointService(opts...)
	r.Deliveries = NewWebhookDeliveryService(opts...)
	r.Events = NewWebhookEventService(opts...)
	return
}
