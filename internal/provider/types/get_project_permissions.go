// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type GetProjectPermissions struct {
	AccessGroup                              []types.String `tfsdk:"access_group"`
	AliasGlobal                              []types.String `tfsdk:"alias_global"`
	AliasProject                             []types.String `tfsdk:"alias_project"`
	AliasProtectionBypass                    []types.String `tfsdk:"alias_protection_bypass"`
	Analytics                                []types.String `tfsdk:"analytics"`
	AnalyticsSampling                        []types.String `tfsdk:"analytics_sampling"`
	AnalyticsUsage                           []types.String `tfsdk:"analytics_usage"`
	AuditLog                                 []types.String `tfsdk:"audit_log"`
	BillingAddress                           []types.String `tfsdk:"billing_address"`
	BillingInformation                       []types.String `tfsdk:"billing_information"`
	BillingInvoice                           []types.String `tfsdk:"billing_invoice"`
	BillingInvoiceEmailRecipient             []types.String `tfsdk:"billing_invoice_email_recipient"`
	BillingInvoiceLanguage                   []types.String `tfsdk:"billing_invoice_language"`
	BillingPlan                              []types.String `tfsdk:"billing_plan"`
	BillingPurchaseOrder                     []types.String `tfsdk:"billing_purchase_order"`
	BillingTaxID                             []types.String `tfsdk:"billing_tax_id"`
	Blob                                     []types.String `tfsdk:"blob"`
	BlobStoreTokenSet                        []types.String `tfsdk:"blob_store_token_set"`
	Budget                                   []types.String `tfsdk:"budget"`
	CacheArtifact                            []types.String `tfsdk:"cache_artifact"`
	CacheArtifactUsageEvent                  []types.String `tfsdk:"cache_artifact_usage_event"`
	CodeChecks                               []types.String `tfsdk:"code_checks"`
	ConcurrentBuilds                         []types.String `tfsdk:"concurrent_builds"`
	Connect                                  []types.String `tfsdk:"connect"`
	ConnectConfiguration                     []types.String `tfsdk:"connect_configuration"`
	ConnectConfigurationLink                 []types.String `tfsdk:"connect_configuration_link"`
	DataCacheBillingSettings                 []types.String `tfsdk:"data_cache_billing_settings"`
	DataCacheNamespace                       []types.String `tfsdk:"data_cache_namespace"`
	Deployment                               []types.String `tfsdk:"deployment"`
	DeploymentCheck                          []types.String `tfsdk:"deployment_check"`
	DeploymentCheckPreview                   []types.String `tfsdk:"deployment_check_preview"`
	DeploymentCheckReRunFromProductionBranch []types.String `tfsdk:"deployment_check_re_run_from_production_branch"`
	DeploymentPreview                        []types.String `tfsdk:"deployment_preview"`
	DeploymentPrivate                        []types.String `tfsdk:"deployment_private"`
	DeploymentProductionGit                  []types.String `tfsdk:"deployment_production_git"`
	DeploymentPromote                        []types.String `tfsdk:"deployment_promote"`
	DeploymentRollback                       []types.String `tfsdk:"deployment_rollback"`
	Domain                                   []types.String `tfsdk:"domain"`
	DomainAcceptDelegation                   []types.String `tfsdk:"domain_accept_delegation"`
	DomainAuthCodes                          []types.String `tfsdk:"domain_auth_codes"`
	DomainCertificate                        []types.String `tfsdk:"domain_certificate"`
	DomainCheckConfig                        []types.String `tfsdk:"domain_check_config"`
	DomainMove                               []types.String `tfsdk:"domain_move"`
	DomainPurchase                           []types.String `tfsdk:"domain_purchase"`
	DomainRecord                             []types.String `tfsdk:"domain_record"`
	DomainTransferIn                         []types.String `tfsdk:"domain_transfer_in"`
	EdgeConfig                               []types.String `tfsdk:"edge_config"`
	EdgeConfigItem                           []types.String `tfsdk:"edge_config_item"`
	EdgeConfigSchema                         []types.String `tfsdk:"edge_config_schema"`
	EdgeConfigToken                          []types.String `tfsdk:"edge_config_token"`
	EndpointVerification                     []types.String `tfsdk:"endpoint_verification"`
	Environments                             []types.String `tfsdk:"environments"`
	Event                                    []types.String `tfsdk:"event"`
	FileUpload                               []types.String `tfsdk:"file_upload"`
	GitRepository                            []types.String `tfsdk:"git_repository"`
	Integration                              []types.String `tfsdk:"integration"`
	IntegrationConfiguration                 []types.String `tfsdk:"integration_configuration"`
	IntegrationConfigurationProjects         []types.String `tfsdk:"integration_configuration_projects"`
	IntegrationConfigurationTransfer         []types.String `tfsdk:"integration_configuration_transfer"`
	IntegrationEvent                         []types.String `tfsdk:"integration_event"`
	IntegrationResourceSecrets               []types.String `tfsdk:"integration_resource_secrets"`
	IntegrationStore                         []types.String `tfsdk:"integration_store"`
	IntegrationStoreTokenSet                 []types.String `tfsdk:"integration_store_token_set"`
	IntegrationVercelConfigurationOverride   []types.String `tfsdk:"integration_vercel_configuration_override"`
	IPBlocking                               []types.String `tfsdk:"ip_blocking"`
	Job                                      []types.String `tfsdk:"job"`
	JobGlobal                                []types.String `tfsdk:"job_global"`
	LogDrain                                 []types.String `tfsdk:"log_drain"`
	Logs                                     []types.String `tfsdk:"logs"`
	LogsPreset                               []types.String `tfsdk:"logs_preset"`
	MarketplaceBillingData                   []types.String `tfsdk:"marketplace_billing_data"`
	Monitoring                               []types.String `tfsdk:"monitoring"`
	MonitoringAlert                          []types.String `tfsdk:"monitoring_alert"`
	MonitoringChart                          []types.String `tfsdk:"monitoring_chart"`
	MonitoringQuery                          []types.String `tfsdk:"monitoring_query"`
	MonitoringSettings                       []types.String `tfsdk:"monitoring_settings"`
	NotificationCustomerBudget               []types.String `tfsdk:"notification_customer_budget"`
	NotificationDeploymentFailed             []types.String `tfsdk:"notification_deployment_failed"`
	NotificationDomainConfiguration          []types.String `tfsdk:"notification_domain_configuration"`
	NotificationDomainExpire                 []types.String `tfsdk:"notification_domain_expire"`
	NotificationDomainMoved                  []types.String `tfsdk:"notification_domain_moved"`
	NotificationDomainPurchase               []types.String `tfsdk:"notification_domain_purchase"`
	NotificationDomainRenewal                []types.String `tfsdk:"notification_domain_renewal"`
	NotificationDomainTransfer               []types.String `tfsdk:"notification_domain_transfer"`
	NotificationDomainUnverified             []types.String `tfsdk:"notification_domain_unverified"`
	NotificationMonitoringAlert              []types.String `tfsdk:"notification_monitoring_alert"`
	NotificationPaymentFailed                []types.String `tfsdk:"notification_payment_failed"`
	NotificationStatementOfReasons           []types.String `tfsdk:"notification_statement_of_reasons"`
	NotificationUsageAlert                   []types.String `tfsdk:"notification_usage_alert"`
	Oauth2Application                        []types.String `tfsdk:"oauth2_application"`
	Oauth2Connection                         []types.String `tfsdk:"oauth2_connection"`
	OpenTelemetryEndpoint                    []types.String `tfsdk:"open_telemetry_endpoint"`
	OptionsAllowlist                         []types.String `tfsdk:"options_allowlist"`
	OwnEvent                                 []types.String `tfsdk:"own_event"`
	PasswordProtection                       []types.String `tfsdk:"password_protection"`
	PasswordProtectionInvoiceItem            []types.String `tfsdk:"password_protection_invoice_item"`
	PaymentMethod                            []types.String `tfsdk:"payment_method"`
	Permissions                              []types.String `tfsdk:"permissions"`
	Postgres                                 []types.String `tfsdk:"postgres"`
	PostgresStoreTokenSet                    []types.String `tfsdk:"postgres_store_token_set"`
	PreviewDeploymentSuffix                  []types.String `tfsdk:"preview_deployment_suffix"`
	ProductionAliasProtectionBypass          []types.String `tfsdk:"production_alias_protection_bypass"`
	Project                                  []types.String `tfsdk:"project"`
	ProjectAccessGroup                       []types.String `tfsdk:"project_access_group"`
	ProjectAnalyticsSampling                 []types.String `tfsdk:"project_analytics_sampling"`
	ProjectAnalyticsUsage                    []types.String `tfsdk:"project_analytics_usage"`
	ProjectDeploymentExpiration              []types.String `tfsdk:"project_deployment_expiration"`
	ProjectDeploymentHook                    []types.String `tfsdk:"project_deployment_hook"`
	ProjectDomain                            []types.String `tfsdk:"project_domain"`
	ProjectDomainCheckConfig                 []types.String `tfsdk:"project_domain_check_config"`
	ProjectDomainMove                        []types.String `tfsdk:"project_domain_move"`
	ProjectEnvVars                           []types.String `tfsdk:"project_env_vars"`
	ProjectEnvVarsProduction                 []types.String `tfsdk:"project_env_vars_production"`
	ProjectEnvVarsUnownedByIntegration       []types.String `tfsdk:"project_env_vars_unowned_by_integration"`
	ProjectID                                []types.String `tfsdk:"project_id"`
	ProjectIntegrationConfiguration          []types.String `tfsdk:"project_integration_configuration"`
	ProjectLink                              []types.String `tfsdk:"project_link"`
	ProjectMember                            []types.String `tfsdk:"project_member"`
	ProjectMonitoring                        []types.String `tfsdk:"project_monitoring"`
	ProjectPermissions                       []types.String `tfsdk:"project_permissions"`
	ProjectProductionBranch                  []types.String `tfsdk:"project_production_branch"`
	ProjectProtectionBypass                  []types.String `tfsdk:"project_protection_bypass"`
	ProjectSupportCase                       []types.String `tfsdk:"project_support_case"`
	ProjectSupportCaseComment                []types.String `tfsdk:"project_support_case_comment"`
	ProjectTransfer                          []types.String `tfsdk:"project_transfer"`
	ProjectTransferIn                        []types.String `tfsdk:"project_transfer_in"`
	ProjectTransferOut                       []types.String `tfsdk:"project_transfer_out"`
	ProjectUsage                             []types.String `tfsdk:"project_usage"`
	ProTrialOnboarding                       []types.String `tfsdk:"pro_trial_onboarding"`
	RateLimit                                []types.String `tfsdk:"rate_limit"`
	Redis                                    []types.String `tfsdk:"redis"`
	RedisStoreTokenSet                       []types.String `tfsdk:"redis_store_token_set"`
	RemoteCaching                            []types.String `tfsdk:"remote_caching"`
	Repository                               []types.String `tfsdk:"repository"`
	SamlConfig                               []types.String `tfsdk:"saml_config"`
	SeawallConfig                            []types.String `tfsdk:"seawall_config"`
	Secret                                   []types.String `tfsdk:"secret"`
	SensitiveEnvironmentVariablePolicy       []types.String `tfsdk:"sensitive_environment_variable_policy"`
	SharedEnvVarConnection                   []types.String `tfsdk:"shared_env_var_connection"`
	SharedEnvVars                            []types.String `tfsdk:"shared_env_vars"`
	SharedEnvVarsProduction                  []types.String `tfsdk:"shared_env_vars_production"`
	SkewProtection                           []types.String `tfsdk:"skew_protection"`
	Space                                    []types.String `tfsdk:"space"`
	SpaceRun                                 []types.String `tfsdk:"space_run"`
	SupportCase                              []types.String `tfsdk:"support_case"`
	SupportCaseComment                       []types.String `tfsdk:"support_case_comment"`
	Team                                     []types.String `tfsdk:"team"`
	TeamAccessRequest                        []types.String `tfsdk:"team_access_request"`
	TeamFellowMembership                     []types.String `tfsdk:"team_fellow_membership"`
	TeamInvite                               []types.String `tfsdk:"team_invite"`
	TeamInviteCode                           []types.String `tfsdk:"team_invite_code"`
	TeamJoin                                 []types.String `tfsdk:"team_join"`
	TeamOwnMembership                        []types.String `tfsdk:"team_own_membership"`
	TeamOwnMembershipDisconnectSAML          []types.String `tfsdk:"team_own_membership_disconnect_saml"`
	Token                                    []types.String `tfsdk:"token"`
	TrustedIps                               []types.String `tfsdk:"trusted_ips"`
	Usage                                    []types.String `tfsdk:"usage"`
	UsageCycle                               []types.String `tfsdk:"usage_cycle"`
	User                                     []types.String `tfsdk:"user"`
	UserConnection                           []types.String `tfsdk:"user_connection"`
	WebAnalytics                             []types.String `tfsdk:"web_analytics"`
	WebAnalyticsPlan                         []types.String `tfsdk:"web_analytics_plan"`
	WebAuthn                                 []types.String `tfsdk:"web_authn"`
	Webhook                                  []types.String `tfsdk:"webhook"`
	WebhookEvent                             []types.String `tfsdk:"webhook_event"`
}
