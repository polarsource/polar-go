# Oauth2AuthorizeResponseOauth2Authorize

Successful Response


## Supported Types

### AuthorizeResponseUser

```go
oauth2AuthorizeResponseOauth2Authorize := operations.CreateOauth2AuthorizeResponseOauth2AuthorizeUser(components.AuthorizeResponseUser{/* values here */})
```

### AuthorizeResponseOrganization

```go
oauth2AuthorizeResponseOauth2Authorize := operations.CreateOauth2AuthorizeResponseOauth2AuthorizeOrganization(components.AuthorizeResponseOrganization{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch oauth2AuthorizeResponseOauth2Authorize.Type {
	case operations.Oauth2AuthorizeResponseOauth2AuthorizeTypeUser:
		// oauth2AuthorizeResponseOauth2Authorize.AuthorizeResponseUser is populated
	case operations.Oauth2AuthorizeResponseOauth2AuthorizeTypeOrganization:
		// oauth2AuthorizeResponseOauth2Authorize.AuthorizeResponseOrganization is populated
}
```
