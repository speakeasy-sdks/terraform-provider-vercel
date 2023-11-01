// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GetAllChecksRequest struct {
	// The deployment to get all checks for
	DeploymentID string `pathParam:"style=simple,explode=false,name=deploymentId"`
	// The Team identifier or slug to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
}

func (o *GetAllChecksRequest) GetDeploymentID() string {
	if o == nil {
		return ""
	}
	return o.DeploymentID
}

func (o *GetAllChecksRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

type GetAllChecks200ApplicationJSONChecksConclusion string

const (
	GetAllChecks200ApplicationJSONChecksConclusionCanceled  GetAllChecks200ApplicationJSONChecksConclusion = "canceled"
	GetAllChecks200ApplicationJSONChecksConclusionFailed    GetAllChecks200ApplicationJSONChecksConclusion = "failed"
	GetAllChecks200ApplicationJSONChecksConclusionNeutral   GetAllChecks200ApplicationJSONChecksConclusion = "neutral"
	GetAllChecks200ApplicationJSONChecksConclusionSucceeded GetAllChecks200ApplicationJSONChecksConclusion = "succeeded"
	GetAllChecks200ApplicationJSONChecksConclusionSkipped   GetAllChecks200ApplicationJSONChecksConclusion = "skipped"
	GetAllChecks200ApplicationJSONChecksConclusionStale     GetAllChecks200ApplicationJSONChecksConclusion = "stale"
)

func (e GetAllChecks200ApplicationJSONChecksConclusion) ToPointer() *GetAllChecks200ApplicationJSONChecksConclusion {
	return &e
}

func (e *GetAllChecks200ApplicationJSONChecksConclusion) UnmarshalJSON(data []byte) error {
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
		*e = GetAllChecks200ApplicationJSONChecksConclusion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetAllChecks200ApplicationJSONChecksConclusion: %v", v)
	}
}

type GetAllChecks200ApplicationJSONChecksOutputMetricsCLSSource string

const (
	GetAllChecks200ApplicationJSONChecksOutputMetricsCLSSourceWebVitals GetAllChecks200ApplicationJSONChecksOutputMetricsCLSSource = "web-vitals"
)

func (e GetAllChecks200ApplicationJSONChecksOutputMetricsCLSSource) ToPointer() *GetAllChecks200ApplicationJSONChecksOutputMetricsCLSSource {
	return &e
}

func (e *GetAllChecks200ApplicationJSONChecksOutputMetricsCLSSource) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "web-vitals":
		*e = GetAllChecks200ApplicationJSONChecksOutputMetricsCLSSource(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetAllChecks200ApplicationJSONChecksOutputMetricsCLSSource: %v", v)
	}
}

type GetAllChecks200ApplicationJSONChecksOutputMetricsCLS struct {
	PreviousValue *int64                                                     `json:"previousValue,omitempty"`
	Source        GetAllChecks200ApplicationJSONChecksOutputMetricsCLSSource `json:"source"`
	Value         *int64                                                     `json:"value"`
}

func (o *GetAllChecks200ApplicationJSONChecksOutputMetricsCLS) GetPreviousValue() *int64 {
	if o == nil {
		return nil
	}
	return o.PreviousValue
}

func (o *GetAllChecks200ApplicationJSONChecksOutputMetricsCLS) GetSource() GetAllChecks200ApplicationJSONChecksOutputMetricsCLSSource {
	if o == nil {
		return GetAllChecks200ApplicationJSONChecksOutputMetricsCLSSource("")
	}
	return o.Source
}

func (o *GetAllChecks200ApplicationJSONChecksOutputMetricsCLS) GetValue() *int64 {
	if o == nil {
		return nil
	}
	return o.Value
}

type GetAllChecks200ApplicationJSONChecksOutputMetricsFCPSource string

const (
	GetAllChecks200ApplicationJSONChecksOutputMetricsFCPSourceWebVitals GetAllChecks200ApplicationJSONChecksOutputMetricsFCPSource = "web-vitals"
)

func (e GetAllChecks200ApplicationJSONChecksOutputMetricsFCPSource) ToPointer() *GetAllChecks200ApplicationJSONChecksOutputMetricsFCPSource {
	return &e
}

func (e *GetAllChecks200ApplicationJSONChecksOutputMetricsFCPSource) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "web-vitals":
		*e = GetAllChecks200ApplicationJSONChecksOutputMetricsFCPSource(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetAllChecks200ApplicationJSONChecksOutputMetricsFCPSource: %v", v)
	}
}

