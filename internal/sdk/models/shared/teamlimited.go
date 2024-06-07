// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/internal/utils"
)

// Connection - Information for the SAML Single Sign-On configuration.
type Connection struct {
	// The Identity Provider "type", for example Okta.
	Type string `json:"type"`
	// Current status of the connection.
	Status string `json:"status"`
	// Current state of the connection.
	State string `json:"state"`
	// Timestamp (in milliseconds) of when the configuration was connected.
	ConnectedAt float64 `json:"connectedAt"`
	// Timestamp (in milliseconds) of when the last webhook event was received from WorkOS.
	LastReceivedWebhookEvent *float64 `json:"lastReceivedWebhookEvent,omitempty"`
}

func (o *Connection) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

func (o *Connection) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

func (o *Connection) GetState() string {
	if o == nil {
		return ""
	}
	return o.State
}

func (o *Connection) GetConnectedAt() float64 {
	if o == nil {
		return 0.0
	}
	return o.ConnectedAt
}

func (o *Connection) GetLastReceivedWebhookEvent() *float64 {
	if o == nil {
		return nil
	}
	return o.LastReceivedWebhookEvent
}

// Directory - Information for the SAML Single Sign-On configuration.
type Directory struct {
	// The Identity Provider "type", for example Okta.
	Type string `json:"type"`
	// Current status of the connection.
	Status string `json:"status"`
	// Current state of the connection.
	State string `json:"state"`
	// Timestamp (in milliseconds) of when the configuration was connected.
	ConnectedAt float64 `json:"connectedAt"`
	// Timestamp (in milliseconds) of when the last webhook event was received from WorkOS.
	LastReceivedWebhookEvent *float64 `json:"lastReceivedWebhookEvent,omitempty"`
}

func (o *Directory) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

func (o *Directory) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

func (o *Directory) GetState() string {
	if o == nil {
		return ""
	}
	return o.State
}

func (o *Directory) GetConnectedAt() float64 {
	if o == nil {
		return 0.0
	}
	return o.ConnectedAt
}

func (o *Directory) GetLastReceivedWebhookEvent() *float64 {
	if o == nil {
		return nil
	}
	return o.LastReceivedWebhookEvent
}

// Saml - When "Single Sign-On (SAML)" is configured, this object contains information that allows the client-side to identify whether or not this Team has SAML enforced.
type Saml struct {
	// Information for the SAML Single Sign-On configuration.
	Connection *Connection `json:"connection,omitempty"`
	// Information for the SAML Single Sign-On configuration.
	Directory *Directory `json:"directory,omitempty"`
	// When `true`, interactions with the Team **must** be done with an authentication token that has been authenticated with the Team's SAML Single Sign-On provider.
	Enforced bool `json:"enforced"`
}

func (o *Saml) GetConnection() *Connection {
	if o == nil {
		return nil
	}
	return o.Connection
}

func (o *Saml) GetDirectory() *Directory {
	if o == nil {
		return nil
	}
	return o.Directory
}

func (o *Saml) GetEnforced() bool {
	if o == nil {
		return false
	}
	return o.Enforced
}

type MembershipRole string

const (
	MembershipRoleOwner       MembershipRole = "OWNER"
	MembershipRoleMember      MembershipRole = "MEMBER"
	MembershipRoleDeveloper   MembershipRole = "DEVELOPER"
	MembershipRoleBilling     MembershipRole = "BILLING"
	MembershipRoleViewer      MembershipRole = "VIEWER"
	MembershipRoleContributor MembershipRole = "CONTRIBUTOR"
)

func (e MembershipRole) ToPointer() *MembershipRole {
	return &e
}
func (e *MembershipRole) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "OWNER":
		fallthrough
	case "MEMBER":
		fallthrough
	case "DEVELOPER":
		fallthrough
	case "BILLING":
		fallthrough
	case "VIEWER":
		fallthrough
	case "CONTRIBUTOR":
		*e = MembershipRole(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MembershipRole: %v", v)
	}
}

type MembershipOrigin string

