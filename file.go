// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package polar

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/polarsource/polar-go/internal/apijson"
	"github.com/polarsource/polar-go/internal/apiquery"
	"github.com/polarsource/polar-go/internal/pagination"
	"github.com/polarsource/polar-go/internal/param"
	"github.com/polarsource/polar-go/internal/requestconfig"
	"github.com/polarsource/polar-go/option"
	"github.com/tidwall/gjson"
)

// FileService contains methods and other services that help with interacting with
// the polar API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFileService] method instead.
type FileService struct {
	Options []option.RequestOption
}

// NewFileService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFileService(opts ...option.RequestOption) (r *FileService) {
	r = &FileService{}
	r.Options = opts
	return
}

// Create a file.
func (r *FileService) New(ctx context.Context, body FileNewParams, opts ...option.RequestOption) (res *FileUpload, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/files/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update a file.
func (r *FileService) Update(ctx context.Context, id string, body FileUpdateParams, opts ...option.RequestOption) (res *FileUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/files/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List files.
func (r *FileService) List(ctx context.Context, query FileListParams, opts ...option.RequestOption) (res *pagination.PolarPagination[FileListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/files/"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List files.
func (r *FileService) ListAutoPaging(ctx context.Context, query FileListParams, opts ...option.RequestOption) *pagination.PolarPaginationAutoPager[FileListResponse] {
	return pagination.NewPolarPaginationAutoPager(r.List(ctx, query, opts...))
}

// Delete a file.
func (r *FileService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/files/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Complete a file upload.
func (r *FileService) Uploaded(ctx context.Context, params FileUploadedParams, opts ...option.RequestOption) (res *FileUploadedResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.PathID.Value == "" {
		err = errors.New("missing required path_id parameter")
		return
	}
	path := fmt.Sprintf("v1/files/%s/uploaded", params.PathID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// File to be associated with the downloadables benefit.
type DownloadableFileRead struct {
	ID                   string                      `json:"id,required" format:"uuid4"`
	ChecksumEtag         string                      `json:"checksum_etag,required,nullable"`
	ChecksumSha256Base64 string                      `json:"checksum_sha256_base64,required,nullable"`
	ChecksumSha256Hex    string                      `json:"checksum_sha256_hex,required,nullable"`
	CreatedAt            time.Time                   `json:"created_at,required" format:"date-time"`
	IsUploaded           bool                        `json:"is_uploaded,required"`
	LastModifiedAt       time.Time                   `json:"last_modified_at,required,nullable" format:"date-time"`
	MimeType             string                      `json:"mime_type,required"`
	Name                 string                      `json:"name,required"`
	OrganizationID       string                      `json:"organization_id,required" format:"uuid4"`
	Path                 string                      `json:"path,required"`
	Service              DownloadableFileReadService `json:"service,required"`
	Size                 int64                       `json:"size,required"`
	SizeReadable         string                      `json:"size_readable,required"`
	StorageVersion       string                      `json:"storage_version,required,nullable"`
	Version              string                      `json:"version,required,nullable"`
	JSON                 downloadableFileReadJSON    `json:"-"`
}

// downloadableFileReadJSON contains the JSON metadata for the struct
// [DownloadableFileRead]
type downloadableFileReadJSON struct {
	ID                   apijson.Field
	ChecksumEtag         apijson.Field
	ChecksumSha256Base64 apijson.Field
	ChecksumSha256Hex    apijson.Field
	CreatedAt            apijson.Field
	IsUploaded           apijson.Field
	LastModifiedAt       apijson.Field
	MimeType             apijson.Field
	Name                 apijson.Field
	OrganizationID       apijson.Field
	Path                 apijson.Field
	Service              apijson.Field
	Size                 apijson.Field
	SizeReadable         apijson.Field
	StorageVersion       apijson.Field
	Version              apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *DownloadableFileRead) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r downloadableFileReadJSON) RawJSON() string {
	return r.raw
}

func (r DownloadableFileRead) implementsListResourceAnnotatedUnionItem() {}

func (r DownloadableFileRead) implementsFileUpdateResponse() {}

func (r DownloadableFileRead) implementsFileListResponse() {}

func (r DownloadableFileRead) implementsFileUploadedResponse() {}

type DownloadableFileReadService string

const (
	DownloadableFileReadServiceDownloadable DownloadableFileReadService = "downloadable"
)

func (r DownloadableFileReadService) IsKnown() bool {
	switch r {
	case DownloadableFileReadServiceDownloadable:
		return true
	}
	return false
}

type FileUpload struct {
	ID                   string            `json:"id,required" format:"uuid4"`
	ChecksumEtag         string            `json:"checksum_etag,required,nullable"`
	ChecksumSha256Base64 string            `json:"checksum_sha256_base64,required,nullable"`
	ChecksumSha256Hex    string            `json:"checksum_sha256_hex,required,nullable"`
	LastModifiedAt       time.Time         `json:"last_modified_at,required,nullable" format:"date-time"`
	MimeType             string            `json:"mime_type,required"`
	Name                 string            `json:"name,required"`
	OrganizationID       string            `json:"organization_id,required" format:"uuid4"`
	Path                 string            `json:"path,required"`
	Service              FileUploadService `json:"service,required"`
	Size                 int64             `json:"size,required"`
	SizeReadable         string            `json:"size_readable,required"`
	StorageVersion       string            `json:"storage_version,required,nullable"`
	Upload               FileUploadUpload  `json:"upload,required"`
	Version              string            `json:"version,required,nullable"`
	IsUploaded           bool              `json:"is_uploaded"`
	JSON                 fileUploadJSON    `json:"-"`
}

// fileUploadJSON contains the JSON metadata for the struct [FileUpload]
type fileUploadJSON struct {
	ID                   apijson.Field
	ChecksumEtag         apijson.Field
	ChecksumSha256Base64 apijson.Field
	ChecksumSha256Hex    apijson.Field
	LastModifiedAt       apijson.Field
	MimeType             apijson.Field
	Name                 apijson.Field
	OrganizationID       apijson.Field
	Path                 apijson.Field
	Service              apijson.Field
	Size                 apijson.Field
	SizeReadable         apijson.Field
	StorageVersion       apijson.Field
	Upload               apijson.Field
	Version              apijson.Field
	IsUploaded           apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *FileUpload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fileUploadJSON) RawJSON() string {
	return r.raw
}

type FileUploadService string

const (
	FileUploadServiceDownloadable       FileUploadService = "downloadable"
	FileUploadServiceProductMedia       FileUploadService = "product_media"
	FileUploadServiceOrganizationAvatar FileUploadService = "organization_avatar"
)

func (r FileUploadService) IsKnown() bool {
	switch r {
	case FileUploadServiceDownloadable, FileUploadServiceProductMedia, FileUploadServiceOrganizationAvatar:
		return true
	}
	return false
}

type FileUploadUpload struct {
	ID    string                 `json:"id,required"`
	Parts []FileUploadUploadPart `json:"parts,required"`
	Path  string                 `json:"path,required"`
	JSON  fileUploadUploadJSON   `json:"-"`
}

// fileUploadUploadJSON contains the JSON metadata for the struct
// [FileUploadUpload]
type fileUploadUploadJSON struct {
	ID          apijson.Field
	Parts       apijson.Field
	Path        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FileUploadUpload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fileUploadUploadJSON) RawJSON() string {
	return r.raw
}

type FileUploadUploadPart struct {
	ChunkEnd             int64                    `json:"chunk_end,required"`
	ChunkStart           int64                    `json:"chunk_start,required"`
	ExpiresAt            time.Time                `json:"expires_at,required" format:"date-time"`
	Number               int64                    `json:"number,required"`
	URL                  string                   `json:"url,required"`
	ChecksumSha256Base64 string                   `json:"checksum_sha256_base64,nullable"`
	Headers              map[string]string        `json:"headers"`
	JSON                 fileUploadUploadPartJSON `json:"-"`
}

// fileUploadUploadPartJSON contains the JSON metadata for the struct
// [FileUploadUploadPart]
type fileUploadUploadPartJSON struct {
	ChunkEnd             apijson.Field
	ChunkStart           apijson.Field
	ExpiresAt            apijson.Field
	Number               apijson.Field
	URL                  apijson.Field
	ChecksumSha256Base64 apijson.Field
	Headers              apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *FileUploadUploadPart) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fileUploadUploadPartJSON) RawJSON() string {
	return r.raw
}

type ListResourceAnnotatedUnion struct {
	Pagination ListResourceAnnotatedUnionPagination `json:"pagination,required"`
	Items      []ListResourceAnnotatedUnionItem     `json:"items"`
	JSON       listResourceAnnotatedUnionJSON       `json:"-"`
}

// listResourceAnnotatedUnionJSON contains the JSON metadata for the struct
// [ListResourceAnnotatedUnion]
type listResourceAnnotatedUnionJSON struct {
	Pagination  apijson.Field
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListResourceAnnotatedUnion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResourceAnnotatedUnionJSON) RawJSON() string {
	return r.raw
}

type ListResourceAnnotatedUnionPagination struct {
	MaxPage    int64                                    `json:"max_page,required"`
	TotalCount int64                                    `json:"total_count,required"`
	JSON       listResourceAnnotatedUnionPaginationJSON `json:"-"`
}

// listResourceAnnotatedUnionPaginationJSON contains the JSON metadata for the
// struct [ListResourceAnnotatedUnionPagination]
type listResourceAnnotatedUnionPaginationJSON struct {
	MaxPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListResourceAnnotatedUnionPagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResourceAnnotatedUnionPaginationJSON) RawJSON() string {
	return r.raw
}

// File to be associated with the downloadables benefit.
type ListResourceAnnotatedUnionItem struct {
	ID                   string                                 `json:"id,required" format:"uuid4"`
	OrganizationID       string                                 `json:"organization_id,required" format:"uuid4"`
	Name                 string                                 `json:"name,required"`
	Path                 string                                 `json:"path,required"`
	MimeType             string                                 `json:"mime_type,required"`
	Size                 int64                                  `json:"size,required"`
	StorageVersion       string                                 `json:"storage_version,required,nullable"`
	ChecksumEtag         string                                 `json:"checksum_etag,required,nullable"`
	ChecksumSha256Base64 string                                 `json:"checksum_sha256_base64,required,nullable"`
	ChecksumSha256Hex    string                                 `json:"checksum_sha256_hex,required,nullable"`
	LastModifiedAt       time.Time                              `json:"last_modified_at,required,nullable" format:"date-time"`
	Version              string                                 `json:"version,required,nullable"`
	Service              ListResourceAnnotatedUnionItemsService `json:"service,required"`
	IsUploaded           bool                                   `json:"is_uploaded,required"`
	CreatedAt            time.Time                              `json:"created_at,required" format:"date-time"`
	SizeReadable         string                                 `json:"size_readable,required"`
	PublicURL            string                                 `json:"public_url"`
	JSON                 listResourceAnnotatedUnionItemJSON     `json:"-"`
	union                ListResourceAnnotatedUnionItemsUnion
}

// listResourceAnnotatedUnionItemJSON contains the JSON metadata for the struct
// [ListResourceAnnotatedUnionItem]
type listResourceAnnotatedUnionItemJSON struct {
	ID                   apijson.Field
	OrganizationID       apijson.Field
	Name                 apijson.Field
	Path                 apijson.Field
	MimeType             apijson.Field
	Size                 apijson.Field
	StorageVersion       apijson.Field
	ChecksumEtag         apijson.Field
	ChecksumSha256Base64 apijson.Field
	ChecksumSha256Hex    apijson.Field
	LastModifiedAt       apijson.Field
	Version              apijson.Field
	Service              apijson.Field
	IsUploaded           apijson.Field
	CreatedAt            apijson.Field
	SizeReadable         apijson.Field
	PublicURL            apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r listResourceAnnotatedUnionItemJSON) RawJSON() string {
	return r.raw
}

func (r *ListResourceAnnotatedUnionItem) UnmarshalJSON(data []byte) (err error) {
	*r = ListResourceAnnotatedUnionItem{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ListResourceAnnotatedUnionItemsUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [DownloadableFileRead],
// [ProductMediaFileReadOutput], [OrganizationAvatarFileRead].
func (r ListResourceAnnotatedUnionItem) AsUnion() ListResourceAnnotatedUnionItemsUnion {
	return r.union
}

// File to be associated with the downloadables benefit.
//
// Union satisfied by [DownloadableFileRead], [ProductMediaFileReadOutput] or
// [OrganizationAvatarFileRead].
type ListResourceAnnotatedUnionItemsUnion interface {
	implementsListResourceAnnotatedUnionItem()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ListResourceAnnotatedUnionItemsUnion)(nil)).Elem(),
		"service",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DownloadableFileRead{}),
			DiscriminatorValue: "downloadable",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProductMediaFileReadOutput{}),
			DiscriminatorValue: "product_media",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(OrganizationAvatarFileRead{}),
			DiscriminatorValue: "organization_avatar",
		},
	)
}

type ListResourceAnnotatedUnionItemsService string

const (
	ListResourceAnnotatedUnionItemsServiceDownloadable       ListResourceAnnotatedUnionItemsService = "downloadable"
	ListResourceAnnotatedUnionItemsServiceProductMedia       ListResourceAnnotatedUnionItemsService = "product_media"
	ListResourceAnnotatedUnionItemsServiceOrganizationAvatar ListResourceAnnotatedUnionItemsService = "organization_avatar"
)

func (r ListResourceAnnotatedUnionItemsService) IsKnown() bool {
	switch r {
	case ListResourceAnnotatedUnionItemsServiceDownloadable, ListResourceAnnotatedUnionItemsServiceProductMedia, ListResourceAnnotatedUnionItemsServiceOrganizationAvatar:
		return true
	}
	return false
}

// File to be used as an organization avatar.
type OrganizationAvatarFileRead struct {
	ID                   string                            `json:"id,required" format:"uuid4"`
	ChecksumEtag         string                            `json:"checksum_etag,required,nullable"`
	ChecksumSha256Base64 string                            `json:"checksum_sha256_base64,required,nullable"`
	ChecksumSha256Hex    string                            `json:"checksum_sha256_hex,required,nullable"`
	CreatedAt            time.Time                         `json:"created_at,required" format:"date-time"`
	IsUploaded           bool                              `json:"is_uploaded,required"`
	LastModifiedAt       time.Time                         `json:"last_modified_at,required,nullable" format:"date-time"`
	MimeType             string                            `json:"mime_type,required"`
	Name                 string                            `json:"name,required"`
	OrganizationID       string                            `json:"organization_id,required" format:"uuid4"`
	Path                 string                            `json:"path,required"`
	PublicURL            string                            `json:"public_url,required"`
	Service              OrganizationAvatarFileReadService `json:"service,required"`
	Size                 int64                             `json:"size,required"`
	SizeReadable         string                            `json:"size_readable,required"`
	StorageVersion       string                            `json:"storage_version,required,nullable"`
	Version              string                            `json:"version,required,nullable"`
	JSON                 organizationAvatarFileReadJSON    `json:"-"`
}

// organizationAvatarFileReadJSON contains the JSON metadata for the struct
// [OrganizationAvatarFileRead]
type organizationAvatarFileReadJSON struct {
	ID                   apijson.Field
	ChecksumEtag         apijson.Field
	ChecksumSha256Base64 apijson.Field
	ChecksumSha256Hex    apijson.Field
	CreatedAt            apijson.Field
	IsUploaded           apijson.Field
	LastModifiedAt       apijson.Field
	MimeType             apijson.Field
	Name                 apijson.Field
	OrganizationID       apijson.Field
	Path                 apijson.Field
	PublicURL            apijson.Field
	Service              apijson.Field
	Size                 apijson.Field
	SizeReadable         apijson.Field
	StorageVersion       apijson.Field
	Version              apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *OrganizationAvatarFileRead) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationAvatarFileReadJSON) RawJSON() string {
	return r.raw
}

