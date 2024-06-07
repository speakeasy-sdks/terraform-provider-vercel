// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type RemoveProjectMemberRequest struct {
	// The ID or name of the Project.
	IDOrName string `pathParam:"style=simple,explode=false,name=idOrName"`
	// The Team slug to perform the request on behalf of.
	Slug *string `queryParam:"style=form,explode=true,name=slug"`
	// The Team identifier to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
	// The user ID of the member.
	UID string `pathParam:"style=simple,explode=false,name=uid"`
}

func (o *RemoveProjectMemberRequest) GetIDOrName() string {
	if o == nil {
		return ""
	}
	return o.IDOrName
}

func (o *RemoveProjectMemberRequest) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *RemoveProjectMemberRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

func (o *RemoveProjectMemberRequest) GetUID() string {
	if o == nil {
		return ""
	}
	return o.UID
}

type RemoveProjectMemberResponseBody struct {
	ID string `json:"id"`
}

func (o *RemoveProjectMemberResponseBody) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type RemoveProjectMemberResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	Object      *RemoveProjectMemberResponseBody
}

func (o *RemoveProjectMemberResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RemoveProjectMemberResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RemoveProjectMemberResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *RemoveProjectMemberResponse) GetObject() *RemoveProjectMemberResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
