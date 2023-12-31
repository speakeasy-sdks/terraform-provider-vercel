// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GetCheckRequest struct {
	// The check to fetch
	CheckID string `pathParam:"style=simple,explode=false,name=checkId"`
	// The deployment to get the check for.
	DeploymentID string `pathParam:"style=simple,explode=false,name=deploymentId"`
	// The Team identifier or slug to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
}

func (o *GetCheckRequest) GetCheckID() string {
	if o == nil {
		return ""
	}
	return o.CheckID
}

func (o *GetCheckRequest) GetDeploymentID() string {
	if o == nil {
		return ""
	}
	return o.DeploymentID
}

func (o *GetCheckRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

type GetCheckConclusion string

const (
	GetCheckConclusionCanceled  GetCheckConclusion = "canceled"
	GetCheckConclusionFailed    GetCheckConclusion = "failed"
	GetCheckConclusionNeutral   GetCheckConclusion = "neutral"
	GetCheckConclusionSucceeded GetCheckConclusion = "succeeded"
	GetCheckConclusionSkipped   GetCheckConclusion = "skipped"
	GetCheckConclusionStale     GetCheckConclusion = "stale"
)

func (e GetCheckConclusion) ToPointer() *GetCheckConclusion {
	return &e
}

func (e *GetCheckConclusion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "canceled":
		fallthrough
	case "failed":
		fallthrough
	case "neutral":
		fallthrough
	case "succeeded":
		fallthrough
	case "skipped":
		fallthrough
	case "stale":
		*e = GetCheckConclusion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetCheckConclusion: %v", v)
	}
}

type GetCheckSource string

const (
	GetCheckSourceWebVitals GetCheckSource = "web-vitals"
)

func (e GetCheckSource) ToPointer() *GetCheckSource {
	return &e
}

func (e *GetCheckSource) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "web-vitals":
		*e = GetCheckSource(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetCheckSource: %v", v)
	}
}

type GetCheckCLS struct {
	PreviousValue *int64         `json:"previousValue,omitempty"`
	Source        GetCheckSource `json:"source"`
	Value         *int64         `json:"value"`
}

func (o *GetCheckCLS) GetPreviousValue() *int64 {
	if o == nil {
		return nil
	}
	return o.PreviousValue
}

func (o *GetCheckCLS) GetSource() GetCheckSource {
	if o == nil {
		return GetCheckSource("")
	}
	return o.Source
}

func (o *GetCheckCLS) GetValue() *int64 {
	if o == nil {
		return nil
	}
	return o.Value
}

type GetCheckChecksSource string

const (
	GetCheckChecksSourceWebVitals GetCheckChecksSource = "web-vitals"
)

func (e GetCheckChecksSource) ToPointer() *GetCheckChecksSource {
	return &e
}

func (e *GetCheckChecksSource) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "web-vitals":
		*e = GetCheckChecksSource(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetCheckChecksSource: %v", v)
	}
}

type GetCheckFCP struct {
	PreviousValue *int64               `json:"previousValue,omitempty"`
	Source        GetCheckChecksSource `json:"source"`
	Value         *int64               `json:"value"`
}

func (o *GetCheckFCP) GetPreviousValue() *int64 {
	if o == nil {
		return nil
	}
	return o.PreviousValue
}

func (o *GetCheckFCP) GetSource() GetCheckChecksSource {
	if o == nil {
		return GetCheckChecksSource("")
	}
	return o.Source
}

func (o *GetCheckFCP) GetValue() *int64 {
	if o == nil {
		return nil
	}
	return o.Value
}

type GetCheckChecksResponseSource string

const (
	GetCheckChecksResponseSourceWebVitals GetCheckChecksResponseSource = "web-vitals"
)

func (e GetCheckChecksResponseSource) ToPointer() *GetCheckChecksResponseSource {
	return &e
}

func (e *GetCheckChecksResponseSource) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "web-vitals":
		*e = GetCheckChecksResponseSource(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetCheckChecksResponseSource: %v", v)
	}
}

