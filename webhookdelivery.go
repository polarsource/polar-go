// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package polar

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/polarsource/polar-go/internal/apijson"
	"github.com/polarsource/polar-go/internal/apiquery"
	"github.com/polarsource/polar-go/internal/param"
	"github.com/polarsource/polar-go/internal/requestconfig"
	"github.com/polarsource/polar-go/option"
)

// WebhookDeliveryService contains methods and other services that help with
// interacting with the polar API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookDeliveryService] method instead.
type WebhookDeliveryService struct {
	Options []option.RequestOption
}

// NewWebhookDeliveryService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWebhookDeliveryService(opts ...option.RequestOption) (r *WebhookDeliveryService) {
	r = &WebhookDeliveryService{}
	r.Options = opts
	return
}

// List webhook deliveries.
//
// Deliveries are all the attempts to deliver a webhook event to an endpoint.
func (r *WebhookDeliveryService) List(ctx context.Context, query WebhookDeliveryListParams, opts ...option.RequestOption) (res *ListResourceWebhookDelivery, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/webhooks/deliveries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ListResourceWebhookDelivery struct {
	Pagination ListResourceWebhookDeliveryPagination `json:"pagination,required"`
	Items      []ListResourceWebhookDeliveryItem     `json:"items"`
	JSON       listResourceWebhookDeliveryJSON       `json:"-"`
}

// listResourceWebhookDeliveryJSON contains the JSON metadata for the struct
// [ListResourceWebhookDelivery]
type listResourceWebhookDeliveryJSON struct {
	Pagination  apijson.Field
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListResourceWebhookDelivery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResourceWebhookDeliveryJSON) RawJSON() string {
	return r.raw
}

type ListResourceWebhookDeliveryPagination struct {
	MaxPage    int64                                     `json:"max_page,required"`
	TotalCount int64                                     `json:"total_count,required"`
	JSON       listResourceWebhookDeliveryPaginationJSON `json:"-"`
}

// listResourceWebhookDeliveryPaginationJSON contains the JSON metadata for the
// struct [ListResourceWebhookDeliveryPagination]
type listResourceWebhookDeliveryPaginationJSON struct {
	MaxPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListResourceWebhookDeliveryPagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResourceWebhookDeliveryPaginationJSON) RawJSON() string {
	return r.raw
}

// A webhook delivery for a webhook event.
type ListResourceWebhookDeliveryItem struct {
	// The webhook delivery ID.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the delivery was successful.
	Succeeded bool `json:"succeeded,required"`
	// A webhook event.
	//
	// An event represent something that happened in the system that should be sent to
	// the webhook endpoint.
	//
	// It can be delivered multiple times until it's marked as succeeded, each one
	// creating a new delivery.
	WebhookEvent ListResourceWebhookDeliveryItemsWebhookEvent `json:"webhook_event,required"`
	// The HTTP code returned by the URL. `null` if the endpoint was unreachable.
	HTTPCode int64 `json:"http_code,nullable"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time                           `json:"modified_at,nullable" format:"date-time"`
	JSON       listResourceWebhookDeliveryItemJSON `json:"-"`
}

// listResourceWebhookDeliveryItemJSON contains the JSON metadata for the struct
// [ListResourceWebhookDeliveryItem]
type listResourceWebhookDeliveryItemJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	Succeeded    apijson.Field
	WebhookEvent apijson.Field
	HTTPCode     apijson.Field
	ModifiedAt   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ListResourceWebhookDeliveryItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResourceWebhookDeliveryItemJSON) RawJSON() string {
	return r.raw
}

// A webhook event.
//
// An event represent something that happened in the system that should be sent to
// the webhook endpoint.
//
// It can be delivered multiple times until it's marked as succeeded, each one
// creating a new delivery.
type ListResourceWebhookDeliveryItemsWebhookEvent struct {
	// The webhook event ID.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The payload of the webhook event.
	Payload string `json:"payload,required"`
	// Last HTTP code returned by the URL. `null` if no delviery has been attempted or
	// if the endpoint was unreachable.
	LastHTTPCode int64 `json:"last_http_code,nullable"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,nullable" format:"date-time"`
	// Whether this event was successfully delivered. `null` if no delivery has been
	// attempted.
	Succeeded bool                                             `json:"succeeded,nullable"`
	JSON      listResourceWebhookDeliveryItemsWebhookEventJSON `json:"-"`
}

// listResourceWebhookDeliveryItemsWebhookEventJSON contains the JSON metadata for
// the struct [ListResourceWebhookDeliveryItemsWebhookEvent]
type listResourceWebhookDeliveryItemsWebhookEventJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	Payload      apijson.Field
	LastHTTPCode apijson.Field
	ModifiedAt   apijson.Field
	Succeeded    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ListResourceWebhookDeliveryItemsWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResourceWebhookDeliveryItemsWebhookEventJSON) RawJSON() string {
	return r.raw
}

type WebhookDeliveryListParams struct {
	// Filter by webhook endpoint ID.
	EndpointID param.Field[string] `query:"endpoint_id" format:"uuid4"`
	// Size of a page, defaults to 10. Maximum is 100.
	Limit param.Field[int64] `query:"limit"`
	// Page number, defaults to 1.
	Page param.Field[int64] `query:"page"`
}

// URLQuery serializes [WebhookDeliveryListParams]'s query parameters as
// `url.Values`.
func (r WebhookDeliveryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
