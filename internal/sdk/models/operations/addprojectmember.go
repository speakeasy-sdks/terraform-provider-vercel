// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/vercel/terraform-provider-terraform/internal/sdk/internal/utils"
	"net/http"
)

// AddProjectMemberRequestBodyRole - The project role of the member that will be added.
type AddProjectMemberRequestBodyRole string

const (
	AddProjectMemberRequestBodyRoleAdmin            AddProjectMemberRequestBodyRole = "ADMIN"
	AddProjectMemberRequestBodyRoleProjectDeveloper AddProjectMemberRequestBodyRole = "PROJECT_DEVELOPER"
	AddProjectMemberRequestBodyRoleProjectViewer    AddProjectMemberRequestBodyRole = "PROJECT_VIEWER"
)

func (e AddProjectMemberRequestBodyRole) ToPointer() *AddProjectMemberRequestBodyRole {
	return &e
}
func (e *AddProjectMemberRequestBodyRole) UnmarshalJSON(data []byte) error {
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
		*e = AddProjectMemberRequestBodyRole(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AddProjectMemberRequestBodyRole: %v", v)
	}
}

type AddProjectMemberRequestBody3 struct {
	// The ID of the team member that should be added to this project.
	UID *string `json:"uid,omitempty"`
	// The username of the team member that should be added to this project.
	Username *string `json:"username,omitempty"`
	// The email of the team member that should be added to this project.
	Email string `json:"email"`
	// The project role of the member that will be added.
	Role AddProjectMemberRequestBodyRole `json:"role"`
}

func (o *AddProjectMemberRequestBody3) GetUID() *string {
	if o == nil {
		return nil
	}
	return o.UID
}

func (o *AddProjectMemberRequestBody3) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

func (o *AddProjectMemberRequestBody3) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

func (o *AddProjectMemberRequestBody3) GetRole() AddProjectMemberRequestBodyRole {
	if o == nil {
		return AddProjectMemberRequestBodyRole("")
	}
	return o.Role
}

// RequestBodyRole - The project role of the member that will be added.
type RequestBodyRole string

const (
	RequestBodyRoleAdmin            RequestBodyRole = "ADMIN"
	RequestBodyRoleProjectDeveloper RequestBodyRole = "PROJECT_DEVELOPER"
	RequestBodyRoleProjectViewer    RequestBodyRole = "PROJECT_VIEWER"
)

func (e RequestBodyRole) ToPointer() *RequestBodyRole {
	return &e
}
func (e *RequestBodyRole) UnmarshalJSON(data []byte) error {
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
		*e = RequestBodyRole(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RequestBodyRole: %v", v)
	}
}

type AddProjectMemberRequestBody2 struct {
	// The ID of the team member that should be added to this project.
	UID *string `json:"uid,omitempty"`
	// The username of the team member that should be added to this project.
	Username string `json:"username"`
	// The email of the team member that should be added to this project.
	Email *string `json:"email,omitempty"`
	// The project role of the member that will be added.
	Role RequestBodyRole `json:"role"`
}

func (o *AddProjectMemberRequestBody2) GetUID() *string {
	if o == nil {
		return nil
	}
	return o.UID
}

func (o *AddProjectMemberRequestBody2) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

func (o *AddProjectMemberRequestBody2) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *AddProjectMemberRequestBody2) GetRole() RequestBodyRole {
	if o == nil {
		return RequestBodyRole("")
	}
	return o.Role
}

// AddProjectMemberRequestBodyProjectMembersRole - The project role of the member that will be added.
type AddProjectMemberRequestBodyProjectMembersRole string

const (
	AddProjectMemberRequestBodyProjectMembersRoleAdmin            AddProjectMemberRequestBodyProjectMembersRole = "ADMIN"
	AddProjectMemberRequestBodyProjectMembersRoleProjectDeveloper AddProjectMemberRequestBodyProjectMembersRole = "PROJECT_DEVELOPER"
	AddProjectMemberRequestBodyProjectMembersRoleProjectViewer    AddProjectMemberRequestBodyProjectMembersRole = "PROJECT_VIEWER"
)

func (e AddProjectMemberRequestBodyProjectMembersRole) ToPointer() *AddProjectMemberRequestBodyProjectMembersRole {
	return &e
}
func (e *AddProjectMemberRequestBodyProjectMembersRole) UnmarshalJSON(data []byte) error {
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
		*e = AddProjectMemberRequestBodyProjectMembersRole(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AddProjectMemberRequestBodyProjectMembersRole: %v", v)
	}
}

type AddProjectMemberRequestBody1 struct {
	// The ID of the team member that should be added to this project.
	UID string `json:"uid"`
	// The username of the team member that should be added to this project.
	Username *string `json:"username,omitempty"`
	// The email of the team member that should be added to this project.
	Email *string `json:"email,omitempty"`
	// The project role of the member that will be added.
	Role AddProjectMemberRequestBodyProjectMembersRole `json:"role"`
}

func (o *AddProjectMemberRequestBody1) GetUID() string {
	if o == nil {
		return ""
	}
	return o.UID
}

func (o *AddProjectMemberRequestBody1) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

