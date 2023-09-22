// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"vercel/internal/sdk/pkg/models/shared"
)

// GetDeploymentsTarget - Filter deployments based on the environment.
type GetDeploymentsTarget string

const (
	GetDeploymentsTargetProduction GetDeploymentsTarget = "production"
	GetDeploymentsTargetPreview    GetDeploymentsTarget = "preview"
)

func (e GetDeploymentsTarget) ToPointer() *GetDeploymentsTarget {
	return &e
}

func (e *GetDeploymentsTarget) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "production":
		fallthrough
	case "preview":
		*e = GetDeploymentsTarget(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetDeploymentsTarget: %v", v)
	}
}

type GetDeploymentsRequest struct {
	// Name of the deployment.
	App *string `queryParam:"style=form,explode=true,name=app"`
	// Gets the deployment created after this Date timestamp. (default: current time)
	From *int64 `queryParam:"style=form,explode=true,name=from"`
	// Maximum number of deployments to list from a request.
	Limit *int64 `queryParam:"style=form,explode=true,name=limit"`
	// Filter deployments from the given `projectId`.
	ProjectID *string `queryParam:"style=form,explode=true,name=projectId"`
	// Filter deployments based on their rollback candidacy
	RollbackCandidate *bool `queryParam:"style=form,explode=true,name=rollbackCandidate"`
	// Get Deployments created after this JavaScript timestamp.
	Since *int64 `queryParam:"style=form,explode=true,name=since"`
	// Filter deployments based on their state (`BUILDING`, `ERROR`, `INITIALIZING`, `QUEUED`, `READY`, `CANCELED`)
	State *string `queryParam:"style=form,explode=true,name=state"`
	// Filter deployments based on the environment.
	Target *GetDeploymentsTarget `queryParam:"style=form,explode=true,name=target"`
	// The Team identifier or slug to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
	// Gets the deployment created before this Date timestamp. (default: current time)
	To *int64 `queryParam:"style=form,explode=true,name=to"`
	// Get Deployments created before this JavaScript timestamp.
	Until *int64 `queryParam:"style=form,explode=true,name=until"`
	// Filter out deployments based on users who have created the deployment.
	Users *string `queryParam:"style=form,explode=true,name=users"`
}

type GetDeployments200ApplicationJSONDeploymentsAliasAssignedType string

const (
	GetDeployments200ApplicationJSONDeploymentsAliasAssignedTypeInteger GetDeployments200ApplicationJSONDeploymentsAliasAssignedType = "integer"
	GetDeployments200ApplicationJSONDeploymentsAliasAssignedTypeBoolean GetDeployments200ApplicationJSONDeploymentsAliasAssignedType = "boolean"
)

type GetDeployments200ApplicationJSONDeploymentsAliasAssigned struct {
	Integer *int64
	Boolean *bool

	Type GetDeployments200ApplicationJSONDeploymentsAliasAssignedType
}

func CreateGetDeployments200ApplicationJSONDeploymentsAliasAssignedInteger(integer int64) GetDeployments200ApplicationJSONDeploymentsAliasAssigned {
	typ := GetDeployments200ApplicationJSONDeploymentsAliasAssignedTypeInteger

	return GetDeployments200ApplicationJSONDeploymentsAliasAssigned{
		Integer: &integer,
		Type:    typ,
	}
}

func CreateGetDeployments200ApplicationJSONDeploymentsAliasAssignedBoolean(boolean bool) GetDeployments200ApplicationJSONDeploymentsAliasAssigned {
	typ := GetDeployments200ApplicationJSONDeploymentsAliasAssignedTypeBoolean

	return GetDeployments200ApplicationJSONDeploymentsAliasAssigned{
		Boolean: &boolean,
		Type:    typ,
	}
}

