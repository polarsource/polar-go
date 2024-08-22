// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package polar

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/polar-go/internal/apijson"
	"github.com/stainless-sdks/polar-go/internal/param"
	"github.com/stainless-sdks/polar-go/internal/requestconfig"
	"github.com/stainless-sdks/polar-go/option"
)

// Oauth2RegisterService contains methods and other services that help with
// interacting with the polar API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOauth2RegisterService] method instead.
type Oauth2RegisterService struct {
	Options []option.RequestOption
}

// NewOauth2RegisterService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOauth2RegisterService(opts ...option.RequestOption) (r *Oauth2RegisterService) {
	r = &Oauth2RegisterService{}
	r.Options = opts
	return
}

// Create an OAuth2 client.
func (r *Oauth2RegisterService) New(ctx context.Context, body Oauth2RegisterNewParams, opts ...option.RequestOption) (res *Oauth2RegisterNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/oauth2/register"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get an OAuth2 client by Client ID.
func (r *Oauth2RegisterService) Get(ctx context.Context, clientID string, opts ...option.RequestOption) (res *Oauth2RegisterGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if clientID == "" {
		err = errors.New("missing required client_id parameter")
		return
	}
	path := fmt.Sprintf("v1/oauth2/register/%s", clientID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an OAuth2 client.
func (r *Oauth2RegisterService) Update(ctx context.Context, params Oauth2RegisterUpdateParams, opts ...option.RequestOption) (res *Oauth2RegisterUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.PathClientID.Value == "" {
		err = errors.New("missing required path_client_id parameter")
		return
	}
	path := fmt.Sprintf("v1/oauth2/register/%s", params.PathClientID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Delete an OAuth2 client.
func (r *Oauth2RegisterService) Delete(ctx context.Context, clientID string, opts ...option.RequestOption) (res *Oauth2RegisterDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if clientID == "" {
		err = errors.New("missing required client_id parameter")
		return
	}
	path := fmt.Sprintf("v1/oauth2/register/%s", clientID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type Oauth2RegisterNewResponse = interface{}

type Oauth2RegisterGetResponse = interface{}

type Oauth2RegisterUpdateResponse = interface{}

type Oauth2RegisterDeleteResponse = interface{}

type Oauth2RegisterNewParams struct {
	ClientName              param.Field[string]                                         `json:"client_name,required"`
	RedirectUris            param.Field[[]string]                                       `json:"redirect_uris,required" format:"uri"`
	ClientUri               param.Field[string]                                         `json:"client_uri"`
	GrantTypes              param.Field[[]Oauth2RegisterNewParamsGrantType]             `json:"grant_types"`
	LogoUri                 param.Field[string]                                         `json:"logo_uri" format:"uri"`
	PolicyUri               param.Field[string]                                         `json:"policy_uri" format:"uri"`
	ResponseTypes           param.Field[[]Oauth2RegisterNewParamsResponseType]          `json:"response_types"`
	Scope                   param.Field[string]                                         `json:"scope"`
	TokenEndpointAuthMethod param.Field[Oauth2RegisterNewParamsTokenEndpointAuthMethod] `json:"token_endpoint_auth_method"`
	TosUri                  param.Field[string]                                         `json:"tos_uri" format:"uri"`
}

func (r Oauth2RegisterNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Oauth2RegisterNewParamsGrantType string

const (
	Oauth2RegisterNewParamsGrantTypeAuthorizationCode Oauth2RegisterNewParamsGrantType = "authorization_code"
	Oauth2RegisterNewParamsGrantTypeRefreshToken      Oauth2RegisterNewParamsGrantType = "refresh_token"
)

func (r Oauth2RegisterNewParamsGrantType) IsKnown() bool {
	switch r {
	case Oauth2RegisterNewParamsGrantTypeAuthorizationCode, Oauth2RegisterNewParamsGrantTypeRefreshToken:
		return true
	}
	return false
}

type Oauth2RegisterNewParamsResponseType string

const (
	Oauth2RegisterNewParamsResponseTypeCode Oauth2RegisterNewParamsResponseType = "code"
)

func (r Oauth2RegisterNewParamsResponseType) IsKnown() bool {
	switch r {
	case Oauth2RegisterNewParamsResponseTypeCode:
		return true
	}
	return false
}

type Oauth2RegisterNewParamsTokenEndpointAuthMethod string

const (
	Oauth2RegisterNewParamsTokenEndpointAuthMethodClientSecretBasic Oauth2RegisterNewParamsTokenEndpointAuthMethod = "client_secret_basic"
	Oauth2RegisterNewParamsTokenEndpointAuthMethodClientSecretPost  Oauth2RegisterNewParamsTokenEndpointAuthMethod = "client_secret_post"
	Oauth2RegisterNewParamsTokenEndpointAuthMethodNone              Oauth2RegisterNewParamsTokenEndpointAuthMethod = "none"
)

func (r Oauth2RegisterNewParamsTokenEndpointAuthMethod) IsKnown() bool {
	switch r {
	case Oauth2RegisterNewParamsTokenEndpointAuthMethodClientSecretBasic, Oauth2RegisterNewParamsTokenEndpointAuthMethodClientSecretPost, Oauth2RegisterNewParamsTokenEndpointAuthMethodNone:
		return true
	}
	return false
}

type Oauth2RegisterUpdateParams struct {
	PathClientID            param.Field[string]                                            `path:"client_id,required"`
	BodyClientID            param.Field[string]                                            `json:"client_id,required"`
	ClientName              param.Field[string]                                            `json:"client_name,required"`
	RedirectUris            param.Field[[]string]                                          `json:"redirect_uris,required" format:"uri"`
	ClientUri               param.Field[string]                                            `json:"client_uri"`
	GrantTypes              param.Field[[]Oauth2RegisterUpdateParamsGrantType]             `json:"grant_types"`
	LogoUri                 param.Field[string]                                            `json:"logo_uri" format:"uri"`
	PolicyUri               param.Field[string]                                            `json:"policy_uri" format:"uri"`
	ResponseTypes           param.Field[[]Oauth2RegisterUpdateParamsResponseType]          `json:"response_types"`
	Scope                   param.Field[string]                                            `json:"scope"`
	TokenEndpointAuthMethod param.Field[Oauth2RegisterUpdateParamsTokenEndpointAuthMethod] `json:"token_endpoint_auth_method"`
	TosUri                  param.Field[string]                                            `json:"tos_uri" format:"uri"`
}

func (r Oauth2RegisterUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Oauth2RegisterUpdateParamsGrantType string

const (
	Oauth2RegisterUpdateParamsGrantTypeAuthorizationCode Oauth2RegisterUpdateParamsGrantType = "authorization_code"
	Oauth2RegisterUpdateParamsGrantTypeRefreshToken      Oauth2RegisterUpdateParamsGrantType = "refresh_token"
)

func (r Oauth2RegisterUpdateParamsGrantType) IsKnown() bool {
	switch r {
	case Oauth2RegisterUpdateParamsGrantTypeAuthorizationCode, Oauth2RegisterUpdateParamsGrantTypeRefreshToken:
		return true
	}
	return false
}

type Oauth2RegisterUpdateParamsResponseType string

const (
	Oauth2RegisterUpdateParamsResponseTypeCode Oauth2RegisterUpdateParamsResponseType = "code"
)

func (r Oauth2RegisterUpdateParamsResponseType) IsKnown() bool {
	switch r {
	case Oauth2RegisterUpdateParamsResponseTypeCode:
		return true
	}
	return false
}

type Oauth2RegisterUpdateParamsTokenEndpointAuthMethod string

const (
	Oauth2RegisterUpdateParamsTokenEndpointAuthMethodClientSecretBasic Oauth2RegisterUpdateParamsTokenEndpointAuthMethod = "client_secret_basic"
	Oauth2RegisterUpdateParamsTokenEndpointAuthMethodClientSecretPost  Oauth2RegisterUpdateParamsTokenEndpointAuthMethod = "client_secret_post"
	Oauth2RegisterUpdateParamsTokenEndpointAuthMethodNone              Oauth2RegisterUpdateParamsTokenEndpointAuthMethod = "none"
)

func (r Oauth2RegisterUpdateParamsTokenEndpointAuthMethod) IsKnown() bool {
	switch r {
	case Oauth2RegisterUpdateParamsTokenEndpointAuthMethodClientSecretBasic, Oauth2RegisterUpdateParamsTokenEndpointAuthMethodClientSecretPost, Oauth2RegisterUpdateParamsTokenEndpointAuthMethodNone:
		return true
	}
	return false
}
