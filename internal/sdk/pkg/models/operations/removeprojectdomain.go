// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type RemoveProjectDomainRequest struct {
	// The project domain name
	Domain string `pathParam:"style=simple,explode=false,name=domain"`
	// The unique project identifier or the project name
	IDOrName string `pathParam:"style=simple,explode=false,name=idOrName"`
	// The Team identifier or slug to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
}

func (o *RemoveProjectDomainRequest) GetDomain() string {
	if o == nil {
		return ""
	}
	return o.Domain
}

func (o *RemoveProjectDomainRequest) GetIDOrName() string {
	if o == nil {
		return ""
	}
	return o.IDOrName
}

func (o *RemoveProjectDomainRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

// RemoveProjectDomainResponseBody - The domain was succesfully removed from the project
type RemoveProjectDomainResponseBody struct {
}

type RemoveProjectDomainResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The domain was succesfully removed from the project
	Object *RemoveProjectDomainResponseBody
}

func (o *RemoveProjectDomainResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RemoveProjectDomainResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RemoveProjectDomainResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *RemoveProjectDomainResponse) GetObject() *RemoveProjectDomainResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