func (u *GetDeployments200ApplicationJSONDeploymentsAliasAssigned) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	integer := new(int64)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&integer); err == nil {
		u.Integer = integer
		u.Type = GetDeployments200ApplicationJSONDeploymentsAliasAssignedTypeInteger
		return nil
	}

	boolean := new(bool)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&boolean); err == nil {
		u.Boolean = boolean
		u.Type = GetDeployments200ApplicationJSONDeploymentsAliasAssignedTypeBoolean
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetDeployments200ApplicationJSONDeploymentsAliasAssigned) MarshalJSON() ([]byte, error) {
	if u.Integer != nil {
		return json.Marshal(u.Integer)
	}

	if u.Boolean != nil {
		return json.Marshal(u.Boolean)
	}

	return nil, nil
}

// GetDeployments200ApplicationJSONDeploymentsAliasError - An error object in case aliasing of the deployment failed.
type GetDeployments200ApplicationJSONDeploymentsAliasError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// GetDeployments200ApplicationJSONDeploymentsChecksConclusion - Conclusion for checks
type GetDeployments200ApplicationJSONDeploymentsChecksConclusion string

const (
	GetDeployments200ApplicationJSONDeploymentsChecksConclusionSucceeded GetDeployments200ApplicationJSONDeploymentsChecksConclusion = "succeeded"
	GetDeployments200ApplicationJSONDeploymentsChecksConclusionFailed    GetDeployments200ApplicationJSONDeploymentsChecksConclusion = "failed"
	GetDeployments200ApplicationJSONDeploymentsChecksConclusionSkipped   GetDeployments200ApplicationJSONDeploymentsChecksConclusion = "skipped"
	GetDeployments200ApplicationJSONDeploymentsChecksConclusionCanceled  GetDeployments200ApplicationJSONDeploymentsChecksConclusion = "canceled"
)

func (e GetDeployments200ApplicationJSONDeploymentsChecksConclusion) ToPointer() *GetDeployments200ApplicationJSONDeploymentsChecksConclusion {
	return &e
}

func (e *GetDeployments200ApplicationJSONDeploymentsChecksConclusion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "succeeded":
		fallthrough
	case "failed":
		fallthrough
	case "skipped":
		fallthrough
	case "canceled":
		*e = GetDeployments200ApplicationJSONDeploymentsChecksConclusion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetDeployments200ApplicationJSONDeploymentsChecksConclusion: %v", v)
	}
}

// GetDeployments200ApplicationJSONDeploymentsChecksState - State of all registered checks
type GetDeployments200ApplicationJSONDeploymentsChecksState string

const (
	GetDeployments200ApplicationJSONDeploymentsChecksStateRegistered GetDeployments200ApplicationJSONDeploymentsChecksState = "registered"
	GetDeployments200ApplicationJSONDeploymentsChecksStateRunning    GetDeployments200ApplicationJSONDeploymentsChecksState = "running"
	GetDeployments200ApplicationJSONDeploymentsChecksStateCompleted  GetDeployments200ApplicationJSONDeploymentsChecksState = "completed"
)

func (e GetDeployments200ApplicationJSONDeploymentsChecksState) ToPointer() *GetDeployments200ApplicationJSONDeploymentsChecksState {
	return &e
}

func (e *GetDeployments200ApplicationJSONDeploymentsChecksState) UnmarshalJSON(data []byte) error {
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
		*e = GetDeployments200ApplicationJSONDeploymentsChecksState(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetDeployments200ApplicationJSONDeploymentsChecksState: %v", v)
	}
}

// GetDeployments200ApplicationJSONDeploymentsCreator - Metadata information of the user who created the deployment.
type GetDeployments200ApplicationJSONDeploymentsCreator struct {
	// The email address of the user.
	Email *string `json:"email,omitempty"`
	// The GitHub login of the user.
	GithubLogin *string `json:"githubLogin,omitempty"`
	// The GitLab login of the user.
	GitlabLogin *string `json:"gitlabLogin,omitempty"`
	// The unique identifier of the user.
	UID string `json:"uid"`
	// The username of the user.
	Username *string `json:"username,omitempty"`
}

type GetDeployments200ApplicationJSONDeploymentsProjectSettingsFramework string

