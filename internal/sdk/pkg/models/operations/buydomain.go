// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type BuyDomainRequestBody struct {
	// The price you expect to be charged for the purchase.
	ExpectedPrice *int64 `json:"expectedPrice,omitempty"`
	// The domain name to purchase.
	Name string `json:"name"`
	// Indicates whether the domain should be automatically renewed.
	Renew *bool `json:"renew,omitempty"`
}

func (o *BuyDomainRequestBody) GetExpectedPrice() *int64 {
	if o == nil {
		return nil
	}
	return o.ExpectedPrice
}

func (o *BuyDomainRequestBody) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *BuyDomainRequestBody) GetRenew() *bool {
	if o == nil {
		return nil
	}
	return o.Renew
}

type BuyDomainRequest struct {
	RequestBody *BuyDomainRequestBody `request:"mediaType=application/json"`
	// The Team identifier or slug to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
}

func (o *BuyDomainRequest) GetRequestBody() *BuyDomainRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *BuyDomainRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

type BuyDomainDomain struct {
	Created  int64    `json:"created"`
	Ns       []string `json:"ns"`
	Pending  bool     `json:"pending"`
	UID      string   `json:"uid"`
	Verified bool     `json:"verified"`
}

func (o *BuyDomainDomain) GetCreated() int64 {
	if o == nil {
		return 0
	}
	return o.Created
}

func (o *BuyDomainDomain) GetNs() []string {
	if o == nil {
		return []string{}
	}
	return o.Ns
}

func (o *BuyDomainDomain) GetPending() bool {
	if o == nil {
		return false
	}
	return o.Pending
}

func (o *BuyDomainDomain) GetUID() string {
	if o == nil {
		return ""
	}
	return o.UID
}

func (o *BuyDomainDomain) GetVerified() bool {
	if o == nil {
		return false
	}
	return o.Verified
}

// BuyDomainDomainsResponseBody - Domain purchase is being processed asynchronously.
type BuyDomainDomainsResponseBody struct {
	Domain BuyDomainDomain `json:"domain"`
}

func (o *BuyDomainDomainsResponseBody) GetDomain() BuyDomainDomain {
	if o == nil {
		return BuyDomainDomain{}
	}
	return o.Domain
}

type Domain struct {
	Created  int64    `json:"created"`
	Ns       []string `json:"ns"`
	Pending  bool     `json:"pending"`
	UID      string   `json:"uid"`
	Verified bool     `json:"verified"`
}

func (o *Domain) GetCreated() int64 {
	if o == nil {
		return 0
	}
	return o.Created
}

func (o *Domain) GetNs() []string {
	if o == nil {
		return []string{}
	}
	return o.Ns
}

func (o *Domain) GetPending() bool {
	if o == nil {
		return false
	}
	return o.Pending
}

func (o *Domain) GetUID() string {
	if o == nil {
		return ""
	}
	return o.UID
}

func (o *Domain) GetVerified() bool {
	if o == nil {
		return false
	}
	return o.Verified
}

// BuyDomainResponseBody - Successful response for purchasing a Domain.
type BuyDomainResponseBody struct {
	Domain Domain `json:"domain"`
}

func (o *BuyDomainResponseBody) GetDomain() Domain {
	if o == nil {
		return Domain{}
	}
	return o.Domain
}

type BuyDomainResponse struct {
	// Successful response for purchasing a Domain.
	TwoHundredAndOneApplicationJSONObject *BuyDomainResponseBody
	// Domain purchase is being processed asynchronously.
	TwoHundredAndTwoApplicationJSONObject *BuyDomainDomainsResponseBody
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *BuyDomainResponse) GetTwoHundredAndOneApplicationJSONObject() *BuyDomainResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredAndOneApplicationJSONObject
}

func (o *BuyDomainResponse) GetTwoHundredAndTwoApplicationJSONObject() *BuyDomainDomainsResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredAndTwoApplicationJSONObject
}

func (o *BuyDomainResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *BuyDomainResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *BuyDomainResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
