// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"errors"
	"fmt"
	"github.com/zchee/terraform-provider-vercel/internal/sdk/internal/utils"
	"github.com/zchee/terraform-provider-vercel/internal/sdk/models/shared"
	"net/http"
)

type GetEequestPromoteAliasesRequest struct {
	// Filter results down to aliases that failed to map to the requested deployment
	FailedOnly *bool `queryParam:"style=form,explode=true,name=failedOnly"`
	// Maximum number of aliases to list from a request (max 100).
	Limit     *float64 `queryParam:"style=form,explode=true,name=limit"`
	ProjectID string   `pathParam:"style=simple,explode=false,name=projectId"`
	// Get aliases created after this epoch timestamp.
	Since *float64 `queryParam:"style=form,explode=true,name=since"`
	// The Team slug to perform the request on behalf of.
	Slug *string `queryParam:"style=form,explode=true,name=slug"`
	// The Team identifier to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
	// Get aliases created before this epoch timestamp.
	Until *float64 `queryParam:"style=form,explode=true,name=until"`
}

func (o *GetEequestPromoteAliasesRequest) GetFailedOnly() *bool {
	if o == nil {
		return nil
	}
	return o.FailedOnly
}

func (o *GetEequestPromoteAliasesRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *GetEequestPromoteAliasesRequest) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

func (o *GetEequestPromoteAliasesRequest) GetSince() *float64 {
	if o == nil {
		return nil
	}
	return o.Since
}

func (o *GetEequestPromoteAliasesRequest) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *GetEequestPromoteAliasesRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

func (o *GetEequestPromoteAliasesRequest) GetUntil() *float64 {
	if o == nil {
		return nil
	}
	return o.Until
}

type GetEequestPromoteAliasesAliases struct {
	Alias  string `json:"alias"`
	ID     string `json:"id"`
	Status string `json:"status"`
}

func (o *GetEequestPromoteAliasesAliases) GetAlias() string {
	if o == nil {
		return ""
	}
	return o.Alias
}

func (o *GetEequestPromoteAliasesAliases) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetEequestPromoteAliasesAliases) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

type GetEequestPromoteAliases2 struct {
	Aliases []GetEequestPromoteAliasesAliases `json:"aliases"`
	// This object contains information related to the pagination of the current request, including the necessary parameters to get the next or previous page of data.
	Pagination shared.Pagination `json:"pagination"`
}

func (o *GetEequestPromoteAliases2) GetAliases() []GetEequestPromoteAliasesAliases {
	if o == nil {
		return []GetEequestPromoteAliasesAliases{}
	}
	return o.Aliases
}

func (o *GetEequestPromoteAliases2) GetPagination() shared.Pagination {
	if o == nil {
		return shared.Pagination{}
	}
	return o.Pagination
}

type GetEequestPromoteAliases1 struct {
}

type GetEequestPromoteAliasesResponseBodyType string

const (
	GetEequestPromoteAliasesResponseBodyTypeGetEequestPromoteAliases1 GetEequestPromoteAliasesResponseBodyType = "getEequestPromoteAliases_1"
	GetEequestPromoteAliasesResponseBodyTypeGetEequestPromoteAliases2 GetEequestPromoteAliasesResponseBodyType = "getEequestPromoteAliases_2"
)

type GetEequestPromoteAliasesResponseBody struct {
	GetEequestPromoteAliases1 *GetEequestPromoteAliases1
	GetEequestPromoteAliases2 *GetEequestPromoteAliases2

	Type GetEequestPromoteAliasesResponseBodyType
}

func CreateGetEequestPromoteAliasesResponseBodyGetEequestPromoteAliases1(getEequestPromoteAliases1 GetEequestPromoteAliases1) GetEequestPromoteAliasesResponseBody {
	typ := GetEequestPromoteAliasesResponseBodyTypeGetEequestPromoteAliases1

	return GetEequestPromoteAliasesResponseBody{
		GetEequestPromoteAliases1: &getEequestPromoteAliases1,
		Type:                      typ,
	}
}

func CreateGetEequestPromoteAliasesResponseBodyGetEequestPromoteAliases2(getEequestPromoteAliases2 GetEequestPromoteAliases2) GetEequestPromoteAliasesResponseBody {
	typ := GetEequestPromoteAliasesResponseBodyTypeGetEequestPromoteAliases2

	return GetEequestPromoteAliasesResponseBody{
		GetEequestPromoteAliases2: &getEequestPromoteAliases2,
		Type:                      typ,
	}
}

func (u *GetEequestPromoteAliasesResponseBody) UnmarshalJSON(data []byte) error {

	var getEequestPromoteAliases1 GetEequestPromoteAliases1 = GetEequestPromoteAliases1{}
	if err := utils.UnmarshalJSON(data, &getEequestPromoteAliases1, "", true, true); err == nil {
		u.GetEequestPromoteAliases1 = &getEequestPromoteAliases1
		u.Type = GetEequestPromoteAliasesResponseBodyTypeGetEequestPromoteAliases1
		return nil
	}

	var getEequestPromoteAliases2 GetEequestPromoteAliases2 = GetEequestPromoteAliases2{}
	if err := utils.UnmarshalJSON(data, &getEequestPromoteAliases2, "", true, true); err == nil {
		u.GetEequestPromoteAliases2 = &getEequestPromoteAliases2
		u.Type = GetEequestPromoteAliasesResponseBodyTypeGetEequestPromoteAliases2
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for GetEequestPromoteAliasesResponseBody", string(data))
}

func (u GetEequestPromoteAliasesResponseBody) MarshalJSON() ([]byte, error) {
	if u.GetEequestPromoteAliases1 != nil {
		return utils.MarshalJSON(u.GetEequestPromoteAliases1, "", true)
	}

	if u.GetEequestPromoteAliases2 != nil {
		return utils.MarshalJSON(u.GetEequestPromoteAliases2, "", true)
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
