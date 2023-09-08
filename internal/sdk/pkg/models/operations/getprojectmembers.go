// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type GetProjectMembersSecurity struct {
	BearerToken string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

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

type GetProjectMembers200ApplicationJSON2Pagination struct {
	// Amount of items in the current page.
	Count   int64 `json:"count"`
	HasNext bool  `json:"hasNext"`
	// Timestamp that must be used to request the next page.
	Next *int64 `json:"next"`
	// Timestamp that must be used to request the previous page.
	Prev *int64 `json:"prev"`
}

// GetProjectMembers200ApplicationJSON2 - Paginated list of members for the project.
type GetProjectMembers200ApplicationJSON2 struct {
	Members    []GetProjectMembers200ApplicationJSON2Members  `json:"members"`
	Pagination GetProjectMembers200ApplicationJSON2Pagination `json:"pagination"`
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
	var d *json.Decoder

	getProjectMembers200ApplicationJSON1 := new(GetProjectMembers200ApplicationJSON1)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&getProjectMembers200ApplicationJSON1); err == nil {
		u.GetProjectMembers200ApplicationJSON1 = getProjectMembers200ApplicationJSON1
		u.Type = GetProjectMembers200ApplicationJSONTypeGetProjectMembers200ApplicationJSON1
		return nil
	}

	getProjectMembers200ApplicationJSON2 := new(GetProjectMembers200ApplicationJSON2)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&getProjectMembers200ApplicationJSON2); err == nil {
		u.GetProjectMembers200ApplicationJSON2 = getProjectMembers200ApplicationJSON2
		u.Type = GetProjectMembers200ApplicationJSONTypeGetProjectMembers200ApplicationJSON2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetProjectMembers200ApplicationJSON) MarshalJSON() ([]byte, error) {
	if u.GetProjectMembers200ApplicationJSON1 != nil {
		return json.Marshal(u.GetProjectMembers200ApplicationJSON1)
	}

	if u.GetProjectMembers200ApplicationJSON2 != nil {
		return json.Marshal(u.GetProjectMembers200ApplicationJSON2)
	}

	return nil, nil
}

type GetProjectMembersResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Paginated list of members for the project.
	GetProjectMembers200ApplicationJSONOneOf *GetProjectMembers200ApplicationJSON
}
