// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package polar

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/stainless-sdks/polar-go/internal/apijson"
	"github.com/stainless-sdks/polar-go/internal/apiquery"
	"github.com/stainless-sdks/polar-go/internal/param"
	"github.com/stainless-sdks/polar-go/internal/requestconfig"
	"github.com/stainless-sdks/polar-go/option"
	"github.com/tidwall/gjson"
)

// TransactionPayoutService contains methods and other services that help with
// interacting with the polar API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransactionPayoutService] method instead.
type TransactionPayoutService struct {
	Options []option.RequestOption
}

// NewTransactionPayoutService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTransactionPayoutService(opts ...option.RequestOption) (r *TransactionPayoutService) {
	r = &TransactionPayoutService{}
	r.Options = opts
	return
}

// Create Payout
func (r *TransactionPayoutService) New(ctx context.Context, body TransactionPayoutNewParams, opts ...option.RequestOption) (res *Transaction, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/transactions/payouts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get Payout Estimate
func (r *TransactionPayoutService) List(ctx context.Context, query TransactionPayoutListParams, opts ...option.RequestOption) (res *PayoutEstimate, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/transactions/payouts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get Payout Csv
func (r *TransactionPayoutService) Csv(ctx context.Context, id string, opts ...option.RequestOption) (res *TransactionPayoutCsvResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/transactions/payouts/%s/csv", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type PayoutEstimate struct {
	AccountID   string             `json:"account_id,required" format:"uuid4"`
	FeesAmount  int64              `json:"fees_amount,required"`
	GrossAmount int64              `json:"gross_amount,required"`
	NetAmount   int64              `json:"net_amount,required"`
	JSON        payoutEstimateJSON `json:"-"`
}

// payoutEstimateJSON contains the JSON metadata for the struct [PayoutEstimate]
type payoutEstimateJSON struct {
	AccountID   apijson.Field
	FeesAmount  apijson.Field
	GrossAmount apijson.Field
	NetAmount   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PayoutEstimate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r payoutEstimateJSON) RawJSON() string {
	return r.raw
}

type Transaction struct {
	ID                          string                                  `json:"id,required" format:"uuid4"`
	AccountAmount               int64                                   `json:"account_amount,required"`
	AccountCurrency             string                                  `json:"account_currency,required"`
	AccountIncurredTransactions []TransactionAccountIncurredTransaction `json:"account_incurred_transactions,required"`
	Amount                      int64                                   `json:"amount,required"`
	// Creation timestamp of the object.
	CreatedAt      time.Time `json:"created_at,required" format:"date-time"`
	Currency       string    `json:"currency,required"`
	GrossAmount    int64     `json:"gross_amount,required"`
	IncurredAmount int64     `json:"incurred_amount,required"`
	NetAmount      int64     `json:"net_amount,required"`
	// Type of transactions.
	Type                    TransactionType        `json:"type,required"`
	Donation                TransactionDonation    `json:"donation,nullable"`
	IncurredByTransactionID string                 `json:"incurred_by_transaction_id,nullable" format:"uuid4"`
	IssueReward             TransactionIssueReward `json:"issue_reward,nullable"`
	IssueRewardID           string                 `json:"issue_reward_id,nullable" format:"uuid4"`
	// Last modification timestamp of the object.
	ModifiedAt          time.Time        `json:"modified_at,nullable" format:"date-time"`
	Order               TransactionOrder `json:"order,nullable"`
	PayoutTransactionID string           `json:"payout_transaction_id,nullable" format:"uuid4"`
	// Type of fees applied by Polar, and billed to the users.
	PlatformFeeType TransactionPlatformFeeType `json:"platform_fee_type,nullable"`
	Pledge          TransactionPledge          `json:"pledge,nullable"`
	PledgeID        string                     `json:"pledge_id,nullable" format:"uuid4"`
	// Supported payment processors.
	Processor      TransactionProcessor `json:"processor,nullable"`
	ProductPriceID string               `json:"product_price_id,nullable" format:"uuid4"`
	SubscriptionID string               `json:"subscription_id,nullable" format:"uuid4"`
	JSON           transactionJSON      `json:"-"`
}

// transactionJSON contains the JSON metadata for the struct [Transaction]
type transactionJSON struct {
	ID                          apijson.Field
	AccountAmount               apijson.Field
	AccountCurrency             apijson.Field
	AccountIncurredTransactions apijson.Field
	Amount                      apijson.Field
	CreatedAt                   apijson.Field
	Currency                    apijson.Field
	GrossAmount                 apijson.Field
	IncurredAmount              apijson.Field
	NetAmount                   apijson.Field
	Type                        apijson.Field
	Donation                    apijson.Field
	IncurredByTransactionID     apijson.Field
	IssueReward                 apijson.Field
	IssueRewardID               apijson.Field
	ModifiedAt                  apijson.Field
	Order                       apijson.Field
	PayoutTransactionID         apijson.Field
	PlatformFeeType             apijson.Field
	Pledge                      apijson.Field
	PledgeID                    apijson.Field
	Processor                   apijson.Field
	ProductPriceID              apijson.Field
	SubscriptionID              apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *Transaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionJSON) RawJSON() string {
	return r.raw
}

type TransactionAccountIncurredTransaction struct {
	ID              string `json:"id,required" format:"uuid4"`
	AccountAmount   int64  `json:"account_amount,required"`
	AccountCurrency string `json:"account_currency,required"`
	Amount          int64  `json:"amount,required"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	Currency  string    `json:"currency,required"`
	// Type of transactions.
	Type                    TransactionAccountIncurredTransactionsType `json:"type,required"`
	IncurredByTransactionID string                                     `json:"incurred_by_transaction_id,nullable" format:"uuid4"`
	IssueRewardID           string                                     `json:"issue_reward_id,nullable" format:"uuid4"`
	// Last modification timestamp of the object.
	ModifiedAt          time.Time `json:"modified_at,nullable" format:"date-time"`
	PayoutTransactionID string    `json:"payout_transaction_id,nullable" format:"uuid4"`
	// Type of fees applied by Polar, and billed to the users.
	PlatformFeeType TransactionAccountIncurredTransactionsPlatformFeeType `json:"platform_fee_type,nullable"`
	PledgeID        string                                                `json:"pledge_id,nullable" format:"uuid4"`
	// Supported payment processors.
	Processor      TransactionAccountIncurredTransactionsProcessor `json:"processor,nullable"`
	ProductPriceID string                                          `json:"product_price_id,nullable" format:"uuid4"`
	SubscriptionID string                                          `json:"subscription_id,nullable" format:"uuid4"`
	JSON           transactionAccountIncurredTransactionJSON       `json:"-"`
}

// transactionAccountIncurredTransactionJSON contains the JSON metadata for the
// struct [TransactionAccountIncurredTransaction]
type transactionAccountIncurredTransactionJSON struct {
	ID                      apijson.Field
	AccountAmount           apijson.Field
	AccountCurrency         apijson.Field
	Amount                  apijson.Field
	CreatedAt               apijson.Field
	Currency                apijson.Field
	Type                    apijson.Field
	IncurredByTransactionID apijson.Field
	IssueRewardID           apijson.Field
	ModifiedAt              apijson.Field
	PayoutTransactionID     apijson.Field
	PlatformFeeType         apijson.Field
	PledgeID                apijson.Field
	Processor               apijson.Field
	ProductPriceID          apijson.Field
	SubscriptionID          apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *TransactionAccountIncurredTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionAccountIncurredTransactionJSON) RawJSON() string {
	return r.raw
}

// Type of transactions.
type TransactionAccountIncurredTransactionsType string

const (
	TransactionAccountIncurredTransactionsTypePayment         TransactionAccountIncurredTransactionsType = "payment"
	TransactionAccountIncurredTransactionsTypeProcessorFee    TransactionAccountIncurredTransactionsType = "processor_fee"
	TransactionAccountIncurredTransactionsTypeRefund          TransactionAccountIncurredTransactionsType = "refund"
	TransactionAccountIncurredTransactionsTypeDispute         TransactionAccountIncurredTransactionsType = "dispute"
	TransactionAccountIncurredTransactionsTypeDisputeReversal TransactionAccountIncurredTransactionsType = "dispute_reversal"
	TransactionAccountIncurredTransactionsTypeBalance         TransactionAccountIncurredTransactionsType = "balance"
	TransactionAccountIncurredTransactionsTypePayout          TransactionAccountIncurredTransactionsType = "payout"
)

func (r TransactionAccountIncurredTransactionsType) IsKnown() bool {
	switch r {
	case TransactionAccountIncurredTransactionsTypePayment, TransactionAccountIncurredTransactionsTypeProcessorFee, TransactionAccountIncurredTransactionsTypeRefund, TransactionAccountIncurredTransactionsTypeDispute, TransactionAccountIncurredTransactionsTypeDisputeReversal, TransactionAccountIncurredTransactionsTypeBalance, TransactionAccountIncurredTransactionsTypePayout:
		return true
	}
	return false
}

// Type of fees applied by Polar, and billed to the users.
type TransactionAccountIncurredTransactionsPlatformFeeType string

const (
	TransactionAccountIncurredTransactionsPlatformFeeTypePlatform            TransactionAccountIncurredTransactionsPlatformFeeType = "platform"
	TransactionAccountIncurredTransactionsPlatformFeeTypePayment             TransactionAccountIncurredTransactionsPlatformFeeType = "payment"
	TransactionAccountIncurredTransactionsPlatformFeeTypeSubscription        TransactionAccountIncurredTransactionsPlatformFeeType = "subscription"
	TransactionAccountIncurredTransactionsPlatformFeeTypeInvoice             TransactionAccountIncurredTransactionsPlatformFeeType = "invoice"
	TransactionAccountIncurredTransactionsPlatformFeeTypeCrossBorderTransfer TransactionAccountIncurredTransactionsPlatformFeeType = "cross_border_transfer"
	TransactionAccountIncurredTransactionsPlatformFeeTypePayout              TransactionAccountIncurredTransactionsPlatformFeeType = "payout"
	TransactionAccountIncurredTransactionsPlatformFeeTypeAccount             TransactionAccountIncurredTransactionsPlatformFeeType = "account"
)

func (r TransactionAccountIncurredTransactionsPlatformFeeType) IsKnown() bool {
	switch r {
	case TransactionAccountIncurredTransactionsPlatformFeeTypePlatform, TransactionAccountIncurredTransactionsPlatformFeeTypePayment, TransactionAccountIncurredTransactionsPlatformFeeTypeSubscription, TransactionAccountIncurredTransactionsPlatformFeeTypeInvoice, TransactionAccountIncurredTransactionsPlatformFeeTypeCrossBorderTransfer, TransactionAccountIncurredTransactionsPlatformFeeTypePayout, TransactionAccountIncurredTransactionsPlatformFeeTypeAccount:
		return true
	}
	return false
}

// Supported payment processors.
type TransactionAccountIncurredTransactionsProcessor string

const (
	TransactionAccountIncurredTransactionsProcessorStripe         TransactionAccountIncurredTransactionsProcessor = "stripe"
	TransactionAccountIncurredTransactionsProcessorOpenCollective TransactionAccountIncurredTransactionsProcessor = "open_collective"
)

func (r TransactionAccountIncurredTransactionsProcessor) IsKnown() bool {
	switch r {
	case TransactionAccountIncurredTransactionsProcessorStripe, TransactionAccountIncurredTransactionsProcessorOpenCollective:
		return true
	}
	return false
}

// Type of transactions.
type TransactionType string

const (
	TransactionTypePayment         TransactionType = "payment"
	TransactionTypeProcessorFee    TransactionType = "processor_fee"
	TransactionTypeRefund          TransactionType = "refund"
	TransactionTypeDispute         TransactionType = "dispute"
	TransactionTypeDisputeReversal TransactionType = "dispute_reversal"
	TransactionTypeBalance         TransactionType = "balance"
	TransactionTypePayout          TransactionType = "payout"
)

func (r TransactionType) IsKnown() bool {
	switch r {
	case TransactionTypePayment, TransactionTypeProcessorFee, TransactionTypeRefund, TransactionTypeDispute, TransactionTypeDisputeReversal, TransactionTypeBalance, TransactionTypePayout:
		return true
	}
	return false
}

type TransactionDonation struct {
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Last modification timestamp of the object.
	ModifiedAt     time.Time                         `json:"modified_at,nullable" format:"date-time"`
	ToOrganization TransactionDonationToOrganization `json:"to_organization,nullable"`
	JSON           transactionDonationJSON           `json:"-"`
}

// transactionDonationJSON contains the JSON metadata for the struct
// [TransactionDonation]
type transactionDonationJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	ModifiedAt     apijson.Field
	ToOrganization apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TransactionDonation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionDonationJSON) RawJSON() string {
	return r.raw
}

type TransactionDonationToOrganization struct {
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	Name      string    `json:"name,required"`
	Slug      string    `json:"slug,required"`
	AvatarURL string    `json:"avatar_url,nullable"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time                             `json:"modified_at,nullable" format:"date-time"`
	JSON       transactionDonationToOrganizationJSON `json:"-"`
}

// transactionDonationToOrganizationJSON contains the JSON metadata for the struct
// [TransactionDonationToOrganization]
type transactionDonationToOrganizationJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	Slug        apijson.Field
	AvatarURL   apijson.Field
	ModifiedAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionDonationToOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionDonationToOrganizationJSON) RawJSON() string {
	return r.raw
}

type TransactionIssueReward struct {
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt      time.Time `json:"created_at,required" format:"date-time"`
	IssueID        string    `json:"issue_id,required" format:"uuid4"`
	ShareThousands int64     `json:"share_thousands,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time                  `json:"modified_at,nullable" format:"date-time"`
	JSON       transactionIssueRewardJSON `json:"-"`
}

// transactionIssueRewardJSON contains the JSON metadata for the struct
// [TransactionIssueReward]
type transactionIssueRewardJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	IssueID        apijson.Field
	ShareThousands apijson.Field
	ModifiedAt     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TransactionIssueReward) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionIssueRewardJSON) RawJSON() string {
	return r.raw
}

type TransactionOrder struct {
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time               `json:"created_at,required" format:"date-time"`
	Product   TransactionOrderProduct `json:"product,required"`
	// A recurring price for a product, i.e. a subscription.
	ProductPrice TransactionOrderProductPrice `json:"product_price,required"`
	// Last modification timestamp of the object.
	ModifiedAt     time.Time            `json:"modified_at,nullable" format:"date-time"`
	SubscriptionID string               `json:"subscription_id,nullable" format:"uuid4"`
	JSON           transactionOrderJSON `json:"-"`
}

// transactionOrderJSON contains the JSON metadata for the struct
// [TransactionOrder]
type transactionOrderJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Product        apijson.Field
	ProductPrice   apijson.Field
	ModifiedAt     apijson.Field
	SubscriptionID apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TransactionOrder) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionOrderJSON) RawJSON() string {
	return r.raw
}

type TransactionOrderProduct struct {
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	Name      string    `json:"name,required"`
	// Last modification timestamp of the object.
	ModifiedAt     time.Time                           `json:"modified_at,nullable" format:"date-time"`
	Organization   TransactionOrderProductOrganization `json:"organization,nullable"`
	OrganizationID string                              `json:"organization_id,nullable" format:"uuid4"`
	Type           TransactionOrderProductType         `json:"type,nullable"`
	JSON           transactionOrderProductJSON         `json:"-"`
}

// transactionOrderProductJSON contains the JSON metadata for the struct
// [TransactionOrderProduct]
type transactionOrderProductJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Name           apijson.Field
	ModifiedAt     apijson.Field
	Organization   apijson.Field
	OrganizationID apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TransactionOrderProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionOrderProductJSON) RawJSON() string {
	return r.raw
}

type TransactionOrderProductOrganization struct {
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	Name      string    `json:"name,required"`
	Slug      string    `json:"slug,required"`
	AvatarURL string    `json:"avatar_url,nullable"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time                               `json:"modified_at,nullable" format:"date-time"`
	JSON       transactionOrderProductOrganizationJSON `json:"-"`
}

// transactionOrderProductOrganizationJSON contains the JSON metadata for the
// struct [TransactionOrderProductOrganization]
type transactionOrderProductOrganizationJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	Slug        apijson.Field
	AvatarURL   apijson.Field
	ModifiedAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionOrderProductOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionOrderProductOrganizationJSON) RawJSON() string {
	return r.raw
}

