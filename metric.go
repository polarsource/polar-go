// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package polar

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/polar-go/internal/apijson"
	"github.com/stainless-sdks/polar-go/internal/apiquery"
	"github.com/stainless-sdks/polar-go/internal/param"
	"github.com/stainless-sdks/polar-go/internal/requestconfig"
	"github.com/stainless-sdks/polar-go/option"
)

// MetricService contains methods and other services that help with interacting
// with the polar API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMetricService] method instead.
type MetricService struct {
	Options []option.RequestOption
	Limits  *MetricLimitService
}

// NewMetricService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMetricService(opts ...option.RequestOption) (r *MetricService) {
	r = &MetricService{}
	r.Options = opts
	r.Limits = NewMetricLimitService(opts...)
	return
}

// Get metrics about your orders and subscriptions.
func (r *MetricService) List(ctx context.Context, query MetricListParams, opts ...option.RequestOption) (res *MetricsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/metrics/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Metrics response schema.
type MetricsResponse struct {
	// Information about the returned metrics.
	Metrics MetricsResponseMetrics `json:"metrics,required"`
	// List of data for each timestamp.
	Periods []MetricsResponsePeriod `json:"periods,required"`
	JSON    metricsResponseJSON     `json:"-"`
}

// metricsResponseJSON contains the JSON metadata for the struct [MetricsResponse]
type metricsResponseJSON struct {
	Metrics     apijson.Field
	Periods     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MetricsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r metricsResponseJSON) RawJSON() string {
	return r.raw
}

// Information about the returned metrics.
type MetricsResponseMetrics struct {
	// Information about a metric.
	ActiveSubscriptions MetricsResponseMetricsActiveSubscriptions `json:"active_subscriptions,required"`
	// Information about a metric.
	AverageOrderValue MetricsResponseMetricsAverageOrderValue `json:"average_order_value,required"`
	// Information about a metric.
	MonthlyRecurringRevenue MetricsResponseMetricsMonthlyRecurringRevenue `json:"monthly_recurring_revenue,required"`
	// Information about a metric.
	NewSubscriptions MetricsResponseMetricsNewSubscriptions `json:"new_subscriptions,required"`
	// Information about a metric.
	NewSubscriptionsRevenue MetricsResponseMetricsNewSubscriptionsRevenue `json:"new_subscriptions_revenue,required"`
	// Information about a metric.
	OneTimeProducts MetricsResponseMetricsOneTimeProducts `json:"one_time_products,required"`
	// Information about a metric.
	OneTimeProductsRevenue MetricsResponseMetricsOneTimeProductsRevenue `json:"one_time_products_revenue,required"`
	// Information about a metric.
	Orders MetricsResponseMetricsOrders `json:"orders,required"`
	// Information about a metric.
	RenewedSubscriptions MetricsResponseMetricsRenewedSubscriptions `json:"renewed_subscriptions,required"`
	// Information about a metric.
	RenewedSubscriptionsRevenue MetricsResponseMetricsRenewedSubscriptionsRevenue `json:"renewed_subscriptions_revenue,required"`
	// Information about a metric.
	Revenue MetricsResponseMetricsRevenue `json:"revenue,required"`
	JSON    metricsResponseMetricsJSON    `json:"-"`
}

// metricsResponseMetricsJSON contains the JSON metadata for the struct
// [MetricsResponseMetrics]
type metricsResponseMetricsJSON struct {
	ActiveSubscriptions         apijson.Field
	AverageOrderValue           apijson.Field
	MonthlyRecurringRevenue     apijson.Field
	NewSubscriptions            apijson.Field
	NewSubscriptionsRevenue     apijson.Field
	OneTimeProducts             apijson.Field
	OneTimeProductsRevenue      apijson.Field
	Orders                      apijson.Field
	RenewedSubscriptions        apijson.Field
	RenewedSubscriptionsRevenue apijson.Field
	Revenue                     apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *MetricsResponseMetrics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r metricsResponseMetricsJSON) RawJSON() string {
	return r.raw
}

// Information about a metric.
type MetricsResponseMetricsActiveSubscriptions struct {
	// Human-readable name for the metric.
	DisplayName string `json:"display_name,required"`
	// Unique identifier for the metric.
	Slug string `json:"slug,required"`
	// Type of the metric, useful to know the unit or format of the value.
	Type MetricsResponseMetricsActiveSubscriptionsType `json:"type,required"`
	JSON metricsResponseMetricsActiveSubscriptionsJSON `json:"-"`
}

// metricsResponseMetricsActiveSubscriptionsJSON contains the JSON metadata for the
// struct [MetricsResponseMetricsActiveSubscriptions]
type metricsResponseMetricsActiveSubscriptionsJSON struct {
	DisplayName apijson.Field
	Slug        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MetricsResponseMetricsActiveSubscriptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r metricsResponseMetricsActiveSubscriptionsJSON) RawJSON() string {
	return r.raw
}

// Type of the metric, useful to know the unit or format of the value.
type MetricsResponseMetricsActiveSubscriptionsType string

const (
	MetricsResponseMetricsActiveSubscriptionsTypeScalar   MetricsResponseMetricsActiveSubscriptionsType = "scalar"
	MetricsResponseMetricsActiveSubscriptionsTypeCurrency MetricsResponseMetricsActiveSubscriptionsType = "currency"
)

func (r MetricsResponseMetricsActiveSubscriptionsType) IsKnown() bool {
	switch r {
	case MetricsResponseMetricsActiveSubscriptionsTypeScalar, MetricsResponseMetricsActiveSubscriptionsTypeCurrency:
		return true
	}
	return false
}

// Information about a metric.
type MetricsResponseMetricsAverageOrderValue struct {
	// Human-readable name for the metric.
	DisplayName string `json:"display_name,required"`
	// Unique identifier for the metric.
	Slug string `json:"slug,required"`
	// Type of the metric, useful to know the unit or format of the value.
	Type MetricsResponseMetricsAverageOrderValueType `json:"type,required"`
	JSON metricsResponseMetricsAverageOrderValueJSON `json:"-"`
}

// metricsResponseMetricsAverageOrderValueJSON contains the JSON metadata for the
// struct [MetricsResponseMetricsAverageOrderValue]
type metricsResponseMetricsAverageOrderValueJSON struct {
	DisplayName apijson.Field
	Slug        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MetricsResponseMetricsAverageOrderValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r metricsResponseMetricsAverageOrderValueJSON) RawJSON() string {
	return r.raw
}

// Type of the metric, useful to know the unit or format of the value.
type MetricsResponseMetricsAverageOrderValueType string

const (
	MetricsResponseMetricsAverageOrderValueTypeScalar   MetricsResponseMetricsAverageOrderValueType = "scalar"
	MetricsResponseMetricsAverageOrderValueTypeCurrency MetricsResponseMetricsAverageOrderValueType = "currency"
)

func (r MetricsResponseMetricsAverageOrderValueType) IsKnown() bool {
	switch r {
	case MetricsResponseMetricsAverageOrderValueTypeScalar, MetricsResponseMetricsAverageOrderValueTypeCurrency:
		return true
	}
	return false
}

// Information about a metric.
type MetricsResponseMetricsMonthlyRecurringRevenue struct {
	// Human-readable name for the metric.
	DisplayName string `json:"display_name,required"`
	// Unique identifier for the metric.
	Slug string `json:"slug,required"`
	// Type of the metric, useful to know the unit or format of the value.
	Type MetricsResponseMetricsMonthlyRecurringRevenueType `json:"type,required"`
	JSON metricsResponseMetricsMonthlyRecurringRevenueJSON `json:"-"`
}

// metricsResponseMetricsMonthlyRecurringRevenueJSON contains the JSON metadata for
// the struct [MetricsResponseMetricsMonthlyRecurringRevenue]
type metricsResponseMetricsMonthlyRecurringRevenueJSON struct {
	DisplayName apijson.Field
	Slug        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MetricsResponseMetricsMonthlyRecurringRevenue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r metricsResponseMetricsMonthlyRecurringRevenueJSON) RawJSON() string {
	return r.raw
}

// Type of the metric, useful to know the unit or format of the value.
type MetricsResponseMetricsMonthlyRecurringRevenueType string

const (
	MetricsResponseMetricsMonthlyRecurringRevenueTypeScalar   MetricsResponseMetricsMonthlyRecurringRevenueType = "scalar"
	MetricsResponseMetricsMonthlyRecurringRevenueTypeCurrency MetricsResponseMetricsMonthlyRecurringRevenueType = "currency"
)

func (r MetricsResponseMetricsMonthlyRecurringRevenueType) IsKnown() bool {
	switch r {
	case MetricsResponseMetricsMonthlyRecurringRevenueTypeScalar, MetricsResponseMetricsMonthlyRecurringRevenueTypeCurrency:
		return true
	}
	return false
}

// Information about a metric.
type MetricsResponseMetricsNewSubscriptions struct {
	// Human-readable name for the metric.
	DisplayName string `json:"display_name,required"`
	// Unique identifier for the metric.
	Slug string `json:"slug,required"`
	// Type of the metric, useful to know the unit or format of the value.
	Type MetricsResponseMetricsNewSubscriptionsType `json:"type,required"`
	JSON metricsResponseMetricsNewSubscriptionsJSON `json:"-"`
}

// metricsResponseMetricsNewSubscriptionsJSON contains the JSON metadata for the
// struct [MetricsResponseMetricsNewSubscriptions]
type metricsResponseMetricsNewSubscriptionsJSON struct {
	DisplayName apijson.Field
	Slug        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MetricsResponseMetricsNewSubscriptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r metricsResponseMetricsNewSubscriptionsJSON) RawJSON() string {
	return r.raw
}

// Type of the metric, useful to know the unit or format of the value.
type MetricsResponseMetricsNewSubscriptionsType string

const (
	MetricsResponseMetricsNewSubscriptionsTypeScalar   MetricsResponseMetricsNewSubscriptionsType = "scalar"
	MetricsResponseMetricsNewSubscriptionsTypeCurrency MetricsResponseMetricsNewSubscriptionsType = "currency"
)

func (r MetricsResponseMetricsNewSubscriptionsType) IsKnown() bool {
	switch r {
	case MetricsResponseMetricsNewSubscriptionsTypeScalar, MetricsResponseMetricsNewSubscriptionsTypeCurrency:
		return true
	}
	return false
}

// Information about a metric.
type MetricsResponseMetricsNewSubscriptionsRevenue struct {
	// Human-readable name for the metric.
	DisplayName string `json:"display_name,required"`
	// Unique identifier for the metric.
	Slug string `json:"slug,required"`
	// Type of the metric, useful to know the unit or format of the value.
	Type MetricsResponseMetricsNewSubscriptionsRevenueType `json:"type,required"`
	JSON metricsResponseMetricsNewSubscriptionsRevenueJSON `json:"-"`
}

// metricsResponseMetricsNewSubscriptionsRevenueJSON contains the JSON metadata for
// the struct [MetricsResponseMetricsNewSubscriptionsRevenue]
type metricsResponseMetricsNewSubscriptionsRevenueJSON struct {
	DisplayName apijson.Field
	Slug        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MetricsResponseMetricsNewSubscriptionsRevenue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r metricsResponseMetricsNewSubscriptionsRevenueJSON) RawJSON() string {
	return r.raw
}

// Type of the metric, useful to know the unit or format of the value.
type MetricsResponseMetricsNewSubscriptionsRevenueType string

const (
	MetricsResponseMetricsNewSubscriptionsRevenueTypeScalar   MetricsResponseMetricsNewSubscriptionsRevenueType = "scalar"
	MetricsResponseMetricsNewSubscriptionsRevenueTypeCurrency MetricsResponseMetricsNewSubscriptionsRevenueType = "currency"
)

func (r MetricsResponseMetricsNewSubscriptionsRevenueType) IsKnown() bool {
	switch r {
	case MetricsResponseMetricsNewSubscriptionsRevenueTypeScalar, MetricsResponseMetricsNewSubscriptionsRevenueTypeCurrency:
		return true
	}
	return false
}

// Information about a metric.
type MetricsResponseMetricsOneTimeProducts struct {
	// Human-readable name for the metric.
	DisplayName string `json:"display_name,required"`
	// Unique identifier for the metric.
	Slug string `json:"slug,required"`
	// Type of the metric, useful to know the unit or format of the value.
	Type MetricsResponseMetricsOneTimeProductsType `json:"type,required"`
	JSON metricsResponseMetricsOneTimeProductsJSON `json:"-"`
}

// metricsResponseMetricsOneTimeProductsJSON contains the JSON metadata for the
// struct [MetricsResponseMetricsOneTimeProducts]
type metricsResponseMetricsOneTimeProductsJSON struct {
	DisplayName apijson.Field
	Slug        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MetricsResponseMetricsOneTimeProducts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r metricsResponseMetricsOneTimeProductsJSON) RawJSON() string {
	return r.raw
}

// Type of the metric, useful to know the unit or format of the value.
type MetricsResponseMetricsOneTimeProductsType string

const (
	MetricsResponseMetricsOneTimeProductsTypeScalar   MetricsResponseMetricsOneTimeProductsType = "scalar"
	MetricsResponseMetricsOneTimeProductsTypeCurrency MetricsResponseMetricsOneTimeProductsType = "currency"
)

func (r MetricsResponseMetricsOneTimeProductsType) IsKnown() bool {
	switch r {
	case MetricsResponseMetricsOneTimeProductsTypeScalar, MetricsResponseMetricsOneTimeProductsTypeCurrency:
		return true
	}
	return false
}

// Information about a metric.
type MetricsResponseMetricsOneTimeProductsRevenue struct {
	// Human-readable name for the metric.
	DisplayName string `json:"display_name,required"`
	// Unique identifier for the metric.
	Slug string `json:"slug,required"`
	// Type of the metric, useful to know the unit or format of the value.
	Type MetricsResponseMetricsOneTimeProductsRevenueType `json:"type,required"`
	JSON metricsResponseMetricsOneTimeProductsRevenueJSON `json:"-"`
}

// metricsResponseMetricsOneTimeProductsRevenueJSON contains the JSON metadata for
// the struct [MetricsResponseMetricsOneTimeProductsRevenue]
type metricsResponseMetricsOneTimeProductsRevenueJSON struct {
	DisplayName apijson.Field
	Slug        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MetricsResponseMetricsOneTimeProductsRevenue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r metricsResponseMetricsOneTimeProductsRevenueJSON) RawJSON() string {
	return r.raw
}

// Type of the metric, useful to know the unit or format of the value.
type MetricsResponseMetricsOneTimeProductsRevenueType string

const (
	MetricsResponseMetricsOneTimeProductsRevenueTypeScalar   MetricsResponseMetricsOneTimeProductsRevenueType = "scalar"
	MetricsResponseMetricsOneTimeProductsRevenueTypeCurrency MetricsResponseMetricsOneTimeProductsRevenueType = "currency"
)

func (r MetricsResponseMetricsOneTimeProductsRevenueType) IsKnown() bool {
	switch r {
	case MetricsResponseMetricsOneTimeProductsRevenueTypeScalar, MetricsResponseMetricsOneTimeProductsRevenueTypeCurrency:
		return true
	}
	return false
}

// Information about a metric.
type MetricsResponseMetricsOrders struct {
	// Human-readable name for the metric.
	DisplayName string `json:"display_name,required"`
	// Unique identifier for the metric.
	Slug string `json:"slug,required"`
	// Type of the metric, useful to know the unit or format of the value.
	Type MetricsResponseMetricsOrdersType `json:"type,required"`
	JSON metricsResponseMetricsOrdersJSON `json:"-"`
}

// metricsResponseMetricsOrdersJSON contains the JSON metadata for the struct
// [MetricsResponseMetricsOrders]
type metricsResponseMetricsOrdersJSON struct {
	DisplayName apijson.Field
	Slug        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MetricsResponseMetricsOrders) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r metricsResponseMetricsOrdersJSON) RawJSON() string {
	return r.raw
}

