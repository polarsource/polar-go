# SucceededCheckouts


## Supported Types

### 

```go
succeededCheckouts := components.CreateSucceededCheckoutsInteger(int64{/* values here */})
```

### 

```go
succeededCheckouts := components.CreateSucceededCheckoutsNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch succeededCheckouts.Type {
	case components.SucceededCheckoutsTypeInteger:
		// succeededCheckouts.Integer is populated
	case components.SucceededCheckoutsTypeNumber:
		// succeededCheckouts.Number is populated
}
```
