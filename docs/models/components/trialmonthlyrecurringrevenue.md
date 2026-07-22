# TrialMonthlyRecurringRevenue


## Supported Types

### 

```go
trialMonthlyRecurringRevenue := components.CreateTrialMonthlyRecurringRevenueInteger(int64{/* values here */})
```

### 

```go
trialMonthlyRecurringRevenue := components.CreateTrialMonthlyRecurringRevenueNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch trialMonthlyRecurringRevenue.Type {
	case components.TrialMonthlyRecurringRevenueTypeInteger:
		// trialMonthlyRecurringRevenue.Integer is populated
	case components.TrialMonthlyRecurringRevenueTypeNumber:
		// trialMonthlyRecurringRevenue.Number is populated
}
```
