// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zchee/terraform-provider-vercel/internal/sdk/internal/utils"
	"net/http"
)

// AddProjectMemberProjectMembersRequestRole - The project role of the member that will be added.
type AddProjectMemberProjectMembersRequestRole string

const (
	AddProjectMemberProjectMembersRequestRoleAdmin            AddProjectMemberProjectMembersRequestRole = "ADMIN"
	AddProjectMemberProjectMembersRequestRoleProjectDeveloper AddProjectMemberProjectMembersRequestRole = "PROJECT_DEVELOPER"
	AddProjectMemberProjectMembersRequestRoleProjectViewer    AddProjectMemberProjectMembersRequestRole = "PROJECT_VIEWER"
)

func (e AddProjectMemberProjectMembersRequestRole) ToPointer() *AddProjectMemberProjectMembersRequestRole {
	return &e
}

func (e *AddProjectMemberProjectMembersRequestRole) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ADMIN":
		fallthrough
	case "PROJECT_DEVELOPER":
		fallthrough
	case "PROJECT_VIEWER":
		*e = AddProjectMemberProjectMembersRequestRole(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AddProjectMemberProjectMembersRequestRole: %v", v)
	}
}

type AddProjectMember3 struct {
	// The email of the team member that should be added to this project.
	Email string `json:"email"`
	// The project role of the member that will be added.
	Role AddProjectMemberProjectMembersRequestRole `json:"role"`
	// The ID of the team member that should be added to this project.
	UID *string `json:"uid,omitempty"`
	// The username of the team member that should be added to this project.
	Username *string `json:"username,omitempty"`
}

func (o *AddProjectMember3) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

func (o *AddProjectMember3) GetRole() AddProjectMemberProjectMembersRequestRole {
	if o == nil {
		return AddProjectMemberProjectMembersRequestRole("")
	}
	return o.Role
}

func (o *AddProjectMember3) GetUID() *string {
	if o == nil {
		return nil
	}
	return o.UID
}

func (o *AddProjectMember3) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

// AddProjectMemberProjectMembersRole - The project role of the member that will be added.
type AddProjectMemberProjectMembersRole string

const (
	AddProjectMemberProjectMembersRoleAdmin            AddProjectMemberProjectMembersRole = "ADMIN"
	AddProjectMemberProjectMembersRoleProjectDeveloper AddProjectMemberProjectMembersRole = "PROJECT_DEVELOPER"
	AddProjectMemberProjectMembersRoleProjectViewer    AddProjectMemberProjectMembersRole = "PROJECT_VIEWER"
)

func (e AddProjectMemberProjectMembersRole) ToPointer() *AddProjectMemberProjectMembersRole {
	return &e
}

func (e *AddProjectMemberProjectMembersRole) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ADMIN":
		fallthrough
	case "PROJECT_DEVELOPER":
		fallthrough
	case "PROJECT_VIEWER":
		*e = AddProjectMemberProjectMembersRole(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AddProjectMemberProjectMembersRole: %v", v)
	}
}

type AddProjectMember2 struct {
	// The email of the team member that should be added to this project.
	Email *string `json:"email,omitempty"`
	// The project role of the member that will be added.
	Role AddProjectMemberProjectMembersRole `json:"role"`
	// The ID of the team member that should be added to this project.
	UID *string `json:"uid,omitempty"`
	// The username of the team member that should be added to this project.
	Username string `json:"username"`
}

func (o *AddProjectMember2) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *AddProjectMember2) GetRole() AddProjectMemberProjectMembersRole {
	if o == nil {
		return AddProjectMemberProjectMembersRole("")
	}
	return o.Role
}

func (o *AddProjectMember2) GetUID() *string {
	if o == nil {
		return nil
	}
	return o.UID
}

func (o *AddProjectMember2) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

// AddProjectMemberRole - The project role of the member that will be added.
type AddProjectMemberRole string

const (
	AddProjectMemberRoleAdmin            AddProjectMemberRole = "ADMIN"
	AddProjectMemberRoleProjectDeveloper AddProjectMemberRole = "PROJECT_DEVELOPER"
	AddProjectMemberRoleProjectViewer    AddProjectMemberRole = "PROJECT_VIEWER"
)

func (e AddProjectMemberRole) ToPointer() *AddProjectMemberRole {
	return &e
}

func (e *AddProjectMemberRole) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ADMIN":
		fallthrough
	case "PROJECT_DEVELOPER":
		fallthrough
	case "PROJECT_VIEWER":
		*e = AddProjectMemberRole(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AddProjectMemberRole: %v", v)
	}
}

type AddProjectMember1 struct {
	// The email of the team member that should be added to this project.
	Email *string `json:"email,omitempty"`
	// The project role of the member that will be added.
	Role AddProjectMemberRole `json:"role"`
	// The ID of the team member that should be added to this project.
	UID string `json:"uid"`
	// The username of the team member that should be added to this project.
	Username *string `json:"username,omitempty"`
}

func (o *AddProjectMember1) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *AddProjectMember1) GetRole() AddProjectMemberRole {
	if o == nil {
		return AddProjectMemberRole("")
	}
	return o.Role
}

func (o *AddProjectMember1) GetUID() string {
	if o == nil {
		return ""
	}
	return o.UID
}

func (o *AddProjectMember1) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

type AddProjectMemberRequestBodyType string

const (
	AddProjectMemberRequestBodyTypeAddProjectMember1 AddProjectMemberRequestBodyType = "addProjectMember_1"
	AddProjectMemberRequestBodyTypeAddProjectMember2 AddProjectMemberRequestBodyType = "addProjectMember_2"
	AddProjectMemberRequestBodyTypeAddProjectMember3 AddProjectMemberRequestBodyType = "addProjectMember_3"
)

type AddProjectMemberRequestBody struct {
	AddProjectMember1 *AddProjectMember1
	AddProjectMember2 *AddProjectMember2
	AddProjectMember3 *AddProjectMember3

	Type AddProjectMemberRequestBodyType
}

func CreateAddProjectMemberRequestBodyAddProjectMember1(addProjectMember1 AddProjectMember1) AddProjectMemberRequestBody {
	typ := AddProjectMemberRequestBodyTypeAddProjectMember1

	return AddProjectMemberRequestBody{
		AddProjectMember1: &addProjectMember1,
		Type:              typ,
	}
}

func CreateAddProjectMemberRequestBodyAddProjectMember2(addProjectMember2 AddProjectMember2) AddProjectMemberRequestBody {
	typ := AddProjectMemberRequestBodyTypeAddProjectMember2

	return AddProjectMemberRequestBody{
		AddProjectMember2: &addProjectMember2,
		Type:              typ,
	}
}

func CreateAddProjectMemberRequestBodyAddProjectMember3(addProjectMember3 AddProjectMember3) AddProjectMemberRequestBody {
	typ := AddProjectMemberRequestBodyTypeAddProjectMember3

	return AddProjectMemberRequestBody{
		AddProjectMember3: &addProjectMember3,
		Type:              typ,
	}
}

func (u *AddProjectMemberRequestBody) UnmarshalJSON(data []byte) error {

	addProjectMember1 := AddProjectMember1{}
	if err := utils.UnmarshalJSON(data, &addProjectMember1, "", true, true); err == nil {
		u.AddProjectMember1 = &addProjectMember1
		u.Type = AddProjectMemberRequestBodyTypeAddProjectMember1
		return nil
	}

	addProjectMember2 := AddProjectMember2{}
	if err := utils.UnmarshalJSON(data, &addProjectMember2, "", true, true); err == nil {
		u.AddProjectMember2 = &addProjectMember2
		u.Type = AddProjectMemberRequestBodyTypeAddProjectMember2
		return nil
	}

	addProjectMember3 := AddProjectMember3{}
	if err := utils.UnmarshalJSON(data, &addProjectMember3, "", true, true); err == nil {
		u.AddProjectMember3 = &addProjectMember3
		u.Type = AddProjectMemberRequestBodyTypeAddProjectMember3
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u AddProjectMemberRequestBody) MarshalJSON() ([]byte, error) {
	if u.AddProjectMember1 != nil {
		return utils.MarshalJSON(u.AddProjectMember1, "", true)
	}

	if u.AddProjectMember2 != nil {
		return utils.MarshalJSON(u.AddProjectMember2, "", true)
	}

	if u.AddProjectMember3 != nil {
		return utils.MarshalJSON(u.AddProjectMember3, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type AddProjectMemberRequest struct {
	RequestBody *AddProjectMemberRequestBody `request:"mediaType=application/json"`
	// The ID or name of the Project.
	IDOrName string `pathParam:"style=simple,explode=false,name=idOrName"`
	// The Team identifier or slug to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
}

func (o *AddProjectMemberRequest) GetRequestBody() *AddProjectMemberRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *AddProjectMemberRequest) GetIDOrName() string {
	if o == nil {
		return ""
	}
	return o.IDOrName
}

func (o *AddProjectMemberRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

// AddProjectMemberResponseBody - Responds with the project ID on success.
type AddProjectMemberResponseBody struct {
	ID string `json:"id"`
}

func (o *AddProjectMemberResponseBody) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type AddProjectMemberResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Responds with the project ID on success.
	Object *AddProjectMemberResponseBody
}

func (o *AddProjectMemberResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AddProjectMemberResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AddProjectMemberResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *AddProjectMemberResponse) GetObject() *AddProjectMemberResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
