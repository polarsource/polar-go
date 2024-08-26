// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/polarsource/polar-go/internal/apijson"
	"github.com/polarsource/polar-go/internal/requestconfig"
)

type PolarPaginationPagination struct {
	MaxPage    int64                         `json:"max_page"`
	TotalCount int64                         `json:"total_count"`
	JSON       polarPaginationPaginationJSON `json:"-"`
}

// polarPaginationPaginationJSON contains the JSON metadata for the struct
// [PolarPaginationPagination]
type polarPaginationPaginationJSON struct {
	MaxPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PolarPaginationPagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r polarPaginationPaginationJSON) RawJSON() string {
	return r.raw
}

type PolarPagination[T any] struct {
	Items      []T                       `json:"items"`
	Pagination PolarPaginationPagination `json:"pagination"`
	JSON       polarPaginationJSON       `json:"-"`
	cfg        *requestconfig.RequestConfig
	res        *http.Response
}

// polarPaginationJSON contains the JSON metadata for the struct
// [PolarPagination[T]]
type polarPaginationJSON struct {
	Items       apijson.Field
	Pagination  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PolarPagination[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r polarPaginationJSON) RawJSON() string {
	return r.raw
}

// NextPage returns the next page as defined by this pagination style. When there
// is no next page, this function will return a 'nil' for the page value, but will
// not return an error
func (r *PolarPagination[T]) GetNextPage() (res *PolarPagination[T], err error) {
	u := r.cfg.Request.URL
	currentPage, err := strconv.Atoi(u.Query().Get("page"))
	if err != nil {
		currentPage = 1
	}
	if currentPage >= r.Pagination.MaxPage {
		return nil, nil
	}
	cfg := r.cfg.Clone(context.Background())
	query := cfg.Request.URL.Query()
	query.Set("page", fmt.Sprintf("%d", currentPage+1))
	cfg.Request.URL.RawQuery = query.Encode()
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *PolarPagination[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &PolarPagination[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type PolarPaginationAutoPager[T any] struct {
	page *PolarPagination[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewPolarPaginationAutoPager[T any](page *PolarPagination[T], err error) *PolarPaginationAutoPager[T] {
	return &PolarPaginationAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *PolarPaginationAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Items) == 0 {
		return false
	}
	if r.idx >= len(r.page.Items) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Items) == 0 {
			return false
		}
	}
	r.cur = r.page.Items[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *PolarPaginationAutoPager[T]) Current() T {
	return r.cur
}

func (r *PolarPaginationAutoPager[T]) Err() error {
	return r.err
}

func (r *PolarPaginationAutoPager[T]) Index() int {
	return r.run
}
