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

func (o *UploadCertRequestBody) GetCa() string {
	if o == nil {
		return ""
	}
	return o.Ca
}

func (o *UploadCertRequestBody) GetCert() string {
	if o == nil {
		return ""
	}
	return o.Cert
}

func (o *UploadCertRequestBody) GetKey() string {
	if o == nil {
		return ""
	}
	return o.Key
}

func (o *UploadCertRequestBody) GetSkipValidation() *bool {
	if o == nil {
		return nil
	}
	return o.SkipValidation
}

type UploadCertRequest struct {
	RequestBody *UploadCertRequestBody `request:"mediaType=application/json"`
	// The Team slug to perform the request on behalf of.
	Slug *string `queryParam:"style=form,explode=true,name=slug"`
	// The Team identifier to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
}

func (o *UploadCertRequest) GetRequestBody() *UploadCertRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *UploadCertRequest) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *UploadCertRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

type UploadCertResponseBody struct {
	AutoRenew bool     `json:"autoRenew"`
	Cns       []string `json:"cns"`
	CreatedAt float64  `json:"createdAt"`
	ExpiresAt float64  `json:"expiresAt"`
	ID        string   `json:"id"`
}

func (o *UploadCertResponseBody) GetAutoRenew() bool {
	if o == nil {
		return false
	}
	return o.AutoRenew
}

func (o *UploadCertResponseBody) GetCns() []string {
	if o == nil {
		return []string{}
	}
	return o.Cns
}

func (o *UploadCertResponseBody) GetCreatedAt() float64 {
	if o == nil {
		return 0.0
	}
	return o.CreatedAt
}

func (o *UploadCertResponseBody) GetExpiresAt() float64 {
	if o == nil {
		return 0.0
	}
	return o.ExpiresAt
}

func (o *UploadCertResponseBody) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type UploadCertResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	Object      *UploadCertResponseBody
}

func (o *UploadCertResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UploadCertResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UploadCertResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UploadCertResponse) GetObject() *UploadCertResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
