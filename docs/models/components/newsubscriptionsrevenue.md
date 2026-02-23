# NewSubscriptionsRevenue


## Supported Types

### 

```go
newSubscriptionsRevenue := components.CreateNewSubscriptionsRevenueInteger(int64{/* values here */})
```

### 

```go
newSubscriptionsRevenue := components.CreateNewSubscriptionsRevenueNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch newSubscriptionsRevenue.Type {
	case components.NewSubscriptionsRevenueTypeInteger:
		// newSubscriptionsRevenue.Integer is populated
	case components.NewSubscriptionsRevenueTypeNumber:
		// newSubscriptionsRevenue.Number is populated
}
```
