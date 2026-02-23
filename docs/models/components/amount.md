# Amount

The amount in cents.


## Supported Types

### 

```go
amount := components.CreateAmountNumber(float64{/* values here */})
```

### 

```go
amount := components.CreateAmountStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch amount.Type {
	case components.AmountTypeNumber:
		// amount.Number is populated
	case components.AmountTypeStr:
		// amount.Str is populated
}
```
