# Orders


## Supported Types

### 

```go
orders := components.CreateOrdersInteger(int64{/* values here */})
```

### 

```go
orders := components.CreateOrdersNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch orders.Type {
	case components.OrdersTypeInteger:
		// orders.Integer is populated
	case components.OrdersTypeNumber:
		// orders.Number is populated
}
```
