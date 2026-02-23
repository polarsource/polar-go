# CheckoutLinksListQueryParamOrganizationIDFilter

Filter by organization ID.


## Supported Types

### 

```go
checkoutLinksListQueryParamOrganizationIDFilter := operations.CreateCheckoutLinksListQueryParamOrganizationIDFilterStr(string{/* values here */})
```

### 

```go
checkoutLinksListQueryParamOrganizationIDFilter := operations.CreateCheckoutLinksListQueryParamOrganizationIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch checkoutLinksListQueryParamOrganizationIDFilter.Type {
	case operations.CheckoutLinksListQueryParamOrganizationIDFilterTypeStr:
		// checkoutLinksListQueryParamOrganizationIDFilter.Str is populated
	case operations.CheckoutLinksListQueryParamOrganizationIDFilterTypeArrayOfStr:
		// checkoutLinksListQueryParamOrganizationIDFilter.ArrayOfStr is populated
}
```
