# SeatUtilizationRate


## Supported Types

### 

```go
seatUtilizationRate := components.CreateSeatUtilizationRateInteger(int64{/* values here */})
```

### 

```go
seatUtilizationRate := components.CreateSeatUtilizationRateNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch seatUtilizationRate.Type {
	case components.SeatUtilizationRateTypeInteger:
		// seatUtilizationRate.Integer is populated
	case components.SeatUtilizationRateTypeNumber:
		// seatUtilizationRate.Number is populated
}
```
