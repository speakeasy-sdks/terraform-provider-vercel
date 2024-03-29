// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/zchee/terraform-provider-vercel/internal/sdk/models/shared"
	"net/http"
)

type GetEdgeConfigTokensRequest struct {
	EdgeConfigID string `pathParam:"style=simple,explode=false,name=edgeConfigId"`
	// The Team identifier or slug to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
}

func (o *GetEdgeConfigTokensRequest) GetEdgeConfigID() string {
	if o == nil {
		return ""
	}
	return o.EdgeConfigID
}

func (o *GetEdgeConfigTokensRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

type GetEdgeConfigTokensResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// The EdgeConfig.
	EdgeConfigToken *shared.EdgeConfigToken
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetEdgeConfigTokensResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetEdgeConfigTokensResponse) GetEdgeConfigToken() *shared.EdgeConfigToken {
	if o == nil {
		return nil
	}
	return o.EdgeConfigToken
}

func (o *GetEdgeConfigTokensResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetEdgeConfigTokensResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
