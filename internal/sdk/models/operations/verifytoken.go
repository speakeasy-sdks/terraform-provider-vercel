// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// TeamPlan - The plan for this user's team (pro or hobby).
type TeamPlan string

const (
	TeamPlanPro   TeamPlan = "pro"
	TeamPlanHobby TeamPlan = "hobby"
)

func (e TeamPlan) ToPointer() *TeamPlan {
	return &e
}
func (e *TeamPlan) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "pro":
		fallthrough
	case "hobby":
		*e = TeamPlan(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TeamPlan: %v", v)
	}
}

type VerifyTokenRequest struct {
	// Email to verify the login.
	Email *string `queryParam:"style=form,explode=true,name=email"`
	// The page on which the user started their session.
	LandingPage *string `queryParam:"style=form,explode=true,name=landingPage"`
	OppID       *string `queryParam:"style=form,explode=true,name=oppId"`
	// The page that sent the user to the signup page.
	PageBeforeConversionPage *string `queryParam:"style=form,explode=true,name=pageBeforeConversionPage"`
	// Referrer to the session.
	SessionReferrer *string `queryParam:"style=form,explode=true,name=sessionReferrer"`
	// The SAML Profile ID, when connecting a SAML Profile to a Team member for the first time.
	SsoUserID *string `queryParam:"style=form,explode=true,name=ssoUserId"`
	// The name of this user's team.
	TeamName *string `queryParam:"style=form,explode=true,name=teamName"`
	// The plan for this user's team (pro or hobby).
	TeamPlan *TeamPlan `queryParam:"style=form,explode=true,name=teamPlan"`
	// The slug for this user's team.
	TeamSlug *string `queryParam:"style=form,explode=true,name=teamSlug"`
	// The token returned when the login was requested.
	Token string `queryParam:"style=form,explode=true,name=token"`
	// The desired name for the token. It will be displayed on the user account details.
	TokenName   *string `queryParam:"style=form,explode=true,name=tokenName"`
	UtmCampaign *string `queryParam:"style=form,explode=true,name=utmCampaign"`
	UtmMedium   *string `queryParam:"style=form,explode=true,name=utmMedium"`
	UtmSource   *string `queryParam:"style=form,explode=true,name=utmSource"`
	UtmTerm     *string `queryParam:"style=form,explode=true,name=utmTerm"`
}

func (o *VerifyTokenRequest) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *VerifyTokenRequest) GetLandingPage() *string {
	if o == nil {
		return nil
	}
	return o.LandingPage
}

func (o *VerifyTokenRequest) GetOppID() *string {
	if o == nil {
		return nil
	}
	return o.OppID
}

func (o *VerifyTokenRequest) GetPageBeforeConversionPage() *string {
	if o == nil {
		return nil
	}
	return o.PageBeforeConversionPage
}

func (o *VerifyTokenRequest) GetSessionReferrer() *string {
	if o == nil {
		return nil
	}
	return o.SessionReferrer
}

func (o *VerifyTokenRequest) GetSsoUserID() *string {
	if o == nil {
		return nil
	}
	return o.SsoUserID
}

func (o *VerifyTokenRequest) GetTeamName() *string {
	if o == nil {
		return nil
	}
	return o.TeamName
}

func (o *VerifyTokenRequest) GetTeamPlan() *TeamPlan {
	if o == nil {
		return nil
	}
	return o.TeamPlan
}

func (o *VerifyTokenRequest) GetTeamSlug() *string {
	if o == nil {
		return nil
	}
	return o.TeamSlug
}

func (o *VerifyTokenRequest) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

func (o *VerifyTokenRequest) GetTokenName() *string {
	if o == nil {
		return nil
	}
	return o.TokenName
}

func (o *VerifyTokenRequest) GetUtmCampaign() *string {
	if o == nil {
		return nil
	}
	return o.UtmCampaign
}

func (o *VerifyTokenRequest) GetUtmMedium() *string {
	if o == nil {
		return nil
	}
	return o.UtmMedium
}

func (o *VerifyTokenRequest) GetUtmSource() *string {
	if o == nil {
		return nil
	}
	return o.UtmSource
}

func (o *VerifyTokenRequest) GetUtmTerm() *string {
	if o == nil {
		return nil
	}
	return o.UtmTerm
}

// VerifyTokenResponseBody - The verification was successful.
type VerifyTokenResponseBody struct {
	// Email address of the authenticated user.
	Email string `json:"email"`
	// When completing SAML Single Sign-On authentication, this will be the ID of the Team that was authenticated for.
	TeamID *string `json:"teamId,omitempty"`
	// The user authentication token that can be used to perform API requests.
	Token string `json:"token"`
}

func (o *VerifyTokenResponseBody) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

func (o *VerifyTokenResponseBody) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

func (o *VerifyTokenResponseBody) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

type VerifyTokenResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The verification was successful.
	Object *VerifyTokenResponseBody
}

func (o *VerifyTokenResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *VerifyTokenResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *VerifyTokenResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *VerifyTokenResponse) GetObject() *VerifyTokenResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
