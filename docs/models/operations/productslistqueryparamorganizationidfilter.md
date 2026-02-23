# ProductsListQueryParamOrganizationIDFilter

Filter by organization ID.


## Supported Types

### 

```go
productsListQueryParamOrganizationIDFilter := operations.CreateProductsListQueryParamOrganizationIDFilterStr(string{/* values here */})
```

### 

```go
productsListQueryParamOrganizationIDFilter := operations.CreateProductsListQueryParamOrganizationIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch productsListQueryParamOrganizationIDFilter.Type {
	case operations.ProductsListQueryParamOrganizationIDFilterTypeStr:
		// productsListQueryParamOrganizationIDFilter.Str is populated
	case operations.ProductsListQueryParamOrganizationIDFilterTypeArrayOfStr:
		// productsListQueryParamOrganizationIDFilter.ArrayOfStr is populated
}
```
