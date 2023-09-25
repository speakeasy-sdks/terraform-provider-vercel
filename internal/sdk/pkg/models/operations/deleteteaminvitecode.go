// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteTeamInviteCodeRequest struct {
	// The Team invite code ID.
	InviteID string `pathParam:"style=simple,explode=false,name=inviteId"`
	// The Team identifier or slug to perform the request on behalf of.
	TeamID string `pathParam:"style=simple,explode=false,name=teamId"`
}

// DeleteTeamInviteCode200ApplicationJSON - Successfully deleted Team invite code.
type DeleteTeamInviteCode200ApplicationJSON struct {
	// ID of the team.
	ID string `json:"id"`
}

type DeleteTeamInviteCodeResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successfully deleted Team invite code.
	DeleteTeamInviteCode200ApplicationJSONObject *DeleteTeamInviteCode200ApplicationJSON
}
