// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"vercel/internal/sdk/pkg/utils"
)

type RequestAccessToTeamRequestBodyJoinedFromGitUserIDType string

const (
	RequestAccessToTeamRequestBodyJoinedFromGitUserIDTypeStr     RequestAccessToTeamRequestBodyJoinedFromGitUserIDType = "str"
	RequestAccessToTeamRequestBodyJoinedFromGitUserIDTypeInteger RequestAccessToTeamRequestBodyJoinedFromGitUserIDType = "integer"
)

type RequestAccessToTeamRequestBodyJoinedFromGitUserID struct {
	Str     *string
	Integer *int64

	Type RequestAccessToTeamRequestBodyJoinedFromGitUserIDType
}

func CreateRequestAccessToTeamRequestBodyJoinedFromGitUserIDStr(str string) RequestAccessToTeamRequestBodyJoinedFromGitUserID {
	typ := RequestAccessToTeamRequestBodyJoinedFromGitUserIDTypeStr

	return RequestAccessToTeamRequestBodyJoinedFromGitUserID{
		Str:  &str,
		Type: typ,
	}
}

func CreateRequestAccessToTeamRequestBodyJoinedFromGitUserIDInteger(integer int64) RequestAccessToTeamRequestBodyJoinedFromGitUserID {
	typ := RequestAccessToTeamRequestBodyJoinedFromGitUserIDTypeInteger

	return RequestAccessToTeamRequestBodyJoinedFromGitUserID{
		Integer: &integer,
		Type:    typ,
	}
}

func (u *RequestAccessToTeamRequestBodyJoinedFromGitUserID) UnmarshalJSON(data []byte) error {

	str := new(string)
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = str
		u.Type = RequestAccessToTeamRequestBodyJoinedFromGitUserIDTypeStr
		return nil
	}

	integer := new(int64)
	if err := utils.UnmarshalJSON(data, &integer, "", true, true); err == nil {
		u.Integer = integer
		u.Type = RequestAccessToTeamRequestBodyJoinedFromGitUserIDTypeInteger
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u RequestAccessToTeamRequestBodyJoinedFromGitUserID) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.Integer != nil {
		return utils.MarshalJSON(u.Integer, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// RequestAccessToTeamRequestBodyJoinedFromOrigin - The origin of the request.
type RequestAccessToTeamRequestBodyJoinedFromOrigin string

const (
	RequestAccessToTeamRequestBodyJoinedFromOriginImport            RequestAccessToTeamRequestBodyJoinedFromOrigin = "import"
	RequestAccessToTeamRequestBodyJoinedFromOriginTeams             RequestAccessToTeamRequestBodyJoinedFromOrigin = "teams"
	RequestAccessToTeamRequestBodyJoinedFromOriginGithub            RequestAccessToTeamRequestBodyJoinedFromOrigin = "github"
	RequestAccessToTeamRequestBodyJoinedFromOriginGitlab            RequestAccessToTeamRequestBodyJoinedFromOrigin = "gitlab"
	RequestAccessToTeamRequestBodyJoinedFromOriginBitbucket         RequestAccessToTeamRequestBodyJoinedFromOrigin = "bitbucket"
	RequestAccessToTeamRequestBodyJoinedFromOriginFeedback          RequestAccessToTeamRequestBodyJoinedFromOrigin = "feedback"
	RequestAccessToTeamRequestBodyJoinedFromOriginOrganizationTeams RequestAccessToTeamRequestBodyJoinedFromOrigin = "organization-teams"
)

func (e RequestAccessToTeamRequestBodyJoinedFromOrigin) ToPointer() *RequestAccessToTeamRequestBodyJoinedFromOrigin {
	return &e
}

func (e *RequestAccessToTeamRequestBodyJoinedFromOrigin) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
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
	case "feedback":
		fallthrough
	case "organization-teams":
		*e = RequestAccessToTeamRequestBodyJoinedFromOrigin(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RequestAccessToTeamRequestBodyJoinedFromOrigin: %v", v)
	}
}

type RequestAccessToTeamRequestBodyJoinedFrom struct {
	// The commit sha if the origin is a git provider.
	CommitID *string `json:"commitId,omitempty"`
	// The ID of the Git account of the user who requests access.
	GitUserID *RequestAccessToTeamRequestBodyJoinedFromGitUserID `json:"gitUserId,omitempty"`
	// The login name for the Git account of the user who requests access.
	GitUserLogin *string `json:"gitUserLogin,omitempty"`
	// The origin of the request.
	Origin RequestAccessToTeamRequestBodyJoinedFromOrigin `json:"origin"`
	// The ID of the repository for the given Git provider.
	RepoID *string `json:"repoId,omitempty"`
	// The path to the repository for the given Git provider.
	RepoPath *string `json:"repoPath,omitempty"`
	UserID   *string `json:"userId,omitempty"`
}

func (o *RequestAccessToTeamRequestBodyJoinedFrom) GetCommitID() *string {
	if o == nil {
		return nil
	}
	return o.CommitID
}

func (o *RequestAccessToTeamRequestBodyJoinedFrom) GetGitUserID() *RequestAccessToTeamRequestBodyJoinedFromGitUserID {
	if o == nil {
		return nil
	}
	return o.GitUserID
}

func (o *RequestAccessToTeamRequestBodyJoinedFrom) GetGitUserLogin() *string {
	if o == nil {
		return nil
	}
	return o.GitUserLogin
}

func (o *RequestAccessToTeamRequestBodyJoinedFrom) GetOrigin() RequestAccessToTeamRequestBodyJoinedFromOrigin {
	if o == nil {
		return RequestAccessToTeamRequestBodyJoinedFromOrigin("")
	}
	return o.Origin
}

func (o *RequestAccessToTeamRequestBodyJoinedFrom) GetRepoID() *string {
	if o == nil {
		return nil
	}
	return o.RepoID
}

func (o *RequestAccessToTeamRequestBodyJoinedFrom) GetRepoPath() *string {
	if o == nil {
		return nil
	}
	return o.RepoPath
}

func (o *RequestAccessToTeamRequestBodyJoinedFrom) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}