// Type of the metric, useful to know the unit or format of the value.
type MetricsResponseMetricsOrdersType string

const (
	MetricsResponseMetricsOrdersTypeScalar   MetricsResponseMetricsOrdersType = "scalar"
	MetricsResponseMetricsOrdersTypeCurrency MetricsResponseMetricsOrdersType = "currency"
)

func (r MetricsResponseMetricsOrdersType) IsKnown() bool {
	switch r {
	case MetricsResponseMetricsOrdersTypeScalar, MetricsResponseMetricsOrdersTypeCurrency:
		return true
	}
	return false
}

// Information about a metric.
type MetricsResponseMetricsRenewedSubscriptions struct {
	// Human-readable name for the metric.
	DisplayName string `json:"display_name,required"`
	// Unique identifier for the metric.
	Slug string `json:"slug,required"`
	// Type of the metric, useful to know the unit or format of the value.
	Type MetricsResponseMetricsRenewedSubscriptionsType `json:"type,required"`
	JSON metricsResponseMetricsRenewedSubscriptionsJSON `json:"-"`
}

// metricsResponseMetricsRenewedSubscriptionsJSON contains the JSON metadata for
// the struct [MetricsResponseMetricsRenewedSubscriptions]
type metricsResponseMetricsRenewedSubscriptionsJSON struct {
	DisplayName apijson.Field
	Slug        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MetricsResponseMetricsRenewedSubscriptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r metricsResponseMetricsRenewedSubscriptionsJSON) RawJSON() string {
	return r.raw
}