type TransactionOrderProductType string

const (
	TransactionOrderProductTypeFree       TransactionOrderProductType = "free"
	TransactionOrderProductTypeIndividual TransactionOrderProductType = "individual"
	TransactionOrderProductTypeBusiness   TransactionOrderProductType = "business"
)

func (r TransactionOrderProductType) IsKnown() bool {
	switch r {
	case TransactionOrderProductTypeFree, TransactionOrderProductTypeIndividual, TransactionOrderProductTypeBusiness:
		return true
	}
	return false
}

// A recurring price for a product, i.e. a subscription.
type TransactionOrderProductPrice struct {
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
	Type TransactionOrderProductPriceType `json:"type,required"`
	// The recurring interval of the price, if type is `recurring`.
	RecurringInterval TransactionOrderProductPriceRecurringInterval `json:"recurring_interval,nullable"`
	JSON              transactionOrderProductPriceJSON              `json:"-"`
	union             TransactionOrderProductPriceUnion
}

// transactionOrderProductPriceJSON contains the JSON metadata for the struct
// [TransactionOrderProductPrice]
type transactionOrderProductPriceJSON struct {
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

func (r transactionOrderProductPriceJSON) RawJSON() string {
	return r.raw
}

func (r *TransactionOrderProductPrice) UnmarshalJSON(data []byte) (err error) {
	*r = TransactionOrderProductPrice{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [TransactionOrderProductPriceUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [TransactionOrderProductPriceProductPriceRecurring],
// [TransactionOrderProductPriceProductPriceOneTime].
func (r TransactionOrderProductPrice) AsUnion() TransactionOrderProductPriceUnion {
	return r.union
}

// A recurring price for a product, i.e. a subscription.
//
// Union satisfied by [TransactionOrderProductPriceProductPriceRecurring] or
// [TransactionOrderProductPriceProductPriceOneTime].
type TransactionOrderProductPriceUnion interface {
	implementsTransactionOrderProductPrice()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TransactionOrderProductPriceUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TransactionOrderProductPriceProductPriceRecurring{}),
			DiscriminatorValue: "recurring",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TransactionOrderProductPriceProductPriceOneTime{}),
			DiscriminatorValue: "one_time",
		},
	)
}

