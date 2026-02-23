# NetCumulativeRevenue


## Supported Types

### 

```go
netCumulativeRevenue := components.CreateNetCumulativeRevenueInteger(int64{/* values here */})
```

### 

```go
netCumulativeRevenue := components.CreateNetCumulativeRevenueNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch netCumulativeRevenue.Type {
	case components.NetCumulativeRevenueTypeInteger:
		// netCumulativeRevenue.Integer is populated
	case components.NetCumulativeRevenueTypeNumber:
		// netCumulativeRevenue.Number is populated
}
```
