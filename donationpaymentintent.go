// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package polar

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/polar-go/internal/apijson"
	"github.com/stainless-sdks/polar-go/internal/param"
	"github.com/stainless-sdks/polar-go/internal/requestconfig"
	"github.com/stainless-sdks/polar-go/option"
)

// DonationPaymentIntentService contains methods and other services that help with
// interacting with the polar API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDonationPaymentIntentService] method instead.
type DonationPaymentIntentService struct {
	Options []option.RequestOption
}

// NewDonationPaymentIntentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDonationPaymentIntentService(opts ...option.RequestOption) (r *DonationPaymentIntentService) {
	r = &DonationPaymentIntentService{}
	r.Options = opts
	return
}

// Create Payment Intent
func (r *DonationPaymentIntentService) New(ctx context.Context, body DonationPaymentIntentNewParams, opts ...option.RequestOption) (res *DonationStripePaymentIntentMutationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/donations/payment_intent"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update Payment Intent
func (r *DonationPaymentIntentService) Update(ctx context.Context, id string, body DonationPaymentIntentUpdateParams, opts ...option.RequestOption) (res *DonationStripePaymentIntentMutationResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/donations/payment_intent/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type DonationStripePaymentIntentMutationResponse struct {
	Amount          int64                                           `json:"amount,required"`
	Currency        string                                          `json:"currency,required"`
	PaymentIntentID string                                          `json:"payment_intent_id,required"`
	ClientSecret    string                                          `json:"client_secret,nullable"`
	JSON            donationStripePaymentIntentMutationResponseJSON `json:"-"`
}

// donationStripePaymentIntentMutationResponseJSON contains the JSON metadata for
// the struct [DonationStripePaymentIntentMutationResponse]
type donationStripePaymentIntentMutationResponseJSON struct {
	Amount          apijson.Field
	Currency        apijson.Field
	PaymentIntentID apijson.Field
	ClientSecret    apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DonationStripePaymentIntentMutationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r donationStripePaymentIntentMutationResponseJSON) RawJSON() string {
	return r.raw
}

type DonationPaymentIntentNewParams struct {
	// The amount in cents.
	Amount param.Field[int64] `json:"amount,required"`
	// The donators email address. Receipts will be sent to this address.
	Email            param.Field[string] `json:"email,required" format:"email"`
	ToOrganizationID param.Field[string] `json:"to_organization_id,required" format:"uuid4"`
	// The currency. Currently, only `usd` is supported.
	Currency param.Field[string] `json:"currency"`
	IssueID  param.Field[string] `json:"issue_id" format:"uuid4"`
	// Message included with the donation
	Message param.Field[string] `json:"message"`
	// The organization to give credit to. The pledge will be paid by the authenticated
	// user.
	OnBehalfOfOrganizationID param.Field[string] `json:"on_behalf_of_organization_id" format:"uuid4"`
	// If the payment method should be saved for future usage.
	SetupFutureUsage param.Field[DonationPaymentIntentNewParamsSetupFutureUsage] `json:"setup_future_usage"`
}

func (r DonationPaymentIntentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If the payment method should be saved for future usage.
type DonationPaymentIntentNewParamsSetupFutureUsage string

const (
	DonationPaymentIntentNewParamsSetupFutureUsageOnSession DonationPaymentIntentNewParamsSetupFutureUsage = "on_session"
)

func (r DonationPaymentIntentNewParamsSetupFutureUsage) IsKnown() bool {
	switch r {
	case DonationPaymentIntentNewParamsSetupFutureUsageOnSession:
		return true
	}
	return false
}

type DonationPaymentIntentUpdateParams struct {
	// The amount in cents.
	Amount param.Field[int64] `json:"amount,required"`
	// The donators email address. Receipts will be sent to this address.
	Email param.Field[string] `json:"email,required"`
	// The currency. Currently, only `usd` is supported.
	Currency param.Field[string] `json:"currency"`
	IssueID  param.Field[string] `json:"issue_id" format:"uuid4"`
	// Message included with the donation
	Message param.Field[string] `json:"message"`
	// The organization to give credit to. The pledge will be paid by the authenticated
	// user.
	OnBehalfOfOrganizationID param.Field[string] `json:"on_behalf_of_organization_id" format:"uuid4"`
	// If the payment method should be saved for future usage.
	SetupFutureUsage param.Field[DonationPaymentIntentUpdateParamsSetupFutureUsage] `json:"setup_future_usage"`
}

func (r DonationPaymentIntentUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If the payment method should be saved for future usage.
type DonationPaymentIntentUpdateParamsSetupFutureUsage string

const (
	DonationPaymentIntentUpdateParamsSetupFutureUsageOnSession DonationPaymentIntentUpdateParamsSetupFutureUsage = "on_session"
)

func (r DonationPaymentIntentUpdateParamsSetupFutureUsage) IsKnown() bool {
	switch r {
	case DonationPaymentIntentUpdateParamsSetupFutureUsageOnSession:
		return true
	}
	return false
}
