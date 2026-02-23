# RefundsListQueryParamExternalCustomerIDFilter

Filter by customer external ID.


## Supported Types

### 

```go
refundsListQueryParamExternalCustomerIDFilter := operations.CreateRefundsListQueryParamExternalCustomerIDFilterStr(string{/* values here */})
```

### 

```go
refundsListQueryParamExternalCustomerIDFilter := operations.CreateRefundsListQueryParamExternalCustomerIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch refundsListQueryParamExternalCustomerIDFilter.Type {
	case operations.RefundsListQueryParamExternalCustomerIDFilterTypeStr:
		// refundsListQueryParamExternalCustomerIDFilter.Str is populated
	case operations.RefundsListQueryParamExternalCustomerIDFilterTypeArrayOfStr:
		// refundsListQueryParamExternalCustomerIDFilter.ArrayOfStr is populated
}
```