type RequestAccessToTeamRequestBody struct {
	JoinedFrom RequestAccessToTeamRequestBodyJoinedFrom `json:"joinedFrom"`
}

func (o *RequestAccessToTeamRequestBody) GetJoinedFrom() RequestAccessToTeamRequestBodyJoinedFrom {
	if o == nil {
		return RequestAccessToTeamRequestBodyJoinedFrom{}
	}
	return o.JoinedFrom
}

type RequestAccessToTeamRequest struct {
	RequestBody *RequestAccessToTeamRequestBody `request:"mediaType=application/json"`
	TeamID      string                          `pathParam:"style=simple,explode=false,name=teamId"`
}

func (o *RequestAccessToTeamRequest) GetRequestBody() *RequestAccessToTeamRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *RequestAccessToTeamRequest) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

type RequestAccessToTeam200ApplicationJSONBitbucket struct {
	Login *string `json:"login,omitempty"`
}

func (o *RequestAccessToTeam200ApplicationJSONBitbucket) GetLogin() *string {
	if o == nil {
		return nil
	}
	return o.Login
}

type RequestAccessToTeam200ApplicationJSONGithub struct {
	Login *string `json:"login,omitempty"`
}

func (o *RequestAccessToTeam200ApplicationJSONGithub) GetLogin() *string {
	if o == nil {
		return nil
	}
	return o.Login
}

type RequestAccessToTeam200ApplicationJSONGitlab struct {
	Login *string `json:"login,omitempty"`
}

func (o *RequestAccessToTeam200ApplicationJSONGitlab) GetLogin() *string {
	if o == nil {
		return nil
	}
	return o.Login
}

type RequestAccessToTeam200ApplicationJSONJoinedFromGitUserIDType string

const (
	RequestAccessToTeam200ApplicationJSONJoinedFromGitUserIDTypeStr     RequestAccessToTeam200ApplicationJSONJoinedFromGitUserIDType = "str"
	RequestAccessToTeam200ApplicationJSONJoinedFromGitUserIDTypeInteger RequestAccessToTeam200ApplicationJSONJoinedFromGitUserIDType = "integer"
)

type RequestAccessToTeam200ApplicationJSONJoinedFromGitUserID struct {
	Str     *string
	Integer *int64

	Type RequestAccessToTeam200ApplicationJSONJoinedFromGitUserIDType
}

func CreateRequestAccessToTeam200ApplicationJSONJoinedFromGitUserIDStr(str string) RequestAccessToTeam200ApplicationJSONJoinedFromGitUserID {
	typ := RequestAccessToTeam200ApplicationJSONJoinedFromGitUserIDTypeStr

	return RequestAccessToTeam200ApplicationJSONJoinedFromGitUserID{
		Str:  &str,
		Type: typ,
	}
}

func CreateRequestAccessToTeam200ApplicationJSONJoinedFromGitUserIDInteger(integer int64) RequestAccessToTeam200ApplicationJSONJoinedFromGitUserID {
	typ := RequestAccessToTeam200ApplicationJSONJoinedFromGitUserIDTypeInteger

	return RequestAccessToTeam200ApplicationJSONJoinedFromGitUserID{
		Integer: &integer,
		Type:    typ,
	}
}

