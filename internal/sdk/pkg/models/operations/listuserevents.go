// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"vercel/internal/sdk/pkg/models/shared"
)

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

func (o *ListUserEventsRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListUserEventsRequest) GetSince() *string {
	if o == nil {
		return nil
	}
	return o.Since
}

func (o *ListUserEventsRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

func (o *ListUserEventsRequest) GetTypes() *string {
	if o == nil {
		return nil
	}
	return o.Types
}

func (o *ListUserEventsRequest) GetUntil() *string {
	if o == nil {
		return nil
	}
	return o.Until
}

func (o *ListUserEventsRequest) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}

// ListUserEvents200ApplicationJSON - Successful response.
type ListUserEvents200ApplicationJSON struct {
	// Array of events generated by the User.
	Events []shared.UserEvent `json:"events"`
}

func (o *ListUserEvents200ApplicationJSON) GetEvents() []shared.UserEvent {
	if o == nil {
		return []shared.UserEvent{}
	}
	return o.Events
}

type ListUserEventsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful response.
	ListUserEvents200ApplicationJSONObject *ListUserEvents200ApplicationJSON
}

func (o *ListUserEventsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListUserEventsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListUserEventsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListUserEventsResponse) GetListUserEvents200ApplicationJSONObject() *ListUserEvents200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ListUserEvents200ApplicationJSONObject
}
