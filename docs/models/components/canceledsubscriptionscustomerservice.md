# CanceledSubscriptionsCustomerService


## Supported Types

### 

```go
canceledSubscriptionsCustomerService := components.CreateCanceledSubscriptionsCustomerServiceInteger(int64{/* values here */})
```

### 

```go
canceledSubscriptionsCustomerService := components.CreateCanceledSubscriptionsCustomerServiceNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch canceledSubscriptionsCustomerService.Type {
	case components.CanceledSubscriptionsCustomerServiceTypeInteger:
		// canceledSubscriptionsCustomerService.Integer is populated
	case components.CanceledSubscriptionsCustomerServiceTypeNumber:
		// canceledSubscriptionsCustomerService.Number is populated
}
```
