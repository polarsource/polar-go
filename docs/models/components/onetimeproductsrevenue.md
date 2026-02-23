# OneTimeProductsRevenue


## Supported Types

### 

```go
oneTimeProductsRevenue := components.CreateOneTimeProductsRevenueInteger(int64{/* values here */})
```

### 

```go
oneTimeProductsRevenue := components.CreateOneTimeProductsRevenueNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch oneTimeProductsRevenue.Type {
	case components.OneTimeProductsRevenueTypeInteger:
		// oneTimeProductsRevenue.Integer is populated
	case components.OneTimeProductsRevenueTypeNumber:
		// oneTimeProductsRevenue.Number is populated
}
```