func (r OrganizationAvatarFileRead) implementsListResourceAnnotatedUnionItem() {}

func (r OrganizationAvatarFileRead) implementsFileUpdateResponse() {}

func (r OrganizationAvatarFileRead) implementsFileListResponse() {}

func (r OrganizationAvatarFileRead) implementsFileUploadedResponse() {}

type OrganizationAvatarFileReadService string

const (
	OrganizationAvatarFileReadServiceOrganizationAvatar OrganizationAvatarFileReadService = "organization_avatar"
)

func (r OrganizationAvatarFileReadService) IsKnown() bool {
	switch r {
	case OrganizationAvatarFileReadServiceOrganizationAvatar:
		return true
	}
	return false
}

// File to be used as a product media file.
type ProductMediaFileReadOutput struct {
	ID                   string                            `json:"id,required" format:"uuid4"`
	ChecksumEtag         string                            `json:"checksum_etag,required,nullable"`
	ChecksumSha256Base64 string                            `json:"checksum_sha256_base64,required,nullable"`
	ChecksumSha256Hex    string                            `json:"checksum_sha256_hex,required,nullable"`
	CreatedAt            time.Time                         `json:"created_at,required" format:"date-time"`
	IsUploaded           bool                              `json:"is_uploaded,required"`
	LastModifiedAt       time.Time                         `json:"last_modified_at,required,nullable" format:"date-time"`
	MimeType             string                            `json:"mime_type,required"`
	Name                 string                            `json:"name,required"`
	OrganizationID       string                            `json:"organization_id,required" format:"uuid4"`
	Path                 string                            `json:"path,required"`
	PublicURL            string                            `json:"public_url,required"`
	Service              ProductMediaFileReadOutputService `json:"service,required"`
	Size                 int64                             `json:"size,required"`
	SizeReadable         string                            `json:"size_readable,required"`
	StorageVersion       string                            `json:"storage_version,required,nullable"`
	Version              string                            `json:"version,required,nullable"`
	JSON                 productMediaFileReadOutputJSON    `json:"-"`
}

