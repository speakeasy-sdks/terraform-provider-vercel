// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zchee/terraform-provider-vercel/internal/sdk/pkg/utils"
	"net/http"
)

// CreateRecordDNSRequestRequestBody9Type - The type of record, it could be one of the valid DNS records.
type CreateRecordDNSRequestRequestBody9Type string

const (
	CreateRecordDNSRequestRequestBody9TypeA     CreateRecordDNSRequestRequestBody9Type = "A"
	CreateRecordDNSRequestRequestBody9TypeAaaa  CreateRecordDNSRequestRequestBody9Type = "AAAA"
	CreateRecordDNSRequestRequestBody9TypeAlias CreateRecordDNSRequestRequestBody9Type = "ALIAS"
	CreateRecordDNSRequestRequestBody9TypeCaa   CreateRecordDNSRequestRequestBody9Type = "CAA"
	CreateRecordDNSRequestRequestBody9TypeCname CreateRecordDNSRequestRequestBody9Type = "CNAME"
	CreateRecordDNSRequestRequestBody9TypeMx    CreateRecordDNSRequestRequestBody9Type = "MX"
	CreateRecordDNSRequestRequestBody9TypeSrv   CreateRecordDNSRequestRequestBody9Type = "SRV"
	CreateRecordDNSRequestRequestBody9TypeTxt   CreateRecordDNSRequestRequestBody9Type = "TXT"
	CreateRecordDNSRequestRequestBody9TypeNs    CreateRecordDNSRequestRequestBody9Type = "NS"
)

func (e CreateRecordDNSRequestRequestBody9Type) ToPointer() *CreateRecordDNSRequestRequestBody9Type {
	return &e
}

func (e *CreateRecordDNSRequestRequestBody9Type) UnmarshalJSON(data []byte) error {
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
		*e = CreateRecordDNSRequestRequestBody9Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateRecordDNSRequestRequestBody9Type: %v", v)
	}
}

type Nine struct {
	// A comment to add context on what this DNS record is for
	Comment *string `json:"comment,omitempty"`
	// A subdomain name.
	Name string `json:"name"`
	// The TTL value. Must be a number between 60 and 2147483647. Default value is 60.
	TTL *int64 `json:"ttl,omitempty"`
	// The type of record, it could be one of the valid DNS records.
	Type CreateRecordDNSRequestRequestBody9Type `json:"type"`
	// An NS domain value.
	Value *string `json:"value,omitempty"`
}

func (o *Nine) GetComment() *string {
	if o == nil {
		return nil
	}
	return o.Comment
}

func (o *Nine) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Nine) GetTTL() *int64 {
	if o == nil {
		return nil
	}
	return o.TTL
}

func (o *Nine) GetType() CreateRecordDNSRequestRequestBody9Type {
	if o == nil {
		return CreateRecordDNSRequestRequestBody9Type("")
	}
	return o.Type
}

func (o *Nine) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

// CreateRecordDNSRequestRequestBody8Type - The type of record, it could be one of the valid DNS records.
type CreateRecordDNSRequestRequestBody8Type string

const (
	CreateRecordDNSRequestRequestBody8TypeA     CreateRecordDNSRequestRequestBody8Type = "A"
	CreateRecordDNSRequestRequestBody8TypeAaaa  CreateRecordDNSRequestRequestBody8Type = "AAAA"
	CreateRecordDNSRequestRequestBody8TypeAlias CreateRecordDNSRequestRequestBody8Type = "ALIAS"
	CreateRecordDNSRequestRequestBody8TypeCaa   CreateRecordDNSRequestRequestBody8Type = "CAA"
	CreateRecordDNSRequestRequestBody8TypeCname CreateRecordDNSRequestRequestBody8Type = "CNAME"
	CreateRecordDNSRequestRequestBody8TypeMx    CreateRecordDNSRequestRequestBody8Type = "MX"
	CreateRecordDNSRequestRequestBody8TypeSrv   CreateRecordDNSRequestRequestBody8Type = "SRV"
	CreateRecordDNSRequestRequestBody8TypeTxt   CreateRecordDNSRequestRequestBody8Type = "TXT"
	CreateRecordDNSRequestRequestBody8TypeNs    CreateRecordDNSRequestRequestBody8Type = "NS"
)

func (e CreateRecordDNSRequestRequestBody8Type) ToPointer() *CreateRecordDNSRequestRequestBody8Type {
	return &e
}

