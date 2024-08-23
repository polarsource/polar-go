// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package polar

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/polarsource/polar-go/internal/apijson"
	"github.com/polarsource/polar-go/internal/apiquery"
	"github.com/polarsource/polar-go/internal/param"
	"github.com/polarsource/polar-go/internal/requestconfig"
	"github.com/polarsource/polar-go/option"
	"github.com/tidwall/gjson"
)

// Oauth2Service contains methods and other services that help with interacting
// with the polar API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOauth2Service] method instead.
type Oauth2Service struct {
	Options  []option.RequestOption
	Register *Oauth2RegisterService
	Userinfo *Oauth2UserinfoService
}

// NewOauth2Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewOauth2Service(opts ...option.RequestOption) (r *Oauth2Service) {
	r = &Oauth2Service{}
	r.Options = opts
	r.Register = NewOauth2RegisterService(opts...)
	r.Userinfo = NewOauth2UserinfoService(opts...)
	return
}

// List OAuth2 clients.
func (r *Oauth2Service) List(ctx context.Context, query Oauth2ListParams, opts ...option.RequestOption) (res *Oauth2ListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/oauth2/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Authorize
func (r *Oauth2Service) Authorize(ctx context.Context, opts ...option.RequestOption) (res *Oauth2AuthorizeResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/oauth2/authorize"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get information about an access token.
func (r *Oauth2Service) Introspect(ctx context.Context, body Oauth2IntrospectParams, opts ...option.RequestOption) (res *IntrospectTokenResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/oauth2/introspect"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Revoke an access token or a refresh token.
func (r *Oauth2Service) Revoke(ctx context.Context, body Oauth2RevokeParams, opts ...option.RequestOption) (res *RevokeTokenResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/oauth2/revoke"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Request an access token using a valid grant.
func (r *Oauth2Service) Token(ctx context.Context, body Oauth2TokenParams, opts ...option.RequestOption) (res *TokenResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/oauth2/token"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AuthorizeResponseOrganization struct {
	Client        AuthorizeResponseOrganizationClient         `json:"client,required"`
	Organizations []AuthorizeResponseOrganizationOrganization `json:"organizations,required"`
	Scopes        []AuthorizeResponseOrganizationScope        `json:"scopes,required"`
	Sub           AuthorizeResponseOrganizationSub            `json:"sub,required,nullable"`
	SubType       AuthorizeResponseOrganizationSubType        `json:"sub_type,required"`
	JSON          authorizeResponseOrganizationJSON           `json:"-"`
}

// authorizeResponseOrganizationJSON contains the JSON metadata for the struct
// [AuthorizeResponseOrganization]
type authorizeResponseOrganizationJSON struct {
	Client        apijson.Field
	Organizations apijson.Field
	Scopes        apijson.Field
	Sub           apijson.Field
	SubType       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AuthorizeResponseOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authorizeResponseOrganizationJSON) RawJSON() string {
	return r.raw
}

func (r AuthorizeResponseOrganization) implementsOauth2AuthorizeResponse() {}

type AuthorizeResponseOrganizationClient struct {
	ClientID   string `json:"client_id,required"`
	ClientName string `json:"client_name,required,nullable"`
	ClientUri  string `json:"client_uri,required,nullable"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	LogoUri   string    `json:"logo_uri,required,nullable"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time                               `json:"modified_at,required,nullable" format:"date-time"`
	PolicyUri  string                                  `json:"policy_uri,required,nullable"`
	TosUri     string                                  `json:"tos_uri,required,nullable"`
	JSON       authorizeResponseOrganizationClientJSON `json:"-"`
}

// authorizeResponseOrganizationClientJSON contains the JSON metadata for the
// struct [AuthorizeResponseOrganizationClient]
type authorizeResponseOrganizationClientJSON struct {
	ClientID    apijson.Field
	ClientName  apijson.Field
	ClientUri   apijson.Field
	CreatedAt   apijson.Field
	LogoUri     apijson.Field
	ModifiedAt  apijson.Field
	PolicyUri   apijson.Field
	TosUri      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthorizeResponseOrganizationClient) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authorizeResponseOrganizationClientJSON) RawJSON() string {
	return r.raw
}

type AuthorizeResponseOrganizationOrganization struct {
	ID        string                                        `json:"id,required" format:"uuid4"`
	AvatarURL string                                        `json:"avatar_url,required,nullable"`
	Slug      string                                        `json:"slug,required"`
	JSON      authorizeResponseOrganizationOrganizationJSON `json:"-"`
}

// authorizeResponseOrganizationOrganizationJSON contains the JSON metadata for the
// struct [AuthorizeResponseOrganizationOrganization]
type authorizeResponseOrganizationOrganizationJSON struct {
	ID          apijson.Field
	AvatarURL   apijson.Field
	Slug        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthorizeResponseOrganizationOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authorizeResponseOrganizationOrganizationJSON) RawJSON() string {
	return r.raw
}

type AuthorizeResponseOrganizationScope string

const (
	AuthorizeResponseOrganizationScopeOpenid                          AuthorizeResponseOrganizationScope = "openid"
	AuthorizeResponseOrganizationScopeProfile                         AuthorizeResponseOrganizationScope = "profile"
	AuthorizeResponseOrganizationScopeEmail                           AuthorizeResponseOrganizationScope = "email"
	AuthorizeResponseOrganizationScopeUserRead                        AuthorizeResponseOrganizationScope = "user:read"
	AuthorizeResponseOrganizationScopeAdmin                           AuthorizeResponseOrganizationScope = "admin"
	AuthorizeResponseOrganizationScopeWebDefault                      AuthorizeResponseOrganizationScope = "web_default"
	AuthorizeResponseOrganizationScopeOrganizationsRead               AuthorizeResponseOrganizationScope = "organizations:read"
	AuthorizeResponseOrganizationScopeOrganizationsWrite              AuthorizeResponseOrganizationScope = "organizations:write"
	AuthorizeResponseOrganizationScopeProductsRead                    AuthorizeResponseOrganizationScope = "products:read"
	AuthorizeResponseOrganizationScopeProductsWrite                   AuthorizeResponseOrganizationScope = "products:write"
	AuthorizeResponseOrganizationScopeBenefitsRead                    AuthorizeResponseOrganizationScope = "benefits:read"
	AuthorizeResponseOrganizationScopeBenefitsWrite                   AuthorizeResponseOrganizationScope = "benefits:write"
	AuthorizeResponseOrganizationScopeFilesRead                       AuthorizeResponseOrganizationScope = "files:read"
	AuthorizeResponseOrganizationScopeFilesWrite                      AuthorizeResponseOrganizationScope = "files:write"
	AuthorizeResponseOrganizationScopeSubscriptionsRead               AuthorizeResponseOrganizationScope = "subscriptions:read"
	AuthorizeResponseOrganizationScopeSubscriptionsWrite              AuthorizeResponseOrganizationScope = "subscriptions:write"
	AuthorizeResponseOrganizationScopeOrdersRead                      AuthorizeResponseOrganizationScope = "orders:read"
	AuthorizeResponseOrganizationScopeMetricsRead                     AuthorizeResponseOrganizationScope = "metrics:read"
	AuthorizeResponseOrganizationScopeArticlesRead                    AuthorizeResponseOrganizationScope = "articles:read"
	AuthorizeResponseOrganizationScopeArticlesWrite                   AuthorizeResponseOrganizationScope = "articles:write"
	AuthorizeResponseOrganizationScopeWebhooksRead                    AuthorizeResponseOrganizationScope = "webhooks:read"
	AuthorizeResponseOrganizationScopeWebhooksWrite                   AuthorizeResponseOrganizationScope = "webhooks:write"
	AuthorizeResponseOrganizationScopeExternalOrganizationsRead       AuthorizeResponseOrganizationScope = "external_organizations:read"
	AuthorizeResponseOrganizationScopeRepositoriesRead                AuthorizeResponseOrganizationScope = "repositories:read"
	AuthorizeResponseOrganizationScopeRepositoriesWrite               AuthorizeResponseOrganizationScope = "repositories:write"
	AuthorizeResponseOrganizationScopeIssuesRead                      AuthorizeResponseOrganizationScope = "issues:read"
	AuthorizeResponseOrganizationScopeIssuesWrite                     AuthorizeResponseOrganizationScope = "issues:write"
	AuthorizeResponseOrganizationScopeUserBenefitsRead                AuthorizeResponseOrganizationScope = "user:benefits:read"
	AuthorizeResponseOrganizationScopeUserOrdersRead                  AuthorizeResponseOrganizationScope = "user:orders:read"
	AuthorizeResponseOrganizationScopeUserSubscriptionsRead           AuthorizeResponseOrganizationScope = "user:subscriptions:read"
	AuthorizeResponseOrganizationScopeUserSubscriptionsWrite          AuthorizeResponseOrganizationScope = "user:subscriptions:write"
	AuthorizeResponseOrganizationScopeUserDownloadablesRead           AuthorizeResponseOrganizationScope = "user:downloadables:read"
	AuthorizeResponseOrganizationScopeUserAdvertisementCampaignsRead  AuthorizeResponseOrganizationScope = "user:advertisement_campaigns:read"
	AuthorizeResponseOrganizationScopeUserAdvertisementCampaignsWrite AuthorizeResponseOrganizationScope = "user:advertisement_campaigns:write"
)

func (r AuthorizeResponseOrganizationScope) IsKnown() bool {
	switch r {
	case AuthorizeResponseOrganizationScopeOpenid, AuthorizeResponseOrganizationScopeProfile, AuthorizeResponseOrganizationScopeEmail, AuthorizeResponseOrganizationScopeUserRead, AuthorizeResponseOrganizationScopeAdmin, AuthorizeResponseOrganizationScopeWebDefault, AuthorizeResponseOrganizationScopeOrganizationsRead, AuthorizeResponseOrganizationScopeOrganizationsWrite, AuthorizeResponseOrganizationScopeProductsRead, AuthorizeResponseOrganizationScopeProductsWrite, AuthorizeResponseOrganizationScopeBenefitsRead, AuthorizeResponseOrganizationScopeBenefitsWrite, AuthorizeResponseOrganizationScopeFilesRead, AuthorizeResponseOrganizationScopeFilesWrite, AuthorizeResponseOrganizationScopeSubscriptionsRead, AuthorizeResponseOrganizationScopeSubscriptionsWrite, AuthorizeResponseOrganizationScopeOrdersRead, AuthorizeResponseOrganizationScopeMetricsRead, AuthorizeResponseOrganizationScopeArticlesRead, AuthorizeResponseOrganizationScopeArticlesWrite, AuthorizeResponseOrganizationScopeWebhooksRead, AuthorizeResponseOrganizationScopeWebhooksWrite, AuthorizeResponseOrganizationScopeExternalOrganizationsRead, AuthorizeResponseOrganizationScopeRepositoriesRead, AuthorizeResponseOrganizationScopeRepositoriesWrite, AuthorizeResponseOrganizationScopeIssuesRead, AuthorizeResponseOrganizationScopeIssuesWrite, AuthorizeResponseOrganizationScopeUserBenefitsRead, AuthorizeResponseOrganizationScopeUserOrdersRead, AuthorizeResponseOrganizationScopeUserSubscriptionsRead, AuthorizeResponseOrganizationScopeUserSubscriptionsWrite, AuthorizeResponseOrganizationScopeUserDownloadablesRead, AuthorizeResponseOrganizationScopeUserAdvertisementCampaignsRead, AuthorizeResponseOrganizationScopeUserAdvertisementCampaignsWrite:
		return true
	}
	return false
}

type AuthorizeResponseOrganizationSub struct {
	ID        string                               `json:"id,required" format:"uuid4"`
	AvatarURL string                               `json:"avatar_url,required,nullable"`
	Slug      string                               `json:"slug,required"`
	JSON      authorizeResponseOrganizationSubJSON `json:"-"`
}

// authorizeResponseOrganizationSubJSON contains the JSON metadata for the struct
// [AuthorizeResponseOrganizationSub]
type authorizeResponseOrganizationSubJSON struct {
	ID          apijson.Field
	AvatarURL   apijson.Field
	Slug        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthorizeResponseOrganizationSub) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authorizeResponseOrganizationSubJSON) RawJSON() string {
	return r.raw
}

type AuthorizeResponseOrganizationSubType string

const (
	AuthorizeResponseOrganizationSubTypeOrganization AuthorizeResponseOrganizationSubType = "organization"
)

func (r AuthorizeResponseOrganizationSubType) IsKnown() bool {
	switch r {
	case AuthorizeResponseOrganizationSubTypeOrganization:
		return true
	}
	return false
}

type AuthorizeResponseUser struct {
	Client  AuthorizeResponseUserClient  `json:"client,required"`
	Scopes  []AuthorizeResponseUserScope `json:"scopes,required"`
	Sub     AuthorizeResponseUserSub     `json:"sub,required,nullable"`
	SubType AuthorizeResponseUserSubType `json:"sub_type,required"`
	JSON    authorizeResponseUserJSON    `json:"-"`
}

// authorizeResponseUserJSON contains the JSON metadata for the struct
// [AuthorizeResponseUser]
type authorizeResponseUserJSON struct {
	Client      apijson.Field
	Scopes      apijson.Field
	Sub         apijson.Field
	SubType     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthorizeResponseUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authorizeResponseUserJSON) RawJSON() string {
	return r.raw
}

func (r AuthorizeResponseUser) implementsOauth2AuthorizeResponse() {}

type AuthorizeResponseUserClient struct {
	ClientID   string `json:"client_id,required"`
	ClientName string `json:"client_name,required,nullable"`
	ClientUri  string `json:"client_uri,required,nullable"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	LogoUri   string    `json:"logo_uri,required,nullable"`
	// Last modification timestamp of the object.
	ModifiedAt time.Time                       `json:"modified_at,required,nullable" format:"date-time"`
	PolicyUri  string                          `json:"policy_uri,required,nullable"`
	TosUri     string                          `json:"tos_uri,required,nullable"`
	JSON       authorizeResponseUserClientJSON `json:"-"`
}

// authorizeResponseUserClientJSON contains the JSON metadata for the struct
// [AuthorizeResponseUserClient]
type authorizeResponseUserClientJSON struct {
	ClientID    apijson.Field
	ClientName  apijson.Field
	ClientUri   apijson.Field
	CreatedAt   apijson.Field
	LogoUri     apijson.Field
	ModifiedAt  apijson.Field
	PolicyUri   apijson.Field
	TosUri      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthorizeResponseUserClient) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authorizeResponseUserClientJSON) RawJSON() string {
	return r.raw
}

type AuthorizeResponseUserScope string

const (
	AuthorizeResponseUserScopeOpenid                          AuthorizeResponseUserScope = "openid"
	AuthorizeResponseUserScopeProfile                         AuthorizeResponseUserScope = "profile"
	AuthorizeResponseUserScopeEmail                           AuthorizeResponseUserScope = "email"
	AuthorizeResponseUserScopeUserRead                        AuthorizeResponseUserScope = "user:read"
	AuthorizeResponseUserScopeAdmin                           AuthorizeResponseUserScope = "admin"
	AuthorizeResponseUserScopeWebDefault                      AuthorizeResponseUserScope = "web_default"
	AuthorizeResponseUserScopeOrganizationsRead               AuthorizeResponseUserScope = "organizations:read"
	AuthorizeResponseUserScopeOrganizationsWrite              AuthorizeResponseUserScope = "organizations:write"
	AuthorizeResponseUserScopeProductsRead                    AuthorizeResponseUserScope = "products:read"
	AuthorizeResponseUserScopeProductsWrite                   AuthorizeResponseUserScope = "products:write"
	AuthorizeResponseUserScopeBenefitsRead                    AuthorizeResponseUserScope = "benefits:read"
	AuthorizeResponseUserScopeBenefitsWrite                   AuthorizeResponseUserScope = "benefits:write"
	AuthorizeResponseUserScopeFilesRead                       AuthorizeResponseUserScope = "files:read"
	AuthorizeResponseUserScopeFilesWrite                      AuthorizeResponseUserScope = "files:write"
	AuthorizeResponseUserScopeSubscriptionsRead               AuthorizeResponseUserScope = "subscriptions:read"
	AuthorizeResponseUserScopeSubscriptionsWrite              AuthorizeResponseUserScope = "subscriptions:write"
	AuthorizeResponseUserScopeOrdersRead                      AuthorizeResponseUserScope = "orders:read"
	AuthorizeResponseUserScopeMetricsRead                     AuthorizeResponseUserScope = "metrics:read"
	AuthorizeResponseUserScopeArticlesRead                    AuthorizeResponseUserScope = "articles:read"
	AuthorizeResponseUserScopeArticlesWrite                   AuthorizeResponseUserScope = "articles:write"
	AuthorizeResponseUserScopeWebhooksRead                    AuthorizeResponseUserScope = "webhooks:read"
	AuthorizeResponseUserScopeWebhooksWrite                   AuthorizeResponseUserScope = "webhooks:write"
	AuthorizeResponseUserScopeExternalOrganizationsRead       AuthorizeResponseUserScope = "external_organizations:read"
	AuthorizeResponseUserScopeRepositoriesRead                AuthorizeResponseUserScope = "repositories:read"
	AuthorizeResponseUserScopeRepositoriesWrite               AuthorizeResponseUserScope = "repositories:write"
	AuthorizeResponseUserScopeIssuesRead                      AuthorizeResponseUserScope = "issues:read"
	AuthorizeResponseUserScopeIssuesWrite                     AuthorizeResponseUserScope = "issues:write"
	AuthorizeResponseUserScopeUserBenefitsRead                AuthorizeResponseUserScope = "user:benefits:read"
	AuthorizeResponseUserScopeUserOrdersRead                  AuthorizeResponseUserScope = "user:orders:read"
	AuthorizeResponseUserScopeUserSubscriptionsRead           AuthorizeResponseUserScope = "user:subscriptions:read"
	AuthorizeResponseUserScopeUserSubscriptionsWrite          AuthorizeResponseUserScope = "user:subscriptions:write"
	AuthorizeResponseUserScopeUserDownloadablesRead           AuthorizeResponseUserScope = "user:downloadables:read"
	AuthorizeResponseUserScopeUserAdvertisementCampaignsRead  AuthorizeResponseUserScope = "user:advertisement_campaigns:read"
	AuthorizeResponseUserScopeUserAdvertisementCampaignsWrite AuthorizeResponseUserScope = "user:advertisement_campaigns:write"
)

func (r AuthorizeResponseUserScope) IsKnown() bool {
	switch r {
	case AuthorizeResponseUserScopeOpenid, AuthorizeResponseUserScopeProfile, AuthorizeResponseUserScopeEmail, AuthorizeResponseUserScopeUserRead, AuthorizeResponseUserScopeAdmin, AuthorizeResponseUserScopeWebDefault, AuthorizeResponseUserScopeOrganizationsRead, AuthorizeResponseUserScopeOrganizationsWrite, AuthorizeResponseUserScopeProductsRead, AuthorizeResponseUserScopeProductsWrite, AuthorizeResponseUserScopeBenefitsRead, AuthorizeResponseUserScopeBenefitsWrite, AuthorizeResponseUserScopeFilesRead, AuthorizeResponseUserScopeFilesWrite, AuthorizeResponseUserScopeSubscriptionsRead, AuthorizeResponseUserScopeSubscriptionsWrite, AuthorizeResponseUserScopeOrdersRead, AuthorizeResponseUserScopeMetricsRead, AuthorizeResponseUserScopeArticlesRead, AuthorizeResponseUserScopeArticlesWrite, AuthorizeResponseUserScopeWebhooksRead, AuthorizeResponseUserScopeWebhooksWrite, AuthorizeResponseUserScopeExternalOrganizationsRead, AuthorizeResponseUserScopeRepositoriesRead, AuthorizeResponseUserScopeRepositoriesWrite, AuthorizeResponseUserScopeIssuesRead, AuthorizeResponseUserScopeIssuesWrite, AuthorizeResponseUserScopeUserBenefitsRead, AuthorizeResponseUserScopeUserOrdersRead, AuthorizeResponseUserScopeUserSubscriptionsRead, AuthorizeResponseUserScopeUserSubscriptionsWrite, AuthorizeResponseUserScopeUserDownloadablesRead, AuthorizeResponseUserScopeUserAdvertisementCampaignsRead, AuthorizeResponseUserScopeUserAdvertisementCampaignsWrite:
		return true
	}
	return false
}

type AuthorizeResponseUserSub struct {
	ID        string                       `json:"id,required" format:"uuid4"`
	AvatarURL string                       `json:"avatar_url,required,nullable"`
	Email     string                       `json:"email,required" format:"email"`
	Username  string                       `json:"username,required"`
	JSON      authorizeResponseUserSubJSON `json:"-"`
}

// authorizeResponseUserSubJSON contains the JSON metadata for the struct
// [AuthorizeResponseUserSub]
type authorizeResponseUserSubJSON struct {
	ID          apijson.Field
	AvatarURL   apijson.Field
	Email       apijson.Field
	Username    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthorizeResponseUserSub) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authorizeResponseUserSubJSON) RawJSON() string {
	return r.raw
}

type AuthorizeResponseUserSubType string

const (
	AuthorizeResponseUserSubTypeUser AuthorizeResponseUserSubType = "user"
)

func (r AuthorizeResponseUserSubType) IsKnown() bool {
	switch r {
	case AuthorizeResponseUserSubTypeUser:
		return true
	}
	return false
}

type IntrospectTokenResponse struct {
	Active    bool                             `json:"active,required"`
	Aud       string                           `json:"aud,required"`
	ClientID  string                           `json:"client_id,required"`
	Exp       int64                            `json:"exp,required"`
	Iat       int64                            `json:"iat,required"`
	Iss       string                           `json:"iss,required"`
	Scope     string                           `json:"scope,required"`
	Sub       string                           `json:"sub,required"`
	SubType   IntrospectTokenResponseSubType   `json:"sub_type,required"`
	TokenType IntrospectTokenResponseTokenType `json:"token_type,required"`
	JSON      introspectTokenResponseJSON      `json:"-"`
}

// introspectTokenResponseJSON contains the JSON metadata for the struct
// [IntrospectTokenResponse]
type introspectTokenResponseJSON struct {
	Active      apijson.Field
	Aud         apijson.Field
	ClientID    apijson.Field
	Exp         apijson.Field
	Iat         apijson.Field
	Iss         apijson.Field
	Scope       apijson.Field
	Sub         apijson.Field
	SubType     apijson.Field
	TokenType   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntrospectTokenResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r introspectTokenResponseJSON) RawJSON() string {
	return r.raw
}

type IntrospectTokenResponseSubType string

const (
	IntrospectTokenResponseSubTypeUser         IntrospectTokenResponseSubType = "user"
	IntrospectTokenResponseSubTypeOrganization IntrospectTokenResponseSubType = "organization"
)

func (r IntrospectTokenResponseSubType) IsKnown() bool {
	switch r {
	case IntrospectTokenResponseSubTypeUser, IntrospectTokenResponseSubTypeOrganization:
		return true
	}
	return false
}

type IntrospectTokenResponseTokenType string

const (
	IntrospectTokenResponseTokenTypeAccessToken  IntrospectTokenResponseTokenType = "access_token"
	IntrospectTokenResponseTokenTypeRefreshToken IntrospectTokenResponseTokenType = "refresh_token"
)

func (r IntrospectTokenResponseTokenType) IsKnown() bool {
	switch r {
	case IntrospectTokenResponseTokenTypeAccessToken, IntrospectTokenResponseTokenTypeRefreshToken:
		return true
	}
	return false
}

type RevokeTokenResponse = interface{}

type TokenResponse struct {
	AccessToken  string                 `json:"access_token,required"`
	ExpiresIn    int64                  `json:"expires_in,required"`
	IDToken      string                 `json:"id_token,required"`
	RefreshToken string                 `json:"refresh_token,required,nullable"`
	Scope        string                 `json:"scope,required"`
	TokenType    TokenResponseTokenType `json:"token_type,required"`
	JSON         tokenResponseJSON      `json:"-"`
}

// tokenResponseJSON contains the JSON metadata for the struct [TokenResponse]
type tokenResponseJSON struct {
	AccessToken  apijson.Field
	ExpiresIn    apijson.Field
	IDToken      apijson.Field
	RefreshToken apijson.Field
	Scope        apijson.Field
	TokenType    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *TokenResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenResponseJSON) RawJSON() string {
	return r.raw
}

type TokenResponseTokenType string

const (
	TokenResponseTokenTypeBearer TokenResponseTokenType = "Bearer"
)

func (r TokenResponseTokenType) IsKnown() bool {
	switch r {
	case TokenResponseTokenTypeBearer:
		return true
	}
	return false
}

type Oauth2ListResponse struct {
	Pagination Oauth2ListResponsePagination `json:"pagination,required"`
	Items      []Oauth2ListResponseItem     `json:"items"`
	JSON       oauth2ListResponseJSON       `json:"-"`
}

// oauth2ListResponseJSON contains the JSON metadata for the struct
// [Oauth2ListResponse]
type oauth2ListResponseJSON struct {
	Pagination  apijson.Field
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Oauth2ListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oauth2ListResponseJSON) RawJSON() string {
	return r.raw
}

type Oauth2ListResponsePagination struct {
	MaxPage    int64                            `json:"max_page,required"`
	TotalCount int64                            `json:"total_count,required"`
	JSON       oauth2ListResponsePaginationJSON `json:"-"`
}

// oauth2ListResponsePaginationJSON contains the JSON metadata for the struct
// [Oauth2ListResponsePagination]
type oauth2ListResponsePaginationJSON struct {
	MaxPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Oauth2ListResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oauth2ListResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type Oauth2ListResponseItem struct {
	ClientID              string `json:"client_id,required"`
	ClientIDIssuedAt      int64  `json:"client_id_issued_at,required"`
	ClientName            string `json:"client_name,required"`
	ClientSecret          string `json:"client_secret,required"`
	ClientSecretExpiresAt int64  `json:"client_secret_expires_at,required"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Last modification timestamp of the object.
	ModifiedAt              time.Time                                      `json:"modified_at,required,nullable" format:"date-time"`
	RedirectUris            []string                                       `json:"redirect_uris,required" format:"uri"`
	ClientUri               string                                         `json:"client_uri,nullable"`
	GrantTypes              []Oauth2ListResponseItemsGrantType             `json:"grant_types"`
	LogoUri                 string                                         `json:"logo_uri,nullable" format:"uri"`
	PolicyUri               string                                         `json:"policy_uri,nullable" format:"uri"`
	ResponseTypes           []Oauth2ListResponseItemsResponseType          `json:"response_types"`
	Scope                   string                                         `json:"scope"`
	TokenEndpointAuthMethod Oauth2ListResponseItemsTokenEndpointAuthMethod `json:"token_endpoint_auth_method"`
	TosUri                  string                                         `json:"tos_uri,nullable" format:"uri"`
	JSON                    oauth2ListResponseItemJSON                     `json:"-"`
}

// oauth2ListResponseItemJSON contains the JSON metadata for the struct
// [Oauth2ListResponseItem]
type oauth2ListResponseItemJSON struct {
	ClientID                apijson.Field
	ClientIDIssuedAt        apijson.Field
	ClientName              apijson.Field
	ClientSecret            apijson.Field
	ClientSecretExpiresAt   apijson.Field
	CreatedAt               apijson.Field
	ModifiedAt              apijson.Field
	RedirectUris            apijson.Field
	ClientUri               apijson.Field
	GrantTypes              apijson.Field
	LogoUri                 apijson.Field
	PolicyUri               apijson.Field
	ResponseTypes           apijson.Field
	Scope                   apijson.Field
	TokenEndpointAuthMethod apijson.Field
	TosUri                  apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *Oauth2ListResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oauth2ListResponseItemJSON) RawJSON() string {
	return r.raw
}

type Oauth2ListResponseItemsGrantType string

const (
	Oauth2ListResponseItemsGrantTypeAuthorizationCode Oauth2ListResponseItemsGrantType = "authorization_code"
	Oauth2ListResponseItemsGrantTypeRefreshToken      Oauth2ListResponseItemsGrantType = "refresh_token"
)

func (r Oauth2ListResponseItemsGrantType) IsKnown() bool {
	switch r {
	case Oauth2ListResponseItemsGrantTypeAuthorizationCode, Oauth2ListResponseItemsGrantTypeRefreshToken:
		return true
	}
	return false
}

type Oauth2ListResponseItemsResponseType string

const (
	Oauth2ListResponseItemsResponseTypeCode Oauth2ListResponseItemsResponseType = "code"
)

func (r Oauth2ListResponseItemsResponseType) IsKnown() bool {
	switch r {
	case Oauth2ListResponseItemsResponseTypeCode:
		return true
	}
	return false
}

type Oauth2ListResponseItemsTokenEndpointAuthMethod string

const (
	Oauth2ListResponseItemsTokenEndpointAuthMethodClientSecretBasic Oauth2ListResponseItemsTokenEndpointAuthMethod = "client_secret_basic"
	Oauth2ListResponseItemsTokenEndpointAuthMethodClientSecretPost  Oauth2ListResponseItemsTokenEndpointAuthMethod = "client_secret_post"
	Oauth2ListResponseItemsTokenEndpointAuthMethodNone              Oauth2ListResponseItemsTokenEndpointAuthMethod = "none"
)

func (r Oauth2ListResponseItemsTokenEndpointAuthMethod) IsKnown() bool {
	switch r {
	case Oauth2ListResponseItemsTokenEndpointAuthMethodClientSecretBasic, Oauth2ListResponseItemsTokenEndpointAuthMethodClientSecretPost, Oauth2ListResponseItemsTokenEndpointAuthMethodNone:
		return true
	}
	return false
}

type Oauth2AuthorizeResponse struct {
	// This field can have the runtime type of [AuthorizeResponseUserClient],
	// [AuthorizeResponseOrganizationClient].
	Client  interface{}                    `json:"client"`
	SubType Oauth2AuthorizeResponseSubType `json:"sub_type,required"`
	// This field can have the runtime type of [AuthorizeResponseUserSub],
	// [AuthorizeResponseOrganizationSub].
	Sub interface{} `json:"sub"`
	// This field can have the runtime type of [[]AuthorizeResponseUserScope],
	// [[]AuthorizeResponseOrganizationScope].
	Scopes interface{} `json:"scopes"`
	// This field can have the runtime type of
	// [[]AuthorizeResponseOrganizationOrganization].
	Organizations interface{}                 `json:"organizations,required"`
	JSON          oauth2AuthorizeResponseJSON `json:"-"`
	union         Oauth2AuthorizeResponseUnion
}

// oauth2AuthorizeResponseJSON contains the JSON metadata for the struct
// [Oauth2AuthorizeResponse]
type oauth2AuthorizeResponseJSON struct {
	Client        apijson.Field
	SubType       apijson.Field
	Sub           apijson.Field
	Scopes        apijson.Field
	Organizations apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r oauth2AuthorizeResponseJSON) RawJSON() string {
	return r.raw
}

func (r *Oauth2AuthorizeResponse) UnmarshalJSON(data []byte) (err error) {
	*r = Oauth2AuthorizeResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [Oauth2AuthorizeResponseUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [AuthorizeResponseUser],
// [AuthorizeResponseOrganization].
func (r Oauth2AuthorizeResponse) AsUnion() Oauth2AuthorizeResponseUnion {
	return r.union
}

// Union satisfied by [AuthorizeResponseUser] or [AuthorizeResponseOrganization].
type Oauth2AuthorizeResponseUnion interface {
	implementsOauth2AuthorizeResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*Oauth2AuthorizeResponseUnion)(nil)).Elem(),
		"sub_type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AuthorizeResponseUser{}),
			DiscriminatorValue: "user",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AuthorizeResponseOrganization{}),
			DiscriminatorValue: "organization",
		},
	)
}

