// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// RedirectStatusCode - Status code for domain redirect
type RedirectStatusCode int64

const (
	RedirectStatusCodeThreeHundredAndOne   RedirectStatusCode = 301
	RedirectStatusCodeThreeHundredAndTwo   RedirectStatusCode = 302
	RedirectStatusCodeThreeHundredAndSeven RedirectStatusCode = 307
	RedirectStatusCodeThreeHundredAndEight RedirectStatusCode = 308
)

func (e RedirectStatusCode) ToPointer() *RedirectStatusCode {
	return &e
}

func (e *RedirectStatusCode) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 301:
		fallthrough
	case 302:
		fallthrough
	case 307:
		fallthrough
	case 308:
		*e = RedirectStatusCode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RedirectStatusCode: %v", v)
	}
}

type AddProjectDomainRequestBody struct {
	// Git branch to link the project domain
	GitBranch *string `json:"gitBranch,omitempty"`
	// The project domain name
	Name string `json:"name"`
	// Target destination domain for redirect
	Redirect *string `json:"redirect,omitempty"`
	// Status code for domain redirect
	RedirectStatusCode *RedirectStatusCode `json:"redirectStatusCode,omitempty"`
}

func (o *AddProjectDomainRequestBody) GetGitBranch() *string {
	if o == nil {
		return nil
	}
	return o.GitBranch
}

func (o *AddProjectDomainRequestBody) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *AddProjectDomainRequestBody) GetRedirect() *string {
	if o == nil {
		return nil
	}
	return o.Redirect
}

func (o *AddProjectDomainRequestBody) GetRedirectStatusCode() *RedirectStatusCode {
	if o == nil {
		return nil
	}
	return o.RedirectStatusCode
}

type AddProjectDomainRequest struct {
	RequestBody *AddProjectDomainRequestBody `request:"mediaType=application/json"`
	// The unique project identifier or the project name
	IDOrName string `pathParam:"style=simple,explode=false,name=idOrName"`
	// The Team identifier or slug to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
}

func (o *AddProjectDomainRequest) GetRequestBody() *AddProjectDomainRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *AddProjectDomainRequest) GetIDOrName() string {
	if o == nil {
		return ""
	}
	return o.IDOrName
}

func (o *AddProjectDomainRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

type AddProjectDomainRedirectStatusCode int64

const (
	AddProjectDomainRedirectStatusCodeThreeHundredAndSeven AddProjectDomainRedirectStatusCode = 307
	AddProjectDomainRedirectStatusCodeThreeHundredAndOne   AddProjectDomainRedirectStatusCode = 301
	AddProjectDomainRedirectStatusCodeThreeHundredAndTwo   AddProjectDomainRedirectStatusCode = 302
	AddProjectDomainRedirectStatusCodeThreeHundredAndEight AddProjectDomainRedirectStatusCode = 308
)

func (e AddProjectDomainRedirectStatusCode) ToPointer() *AddProjectDomainRedirectStatusCode {
	return &e
}

func (e *AddProjectDomainRedirectStatusCode) UnmarshalJSON(data []byte) error {
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
		*e = AddProjectDomainRedirectStatusCode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AddProjectDomainRedirectStatusCode: %v", v)
	}
}

// Verification - A list of verification challenges, one of which must be completed to verify the domain for use on the project. After the challenge is complete `POST /projects/:idOrName/domains/:domain/verify` to verify the domain. Possible challenges: - If `verification.type = TXT` the `verification.domain` will be checked for a TXT record matching `verification.value`.
type Verification struct {
	Domain string `json:"domain"`
	Reason string `json:"reason"`
	Type   string `json:"type"`
	Value  string `json:"value"`
}

func (o *Verification) GetDomain() string {
	if o == nil {
		return ""
	}
	return o.Domain
}

func (o *Verification) GetReason() string {
	if o == nil {
		return ""
	}
	return o.Reason
}

func (o *Verification) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

func (o *Verification) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

// AddProjectDomainResponseBody - The domain was successfully added to the project
type AddProjectDomainResponseBody struct {
	ApexName           string                              `json:"apexName"`
	CreatedAt          *int64                              `json:"createdAt,omitempty"`
	GitBranch          *string                             `json:"gitBranch,omitempty"`
	Name               string                              `json:"name"`
	ProjectID          string                              `json:"projectId"`
	Redirect           *string                             `json:"redirect,omitempty"`
	RedirectStatusCode *AddProjectDomainRedirectStatusCode `json:"redirectStatusCode,omitempty"`
	UpdatedAt          *int64                              `json:"updatedAt,omitempty"`
	// A list of verification challenges, one of which must be completed to verify the domain for use on the project. After the challenge is complete `POST /projects/:idOrName/domains/:domain/verify` to verify the domain. Possible challenges: - If `verification.type = TXT` the `verification.domain` will be checked for a TXT record matching `verification.value`.
	Verification []Verification `json:"verification,omitempty"`
	// `true` if the domain is verified for use with the project. If `false` it will not be used as an alias on this project until the challenge in `verification` is completed.
	Verified bool `json:"verified"`
}

func (o *AddProjectDomainResponseBody) GetApexName() string {
	if o == nil {
		return ""
	}
	return o.ApexName
}

func (o *AddProjectDomainResponseBody) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *AddProjectDomainResponseBody) GetGitBranch() *string {
	if o == nil {
		return nil
	}
	return o.GitBranch
}

func (o *AddProjectDomainResponseBody) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *AddProjectDomainResponseBody) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

func (o *AddProjectDomainResponseBody) GetRedirect() *string {
	if o == nil {
		return nil
	}
	return o.Redirect
}

func (o *AddProjectDomainResponseBody) GetRedirectStatusCode() *AddProjectDomainRedirectStatusCode {
	if o == nil {
		return nil
	}
	return o.RedirectStatusCode
}

func (o *AddProjectDomainResponseBody) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *AddProjectDomainResponseBody) GetVerification() []Verification {
	if o == nil {
		return nil
	}
	return o.Verification
}

func (o *AddProjectDomainResponseBody) GetVerified() bool {
	if o == nil {
		return false
	}
	return o.Verified
}

type AddProjectDomainResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The domain was successfully added to the project
	Object *AddProjectDomainResponseBody
}

func (o *AddProjectDomainResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AddProjectDomainResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AddProjectDomainResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *AddProjectDomainResponse) GetObject() *AddProjectDomainResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