func (e *CreateRecordDNSRequestRequestBody8Type) UnmarshalJSON(data []byte) error {
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
		*e = CreateRecordDNSRequestRequestBody8Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateRecordDNSRequestRequestBody8Type: %v", v)
	}
}

type Eight struct {
	// A comment to add context on what this DNS record is for
	Comment *string `json:"comment,omitempty"`
	// The TTL value. Must be a number between 60 and 2147483647. Default value is 60.
	TTL *int64 `json:"ttl,omitempty"`
	// The type of record, it could be one of the valid DNS records.
	Type CreateRecordDNSRequestRequestBody8Type `json:"type"`
	// A TXT record containing arbitrary text.
	Value string `json:"value"`
}

func (o *Eight) GetComment() *string {
	if o == nil {
		return nil
	}
	return o.Comment
}

func (o *Eight) GetTTL() *int64 {
	if o == nil {
		return nil
	}
	return o.TTL
}

func (o *Eight) GetType() CreateRecordDNSRequestRequestBody8Type {
	if o == nil {
		return CreateRecordDNSRequestRequestBody8Type("")
	}
	return o.Type
}

func (o *Eight) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

type CreateRecordSrv struct {
	Port     int64   `json:"port"`
	Priority int64   `json:"priority"`
	Target   *string `json:"target,omitempty"`
	Weight   int64   `json:"weight"`
}

func (o *CreateRecordSrv) GetPort() int64 {
	if o == nil {
		return 0
	}
	return o.Port
}

func (o *CreateRecordSrv) GetPriority() int64 {
	if o == nil {
		return 0
	}
	return o.Priority
}

func (o *CreateRecordSrv) GetTarget() *string {
	if o == nil {
		return nil
	}
	return o.Target
}

func (o *CreateRecordSrv) GetWeight() int64 {
	if o == nil {
		return 0
	}
	return o.Weight
}

// CreateRecordDNSRequestRequestBody7Type - The type of record, it could be one of the valid DNS records.
type CreateRecordDNSRequestRequestBody7Type string

const (
	CreateRecordDNSRequestRequestBody7TypeA     CreateRecordDNSRequestRequestBody7Type = "A"
	CreateRecordDNSRequestRequestBody7TypeAaaa  CreateRecordDNSRequestRequestBody7Type = "AAAA"
	CreateRecordDNSRequestRequestBody7TypeAlias CreateRecordDNSRequestRequestBody7Type = "ALIAS"
	CreateRecordDNSRequestRequestBody7TypeCaa   CreateRecordDNSRequestRequestBody7Type = "CAA"
	CreateRecordDNSRequestRequestBody7TypeCname CreateRecordDNSRequestRequestBody7Type = "CNAME"
	CreateRecordDNSRequestRequestBody7TypeMx    CreateRecordDNSRequestRequestBody7Type = "MX"
	CreateRecordDNSRequestRequestBody7TypeSrv   CreateRecordDNSRequestRequestBody7Type = "SRV"
	CreateRecordDNSRequestRequestBody7TypeTxt   CreateRecordDNSRequestRequestBody7Type = "TXT"
	CreateRecordDNSRequestRequestBody7TypeNs    CreateRecordDNSRequestRequestBody7Type = "NS"
)

func (e CreateRecordDNSRequestRequestBody7Type) ToPointer() *CreateRecordDNSRequestRequestBody7Type {
	return &e
}

func (e *CreateRecordDNSRequestRequestBody7Type) UnmarshalJSON(data []byte) error {
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
		*e = CreateRecordDNSRequestRequestBody7Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateRecordDNSRequestRequestBody7Type: %v", v)
	}
}

type Seven struct {
	// A comment to add context on what this DNS record is for
	Comment *string         `json:"comment,omitempty"`
	Srv     CreateRecordSrv `json:"srv"`
	// The TTL value. Must be a number between 60 and 2147483647. Default value is 60.
	TTL *int64 `json:"ttl,omitempty"`
	// The type of record, it could be one of the valid DNS records.
	Type CreateRecordDNSRequestRequestBody7Type `json:"type"`
}

func (o *Seven) GetComment() *string {
	if o == nil {
		return nil
	}
	return o.Comment
}

func (o *Seven) GetSrv() CreateRecordSrv {
	if o == nil {
		return CreateRecordSrv{}
	}
	return o.Srv
}

func (o *Seven) GetTTL() *int64 {
	if o == nil {
		return nil
	}
	return o.TTL
}

