# SeatsClaimed


## Supported Types

### 

```go
seatsClaimed := components.CreateSeatsClaimedInteger(int64{/* values here */})
```

### 

```go
seatsClaimed := components.CreateSeatsClaimedNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch seatsClaimed.Type {
	case components.SeatsClaimedTypeInteger:
		// seatsClaimed.Integer is populated
	case components.SeatsClaimedTypeNumber:
		// seatsClaimed.Number is populated
}
```
