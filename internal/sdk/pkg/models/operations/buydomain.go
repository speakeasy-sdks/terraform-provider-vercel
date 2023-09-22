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

type BuyDomainRequest struct {
	RequestBody *BuyDomainRequestBody `request:"mediaType=application/json"`
	// The Team identifier or slug to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
}

type BuyDomain202ApplicationJSONDomain struct {
	Created  int64    `json:"created"`
	Ns       []string `json:"ns"`
	Pending  bool     `json:"pending"`
	UID      string   `json:"uid"`
	Verified bool     `json:"verified"`
}

// BuyDomain202ApplicationJSON - Domain purchase is being processed asynchronously.
type BuyDomain202ApplicationJSON struct {
	Domain BuyDomain202ApplicationJSONDomain `json:"domain"`
}

type BuyDomain201ApplicationJSONDomain struct {
	Created  int64    `json:"created"`
	Ns       []string `json:"ns"`
	Pending  bool     `json:"pending"`
	UID      string   `json:"uid"`
	Verified bool     `json:"verified"`
}

// BuyDomain201ApplicationJSON - Successful response for purchasing a Domain.
type BuyDomain201ApplicationJSON struct {
	Domain BuyDomain201ApplicationJSONDomain `json:"domain"`
}

type BuyDomainResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful response for purchasing a Domain.
	BuyDomain201ApplicationJSONObject *BuyDomain201ApplicationJSON
	// Domain purchase is being processed asynchronously.
	BuyDomain202ApplicationJSONObject *BuyDomain202ApplicationJSON
}