// productMediaFileReadOutputJSON contains the JSON metadata for the struct
// [ProductMediaFileReadOutput]
type productMediaFileReadOutputJSON struct {
	ID                   apijson.Field
	ChecksumEtag         apijson.Field
	ChecksumSha256Base64 apijson.Field
	ChecksumSha256Hex    apijson.Field
	CreatedAt            apijson.Field
	IsUploaded           apijson.Field
	LastModifiedAt       apijson.Field
	MimeType             apijson.Field
	Name                 apijson.Field
	OrganizationID       apijson.Field
	Path                 apijson.Field
	PublicURL            apijson.Field
	Service              apijson.Field
	Size                 apijson.Field
	SizeReadable         apijson.Field
	StorageVersion       apijson.Field
	Version              apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ProductMediaFileReadOutput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productMediaFileReadOutputJSON) RawJSON() string {
	return r.raw
}

func (r ProductMediaFileReadOutput) implementsListResourceAnnotatedUnionItem() {}

func (r ProductMediaFileReadOutput) implementsFileUpdateResponse() {}

func (r ProductMediaFileReadOutput) implementsFileListResponse() {}

func (r ProductMediaFileReadOutput) implementsFileUploadedResponse() {}

type ProductMediaFileReadOutputService string

const (
	ProductMediaFileReadOutputServiceProductMedia ProductMediaFileReadOutputService = "product_media"
)

