// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type CheckDomainStatusRequest struct {
	// The name of the domain for which we would like to check the status.
	Name string `queryParam:"style=form,explode=true,name=name"`
	// The Team identifier or slug to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
}

func (o *CheckDomainStatusRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CheckDomainStatusRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

// CheckDomainStatusResponseBody - Successful response checking if a Domain's name is available.
type CheckDomainStatusResponseBody struct {
	Available bool `json:"available"`
}

func (o *CheckDomainStatusResponseBody) GetAvailable() bool {
	if o == nil {
		return false
	}
	return o.Available
}

type CheckDomainStatusResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful response checking if a Domain's name is available.
	Object *CheckDomainStatusResponseBody
}

func (o *CheckDomainStatusResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CheckDomainStatusResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CheckDomainStatusResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CheckDomainStatusResponse) GetObject() *CheckDomainStatusResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