const (
	GetDeployments200ApplicationJSONDeploymentsProjectSettingsFrameworkBlitzjs        GetDeployments200ApplicationJSONDeploymentsProjectSettingsFramework = "blitzjs"
	GetDeployments200ApplicationJSONDeploymentsProjectSettingsFrameworkNextjs         GetDeployments200ApplicationJSONDeploymentsProjectSettingsFramework = "nextjs"
	GetDeployments200ApplicationJSONDeploymentsProjectSettingsFrameworkGatsby         GetDeployments200ApplicationJSONDeploymentsProjectSettingsFramework = "gatsby"
	GetDeployments200ApplicationJSONDeploymentsProjectSettingsFrameworkRemix          GetDeployments200ApplicationJSONDeploymentsProjectSettingsFramework = "remix"
	GetDeployments200ApplicationJSONDeploymentsProjectSettingsFrameworkAstro          GetDeployments200ApplicationJSONDeploymentsProjectSettingsFramework = "astro"
	GetDeployments200ApplicationJSONDeploymentsProjectSettingsFrameworkHexo           GetDeployments200ApplicationJSONDeploymentsProjectSettingsFramework = "hexo"
	GetDeployments200ApplicationJSONDeploymentsProjectSettingsFrameworkEleventy       GetDeployments200ApplicationJSONDeploymentsProjectSettingsFramework = "eleventy"
	GetDeployments200ApplicationJSONDeploymentsProjectSettingsFrameworkDocusaurus2    GetDeployments200ApplicationJSONDeploymentsProjectSettingsFramework = "docusaurus-2"
	GetDeployments200ApplicationJSONDeploymentsProjectSettingsFrameworkDocusaurus     GetDeployments200ApplicationJSONDeploymentsProjectSettingsFramework = "docusaurus"
	GetDeployments200ApplicationJSONDeploymentsProjectSettingsFrameworkPreact         GetDeployments200ApplicationJSONDeploymentsProjectSettingsFramework = "preact"
	GetDeployments200ApplicationJSONDeploymentsProjectSettingsFrameworkSolidstart     GetDeployments200ApplicationJSONDeploymentsProjectSettingsFramework = "solidstart"
	GetDeployments200ApplicationJSONDeploymentsProjectSettingsFrameworkDojo           GetDeployments200ApplicationJSONDeploymentsProjectSettingsFramework = "dojo"
	GetDeployments200ApplicationJSONDeploymentsProjectSettingsFrameworkEmber          GetDeployments200ApplicationJSONDeploymentsProjectSettingsFramework = "ember"
	GetDeployments200ApplicationJSONDeploymentsProjectSettingsFrameworkVue            GetDeployments200ApplicationJSONDeploymentsProjectSettingsFramework = "vue"
	GetDeployments200ApplicationJSONDeploymentsProjectSettingsFrameworkScully         GetDeployments200ApplicationJSONDeploymentsProjectSettingsFramework = "scully"
	GetDeployments200ApplicationJSONDeploymentsProjectSettingsFrameworkIonicAngular   GetDeployments200ApplicationJSONDeploymentsProjectSettingsFramework = "ionic-angular"
	GetDeployments200ApplicationJSONDeploymentsProjectSettingsFrameworkAngular        GetDeployments200ApplicationJSONDeploymentsProjectSettingsFramework = "angular"
	GetDeployments200ApplicationJSONDeploymentsProjectSettingsFrameworkPolymer        GetDeployments200ApplicationJSONDeploymentsProjectSettingsFramework = "polymer"
	GetDeployments200ApplicationJSONDeploymentsProjectSettingsFrameworkSvelte         GetDeployments200ApplicationJSONDeploymentsProjectSettingsFramework = "svelte"
	GetDeployments200ApplicationJSONDeploymentsProjectSettingsFrameworkSveltekit      GetDeployments200ApplicationJSONDeploymentsProjectSettingsFramework = "sveltekit"
	GetDeployments200ApplicationJSONDeploymentsProjectSettingsFrameworkSveltekit1     GetDeployments200ApplicationJSONDeploymentsProjectSettingsFramework = "sveltekit-1"
	GetDeployments200ApplicationJSONDeploymentsProjectSettingsFrameworkIonicReact     GetDeployments200ApplicationJSONDeploymentsProjectSettingsFramework = "ionic-react"
	GetDeployments200ApplicationJSONDeploymentsProjectSettingsFrameworkCreateReactApp GetDeployments200ApplicationJSONDeploymentsProjectSettingsFramework = "create-react-app"
	GetDeployments200ApplicationJSONDeploymentsProjectSettingsFrameworkGridsome       GetDeployments200ApplicationJSONDeploymentsProjectSettingsFramework = "gridsome"
	GetDeployments200ApplicationJSONDeploymentsProjectSettingsFrameworkUmijs          GetDeployments200ApplicationJSONDeploymentsProjectSettingsFramework = "umijs"
	GetDeployments200ApplicationJSONDeploymentsProjectSettingsFrameworkSapper         GetDeployments200ApplicationJSONDeploymentsProjectSettingsFramework = "sapper"
	GetDeployments200ApplicationJSONDeploymentsProjectSettingsFrameworkSaber          GetDeployments200ApplicationJSONDeploymentsProjectSettingsFramework = "saber"
	GetDeployments200ApplicationJSONDeploymentsProjectSettingsFrameworkStencil        GetDeployments200ApplicationJSONDeploymentsProjectSettingsFramework = "stencil"
	GetDeployments200ApplicationJSONDeploymentsProjectSettingsFrameworkNuxtjs         GetDeployments200ApplicationJSONDeploymentsProjectSettingsFramework = "nuxtjs"
	GetDeployments200ApplicationJSONDeploymentsProjectSettingsFrameworkRedwoodjs      GetDeployments200ApplicationJSONDeploymentsProjectSettingsFramework = "redwoodjs"
	GetDeployments200ApplicationJSONDeploymentsProjectSettingsFrameworkHugo           GetDeployments200ApplicationJSONDeploymentsProjectSettingsFramework = "hugo"
	GetDeployments200ApplicationJSONDeploymentsProjectSettingsFrameworkJekyll         GetDeployments200ApplicationJSONDeploymentsProjectSettingsFramework = "jekyll"
	GetDeployments200ApplicationJSONDeploymentsProjectSettingsFrameworkBrunch         GetDeployments200ApplicationJSONDeploymentsProjectSettingsFramework = "brunch"
	GetDeployments200ApplicationJSONDeploymentsProjectSettingsFrameworkMiddleman      GetDeployments200ApplicationJSONDeploymentsProjectSettingsFramework = "middleman"
	GetDeployments200ApplicationJSONDeploymentsProjectSettingsFrameworkZola           GetDeployments200ApplicationJSONDeploymentsProjectSettingsFramework = "zola"
	GetDeployments200ApplicationJSONDeploymentsProjectSettingsFrameworkHydrogen       GetDeployments200ApplicationJSONDeploymentsProjectSettingsFramework = "hydrogen"
	GetDeployments200ApplicationJSONDeploymentsProjectSettingsFrameworkVite           GetDeployments200ApplicationJSONDeploymentsProjectSettingsFramework = "vite"
	GetDeployments200ApplicationJSONDeploymentsProjectSettingsFrameworkVitepress      GetDeployments200ApplicationJSONDeploymentsProjectSettingsFramework = "vitepress"
	GetDeployments200ApplicationJSONDeploymentsProjectSettingsFrameworkVuepress       GetDeployments200ApplicationJSONDeploymentsProjectSettingsFramework = "vuepress"
	GetDeployments200ApplicationJSONDeploymentsProjectSettingsFrameworkParcel         GetDeployments200ApplicationJSONDeploymentsProjectSettingsFramework = "parcel"
	GetDeployments200ApplicationJSONDeploymentsProjectSettingsFrameworkSanity         GetDeployments200ApplicationJSONDeploymentsProjectSettingsFramework = "sanity"
	GetDeployments200ApplicationJSONDeploymentsProjectSettingsFrameworkStorybook      GetDeployments200ApplicationJSONDeploymentsProjectSettingsFramework = "storybook"
)