func (o *Seven) GetType() CreateRecordDNSRequestRequestBody7Type {
	if o == nil {
		return CreateRecordDNSRequestRequestBody7Type("")
	}
	return o.Type
}

// CreateRecordDNSRequestRequestBody6Type - The type of record, it could be one of the valid DNS records.
type CreateRecordDNSRequestRequestBody6Type string

const (
	CreateRecordDNSRequestRequestBody6TypeA     CreateRecordDNSRequestRequestBody6Type = "A"
	CreateRecordDNSRequestRequestBody6TypeAaaa  CreateRecordDNSRequestRequestBody6Type = "AAAA"
	CreateRecordDNSRequestRequestBody6TypeAlias CreateRecordDNSRequestRequestBody6Type = "ALIAS"
	CreateRecordDNSRequestRequestBody6TypeCaa   CreateRecordDNSRequestRequestBody6Type = "CAA"
	CreateRecordDNSRequestRequestBody6TypeCname CreateRecordDNSRequestRequestBody6Type = "CNAME"
	CreateRecordDNSRequestRequestBody6TypeMx    CreateRecordDNSRequestRequestBody6Type = "MX"
	CreateRecordDNSRequestRequestBody6TypeSrv   CreateRecordDNSRequestRequestBody6Type = "SRV"
	CreateRecordDNSRequestRequestBody6TypeTxt   CreateRecordDNSRequestRequestBody6Type = "TXT"
	CreateRecordDNSRequestRequestBody6TypeNs    CreateRecordDNSRequestRequestBody6Type = "NS"
)

func (e CreateRecordDNSRequestRequestBody6Type) ToPointer() *CreateRecordDNSRequestRequestBody6Type {
	return &e
}

func (e *CreateRecordDNSRequestRequestBody6Type) UnmarshalJSON(data []byte) error {
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
		*e = CreateRecordDNSRequestRequestBody6Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateRecordDNSRequestRequestBody6Type: %v", v)
	}
}

type Six struct {
	// A comment to add context on what this DNS record is for
	Comment    *string `json:"comment,omitempty"`
	MxPriority int64   `json:"mxPriority"`
	// A subdomain name or an empty string for the root domain.
	Name string `json:"name"`
	// The TTL value. Must be a number between 60 and 2147483647. Default value is 60.
	TTL *int64 `json:"ttl,omitempty"`
	// The type of record, it could be one of the valid DNS records.
	Type CreateRecordDNSRequestRequestBody6Type `json:"type"`
	// An MX record specifying the mail server responsible for accepting messages on behalf of the domain name.
	Value string `json:"value"`
}

func (o *Six) GetComment() *string {
	if o == nil {
		return nil
	}
	return o.Comment
}

func (o *Six) GetMxPriority() int64 {
	if o == nil {
		return 0
	}
	return o.MxPriority
}

func (o *Six) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Six) GetTTL() *int64 {
	if o == nil {
		return nil
	}
	return o.TTL
}

func (o *Six) GetType() CreateRecordDNSRequestRequestBody6Type {
	if o == nil {
		return CreateRecordDNSRequestRequestBody6Type("")
	}
	return o.Type
}

func (o *Six) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

// CreateRecordDNSRequestRequestBody5Type - The type of record, it could be one of the valid DNS records.
type CreateRecordDNSRequestRequestBody5Type string

const (
	CreateRecordDNSRequestRequestBody5TypeA     CreateRecordDNSRequestRequestBody5Type = "A"
	CreateRecordDNSRequestRequestBody5TypeAaaa  CreateRecordDNSRequestRequestBody5Type = "AAAA"
	CreateRecordDNSRequestRequestBody5TypeAlias CreateRecordDNSRequestRequestBody5Type = "ALIAS"
	CreateRecordDNSRequestRequestBody5TypeCaa   CreateRecordDNSRequestRequestBody5Type = "CAA"
	CreateRecordDNSRequestRequestBody5TypeCname CreateRecordDNSRequestRequestBody5Type = "CNAME"
	CreateRecordDNSRequestRequestBody5TypeMx    CreateRecordDNSRequestRequestBody5Type = "MX"
	CreateRecordDNSRequestRequestBody5TypeSrv   CreateRecordDNSRequestRequestBody5Type = "SRV"
	CreateRecordDNSRequestRequestBody5TypeTxt   CreateRecordDNSRequestRequestBody5Type = "TXT"
	CreateRecordDNSRequestRequestBody5TypeNs    CreateRecordDNSRequestRequestBody5Type = "NS"
)

