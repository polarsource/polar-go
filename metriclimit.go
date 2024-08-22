// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package polar

import (
	"context"
	"net/http"
	"time"

	"github.com/stainless-sdks/polar-go/internal/apijson"
	"github.com/stainless-sdks/polar-go/internal/requestconfig"
	"github.com/stainless-sdks/polar-go/option"
)

// MetricLimitService contains methods and other services that help with
// interacting with the polar API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMetricLimitService] method instead.
type MetricLimitService struct {
	Options []option.RequestOption
}

// NewMetricLimitService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMetricLimitService(opts ...option.RequestOption) (r *MetricLimitService) {
	r = &MetricLimitService{}
	r.Options = opts
	return
}

// Get the interval limits for the metrics endpoint.
func (r *MetricLimitService) Get(ctx context.Context, opts ...option.RequestOption) (res *MetricsLimits, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/metrics/limits"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Date limits to get metrics.
type MetricsLimits struct {
	// Date interval limits to get metrics for each interval.
	Intervals MetricsLimitsIntervals `json:"intervals,required"`
	// Minimum date to get metrics.
	MinDate time.Time         `json:"min_date,required" format:"date"`
	JSON    metricsLimitsJSON `json:"-"`
}

// metricsLimitsJSON contains the JSON metadata for the struct [MetricsLimits]
type metricsLimitsJSON struct {
	Intervals   apijson.Field
	MinDate     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MetricsLimits) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r metricsLimitsJSON) RawJSON() string {
	return r.raw
}

// Date interval limits to get metrics for each interval.
type MetricsLimitsIntervals struct {
	// Date interval limit to get metrics for a given interval.
	Day MetricsLimitsIntervalsDay `json:"day,required"`
	// Date interval limit to get metrics for a given interval.
	Hour MetricsLimitsIntervalsHour `json:"hour,required"`
	// Date interval limit to get metrics for a given interval.
	Month MetricsLimitsIntervalsMonth `json:"month,required"`
	// Date interval limit to get metrics for a given interval.
	Week MetricsLimitsIntervalsWeek `json:"week,required"`
	// Date interval limit to get metrics for a given interval.
	Year MetricsLimitsIntervalsYear `json:"year,required"`
	JSON metricsLimitsIntervalsJSON `json:"-"`
}

// metricsLimitsIntervalsJSON contains the JSON metadata for the struct
// [MetricsLimitsIntervals]
type metricsLimitsIntervalsJSON struct {
	Day         apijson.Field
	Hour        apijson.Field
	Month       apijson.Field
	Week        apijson.Field
	Year        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MetricsLimitsIntervals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r metricsLimitsIntervalsJSON) RawJSON() string {
	return r.raw
}

// Date interval limit to get metrics for a given interval.
type MetricsLimitsIntervalsDay struct {
	// Maximum number of days for this interval.
	MaxDays int64                         `json:"max_days,required"`
	JSON    metricsLimitsIntervalsDayJSON `json:"-"`
}

// metricsLimitsIntervalsDayJSON contains the JSON metadata for the struct
// [MetricsLimitsIntervalsDay]
type metricsLimitsIntervalsDayJSON struct {
	MaxDays     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MetricsLimitsIntervalsDay) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r metricsLimitsIntervalsDayJSON) RawJSON() string {
	return r.raw
}

// Date interval limit to get metrics for a given interval.
type MetricsLimitsIntervalsHour struct {
	// Maximum number of days for this interval.
	MaxDays int64                          `json:"max_days,required"`
	JSON    metricsLimitsIntervalsHourJSON `json:"-"`
}

// metricsLimitsIntervalsHourJSON contains the JSON metadata for the struct
// [MetricsLimitsIntervalsHour]
type metricsLimitsIntervalsHourJSON struct {
	MaxDays     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MetricsLimitsIntervalsHour) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r metricsLimitsIntervalsHourJSON) RawJSON() string {
	return r.raw
}

// Date interval limit to get metrics for a given interval.
type MetricsLimitsIntervalsMonth struct {
	// Maximum number of days for this interval.
	MaxDays int64                           `json:"max_days,required"`
	JSON    metricsLimitsIntervalsMonthJSON `json:"-"`
}

// metricsLimitsIntervalsMonthJSON contains the JSON metadata for the struct
// [MetricsLimitsIntervalsMonth]
type metricsLimitsIntervalsMonthJSON struct {
	MaxDays     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MetricsLimitsIntervalsMonth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r metricsLimitsIntervalsMonthJSON) RawJSON() string {
	return r.raw
}

// Date interval limit to get metrics for a given interval.
type MetricsLimitsIntervalsWeek struct {
	// Maximum number of days for this interval.
	MaxDays int64                          `json:"max_days,required"`
	JSON    metricsLimitsIntervalsWeekJSON `json:"-"`
}

// metricsLimitsIntervalsWeekJSON contains the JSON metadata for the struct
// [MetricsLimitsIntervalsWeek]
type metricsLimitsIntervalsWeekJSON struct {
	MaxDays     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MetricsLimitsIntervalsWeek) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r metricsLimitsIntervalsWeekJSON) RawJSON() string {
	return r.raw
}

// Date interval limit to get metrics for a given interval.
type MetricsLimitsIntervalsYear struct {
	// Maximum number of days for this interval.
	MaxDays int64                          `json:"max_days,required"`
	JSON    metricsLimitsIntervalsYearJSON `json:"-"`
}

// metricsLimitsIntervalsYearJSON contains the JSON metadata for the struct
// [MetricsLimitsIntervalsYear]
type metricsLimitsIntervalsYearJSON struct {
	MaxDays     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MetricsLimitsIntervalsYear) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r metricsLimitsIntervalsYearJSON) RawJSON() string {
	return r.raw
}
