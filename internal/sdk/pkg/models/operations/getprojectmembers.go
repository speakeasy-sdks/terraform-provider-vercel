// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zchee/terraform-provider-vercel/internal/sdk/pkg/utils"
	"net/http"
)

type GetProjectMembersRequest struct {
	// The ID or name of the Project.
	IDOrName string `pathParam:"style=simple,explode=false,name=idOrName"`
	// Limit how many project members should be returned
	Limit *int64 `queryParam:"style=form,explode=true,name=limit"`
	// Search project members by their name, username, and email.
	Search *string `queryParam:"style=form,explode=true,name=search"`
	// Timestamp in milliseconds to only include members added since then.
	Since *int64 `queryParam:"style=form,explode=true,name=since"`
	// The Team identifier or slug to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
	// Timestamp in milliseconds to only include members added until then.
	Until *int64 `queryParam:"style=form,explode=true,name=until"`
}

func (o *GetProjectMembersRequest) GetIDOrName() string {
	if o == nil {
		return ""
	}
	return o.IDOrName
}

func (o *GetProjectMembersRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *GetProjectMembersRequest) GetSearch() *string {
	if o == nil {
		return nil
	}
	return o.Search
}

func (o *GetProjectMembersRequest) GetSince() *int64 {
	if o == nil {
		return nil
	}
	return o.Since
}

func (o *GetProjectMembersRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

func (o *GetProjectMembersRequest) GetUntil() *int64 {
	if o == nil {
		return nil
	}
	return o.Until
}

// GetProjectMembersRole - Role of this user in the project.
type GetProjectMembersRole string

const (
	GetProjectMembersRoleAdmin            GetProjectMembersRole = "ADMIN"
	GetProjectMembersRoleProjectDeveloper GetProjectMembersRole = "PROJECT_DEVELOPER"
	GetProjectMembersRoleProjectViewer    GetProjectMembersRole = "PROJECT_VIEWER"
)

func (e GetProjectMembersRole) ToPointer() *GetProjectMembersRole {
	return &e
}

func (e *GetProjectMembersRole) UnmarshalJSON(data []byte) error {
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
		*e = GetProjectMembersRole(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetProjectMembersRole: %v", v)
	}
}

type GetProjectMembersMembers struct {
	// ID of the file for the Avatar of this member.
	Avatar *string `json:"avatar,omitempty"`
	// Timestamp in milliseconds when this member was added.
	CreatedAt int64 `json:"createdAt"`
	// The email of this member.
	Email string `json:"email"`
	// The name of this user.
	Name *string `json:"name,omitempty"`
	// Role of this user in the project.
	Role GetProjectMembersRole `json:"role"`
	// The ID of this user.
	UID string `json:"uid"`
	// The unique username of this user.
	Username string `json:"username"`
}

func (o *GetProjectMembersMembers) GetAvatar() *string {
	if o == nil {
		return nil
	}
	return o.Avatar
}

func (o *GetProjectMembersMembers) GetCreatedAt() int64 {
	if o == nil {
		return 0
	}
	return o.CreatedAt
}

func (o *GetProjectMembersMembers) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

func (o *GetProjectMembersMembers) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *GetProjectMembersMembers) GetRole() GetProjectMembersRole {
	if o == nil {
		return GetProjectMembersRole("")
	}
	return o.Role
}

func (o *GetProjectMembersMembers) GetUID() string {
	if o == nil {
		return ""
	}
	return o.UID
}

func (o *GetProjectMembersMembers) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

type GetProjectMembersPagination struct {
	// Amount of items in the current page.
	Count   int64 `json:"count"`
	HasNext bool  `json:"hasNext"`
	// Timestamp that must be used to request the next page.
	Next *int64 `json:"next"`
	// Timestamp that must be used to request the previous page.
	Prev *int64 `json:"prev"`
}

func (o *GetProjectMembersPagination) GetCount() int64 {
	if o == nil {
		return 0
	}
	return o.Count
}

func (o *GetProjectMembersPagination) GetHasNext() bool {
	if o == nil {
		return false
	}
	return o.HasNext
}

func (o *GetProjectMembersPagination) GetNext() *int64 {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *GetProjectMembersPagination) GetPrev() *int64 {
	if o == nil {
		return nil
	}
	return o.Prev
}

// GetProjectMembers2 - Paginated list of members for the project.
type GetProjectMembers2 struct {
	Members    []GetProjectMembersMembers  `json:"members"`
	Pagination GetProjectMembersPagination `json:"pagination"`
}

func (o *GetProjectMembers2) GetMembers() []GetProjectMembersMembers {
	if o == nil {
		return []GetProjectMembersMembers{}
	}
	return o.Members
}

func (o *GetProjectMembers2) GetPagination() GetProjectMembersPagination {
	if o == nil {
		return GetProjectMembersPagination{}
	}
	return o.Pagination
}

type GetProjectMembers1 struct {
}

type GetProjectMembersResponseBodyType string

const (
	GetProjectMembersResponseBodyTypeGetProjectMembers1 GetProjectMembersResponseBodyType = "getProjectMembers_1"
	GetProjectMembersResponseBodyTypeGetProjectMembers2 GetProjectMembersResponseBodyType = "getProjectMembers_2"
)

// GetProjectMembersResponseBody - Paginated list of members for the project.
type GetProjectMembersResponseBody struct {
	GetProjectMembers1 *GetProjectMembers1
	GetProjectMembers2 *GetProjectMembers2

	Type GetProjectMembersResponseBodyType
}

func CreateGetProjectMembersResponseBodyGetProjectMembers1(getProjectMembers1 GetProjectMembers1) GetProjectMembersResponseBody {
	typ := GetProjectMembersResponseBodyTypeGetProjectMembers1

	return GetProjectMembersResponseBody{
		GetProjectMembers1: &getProjectMembers1,
		Type:               typ,
	}
}

func CreateGetProjectMembersResponseBodyGetProjectMembers2(getProjectMembers2 GetProjectMembers2) GetProjectMembersResponseBody {
	typ := GetProjectMembersResponseBodyTypeGetProjectMembers2

	return GetProjectMembersResponseBody{
		GetProjectMembers2: &getProjectMembers2,
		Type:               typ,
	}
}

func (u *GetProjectMembersResponseBody) UnmarshalJSON(data []byte) error {

	getProjectMembers1 := new(GetProjectMembers1)
	if err := utils.UnmarshalJSON(data, &getProjectMembers1, "", true, true); err == nil {
		u.GetProjectMembers1 = getProjectMembers1
		u.Type = GetProjectMembersResponseBodyTypeGetProjectMembers1
		return nil
	}

	getProjectMembers2 := new(GetProjectMembers2)
	if err := utils.UnmarshalJSON(data, &getProjectMembers2, "", true, true); err == nil {
		u.GetProjectMembers2 = getProjectMembers2
		u.Type = GetProjectMembersResponseBodyTypeGetProjectMembers2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetProjectMembersResponseBody) MarshalJSON() ([]byte, error) {
	if u.GetProjectMembers1 != nil {
		return utils.MarshalJSON(u.GetProjectMembers1, "", true)
	}

	if u.GetProjectMembers2 != nil {
		return utils.MarshalJSON(u.GetProjectMembers2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetProjectMembersResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Paginated list of members for the project.
	OneOf *GetProjectMembersResponseBody
}

func (o *GetProjectMembersResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetProjectMembersResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetProjectMembersResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetProjectMembersResponse) GetOneOf() *GetProjectMembersResponseBody {
	if o == nil {
		return nil
	}
	return o.OneOf
}
