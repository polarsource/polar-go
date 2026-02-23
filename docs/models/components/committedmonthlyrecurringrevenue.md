# CommittedMonthlyRecurringRevenue


## Supported Types

### 

```go
committedMonthlyRecurringRevenue := components.CreateCommittedMonthlyRecurringRevenueInteger(int64{/* values here */})
```

### 

```go
committedMonthlyRecurringRevenue := components.CreateCommittedMonthlyRecurringRevenueNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch committedMonthlyRecurringRevenue.Type {
	case components.CommittedMonthlyRecurringRevenueTypeInteger:
		// committedMonthlyRecurringRevenue.Integer is populated
	case components.CommittedMonthlyRecurringRevenueTypeNumber:
		// committedMonthlyRecurringRevenue.Number is populated
}
```
