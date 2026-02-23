# Checkouts


## Supported Types

### 

```go
checkouts := components.CreateCheckoutsInteger(int64{/* values here */})
```

### 

```go
checkouts := components.CreateCheckoutsNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch checkouts.Type {
	case components.CheckoutsTypeInteger:
		// checkouts.Integer is populated
	case components.CheckoutsTypeNumber:
		// checkouts.Number is populated
}
```
