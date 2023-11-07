// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"time"
	"vercel/internal/sdk/pkg/utils"
)

type AssignAliasRequestBody struct {
	// The alias we want to assign to the deployment defined in the URL
	Alias *string `json:"alias,omitempty"`
	// The redirect property will take precedence over the deployment id from the URL and consists of a hostname (like test.com) to which the alias should redirect using status code 307
	Redirect *string `json:"redirect,omitempty"`
}

func (o *AssignAliasRequestBody) GetAlias() *string {
	if o == nil {
		return nil
	}
	return o.Alias
}

func (o *AssignAliasRequestBody) GetRedirect() *string {
	if o == nil {
		return nil
	}
	return o.Redirect
}

type AssignAliasRequest struct {
	RequestBody *AssignAliasRequestBody `request:"mediaType=application/json"`
	// The ID of the deployment the aliases should be listed for
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// The Team identifier or slug to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
}

func (o *AssignAliasRequest) GetRequestBody() *AssignAliasRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *AssignAliasRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *AssignAliasRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

// AssignAliasResponseBody - The alias was successfully assigned to the deployment
type AssignAliasResponseBody struct {
	// The assigned alias name
	Alias string `json:"alias"`
	// The date when the alias was created
	Created time.Time `json:"created"`
	// The unique identifier of the previously aliased deployment, only received when the alias was used before
	OldDeploymentID *string `json:"oldDeploymentId,omitempty"`
	// The unique identifier of the alias
	UID string `json:"uid"`
}

func (a AssignAliasResponseBody) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AssignAliasResponseBody) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AssignAliasResponseBody) GetAlias() string {
	if o == nil {
		return ""
	}
	return o.Alias
}

func (o *AssignAliasResponseBody) GetCreated() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.Created
}

func (o *AssignAliasResponseBody) GetOldDeploymentID() *string {
	if o == nil {
		return nil
	}
	return o.OldDeploymentID
}

func (o *AssignAliasResponseBody) GetUID() string {
	if o == nil {
		return ""
	}
	return o.UID
}

type AssignAliasResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The alias was successfully assigned to the deployment
	Object *AssignAliasResponseBody
}

func (o *AssignAliasResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AssignAliasResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AssignAliasResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *AssignAliasResponse) GetObject() *AssignAliasResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