func (r ProductMediaFileReadOutputService) IsKnown() bool {
	switch r {
	case ProductMediaFileReadOutputServiceProductMedia:
		return true
	}
	return false
}

// File to be associated with the downloadables benefit.
type FileUpdateResponse struct {
	ID                   string                    `json:"id,required" format:"uuid4"`
	OrganizationID       string                    `json:"organization_id,required" format:"uuid4"`
	Name                 string                    `json:"name,required"`
	Path                 string                    `json:"path,required"`
	MimeType             string                    `json:"mime_type,required"`
	Size                 int64                     `json:"size,required"`
	StorageVersion       string                    `json:"storage_version,required,nullable"`
	ChecksumEtag         string                    `json:"checksum_etag,required,nullable"`
	ChecksumSha256Base64 string                    `json:"checksum_sha256_base64,required,nullable"`
	ChecksumSha256Hex    string                    `json:"checksum_sha256_hex,required,nullable"`
	LastModifiedAt       time.Time                 `json:"last_modified_at,required,nullable" format:"date-time"`
	Version              string                    `json:"version,required,nullable"`
	Service              FileUpdateResponseService `json:"service,required"`
	IsUploaded           bool                      `json:"is_uploaded,required"`
	CreatedAt            time.Time                 `json:"created_at,required" format:"date-time"`
	SizeReadable         string                    `json:"size_readable,required"`
	PublicURL            string                    `json:"public_url"`
	JSON                 fileUpdateResponseJSON    `json:"-"`
	union                FileUpdateResponseUnion
}