const (
	MembershipOriginLink              MembershipOrigin = "link"
	MembershipOriginSaml              MembershipOrigin = "saml"
	MembershipOriginMail              MembershipOrigin = "mail"
	MembershipOriginImport            MembershipOrigin = "import"
	MembershipOriginTeams             MembershipOrigin = "teams"
	MembershipOriginGithub            MembershipOrigin = "github"
	MembershipOriginGitlab            MembershipOrigin = "gitlab"
	MembershipOriginBitbucket         MembershipOrigin = "bitbucket"
	MembershipOriginDsync             MembershipOrigin = "dsync"
	MembershipOriginFeedback          MembershipOrigin = "feedback"
	MembershipOriginOrganizationTeams MembershipOrigin = "organization-teams"
)

func (e MembershipOrigin) ToPointer() *MembershipOrigin {
	return &e
}
func (e *MembershipOrigin) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "link":
		fallthrough
	case "saml":
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
	case "dsync":
		fallthrough
	case "feedback":
		fallthrough
	case "organization-teams":
		*e = MembershipOrigin(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MembershipOrigin: %v", v)
	}
}

type MembershipGitUserIDType string

const (
	MembershipGitUserIDTypeStr    MembershipGitUserIDType = "str"
	MembershipGitUserIDTypeNumber MembershipGitUserIDType = "number"
)

type MembershipGitUserID struct {
	Str    *string
	Number *float64

	Type MembershipGitUserIDType
}

func CreateMembershipGitUserIDStr(str string) MembershipGitUserID {
	typ := MembershipGitUserIDTypeStr

	return MembershipGitUserID{
		Str:  &str,
		Type: typ,
	}
}

func CreateMembershipGitUserIDNumber(number float64) MembershipGitUserID {
	typ := MembershipGitUserIDTypeNumber

	return MembershipGitUserID{
		Number: &number,
		Type:   typ,
	}
}

func (u *MembershipGitUserID) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = MembershipGitUserIDTypeStr
		return nil
	}

	var number float64 = float64(0)
	if err := utils.UnmarshalJSON(data, &number, "", true, true); err == nil {
		u.Number = &number
		u.Type = MembershipGitUserIDTypeNumber
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for MembershipGitUserID", string(data))
}

func (u MembershipGitUserID) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.Number != nil {
		return utils.MarshalJSON(u.Number, "", true)
	}

	return nil, errors.New("could not marshal union type MembershipGitUserID: all fields are null")
}

type MembershipJoinedFrom struct {
	Origin           MembershipOrigin     `json:"origin"`
	CommitID         *string              `json:"commitId,omitempty"`
	RepoID           *string              `json:"repoId,omitempty"`
	RepoPath         *string              `json:"repoPath,omitempty"`
	GitUserID        *MembershipGitUserID `json:"gitUserId,omitempty"`
	GitUserLogin     *string              `json:"gitUserLogin,omitempty"`
	SsoUserID        *string              `json:"ssoUserId,omitempty"`
	SsoConnectedAt   *float64             `json:"ssoConnectedAt,omitempty"`
	IdpUserID        *string              `json:"idpUserId,omitempty"`
	DsyncUserID      *string              `json:"dsyncUserId,omitempty"`
	DsyncConnectedAt *float64             `json:"dsyncConnectedAt,omitempty"`
}

func (o *MembershipJoinedFrom) GetOrigin() MembershipOrigin {
	if o == nil {
		return MembershipOrigin("")
	}
	return o.Origin
}

func (o *MembershipJoinedFrom) GetCommitID() *string {
	if o == nil {
		return nil
	}
	return o.CommitID
}

func (o *MembershipJoinedFrom) GetRepoID() *string {
	if o == nil {
		return nil
	}
	return o.RepoID
}

func (o *MembershipJoinedFrom) GetRepoPath() *string {
	if o == nil {
		return nil
	}
	return o.RepoPath
}

func (o *MembershipJoinedFrom) GetGitUserID() *MembershipGitUserID {
	if o == nil {
		return nil
	}
	return o.GitUserID
}

func (o *MembershipJoinedFrom) GetGitUserLogin() *string {
	if o == nil {
		return nil
	}
	return o.GitUserLogin
}

func (o *MembershipJoinedFrom) GetSsoUserID() *string {
	if o == nil {
		return nil
	}
	return o.SsoUserID
}

func (o *MembershipJoinedFrom) GetSsoConnectedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.SsoConnectedAt
}

