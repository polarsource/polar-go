// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package polar

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/polarsource/polar-go/internal/apijson"
	"github.com/polarsource/polar-go/internal/apiquery"
	"github.com/polarsource/polar-go/internal/param"
	"github.com/polarsource/polar-go/internal/requestconfig"
	"github.com/polarsource/polar-go/option"
	"github.com/tidwall/gjson"
)

// TransactionService contains methods and other services that help with
// interacting with the polar API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransactionService] method instead.
type TransactionService struct {
	Options []option.RequestOption
	Payouts *TransactionPayoutService
}

// NewTransactionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTransactionService(opts ...option.RequestOption) (r *TransactionService) {
	r = &TransactionService{}
	r.Options = opts
	r.Payouts = NewTransactionPayoutService(opts...)
	return
}

// Lookup Transaction
func (r *TransactionService) Lookup(ctx context.Context, query TransactionLookupParams, opts ...option.RequestOption) (res *TransactionDetails, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/transactions/lookup"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Search Transactions
func (r *TransactionService) Search(ctx context.Context, query TransactionSearchParams, opts ...option.RequestOption) (res *ListResourceTransaction, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/transactions/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get Summary
func (r *TransactionService) Summary(ctx context.Context, query TransactionSummaryParams, opts ...option.RequestOption) (res *TransactionsSummary, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/transactions/summary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ListResourceTransaction struct {
	Pagination ListResourceTransactionPagination `json:"pagination,required"`
	Items      []Transaction                     `json:"items"`
	JSON       listResourceTransactionJSON       `json:"-"`
}

// listResourceTransactionJSON contains the JSON metadata for the struct
// [ListResourceTransaction]
type listResourceTransactionJSON struct {
	Pagination  apijson.Field
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListResourceTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResourceTransactionJSON) RawJSON() string {
	return r.raw
}

type ListResourceTransactionPagination struct {
	MaxPage    int64                                 `json:"max_page,required"`
	TotalCount int64                                 `json:"total_count,required"`
	JSON       listResourceTransactionPaginationJSON `json:"-"`
}

// listResourceTransactionPaginationJSON contains the JSON metadata for the struct
// [ListResourceTransactionPagination]
type listResourceTransactionPaginationJSON struct {
	MaxPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListResourceTransactionPagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResourceTransactionPaginationJSON) RawJSON() string {
	return r.raw
}

type TransactionDetails struct {
	ID                          string                                         `json:"id,required" format:"uuid4"`
	AccountAmount               int64                                          `json:"account_amount,required"`
	AccountCurrency             string                                         `json:"account_currency,required"`
	AccountIncurredTransactions []TransactionDetailsAccountIncurredTransaction `json:"account_incurred_transactions,required"`
	Amount                      int64                                          `json:"amount,required"`
	// Creation timestamp of the object.
	CreatedAt        time.Time     `json:"created_at,required" format:"date-time"`
	Currency         string        `json:"currency,required"`
	GrossAmount      int64         `json:"gross_amount,required"`
	IncurredAmount   int64         `json:"incurred_amount,required"`
	NetAmount        int64         `json:"net_amount,required"`
	PaidTransactions []Transaction `json:"paid_transactions,required"`
	// Type of transactions.
	Type                    TransactionDetailsType        `json:"type,required"`
	Donation                TransactionDetailsDonation    `json:"donation,nullable"`
	IncurredByTransactionID string                        `json:"incurred_by_transaction_id,nullable" format:"uuid4"`
	IssueReward             TransactionDetailsIssueReward `json:"issue_reward,nullable"`
	IssueRewardID           string                        `json:"issue_reward_id,nullable" format:"uuid4"`
	// Last modification timestamp of the object.
	ModifiedAt          time.Time               `json:"modified_at,nullable" format:"date-time"`
	Order               TransactionDetailsOrder `json:"order,nullable"`
	PayoutTransactionID string                  `json:"payout_transaction_id,nullable" format:"uuid4"`
	// Type of fees applied by Polar, and billed to the users.
	PlatformFeeType TransactionDetailsPlatformFeeType `json:"platform_fee_type,nullable"`
	Pledge          TransactionDetailsPledge          `json:"pledge,nullable"`
	PledgeID        string                            `json:"pledge_id,nullable" format:"uuid4"`
	// Supported payment processors.
	Processor      TransactionDetailsProcessor `json:"processor,nullable"`
	ProductPriceID string                      `json:"product_price_id,nullable" format:"uuid4"`
	SubscriptionID string                      `json:"subscription_id,nullable" format:"uuid4"`
	JSON           transactionDetailsJSON      `json:"-"`
}

// transactionDetailsJSON contains the JSON metadata for the struct
// [TransactionDetails]
type transactionDetailsJSON struct {
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
	PaidTransactions            apijson.Field
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

func (r *TransactionDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionDetailsJSON) RawJSON() string {
	return r.raw
}

type TransactionDetailsAccountIncurredTransaction struct {
	ID              string `json:"id,required" format:"uuid4"`
	AccountAmount   int64  `json:"account_amount,required"`
	AccountCurrency string `json:"account_currency,required"`
	Amount          int64  `json:"amount,required"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	Currency  string    `json:"currency,required"`
	// Type of transactions.
	Type                    TransactionDetailsAccountIncurredTransactionsType `json:"type,required"`
	IncurredByTransactionID string                                            `json:"incurred_by_transaction_id,nullable" format:"uuid4"`
	IssueRewardID           string                                            `json:"issue_reward_id,nullable" format:"uuid4"`
	// Last modification timestamp of the object.
	ModifiedAt          time.Time `json:"modified_at,nullable" format:"date-time"`
	PayoutTransactionID string    `json:"payout_transaction_id,nullable" format:"uuid4"`
	// Type of fees applied by Polar, and billed to the users.
	PlatformFeeType TransactionDetailsAccountIncurredTransactionsPlatformFeeType `json:"platform_fee_type,nullable"`
	PledgeID        string                                                       `json:"pledge_id,nullable" format:"uuid4"`
	// Supported payment processors.
	Processor      TransactionDetailsAccountIncurredTransactionsProcessor `json:"processor,nullable"`
	ProductPriceID string                                                 `json:"product_price_id,nullable" format:"uuid4"`
	SubscriptionID string                                                 `json:"subscription_id,nullable" format:"uuid4"`
	JSON           transactionDetailsAccountIncurredTransactionJSON       `json:"-"`
}

// transactionDetailsAccountIncurredTransactionJSON contains the JSON metadata for
// the struct [TransactionDetailsAccountIncurredTransaction]
type transactionDetailsAccountIncurredTransactionJSON struct {
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

func (r *TransactionDetailsAccountIncurredTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionDetailsAccountIncurredTransactionJSON) RawJSON() string {
	return r.raw
}

// Type of transactions.
type TransactionDetailsAccountIncurredTransactionsType string

const (
	TransactionDetailsAccountIncurredTransactionsTypePayment         TransactionDetailsAccountIncurredTransactionsType = "payment"
	TransactionDetailsAccountIncurredTransactionsTypeProcessorFee    TransactionDetailsAccountIncurredTransactionsType = "processor_fee"
	TransactionDetailsAccountIncurredTransactionsTypeRefund          TransactionDetailsAccountIncurredTransactionsType = "refund"
	TransactionDetailsAccountIncurredTransactionsTypeDispute         TransactionDetailsAccountIncurredTransactionsType = "dispute"
	TransactionDetailsAccountIncurredTransactionsTypeDisputeReversal TransactionDetailsAccountIncurredTransactionsType = "dispute_reversal"
	TransactionDetailsAccountIncurredTransactionsTypeBalance         TransactionDetailsAccountIncurredTransactionsType = "balance"
	TransactionDetailsAccountIncurredTransactionsTypePayout          TransactionDetailsAccountIncurredTransactionsType = "payout"
)

func (r TransactionDetailsAccountIncurredTransactionsType) IsKnown() bool {
	switch r {
	case TransactionDetailsAccountIncurredTransactionsTypePayment, TransactionDetailsAccountIncurredTransactionsTypeProcessorFee, TransactionDetailsAccountIncurredTransactionsTypeRefund, TransactionDetailsAccountIncurredTransactionsTypeDispute, TransactionDetailsAccountIncurredTransactionsTypeDisputeReversal, TransactionDetailsAccountIncurredTransactionsTypeBalance, TransactionDetailsAccountIncurredTransactionsTypePayout:
		return true
	}
	return false
}

// Type of fees applied by Polar, and billed to the users.
type TransactionDetailsAccountIncurredTransactionsPlatformFeeType string

const (
	TransactionDetailsAccountIncurredTransactionsPlatformFeeTypePlatform            TransactionDetailsAccountIncurredTransactionsPlatformFeeType = "platform"
	TransactionDetailsAccountIncurredTransactionsPlatformFeeTypePayment             TransactionDetailsAccountIncurredTransactionsPlatformFeeType = "payment"
	TransactionDetailsAccountIncurredTransactionsPlatformFeeTypeSubscription        TransactionDetailsAccountIncurredTransactionsPlatformFeeType = "subscription"
	TransactionDetailsAccountIncurredTransactionsPlatformFeeTypeInvoice             TransactionDetailsAccountIncurredTransactionsPlatformFeeType = "invoice"
	TransactionDetailsAccountIncurredTransactionsPlatformFeeTypeCrossBorderTransfer TransactionDetailsAccountIncurredTransactionsPlatformFeeType = "cross_border_transfer"
	TransactionDetailsAccountIncurredTransactionsPlatformFeeTypePayout              TransactionDetailsAccountIncurredTransactionsPlatformFeeType = "payout"
	TransactionDetailsAccountIncurredTransactionsPlatformFeeTypeAccount             TransactionDetailsAccountIncurredTransactionsPlatformFeeType = "account"
)

func (r TransactionDetailsAccountIncurredTransactionsPlatformFeeType) IsKnown() bool {
	switch r {
	case TransactionDetailsAccountIncurredTransactionsPlatformFeeTypePlatform, TransactionDetailsAccountIncurredTransactionsPlatformFeeTypePayment, TransactionDetailsAccountIncurredTransactionsPlatformFeeTypeSubscription, TransactionDetailsAccountIncurredTransactionsPlatformFeeTypeInvoice, TransactionDetailsAccountIncurredTransactionsPlatformFeeTypeCrossBorderTransfer, TransactionDetailsAccountIncurredTransactionsPlatformFeeTypePayout, TransactionDetailsAccountIncurredTransactionsPlatformFeeTypeAccount:
		return true
	}
	return false
}

// Supported payment processors.
type TransactionDetailsAccountIncurredTransactionsProcessor string

const (
	TransactionDetailsAccountIncurredTransactionsProcessorStripe         TransactionDetailsAccountIncurredTransactionsProcessor = "stripe"
	TransactionDetailsAccountIncurredTransactionsProcessorOpenCollective TransactionDetailsAccountIncurredTransactionsProcessor = "open_collective"
)

func (r TransactionDetailsAccountIncurredTransactionsProcessor) IsKnown() bool {
	switch r {
	case TransactionDetailsAccountIncurredTransactionsProcessorStripe, TransactionDetailsAccountIncurredTransactionsProcessorOpenCollective:
		return true
	}
	return false
}

// Type of transactions.
type TransactionDetailsType string

const (
	TransactionDetailsTypePayment         TransactionDetailsType = "payment"
	TransactionDetailsTypeProcessorFee    TransactionDetailsType = "processor_fee"
	TransactionDetailsTypeRefund          TransactionDetailsType = "refund"
	TransactionDetailsTypeDispute         TransactionDetailsType = "dispute"
	TransactionDetailsTypeDisputeReversal TransactionDetailsType = "dispute_reversal"
	TransactionDetailsTypeBalance         TransactionDetailsType = "balance"
	TransactionDetailsTypePayout          TransactionDetailsType = "payout"
)

func (r TransactionDetailsType) IsKnown() bool {
	switch r {
	case TransactionDetailsTypePayment, TransactionDetailsTypeProcessorFee, TransactionDetailsTypeRefund, TransactionDetailsTypeDispute, TransactionDetailsTypeDisputeReversal, TransactionDetailsTypeBalance, TransactionDetailsTypePayout:
		return true
	}
	return false
}

type TransactionDetailsDonation struct {
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Last modification timestamp of the object.
	ModifiedAt     time.Time                                `json:"modified_at,nullable" format:"date-time"`
	ToOrganization TransactionDetailsDonationToOrganization `json:"to_organization,nullable"`
	JSON           transactionDetailsDonationJSON           `json:"-"`
}

// transactionDetailsDonationJSON contains the JSON metadata for the struct
// [TransactionDetailsDonation]
type transactionDetailsDonationJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	ModifiedAt     apijson.Field
	ToOrganization apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TransactionDetailsDonation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionDetailsDonationJSON) RawJSON() string {
	return r.raw
}

type TransactionDetailsDonationToOrganization struct {
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	Name      string    `json:"name,required"`
	Slug      string    `json:"slug,required"`
	AvatarURL string    `json:"avatar_url,nullable"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time                                    `json:"modified_at,nullable" format:"date-time"`
	JSON       transactionDetailsDonationToOrganizationJSON `json:"-"`
}

// transactionDetailsDonationToOrganizationJSON contains the JSON metadata for the
// struct [TransactionDetailsDonationToOrganization]
type transactionDetailsDonationToOrganizationJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	Slug        apijson.Field
	AvatarURL   apijson.Field
	ModifiedAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionDetailsDonationToOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionDetailsDonationToOrganizationJSON) RawJSON() string {
	return r.raw
}

type TransactionDetailsIssueReward struct {
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt      time.Time `json:"created_at,required" format:"date-time"`
	IssueID        string    `json:"issue_id,required" format:"uuid4"`
	ShareThousands int64     `json:"share_thousands,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time                         `json:"modified_at,nullable" format:"date-time"`
	JSON       transactionDetailsIssueRewardJSON `json:"-"`
}

// transactionDetailsIssueRewardJSON contains the JSON metadata for the struct
// [TransactionDetailsIssueReward]
type transactionDetailsIssueRewardJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	IssueID        apijson.Field
	ShareThousands apijson.Field
	ModifiedAt     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TransactionDetailsIssueReward) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionDetailsIssueRewardJSON) RawJSON() string {
	return r.raw
}

type TransactionDetailsOrder struct {
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time                      `json:"created_at,required" format:"date-time"`
	Product   TransactionDetailsOrderProduct `json:"product,required"`
	// A recurring price for a product, i.e. a subscription.
	ProductPrice TransactionDetailsOrderProductPrice `json:"product_price,required"`
	// Last modification timestamp of the object.
	ModifiedAt     time.Time                   `json:"modified_at,nullable" format:"date-time"`
	SubscriptionID string                      `json:"subscription_id,nullable" format:"uuid4"`
	JSON           transactionDetailsOrderJSON `json:"-"`
}

// transactionDetailsOrderJSON contains the JSON metadata for the struct
// [TransactionDetailsOrder]
type transactionDetailsOrderJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Product        apijson.Field
	ProductPrice   apijson.Field
	ModifiedAt     apijson.Field
	SubscriptionID apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TransactionDetailsOrder) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionDetailsOrderJSON) RawJSON() string {
	return r.raw
}

type TransactionDetailsOrderProduct struct {
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	Name      string    `json:"name,required"`
	// Last modification timestamp of the object.
	ModifiedAt     time.Time                                  `json:"modified_at,nullable" format:"date-time"`
	Organization   TransactionDetailsOrderProductOrganization `json:"organization,nullable"`
	OrganizationID string                                     `json:"organization_id,nullable" format:"uuid4"`
	Type           TransactionDetailsOrderProductType         `json:"type,nullable"`
	JSON           transactionDetailsOrderProductJSON         `json:"-"`
}

// transactionDetailsOrderProductJSON contains the JSON metadata for the struct
// [TransactionDetailsOrderProduct]
type transactionDetailsOrderProductJSON struct {
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

func (r *TransactionDetailsOrderProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionDetailsOrderProductJSON) RawJSON() string {
	return r.raw
}

type TransactionDetailsOrderProductOrganization struct {
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	Name      string    `json:"name,required"`
	Slug      string    `json:"slug,required"`
	AvatarURL string    `json:"avatar_url,nullable"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time                                      `json:"modified_at,nullable" format:"date-time"`
	JSON       transactionDetailsOrderProductOrganizationJSON `json:"-"`
}

// transactionDetailsOrderProductOrganizationJSON contains the JSON metadata for
// the struct [TransactionDetailsOrderProductOrganization]
type transactionDetailsOrderProductOrganizationJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	Slug        apijson.Field
	AvatarURL   apijson.Field
	ModifiedAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionDetailsOrderProductOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionDetailsOrderProductOrganizationJSON) RawJSON() string {
	return r.raw
}

type TransactionDetailsOrderProductType string

const (
	TransactionDetailsOrderProductTypeFree       TransactionDetailsOrderProductType = "free"
	TransactionDetailsOrderProductTypeIndividual TransactionDetailsOrderProductType = "individual"
	TransactionDetailsOrderProductTypeBusiness   TransactionDetailsOrderProductType = "business"
)

func (r TransactionDetailsOrderProductType) IsKnown() bool {
	switch r {
	case TransactionDetailsOrderProductTypeFree, TransactionDetailsOrderProductTypeIndividual, TransactionDetailsOrderProductTypeBusiness:
		return true
	}
	return false
}

// A recurring price for a product, i.e. a subscription.
type TransactionDetailsOrderProductPrice struct {
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
	Type TransactionDetailsOrderProductPriceType `json:"type,required"`
	// The recurring interval of the price, if type is `recurring`.
	RecurringInterval TransactionDetailsOrderProductPriceRecurringInterval `json:"recurring_interval,nullable"`
	JSON              transactionDetailsOrderProductPriceJSON              `json:"-"`
	union             TransactionDetailsOrderProductPriceUnion
}

// transactionDetailsOrderProductPriceJSON contains the JSON metadata for the
// struct [TransactionDetailsOrderProductPrice]
type transactionDetailsOrderProductPriceJSON struct {
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

func (r transactionDetailsOrderProductPriceJSON) RawJSON() string {
	return r.raw
}

func (r *TransactionDetailsOrderProductPrice) UnmarshalJSON(data []byte) (err error) {
	*r = TransactionDetailsOrderProductPrice{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [TransactionDetailsOrderProductPriceUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [TransactionDetailsOrderProductPriceProductPriceRecurring],
// [TransactionDetailsOrderProductPriceProductPriceOneTime].
func (r TransactionDetailsOrderProductPrice) AsUnion() TransactionDetailsOrderProductPriceUnion {
	return r.union
}

// A recurring price for a product, i.e. a subscription.
//
// Union satisfied by [TransactionDetailsOrderProductPriceProductPriceRecurring] or
// [TransactionDetailsOrderProductPriceProductPriceOneTime].
type TransactionDetailsOrderProductPriceUnion interface {
	implementsTransactionDetailsOrderProductPrice()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TransactionDetailsOrderProductPriceUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TransactionDetailsOrderProductPriceProductPriceRecurring{}),
			DiscriminatorValue: "recurring",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TransactionDetailsOrderProductPriceProductPriceOneTime{}),
			DiscriminatorValue: "one_time",
		},
	)
}

// A recurring price for a product, i.e. a subscription.
type TransactionDetailsOrderProductPriceProductPriceRecurring struct {
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
	Type TransactionDetailsOrderProductPriceProductPriceRecurringType `json:"type,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time `json:"modified_at,nullable" format:"date-time"`
	// The recurring interval of the price, if type is `recurring`.
	RecurringInterval TransactionDetailsOrderProductPriceProductPriceRecurringRecurringInterval `json:"recurring_interval,nullable"`
	JSON              transactionDetailsOrderProductPriceProductPriceRecurringJSON              `json:"-"`
}

// transactionDetailsOrderProductPriceProductPriceRecurringJSON contains the JSON
// metadata for the struct
// [TransactionDetailsOrderProductPriceProductPriceRecurring]
type transactionDetailsOrderProductPriceProductPriceRecurringJSON struct {
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

func (r *TransactionDetailsOrderProductPriceProductPriceRecurring) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionDetailsOrderProductPriceProductPriceRecurringJSON) RawJSON() string {
	return r.raw
}

func (r TransactionDetailsOrderProductPriceProductPriceRecurring) implementsTransactionDetailsOrderProductPrice() {
}

// The type of the price.
type TransactionDetailsOrderProductPriceProductPriceRecurringType string

const (
	TransactionDetailsOrderProductPriceProductPriceRecurringTypeRecurring TransactionDetailsOrderProductPriceProductPriceRecurringType = "recurring"
)

func (r TransactionDetailsOrderProductPriceProductPriceRecurringType) IsKnown() bool {
	switch r {
	case TransactionDetailsOrderProductPriceProductPriceRecurringTypeRecurring:
		return true
	}
	return false
}

// The recurring interval of the price, if type is `recurring`.
type TransactionDetailsOrderProductPriceProductPriceRecurringRecurringInterval string

const (
	TransactionDetailsOrderProductPriceProductPriceRecurringRecurringIntervalMonth TransactionDetailsOrderProductPriceProductPriceRecurringRecurringInterval = "month"
	TransactionDetailsOrderProductPriceProductPriceRecurringRecurringIntervalYear  TransactionDetailsOrderProductPriceProductPriceRecurringRecurringInterval = "year"
)

func (r TransactionDetailsOrderProductPriceProductPriceRecurringRecurringInterval) IsKnown() bool {
	switch r {
	case TransactionDetailsOrderProductPriceProductPriceRecurringRecurringIntervalMonth, TransactionDetailsOrderProductPriceProductPriceRecurringRecurringIntervalYear:
		return true
	}
	return false
}

// A one-time price for a product.
type TransactionDetailsOrderProductPriceProductPriceOneTime struct {
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
	Type TransactionDetailsOrderProductPriceProductPriceOneTimeType `json:"type,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time                                                  `json:"modified_at,nullable" format:"date-time"`
	JSON       transactionDetailsOrderProductPriceProductPriceOneTimeJSON `json:"-"`
}

// transactionDetailsOrderProductPriceProductPriceOneTimeJSON contains the JSON
// metadata for the struct [TransactionDetailsOrderProductPriceProductPriceOneTime]
type transactionDetailsOrderProductPriceProductPriceOneTimeJSON struct {
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

func (r *TransactionDetailsOrderProductPriceProductPriceOneTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionDetailsOrderProductPriceProductPriceOneTimeJSON) RawJSON() string {
	return r.raw
}

func (r TransactionDetailsOrderProductPriceProductPriceOneTime) implementsTransactionDetailsOrderProductPrice() {
}

// The type of the price.
type TransactionDetailsOrderProductPriceProductPriceOneTimeType string

const (
	TransactionDetailsOrderProductPriceProductPriceOneTimeTypeOneTime TransactionDetailsOrderProductPriceProductPriceOneTimeType = "one_time"
)

func (r TransactionDetailsOrderProductPriceProductPriceOneTimeType) IsKnown() bool {
	switch r {
	case TransactionDetailsOrderProductPriceProductPriceOneTimeTypeOneTime:
		return true
	}
	return false
}

// The type of the price.
type TransactionDetailsOrderProductPriceType string

const (
	TransactionDetailsOrderProductPriceTypeRecurring TransactionDetailsOrderProductPriceType = "recurring"
	TransactionDetailsOrderProductPriceTypeOneTime   TransactionDetailsOrderProductPriceType = "one_time"
)

func (r TransactionDetailsOrderProductPriceType) IsKnown() bool {
	switch r {
	case TransactionDetailsOrderProductPriceTypeRecurring, TransactionDetailsOrderProductPriceTypeOneTime:
		return true
	}
	return false
}

// The recurring interval of the price, if type is `recurring`.
type TransactionDetailsOrderProductPriceRecurringInterval string

const (
	TransactionDetailsOrderProductPriceRecurringIntervalMonth TransactionDetailsOrderProductPriceRecurringInterval = "month"
	TransactionDetailsOrderProductPriceRecurringIntervalYear  TransactionDetailsOrderProductPriceRecurringInterval = "year"
)

func (r TransactionDetailsOrderProductPriceRecurringInterval) IsKnown() bool {
	switch r {
	case TransactionDetailsOrderProductPriceRecurringIntervalMonth, TransactionDetailsOrderProductPriceRecurringIntervalYear:
		return true
	}
	return false
}

// Type of fees applied by Polar, and billed to the users.
type TransactionDetailsPlatformFeeType string

const (
	TransactionDetailsPlatformFeeTypePlatform            TransactionDetailsPlatformFeeType = "platform"
	TransactionDetailsPlatformFeeTypePayment             TransactionDetailsPlatformFeeType = "payment"
	TransactionDetailsPlatformFeeTypeSubscription        TransactionDetailsPlatformFeeType = "subscription"
	TransactionDetailsPlatformFeeTypeInvoice             TransactionDetailsPlatformFeeType = "invoice"
	TransactionDetailsPlatformFeeTypeCrossBorderTransfer TransactionDetailsPlatformFeeType = "cross_border_transfer"
	TransactionDetailsPlatformFeeTypePayout              TransactionDetailsPlatformFeeType = "payout"
	TransactionDetailsPlatformFeeTypeAccount             TransactionDetailsPlatformFeeType = "account"
)

func (r TransactionDetailsPlatformFeeType) IsKnown() bool {
	switch r {
	case TransactionDetailsPlatformFeeTypePlatform, TransactionDetailsPlatformFeeTypePayment, TransactionDetailsPlatformFeeTypeSubscription, TransactionDetailsPlatformFeeTypeInvoice, TransactionDetailsPlatformFeeTypeCrossBorderTransfer, TransactionDetailsPlatformFeeTypePayout, TransactionDetailsPlatformFeeTypeAccount:
		return true
	}
	return false
}

type TransactionDetailsPledge struct {
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt time.Time                     `json:"created_at,required" format:"date-time"`
	Issue     TransactionDetailsPledgeIssue `json:"issue,required"`
	State     TransactionDetailsPledgeState `json:"state,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time                    `json:"modified_at,nullable" format:"date-time"`
	JSON       transactionDetailsPledgeJSON `json:"-"`
}

// transactionDetailsPledgeJSON contains the JSON metadata for the struct
// [TransactionDetailsPledge]
type transactionDetailsPledgeJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Issue       apijson.Field
	State       apijson.Field
	ModifiedAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionDetailsPledge) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionDetailsPledgeJSON) RawJSON() string {
	return r.raw
}

type TransactionDetailsPledgeIssue struct {
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt      time.Time                                 `json:"created_at,required" format:"date-time"`
	Number         int64                                     `json:"number,required"`
	Organization   TransactionDetailsPledgeIssueOrganization `json:"organization,required"`
	OrganizationID string                                    `json:"organization_id,required" format:"uuid4"`
	Platform       TransactionDetailsPledgeIssuePlatform     `json:"platform,required"`
	Repository     TransactionDetailsPledgeIssueRepository   `json:"repository,required"`
	RepositoryID   string                                    `json:"repository_id,required" format:"uuid4"`
	Title          string                                    `json:"title,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time                         `json:"modified_at,nullable" format:"date-time"`
	JSON       transactionDetailsPledgeIssueJSON `json:"-"`
}

// transactionDetailsPledgeIssueJSON contains the JSON metadata for the struct
// [TransactionDetailsPledgeIssue]
type transactionDetailsPledgeIssueJSON struct {
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

func (r *TransactionDetailsPledgeIssue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionDetailsPledgeIssueJSON) RawJSON() string {
	return r.raw
}

type TransactionDetailsPledgeIssueOrganization struct {
	ID        string `json:"id,required" format:"uuid4"`
	AvatarURL string `json:"avatar_url,required"`
	// Creation timestamp of the object.
	CreatedAt  time.Time                                         `json:"created_at,required" format:"date-time"`
	IsPersonal bool                                              `json:"is_personal,required"`
	Name       string                                            `json:"name,required"`
	Platform   TransactionDetailsPledgeIssueOrganizationPlatform `json:"platform,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time                                     `json:"modified_at,nullable" format:"date-time"`
	JSON       transactionDetailsPledgeIssueOrganizationJSON `json:"-"`
}

// transactionDetailsPledgeIssueOrganizationJSON contains the JSON metadata for the
// struct [TransactionDetailsPledgeIssueOrganization]
type transactionDetailsPledgeIssueOrganizationJSON struct {
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

func (r *TransactionDetailsPledgeIssueOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionDetailsPledgeIssueOrganizationJSON) RawJSON() string {
	return r.raw
}

type TransactionDetailsPledgeIssueOrganizationPlatform string

const (
	TransactionDetailsPledgeIssueOrganizationPlatformGitHub TransactionDetailsPledgeIssueOrganizationPlatform = "github"
)

func (r TransactionDetailsPledgeIssueOrganizationPlatform) IsKnown() bool {
	switch r {
	case TransactionDetailsPledgeIssueOrganizationPlatformGitHub:
		return true
	}
	return false
}

type TransactionDetailsPledgeIssuePlatform string

const (
	TransactionDetailsPledgeIssuePlatformGitHub TransactionDetailsPledgeIssuePlatform = "github"
)

func (r TransactionDetailsPledgeIssuePlatform) IsKnown() bool {
	switch r {
	case TransactionDetailsPledgeIssuePlatformGitHub:
		return true
	}
	return false
}

type TransactionDetailsPledgeIssueRepository struct {
	ID string `json:"id,required" format:"uuid4"`
	// Creation timestamp of the object.
	CreatedAt      time.Time                                       `json:"created_at,required" format:"date-time"`
	Name           string                                          `json:"name,required"`
	OrganizationID string                                          `json:"organization_id,required" format:"uuid4"`
	Platform       TransactionDetailsPledgeIssueRepositoryPlatform `json:"platform,required"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time                                   `json:"modified_at,nullable" format:"date-time"`
	JSON       transactionDetailsPledgeIssueRepositoryJSON `json:"-"`
}

// transactionDetailsPledgeIssueRepositoryJSON contains the JSON metadata for the
// struct [TransactionDetailsPledgeIssueRepository]
type transactionDetailsPledgeIssueRepositoryJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Name           apijson.Field
	OrganizationID apijson.Field
	Platform       apijson.Field
	ModifiedAt     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TransactionDetailsPledgeIssueRepository) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionDetailsPledgeIssueRepositoryJSON) RawJSON() string {
	return r.raw
}

type TransactionDetailsPledgeIssueRepositoryPlatform string

const (
	TransactionDetailsPledgeIssueRepositoryPlatformGitHub TransactionDetailsPledgeIssueRepositoryPlatform = "github"
)

func (r TransactionDetailsPledgeIssueRepositoryPlatform) IsKnown() bool {
	switch r {
	case TransactionDetailsPledgeIssueRepositoryPlatformGitHub:
		return true
	}
	return false
}

type TransactionDetailsPledgeState string

const (
	TransactionDetailsPledgeStateInitiated      TransactionDetailsPledgeState = "initiated"
	TransactionDetailsPledgeStateCreated        TransactionDetailsPledgeState = "created"
	TransactionDetailsPledgeStatePending        TransactionDetailsPledgeState = "pending"
	TransactionDetailsPledgeStateRefunded       TransactionDetailsPledgeState = "refunded"
	TransactionDetailsPledgeStateDisputed       TransactionDetailsPledgeState = "disputed"
	TransactionDetailsPledgeStateChargeDisputed TransactionDetailsPledgeState = "charge_disputed"
	TransactionDetailsPledgeStateCancelled      TransactionDetailsPledgeState = "cancelled"
)

func (r TransactionDetailsPledgeState) IsKnown() bool {
	switch r {
	case TransactionDetailsPledgeStateInitiated, TransactionDetailsPledgeStateCreated, TransactionDetailsPledgeStatePending, TransactionDetailsPledgeStateRefunded, TransactionDetailsPledgeStateDisputed, TransactionDetailsPledgeStateChargeDisputed, TransactionDetailsPledgeStateCancelled:
		return true
	}
	return false
}

// Supported payment processors.
type TransactionDetailsProcessor string

const (
	TransactionDetailsProcessorStripe         TransactionDetailsProcessor = "stripe"
	TransactionDetailsProcessorOpenCollective TransactionDetailsProcessor = "open_collective"
)

func (r TransactionDetailsProcessor) IsKnown() bool {
	switch r {
	case TransactionDetailsProcessorStripe, TransactionDetailsProcessorOpenCollective:
		return true
	}
	return false
}

type TransactionsSummary struct {
	Balance TransactionsSummaryBalance `json:"balance,required"`
	Payout  TransactionsSummaryPayout  `json:"payout,required"`
	JSON    transactionsSummaryJSON    `json:"-"`
}

// transactionsSummaryJSON contains the JSON metadata for the struct
// [TransactionsSummary]
type transactionsSummaryJSON struct {
	Balance     apijson.Field
	Payout      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionsSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionsSummaryJSON) RawJSON() string {
	return r.raw
}

type TransactionsSummaryBalance struct {
	AccountAmount   int64                          `json:"account_amount,required"`
	AccountCurrency string                         `json:"account_currency,required"`
	Amount          int64                          `json:"amount,required"`
	Currency        string                         `json:"currency,required"`
	JSON            transactionsSummaryBalanceJSON `json:"-"`
}

// transactionsSummaryBalanceJSON contains the JSON metadata for the struct
// [TransactionsSummaryBalance]
type transactionsSummaryBalanceJSON struct {
	AccountAmount   apijson.Field
	AccountCurrency apijson.Field
	Amount          apijson.Field
	Currency        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *TransactionsSummaryBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionsSummaryBalanceJSON) RawJSON() string {
	return r.raw
}

type TransactionsSummaryPayout struct {
	AccountAmount   int64                         `json:"account_amount,required"`
	AccountCurrency string                        `json:"account_currency,required"`
	Amount          int64                         `json:"amount,required"`
	Currency        string                        `json:"currency,required"`
	JSON            transactionsSummaryPayoutJSON `json:"-"`
}

// transactionsSummaryPayoutJSON contains the JSON metadata for the struct
// [TransactionsSummaryPayout]
type transactionsSummaryPayoutJSON struct {
	AccountAmount   apijson.Field
	AccountCurrency apijson.Field
	Amount          apijson.Field
	Currency        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *TransactionsSummaryPayout) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionsSummaryPayoutJSON) RawJSON() string {
	return r.raw
}

type TransactionLookupParams struct {
	TransactionID param.Field[string] `query:"transaction_id,required" format:"uuid4"`
}

// URLQuery serializes [TransactionLookupParams]'s query parameters as
// `url.Values`.
func (r TransactionLookupParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TransactionSearchParams struct {
	AccountID           param.Field[string] `query:"account_id" format:"uuid4"`
	ExcludePlatformFees param.Field[bool]   `query:"exclude_platform_fees"`
	// Size of a page, defaults to 10. Maximum is 100.
	Limit param.Field[int64] `query:"limit"`
	// Page number, defaults to 1.
	Page                  param.Field[int64]  `query:"page"`
	PaymentOrganizationID param.Field[string] `query:"payment_organization_id" format:"uuid4"`
	PaymentUserID         param.Field[string] `query:"payment_user_id" format:"uuid4"`
	// Sorting criterion. Several criteria can be used simultaneously and will be
	// applied in order. Add a minus sign `-` before the criteria name to sort by
	// descending order.
	Sorting param.Field[[]string] `query:"sorting"`
	// Type of transactions.
	Type param.Field[TransactionSearchParamsType] `query:"type"`
}

// URLQuery serializes [TransactionSearchParams]'s query parameters as
// `url.Values`.
func (r TransactionSearchParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Type of transactions.
type TransactionSearchParamsType string

const (
	TransactionSearchParamsTypePayment         TransactionSearchParamsType = "payment"
	TransactionSearchParamsTypeProcessorFee    TransactionSearchParamsType = "processor_fee"
	TransactionSearchParamsTypeRefund          TransactionSearchParamsType = "refund"
	TransactionSearchParamsTypeDispute         TransactionSearchParamsType = "dispute"
	TransactionSearchParamsTypeDisputeReversal TransactionSearchParamsType = "dispute_reversal"
	TransactionSearchParamsTypeBalance         TransactionSearchParamsType = "balance"
	TransactionSearchParamsTypePayout          TransactionSearchParamsType = "payout"
)

func (r TransactionSearchParamsType) IsKnown() bool {
	switch r {
	case TransactionSearchParamsTypePayment, TransactionSearchParamsTypeProcessorFee, TransactionSearchParamsTypeRefund, TransactionSearchParamsTypeDispute, TransactionSearchParamsTypeDisputeReversal, TransactionSearchParamsTypeBalance, TransactionSearchParamsTypePayout:
		return true
	}
	return false
}

type TransactionSummaryParams struct {
	AccountID param.Field[string] `query:"account_id,required" format:"uuid4"`
}

// URLQuery serializes [TransactionSummaryParams]'s query parameters as
// `url.Values`.
func (r TransactionSummaryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
