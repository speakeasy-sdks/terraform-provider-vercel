// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Conclusion - The result of the check being run
type Conclusion string

const (
	ConclusionCanceled  Conclusion = "canceled"
	ConclusionFailed    Conclusion = "failed"
	ConclusionNeutral   Conclusion = "neutral"
	ConclusionSucceeded Conclusion = "succeeded"
	ConclusionSkipped   Conclusion = "skipped"
)

func (e Conclusion) ToPointer() *Conclusion {
	return &e
}

func (e *Conclusion) UnmarshalJSON(data []byte) error {
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
		*e = Conclusion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Conclusion: %v", v)
	}
}

type UpdateCheckSource string

const (
	UpdateCheckSourceWebVitals UpdateCheckSource = "web-vitals"
)

func (e UpdateCheckSource) ToPointer() *UpdateCheckSource {
	return &e
}

func (e *UpdateCheckSource) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "web-vitals":
		*e = UpdateCheckSource(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateCheckSource: %v", v)
	}
}

type Cls struct {
	// Previous Cumulative Layout Shift value to display a delta
	PreviousValue *int64            `json:"previousValue,omitempty"`
	Source        UpdateCheckSource `json:"source"`
	// Cumulative Layout Shift value
	Value *int64 `json:"value"`
}

func (o *Cls) GetPreviousValue() *int64 {
	if o == nil {
		return nil
	}
	return o.PreviousValue
}

func (o *Cls) GetSource() UpdateCheckSource {
	if o == nil {
		return UpdateCheckSource("")
	}
	return o.Source
}

func (o *Cls) GetValue() *int64 {
	if o == nil {
		return nil
	}
	return o.Value
}

type UpdateCheckChecksSource string

const (
	UpdateCheckChecksSourceWebVitals UpdateCheckChecksSource = "web-vitals"
)

func (e UpdateCheckChecksSource) ToPointer() *UpdateCheckChecksSource {
	return &e
}

func (e *UpdateCheckChecksSource) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "web-vitals":
		*e = UpdateCheckChecksSource(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateCheckChecksSource: %v", v)
	}
}

type Fcp struct {
	// Previous First Contentful Paint value to display a delta
	PreviousValue *int64                  `json:"previousValue,omitempty"`
	Source        UpdateCheckChecksSource `json:"source"`
	// First Contentful Paint value
	Value *int64 `json:"value"`
}

func (o *Fcp) GetPreviousValue() *int64 {
	if o == nil {
		return nil
	}
	return o.PreviousValue
}

func (o *Fcp) GetSource() UpdateCheckChecksSource {
	if o == nil {
		return UpdateCheckChecksSource("")
	}
	return o.Source
}

func (o *Fcp) GetValue() *int64 {
	if o == nil {
		return nil
	}
	return o.Value
}

type UpdateCheckChecksRequestSource string

const (
	UpdateCheckChecksRequestSourceWebVitals UpdateCheckChecksRequestSource = "web-vitals"
)

func (e UpdateCheckChecksRequestSource) ToPointer() *UpdateCheckChecksRequestSource {
	return &e
}

func (e *UpdateCheckChecksRequestSource) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "web-vitals":
		*e = UpdateCheckChecksRequestSource(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateCheckChecksRequestSource: %v", v)
	}
}

type Lcp struct {
	// Previous Largest Contentful Paint value to display a delta
	PreviousValue *int64                         `json:"previousValue,omitempty"`
	Source        UpdateCheckChecksRequestSource `json:"source"`
	// Largest Contentful Paint value
	Value *int64 `json:"value"`
}

func (o *Lcp) GetPreviousValue() *int64 {
	if o == nil {
		return nil
	}
	return o.PreviousValue
}

func (o *Lcp) GetSource() UpdateCheckChecksRequestSource {
	if o == nil {
		return UpdateCheckChecksRequestSource("")
	}
	return o.Source
}

func (o *Lcp) GetValue() *int64 {
	if o == nil {
		return nil
	}
	return o.Value
}

type UpdateCheckChecksRequestRequestBodySource string