// Type of the metric, useful to know the unit or format of the value.
type MetricsResponseMetricsRenewedSubscriptionsType string

const (
	MetricsResponseMetricsRenewedSubscriptionsTypeScalar   MetricsResponseMetricsRenewedSubscriptionsType = "scalar"
	MetricsResponseMetricsRenewedSubscriptionsTypeCurrency MetricsResponseMetricsRenewedSubscriptionsType = "currency"
)

func (r MetricsResponseMetricsRenewedSubscriptionsType) IsKnown() bool {
	switch r {
	case MetricsResponseMetricsRenewedSubscriptionsTypeScalar, MetricsResponseMetricsRenewedSubscriptionsTypeCurrency:
		return true
	}
	return false
}

// Information about a metric.
type MetricsResponseMetricsRenewedSubscriptionsRevenue struct {
	// Human-readable name for the metric.
	DisplayName string `json:"display_name,required"`
	// Unique identifier for the metric.
	Slug string `json:"slug,required"`
	// Type of the metric, useful to know the unit or format of the value.
	Type MetricsResponseMetricsRenewedSubscriptionsRevenueType `json:"type,required"`
	JSON metricsResponseMetricsRenewedSubscriptionsRevenueJSON `json:"-"`
}

// metricsResponseMetricsRenewedSubscriptionsRevenueJSON contains the JSON metadata
// for the struct [MetricsResponseMetricsRenewedSubscriptionsRevenue]
type metricsResponseMetricsRenewedSubscriptionsRevenueJSON struct {
	DisplayName apijson.Field
	Slug        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MetricsResponseMetricsRenewedSubscriptionsRevenue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r metricsResponseMetricsRenewedSubscriptionsRevenueJSON) RawJSON() string {
	return r.raw
}