func (u *RequestAccessToTeam200ApplicationJSONJoinedFromGitUserID) UnmarshalJSON(data []byte) error {

	str := new(string)
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = str
		u.Type = RequestAccessToTeam200ApplicationJSONJoinedFromGitUserIDTypeStr
		return nil
	}

	integer := new(int64)
	if err := utils.UnmarshalJSON(data, &integer, "", true, true); err == nil {
		u.Integer = integer
		u.Type = RequestAccessToTeam200ApplicationJSONJoinedFromGitUserIDTypeInteger
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u RequestAccessToTeam200ApplicationJSONJoinedFromGitUserID) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.Integer != nil {
		return utils.MarshalJSON(u.Integer, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type RequestAccessToTeam200ApplicationJSONJoinedFromOrigin string

const (
	RequestAccessToTeam200ApplicationJSONJoinedFromOriginLink              RequestAccessToTeam200ApplicationJSONJoinedFromOrigin = "link"
	RequestAccessToTeam200ApplicationJSONJoinedFromOriginImport            RequestAccessToTeam200ApplicationJSONJoinedFromOrigin = "import"
	RequestAccessToTeam200ApplicationJSONJoinedFromOriginTeams             RequestAccessToTeam200ApplicationJSONJoinedFromOrigin = "teams"
	RequestAccessToTeam200ApplicationJSONJoinedFromOriginGithub            RequestAccessToTeam200ApplicationJSONJoinedFromOrigin = "github"
	RequestAccessToTeam200ApplicationJSONJoinedFromOriginGitlab            RequestAccessToTeam200ApplicationJSONJoinedFromOrigin = "gitlab"
	RequestAccessToTeam200ApplicationJSONJoinedFromOriginBitbucket         RequestAccessToTeam200ApplicationJSONJoinedFromOrigin = "bitbucket"
	RequestAccessToTeam200ApplicationJSONJoinedFromOriginFeedback          RequestAccessToTeam200ApplicationJSONJoinedFromOrigin = "feedback"
	RequestAccessToTeam200ApplicationJSONJoinedFromOriginOrganizationTeams RequestAccessToTeam200ApplicationJSONJoinedFromOrigin = "organization-teams"
	RequestAccessToTeam200ApplicationJSONJoinedFromOriginMail              RequestAccessToTeam200ApplicationJSONJoinedFromOrigin = "mail"
	RequestAccessToTeam200ApplicationJSONJoinedFromOriginSaml              RequestAccessToTeam200ApplicationJSONJoinedFromOrigin = "saml"
	RequestAccessToTeam200ApplicationJSONJoinedFromOriginDsync             RequestAccessToTeam200ApplicationJSONJoinedFromOrigin = "dsync"
)

func (e RequestAccessToTeam200ApplicationJSONJoinedFromOrigin) ToPointer() *RequestAccessToTeam200ApplicationJSONJoinedFromOrigin {
	return &e
}

func (e *RequestAccessToTeam200ApplicationJSONJoinedFromOrigin) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "link":
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
	case "feedback":
		fallthrough
	case "organization-teams":
		fallthrough
	case "mail":
		fallthrough
	case "saml":
		fallthrough
	case "dsync":
		*e = RequestAccessToTeam200ApplicationJSONJoinedFromOrigin(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RequestAccessToTeam200ApplicationJSONJoinedFromOrigin: %v", v)
	}
}

type RequestAccessToTeam200ApplicationJSONJoinedFrom struct {
	CommitID         *string                                                   `json:"commitId,omitempty"`
	DsyncConnectedAt *int64                                                    `json:"dsyncConnectedAt,omitempty"`
	DsyncUserID      *string                                                   `json:"dsyncUserId,omitempty"`
	GitUserID        *RequestAccessToTeam200ApplicationJSONJoinedFromGitUserID `json:"gitUserId,omitempty"`
	GitUserLogin     *string                                                   `json:"gitUserLogin,omitempty"`
	IdpUserID        *string                                                   `json:"idpUserId,omitempty"`
	Origin           RequestAccessToTeam200ApplicationJSONJoinedFromOrigin     `json:"origin"`
	RepoID           *string                                                   `json:"repoId,omitempty"`
	RepoPath         *string                                                   `json:"repoPath,omitempty"`
	SsoConnectedAt   *int64                                                    `json:"ssoConnectedAt,omitempty"`
	SsoUserID        *string                                                   `json:"ssoUserId,omitempty"`
}

func (o *RequestAccessToTeam200ApplicationJSONJoinedFrom) GetCommitID() *string {
	if o == nil {
		return nil
	}
	return o.CommitID
}

func (o *RequestAccessToTeam200ApplicationJSONJoinedFrom) GetDsyncConnectedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.DsyncConnectedAt
}

