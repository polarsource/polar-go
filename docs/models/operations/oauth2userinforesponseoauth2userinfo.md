# Oauth2UserinfoResponseOauth2Userinfo

Successful Response


## Supported Types

### UserInfoUser

```go
oauth2UserinfoResponseOauth2Userinfo := operations.CreateOauth2UserinfoResponseOauth2UserinfoUserInfoUser(components.UserInfoUser{/* values here */})
```

### UserInfoOrganization

```go
oauth2UserinfoResponseOauth2Userinfo := operations.CreateOauth2UserinfoResponseOauth2UserinfoUserInfoOrganization(components.UserInfoOrganization{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch oauth2UserinfoResponseOauth2Userinfo.Type {
	case operations.Oauth2UserinfoResponseOauth2UserinfoTypeUserInfoUser:
		// oauth2UserinfoResponseOauth2Userinfo.UserInfoUser is populated
	case operations.Oauth2UserinfoResponseOauth2UserinfoTypeUserInfoOrganization:
		// oauth2UserinfoResponseOauth2Userinfo.UserInfoOrganization is populated
}
```
