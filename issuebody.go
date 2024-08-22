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

// IssueBodyService contains methods and other services that help with interacting
// with the polar API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIssueBodyService] method instead.
type IssueBodyService struct {
	Options []option.RequestOption
}

// NewIssueBodyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewIssueBodyService(opts ...option.RequestOption) (r *IssueBodyService) {
	r = &IssueBodyService{}
	r.Options = opts
	return
}

// Get Body
func (r *IssueBodyService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *IssueBodyGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/issues/%s/body", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type IssueBodyGetResponse = interface{}
