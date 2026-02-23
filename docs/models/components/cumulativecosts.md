# CumulativeCosts


## Supported Types

### 

```go
cumulativeCosts := components.CreateCumulativeCostsInteger(int64{/* values here */})
```

### 

```go
cumulativeCosts := components.CreateCumulativeCostsNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch cumulativeCosts.Type {
	case components.CumulativeCostsTypeInteger:
		// cumulativeCosts.Integer is populated
	case components.CumulativeCostsTypeNumber:
		// cumulativeCosts.Number is populated
}
```
