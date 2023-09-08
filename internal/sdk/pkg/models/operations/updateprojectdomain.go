// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type UpdateProjectDomainSecurity struct {
	BearerToken string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

// UpdateProjectDomainRequestBodyRedirectStatusCode - Status code for domain redirect
type UpdateProjectDomainRequestBodyRedirectStatusCode int64

const (
	UpdateProjectDomainRequestBodyRedirectStatusCodeThreeHundredAndOne   UpdateProjectDomainRequestBodyRedirectStatusCode = 301
	UpdateProjectDomainRequestBodyRedirectStatusCodeThreeHundredAndTwo   UpdateProjectDomainRequestBodyRedirectStatusCode = 302
	UpdateProjectDomainRequestBodyRedirectStatusCodeThreeHundredAndSeven UpdateProjectDomainRequestBodyRedirectStatusCode = 307
	UpdateProjectDomainRequestBodyRedirectStatusCodeThreeHundredAndEight UpdateProjectDomainRequestBodyRedirectStatusCode = 308
)

func (e UpdateProjectDomainRequestBodyRedirectStatusCode) ToPointer() *UpdateProjectDomainRequestBodyRedirectStatusCode {
	return &e
}

func (e *UpdateProjectDomainRequestBodyRedirectStatusCode) UnmarshalJSON(data []byte) error {
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
		*e = UpdateProjectDomainRequestBodyRedirectStatusCode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateProjectDomainRequestBodyRedirectStatusCode: %v", v)
	}
}

type UpdateProjectDomainRequestBody struct {
	// Git branch to link the project domain
	GitBranch *string `json:"gitBranch,omitempty"`
	// Target destination domain for redirect
	Redirect *string `json:"redirect,omitempty"`
	// Status code for domain redirect
	RedirectStatusCode *UpdateProjectDomainRequestBodyRedirectStatusCode `json:"redirectStatusCode,omitempty"`
}

type UpdateProjectDomainRequest struct {
	RequestBody *UpdateProjectDomainRequestBody `request:"mediaType=application/json"`
	// The project domain name
	Domain string `pathParam:"style=simple,explode=false,name=domain"`
	// The unique project identifier or the project name
	IDOrName string `pathParam:"style=simple,explode=false,name=idOrName"`
	// The Team identifier or slug to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
}

type UpdateProjectDomain200ApplicationJSONRedirectStatusCode int64

const (
	UpdateProjectDomain200ApplicationJSONRedirectStatusCodeThreeHundredAndSeven UpdateProjectDomain200ApplicationJSONRedirectStatusCode = 307
	UpdateProjectDomain200ApplicationJSONRedirectStatusCodeThreeHundredAndOne   UpdateProjectDomain200ApplicationJSONRedirectStatusCode = 301
	UpdateProjectDomain200ApplicationJSONRedirectStatusCodeThreeHundredAndTwo   UpdateProjectDomain200ApplicationJSONRedirectStatusCode = 302
	UpdateProjectDomain200ApplicationJSONRedirectStatusCodeThreeHundredAndEight UpdateProjectDomain200ApplicationJSONRedirectStatusCode = 308
)

func (e UpdateProjectDomain200ApplicationJSONRedirectStatusCode) ToPointer() *UpdateProjectDomain200ApplicationJSONRedirectStatusCode {
	return &e
}

func (e *UpdateProjectDomain200ApplicationJSONRedirectStatusCode) UnmarshalJSON(data []byte) error {
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
		*e = UpdateProjectDomain200ApplicationJSONRedirectStatusCode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateProjectDomain200ApplicationJSONRedirectStatusCode: %v", v)
	}
}

// UpdateProjectDomain200ApplicationJSONVerification - A list of verification challenges, one of which must be completed to verify the domain for use on the project. After the challenge is complete `POST /projects/:idOrName/domains/:domain/verify` to verify the domain. Possible challenges: - If `verification.type = TXT` the `verification.domain` will be checked for a TXT record matching `verification.value`.
type UpdateProjectDomain200ApplicationJSONVerification struct {
	Domain string `json:"domain"`
	Reason string `json:"reason"`
	Type   string `json:"type"`
	Value  string `json:"value"`
}

// UpdateProjectDomain200ApplicationJSON - The domain was updated successfuly
type UpdateProjectDomain200ApplicationJSON struct {
	ApexName           string                                                   `json:"apexName"`
	CreatedAt          *int64                                                   `json:"createdAt,omitempty"`
	GitBranch          *string                                                  `json:"gitBranch,omitempty"`
	Name               string                                                   `json:"name"`
	ProjectID          string                                                   `json:"projectId"`
	Redirect           *string                                                  `json:"redirect,omitempty"`
	RedirectStatusCode *UpdateProjectDomain200ApplicationJSONRedirectStatusCode `json:"redirectStatusCode,omitempty"`
	UpdatedAt          *int64                                                   `json:"updatedAt,omitempty"`
	// A list of verification challenges, one of which must be completed to verify the domain for use on the project. After the challenge is complete `POST /projects/:idOrName/domains/:domain/verify` to verify the domain. Possible challenges: - If `verification.type = TXT` the `verification.domain` will be checked for a TXT record matching `verification.value`.
	Verification []UpdateProjectDomain200ApplicationJSONVerification `json:"verification,omitempty"`
	// `true` if the domain is verified for use with the project. If `false` it will not be used as an alias on this project until the challenge in `verification` is completed.
	Verified bool `json:"verified"`
}

type UpdateProjectDomainResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// The domain was updated successfuly
	UpdateProjectDomain200ApplicationJSONObject *UpdateProjectDomain200ApplicationJSON
}
