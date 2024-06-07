// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type UnpauseProjectRequest struct {
	// The unique project identifier
	ProjectID string `pathParam:"style=simple,explode=false,name=projectId"`
	// The Team slug to perform the request on behalf of.
	Slug *string `queryParam:"style=form,explode=true,name=slug"`
	// The Team identifier to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
}

func (o *UnpauseProjectRequest) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

func (o *UnpauseProjectRequest) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *UnpauseProjectRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

type UnpauseProjectResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UnpauseProjectResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UnpauseProjectResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UnpauseProjectResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
