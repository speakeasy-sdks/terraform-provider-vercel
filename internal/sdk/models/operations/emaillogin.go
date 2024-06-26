// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"errors"
	"fmt"
	"github.com/vercel/terraform-provider-vercel/internal/sdk/internal/utils"
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

type EmailLoginResponseBody2 struct {
	// The token used to verify the user accepted the login request
	Token string `json:"token"`
	// The code the user is going to receive on the email. **Must** be displayed to the user so they can verify the request is the correct.
	SecurityCode string `json:"securityCode"`
}

func (o *EmailLoginResponseBody2) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

func (o *EmailLoginResponseBody2) GetSecurityCode() string {
	if o == nil {
		return ""
	}
	return o.SecurityCode
}

type EmailLoginResponseBody1 struct {
	Token        string `json:"token"`
	SecurityCode string `json:"securityCode"`
}

func (o *EmailLoginResponseBody1) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

func (o *EmailLoginResponseBody1) GetSecurityCode() string {
	if o == nil {
		return ""
	}
	return o.SecurityCode
}

type EmailLoginResponseBodyType string

const (
	EmailLoginResponseBodyTypeEmailLoginResponseBody1 EmailLoginResponseBodyType = "emailLogin_responseBody_1"
	EmailLoginResponseBodyTypeEmailLoginResponseBody2 EmailLoginResponseBodyType = "emailLogin_responseBody_2"
)

// EmailLoginResponseBody - The request was successful and an email was sent
type EmailLoginResponseBody struct {
	EmailLoginResponseBody1 *EmailLoginResponseBody1
	EmailLoginResponseBody2 *EmailLoginResponseBody2

	Type EmailLoginResponseBodyType
}

func CreateEmailLoginResponseBodyEmailLoginResponseBody1(emailLoginResponseBody1 EmailLoginResponseBody1) EmailLoginResponseBody {
	typ := EmailLoginResponseBodyTypeEmailLoginResponseBody1

	return EmailLoginResponseBody{
		EmailLoginResponseBody1: &emailLoginResponseBody1,
		Type:                    typ,
	}
}

func CreateEmailLoginResponseBodyEmailLoginResponseBody2(emailLoginResponseBody2 EmailLoginResponseBody2) EmailLoginResponseBody {
	typ := EmailLoginResponseBodyTypeEmailLoginResponseBody2

	return EmailLoginResponseBody{
		EmailLoginResponseBody2: &emailLoginResponseBody2,
		Type:                    typ,
	}
}

func (u *EmailLoginResponseBody) UnmarshalJSON(data []byte) error {

	var emailLoginResponseBody1 EmailLoginResponseBody1 = EmailLoginResponseBody1{}
	if err := utils.UnmarshalJSON(data, &emailLoginResponseBody1, "", true, true); err == nil {
		u.EmailLoginResponseBody1 = &emailLoginResponseBody1
		u.Type = EmailLoginResponseBodyTypeEmailLoginResponseBody1
		return nil
	}

	var emailLoginResponseBody2 EmailLoginResponseBody2 = EmailLoginResponseBody2{}
	if err := utils.UnmarshalJSON(data, &emailLoginResponseBody2, "", true, true); err == nil {
		u.EmailLoginResponseBody2 = &emailLoginResponseBody2
		u.Type = EmailLoginResponseBodyTypeEmailLoginResponseBody2
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for EmailLoginResponseBody", string(data))
}

func (u EmailLoginResponseBody) MarshalJSON() ([]byte, error) {
	if u.EmailLoginResponseBody1 != nil {
		return utils.MarshalJSON(u.EmailLoginResponseBody1, "", true)
	}

	if u.EmailLoginResponseBody2 != nil {
		return utils.MarshalJSON(u.EmailLoginResponseBody2, "", true)
	}

	return nil, errors.New("could not marshal union type EmailLoginResponseBody: all fields are null")
}

type EmailLoginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The request was successful and an email was sent
	OneOf *EmailLoginResponseBody
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

func (o *EmailLoginResponse) GetOneOf() *EmailLoginResponseBody {
	if o == nil {
		return nil
	}
	return o.OneOf
}