type GetCheckLCP struct {
	PreviousValue *int64                       `json:"previousValue,omitempty"`
	Source        GetCheckChecksResponseSource `json:"source"`
	Value         *int64                       `json:"value"`
}

func (o *GetCheckLCP) GetPreviousValue() *int64 {
	if o == nil {
		return nil
	}
	return o.PreviousValue
}

func (o *GetCheckLCP) GetSource() GetCheckChecksResponseSource {
	if o == nil {
		return GetCheckChecksResponseSource("")
	}
	return o.Source
}

func (o *GetCheckLCP) GetValue() *int64 {
	if o == nil {
		return nil
	}
	return o.Value
}

type GetCheckChecksResponse200Source string

const (
	GetCheckChecksResponse200SourceWebVitals GetCheckChecksResponse200Source = "web-vitals"
)

func (e GetCheckChecksResponse200Source) ToPointer() *GetCheckChecksResponse200Source {
	return &e
}

func (e *GetCheckChecksResponse200Source) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "web-vitals":
		*e = GetCheckChecksResponse200Source(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetCheckChecksResponse200Source: %v", v)
	}
}

type GetCheckTBT struct {
	PreviousValue *int64                          `json:"previousValue,omitempty"`
	Source        GetCheckChecksResponse200Source `json:"source"`
	Value         *int64                          `json:"value"`
}

func (o *GetCheckTBT) GetPreviousValue() *int64 {
	if o == nil {
		return nil
	}
	return o.PreviousValue
}

func (o *GetCheckTBT) GetSource() GetCheckChecksResponse200Source {
	if o == nil {
		return GetCheckChecksResponse200Source("")
	}
	return o.Source
}

func (o *GetCheckTBT) GetValue() *int64 {
	if o == nil {
		return nil
	}
	return o.Value
}

type GetCheckChecksResponse200ApplicationJSONSource string

const (
	GetCheckChecksResponse200ApplicationJSONSourceWebVitals GetCheckChecksResponse200ApplicationJSONSource = "web-vitals"
)

func (e GetCheckChecksResponse200ApplicationJSONSource) ToPointer() *GetCheckChecksResponse200ApplicationJSONSource {
	return &e
}

func (e *GetCheckChecksResponse200ApplicationJSONSource) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "web-vitals":
		*e = GetCheckChecksResponse200ApplicationJSONSource(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetCheckChecksResponse200ApplicationJSONSource: %v", v)
	}
}

type GetCheckVirtualExperienceScore struct {
	PreviousValue *int64                                         `json:"previousValue,omitempty"`
	Source        GetCheckChecksResponse200ApplicationJSONSource `json:"source"`
	Value         *int64                                         `json:"value"`
}

func (o *GetCheckVirtualExperienceScore) GetPreviousValue() *int64 {
	if o == nil {
		return nil
	}
	return o.PreviousValue
}

func (o *GetCheckVirtualExperienceScore) GetSource() GetCheckChecksResponse200ApplicationJSONSource {
	if o == nil {
		return GetCheckChecksResponse200ApplicationJSONSource("")
	}
	return o.Source
}

func (o *GetCheckVirtualExperienceScore) GetValue() *int64 {
	if o == nil {
		return nil
	}
	return o.Value
}

type GetCheckMetrics struct {
	Cls                    GetCheckCLS                     `json:"CLS"`
	Fcp                    GetCheckFCP                     `json:"FCP"`
	Lcp                    GetCheckLCP                     `json:"LCP"`
	Tbt                    GetCheckTBT                     `json:"TBT"`
	VirtualExperienceScore *GetCheckVirtualExperienceScore `json:"virtualExperienceScore,omitempty"`
}

func (o *GetCheckMetrics) GetCls() GetCheckCLS {
	if o == nil {
		return GetCheckCLS{}
	}
	return o.Cls
}

func (o *GetCheckMetrics) GetFcp() GetCheckFCP {
	if o == nil {
		return GetCheckFCP{}
	}
	return o.Fcp
}

func (o *GetCheckMetrics) GetLcp() GetCheckLCP {
	if o == nil {
		return GetCheckLCP{}
	}
	return o.Lcp
}