func (e CreateRecordDNSRequestRequestBody5Type) ToPointer() *CreateRecordDNSRequestRequestBody5Type {
	return &e
}

func (e *CreateRecordDNSRequestRequestBody5Type) UnmarshalJSON(data []byte) error {
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
		*e = CreateRecordDNSRequestRequestBody5Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateRecordDNSRequestRequestBody5Type: %v", v)
	}
}

type CreateRecord5 struct {
	// A comment to add context on what this DNS record is for
	Comment *string `json:"comment,omitempty"`
	// A subdomain name or an empty string for the root domain.
	Name string `json:"name"`
	// The TTL value. Must be a number between 60 and 2147483647. Default value is 60.
	TTL *int64 `json:"ttl,omitempty"`
	// The type of record, it could be one of the valid DNS records.
	Type CreateRecordDNSRequestRequestBody5Type `json:"type"`
	// A CNAME record mapping to another domain name.
	Value *string `json:"value,omitempty"`
}

func (o *CreateRecord5) GetComment() *string {
	if o == nil {
		return nil
	}
	return o.Comment
}

func (o *CreateRecord5) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateRecord5) GetTTL() *int64 {
	if o == nil {
		return nil
	}
	return o.TTL
}

func (o *CreateRecord5) GetType() CreateRecordDNSRequestRequestBody5Type {
	if o == nil {
		return CreateRecordDNSRequestRequestBody5Type("")
	}
	return o.Type
}

func (o *CreateRecord5) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

// CreateRecordDNSRequestRequestBodyType - The type of record, it could be one of the valid DNS records.
type CreateRecordDNSRequestRequestBodyType string

const (
	CreateRecordDNSRequestRequestBodyTypeA     CreateRecordDNSRequestRequestBodyType = "A"
	CreateRecordDNSRequestRequestBodyTypeAaaa  CreateRecordDNSRequestRequestBodyType = "AAAA"
	CreateRecordDNSRequestRequestBodyTypeAlias CreateRecordDNSRequestRequestBodyType = "ALIAS"
	CreateRecordDNSRequestRequestBodyTypeCaa   CreateRecordDNSRequestRequestBodyType = "CAA"
	CreateRecordDNSRequestRequestBodyTypeCname CreateRecordDNSRequestRequestBodyType = "CNAME"
	CreateRecordDNSRequestRequestBodyTypeMx    CreateRecordDNSRequestRequestBodyType = "MX"
	CreateRecordDNSRequestRequestBodyTypeSrv   CreateRecordDNSRequestRequestBodyType = "SRV"
	CreateRecordDNSRequestRequestBodyTypeTxt   CreateRecordDNSRequestRequestBodyType = "TXT"
	CreateRecordDNSRequestRequestBodyTypeNs    CreateRecordDNSRequestRequestBodyType = "NS"
)

func (e CreateRecordDNSRequestRequestBodyType) ToPointer() *CreateRecordDNSRequestRequestBodyType {
	return &e
}

func (e *CreateRecordDNSRequestRequestBodyType) UnmarshalJSON(data []byte) error {
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
		*e = CreateRecordDNSRequestRequestBodyType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateRecordDNSRequestRequestBodyType: %v", v)
	}
}

type CreateRecord4 struct {
	// A comment to add context on what this DNS record is for
	Comment *string `json:"comment,omitempty"`
	// A subdomain name or an empty string for the root domain.
	Name string `json:"name"`
	// The TTL value. Must be a number between 60 and 2147483647. Default value is 60.
	TTL *int64 `json:"ttl,omitempty"`
	// The type of record, it could be one of the valid DNS records.
	Type CreateRecordDNSRequestRequestBodyType `json:"type"`
	// A CAA record to specify which Certificate Authorities (CAs) are allowed to issue certificates for the domain.
	Value string `json:"value"`
}

func (o *CreateRecord4) GetComment() *string {
	if o == nil {
		return nil
	}
	return o.Comment
}

func (o *CreateRecord4) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateRecord4) GetTTL() *int64 {
	if o == nil {
		return nil
	}
	return o.TTL
}

func (o *CreateRecord4) GetType() CreateRecordDNSRequestRequestBodyType {
	if o == nil {
		return CreateRecordDNSRequestRequestBodyType("")
	}
	return o.Type
}

func (o *CreateRecord4) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

