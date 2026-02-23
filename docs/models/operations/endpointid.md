# EndpointID

Filter by webhook endpoint ID.


## Supported Types

### 

```go
endpointID := operations.CreateEndpointIDStr(string{/* values here */})
```

### 

```go
endpointID := operations.CreateEndpointIDArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch endpointID.Type {
	case operations.EndpointIDTypeStr:
		// endpointID.Str is populated
	case operations.EndpointIDTypeArrayOfStr:
		// endpointID.ArrayOfStr is populated
}
```
