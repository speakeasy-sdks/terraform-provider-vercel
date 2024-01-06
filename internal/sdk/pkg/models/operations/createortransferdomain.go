// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zchee/terraform-provider-vercel/internal/sdk/pkg/utils"
	"net/http"
)

// CreateOrTransferDomain3 - transfer-in
type CreateOrTransferDomain3 struct {
	// The authorization code assigned to the domain.
	AuthCode *string `json:"authCode,omitempty"`
	// The price you expect to be charged for the required 1 year renewal.
	ExpectedPrice *int64 `json:"expectedPrice,omitempty"`
	// The domain operation to perform. It can be either `add` or `transfer-in`.
	Method string `json:"method"`
	// The domain name you want to add.
	Name string `json:"name"`
}

func (o *CreateOrTransferDomain3) GetAuthCode() *string {
	if o == nil {
		return nil
	}
	return o.AuthCode
}

func (o *CreateOrTransferDomain3) GetExpectedPrice() *int64 {
	if o == nil {
		return nil
	}
	return o.ExpectedPrice
}

func (o *CreateOrTransferDomain3) GetMethod() string {
	if o == nil {
		return ""
	}
	return o.Method
}

func (o *CreateOrTransferDomain3) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

// CreateOrTransferDomain2 - move-in
type CreateOrTransferDomain2 struct {
	// The domain operation to perform. It can be either `add` or `transfer-in`.
	Method string `json:"method"`
	// The domain name you want to add.
	Name string `json:"name"`
	// The move-in token from Move Requested email.
	Token *string `json:"token,omitempty"`
}

func (o *CreateOrTransferDomain2) GetMethod() string {
	if o == nil {
		return ""
	}
	return o.Method
}

func (o *CreateOrTransferDomain2) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateOrTransferDomain2) GetToken() *string {
	if o == nil {
		return nil
	}
	return o.Token
}

// CreateOrTransferDomain1 - add
type CreateOrTransferDomain1 struct {
	// Whether the domain has the Vercel Edge Network enabled or not.
	CdnEnabled *bool `json:"cdnEnabled,omitempty"`
	// The domain operation to perform. It can be either `add` or `transfer-in`.
	Method *string `json:"method,omitempty"`
	// The domain name you want to add.
	Name string `json:"name"`
	Zone *bool  `json:"zone,omitempty"`
}

func (o *CreateOrTransferDomain1) GetCdnEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.CdnEnabled
}

func (o *CreateOrTransferDomain1) GetMethod() *string {
	if o == nil {
		return nil
	}
	return o.Method
}

func (o *CreateOrTransferDomain1) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateOrTransferDomain1) GetZone() *bool {
	if o == nil {
		return nil
	}
	return o.Zone
}

type CreateOrTransferDomainRequestBodyType string

const (
	CreateOrTransferDomainRequestBodyTypeCreateOrTransferDomain1 CreateOrTransferDomainRequestBodyType = "createOrTransferDomain_1"
	CreateOrTransferDomainRequestBodyTypeCreateOrTransferDomain2 CreateOrTransferDomainRequestBodyType = "createOrTransferDomain_2"
	CreateOrTransferDomainRequestBodyTypeCreateOrTransferDomain3 CreateOrTransferDomainRequestBodyType = "createOrTransferDomain_3"
)

type CreateOrTransferDomainRequestBody struct {
	CreateOrTransferDomain1 *CreateOrTransferDomain1
	CreateOrTransferDomain2 *CreateOrTransferDomain2
	CreateOrTransferDomain3 *CreateOrTransferDomain3

	Type CreateOrTransferDomainRequestBodyType
}

func CreateCreateOrTransferDomainRequestBodyCreateOrTransferDomain1(createOrTransferDomain1 CreateOrTransferDomain1) CreateOrTransferDomainRequestBody {
	typ := CreateOrTransferDomainRequestBodyTypeCreateOrTransferDomain1

	return CreateOrTransferDomainRequestBody{
		CreateOrTransferDomain1: &createOrTransferDomain1,
		Type:                    typ,
	}
}

func CreateCreateOrTransferDomainRequestBodyCreateOrTransferDomain2(createOrTransferDomain2 CreateOrTransferDomain2) CreateOrTransferDomainRequestBody {
	typ := CreateOrTransferDomainRequestBodyTypeCreateOrTransferDomain2

	return CreateOrTransferDomainRequestBody{
		CreateOrTransferDomain2: &createOrTransferDomain2,
		Type:                    typ,
	}
}