func (e GetDeployments200ApplicationJSONDeploymentsProjectSettingsFramework) ToPointer() *GetDeployments200ApplicationJSONDeploymentsProjectSettingsFramework {
	return &e
}

func (e *GetDeployments200ApplicationJSONDeploymentsProjectSettingsFramework) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "blitzjs":
		fallthrough
	case "nextjs":
		fallthrough
	case "gatsby":
		fallthrough
	case "remix":
		fallthrough
	case "astro":
		fallthrough
	case "hexo":
		fallthrough
	case "eleventy":
		fallthrough
	case "docusaurus-2":
		fallthrough
	case "docusaurus":
		fallthrough
	case "preact":
		fallthrough
	case "solidstart":
		fallthrough
	case "dojo":
		fallthrough
	case "ember":
		fallthrough
	case "vue":
		fallthrough
	case "scully":
		fallthrough
	case "ionic-angular":
		fallthrough
	case "angular":
		fallthrough
	case "polymer":
		fallthrough
	case "svelte":
		fallthrough
	case "sveltekit":
		fallthrough
	case "sveltekit-1":
		fallthrough
	case "ionic-react":
		fallthrough
	case "create-react-app":
		fallthrough
	case "gridsome":
		fallthrough
	case "umijs":
		fallthrough
	case "sapper":
		fallthrough
	case "saber":
		fallthrough
	case "stencil":
		fallthrough
	case "nuxtjs":
		fallthrough
	case "redwoodjs":
		fallthrough
	case "hugo":
		fallthrough
	case "jekyll":
		fallthrough
	case "brunch":
		fallthrough
	case "middleman":
		fallthrough
	case "zola":
		fallthrough
	case "hydrogen":
		fallthrough
	case "vite":
		fallthrough
	case "vitepress":
		fallthrough
	case "vuepress":
		fallthrough
	case "parcel":
		fallthrough
	case "sanity":
		fallthrough
	case "storybook":
		*e = GetDeployments200ApplicationJSONDeploymentsProjectSettingsFramework(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetDeployments200ApplicationJSONDeploymentsProjectSettingsFramework: %v", v)
	}
}

