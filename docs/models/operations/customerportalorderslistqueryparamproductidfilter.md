# CustomerPortalOrdersListQueryParamProductIDFilter

Filter by product ID.


## Supported Types

### 

```go
customerPortalOrdersListQueryParamProductIDFilter := operations.CreateCustomerPortalOrdersListQueryParamProductIDFilterStr(string{/* values here */})
```

### 

```go
customerPortalOrdersListQueryParamProductIDFilter := operations.CreateCustomerPortalOrdersListQueryParamProductIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customerPortalOrdersListQueryParamProductIDFilter.Type {
	case operations.CustomerPortalOrdersListQueryParamProductIDFilterTypeStr:
		// customerPortalOrdersListQueryParamProductIDFilter.Str is populated
	case operations.CustomerPortalOrdersListQueryParamProductIDFilterTypeArrayOfStr:
		// customerPortalOrdersListQueryParamProductIDFilter.ArrayOfStr is populated
}
```
