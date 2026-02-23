# FileCreate


## Supported Types

### DownloadableFileCreate

```go
fileCreate := components.CreateFileCreateDownloadable(components.DownloadableFileCreate{/* values here */})
```

### OrganizationAvatarFileCreate

```go
fileCreate := components.CreateFileCreateOrganizationAvatar(components.OrganizationAvatarFileCreate{/* values here */})
```

### ProductMediaFileCreate

```go
fileCreate := components.CreateFileCreateProductMedia(components.ProductMediaFileCreate{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch fileCreate.Type {
	case components.FileCreateTypeDownloadable:
		// fileCreate.DownloadableFileCreate is populated
	case components.FileCreateTypeOrganizationAvatar:
		// fileCreate.OrganizationAvatarFileCreate is populated
	case components.FileCreateTypeProductMedia:
		// fileCreate.ProductMediaFileCreate is populated
}
```