func (o *MembershipJoinedFrom) GetIdpUserID() *string {
	if o == nil {
		return nil
	}
	return o.IdpUserID
}

func (o *MembershipJoinedFrom) GetDsyncUserID() *string {
	if o == nil {
		return nil
	}
	return o.DsyncUserID
}

func (o *MembershipJoinedFrom) GetDsyncConnectedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.DsyncConnectedAt
}

// Two - The membership of the authenticated User in relation to the Team.
type Two struct {
	Confirmed         bool                  `json:"confirmed"`
	ConfirmedAt       *float64              `json:"confirmedAt,omitempty"`
	AccessRequestedAt float64               `json:"accessRequestedAt"`
	Role              MembershipRole        `json:"role"`
	TeamID            *string               `json:"teamId,omitempty"`
	UID               string                `json:"uid"`
	CreatedAt         float64               `json:"createdAt"`
	Created           float64               `json:"created"`
	JoinedFrom        *MembershipJoinedFrom `json:"joinedFrom,omitempty"`
}

func (o *Two) GetConfirmed() bool {
	if o == nil {
		return false
	}
	return o.Confirmed
}

func (o *Two) GetConfirmedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.ConfirmedAt
}

func (o *Two) GetAccessRequestedAt() float64 {
	if o == nil {
		return 0.0
	}
	return o.AccessRequestedAt
}

func (o *Two) GetRole() MembershipRole {
	if o == nil {
		return MembershipRole("")
	}
	return o.Role
}

func (o *Two) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

func (o *Two) GetUID() string {
	if o == nil {
		return ""
	}
	return o.UID
}

func (o *Two) GetCreatedAt() float64 {
	if o == nil {
		return 0.0
	}
	return o.CreatedAt
}

func (o *Two) GetCreated() float64 {
	if o == nil {
		return 0.0
	}
	return o.Created
}

func (o *Two) GetJoinedFrom() *MembershipJoinedFrom {
	if o == nil {
		return nil
	}
	return o.JoinedFrom
}

type Role string

const (
	RoleOwner       Role = "OWNER"
	RoleMember      Role = "MEMBER"
	RoleDeveloper   Role = "DEVELOPER"
	RoleBilling     Role = "BILLING"
	RoleViewer      Role = "VIEWER"
	RoleContributor Role = "CONTRIBUTOR"
)

func (e Role) ToPointer() *Role {
	return &e
}
func (e *Role) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "OWNER":
		fallthrough
	case "MEMBER":
		fallthrough
	case "DEVELOPER":
		fallthrough
	case "BILLING":
		fallthrough
	case "VIEWER":
		fallthrough
	case "CONTRIBUTOR":
		*e = Role(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Role: %v", v)
	}
}

type TeamLimitedMembershipOrigin string

const (
	TeamLimitedMembershipOriginLink              TeamLimitedMembershipOrigin = "link"
	TeamLimitedMembershipOriginSaml              TeamLimitedMembershipOrigin = "saml"
	TeamLimitedMembershipOriginMail              TeamLimitedMembershipOrigin = "mail"
	TeamLimitedMembershipOriginImport            TeamLimitedMembershipOrigin = "import"
	TeamLimitedMembershipOriginTeams             TeamLimitedMembershipOrigin = "teams"
	TeamLimitedMembershipOriginGithub            TeamLimitedMembershipOrigin = "github"
	TeamLimitedMembershipOriginGitlab            TeamLimitedMembershipOrigin = "gitlab"
	TeamLimitedMembershipOriginBitbucket         TeamLimitedMembershipOrigin = "bitbucket"
	TeamLimitedMembershipOriginDsync             TeamLimitedMembershipOrigin = "dsync"
	TeamLimitedMembershipOriginFeedback          TeamLimitedMembershipOrigin = "feedback"
	TeamLimitedMembershipOriginOrganizationTeams TeamLimitedMembershipOrigin = "organization-teams"
)

func (e TeamLimitedMembershipOrigin) ToPointer() *TeamLimitedMembershipOrigin {
	return &e
}
func (e *TeamLimitedMembershipOrigin) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "link":
		fallthrough
	case "saml":
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
	case "dsync":
		fallthrough
	case "feedback":
		fallthrough
	case "organization-teams":
		*e = TeamLimitedMembershipOrigin(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TeamLimitedMembershipOrigin: %v", v)
	}
}

