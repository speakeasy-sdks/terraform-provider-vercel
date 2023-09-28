// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"vercel/internal/sdk/pkg/models/shared"
)

type ListDeploymentFilesRequest struct {
	// The unique deployment identifier
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// The Team identifier or slug to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
}

type ListDeploymentFilesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Retrieved the file tree successfully
	FileTrees []shared.FileTree
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}
