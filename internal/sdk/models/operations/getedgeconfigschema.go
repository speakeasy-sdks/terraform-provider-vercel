// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetEdgeConfigSchemaRequest struct {
	EdgeConfigID string `pathParam:"style=simple,explode=false,name=edgeConfigId"`
	// The Team identifier to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
	// The Team slug to perform the request on behalf of.
	Slug *string `queryParam:"style=form,explode=true,name=slug"`
}

func (o *GetEdgeConfigSchemaRequest) GetEdgeConfigID() string {
	if o == nil {
		return ""
	}
	return o.EdgeConfigID
}

func (o *GetEdgeConfigSchemaRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

func (o *GetEdgeConfigSchemaRequest) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

// GetEdgeConfigSchemaResponseBody - The EdgeConfig.
type GetEdgeConfigSchemaResponseBody struct {
}

type GetEdgeConfigSchemaResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The EdgeConfig.
	Object *GetEdgeConfigSchemaResponseBody
}

func (o *GetEdgeConfigSchemaResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetEdgeConfigSchemaResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetEdgeConfigSchemaResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetEdgeConfigSchemaResponse) GetObject() *GetEdgeConfigSchemaResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
