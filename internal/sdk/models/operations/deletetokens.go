// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteTokensRequestBody struct {
	Tokens []string `json:"tokens"`
}

func (o *DeleteTokensRequestBody) GetTokens() []string {
	if o == nil {
		return []string{}
	}
	return o.Tokens
}

type DeleteTokensRequest struct {
	EdgeConfigID string `pathParam:"style=simple,explode=false,name=edgeConfigId"`
	// The Team identifier to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
	// The Team slug to perform the request on behalf of.
	Slug        *string                  `queryParam:"style=form,explode=true,name=slug"`
	RequestBody *DeleteTokensRequestBody `request:"mediaType=application/json"`
}

func (o *DeleteTokensRequest) GetEdgeConfigID() string {
	if o == nil {
		return ""
	}
	return o.EdgeConfigID
}

func (o *DeleteTokensRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

func (o *DeleteTokensRequest) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *DeleteTokensRequest) GetRequestBody() *DeleteTokensRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

type DeleteTokensResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteTokensResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteTokensResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteTokensResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