// CreateRecordDNSRequestType - The type of record, it could be one of the valid DNS records.
type CreateRecordDNSRequestType string

const (
	CreateRecordDNSRequestTypeA     CreateRecordDNSRequestType = "A"
	CreateRecordDNSRequestTypeAaaa  CreateRecordDNSRequestType = "AAAA"
	CreateRecordDNSRequestTypeAlias CreateRecordDNSRequestType = "ALIAS"
	CreateRecordDNSRequestTypeCaa   CreateRecordDNSRequestType = "CAA"
	CreateRecordDNSRequestTypeCname CreateRecordDNSRequestType = "CNAME"
	CreateRecordDNSRequestTypeMx    CreateRecordDNSRequestType = "MX"
	CreateRecordDNSRequestTypeSrv   CreateRecordDNSRequestType = "SRV"
	CreateRecordDNSRequestTypeTxt   CreateRecordDNSRequestType = "TXT"
	CreateRecordDNSRequestTypeNs    CreateRecordDNSRequestType = "NS"
)

func (e CreateRecordDNSRequestType) ToPointer() *CreateRecordDNSRequestType {
	return &e
}

func (e *CreateRecordDNSRequestType) UnmarshalJSON(data []byte) error {
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
		*e = CreateRecordDNSRequestType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateRecordDNSRequestType: %v", v)
	}
}

type CreateRecord3 struct {
	// A comment to add context on what this DNS record is for
	Comment *string `json:"comment,omitempty"`
	// A subdomain name or an empty string for the root domain.
	Name string `json:"name"`
	// The TTL value. Must be a number between 60 and 2147483647. Default value is 60.
	TTL *int64 `json:"ttl,omitempty"`
	// The type of record, it could be one of the valid DNS records.
	Type CreateRecordDNSRequestType `json:"type"`
	// An ALIAS virtual record pointing to a hostname resolved to an A record on server side.
	Value string `json:"value"`
}

func (o *CreateRecord3) GetComment() *string {
	if o == nil {
		return nil
	}
	return o.Comment
}

func (o *CreateRecord3) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateRecord3) GetTTL() *int64 {
	if o == nil {
		return nil
	}
	return o.TTL
}

func (o *CreateRecord3) GetType() CreateRecordDNSRequestType {
	if o == nil {
		return CreateRecordDNSRequestType("")
	}
	return o.Type
}

func (o *CreateRecord3) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

// CreateRecordDNSType - The type of record, it could be one of the valid DNS records.
type CreateRecordDNSType string

const (
	CreateRecordDNSTypeA     CreateRecordDNSType = "A"
	CreateRecordDNSTypeAaaa  CreateRecordDNSType = "AAAA"
	CreateRecordDNSTypeAlias CreateRecordDNSType = "ALIAS"
	CreateRecordDNSTypeCaa   CreateRecordDNSType = "CAA"
	CreateRecordDNSTypeCname CreateRecordDNSType = "CNAME"
	CreateRecordDNSTypeMx    CreateRecordDNSType = "MX"
	CreateRecordDNSTypeSrv   CreateRecordDNSType = "SRV"
	CreateRecordDNSTypeTxt   CreateRecordDNSType = "TXT"
	CreateRecordDNSTypeNs    CreateRecordDNSType = "NS"
)

func (e CreateRecordDNSType) ToPointer() *CreateRecordDNSType {
	return &e
}

func (e *CreateRecordDNSType) UnmarshalJSON(data []byte) error {
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
		*e = CreateRecordDNSType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateRecordDNSType: %v", v)
	}
}

type CreateRecord2 struct {
	// A comment to add context on what this DNS record is for
	Comment *string `json:"comment,omitempty"`
	// A subdomain name or an empty string for the root domain.
	Name string `json:"name"`
	// The TTL value. Must be a number between 60 and 2147483647. Default value is 60.
	TTL *int64 `json:"ttl,omitempty"`
	// The type of record, it could be one of the valid DNS records.
	Type CreateRecordDNSType `json:"type"`
	// An AAAA record pointing to an IPv6 address.
	Value string `json:"value"`
}

func (o *CreateRecord2) GetComment() *string {
	if o == nil {
		return nil
	}
	return o.Comment
}

func (o *CreateRecord2) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateRecord2) GetTTL() *int64 {
	if o == nil {
		return nil
	}
	return o.TTL
}

