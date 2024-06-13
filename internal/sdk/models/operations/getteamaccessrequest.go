// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/vercel/terraform-provider-vercel/internal/sdk/internal/utils"
	"net/http"
)

type GetTeamAccessRequestRequest struct {
	TeamID string `pathParam:"style=simple,explode=false,name=teamId"`
	UserID string `pathParam:"style=simple,explode=false,name=userId"`
}

func (o *GetTeamAccessRequestRequest) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

func (o *GetTeamAccessRequestRequest) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

type GetTeamAccessRequestOrigin string

const (
	GetTeamAccessRequestOriginMail              GetTeamAccessRequestOrigin = "mail"
	GetTeamAccessRequestOriginLink              GetTeamAccessRequestOrigin = "link"
	GetTeamAccessRequestOriginImport            GetTeamAccessRequestOrigin = "import"
	GetTeamAccessRequestOriginTeams             GetTeamAccessRequestOrigin = "teams"
	GetTeamAccessRequestOriginGithub            GetTeamAccessRequestOrigin = "github"
	GetTeamAccessRequestOriginGitlab            GetTeamAccessRequestOrigin = "gitlab"
	GetTeamAccessRequestOriginBitbucket         GetTeamAccessRequestOrigin = "bitbucket"
	GetTeamAccessRequestOriginSaml              GetTeamAccessRequestOrigin = "saml"
	GetTeamAccessRequestOriginDsync             GetTeamAccessRequestOrigin = "dsync"
	GetTeamAccessRequestOriginFeedback          GetTeamAccessRequestOrigin = "feedback"
	GetTeamAccessRequestOriginOrganizationTeams GetTeamAccessRequestOrigin = "organization-teams"
)

func (e GetTeamAccessRequestOrigin) ToPointer() *GetTeamAccessRequestOrigin {
	return &e
}
func (e *GetTeamAccessRequestOrigin) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "mail":
		fallthrough
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
	case "saml":
		fallthrough
	case "dsync":
		fallthrough
	case "feedback":
		fallthrough
	case "organization-teams":
		*e = GetTeamAccessRequestOrigin(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetTeamAccessRequestOrigin: %v", v)
	}
}

type GetTeamAccessRequestGitUserIDType string

const (
	GetTeamAccessRequestGitUserIDTypeStr    GetTeamAccessRequestGitUserIDType = "str"
	GetTeamAccessRequestGitUserIDTypeNumber GetTeamAccessRequestGitUserIDType = "number"
)

type GetTeamAccessRequestGitUserID struct {
	Str    *string
	Number *float64

	Type GetTeamAccessRequestGitUserIDType
}

func CreateGetTeamAccessRequestGitUserIDStr(str string) GetTeamAccessRequestGitUserID {
	typ := GetTeamAccessRequestGitUserIDTypeStr

	return GetTeamAccessRequestGitUserID{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetTeamAccessRequestGitUserIDNumber(number float64) GetTeamAccessRequestGitUserID {
	typ := GetTeamAccessRequestGitUserIDTypeNumber

	return GetTeamAccessRequestGitUserID{
		Number: &number,
		Type:   typ,
	}
}

func (u *GetTeamAccessRequestGitUserID) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = GetTeamAccessRequestGitUserIDTypeStr
		return nil
	}

	var number float64 = float64(0)
	if err := utils.UnmarshalJSON(data, &number, "", true, true); err == nil {
		u.Number = &number
		u.Type = GetTeamAccessRequestGitUserIDTypeNumber
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for GetTeamAccessRequestGitUserID", string(data))
}

func (u GetTeamAccessRequestGitUserID) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.Number != nil {
		return utils.MarshalJSON(u.Number, "", true)
	}

	return nil, errors.New("could not marshal union type GetTeamAccessRequestGitUserID: all fields are null")
}

// GetTeamAccessRequestJoinedFrom - A map that describes the origin from where the user joined.
type GetTeamAccessRequestJoinedFrom struct {
	Origin           GetTeamAccessRequestOrigin     `json:"origin"`
	CommitID         *string                        `json:"commitId,omitempty"`
	RepoID           *string                        `json:"repoId,omitempty"`
	RepoPath         *string                        `json:"repoPath,omitempty"`
	GitUserID        *GetTeamAccessRequestGitUserID `json:"gitUserId,omitempty"`
	GitUserLogin     *string                        `json:"gitUserLogin,omitempty"`
	SsoUserID        *string                        `json:"ssoUserId,omitempty"`
	SsoConnectedAt   *float64                       `json:"ssoConnectedAt,omitempty"`
	IdpUserID        *string                        `json:"idpUserId,omitempty"`
	DsyncUserID      *string                        `json:"dsyncUserId,omitempty"`
	DsyncConnectedAt *float64                       `json:"dsyncConnectedAt,omitempty"`
}

func (o *GetTeamAccessRequestJoinedFrom) GetOrigin() GetTeamAccessRequestOrigin {
	if o == nil {
		return GetTeamAccessRequestOrigin("")
	}
	return o.Origin
}

func (o *GetTeamAccessRequestJoinedFrom) GetCommitID() *string {
	if o == nil {
		return nil
	}
	return o.CommitID
}

