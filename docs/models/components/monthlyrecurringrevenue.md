# MonthlyRecurringRevenue


## Supported Types

### 

```go
monthlyRecurringRevenue := components.CreateMonthlyRecurringRevenueInteger(int64{/* values here */})
```

### 

```go
monthlyRecurringRevenue := components.CreateMonthlyRecurringRevenueNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch monthlyRecurringRevenue.Type {
	case components.MonthlyRecurringRevenueTypeInteger:
		// monthlyRecurringRevenue.Integer is populated
	case components.MonthlyRecurringRevenueTypeNumber:
		// monthlyRecurringRevenue.Number is populated
}
```
