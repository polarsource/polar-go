# Oauth2RequestTokenRequestBody


## Supported Types

### AuthorizationCodeTokenRequest

```go
oauth2RequestTokenRequestBody := operations.CreateOauth2RequestTokenRequestBodyAuthorizationCodeTokenRequest(components.AuthorizationCodeTokenRequest{/* values here */})
```

### RefreshTokenRequest

```go
oauth2RequestTokenRequestBody := operations.CreateOauth2RequestTokenRequestBodyRefreshTokenRequest(components.RefreshTokenRequest{/* values here */})
```

### WebTokenRequest

```go
oauth2RequestTokenRequestBody := operations.CreateOauth2RequestTokenRequestBodyWebTokenRequest(components.WebTokenRequest{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch oauth2RequestTokenRequestBody.Type {
	case operations.Oauth2RequestTokenRequestBodyTypeAuthorizationCodeTokenRequest:
		// oauth2RequestTokenRequestBody.AuthorizationCodeTokenRequest is populated
	case operations.Oauth2RequestTokenRequestBodyTypeRefreshTokenRequest:
		// oauth2RequestTokenRequestBody.RefreshTokenRequest is populated
	case operations.Oauth2RequestTokenRequestBodyTypeWebTokenRequest:
		// oauth2RequestTokenRequestBody.WebTokenRequest is populated
}
```