const (
	UpdateCheckChecksRequestRequestBodySourceWebVitals UpdateCheckChecksRequestRequestBodySource = "web-vitals"
)

func (e UpdateCheckChecksRequestRequestBodySource) ToPointer() *UpdateCheckChecksRequestRequestBodySource {
	return &e
}

func (e *UpdateCheckChecksRequestRequestBodySource) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "web-vitals":
		*e = UpdateCheckChecksRequestRequestBodySource(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateCheckChecksRequestRequestBodySource: %v", v)
	}
}

type Tbt struct {
	// Previous Total Blocking Time value to display a delta
	PreviousValue *int64                                    `json:"previousValue,omitempty"`
	Source        UpdateCheckChecksRequestRequestBodySource `json:"source"`
	// Total Blocking Time value
	Value *int64 `json:"value"`
}

func (o *Tbt) GetPreviousValue() *int64 {
	if o == nil {
		return nil
	}
	return o.PreviousValue
}

func (o *Tbt) GetSource() UpdateCheckChecksRequestRequestBodySource {
	if o == nil {
		return UpdateCheckChecksRequestRequestBodySource("")
	}
	return o.Source
}

func (o *Tbt) GetValue() *int64 {
	if o == nil {
		return nil
	}
	return o.Value
}

type UpdateCheckChecksRequestRequestBodyOutputSource string

const (
	UpdateCheckChecksRequestRequestBodyOutputSourceWebVitals UpdateCheckChecksRequestRequestBodyOutputSource = "web-vitals"
)

func (e UpdateCheckChecksRequestRequestBodyOutputSource) ToPointer() *UpdateCheckChecksRequestRequestBodyOutputSource {
	return &e
}

func (e *UpdateCheckChecksRequestRequestBodyOutputSource) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "web-vitals":
		*e = UpdateCheckChecksRequestRequestBodyOutputSource(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateCheckChecksRequestRequestBodyOutputSource: %v", v)
	}
}

type VirtualExperienceScore struct {
	// A previous Virtual Experience Score value to display a delta, between 0 and 100
	PreviousValue *int64                                          `json:"previousValue,omitempty"`
	Source        UpdateCheckChecksRequestRequestBodyOutputSource `json:"source"`
	// The calculated Virtual Experience Score value, between 0 and 100
	Value *int64 `json:"value"`
}

func (o *VirtualExperienceScore) GetPreviousValue() *int64 {
	if o == nil {
		return nil
	}
	return o.PreviousValue
}

func (o *VirtualExperienceScore) GetSource() UpdateCheckChecksRequestRequestBodyOutputSource {
	if o == nil {
		return UpdateCheckChecksRequestRequestBodyOutputSource("")
	}
	return o.Source
}

func (o *VirtualExperienceScore) GetValue() *int64 {
	if o == nil {
		return nil
	}
	return o.Value
}

// Metrics about the page
type Metrics struct {
	Cls                    Cls                     `json:"CLS"`
	Fcp                    Fcp                     `json:"FCP"`
	Lcp                    Lcp                     `json:"LCP"`
	Tbt                    Tbt                     `json:"TBT"`
	VirtualExperienceScore *VirtualExperienceScore `json:"virtualExperienceScore,omitempty"`
}

func (o *Metrics) GetCls() Cls {
	if o == nil {
		return Cls{}
	}
	return o.Cls
}

func (o *Metrics) GetFcp() Fcp {
	if o == nil {
		return Fcp{}
	}
	return o.Fcp
}

func (o *Metrics) GetLcp() Lcp {
	if o == nil {
		return Lcp{}
	}
	return o.Lcp
}

func (o *Metrics) GetTbt() Tbt {
	if o == nil {
		return Tbt{}
	}
	return o.Tbt
}

func (o *Metrics) GetVirtualExperienceScore() *VirtualExperienceScore {
	if o == nil {
		return nil
	}
	return o.VirtualExperienceScore
}

// Output - The results of the check Run
type Output struct {
	// Metrics about the page
	Metrics *Metrics `json:"metrics,omitempty"`
}

