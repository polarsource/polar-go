# NewSubscriptionsNetRevenue


## Supported Types

### 

```go
newSubscriptionsNetRevenue := components.CreateNewSubscriptionsNetRevenueInteger(int64{/* values here */})
```

### 

```go
newSubscriptionsNetRevenue := components.CreateNewSubscriptionsNetRevenueNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch newSubscriptionsNetRevenue.Type {
	case components.NewSubscriptionsNetRevenueTypeInteger:
		// newSubscriptionsNetRevenue.Integer is populated
	case components.NewSubscriptionsNetRevenueTypeNumber:
		// newSubscriptionsNetRevenue.Number is populated
}
```