// fileUpdateResponseJSON contains the JSON metadata for the struct
// [FileUpdateResponse]
type fileUpdateResponseJSON struct {
	ID                   apijson.Field
	OrganizationID       apijson.Field
	Name                 apijson.Field
	Path                 apijson.Field
	MimeType             apijson.Field
	Size                 apijson.Field
	StorageVersion       apijson.Field
	ChecksumEtag         apijson.Field
	ChecksumSha256Base64 apijson.Field
	ChecksumSha256Hex    apijson.Field
	LastModifiedAt       apijson.Field
	Version              apijson.Field
	Service              apijson.Field
	IsUploaded           apijson.Field
	CreatedAt            apijson.Field
	SizeReadable         apijson.Field
	PublicURL            apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r fileUpdateResponseJSON) RawJSON() string {
	return r.raw
}

func (r *FileUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	*r = FileUpdateResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [FileUpdateResponseUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [DownloadableFileRead],
// [ProductMediaFileReadOutput], [OrganizationAvatarFileRead].
func (r FileUpdateResponse) AsUnion() FileUpdateResponseUnion {
	return r.union
}

// File to be associated with the downloadables benefit.
//
// Union satisfied by [DownloadableFileRead], [ProductMediaFileReadOutput] or
// [OrganizationAvatarFileRead].
type FileUpdateResponseUnion interface {
	implementsFileUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FileUpdateResponseUnion)(nil)).Elem(),
		"service",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DownloadableFileRead{}),
			DiscriminatorValue: "downloadable",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProductMediaFileReadOutput{}),
			DiscriminatorValue: "product_media",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(OrganizationAvatarFileRead{}),
			DiscriminatorValue: "organization_avatar",
		},
	)
}

type FileUpdateResponseService string

const (
	FileUpdateResponseServiceDownloadable       FileUpdateResponseService = "downloadable"
	FileUpdateResponseServiceProductMedia       FileUpdateResponseService = "product_media"
	FileUpdateResponseServiceOrganizationAvatar FileUpdateResponseService = "organization_avatar"
)

func (r FileUpdateResponseService) IsKnown() bool {
	switch r {
	case FileUpdateResponseServiceDownloadable, FileUpdateResponseServiceProductMedia, FileUpdateResponseServiceOrganizationAvatar:
		return true
	}
	return false
}

// File to be associated with the downloadables benefit.
type FileListResponse struct {
	ID                   string                  `json:"id,required" format:"uuid4"`
	OrganizationID       string                  `json:"organization_id,required" format:"uuid4"`
	Name                 string                  `json:"name,required"`
	Path                 string                  `json:"path,required"`
	MimeType             string                  `json:"mime_type,required"`
	Size                 int64                   `json:"size,required"`
	StorageVersion       string                  `json:"storage_version,required,nullable"`
	ChecksumEtag         string                  `json:"checksum_etag,required,nullable"`
	ChecksumSha256Base64 string                  `json:"checksum_sha256_base64,required,nullable"`
	ChecksumSha256Hex    string                  `json:"checksum_sha256_hex,required,nullable"`
	LastModifiedAt       time.Time               `json:"last_modified_at,required,nullable" format:"date-time"`
	Version              string                  `json:"version,required,nullable"`
	Service              FileListResponseService `json:"service,required"`
	IsUploaded           bool                    `json:"is_uploaded,required"`
	CreatedAt            time.Time               `json:"created_at,required" format:"date-time"`
	SizeReadable         string                  `json:"size_readable,required"`
	PublicURL            string                  `json:"public_url"`
	JSON                 fileListResponseJSON    `json:"-"`
	union                FileListResponseUnion
}

// fileListResponseJSON contains the JSON metadata for the struct
// [FileListResponse]
type fileListResponseJSON struct {
	ID                   apijson.Field
	OrganizationID       apijson.Field
	Name                 apijson.Field
	Path                 apijson.Field
	MimeType             apijson.Field
	Size                 apijson.Field
	StorageVersion       apijson.Field
	ChecksumEtag         apijson.Field
	ChecksumSha256Base64 apijson.Field
	ChecksumSha256Hex    apijson.Field
	LastModifiedAt       apijson.Field
	Version              apijson.Field
	Service              apijson.Field
	IsUploaded           apijson.Field
	CreatedAt            apijson.Field
	SizeReadable         apijson.Field
	PublicURL            apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r fileListResponseJSON) RawJSON() string {
	return r.raw
}

