# Clauses


## Supported Types

### FilterClause

```go
clauses := components.CreateClausesFilterClause(components.FilterClause{/* values here */})
```

### Filter

```go
clauses := components.CreateClausesFilter(components.Filter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch clauses.Type {
	case components.ClausesTypeFilterClause:
		// clauses.FilterClause is populated
	case components.ClausesTypeFilter:
		// clauses.Filter is populated
}
```
