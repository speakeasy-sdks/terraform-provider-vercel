// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"vercel/internal/sdk/pkg/utils"
)

type TeamLimitedGitUserIDType string

const (
	TeamLimitedGitUserIDTypeStr     TeamLimitedGitUserIDType = "str"
	TeamLimitedGitUserIDTypeInteger TeamLimitedGitUserIDType = "integer"
)

type TeamLimitedGitUserID struct {
	Str     *string
	Integer *int64

	Type TeamLimitedGitUserIDType
}

func CreateTeamLimitedGitUserIDStr(str string) TeamLimitedGitUserID {
	typ := TeamLimitedGitUserIDTypeStr

	return TeamLimitedGitUserID{
		Str:  &str,
		Type: typ,
	}
}

func CreateTeamLimitedGitUserIDInteger(integer int64) TeamLimitedGitUserID {
	typ := TeamLimitedGitUserIDTypeInteger

	return TeamLimitedGitUserID{
		Integer: &integer,
		Type:    typ,
	}
}

func (u *TeamLimitedGitUserID) UnmarshalJSON(data []byte) error {

	str := new(string)
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = str
		u.Type = TeamLimitedGitUserIDTypeStr
		return nil
	}

	integer := new(int64)
	if err := utils.UnmarshalJSON(data, &integer, "", true, true); err == nil {
		u.Integer = integer
		u.Type = TeamLimitedGitUserIDTypeInteger
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u TeamLimitedGitUserID) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.Integer != nil {
		return utils.MarshalJSON(u.Integer, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type TeamLimitedSchemasOrigin string

const (
	TeamLimitedSchemasOriginLink              TeamLimitedSchemasOrigin = "link"
	TeamLimitedSchemasOriginSaml              TeamLimitedSchemasOrigin = "saml"
	TeamLimitedSchemasOriginMail              TeamLimitedSchemasOrigin = "mail"
	TeamLimitedSchemasOriginImport            TeamLimitedSchemasOrigin = "import"
	TeamLimitedSchemasOriginTeams             TeamLimitedSchemasOrigin = "teams"
	TeamLimitedSchemasOriginGithub            TeamLimitedSchemasOrigin = "github"
	TeamLimitedSchemasOriginGitlab            TeamLimitedSchemasOrigin = "gitlab"
	TeamLimitedSchemasOriginBitbucket         TeamLimitedSchemasOrigin = "bitbucket"
	TeamLimitedSchemasOriginDsync             TeamLimitedSchemasOrigin = "dsync"
	TeamLimitedSchemasOriginFeedback          TeamLimitedSchemasOrigin = "feedback"
	TeamLimitedSchemasOriginOrganizationTeams TeamLimitedSchemasOrigin = "organization-teams"
)

func (e TeamLimitedSchemasOrigin) ToPointer() *TeamLimitedSchemasOrigin {
	return &e
}

func (e *TeamLimitedSchemasOrigin) UnmarshalJSON(data []byte) error {
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
		*e = TeamLimitedSchemasOrigin(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TeamLimitedSchemasOrigin: %v", v)
	}
}

type TeamLimitedJoinedFrom struct {
	CommitID         *string                  `json:"commitId,omitempty"`
	DsyncConnectedAt *int64                   `json:"dsyncConnectedAt,omitempty"`
	DsyncUserID      *string                  `json:"dsyncUserId,omitempty"`
	GitUserID        *TeamLimitedGitUserID    `json:"gitUserId,omitempty"`
	GitUserLogin     *string                  `json:"gitUserLogin,omitempty"`
	IdpUserID        *string                  `json:"idpUserId,omitempty"`
	Origin           TeamLimitedSchemasOrigin `json:"origin"`
	RepoID           *string                  `json:"repoId,omitempty"`
	RepoPath         *string                  `json:"repoPath,omitempty"`
	SsoConnectedAt   *int64                   `json:"ssoConnectedAt,omitempty"`
	SsoUserID        *string                  `json:"ssoUserId,omitempty"`
}

func (o *TeamLimitedJoinedFrom) GetCommitID() *string {
	if o == nil {
		return nil
	}
	return o.CommitID
}

func (o *TeamLimitedJoinedFrom) GetDsyncConnectedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.DsyncConnectedAt
}

func (o *TeamLimitedJoinedFrom) GetDsyncUserID() *string {
	if o == nil {
		return nil
	}
	return o.DsyncUserID
}

func (o *TeamLimitedJoinedFrom) GetGitUserID() *TeamLimitedGitUserID {
	if o == nil {
		return nil
	}
	return o.GitUserID
}

func (o *TeamLimitedJoinedFrom) GetGitUserLogin() *string {
	if o == nil {
		return nil
	}
	return o.GitUserLogin
}

func (o *TeamLimitedJoinedFrom) GetIdpUserID() *string {
	if o == nil {
		return nil
	}
	return o.IdpUserID
}

func (o *TeamLimitedJoinedFrom) GetOrigin() TeamLimitedSchemasOrigin {
	if o == nil {
		return TeamLimitedSchemasOrigin("")
	}
	return o.Origin
}

func (o *TeamLimitedJoinedFrom) GetRepoID() *string {
	if o == nil {
		return nil
	}
	return o.RepoID
}

func (o *TeamLimitedJoinedFrom) GetRepoPath() *string {
	if o == nil {
		return nil
	}
	return o.RepoPath
}

func (o *TeamLimitedJoinedFrom) GetSsoConnectedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.SsoConnectedAt
}

