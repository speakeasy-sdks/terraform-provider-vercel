// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type VerifyProjectDomainRequest struct {
	// The domain name you want to verify
	Domain string `pathParam:"style=simple,explode=false,name=domain"`
	// The unique project identifier or the project name
	IDOrName string `pathParam:"style=simple,explode=false,name=idOrName"`
	// The Team identifier or slug to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
}

func (o *VerifyProjectDomainRequest) GetDomain() string {
	if o == nil {
		return ""
	}
	return o.Domain
}

func (o *VerifyProjectDomainRequest) GetIDOrName() string {
	if o == nil {
		return ""
	}
	return o.IDOrName
}

func (o *VerifyProjectDomainRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

type VerifyProjectDomain200ApplicationJSONRedirectStatusCode int64

const (
	VerifyProjectDomain200ApplicationJSONRedirectStatusCodeThreeHundredAndSeven VerifyProjectDomain200ApplicationJSONRedirectStatusCode = 307
	VerifyProjectDomain200ApplicationJSONRedirectStatusCodeThreeHundredAndOne   VerifyProjectDomain200ApplicationJSONRedirectStatusCode = 301
	VerifyProjectDomain200ApplicationJSONRedirectStatusCodeThreeHundredAndTwo   VerifyProjectDomain200ApplicationJSONRedirectStatusCode = 302
	VerifyProjectDomain200ApplicationJSONRedirectStatusCodeThreeHundredAndEight VerifyProjectDomain200ApplicationJSONRedirectStatusCode = 308
)

func (e VerifyProjectDomain200ApplicationJSONRedirectStatusCode) ToPointer() *VerifyProjectDomain200ApplicationJSONRedirectStatusCode {
	return &e
}

func (e *VerifyProjectDomain200ApplicationJSONRedirectStatusCode) UnmarshalJSON(data []byte) error {
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
		*e = VerifyProjectDomain200ApplicationJSONRedirectStatusCode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for VerifyProjectDomain200ApplicationJSONRedirectStatusCode: %v", v)
	}
}

// VerifyProjectDomain200ApplicationJSONVerification - A list of verification challenges, one of which must be completed to verify the domain for use on the project. After the challenge is complete `POST /projects/:idOrName/domains/:domain/verify` to verify the domain. Possible challenges: - If `verification.type = TXT` the `verification.domain` will be checked for a TXT record matching `verification.value`.
type VerifyProjectDomain200ApplicationJSONVerification struct {
	Domain string `json:"domain"`
	Reason string `json:"reason"`
	Type   string `json:"type"`
	Value  string `json:"value"`
}

func (o *VerifyProjectDomain200ApplicationJSONVerification) GetDomain() string {
	if o == nil {
		return ""
	}
	return o.Domain
}

func (o *VerifyProjectDomain200ApplicationJSONVerification) GetReason() string {
	if o == nil {
		return ""
	}
	return o.Reason
}

func (o *VerifyProjectDomain200ApplicationJSONVerification) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

func (o *VerifyProjectDomain200ApplicationJSONVerification) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

// VerifyProjectDomain200ApplicationJSON - The project domain was verified successfully
// Domain is already verified
type VerifyProjectDomain200ApplicationJSON struct {
	ApexName           string                                                   `json:"apexName"`
	CreatedAt          *int64                                                   `json:"createdAt,omitempty"`
	GitBranch          *string                                                  `json:"gitBranch,omitempty"`
	Name               string                                                   `json:"name"`
	ProjectID          string                                                   `json:"projectId"`
	Redirect           *string                                                  `json:"redirect,omitempty"`
	RedirectStatusCode *VerifyProjectDomain200ApplicationJSONRedirectStatusCode `json:"redirectStatusCode,omitempty"`
	UpdatedAt          *int64                                                   `json:"updatedAt,omitempty"`
	// A list of verification challenges, one of which must be completed to verify the domain for use on the project. After the challenge is complete `POST /projects/:idOrName/domains/:domain/verify` to verify the domain. Possible challenges: - If `verification.type = TXT` the `verification.domain` will be checked for a TXT record matching `verification.value`.
	Verification []VerifyProjectDomain200ApplicationJSONVerification `json:"verification,omitempty"`
	// `true` if the domain is verified for use with the project. If `false` it will not be used as an alias on this project until the challenge in `verification` is completed.
	Verified bool `json:"verified"`
}

func (o *VerifyProjectDomain200ApplicationJSON) GetApexName() string {
	if o == nil {
		return ""
	}
	return o.ApexName
}

func (o *VerifyProjectDomain200ApplicationJSON) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *VerifyProjectDomain200ApplicationJSON) GetGitBranch() *string {
	if o == nil {
		return nil
	}
	return o.GitBranch
}

func (o *VerifyProjectDomain200ApplicationJSON) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *VerifyProjectDomain200ApplicationJSON) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

func (o *VerifyProjectDomain200ApplicationJSON) GetRedirect() *string {
	if o == nil {
		return nil
	}
	return o.Redirect
}

func (o *VerifyProjectDomain200ApplicationJSON) GetRedirectStatusCode() *VerifyProjectDomain200ApplicationJSONRedirectStatusCode {
	if o == nil {
		return nil
	}
	return o.RedirectStatusCode
}

func (o *VerifyProjectDomain200ApplicationJSON) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *VerifyProjectDomain200ApplicationJSON) GetVerification() []VerifyProjectDomain200ApplicationJSONVerification {
	if o == nil {
		return nil
	}
	return o.Verification
}

func (o *VerifyProjectDomain200ApplicationJSON) GetVerified() bool {
	if o == nil {
		return false
	}
	return o.Verified
}

type VerifyProjectDomainResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The project domain was verified successfully
	// Domain is already verified
	VerifyProjectDomain200ApplicationJSONObject *VerifyProjectDomain200ApplicationJSON
}

func (o *VerifyProjectDomainResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *VerifyProjectDomainResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *VerifyProjectDomainResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *VerifyProjectDomainResponse) GetVerifyProjectDomain200ApplicationJSONObject() *VerifyProjectDomain200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.VerifyProjectDomain200ApplicationJSONObject
}