func (o *GetTeamAccessRequestJoinedFrom) GetRepoID() *string {
	if o == nil {
		return nil
	}
	return o.RepoID
}

func (o *GetTeamAccessRequestJoinedFrom) GetRepoPath() *string {
	if o == nil {
		return nil
	}
	return o.RepoPath
}

func (o *GetTeamAccessRequestJoinedFrom) GetGitUserID() *GetTeamAccessRequestGitUserID {
	if o == nil {
		return nil
	}
	return o.GitUserID
}

func (o *GetTeamAccessRequestJoinedFrom) GetGitUserLogin() *string {
	if o == nil {
		return nil
	}
	return o.GitUserLogin
}

func (o *GetTeamAccessRequestJoinedFrom) GetSsoUserID() *string {
	if o == nil {
		return nil
	}
	return o.SsoUserID
}

func (o *GetTeamAccessRequestJoinedFrom) GetSsoConnectedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.SsoConnectedAt
}

func (o *GetTeamAccessRequestJoinedFrom) GetIdpUserID() *string {
	if o == nil {
		return nil
	}
	return o.IdpUserID
}

func (o *GetTeamAccessRequestJoinedFrom) GetDsyncUserID() *string {
	if o == nil {
		return nil
	}
	return o.DsyncUserID
}

func (o *GetTeamAccessRequestJoinedFrom) GetDsyncConnectedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.DsyncConnectedAt
}

// GetTeamAccessRequestGithub - Map of the connected GitHub account.
type GetTeamAccessRequestGithub struct {
	Login *string `json:"login,omitempty"`
}

func (o *GetTeamAccessRequestGithub) GetLogin() *string {
	if o == nil {
		return nil
	}
	return o.Login
}

// GetTeamAccessRequestGitlab - Map of the connected GitLab account.
type GetTeamAccessRequestGitlab struct {
	Login *string `json:"login,omitempty"`
}

func (o *GetTeamAccessRequestGitlab) GetLogin() *string {
	if o == nil {
		return nil
	}
	return o.Login
}

// GetTeamAccessRequestBitbucket - Map of the connected Bitbucket account.
type GetTeamAccessRequestBitbucket struct {
	Login *string `json:"login,omitempty"`
}

func (o *GetTeamAccessRequestBitbucket) GetLogin() *string {
	if o == nil {
		return nil
	}
	return o.Login
}

// GetTeamAccessRequestResponseBody - Successfully
type GetTeamAccessRequestResponseBody struct {
	// The slug of the team.
	TeamSlug string `json:"teamSlug"`
	// The name of the team.
	TeamName string `json:"teamName"`
	// Current status of the membership. Will be `true` if confirmed, if pending it'll be `false`.
	Confirmed bool `json:"confirmed"`
	// A map that describes the origin from where the user joined.
	JoinedFrom GetTeamAccessRequestJoinedFrom `json:"joinedFrom"`
	// Timestamp in milliseconds when the user requested access to the team.
	AccessRequestedAt float64 `json:"accessRequestedAt"`
	// Map of the connected GitHub account.
	Github *GetTeamAccessRequestGithub `json:"github"`
	// Map of the connected GitLab account.
	Gitlab *GetTeamAccessRequestGitlab `json:"gitlab"`
	// Map of the connected Bitbucket account.
	Bitbucket *GetTeamAccessRequestBitbucket `json:"bitbucket"`
}

func (o *GetTeamAccessRequestResponseBody) GetTeamSlug() string {
	if o == nil {
		return ""
	}
	return o.TeamSlug
}

func (o *GetTeamAccessRequestResponseBody) GetTeamName() string {
	if o == nil {
		return ""
	}
	return o.TeamName
}

func (o *GetTeamAccessRequestResponseBody) GetConfirmed() bool {
	if o == nil {
		return false
	}
	return o.Confirmed
}

func (o *GetTeamAccessRequestResponseBody) GetJoinedFrom() GetTeamAccessRequestJoinedFrom {
	if o == nil {
		return GetTeamAccessRequestJoinedFrom{}
	}
	return o.JoinedFrom
}

func (o *GetTeamAccessRequestResponseBody) GetAccessRequestedAt() float64 {
	if o == nil {
		return 0.0
	}
	return o.AccessRequestedAt
}

func (o *GetTeamAccessRequestResponseBody) GetGithub() *GetTeamAccessRequestGithub {
	if o == nil {
		return nil
	}
	return o.Github
}

func (o *GetTeamAccessRequestResponseBody) GetGitlab() *GetTeamAccessRequestGitlab {
	if o == nil {
		return nil
	}
	return o.Gitlab
}

func (o *GetTeamAccessRequestResponseBody) GetBitbucket() *GetTeamAccessRequestBitbucket {
	if o == nil {
		return nil
	}
	return o.Bitbucket
}

type GetTeamAccessRequestResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully
	Object *GetTeamAccessRequestResponseBody
}

func (o *GetTeamAccessRequestResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetTeamAccessRequestResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetTeamAccessRequestResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetTeamAccessRequestResponse) GetObject() *GetTeamAccessRequestResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
