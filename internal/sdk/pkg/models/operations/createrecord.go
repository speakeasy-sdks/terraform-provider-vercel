// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// CreateRecordRequestBody9Type - Must be of type `NS`.
type CreateRecordRequestBody9Type string

const (
	CreateRecordRequestBody9TypeNs CreateRecordRequestBody9Type = "NS"
)

func (e CreateRecordRequestBody9Type) ToPointer() *CreateRecordRequestBody9Type {
	return &e
}

func (e *CreateRecordRequestBody9Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NS":
		*e = CreateRecordRequestBody9Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateRecordRequestBody9Type: %v", v)
	}
}

type CreateRecordRequestBody9 struct {
	// A comment to add context on what this DNS record is for
	Comment *string `json:"comment,omitempty"`
	// A subdomain name.
	Name string `json:"name"`
	// The TTL value. Must be a number between 60 and 2147483647. Default value is 60.
	TTL *int64 `json:"ttl,omitempty"`
	// Must be of type `NS`.
	Type CreateRecordRequestBody9Type `json:"type"`
	// An NS domain value.
	Value *string `json:"value,omitempty"`
}

// CreateRecordRequestBody8Type - Must be of type `TXT`.
type CreateRecordRequestBody8Type string

const (
	CreateRecordRequestBody8TypeTxt CreateRecordRequestBody8Type = "TXT"
)

func (e CreateRecordRequestBody8Type) ToPointer() *CreateRecordRequestBody8Type {
	return &e
}

func (e *CreateRecordRequestBody8Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TXT":
		*e = CreateRecordRequestBody8Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateRecordRequestBody8Type: %v", v)
	}
}

type CreateRecordRequestBody8 struct {
	// A comment to add context on what this DNS record is for
	Comment *string `json:"comment,omitempty"`
	// The TTL value. Must be a number between 60 and 2147483647. Default value is 60.
	TTL *int64 `json:"ttl,omitempty"`
	// Must be of type `TXT`.
	Type CreateRecordRequestBody8Type `json:"type"`
	// A TXT record containing arbitrary text.
	Value string `json:"value"`
}

type CreateRecordRequestBody7Srv struct {
	Port     int64   `json:"port"`
	Priority int64   `json:"priority"`
	Target   *string `json:"target,omitempty"`
	Weight   int64   `json:"weight"`
}

// CreateRecordRequestBody7Type - Must be of type `SRV`.
type CreateRecordRequestBody7Type string

const (
	CreateRecordRequestBody7TypeSrv CreateRecordRequestBody7Type = "SRV"
)

func (e CreateRecordRequestBody7Type) ToPointer() *CreateRecordRequestBody7Type {
	return &e
}

func (e *CreateRecordRequestBody7Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SRV":
		*e = CreateRecordRequestBody7Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateRecordRequestBody7Type: %v", v)
	}
}

type CreateRecordRequestBody7 struct {
	// A comment to add context on what this DNS record is for
	Comment *string                     `json:"comment,omitempty"`
	Srv     CreateRecordRequestBody7Srv `json:"srv"`
	// The TTL value. Must be a number between 60 and 2147483647. Default value is 60.
	TTL *int64 `json:"ttl,omitempty"`
	// Must be of type `SRV`.
	Type CreateRecordRequestBody7Type `json:"type"`
}

// CreateRecordRequestBody6Type - Must be of type `MX`.
type CreateRecordRequestBody6Type string

const (
	CreateRecordRequestBody6TypeMx CreateRecordRequestBody6Type = "MX"
)

func (e CreateRecordRequestBody6Type) ToPointer() *CreateRecordRequestBody6Type {
	return &e
}

func (e *CreateRecordRequestBody6Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "MX":
		*e = CreateRecordRequestBody6Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateRecordRequestBody6Type: %v", v)
	}
}

type CreateRecordRequestBody6 struct {
	// A comment to add context on what this DNS record is for
	Comment    *string `json:"comment,omitempty"`
	MxPriority int64   `json:"mxPriority"`
	// A subdomain name or an empty string for the root domain.
	Name string `json:"name"`
	// The TTL value. Must be a number between 60 and 2147483647. Default value is 60.
	TTL *int64 `json:"ttl,omitempty"`
	// Must be of type `MX`.
	Type CreateRecordRequestBody6Type `json:"type"`
	// An MX record specifying the mail server responsible for accepting messages on behalf of the domain name.
	Value string `json:"value"`
}

// CreateRecordRequestBody5Type - Must be of type `CNAME`.
type CreateRecordRequestBody5Type string

