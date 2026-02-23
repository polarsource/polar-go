# FilesUpdateResponseFilesUpdate

File updated.


## Supported Types

### DownloadableFileRead

```go
filesUpdateResponseFilesUpdate := operations.CreateFilesUpdateResponseFilesUpdateDownloadable(components.DownloadableFileRead{/* values here */})
```

### ProductMediaFileRead

```go
filesUpdateResponseFilesUpdate := operations.CreateFilesUpdateResponseFilesUpdateProductMedia(components.ProductMediaFileRead{/* values here */})
```

### OrganizationAvatarFileRead

```go
filesUpdateResponseFilesUpdate := operations.CreateFilesUpdateResponseFilesUpdateOrganizationAvatar(components.OrganizationAvatarFileRead{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch filesUpdateResponseFilesUpdate.Type {
	case operations.FilesUpdateResponseFilesUpdateTypeDownloadable:
		// filesUpdateResponseFilesUpdate.DownloadableFileRead is populated
	case operations.FilesUpdateResponseFilesUpdateTypeProductMedia:
		// filesUpdateResponseFilesUpdate.ProductMediaFileRead is populated
	case operations.FilesUpdateResponseFilesUpdateTypeOrganizationAvatar:
		// filesUpdateResponseFilesUpdate.OrganizationAvatarFileRead is populated
}
```
