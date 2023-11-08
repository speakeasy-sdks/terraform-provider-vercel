// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zchee/terraform-provider-vercel/internal/sdk/pkg/utils"
)

type AuthTokenOrigin string

const (
	AuthTokenOriginSaml      AuthTokenOrigin = "saml"
	AuthTokenOriginGithub    AuthTokenOrigin = "github"
	AuthTokenOriginGitlab    AuthTokenOrigin = "gitlab"
	AuthTokenOriginBitbucket AuthTokenOrigin = "bitbucket"
	AuthTokenOriginEmail     AuthTokenOrigin = "email"
	AuthTokenOriginManual    AuthTokenOrigin = "manual"
)

func (e AuthTokenOrigin) ToPointer() *AuthTokenOrigin {
	return &e
}

func (e *AuthTokenOrigin) UnmarshalJSON(data []byte) error {
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
		*e = AuthTokenOrigin(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AuthTokenOrigin: %v", v)
	}
}

type AuthTokenType string

const (
	AuthTokenTypeTeam AuthTokenType = "team"
)

func (e AuthTokenType) ToPointer() *AuthTokenType {
	return &e
}

func (e *AuthTokenType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "team":
		*e = AuthTokenType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AuthTokenType: %v", v)
	}
}

// Two - The access scopes granted to the token.
type Two struct {
	CreatedAt int64           `json:"createdAt"`
	ExpiresAt *int64          `json:"expiresAt,omitempty"`
	Origin    AuthTokenOrigin `json:"origin"`
	TeamID    string          `json:"teamId"`
	Type      AuthTokenType   `json:"type"`
}

func (o *Two) GetCreatedAt() int64 {
	if o == nil {
		return 0
	}
	return o.CreatedAt
}

func (o *Two) GetExpiresAt() *int64 {
	if o == nil {
		return nil
	}
	return o.ExpiresAt
}

func (o *Two) GetOrigin() AuthTokenOrigin {
	if o == nil {
		return AuthTokenOrigin("")
	}
	return o.Origin
}

func (o *Two) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

func (o *Two) GetType() AuthTokenType {
	if o == nil {
		return AuthTokenType("")
	}
	return o.Type
}

type Origin string

const (
	OriginSaml      Origin = "saml"
	OriginGithub    Origin = "github"
	OriginGitlab    Origin = "gitlab"
	OriginBitbucket Origin = "bitbucket"
	OriginEmail     Origin = "email"
	OriginManual    Origin = "manual"
)

func (e Origin) ToPointer() *Origin {
	return &e
}

func (e *Origin) UnmarshalJSON(data []byte) error {
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
		*e = Origin(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Origin: %v", v)
	}
}

type AuthTokenSchemasType string

const (
	AuthTokenSchemasTypeUser AuthTokenSchemasType = "user"
)

func (e AuthTokenSchemasType) ToPointer() *AuthTokenSchemasType {
	return &e
}

func (e *AuthTokenSchemasType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "user":
		*e = AuthTokenSchemasType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AuthTokenSchemasType: %v", v)
	}
}

// One - The access scopes granted to the token.
type One struct {
	CreatedAt int64                `json:"createdAt"`
	ExpiresAt *int64               `json:"expiresAt,omitempty"`
	Origin    Origin               `json:"origin"`
	Type      AuthTokenSchemasType `json:"type"`
}

func (o *One) GetCreatedAt() int64 {
	if o == nil {
		return 0
	}
	return o.CreatedAt
}

func (o *One) GetExpiresAt() *int64 {
	if o == nil {
		return nil
	}
	return o.ExpiresAt
}

func (o *One) GetOrigin() Origin {
	if o == nil {
		return Origin("")
	}
	return o.Origin
}

func (o *One) GetType() AuthTokenSchemasType {
	if o == nil {
		return AuthTokenSchemasType("")
	}
	return o.Type
}

type ScopesType string

const (
	ScopesTypeOne ScopesType = "1"
	ScopesTypeTwo ScopesType = "2"
)

type Scopes struct {
	One *One
	Two *Two

	Type ScopesType
}

func CreateScopesOne(one One) Scopes {
	typ := ScopesTypeOne

	return Scopes{
		One:  &one,
		Type: typ,
	}
}

func CreateScopesTwo(two Two) Scopes {
	typ := ScopesTypeTwo

	return Scopes{
		Two:  &two,
		Type: typ,
	}
}

func (u *Scopes) UnmarshalJSON(data []byte) error {

	one := new(One)
	if err := utils.UnmarshalJSON(data, &one, "", true, true); err == nil {
		u.One = one
		u.Type = ScopesTypeOne
		return nil
	}

	two := new(Two)
	if err := utils.UnmarshalJSON(data, &two, "", true, true); err == nil {
		u.Two = two
		u.Type = ScopesTypeTwo
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u Scopes) MarshalJSON() ([]byte, error) {
	if u.One != nil {
		return utils.MarshalJSON(u.One, "", true)
	}

	if u.Two != nil {
		return utils.MarshalJSON(u.Two, "", true)
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
	Scopes []Scopes `json:"scopes,omitempty"`
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

func (o *AuthToken) GetScopes() []Scopes {
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