// GetDeployments200ApplicationJSONDeploymentsProjectSettingsGitComments - Since June '23
type GetDeployments200ApplicationJSONDeploymentsProjectSettingsGitComments struct {
	// Whether the Vercel bot should comment on commits
	OnCommit bool `json:"onCommit"`
	// Whether the Vercel bot should comment on PRs
	OnPullRequest bool `json:"onPullRequest"`
}

type GetDeployments200ApplicationJSONDeploymentsProjectSettingsNodeVersion string

const (
	GetDeployments200ApplicationJSONDeploymentsProjectSettingsNodeVersionEighteenX GetDeployments200ApplicationJSONDeploymentsProjectSettingsNodeVersion = "18.x"
	GetDeployments200ApplicationJSONDeploymentsProjectSettingsNodeVersionSixteenX  GetDeployments200ApplicationJSONDeploymentsProjectSettingsNodeVersion = "16.x"
	GetDeployments200ApplicationJSONDeploymentsProjectSettingsNodeVersionFourteenX GetDeployments200ApplicationJSONDeploymentsProjectSettingsNodeVersion = "14.x"
	GetDeployments200ApplicationJSONDeploymentsProjectSettingsNodeVersionTwelveX   GetDeployments200ApplicationJSONDeploymentsProjectSettingsNodeVersion = "12.x"
	GetDeployments200ApplicationJSONDeploymentsProjectSettingsNodeVersionTenX      GetDeployments200ApplicationJSONDeploymentsProjectSettingsNodeVersion = "10.x"
)

func (e GetDeployments200ApplicationJSONDeploymentsProjectSettingsNodeVersion) ToPointer() *GetDeployments200ApplicationJSONDeploymentsProjectSettingsNodeVersion {
	return &e
}

