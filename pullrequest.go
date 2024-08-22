// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package polar

import (
	"context"
	"net/http"
	"net/url"

	"github.com/polarsource/polar-go/internal/apijson"
	"github.com/polarsource/polar-go/internal/apiquery"
	"github.com/polarsource/polar-go/internal/param"
	"github.com/polarsource/polar-go/internal/requestconfig"
	"github.com/polarsource/polar-go/option"
)

// PullRequestService contains methods and other services that help with
// interacting with the polar API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPullRequestService] method instead.
type PullRequestService struct {
	Options []option.RequestOption
}

// NewPullRequestService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPullRequestService(opts ...option.RequestOption) (r *PullRequestService) {
	r = &PullRequestService{}
	r.Options = opts
	return
}

// Search pull requests.
func (r *PullRequestService) Search(ctx context.Context, query PullRequestSearchParams, opts ...option.RequestOption) (res *PullRequestSearchResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/pull_requests/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type PullRequestSearchResponse struct {
	Pagination PullRequestSearchResponsePagination `json:"pagination,required"`
	Items      []PullRequestSearchResponseItem     `json:"items"`
	JSON       pullRequestSearchResponseJSON       `json:"-"`
}

// pullRequestSearchResponseJSON contains the JSON metadata for the struct
// [PullRequestSearchResponse]
type pullRequestSearchResponseJSON struct {
	Pagination  apijson.Field
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PullRequestSearchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pullRequestSearchResponseJSON) RawJSON() string {
	return r.raw
}

type PullRequestSearchResponsePagination struct {
	MaxPage    int64                                   `json:"max_page,required"`
	TotalCount int64                                   `json:"total_count,required"`
	JSON       pullRequestSearchResponsePaginationJSON `json:"-"`
}

// pullRequestSearchResponsePaginationJSON contains the JSON metadata for the
// struct [PullRequestSearchResponsePagination]
type pullRequestSearchResponsePaginationJSON struct {
	MaxPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PullRequestSearchResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pullRequestSearchResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type PullRequestSearchResponseItem struct {
	ID        string                               `json:"id,required" format:"uuid"`
	Additions int64                                `json:"additions,required"`
	Deletions int64                                `json:"deletions,required"`
	IsClosed  bool                                 `json:"is_closed,required"`
	IsMerged  bool                                 `json:"is_merged,required"`
	Number    int64                                `json:"number,required"`
	Title     string                               `json:"title,required"`
	Author    PullRequestSearchResponseItemsAuthor `json:"author,nullable"`
	JSON      pullRequestSearchResponseItemJSON    `json:"-"`
}

// pullRequestSearchResponseItemJSON contains the JSON metadata for the struct
// [PullRequestSearchResponseItem]
type pullRequestSearchResponseItemJSON struct {
	ID          apijson.Field
	Additions   apijson.Field
	Deletions   apijson.Field
	IsClosed    apijson.Field
	IsMerged    apijson.Field
	Number      apijson.Field
	Title       apijson.Field
	Author      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PullRequestSearchResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pullRequestSearchResponseItemJSON) RawJSON() string {
	return r.raw
}

type PullRequestSearchResponseItemsAuthor struct {
	ID        int64                                    `json:"id,required"`
	AvatarURL string                                   `json:"avatar_url,required" format:"uri"`
	HTMLURL   string                                   `json:"html_url,required" format:"uri"`
	Login     string                                   `json:"login,required"`
	JSON      pullRequestSearchResponseItemsAuthorJSON `json:"-"`
}

// pullRequestSearchResponseItemsAuthorJSON contains the JSON metadata for the
// struct [PullRequestSearchResponseItemsAuthor]
type pullRequestSearchResponseItemsAuthorJSON struct {
	ID          apijson.Field
	AvatarURL   apijson.Field
	HTMLURL     apijson.Field
	Login       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PullRequestSearchResponseItemsAuthor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pullRequestSearchResponseItemsAuthorJSON) RawJSON() string {
	return r.raw
}

type PullRequestSearchParams struct {
	// Search pull requests that are mentioning this issue
	ReferencesIssueID param.Field[string] `query:"references_issue_id" format:"uuid"`
}

// URLQuery serializes [PullRequestSearchParams]'s query parameters as
// `url.Values`.
func (r PullRequestSearchParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
