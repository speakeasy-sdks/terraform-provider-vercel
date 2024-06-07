// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteTeamInviteCodeRequest struct {
	// the team id related to the invite code
	TeamID string `pathParam:"style=simple,explode=false,name=teamId"`
	// The Team invite code ID.
	InviteID string `pathParam:"style=simple,explode=false,name=inviteId"`
}

func (o *DeleteTeamInviteCodeRequest) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

func (o *DeleteTeamInviteCodeRequest) GetInviteID() string {
	if o == nil {
		return ""
	}
	return o.InviteID
}

// DeleteTeamInviteCodeResponseBody - Successfully deleted Team invite code.
type DeleteTeamInviteCodeResponseBody struct {
	// ID of the team.
	ID string `json:"id"`
}

func (o *DeleteTeamInviteCodeResponseBody) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type DeleteTeamInviteCodeResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully deleted Team invite code.
	Object *DeleteTeamInviteCodeResponseBody
}

func (o *DeleteTeamInviteCodeResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteTeamInviteCodeResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteTeamInviteCodeResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteTeamInviteCodeResponse) GetObject() *DeleteTeamInviteCodeResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
