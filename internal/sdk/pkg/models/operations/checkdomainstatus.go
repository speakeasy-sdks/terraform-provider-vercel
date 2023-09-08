// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type CheckDomainStatusSecurity struct {
	BearerToken string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

type CheckDomainStatusRequest struct {
	// The name of the domain for which we would like to check the status.
	Name string `queryParam:"style=form,explode=true,name=name"`
	// The Team identifier or slug to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
}

// CheckDomainStatus200ApplicationJSON - Successful response checking if a Domain's name is available.
type CheckDomainStatus200ApplicationJSON struct {
	Available bool `json:"available"`
}

type CheckDomainStatusResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful response checking if a Domain's name is available.
	CheckDomainStatus200ApplicationJSONObject *CheckDomainStatus200ApplicationJSON
}
