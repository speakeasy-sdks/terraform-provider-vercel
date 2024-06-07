// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GetAllLogDrainsRequest struct {
	ProjectID *string `queryParam:"style=form,explode=true,name=projectId"`
	// The Team slug to perform the request on behalf of.
	Slug *string `queryParam:"style=form,explode=true,name=slug"`
	// The Team identifier to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
}

func (o *GetAllLogDrainsRequest) GetProjectID() *string {
	if o == nil {
		return nil
	}
	return o.ProjectID
}

func (o *GetAllLogDrainsRequest) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *GetAllLogDrainsRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

type GetAllLogDrainsCompression string

const (
	GetAllLogDrainsCompressionGzip GetAllLogDrainsCompression = "gzip"
	GetAllLogDrainsCompressionZstd GetAllLogDrainsCompression = "zstd"
	GetAllLogDrainsCompressionNone GetAllLogDrainsCompression = "none"
)

func (e GetAllLogDrainsCompression) ToPointer() *GetAllLogDrainsCompression {
	return &e
}
func (e *GetAllLogDrainsCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "gzip":
		fallthrough
	case "zstd":
		fallthrough
	case "none":
		*e = GetAllLogDrainsCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetAllLogDrainsCompression: %v", v)
	}
}

type GetAllLogDrainsCreatedFrom string

const (
	GetAllLogDrainsCreatedFromSelfServed GetAllLogDrainsCreatedFrom = "self-served"
)

func (e GetAllLogDrainsCreatedFrom) ToPointer() *GetAllLogDrainsCreatedFrom {
	return &e
}
func (e *GetAllLogDrainsCreatedFrom) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "self-served":
		*e = GetAllLogDrainsCreatedFrom(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetAllLogDrainsCreatedFrom: %v", v)
	}
}

type GetAllLogDrainsDeliveryFormat string

const (
	GetAllLogDrainsDeliveryFormatJSON   GetAllLogDrainsDeliveryFormat = "json"
	GetAllLogDrainsDeliveryFormatNdjson GetAllLogDrainsDeliveryFormat = "ndjson"
	GetAllLogDrainsDeliveryFormatSyslog GetAllLogDrainsDeliveryFormat = "syslog"
)

func (e GetAllLogDrainsDeliveryFormat) ToPointer() *GetAllLogDrainsDeliveryFormat {
	return &e
}
func (e *GetAllLogDrainsDeliveryFormat) UnmarshalJSON(data []byte) error {
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
		*e = GetAllLogDrainsDeliveryFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetAllLogDrainsDeliveryFormat: %v", v)
	}
}

type GetAllLogDrainsDisabledReason string

const (
	GetAllLogDrainsDisabledReasonDisabledByOwner     GetAllLogDrainsDisabledReason = "disabled-by-owner"
	GetAllLogDrainsDisabledReasonFeatureNotAvailable GetAllLogDrainsDisabledReason = "feature-not-available"
	GetAllLogDrainsDisabledReasonDisabledByAdmin     GetAllLogDrainsDisabledReason = "disabled-by-admin"
)

func (e GetAllLogDrainsDisabledReason) ToPointer() *GetAllLogDrainsDisabledReason {
	return &e
}
func (e *GetAllLogDrainsDisabledReason) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "disabled-by-owner":
		fallthrough
	case "feature-not-available":
		fallthrough
	case "disabled-by-admin":
		*e = GetAllLogDrainsDisabledReason(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetAllLogDrainsDisabledReason: %v", v)
	}
}

type GetAllLogDrainsEnvironments string

const (
	GetAllLogDrainsEnvironmentsProduction GetAllLogDrainsEnvironments = "production"
	GetAllLogDrainsEnvironmentsPreview    GetAllLogDrainsEnvironments = "preview"
)

func (e GetAllLogDrainsEnvironments) ToPointer() *GetAllLogDrainsEnvironments {
	return &e
}
func (e *GetAllLogDrainsEnvironments) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "production":
		fallthrough
	case "preview":
		*e = GetAllLogDrainsEnvironments(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetAllLogDrainsEnvironments: %v", v)
	}
}

type GetAllLogDrainsSources string

const (
	GetAllLogDrainsSourcesBuild    GetAllLogDrainsSources = "build"
	GetAllLogDrainsSourcesEdge     GetAllLogDrainsSources = "edge"
	GetAllLogDrainsSourcesLambda   GetAllLogDrainsSources = "lambda"
	GetAllLogDrainsSourcesStatic   GetAllLogDrainsSources = "static"
	GetAllLogDrainsSourcesExternal GetAllLogDrainsSources = "external"
)

func (e GetAllLogDrainsSources) ToPointer() *GetAllLogDrainsSources {
	return &e
}
func (e *GetAllLogDrainsSources) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "build":
		fallthrough
	case "edge":
		fallthrough
	case "lambda":
		fallthrough
	case "static":
		fallthrough
	case "external":
		*e = GetAllLogDrainsSources(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetAllLogDrainsSources: %v", v)
	}
}

type GetAllLogDrainsStatus string

const (
	GetAllLogDrainsStatusEnabled  GetAllLogDrainsStatus = "enabled"
	GetAllLogDrainsStatusDisabled GetAllLogDrainsStatus = "disabled"
	GetAllLogDrainsStatusErrored  GetAllLogDrainsStatus = "errored"
)

