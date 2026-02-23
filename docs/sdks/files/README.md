# Files

## Overview

### Available Operations

* [List](#list) - List Files
* [Create](#create) - Create File
* [Uploaded](#uploaded) - Complete File Upload
* [Delete](#delete) - Delete File
* [Update](#update) - Update File

## List

List files.

**Scopes**: `files:read` `files:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="files:list" method="get" path="/v1/files/" -->
```go
package main

import(
	"context"
	"os"
	polargo "github.com/polarsource/polar-go"
	"github.com/polarsource/polar-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := polargo.New(
        polargo.WithSecurity(os.Getenv("POLAR_ACCESS_TOKEN")),
    )

    res, err := s.Files.List(ctx, polargo.Pointer(operations.CreateFilesListQueryParamOrganizationIDFilterStr(
        "1dbfc517-0bbf-4301-9ba8-555ca42b9737",
    )), nil, polargo.Pointer[int64](1), polargo.Pointer[int64](10))
    if err != nil {
        log.Fatal(err)
    }
    if res.ListResourceFileRead != nil {
        for {
            // handle items

            res, err = res.Next()

            if err != nil {
                // handle error
            }

            if res == nil {
                break
            }
        }
    }
}
```

### Parameters

| Parameter                                                                                                                 | Type                                                                                                                      | Required                                                                                                                  | Description                                                                                                               |
| ------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                                                     | :heavy_check_mark:                                                                                                        | The context to use for the request.                                                                                       |
| `organizationID`                                                                                                          | [*operations.FilesListQueryParamOrganizationIDFilter](../../models/operations/fileslistqueryparamorganizationidfilter.md) | :heavy_minus_sign:                                                                                                        | Filter by organization ID.                                                                                                |
| `ids`                                                                                                                     | [*operations.FileIDFilter](../../models/operations/fileidfilter.md)                                                       | :heavy_minus_sign:                                                                                                        | Filter by file ID.                                                                                                        |
| `page`                                                                                                                    | **int64*                                                                                                                  | :heavy_minus_sign:                                                                                                        | Page number, defaults to 1.                                                                                               |
| `limit`                                                                                                                   | **int64*                                                                                                                  | :heavy_minus_sign:                                                                                                        | Size of a page, defaults to 10. Maximum is 100.                                                                           |
| `opts`                                                                                                                    | [][operations.Option](../../models/operations/option.md)                                                                  | :heavy_minus_sign:                                                                                                        | The options for this request.                                                                                             |

### Response

**[*operations.FilesListResponse](../../models/operations/fileslistresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Create

Create a file.

**Scopes**: `files:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="files:create" method="post" path="/v1/files/" -->
```go
package main

import(
	"context"
	"os"
	polargo "github.com/polarsource/polar-go"
	"github.com/polarsource/polar-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := polargo.New(
        polargo.WithSecurity(os.Getenv("POLAR_ACCESS_TOKEN")),
    )

    res, err := s.Files.Create(ctx, components.CreateFileCreateDownloadable(
        components.DownloadableFileCreate{
            OrganizationID: polargo.Pointer("1dbfc517-0bbf-4301-9ba8-555ca42b9737"),
            Name: "<value>",
            MimeType: "<value>",
            Size: 612128,
            Upload: components.S3FileCreateMultipart{
                Parts: []components.S3FileCreatePart{},
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.FileUpload != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |
| `request`                                                      | [components.FileCreate](../../models/components/filecreate.md) | :heavy_check_mark:                                             | The request object to use for the request.                     |
| `opts`                                                         | [][operations.Option](../../models/operations/option.md)       | :heavy_minus_sign:                                             | The options for this request.                                  |

### Response

**[*operations.FilesCreateResponse](../../models/operations/filescreateresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Uploaded

Complete a file upload.

**Scopes**: `files:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="files:uploaded" method="post" path="/v1/files/{id}/uploaded" -->
```go
package main

import(
	"context"
	"os"
	polargo "github.com/polarsource/polar-go"
	"github.com/polarsource/polar-go/models/components"
	"log"
	"github.com/polarsource/polar-go/models/operations"
)

func main() {
    ctx := context.Background()

    s := polargo.New(
        polargo.WithSecurity(os.Getenv("POLAR_ACCESS_TOKEN")),
    )

    res, err := s.Files.Uploaded(ctx, "<value>", components.FileUploadCompleted{
        ID: "<id>",
        Path: "/boot",
        Parts: []components.S3FileUploadCompletedPart{
            components.S3FileUploadCompletedPart{
                Number: 979613,
                ChecksumEtag: "<value>",
                ChecksumSha256Base64: polargo.Pointer("<value>"),
            },
            components.S3FileUploadCompletedPart{
                Number: 979613,
                ChecksumEtag: "<value>",
                ChecksumSha256Base64: polargo.Pointer("<value>"),
            },
            components.S3FileUploadCompletedPart{
                Number: 979613,
                ChecksumEtag: "<value>",
                ChecksumSha256Base64: polargo.Pointer("<value>"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseFilesUploaded != nil {
        switch res.ResponseFilesUploaded.Type {
            case operations.FilesUploadedResponseFilesUploadedTypeDownloadable:
                // res.ResponseFilesUploaded.DownloadableFileRead is populated
            case operations.FilesUploadedResponseFilesUploadedTypeProductMedia:
                // res.ResponseFilesUploaded.ProductMediaFileRead is populated
            case operations.FilesUploadedResponseFilesUploadedTypeOrganizationAvatar:
                // res.ResponseFilesUploaded.OrganizationAvatarFileRead is populated
        }

    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `id`                                                                             | *string*                                                                         | :heavy_check_mark:                                                               | The file ID.                                                                     |
| `fileUploadCompleted`                                                            | [components.FileUploadCompleted](../../models/components/fileuploadcompleted.md) | :heavy_check_mark:                                                               | N/A                                                                              |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.FilesUploadedResponse](../../models/operations/filesuploadedresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.NotPermitted        | 403                           | application/json              |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Delete

Delete a file.

**Scopes**: `files:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="files:delete" method="delete" path="/v1/files/{id}" -->
```go
package main

import(
	"context"
	"os"
	polargo "github.com/polarsource/polar-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := polargo.New(
        polargo.WithSecurity(os.Getenv("POLAR_ACCESS_TOKEN")),
    )

    res, err := s.Files.Delete(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.FilesDeleteResponse](../../models/operations/filesdeleteresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.NotPermitted        | 403                           | application/json              |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Update

Update a file.

**Scopes**: `files:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="files:update" method="patch" path="/v1/files/{id}" -->
```go
package main

import(
	"context"
	"os"
	polargo "github.com/polarsource/polar-go"
	"github.com/polarsource/polar-go/models/components"
	"log"
	"github.com/polarsource/polar-go/models/operations"
)

func main() {
    ctx := context.Background()

    s := polargo.New(
        polargo.WithSecurity(os.Getenv("POLAR_ACCESS_TOKEN")),
    )

    res, err := s.Files.Update(ctx, "<value>", components.FilePatch{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseFilesUpdate != nil {
        switch res.ResponseFilesUpdate.Type {
            case operations.FilesUpdateResponseFilesUpdateTypeDownloadable:
                // res.ResponseFilesUpdate.DownloadableFileRead is populated
            case operations.FilesUpdateResponseFilesUpdateTypeProductMedia:
                // res.ResponseFilesUpdate.ProductMediaFileRead is populated
            case operations.FilesUpdateResponseFilesUpdateTypeOrganizationAvatar:
                // res.ResponseFilesUpdate.OrganizationAvatarFileRead is populated
        }

    }
}
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `id`                                                         | *string*                                                     | :heavy_check_mark:                                           | The file ID.                                                 |
| `filePatch`                                                  | [components.FilePatch](../../models/components/filepatch.md) | :heavy_check_mark:                                           | N/A                                                          |
| `opts`                                                       | [][operations.Option](../../models/operations/option.md)     | :heavy_minus_sign:                                           | The options for this request.                                |

### Response

**[*operations.FilesUpdateResponse](../../models/operations/filesupdateresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.NotPermitted        | 403                           | application/json              |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |