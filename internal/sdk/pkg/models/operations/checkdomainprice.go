// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// QueryParamType - In which status of the domain the price needs to be checked.
type QueryParamType string

const (
	QueryParamTypeNew     QueryParamType = "new"
	QueryParamTypeRenewal QueryParamType = "renewal"
)

func (e QueryParamType) ToPointer() *QueryParamType {
	return &e
}

func (e *QueryParamType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "new":
		fallthrough
	case "renewal":
		*e = QueryParamType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for QueryParamType: %v", v)
	}
}

type CheckDomainPriceRequest struct {
	// The name of the domain for which the price needs to be checked.
	Name string `queryParam:"style=form,explode=true,name=name"`
	// The Team identifier or slug to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
	// In which status of the domain the price needs to be checked.
	Type *QueryParamType `queryParam:"style=form,explode=true,name=type"`
}

func (o *CheckDomainPriceRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CheckDomainPriceRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

func (o *CheckDomainPriceRequest) GetType() *QueryParamType {
	if o == nil {
		return nil
	}
	return o.Type
}

// CheckDomainPriceResponseBody - Successful response which returns the price of the domain and the period.
type CheckDomainPriceResponseBody struct {
	// The number of years the domain could be held before paying again.
	Period int64 `json:"period"`
	// The domain price in USD.
	Price int64 `json:"price"`
}

func (o *CheckDomainPriceResponseBody) GetPeriod() int64 {
	if o == nil {
		return 0
	}
	return o.Period
}

func (o *CheckDomainPriceResponseBody) GetPrice() int64 {
	if o == nil {
		return 0
	}
	return o.Price
}

type CheckDomainPriceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful response which returns the price of the domain and the period.
	Object *CheckDomainPriceResponseBody
}

func (o *CheckDomainPriceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CheckDomainPriceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CheckDomainPriceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CheckDomainPriceResponse) GetObject() *CheckDomainPriceResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