type Oauth2AuthorizeResponseSubType string

const (
	Oauth2AuthorizeResponseSubTypeUser         Oauth2AuthorizeResponseSubType = "user"
	Oauth2AuthorizeResponseSubTypeOrganization Oauth2AuthorizeResponseSubType = "organization"
)

func (r Oauth2AuthorizeResponseSubType) IsKnown() bool {
	switch r {
	case Oauth2AuthorizeResponseSubTypeUser, Oauth2AuthorizeResponseSubTypeOrganization:
		return true
	}
	return false
}

type Oauth2ListParams struct {
	// Size of a page, defaults to 10. Maximum is 100.
	Limit param.Field[int64] `query:"limit"`
	// Page number, defaults to 1.
	Page param.Field[int64] `query:"page"`
}

// URLQuery serializes [Oauth2ListParams]'s query parameters as `url.Values`.
func (r Oauth2ListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type Oauth2IntrospectParams struct {
	Token         param.Field[string]                              `json:"token,required"`
	ClientID      param.Field[string]                              `json:"client_id,required"`
	ClientSecret  param.Field[string]                              `json:"client_secret,required"`
	TokenTypeHint param.Field[Oauth2IntrospectParamsTokenTypeHint] `json:"token_type_hint"`
}

func (r Oauth2IntrospectParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Oauth2IntrospectParamsTokenTypeHint string

const (
	Oauth2IntrospectParamsTokenTypeHintAccessToken  Oauth2IntrospectParamsTokenTypeHint = "access_token"
	Oauth2IntrospectParamsTokenTypeHintRefreshToken Oauth2IntrospectParamsTokenTypeHint = "refresh_token"
)

func (r Oauth2IntrospectParamsTokenTypeHint) IsKnown() bool {
	switch r {
	case Oauth2IntrospectParamsTokenTypeHintAccessToken, Oauth2IntrospectParamsTokenTypeHintRefreshToken:
		return true
	}
	return false
}

type Oauth2RevokeParams struct {
	Token         param.Field[string]                          `json:"token,required"`
	ClientID      param.Field[string]                          `json:"client_id,required"`
	ClientSecret  param.Field[string]                          `json:"client_secret,required"`
	TokenTypeHint param.Field[Oauth2RevokeParamsTokenTypeHint] `json:"token_type_hint"`
}

func (r Oauth2RevokeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Oauth2RevokeParamsTokenTypeHint string

const (
	Oauth2RevokeParamsTokenTypeHintAccessToken  Oauth2RevokeParamsTokenTypeHint = "access_token"
	Oauth2RevokeParamsTokenTypeHintRefreshToken Oauth2RevokeParamsTokenTypeHint = "refresh_token"
)

func (r Oauth2RevokeParamsTokenTypeHint) IsKnown() bool {
	switch r {
	case Oauth2RevokeParamsTokenTypeHintAccessToken, Oauth2RevokeParamsTokenTypeHintRefreshToken:
		return true
	}
	return false
}

type Oauth2TokenParams struct {
	Body Oauth2TokenParamsBodyUnion `json:"body,required"`
}

func (r Oauth2TokenParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// Satisfied by [Oauth2TokenParamsBodyUnknown], [Oauth2TokenParamsBodyUnknown].
type Oauth2TokenParamsBodyUnion interface {
	implementsOauth2TokenParamsBodyUnion()
}
