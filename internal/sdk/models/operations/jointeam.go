// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type JoinTeamRequestBody struct {
	// The invite code to join the team.
	InviteCode *string `json:"inviteCode,omitempty"`
	// The team ID.
	TeamID *string `json:"teamId,omitempty"`
}

func (o *JoinTeamRequestBody) GetInviteCode() *string {
	if o == nil {
		return nil
	}
	return o.InviteCode
}

func (o *JoinTeamRequestBody) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

type JoinTeamRequest struct {
	RequestBody *JoinTeamRequestBody `request:"mediaType=application/json"`
	TeamID      string               `pathParam:"style=simple,explode=false,name=teamId"`
}

func (o *JoinTeamRequest) GetRequestBody() *JoinTeamRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *JoinTeamRequest) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

// JoinTeamResponseBody - Successfully joined a team.
type JoinTeamResponseBody struct {
	// The origin of how the user joined.
	From string `json:"from"`
	// The name of the team the user joined.
	Name string `json:"name"`
	// The slug of the team the user joined.
	Slug string `json:"slug"`
	// The ID of the team the user joined.
	TeamID string `json:"teamId"`
}

func (o *JoinTeamResponseBody) GetFrom() string {
	if o == nil {
		return ""
	}
	return o.From
}

func (o *JoinTeamResponseBody) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *JoinTeamResponseBody) GetSlug() string {
	if o == nil {
		return ""
	}
	return o.Slug
}

func (o *JoinTeamResponseBody) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

type JoinTeamResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully joined a team.
	Object *JoinTeamResponseBody
}

func (o *JoinTeamResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *JoinTeamResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *JoinTeamResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *JoinTeamResponse) GetObject() *JoinTeamResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}