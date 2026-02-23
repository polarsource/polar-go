# ProductIDFilter

Filter by product ID.


## Supported Types

### 

```go
productIDFilter := operations.CreateProductIDFilterStr(string{/* values here */})
```

### 

```go
productIDFilter := operations.CreateProductIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch productIDFilter.Type {
	case operations.ProductIDFilterTypeStr:
		// productIDFilter.Str is populated
	case operations.ProductIDFilterTypeArrayOfStr:
		// productIDFilter.ArrayOfStr is populated
}
```
