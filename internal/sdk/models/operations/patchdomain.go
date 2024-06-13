// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"errors"
	"fmt"
	"github.com/vercel/terraform-provider-vercel/internal/sdk/internal/utils"
	"net/http"
)

// PatchDomainRequestBody2 - move-out
type PatchDomainRequestBody2 struct {
	Op *string `json:"op,omitempty"`
	// User or team to move domain to
	Destination *string `json:"destination,omitempty"`
}

func (o *PatchDomainRequestBody2) GetOp() *string {
	if o == nil {
		return nil
	}
	return o.Op
}

func (o *PatchDomainRequestBody2) GetDestination() *string {
	if o == nil {
		return nil
	}
	return o.Destination
}

// PatchDomainRequestBody1 - update
type PatchDomainRequestBody1 struct {
	Op *string `json:"op,omitempty"`
	// Specifies whether domain should be renewed.
	Renew *bool `json:"renew,omitempty"`
	// The custom nameservers for this project.
	CustomNameservers []string `json:"customNameservers,omitempty"`
	// Specifies whether this is a DNS zone that intends to use Vercel's nameservers.
	Zone *bool `json:"zone,omitempty"`
}

func (o *PatchDomainRequestBody1) GetOp() *string {
	if o == nil {
		return nil
	}
	return o.Op
}

func (o *PatchDomainRequestBody1) GetRenew() *bool {
	if o == nil {
		return nil
	}
	return o.Renew
}

func (o *PatchDomainRequestBody1) GetCustomNameservers() []string {
	if o == nil {
		return nil
	}
	return o.CustomNameservers
}

func (o *PatchDomainRequestBody1) GetZone() *bool {
	if o == nil {
		return nil
	}
	return o.Zone
}

type PatchDomainRequestBodyType string

const (
	PatchDomainRequestBodyTypePatchDomainRequestBody1 PatchDomainRequestBodyType = "patchDomain_requestBody_1"
	PatchDomainRequestBodyTypePatchDomainRequestBody2 PatchDomainRequestBodyType = "patchDomain_requestBody_2"
)

type PatchDomainRequestBody struct {
	PatchDomainRequestBody1 *PatchDomainRequestBody1
	PatchDomainRequestBody2 *PatchDomainRequestBody2

	Type PatchDomainRequestBodyType
}

func CreatePatchDomainRequestBodyPatchDomainRequestBody1(patchDomainRequestBody1 PatchDomainRequestBody1) PatchDomainRequestBody {
	typ := PatchDomainRequestBodyTypePatchDomainRequestBody1

	return PatchDomainRequestBody{
		PatchDomainRequestBody1: &patchDomainRequestBody1,
		Type:                    typ,
	}
}

func CreatePatchDomainRequestBodyPatchDomainRequestBody2(patchDomainRequestBody2 PatchDomainRequestBody2) PatchDomainRequestBody {
	typ := PatchDomainRequestBodyTypePatchDomainRequestBody2

	return PatchDomainRequestBody{
		PatchDomainRequestBody2: &patchDomainRequestBody2,
		Type:                    typ,
	}
}