func (o *Output) GetMetrics() *Metrics {
	if o == nil {
		return nil
	}
	return o.Metrics
}

// Status - The current status of the check
type Status string

const (
	StatusRunning   Status = "running"
	StatusCompleted Status = "completed"
)

func (e Status) ToPointer() *Status {
	return &e
}

func (e *Status) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "running":
		fallthrough
	case "completed":
		*e = Status(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Status: %v", v)
	}
}

type UpdateCheckRequestBody struct {
	// The result of the check being run
	Conclusion *Conclusion `json:"conclusion,omitempty"`
	// A URL a user may visit to see more information about the check
	DetailsURL *string `json:"detailsUrl,omitempty"`
	// An identifier that can be used as an external reference
	ExternalID *string `json:"externalId,omitempty"`
	// The name of the check being created
	Name *string `json:"name,omitempty"`
	// The results of the check Run
	Output *Output `json:"output,omitempty"`
	// Path of the page that is being checked
	Path *string `json:"path,omitempty"`
	// The current status of the check
	Status *Status `json:"status,omitempty"`
}

func (o *UpdateCheckRequestBody) GetConclusion() *Conclusion {
	if o == nil {
		return nil
	}
	return o.Conclusion
}

func (o *UpdateCheckRequestBody) GetDetailsURL() *string {
	if o == nil {
		return nil
	}
	return o.DetailsURL
}

func (o *UpdateCheckRequestBody) GetExternalID() *string {
	if o == nil {
		return nil
	}
	return o.ExternalID
}

func (o *UpdateCheckRequestBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *UpdateCheckRequestBody) GetOutput() *Output {
	if o == nil {
		return nil
	}
	return o.Output
}

func (o *UpdateCheckRequestBody) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

func (o *UpdateCheckRequestBody) GetStatus() *Status {
	if o == nil {
		return nil
	}
	return o.Status
}

type UpdateCheckRequest struct {
	RequestBody *UpdateCheckRequestBody `request:"mediaType=application/json"`
	// The check being updated
	CheckID string `pathParam:"style=simple,explode=false,name=checkId"`
	// The deployment to update the check for.
	DeploymentID string `pathParam:"style=simple,explode=false,name=deploymentId"`
	// The Team identifier or slug to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
}

func (o *UpdateCheckRequest) GetRequestBody() *UpdateCheckRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *UpdateCheckRequest) GetCheckID() string {
	if o == nil {
		return ""
	}
	return o.CheckID
}

func (o *UpdateCheckRequest) GetDeploymentID() string {
	if o == nil {
		return ""
	}
	return o.DeploymentID
}

func (o *UpdateCheckRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

type UpdateCheckConclusion string

const (
	UpdateCheckConclusionCanceled  UpdateCheckConclusion = "canceled"
	UpdateCheckConclusionFailed    UpdateCheckConclusion = "failed"
	UpdateCheckConclusionNeutral   UpdateCheckConclusion = "neutral"
	UpdateCheckConclusionSucceeded UpdateCheckConclusion = "succeeded"
	UpdateCheckConclusionSkipped   UpdateCheckConclusion = "skipped"
	UpdateCheckConclusionStale     UpdateCheckConclusion = "stale"
)

func (e UpdateCheckConclusion) ToPointer() *UpdateCheckConclusion {
	return &e
}

func (e *UpdateCheckConclusion) UnmarshalJSON(data []byte) error {
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
		*e = UpdateCheckConclusion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateCheckConclusion: %v", v)
	}
}

type UpdateCheckChecksResponseSource string

const (
	UpdateCheckChecksResponseSourceWebVitals UpdateCheckChecksResponseSource = "web-vitals"
)

func (e UpdateCheckChecksResponseSource) ToPointer() *UpdateCheckChecksResponseSource {
	return &e
}

func (e *UpdateCheckChecksResponseSource) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "web-vitals":
		*e = UpdateCheckChecksResponseSource(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateCheckChecksResponseSource: %v", v)
	}
}