// Type of the metric, useful to know the unit or format of the value.
type MetricsResponseMetricsRenewedSubscriptionsRevenueType string

const (
	MetricsResponseMetricsRenewedSubscriptionsRevenueTypeScalar   MetricsResponseMetricsRenewedSubscriptionsRevenueType = "scalar"
	MetricsResponseMetricsRenewedSubscriptionsRevenueTypeCurrency MetricsResponseMetricsRenewedSubscriptionsRevenueType = "currency"
)

func (r MetricsResponseMetricsRenewedSubscriptionsRevenueType) IsKnown() bool {
	switch r {
	case MetricsResponseMetricsRenewedSubscriptionsRevenueTypeScalar, MetricsResponseMetricsRenewedSubscriptionsRevenueTypeCurrency:
		return true
	}
	return false
}

// Information about a metric.
type MetricsResponseMetricsRevenue struct {
	// Human-readable name for the metric.
	DisplayName string `json:"display_name,required"`
	// Unique identifier for the metric.
	Slug string `json:"slug,required"`
	// Type of the metric, useful to know the unit or format of the value.
	Type MetricsResponseMetricsRevenueType `json:"type,required"`
	JSON metricsResponseMetricsRevenueJSON `json:"-"`
}

// metricsResponseMetricsRevenueJSON contains the JSON metadata for the struct
// [MetricsResponseMetricsRevenue]
type metricsResponseMetricsRevenueJSON struct {
	DisplayName apijson.Field
	Slug        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MetricsResponseMetricsRevenue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r metricsResponseMetricsRevenueJSON) RawJSON() string {
	return r.raw
}

