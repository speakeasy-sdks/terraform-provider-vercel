// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Srv struct {
	Port     *int64  `json:"port"`
	Priority *int64  `json:"priority"`
	Target   *string `json:"target"`
	Weight   *int64  `json:"weight"`
}

func (o *Srv) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *Srv) GetPriority() *int64 {
	if o == nil {
		return nil
	}
	return o.Priority
}

func (o *Srv) GetTarget() *string {
	if o == nil {
		return nil
	}
	return o.Target
}

func (o *Srv) GetWeight() *int64 {
	if o == nil {
		return nil
	}
	return o.Weight
}

// UpdateRecordType - The type of the DNS record
type UpdateRecordType string

const (
	UpdateRecordTypeA     UpdateRecordType = "A"
	UpdateRecordTypeAaaa  UpdateRecordType = "AAAA"
	UpdateRecordTypeAlias UpdateRecordType = "ALIAS"
	UpdateRecordTypeCaa   UpdateRecordType = "CAA"
	UpdateRecordTypeCname UpdateRecordType = "CNAME"
	UpdateRecordTypeMx    UpdateRecordType = "MX"
	UpdateRecordTypeSrv   UpdateRecordType = "SRV"
	UpdateRecordTypeTxt   UpdateRecordType = "TXT"
	UpdateRecordTypeNs    UpdateRecordType = "NS"
)

func (e UpdateRecordType) ToPointer() *UpdateRecordType {
	return &e
}

func (e *UpdateRecordType) UnmarshalJSON(data []byte) error {
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
	case "MX":
		fallthrough
	case "SRV":
		fallthrough
	case "TXT":
		fallthrough
	case "NS":
		*e = UpdateRecordType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateRecordType: %v", v)
	}
}

type UpdateRecordRequestBody struct {
	// A comment to add context on what this DNS record is for
	Comment *string `json:"comment,omitempty"`
	// The MX priority value of the DNS record
	MxPriority *int64 `json:"mxPriority,omitempty"`
	// The name of the DNS record
	Name *string `json:"name,omitempty"`
	Srv  *Srv    `json:"srv,omitempty"`
	// The Time to live (TTL) value of the DNS record
	TTL *int64 `json:"ttl,omitempty"`
	// The type of the DNS record
	Type *UpdateRecordType `json:"type,omitempty"`
	// The value of the DNS record
	Value *string `json:"value,omitempty"`
}

func (o *UpdateRecordRequestBody) GetComment() *string {
	if o == nil {
		return nil
	}
	return o.Comment
}

func (o *UpdateRecordRequestBody) GetMxPriority() *int64 {
	if o == nil {
		return nil
	}
	return o.MxPriority
}

func (o *UpdateRecordRequestBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *UpdateRecordRequestBody) GetSrv() *Srv {
	if o == nil {
		return nil
	}
	return o.Srv
}

func (o *UpdateRecordRequestBody) GetTTL() *int64 {
	if o == nil {
		return nil
	}
	return o.TTL
}

func (o *UpdateRecordRequestBody) GetType() *UpdateRecordType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *UpdateRecordRequestBody) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

type UpdateRecordRequest struct {
	RequestBody *UpdateRecordRequestBody `request:"mediaType=application/json"`
	// The id of the DNS record
	RecordID string `pathParam:"style=simple,explode=false,name=recordId"`
	// The Team identifier or slug to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
}

func (o *UpdateRecordRequest) GetRequestBody() *UpdateRecordRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *UpdateRecordRequest) GetRecordID() string {
	if o == nil {
		return ""
	}
	return o.RecordID
}

func (o *UpdateRecordRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

type RecordType string

const (
	RecordTypeA     RecordType = "A"
	RecordTypeAaaa  RecordType = "AAAA"
	RecordTypeAlias RecordType = "ALIAS"
	RecordTypeCaa   RecordType = "CAA"
	RecordTypeCname RecordType = "CNAME"
	RecordTypeMx    RecordType = "MX"
	RecordTypeSrv   RecordType = "SRV"
	RecordTypeTxt   RecordType = "TXT"
	RecordTypeNs    RecordType = "NS"
)

func (e RecordType) ToPointer() *RecordType {
	return &e
}

func (e *RecordType) UnmarshalJSON(data []byte) error {
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
	case "MX":
		fallthrough
	case "SRV":
		fallthrough
	case "TXT":
		fallthrough
	case "NS":
		*e = RecordType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RecordType: %v", v)
	}
}

type UpdateRecordDNSType string

const (
	UpdateRecordDNSTypeRecord    UpdateRecordDNSType = "record"
	UpdateRecordDNSTypeRecordSys UpdateRecordDNSType = "record-sys"
)

func (e UpdateRecordDNSType) ToPointer() *UpdateRecordDNSType {
	return &e
}

func (e *UpdateRecordDNSType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "record":
		fallthrough
	case "record-sys":
		*e = UpdateRecordDNSType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateRecordDNSType: %v", v)
	}
}

type UpdateRecordResponseBody struct {
	Comment    *string             `json:"comment,omitempty"`
	CreatedAt  *int64              `json:"createdAt,omitempty"`
	Creator    string              `json:"creator"`
	Domain     string              `json:"domain"`
	ID         string              `json:"id"`
	Name       string              `json:"name"`
	RecordType RecordType          `json:"recordType"`
	TTL        *int64              `json:"ttl,omitempty"`
	Type       UpdateRecordDNSType `json:"type"`
	Value      string              `json:"value"`
}

func (o *UpdateRecordResponseBody) GetComment() *string {
	if o == nil {
		return nil
	}
	return o.Comment
}

func (o *UpdateRecordResponseBody) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *UpdateRecordResponseBody) GetCreator() string {
	if o == nil {
		return ""
	}
	return o.Creator
}

func (o *UpdateRecordResponseBody) GetDomain() string {
	if o == nil {
		return ""
	}
	return o.Domain
}

func (o *UpdateRecordResponseBody) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateRecordResponseBody) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *UpdateRecordResponseBody) GetRecordType() RecordType {
	if o == nil {
		return RecordType("")
	}
	return o.RecordType
}

func (o *UpdateRecordResponseBody) GetTTL() *int64 {
	if o == nil {
		return nil
	}
	return o.TTL
}

func (o *UpdateRecordResponseBody) GetType() UpdateRecordDNSType {
	if o == nil {
		return UpdateRecordDNSType("")
	}
	return o.Type
}

func (o *UpdateRecordResponseBody) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

type UpdateRecordResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	Object      *UpdateRecordResponseBody
}

func (o *UpdateRecordResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateRecordResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateRecordResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateRecordResponse) GetObject() *UpdateRecordResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}