// A recurring price for a product, i.e. a subscription.
type TransactionOrderProductPriceProductPriceRecurring struct {
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
	Type TransactionOrderProductPriceProductPriceRecurringType `json:"type,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,nullable" format:"date-time"`
	// The recurring interval of the price, if type is `recurring`.
	RecurringInterval TransactionOrderProductPriceProductPriceRecurringRecurringInterval `json:"recurring_interval,nullable"`
	JSON              transactionOrderProductPriceProductPriceRecurringJSON              `json:"-"`
}

// transactionOrderProductPriceProductPriceRecurringJSON contains the JSON metadata
// for the struct [TransactionOrderProductPriceProductPriceRecurring]
type transactionOrderProductPriceProductPriceRecurringJSON struct {
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

func (r *TransactionOrderProductPriceProductPriceRecurring) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionOrderProductPriceProductPriceRecurringJSON) RawJSON() string {
	return r.raw
}

func (r TransactionOrderProductPriceProductPriceRecurring) implementsTransactionOrderProductPrice() {}

// The type of the price.
type TransactionOrderProductPriceProductPriceRecurringType string

const (
	TransactionOrderProductPriceProductPriceRecurringTypeRecurring TransactionOrderProductPriceProductPriceRecurringType = "recurring"
)

func (r TransactionOrderProductPriceProductPriceRecurringType) IsKnown() bool {
	switch r {
	case TransactionOrderProductPriceProductPriceRecurringTypeRecurring:
		return true
	}
	return false
}