func (u *PatchDomainRequestBody) UnmarshalJSON(data []byte) error {

	var patchDomainRequestBody2 PatchDomainRequestBody2 = PatchDomainRequestBody2{}
	if err := utils.UnmarshalJSON(data, &patchDomainRequestBody2, "", true, true); err == nil {
		u.PatchDomainRequestBody2 = &patchDomainRequestBody2
		u.Type = PatchDomainRequestBodyTypePatchDomainRequestBody2
		return nil
	}

	var patchDomainRequestBody1 PatchDomainRequestBody1 = PatchDomainRequestBody1{}
	if err := utils.UnmarshalJSON(data, &patchDomainRequestBody1, "", true, true); err == nil {
		u.PatchDomainRequestBody1 = &patchDomainRequestBody1
		u.Type = PatchDomainRequestBodyTypePatchDomainRequestBody1
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for PatchDomainRequestBody", string(data))
}

func (u PatchDomainRequestBody) MarshalJSON() ([]byte, error) {
	if u.PatchDomainRequestBody1 != nil {
		return utils.MarshalJSON(u.PatchDomainRequestBody1, "", true)
	}

	if u.PatchDomainRequestBody2 != nil {
		return utils.MarshalJSON(u.PatchDomainRequestBody2, "", true)
	}

	return nil, errors.New("could not marshal union type PatchDomainRequestBody: all fields are null")
}

type PatchDomainRequest struct {
	Domain string `pathParam:"style=simple,explode=false,name=domain"`
	// The Team identifier to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
	// The Team slug to perform the request on behalf of.
	Slug        *string                 `queryParam:"style=form,explode=true,name=slug"`
	RequestBody *PatchDomainRequestBody `request:"mediaType=application/json"`
}

func (o *PatchDomainRequest) GetDomain() string {
	if o == nil {
		return ""
	}
	return o.Domain
}

func (o *PatchDomainRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

func (o *PatchDomainRequest) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *PatchDomainRequest) GetRequestBody() *PatchDomainRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

type PatchDomainResponseBody3 struct {
	Renew             *bool    `json:"renew,omitempty"`
	CustomNameservers []string `json:"customNameservers,omitempty"`
	Zone              *bool    `json:"zone,omitempty"`
}

func (o *PatchDomainResponseBody3) GetRenew() *bool {
	if o == nil {
		return nil
	}
	return o.Renew
}

func (o *PatchDomainResponseBody3) GetCustomNameservers() []string {
	if o == nil {
		return nil
	}
	return o.CustomNameservers
}

func (o *PatchDomainResponseBody3) GetZone() *bool {
	if o == nil {
		return nil
	}
	return o.Zone
}

type PatchDomainResponseBody2 struct {
	Moved bool   `json:"moved"`
	Token string `json:"token"`
}

func (o *PatchDomainResponseBody2) GetMoved() bool {
	if o == nil {
		return false
	}
	return o.Moved
}

func (o *PatchDomainResponseBody2) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

type PatchDomainResponseBody1 struct {
	Moved bool `json:"moved"`
}

func (o *PatchDomainResponseBody1) GetMoved() bool {
	if o == nil {
		return false
	}
	return o.Moved
}

type PatchDomainResponseBodyType string

const (
	PatchDomainResponseBodyTypePatchDomainResponseBody1 PatchDomainResponseBodyType = "patchDomain_responseBody_1"
	PatchDomainResponseBodyTypePatchDomainResponseBody2 PatchDomainResponseBodyType = "patchDomain_responseBody_2"
	PatchDomainResponseBodyTypePatchDomainResponseBody3 PatchDomainResponseBodyType = "patchDomain_responseBody_3"
)

type PatchDomainResponseBody struct {
	PatchDomainResponseBody1 *PatchDomainResponseBody1
	PatchDomainResponseBody2 *PatchDomainResponseBody2
	PatchDomainResponseBody3 *PatchDomainResponseBody3

	Type PatchDomainResponseBodyType
}

func CreatePatchDomainResponseBodyPatchDomainResponseBody1(patchDomainResponseBody1 PatchDomainResponseBody1) PatchDomainResponseBody {
	typ := PatchDomainResponseBodyTypePatchDomainResponseBody1

	return PatchDomainResponseBody{
		PatchDomainResponseBody1: &patchDomainResponseBody1,
		Type:                     typ,
	}
}

func CreatePatchDomainResponseBodyPatchDomainResponseBody2(patchDomainResponseBody2 PatchDomainResponseBody2) PatchDomainResponseBody {
	typ := PatchDomainResponseBodyTypePatchDomainResponseBody2

	return PatchDomainResponseBody{
		PatchDomainResponseBody2: &patchDomainResponseBody2,
		Type:                     typ,
	}
}

func CreatePatchDomainResponseBodyPatchDomainResponseBody3(patchDomainResponseBody3 PatchDomainResponseBody3) PatchDomainResponseBody {
	typ := PatchDomainResponseBodyTypePatchDomainResponseBody3

	return PatchDomainResponseBody{
		PatchDomainResponseBody3: &patchDomainResponseBody3,
		Type:                     typ,
	}
}

func (u *PatchDomainResponseBody) UnmarshalJSON(data []byte) error {

	var patchDomainResponseBody1 PatchDomainResponseBody1 = PatchDomainResponseBody1{}
	if err := utils.UnmarshalJSON(data, &patchDomainResponseBody1, "", true, true); err == nil {
		u.PatchDomainResponseBody1 = &patchDomainResponseBody1
		u.Type = PatchDomainResponseBodyTypePatchDomainResponseBody1
		return nil
	}

	var patchDomainResponseBody2 PatchDomainResponseBody2 = PatchDomainResponseBody2{}
	if err := utils.UnmarshalJSON(data, &patchDomainResponseBody2, "", true, true); err == nil {
		u.PatchDomainResponseBody2 = &patchDomainResponseBody2
		u.Type = PatchDomainResponseBodyTypePatchDomainResponseBody2
		return nil
	}

	var patchDomainResponseBody3 PatchDomainResponseBody3 = PatchDomainResponseBody3{}
	if err := utils.UnmarshalJSON(data, &patchDomainResponseBody3, "", true, true); err == nil {
		u.PatchDomainResponseBody3 = &patchDomainResponseBody3
		u.Type = PatchDomainResponseBodyTypePatchDomainResponseBody3
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for PatchDomainResponseBody", string(data))
}

func (u PatchDomainResponseBody) MarshalJSON() ([]byte, error) {
	if u.PatchDomainResponseBody1 != nil {
		return utils.MarshalJSON(u.PatchDomainResponseBody1, "", true)
	}

	if u.PatchDomainResponseBody2 != nil {
		return utils.MarshalJSON(u.PatchDomainResponseBody2, "", true)
	}

	if u.PatchDomainResponseBody3 != nil {
		return utils.MarshalJSON(u.PatchDomainResponseBody3, "", true)
	}

	return nil, errors.New("could not marshal union type PatchDomainResponseBody: all fields are null")
}

type PatchDomainResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	OneOf       *PatchDomainResponseBody
}

func (o *PatchDomainResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchDomainResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchDomainResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PatchDomainResponse) GetOneOf() *PatchDomainResponseBody {
	if o == nil {
		return nil
	}
	return o.OneOf
}
