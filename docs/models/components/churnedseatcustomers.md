# ChurnedSeatCustomers


## Supported Types

### 

```go
churnedSeatCustomers := components.CreateChurnedSeatCustomersInteger(int64{/* values here */})
```

### 

```go
churnedSeatCustomers := components.CreateChurnedSeatCustomersNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch churnedSeatCustomers.Type {
	case components.ChurnedSeatCustomersTypeInteger:
		// churnedSeatCustomers.Integer is populated
	case components.ChurnedSeatCustomersTypeNumber:
		// churnedSeatCustomers.Number is populated
}
```