type GetAllChecks200ApplicationJSONChecksOutputMetricsFCP struct {
	PreviousValue *int64                                                     `json:"previousValue,omitempty"`
	Source        GetAllChecks200ApplicationJSONChecksOutputMetricsFCPSource `json:"source"`
	Value         *int64                                                     `json:"value"`
}

func (o *GetAllChecks200ApplicationJSONChecksOutputMetricsFCP) GetPreviousValue() *int64 {
	if o == nil {
		return nil
	}
	return o.PreviousValue
}

func (o *GetAllChecks200ApplicationJSONChecksOutputMetricsFCP) GetSource() GetAllChecks200ApplicationJSONChecksOutputMetricsFCPSource {
	if o == nil {
		return GetAllChecks200ApplicationJSONChecksOutputMetricsFCPSource("")
	}
	return o.Source
}

func (o *GetAllChecks200ApplicationJSONChecksOutputMetricsFCP) GetValue() *int64 {
	if o == nil {
		return nil
	}
	return o.Value
}

type GetAllChecks200ApplicationJSONChecksOutputMetricsLCPSource string

const (
	GetAllChecks200ApplicationJSONChecksOutputMetricsLCPSourceWebVitals GetAllChecks200ApplicationJSONChecksOutputMetricsLCPSource = "web-vitals"
)

func (e GetAllChecks200ApplicationJSONChecksOutputMetricsLCPSource) ToPointer() *GetAllChecks200ApplicationJSONChecksOutputMetricsLCPSource {
	return &e
}

func (e *GetAllChecks200ApplicationJSONChecksOutputMetricsLCPSource) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "web-vitals":
		*e = GetAllChecks200ApplicationJSONChecksOutputMetricsLCPSource(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetAllChecks200ApplicationJSONChecksOutputMetricsLCPSource: %v", v)
	}
}

type GetAllChecks200ApplicationJSONChecksOutputMetricsLCP struct {
	PreviousValue *int64                                                     `json:"previousValue,omitempty"`
	Source        GetAllChecks200ApplicationJSONChecksOutputMetricsLCPSource `json:"source"`
	Value         *int64                                                     `json:"value"`
}

func (o *GetAllChecks200ApplicationJSONChecksOutputMetricsLCP) GetPreviousValue() *int64 {
	if o == nil {
		return nil
	}
	return o.PreviousValue
}

func (o *GetAllChecks200ApplicationJSONChecksOutputMetricsLCP) GetSource() GetAllChecks200ApplicationJSONChecksOutputMetricsLCPSource {
	if o == nil {
		return GetAllChecks200ApplicationJSONChecksOutputMetricsLCPSource("")
	}
	return o.Source
}

func (o *GetAllChecks200ApplicationJSONChecksOutputMetricsLCP) GetValue() *int64 {
	if o == nil {
		return nil
	}
	return o.Value
}

type GetAllChecks200ApplicationJSONChecksOutputMetricsTBTSource string

const (
	GetAllChecks200ApplicationJSONChecksOutputMetricsTBTSourceWebVitals GetAllChecks200ApplicationJSONChecksOutputMetricsTBTSource = "web-vitals"
)

func (e GetAllChecks200ApplicationJSONChecksOutputMetricsTBTSource) ToPointer() *GetAllChecks200ApplicationJSONChecksOutputMetricsTBTSource {
	return &e
}

func (e *GetAllChecks200ApplicationJSONChecksOutputMetricsTBTSource) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "web-vitals":
		*e = GetAllChecks200ApplicationJSONChecksOutputMetricsTBTSource(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetAllChecks200ApplicationJSONChecksOutputMetricsTBTSource: %v", v)
	}
}

type GetAllChecks200ApplicationJSONChecksOutputMetricsTBT struct {
	PreviousValue *int64                                                     `json:"previousValue,omitempty"`
	Source        GetAllChecks200ApplicationJSONChecksOutputMetricsTBTSource `json:"source"`
	Value         *int64                                                     `json:"value"`
}

func (o *GetAllChecks200ApplicationJSONChecksOutputMetricsTBT) GetPreviousValue() *int64 {
	if o == nil {
		return nil
	}
	return o.PreviousValue
}

func (o *GetAllChecks200ApplicationJSONChecksOutputMetricsTBT) GetSource() GetAllChecks200ApplicationJSONChecksOutputMetricsTBTSource {
	if o == nil {
		return GetAllChecks200ApplicationJSONChecksOutputMetricsTBTSource("")
	}
	return o.Source
}

func (o *GetAllChecks200ApplicationJSONChecksOutputMetricsTBT) GetValue() *int64 {
	if o == nil {
		return nil
	}
	return o.Value
}

type GetAllChecks200ApplicationJSONChecksOutputMetricsVirtualExperienceScoreSource string

