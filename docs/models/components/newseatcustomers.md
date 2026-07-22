# NewSeatCustomers


## Supported Types

### 

```go
newSeatCustomers := components.CreateNewSeatCustomersInteger(int64{/* values here */})
```

### 

```go
newSeatCustomers := components.CreateNewSeatCustomersNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch newSeatCustomers.Type {
	case components.NewSeatCustomersTypeInteger:
		// newSeatCustomers.Integer is populated
	case components.NewSeatCustomersTypeNumber:
		// newSeatCustomers.Number is populated
}
```
