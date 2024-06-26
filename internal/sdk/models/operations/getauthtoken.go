// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/vercel/terraform-provider-vercel/internal/sdk/models/shared"
	"net/http"
)

type GetAuthTokenRequest struct {
	// The identifier of the token to retrieve. The special value \"current\" may be supplied, which returns the metadata for the token that the current HTTP request is authenticated with.
	TokenID string `pathParam:"style=simple,explode=false,name=tokenId"`
}

func (o *GetAuthTokenRequest) GetTokenID() string {
	if o == nil {
		return ""
	}
	return o.TokenID
}

// GetAuthTokenResponseBody - Successful response.
type GetAuthTokenResponseBody struct {
	// Authentication token metadata.
	Token shared.AuthToken `json:"token"`
}

func (o *GetAuthTokenResponseBody) GetToken() shared.AuthToken {
	if o == nil {
		return shared.AuthToken{}
	}
	return o.Token
}

type GetAuthTokenResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful response.
	Object *GetAuthTokenResponseBody
}

func (o *GetAuthTokenResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAuthTokenResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAuthTokenResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetAuthTokenResponse) GetObject() *GetAuthTokenResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