func (e *GetDeployments200ApplicationJSONDeploymentsProjectSettingsNodeVersion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "18.x":
		fallthrough
	case "16.x":
		fallthrough
	case "14.x":
		fallthrough
	case "12.x":
		fallthrough
	case "10.x":
		*e = GetDeployments200ApplicationJSONDeploymentsProjectSettingsNodeVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetDeployments200ApplicationJSONDeploymentsProjectSettingsNodeVersion: %v", v)
	}
}

// GetDeployments200ApplicationJSONDeploymentsProjectSettings - The project settings which was used for this deployment
type GetDeployments200ApplicationJSONDeploymentsProjectSettings struct {
	BuildCommand                  *string                                                              `json:"buildCommand,omitempty"`
	CommandForIgnoringBuildStep   *string                                                              `json:"commandForIgnoringBuildStep,omitempty"`
	CreatedAt                     *int64                                                               `json:"createdAt,omitempty"`
	CustomerSupportCodeVisibility *bool                                                                `json:"customerSupportCodeVisibility,omitempty"`
	DevCommand                    *string                                                              `json:"devCommand,omitempty"`
	Framework                     *GetDeployments200ApplicationJSONDeploymentsProjectSettingsFramework `json:"framework,omitempty"`
	// Since June '23
	GitComments                     *GetDeployments200ApplicationJSONDeploymentsProjectSettingsGitComments `json:"gitComments,omitempty"`
	GitForkProtection               *bool                                                                  `json:"gitForkProtection,omitempty"`
	GitLFS                          *bool                                                                  `json:"gitLFS,omitempty"`
	InstallCommand                  *string                                                                `json:"installCommand,omitempty"`
	NodeVersion                     *GetDeployments200ApplicationJSONDeploymentsProjectSettingsNodeVersion `json:"nodeVersion,omitempty"`
	OutputDirectory                 *string                                                                `json:"outputDirectory,omitempty"`
	PublicSource                    *bool                                                                  `json:"publicSource,omitempty"`
	RootDirectory                   *string                                                                `json:"rootDirectory,omitempty"`
	ServerlessFunctionRegion        *string                                                                `json:"serverlessFunctionRegion,omitempty"`
	SkipGitConnectDuringLink        *bool                                                                  `json:"skipGitConnectDuringLink,omitempty"`
	SourceFilesOutsideRootDirectory *bool                                                                  `json:"sourceFilesOutsideRootDirectory,omitempty"`
}

// GetDeployments200ApplicationJSONDeploymentsReadySubstate - Since June 2023 Substate of deployment when readyState is 'READY' Tracks whether or not deployment has seen production traffic: - STAGED: never seen production traffic - PROMOTED: has seen production traffic
type GetDeployments200ApplicationJSONDeploymentsReadySubstate string

const (
	GetDeployments200ApplicationJSONDeploymentsReadySubstateStaged   GetDeployments200ApplicationJSONDeploymentsReadySubstate = "STAGED"
	GetDeployments200ApplicationJSONDeploymentsReadySubstatePromoted GetDeployments200ApplicationJSONDeploymentsReadySubstate = "PROMOTED"
)

func (e GetDeployments200ApplicationJSONDeploymentsReadySubstate) ToPointer() *GetDeployments200ApplicationJSONDeploymentsReadySubstate {
	return &e
}

func (e *GetDeployments200ApplicationJSONDeploymentsReadySubstate) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "STAGED":
		fallthrough
	case "PROMOTED":
		*e = GetDeployments200ApplicationJSONDeploymentsReadySubstate(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetDeployments200ApplicationJSONDeploymentsReadySubstate: %v", v)
	}
}

// GetDeployments200ApplicationJSONDeploymentsSource - The source of the deployment.
type GetDeployments200ApplicationJSONDeploymentsSource string