const (
	GetAllChecks200ApplicationJSONChecksOutputMetricsVirtualExperienceScoreSourceWebVitals GetAllChecks200ApplicationJSONChecksOutputMetricsVirtualExperienceScoreSource = "web-vitals"
)

func (e GetAllChecks200ApplicationJSONChecksOutputMetricsVirtualExperienceScoreSource) ToPointer() *GetAllChecks200ApplicationJSONChecksOutputMetricsVirtualExperienceScoreSource {
	return &e
}

func (e *GetAllChecks200ApplicationJSONChecksOutputMetricsVirtualExperienceScoreSource) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "web-vitals":
		*e = GetAllChecks200ApplicationJSONChecksOutputMetricsVirtualExperienceScoreSource(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetAllChecks200ApplicationJSONChecksOutputMetricsVirtualExperienceScoreSource: %v", v)
	}
}

type GetAllChecks200ApplicationJSONChecksOutputMetricsVirtualExperienceScore struct {
	PreviousValue *int64                                                                        `json:"previousValue,omitempty"`
	Source        GetAllChecks200ApplicationJSONChecksOutputMetricsVirtualExperienceScoreSource `json:"source"`
	Value         *int64                                                                        `json:"value"`
}

func (o *GetAllChecks200ApplicationJSONChecksOutputMetricsVirtualExperienceScore) GetPreviousValue() *int64 {
	if o == nil {
		return nil
	}
	return o.PreviousValue
}

func (o *GetAllChecks200ApplicationJSONChecksOutputMetricsVirtualExperienceScore) GetSource() GetAllChecks200ApplicationJSONChecksOutputMetricsVirtualExperienceScoreSource {
	if o == nil {
		return GetAllChecks200ApplicationJSONChecksOutputMetricsVirtualExperienceScoreSource("")
	}
	return o.Source
}

func (o *GetAllChecks200ApplicationJSONChecksOutputMetricsVirtualExperienceScore) GetValue() *int64 {
	if o == nil {
		return nil
	}
	return o.Value
}

type GetAllChecks200ApplicationJSONChecksOutputMetrics struct {
	Cls                    GetAllChecks200ApplicationJSONChecksOutputMetricsCLS                     `json:"CLS"`
	Fcp                    GetAllChecks200ApplicationJSONChecksOutputMetricsFCP                     `json:"FCP"`
	Lcp                    GetAllChecks200ApplicationJSONChecksOutputMetricsLCP                     `json:"LCP"`
	Tbt                    GetAllChecks200ApplicationJSONChecksOutputMetricsTBT                     `json:"TBT"`
	VirtualExperienceScore *GetAllChecks200ApplicationJSONChecksOutputMetricsVirtualExperienceScore `json:"virtualExperienceScore,omitempty"`
}

func (o *GetAllChecks200ApplicationJSONChecksOutputMetrics) GetCls() GetAllChecks200ApplicationJSONChecksOutputMetricsCLS {
	if o == nil {
		return GetAllChecks200ApplicationJSONChecksOutputMetricsCLS{}
	}
	return o.Cls
}

func (o *GetAllChecks200ApplicationJSONChecksOutputMetrics) GetFcp() GetAllChecks200ApplicationJSONChecksOutputMetricsFCP {
	if o == nil {
		return GetAllChecks200ApplicationJSONChecksOutputMetricsFCP{}
	}
	return o.Fcp
}

func (o *GetAllChecks200ApplicationJSONChecksOutputMetrics) GetLcp() GetAllChecks200ApplicationJSONChecksOutputMetricsLCP {
	if o == nil {
		return GetAllChecks200ApplicationJSONChecksOutputMetricsLCP{}
	}
	return o.Lcp
}

func (o *GetAllChecks200ApplicationJSONChecksOutputMetrics) GetTbt() GetAllChecks200ApplicationJSONChecksOutputMetricsTBT {
	if o == nil {
		return GetAllChecks200ApplicationJSONChecksOutputMetricsTBT{}
	}
	return o.Tbt
}

func (o *GetAllChecks200ApplicationJSONChecksOutputMetrics) GetVirtualExperienceScore() *GetAllChecks200ApplicationJSONChecksOutputMetricsVirtualExperienceScore {
	if o == nil {
		return nil
	}
	return o.VirtualExperienceScore
}

type GetAllChecks200ApplicationJSONChecksOutput struct {
	Metrics *GetAllChecks200ApplicationJSONChecksOutputMetrics `json:"metrics,omitempty"`
}

func (o *GetAllChecks200ApplicationJSONChecksOutput) GetMetrics() *GetAllChecks200ApplicationJSONChecksOutputMetrics {
	if o == nil {
		return nil
	}
	return o.Metrics
}

