# MetersQuantitiesQueryParamCustomerIDFilter

Filter by customer ID.


## Supported Types

### 

```go
metersQuantitiesQueryParamCustomerIDFilter := operations.CreateMetersQuantitiesQueryParamCustomerIDFilterStr(string{/* values here */})
```

### 

```go
metersQuantitiesQueryParamCustomerIDFilter := operations.CreateMetersQuantitiesQueryParamCustomerIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metersQuantitiesQueryParamCustomerIDFilter.Type {
	case operations.MetersQuantitiesQueryParamCustomerIDFilterTypeStr:
		// metersQuantitiesQueryParamCustomerIDFilter.Str is populated
	case operations.MetersQuantitiesQueryParamCustomerIDFilterTypeArrayOfStr:
		// metersQuantitiesQueryParamCustomerIDFilter.ArrayOfStr is populated
}
```
