# TrialCommittedMonthlyRecurringRevenue


## Supported Types

### 

```go
trialCommittedMonthlyRecurringRevenue := components.CreateTrialCommittedMonthlyRecurringRevenueInteger(int64{/* values here */})
```

### 

```go
trialCommittedMonthlyRecurringRevenue := components.CreateTrialCommittedMonthlyRecurringRevenueNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch trialCommittedMonthlyRecurringRevenue.Type {
	case components.TrialCommittedMonthlyRecurringRevenueTypeInteger:
		// trialCommittedMonthlyRecurringRevenue.Integer is populated
	case components.TrialCommittedMonthlyRecurringRevenueTypeNumber:
		// trialCommittedMonthlyRecurringRevenue.Number is populated
}
```
