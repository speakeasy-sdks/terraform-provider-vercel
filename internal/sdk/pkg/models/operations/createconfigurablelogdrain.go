// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// CreateConfigurableLogDrainRequestBodyDeliveryFormat - The delivery log format
type CreateConfigurableLogDrainRequestBodyDeliveryFormat string

const (
	CreateConfigurableLogDrainRequestBodyDeliveryFormatJSON   CreateConfigurableLogDrainRequestBodyDeliveryFormat = "json"
	CreateConfigurableLogDrainRequestBodyDeliveryFormatNdjson CreateConfigurableLogDrainRequestBodyDeliveryFormat = "ndjson"
)

func (e CreateConfigurableLogDrainRequestBodyDeliveryFormat) ToPointer() *CreateConfigurableLogDrainRequestBodyDeliveryFormat {
	return &e
}

func (e *CreateConfigurableLogDrainRequestBodyDeliveryFormat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "json":
		fallthrough
	case "ndjson":
		*e = CreateConfigurableLogDrainRequestBodyDeliveryFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateConfigurableLogDrainRequestBodyDeliveryFormat: %v", v)
	}
}

// CreateConfigurableLogDrainRequestBodyEnvironment - The environment of log drain
type CreateConfigurableLogDrainRequestBodyEnvironment string

const (
	CreateConfigurableLogDrainRequestBodyEnvironmentPreview    CreateConfigurableLogDrainRequestBodyEnvironment = "preview"
	CreateConfigurableLogDrainRequestBodyEnvironmentProduction CreateConfigurableLogDrainRequestBodyEnvironment = "production"
)

func (e CreateConfigurableLogDrainRequestBodyEnvironment) ToPointer() *CreateConfigurableLogDrainRequestBodyEnvironment {
	return &e
}

func (e *CreateConfigurableLogDrainRequestBodyEnvironment) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "preview":
		fallthrough
	case "production":
		*e = CreateConfigurableLogDrainRequestBodyEnvironment(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateConfigurableLogDrainRequestBodyEnvironment: %v", v)
	}
}

type CreateConfigurableLogDrainRequestBodySources string

const (
	CreateConfigurableLogDrainRequestBodySourcesStatic   CreateConfigurableLogDrainRequestBodySources = "static"
	CreateConfigurableLogDrainRequestBodySourcesLambda   CreateConfigurableLogDrainRequestBodySources = "lambda"
	CreateConfigurableLogDrainRequestBodySourcesBuild    CreateConfigurableLogDrainRequestBodySources = "build"
	CreateConfigurableLogDrainRequestBodySourcesEdge     CreateConfigurableLogDrainRequestBodySources = "edge"
	CreateConfigurableLogDrainRequestBodySourcesExternal CreateConfigurableLogDrainRequestBodySources = "external"
)

func (e CreateConfigurableLogDrainRequestBodySources) ToPointer() *CreateConfigurableLogDrainRequestBodySources {
	return &e
}

func (e *CreateConfigurableLogDrainRequestBodySources) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "static":
		fallthrough
	case "lambda":
		fallthrough
	case "build":
		fallthrough
	case "edge":
		fallthrough
	case "external":
		*e = CreateConfigurableLogDrainRequestBodySources(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateConfigurableLogDrainRequestBodySources: %v", v)
	}
}

type CreateConfigurableLogDrainRequestBody struct {
	// The branch regexp of log drain
	Branch *string `json:"branch,omitempty"`
	// The delivery log format
	DeliveryFormat CreateConfigurableLogDrainRequestBodyDeliveryFormat `json:"deliveryFormat"`
	// The environment of log drain
	Environment *CreateConfigurableLogDrainRequestBodyEnvironment `json:"environment,omitempty"`
	// Headers to be sent together with the request
	Headers    map[string]string `json:"headers,omitempty"`
	ProjectIds []string          `json:"projectIds,omitempty"`
	// Custom secret of log drain
	Secret  *string                                        `json:"secret,omitempty"`
	Sources []CreateConfigurableLogDrainRequestBodySources `json:"sources"`
	// The log drain url
	URL string `json:"url"`

	AdditionalProperties interface{} `json:"-"`
}
type _CreateConfigurableLogDrainRequestBody CreateConfigurableLogDrainRequestBody

func (c *CreateConfigurableLogDrainRequestBody) UnmarshalJSON(bs []byte) error {
	data := _CreateConfigurableLogDrainRequestBody{}

	if err := json.Unmarshal(bs, &data); err != nil {
		return err
	}
	*c = CreateConfigurableLogDrainRequestBody(data)

	additionalFields := make(map[string]interface{})

	if err := json.Unmarshal(bs, &additionalFields); err != nil {
		return err
	}
	delete(additionalFields, "branch")
	delete(additionalFields, "deliveryFormat")
	delete(additionalFields, "environment")
	delete(additionalFields, "headers")
	delete(additionalFields, "projectIds")
	delete(additionalFields, "secret")
	delete(additionalFields, "sources")
	delete(additionalFields, "url")

	c.AdditionalProperties = additionalFields

	return nil
}

