# OrderIDFilter

Filter by order ID.


## Supported Types

### 

```go
orderIDFilter := operations.CreateOrderIDFilterStr(string{/* values here */})
```

### 

```go
orderIDFilter := operations.CreateOrderIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch orderIDFilter.Type {
	case operations.OrderIDFilterTypeStr:
		// orderIDFilter.Str is populated
	case operations.OrderIDFilterTypeArrayOfStr:
		// orderIDFilter.ArrayOfStr is populated
}
```
