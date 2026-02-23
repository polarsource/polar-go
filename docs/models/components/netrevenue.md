# NetRevenue


## Supported Types

### 

```go
netRevenue := components.CreateNetRevenueInteger(int64{/* values here */})
```

### 

```go
netRevenue := components.CreateNetRevenueNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch netRevenue.Type {
	case components.NetRevenueTypeInteger:
		// netRevenue.Integer is populated
	case components.NetRevenueTypeNumber:
		// netRevenue.Number is populated
}
```