func (o *CreateRecord2) GetType() CreateRecordDNSType {
	if o == nil {
		return CreateRecordDNSType("")
	}
	return o.Type
}

func (o *CreateRecord2) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

// CreateRecordType - The type of record, it could be one of the valid DNS records.
type CreateRecordType string

const (
	CreateRecordTypeA     CreateRecordType = "A"
	CreateRecordTypeAaaa  CreateRecordType = "AAAA"
	CreateRecordTypeAlias CreateRecordType = "ALIAS"
	CreateRecordTypeCaa   CreateRecordType = "CAA"
	CreateRecordTypeCname CreateRecordType = "CNAME"
	CreateRecordTypeMx    CreateRecordType = "MX"
	CreateRecordTypeSrv   CreateRecordType = "SRV"
	CreateRecordTypeTxt   CreateRecordType = "TXT"
	CreateRecordTypeNs    CreateRecordType = "NS"
)

func (e CreateRecordType) ToPointer() *CreateRecordType {
	return &e
}

func (e *CreateRecordType) UnmarshalJSON(data []byte) error {
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
		*e = CreateRecordType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateRecordType: %v", v)
	}
}

type CreateRecord1 struct {
	// A comment to add context on what this DNS record is for
	Comment *string `json:"comment,omitempty"`
	// A subdomain name or an empty string for the root domain.
	Name string `json:"name"`
	// The TTL value. Must be a number between 60 and 2147483647. Default value is 60.
	TTL *int64 `json:"ttl,omitempty"`
	// The type of record, it could be one of the valid DNS records.
	Type CreateRecordType `json:"type"`
	// The record value must be a valid IPv4 address.
	Value string `json:"value"`
}

func (o *CreateRecord1) GetComment() *string {
	if o == nil {
		return nil
	}
	return o.Comment
}

func (o *CreateRecord1) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateRecord1) GetTTL() *int64 {
	if o == nil {
		return nil
	}
	return o.TTL
}

func (o *CreateRecord1) GetType() CreateRecordType {
	if o == nil {
		return CreateRecordType("")
	}
	return o.Type
}

func (o *CreateRecord1) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

type CreateRecordRequestBodyType string

const (
	CreateRecordRequestBodyTypeCreateRecord1 CreateRecordRequestBodyType = "createRecord_1"
	CreateRecordRequestBodyTypeCreateRecord2 CreateRecordRequestBodyType = "createRecord_2"
	CreateRecordRequestBodyTypeCreateRecord3 CreateRecordRequestBodyType = "createRecord_3"
	CreateRecordRequestBodyTypeCreateRecord4 CreateRecordRequestBodyType = "createRecord_4"
	CreateRecordRequestBodyTypeCreateRecord5 CreateRecordRequestBodyType = "createRecord_5"
	CreateRecordRequestBodyTypeSix           CreateRecordRequestBodyType = "6"
	CreateRecordRequestBodyTypeSeven         CreateRecordRequestBodyType = "7"
	CreateRecordRequestBodyTypeEight         CreateRecordRequestBodyType = "8"
	CreateRecordRequestBodyTypeNine          CreateRecordRequestBodyType = "9"
)

type CreateRecordRequestBody struct {
	CreateRecord1 *CreateRecord1
	CreateRecord2 *CreateRecord2
	CreateRecord3 *CreateRecord3
	CreateRecord4 *CreateRecord4
	CreateRecord5 *CreateRecord5
	Six           *Six
	Seven         *Seven
	Eight         *Eight
	Nine          *Nine

	Type CreateRecordRequestBodyType
}

func CreateCreateRecordRequestBodyCreateRecord1(createRecord1 CreateRecord1) CreateRecordRequestBody {
	typ := CreateRecordRequestBodyTypeCreateRecord1

	return CreateRecordRequestBody{
		CreateRecord1: &createRecord1,
		Type:          typ,
	}
}

func CreateCreateRecordRequestBodyCreateRecord2(createRecord2 CreateRecord2) CreateRecordRequestBody {
	typ := CreateRecordRequestBodyTypeCreateRecord2

	return CreateRecordRequestBody{
		CreateRecord2: &createRecord2,
		Type:          typ,
	}
}

func CreateCreateRecordRequestBodyCreateRecord3(createRecord3 CreateRecord3) CreateRecordRequestBody {
	typ := CreateRecordRequestBodyTypeCreateRecord3

	return CreateRecordRequestBody{
		CreateRecord3: &createRecord3,
		Type:          typ,
	}
}

