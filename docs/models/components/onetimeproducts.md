# OneTimeProducts


## Supported Types

### 

```go
oneTimeProducts := components.CreateOneTimeProductsInteger(int64{/* values here */})
```

### 

```go
oneTimeProducts := components.CreateOneTimeProductsNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch oneTimeProducts.Type {
	case components.OneTimeProductsTypeInteger:
		// oneTimeProducts.Integer is populated
	case components.OneTimeProductsTypeNumber:
		// oneTimeProducts.Number is populated
}
```
