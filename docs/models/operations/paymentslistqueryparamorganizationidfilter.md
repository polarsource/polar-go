# PaymentsListQueryParamOrganizationIDFilter

Filter by organization ID.


## Supported Types

### 

```go
paymentsListQueryParamOrganizationIDFilter := operations.CreatePaymentsListQueryParamOrganizationIDFilterStr(string{/* values here */})
```

### 

```go
paymentsListQueryParamOrganizationIDFilter := operations.CreatePaymentsListQueryParamOrganizationIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch paymentsListQueryParamOrganizationIDFilter.Type {
	case operations.PaymentsListQueryParamOrganizationIDFilterTypeStr:
		// paymentsListQueryParamOrganizationIDFilter.Str is populated
	case operations.PaymentsListQueryParamOrganizationIDFilterTypeArrayOfStr:
		// paymentsListQueryParamOrganizationIDFilter.ArrayOfStr is populated
}
```