// The recurring interval of the price, if type is `recurring`.
type TransactionOrderProductPriceProductPriceRecurringRecurringInterval string

const (
	TransactionOrderProductPriceProductPriceRecurringRecurringIntervalMonth TransactionOrderProductPriceProductPriceRecurringRecurringInterval = "month"
	TransactionOrderProductPriceProductPriceRecurringRecurringIntervalYear  TransactionOrderProductPriceProductPriceRecurringRecurringInterval = "year"
)

func (r TransactionOrderProductPriceProductPriceRecurringRecurringInterval) IsKnown() bool {
	switch r {
	case TransactionOrderProductPriceProductPriceRecurringRecurringIntervalMonth, TransactionOrderProductPriceProductPriceRecurringRecurringIntervalYear:
		return true
	}
	return false
}

// A one-time price for a product.
type TransactionOrderProductPriceProductPriceOneTime struct {
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
	Type TransactionOrderProductPriceProductPriceOneTimeType `json:"type,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time                                           `json:"modified_at,nullable" format:"date-time"`
	JSON       transactionOrderProductPriceProductPriceOneTimeJSON `json:"-"`
}

// transactionOrderProductPriceProductPriceOneTimeJSON contains the JSON metadata
// for the struct [TransactionOrderProductPriceProductPriceOneTime]
type transactionOrderProductPriceProductPriceOneTimeJSON struct {
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

func (r *TransactionOrderProductPriceProductPriceOneTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionOrderProductPriceProductPriceOneTimeJSON) RawJSON() string {
	return r.raw
}

func (r TransactionOrderProductPriceProductPriceOneTime) implementsTransactionOrderProductPrice() {}

// The type of the price.
type TransactionOrderProductPriceProductPriceOneTimeType string

const (
	TransactionOrderProductPriceProductPriceOneTimeTypeOneTime TransactionOrderProductPriceProductPriceOneTimeType = "one_time"
)

func (r TransactionOrderProductPriceProductPriceOneTimeType) IsKnown() bool {
	switch r {
	case TransactionOrderProductPriceProductPriceOneTimeTypeOneTime:
		return true
	}
	return false
}

// The type of the price.
type TransactionOrderProductPriceType string

const (
	TransactionOrderProductPriceTypeRecurring TransactionOrderProductPriceType = "recurring"
	TransactionOrderProductPriceTypeOneTime   TransactionOrderProductPriceType = "one_time"
)

func (r TransactionOrderProductPriceType) IsKnown() bool {
	switch r {
	case TransactionOrderProductPriceTypeRecurring, TransactionOrderProductPriceTypeOneTime:
		return true
	}
	return false
}

// The recurring interval of the price, if type is `recurring`.
type TransactionOrderProductPriceRecurringInterval string

const (
	TransactionOrderProductPriceRecurringIntervalMonth TransactionOrderProductPriceRecurringInterval = "month"
	TransactionOrderProductPriceRecurringIntervalYear  TransactionOrderProductPriceRecurringInterval = "year"
)

func (r TransactionOrderProductPriceRecurringInterval) IsKnown() bool {
	switch r {
	case TransactionOrderProductPriceRecurringIntervalMonth, TransactionOrderProductPriceRecurringIntervalYear:
		return true
	}
	return false
}

// Type of fees applied by Polar, and billed to the users.
type TransactionPlatformFeeType string

const (
	TransactionPlatformFeeTypePlatform            TransactionPlatformFeeType = "platform"
	TransactionPlatformFeeTypePayment             TransactionPlatformFeeType = "payment"
	TransactionPlatformFeeTypeSubscription        TransactionPlatformFeeType = "subscription"
	TransactionPlatformFeeTypeInvoice             TransactionPlatformFeeType = "invoice"
	TransactionPlatformFeeTypeCrossBorderTransfer TransactionPlatformFeeType = "cross_border_transfer"
	TransactionPlatformFeeTypePayout              TransactionPlatformFeeType = "payout"
	TransactionPlatformFeeTypeAccount             TransactionPlatformFeeType = "account"
)

func (r TransactionPlatformFeeType) IsKnown() bool {
	switch r {
	case TransactionPlatformFeeTypePlatform, TransactionPlatformFeeTypePayment, TransactionPlatformFeeTypeSubscription, TransactionPlatformFeeTypeInvoice, TransactionPlatformFeeTypeCrossBorderTransfer, TransactionPlatformFeeTypePayout, TransactionPlatformFeeTypeAccount:
		return true
	}
	return false
}

type TransactionPledge struct {
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time              `json:"created_at,required" format:"date-time"`
	Issue     TransactionPledgeIssue `json:"issue,required"`
	State     TransactionPledgeState `json:"state,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time             `json:"modified_at,nullable" format:"date-time"`
	JSON       transactionPledgeJSON `json:"-"`
}