const (
	GetDeployments200ApplicationJSONDeploymentsSourceCli        GetDeployments200ApplicationJSONDeploymentsSource = "cli"
	GetDeployments200ApplicationJSONDeploymentsSourceGit        GetDeployments200ApplicationJSONDeploymentsSource = "git"
	GetDeployments200ApplicationJSONDeploymentsSourceImport     GetDeployments200ApplicationJSONDeploymentsSource = "import"
	GetDeployments200ApplicationJSONDeploymentsSourceImportRepo GetDeployments200ApplicationJSONDeploymentsSource = "import/repo"
	GetDeployments200ApplicationJSONDeploymentsSourceCloneRepo  GetDeployments200ApplicationJSONDeploymentsSource = "clone/repo"
)

func (e GetDeployments200ApplicationJSONDeploymentsSource) ToPointer() *GetDeployments200ApplicationJSONDeploymentsSource {
	return &e
}

func (e *GetDeployments200ApplicationJSONDeploymentsSource) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "cli":
		fallthrough
	case "git":
		fallthrough
	case "import":
		fallthrough
	case "import/repo":
		fallthrough
	case "clone/repo":
		*e = GetDeployments200ApplicationJSONDeploymentsSource(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetDeployments200ApplicationJSONDeploymentsSource: %v", v)
	}
}

// GetDeployments200ApplicationJSONDeploymentsState - In which state is the deployment.
type GetDeployments200ApplicationJSONDeploymentsState string

const (
	GetDeployments200ApplicationJSONDeploymentsStateBuilding     GetDeployments200ApplicationJSONDeploymentsState = "BUILDING"
	GetDeployments200ApplicationJSONDeploymentsStateError        GetDeployments200ApplicationJSONDeploymentsState = "ERROR"
	GetDeployments200ApplicationJSONDeploymentsStateInitializing GetDeployments200ApplicationJSONDeploymentsState = "INITIALIZING"
	GetDeployments200ApplicationJSONDeploymentsStateQueued       GetDeployments200ApplicationJSONDeploymentsState = "QUEUED"
	GetDeployments200ApplicationJSONDeploymentsStateReady        GetDeployments200ApplicationJSONDeploymentsState = "READY"
	GetDeployments200ApplicationJSONDeploymentsStateCanceled     GetDeployments200ApplicationJSONDeploymentsState = "CANCELED"
)

func (e GetDeployments200ApplicationJSONDeploymentsState) ToPointer() *GetDeployments200ApplicationJSONDeploymentsState {
	return &e
}

func (e *GetDeployments200ApplicationJSONDeploymentsState) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "BUILDING":
		fallthrough
	case "ERROR":
		fallthrough
	case "INITIALIZING":
		fallthrough
	case "QUEUED":
		fallthrough
	case "READY":
		fallthrough
	case "CANCELED":
		*e = GetDeployments200ApplicationJSONDeploymentsState(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetDeployments200ApplicationJSONDeploymentsState: %v", v)
	}
}

// GetDeployments200ApplicationJSONDeploymentsTarget - On which environment has the deployment been deployed to.
type GetDeployments200ApplicationJSONDeploymentsTarget string

const (
	GetDeployments200ApplicationJSONDeploymentsTargetProduction GetDeployments200ApplicationJSONDeploymentsTarget = "production"
	GetDeployments200ApplicationJSONDeploymentsTargetStaging    GetDeployments200ApplicationJSONDeploymentsTarget = "staging"
)

func (e GetDeployments200ApplicationJSONDeploymentsTarget) ToPointer() *GetDeployments200ApplicationJSONDeploymentsTarget {
	return &e
}

func (e *GetDeployments200ApplicationJSONDeploymentsTarget) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "production":
		fallthrough
	case "staging":
		*e = GetDeployments200ApplicationJSONDeploymentsTarget(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetDeployments200ApplicationJSONDeploymentsTarget: %v", v)
	}
}

// GetDeployments200ApplicationJSONDeploymentsType - The type of the deployment.
type GetDeployments200ApplicationJSONDeploymentsType string

const (
	GetDeployments200ApplicationJSONDeploymentsTypeLambdas GetDeployments200ApplicationJSONDeploymentsType = "LAMBDAS"
)

