// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"vercel/internal/sdk/pkg/models/shared"
)

type GetTeamSecurity struct {
	BearerToken string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

type GetTeamRequest struct {
	Slug *string `queryParam:"style=form,explode=true,name=slug"`
	// The Team identifier or slug to perform the request on behalf of.
	TeamID string `pathParam:"style=simple,explode=false,name=teamId"`
}

type GetTeamResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// The requested team
	Team *shared.Team
}