// transactionPledgeJSON contains the JSON metadata for the struct
// [TransactionPledge]
type transactionPledgeJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Issue       apijson.Field
	State       apijson.Field
	ModifiedAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionPledge) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionPledgeJSON) RawJSON() string {
	return r.raw
}

type TransactionPledgeIssue struct {
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt      time.Time                          `json:"created_at,required" format:"date-time"`
	Number         int64                              `json:"number,required"`
	Organization   TransactionPledgeIssueOrganization `json:"organization,required"`
	OrganizationID string                             `json:"organization_id,required" format:"uuid4"`
	Platform       TransactionPledgeIssuePlatform     `json:"platform,required"`
	Repository     TransactionPledgeIssueRepository   `json:"repository,required"`
	RepositoryID   string                             `json:"repository_id,required" format:"uuid4"`
	Title          string                             `json:"title,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time                  `json:"modified_at,nullable" format:"date-time"`
	JSON       transactionPledgeIssueJSON `json:"-"`
}

// transactionPledgeIssueJSON contains the JSON metadata for the struct
// [TransactionPledgeIssue]
type transactionPledgeIssueJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Number         apijson.Field
	Organization   apijson.Field
	OrganizationID apijson.Field
	Platform       apijson.Field
	Repository     apijson.Field
	RepositoryID   apijson.Field
	Title          apijson.Field
	ModifiedAt     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TransactionPledgeIssue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionPledgeIssueJSON) RawJSON() string {
	return r.raw
}

type TransactionPledgeIssueOrganization struct {
	ID        string `json:"id,required" format:"uuid4"`
	AvatarURL string `json:"avatar_url,required"`
	// Creation timestamp of the object.
	CreatedAt  time.Time                                  `json:"created_at,required" format:"date-time"`
	IsPersonal bool                                       `json:"is_personal,required"`
	Name       string                                     `json:"name,required"`
	Platform   TransactionPledgeIssueOrganizationPlatform `json:"platform,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time                              `json:"modified_at,nullable" format:"date-time"`
	JSON       transactionPledgeIssueOrganizationJSON `json:"-"`
}