func (e GetDeployments200ApplicationJSONDeploymentsType) ToPointer() *GetDeployments200ApplicationJSONDeploymentsType {
	return &e
}

func (e *GetDeployments200ApplicationJSONDeploymentsType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "LAMBDAS":
		*e = GetDeployments200ApplicationJSONDeploymentsType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetDeployments200ApplicationJSONDeploymentsType: %v", v)
	}
}

type GetDeployments200ApplicationJSONDeployments struct {
	AliasAssigned *GetDeployments200ApplicationJSONDeploymentsAliasAssigned `json:"aliasAssigned,omitempty"`
	// An error object in case aliasing of the deployment failed.
	AliasError *GetDeployments200ApplicationJSONDeploymentsAliasError `json:"aliasError,omitempty"`
	// Timestamp of when the deployment started building at.
	BuildingAt *int64 `json:"buildingAt,omitempty"`
	// Conclusion for checks
	ChecksConclusion *GetDeployments200ApplicationJSONDeploymentsChecksConclusion `json:"checksConclusion,omitempty"`
	// State of all registered checks
	ChecksState *GetDeployments200ApplicationJSONDeploymentsChecksState `json:"checksState,omitempty"`
	// The flag saying if Vercel Connect configuration is used for builds
	ConnectBuildsEnabled *bool `json:"connectBuildsEnabled,omitempty"`
	// The ID of Vercel Connect configuration used for this deployment
	ConnectConfigurationID *string `json:"connectConfigurationId,omitempty"`
	// Timestamp of when the deployment got created.
	Created int64 `json:"created"`
	// Timestamp of when the deployment got created.
	CreatedAt *int64 `json:"createdAt,omitempty"`
	// Metadata information of the user who created the deployment.
	Creator GetDeployments200ApplicationJSONDeploymentsCreator `json:"creator"`
	// Vercel URL to inspect the deployment.
	InspectorURL *string `json:"inspectorUrl"`
	// Deployment can be used for instant rollback
	IsRollbackCandidate *bool `json:"isRollbackCandidate,omitempty"`
	// Metadata information from the Git provider.
	Meta map[string]string `json:"meta,omitempty"`
	// The name of the deployment.
	Name string `json:"name"`
	// The project settings which was used for this deployment
	ProjectSettings *GetDeployments200ApplicationJSONDeploymentsProjectSettings `json:"projectSettings,omitempty"`
	// Timestamp of when the deployment got ready.
	Ready *int64 `json:"ready,omitempty"`
	// Since June 2023 Substate of deployment when readyState is 'READY' Tracks whether or not deployment has seen production traffic: - STAGED: never seen production traffic - PROMOTED: has seen production traffic
	ReadySubstate *GetDeployments200ApplicationJSONDeploymentsReadySubstate `json:"readySubstate,omitempty"`
	// The source of the deployment.
	Source *GetDeployments200ApplicationJSONDeploymentsSource `json:"source,omitempty"`
	// In which state is the deployment.
	State *GetDeployments200ApplicationJSONDeploymentsState `json:"state,omitempty"`
	// On which environment has the deployment been deployed to.
	Target *GetDeployments200ApplicationJSONDeploymentsTarget `json:"target,omitempty"`
	// The type of the deployment.
	Type GetDeployments200ApplicationJSONDeploymentsType `json:"type"`
	// The unique identifier of the deployment.
	UID string `json:"uid"`
	// The URL of the deployment.
	URL string `json:"url"`
}

type GetDeployments200ApplicationJSON struct {
	Deployments []GetDeployments200ApplicationJSONDeployments `json:"deployments"`
	// This object contains information related to the pagination of the current request, including the necessary parameters to get the next or previous page of data.
	Pagination shared.Pagination `json:"pagination"`
}

type GetDeploymentsResponse struct {
	ContentType                            string
	StatusCode                             int
	RawResponse                            *http.Response
	GetDeployments200ApplicationJSONObject *GetDeployments200ApplicationJSON
}
