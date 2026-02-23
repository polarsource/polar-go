# CumulativeRevenue


## Supported Types

### 

```go
cumulativeRevenue := components.CreateCumulativeRevenueInteger(int64{/* values here */})
```

### 

```go
cumulativeRevenue := components.CreateCumulativeRevenueNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch cumulativeRevenue.Type {
	case components.CumulativeRevenueTypeInteger:
		// cumulativeRevenue.Integer is populated
	case components.CumulativeRevenueTypeNumber:
		// cumulativeRevenue.Number is populated
}
```
