// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type GetTeamAccessRequestSecurity struct {
	BearerToken string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

type GetTeamAccessRequestRequest struct {
	TeamID string `pathParam:"style=simple,explode=false,name=teamId"`
	UserID string `pathParam:"style=simple,explode=false,name=userId"`
}

// GetTeamAccessRequest200ApplicationJSONBitbucket - Map of the connected Bitbucket account.
type GetTeamAccessRequest200ApplicationJSONBitbucket struct {
	Login *string `json:"login,omitempty"`
}

// GetTeamAccessRequest200ApplicationJSONGithub - Map of the connected GitHub account.
type GetTeamAccessRequest200ApplicationJSONGithub struct {
	Login *string `json:"login,omitempty"`
}

// GetTeamAccessRequest200ApplicationJSONGitlab - Map of the connected GitLab account.
type GetTeamAccessRequest200ApplicationJSONGitlab struct {
	Login *string `json:"login,omitempty"`
}

type GetTeamAccessRequest200ApplicationJSONJoinedFromGitUserIDType string

const (
	GetTeamAccessRequest200ApplicationJSONJoinedFromGitUserIDTypeStr     GetTeamAccessRequest200ApplicationJSONJoinedFromGitUserIDType = "str"
	GetTeamAccessRequest200ApplicationJSONJoinedFromGitUserIDTypeInteger GetTeamAccessRequest200ApplicationJSONJoinedFromGitUserIDType = "integer"
)

type GetTeamAccessRequest200ApplicationJSONJoinedFromGitUserID struct {
	Str     *string
	Integer *int64

	Type GetTeamAccessRequest200ApplicationJSONJoinedFromGitUserIDType
}

func CreateGetTeamAccessRequest200ApplicationJSONJoinedFromGitUserIDStr(str string) GetTeamAccessRequest200ApplicationJSONJoinedFromGitUserID {
	typ := GetTeamAccessRequest200ApplicationJSONJoinedFromGitUserIDTypeStr

	return GetTeamAccessRequest200ApplicationJSONJoinedFromGitUserID{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetTeamAccessRequest200ApplicationJSONJoinedFromGitUserIDInteger(integer int64) GetTeamAccessRequest200ApplicationJSONJoinedFromGitUserID {
	typ := GetTeamAccessRequest200ApplicationJSONJoinedFromGitUserIDTypeInteger

	return GetTeamAccessRequest200ApplicationJSONJoinedFromGitUserID{
		Integer: &integer,
		Type:    typ,
	}
}

func (u *GetTeamAccessRequest200ApplicationJSONJoinedFromGitUserID) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	str := new(string)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&str); err == nil {
		u.Str = str
		u.Type = GetTeamAccessRequest200ApplicationJSONJoinedFromGitUserIDTypeStr
		return nil
	}

	integer := new(int64)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&integer); err == nil {
		u.Integer = integer
		u.Type = GetTeamAccessRequest200ApplicationJSONJoinedFromGitUserIDTypeInteger
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetTeamAccessRequest200ApplicationJSONJoinedFromGitUserID) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return json.Marshal(u.Str)
	}

	if u.Integer != nil {
		return json.Marshal(u.Integer)
	}

	return nil, nil
}

type GetTeamAccessRequest200ApplicationJSONJoinedFromOrigin string

const (
	GetTeamAccessRequest200ApplicationJSONJoinedFromOriginLink              GetTeamAccessRequest200ApplicationJSONJoinedFromOrigin = "link"
	GetTeamAccessRequest200ApplicationJSONJoinedFromOriginMail              GetTeamAccessRequest200ApplicationJSONJoinedFromOrigin = "mail"
	GetTeamAccessRequest200ApplicationJSONJoinedFromOriginImport            GetTeamAccessRequest200ApplicationJSONJoinedFromOrigin = "import"
	GetTeamAccessRequest200ApplicationJSONJoinedFromOriginTeams             GetTeamAccessRequest200ApplicationJSONJoinedFromOrigin = "teams"
	GetTeamAccessRequest200ApplicationJSONJoinedFromOriginGithub            GetTeamAccessRequest200ApplicationJSONJoinedFromOrigin = "github"
	GetTeamAccessRequest200ApplicationJSONJoinedFromOriginGitlab            GetTeamAccessRequest200ApplicationJSONJoinedFromOrigin = "gitlab"
	GetTeamAccessRequest200ApplicationJSONJoinedFromOriginBitbucket         GetTeamAccessRequest200ApplicationJSONJoinedFromOrigin = "bitbucket"
	GetTeamAccessRequest200ApplicationJSONJoinedFromOriginSaml              GetTeamAccessRequest200ApplicationJSONJoinedFromOrigin = "saml"
	GetTeamAccessRequest200ApplicationJSONJoinedFromOriginDsync             GetTeamAccessRequest200ApplicationJSONJoinedFromOrigin = "dsync"
	GetTeamAccessRequest200ApplicationJSONJoinedFromOriginFeedback          GetTeamAccessRequest200ApplicationJSONJoinedFromOrigin = "feedback"
	GetTeamAccessRequest200ApplicationJSONJoinedFromOriginOrganizationTeams GetTeamAccessRequest200ApplicationJSONJoinedFromOrigin = "organization-teams"
)

