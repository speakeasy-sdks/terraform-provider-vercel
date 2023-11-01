// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"vercel/internal/sdk/pkg/models/shared"
)

type GetEdgeConfigTokenRequest struct {
	EdgeConfigID string `pathParam:"style=simple,explode=false,name=edgeConfigId"`
	// The Team identifier or slug to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
	Token  string  `pathParam:"style=simple,explode=false,name=token"`
}

func (o *GetEdgeConfigTokenRequest) GetEdgeConfigID() string {
	if o == nil {
		return ""
	}
	return o.EdgeConfigID
}

func (o *GetEdgeConfigTokenRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

func (o *GetEdgeConfigTokenRequest) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

type GetEdgeConfigTokenResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// The EdgeConfig.
	EdgeConfigToken *shared.EdgeConfigToken
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetEdgeConfigTokenResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetEdgeConfigTokenResponse) GetEdgeConfigToken() *shared.EdgeConfigToken {
	if o == nil {
		return nil
	}
	return o.EdgeConfigToken
}

func (o *GetEdgeConfigTokenResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetEdgeConfigTokenResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
