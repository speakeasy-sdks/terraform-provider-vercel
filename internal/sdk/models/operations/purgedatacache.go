// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type PurgeDataCacheRequest struct {
	ProjectIDOrName string `queryParam:"style=form,explode=true,name=projectIdOrName"`
}

func (o *PurgeDataCacheRequest) GetProjectIDOrName() string {
	if o == nil {
		return ""
	}
	return o.ProjectIDOrName
}

type PurgeDataCacheResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PurgeDataCacheResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PurgeDataCacheResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PurgeDataCacheResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
