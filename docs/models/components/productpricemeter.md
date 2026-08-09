# ProductPriceMeter

A meter associated to a metered price.


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ID`                                                         | `string`                                                     | :heavy_check_mark:                                           | The ID of the object.                                        |
| `Name`                                                       | `string`                                                     | :heavy_check_mark:                                           | The name of the meter.                                       |
| `Unit`                                                       | [components.MeterUnit](../../models/components/meterunit.md) | :heavy_check_mark:                                           | N/A                                                          |
| `CustomLabel`                                                | `*string`                                                    | :heavy_check_mark:                                           | The label for the custom unit.                               |
| `CustomMultiplier`                                           | `*int64`                                                     | :heavy_check_mark:                                           | The multiplier to convert from base unit to display scale.   |