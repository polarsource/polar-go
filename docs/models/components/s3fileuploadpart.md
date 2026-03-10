# S3FileUploadPart


## Fields

| Field                                     | Type                                      | Required                                  | Description                               |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| `Number`                                  | `int64`                                   | :heavy_check_mark:                        | N/A                                       |
| `ChunkStart`                              | `int64`                                   | :heavy_check_mark:                        | N/A                                       |
| `ChunkEnd`                                | `int64`                                   | :heavy_check_mark:                        | N/A                                       |
| `ChecksumSha256Base64`                    | `*string`                                 | :heavy_minus_sign:                        | N/A                                       |
| `URL`                                     | `string`                                  | :heavy_check_mark:                        | N/A                                       |
| `ExpiresAt`                               | [time.Time](https://pkg.go.dev/time#Time) | :heavy_check_mark:                        | N/A                                       |
| `Headers`                                 | map[string]`string`                       | :heavy_minus_sign:                        | N/A                                       |