func (e GetAllLogDrainsStatus) ToPointer() *GetAllLogDrainsStatus {
	return &e
}
func (e *GetAllLogDrainsStatus) UnmarshalJSON(data []byte) error {
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
		*e = GetAllLogDrainsStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetAllLogDrainsStatus: %v", v)
	}
}

type GetAllLogDrainsResponseBody struct {
	ClientID            *string                        `json:"clientId,omitempty"`
	Compression         *GetAllLogDrainsCompression    `json:"compression,omitempty"`
	ConfigurationID     *string                        `json:"configurationId,omitempty"`
	CreatedAt           float64                        `json:"createdAt"`
	CreatedFrom         *GetAllLogDrainsCreatedFrom    `json:"createdFrom,omitempty"`
	DeliveryFormat      GetAllLogDrainsDeliveryFormat  `json:"deliveryFormat"`
	DisabledAt          *float64                       `json:"disabledAt,omitempty"`
	DisabledBy          *string                        `json:"disabledBy,omitempty"`
	DisabledReason      *GetAllLogDrainsDisabledReason `json:"disabledReason,omitempty"`
	Environments        []GetAllLogDrainsEnvironments  `json:"environments"`
	FirstErrorTimestamp *float64                       `json:"firstErrorTimestamp,omitempty"`
	Headers             map[string]string              `json:"headers,omitempty"`
	ID                  string                         `json:"id"`
	Name                string                         `json:"name"`
	OwnerID             string                         `json:"ownerId"`
	ProjectIds          []string                       `json:"projectIds,omitempty"`
	SamplingRate        *float64                       `json:"samplingRate,omitempty"`
	Secret              string                         `json:"secret"`
	Sources             []GetAllLogDrainsSources       `json:"sources,omitempty"`
	Status              *GetAllLogDrainsStatus         `json:"status,omitempty"`
	TeamID              *string                        `json:"teamId,omitempty"`
	URL                 string                         `json:"url"`
}

func (o *GetAllLogDrainsResponseBody) GetClientID() *string {
	if o == nil {
		return nil
	}
	return o.ClientID
}

func (o *GetAllLogDrainsResponseBody) GetCompression() *GetAllLogDrainsCompression {
	if o == nil {
		return nil
	}
	return o.Compression
}

func (o *GetAllLogDrainsResponseBody) GetConfigurationID() *string {
	if o == nil {
		return nil
	}
	return o.ConfigurationID
}

func (o *GetAllLogDrainsResponseBody) GetCreatedAt() float64 {
	if o == nil {
		return 0.0
	}
	return o.CreatedAt
}

func (o *GetAllLogDrainsResponseBody) GetCreatedFrom() *GetAllLogDrainsCreatedFrom {
	if o == nil {
		return nil
	}
	return o.CreatedFrom
}

func (o *GetAllLogDrainsResponseBody) GetDeliveryFormat() GetAllLogDrainsDeliveryFormat {
	if o == nil {
		return GetAllLogDrainsDeliveryFormat("")
	}
	return o.DeliveryFormat
}

func (o *GetAllLogDrainsResponseBody) GetDisabledAt() *float64 {
	if o == nil {
		return nil
	}
	return o.DisabledAt
}

func (o *GetAllLogDrainsResponseBody) GetDisabledBy() *string {
	if o == nil {
		return nil
	}
	return o.DisabledBy
}

func (o *GetAllLogDrainsResponseBody) GetDisabledReason() *GetAllLogDrainsDisabledReason {
	if o == nil {
		return nil
	}
	return o.DisabledReason
}

func (o *GetAllLogDrainsResponseBody) GetEnvironments() []GetAllLogDrainsEnvironments {
	if o == nil {
		return []GetAllLogDrainsEnvironments{}
	}
	return o.Environments
}

func (o *GetAllLogDrainsResponseBody) GetFirstErrorTimestamp() *float64 {
	if o == nil {
		return nil
	}
	return o.FirstErrorTimestamp
}

func (o *GetAllLogDrainsResponseBody) GetHeaders() map[string]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *GetAllLogDrainsResponseBody) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetAllLogDrainsResponseBody) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *GetAllLogDrainsResponseBody) GetOwnerID() string {
	if o == nil {
		return ""
	}
	return o.OwnerID
}

func (o *GetAllLogDrainsResponseBody) GetProjectIds() []string {
	if o == nil {
		return nil
	}
	return o.ProjectIds
}

func (o *GetAllLogDrainsResponseBody) GetSamplingRate() *float64 {
	if o == nil {
		return nil
	}
	return o.SamplingRate
}

func (o *GetAllLogDrainsResponseBody) GetSecret() string {
	if o == nil {
		return ""
	}
	return o.Secret
}

func (o *GetAllLogDrainsResponseBody) GetSources() []GetAllLogDrainsSources {
	if o == nil {
		return nil
	}
	return o.Sources
}

func (o *GetAllLogDrainsResponseBody) GetStatus() *GetAllLogDrainsStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *GetAllLogDrainsResponseBody) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

func (o *GetAllLogDrainsResponseBody) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

type GetAllLogDrainsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	Classes     []GetAllLogDrainsResponseBody
}

func (o *GetAllLogDrainsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAllLogDrainsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAllLogDrainsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetAllLogDrainsResponse) GetClasses() []GetAllLogDrainsResponseBody {
	if o == nil {
		return nil
	}
	return o.Classes
}