type GitUserIDType string

const (
	GitUserIDTypeStr    GitUserIDType = "str"
	GitUserIDTypeNumber GitUserIDType = "number"
)

type GitUserID struct {
	Str    *string
	Number *float64

	Type GitUserIDType
}

func CreateGitUserIDStr(str string) GitUserID {
	typ := GitUserIDTypeStr

	return GitUserID{
		Str:  &str,
		Type: typ,
	}
}

func CreateGitUserIDNumber(number float64) GitUserID {
	typ := GitUserIDTypeNumber

	return GitUserID{
		Number: &number,
		Type:   typ,
	}
}

func (u *GitUserID) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = GitUserIDTypeStr
		return nil
	}

	var number float64 = float64(0)
	if err := utils.UnmarshalJSON(data, &number, "", true, true); err == nil {
		u.Number = &number
		u.Type = GitUserIDTypeNumber
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for GitUserID", string(data))
}

func (u GitUserID) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.Number != nil {
		return utils.MarshalJSON(u.Number, "", true)
	}

	return nil, errors.New("could not marshal union type GitUserID: all fields are null")
}

type JoinedFrom struct {
	Origin           TeamLimitedMembershipOrigin `json:"origin"`
	CommitID         *string                     `json:"commitId,omitempty"`
	RepoID           *string                     `json:"repoId,omitempty"`
	RepoPath         *string                     `json:"repoPath,omitempty"`
	GitUserID        *GitUserID                  `json:"gitUserId,omitempty"`
	GitUserLogin     *string                     `json:"gitUserLogin,omitempty"`
	SsoUserID        *string                     `json:"ssoUserId,omitempty"`
	SsoConnectedAt   *float64                    `json:"ssoConnectedAt,omitempty"`
	IdpUserID        *string                     `json:"idpUserId,omitempty"`
	DsyncUserID      *string                     `json:"dsyncUserId,omitempty"`
	DsyncConnectedAt *float64                    `json:"dsyncConnectedAt,omitempty"`
}

func (o *JoinedFrom) GetOrigin() TeamLimitedMembershipOrigin {
	if o == nil {
		return TeamLimitedMembershipOrigin("")
	}
	return o.Origin
}

func (o *JoinedFrom) GetCommitID() *string {
	if o == nil {
		return nil
	}
	return o.CommitID
}

func (o *JoinedFrom) GetRepoID() *string {
	if o == nil {
		return nil
	}
	return o.RepoID
}

func (o *JoinedFrom) GetRepoPath() *string {
	if o == nil {
		return nil
	}
	return o.RepoPath
}

func (o *JoinedFrom) GetGitUserID() *GitUserID {
	if o == nil {
		return nil
	}
	return o.GitUserID
}

func (o *JoinedFrom) GetGitUserLogin() *string {
	if o == nil {
		return nil
	}
	return o.GitUserLogin
}

func (o *JoinedFrom) GetSsoUserID() *string {
	if o == nil {
		return nil
	}
	return o.SsoUserID
}

func (o *JoinedFrom) GetSsoConnectedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.SsoConnectedAt
}

func (o *JoinedFrom) GetIdpUserID() *string {
	if o == nil {
		return nil
	}
	return o.IdpUserID
}

func (o *JoinedFrom) GetDsyncUserID() *string {
	if o == nil {
		return nil
	}
	return o.DsyncUserID
}

func (o *JoinedFrom) GetDsyncConnectedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.DsyncConnectedAt
}

// One - The membership of the authenticated User in relation to the Team.
type One struct {
	Confirmed         bool        `json:"confirmed"`
	ConfirmedAt       float64     `json:"confirmedAt"`
	AccessRequestedAt *float64    `json:"accessRequestedAt,omitempty"`
	Role              Role        `json:"role"`
	TeamID            *string     `json:"teamId,omitempty"`
	UID               string      `json:"uid"`
	CreatedAt         float64     `json:"createdAt"`
	Created           float64     `json:"created"`
	JoinedFrom        *JoinedFrom `json:"joinedFrom,omitempty"`
}

func (o *One) GetConfirmed() bool {
	if o == nil {
		return false
	}
	return o.Confirmed
}

