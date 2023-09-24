// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type UploadCertRequestBody struct {
	// The certificate authority
	Ca string `json:"ca"`
	// The certificate
	Cert string `json:"cert"`
	// The certificate key
	Key string `json:"key"`
	// Skip validation of the certificate
	SkipValidation *bool `json:"skipValidation,omitempty"`
}

type UploadCertRequest struct {
	RequestBody *UploadCertRequestBody `request:"mediaType=application/json"`
	// The Team identifier or slug to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
}

type UploadCert200ApplicationJSON struct {
	AutoRenew bool     `json:"autoRenew"`
	Cns       []string `json:"cns"`
	CreatedAt int64    `json:"createdAt"`
	ExpiresAt int64    `json:"expiresAt"`
	ID        string   `json:"id"`
}

type UploadCertResponse struct {
	ContentType                        string
	StatusCode                         int
	RawResponse                        *http.Response
	UploadCert200ApplicationJSONObject *UploadCert200ApplicationJSON
}
