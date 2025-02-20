# QueryParamProductPriceTypeFilter

Filter by product price type. `recurring` will return orders corresponding to subscriptions creations or renewals. `one_time` will return orders corresponding to one-time purchases.


## Supported Types

### ProductPriceType

```go
queryParamProductPriceTypeFilter := operations.CreateQueryParamProductPriceTypeFilterProductPriceType(components.ProductPriceType{/* values here */})
```

### 

```go
queryParamProductPriceTypeFilter := operations.CreateQueryParamProductPriceTypeFilterArrayOfProductPriceType([]components.ProductPriceType{/* values here */})
```