func (o *RequestAccessToTeam200ApplicationJSONJoinedFrom) GetDsyncUserID() *string {
	if o == nil {
		return nil
	}
	return o.DsyncUserID
}

func (o *RequestAccessToTeam200ApplicationJSONJoinedFrom) GetGitUserID() *RequestAccessToTeam200ApplicationJSONJoinedFromGitUserID {
	if o == nil {
		return nil
	}
	return o.GitUserID
}

func (o *RequestAccessToTeam200ApplicationJSONJoinedFrom) GetGitUserLogin() *string {
	if o == nil {
		return nil
	}
	return o.GitUserLogin
}

func (o *RequestAccessToTeam200ApplicationJSONJoinedFrom) GetIdpUserID() *string {
	if o == nil {
		return nil
	}
	return o.IdpUserID
}

func (o *RequestAccessToTeam200ApplicationJSONJoinedFrom) GetOrigin() RequestAccessToTeam200ApplicationJSONJoinedFromOrigin {
	if o == nil {
		return RequestAccessToTeam200ApplicationJSONJoinedFromOrigin("")
	}
	return o.Origin
}

func (o *RequestAccessToTeam200ApplicationJSONJoinedFrom) GetRepoID() *string {
	if o == nil {
		return nil
	}
	return o.RepoID
}

func (o *RequestAccessToTeam200ApplicationJSONJoinedFrom) GetRepoPath() *string {
	if o == nil {
		return nil
	}
	return o.RepoPath
}

func (o *RequestAccessToTeam200ApplicationJSONJoinedFrom) GetSsoConnectedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.SsoConnectedAt
}

func (o *RequestAccessToTeam200ApplicationJSONJoinedFrom) GetSsoUserID() *string {
	if o == nil {
		return nil
	}
	return o.SsoUserID
}

// RequestAccessToTeam200ApplicationJSON - Successfuly requested access to the team.
type RequestAccessToTeam200ApplicationJSON struct {
	AccessRequestedAt *int64                                           `json:"accessRequestedAt,omitempty"`
	Bitbucket         *RequestAccessToTeam200ApplicationJSONBitbucket  `json:"bitbucket"`
	Confirmed         *bool                                            `json:"confirmed,omitempty"`
	Github            *RequestAccessToTeam200ApplicationJSONGithub     `json:"github"`
	Gitlab            *RequestAccessToTeam200ApplicationJSONGitlab     `json:"gitlab"`
	JoinedFrom        *RequestAccessToTeam200ApplicationJSONJoinedFrom `json:"joinedFrom,omitempty"`
	TeamName          string                                           `json:"teamName"`
	TeamSlug          string                                           `json:"teamSlug"`
}

func (o *RequestAccessToTeam200ApplicationJSON) GetAccessRequestedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.AccessRequestedAt
}

func (o *RequestAccessToTeam200ApplicationJSON) GetBitbucket() *RequestAccessToTeam200ApplicationJSONBitbucket {
	if o == nil {
		return nil
	}
	return o.Bitbucket
}

func (o *RequestAccessToTeam200ApplicationJSON) GetConfirmed() *bool {
	if o == nil {
		return nil
	}
	return o.Confirmed
}

func (o *RequestAccessToTeam200ApplicationJSON) GetGithub() *RequestAccessToTeam200ApplicationJSONGithub {
	if o == nil {
		return nil
	}
	return o.Github
}

func (o *RequestAccessToTeam200ApplicationJSON) GetGitlab() *RequestAccessToTeam200ApplicationJSONGitlab {
	if o == nil {
		return nil
	}
	return o.Gitlab
}

func (o *RequestAccessToTeam200ApplicationJSON) GetJoinedFrom() *RequestAccessToTeam200ApplicationJSONJoinedFrom {
	if o == nil {
		return nil
	}
	return o.JoinedFrom
}

func (o *RequestAccessToTeam200ApplicationJSON) GetTeamName() string {
	if o == nil {
		return ""
	}
	return o.TeamName
}

func (o *RequestAccessToTeam200ApplicationJSON) GetTeamSlug() string {
	if o == nil {
		return ""
	}
	return o.TeamSlug
}

type RequestAccessToTeamResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfuly requested access to the team.
	RequestAccessToTeam200ApplicationJSONObject *RequestAccessToTeam200ApplicationJSON
}

func (o *RequestAccessToTeamResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RequestAccessToTeamResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RequestAccessToTeamResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *RequestAccessToTeamResponse) GetRequestAccessToTeam200ApplicationJSONObject() *RequestAccessToTeam200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.RequestAccessToTeam200ApplicationJSONObject
}
