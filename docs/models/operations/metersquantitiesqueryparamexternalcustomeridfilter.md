# MetersQuantitiesQueryParamExternalCustomerIDFilter

Filter by external customer ID.


## Supported Types

### 

```go
metersQuantitiesQueryParamExternalCustomerIDFilter := operations.CreateMetersQuantitiesQueryParamExternalCustomerIDFilterStr(string{/* values here */})
```

### 

```go
metersQuantitiesQueryParamExternalCustomerIDFilter := operations.CreateMetersQuantitiesQueryParamExternalCustomerIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metersQuantitiesQueryParamExternalCustomerIDFilter.Type {
	case operations.MetersQuantitiesQueryParamExternalCustomerIDFilterTypeStr:
		// metersQuantitiesQueryParamExternalCustomerIDFilter.Str is populated
	case operations.MetersQuantitiesQueryParamExternalCustomerIDFilterTypeArrayOfStr:
		// metersQuantitiesQueryParamExternalCustomerIDFilter.ArrayOfStr is populated
}
```