func (r *FileListResponse) UnmarshalJSON(data []byte) (err error) {
	*r = FileListResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [FileListResponseUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [DownloadableFileRead],
// [ProductMediaFileReadOutput], [OrganizationAvatarFileRead].
func (r FileListResponse) AsUnion() FileListResponseUnion {
	return r.union
}

// File to be associated with the downloadables benefit.
//
// Union satisfied by [DownloadableFileRead], [ProductMediaFileReadOutput] or
// [OrganizationAvatarFileRead].
type FileListResponseUnion interface {
	implementsFileListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FileListResponseUnion)(nil)).Elem(),
		"service",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DownloadableFileRead{}),
			DiscriminatorValue: "downloadable",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProductMediaFileReadOutput{}),
			DiscriminatorValue: "product_media",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(OrganizationAvatarFileRead{}),
			DiscriminatorValue: "organization_avatar",
		},
	)
}

type FileListResponseService string

const (
	FileListResponseServiceDownloadable       FileListResponseService = "downloadable"
	FileListResponseServiceProductMedia       FileListResponseService = "product_media"
	FileListResponseServiceOrganizationAvatar FileListResponseService = "organization_avatar"
)

func (r FileListResponseService) IsKnown() bool {
	switch r {
	case FileListResponseServiceDownloadable, FileListResponseServiceProductMedia, FileListResponseServiceOrganizationAvatar:
		return true
	}
	return false
}

// File to be associated with the downloadables benefit.
type FileUploadedResponse struct {
	ID                   string                      `json:"id,required" format:"uuid4"`
	OrganizationID       string                      `json:"organization_id,required" format:"uuid4"`
	Name                 string                      `json:"name,required"`
	Path                 string                      `json:"path,required"`
	MimeType             string                      `json:"mime_type,required"`
	Size                 int64                       `json:"size,required"`
	StorageVersion       string                      `json:"storage_version,required,nullable"`
	ChecksumEtag         string                      `json:"checksum_etag,required,nullable"`
	ChecksumSha256Base64 string                      `json:"checksum_sha256_base64,required,nullable"`
	ChecksumSha256Hex    string                      `json:"checksum_sha256_hex,required,nullable"`
	LastModifiedAt       time.Time                   `json:"last_modified_at,required,nullable" format:"date-time"`
	Version              string                      `json:"version,required,nullable"`
	Service              FileUploadedResponseService `json:"service,required"`
	IsUploaded           bool                        `json:"is_uploaded,required"`
	CreatedAt            time.Time                   `json:"created_at,required" format:"date-time"`
	SizeReadable         string                      `json:"size_readable,required"`
	PublicURL            string                      `json:"public_url"`
	JSON                 fileUploadedResponseJSON    `json:"-"`
	union                FileUploadedResponseUnion
}

// fileUploadedResponseJSON contains the JSON metadata for the struct
// [FileUploadedResponse]
type fileUploadedResponseJSON struct {
	ID                   apijson.Field
	OrganizationID       apijson.Field
	Name                 apijson.Field
	Path                 apijson.Field
	MimeType             apijson.Field
	Size                 apijson.Field
	StorageVersion       apijson.Field
	ChecksumEtag         apijson.Field
	ChecksumSha256Base64 apijson.Field
	ChecksumSha256Hex    apijson.Field
	LastModifiedAt       apijson.Field
	Version              apijson.Field
	Service              apijson.Field
	IsUploaded           apijson.Field
	CreatedAt            apijson.Field
	SizeReadable         apijson.Field
	PublicURL            apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r fileUploadedResponseJSON) RawJSON() string {
	return r.raw
}

func (r *FileUploadedResponse) UnmarshalJSON(data []byte) (err error) {
	*r = FileUploadedResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [FileUploadedResponseUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [DownloadableFileRead],
// [ProductMediaFileReadOutput], [OrganizationAvatarFileRead].
func (r FileUploadedResponse) AsUnion() FileUploadedResponseUnion {
	return r.union
}

// File to be associated with the downloadables benefit.
//
// Union satisfied by [DownloadableFileRead], [ProductMediaFileReadOutput] or
// [OrganizationAvatarFileRead].
type FileUploadedResponseUnion interface {
	implementsFileUploadedResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FileUploadedResponseUnion)(nil)).Elem(),
		"service",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DownloadableFileRead{}),
			DiscriminatorValue: "downloadable",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProductMediaFileReadOutput{}),
			DiscriminatorValue: "product_media",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(OrganizationAvatarFileRead{}),
			DiscriminatorValue: "organization_avatar",
		},
	)
}

type FileUploadedResponseService string

const (
	FileUploadedResponseServiceDownloadable       FileUploadedResponseService = "downloadable"
	FileUploadedResponseServiceProductMedia       FileUploadedResponseService = "product_media"
	FileUploadedResponseServiceOrganizationAvatar FileUploadedResponseService = "organization_avatar"
)