const (
	CreateRecordRequestBody5TypeCname CreateRecordRequestBody5Type = "CNAME"
)

func (e CreateRecordRequestBody5Type) ToPointer() *CreateRecordRequestBody5Type {
	return &e
}

func (e *CreateRecordRequestBody5Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "CNAME":
		*e = CreateRecordRequestBody5Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateRecordRequestBody5Type: %v", v)
	}
}

type CreateRecordRequestBody5 struct {
	// A comment to add context on what this DNS record is for
	Comment *string `json:"comment,omitempty"`
	// A subdomain name or an empty string for the root domain.
	Name string `json:"name"`
	// The TTL value. Must be a number between 60 and 2147483647. Default value is 60.
	TTL *int64 `json:"ttl,omitempty"`
	// Must be of type `CNAME`.
	Type CreateRecordRequestBody5Type `json:"type"`
	// A CNAME record mapping to another domain name.
	Value *string `json:"value,omitempty"`
}

// CreateRecordRequestBody4Type - Must be of type `CAA`.
type CreateRecordRequestBody4Type string

const (
	CreateRecordRequestBody4TypeCaa CreateRecordRequestBody4Type = "CAA"
)

func (e CreateRecordRequestBody4Type) ToPointer() *CreateRecordRequestBody4Type {
	return &e
}

func (e *CreateRecordRequestBody4Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "CAA":
		*e = CreateRecordRequestBody4Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateRecordRequestBody4Type: %v", v)
	}
}

type CreateRecordRequestBody4 struct {
	// A comment to add context on what this DNS record is for
	Comment *string `json:"comment,omitempty"`
	// A subdomain name or an empty string for the root domain.
	Name string `json:"name"`
	// The TTL value. Must be a number between 60 and 2147483647. Default value is 60.
	TTL *int64 `json:"ttl,omitempty"`
	// Must be of type `CAA`.
	Type CreateRecordRequestBody4Type `json:"type"`
	// A CAA record to specify which Certificate Authorities (CAs) are allowed to issue certificates for the domain.
	Value string `json:"value"`
}

// CreateRecordRequestBody3Type - Must be of type `ALIAS`.
type CreateRecordRequestBody3Type string

const (
	CreateRecordRequestBody3TypeAlias CreateRecordRequestBody3Type = "ALIAS"
)

func (e CreateRecordRequestBody3Type) ToPointer() *CreateRecordRequestBody3Type {
	return &e
}

func (e *CreateRecordRequestBody3Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ALIAS":
		*e = CreateRecordRequestBody3Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateRecordRequestBody3Type: %v", v)
	}
}

type CreateRecordRequestBody3 struct {
	// A comment to add context on what this DNS record is for
	Comment *string `json:"comment,omitempty"`
	// A subdomain name or an empty string for the root domain.
	Name string `json:"name"`
	// The TTL value. Must be a number between 60 and 2147483647. Default value is 60.
	TTL *int64 `json:"ttl,omitempty"`
	// Must be of type `ALIAS`.
	Type CreateRecordRequestBody3Type `json:"type"`
	// An ALIAS virtual record pointing to a hostname resolved to an A record on server side.
	Value string `json:"value"`
}

// CreateRecordRequestBody2Type - Must be of type `AAAA`.
type CreateRecordRequestBody2Type string

const (
	CreateRecordRequestBody2TypeAaaa CreateRecordRequestBody2Type = "AAAA"
)

func (e CreateRecordRequestBody2Type) ToPointer() *CreateRecordRequestBody2Type {
	return &e
}

func (e *CreateRecordRequestBody2Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "AAAA":
		*e = CreateRecordRequestBody2Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateRecordRequestBody2Type: %v", v)
	}
}

type CreateRecordRequestBody2 struct {
	// A comment to add context on what this DNS record is for
	Comment *string `json:"comment,omitempty"`
	// A subdomain name or an empty string for the root domain.
	Name string `json:"name"`
	// The TTL value. Must be a number between 60 and 2147483647. Default value is 60.
	TTL *int64 `json:"ttl,omitempty"`
	// Must be of type `AAAA`.
	Type CreateRecordRequestBody2Type `json:"type"`
	// An AAAA record pointing to an IPv6 address.
	Value string `json:"value"`
}

// CreateRecordRequestBody1Type - Must be of type `A`.
type CreateRecordRequestBody1Type string

const (
	CreateRecordRequestBody1TypeA CreateRecordRequestBody1Type = "A"
)

func (e CreateRecordRequestBody1Type) ToPointer() *CreateRecordRequestBody1Type {
	return &e
}