func (c CreateConfigurableLogDrainRequestBody) MarshalJSON() ([]byte, error) {
	out := map[string]interface{}{}
	bs, err := json.Marshal(_CreateConfigurableLogDrainRequestBody(c))
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(bs), &out); err != nil {
		return nil, err
	}

	bs, err = json.Marshal(c.AdditionalProperties)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(bs), &out); err != nil {
		return nil, err
	}

	return json.Marshal(out)
}

type CreateConfigurableLogDrainRequest struct {
	RequestBody *CreateConfigurableLogDrainRequestBody `request:"mediaType=application/json"`
	// The Team identifier or slug to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
}

type CreateConfigurableLogDrain200ApplicationJSONCreatedFrom string

const (
	CreateConfigurableLogDrain200ApplicationJSONCreatedFromSelfServed CreateConfigurableLogDrain200ApplicationJSONCreatedFrom = "self-served"
)

func (e CreateConfigurableLogDrain200ApplicationJSONCreatedFrom) ToPointer() *CreateConfigurableLogDrain200ApplicationJSONCreatedFrom {
	return &e
}

func (e *CreateConfigurableLogDrain200ApplicationJSONCreatedFrom) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "self-served":
		*e = CreateConfigurableLogDrain200ApplicationJSONCreatedFrom(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateConfigurableLogDrain200ApplicationJSONCreatedFrom: %v", v)
	}
}

type CreateConfigurableLogDrain200ApplicationJSONDeliveryFormat string

const (
	CreateConfigurableLogDrain200ApplicationJSONDeliveryFormatJSON   CreateConfigurableLogDrain200ApplicationJSONDeliveryFormat = "json"
	CreateConfigurableLogDrain200ApplicationJSONDeliveryFormatNdjson CreateConfigurableLogDrain200ApplicationJSONDeliveryFormat = "ndjson"
	CreateConfigurableLogDrain200ApplicationJSONDeliveryFormatSyslog CreateConfigurableLogDrain200ApplicationJSONDeliveryFormat = "syslog"
)

func (e CreateConfigurableLogDrain200ApplicationJSONDeliveryFormat) ToPointer() *CreateConfigurableLogDrain200ApplicationJSONDeliveryFormat {
	return &e
}

func (e *CreateConfigurableLogDrain200ApplicationJSONDeliveryFormat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "json":
		fallthrough
	case "ndjson":
		fallthrough
	case "syslog":
		*e = CreateConfigurableLogDrain200ApplicationJSONDeliveryFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateConfigurableLogDrain200ApplicationJSONDeliveryFormat: %v", v)
	}
}

type CreateConfigurableLogDrain200ApplicationJSONDisabledReason string

const (
	CreateConfigurableLogDrain200ApplicationJSONDisabledReasonLogDrainHighErrorRate         CreateConfigurableLogDrain200ApplicationJSONDisabledReason = "log-drain-high-error-rate"
	CreateConfigurableLogDrain200ApplicationJSONDisabledReasonLogDrainsAddOnDisabledByOwner CreateConfigurableLogDrain200ApplicationJSONDisabledReason = "log-drains-add-on-disabled-by-owner"
	CreateConfigurableLogDrain200ApplicationJSONDisabledReasonDisabledByAdmin               CreateConfigurableLogDrain200ApplicationJSONDisabledReason = "disabled-by-admin"
	CreateConfigurableLogDrain200ApplicationJSONDisabledReasonAccountPlanDowngrade          CreateConfigurableLogDrain200ApplicationJSONDisabledReason = "account-plan-downgrade"
)

func (e CreateConfigurableLogDrain200ApplicationJSONDisabledReason) ToPointer() *CreateConfigurableLogDrain200ApplicationJSONDisabledReason {
	return &e
}

func (e *CreateConfigurableLogDrain200ApplicationJSONDisabledReason) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "log-drain-high-error-rate":
		fallthrough
	case "log-drains-add-on-disabled-by-owner":
		fallthrough
	case "disabled-by-admin":
		fallthrough
	case "account-plan-downgrade":
		*e = CreateConfigurableLogDrain200ApplicationJSONDisabledReason(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateConfigurableLogDrain200ApplicationJSONDisabledReason: %v", v)
	}
}

type CreateConfigurableLogDrain200ApplicationJSONEnvironment string

const (
	CreateConfigurableLogDrain200ApplicationJSONEnvironmentProduction CreateConfigurableLogDrain200ApplicationJSONEnvironment = "production"
	CreateConfigurableLogDrain200ApplicationJSONEnvironmentPreview    CreateConfigurableLogDrain200ApplicationJSONEnvironment = "preview"
)

func (e CreateConfigurableLogDrain200ApplicationJSONEnvironment) ToPointer() *CreateConfigurableLogDrain200ApplicationJSONEnvironment {
	return &e
}

