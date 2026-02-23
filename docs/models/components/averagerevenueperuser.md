# AverageRevenuePerUser


## Supported Types

### 

```go
averageRevenuePerUser := components.CreateAverageRevenuePerUserInteger(int64{/* values here */})
```

### 

```go
averageRevenuePerUser := components.CreateAverageRevenuePerUserNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch averageRevenuePerUser.Type {
	case components.AverageRevenuePerUserTypeInteger:
		// averageRevenuePerUser.Integer is populated
	case components.AverageRevenuePerUserTypeNumber:
		// averageRevenuePerUser.Number is populated
}
```
