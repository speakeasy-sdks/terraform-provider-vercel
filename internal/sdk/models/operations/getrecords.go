// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/internal/utils"
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/models/shared"
	"net/http"
)

type GetRecordsRequest struct {
	Domain string `pathParam:"style=simple,explode=false,name=domain"`
	// Maximum number of records to list from a request.
	Limit *string `queryParam:"style=form,explode=true,name=limit"`
	// Get records created after this JavaScript timestamp.
	Since *string `queryParam:"style=form,explode=true,name=since"`
	// Get records created before this JavaScript timestamp.
	Until *string `queryParam:"style=form,explode=true,name=until"`
	// The Team identifier to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
	// The Team slug to perform the request on behalf of.
	Slug *string `queryParam:"style=form,explode=true,name=slug"`
}

func (o *GetRecordsRequest) GetDomain() string {
	if o == nil {
		return ""
	}
	return o.Domain
}

func (o *GetRecordsRequest) GetLimit() *string {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *GetRecordsRequest) GetSince() *string {
	if o == nil {
		return nil
	}
	return o.Since
}

func (o *GetRecordsRequest) GetUntil() *string {
	if o == nil {
		return nil
	}
	return o.Until
}

func (o *GetRecordsRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

func (o *GetRecordsRequest) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

type GetRecordsResponseBodyDNSType string

const (
	GetRecordsResponseBodyDNSTypeA     GetRecordsResponseBodyDNSType = "A"
	GetRecordsResponseBodyDNSTypeAaaa  GetRecordsResponseBodyDNSType = "AAAA"
	GetRecordsResponseBodyDNSTypeAlias GetRecordsResponseBodyDNSType = "ALIAS"
	GetRecordsResponseBodyDNSTypeCaa   GetRecordsResponseBodyDNSType = "CAA"
	GetRecordsResponseBodyDNSTypeCname GetRecordsResponseBodyDNSType = "CNAME"
	GetRecordsResponseBodyDNSTypeHTTPS GetRecordsResponseBodyDNSType = "HTTPS"
	GetRecordsResponseBodyDNSTypeMx    GetRecordsResponseBodyDNSType = "MX"
	GetRecordsResponseBodyDNSTypeSrv   GetRecordsResponseBodyDNSType = "SRV"
	GetRecordsResponseBodyDNSTypeTxt   GetRecordsResponseBodyDNSType = "TXT"
	GetRecordsResponseBodyDNSTypeNs    GetRecordsResponseBodyDNSType = "NS"
)

func (e GetRecordsResponseBodyDNSType) ToPointer() *GetRecordsResponseBodyDNSType {
	return &e
}
func (e *GetRecordsResponseBodyDNSType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "A":
		fallthrough
	case "AAAA":
		fallthrough
	case "ALIAS":
		fallthrough
	case "CAA":
		fallthrough
	case "CNAME":
		fallthrough
	case "HTTPS":
		fallthrough
	case "MX":
		fallthrough
	case "SRV":
		fallthrough
	case "TXT":
		fallthrough
	case "NS":
		*e = GetRecordsResponseBodyDNSType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetRecordsResponseBodyDNSType: %v", v)
	}
}

type ResponseBodyRecords struct {
	ID         string                        `json:"id"`
	Slug       string                        `json:"slug"`
	Name       string                        `json:"name"`
	Type       GetRecordsResponseBodyDNSType `json:"type"`
	Value      string                        `json:"value"`
	MxPriority *float64                      `json:"mxPriority,omitempty"`
	Priority   *float64                      `json:"priority,omitempty"`
	Creator    string                        `json:"creator"`
	Created    *float64                      `json:"created"`
	Updated    *float64                      `json:"updated"`
	CreatedAt  *float64                      `json:"createdAt"`
	UpdatedAt  *float64                      `json:"updatedAt"`
}

func (o *ResponseBodyRecords) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ResponseBodyRecords) GetSlug() string {
	if o == nil {
		return ""
	}
	return o.Slug
}

func (o *ResponseBodyRecords) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ResponseBodyRecords) GetType() GetRecordsResponseBodyDNSType {
	if o == nil {
		return GetRecordsResponseBodyDNSType("")
	}
	return o.Type
}

func (o *ResponseBodyRecords) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