func CreateCreateRecordRequestBodyCreateRecord4(createRecord4 CreateRecord4) CreateRecordRequestBody {
	typ := CreateRecordRequestBodyTypeCreateRecord4

	return CreateRecordRequestBody{
		CreateRecord4: &createRecord4,
		Type:          typ,
	}
}

func CreateCreateRecordRequestBodyCreateRecord5(createRecord5 CreateRecord5) CreateRecordRequestBody {
	typ := CreateRecordRequestBodyTypeCreateRecord5

	return CreateRecordRequestBody{
		CreateRecord5: &createRecord5,
		Type:          typ,
	}
}

func CreateCreateRecordRequestBodySix(six Six) CreateRecordRequestBody {
	typ := CreateRecordRequestBodyTypeSix

	return CreateRecordRequestBody{
		Six:  &six,
		Type: typ,
	}
}

func CreateCreateRecordRequestBodySeven(seven Seven) CreateRecordRequestBody {
	typ := CreateRecordRequestBodyTypeSeven

	return CreateRecordRequestBody{
		Seven: &seven,
		Type:  typ,
	}
}

func CreateCreateRecordRequestBodyEight(eight Eight) CreateRecordRequestBody {
	typ := CreateRecordRequestBodyTypeEight

	return CreateRecordRequestBody{
		Eight: &eight,
		Type:  typ,
	}
}

func CreateCreateRecordRequestBodyNine(nine Nine) CreateRecordRequestBody {
	typ := CreateRecordRequestBodyTypeNine

	return CreateRecordRequestBody{
		Nine: &nine,
		Type: typ,
	}
}

