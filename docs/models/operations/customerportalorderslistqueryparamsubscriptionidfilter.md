# CustomerPortalOrdersListQueryParamSubscriptionIDFilter

Filter by subscription ID.


## Supported Types

### 

```go
customerPortalOrdersListQueryParamSubscriptionIDFilter := operations.CreateCustomerPortalOrdersListQueryParamSubscriptionIDFilterStr(string{/* values here */})
```

### 

```go
customerPortalOrdersListQueryParamSubscriptionIDFilter := operations.CreateCustomerPortalOrdersListQueryParamSubscriptionIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customerPortalOrdersListQueryParamSubscriptionIDFilter.Type {
	case operations.CustomerPortalOrdersListQueryParamSubscriptionIDFilterTypeStr:
		// customerPortalOrdersListQueryParamSubscriptionIDFilter.Str is populated
	case operations.CustomerPortalOrdersListQueryParamSubscriptionIDFilterTypeArrayOfStr:
		// customerPortalOrdersListQueryParamSubscriptionIDFilter.ArrayOfStr is populated
}
```
