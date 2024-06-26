// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DownloadArtifactRequest struct {
	// The continuous integration or delivery environment where this artifact is downloaded.
	XArtifactClientCi *string `header:"style=simple,explode=false,name=x-artifact-client-ci"`
	// 1 if the client is an interactive shell. Otherwise 0
	XArtifactClientInteractive *int64 `header:"style=simple,explode=false,name=x-artifact-client-interactive"`
	// The artifact hash
	Hash string `pathParam:"style=simple,explode=false,name=hash"`
	// The Team identifier to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
	// The Team slug to perform the request on behalf of.
	Slug *string `queryParam:"style=form,explode=true,name=slug"`
}

func (o *DownloadArtifactRequest) GetXArtifactClientCi() *string {
	if o == nil {
		return nil
	}
	return o.XArtifactClientCi
}

func (o *DownloadArtifactRequest) GetXArtifactClientInteractive() *int64 {
	if o == nil {
		return nil
	}
	return o.XArtifactClientInteractive
}

func (o *DownloadArtifactRequest) GetHash() string {
	if o == nil {
		return ""
	}
	return o.Hash
}

func (o *DownloadArtifactRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

func (o *DownloadArtifactRequest) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

type DownloadArtifactResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The artifact was found and is downloaded as a stream. Content-Length should be verified.
	Bytes []byte
}

func (o *DownloadArtifactResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DownloadArtifactResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DownloadArtifactResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DownloadArtifactResponse) GetBytes() []byte {
	if o == nil {
		return nil
	}
	return o.Bytes
}