type UpdateCheckCLS struct {
	PreviousValue *int64                          `json:"previousValue,omitempty"`
	Source        UpdateCheckChecksResponseSource `json:"source"`
	Value         *int64                          `json:"value"`
}

func (o *UpdateCheckCLS) GetPreviousValue() *int64 {
	if o == nil {
		return nil
	}
	return o.PreviousValue
}

func (o *UpdateCheckCLS) GetSource() UpdateCheckChecksResponseSource {
	if o == nil {
		return UpdateCheckChecksResponseSource("")
	}
	return o.Source
}

func (o *UpdateCheckCLS) GetValue() *int64 {
	if o == nil {
		return nil
	}
	return o.Value
}

type UpdateCheckChecksResponse200Source string

const (
	UpdateCheckChecksResponse200SourceWebVitals UpdateCheckChecksResponse200Source = "web-vitals"
)

func (e UpdateCheckChecksResponse200Source) ToPointer() *UpdateCheckChecksResponse200Source {
	return &e
}

func (e *UpdateCheckChecksResponse200Source) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "web-vitals":
		*e = UpdateCheckChecksResponse200Source(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateCheckChecksResponse200Source: %v", v)
	}
}

type UpdateCheckFCP struct {
	PreviousValue *int64                             `json:"previousValue,omitempty"`
	Source        UpdateCheckChecksResponse200Source `json:"source"`
	Value         *int64                             `json:"value"`
}

func (o *UpdateCheckFCP) GetPreviousValue() *int64 {
	if o == nil {
		return nil
	}
	return o.PreviousValue
}

func (o *UpdateCheckFCP) GetSource() UpdateCheckChecksResponse200Source {
	if o == nil {
		return UpdateCheckChecksResponse200Source("")
	}
	return o.Source
}

func (o *UpdateCheckFCP) GetValue() *int64 {
	if o == nil {
		return nil
	}
	return o.Value
}

type UpdateCheckChecksResponse200ApplicationJSONSource string

const (
	UpdateCheckChecksResponse200ApplicationJSONSourceWebVitals UpdateCheckChecksResponse200ApplicationJSONSource = "web-vitals"
)

func (e UpdateCheckChecksResponse200ApplicationJSONSource) ToPointer() *UpdateCheckChecksResponse200ApplicationJSONSource {
	return &e
}

func (e *UpdateCheckChecksResponse200ApplicationJSONSource) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "web-vitals":
		*e = UpdateCheckChecksResponse200ApplicationJSONSource(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateCheckChecksResponse200ApplicationJSONSource: %v", v)
	}
}

type UpdateCheckLCP struct {
	PreviousValue *int64                                            `json:"previousValue,omitempty"`
	Source        UpdateCheckChecksResponse200ApplicationJSONSource `json:"source"`
	Value         *int64                                            `json:"value"`
}

func (o *UpdateCheckLCP) GetPreviousValue() *int64 {
	if o == nil {
		return nil
	}
	return o.PreviousValue
}

func (o *UpdateCheckLCP) GetSource() UpdateCheckChecksResponse200ApplicationJSONSource {
	if o == nil {
		return UpdateCheckChecksResponse200ApplicationJSONSource("")
	}
	return o.Source
}

func (o *UpdateCheckLCP) GetValue() *int64 {
	if o == nil {
		return nil
	}
	return o.Value
}

type UpdateCheckChecksResponse200ApplicationJSONResponseBodySource string

const (
	UpdateCheckChecksResponse200ApplicationJSONResponseBodySourceWebVitals UpdateCheckChecksResponse200ApplicationJSONResponseBodySource = "web-vitals"
)

func (e UpdateCheckChecksResponse200ApplicationJSONResponseBodySource) ToPointer() *UpdateCheckChecksResponse200ApplicationJSONResponseBodySource {
	return &e
}

func (e *UpdateCheckChecksResponse200ApplicationJSONResponseBodySource) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "web-vitals":
		*e = UpdateCheckChecksResponse200ApplicationJSONResponseBodySource(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateCheckChecksResponse200ApplicationJSONResponseBodySource: %v", v)
	}
}

