# DiscountIDFilter

Filter by discount ID.


## Supported Types

### 

```go
discountIDFilter := operations.CreateDiscountIDFilterStr(string{/* values here */})
```

### 

```go
discountIDFilter := operations.CreateDiscountIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch discountIDFilter.Type {
	case operations.DiscountIDFilterTypeStr:
		// discountIDFilter.Str is populated
	case operations.DiscountIDFilterTypeArrayOfStr:
		// discountIDFilter.ArrayOfStr is populated
}
```
