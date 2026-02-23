# OneTimeProductsNetRevenue


## Supported Types

### 

```go
oneTimeProductsNetRevenue := components.CreateOneTimeProductsNetRevenueInteger(int64{/* values here */})
```

### 

```go
oneTimeProductsNetRevenue := components.CreateOneTimeProductsNetRevenueNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch oneTimeProductsNetRevenue.Type {
	case components.OneTimeProductsNetRevenueTypeInteger:
		// oneTimeProductsNetRevenue.Integer is populated
	case components.OneTimeProductsNetRevenueTypeNumber:
		// oneTimeProductsNetRevenue.Number is populated
}
```