func (e *CreateConfigurableLogDrain200ApplicationJSONEnvironment) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "production":
		fallthrough
	case "preview":
		*e = CreateConfigurableLogDrain200ApplicationJSONEnvironment(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateConfigurableLogDrain200ApplicationJSONEnvironment: %v", v)
	}
}

type CreateConfigurableLogDrain200ApplicationJSONSources string

const (
	CreateConfigurableLogDrain200ApplicationJSONSourcesStatic     CreateConfigurableLogDrain200ApplicationJSONSources = "static"
	CreateConfigurableLogDrain200ApplicationJSONSourcesLambda     CreateConfigurableLogDrain200ApplicationJSONSources = "lambda"
	CreateConfigurableLogDrain200ApplicationJSONSourcesBuild      CreateConfigurableLogDrain200ApplicationJSONSources = "build"
	CreateConfigurableLogDrain200ApplicationJSONSourcesEdge       CreateConfigurableLogDrain200ApplicationJSONSources = "edge"
	CreateConfigurableLogDrain200ApplicationJSONSourcesExternal   CreateConfigurableLogDrain200ApplicationJSONSources = "external"
	CreateConfigurableLogDrain200ApplicationJSONSourcesDeployment CreateConfigurableLogDrain200ApplicationJSONSources = "deployment"
)

func (e CreateConfigurableLogDrain200ApplicationJSONSources) ToPointer() *CreateConfigurableLogDrain200ApplicationJSONSources {
	return &e
}

func (e *CreateConfigurableLogDrain200ApplicationJSONSources) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "static":
		fallthrough
	case "lambda":
		fallthrough
	case "build":
		fallthrough
	case "edge":
		fallthrough
	case "external":
		fallthrough
	case "deployment":
		*e = CreateConfigurableLogDrain200ApplicationJSONSources(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateConfigurableLogDrain200ApplicationJSONSources: %v", v)
	}
}

type CreateConfigurableLogDrain200ApplicationJSONStatus string

const (
	CreateConfigurableLogDrain200ApplicationJSONStatusEnabled  CreateConfigurableLogDrain200ApplicationJSONStatus = "enabled"
	CreateConfigurableLogDrain200ApplicationJSONStatusDisabled CreateConfigurableLogDrain200ApplicationJSONStatus = "disabled"
	CreateConfigurableLogDrain200ApplicationJSONStatusErrored  CreateConfigurableLogDrain200ApplicationJSONStatus = "errored"
)

func (e CreateConfigurableLogDrain200ApplicationJSONStatus) ToPointer() *CreateConfigurableLogDrain200ApplicationJSONStatus {
	return &e
}

func (e *CreateConfigurableLogDrain200ApplicationJSONStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "enabled":
		fallthrough
	case "disabled":
		fallthrough
	case "errored":
		*e = CreateConfigurableLogDrain200ApplicationJSONStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateConfigurableLogDrain200ApplicationJSONStatus: %v", v)
	}
}

type CreateConfigurableLogDrain200ApplicationJSON struct {
	Branch              *string                                                     `json:"branch,omitempty"`
	ClientID            *string                                                     `json:"clientId,omitempty"`
	ConfigurationID     *string                                                     `json:"configurationId,omitempty"`
	CreatedAt           int64                                                       `json:"createdAt"`
	CreatedFrom         *CreateConfigurableLogDrain200ApplicationJSONCreatedFrom    `json:"createdFrom,omitempty"`
	DeliveryFormat      CreateConfigurableLogDrain200ApplicationJSONDeliveryFormat  `json:"deliveryFormat"`
	DisabledAt          *int64                                                      `json:"disabledAt,omitempty"`
	DisabledBy          *string                                                     `json:"disabledBy,omitempty"`
	DisabledReason      *CreateConfigurableLogDrain200ApplicationJSONDisabledReason `json:"disabledReason,omitempty"`
	Environment         *CreateConfigurableLogDrain200ApplicationJSONEnvironment    `json:"environment,omitempty"`
	FirstErrorTimestamp *int64                                                      `json:"firstErrorTimestamp,omitempty"`
	Headers             map[string]string                                           `json:"headers,omitempty"`
	ID                  string                                                      `json:"id"`
	Name                string                                                      `json:"name"`
	OwnerID             string                                                      `json:"ownerId"`
	ProjectIds          []string                                                    `json:"projectIds,omitempty"`
	// The secret to validate the log-drain payload
	Secret  *string                                               `json:"secret,omitempty"`
	Sources []CreateConfigurableLogDrain200ApplicationJSONSources `json:"sources,omitempty"`
	Status  *CreateConfigurableLogDrain200ApplicationJSONStatus   `json:"status,omitempty"`
	TeamID  *string                                               `json:"teamId,omitempty"`
	URL     string                                                `json:"url"`
}

type CreateConfigurableLogDrainResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse                                        *http.Response
	CreateConfigurableLogDrain200ApplicationJSONObject *CreateConfigurableLogDrain200ApplicationJSON
}
