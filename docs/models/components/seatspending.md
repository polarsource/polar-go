# SeatsPending


## Supported Types

### 

```go
seatsPending := components.CreateSeatsPendingInteger(int64{/* values here */})
```

### 

```go
seatsPending := components.CreateSeatsPendingNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch seatsPending.Type {
	case components.SeatsPendingTypeInteger:
		// seatsPending.Integer is populated
	case components.SeatsPendingTypeNumber:
		// seatsPending.Number is populated
}
```