func CreateCreateOrTransferDomainRequestBodyCreateOrTransferDomain3(createOrTransferDomain3 CreateOrTransferDomain3) CreateOrTransferDomainRequestBody {
	typ := CreateOrTransferDomainRequestBodyTypeCreateOrTransferDomain3

	return CreateOrTransferDomainRequestBody{
		CreateOrTransferDomain3: &createOrTransferDomain3,
		Type:                    typ,
	}
}

func (u *CreateOrTransferDomainRequestBody) UnmarshalJSON(data []byte) error {

	createOrTransferDomain2 := new(CreateOrTransferDomain2)
	if err := utils.UnmarshalJSON(data, &createOrTransferDomain2, "", true, true); err == nil {
		u.CreateOrTransferDomain2 = createOrTransferDomain2
		u.Type = CreateOrTransferDomainRequestBodyTypeCreateOrTransferDomain2
		return nil
	}

	createOrTransferDomain1 := new(CreateOrTransferDomain1)
	if err := utils.UnmarshalJSON(data, &createOrTransferDomain1, "", true, true); err == nil {
		u.CreateOrTransferDomain1 = createOrTransferDomain1
		u.Type = CreateOrTransferDomainRequestBodyTypeCreateOrTransferDomain1
		return nil
	}

	createOrTransferDomain3 := new(CreateOrTransferDomain3)
	if err := utils.UnmarshalJSON(data, &createOrTransferDomain3, "", true, true); err == nil {
		u.CreateOrTransferDomain3 = createOrTransferDomain3
		u.Type = CreateOrTransferDomainRequestBodyTypeCreateOrTransferDomain3
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u CreateOrTransferDomainRequestBody) MarshalJSON() ([]byte, error) {
	if u.CreateOrTransferDomain1 != nil {
		return utils.MarshalJSON(u.CreateOrTransferDomain1, "", true)
	}

	if u.CreateOrTransferDomain2 != nil {
		return utils.MarshalJSON(u.CreateOrTransferDomain2, "", true)
	}

	if u.CreateOrTransferDomain3 != nil {
		return utils.MarshalJSON(u.CreateOrTransferDomain3, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type CreateOrTransferDomainRequest struct {
	RequestBody *CreateOrTransferDomainRequestBody `request:"mediaType=application/json"`
	// The Team identifier or slug to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
}

func (o *CreateOrTransferDomainRequest) GetRequestBody() *CreateOrTransferDomainRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *CreateOrTransferDomainRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

// CreateOrTransferDomainCreator - An object containing information of the domain creator, including the user's id, username, and email.
type CreateOrTransferDomainCreator struct {
	CustomerID       *string `json:"customerId,omitempty"`
	Email            string  `json:"email"`
	ID               string  `json:"id"`
	IsDomainReseller *bool   `json:"isDomainReseller,omitempty"`
	Username         string  `json:"username"`
}

func (o *CreateOrTransferDomainCreator) GetCustomerID() *string {
	if o == nil {
		return nil
	}
	return o.CustomerID
}

func (o *CreateOrTransferDomainCreator) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

func (o *CreateOrTransferDomainCreator) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *CreateOrTransferDomainCreator) GetIsDomainReseller() *bool {
	if o == nil {
		return nil
	}
	return o.IsDomainReseller
}

func (o *CreateOrTransferDomainCreator) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

// ServiceType - The type of service the domain is handled by. `external` if the DNS is externally handled, `zeit.world` if handled with Vercel, or `na` if the service is not available.
type ServiceType string

const (
	ServiceTypeZeitWorld ServiceType = "zeit.world"
	ServiceTypeExternal  ServiceType = "external"
	ServiceTypeNa        ServiceType = "na"
)

func (e ServiceType) ToPointer() *ServiceType {
	return &e
}

func (e *ServiceType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "zeit.world":
		fallthrough
	case "external":
		fallthrough
	case "na":
		*e = ServiceType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ServiceType: %v", v)
	}
}

type CreateOrTransferDomainDomain struct {
	// If it was purchased through Vercel, the timestamp in milliseconds when it was purchased.
	BoughtAt *int64 `json:"boughtAt"`
	// Timestamp in milliseconds when the domain was created in the registry.
	CreatedAt int64 `json:"createdAt"`
	// An object containing information of the domain creator, including the user's id, username, and email.
	Creator CreateOrTransferDomainCreator `json:"creator"`
	// A list of custom nameservers for the domain to point to. Only applies to domains purchased with Vercel.
	CustomNameservers []string `json:"customNameservers,omitempty"`
	// Timestamp in milliseconds at which the domain is set to expire. `null` if not bought with Vercel.
	ExpiresAt *int64 `json:"expiresAt"`
	// The unique identifier of the domain.
	ID string `json:"id"`
	// A list of the intended nameservers for the domain to point to Vercel DNS.
	IntendedNameservers []string `json:"intendedNameservers"`
	// The domain name.
	Name string `json:"name"`
	// A list of the current nameservers of the domain.
	Nameservers []string `json:"nameservers"`
	// Timestamp in milliseconds at which the domain was ordered.
	OrderedAt *int64 `json:"orderedAt,omitempty"`
	// Indicates whether the domain is set to automatically renew.
	Renew *bool `json:"renew,omitempty"`
	// The type of service the domain is handled by. `external` if the DNS is externally handled, `zeit.world` if handled with Vercel, or `na` if the service is not available.
	ServiceType ServiceType `json:"serviceType"`
	// If transferred into Vercel, timestamp in milliseconds when the domain transfer was initiated.
	TransferStartedAt *int64 `json:"transferStartedAt,omitempty"`
	// Timestamp in milliseconds at which the domain was successfully transferred into Vercel. `null` if the transfer is still processing or was never transferred in.
	TransferredAt *int64 `json:"transferredAt,omitempty"`
	// If the domain has the ownership verified.
	Verified bool `json:"verified"`
}

func (o *CreateOrTransferDomainDomain) GetBoughtAt() *int64 {
	if o == nil {
		return nil
	}
	return o.BoughtAt
}

func (o *CreateOrTransferDomainDomain) GetCreatedAt() int64 {
	if o == nil {
		return 0
	}
	return o.CreatedAt
}

func (o *CreateOrTransferDomainDomain) GetCreator() CreateOrTransferDomainCreator {
	if o == nil {
		return CreateOrTransferDomainCreator{}
	}
	return o.Creator
}

func (o *CreateOrTransferDomainDomain) GetCustomNameservers() []string {
	if o == nil {
		return nil
	}
	return o.CustomNameservers
}

func (o *CreateOrTransferDomainDomain) GetExpiresAt() *int64 {
	if o == nil {
		return nil
	}
	return o.ExpiresAt
}

func (o *CreateOrTransferDomainDomain) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *CreateOrTransferDomainDomain) GetIntendedNameservers() []string {
	if o == nil {
		return []string{}
	}
	return o.IntendedNameservers
}

func (o *CreateOrTransferDomainDomain) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateOrTransferDomainDomain) GetNameservers() []string {
	if o == nil {
		return []string{}
	}
	return o.Nameservers
}

