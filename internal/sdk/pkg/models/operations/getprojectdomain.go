// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GetProjectDomainRequest struct {
	// The project domain name
	Domain string `pathParam:"style=simple,explode=false,name=domain"`
	// The unique project identifier or the project name
	IDOrName string `pathParam:"style=simple,explode=false,name=idOrName"`
	// The Team identifier or slug to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
}

func (o *GetProjectDomainRequest) GetDomain() string {
	if o == nil {
		return ""
	}
	return o.Domain
}

func (o *GetProjectDomainRequest) GetIDOrName() string {
	if o == nil {
		return ""
	}
	return o.IDOrName
}

func (o *GetProjectDomainRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

type GetProjectDomain200ApplicationJSONRedirectStatusCode int64

const (
	GetProjectDomain200ApplicationJSONRedirectStatusCodeThreeHundredAndSeven GetProjectDomain200ApplicationJSONRedirectStatusCode = 307
	GetProjectDomain200ApplicationJSONRedirectStatusCodeThreeHundredAndOne   GetProjectDomain200ApplicationJSONRedirectStatusCode = 301
	GetProjectDomain200ApplicationJSONRedirectStatusCodeThreeHundredAndTwo   GetProjectDomain200ApplicationJSONRedirectStatusCode = 302
	GetProjectDomain200ApplicationJSONRedirectStatusCodeThreeHundredAndEight GetProjectDomain200ApplicationJSONRedirectStatusCode = 308
)

func (e GetProjectDomain200ApplicationJSONRedirectStatusCode) ToPointer() *GetProjectDomain200ApplicationJSONRedirectStatusCode {
	return &e
}

func (e *GetProjectDomain200ApplicationJSONRedirectStatusCode) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 307:
		fallthrough
	case 301:
		fallthrough
	case 302:
		fallthrough
	case 308:
		*e = GetProjectDomain200ApplicationJSONRedirectStatusCode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetProjectDomain200ApplicationJSONRedirectStatusCode: %v", v)
	}
}

// GetProjectDomain200ApplicationJSONVerification - A list of verification challenges, one of which must be completed to verify the domain for use on the project. After the challenge is complete `POST /projects/:idOrName/domains/:domain/verify` to verify the domain. Possible challenges: - If `verification.type = TXT` the `verification.domain` will be checked for a TXT record matching `verification.value`.
type GetProjectDomain200ApplicationJSONVerification struct {
	Domain string `json:"domain"`
	Reason string `json:"reason"`
	Type   string `json:"type"`
	Value  string `json:"value"`
}

func (o *GetProjectDomain200ApplicationJSONVerification) GetDomain() string {
	if o == nil {
		return ""
	}
	return o.Domain
}

func (o *GetProjectDomain200ApplicationJSONVerification) GetReason() string {
	if o == nil {
		return ""
	}
	return o.Reason
}

func (o *GetProjectDomain200ApplicationJSONVerification) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

func (o *GetProjectDomain200ApplicationJSONVerification) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

type GetProjectDomain200ApplicationJSON struct {
	ApexName           string                                                `json:"apexName"`
	CreatedAt          *int64                                                `json:"createdAt,omitempty"`
	GitBranch          *string                                               `json:"gitBranch,omitempty"`
	Name               string                                                `json:"name"`
	ProjectID          string                                                `json:"projectId"`
	Redirect           *string                                               `json:"redirect,omitempty"`
	RedirectStatusCode *GetProjectDomain200ApplicationJSONRedirectStatusCode `json:"redirectStatusCode,omitempty"`
	UpdatedAt          *int64                                                `json:"updatedAt,omitempty"`
	// A list of verification challenges, one of which must be completed to verify the domain for use on the project. After the challenge is complete `POST /projects/:idOrName/domains/:domain/verify` to verify the domain. Possible challenges: - If `verification.type = TXT` the `verification.domain` will be checked for a TXT record matching `verification.value`.
	Verification []GetProjectDomain200ApplicationJSONVerification `json:"verification,omitempty"`
	// `true` if the domain is verified for use with the project. If `false` it will not be used as an alias on this project until the challenge in `verification` is completed.
	Verified bool `json:"verified"`
}

func (o *GetProjectDomain200ApplicationJSON) GetApexName() string {
	if o == nil {
		return ""
	}
	return o.ApexName
}

func (o *GetProjectDomain200ApplicationJSON) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *GetProjectDomain200ApplicationJSON) GetGitBranch() *string {
	if o == nil {
		return nil
	}
	return o.GitBranch
}

func (o *GetProjectDomain200ApplicationJSON) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *GetProjectDomain200ApplicationJSON) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

func (o *GetProjectDomain200ApplicationJSON) GetRedirect() *string {
	if o == nil {
		return nil
	}
	return o.Redirect
}

func (o *GetProjectDomain200ApplicationJSON) GetRedirectStatusCode() *GetProjectDomain200ApplicationJSONRedirectStatusCode {
	if o == nil {
		return nil
	}
	return o.RedirectStatusCode
}

func (o *GetProjectDomain200ApplicationJSON) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *GetProjectDomain200ApplicationJSON) GetVerification() []GetProjectDomain200ApplicationJSONVerification {
	if o == nil {
		return nil
	}
	return o.Verification
}

func (o *GetProjectDomain200ApplicationJSON) GetVerified() bool {
	if o == nil {
		return false
	}
	return o.Verified
}

type GetProjectDomainResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse                              *http.Response
	GetProjectDomain200ApplicationJSONObject *GetProjectDomain200ApplicationJSON
}

func (o *GetProjectDomainResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetProjectDomainResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetProjectDomainResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetProjectDomainResponse) GetGetProjectDomain200ApplicationJSONObject() *GetProjectDomain200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetProjectDomain200ApplicationJSONObject
}