// Type of the metric, useful to know the unit or format of the value.
type MetricsResponseMetricsRevenueType string

const (
	MetricsResponseMetricsRevenueTypeScalar   MetricsResponseMetricsRevenueType = "scalar"
	MetricsResponseMetricsRevenueTypeCurrency MetricsResponseMetricsRevenueType = "currency"
)

func (r MetricsResponseMetricsRevenueType) IsKnown() bool {
	switch r {
	case MetricsResponseMetricsRevenueTypeScalar, MetricsResponseMetricsRevenueTypeCurrency:
		return true
	}
	return false
}

type MetricsResponsePeriod struct {
	ActiveSubscriptions         int64 `json:"active_subscriptions,required"`
	AverageOrderValue           int64 `json:"average_order_value,required"`
	MonthlyRecurringRevenue     int64 `json:"monthly_recurring_revenue,required"`
	NewSubscriptions            int64 `json:"new_subscriptions,required"`
	NewSubscriptionsRevenue     int64 `json:"new_subscriptions_revenue,required"`
	OneTimeProducts             int64 `json:"one_time_products,required"`
	OneTimeProductsRevenue      int64 `json:"one_time_products_revenue,required"`
	Orders                      int64 `json:"orders,required"`
	RenewedSubscriptions        int64 `json:"renewed_subscriptions,required"`
	RenewedSubscriptionsRevenue int64 `json:"renewed_subscriptions_revenue,required"`
	Revenue                     int64 `json:"revenue,required"`
	// Timestamp of this period data.
	Timestamp time.Time                 `json:"timestamp,required" format:"date-time"`
	JSON      metricsResponsePeriodJSON `json:"-"`
}

