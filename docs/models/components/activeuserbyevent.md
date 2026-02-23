# ActiveUserByEvent


## Supported Types

### 

```go
activeUserByEvent := components.CreateActiveUserByEventInteger(int64{/* values here */})
```

### 

```go
activeUserByEvent := components.CreateActiveUserByEventNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch activeUserByEvent.Type {
	case components.ActiveUserByEventTypeInteger:
		// activeUserByEvent.Integer is populated
	case components.ActiveUserByEventTypeNumber:
		// activeUserByEvent.Number is populated
}
```
