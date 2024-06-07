// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteEdgeConfigRequest struct {
	EdgeConfigID string `pathParam:"style=simple,explode=false,name=edgeConfigId"`
	// The Team identifier to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
	// The Team slug to perform the request on behalf of.
	Slug *string `queryParam:"style=form,explode=true,name=slug"`
}

func (o *DeleteEdgeConfigRequest) GetEdgeConfigID() string {
	if o == nil {
		return ""
	}
	return o.EdgeConfigID
}

func (o *DeleteEdgeConfigRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

func (o *DeleteEdgeConfigRequest) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

type DeleteEdgeConfigResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteEdgeConfigResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteEdgeConfigResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteEdgeConfigResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