func (o *ResponseBodyRecords) GetMxPriority() *float64 {
	if o == nil {
		return nil
	}
	return o.MxPriority
}

func (o *ResponseBodyRecords) GetPriority() *float64 {
	if o == nil {
		return nil
	}
	return o.Priority
}

func (o *ResponseBodyRecords) GetCreator() string {
	if o == nil {
		return ""
	}
	return o.Creator
}

func (o *ResponseBodyRecords) GetCreated() *float64 {
	if o == nil {
		return nil
	}
	return o.Created
}

func (o *ResponseBodyRecords) GetUpdated() *float64 {
	if o == nil {
		return nil
	}
	return o.Updated
}

func (o *ResponseBodyRecords) GetCreatedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *ResponseBodyRecords) GetUpdatedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

// GetRecordsResponseBody3 - Successful response retrieving a list of paginated DNS records.
type GetRecordsResponseBody3 struct {
	Records []ResponseBodyRecords `json:"records"`
	// This object contains information related to the pagination of the current request, including the necessary parameters to get the next or previous page of data.
	Pagination shared.Pagination `json:"pagination"`
}

func (o *GetRecordsResponseBody3) GetRecords() []ResponseBodyRecords {
	if o == nil {
		return []ResponseBodyRecords{}
	}
	return o.Records
}

func (o *GetRecordsResponseBody3) GetPagination() shared.Pagination {
	if o == nil {
		return shared.Pagination{}
	}
	return o.Pagination
}

type GetRecordsResponseBodyType string

const (
	GetRecordsResponseBodyTypeA     GetRecordsResponseBodyType = "A"
	GetRecordsResponseBodyTypeAaaa  GetRecordsResponseBodyType = "AAAA"
	GetRecordsResponseBodyTypeAlias GetRecordsResponseBodyType = "ALIAS"
	GetRecordsResponseBodyTypeCaa   GetRecordsResponseBodyType = "CAA"
	GetRecordsResponseBodyTypeCname GetRecordsResponseBodyType = "CNAME"
	GetRecordsResponseBodyTypeHTTPS GetRecordsResponseBodyType = "HTTPS"
	GetRecordsResponseBodyTypeMx    GetRecordsResponseBodyType = "MX"
	GetRecordsResponseBodyTypeSrv   GetRecordsResponseBodyType = "SRV"
	GetRecordsResponseBodyTypeTxt   GetRecordsResponseBodyType = "TXT"
	GetRecordsResponseBodyTypeNs    GetRecordsResponseBodyType = "NS"
)

func (e GetRecordsResponseBodyType) ToPointer() *GetRecordsResponseBodyType {
	return &e
}
func (e *GetRecordsResponseBodyType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "A":
		fallthrough
	case "AAAA":
		fallthrough
	case "ALIAS":
		fallthrough
	case "CAA":
		fallthrough
	case "CNAME":
		fallthrough
	case "HTTPS":
		fallthrough
	case "MX":
		fallthrough
	case "SRV":
		fallthrough
	case "TXT":
		fallthrough
	case "NS":
		*e = GetRecordsResponseBodyType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetRecordsResponseBodyType: %v", v)
	}
}

type Records struct {
	ID         string                     `json:"id"`
	Slug       string                     `json:"slug"`
	Name       string                     `json:"name"`
	Type       GetRecordsResponseBodyType `json:"type"`
	Value      string                     `json:"value"`
	MxPriority *float64                   `json:"mxPriority,omitempty"`
	Priority   *float64                   `json:"priority,omitempty"`
	Creator    string                     `json:"creator"`
	Created    *float64                   `json:"created"`
	Updated    *float64                   `json:"updated"`
	CreatedAt  *float64                   `json:"createdAt"`
	UpdatedAt  *float64                   `json:"updatedAt"`
}

func (o *Records) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Records) GetSlug() string {
	if o == nil {
		return ""
	}
	return o.Slug
}

func (o *Records) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Records) GetType() GetRecordsResponseBodyType {
	if o == nil {
		return GetRecordsResponseBodyType("")
	}
	return o.Type
}

func (o *Records) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

func (o *Records) GetMxPriority() *float64 {
	if o == nil {
		return nil
	}
	return o.MxPriority
}

