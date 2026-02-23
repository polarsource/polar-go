# ProductCreate


## Supported Types

### ProductCreateRecurring

```go
productCreate := components.CreateProductCreateProductCreateRecurring(components.ProductCreateRecurring{/* values here */})
```

### ProductCreateOneTime

```go
productCreate := components.CreateProductCreateProductCreateOneTime(components.ProductCreateOneTime{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch productCreate.Type {
	case components.ProductCreateTypeProductCreateRecurring:
		// productCreate.ProductCreateRecurring is populated
	case components.ProductCreateTypeProductCreateOneTime:
		// productCreate.ProductCreateOneTime is populated
}
```
