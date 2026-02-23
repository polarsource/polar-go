# CustomerPortalSubscriptionsListQueryParamProductIDFilter

Filter by product ID.


## Supported Types

### 

```go
customerPortalSubscriptionsListQueryParamProductIDFilter := operations.CreateCustomerPortalSubscriptionsListQueryParamProductIDFilterStr(string{/* values here */})
```

### 

```go
customerPortalSubscriptionsListQueryParamProductIDFilter := operations.CreateCustomerPortalSubscriptionsListQueryParamProductIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customerPortalSubscriptionsListQueryParamProductIDFilter.Type {
	case operations.CustomerPortalSubscriptionsListQueryParamProductIDFilterTypeStr:
		// customerPortalSubscriptionsListQueryParamProductIDFilter.Str is populated
	case operations.CustomerPortalSubscriptionsListQueryParamProductIDFilterTypeArrayOfStr:
		// customerPortalSubscriptionsListQueryParamProductIDFilter.ArrayOfStr is populated
}
```
