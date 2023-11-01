// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GetDeploymentsDeploymentIDBuildsRequest struct {
	// The deployment unique identifier
	DeploymentID string `pathParam:"style=simple,explode=false,name=deploymentId"`
}

func (o *GetDeploymentsDeploymentIDBuildsRequest) GetDeploymentID() string {
	if o == nil {
		return ""
	}
	return o.DeploymentID
}

// GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsConfig - An object that contains the Build's configuration
type GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsConfig struct {
	DistDir           *string `json:"distDir,omitempty"`
	ForceBuildIn      *string `json:"forceBuildIn,omitempty"`
	ReuseWorkPathFrom *string `json:"reuseWorkPathFrom,omitempty"`
	ZeroConfig        *bool   `json:"zeroConfig,omitempty"`
}

func (o *GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsConfig) GetDistDir() *string {
	if o == nil {
		return nil
	}
	return o.DistDir
}

func (o *GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsConfig) GetForceBuildIn() *string {
	if o == nil {
		return nil
	}
	return o.ForceBuildIn
}

func (o *GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsConfig) GetReuseWorkPathFrom() *string {
	if o == nil {
		return nil
	}
	return o.ReuseWorkPathFrom
}

func (o *GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsConfig) GetZeroConfig() *bool {
	if o == nil {
		return nil
	}
	return o.ZeroConfig
}

// GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsOutputEdge - Exists if the output is an edge function.
type GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsOutputEdge struct {
	// The regions where the edge function will be invoked. Only exists if the edge function as a regional edge function, see: https://vercel.com/docs/concepts/edge-network/regions#setting-edge-function-regions
	Regions []string `json:"regions"`
}

func (o *GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsOutputEdge) GetRegions() []string {
	if o == nil {
		return nil
	}
	return o.Regions
}

// GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsOutputLambda - If the output is a Serverless Function, an object containing the name, location and memory size of the function
type GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsOutputLambda struct {
	DeployedTo   []string `json:"deployedTo"`
	FunctionName string   `json:"functionName"`
	Layers       []string `json:"layers,omitempty"`
	MemorySize   *int64   `json:"memorySize,omitempty"`
	Timeout      *int64   `json:"timeout,omitempty"`
}

func (o *GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsOutputLambda) GetDeployedTo() []string {
	if o == nil {
		return []string{}
	}
	return o.DeployedTo
}

func (o *GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsOutputLambda) GetFunctionName() string {
	if o == nil {
		return ""
	}
	return o.FunctionName
}

func (o *GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsOutputLambda) GetLayers() []string {
	if o == nil {
		return nil
	}
	return o.Layers
}

func (o *GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsOutputLambda) GetMemorySize() *int64 {
	if o == nil {
		return nil
	}
	return o.MemorySize
}

func (o *GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsOutputLambda) GetTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.Timeout
}

// GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsOutputType - The type of the output
type GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsOutputType string

const (
	GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsOutputTypeLambda GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsOutputType = "lambda"
	GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsOutputTypeFile   GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsOutputType = "file"
	GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsOutputTypeEdge   GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsOutputType = "edge"
)

func (e GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsOutputType) ToPointer() *GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsOutputType {
	return &e
}

func (e *GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsOutputType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "lambda":
		fallthrough
	case "file":
		fallthrough
	case "edge":
		*e = GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsOutputType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsOutputType: %v", v)
	}
}

// GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsOutput - A list of outputs for the Build that can be either Serverless Functions or static files
type GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsOutput struct {
	// The SHA1 of the file
	Digest string `json:"digest"`
	// Exists if the output is an edge function.
	Edge *GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsOutputEdge `json:"edge,omitempty"`
	// If the output is a Serverless Function, an object containing the name, location and memory size of the function
	Lambda *GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsOutputLambda `json:"lambda,omitempty"`
	// The POSIX file permissions
	Mode int64 `json:"mode"`
	// The absolute path of the file or Serverless Function
	Path string `json:"path"`
	// The size of the file in bytes
	Size *int64 `json:"size,omitempty"`
	// The type of the output
	Type *GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsOutputType `json:"type,omitempty"`
}

func (o *GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsOutput) GetDigest() string {
	if o == nil {
		return ""
	}
	return o.Digest
}

func (o *GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsOutput) GetEdge() *GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsOutputEdge {
	if o == nil {
		return nil
	}
	return o.Edge
}

func (o *GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsOutput) GetLambda() *GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsOutputLambda {
	if o == nil {
		return nil
	}
	return o.Lambda
}

func (o *GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsOutput) GetMode() int64 {
	if o == nil {
		return 0
	}
	return o.Mode
}

func (o *GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsOutput) GetPath() string {
	if o == nil {
		return ""
	}
	return o.Path
}

func (o *GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsOutput) GetSize() *int64 {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsOutput) GetType() *GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsOutputType {
	if o == nil {
		return nil
	}
	return o.Type
}

// GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsReadyState - The state of the deployment depending on the process of deploying, or if it is ready or in an error state
type GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsReadyState string