// transactionPledgeIssueOrganizationJSON contains the JSON metadata for the struct
// [TransactionPledgeIssueOrganization]
type transactionPledgeIssueOrganizationJSON struct {
	ID          apijson.Field
	AvatarURL   apijson.Field
	CreatedAt   apijson.Field
	IsPersonal  apijson.Field
	Name        apijson.Field
	Platform    apijson.Field
	ModifiedAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionPledgeIssueOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionPledgeIssueOrganizationJSON) RawJSON() string {
	return r.raw
}

type TransactionPledgeIssueOrganizationPlatform string

const (
	TransactionPledgeIssueOrganizationPlatformGitHub TransactionPledgeIssueOrganizationPlatform = "github"
)

func (r TransactionPledgeIssueOrganizationPlatform) IsKnown() bool {
	switch r {
	case TransactionPledgeIssueOrganizationPlatformGitHub:
		return true
	}
	return false
}

type TransactionPledgeIssuePlatform string

const (
	TransactionPledgeIssuePlatformGitHub TransactionPledgeIssuePlatform = "github"
)

func (r TransactionPledgeIssuePlatform) IsKnown() bool {
	switch r {
	case TransactionPledgeIssuePlatformGitHub:
		return true
	}
	return false
}

type TransactionPledgeIssueRepository struct {
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt      time.Time                                `json:"created_at,required" format:"date-time"`
	Name           string                                   `json:"name,required"`
	OrganizationID string                                   `json:"organization_id,required" format:"uuid4"`
	Platform       TransactionPledgeIssueRepositoryPlatform `json:"platform,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time                            `json:"modified_at,nullable" format:"date-time"`
	JSON       transactionPledgeIssueRepositoryJSON `json:"-"`
}

// transactionPledgeIssueRepositoryJSON contains the JSON metadata for the struct
// [TransactionPledgeIssueRepository]
type transactionPledgeIssueRepositoryJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Name           apijson.Field
	OrganizationID apijson.Field
	Platform       apijson.Field
	ModifiedAt     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TransactionPledgeIssueRepository) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionPledgeIssueRepositoryJSON) RawJSON() string {
	return r.raw
}

type TransactionPledgeIssueRepositoryPlatform string

const (
	TransactionPledgeIssueRepositoryPlatformGitHub TransactionPledgeIssueRepositoryPlatform = "github"
)

func (r TransactionPledgeIssueRepositoryPlatform) IsKnown() bool {
	switch r {
	case TransactionPledgeIssueRepositoryPlatformGitHub:
		return true
	}
	return false
}

type TransactionPledgeState string

const (
	TransactionPledgeStateInitiated      TransactionPledgeState = "initiated"
	TransactionPledgeStateCreated        TransactionPledgeState = "created"
	TransactionPledgeStatePending        TransactionPledgeState = "pending"
	TransactionPledgeStateRefunded       TransactionPledgeState = "refunded"
	TransactionPledgeStateDisputed       TransactionPledgeState = "disputed"
	TransactionPledgeStateChargeDisputed TransactionPledgeState = "charge_disputed"
	TransactionPledgeStateCancelled      TransactionPledgeState = "cancelled"
)

func (r TransactionPledgeState) IsKnown() bool {
	switch r {
	case TransactionPledgeStateInitiated, TransactionPledgeStateCreated, TransactionPledgeStatePending, TransactionPledgeStateRefunded, TransactionPledgeStateDisputed, TransactionPledgeStateChargeDisputed, TransactionPledgeStateCancelled:
		return true
	}
	return false
}

// Supported payment processors.
type TransactionProcessor string

const (
	TransactionProcessorStripe         TransactionProcessor = "stripe"
	TransactionProcessorOpenCollective TransactionProcessor = "open_collective"
)

func (r TransactionProcessor) IsKnown() bool {
	switch r {
	case TransactionProcessorStripe, TransactionProcessorOpenCollective:
		return true
	}
	return false
}

type TransactionPayoutCsvResponse = interface{}

type TransactionPayoutNewParams struct {
	AccountID param.Field[string] `json:"account_id,required" format:"uuid4"`
}

func (r TransactionPayoutNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TransactionPayoutListParams struct {
	AccountID param.Field[string] `query:"account_id,required" format:"uuid4"`
}

// URLQuery serializes [TransactionPayoutListParams]'s query parameters as
// `url.Values`.
func (r TransactionPayoutListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
