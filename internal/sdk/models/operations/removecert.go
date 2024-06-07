// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type RemoveCertRequest struct {
	// The cert id to remove
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// The Team slug to perform the request on behalf of.
	Slug *string `queryParam:"style=form,explode=true,name=slug"`
	// The Team identifier to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
}

func (o *RemoveCertRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *RemoveCertRequest) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *RemoveCertRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

type RemoveCertResponseBody struct {
}

type RemoveCertResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	Object      *RemoveCertResponseBody
}

func (o *RemoveCertResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RemoveCertResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RemoveCertResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *RemoveCertResponse) GetObject() *RemoveCertResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
