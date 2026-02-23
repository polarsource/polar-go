# Revenue


## Supported Types

### 

```go
revenue := components.CreateRevenueInteger(int64{/* values here */})
```

### 

```go
revenue := components.CreateRevenueNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch revenue.Type {
	case components.RevenueTypeInteger:
		// revenue.Integer is populated
	case components.RevenueTypeNumber:
		// revenue.Number is populated
}
```
