# AnnualRecurringRevenue


## Supported Types

### 

```go
annualRecurringRevenue := components.CreateAnnualRecurringRevenueInteger(int64{/* values here */})
```

### 

```go
annualRecurringRevenue := components.CreateAnnualRecurringRevenueNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch annualRecurringRevenue.Type {
	case components.AnnualRecurringRevenueTypeInteger:
		// annualRecurringRevenue.Integer is populated
	case components.AnnualRecurringRevenueTypeNumber:
		// annualRecurringRevenue.Number is populated
}
```
