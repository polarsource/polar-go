# LLMMetadata


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `Vendor`                                                      | `string`                                                      | :heavy_check_mark:                                            | The vendor of the event.                                      |
| `Model`                                                       | `string`                                                      | :heavy_check_mark:                                            | The model used for the event.                                 |
| `Prompt`                                                      | `*string`                                                     | :heavy_minus_sign:                                            | The LLM prompt used for the event.                            |
| `Response`                                                    | `*string`                                                     | :heavy_minus_sign:                                            | The LLM response used for the event.                          |
| `InputTokens`                                                 | `int64`                                                       | :heavy_check_mark:                                            | The number of LLM input tokens used for the event.            |
| `CachedInputTokens`                                           | `*int64`                                                      | :heavy_minus_sign:                                            | The number of LLM cached tokens that were used for the event. |
| `OutputTokens`                                                | `int64`                                                       | :heavy_check_mark:                                            | The number of LLM output tokens used for the event.           |
| `TotalTokens`                                                 | `int64`                                                       | :heavy_check_mark:                                            | The total number of LLM tokens used for the event.            |