func (e *CreateRecordRequestBody1Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "A":
		*e = CreateRecordRequestBody1Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateRecordRequestBody1Type: %v", v)
	}
}

type CreateRecordRequestBody1 struct {
	// A comment to add context on what this DNS record is for
	Comment *string `json:"comment,omitempty"`
	// A subdomain name or an empty string for the root domain.
	Name string `json:"name"`
	// The TTL value. Must be a number between 60 and 2147483647. Default value is 60.
	TTL *int64 `json:"ttl,omitempty"`
	// Must be of type `A`.
	Type CreateRecordRequestBody1Type `json:"type"`
	// The record value must be a valid IPv4 address.
	Value string `json:"value"`
}

type CreateRecordRequestBodyType string

const (
	CreateRecordRequestBodyTypeCreateRecordRequestBody1 CreateRecordRequestBodyType = "createRecord_requestBody_1"
	CreateRecordRequestBodyTypeCreateRecordRequestBody2 CreateRecordRequestBodyType = "createRecord_requestBody_2"
	CreateRecordRequestBodyTypeCreateRecordRequestBody3 CreateRecordRequestBodyType = "createRecord_requestBody_3"
	CreateRecordRequestBodyTypeCreateRecordRequestBody4 CreateRecordRequestBodyType = "createRecord_requestBody_4"
	CreateRecordRequestBodyTypeCreateRecordRequestBody5 CreateRecordRequestBodyType = "createRecord_requestBody_5"
	CreateRecordRequestBodyTypeCreateRecordRequestBody6 CreateRecordRequestBodyType = "createRecord_requestBody_6"
	CreateRecordRequestBodyTypeCreateRecordRequestBody7 CreateRecordRequestBodyType = "createRecord_requestBody_7"
	CreateRecordRequestBodyTypeCreateRecordRequestBody8 CreateRecordRequestBodyType = "createRecord_requestBody_8"
	CreateRecordRequestBodyTypeCreateRecordRequestBody9 CreateRecordRequestBodyType = "createRecord_requestBody_9"
)

type CreateRecordRequestBody struct {
	CreateRecordRequestBody1 *CreateRecordRequestBody1
	CreateRecordRequestBody2 *CreateRecordRequestBody2
	CreateRecordRequestBody3 *CreateRecordRequestBody3
	CreateRecordRequestBody4 *CreateRecordRequestBody4
	CreateRecordRequestBody5 *CreateRecordRequestBody5
	CreateRecordRequestBody6 *CreateRecordRequestBody6
	CreateRecordRequestBody7 *CreateRecordRequestBody7
	CreateRecordRequestBody8 *CreateRecordRequestBody8
	CreateRecordRequestBody9 *CreateRecordRequestBody9

	Type CreateRecordRequestBodyType
}

func CreateCreateRecordRequestBodyCreateRecordRequestBody1(createRecordRequestBody1 CreateRecordRequestBody1) CreateRecordRequestBody {
	typ := CreateRecordRequestBodyTypeCreateRecordRequestBody1

	return CreateRecordRequestBody{
		CreateRecordRequestBody1: &createRecordRequestBody1,
		Type:                     typ,
	}
}

func CreateCreateRecordRequestBodyCreateRecordRequestBody2(createRecordRequestBody2 CreateRecordRequestBody2) CreateRecordRequestBody {
	typ := CreateRecordRequestBodyTypeCreateRecordRequestBody2

	return CreateRecordRequestBody{
		CreateRecordRequestBody2: &createRecordRequestBody2,
		Type:                     typ,
	}
}

func CreateCreateRecordRequestBodyCreateRecordRequestBody3(createRecordRequestBody3 CreateRecordRequestBody3) CreateRecordRequestBody {
	typ := CreateRecordRequestBodyTypeCreateRecordRequestBody3

	return CreateRecordRequestBody{
		CreateRecordRequestBody3: &createRecordRequestBody3,
		Type:                     typ,
	}
}

func CreateCreateRecordRequestBodyCreateRecordRequestBody4(createRecordRequestBody4 CreateRecordRequestBody4) CreateRecordRequestBody {
	typ := CreateRecordRequestBodyTypeCreateRecordRequestBody4

	return CreateRecordRequestBody{
		CreateRecordRequestBody4: &createRecordRequestBody4,
		Type:                     typ,
	}
}

func CreateCreateRecordRequestBodyCreateRecordRequestBody5(createRecordRequestBody5 CreateRecordRequestBody5) CreateRecordRequestBody {
	typ := CreateRecordRequestBodyTypeCreateRecordRequestBody5

	return CreateRecordRequestBody{
		CreateRecordRequestBody5: &createRecordRequestBody5,
		Type:                     typ,
	}
}

