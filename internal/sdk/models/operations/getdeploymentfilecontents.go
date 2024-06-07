// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetDeploymentFileContentsRequest struct {
	// The unique deployment identifier
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// The unique file identifier
	FileID string `pathParam:"style=simple,explode=false,name=fileId"`
	// Path to the file to fetch (only for Git deployments)
	Path *string `queryParam:"style=form,explode=true,name=path"`
	// The Team identifier to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
	// The Team slug to perform the request on behalf of.
	Slug *string `queryParam:"style=form,explode=true,name=slug"`
}

func (o *GetDeploymentFileContentsRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetDeploymentFileContentsRequest) GetFileID() string {
	if o == nil {
		return ""
	}
	return o.FileID
}

func (o *GetDeploymentFileContentsRequest) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

func (o *GetDeploymentFileContentsRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

func (o *GetDeploymentFileContentsRequest) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

type GetDeploymentFileContentsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetDeploymentFileContentsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetDeploymentFileContentsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetDeploymentFileContentsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
