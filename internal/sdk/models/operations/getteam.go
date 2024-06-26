// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/vercel/terraform-provider-vercel/internal/sdk/models/shared"
	"net/http"
)

type GetTeamRequest struct {
	// The Team identifier to perform the request on behalf of.
	TeamID string `pathParam:"style=simple,explode=false,name=teamId"`
	// The Team slug to perform the request on behalf of.
	Slug *string `queryParam:"style=form,explode=true,name=slug"`
}

func (o *GetTeamRequest) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

func (o *GetTeamRequest) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

type GetTeamResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The requested team
	Team *shared.Team
}

func (o *GetTeamResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetTeamResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetTeamResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetTeamResponse) GetTeam() *shared.Team {
	if o == nil {
		return nil
	}
	return o.Team
}
