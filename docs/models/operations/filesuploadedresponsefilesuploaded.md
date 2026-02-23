# FilesUploadedResponseFilesUploaded

File upload completed.


## Supported Types

### DownloadableFileRead

```go
filesUploadedResponseFilesUploaded := operations.CreateFilesUploadedResponseFilesUploadedDownloadable(components.DownloadableFileRead{/* values here */})
```

### ProductMediaFileRead

```go
filesUploadedResponseFilesUploaded := operations.CreateFilesUploadedResponseFilesUploadedProductMedia(components.ProductMediaFileRead{/* values here */})
```

### OrganizationAvatarFileRead

```go
filesUploadedResponseFilesUploaded := operations.CreateFilesUploadedResponseFilesUploadedOrganizationAvatar(components.OrganizationAvatarFileRead{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch filesUploadedResponseFilesUploaded.Type {
	case operations.FilesUploadedResponseFilesUploadedTypeDownloadable:
		// filesUploadedResponseFilesUploaded.DownloadableFileRead is populated
	case operations.FilesUploadedResponseFilesUploadedTypeProductMedia:
		// filesUploadedResponseFilesUploaded.ProductMediaFileRead is populated
	case operations.FilesUploadedResponseFilesUploadedTypeOrganizationAvatar:
		// filesUploadedResponseFilesUploaded.OrganizationAvatarFileRead is populated
}
```