func (e GetTeamAccessRequest200ApplicationJSONJoinedFromOrigin) ToPointer() *GetTeamAccessRequest200ApplicationJSONJoinedFromOrigin {
	return &e
}

func (e *GetTeamAccessRequest200ApplicationJSONJoinedFromOrigin) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "link":
		fallthrough
	case "mail":
		fallthrough
	case "import":
		fallthrough
	case "teams":
		fallthrough
	case "github":
		fallthrough
	case "gitlab":
		fallthrough
	case "bitbucket":
		fallthrough
	case "saml":
		fallthrough
	case "dsync":
		fallthrough
	case "feedback":
		fallthrough
	case "organization-teams":
		*e = GetTeamAccessRequest200ApplicationJSONJoinedFromOrigin(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetTeamAccessRequest200ApplicationJSONJoinedFromOrigin: %v", v)
	}
}

// GetTeamAccessRequest200ApplicationJSONJoinedFrom - A map that describes the origin from where the user joined.
type GetTeamAccessRequest200ApplicationJSONJoinedFrom struct {
	CommitID         *string                                                    `json:"commitId,omitempty"`
	DsyncConnectedAt *int64                                                     `json:"dsyncConnectedAt,omitempty"`
	DsyncUserID      *string                                                    `json:"dsyncUserId,omitempty"`
	GitUserID        *GetTeamAccessRequest200ApplicationJSONJoinedFromGitUserID `json:"gitUserId,omitempty"`
	GitUserLogin     *string                                                    `json:"gitUserLogin,omitempty"`
	IdpUserID        *string                                                    `json:"idpUserId,omitempty"`
	Origin           GetTeamAccessRequest200ApplicationJSONJoinedFromOrigin     `json:"origin"`
	RepoID           *string                                                    `json:"repoId,omitempty"`
	RepoPath         *string                                                    `json:"repoPath,omitempty"`
	SsoConnectedAt   *int64                                                     `json:"ssoConnectedAt,omitempty"`
	SsoUserID        *string                                                    `json:"ssoUserId,omitempty"`
}

// GetTeamAccessRequest200ApplicationJSON - Successfully
type GetTeamAccessRequest200ApplicationJSON struct {
	// Timestamp in milliseconds when the user requested access to the team.
	AccessRequestedAt int64 `json:"accessRequestedAt"`
	// Map of the connected Bitbucket account.
	Bitbucket *GetTeamAccessRequest200ApplicationJSONBitbucket `json:"bitbucket"`
	// Current status of the membership. Will be `true` if confirmed, if pending it'll be `false`.
	Confirmed bool `json:"confirmed"`
	// Map of the connected GitHub account.
	Github *GetTeamAccessRequest200ApplicationJSONGithub `json:"github"`
	// Map of the connected GitLab account.
	Gitlab *GetTeamAccessRequest200ApplicationJSONGitlab `json:"gitlab"`
	// A map that describes the origin from where the user joined.
	JoinedFrom GetTeamAccessRequest200ApplicationJSONJoinedFrom `json:"joinedFrom"`
	// The name of the team.
	TeamName string `json:"teamName"`
	// The slug of the team.
	TeamSlug string `json:"teamSlug"`
}

type GetTeamAccessRequestResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successfully
	GetTeamAccessRequest200ApplicationJSONObject *GetTeamAccessRequest200ApplicationJSON
}