func (u *CreateRecordRequestBody) UnmarshalJSON(data []byte) error {

	seven := new(Seven)
	if err := utils.UnmarshalJSON(data, &seven, "", true, true); err == nil {
		u.Seven = seven
		u.Type = CreateRecordRequestBodyTypeSeven
		return nil
	}

	eight := new(Eight)
	if err := utils.UnmarshalJSON(data, &eight, "", true, true); err == nil {
		u.Eight = eight
		u.Type = CreateRecordRequestBodyTypeEight
		return nil
	}

	createRecord1 := new(CreateRecord1)
	if err := utils.UnmarshalJSON(data, &createRecord1, "", true, true); err == nil {
		u.CreateRecord1 = createRecord1
		u.Type = CreateRecordRequestBodyTypeCreateRecord1
		return nil
	}

	createRecord2 := new(CreateRecord2)
	if err := utils.UnmarshalJSON(data, &createRecord2, "", true, true); err == nil {
		u.CreateRecord2 = createRecord2
		u.Type = CreateRecordRequestBodyTypeCreateRecord2
		return nil
	}

	createRecord3 := new(CreateRecord3)
	if err := utils.UnmarshalJSON(data, &createRecord3, "", true, true); err == nil {
		u.CreateRecord3 = createRecord3
		u.Type = CreateRecordRequestBodyTypeCreateRecord3
		return nil
	}

	createRecord4 := new(CreateRecord4)
	if err := utils.UnmarshalJSON(data, &createRecord4, "", true, true); err == nil {
		u.CreateRecord4 = createRecord4
		u.Type = CreateRecordRequestBodyTypeCreateRecord4
		return nil
	}

	createRecord5 := new(CreateRecord5)
	if err := utils.UnmarshalJSON(data, &createRecord5, "", true, true); err == nil {
		u.CreateRecord5 = createRecord5
		u.Type = CreateRecordRequestBodyTypeCreateRecord5
		return nil
	}

	nine := new(Nine)
	if err := utils.UnmarshalJSON(data, &nine, "", true, true); err == nil {
		u.Nine = nine
		u.Type = CreateRecordRequestBodyTypeNine
		return nil
	}

	six := new(Six)
	if err := utils.UnmarshalJSON(data, &six, "", true, true); err == nil {
		u.Six = six
		u.Type = CreateRecordRequestBodyTypeSix
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u CreateRecordRequestBody) MarshalJSON() ([]byte, error) {
	if u.CreateRecord1 != nil {
		return utils.MarshalJSON(u.CreateRecord1, "", true)
	}

	if u.CreateRecord2 != nil {
		return utils.MarshalJSON(u.CreateRecord2, "", true)
	}

	if u.CreateRecord3 != nil {
		return utils.MarshalJSON(u.CreateRecord3, "", true)
	}

	if u.CreateRecord4 != nil {
		return utils.MarshalJSON(u.CreateRecord4, "", true)
	}

	if u.CreateRecord5 != nil {
		return utils.MarshalJSON(u.CreateRecord5, "", true)
	}

	if u.Six != nil {
		return utils.MarshalJSON(u.Six, "", true)
	}

	if u.Seven != nil {
		return utils.MarshalJSON(u.Seven, "", true)
	}

	if u.Eight != nil {
		return utils.MarshalJSON(u.Eight, "", true)
	}

	if u.Nine != nil {
		return utils.MarshalJSON(u.Nine, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type CreateRecordRequest struct {
	RequestBody *CreateRecordRequestBody `request:"mediaType=application/json"`
	// The domain used to create the DNS record.
	Domain string `pathParam:"style=simple,explode=false,name=domain"`
	// The Team identifier or slug to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
}

func (o *CreateRecordRequest) GetRequestBody() *CreateRecordRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *CreateRecordRequest) GetDomain() string {
	if o == nil {
		return ""
	}
	return o.Domain
}

func (o *CreateRecordRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

type CreateRecordDNS2 struct {
	// The id of the newly created DNS record
	UID string `json:"uid"`
}

func (o *CreateRecordDNS2) GetUID() string {
	if o == nil {
		return ""
	}
	return o.UID
}

type CreateRecordDNS1 struct {
	UID     string `json:"uid"`
	Updated int64  `json:"updated"`
}

func (o *CreateRecordDNS1) GetUID() string {
	if o == nil {
		return ""
	}
	return o.UID
}

func (o *CreateRecordDNS1) GetUpdated() int64 {
	if o == nil {
		return 0
	}
	return o.Updated
}

type CreateRecordResponseBodyType string

const (
	CreateRecordResponseBodyTypeCreateRecordDNS1 CreateRecordResponseBodyType = "createRecord_dns_1"
	CreateRecordResponseBodyTypeCreateRecordDNS2 CreateRecordResponseBodyType = "createRecord_dns_2"
)

// CreateRecordResponseBody - Successful response showing the uid of the newly created DNS record.
type CreateRecordResponseBody struct {
	CreateRecordDNS1 *CreateRecordDNS1
	CreateRecordDNS2 *CreateRecordDNS2

	Type CreateRecordResponseBodyType
}

func CreateCreateRecordResponseBodyCreateRecordDNS1(createRecordDNS1 CreateRecordDNS1) CreateRecordResponseBody {
	typ := CreateRecordResponseBodyTypeCreateRecordDNS1

	return CreateRecordResponseBody{
		CreateRecordDNS1: &createRecordDNS1,
		Type:             typ,
	}
}

func CreateCreateRecordResponseBodyCreateRecordDNS2(createRecordDNS2 CreateRecordDNS2) CreateRecordResponseBody {
	typ := CreateRecordResponseBodyTypeCreateRecordDNS2

	return CreateRecordResponseBody{
		CreateRecordDNS2: &createRecordDNS2,
		Type:             typ,
	}
}

func (u *CreateRecordResponseBody) UnmarshalJSON(data []byte) error {

	createRecordDNS2 := new(CreateRecordDNS2)
	if err := utils.UnmarshalJSON(data, &createRecordDNS2, "", true, true); err == nil {
		u.CreateRecordDNS2 = createRecordDNS2
		u.Type = CreateRecordResponseBodyTypeCreateRecordDNS2
		return nil
	}

	createRecordDNS1 := new(CreateRecordDNS1)
	if err := utils.UnmarshalJSON(data, &createRecordDNS1, "", true, true); err == nil {
		u.CreateRecordDNS1 = createRecordDNS1
		u.Type = CreateRecordResponseBodyTypeCreateRecordDNS1
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u CreateRecordResponseBody) MarshalJSON() ([]byte, error) {
	if u.CreateRecordDNS1 != nil {
		return utils.MarshalJSON(u.CreateRecordDNS1, "", true)
	}

	if u.CreateRecordDNS2 != nil {
		return utils.MarshalJSON(u.CreateRecordDNS2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type CreateRecordResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful response showing the uid of the newly created DNS record.
	OneOf *CreateRecordResponseBody
}

func (o *CreateRecordResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateRecordResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateRecordResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateRecordResponse) GetOneOf() *CreateRecordResponseBody {
	if o == nil {
		return nil
	}
	return o.OneOf
}
