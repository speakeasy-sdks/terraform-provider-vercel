// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteConfigurationRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// The Team identifier or slug to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
}

func (o *DeleteConfigurationRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *DeleteConfigurationRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

type DeleteConfigurationResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteConfigurationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteConfigurationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteConfigurationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