const (
	GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsReadyStateInitializing GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsReadyState = "INITIALIZING"
	GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsReadyStateBuilding     GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsReadyState = "BUILDING"
	GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsReadyStateUploading    GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsReadyState = "UPLOADING"
	GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsReadyStateDeploying    GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsReadyState = "DEPLOYING"
	GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsReadyStateReady        GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsReadyState = "READY"
	GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsReadyStateArchived     GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsReadyState = "ARCHIVED"
	GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsReadyStateError        GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsReadyState = "ERROR"
	GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsReadyStateQueued       GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsReadyState = "QUEUED"
	GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsReadyStateCanceled     GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsReadyState = "CANCELED"
)

func (e GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsReadyState) ToPointer() *GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsReadyState {
	return &e
}

func (e *GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsReadyState) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "INITIALIZING":
		fallthrough
	case "BUILDING":
		fallthrough
	case "UPLOADING":
		fallthrough
	case "DEPLOYING":
		fallthrough
	case "READY":
		fallthrough
	case "ARCHIVED":
		fallthrough
	case "ERROR":
		fallthrough
	case "QUEUED":
		fallthrough
	case "CANCELED":
		*e = GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsReadyState(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsReadyState: %v", v)
	}
}

// GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuilds - An object representing a Build on Vercel
type GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuilds struct {
	// An object that contains the Build's configuration
	Config     *GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsConfig `json:"config,omitempty"`
	CopiedFrom *string                                                         `json:"copiedFrom,omitempty"`
	// The time at which the Build was created
	CreatedAt *int64 `json:"createdAt,omitempty"`
	// The region where the Build was first created
	CreatedIn *string `json:"createdIn,omitempty"`
	// The time at which the Build was deployed
	DeployedAt *int64 `json:"deployedAt,omitempty"`
	// The unique identifier of the deployment
	DeploymentID string `json:"deploymentId"`
	// The entrypoint of the deployment
	Entrypoint string `json:"entrypoint"`
	// If the Build uses the `@vercel/static` Runtime, it contains a hashed string of all outputs
	Fingerprint *string `json:"fingerprint,omitempty"`
	// The unique identifier of the Build
	ID string `json:"id"`
	// A list of outputs for the Build that can be either Serverless Functions or static files
	Output []GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsOutput `json:"output"`
	// The state of the deployment depending on the process of deploying, or if it is ready or in an error state
	ReadyState GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsReadyState `json:"readyState"`
	// The time at which the Build state was last modified
	ReadyStateAt *int64 `json:"readyStateAt,omitempty"`
	// The time at which the Build was scheduled to be built
	ScheduledAt *int64 `json:"scheduledAt,omitempty"`
	// The Runtime the Build used to generate the output
	Use *string `json:"use,omitempty"`
}

func (o *GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuilds) GetConfig() *GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuilds) GetCopiedFrom() *string {
	if o == nil {
		return nil
	}
	return o.CopiedFrom
}

func (o *GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuilds) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuilds) GetCreatedIn() *string {
	if o == nil {
		return nil
	}
	return o.CreatedIn
}

func (o *GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuilds) GetDeployedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.DeployedAt
}

func (o *GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuilds) GetDeploymentID() string {
	if o == nil {
		return ""
	}
	return o.DeploymentID
}

func (o *GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuilds) GetEntrypoint() string {
	if o == nil {
		return ""
	}
	return o.Entrypoint
}

func (o *GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuilds) GetFingerprint() *string {
	if o == nil {
		return nil
	}
	return o.Fingerprint
}

func (o *GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuilds) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuilds) GetOutput() []GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsOutput {
	if o == nil {
		return []GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsOutput{}
	}
	return o.Output
}

func (o *GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuilds) GetReadyState() GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsReadyState {
	if o == nil {
		return GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuildsReadyState("")
	}
	return o.ReadyState
}

func (o *GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuilds) GetReadyStateAt() *int64 {
	if o == nil {
		return nil
	}
	return o.ReadyStateAt
}

func (o *GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuilds) GetScheduledAt() *int64 {
	if o == nil {
		return nil
	}
	return o.ScheduledAt
}

func (o *GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuilds) GetUse() *string {
	if o == nil {
		return nil
	}
	return o.Use
}

type GetDeploymentsDeploymentIDBuilds200ApplicationJSON struct {
	Builds []GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuilds `json:"builds"`
}

func (o *GetDeploymentsDeploymentIDBuilds200ApplicationJSON) GetBuilds() []GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuilds {
	if o == nil {
		return []GetDeploymentsDeploymentIDBuilds200ApplicationJSONBuilds{}
	}
	return o.Builds
}

type GetDeploymentsDeploymentIDBuildsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse                                              *http.Response
	GetDeploymentsDeploymentIDBuilds200ApplicationJSONObject *GetDeploymentsDeploymentIDBuilds200ApplicationJSON
}

func (o *GetDeploymentsDeploymentIDBuildsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetDeploymentsDeploymentIDBuildsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetDeploymentsDeploymentIDBuildsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetDeploymentsDeploymentIDBuildsResponse) GetGetDeploymentsDeploymentIDBuilds200ApplicationJSONObject() *GetDeploymentsDeploymentIDBuilds200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetDeploymentsDeploymentIDBuilds200ApplicationJSONObject
}
