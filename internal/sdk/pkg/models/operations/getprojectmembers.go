// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"vercel/internal/sdk/pkg/utils"
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

// GetProjectMembers200ApplicationJSON2MembersRole - Role of this user in the project.
type GetProjectMembers200ApplicationJSON2MembersRole string

const (
	GetProjectMembers200ApplicationJSON2MembersRoleAdmin            GetProjectMembers200ApplicationJSON2MembersRole = "ADMIN"
	GetProjectMembers200ApplicationJSON2MembersRoleProjectDeveloper GetProjectMembers200ApplicationJSON2MembersRole = "PROJECT_DEVELOPER"
	GetProjectMembers200ApplicationJSON2MembersRoleProjectViewer    GetProjectMembers200ApplicationJSON2MembersRole = "PROJECT_VIEWER"
)

func (e GetProjectMembers200ApplicationJSON2MembersRole) ToPointer() *GetProjectMembers200ApplicationJSON2MembersRole {
	return &e
}

func (e *GetProjectMembers200ApplicationJSON2MembersRole) UnmarshalJSON(data []byte) error {
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
		*e = GetProjectMembers200ApplicationJSON2MembersRole(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetProjectMembers200ApplicationJSON2MembersRole: %v", v)
	}
}

type GetProjectMembers200ApplicationJSON2Members struct {
	// ID of the file for the Avatar of this member.
	Avatar *string `json:"avatar,omitempty"`
	// Timestamp in milliseconds when this member was added.
	CreatedAt int64 `json:"createdAt"`
	// The email of this member.
	Email string `json:"email"`
	// The name of this user.
	Name *string `json:"name,omitempty"`
	// Role of this user in the project.
	Role GetProjectMembers200ApplicationJSON2MembersRole `json:"role"`
	// The ID of this user.
	UID string `json:"uid"`
	// The unique username of this user.
	Username string `json:"username"`
}

func (o *GetProjectMembers200ApplicationJSON2Members) GetAvatar() *string {
	if o == nil {
		return nil
	}
	return o.Avatar
}

func (o *GetProjectMembers200ApplicationJSON2Members) GetCreatedAt() int64 {
	if o == nil {
		return 0
	}
	return o.CreatedAt
}

func (o *GetProjectMembers200ApplicationJSON2Members) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

func (o *GetProjectMembers200ApplicationJSON2Members) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *GetProjectMembers200ApplicationJSON2Members) GetRole() GetProjectMembers200ApplicationJSON2MembersRole {
	if o == nil {
		return GetProjectMembers200ApplicationJSON2MembersRole("")
	}
	return o.Role
}

func (o *GetProjectMembers200ApplicationJSON2Members) GetUID() string {
	if o == nil {
		return ""
	}
	return o.UID
}

func (o *GetProjectMembers200ApplicationJSON2Members) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

type GetProjectMembers200ApplicationJSON2Pagination struct {
	// Amount of items in the current page.
	Count   int64 `json:"count"`
	HasNext bool  `json:"hasNext"`
	// Timestamp that must be used to request the next page.
	Next *int64 `json:"next"`
	// Timestamp that must be used to request the previous page.
	Prev *int64 `json:"prev"`
}

func (o *GetProjectMembers200ApplicationJSON2Pagination) GetCount() int64 {
	if o == nil {
		return 0
	}
	return o.Count
}

func (o *GetProjectMembers200ApplicationJSON2Pagination) GetHasNext() bool {
	if o == nil {
		return false
	}
	return o.HasNext
}

func (o *GetProjectMembers200ApplicationJSON2Pagination) GetNext() *int64 {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *GetProjectMembers200ApplicationJSON2Pagination) GetPrev() *int64 {
	if o == nil {
		return nil
	}
	return o.Prev
}

// GetProjectMembers200ApplicationJSON2 - Paginated list of members for the project.
type GetProjectMembers200ApplicationJSON2 struct {
	Members    []GetProjectMembers200ApplicationJSON2Members  `json:"members"`
	Pagination GetProjectMembers200ApplicationJSON2Pagination `json:"pagination"`
}

func (o *GetProjectMembers200ApplicationJSON2) GetMembers() []GetProjectMembers200ApplicationJSON2Members {
	if o == nil {
		return []GetProjectMembers200ApplicationJSON2Members{}
	}
	return o.Members
}

func (o *GetProjectMembers200ApplicationJSON2) GetPagination() GetProjectMembers200ApplicationJSON2Pagination {
	if o == nil {
		return GetProjectMembers200ApplicationJSON2Pagination{}
	}
	return o.Pagination
}

type GetProjectMembers200ApplicationJSON1 struct {
}

type GetProjectMembers200ApplicationJSONType string

const (
	GetProjectMembers200ApplicationJSONTypeGetProjectMembers200ApplicationJSON1 GetProjectMembers200ApplicationJSONType = "getProjectMembers_200ApplicationJSON_1"
	GetProjectMembers200ApplicationJSONTypeGetProjectMembers200ApplicationJSON2 GetProjectMembers200ApplicationJSONType = "getProjectMembers_200ApplicationJSON_2"
)

type GetProjectMembers200ApplicationJSON struct {
	GetProjectMembers200ApplicationJSON1 *GetProjectMembers200ApplicationJSON1
	GetProjectMembers200ApplicationJSON2 *GetProjectMembers200ApplicationJSON2

	Type GetProjectMembers200ApplicationJSONType
}

func CreateGetProjectMembers200ApplicationJSONGetProjectMembers200ApplicationJSON1(getProjectMembers200ApplicationJSON1 GetProjectMembers200ApplicationJSON1) GetProjectMembers200ApplicationJSON {
	typ := GetProjectMembers200ApplicationJSONTypeGetProjectMembers200ApplicationJSON1

	return GetProjectMembers200ApplicationJSON{
		GetProjectMembers200ApplicationJSON1: &getProjectMembers200ApplicationJSON1,
		Type:                                 typ,
	}
}

func CreateGetProjectMembers200ApplicationJSONGetProjectMembers200ApplicationJSON2(getProjectMembers200ApplicationJSON2 GetProjectMembers200ApplicationJSON2) GetProjectMembers200ApplicationJSON {
	typ := GetProjectMembers200ApplicationJSONTypeGetProjectMembers200ApplicationJSON2

	return GetProjectMembers200ApplicationJSON{
		GetProjectMembers200ApplicationJSON2: &getProjectMembers200ApplicationJSON2,
		Type:                                 typ,
	}
}

func (u *GetProjectMembers200ApplicationJSON) UnmarshalJSON(data []byte) error {

	getProjectMembers200ApplicationJSON1 := new(GetProjectMembers200ApplicationJSON1)
	if err := utils.UnmarshalJSON(data, &getProjectMembers200ApplicationJSON1, "", true, true); err == nil {
		u.GetProjectMembers200ApplicationJSON1 = getProjectMembers200ApplicationJSON1
		u.Type = GetProjectMembers200ApplicationJSONTypeGetProjectMembers200ApplicationJSON1
		return nil
	}

	getProjectMembers200ApplicationJSON2 := new(GetProjectMembers200ApplicationJSON2)
	if err := utils.UnmarshalJSON(data, &getProjectMembers200ApplicationJSON2, "", true, true); err == nil {
		u.GetProjectMembers200ApplicationJSON2 = getProjectMembers200ApplicationJSON2
		u.Type = GetProjectMembers200ApplicationJSONTypeGetProjectMembers200ApplicationJSON2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetProjectMembers200ApplicationJSON) MarshalJSON() ([]byte, error) {
	if u.GetProjectMembers200ApplicationJSON1 != nil {
		return utils.MarshalJSON(u.GetProjectMembers200ApplicationJSON1, "", true)
	}

	if u.GetProjectMembers200ApplicationJSON2 != nil {
		return utils.MarshalJSON(u.GetProjectMembers200ApplicationJSON2, "", true)
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
	GetProjectMembers200ApplicationJSONOneOf *GetProjectMembers200ApplicationJSON
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

func (o *GetProjectMembersResponse) GetGetProjectMembers200ApplicationJSONOneOf() *GetProjectMembers200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetProjectMembers200ApplicationJSONOneOf
}
