# RefundsListQueryParamCustomerIDFilter

Filter by customer ID.


## Supported Types

### 

```go
refundsListQueryParamCustomerIDFilter := operations.CreateRefundsListQueryParamCustomerIDFilterStr(string{/* values here */})
```

### 

```go
refundsListQueryParamCustomerIDFilter := operations.CreateRefundsListQueryParamCustomerIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch refundsListQueryParamCustomerIDFilter.Type {
	case operations.RefundsListQueryParamCustomerIDFilterTypeStr:
		// refundsListQueryParamCustomerIDFilter.Str is populated
	case operations.RefundsListQueryParamCustomerIDFilterTypeArrayOfStr:
		// refundsListQueryParamCustomerIDFilter.ArrayOfStr is populated
}
```
