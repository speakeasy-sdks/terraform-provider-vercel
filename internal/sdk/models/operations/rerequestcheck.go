// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type RerequestCheckRequest struct {
	// The deployment to rerun the check for.
	DeploymentID string `pathParam:"style=simple,explode=false,name=deploymentId"`
	// The check to rerun
	CheckID string `pathParam:"style=simple,explode=false,name=checkId"`
	// The Team identifier to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
	// The Team slug to perform the request on behalf of.
	Slug *string `queryParam:"style=form,explode=true,name=slug"`
}

func (o *RerequestCheckRequest) GetDeploymentID() string {
	if o == nil {
		return ""
	}
	return o.DeploymentID
}

func (o *RerequestCheckRequest) GetCheckID() string {
	if o == nil {
		return ""
	}
	return o.CheckID
}

func (o *RerequestCheckRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

func (o *RerequestCheckRequest) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

type RerequestCheckResponseBody struct {
}

type RerequestCheckResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	Object      *RerequestCheckResponseBody
}

func (o *RerequestCheckResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RerequestCheckResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RerequestCheckResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *RerequestCheckResponse) GetObject() *RerequestCheckResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
