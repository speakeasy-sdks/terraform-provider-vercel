// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/internal/utils"
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/models/shared"
	"net/http"
	"time"
)

type GetSecretsRequest struct {
	// Filter out secrets based on comma separated secret ids.
	ID *string `queryParam:"style=form,explode=true,name=id"`
	// Filter out secrets that belong to a project.
	ProjectID *string `queryParam:"style=form,explode=true,name=projectId"`
	// The Team identifier to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
	// The Team slug to perform the request on behalf of.
	Slug *string `queryParam:"style=form,explode=true,name=slug"`
}

func (o *GetSecretsRequest) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetSecretsRequest) GetProjectID() *string {
	if o == nil {
		return nil
	}
	return o.ProjectID
}

func (o *GetSecretsRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

func (o *GetSecretsRequest) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

// Secrets - Data representing a secret.
type Secrets struct {
	// The date when the secret was created.
	Created time.Time `json:"created"`
	// The name of the secret.
	Name string `json:"name"`
	// The unique identifier of the team the secret was created for.
	TeamID *string `json:"teamId,omitempty"`
	// The unique identifier of the secret.
	UID string `json:"uid"`
	// The unique identifier of the user who created the secret.
	UserID *string `json:"userId,omitempty"`
	// The value of the secret.
	Value *string `json:"value,omitempty"`
	// Timestamp for when the secret was created.
	CreatedAt *float64 `json:"createdAt,omitempty"`
	// The unique identifier of the project which the secret belongs to.
	ProjectID *string `json:"projectId,omitempty"`
	// Indicates whether the secret value can be decrypted after it has been created.
	Decryptable *bool `json:"decryptable,omitempty"`
}

func (s Secrets) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *Secrets) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Secrets) GetCreated() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.Created
}

func (o *Secrets) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Secrets) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

func (o *Secrets) GetUID() string {
	if o == nil {
		return ""
	}
	return o.UID
}

func (o *Secrets) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}

func (o *Secrets) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

func (o *Secrets) GetCreatedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *Secrets) GetProjectID() *string {
	if o == nil {
		return nil
	}
	return o.ProjectID
}

func (o *Secrets) GetDecryptable() *bool {
	if o == nil {
		return nil
	}
	return o.Decryptable
}

// GetSecretsResponseBody - Successful response retrieving a list of secrets.
type GetSecretsResponseBody struct {
	Secrets []Secrets `json:"secrets"`
	// This object contains information related to the pagination of the current request, including the necessary parameters to get the next or previous page of data.
	Pagination shared.Pagination `json:"pagination"`
}

func (o *GetSecretsResponseBody) GetSecrets() []Secrets {
	if o == nil {
		return []Secrets{}
	}
	return o.Secrets
}

func (o *GetSecretsResponseBody) GetPagination() shared.Pagination {
	if o == nil {
		return shared.Pagination{}
	}
	return o.Pagination
}

type GetSecretsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful response retrieving a list of secrets.
	Object *GetSecretsResponseBody
}

func (o *GetSecretsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSecretsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSecretsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetSecretsResponse) GetObject() *GetSecretsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