func (r FileUploadedResponseService) IsKnown() bool {
	switch r {
	case FileUploadedResponseServiceDownloadable, FileUploadedResponseServiceProductMedia, FileUploadedResponseServiceOrganizationAvatar:
		return true
	}
	return false
}

type FileNewParams struct {
	// Schema to create a file to be associated with the downloadables benefit.
	Body FileNewParamsBodyUnion `json:"body,required"`
}

func (r FileNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// Schema to create a file to be associated with the downloadables benefit.
type FileNewParamsBody struct {
	// The organization ID.
	OrganizationID       param.Field[string]                   `json:"organization_id" format:"uuid4"`
	Name                 param.Field[string]                   `json:"name,required"`
	MimeType             param.Field[string]                   `json:"mime_type,required"`
	Size                 param.Field[int64]                    `json:"size,required"`
	ChecksumSha256Base64 param.Field[string]                   `json:"checksum_sha256_base64"`
	Upload               param.Field[interface{}]              `json:"upload"`
	Service              param.Field[FileNewParamsBodyService] `json:"service,required"`
	Version              param.Field[string]                   `json:"version"`
}

func (r FileNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FileNewParamsBody) implementsFileNewParamsBodyUnion() {}

// Schema to create a file to be associated with the downloadables benefit.
//
// Satisfied by [FileNewParamsBodyDownloadableFileCreate],
// [FileNewParamsBodyProductMediaFileCreate],
// [FileNewParamsBodyOrganizationAvatarFileCreate], [FileNewParamsBody].
type FileNewParamsBodyUnion interface {
	implementsFileNewParamsBodyUnion()
}

// Schema to create a file to be associated with the downloadables benefit.
type FileNewParamsBodyDownloadableFileCreate struct {
	MimeType             param.Field[string]                                         `json:"mime_type,required"`
	Name                 param.Field[string]                                         `json:"name,required"`
	Service              param.Field[FileNewParamsBodyDownloadableFileCreateService] `json:"service,required"`
	Size                 param.Field[int64]                                          `json:"size,required"`
	Upload               param.Field[FileNewParamsBodyDownloadableFileCreateUpload]  `json:"upload,required"`
	ChecksumSha256Base64 param.Field[string]                                         `json:"checksum_sha256_base64"`
	// The organization ID.
	OrganizationID param.Field[string] `json:"organization_id" format:"uuid4"`
	Version        param.Field[string] `json:"version"`
}

func (r FileNewParamsBodyDownloadableFileCreate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FileNewParamsBodyDownloadableFileCreate) implementsFileNewParamsBodyUnion() {}

type FileNewParamsBodyDownloadableFileCreateService string

const (
	FileNewParamsBodyDownloadableFileCreateServiceDownloadable FileNewParamsBodyDownloadableFileCreateService = "downloadable"
)

func (r FileNewParamsBodyDownloadableFileCreateService) IsKnown() bool {
	switch r {
	case FileNewParamsBodyDownloadableFileCreateServiceDownloadable:
		return true
	}
	return false
}

type FileNewParamsBodyDownloadableFileCreateUpload struct {
	Parts param.Field[[]FileNewParamsBodyDownloadableFileCreateUploadPart] `json:"parts,required"`
}

func (r FileNewParamsBodyDownloadableFileCreateUpload) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FileNewParamsBodyDownloadableFileCreateUploadPart struct {
	ChunkEnd             param.Field[int64]  `json:"chunk_end,required"`
	ChunkStart           param.Field[int64]  `json:"chunk_start,required"`
	Number               param.Field[int64]  `json:"number,required"`
	ChecksumSha256Base64 param.Field[string] `json:"checksum_sha256_base64"`
}

func (r FileNewParamsBodyDownloadableFileCreateUploadPart) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Schema to create a file to be used as a product media file.
type FileNewParamsBodyProductMediaFileCreate struct {
	// MIME type of the file. Only images are supported for this type of file.
	MimeType param.Field[string]                                         `json:"mime_type,required"`
	Name     param.Field[string]                                         `json:"name,required"`
	Service  param.Field[FileNewParamsBodyProductMediaFileCreateService] `json:"service,required"`
	// Size of the file. A maximum of 10 MB is allowed for this type of file.
	Size                 param.Field[int64]                                         `json:"size,required"`
	Upload               param.Field[FileNewParamsBodyProductMediaFileCreateUpload] `json:"upload,required"`
	ChecksumSha256Base64 param.Field[string]                                        `json:"checksum_sha256_base64"`
	// The organization ID.
	OrganizationID param.Field[string] `json:"organization_id" format:"uuid4"`
	Version        param.Field[string] `json:"version"`
}

func (r FileNewParamsBodyProductMediaFileCreate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FileNewParamsBodyProductMediaFileCreate) implementsFileNewParamsBodyUnion() {}

type FileNewParamsBodyProductMediaFileCreateService string

const (
	FileNewParamsBodyProductMediaFileCreateServiceProductMedia FileNewParamsBodyProductMediaFileCreateService = "product_media"
)

func (r FileNewParamsBodyProductMediaFileCreateService) IsKnown() bool {
	switch r {
	case FileNewParamsBodyProductMediaFileCreateServiceProductMedia:
		return true
	}
	return false
}

type FileNewParamsBodyProductMediaFileCreateUpload struct {
	Parts param.Field[[]FileNewParamsBodyProductMediaFileCreateUploadPart] `json:"parts,required"`
}

func (r FileNewParamsBodyProductMediaFileCreateUpload) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FileNewParamsBodyProductMediaFileCreateUploadPart struct {
	ChunkEnd             param.Field[int64]  `json:"chunk_end,required"`
	ChunkStart           param.Field[int64]  `json:"chunk_start,required"`
	Number               param.Field[int64]  `json:"number,required"`
	ChecksumSha256Base64 param.Field[string] `json:"checksum_sha256_base64"`
}

func (r FileNewParamsBodyProductMediaFileCreateUploadPart) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Schema to create a file to be used as an organization avatar.
type FileNewParamsBodyOrganizationAvatarFileCreate struct {
	// MIME type of the file. Only images are supported for this type of file.
	MimeType param.Field[string]                                               `json:"mime_type,required"`
	Name     param.Field[string]                                               `json:"name,required"`
	Service  param.Field[FileNewParamsBodyOrganizationAvatarFileCreateService] `json:"service,required"`
	// Size of the file. A maximum of 1 MB is allowed for this type of file.
	Size                 param.Field[int64]                                               `json:"size,required"`
	Upload               param.Field[FileNewParamsBodyOrganizationAvatarFileCreateUpload] `json:"upload,required"`
	ChecksumSha256Base64 param.Field[string]                                              `json:"checksum_sha256_base64"`
	// The organization ID.
	OrganizationID param.Field[string] `json:"organization_id" format:"uuid4"`
	Version        param.Field[string] `json:"version"`
}

func (r FileNewParamsBodyOrganizationAvatarFileCreate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FileNewParamsBodyOrganizationAvatarFileCreate) implementsFileNewParamsBodyUnion() {}

type FileNewParamsBodyOrganizationAvatarFileCreateService string

const (
	FileNewParamsBodyOrganizationAvatarFileCreateServiceOrganizationAvatar FileNewParamsBodyOrganizationAvatarFileCreateService = "organization_avatar"
)

func (r FileNewParamsBodyOrganizationAvatarFileCreateService) IsKnown() bool {
	switch r {
	case FileNewParamsBodyOrganizationAvatarFileCreateServiceOrganizationAvatar:
		return true
	}
	return false
}

type FileNewParamsBodyOrganizationAvatarFileCreateUpload struct {
	Parts param.Field[[]FileNewParamsBodyOrganizationAvatarFileCreateUploadPart] `json:"parts,required"`
}

func (r FileNewParamsBodyOrganizationAvatarFileCreateUpload) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FileNewParamsBodyOrganizationAvatarFileCreateUploadPart struct {
	ChunkEnd             param.Field[int64]  `json:"chunk_end,required"`
	ChunkStart           param.Field[int64]  `json:"chunk_start,required"`
	Number               param.Field[int64]  `json:"number,required"`
	ChecksumSha256Base64 param.Field[string] `json:"checksum_sha256_base64"`
}

func (r FileNewParamsBodyOrganizationAvatarFileCreateUploadPart) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FileNewParamsBodyService string

const (
	FileNewParamsBodyServiceDownloadable       FileNewParamsBodyService = "downloadable"
	FileNewParamsBodyServiceProductMedia       FileNewParamsBodyService = "product_media"
	FileNewParamsBodyServiceOrganizationAvatar FileNewParamsBodyService = "organization_avatar"
)

func (r FileNewParamsBodyService) IsKnown() bool {
	switch r {
	case FileNewParamsBodyServiceDownloadable, FileNewParamsBodyServiceProductMedia, FileNewParamsBodyServiceOrganizationAvatar:
		return true
	}
	return false
}

type FileUpdateParams struct {
	Name    param.Field[string] `json:"name"`
	Version param.Field[string] `json:"version"`
}

func (r FileUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FileListParams struct {
	// List of file IDs to get.
	IDs param.Field[[]string] `query:"ids" format:"uuid4"`
	// Size of a page, defaults to 10. Maximum is 100.
	Limit param.Field[int64] `query:"limit"`
	// The organization ID.
	OrganizationID param.Field[string] `query:"organization_id" format:"uuid4"`
	// Page number, defaults to 1.
	Page param.Field[int64] `query:"page"`
}

// URLQuery serializes [FileListParams]'s query parameters as `url.Values`.
func (r FileListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FileUploadedParams struct {
	// The file ID.
	PathID param.Field[string]                   `path:"id,required" format:"uuid4"`
	BodyID param.Field[string]                   `json:"id,required"`
	Parts  param.Field[[]FileUploadedParamsPart] `json:"parts,required"`
	Path   param.Field[string]                   `json:"path,required"`
}

func (r FileUploadedParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FileUploadedParamsPart struct {
	ChecksumEtag         param.Field[string] `json:"checksum_etag,required"`
	ChecksumSha256Base64 param.Field[string] `json:"checksum_sha256_base64,required"`
	Number               param.Field[int64]  `json:"number,required"`
}

func (r FileUploadedParamsPart) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
