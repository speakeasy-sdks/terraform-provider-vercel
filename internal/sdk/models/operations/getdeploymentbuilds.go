// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GetDeploymentBuildsRequest struct {
	// The deployment unique identifier
	DeploymentID string `pathParam:"style=simple,explode=false,name=deploymentId"`
}

func (o *GetDeploymentBuildsRequest) GetDeploymentID() string {
	if o == nil {
		return ""
	}
	return o.DeploymentID
}

// ReadyState - The state of the deployment depending on the process of deploying, or if it is ready or in an error state
type ReadyState string

const (
	ReadyStateInitializing ReadyState = "INITIALIZING"
	ReadyStateBuilding     ReadyState = "BUILDING"
	ReadyStateUploading    ReadyState = "UPLOADING"
	ReadyStateDeploying    ReadyState = "DEPLOYING"
	ReadyStateReady        ReadyState = "READY"
	ReadyStateArchived     ReadyState = "ARCHIVED"
	ReadyStateError        ReadyState = "ERROR"
	ReadyStateQueued       ReadyState = "QUEUED"
	ReadyStateCanceled     ReadyState = "CANCELED"
)

func (e ReadyState) ToPointer() *ReadyState {
	return &e
}
func (e *ReadyState) UnmarshalJSON(data []byte) error {
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
		*e = ReadyState(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ReadyState: %v", v)
	}
}

// Config - An object that contains the Build's configuration
type Config struct {
	DistDir           *string `json:"distDir,omitempty"`
	ForceBuildIn      *string `json:"forceBuildIn,omitempty"`
	ReuseWorkPathFrom *string `json:"reuseWorkPathFrom,omitempty"`
	ZeroConfig        *bool   `json:"zeroConfig,omitempty"`
}

func (o *Config) GetDistDir() *string {
	if o == nil {
		return nil
	}
	return o.DistDir
}

func (o *Config) GetForceBuildIn() *string {
	if o == nil {
		return nil
	}
	return o.ForceBuildIn
}

func (o *Config) GetReuseWorkPathFrom() *string {
	if o == nil {
		return nil
	}
	return o.ReuseWorkPathFrom
}

func (o *Config) GetZeroConfig() *bool {
	if o == nil {
		return nil
	}
	return o.ZeroConfig
}

// GetDeploymentBuildsType - The type of the output
type GetDeploymentBuildsType string

const (
	GetDeploymentBuildsTypeLambda GetDeploymentBuildsType = "lambda"
	GetDeploymentBuildsTypeFile   GetDeploymentBuildsType = "file"
	GetDeploymentBuildsTypeEdge   GetDeploymentBuildsType = "edge"
)

func (e GetDeploymentBuildsType) ToPointer() *GetDeploymentBuildsType {
	return &e
}
func (e *GetDeploymentBuildsType) UnmarshalJSON(data []byte) error {
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
		*e = GetDeploymentBuildsType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetDeploymentBuildsType: %v", v)
	}
}

// Lambda - If the output is a Serverless Function, an object containing the name, location and memory size of the function
type Lambda struct {
	FunctionName string   `json:"functionName"`
	DeployedTo   []string `json:"deployedTo"`
	MemorySize   *float64 `json:"memorySize,omitempty"`
	Timeout      *float64 `json:"timeout,omitempty"`
	Layers       []string `json:"layers,omitempty"`
}

func (o *Lambda) GetFunctionName() string {
	if o == nil {
		return ""
	}
	return o.FunctionName
}

func (o *Lambda) GetDeployedTo() []string {
	if o == nil {
		return []string{}
	}
	return o.DeployedTo
}

func (o *Lambda) GetMemorySize() *float64 {
	if o == nil {
		return nil
	}
	return o.MemorySize
}

func (o *Lambda) GetTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.Timeout
}

func (o *Lambda) GetLayers() []string {
	if o == nil {
		return nil
	}
	return o.Layers
}

// Edge - Exists if the output is an edge function.
type Edge struct {
	// The regions where the edge function will be invoked. Only exists if the edge function as a regional edge function, see: https://vercel.com/docs/concepts/edge-network/regions#setting-edge-function-regions
	Regions []string `json:"regions"`
}

func (o *Edge) GetRegions() []string {
	if o == nil {
		return nil
	}
	return o.Regions
}

// GetDeploymentBuildsOutput - A list of outputs for the Build that can be either Serverless Functions or static files
type GetDeploymentBuildsOutput struct {
	// The type of the output
	Type *GetDeploymentBuildsType `json:"type,omitempty"`
	// The absolute path of the file or Serverless Function
	Path string `json:"path"`
	// The SHA1 of the file
	Digest string `json:"digest"`
	// The POSIX file permissions
	Mode float64 `json:"mode"`
	// The size of the file in bytes
	Size *float64 `json:"size,omitempty"`
	// If the output is a Serverless Function, an object containing the name, location and memory size of the function
	Lambda *Lambda `json:"lambda,omitempty"`
	// Exists if the output is an edge function.
	Edge *Edge `json:"edge,omitempty"`
}

