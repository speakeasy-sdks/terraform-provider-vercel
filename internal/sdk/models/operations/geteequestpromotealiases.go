// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"errors"
	"fmt"
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/internal/utils"
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/models/shared"
	"net/http"
)

type GetEequestPromoteAliasesRequest struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=projectId"`
	// Maximum number of aliases to list from a request (max 100).
	Limit *float64 `queryParam:"style=form,explode=true,name=limit"`
	// Get aliases created after this epoch timestamp.
	Since *float64 `queryParam:"style=form,explode=true,name=since"`
	// Get aliases created before this epoch timestamp.
	Until *float64 `queryParam:"style=form,explode=true,name=until"`
	// Filter results down to aliases that failed to map to the requested deployment
	FailedOnly *bool `queryParam:"style=form,explode=true,name=failedOnly"`
	// The Team identifier to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
	// The Team slug to perform the request on behalf of.
	Slug *string `queryParam:"style=form,explode=true,name=slug"`
}

func (o *GetEequestPromoteAliasesRequest) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

func (o *GetEequestPromoteAliasesRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *GetEequestPromoteAliasesRequest) GetSince() *float64 {
	if o == nil {
		return nil
	}
	return o.Since
}

func (o *GetEequestPromoteAliasesRequest) GetUntil() *float64 {
	if o == nil {
		return nil
	}
	return o.Until
}

func (o *GetEequestPromoteAliasesRequest) GetFailedOnly() *bool {
	if o == nil {
		return nil
	}
	return o.FailedOnly
}

func (o *GetEequestPromoteAliasesRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

func (o *GetEequestPromoteAliasesRequest) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

type ResponseBodyAliases struct {
	Status string `json:"status"`
	Alias  string `json:"alias"`
	ID     string `json:"id"`
}

func (o *ResponseBodyAliases) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

func (o *ResponseBodyAliases) GetAlias() string {
	if o == nil {
		return ""
	}
	return o.Alias
}

func (o *ResponseBodyAliases) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetEequestPromoteAliasesResponseBody2 struct {
	Aliases []ResponseBodyAliases `json:"aliases"`
	// This object contains information related to the pagination of the current request, including the necessary parameters to get the next or previous page of data.
	Pagination shared.Pagination `json:"pagination"`
}

func (o *GetEequestPromoteAliasesResponseBody2) GetAliases() []ResponseBodyAliases {
	if o == nil {
		return []ResponseBodyAliases{}
	}
	return o.Aliases
}

func (o *GetEequestPromoteAliasesResponseBody2) GetPagination() shared.Pagination {
	if o == nil {
		return shared.Pagination{}
	}
	return o.Pagination
}

type GetEequestPromoteAliasesResponseBody1 struct {
}

type GetEequestPromoteAliasesResponseBodyType string

const (
	GetEequestPromoteAliasesResponseBodyTypeGetEequestPromoteAliasesResponseBody1 GetEequestPromoteAliasesResponseBodyType = "getEequestPromoteAliases_responseBody_1"
	GetEequestPromoteAliasesResponseBodyTypeGetEequestPromoteAliasesResponseBody2 GetEequestPromoteAliasesResponseBodyType = "getEequestPromoteAliases_responseBody_2"
)

type GetEequestPromoteAliasesResponseBody struct {
	GetEequestPromoteAliasesResponseBody1 *GetEequestPromoteAliasesResponseBody1
	GetEequestPromoteAliasesResponseBody2 *GetEequestPromoteAliasesResponseBody2

	Type GetEequestPromoteAliasesResponseBodyType
}

func CreateGetEequestPromoteAliasesResponseBodyGetEequestPromoteAliasesResponseBody1(getEequestPromoteAliasesResponseBody1 GetEequestPromoteAliasesResponseBody1) GetEequestPromoteAliasesResponseBody {
	typ := GetEequestPromoteAliasesResponseBodyTypeGetEequestPromoteAliasesResponseBody1

	return GetEequestPromoteAliasesResponseBody{
		GetEequestPromoteAliasesResponseBody1: &getEequestPromoteAliasesResponseBody1,
		Type:                                  typ,
	}
}

func CreateGetEequestPromoteAliasesResponseBodyGetEequestPromoteAliasesResponseBody2(getEequestPromoteAliasesResponseBody2 GetEequestPromoteAliasesResponseBody2) GetEequestPromoteAliasesResponseBody {
	typ := GetEequestPromoteAliasesResponseBodyTypeGetEequestPromoteAliasesResponseBody2

	return GetEequestPromoteAliasesResponseBody{
		GetEequestPromoteAliasesResponseBody2: &getEequestPromoteAliasesResponseBody2,
		Type:                                  typ,
	}
}

func (u *GetEequestPromoteAliasesResponseBody) UnmarshalJSON(data []byte) error {

	var getEequestPromoteAliasesResponseBody1 GetEequestPromoteAliasesResponseBody1 = GetEequestPromoteAliasesResponseBody1{}
	if err := utils.UnmarshalJSON(data, &getEequestPromoteAliasesResponseBody1, "", true, true); err == nil {
		u.GetEequestPromoteAliasesResponseBody1 = &getEequestPromoteAliasesResponseBody1
		u.Type = GetEequestPromoteAliasesResponseBodyTypeGetEequestPromoteAliasesResponseBody1
		return nil
	}

	var getEequestPromoteAliasesResponseBody2 GetEequestPromoteAliasesResponseBody2 = GetEequestPromoteAliasesResponseBody2{}
	if err := utils.UnmarshalJSON(data, &getEequestPromoteAliasesResponseBody2, "", true, true); err == nil {
		u.GetEequestPromoteAliasesResponseBody2 = &getEequestPromoteAliasesResponseBody2
		u.Type = GetEequestPromoteAliasesResponseBodyTypeGetEequestPromoteAliasesResponseBody2
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for GetEequestPromoteAliasesResponseBody", string(data))
}

func (u GetEequestPromoteAliasesResponseBody) MarshalJSON() ([]byte, error) {
	if u.GetEequestPromoteAliasesResponseBody1 != nil {
		return utils.MarshalJSON(u.GetEequestPromoteAliasesResponseBody1, "", true)
	}

	if u.GetEequestPromoteAliasesResponseBody2 != nil {
		return utils.MarshalJSON(u.GetEequestPromoteAliasesResponseBody2, "", true)
	}

	return nil, errors.New("could not marshal union type GetEequestPromoteAliasesResponseBody: all fields are null")
}

type GetEequestPromoteAliasesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	OneOf       *GetEequestPromoteAliasesResponseBody
}

func (o *GetEequestPromoteAliasesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetEequestPromoteAliasesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetEequestPromoteAliasesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetEequestPromoteAliasesResponse) GetOneOf() *GetEequestPromoteAliasesResponseBody {
	if o == nil {
		return nil
	}
	return o.OneOf
}