type GetAllChecks200ApplicationJSONChecksStatus string

const (
	GetAllChecks200ApplicationJSONChecksStatusRegistered GetAllChecks200ApplicationJSONChecksStatus = "registered"
	GetAllChecks200ApplicationJSONChecksStatusRunning    GetAllChecks200ApplicationJSONChecksStatus = "running"
	GetAllChecks200ApplicationJSONChecksStatusCompleted  GetAllChecks200ApplicationJSONChecksStatus = "completed"
)

func (e GetAllChecks200ApplicationJSONChecksStatus) ToPointer() *GetAllChecks200ApplicationJSONChecksStatus {
	return &e
}

func (e *GetAllChecks200ApplicationJSONChecksStatus) UnmarshalJSON(data []byte) error {
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
		*e = GetAllChecks200ApplicationJSONChecksStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetAllChecks200ApplicationJSONChecksStatus: %v", v)
	}
}

type GetAllChecks200ApplicationJSONChecks struct {
	CompletedAt   *int64                                          `json:"completedAt,omitempty"`
	Conclusion    *GetAllChecks200ApplicationJSONChecksConclusion `json:"conclusion,omitempty"`
	CreatedAt     int64                                           `json:"createdAt"`
	DetailsURL    *string                                         `json:"detailsUrl,omitempty"`
	ID            string                                          `json:"id"`
	IntegrationID string                                          `json:"integrationId"`
	Name          string                                          `json:"name"`
	Output        *GetAllChecks200ApplicationJSONChecksOutput     `json:"output,omitempty"`
	Path          *string                                         `json:"path,omitempty"`
	Rerequestable bool                                            `json:"rerequestable"`
	StartedAt     *int64                                          `json:"startedAt,omitempty"`
	Status        GetAllChecks200ApplicationJSONChecksStatus      `json:"status"`
	UpdatedAt     int64                                           `json:"updatedAt"`
}

func (o *GetAllChecks200ApplicationJSONChecks) GetCompletedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CompletedAt
}

func (o *GetAllChecks200ApplicationJSONChecks) GetConclusion() *GetAllChecks200ApplicationJSONChecksConclusion {
	if o == nil {
		return nil
	}
	return o.Conclusion
}

func (o *GetAllChecks200ApplicationJSONChecks) GetCreatedAt() int64 {
	if o == nil {
		return 0
	}
	return o.CreatedAt
}

func (o *GetAllChecks200ApplicationJSONChecks) GetDetailsURL() *string {
	if o == nil {
		return nil
	}
	return o.DetailsURL
}

func (o *GetAllChecks200ApplicationJSONChecks) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetAllChecks200ApplicationJSONChecks) GetIntegrationID() string {
	if o == nil {
		return ""
	}
	return o.IntegrationID
}

func (o *GetAllChecks200ApplicationJSONChecks) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *GetAllChecks200ApplicationJSONChecks) GetOutput() *GetAllChecks200ApplicationJSONChecksOutput {
	if o == nil {
		return nil
	}
	return o.Output
}

func (o *GetAllChecks200ApplicationJSONChecks) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

func (o *GetAllChecks200ApplicationJSONChecks) GetRerequestable() bool {
	if o == nil {
		return false
	}
	return o.Rerequestable
}

func (o *GetAllChecks200ApplicationJSONChecks) GetStartedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.StartedAt
}

func (o *GetAllChecks200ApplicationJSONChecks) GetStatus() GetAllChecks200ApplicationJSONChecksStatus {
	if o == nil {
		return GetAllChecks200ApplicationJSONChecksStatus("")
	}
	return o.Status
}

func (o *GetAllChecks200ApplicationJSONChecks) GetUpdatedAt() int64 {
	if o == nil {
		return 0
	}
	return o.UpdatedAt
}

type GetAllChecks200ApplicationJSON struct {
	Checks []GetAllChecks200ApplicationJSONChecks `json:"checks"`
}

func (o *GetAllChecks200ApplicationJSON) GetChecks() []GetAllChecks200ApplicationJSONChecks {
	if o == nil {
		return []GetAllChecks200ApplicationJSONChecks{}
	}
	return o.Checks
}

type GetAllChecksResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse                          *http.Response
	GetAllChecks200ApplicationJSONObject *GetAllChecks200ApplicationJSON
}

func (o *GetAllChecksResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAllChecksResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAllChecksResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetAllChecksResponse) GetGetAllChecks200ApplicationJSONObject() *GetAllChecks200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetAllChecks200ApplicationJSONObject
}
