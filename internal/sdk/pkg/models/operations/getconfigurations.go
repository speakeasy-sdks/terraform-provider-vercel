// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type GetConfigurationsSecurity struct {
	BearerToken string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

// GetConfigurationsView
type GetConfigurationsView string

const (
	GetConfigurationsViewAccount GetConfigurationsView = "account"
	GetConfigurationsViewProject GetConfigurationsView = "project"
)

func (e GetConfigurationsView) ToPointer() *GetConfigurationsView {
	return &e
}

func (e *GetConfigurationsView) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "account":
		fallthrough
	case "project":
		*e = GetConfigurationsView(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetConfigurationsView: %v", v)
	}
}

type GetConfigurationsRequest struct {
	// The Team identifier or slug to perform the request on behalf of.
	TeamID *string               `queryParam:"style=form,explode=true,name=teamId"`
	View   GetConfigurationsView `queryParam:"style=form,explode=true,name=view"`
}

type GetConfigurations200ApplicationJSON2DisabledReason string

const (
	GetConfigurations200ApplicationJSON2DisabledReasonLogDrainHighErrorRate         GetConfigurations200ApplicationJSON2DisabledReason = "log-drain-high-error-rate"
	GetConfigurations200ApplicationJSON2DisabledReasonLogDrainsAddOnDisabledByOwner GetConfigurations200ApplicationJSON2DisabledReason = "log-drains-add-on-disabled-by-owner"
	GetConfigurations200ApplicationJSON2DisabledReasonAccountPlanDowngrade          GetConfigurations200ApplicationJSON2DisabledReason = "account-plan-downgrade"
	GetConfigurations200ApplicationJSON2DisabledReasonDisabledByAdmin               GetConfigurations200ApplicationJSON2DisabledReason = "disabled-by-admin"
	GetConfigurations200ApplicationJSON2DisabledReasonOriginalOwnerLeftTheTeam      GetConfigurations200ApplicationJSON2DisabledReason = "original-owner-left-the-team"
)

func (e GetConfigurations200ApplicationJSON2DisabledReason) ToPointer() *GetConfigurations200ApplicationJSON2DisabledReason {
	return &e
}

func (e *GetConfigurations200ApplicationJSON2DisabledReason) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "log-drain-high-error-rate":
		fallthrough
	case "log-drains-add-on-disabled-by-owner":
		fallthrough
	case "account-plan-downgrade":
		fallthrough
	case "disabled-by-admin":
		fallthrough
	case "original-owner-left-the-team":
		*e = GetConfigurations200ApplicationJSON2DisabledReason(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetConfigurations200ApplicationJSON2DisabledReason: %v", v)
	}
}

type GetConfigurations200ApplicationJSON2Integration struct {
	AssignedBetaLabelAt *int64   `json:"assignedBetaLabelAt,omitempty"`
	Category            string   `json:"category"`
	Flags               []string `json:"flags,omitempty"`
	Icon                string   `json:"icon"`
	IsLegacy            bool     `json:"isLegacy"`
	Name                string   `json:"name"`
}

type GetConfigurations200ApplicationJSON2ScopesQueueScopesAdded string

const (
	GetConfigurations200ApplicationJSON2ScopesQueueScopesAddedReadIntegrationConfiguration      GetConfigurations200ApplicationJSON2ScopesQueueScopesAdded = "read:integration-configuration"
	GetConfigurations200ApplicationJSON2ScopesQueueScopesAddedReadWriteIntegrationConfiguration GetConfigurations200ApplicationJSON2ScopesQueueScopesAdded = "read-write:integration-configuration"
	GetConfigurations200ApplicationJSON2ScopesQueueScopesAddedReadDeployment                    GetConfigurations200ApplicationJSON2ScopesQueueScopesAdded = "read:deployment"
	GetConfigurations200ApplicationJSON2ScopesQueueScopesAddedReadWriteDeployment               GetConfigurations200ApplicationJSON2ScopesQueueScopesAdded = "read-write:deployment"
	GetConfigurations200ApplicationJSON2ScopesQueueScopesAddedReadWriteDeploymentCheck          GetConfigurations200ApplicationJSON2ScopesQueueScopesAdded = "read-write:deployment-check"
	GetConfigurations200ApplicationJSON2ScopesQueueScopesAddedReadProject                       GetConfigurations200ApplicationJSON2ScopesQueueScopesAdded = "read:project"
	GetConfigurations200ApplicationJSON2ScopesQueueScopesAddedReadWriteProject                  GetConfigurations200ApplicationJSON2ScopesQueueScopesAdded = "read-write:project"
	GetConfigurations200ApplicationJSON2ScopesQueueScopesAddedReadWriteProjectEnvVars           GetConfigurations200ApplicationJSON2ScopesQueueScopesAdded = "read-write:project-env-vars"
	GetConfigurations200ApplicationJSON2ScopesQueueScopesAddedReadWriteGlobalProjectEnvVars     GetConfigurations200ApplicationJSON2ScopesQueueScopesAdded = "read-write:global-project-env-vars"
	GetConfigurations200ApplicationJSON2ScopesQueueScopesAddedReadTeam                          GetConfigurations200ApplicationJSON2ScopesQueueScopesAdded = "read:team"
	GetConfigurations200ApplicationJSON2ScopesQueueScopesAddedReadUser                          GetConfigurations200ApplicationJSON2ScopesQueueScopesAdded = "read:user"
	GetConfigurations200ApplicationJSON2ScopesQueueScopesAddedReadWriteLogDrain                 GetConfigurations200ApplicationJSON2ScopesQueueScopesAdded = "read-write:log-drain"
	GetConfigurations200ApplicationJSON2ScopesQueueScopesAddedReadDomain                        GetConfigurations200ApplicationJSON2ScopesQueueScopesAdded = "read:domain"
	GetConfigurations200ApplicationJSON2ScopesQueueScopesAddedReadWriteDomain                   GetConfigurations200ApplicationJSON2ScopesQueueScopesAdded = "read-write:domain"
	GetConfigurations200ApplicationJSON2ScopesQueueScopesAddedReadWriteEdgeConfig               GetConfigurations200ApplicationJSON2ScopesQueueScopesAdded = "read-write:edge-config"
	GetConfigurations200ApplicationJSON2ScopesQueueScopesAddedReadWriteOtelEndpoint             GetConfigurations200ApplicationJSON2ScopesQueueScopesAdded = "read-write:otel-endpoint"
	GetConfigurations200ApplicationJSON2ScopesQueueScopesAddedReadMonitoring                    GetConfigurations200ApplicationJSON2ScopesQueueScopesAdded = "read:monitoring"
)

func (e GetConfigurations200ApplicationJSON2ScopesQueueScopesAdded) ToPointer() *GetConfigurations200ApplicationJSON2ScopesQueueScopesAdded {
	return &e
}

func (e *GetConfigurations200ApplicationJSON2ScopesQueueScopesAdded) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "read:integration-configuration":
		fallthrough
	case "read-write:integration-configuration":
		fallthrough
	case "read:deployment":
		fallthrough
	case "read-write:deployment":
		fallthrough
	case "read-write:deployment-check":
		fallthrough
	case "read:project":
		fallthrough
	case "read-write:project":
		fallthrough
	case "read-write:project-env-vars":
		fallthrough
	case "read-write:global-project-env-vars":
		fallthrough
	case "read:team":
		fallthrough
	case "read:user":
		fallthrough
	case "read-write:log-drain":
		fallthrough
	case "read:domain":
		fallthrough
	case "read-write:domain":
		fallthrough
	case "read-write:edge-config":
		fallthrough
	case "read-write:otel-endpoint":
		fallthrough
	case "read:monitoring":
		*e = GetConfigurations200ApplicationJSON2ScopesQueueScopesAdded(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetConfigurations200ApplicationJSON2ScopesQueueScopesAdded: %v", v)
	}
}

type GetConfigurations200ApplicationJSON2ScopesQueueScopesUpgraded string

const (
	GetConfigurations200ApplicationJSON2ScopesQueueScopesUpgradedReadIntegrationConfiguration      GetConfigurations200ApplicationJSON2ScopesQueueScopesUpgraded = "read:integration-configuration"
	GetConfigurations200ApplicationJSON2ScopesQueueScopesUpgradedReadWriteIntegrationConfiguration GetConfigurations200ApplicationJSON2ScopesQueueScopesUpgraded = "read-write:integration-configuration"
	GetConfigurations200ApplicationJSON2ScopesQueueScopesUpgradedReadDeployment                    GetConfigurations200ApplicationJSON2ScopesQueueScopesUpgraded = "read:deployment"
	GetConfigurations200ApplicationJSON2ScopesQueueScopesUpgradedReadWriteDeployment               GetConfigurations200ApplicationJSON2ScopesQueueScopesUpgraded = "read-write:deployment"
	GetConfigurations200ApplicationJSON2ScopesQueueScopesUpgradedReadWriteDeploymentCheck          GetConfigurations200ApplicationJSON2ScopesQueueScopesUpgraded = "read-write:deployment-check"
	GetConfigurations200ApplicationJSON2ScopesQueueScopesUpgradedReadProject                       GetConfigurations200ApplicationJSON2ScopesQueueScopesUpgraded = "read:project"
	GetConfigurations200ApplicationJSON2ScopesQueueScopesUpgradedReadWriteProject                  GetConfigurations200ApplicationJSON2ScopesQueueScopesUpgraded = "read-write:project"
	GetConfigurations200ApplicationJSON2ScopesQueueScopesUpgradedReadWriteProjectEnvVars           GetConfigurations200ApplicationJSON2ScopesQueueScopesUpgraded = "read-write:project-env-vars"
	GetConfigurations200ApplicationJSON2ScopesQueueScopesUpgradedReadWriteGlobalProjectEnvVars     GetConfigurations200ApplicationJSON2ScopesQueueScopesUpgraded = "read-write:global-project-env-vars"
	GetConfigurations200ApplicationJSON2ScopesQueueScopesUpgradedReadTeam                          GetConfigurations200ApplicationJSON2ScopesQueueScopesUpgraded = "read:team"
	GetConfigurations200ApplicationJSON2ScopesQueueScopesUpgradedReadUser                          GetConfigurations200ApplicationJSON2ScopesQueueScopesUpgraded = "read:user"
	GetConfigurations200ApplicationJSON2ScopesQueueScopesUpgradedReadWriteLogDrain                 GetConfigurations200ApplicationJSON2ScopesQueueScopesUpgraded = "read-write:log-drain"
	GetConfigurations200ApplicationJSON2ScopesQueueScopesUpgradedReadDomain                        GetConfigurations200ApplicationJSON2ScopesQueueScopesUpgraded = "read:domain"
	GetConfigurations200ApplicationJSON2ScopesQueueScopesUpgradedReadWriteDomain                   GetConfigurations200ApplicationJSON2ScopesQueueScopesUpgraded = "read-write:domain"
	GetConfigurations200ApplicationJSON2ScopesQueueScopesUpgradedReadWriteEdgeConfig               GetConfigurations200ApplicationJSON2ScopesQueueScopesUpgraded = "read-write:edge-config"
	GetConfigurations200ApplicationJSON2ScopesQueueScopesUpgradedReadWriteOtelEndpoint             GetConfigurations200ApplicationJSON2ScopesQueueScopesUpgraded = "read-write:otel-endpoint"
	GetConfigurations200ApplicationJSON2ScopesQueueScopesUpgradedReadMonitoring                    GetConfigurations200ApplicationJSON2ScopesQueueScopesUpgraded = "read:monitoring"
)

func (e GetConfigurations200ApplicationJSON2ScopesQueueScopesUpgraded) ToPointer() *GetConfigurations200ApplicationJSON2ScopesQueueScopesUpgraded {
	return &e
}

func (e *GetConfigurations200ApplicationJSON2ScopesQueueScopesUpgraded) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "read:integration-configuration":
		fallthrough
	case "read-write:integration-configuration":
		fallthrough
	case "read:deployment":
		fallthrough
	case "read-write:deployment":
		fallthrough
	case "read-write:deployment-check":
		fallthrough
	case "read:project":
		fallthrough
	case "read-write:project":
		fallthrough
	case "read-write:project-env-vars":
		fallthrough
	case "read-write:global-project-env-vars":
		fallthrough
	case "read:team":
		fallthrough
	case "read:user":
		fallthrough
	case "read-write:log-drain":
		fallthrough
	case "read:domain":
		fallthrough
	case "read-write:domain":
		fallthrough
	case "read-write:edge-config":
		fallthrough
	case "read-write:otel-endpoint":
		fallthrough
	case "read:monitoring":
		*e = GetConfigurations200ApplicationJSON2ScopesQueueScopesUpgraded(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetConfigurations200ApplicationJSON2ScopesQueueScopesUpgraded: %v", v)
	}
}

type GetConfigurations200ApplicationJSON2ScopesQueueScopes struct {
	Added    []GetConfigurations200ApplicationJSON2ScopesQueueScopesAdded    `json:"added"`
	Upgraded []GetConfigurations200ApplicationJSON2ScopesQueueScopesUpgraded `json:"upgraded"`
}

type GetConfigurations200ApplicationJSON2ScopesQueue struct {
	ConfirmedAt *int64                                                `json:"confirmedAt,omitempty"`
	Note        string                                                `json:"note"`
	RequestedAt int64                                                 `json:"requestedAt"`
	Scopes      GetConfigurations200ApplicationJSON2ScopesQueueScopes `json:"scopes"`
}

// GetConfigurations200ApplicationJSON2Source - Source defines where the configuration was installed from. It is used to analyze user engagement for integration installations in product metrics.
type GetConfigurations200ApplicationJSON2Source string

const (
	GetConfigurations200ApplicationJSON2SourceMarketplace  GetConfigurations200ApplicationJSON2Source = "marketplace"
	GetConfigurations200ApplicationJSON2SourceDeployButton GetConfigurations200ApplicationJSON2Source = "deploy-button"
	GetConfigurations200ApplicationJSON2SourceExternal     GetConfigurations200ApplicationJSON2Source = "external"
)

func (e GetConfigurations200ApplicationJSON2Source) ToPointer() *GetConfigurations200ApplicationJSON2Source {
	return &e
}

func (e *GetConfigurations200ApplicationJSON2Source) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "marketplace":
		fallthrough
	case "deploy-button":
		fallthrough
	case "external":
		*e = GetConfigurations200ApplicationJSON2Source(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetConfigurations200ApplicationJSON2Source: %v", v)
	}
}

type GetConfigurations200ApplicationJSON2Type string

const (
	GetConfigurations200ApplicationJSON2TypeIntegrationConfiguration GetConfigurations200ApplicationJSON2Type = "integration-configuration"
)

func (e GetConfigurations200ApplicationJSON2Type) ToPointer() *GetConfigurations200ApplicationJSON2Type {
	return &e
}

func (e *GetConfigurations200ApplicationJSON2Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "integration-configuration":
		*e = GetConfigurations200ApplicationJSON2Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetConfigurations200ApplicationJSON2Type: %v", v)
	}
}

type GetConfigurations200ApplicationJSON2 struct {
	// A timestamp that tells you when the configuration was installed successfully
	CompletedAt *int64 `json:"completedAt,omitempty"`
	// A timestamp that tells you when the configuration was created
	CreatedAt int64 `json:"createdAt"`
	// A timestamp that tells you when the configuration was updated.
	DeletedAt *int64 `json:"deletedAt,omitempty"`
	// A timestamp that tells you when the configuration was disabled. Note: Configurations can be disabled when the associated user loses access to a team. They do not function during this time until the configuration is 'transferred', meaning the associated user is changed to one with access to the team.
	DisabledAt     *int64                                              `json:"disabledAt,omitempty"`
	DisabledReason *GetConfigurations200ApplicationJSON2DisabledReason `json:"disabledReason,omitempty"`
	// The unique identifier of the configuration
	ID          string                                          `json:"id"`
	Integration GetConfigurations200ApplicationJSON2Integration `json:"integration"`
	// The unique identifier of the app the configuration was created for
	IntegrationID string `json:"integrationId"`
	// The user or team ID that owns the configuration
	OwnerID string `json:"ownerId"`
	// When a configuration is limited to access certain projects, this will contain each of the project ID it is allowed to access. If it is not defined, the configuration has full access.
	Projects             []string `json:"projects,omitempty"`
	RemovedLogDrainsAt   *int64   `json:"removedLogDrainsAt,omitempty"`
	RemovedProjectEnvsAt *int64   `json:"removedProjectEnvsAt,omitempty"`
	RemovedTokensAt      *int64   `json:"removedTokensAt,omitempty"`
	RemovedWebhooksAt    *int64   `json:"removedWebhooksAt,omitempty"`
	// The resources that are allowed to be accessed by the configuration.
	Scopes      []string                                          `json:"scopes"`
	ScopesQueue []GetConfigurations200ApplicationJSON2ScopesQueue `json:"scopesQueue,omitempty"`
	// The slug of the integration the configuration is created for.
	Slug string `json:"slug"`
	// Source defines where the configuration was installed from. It is used to analyze user engagement for integration installations in product metrics.
	Source *GetConfigurations200ApplicationJSON2Source `json:"source,omitempty"`
	// When the configuration was created for a team, this will show the ID of the team.
	TeamID *string                                  `json:"teamId,omitempty"`
	Type   GetConfigurations200ApplicationJSON2Type `json:"type"`
	// A timestamp that tells you when the configuration was updated.
	UpdatedAt int64 `json:"updatedAt"`
	// The ID of the user that created the configuration.
	UserID string `json:"userId"`
}

type GetConfigurations200ApplicationJSON1DisabledReason string

const (
	GetConfigurations200ApplicationJSON1DisabledReasonLogDrainHighErrorRate         GetConfigurations200ApplicationJSON1DisabledReason = "log-drain-high-error-rate"
	GetConfigurations200ApplicationJSON1DisabledReasonLogDrainsAddOnDisabledByOwner GetConfigurations200ApplicationJSON1DisabledReason = "log-drains-add-on-disabled-by-owner"
	GetConfigurations200ApplicationJSON1DisabledReasonAccountPlanDowngrade          GetConfigurations200ApplicationJSON1DisabledReason = "account-plan-downgrade"
	GetConfigurations200ApplicationJSON1DisabledReasonDisabledByAdmin               GetConfigurations200ApplicationJSON1DisabledReason = "disabled-by-admin"
	GetConfigurations200ApplicationJSON1DisabledReasonOriginalOwnerLeftTheTeam      GetConfigurations200ApplicationJSON1DisabledReason = "original-owner-left-the-team"
)

func (e GetConfigurations200ApplicationJSON1DisabledReason) ToPointer() *GetConfigurations200ApplicationJSON1DisabledReason {
	return &e
}

func (e *GetConfigurations200ApplicationJSON1DisabledReason) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "log-drain-high-error-rate":
		fallthrough
	case "log-drains-add-on-disabled-by-owner":
		fallthrough
	case "account-plan-downgrade":
		fallthrough
	case "disabled-by-admin":
		fallthrough
	case "original-owner-left-the-team":
		*e = GetConfigurations200ApplicationJSON1DisabledReason(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetConfigurations200ApplicationJSON1DisabledReason: %v", v)
	}
}

type GetConfigurations200ApplicationJSON1ScopesQueueScopesAdded string

const (
	GetConfigurations200ApplicationJSON1ScopesQueueScopesAddedReadIntegrationConfiguration      GetConfigurations200ApplicationJSON1ScopesQueueScopesAdded = "read:integration-configuration"
	GetConfigurations200ApplicationJSON1ScopesQueueScopesAddedReadWriteIntegrationConfiguration GetConfigurations200ApplicationJSON1ScopesQueueScopesAdded = "read-write:integration-configuration"
	GetConfigurations200ApplicationJSON1ScopesQueueScopesAddedReadDeployment                    GetConfigurations200ApplicationJSON1ScopesQueueScopesAdded = "read:deployment"
	GetConfigurations200ApplicationJSON1ScopesQueueScopesAddedReadWriteDeployment               GetConfigurations200ApplicationJSON1ScopesQueueScopesAdded = "read-write:deployment"
	GetConfigurations200ApplicationJSON1ScopesQueueScopesAddedReadWriteDeploymentCheck          GetConfigurations200ApplicationJSON1ScopesQueueScopesAdded = "read-write:deployment-check"
	GetConfigurations200ApplicationJSON1ScopesQueueScopesAddedReadProject                       GetConfigurations200ApplicationJSON1ScopesQueueScopesAdded = "read:project"
	GetConfigurations200ApplicationJSON1ScopesQueueScopesAddedReadWriteProject                  GetConfigurations200ApplicationJSON1ScopesQueueScopesAdded = "read-write:project"
	GetConfigurations200ApplicationJSON1ScopesQueueScopesAddedReadWriteProjectEnvVars           GetConfigurations200ApplicationJSON1ScopesQueueScopesAdded = "read-write:project-env-vars"
	GetConfigurations200ApplicationJSON1ScopesQueueScopesAddedReadWriteGlobalProjectEnvVars     GetConfigurations200ApplicationJSON1ScopesQueueScopesAdded = "read-write:global-project-env-vars"
	GetConfigurations200ApplicationJSON1ScopesQueueScopesAddedReadTeam                          GetConfigurations200ApplicationJSON1ScopesQueueScopesAdded = "read:team"
	GetConfigurations200ApplicationJSON1ScopesQueueScopesAddedReadUser                          GetConfigurations200ApplicationJSON1ScopesQueueScopesAdded = "read:user"
	GetConfigurations200ApplicationJSON1ScopesQueueScopesAddedReadWriteLogDrain                 GetConfigurations200ApplicationJSON1ScopesQueueScopesAdded = "read-write:log-drain"
	GetConfigurations200ApplicationJSON1ScopesQueueScopesAddedReadDomain                        GetConfigurations200ApplicationJSON1ScopesQueueScopesAdded = "read:domain"
	GetConfigurations200ApplicationJSON1ScopesQueueScopesAddedReadWriteDomain                   GetConfigurations200ApplicationJSON1ScopesQueueScopesAdded = "read-write:domain"
	GetConfigurations200ApplicationJSON1ScopesQueueScopesAddedReadWriteEdgeConfig               GetConfigurations200ApplicationJSON1ScopesQueueScopesAdded = "read-write:edge-config"
	GetConfigurations200ApplicationJSON1ScopesQueueScopesAddedReadWriteOtelEndpoint             GetConfigurations200ApplicationJSON1ScopesQueueScopesAdded = "read-write:otel-endpoint"
	GetConfigurations200ApplicationJSON1ScopesQueueScopesAddedReadMonitoring                    GetConfigurations200ApplicationJSON1ScopesQueueScopesAdded = "read:monitoring"
)

func (e GetConfigurations200ApplicationJSON1ScopesQueueScopesAdded) ToPointer() *GetConfigurations200ApplicationJSON1ScopesQueueScopesAdded {
	return &e
}

func (e *GetConfigurations200ApplicationJSON1ScopesQueueScopesAdded) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "read:integration-configuration":
		fallthrough
	case "read-write:integration-configuration":
		fallthrough
	case "read:deployment":
		fallthrough
	case "read-write:deployment":
		fallthrough
	case "read-write:deployment-check":
		fallthrough
	case "read:project":
		fallthrough
	case "read-write:project":
		fallthrough
	case "read-write:project-env-vars":
		fallthrough
	case "read-write:global-project-env-vars":
		fallthrough
	case "read:team":
		fallthrough
	case "read:user":
		fallthrough
	case "read-write:log-drain":
		fallthrough
	case "read:domain":
		fallthrough
	case "read-write:domain":
		fallthrough
	case "read-write:edge-config":
		fallthrough
	case "read-write:otel-endpoint":
		fallthrough
	case "read:monitoring":
		*e = GetConfigurations200ApplicationJSON1ScopesQueueScopesAdded(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetConfigurations200ApplicationJSON1ScopesQueueScopesAdded: %v", v)
	}
}

type GetConfigurations200ApplicationJSON1ScopesQueueScopesUpgraded string

const (
	GetConfigurations200ApplicationJSON1ScopesQueueScopesUpgradedReadIntegrationConfiguration      GetConfigurations200ApplicationJSON1ScopesQueueScopesUpgraded = "read:integration-configuration"
	GetConfigurations200ApplicationJSON1ScopesQueueScopesUpgradedReadWriteIntegrationConfiguration GetConfigurations200ApplicationJSON1ScopesQueueScopesUpgraded = "read-write:integration-configuration"
	GetConfigurations200ApplicationJSON1ScopesQueueScopesUpgradedReadDeployment                    GetConfigurations200ApplicationJSON1ScopesQueueScopesUpgraded = "read:deployment"
	GetConfigurations200ApplicationJSON1ScopesQueueScopesUpgradedReadWriteDeployment               GetConfigurations200ApplicationJSON1ScopesQueueScopesUpgraded = "read-write:deployment"
	GetConfigurations200ApplicationJSON1ScopesQueueScopesUpgradedReadWriteDeploymentCheck          GetConfigurations200ApplicationJSON1ScopesQueueScopesUpgraded = "read-write:deployment-check"
	GetConfigurations200ApplicationJSON1ScopesQueueScopesUpgradedReadProject                       GetConfigurations200ApplicationJSON1ScopesQueueScopesUpgraded = "read:project"
	GetConfigurations200ApplicationJSON1ScopesQueueScopesUpgradedReadWriteProject                  GetConfigurations200ApplicationJSON1ScopesQueueScopesUpgraded = "read-write:project"
	GetConfigurations200ApplicationJSON1ScopesQueueScopesUpgradedReadWriteProjectEnvVars           GetConfigurations200ApplicationJSON1ScopesQueueScopesUpgraded = "read-write:project-env-vars"
	GetConfigurations200ApplicationJSON1ScopesQueueScopesUpgradedReadWriteGlobalProjectEnvVars     GetConfigurations200ApplicationJSON1ScopesQueueScopesUpgraded = "read-write:global-project-env-vars"
	GetConfigurations200ApplicationJSON1ScopesQueueScopesUpgradedReadTeam                          GetConfigurations200ApplicationJSON1ScopesQueueScopesUpgraded = "read:team"
	GetConfigurations200ApplicationJSON1ScopesQueueScopesUpgradedReadUser                          GetConfigurations200ApplicationJSON1ScopesQueueScopesUpgraded = "read:user"
	GetConfigurations200ApplicationJSON1ScopesQueueScopesUpgradedReadWriteLogDrain                 GetConfigurations200ApplicationJSON1ScopesQueueScopesUpgraded = "read-write:log-drain"
	GetConfigurations200ApplicationJSON1ScopesQueueScopesUpgradedReadDomain                        GetConfigurations200ApplicationJSON1ScopesQueueScopesUpgraded = "read:domain"
	GetConfigurations200ApplicationJSON1ScopesQueueScopesUpgradedReadWriteDomain                   GetConfigurations200ApplicationJSON1ScopesQueueScopesUpgraded = "read-write:domain"
	GetConfigurations200ApplicationJSON1ScopesQueueScopesUpgradedReadWriteEdgeConfig               GetConfigurations200ApplicationJSON1ScopesQueueScopesUpgraded = "read-write:edge-config"
	GetConfigurations200ApplicationJSON1ScopesQueueScopesUpgradedReadWriteOtelEndpoint             GetConfigurations200ApplicationJSON1ScopesQueueScopesUpgraded = "read-write:otel-endpoint"
	GetConfigurations200ApplicationJSON1ScopesQueueScopesUpgradedReadMonitoring                    GetConfigurations200ApplicationJSON1ScopesQueueScopesUpgraded = "read:monitoring"
)

func (e GetConfigurations200ApplicationJSON1ScopesQueueScopesUpgraded) ToPointer() *GetConfigurations200ApplicationJSON1ScopesQueueScopesUpgraded {
	return &e
}

func (e *GetConfigurations200ApplicationJSON1ScopesQueueScopesUpgraded) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "read:integration-configuration":
		fallthrough
	case "read-write:integration-configuration":
		fallthrough
	case "read:deployment":
		fallthrough
	case "read-write:deployment":
		fallthrough
	case "read-write:deployment-check":
		fallthrough
	case "read:project":
		fallthrough
	case "read-write:project":
		fallthrough
	case "read-write:project-env-vars":
		fallthrough
	case "read-write:global-project-env-vars":
		fallthrough
	case "read:team":
		fallthrough
	case "read:user":
		fallthrough
	case "read-write:log-drain":
		fallthrough
	case "read:domain":
		fallthrough
	case "read-write:domain":
		fallthrough
	case "read-write:edge-config":
		fallthrough
	case "read-write:otel-endpoint":
		fallthrough
	case "read:monitoring":
		*e = GetConfigurations200ApplicationJSON1ScopesQueueScopesUpgraded(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetConfigurations200ApplicationJSON1ScopesQueueScopesUpgraded: %v", v)
	}
}

type GetConfigurations200ApplicationJSON1ScopesQueueScopes struct {
	Added    []GetConfigurations200ApplicationJSON1ScopesQueueScopesAdded    `json:"added"`
	Upgraded []GetConfigurations200ApplicationJSON1ScopesQueueScopesUpgraded `json:"upgraded"`
}

type GetConfigurations200ApplicationJSON1ScopesQueue struct {
	ConfirmedAt *int64                                                `json:"confirmedAt,omitempty"`
	Note        string                                                `json:"note"`
	RequestedAt int64                                                 `json:"requestedAt"`
	Scopes      GetConfigurations200ApplicationJSON1ScopesQueueScopes `json:"scopes"`
}

// GetConfigurations200ApplicationJSON1Source - Source defines where the configuration was installed from. It is used to analyze user engagement for integration installations in product metrics.
type GetConfigurations200ApplicationJSON1Source string

const (
	GetConfigurations200ApplicationJSON1SourceMarketplace  GetConfigurations200ApplicationJSON1Source = "marketplace"
	GetConfigurations200ApplicationJSON1SourceDeployButton GetConfigurations200ApplicationJSON1Source = "deploy-button"
	GetConfigurations200ApplicationJSON1SourceExternal     GetConfigurations200ApplicationJSON1Source = "external"
)

func (e GetConfigurations200ApplicationJSON1Source) ToPointer() *GetConfigurations200ApplicationJSON1Source {
	return &e
}

func (e *GetConfigurations200ApplicationJSON1Source) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "marketplace":
		fallthrough
	case "deploy-button":
		fallthrough
	case "external":
		*e = GetConfigurations200ApplicationJSON1Source(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetConfigurations200ApplicationJSON1Source: %v", v)
	}
}

type GetConfigurations200ApplicationJSON1Type string

const (
	GetConfigurations200ApplicationJSON1TypeIntegrationConfiguration GetConfigurations200ApplicationJSON1Type = "integration-configuration"
)

func (e GetConfigurations200ApplicationJSON1Type) ToPointer() *GetConfigurations200ApplicationJSON1Type {
	return &e
}

func (e *GetConfigurations200ApplicationJSON1Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "integration-configuration":
		*e = GetConfigurations200ApplicationJSON1Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetConfigurations200ApplicationJSON1Type: %v", v)
	}
}

type GetConfigurations200ApplicationJSON1 struct {
	// A timestamp that tells you when the configuration was installed successfully
	CompletedAt *int64 `json:"completedAt,omitempty"`
	// A timestamp that tells you when the configuration was created
	CreatedAt int64 `json:"createdAt"`
	// A timestamp that tells you when the configuration was updated.
	DeletedAt *int64 `json:"deletedAt,omitempty"`
	// A timestamp that tells you when the configuration was disabled. Note: Configurations can be disabled when the associated user loses access to a team. They do not function during this time until the configuration is 'transferred', meaning the associated user is changed to one with access to the team.
	DisabledAt     *int64                                              `json:"disabledAt,omitempty"`
	DisabledReason *GetConfigurations200ApplicationJSON1DisabledReason `json:"disabledReason,omitempty"`
	// The unique identifier of the configuration
	ID string `json:"id"`
	// The unique identifier of the app the configuration was created for
	IntegrationID string `json:"integrationId"`
	// The user or team ID that owns the configuration
	OwnerID string `json:"ownerId"`
	// When a configuration is limited to access certain projects, this will contain each of the project ID it is allowed to access. If it is not defined, the configuration has full access.
	Projects             []string `json:"projects,omitempty"`
	RemovedLogDrainsAt   *int64   `json:"removedLogDrainsAt,omitempty"`
	RemovedProjectEnvsAt *int64   `json:"removedProjectEnvsAt,omitempty"`
	RemovedTokensAt      *int64   `json:"removedTokensAt,omitempty"`
	RemovedWebhooksAt    *int64   `json:"removedWebhooksAt,omitempty"`
	// The resources that are allowed to be accessed by the configuration.
	Scopes      []string                                          `json:"scopes"`
	ScopesQueue []GetConfigurations200ApplicationJSON1ScopesQueue `json:"scopesQueue,omitempty"`
	// The slug of the integration the configuration is created for.
	Slug string `json:"slug"`
	// Source defines where the configuration was installed from. It is used to analyze user engagement for integration installations in product metrics.
	Source *GetConfigurations200ApplicationJSON1Source `json:"source,omitempty"`
	// When the configuration was created for a team, this will show the ID of the team.
	TeamID *string                                  `json:"teamId,omitempty"`
	Type   GetConfigurations200ApplicationJSON1Type `json:"type"`
	// A timestamp that tells you when the configuration was updated.
	UpdatedAt int64 `json:"updatedAt"`
	// The ID of the user that created the configuration.
	UserID string `json:"userId"`
}

type GetConfigurations200ApplicationJSONType string

const (
	GetConfigurations200ApplicationJSONTypeArrayOfgetConfigurations200ApplicationJSON1 GetConfigurations200ApplicationJSONType = "arrayOfgetConfigurations_200ApplicationJSON_1"
	GetConfigurations200ApplicationJSONTypeArrayOfgetConfigurations200ApplicationJSON2 GetConfigurations200ApplicationJSONType = "arrayOfgetConfigurations_200ApplicationJSON_2"
)

type GetConfigurations200ApplicationJSON struct {
	ArrayOfgetConfigurations200ApplicationJSON1 []GetConfigurations200ApplicationJSON1
	ArrayOfgetConfigurations200ApplicationJSON2 []GetConfigurations200ApplicationJSON2

	Type GetConfigurations200ApplicationJSONType
}

func CreateGetConfigurations200ApplicationJSONArrayOfgetConfigurations200ApplicationJSON1(arrayOfgetConfigurations200ApplicationJSON1 []GetConfigurations200ApplicationJSON1) GetConfigurations200ApplicationJSON {
	typ := GetConfigurations200ApplicationJSONTypeArrayOfgetConfigurations200ApplicationJSON1

	return GetConfigurations200ApplicationJSON{
		ArrayOfgetConfigurations200ApplicationJSON1: arrayOfgetConfigurations200ApplicationJSON1,
		Type: typ,
	}
}

func CreateGetConfigurations200ApplicationJSONArrayOfgetConfigurations200ApplicationJSON2(arrayOfgetConfigurations200ApplicationJSON2 []GetConfigurations200ApplicationJSON2) GetConfigurations200ApplicationJSON {
	typ := GetConfigurations200ApplicationJSONTypeArrayOfgetConfigurations200ApplicationJSON2

	return GetConfigurations200ApplicationJSON{
		ArrayOfgetConfigurations200ApplicationJSON2: arrayOfgetConfigurations200ApplicationJSON2,
		Type: typ,
	}
}

func (u *GetConfigurations200ApplicationJSON) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	arrayOfgetConfigurations200ApplicationJSON1 := []GetConfigurations200ApplicationJSON1{}
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&arrayOfgetConfigurations200ApplicationJSON1); err == nil {
		u.ArrayOfgetConfigurations200ApplicationJSON1 = arrayOfgetConfigurations200ApplicationJSON1
		u.Type = GetConfigurations200ApplicationJSONTypeArrayOfgetConfigurations200ApplicationJSON1
		return nil
	}

	arrayOfgetConfigurations200ApplicationJSON2 := []GetConfigurations200ApplicationJSON2{}
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&arrayOfgetConfigurations200ApplicationJSON2); err == nil {
		u.ArrayOfgetConfigurations200ApplicationJSON2 = arrayOfgetConfigurations200ApplicationJSON2
		u.Type = GetConfigurations200ApplicationJSONTypeArrayOfgetConfigurations200ApplicationJSON2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetConfigurations200ApplicationJSON) MarshalJSON() ([]byte, error) {
	if u.ArrayOfgetConfigurations200ApplicationJSON1 != nil {
		return json.Marshal(u.ArrayOfgetConfigurations200ApplicationJSON1)
	}

	if u.ArrayOfgetConfigurations200ApplicationJSON2 != nil {
		return json.Marshal(u.ArrayOfgetConfigurations200ApplicationJSON2)
	}

	return nil, nil
}

type GetConfigurationsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// The list of configurations for the authenticated user
	GetConfigurations200ApplicationJSONOneOf *GetConfigurations200ApplicationJSON
}
