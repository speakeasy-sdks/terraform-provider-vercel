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

// EmailLogin200ApplicationJSON - The request was successful and an email was sent
type EmailLogin200ApplicationJSON struct {
	// The code the user is going to receive on the email. **Must** be displayed to the user so they can verify the request is the correct.
	SecurityCode string `json:"securityCode"`
	// The token used to verify the user accepted the login request
	Token string `json:"token"`
}

type EmailLoginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The request was successful and an email was sent
	EmailLogin200ApplicationJSONObject *EmailLogin200ApplicationJSON
}
