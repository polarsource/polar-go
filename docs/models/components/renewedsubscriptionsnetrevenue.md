# RenewedSubscriptionsNetRevenue


## Supported Types

### 

```go
renewedSubscriptionsNetRevenue := components.CreateRenewedSubscriptionsNetRevenueInteger(int64{/* values here */})
```

### 

```go
renewedSubscriptionsNetRevenue := components.CreateRenewedSubscriptionsNetRevenueNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch renewedSubscriptionsNetRevenue.Type {
	case components.RenewedSubscriptionsNetRevenueTypeInteger:
		// renewedSubscriptionsNetRevenue.Integer is populated
	case components.RenewedSubscriptionsNetRevenueTypeNumber:
		// renewedSubscriptionsNetRevenue.Number is populated
}
```
