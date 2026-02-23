# CheckoutsConversion


## Supported Types

### 

```go
checkoutsConversion := components.CreateCheckoutsConversionInteger(int64{/* values here */})
```

### 

```go
checkoutsConversion := components.CreateCheckoutsConversionNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch checkoutsConversion.Type {
	case components.CheckoutsConversionTypeInteger:
		// checkoutsConversion.Integer is populated
	case components.CheckoutsConversionTypeNumber:
		// checkoutsConversion.Number is populated
}
```