// metricsResponsePeriodJSON contains the JSON metadata for the struct
// [MetricsResponsePeriod]
type metricsResponsePeriodJSON struct {
	ActiveSubscriptions         apijson.Field
	AverageOrderValue           apijson.Field
	MonthlyRecurringRevenue     apijson.Field
	NewSubscriptions            apijson.Field
	NewSubscriptionsRevenue     apijson.Field
	OneTimeProducts             apijson.Field
	OneTimeProductsRevenue      apijson.Field
	Orders                      apijson.Field
	RenewedSubscriptions        apijson.Field
	RenewedSubscriptionsRevenue apijson.Field
	Revenue                     apijson.Field
	Timestamp                   apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *MetricsResponsePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r metricsResponsePeriodJSON) RawJSON() string {
	return r.raw
}

type MetricListParams struct {
	// End date.
	EndDate param.Field[time.Time] `query:"end_date,required" format:"date"`
	// Interval between two timestamps.
	Interval param.Field[MetricListParamsInterval] `query:"interval,required"`
	// Start date.
	StartDate param.Field[time.Time] `query:"start_date,required" format:"date"`
	// Filter by organization ID.
	OrganizationID param.Field[MetricListParamsOrganizationIDUnion] `query:"organization_id" format:"uuid4"`
	// Filter by product ID.
	ProductID param.Field[MetricListParamsProductIDUnion] `query:"product_id" format:"uuid4"`
	// Filter by product price type. `recurring` will filter data corresponding to
	// subscriptions creations or renewals. `one_time` will filter data corresponding
	// to one-time purchases.
	ProductPriceType param.Field[MetricListParamsProductPriceTypeUnion] `query:"product_price_type"`
}

