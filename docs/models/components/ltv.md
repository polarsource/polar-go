# Ltv


## Supported Types

### 

```go
ltv := components.CreateLtvInteger(int64{/* values here */})
```

### 

```go
ltv := components.CreateLtvNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch ltv.Type {
	case components.LtvTypeInteger:
		// ltv.Integer is populated
	case components.LtvTypeNumber:
		// ltv.Number is populated
}
```
