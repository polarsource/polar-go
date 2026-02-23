# CheckoutLinksListQueryParamProductIDFilter

Filter by product ID.


## Supported Types

### 

```go
checkoutLinksListQueryParamProductIDFilter := operations.CreateCheckoutLinksListQueryParamProductIDFilterStr(string{/* values here */})
```

### 

```go
checkoutLinksListQueryParamProductIDFilter := operations.CreateCheckoutLinksListQueryParamProductIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch checkoutLinksListQueryParamProductIDFilter.Type {
	case operations.CheckoutLinksListQueryParamProductIDFilterTypeStr:
		// checkoutLinksListQueryParamProductIDFilter.Str is populated
	case operations.CheckoutLinksListQueryParamProductIDFilterTypeArrayOfStr:
		// checkoutLinksListQueryParamProductIDFilter.ArrayOfStr is populated
}
```
