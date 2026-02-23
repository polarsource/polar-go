# UnitAmount

The price per unit in cents. Supports up to 12 decimal places.


## Supported Types

### 

```go
unitAmount := components.CreateUnitAmountNumber(float64{/* values here */})
```

### 

```go
unitAmount := components.CreateUnitAmountStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch unitAmount.Type {
	case components.UnitAmountTypeNumber:
		// unitAmount.Number is populated
	case components.UnitAmountTypeStr:
		// unitAmount.Str is populated
}
```
