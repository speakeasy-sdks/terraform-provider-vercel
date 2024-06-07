// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// UpdateProjectDomainRedirectStatusCode - Status code for domain redirect
type UpdateProjectDomainRedirectStatusCode int64

const (
	UpdateProjectDomainRedirectStatusCodeThreeHundredAndOne   UpdateProjectDomainRedirectStatusCode = 301
	UpdateProjectDomainRedirectStatusCodeThreeHundredAndTwo   UpdateProjectDomainRedirectStatusCode = 302
	UpdateProjectDomainRedirectStatusCodeThreeHundredAndSeven UpdateProjectDomainRedirectStatusCode = 307
	UpdateProjectDomainRedirectStatusCodeThreeHundredAndEight UpdateProjectDomainRedirectStatusCode = 308
)

func (e UpdateProjectDomainRedirectStatusCode) ToPointer() *UpdateProjectDomainRedirectStatusCode {
	return &e
}
func (e *UpdateProjectDomainRedirectStatusCode) UnmarshalJSON(data []byte) error {
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
		*e = UpdateProjectDomainRedirectStatusCode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateProjectDomainRedirectStatusCode: %v", v)
	}
}

type UpdateProjectDomainRequestBody struct {
	// The unique custom environment identifier within the project
	CustomEnvironmentID *string `json:"customEnvironmentId,omitempty"`
	// Git branch to link the project domain
	GitBranch *string `json:"gitBranch,omitempty"`
	// Target destination domain for redirect
	Redirect *string `json:"redirect,omitempty"`
	// Status code for domain redirect
	RedirectStatusCode *UpdateProjectDomainRedirectStatusCode `json:"redirectStatusCode,omitempty"`
}

func (o *UpdateProjectDomainRequestBody) GetCustomEnvironmentID() *string {
	if o == nil {
		return nil
	}
	return o.CustomEnvironmentID
}

func (o *UpdateProjectDomainRequestBody) GetGitBranch() *string {
	if o == nil {
		return nil
	}
	return o.GitBranch
}

func (o *UpdateProjectDomainRequestBody) GetRedirect() *string {
	if o == nil {
		return nil
	}
	return o.Redirect
}

func (o *UpdateProjectDomainRequestBody) GetRedirectStatusCode() *UpdateProjectDomainRedirectStatusCode {
	if o == nil {
		return nil
	}
	return o.RedirectStatusCode
}

type UpdateProjectDomainRequest struct {
	RequestBody *UpdateProjectDomainRequestBody `request:"mediaType=application/json"`
	// The project domain name
	Domain string `pathParam:"style=simple,explode=false,name=domain"`
	// The unique project identifier or the project name
	IDOrName string `pathParam:"style=simple,explode=false,name=idOrName"`
	// The Team slug to perform the request on behalf of.
	Slug *string `queryParam:"style=form,explode=true,name=slug"`
	// The Team identifier to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
}

func (o *UpdateProjectDomainRequest) GetRequestBody() *UpdateProjectDomainRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *UpdateProjectDomainRequest) GetDomain() string {
	if o == nil {
		return ""
	}
	return o.Domain
}

func (o *UpdateProjectDomainRequest) GetIDOrName() string {
	if o == nil {
		return ""
	}
	return o.IDOrName
}

func (o *UpdateProjectDomainRequest) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *UpdateProjectDomainRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

// UpdateProjectDomainVerification - A list of verification challenges, one of which must be completed to verify the domain for use on the project. After the challenge is complete `POST /projects/:idOrName/domains/:domain/verify` to verify the domain. Possible challenges: - If `verification.type = TXT` the `verification.domain` will be checked for a TXT record matching `verification.value`.
type UpdateProjectDomainVerification struct {
	Domain string `json:"domain"`
	Reason string `json:"reason"`
	Type   string `json:"type"`
	Value  string `json:"value"`
}

func (o *UpdateProjectDomainVerification) GetDomain() string {
	if o == nil {
		return ""
	}
	return o.Domain
}

func (o *UpdateProjectDomainVerification) GetReason() string {
	if o == nil {
		return ""
	}
	return o.Reason
}

func (o *UpdateProjectDomainVerification) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

func (o *UpdateProjectDomainVerification) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

// UpdateProjectDomainResponseBody - The domain was updated successfuly
type UpdateProjectDomainResponseBody struct {
	ApexName            string   `json:"apexName"`
	CreatedAt           *float64 `json:"createdAt,omitempty"`
	CustomEnvironmentID *string  `json:"customEnvironmentId,omitempty"`
	GitBranch           *string  `json:"gitBranch,omitempty"`
	Name                string   `json:"name"`
	ProjectID           string   `json:"projectId"`
	Redirect            *string  `json:"redirect,omitempty"`
	RedirectStatusCode  *float64 `json:"redirectStatusCode,omitempty"`
	UpdatedAt           *float64 `json:"updatedAt,omitempty"`
	// A list of verification challenges, one of which must be completed to verify the domain for use on the project. After the challenge is complete `POST /projects/:idOrName/domains/:domain/verify` to verify the domain. Possible challenges: - If `verification.type = TXT` the `verification.domain` will be checked for a TXT record matching `verification.value`.
	Verification []UpdateProjectDomainVerification `json:"verification,omitempty"`
	// `true` if the domain is verified for use with the project. If `false` it will not be used as an alias on this project until the challenge in `verification` is completed.
	Verified bool `json:"verified"`
}

func (o *UpdateProjectDomainResponseBody) GetApexName() string {
	if o == nil {
		return ""
	}
	return o.ApexName
}

func (o *UpdateProjectDomainResponseBody) GetCreatedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *UpdateProjectDomainResponseBody) GetCustomEnvironmentID() *string {
	if o == nil {
		return nil
	}
	return o.CustomEnvironmentID
}

func (o *UpdateProjectDomainResponseBody) GetGitBranch() *string {
	if o == nil {
		return nil
	}
	return o.GitBranch
}

func (o *UpdateProjectDomainResponseBody) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *UpdateProjectDomainResponseBody) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

func (o *UpdateProjectDomainResponseBody) GetRedirect() *string {
	if o == nil {
		return nil
	}
	return o.Redirect
}

func (o *UpdateProjectDomainResponseBody) GetRedirectStatusCode() *float64 {
	if o == nil {
		return nil
	}
	return o.RedirectStatusCode
}

func (o *UpdateProjectDomainResponseBody) GetUpdatedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *UpdateProjectDomainResponseBody) GetVerification() []UpdateProjectDomainVerification {
	if o == nil {
		return nil
	}
	return o.Verification
}

func (o *UpdateProjectDomainResponseBody) GetVerified() bool {
	if o == nil {
		return false
	}
	return o.Verified
}

type UpdateProjectDomainResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The domain was updated successfuly
	Object *UpdateProjectDomainResponseBody
}

func (o *UpdateProjectDomainResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateProjectDomainResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateProjectDomainResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateProjectDomainResponse) GetObject() *UpdateProjectDomainResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