func (o *TeamLimitedJoinedFrom) GetSsoUserID() *string {
	if o == nil {
		return nil
	}
	return o.SsoUserID
}

type TeamLimitedRole string

const (
	TeamLimitedRoleOwner       TeamLimitedRole = "OWNER"
	TeamLimitedRoleMember      TeamLimitedRole = "MEMBER"
	TeamLimitedRoleViewer      TeamLimitedRole = "VIEWER"
	TeamLimitedRoleDeveloper   TeamLimitedRole = "DEVELOPER"
	TeamLimitedRoleBilling     TeamLimitedRole = "BILLING"
	TeamLimitedRoleContributor TeamLimitedRole = "CONTRIBUTOR"
)

func (e TeamLimitedRole) ToPointer() *TeamLimitedRole {
	return &e
}

func (e *TeamLimitedRole) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "OWNER":
		fallthrough
	case "MEMBER":
		fallthrough
	case "VIEWER":
		fallthrough
	case "DEVELOPER":
		fallthrough
	case "BILLING":
		fallthrough
	case "CONTRIBUTOR":
		*e = TeamLimitedRole(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TeamLimitedRole: %v", v)
	}
}

// TeamLimited2 - The membership of the authenticated User in relation to the Team.
type TeamLimited2 struct {
	AccessRequestedAt int64                  `json:"accessRequestedAt"`
	Confirmed         bool                   `json:"confirmed"`
	ConfirmedAt       *int64                 `json:"confirmedAt,omitempty"`
	Created           int64                  `json:"created"`
	CreatedAt         int64                  `json:"createdAt"`
	JoinedFrom        *TeamLimitedJoinedFrom `json:"joinedFrom,omitempty"`
	Role              TeamLimitedRole        `json:"role"`
	TeamID            *string                `json:"teamId,omitempty"`
	UID               string                 `json:"uid"`
}

func (o *TeamLimited2) GetAccessRequestedAt() int64 {
	if o == nil {
		return 0
	}
	return o.AccessRequestedAt
}

func (o *TeamLimited2) GetConfirmed() bool {
	if o == nil {
		return false
	}
	return o.Confirmed
}

func (o *TeamLimited2) GetConfirmedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.ConfirmedAt
}

func (o *TeamLimited2) GetCreated() int64 {
	if o == nil {
		return 0
	}
	return o.Created
}

func (o *TeamLimited2) GetCreatedAt() int64 {
	if o == nil {
		return 0
	}
	return o.CreatedAt
}

func (o *TeamLimited2) GetJoinedFrom() *TeamLimitedJoinedFrom {
	if o == nil {
		return nil
	}
	return o.JoinedFrom
}

