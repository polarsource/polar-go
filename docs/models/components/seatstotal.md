# SeatsTotal


## Supported Types

### 

```go
seatsTotal := components.CreateSeatsTotalInteger(int64{/* values here */})
```

### 

```go
seatsTotal := components.CreateSeatsTotalNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch seatsTotal.Type {
	case components.SeatsTotalTypeInteger:
		// seatsTotal.Integer is populated
	case components.SeatsTotalTypeNumber:
		// seatsTotal.Number is populated
}
```