func (o *GetCheckMetrics) GetTbt() GetCheckTBT {
	if o == nil {
		return GetCheckTBT{}
	}
	return o.Tbt
}

func (o *GetCheckMetrics) GetVirtualExperienceScore() *GetCheckVirtualExperienceScore {
	if o == nil {
		return nil
	}
	return o.VirtualExperienceScore
}

type GetCheckOutput struct {
	Metrics *GetCheckMetrics `json:"metrics,omitempty"`
}

func (o *GetCheckOutput) GetMetrics() *GetCheckMetrics {
	if o == nil {
		return nil
	}
	return o.Metrics
}

type GetCheckStatus string

const (
	GetCheckStatusRegistered GetCheckStatus = "registered"
	GetCheckStatusRunning    GetCheckStatus = "running"
	GetCheckStatusCompleted  GetCheckStatus = "completed"
)

func (e GetCheckStatus) ToPointer() *GetCheckStatus {
	return &e
}

func (e *GetCheckStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "registered":
		fallthrough
	case "running":
		fallthrough
	case "completed":
		*e = GetCheckStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetCheckStatus: %v", v)
	}
}

type GetCheckResponseBody struct {
	Blocking      bool                `json:"blocking"`
	CompletedAt   *int64              `json:"completedAt,omitempty"`
	Conclusion    *GetCheckConclusion `json:"conclusion,omitempty"`
	CreatedAt     int64               `json:"createdAt"`
	DeploymentID  string              `json:"deploymentId"`
	DetailsURL    *string             `json:"detailsUrl,omitempty"`
	ExternalID    *string             `json:"externalId,omitempty"`
	ID            string              `json:"id"`
	IntegrationID string              `json:"integrationId"`
	Name          string              `json:"name"`
	Output        *GetCheckOutput     `json:"output,omitempty"`
	Path          *string             `json:"path,omitempty"`
	Rerequestable *bool               `json:"rerequestable,omitempty"`
	StartedAt     *int64              `json:"startedAt,omitempty"`
	Status        GetCheckStatus      `json:"status"`
	UpdatedAt     int64               `json:"updatedAt"`
}

func (o *GetCheckResponseBody) GetBlocking() bool {
	if o == nil {
		return false
	}
	return o.Blocking
}

func (o *GetCheckResponseBody) GetCompletedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CompletedAt
}

func (o *GetCheckResponseBody) GetConclusion() *GetCheckConclusion {
	if o == nil {
		return nil
	}
	return o.Conclusion
}

func (o *GetCheckResponseBody) GetCreatedAt() int64 {
	if o == nil {
		return 0
	}
	return o.CreatedAt
}

func (o *GetCheckResponseBody) GetDeploymentID() string {
	if o == nil {
		return ""
	}
	return o.DeploymentID
}

func (o *GetCheckResponseBody) GetDetailsURL() *string {
	if o == nil {
		return nil
	}
	return o.DetailsURL
}

func (o *GetCheckResponseBody) GetExternalID() *string {
	if o == nil {
		return nil
	}
	return o.ExternalID
}

func (o *GetCheckResponseBody) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetCheckResponseBody) GetIntegrationID() string {
	if o == nil {
		return ""
	}
	return o.IntegrationID
}

func (o *GetCheckResponseBody) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *GetCheckResponseBody) GetOutput() *GetCheckOutput {
	if o == nil {
		return nil
	}
	return o.Output
}

func (o *GetCheckResponseBody) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

func (o *GetCheckResponseBody) GetRerequestable() *bool {
	if o == nil {
		return nil
	}
	return o.Rerequestable
}

func (o *GetCheckResponseBody) GetStartedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.StartedAt
}

func (o *GetCheckResponseBody) GetStatus() GetCheckStatus {
	if o == nil {
		return GetCheckStatus("")
	}
	return o.Status
}

func (o *GetCheckResponseBody) GetUpdatedAt() int64 {
	if o == nil {
		return 0
	}
	return o.UpdatedAt
}

type GetCheckResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	Object      *GetCheckResponseBody
}

func (o *GetCheckResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetCheckResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetCheckResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetCheckResponse) GetObject() *GetCheckResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
