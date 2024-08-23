// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package polar

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/polarsource/polar-go/internal/apijson"
	"github.com/polarsource/polar-go/internal/apiquery"
	"github.com/polarsource/polar-go/internal/param"
	"github.com/polarsource/polar-go/internal/requestconfig"
	"github.com/polarsource/polar-go/option"
)

// WebhookEndpointService contains methods and other services that help with
// interacting with the polar API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookEndpointService] method instead.
type WebhookEndpointService struct {
	Options []option.RequestOption
}

// NewWebhookEndpointService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWebhookEndpointService(opts ...option.RequestOption) (r *WebhookEndpointService) {
	r = &WebhookEndpointService{}
	r.Options = opts
	return
}

// Create a webhook endpoint.
func (r *WebhookEndpointService) New(ctx context.Context, body WebhookEndpointNewParams, opts ...option.RequestOption) (res *WebhookEndpoint, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/webhooks/endpoints"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a webhook endpoint by ID.
func (r *WebhookEndpointService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *WebhookEndpoint, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/webhooks/endpoints/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a webhook endpoint.
func (r *WebhookEndpointService) Update(ctx context.Context, id string, body WebhookEndpointUpdateParams, opts ...option.RequestOption) (res *WebhookEndpoint, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/webhooks/endpoints/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List webhook endpoints.
func (r *WebhookEndpointService) List(ctx context.Context, query WebhookEndpointListParams, opts ...option.RequestOption) (res *ListResourceWebhookEndpoint, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/webhooks/endpoints"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a webhook endpoint.
func (r *WebhookEndpointService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/webhooks/endpoints/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type ListResourceWebhookEndpoint struct {
	Pagination ListResourceWebhookEndpointPagination `json:"pagination,required"`
	Items      []WebhookEndpoint                     `json:"items"`
	JSON       listResourceWebhookEndpointJSON       `json:"-"`
}

// listResourceWebhookEndpointJSON contains the JSON metadata for the struct
// [ListResourceWebhookEndpoint]
type listResourceWebhookEndpointJSON struct {
	Pagination  apijson.Field
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListResourceWebhookEndpoint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResourceWebhookEndpointJSON) RawJSON() string {
	return r.raw
}

type ListResourceWebhookEndpointPagination struct {
	MaxPage    int64                                     `json:"max_page,required"`
	TotalCount int64                                     `json:"total_count,required"`
	JSON       listResourceWebhookEndpointPaginationJSON `json:"-"`
}

// listResourceWebhookEndpointPaginationJSON contains the JSON metadata for the
// struct [ListResourceWebhookEndpointPagination]
type listResourceWebhookEndpointPaginationJSON struct {
	MaxPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListResourceWebhookEndpointPagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResourceWebhookEndpointPaginationJSON) RawJSON() string {
	return r.raw
}

// A webhook endpoint.
type WebhookEndpoint struct {
	// The webhook endpoint ID.
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The events that will trigger the webhook.
	Events []WebhookEndpointEvent `json:"events,required"`
	// The format of the webhook payload.
	Format WebhookEndpointFormat `json:"format,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,required,nullable" format:"date-time"`
	// The URL where the webhook events will be sent.
	URL string `json:"url,required"`
	// The organization ID associated with the webhook endpoint.
	OrganizationID string `json:"organization_id,nullable" format:"uuid4"`
	// The user ID associated with the webhook endpoint.
	UserID string              `json:"user_id,nullable" format:"uuid4"`
	JSON   webhookEndpointJSON `json:"-"`
}

// webhookEndpointJSON contains the JSON metadata for the struct [WebhookEndpoint]
type webhookEndpointJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Events         apijson.Field
	Format         apijson.Field
	ModifiedAt     apijson.Field
	URL            apijson.Field
	OrganizationID apijson.Field
	UserID         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *WebhookEndpoint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookEndpointJSON) RawJSON() string {
	return r.raw
}

type WebhookEndpointEvent string

const (
	WebhookEndpointEventOrderCreated        WebhookEndpointEvent = "order.created"
	WebhookEndpointEventSubscriptionCreated WebhookEndpointEvent = "subscription.created"
	WebhookEndpointEventSubscriptionUpdated WebhookEndpointEvent = "subscription.updated"
	WebhookEndpointEventProductCreated      WebhookEndpointEvent = "product.created"
	WebhookEndpointEventProductUpdated      WebhookEndpointEvent = "product.updated"
	WebhookEndpointEventBenefitCreated      WebhookEndpointEvent = "benefit.created"
	WebhookEndpointEventBenefitUpdated      WebhookEndpointEvent = "benefit.updated"
	WebhookEndpointEventOrganizationUpdated WebhookEndpointEvent = "organization.updated"
	WebhookEndpointEventPledgeCreated       WebhookEndpointEvent = "pledge.created"
	WebhookEndpointEventPledgeUpdated       WebhookEndpointEvent = "pledge.updated"
	WebhookEndpointEventDonationCreated     WebhookEndpointEvent = "donation.created"
)

func (r WebhookEndpointEvent) IsKnown() bool {
	switch r {
	case WebhookEndpointEventOrderCreated, WebhookEndpointEventSubscriptionCreated, WebhookEndpointEventSubscriptionUpdated, WebhookEndpointEventProductCreated, WebhookEndpointEventProductUpdated, WebhookEndpointEventBenefitCreated, WebhookEndpointEventBenefitUpdated, WebhookEndpointEventOrganizationUpdated, WebhookEndpointEventPledgeCreated, WebhookEndpointEventPledgeUpdated, WebhookEndpointEventDonationCreated:
		return true
	}
	return false
}

// The format of the webhook payload.
type WebhookEndpointFormat string

const (
	WebhookEndpointFormatRaw     WebhookEndpointFormat = "raw"
	WebhookEndpointFormatDiscord WebhookEndpointFormat = "discord"
	WebhookEndpointFormatSlack   WebhookEndpointFormat = "slack"
)

func (r WebhookEndpointFormat) IsKnown() bool {
	switch r {
	case WebhookEndpointFormatRaw, WebhookEndpointFormatDiscord, WebhookEndpointFormatSlack:
		return true
	}
	return false
}

type WebhookEndpointNewParams struct {
	// The events that will trigger the webhook.
	Events param.Field[[]WebhookEndpointNewParamsEvent] `json:"events,required"`
	// The format of the webhook payload.
	Format param.Field[WebhookEndpointNewParamsFormat] `json:"format,required"`
	// The secret used to sign the webhook events.
	Secret param.Field[string] `json:"secret,required"`
	// The URL where the webhook events will be sent.
	URL param.Field[string] `json:"url,required" format:"uri"`
	// The organization ID.
	OrganizationID param.Field[string] `json:"organization_id" format:"uuid4"`
}

func (r WebhookEndpointNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WebhookEndpointNewParamsEvent string

const (
	WebhookEndpointNewParamsEventOrderCreated        WebhookEndpointNewParamsEvent = "order.created"
	WebhookEndpointNewParamsEventSubscriptionCreated WebhookEndpointNewParamsEvent = "subscription.created"
	WebhookEndpointNewParamsEventSubscriptionUpdated WebhookEndpointNewParamsEvent = "subscription.updated"
	WebhookEndpointNewParamsEventProductCreated      WebhookEndpointNewParamsEvent = "product.created"
	WebhookEndpointNewParamsEventProductUpdated      WebhookEndpointNewParamsEvent = "product.updated"
	WebhookEndpointNewParamsEventBenefitCreated      WebhookEndpointNewParamsEvent = "benefit.created"
	WebhookEndpointNewParamsEventBenefitUpdated      WebhookEndpointNewParamsEvent = "benefit.updated"
	WebhookEndpointNewParamsEventOrganizationUpdated WebhookEndpointNewParamsEvent = "organization.updated"
	WebhookEndpointNewParamsEventPledgeCreated       WebhookEndpointNewParamsEvent = "pledge.created"
	WebhookEndpointNewParamsEventPledgeUpdated       WebhookEndpointNewParamsEvent = "pledge.updated"
	WebhookEndpointNewParamsEventDonationCreated     WebhookEndpointNewParamsEvent = "donation.created"
)

func (r WebhookEndpointNewParamsEvent) IsKnown() bool {
	switch r {
	case WebhookEndpointNewParamsEventOrderCreated, WebhookEndpointNewParamsEventSubscriptionCreated, WebhookEndpointNewParamsEventSubscriptionUpdated, WebhookEndpointNewParamsEventProductCreated, WebhookEndpointNewParamsEventProductUpdated, WebhookEndpointNewParamsEventBenefitCreated, WebhookEndpointNewParamsEventBenefitUpdated, WebhookEndpointNewParamsEventOrganizationUpdated, WebhookEndpointNewParamsEventPledgeCreated, WebhookEndpointNewParamsEventPledgeUpdated, WebhookEndpointNewParamsEventDonationCreated:
		return true
	}
	return false
}

// The format of the webhook payload.
type WebhookEndpointNewParamsFormat string

const (
	WebhookEndpointNewParamsFormatRaw     WebhookEndpointNewParamsFormat = "raw"
	WebhookEndpointNewParamsFormatDiscord WebhookEndpointNewParamsFormat = "discord"
	WebhookEndpointNewParamsFormatSlack   WebhookEndpointNewParamsFormat = "slack"
)

func (r WebhookEndpointNewParamsFormat) IsKnown() bool {
	switch r {
	case WebhookEndpointNewParamsFormatRaw, WebhookEndpointNewParamsFormatDiscord, WebhookEndpointNewParamsFormatSlack:
		return true
	}
	return false
}

type WebhookEndpointUpdateParams struct {
	// The events that will trigger the webhook.
	Events param.Field[[]WebhookEndpointUpdateParamsEvent] `json:"events"`
	// The format of the webhook payload.
	Format param.Field[WebhookEndpointUpdateParamsFormat] `json:"format"`
	// The secret used to sign the webhook events.
	Secret param.Field[string] `json:"secret"`
	// The URL where the webhook events will be sent.
	URL param.Field[string] `json:"url" format:"uri"`
}

func (r WebhookEndpointUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WebhookEndpointUpdateParamsEvent string

const (
	WebhookEndpointUpdateParamsEventOrderCreated        WebhookEndpointUpdateParamsEvent = "order.created"
	WebhookEndpointUpdateParamsEventSubscriptionCreated WebhookEndpointUpdateParamsEvent = "subscription.created"
	WebhookEndpointUpdateParamsEventSubscriptionUpdated WebhookEndpointUpdateParamsEvent = "subscription.updated"
	WebhookEndpointUpdateParamsEventProductCreated      WebhookEndpointUpdateParamsEvent = "product.created"
	WebhookEndpointUpdateParamsEventProductUpdated      WebhookEndpointUpdateParamsEvent = "product.updated"
	WebhookEndpointUpdateParamsEventBenefitCreated      WebhookEndpointUpdateParamsEvent = "benefit.created"
	WebhookEndpointUpdateParamsEventBenefitUpdated      WebhookEndpointUpdateParamsEvent = "benefit.updated"
	WebhookEndpointUpdateParamsEventOrganizationUpdated WebhookEndpointUpdateParamsEvent = "organization.updated"
	WebhookEndpointUpdateParamsEventPledgeCreated       WebhookEndpointUpdateParamsEvent = "pledge.created"
	WebhookEndpointUpdateParamsEventPledgeUpdated       WebhookEndpointUpdateParamsEvent = "pledge.updated"
	WebhookEndpointUpdateParamsEventDonationCreated     WebhookEndpointUpdateParamsEvent = "donation.created"
)

func (r WebhookEndpointUpdateParamsEvent) IsKnown() bool {
	switch r {
	case WebhookEndpointUpdateParamsEventOrderCreated, WebhookEndpointUpdateParamsEventSubscriptionCreated, WebhookEndpointUpdateParamsEventSubscriptionUpdated, WebhookEndpointUpdateParamsEventProductCreated, WebhookEndpointUpdateParamsEventProductUpdated, WebhookEndpointUpdateParamsEventBenefitCreated, WebhookEndpointUpdateParamsEventBenefitUpdated, WebhookEndpointUpdateParamsEventOrganizationUpdated, WebhookEndpointUpdateParamsEventPledgeCreated, WebhookEndpointUpdateParamsEventPledgeUpdated, WebhookEndpointUpdateParamsEventDonationCreated:
		return true
	}
	return false
}

// The format of the webhook payload.
type WebhookEndpointUpdateParamsFormat string

const (
	WebhookEndpointUpdateParamsFormatRaw     WebhookEndpointUpdateParamsFormat = "raw"
	WebhookEndpointUpdateParamsFormatDiscord WebhookEndpointUpdateParamsFormat = "discord"
	WebhookEndpointUpdateParamsFormatSlack   WebhookEndpointUpdateParamsFormat = "slack"
)

func (r WebhookEndpointUpdateParamsFormat) IsKnown() bool {
	switch r {
	case WebhookEndpointUpdateParamsFormatRaw, WebhookEndpointUpdateParamsFormatDiscord, WebhookEndpointUpdateParamsFormatSlack:
		return true
	}
	return false
}

type WebhookEndpointListParams struct {
	// Size of a page, defaults to 10. Maximum is 100.
	Limit param.Field[int64] `query:"limit"`
	// The organization ID.
	OrganizationID param.Field[string] `query:"organization_id" format:"uuid4"`
	// Page number, defaults to 1.
	Page param.Field[int64] `query:"page"`
	// Filter by user ID.
	UserID param.Field[string] `query:"user_id" format:"uuid4"`
}

// URLQuery serializes [WebhookEndpointListParams]'s query parameters as
// `url.Values`.
func (r WebhookEndpointListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