func (o *CreateOrTransferDomainDomain) GetOrderedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.OrderedAt
}

func (o *CreateOrTransferDomainDomain) GetRenew() *bool {
	if o == nil {
		return nil
	}
	return o.Renew
}

func (o *CreateOrTransferDomainDomain) GetServiceType() ServiceType {
	if o == nil {
		return ServiceType("")
	}
	return o.ServiceType
}

func (o *CreateOrTransferDomainDomain) GetTransferStartedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.TransferStartedAt
}

func (o *CreateOrTransferDomainDomain) GetTransferredAt() *int64 {
	if o == nil {
		return nil
	}
	return o.TransferredAt
}

func (o *CreateOrTransferDomainDomain) GetVerified() bool {
	if o == nil {
		return false
	}
	return o.Verified
}

type CreateOrTransferDomainResponseBody struct {
	Domain CreateOrTransferDomainDomain `json:"domain"`
}

func (o *CreateOrTransferDomainResponseBody) GetDomain() CreateOrTransferDomainDomain {
	if o == nil {
		return CreateOrTransferDomainDomain{}
	}
	return o.Domain
}

type CreateOrTransferDomainResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	Object      *CreateOrTransferDomainResponseBody
}

func (o *CreateOrTransferDomainResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateOrTransferDomainResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateOrTransferDomainResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateOrTransferDomainResponse) GetObject() *CreateOrTransferDomainResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
