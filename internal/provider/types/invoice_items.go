// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

type InvoiceItems struct {
	Analytics                               *CreateTeamAnalytics `tfsdk:"analytics"`
	AnalyticsUsage                          *AnalyticsUsage      `tfsdk:"analytics_usage"`
	Artifacts                               *AnalyticsUsage      `tfsdk:"artifacts"`
	Bandwidth                               *AnalyticsUsage      `tfsdk:"bandwidth"`
	BlobStores                              *AnalyticsUsage      `tfsdk:"blob_stores"`
	BlobTotalAdvancedRequests               *AnalyticsUsage      `tfsdk:"blob_total_advanced_requests"`
	BlobTotalAvgSizeInBytes                 *AnalyticsUsage      `tfsdk:"blob_total_avg_size_in_bytes"`
	BlobTotalGetResponseObjectSizeInBytes   *AnalyticsUsage      `tfsdk:"blob_total_get_response_object_size_in_bytes"`
	BlobTotalSimpleRequests                 *AnalyticsUsage      `tfsdk:"blob_total_simple_requests"`
	BuildMinute                             *AnalyticsUsage      `tfsdk:"build_minute"`
	ConcurrentBuilds                        *CreateTeamAnalytics `tfsdk:"concurrent_builds"`
	DataCacheRead                           *AnalyticsUsage      `tfsdk:"data_cache_read"`
	DataCacheRevalidation                   *AnalyticsUsage      `tfsdk:"data_cache_revalidation"`
	DataCacheWrite                          *AnalyticsUsage      `tfsdk:"data_cache_write"`
	EdgeConfigRead                          *AnalyticsUsage      `tfsdk:"edge_config_read"`
	EdgeConfigWrite                         *AnalyticsUsage      `tfsdk:"edge_config_write"`
	EdgeFunctionExecutionUnits              *AnalyticsUsage      `tfsdk:"edge_function_execution_units"`
	EdgeMiddlewareInvocations               *AnalyticsUsage      `tfsdk:"edge_middleware_invocations"`
	EdgeRequest                             *AnalyticsUsage      `tfsdk:"edge_request"`
	EdgeRequestAdditionalCPUDuration        *AnalyticsUsage      `tfsdk:"edge_request_additional_cpu_duration"`
	Enterprise                              *CreateTeamAnalytics `tfsdk:"enterprise"`
	FastDataTransfer                        *AnalyticsUsage      `tfsdk:"fast_data_transfer"`
	FastOriginTransfer                      *AnalyticsUsage      `tfsdk:"fast_origin_transfer"`
	FunctionDuration                        *AnalyticsUsage      `tfsdk:"function_duration"`
	FunctionInvocation                      *AnalyticsUsage      `tfsdk:"function_invocation"`
	LogDrainsVolume                         *AnalyticsUsage      `tfsdk:"log_drains_volume"`
	Monitoring                              *CreateTeamAnalytics `tfsdk:"monitoring"`
	MonitoringMetric                        *AnalyticsUsage      `tfsdk:"monitoring_metric"`
	PasswordProtection                      *CreateTeamAnalytics `tfsdk:"password_protection"`
	PostgresComputeTime                     *AnalyticsUsage      `tfsdk:"postgres_compute_time"`
	PostgresDatabase                        *AnalyticsUsage      `tfsdk:"postgres_database"`
	PostgresDataStorage                     *AnalyticsUsage      `tfsdk:"postgres_data_storage"`
	PostgresDataTransfer                    *AnalyticsUsage      `tfsdk:"postgres_data_transfer"`
	PostgresWrittenData                     *AnalyticsUsage      `tfsdk:"postgres_written_data"`
	PreviewDeploymentSuffix                 *CreateTeamAnalytics `tfsdk:"preview_deployment_suffix"`
	Pro                                     *CreateTeamAnalytics `tfsdk:"pro"`
	Saml                                    *CreateTeamAnalytics `tfsdk:"saml"`
	ServerlessFunctionExecution             *AnalyticsUsage      `tfsdk:"serverless_function_execution"`
	SourceImages                            *AnalyticsUsage      `tfsdk:"source_images"`
	StorageRedisTotalBandwidthInBytes       *AnalyticsUsage      `tfsdk:"storage_redis_total_bandwidth_in_bytes"`
	StorageRedisTotalCommands               *AnalyticsUsage      `tfsdk:"storage_redis_total_commands"`
	StorageRedisTotalDailyAvgStorageInBytes *AnalyticsUsage      `tfsdk:"storage_redis_total_daily_avg_storage_in_bytes"`
	StorageRedisTotalDatabases              *AnalyticsUsage      `tfsdk:"storage_redis_total_databases"`
	TeamSeats                               *CreateTeamAnalytics `tfsdk:"team_seats"`
	WafOwaspExcessBytes                     *AnalyticsUsage      `tfsdk:"waf_owasp_excess_bytes"`
	WafOwaspRequests                        *AnalyticsUsage      `tfsdk:"waf_owasp_requests"`
	WebAnalytics                            *CreateTeamAnalytics `tfsdk:"web_analytics"`
	WebAnalyticsEvent                       *AnalyticsUsage      `tfsdk:"web_analytics_event"`
}
