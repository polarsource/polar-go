// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package polar

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/polarsource/polar-go/internal/requestconfig"
	"github.com/polarsource/polar-go/option"
)

// ArticleReceiverService contains methods and other services that help with
// interacting with the polar API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewArticleReceiverService] method instead.
type ArticleReceiverService struct {
	Options []option.RequestOption
}

// NewArticleReceiverService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewArticleReceiverService(opts ...option.RequestOption) (r *ArticleReceiverService) {
	r = &ArticleReceiverService{}
	r.Options = opts
	return
}

// Get number of potential receivers for an article.
func (r *ArticleReceiverService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *ArticleReceivers, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/articles/%s/receivers", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
