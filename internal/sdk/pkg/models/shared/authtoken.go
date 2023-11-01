// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"vercel/internal/sdk/pkg/utils"
)

type AuthTokenScopes2Origin string

const (
	AuthTokenScopes2OriginSaml      AuthTokenScopes2Origin = "saml"
	AuthTokenScopes2OriginGithub    AuthTokenScopes2Origin = "github"
	AuthTokenScopes2OriginGitlab    AuthTokenScopes2Origin = "gitlab"
	AuthTokenScopes2OriginBitbucket AuthTokenScopes2Origin = "bitbucket"
	AuthTokenScopes2OriginEmail     AuthTokenScopes2Origin = "email"
	AuthTokenScopes2OriginManual    AuthTokenScopes2Origin = "manual"
)

func (e AuthTokenScopes2Origin) ToPointer() *AuthTokenScopes2Origin {
	return &e
}

func (e *AuthTokenScopes2Origin) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "saml":
		fallthrough
	case "github":
		fallthrough
	case "gitlab":
		fallthrough
	case "bitbucket":
		fallthrough
	case "email":
		fallthrough
	case "manual":
		*e = AuthTokenScopes2Origin(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AuthTokenScopes2Origin: %v", v)
	}
}

type AuthTokenScopes2Type string

const (
	AuthTokenScopes2TypeTeam AuthTokenScopes2Type = "team"
)

func (e AuthTokenScopes2Type) ToPointer() *AuthTokenScopes2Type {
	return &e
}

func (e *AuthTokenScopes2Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "team":
		*e = AuthTokenScopes2Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AuthTokenScopes2Type: %v", v)
	}
}

// AuthTokenScopes2 - The access scopes granted to the token.
type AuthTokenScopes2 struct {
	CreatedAt int64                  `json:"createdAt"`
	ExpiresAt *int64                 `json:"expiresAt,omitempty"`
	Origin    AuthTokenScopes2Origin `json:"origin"`
	TeamID    string                 `json:"teamId"`
	Type      AuthTokenScopes2Type   `json:"type"`
}

func (o *AuthTokenScopes2) GetCreatedAt() int64 {
	if o == nil {
		return 0
	}
	return o.CreatedAt
}

func (o *AuthTokenScopes2) GetExpiresAt() *int64 {
	if o == nil {
		return nil
	}
	return o.ExpiresAt
}

func (o *AuthTokenScopes2) GetOrigin() AuthTokenScopes2Origin {
	if o == nil {
		return AuthTokenScopes2Origin("")
	}
	return o.Origin
}

func (o *AuthTokenScopes2) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

func (o *AuthTokenScopes2) GetType() AuthTokenScopes2Type {
	if o == nil {
		return AuthTokenScopes2Type("")
	}
	return o.Type
}

type AuthTokenScopes1Origin string

const (
	AuthTokenScopes1OriginSaml      AuthTokenScopes1Origin = "saml"
	AuthTokenScopes1OriginGithub    AuthTokenScopes1Origin = "github"
	AuthTokenScopes1OriginGitlab    AuthTokenScopes1Origin = "gitlab"
	AuthTokenScopes1OriginBitbucket AuthTokenScopes1Origin = "bitbucket"
	AuthTokenScopes1OriginEmail     AuthTokenScopes1Origin = "email"
	AuthTokenScopes1OriginManual    AuthTokenScopes1Origin = "manual"
)

func (e AuthTokenScopes1Origin) ToPointer() *AuthTokenScopes1Origin {
	return &e
}

func (e *AuthTokenScopes1Origin) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "saml":
		fallthrough
	case "github":
		fallthrough
	case "gitlab":
		fallthrough
	case "bitbucket":
		fallthrough
	case "email":
		fallthrough
	case "manual":
		*e = AuthTokenScopes1Origin(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AuthTokenScopes1Origin: %v", v)
	}
}

type AuthTokenScopes1Type string

const (
	AuthTokenScopes1TypeUser AuthTokenScopes1Type = "user"
)

func (e AuthTokenScopes1Type) ToPointer() *AuthTokenScopes1Type {
	return &e
}

func (e *AuthTokenScopes1Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "user":
		*e = AuthTokenScopes1Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AuthTokenScopes1Type: %v", v)
	}
}

// AuthTokenScopes1 - The access scopes granted to the token.
type AuthTokenScopes1 struct {
	CreatedAt int64                  `json:"createdAt"`
	ExpiresAt *int64                 `json:"expiresAt,omitempty"`
	Origin    AuthTokenScopes1Origin `json:"origin"`
	Type      AuthTokenScopes1Type   `json:"type"`
}