func (o *One) GetConfirmedAt() float64 {
	if o == nil {
		return 0.0
	}
	return o.ConfirmedAt
}

func (o *One) GetAccessRequestedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.AccessRequestedAt
}

func (o *One) GetRole() Role {
	if o == nil {
		return Role("")
	}
	return o.Role
}

func (o *One) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

func (o *One) GetUID() string {
	if o == nil {
		return ""
	}
	return o.UID
}

func (o *One) GetCreatedAt() float64 {
	if o == nil {
		return 0.0
	}
	return o.CreatedAt
}

func (o *One) GetCreated() float64 {
	if o == nil {
		return 0.0
	}
	return o.Created
}

func (o *One) GetJoinedFrom() *JoinedFrom {
	if o == nil {
		return nil
	}
	return o.JoinedFrom
}

type MembershipType string

const (
	MembershipTypeOne MembershipType = "1"
	MembershipTypeTwo MembershipType = "2"
)

type Membership struct {
	One *One
	Two *Two

	Type MembershipType
}

func CreateMembershipOne(one One) Membership {
	typ := MembershipTypeOne

	return Membership{
		One:  &one,
		Type: typ,
	}
}

func CreateMembershipTwo(two Two) Membership {
	typ := MembershipTypeTwo

	return Membership{
		Two:  &two,
		Type: typ,
	}
}

func (u *Membership) UnmarshalJSON(data []byte) error {

	var one One = One{}
	if err := utils.UnmarshalJSON(data, &one, "", true, true); err == nil {
		u.One = &one
		u.Type = MembershipTypeOne
		return nil
	}

	var two Two = Two{}
	if err := utils.UnmarshalJSON(data, &two, "", true, true); err == nil {
		u.Two = &two
		u.Type = MembershipTypeTwo
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for Membership", string(data))
}

func (u Membership) MarshalJSON() ([]byte, error) {
	if u.One != nil {
		return utils.MarshalJSON(u.One, "", true)
	}

	if u.Two != nil {
		return utils.MarshalJSON(u.Two, "", true)
	}

	return nil, errors.New("could not marshal union type Membership: all fields are null")
}

// TeamLimited - A limited form of data representing a Team, due to the authentication token missing privileges to read the full Team data.
type TeamLimited struct {
	// Property indicating that this Team data contains only limited information, due to the authentication token missing privileges to read the full Team data. Re-login with the Team's configured SAML Single Sign-On provider in order to upgrade the authentication token with the necessary privileges.
	Limited bool `json:"limited"`
	// When "Single Sign-On (SAML)" is configured, this object contains information that allows the client-side to identify whether or not this Team has SAML enforced.
	Saml *Saml `json:"saml,omitempty"`
	// The Team's unique identifier.
	ID string `json:"id"`
	// The Team's slug, which is unique across the Vercel platform.
	Slug string `json:"slug"`
	// Name associated with the Team account, or `null` if none has been provided.
	Name *string `json:"name"`
	// The ID of the file used as avatar for this Team.
	Avatar     *string    `json:"avatar"`
	Membership Membership `json:"membership"`
	// Will remain undocumented. Remove in v3 API.
	Created string `json:"created"`
	// UNIX timestamp (in milliseconds) when the Team was created.
	CreatedAt float64 `json:"createdAt"`
}

func (o *TeamLimited) GetLimited() bool {
	if o == nil {
		return false
	}
	return o.Limited
}

func (o *TeamLimited) GetSaml() *Saml {
	if o == nil {
		return nil
	}
	return o.Saml
}

func (o *TeamLimited) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *TeamLimited) GetSlug() string {
	if o == nil {
		return ""
	}
	return o.Slug
}

func (o *TeamLimited) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *TeamLimited) GetAvatar() *string {
	if o == nil {
		return nil
	}
	return o.Avatar
}

func (o *TeamLimited) GetMembership() Membership {
	if o == nil {
		return Membership{}
	}
	return o.Membership
}

func (o *TeamLimited) GetCreated() string {
	if o == nil {
		return ""
	}
	return o.Created
}

func (o *TeamLimited) GetCreatedAt() float64 {
	if o == nil {
		return 0.0
	}
	return o.CreatedAt
}
