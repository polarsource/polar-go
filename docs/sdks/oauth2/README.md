# Oauth2
(*Oauth2*)

## Overview

### Available Operations

* [Authorize](#authorize) - Authorize
* [Token](#token) - Request Token
* [Revoke](#revoke) - Revoke Token
* [Introspect](#introspect) - Introspect Token
* [Userinfo](#userinfo) - Get User Info

## Authorize

Authorize

### Example Usage

```go
package main

import(
	"context"
	"os"
	"polar"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := polar.New(
        polar.WithSecurity(os.Getenv("POLAR_ACCESS_TOKEN")),
    )

    res, err := s.Oauth2.Authorize(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseOauth2Authorize != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.Oauth2AuthorizeResponse](../../models/operations/oauth2authorizeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Token

Request an access token using a valid grant.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"polar"
	"polar/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := polar.New(
        polar.WithSecurity(os.Getenv("POLAR_ACCESS_TOKEN")),
    )

    res, err := s.Oauth2.Token(ctx, operations.CreateOauth2RequestTokenRequestBodyOnev11oauth21tokenPostXComponentsAuthorizationCodeTokenRequest(
        components.Onev11oauth21tokenPostXComponentsAuthorizationCodeTokenRequest{
            ClientID: "<id>",
            ClientSecret: "<value>",
            Code: "<value>",
            RedirectURI: "https://talkative-barracks.com",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.TokenResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.Oauth2RequestTokenRequestBody](../../models/operations/oauth2requesttokenrequestbody.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.Oauth2RequestTokenResponse](../../models/operations/oauth2requesttokenresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Revoke

Revoke an access token or a refresh token.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"polar"
	"polar/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := polar.New(
        polar.WithSecurity(os.Getenv("POLAR_ACCESS_TOKEN")),
    )

    res, err := s.Oauth2.Revoke(ctx, operations.Oauth2RevokeTokenRevokeTokenRequest{
        Token: "<value>",
        ClientID: "<id>",
        ClientSecret: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RevokeTokenResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.Oauth2RevokeTokenRevokeTokenRequest](../../models/operations/oauth2revoketokenrevoketokenrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.Oauth2RevokeTokenResponse](../../models/operations/oauth2revoketokenresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Introspect

Get information about an access token.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"polar"
	"polar/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := polar.New(
        polar.WithSecurity(os.Getenv("POLAR_ACCESS_TOKEN")),
    )

    res, err := s.Oauth2.Introspect(ctx, operations.Oauth2IntrospectTokenIntrospectTokenRequest{
        Token: "<value>",
        ClientID: "<id>",
        ClientSecret: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.IntrospectTokenResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.Oauth2IntrospectTokenIntrospectTokenRequest](../../models/operations/oauth2introspecttokenintrospecttokenrequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |
| `opts`                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                         | :heavy_minus_sign:                                                                                                               | The options for this request.                                                                                                    |

### Response

**[*operations.Oauth2IntrospectTokenResponse](../../models/operations/oauth2introspecttokenresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Userinfo

Get information about the authenticated user.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"polar"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := polar.New(
        polar.WithSecurity(os.Getenv("POLAR_ACCESS_TOKEN")),
    )

    res, err := s.Oauth2.Userinfo(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseOauth2Userinfo != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.Oauth2UserinfoResponse](../../models/operations/oauth2userinforesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |