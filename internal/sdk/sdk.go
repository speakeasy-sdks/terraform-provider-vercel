// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdk

import (
	"bytes"
	"context"
	"fmt"
	"github.com/zchee/terraform-provider-vercel/internal/sdk/pkg/models/operations"
	"github.com/zchee/terraform-provider-vercel/internal/sdk/pkg/models/sdkerrors"
	"github.com/zchee/terraform-provider-vercel/internal/sdk/pkg/models/shared"
	"github.com/zchee/terraform-provider-vercel/internal/sdk/pkg/utils"
	"io"
	"net/http"
	"strings"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	// Production API
	"https://api.vercel.com",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

type sdkConfiguration struct {
	DefaultClient     HTTPClient
	SecurityClient    HTTPClient
	Security          func(context.Context) (interface{}, error)
	ServerURL         string
	ServerIndex       int
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
	RetryConfig       *utils.RetryConfig
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], nil
}

// Vercel API: Vercel combines the best developer experience with an obsessive focus on end-user performance. Our platform enables frontend teams to do their best work.
type Vercel struct {
	Authentication *Authentication
	Projects       *Projects
	Checks         *Checks
	DNS            *DNS
	EdgeConfig     *EdgeConfig
	Integrations   *Integrations
	LogDrains      *LogDrains
	ProjectMembers *ProjectMembers
	Teams          *Teams
	User           *User
	Deployments    *Deployments
	Aliases        *Aliases
	Secrets        *Secrets
	Domains        *Domains
	Certs          *Certs
	Artifacts      *Artifacts

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*Vercel)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *Vercel) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *Vercel) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *Vercel) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *Vercel) {
		sdk.sdkConfiguration.DefaultClient = client
	}
}

func withSecurity(security interface{}) func(context.Context) (interface{}, error) {
	return func(context.Context) (interface{}, error) {
		return &security, nil
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(security shared.Security) SDKOption {
	return func(sdk *Vercel) {
		sdk.sdkConfiguration.Security = withSecurity(security)
	}
}

// WithSecuritySource configures the SDK to invoke the Security Source function on each method call to determine authentication
func WithSecuritySource(security func(context.Context) (shared.Security, error)) SDKOption {
	return func(sdk *Vercel) {
		sdk.sdkConfiguration.Security = func(ctx context.Context) (interface{}, error) {
			return security(ctx)
		}
	}
}

func WithRetryConfig(retryConfig utils.RetryConfig) SDKOption {
	return func(sdk *Vercel) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *Vercel {
	sdk := &Vercel{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "0.0.1",
			SDKVersion:        "0.7.3",
			GenVersion:        "2.192.1",
			UserAgent:         "speakeasy-sdk/go 0.7.3 2.192.1 0.0.1 vercel",
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.DefaultClient == nil {
		sdk.sdkConfiguration.DefaultClient = &http.Client{Timeout: 60 * time.Second}
	}
	if sdk.sdkConfiguration.SecurityClient == nil {
		if sdk.sdkConfiguration.Security != nil {
			sdk.sdkConfiguration.SecurityClient = utils.ConfigureSecurityClient(sdk.sdkConfiguration.DefaultClient, sdk.sdkConfiguration.Security)
		} else {
			sdk.sdkConfiguration.SecurityClient = sdk.sdkConfiguration.DefaultClient
		}
	}

	sdk.Authentication = newAuthentication(sdk.sdkConfiguration)

	sdk.Projects = newProjects(sdk.sdkConfiguration)

	sdk.Checks = newChecks(sdk.sdkConfiguration)

	sdk.DNS = newDNS(sdk.sdkConfiguration)

	sdk.EdgeConfig = newEdgeConfig(sdk.sdkConfiguration)

	sdk.Integrations = newIntegrations(sdk.sdkConfiguration)

	sdk.LogDrains = newLogDrains(sdk.sdkConfiguration)

	sdk.ProjectMembers = newProjectMembers(sdk.sdkConfiguration)

	sdk.Teams = newTeams(sdk.sdkConfiguration)

	sdk.User = newUser(sdk.sdkConfiguration)

	sdk.Deployments = newDeployments(sdk.sdkConfiguration)

	sdk.Aliases = newAliases(sdk.sdkConfiguration)

	sdk.Secrets = newSecrets(sdk.sdkConfiguration)

	sdk.Domains = newDomains(sdk.sdkConfiguration)

	sdk.Certs = newCerts(sdk.sdkConfiguration)

	sdk.Artifacts = newArtifacts(sdk.sdkConfiguration)

	return sdk
}

func (s *Vercel) DeleteDataCachePurgeAll(ctx context.Context, request operations.DeleteDataCachePurgeAllRequest) (*operations.DeleteDataCachePurgeAllResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/data-cache/purge-all"

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.DeleteDataCachePurgeAllResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		fallthrough
	case httpRes.StatusCode == 400:
		fallthrough
	case httpRes.StatusCode == 401:
		fallthrough
	case httpRes.StatusCode == 403:
		fallthrough
	case httpRes.StatusCode == 404:
	}

	return res, nil
}

// GetDeploymentsDeploymentIDBuilds - Retrieves the list of builds given their deployment's unique identifier. No longer listed as public API as of May 2023.
func (s *Vercel) GetDeploymentsDeploymentIDBuilds(ctx context.Context, request operations.GetDeploymentsDeploymentIDBuildsRequest) (*operations.GetDeploymentsDeploymentIDBuildsResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/deployments/{deploymentId}/builds", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetDeploymentsDeploymentIDBuildsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out operations.GetDeploymentsDeploymentIDBuildsResponseBody
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.Object = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 400:
		fallthrough
	case httpRes.StatusCode == 401:
		fallthrough
	case httpRes.StatusCode == 403:
		fallthrough
	case httpRes.StatusCode == 404:
	}

	return res, nil
}

func (s *Vercel) PatchDataCacheBillingSettings(ctx context.Context, request *operations.PatchDataCacheBillingSettingsRequestBody) (*operations.PatchDataCacheBillingSettingsResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/data-cache/billing-settings"

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, false, true, "Request", "json", `request:"mediaType=application/json"`)
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	debugBody := bytes.NewBuffer([]byte{})
	debugReader := io.TeeReader(bodyReader, debugBody)

	req, err := http.NewRequestWithContext(ctx, "PATCH", url, debugReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

	req.Header.Set("Content-Type", reqContentType)

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Request.Body = io.NopCloser(debugBody)
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.PatchDataCacheBillingSettingsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out operations.PatchDataCacheBillingSettingsResponseBody
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.Object = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 400:
		fallthrough
	case httpRes.StatusCode == 401:
		fallthrough
	case httpRes.StatusCode == 403:
		fallthrough
	case httpRes.StatusCode == 404:
	}

	return res, nil
}
