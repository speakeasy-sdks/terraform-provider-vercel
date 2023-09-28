// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DownloadArtifactRequest struct {
	// The artifact hash
	Hash string `pathParam:"style=simple,explode=false,name=hash"`
	// The Team identifier or slug to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
	// The continuous integration or delivery environment where this artifact is downloaded.
	XArtifactClientCi *string `header:"style=simple,explode=false,name=x-artifact-client-ci"`
	// 1 if the client is an interactive shell. Otherwise 0
	XArtifactClientInteractive *int64 `header:"style=simple,explode=false,name=x-artifact-client-interactive"`
}

type DownloadArtifactResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The artifact was found and is downloaded as a stream. Content-Length should be verified.
	DownloadArtifact200ApplicationJSONBinaryString []byte
}