func (o *TeamLimited2) GetRole() TeamLimitedRole {
	if o == nil {
		return TeamLimitedRole("")
	}
	return o.Role
}

func (o *TeamLimited2) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

func (o *TeamLimited2) GetUID() string {
	if o == nil {
		return ""
	}
	return o.UID
}

type GitUserIDType string

const (
	GitUserIDTypeStr     GitUserIDType = "str"
	GitUserIDTypeInteger GitUserIDType = "integer"
)

type GitUserID struct {
	Str     *string
	Integer *int64

	Type GitUserIDType
}

func CreateGitUserIDStr(str string) GitUserID {
	typ := GitUserIDTypeStr

	return GitUserID{
		Str:  &str,
		Type: typ,
	}
}

func CreateGitUserIDInteger(integer int64) GitUserID {
	typ := GitUserIDTypeInteger

	return GitUserID{
		Integer: &integer,
		Type:    typ,
	}
}

func (u *GitUserID) UnmarshalJSON(data []byte) error {

	str := new(string)
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = str
		u.Type = GitUserIDTypeStr
		return nil
	}

	integer := new(int64)
	if err := utils.UnmarshalJSON(data, &integer, "", true, true); err == nil {
		u.Integer = integer
		u.Type = GitUserIDTypeInteger
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GitUserID) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.Integer != nil {
		return utils.MarshalJSON(u.Integer, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type TeamLimitedOrigin string

const (
	TeamLimitedOriginLink              TeamLimitedOrigin = "link"
	TeamLimitedOriginSaml              TeamLimitedOrigin = "saml"
	TeamLimitedOriginMail              TeamLimitedOrigin = "mail"
	TeamLimitedOriginImport            TeamLimitedOrigin = "import"
	TeamLimitedOriginTeams             TeamLimitedOrigin = "teams"
	TeamLimitedOriginGithub            TeamLimitedOrigin = "github"
	TeamLimitedOriginGitlab            TeamLimitedOrigin = "gitlab"
	TeamLimitedOriginBitbucket         TeamLimitedOrigin = "bitbucket"
	TeamLimitedOriginDsync             TeamLimitedOrigin = "dsync"
	TeamLimitedOriginFeedback          TeamLimitedOrigin = "feedback"
	TeamLimitedOriginOrganizationTeams TeamLimitedOrigin = "organization-teams"
)

func (e TeamLimitedOrigin) ToPointer() *TeamLimitedOrigin {
	return &e
}

func (e *TeamLimitedOrigin) UnmarshalJSON(data []byte) error {
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
		*e = TeamLimitedOrigin(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TeamLimitedOrigin: %v", v)
	}
}

type JoinedFrom struct {
	CommitID         *string           `json:"commitId,omitempty"`
	DsyncConnectedAt *int64            `json:"dsyncConnectedAt,omitempty"`
	DsyncUserID      *string           `json:"dsyncUserId,omitempty"`
	GitUserID        *GitUserID        `json:"gitUserId,omitempty"`
	GitUserLogin     *string           `json:"gitUserLogin,omitempty"`
	IdpUserID        *string           `json:"idpUserId,omitempty"`
	Origin           TeamLimitedOrigin `json:"origin"`
	RepoID           *string           `json:"repoId,omitempty"`
	RepoPath         *string           `json:"repoPath,omitempty"`
	SsoConnectedAt   *int64            `json:"ssoConnectedAt,omitempty"`
	SsoUserID        *string           `json:"ssoUserId,omitempty"`
}

func (o *JoinedFrom) GetCommitID() *string {
	if o == nil {
		return nil
	}
	return o.CommitID
}

func (o *JoinedFrom) GetDsyncConnectedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.DsyncConnectedAt
}

func (o *JoinedFrom) GetDsyncUserID() *string {
	if o == nil {
		return nil
	}
	return o.DsyncUserID
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

func (o *JoinedFrom) GetIdpUserID() *string {
	if o == nil {
		return nil
	}
	return o.IdpUserID
}

func (o *JoinedFrom) GetOrigin() TeamLimitedOrigin {
	if o == nil {
		return TeamLimitedOrigin("")
	}
	return o.Origin
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

func (o *JoinedFrom) GetSsoConnectedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.SsoConnectedAt
}

func (o *JoinedFrom) GetSsoUserID() *string {
	if o == nil {
		return nil
	}
	return o.SsoUserID
}

type Role string

const (
	RoleOwner       Role = "OWNER"
	RoleMember      Role = "MEMBER"
	RoleViewer      Role = "VIEWER"
	RoleDeveloper   Role = "DEVELOPER"
	RoleBilling     Role = "BILLING"
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
	case "VIEWER":
		fallthrough
	case "DEVELOPER":
		fallthrough
	case "BILLING":
		fallthrough
	case "CONTRIBUTOR":
		*e = Role(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Role: %v", v)
	}
}

// TeamLimited1 - The membership of the authenticated User in relation to the Team.
type TeamLimited1 struct {
	AccessRequestedAt *int64      `json:"accessRequestedAt,omitempty"`
	Confirmed         bool        `json:"confirmed"`
	ConfirmedAt       int64       `json:"confirmedAt"`
	Created           int64       `json:"created"`
	CreatedAt         int64       `json:"createdAt"`
	JoinedFrom        *JoinedFrom `json:"joinedFrom,omitempty"`
	Role              Role        `json:"role"`
	TeamID            *string     `json:"teamId,omitempty"`
	UID               string      `json:"uid"`
}

func (o *TeamLimited1) GetAccessRequestedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.AccessRequestedAt
}

func (o *TeamLimited1) GetConfirmed() bool {
	if o == nil {
		return false
	}
	return o.Confirmed
}

func (o *TeamLimited1) GetConfirmedAt() int64 {
	if o == nil {
		return 0
	}
	return o.ConfirmedAt
}

func (o *TeamLimited1) GetCreated() int64 {
	if o == nil {
		return 0
	}
	return o.Created
}

func (o *TeamLimited1) GetCreatedAt() int64 {
	if o == nil {
		return 0
	}
	return o.CreatedAt
}

func (o *TeamLimited1) GetJoinedFrom() *JoinedFrom {
	if o == nil {
		return nil
	}
	return o.JoinedFrom
}

func (o *TeamLimited1) GetRole() Role {
	if o == nil {
		return Role("")
	}
	return o.Role
}

func (o *TeamLimited1) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

func (o *TeamLimited1) GetUID() string {
	if o == nil {
		return ""
	}
	return o.UID
}

type MembershipType string

const (
	MembershipTypeTeamLimited1 MembershipType = "TeamLimited_1"
	MembershipTypeTeamLimited2 MembershipType = "TeamLimited_2"
)

type Membership struct {
	TeamLimited1 *TeamLimited1
	TeamLimited2 *TeamLimited2

	Type MembershipType
}

func CreateMembershipTeamLimited1(teamLimited1 TeamLimited1) Membership {
	typ := MembershipTypeTeamLimited1

	return Membership{
		TeamLimited1: &teamLimited1,
		Type:         typ,
	}
}

func CreateMembershipTeamLimited2(teamLimited2 TeamLimited2) Membership {
	typ := MembershipTypeTeamLimited2

	return Membership{
		TeamLimited2: &teamLimited2,
		Type:         typ,
	}
}

func (u *Membership) UnmarshalJSON(data []byte) error {

	teamLimited1 := new(TeamLimited1)
	if err := utils.UnmarshalJSON(data, &teamLimited1, "", true, true); err == nil {
		u.TeamLimited1 = teamLimited1
		u.Type = MembershipTypeTeamLimited1
		return nil
	}

	teamLimited2 := new(TeamLimited2)
	if err := utils.UnmarshalJSON(data, &teamLimited2, "", true, true); err == nil {
		u.TeamLimited2 = teamLimited2
		u.Type = MembershipTypeTeamLimited2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u Membership) MarshalJSON() ([]byte, error) {
	if u.TeamLimited1 != nil {
		return utils.MarshalJSON(u.TeamLimited1, "", true)
	}

	if u.TeamLimited2 != nil {
		return utils.MarshalJSON(u.TeamLimited2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// Connection - Information for the SAML Single Sign-On configuration.
type Connection struct {
	// Timestamp (in milliseconds) of when the configuration was connected.
	ConnectedAt int64 `json:"connectedAt"`
	// Timestamp (in milliseconds) of when the last webhook event was received from WorkOS.
	LastReceivedWebhookEvent *int64 `json:"lastReceivedWebhookEvent,omitempty"`
	// Current state of the connection.
	State string `json:"state"`
	// Current status of the connection.
	Status string `json:"status"`
	// The Identity Provider "type", for example Okta.
	Type string `json:"type"`
}

func (o *Connection) GetConnectedAt() int64 {
	if o == nil {
		return 0
	}
	return o.ConnectedAt
}

func (o *Connection) GetLastReceivedWebhookEvent() *int64 {
	if o == nil {
		return nil
	}
	return o.LastReceivedWebhookEvent
}

func (o *Connection) GetState() string {
	if o == nil {
		return ""
	}
	return o.State
}

func (o *Connection) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

func (o *Connection) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

// Directory - Information for the SAML Single Sign-On configuration.
type Directory struct {
	// Timestamp (in milliseconds) of when the configuration was connected.
	ConnectedAt int64 `json:"connectedAt"`
	// Timestamp (in milliseconds) of when the last webhook event was received from WorkOS.
	LastReceivedWebhookEvent *int64 `json:"lastReceivedWebhookEvent,omitempty"`
	// Current state of the connection.
	State string `json:"state"`
	// Current status of the connection.
	Status string `json:"status"`
	// The Identity Provider "type", for example Okta.
	Type string `json:"type"`
}

func (o *Directory) GetConnectedAt() int64 {
	if o == nil {
		return 0
	}
	return o.ConnectedAt
}

func (o *Directory) GetLastReceivedWebhookEvent() *int64 {
	if o == nil {
		return nil
	}
	return o.LastReceivedWebhookEvent
}

func (o *Directory) GetState() string {
	if o == nil {
		return ""
	}
	return o.State
}

func (o *Directory) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

func (o *Directory) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
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

// TeamLimited - A limited form of data representing a Team, due to the authentication token missing privileges to read the full Team data.
type TeamLimited struct {
	// The ID of the file used as avatar for this Team.
	Avatar *string `json:"avatar"`
	// Will remain undocumented. Remove in v3 API.
	Created string `json:"created"`
	// UNIX timestamp (in milliseconds) when the Team was created.
	CreatedAt int64 `json:"createdAt"`
	// The Team's unique identifier.
	ID string `json:"id"`
	// Property indicating that this Team data contains only limited information, due to the authentication token missing privileges to read the full Team data. Re-login with the Team's configured SAML Single Sign-On provider in order to upgrade the authentication token with the necessary privileges.
	Limited    bool       `json:"limited"`
	Membership Membership `json:"membership"`
	// Name associated with the Team account, or `null` if none has been provided.
	Name *string `json:"name"`
	// When "Single Sign-On (SAML)" is configured, this object contains information that allows the client-side to identify whether or not this Team has SAML enforced.
	Saml *Saml `json:"saml,omitempty"`
	// The Team's slug, which is unique across the Vercel platform.
	Slug string `json:"slug"`
}

func (o *TeamLimited) GetAvatar() *string {
	if o == nil {
		return nil
	}
	return o.Avatar
}

func (o *TeamLimited) GetCreated() string {
	if o == nil {
		return ""
	}
	return o.Created
}

func (o *TeamLimited) GetCreatedAt() int64 {
	if o == nil {
		return 0
	}
	return o.CreatedAt
}

func (o *TeamLimited) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *TeamLimited) GetLimited() bool {
	if o == nil {
		return false
	}
	return o.Limited
}

func (o *TeamLimited) GetMembership() Membership {
	if o == nil {
		return Membership{}
	}
	return o.Membership
}

func (o *TeamLimited) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *TeamLimited) GetSaml() *Saml {
	if o == nil {
		return nil
	}
	return o.Saml
}

func (o *TeamLimited) GetSlug() string {
	if o == nil {
		return ""
	}
	return o.Slug
}