func (o *AddProjectMemberRequestBody1) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *AddProjectMemberRequestBody1) GetRole() AddProjectMemberRequestBodyProjectMembersRole {
	if o == nil {
		return AddProjectMemberRequestBodyProjectMembersRole("")
	}
	return o.Role
}

type AddProjectMemberRequestBodyType string

const (
	AddProjectMemberRequestBodyTypeAddProjectMemberRequestBody1 AddProjectMemberRequestBodyType = "addProjectMember_requestBody_1"
	AddProjectMemberRequestBodyTypeAddProjectMemberRequestBody2 AddProjectMemberRequestBodyType = "addProjectMember_requestBody_2"
	AddProjectMemberRequestBodyTypeAddProjectMemberRequestBody3 AddProjectMemberRequestBodyType = "addProjectMember_requestBody_3"
)

type AddProjectMemberRequestBody struct {
	AddProjectMemberRequestBody1 *AddProjectMemberRequestBody1
	AddProjectMemberRequestBody2 *AddProjectMemberRequestBody2
	AddProjectMemberRequestBody3 *AddProjectMemberRequestBody3

	Type AddProjectMemberRequestBodyType
}

func CreateAddProjectMemberRequestBodyAddProjectMemberRequestBody1(addProjectMemberRequestBody1 AddProjectMemberRequestBody1) AddProjectMemberRequestBody {
	typ := AddProjectMemberRequestBodyTypeAddProjectMemberRequestBody1

	return AddProjectMemberRequestBody{
		AddProjectMemberRequestBody1: &addProjectMemberRequestBody1,
		Type:                         typ,
	}
}

func CreateAddProjectMemberRequestBodyAddProjectMemberRequestBody2(addProjectMemberRequestBody2 AddProjectMemberRequestBody2) AddProjectMemberRequestBody {
	typ := AddProjectMemberRequestBodyTypeAddProjectMemberRequestBody2

	return AddProjectMemberRequestBody{
		AddProjectMemberRequestBody2: &addProjectMemberRequestBody2,
		Type:                         typ,
	}
}

func CreateAddProjectMemberRequestBodyAddProjectMemberRequestBody3(addProjectMemberRequestBody3 AddProjectMemberRequestBody3) AddProjectMemberRequestBody {
	typ := AddProjectMemberRequestBodyTypeAddProjectMemberRequestBody3

	return AddProjectMemberRequestBody{
		AddProjectMemberRequestBody3: &addProjectMemberRequestBody3,
		Type:                         typ,
	}
}

func (u *AddProjectMemberRequestBody) UnmarshalJSON(data []byte) error {

	var addProjectMemberRequestBody1 AddProjectMemberRequestBody1 = AddProjectMemberRequestBody1{}
	if err := utils.UnmarshalJSON(data, &addProjectMemberRequestBody1, "", true, true); err == nil {
		u.AddProjectMemberRequestBody1 = &addProjectMemberRequestBody1
		u.Type = AddProjectMemberRequestBodyTypeAddProjectMemberRequestBody1
		return nil
	}

	var addProjectMemberRequestBody2 AddProjectMemberRequestBody2 = AddProjectMemberRequestBody2{}
	if err := utils.UnmarshalJSON(data, &addProjectMemberRequestBody2, "", true, true); err == nil {
		u.AddProjectMemberRequestBody2 = &addProjectMemberRequestBody2
		u.Type = AddProjectMemberRequestBodyTypeAddProjectMemberRequestBody2
		return nil
	}

	var addProjectMemberRequestBody3 AddProjectMemberRequestBody3 = AddProjectMemberRequestBody3{}
	if err := utils.UnmarshalJSON(data, &addProjectMemberRequestBody3, "", true, true); err == nil {
		u.AddProjectMemberRequestBody3 = &addProjectMemberRequestBody3
		u.Type = AddProjectMemberRequestBodyTypeAddProjectMemberRequestBody3
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for AddProjectMemberRequestBody", string(data))
}

func (u AddProjectMemberRequestBody) MarshalJSON() ([]byte, error) {
	if u.AddProjectMemberRequestBody1 != nil {
		return utils.MarshalJSON(u.AddProjectMemberRequestBody1, "", true)
	}

	if u.AddProjectMemberRequestBody2 != nil {
		return utils.MarshalJSON(u.AddProjectMemberRequestBody2, "", true)
	}

	if u.AddProjectMemberRequestBody3 != nil {
		return utils.MarshalJSON(u.AddProjectMemberRequestBody3, "", true)
	}

	return nil, errors.New("could not marshal union type AddProjectMemberRequestBody: all fields are null")
}

type AddProjectMemberRequest struct {
	// The ID or name of the Project.
	IDOrName string `pathParam:"style=simple,explode=false,name=idOrName"`
	// The Team identifier to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
	// The Team slug to perform the request on behalf of.
	Slug        *string                      `queryParam:"style=form,explode=true,name=slug"`
	RequestBody *AddProjectMemberRequestBody `request:"mediaType=application/json"`
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

func (o *AddProjectMemberRequest) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *AddProjectMemberRequest) GetRequestBody() *AddProjectMemberRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
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