type UpdateCheckTBT struct {
	PreviousValue *int64                                                        `json:"previousValue,omitempty"`
	Source        UpdateCheckChecksResponse200ApplicationJSONResponseBodySource `json:"source"`
	Value         *int64                                                        `json:"value"`
}

func (o *UpdateCheckTBT) GetPreviousValue() *int64 {
	if o == nil {
		return nil
	}
	return o.PreviousValue
}

func (o *UpdateCheckTBT) GetSource() UpdateCheckChecksResponse200ApplicationJSONResponseBodySource {
	if o == nil {
		return UpdateCheckChecksResponse200ApplicationJSONResponseBodySource("")
	}
	return o.Source
}

func (o *UpdateCheckTBT) GetValue() *int64 {
	if o == nil {
		return nil
	}
	return o.Value
}

type UpdateCheckChecksResponse200ApplicationJSONResponseBodyOutputSource string

const (
	UpdateCheckChecksResponse200ApplicationJSONResponseBodyOutputSourceWebVitals UpdateCheckChecksResponse200ApplicationJSONResponseBodyOutputSource = "web-vitals"
)

func (e UpdateCheckChecksResponse200ApplicationJSONResponseBodyOutputSource) ToPointer() *UpdateCheckChecksResponse200ApplicationJSONResponseBodyOutputSource {
	return &e
}

func (e *UpdateCheckChecksResponse200ApplicationJSONResponseBodyOutputSource) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "web-vitals":
		*e = UpdateCheckChecksResponse200ApplicationJSONResponseBodyOutputSource(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateCheckChecksResponse200ApplicationJSONResponseBodyOutputSource: %v", v)
	}
}

type UpdateCheckVirtualExperienceScore struct {
	PreviousValue *int64                                                              `json:"previousValue,omitempty"`
	Source        UpdateCheckChecksResponse200ApplicationJSONResponseBodyOutputSource `json:"source"`
	Value         *int64                                                              `json:"value"`
}

func (o *UpdateCheckVirtualExperienceScore) GetPreviousValue() *int64 {
	if o == nil {
		return nil
	}
	return o.PreviousValue
}

func (o *UpdateCheckVirtualExperienceScore) GetSource() UpdateCheckChecksResponse200ApplicationJSONResponseBodyOutputSource {
	if o == nil {
		return UpdateCheckChecksResponse200ApplicationJSONResponseBodyOutputSource("")
	}
	return o.Source
}

func (o *UpdateCheckVirtualExperienceScore) GetValue() *int64 {
	if o == nil {
		return nil
	}
	return o.Value
}

type UpdateCheckMetrics struct {
	Cls                    UpdateCheckCLS                     `json:"CLS"`
	Fcp                    UpdateCheckFCP                     `json:"FCP"`
	Lcp                    UpdateCheckLCP                     `json:"LCP"`
	Tbt                    UpdateCheckTBT                     `json:"TBT"`
	VirtualExperienceScore *UpdateCheckVirtualExperienceScore `json:"virtualExperienceScore,omitempty"`
}

func (o *UpdateCheckMetrics) GetCls() UpdateCheckCLS {
	if o == nil {
		return UpdateCheckCLS{}
	}
	return o.Cls
}

func (o *UpdateCheckMetrics) GetFcp() UpdateCheckFCP {
	if o == nil {
		return UpdateCheckFCP{}
	}
	return o.Fcp
}

func (o *UpdateCheckMetrics) GetLcp() UpdateCheckLCP {
	if o == nil {
		return UpdateCheckLCP{}
	}
	return o.Lcp
}

func (o *UpdateCheckMetrics) GetTbt() UpdateCheckTBT {
	if o == nil {
		return UpdateCheckTBT{}
	}
	return o.Tbt
}

func (o *UpdateCheckMetrics) GetVirtualExperienceScore() *UpdateCheckVirtualExperienceScore {
	if o == nil {
		return nil
	}
	return o.VirtualExperienceScore
}

