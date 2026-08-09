# CommittedAnnualRecurringRevenue


## Supported Types

### 

```go
committedAnnualRecurringRevenue := components.CreateCommittedAnnualRecurringRevenueInteger(int64{/* values here */})
```

### 

```go
committedAnnualRecurringRevenue := components.CreateCommittedAnnualRecurringRevenueNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch committedAnnualRecurringRevenue.Type {
	case components.CommittedAnnualRecurringRevenueTypeInteger:
		// committedAnnualRecurringRevenue.Integer is populated
	case components.CommittedAnnualRecurringRevenueTypeNumber:
		// committedAnnualRecurringRevenue.Number is populated
}
```
