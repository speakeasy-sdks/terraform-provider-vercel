// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type EmailLoginRequestBody struct {
	// The user email.
	Email string `json:"email"`
	// The desired name for the token. It will be displayed on the user account details.
	TokenName *string `json:"tokenName,omitempty"`
}

func (o *EmailLoginRequestBody) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

func (o *EmailLoginRequestBody) GetTokenName() *string {
	if o == nil {
		return nil
	}
	return o.TokenName
}

// EmailLoginResponseBody - The request was successful and an email was sent
type EmailLoginResponseBody struct {
	// The code the user is going to receive on the email. **Must** be displayed to the user so they can verify the request is the correct.
	SecurityCode string `json:"securityCode"`
	// The token used to verify the user accepted the login request
	Token string `json:"token"`
}

func (o *EmailLoginResponseBody) GetSecurityCode() string {
	if o == nil {
		return ""
	}
	return o.SecurityCode
}

func (o *EmailLoginResponseBody) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

type EmailLoginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The request was successful and an email was sent
	Object *EmailLoginResponseBody
}

func (o *EmailLoginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *EmailLoginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *EmailLoginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *EmailLoginResponse) GetObject() *EmailLoginResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}