type UpdateCheckOutput struct {
	Metrics *UpdateCheckMetrics `json:"metrics,omitempty"`
}

func (o *UpdateCheckOutput) GetMetrics() *UpdateCheckMetrics {
	if o == nil {
		return nil
	}
	return o.Metrics
}

type UpdateCheckStatus string

const (
	UpdateCheckStatusRegistered UpdateCheckStatus = "registered"
	UpdateCheckStatusRunning    UpdateCheckStatus = "running"
	UpdateCheckStatusCompleted  UpdateCheckStatus = "completed"
)

func (e UpdateCheckStatus) ToPointer() *UpdateCheckStatus {
	return &e
}

func (e *UpdateCheckStatus) UnmarshalJSON(data []byte) error {
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
		*e = UpdateCheckStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateCheckStatus: %v", v)
	}
}

type UpdateCheckResponseBody struct {
	Blocking      bool                   `json:"blocking"`
	CompletedAt   *int64                 `json:"completedAt,omitempty"`
	Conclusion    *UpdateCheckConclusion `json:"conclusion,omitempty"`
	CreatedAt     int64                  `json:"createdAt"`
	DeploymentID  string                 `json:"deploymentId"`
	DetailsURL    *string                `json:"detailsUrl,omitempty"`
	ExternalID    *string                `json:"externalId,omitempty"`
	ID            string                 `json:"id"`
	IntegrationID string                 `json:"integrationId"`
	Name          string                 `json:"name"`
	Output        *UpdateCheckOutput     `json:"output,omitempty"`
	Path          *string                `json:"path,omitempty"`
	Rerequestable *bool                  `json:"rerequestable,omitempty"`
	StartedAt     *int64                 `json:"startedAt,omitempty"`
	Status        UpdateCheckStatus      `json:"status"`
	UpdatedAt     int64                  `json:"updatedAt"`
}

func (o *UpdateCheckResponseBody) GetBlocking() bool {
	if o == nil {
		return false
	}
	return o.Blocking
}

func (o *UpdateCheckResponseBody) GetCompletedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CompletedAt
}

func (o *UpdateCheckResponseBody) GetConclusion() *UpdateCheckConclusion {
	if o == nil {
		return nil
	}
	return o.Conclusion
}

func (o *UpdateCheckResponseBody) GetCreatedAt() int64 {
	if o == nil {
		return 0
	}
	return o.CreatedAt
}

func (o *UpdateCheckResponseBody) GetDeploymentID() string {
	if o == nil {
		return ""
	}
	return o.DeploymentID
}

func (o *UpdateCheckResponseBody) GetDetailsURL() *string {
	if o == nil {
		return nil
	}
	return o.DetailsURL
}

func (o *UpdateCheckResponseBody) GetExternalID() *string {
	if o == nil {
		return nil
	}
	return o.ExternalID
}

func (o *UpdateCheckResponseBody) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateCheckResponseBody) GetIntegrationID() string {
	if o == nil {
		return ""
	}
	return o.IntegrationID
}

func (o *UpdateCheckResponseBody) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *UpdateCheckResponseBody) GetOutput() *UpdateCheckOutput {
	if o == nil {
		return nil
	}
	return o.Output
}

func (o *UpdateCheckResponseBody) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

func (o *UpdateCheckResponseBody) GetRerequestable() *bool {
	if o == nil {
		return nil
	}
	return o.Rerequestable
}

func (o *UpdateCheckResponseBody) GetStartedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.StartedAt
}

func (o *UpdateCheckResponseBody) GetStatus() UpdateCheckStatus {
	if o == nil {
		return UpdateCheckStatus("")
	}
	return o.Status
}

func (o *UpdateCheckResponseBody) GetUpdatedAt() int64 {
	if o == nil {
		return 0
	}
	return o.UpdatedAt
}

type UpdateCheckResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	Object      *UpdateCheckResponseBody
}

func (o *UpdateCheckResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateCheckResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateCheckResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateCheckResponse) GetObject() *UpdateCheckResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
