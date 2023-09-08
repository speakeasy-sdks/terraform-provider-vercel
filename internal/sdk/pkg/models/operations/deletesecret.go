// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteSecretSecurity struct {
	BearerToken string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

type DeleteSecretRequest struct {
	// The name or the unique identifier to which the secret belongs to.
	IDOrName string `pathParam:"style=simple,explode=false,name=idOrName"`
	// The Team identifier or slug to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
}

type DeleteSecret200ApplicationJSON struct {
	// The date when the secret was created.
	Created int64 `json:"created"`
	// The name of the deleted secret.
	Name string `json:"name"`
	// The unique identifier of the deleted secret.
	UID string `json:"uid"`
}

type DeleteSecretResponse struct {
	ContentType                          string
	StatusCode                           int
	RawResponse                          *http.Response
	DeleteSecret200ApplicationJSONObject *DeleteSecret200ApplicationJSON
}