func (o *GetDeploymentBuildsOutput) GetType() *GetDeploymentBuildsType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *GetDeploymentBuildsOutput) GetPath() string {
	if o == nil {
		return ""
	}
	return o.Path
}

func (o *GetDeploymentBuildsOutput) GetDigest() string {
	if o == nil {
		return ""
	}
	return o.Digest
}

func (o *GetDeploymentBuildsOutput) GetMode() float64 {
	if o == nil {
		return 0.0
	}
	return o.Mode
}

func (o *GetDeploymentBuildsOutput) GetSize() *float64 {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *GetDeploymentBuildsOutput) GetLambda() *Lambda {
	if o == nil {
		return nil
	}
	return o.Lambda
}

func (o *GetDeploymentBuildsOutput) GetEdge() *Edge {
	if o == nil {
		return nil
	}
	return o.Edge
}

// Builds - An object representing a Build on Vercel
type Builds struct {
	// The unique identifier of the Build
	ID string `json:"id"`
	// The unique identifier of the deployment
	DeploymentID string `json:"deploymentId"`
	// The entrypoint of the deployment
	Entrypoint string `json:"entrypoint"`
	// The state of the deployment depending on the process of deploying, or if it is ready or in an error state
	ReadyState ReadyState `json:"readyState"`
	// The time at which the Build state was last modified
	ReadyStateAt *float64 `json:"readyStateAt,omitempty"`
	// The time at which the Build was scheduled to be built
	ScheduledAt *float64 `json:"scheduledAt,omitempty"`
	// The time at which the Build was created
	CreatedAt *float64 `json:"createdAt,omitempty"`
	// The time at which the Build was deployed
	DeployedAt *float64 `json:"deployedAt,omitempty"`
	// The region where the Build was first created
	CreatedIn *string `json:"createdIn,omitempty"`
	// The Runtime the Build used to generate the output
	Use *string `json:"use,omitempty"`
	// An object that contains the Build's configuration
	Config *Config `json:"config,omitempty"`
	// A list of outputs for the Build that can be either Serverless Functions or static files
	Output []GetDeploymentBuildsOutput `json:"output"`
	// If the Build uses the `@vercel/static` Runtime, it contains a hashed string of all outputs
	Fingerprint *string `json:"fingerprint,omitempty"`
	CopiedFrom  *string `json:"copiedFrom,omitempty"`
}

func (o *Builds) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Builds) GetDeploymentID() string {
	if o == nil {
		return ""
	}
	return o.DeploymentID
}

func (o *Builds) GetEntrypoint() string {
	if o == nil {
		return ""
	}
	return o.Entrypoint
}

func (o *Builds) GetReadyState() ReadyState {
	if o == nil {
		return ReadyState("")
	}
	return o.ReadyState
}

func (o *Builds) GetReadyStateAt() *float64 {
	if o == nil {
		return nil
	}
	return o.ReadyStateAt
}

func (o *Builds) GetScheduledAt() *float64 {
	if o == nil {
		return nil
	}
	return o.ScheduledAt
}

func (o *Builds) GetCreatedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *Builds) GetDeployedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.DeployedAt
}

func (o *Builds) GetCreatedIn() *string {
	if o == nil {
		return nil
	}
	return o.CreatedIn
}

func (o *Builds) GetUse() *string {
	if o == nil {
		return nil
	}
	return o.Use
}

func (o *Builds) GetConfig() *Config {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *Builds) GetOutput() []GetDeploymentBuildsOutput {
	if o == nil {
		return []GetDeploymentBuildsOutput{}
	}
	return o.Output
}

func (o *Builds) GetFingerprint() *string {
	if o == nil {
		return nil
	}
	return o.Fingerprint
}

func (o *Builds) GetCopiedFrom() *string {
	if o == nil {
		return nil
	}
	return o.CopiedFrom
}

type GetDeploymentBuildsResponseBody struct {
	Builds []Builds `json:"builds"`
}

func (o *GetDeploymentBuildsResponseBody) GetBuilds() []Builds {
	if o == nil {
		return []Builds{}
	}
	return o.Builds
}

type GetDeploymentBuildsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	Object      *GetDeploymentBuildsResponseBody
}

func (o *GetDeploymentBuildsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetDeploymentBuildsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetDeploymentBuildsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetDeploymentBuildsResponse) GetObject() *GetDeploymentBuildsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
