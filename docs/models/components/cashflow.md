# Cashflow


## Supported Types

### 

```go
cashflow := components.CreateCashflowInteger(int64{/* values here */})
```

### 

```go
cashflow := components.CreateCashflowNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch cashflow.Type {
	case components.CashflowTypeInteger:
		// cashflow.Integer is populated
	case components.CashflowTypeNumber:
		// cashflow.Number is populated
}
```
