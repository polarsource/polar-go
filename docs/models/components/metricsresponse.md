# MetricsResponse

Metrics response schema.


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `Periods`                                                            | [][components.MetricPeriod](../../models/components/metricperiod.md) | :heavy_check_mark:                                                   | List of data for each timestamp.                                     |
| `Totals`                                                             | [components.MetricsTotals](../../models/components/metricstotals.md) | :heavy_check_mark:                                                   | N/A                                                                  |
| `Metrics`                                                            | [components.Metrics](../../models/components/metrics.md)             | :heavy_check_mark:                                                   | N/A                                                                  |