# FileRead


## Supported Types

### DownloadableFileRead

```go
fileRead := components.CreateFileReadDownloadable(components.DownloadableFileRead{/* values here */})
```

### OrganizationAvatarFileRead

```go
fileRead := components.CreateFileReadOrganizationAvatar(components.OrganizationAvatarFileRead{/* values here */})
```

### ProductMediaFileRead

```go
fileRead := components.CreateFileReadProductMedia(components.ProductMediaFileRead{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch fileRead.Type {
	case components.FileReadTypeDownloadable:
		// fileRead.DownloadableFileRead is populated
	case components.FileReadTypeOrganizationAvatar:
		// fileRead.OrganizationAvatarFileRead is populated
	case components.FileReadTypeProductMedia:
		// fileRead.ProductMediaFileRead is populated
}
```