func (o *Records) GetPriority() *float64 {
	if o == nil {
		return nil
	}
	return o.Priority
}

func (o *Records) GetCreator() string {
	if o == nil {
		return ""
	}
	return o.Creator
}

func (o *Records) GetCreated() *float64 {
	if o == nil {
		return nil
	}
	return o.Created
}

func (o *Records) GetUpdated() *float64 {
	if o == nil {
		return nil
	}
	return o.Updated
}

func (o *Records) GetCreatedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *Records) GetUpdatedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

type GetRecordsResponseBody2 struct {
	Records []Records `json:"records"`
}

func (o *GetRecordsResponseBody2) GetRecords() []Records {
	if o == nil {
		return []Records{}
	}
	return o.Records
}

type GetRecordsResponseBodyUnionType string

const (
	GetRecordsResponseBodyUnionTypeStr                     GetRecordsResponseBodyUnionType = "str"
	GetRecordsResponseBodyUnionTypeGetRecordsResponseBody2 GetRecordsResponseBodyUnionType = "getRecords_responseBody_2"
	GetRecordsResponseBodyUnionTypeGetRecordsResponseBody3 GetRecordsResponseBodyUnionType = "getRecords_responseBody_3"
)

// GetRecordsResponseBody - Successful response retrieving a list of paginated DNS records.
type GetRecordsResponseBody struct {
	Str                     *string
	GetRecordsResponseBody2 *GetRecordsResponseBody2
	GetRecordsResponseBody3 *GetRecordsResponseBody3

	Type GetRecordsResponseBodyUnionType
}

func CreateGetRecordsResponseBodyStr(str string) GetRecordsResponseBody {
	typ := GetRecordsResponseBodyUnionTypeStr

	return GetRecordsResponseBody{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetRecordsResponseBodyGetRecordsResponseBody2(getRecordsResponseBody2 GetRecordsResponseBody2) GetRecordsResponseBody {
	typ := GetRecordsResponseBodyUnionTypeGetRecordsResponseBody2

	return GetRecordsResponseBody{
		GetRecordsResponseBody2: &getRecordsResponseBody2,
		Type:                    typ,
	}
}

func CreateGetRecordsResponseBodyGetRecordsResponseBody3(getRecordsResponseBody3 GetRecordsResponseBody3) GetRecordsResponseBody {
	typ := GetRecordsResponseBodyUnionTypeGetRecordsResponseBody3

	return GetRecordsResponseBody{
		GetRecordsResponseBody3: &getRecordsResponseBody3,
		Type:                    typ,
	}
}

func (u *GetRecordsResponseBody) UnmarshalJSON(data []byte) error {

	var getRecordsResponseBody2 GetRecordsResponseBody2 = GetRecordsResponseBody2{}
	if err := utils.UnmarshalJSON(data, &getRecordsResponseBody2, "", true, true); err == nil {
		u.GetRecordsResponseBody2 = &getRecordsResponseBody2
		u.Type = GetRecordsResponseBodyUnionTypeGetRecordsResponseBody2
		return nil
	}

	var getRecordsResponseBody3 GetRecordsResponseBody3 = GetRecordsResponseBody3{}
	if err := utils.UnmarshalJSON(data, &getRecordsResponseBody3, "", true, true); err == nil {
		u.GetRecordsResponseBody3 = &getRecordsResponseBody3
		u.Type = GetRecordsResponseBodyUnionTypeGetRecordsResponseBody3
		return nil
	}

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = GetRecordsResponseBodyUnionTypeStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for GetRecordsResponseBody", string(data))
}

func (u GetRecordsResponseBody) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.GetRecordsResponseBody2 != nil {
		return utils.MarshalJSON(u.GetRecordsResponseBody2, "", true)
	}

	if u.GetRecordsResponseBody3 != nil {
		return utils.MarshalJSON(u.GetRecordsResponseBody3, "", true)
	}

	return nil, errors.New("could not marshal union type GetRecordsResponseBody: all fields are null")
}

type GetRecordsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful response retrieving a list of paginated DNS records.
	OneOf *GetRecordsResponseBody
}

func (o *GetRecordsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetRecordsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetRecordsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetRecordsResponse) GetOneOf() *GetRecordsResponseBody {
	if o == nil {
		return nil
	}
	return o.OneOf
}
