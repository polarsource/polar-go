# RenewedSubscriptionsRevenue


## Supported Types

### 

```go
renewedSubscriptionsRevenue := components.CreateRenewedSubscriptionsRevenueInteger(int64{/* values here */})
```

### 

```go
renewedSubscriptionsRevenue := components.CreateRenewedSubscriptionsRevenueNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch renewedSubscriptionsRevenue.Type {
	case components.RenewedSubscriptionsRevenueTypeInteger:
		// renewedSubscriptionsRevenue.Integer is populated
	case components.RenewedSubscriptionsRevenueTypeNumber:
		// renewedSubscriptionsRevenue.Number is populated
}
```