func CreateCreateRecordRequestBodyCreateRecordRequestBody6(createRecordRequestBody6 CreateRecordRequestBody6) CreateRecordRequestBody {
	typ := CreateRecordRequestBodyTypeCreateRecordRequestBody6

	return CreateRecordRequestBody{
		CreateRecordRequestBody6: &createRecordRequestBody6,
		Type:                     typ,
	}
}

func CreateCreateRecordRequestBodyCreateRecordRequestBody7(createRecordRequestBody7 CreateRecordRequestBody7) CreateRecordRequestBody {
	typ := CreateRecordRequestBodyTypeCreateRecordRequestBody7

	return CreateRecordRequestBody{
		CreateRecordRequestBody7: &createRecordRequestBody7,
		Type:                     typ,
	}
}

func CreateCreateRecordRequestBodyCreateRecordRequestBody8(createRecordRequestBody8 CreateRecordRequestBody8) CreateRecordRequestBody {
	typ := CreateRecordRequestBodyTypeCreateRecordRequestBody8

	return CreateRecordRequestBody{
		CreateRecordRequestBody8: &createRecordRequestBody8,
		Type:                     typ,
	}
}

func CreateCreateRecordRequestBodyCreateRecordRequestBody9(createRecordRequestBody9 CreateRecordRequestBody9) CreateRecordRequestBody {
	typ := CreateRecordRequestBodyTypeCreateRecordRequestBody9

	return CreateRecordRequestBody{
		CreateRecordRequestBody9: &createRecordRequestBody9,
		Type:                     typ,
	}
}

func (u *CreateRecordRequestBody) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	createRecordRequestBody7 := new(CreateRecordRequestBody7)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&createRecordRequestBody7); err == nil {
		u.CreateRecordRequestBody7 = createRecordRequestBody7
		u.Type = CreateRecordRequestBodyTypeCreateRecordRequestBody7
		return nil
	}

	createRecordRequestBody8 := new(CreateRecordRequestBody8)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&createRecordRequestBody8); err == nil {
		u.CreateRecordRequestBody8 = createRecordRequestBody8
		u.Type = CreateRecordRequestBodyTypeCreateRecordRequestBody8
		return nil
	}

	createRecordRequestBody1 := new(CreateRecordRequestBody1)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&createRecordRequestBody1); err == nil {
		u.CreateRecordRequestBody1 = createRecordRequestBody1
		u.Type = CreateRecordRequestBodyTypeCreateRecordRequestBody1
		return nil
	}

	createRecordRequestBody2 := new(CreateRecordRequestBody2)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&createRecordRequestBody2); err == nil {
		u.CreateRecordRequestBody2 = createRecordRequestBody2
		u.Type = CreateRecordRequestBodyTypeCreateRecordRequestBody2
		return nil
	}

	createRecordRequestBody3 := new(CreateRecordRequestBody3)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&createRecordRequestBody3); err == nil {
		u.CreateRecordRequestBody3 = createRecordRequestBody3
		u.Type = CreateRecordRequestBodyTypeCreateRecordRequestBody3
		return nil
	}

	createRecordRequestBody4 := new(CreateRecordRequestBody4)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&createRecordRequestBody4); err == nil {
		u.CreateRecordRequestBody4 = createRecordRequestBody4
		u.Type = CreateRecordRequestBodyTypeCreateRecordRequestBody4
		return nil
	}

	createRecordRequestBody5 := new(CreateRecordRequestBody5)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&createRecordRequestBody5); err == nil {
		u.CreateRecordRequestBody5 = createRecordRequestBody5
		u.Type = CreateRecordRequestBodyTypeCreateRecordRequestBody5
		return nil
	}

	createRecordRequestBody9 := new(CreateRecordRequestBody9)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&createRecordRequestBody9); err == nil {
		u.CreateRecordRequestBody9 = createRecordRequestBody9
		u.Type = CreateRecordRequestBodyTypeCreateRecordRequestBody9
		return nil
	}

	createRecordRequestBody6 := new(CreateRecordRequestBody6)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&createRecordRequestBody6); err == nil {
		u.CreateRecordRequestBody6 = createRecordRequestBody6
		u.Type = CreateRecordRequestBodyTypeCreateRecordRequestBody6
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u CreateRecordRequestBody) MarshalJSON() ([]byte, error) {
	if u.CreateRecordRequestBody7 != nil {
		return json.Marshal(u.CreateRecordRequestBody7)
	}

	if u.CreateRecordRequestBody8 != nil {
		return json.Marshal(u.CreateRecordRequestBody8)
	}

	if u.CreateRecordRequestBody1 != nil {
		return json.Marshal(u.CreateRecordRequestBody1)
	}

	if u.CreateRecordRequestBody2 != nil {
		return json.Marshal(u.CreateRecordRequestBody2)
	}

	if u.CreateRecordRequestBody3 != nil {
		return json.Marshal(u.CreateRecordRequestBody3)
	}

	if u.CreateRecordRequestBody4 != nil {
		return json.Marshal(u.CreateRecordRequestBody4)
	}

	if u.CreateRecordRequestBody5 != nil {
		return json.Marshal(u.CreateRecordRequestBody5)
	}

	if u.CreateRecordRequestBody9 != nil {
		return json.Marshal(u.CreateRecordRequestBody9)
	}

	if u.CreateRecordRequestBody6 != nil {
		return json.Marshal(u.CreateRecordRequestBody6)
	}

	return nil, nil
}

