# SeatCustomers


## Supported Types

### 

```go
seatCustomers := components.CreateSeatCustomersInteger(int64{/* values here */})
```

### 

```go
seatCustomers := components.CreateSeatCustomersNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch seatCustomers.Type {
	case components.SeatCustomersTypeInteger:
		// seatCustomers.Integer is populated
	case components.SeatCustomersTypeNumber:
		// seatCustomers.Number is populated
}
```
