# AverageSeatsPerCustomer


## Supported Types

### 

```go
averageSeatsPerCustomer := components.CreateAverageSeatsPerCustomerInteger(int64{/* values here */})
```

### 

```go
averageSeatsPerCustomer := components.CreateAverageSeatsPerCustomerNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch averageSeatsPerCustomer.Type {
	case components.AverageSeatsPerCustomerTypeInteger:
		// averageSeatsPerCustomer.Integer is populated
	case components.AverageSeatsPerCustomerTypeNumber:
		// averageSeatsPerCustomer.Number is populated
}
```