type CreateRecordRequest struct {
	RequestBody *CreateRecordRequestBody `request:"mediaType=application/json"`
	// The domain used to create the DNS record.
	Domain string `pathParam:"style=simple,explode=false,name=domain"`
	// The Team identifier or slug to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
}

type CreateRecord200ApplicationJSON2 struct {
	// The id of the newly created DNS record
	UID string `json:"uid"`
}

type CreateRecord200ApplicationJSON1 struct {
	UID     string `json:"uid"`
	Updated int64  `json:"updated"`
}

type CreateRecord200ApplicationJSONType string

const (
	CreateRecord200ApplicationJSONTypeCreateRecord200ApplicationJSON1 CreateRecord200ApplicationJSONType = "createRecord_200ApplicationJSON_1"
	CreateRecord200ApplicationJSONTypeCreateRecord200ApplicationJSON2 CreateRecord200ApplicationJSONType = "createRecord_200ApplicationJSON_2"
)

type CreateRecord200ApplicationJSON struct {
	CreateRecord200ApplicationJSON1 *CreateRecord200ApplicationJSON1
	CreateRecord200ApplicationJSON2 *CreateRecord200ApplicationJSON2

	Type CreateRecord200ApplicationJSONType
}

func CreateCreateRecord200ApplicationJSONCreateRecord200ApplicationJSON1(createRecord200ApplicationJSON1 CreateRecord200ApplicationJSON1) CreateRecord200ApplicationJSON {
	typ := CreateRecord200ApplicationJSONTypeCreateRecord200ApplicationJSON1

	return CreateRecord200ApplicationJSON{
		CreateRecord200ApplicationJSON1: &createRecord200ApplicationJSON1,
		Type:                            typ,
	}
}

func CreateCreateRecord200ApplicationJSONCreateRecord200ApplicationJSON2(createRecord200ApplicationJSON2 CreateRecord200ApplicationJSON2) CreateRecord200ApplicationJSON {
	typ := CreateRecord200ApplicationJSONTypeCreateRecord200ApplicationJSON2

	return CreateRecord200ApplicationJSON{
		CreateRecord200ApplicationJSON2: &createRecord200ApplicationJSON2,
		Type:                            typ,
	}
}

func (u *CreateRecord200ApplicationJSON) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	createRecord200ApplicationJSON2 := new(CreateRecord200ApplicationJSON2)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&createRecord200ApplicationJSON2); err == nil {
		u.CreateRecord200ApplicationJSON2 = createRecord200ApplicationJSON2
		u.Type = CreateRecord200ApplicationJSONTypeCreateRecord200ApplicationJSON2
		return nil
	}

	createRecord200ApplicationJSON1 := new(CreateRecord200ApplicationJSON1)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&createRecord200ApplicationJSON1); err == nil {
		u.CreateRecord200ApplicationJSON1 = createRecord200ApplicationJSON1
		u.Type = CreateRecord200ApplicationJSONTypeCreateRecord200ApplicationJSON1
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u CreateRecord200ApplicationJSON) MarshalJSON() ([]byte, error) {
	if u.CreateRecord200ApplicationJSON2 != nil {
		return json.Marshal(u.CreateRecord200ApplicationJSON2)
	}

	if u.CreateRecord200ApplicationJSON1 != nil {
		return json.Marshal(u.CreateRecord200ApplicationJSON1)
	}

	return nil, nil
}

type CreateRecordResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful response showing the uid of the newly created DNS record.
	CreateRecord200ApplicationJSONOneOf *CreateRecord200ApplicationJSON
}
