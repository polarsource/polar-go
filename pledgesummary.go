// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package polar

import (
	"context"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/polar-go/internal/apijson"
	"github.com/stainless-sdks/polar-go/internal/apiquery"
	"github.com/stainless-sdks/polar-go/internal/param"
	"github.com/stainless-sdks/polar-go/internal/requestconfig"
	"github.com/stainless-sdks/polar-go/option"
)

// PledgeSummaryService contains methods and other services that help with
// interacting with the polar API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPledgeSummaryService] method instead.
type PledgeSummaryService struct {
	Options []option.RequestOption
}

// NewPledgeSummaryService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPledgeSummaryService(opts ...option.RequestOption) (r *PledgeSummaryService) {
	r = &PledgeSummaryService{}
	r.Options = opts
	return
}

// Get summary of pledges for resource.
func (r *PledgeSummaryService) Get(ctx context.Context, query PledgeSummaryGetParams, opts ...option.RequestOption) (res *PledgeSummaryGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/pledges/summary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type PledgeSummaryGetResponse struct {
	Funding PledgeSummaryGetResponseFunding  `json:"funding,required"`
	Pledges []PledgeSummaryGetResponsePledge `json:"pledges,required"`
	JSON    pledgeSummaryGetResponseJSON     `json:"-"`
}

// pledgeSummaryGetResponseJSON contains the JSON metadata for the struct
// [PledgeSummaryGetResponse]
type pledgeSummaryGetResponseJSON struct {
	Funding     apijson.Field
	Pledges     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PledgeSummaryGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pledgeSummaryGetResponseJSON) RawJSON() string {
	return r.raw
}

type PledgeSummaryGetResponseFunding struct {
	FundingGoal PledgeSummaryGetResponseFundingFundingGoal `json:"funding_goal,nullable"`
	// Sum of pledges to this isuse (including currently open pledges and pledges that
	// have been paid out). Always in USD.
	PledgesSum PledgeSummaryGetResponseFundingPledgesSum `json:"pledges_sum,nullable"`
	JSON       pledgeSummaryGetResponseFundingJSON       `json:"-"`
}

// pledgeSummaryGetResponseFundingJSON contains the JSON metadata for the struct
// [PledgeSummaryGetResponseFunding]
type pledgeSummaryGetResponseFundingJSON struct {
	FundingGoal apijson.Field
	PledgesSum  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PledgeSummaryGetResponseFunding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pledgeSummaryGetResponseFundingJSON) RawJSON() string {
	return r.raw
}

type PledgeSummaryGetResponseFundingFundingGoal struct {
	// Amount in the currencies smallest unit (cents if currency is USD)
	Amount int64 `json:"amount,required"`
	// Three letter currency code (eg: USD)
	Currency string                                         `json:"currency,required"`
	JSON     pledgeSummaryGetResponseFundingFundingGoalJSON `json:"-"`
}

// pledgeSummaryGetResponseFundingFundingGoalJSON contains the JSON metadata for
// the struct [PledgeSummaryGetResponseFundingFundingGoal]
type pledgeSummaryGetResponseFundingFundingGoalJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PledgeSummaryGetResponseFundingFundingGoal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pledgeSummaryGetResponseFundingFundingGoalJSON) RawJSON() string {
	return r.raw
}

// Sum of pledges to this isuse (including currently open pledges and pledges that
// have been paid out). Always in USD.
type PledgeSummaryGetResponseFundingPledgesSum struct {
	// Amount in the currencies smallest unit (cents if currency is USD)
	Amount int64 `json:"amount,required"`
	// Three letter currency code (eg: USD)
	Currency string                                        `json:"currency,required"`
	JSON     pledgeSummaryGetResponseFundingPledgesSumJSON `json:"-"`
}

// pledgeSummaryGetResponseFundingPledgesSumJSON contains the JSON metadata for the
// struct [PledgeSummaryGetResponseFundingPledgesSum]
type pledgeSummaryGetResponseFundingPledgesSumJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PledgeSummaryGetResponseFundingPledgesSum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pledgeSummaryGetResponseFundingPledgesSumJSON) RawJSON() string {
	return r.raw
}

type PledgeSummaryGetResponsePledge struct {
	// Type of pledge
	Type    PledgeSummaryGetResponsePledgesType    `json:"type,required"`
	Pledger PledgeSummaryGetResponsePledgesPledger `json:"pledger,nullable"`
	JSON    pledgeSummaryGetResponsePledgeJSON     `json:"-"`
}

// pledgeSummaryGetResponsePledgeJSON contains the JSON metadata for the struct
// [PledgeSummaryGetResponsePledge]
type pledgeSummaryGetResponsePledgeJSON struct {
	Type        apijson.Field
	Pledger     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PledgeSummaryGetResponsePledge) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pledgeSummaryGetResponsePledgeJSON) RawJSON() string {
	return r.raw
}

// Type of pledge
type PledgeSummaryGetResponsePledgesType string

const (
	PledgeSummaryGetResponsePledgesTypePayUpfront      PledgeSummaryGetResponsePledgesType = "pay_upfront"
	PledgeSummaryGetResponsePledgesTypePayOnCompletion PledgeSummaryGetResponsePledgesType = "pay_on_completion"
	PledgeSummaryGetResponsePledgesTypePayDirectly     PledgeSummaryGetResponsePledgesType = "pay_directly"
)

func (r PledgeSummaryGetResponsePledgesType) IsKnown() bool {
	switch r {
	case PledgeSummaryGetResponsePledgesTypePayUpfront, PledgeSummaryGetResponsePledgesTypePayOnCompletion, PledgeSummaryGetResponsePledgesTypePayDirectly:
		return true
	}
	return false
}

type PledgeSummaryGetResponsePledgesPledger struct {
	Name           string                                     `json:"name,required"`
	AvatarURL      string                                     `json:"avatar_url,nullable"`
	GitHubUsername string                                     `json:"github_username,nullable"`
	JSON           pledgeSummaryGetResponsePledgesPledgerJSON `json:"-"`
}

// pledgeSummaryGetResponsePledgesPledgerJSON contains the JSON metadata for the
// struct [PledgeSummaryGetResponsePledgesPledger]
type pledgeSummaryGetResponsePledgesPledgerJSON struct {
	Name           apijson.Field
	AvatarURL      apijson.Field
	GitHubUsername apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PledgeSummaryGetResponsePledgesPledger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pledgeSummaryGetResponsePledgesPledgerJSON) RawJSON() string {
	return r.raw
}

type PledgeSummaryGetParams struct {
	IssueID param.Field[string] `query:"issue_id,required" format:"uuid"`
}

// URLQuery serializes [PledgeSummaryGetParams]'s query parameters as `url.Values`.
func (r PledgeSummaryGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