func (o *AuthTokenScopes1) GetCreatedAt() int64 {
	if o == nil {
		return 0
	}
	return o.CreatedAt
}

func (o *AuthTokenScopes1) GetExpiresAt() *int64 {
	if o == nil {
		return nil
	}
	return o.ExpiresAt
}

func (o *AuthTokenScopes1) GetOrigin() AuthTokenScopes1Origin {
	if o == nil {
		return AuthTokenScopes1Origin("")
	}
	return o.Origin
}

func (o *AuthTokenScopes1) GetType() AuthTokenScopes1Type {
	if o == nil {
		return AuthTokenScopes1Type("")
	}
	return o.Type
}

type AuthTokenScopesType string

const (
	AuthTokenScopesTypeAuthTokenScopes1 AuthTokenScopesType = "AuthToken_scopes_1"
	AuthTokenScopesTypeAuthTokenScopes2 AuthTokenScopesType = "AuthToken_scopes_2"
)

type AuthTokenScopes struct {
	AuthTokenScopes1 *AuthTokenScopes1
	AuthTokenScopes2 *AuthTokenScopes2

	Type AuthTokenScopesType
}

func CreateAuthTokenScopesAuthTokenScopes1(authTokenScopes1 AuthTokenScopes1) AuthTokenScopes {
	typ := AuthTokenScopesTypeAuthTokenScopes1

	return AuthTokenScopes{
		AuthTokenScopes1: &authTokenScopes1,
		Type:             typ,
	}
}

func CreateAuthTokenScopesAuthTokenScopes2(authTokenScopes2 AuthTokenScopes2) AuthTokenScopes {
	typ := AuthTokenScopesTypeAuthTokenScopes2

	return AuthTokenScopes{
		AuthTokenScopes2: &authTokenScopes2,
		Type:             typ,
	}
}

func (u *AuthTokenScopes) UnmarshalJSON(data []byte) error {

	authTokenScopes1 := new(AuthTokenScopes1)
	if err := utils.UnmarshalJSON(data, &authTokenScopes1, "", true, true); err == nil {
		u.AuthTokenScopes1 = authTokenScopes1
		u.Type = AuthTokenScopesTypeAuthTokenScopes1
		return nil
	}

	authTokenScopes2 := new(AuthTokenScopes2)
	if err := utils.UnmarshalJSON(data, &authTokenScopes2, "", true, true); err == nil {
		u.AuthTokenScopes2 = authTokenScopes2
		u.Type = AuthTokenScopesTypeAuthTokenScopes2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u AuthTokenScopes) MarshalJSON() ([]byte, error) {
	if u.AuthTokenScopes1 != nil {
		return utils.MarshalJSON(u.AuthTokenScopes1, "", true)
	}

	if u.AuthTokenScopes2 != nil {
		return utils.MarshalJSON(u.AuthTokenScopes2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// AuthToken - Authentication token metadata.
type AuthToken struct {
	// Timestamp (in milliseconds) of when the token was most recently used.
	ActiveAt int64 `json:"activeAt"`
	// Timestamp (in milliseconds) of when the token was created.
	CreatedAt int64 `json:"createdAt"`
	// Timestamp (in milliseconds) of when the token expires.
	ExpiresAt *int64 `json:"expiresAt,omitempty"`
	// The unique identifier of the token.
	ID string `json:"id"`
	// The human-readable name of the token.
	Name string `json:"name"`
	// The origin of how the token was created.
	Origin *string `json:"origin,omitempty"`
	// The access scopes granted to the token.
	Scopes []AuthTokenScopes `json:"scopes,omitempty"`
	// The type of the token.
	Type string `json:"type"`
}

func (o *AuthToken) GetActiveAt() int64 {
	if o == nil {
		return 0
	}
	return o.ActiveAt
}

func (o *AuthToken) GetCreatedAt() int64 {
	if o == nil {
		return 0
	}
	return o.CreatedAt
}

func (o *AuthToken) GetExpiresAt() *int64 {
	if o == nil {
		return nil
	}
	return o.ExpiresAt
}

func (o *AuthToken) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *AuthToken) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *AuthToken) GetOrigin() *string {
	if o == nil {
		return nil
	}
	return o.Origin
}

func (o *AuthToken) GetScopes() []AuthTokenScopes {
	if o == nil {
		return nil
	}
	return o.Scopes
}

func (o *AuthToken) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}
