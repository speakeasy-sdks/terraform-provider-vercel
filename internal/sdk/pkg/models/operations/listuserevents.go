// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"vercel/internal/sdk/pkg/models/shared"
)

type ListUserEventsSecurity struct {
	BearerToken string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

type ListUserEventsRequest struct {
	// Maximum number of items which may be returned.
	Limit *int64 `queryParam:"style=form,explode=true,name=limit"`
	// Timestamp to only include items created since then.
	Since *string `queryParam:"style=form,explode=true,name=since"`
	// The Team identifier or slug to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
	// Comma-delimited list of event \"types\" to filter the results by.
	Types *string `queryParam:"style=form,explode=true,name=types"`
	// Timestamp to only include items created until then.
	Until *string `queryParam:"style=form,explode=true,name=until"`
	// When retrieving events for a Team, the `userId` parameter may be specified to filter events generated by a specific member of the Team.
	UserID *string `queryParam:"style=form,explode=true,name=userId"`
}

// ListUserEvents200ApplicationJSON - Successful response.
type ListUserEvents200ApplicationJSON struct {
	// Array of events generated by the User.
	Events []shared.UserEvent `json:"events"`
}

type ListUserEventsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful response.
	ListUserEvents200ApplicationJSONObject *ListUserEvents200ApplicationJSON
}
