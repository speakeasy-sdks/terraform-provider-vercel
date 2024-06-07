// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type StatusRequest struct {
	// The Team slug to perform the request on behalf of.
	Slug *string `queryParam:"style=form,explode=true,name=slug"`
	// The Team identifier to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
}

func (o *StatusRequest) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *StatusRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

type StatusStatus string

const (
	StatusStatusDisabled  StatusStatus = "disabled"
	StatusStatusEnabled   StatusStatus = "enabled"
	StatusStatusOverLimit StatusStatus = "over_limit"
	StatusStatusPaused    StatusStatus = "paused"
)

func (e StatusStatus) ToPointer() *StatusStatus {
	return &e
}
func (e *StatusStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "disabled":
		fallthrough
	case "enabled":
		fallthrough
	case "over_limit":
		fallthrough
	case "paused":
		*e = StatusStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for StatusStatus: %v", v)
	}
}

type StatusResponseBody struct {
	Status StatusStatus `json:"status"`
}

func (o *StatusResponseBody) GetStatus() StatusStatus {
	if o == nil {
		return StatusStatus("")
	}
	return o.Status
}

type StatusResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	Object      *StatusResponseBody
}

func (o *StatusResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *StatusResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *StatusResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *StatusResponse) GetObject() *StatusResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