// URLQuery serializes [MetricListParams]'s query parameters as `url.Values`.
func (r MetricListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Interval between two timestamps.
type MetricListParamsInterval string

const (
	MetricListParamsIntervalYear  MetricListParamsInterval = "year"
	MetricListParamsIntervalMonth MetricListParamsInterval = "month"
	MetricListParamsIntervalWeek  MetricListParamsInterval = "week"
	MetricListParamsIntervalDay   MetricListParamsInterval = "day"
	MetricListParamsIntervalHour  MetricListParamsInterval = "hour"
)

func (r MetricListParamsInterval) IsKnown() bool {
	switch r {
	case MetricListParamsIntervalYear, MetricListParamsIntervalMonth, MetricListParamsIntervalWeek, MetricListParamsIntervalDay, MetricListParamsIntervalHour:
		return true
	}
	return false
}

// Filter by organization ID.
//
// Satisfied by [shared.UnionString], [MetricListParamsOrganizationIDArray].
type MetricListParamsOrganizationIDUnion interface {
	ImplementsMetricListParamsOrganizationIDUnion()
}

type MetricListParamsOrganizationIDArray []string

func (r MetricListParamsOrganizationIDArray) ImplementsMetricListParamsOrganizationIDUnion() {}

// Filter by product ID.
//
// Satisfied by [shared.UnionString], [MetricListParamsProductIDArray].
type MetricListParamsProductIDUnion interface {
	ImplementsMetricListParamsProductIDUnion()
}

type MetricListParamsProductIDArray []string

func (r MetricListParamsProductIDArray) ImplementsMetricListParamsProductIDUnion() {}

// Filter by product price type. `recurring` will filter data corresponding to
// subscriptions creations or renewals. `one_time` will filter data corresponding
// to one-time purchases.
//
// Satisfied by [MetricListParamsProductPriceTypeProductPriceType],
// [MetricListParamsProductPriceTypeArray].
type MetricListParamsProductPriceTypeUnion interface {
	implementsMetricListParamsProductPriceTypeUnion()
}

type MetricListParamsProductPriceTypeProductPriceType string

const (
	MetricListParamsProductPriceTypeProductPriceTypeOneTime   MetricListParamsProductPriceTypeProductPriceType = "one_time"
	MetricListParamsProductPriceTypeProductPriceTypeRecurring MetricListParamsProductPriceTypeProductPriceType = "recurring"
)

func (r MetricListParamsProductPriceTypeProductPriceType) IsKnown() bool {
	switch r {
	case MetricListParamsProductPriceTypeProductPriceTypeOneTime, MetricListParamsProductPriceTypeProductPriceTypeRecurring:
		return true
	}
	return false
}

func (r MetricListParamsProductPriceTypeProductPriceType) implementsMetricListParamsProductPriceTypeUnion() {
}

type MetricListParamsProductPriceTypeArray []MetricListParamsProductPriceTypeArray

func (r MetricListParamsProductPriceTypeArray) implementsMetricListParamsProductPriceTypeUnion() {}
