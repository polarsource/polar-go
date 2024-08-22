// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package polar

import (
	"context"
	"net/http"
	"reflect"

	"github.com/polarsource/polar-go/internal/apijson"
	"github.com/polarsource/polar-go/internal/requestconfig"
	"github.com/polarsource/polar-go/option"
	"github.com/tidwall/gjson"
)

// Oauth2UserinfoService contains methods and other services that help with
// interacting with the polar API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOauth2UserinfoService] method instead.
type Oauth2UserinfoService struct {
	Options []option.RequestOption
}

// NewOauth2UserinfoService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOauth2UserinfoService(opts ...option.RequestOption) (r *Oauth2UserinfoService) {
	r = &Oauth2UserinfoService{}
	r.Options = opts
	return
}

// Get information about the authenticated user.
func (r *Oauth2UserinfoService) New(ctx context.Context, opts ...option.RequestOption) (res *Oauth2UserinfoNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/oauth2/userinfo"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Get information about the authenticated user.
func (r *Oauth2UserinfoService) Get(ctx context.Context, opts ...option.RequestOption) (res *Oauth2UserinfoGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/oauth2/userinfo"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type UserInfoOrganization struct {
	Sub  string                   `json:"sub,required"`
	Name string                   `json:"name,nullable"`
	JSON userInfoOrganizationJSON `json:"-"`
}

// userInfoOrganizationJSON contains the JSON metadata for the struct
// [UserInfoOrganization]
type userInfoOrganizationJSON struct {
	Sub         apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserInfoOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userInfoOrganizationJSON) RawJSON() string {
	return r.raw
}

func (r UserInfoOrganization) implementsOauth2UserinfoNewResponse() {}

func (r UserInfoOrganization) implementsOauth2UserinfoGetResponse() {}

type UserInfoUser struct {
	Sub           string           `json:"sub,required"`
	Email         string           `json:"email,nullable"`
	EmailVerified bool             `json:"email_verified,nullable"`
	Name          string           `json:"name,nullable"`
	JSON          userInfoUserJSON `json:"-"`
}

// userInfoUserJSON contains the JSON metadata for the struct [UserInfoUser]
type userInfoUserJSON struct {
	Sub           apijson.Field
	Email         apijson.Field
	EmailVerified apijson.Field
	Name          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *UserInfoUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userInfoUserJSON) RawJSON() string {
	return r.raw
}

func (r UserInfoUser) implementsOauth2UserinfoNewResponse() {}

func (r UserInfoUser) implementsOauth2UserinfoGetResponse() {}

type Oauth2UserinfoNewResponse struct {
	Sub           string                        `json:"sub,required"`
	Name          string                        `json:"name,nullable"`
	Email         string                        `json:"email,nullable"`
	EmailVerified bool                          `json:"email_verified,nullable"`
	JSON          oauth2UserinfoNewResponseJSON `json:"-"`
	union         Oauth2UserinfoNewResponseUnion
}

// oauth2UserinfoNewResponseJSON contains the JSON metadata for the struct
// [Oauth2UserinfoNewResponse]
type oauth2UserinfoNewResponseJSON struct {
	Sub           apijson.Field
	Name          apijson.Field
	Email         apijson.Field
	EmailVerified apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r oauth2UserinfoNewResponseJSON) RawJSON() string {
	return r.raw
}

func (r *Oauth2UserinfoNewResponse) UnmarshalJSON(data []byte) (err error) {
	*r = Oauth2UserinfoNewResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [Oauth2UserinfoNewResponseUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are [UserInfoUser], [UserInfoOrganization].
func (r Oauth2UserinfoNewResponse) AsUnion() Oauth2UserinfoNewResponseUnion {
	return r.union
}

// Union satisfied by [UserInfoUser] or [UserInfoOrganization].
type Oauth2UserinfoNewResponseUnion interface {
	implementsOauth2UserinfoNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*Oauth2UserinfoNewResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserInfoUser{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserInfoOrganization{}),
		},
	)
}

type Oauth2UserinfoGetResponse struct {
	Sub           string                        `json:"sub,required"`
	Name          string                        `json:"name,nullable"`
	Email         string                        `json:"email,nullable"`
	EmailVerified bool                          `json:"email_verified,nullable"`
	JSON          oauth2UserinfoGetResponseJSON `json:"-"`
	union         Oauth2UserinfoGetResponseUnion
}

// oauth2UserinfoGetResponseJSON contains the JSON metadata for the struct
// [Oauth2UserinfoGetResponse]
type oauth2UserinfoGetResponseJSON struct {
	Sub           apijson.Field
	Name          apijson.Field
	Email         apijson.Field
	EmailVerified apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r oauth2UserinfoGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *Oauth2UserinfoGetResponse) UnmarshalJSON(data []byte) (err error) {
	*r = Oauth2UserinfoGetResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [Oauth2UserinfoGetResponseUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are [UserInfoUser], [UserInfoOrganization].
func (r Oauth2UserinfoGetResponse) AsUnion() Oauth2UserinfoGetResponseUnion {
	return r.union
}

// Union satisfied by [UserInfoUser] or [UserInfoOrganization].
type Oauth2UserinfoGetResponseUnion interface {
	implementsOauth2UserinfoGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*Oauth2UserinfoGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserInfoUser{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserInfoOrganization{}),
		},
	)
}
