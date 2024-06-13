---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vercel_project Data Source - terraform-provider-vercel"
subcategory: ""
description: |-
  Project DataSource
---

# vercel_project (Data Source)

Project DataSource

## Example Usage

```terraform
data "vercel_project" "my_project" {
  id      = "320f70a5-902f-4ce9-b659-240cd9bb1d5f"
  slug    = "...my_slug..."
  team_id = "...my_team_id..."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) The unique project identifier or the project name

### Optional

- `slug` (String) The Team slug to perform the request on behalf of.
- `team_id` (String) The Team identifier to perform the request on behalf of.

### Read-Only

- `account_id` (String)
- `analytics` (Attributes) (see [below for nested schema](#nestedatt--analytics))
- `auto_assign_custom_domains` (Boolean)
- `auto_assign_custom_domains_updated_by` (String)
- `auto_expose_system_envs` (Boolean)
- `build_command` (String)
- `command_for_ignoring_build_step` (String)
- `concurrency_bucket_name` (String)
- `connect_builds_enabled` (Boolean)
- `connect_configuration_id` (String)
- `created_at` (Number)
- `crons` (Attributes) (see [below for nested schema](#nestedatt--crons))
- `customer_support_code_visibility` (Boolean)
- `data_cache` (Attributes) (see [below for nested schema](#nestedatt--data_cache))
- `dev_command` (String)
- `directory_listing` (Boolean)
- `enable_preview_feedback` (Boolean)
- `env` (Attributes List) (see [below for nested schema](#nestedatt--env))
- `framework` (String) must be one of ["blitzjs", "nextjs", "gatsby", "remix", "astro", "hexo", "eleventy", "docusaurus-2", "docusaurus", "preact", "solidstart-1", "solidstart", "dojo", "ember", "vue", "scully", "ionic-angular", "angular", "polymer", "svelte", "sveltekit", "sveltekit-1", "ionic-react", "create-react-app", "gridsome", "umijs", "sapper", "saber", "stencil", "nuxtjs", "redwoodjs", "hugo", "jekyll", "brunch", "middleman", "zola", "hydrogen", "vite", "vitepress", "vuepress", "parcel", "sanity", "storybook"]
- `git_comments` (Attributes) (see [below for nested schema](#nestedatt--git_comments))
- `git_fork_protection` (Boolean)
- `git_lfs` (Boolean)
- `has_active_branches` (Boolean)
- `has_floating_aliases` (Boolean)
- `install_command` (String)
- `last_alias_request` (Attributes) (see [below for nested schema](#nestedatt--last_alias_request))
- `last_rollback_target` (Attributes) (see [below for nested schema](#nestedatt--last_rollback_target))
- `latest_deployments` (Attributes List) (see [below for nested schema](#nestedatt--latest_deployments))
- `link` (Attributes) (see [below for nested schema](#nestedatt--link))
- `live` (Boolean)
- `name` (String)
- `node_version` (String) must be one of ["20.x", "18.x", "16.x", "14.x", "12.x", "10.x", "8.10.x"]
- `oidc_token_config` (Attributes) (see [below for nested schema](#nestedatt--oidc_token_config))
- `options_allowlist` (Attributes) (see [below for nested schema](#nestedatt--options_allowlist))
- `output_directory` (String)
- `passive_connect_configuration_id` (String)
- `password_protection` (Attributes) (see [below for nested schema](#nestedatt--password_protection))
- `paused` (Boolean)
- `permissions` (Attributes) (see [below for nested schema](#nestedatt--permissions))
- `production_deployments_fast_lane` (Boolean)
- `protection_bypass` (Attributes Map) (see [below for nested schema](#nestedatt--protection_bypass))
- `public_source` (Boolean)
- `root_directory` (String)
- `security` (Attributes) (see [below for nested schema](#nestedatt--security))
- `serverless_function_region` (String)
- `serverless_function_zero_config_failover` (Boolean)
- `skew_protection_boundary_at` (Number)
- `skew_protection_max_age` (Number)
- `skip_git_connect_during_link` (Boolean)
- `source_files_outside_root_directory` (Boolean)
- `speed_insights` (Attributes) (see [below for nested schema](#nestedatt--speed_insights))
- `sso_protection` (Attributes) (see [below for nested schema](#nestedatt--sso_protection))
- `targets` (Attributes Map) (see [below for nested schema](#nestedatt--targets))
- `transfer_completed_at` (Number)
- `transfer_started_at` (Number)
- `transfer_to_account_id` (String)
- `transferred_from_account_id` (String)
- `trusted_ips` (Attributes) (see [below for nested schema](#nestedatt--trusted_ips))
- `updated_at` (Number)
- `web_analytics` (Attributes) (see [below for nested schema](#nestedatt--web_analytics))

<a id="nestedatt--analytics"></a>
### Nested Schema for `analytics`

Read-Only:

- `canceled_at` (Number)
- `disabled_at` (Number)
- `enabled_at` (Number)
- `id` (String)
- `paid_at` (Number)
- `sample_rate_percent` (Number)
- `spend_limit_in_dollars` (Number)


<a id="nestedatt--crons"></a>
### Nested Schema for `crons`

Read-Only:

- `definitions` (Attributes List) (see [below for nested schema](#nestedatt--crons--definitions))
- `deployment_id` (String) The ID of the Deployment from which the definitions originated.
- `disabled_at` (Number) The time the feature was disabled for this project.
- `enabled_at` (Number) The time the feature was enabled for this project. Note: It enables automatically with the first Deployment that outputs cronjobs.
- `updated_at` (Number)

<a id="nestedatt--crons--definitions"></a>
### Nested Schema for `crons.definitions`

Read-Only:

- `host` (String) The hostname that should be used.
- `path` (String) The path that should be called for the cronjob.
- `schedule` (String) The cron expression.



<a id="nestedatt--data_cache"></a>
### Nested Schema for `data_cache`

Read-Only:

- `storage_size_bytes` (Number)
- `unlimited` (Boolean)
- `user_disabled` (Boolean)


<a id="nestedatt--env"></a>
### Nested Schema for `env`

Read-Only:

- `comment` (String)
- `configuration_id` (String)
- `content_hint` (Attributes) (see [below for nested schema](#nestedatt--env--content_hint))
- `created_at` (Number)
- `created_by` (String)
- `custom_environment_id` (String)
- `decrypted` (Boolean) Whether `value` is decrypted.
- `edge_config_id` (String)
- `edge_config_token_id` (String)
- `git_branch` (String)
- `id` (String)
- `internal_content_hint` (Attributes) Similar to `contentHints`, but should not be exposed to the user. (see [below for nested schema](#nestedatt--env--internal_content_hint))
- `key` (String)
- `target` (List of String)
- `type` (String) must be one of ["secret", "system", "encrypted", "plain", "sensitive"]
- `updated_at` (Number)
- `updated_by` (String)
- `value` (String)

<a id="nestedatt--env--content_hint"></a>
### Nested Schema for `env.content_hint`

Read-Only:

- `eight` (Attributes) (see [below for nested schema](#nestedatt--env--content_hint--eight))
- `eleven` (Attributes) (see [below for nested schema](#nestedatt--env--content_hint--eleven))
- `five` (Attributes) (see [below for nested schema](#nestedatt--env--content_hint--five))
- `four` (Attributes) (see [below for nested schema](#nestedatt--env--content_hint--four))
- `fourteen` (Attributes) (see [below for nested schema](#nestedatt--env--content_hint--fourteen))
- `nine` (Attributes) (see [below for nested schema](#nestedatt--env--content_hint--nine))
- `one` (Attributes) (see [below for nested schema](#nestedatt--env--content_hint--one))
- `seven` (Attributes) (see [below for nested schema](#nestedatt--env--content_hint--seven))
- `six` (Attributes) (see [below for nested schema](#nestedatt--env--content_hint--six))
- `ten` (Attributes) (see [below for nested schema](#nestedatt--env--content_hint--ten))
- `thirteen` (Attributes) (see [below for nested schema](#nestedatt--env--content_hint--thirteen))
- `three` (Attributes) (see [below for nested schema](#nestedatt--env--content_hint--three))
- `twelve` (Attributes) (see [below for nested schema](#nestedatt--env--content_hint--twelve))
- `two` (Attributes) (see [below for nested schema](#nestedatt--env--content_hint--two))

<a id="nestedatt--env--content_hint--eight"></a>
### Nested Schema for `env.content_hint.eight`

Read-Only:

- `store_id` (String)
- `type` (String) must be one of ["postgres-prisma-url"]


<a id="nestedatt--env--content_hint--eleven"></a>
### Nested Schema for `env.content_hint.eleven`

Read-Only:

- `store_id` (String)
- `type` (String) must be one of ["postgres-password"]


<a id="nestedatt--env--content_hint--five"></a>
### Nested Schema for `env.content_hint.five`

Read-Only:

- `store_id` (String)
- `type` (String) must be one of ["blob-read-write-token"]


<a id="nestedatt--env--content_hint--four"></a>
### Nested Schema for `env.content_hint.four`

Read-Only:

- `store_id` (String)
- `type` (String) must be one of ["redis-rest-api-read-only-token"]


<a id="nestedatt--env--content_hint--fourteen"></a>
### Nested Schema for `env.content_hint.fourteen`

Read-Only:

- `store_id` (String)
- `type` (String) must be one of ["integration-store-secret"]


<a id="nestedatt--env--content_hint--nine"></a>
### Nested Schema for `env.content_hint.nine`

Read-Only:

- `store_id` (String)
- `type` (String) must be one of ["postgres-user"]


<a id="nestedatt--env--content_hint--one"></a>
### Nested Schema for `env.content_hint.one`

Read-Only:

- `store_id` (String)
- `type` (String) must be one of ["redis-url"]


<a id="nestedatt--env--content_hint--seven"></a>
### Nested Schema for `env.content_hint.seven`

Read-Only:

- `store_id` (String)
- `type` (String) must be one of ["postgres-url-non-pooling"]


<a id="nestedatt--env--content_hint--six"></a>
### Nested Schema for `env.content_hint.six`

Read-Only:

- `store_id` (String)
- `type` (String) must be one of ["postgres-url"]


<a id="nestedatt--env--content_hint--ten"></a>
### Nested Schema for `env.content_hint.ten`

Read-Only:

- `store_id` (String)
- `type` (String) must be one of ["postgres-host"]


<a id="nestedatt--env--content_hint--thirteen"></a>
### Nested Schema for `env.content_hint.thirteen`

Read-Only:

- `store_id` (String)
- `type` (String) must be one of ["postgres-url-no-ssl"]


<a id="nestedatt--env--content_hint--three"></a>
### Nested Schema for `env.content_hint.three`

Read-Only:

- `store_id` (String)
- `type` (String) must be one of ["redis-rest-api-token"]


<a id="nestedatt--env--content_hint--twelve"></a>
### Nested Schema for `env.content_hint.twelve`

Read-Only:

- `store_id` (String)
- `type` (String) must be one of ["postgres-database"]


<a id="nestedatt--env--content_hint--two"></a>
### Nested Schema for `env.content_hint.two`

Read-Only:

- `store_id` (String)
- `type` (String) must be one of ["redis-rest-api-url"]



<a id="nestedatt--env--internal_content_hint"></a>
### Nested Schema for `env.internal_content_hint`

Read-Only:

- `encrypted_value` (String) Contains the `value` of the env variable, encrypted with a special key to make decryption possible in the subscriber Lambda.
- `type` (String) must be one of ["flags-secret"]



<a id="nestedatt--git_comments"></a>
### Nested Schema for `git_comments`

Read-Only:

- `on_commit` (Boolean) Whether the Vercel bot should comment on commits
- `on_pull_request` (Boolean) Whether the Vercel bot should comment on PRs


<a id="nestedatt--last_alias_request"></a>
### Nested Schema for `last_alias_request`

Read-Only:

- `from_deployment_id` (String)
- `job_status` (String) must be one of ["succeeded", "failed", "skipped", "pending", "in-progress"]
- `requested_at` (Number)
- `to_deployment_id` (String)
- `type` (String) must be one of ["promote", "rollback"]


<a id="nestedatt--last_rollback_target"></a>
### Nested Schema for `last_rollback_target`


<a id="nestedatt--latest_deployments"></a>
### Nested Schema for `latest_deployments`

Read-Only:

- `alias` (List of String)
- `alias_assigned` (Attributes) (see [below for nested schema](#nestedatt--latest_deployments--alias_assigned))
- `alias_error` (Attributes) (see [below for nested schema](#nestedatt--latest_deployments--alias_error))
- `alias_final` (String)
- `automatic_aliases` (List of String)
- `building_at` (Number)
- `builds` (Attributes List) (see [below for nested schema](#nestedatt--latest_deployments--builds))
- `checks_conclusion` (String) must be one of ["succeeded", "failed", "skipped", "canceled"]
- `checks_state` (String) must be one of ["registered", "running", "completed"]
- `connect_builds_enabled` (Boolean)
- `connect_configuration_id` (String)
- `created_at` (Number)
- `created_in` (String)
- `creator` (Attributes) (see [below for nested schema](#nestedatt--latest_deployments--creator))
- `deployment_hostname` (String)
- `forced` (Boolean)
- `id` (String)
- `meta` (Map of String)
- `monorepo_manager` (String)
- `name` (String)
- `oidc_token_claims` (Attributes Map) (see [below for nested schema](#nestedatt--latest_deployments--oidc_token_claims))
- `plan` (String) must be one of ["pro", "enterprise", "hobby"]
- `preview_comments_enabled` (Boolean) Whether or not preview comments are enabled for the deployment
- `private` (Boolean)
- `ready_at` (Number)
- `ready_state` (String) must be one of ["BUILDING", "ERROR", "INITIALIZING", "QUEUED", "READY", "CANCELED"]
- `ready_substate` (String) must be one of ["STAGED", "PROMOTED"]
- `requested_at` (Number)
- `target` (String)
- `team_id` (String)
- `type` (String) must be one of ["LAMBDAS"]
- `url` (String)
- `user_id` (String)
- `with_cache` (Boolean)

<a id="nestedatt--latest_deployments--alias_assigned"></a>
### Nested Schema for `latest_deployments.alias_assigned`

Read-Only:

- `boolean` (Boolean)
- `number` (Number)


<a id="nestedatt--latest_deployments--alias_error"></a>
### Nested Schema for `latest_deployments.alias_error`

Read-Only:

- `code` (String)
- `message` (String)


<a id="nestedatt--latest_deployments--builds"></a>
### Nested Schema for `latest_deployments.builds`

Read-Only:

- `dest` (String)
- `src` (String)
- `use` (String)


<a id="nestedatt--latest_deployments--creator"></a>
### Nested Schema for `latest_deployments.creator`

Read-Only:

- `email` (String)
- `github_login` (String)
- `gitlab_login` (String)
- `uid` (String)
- `username` (String)


<a id="nestedatt--latest_deployments--oidc_token_claims"></a>
### Nested Schema for `latest_deployments.oidc_token_claims`

Read-Only:

- `array_of_str` (List of String)
- `str` (String)



<a id="nestedatt--link"></a>
### Nested Schema for `link`

Read-Only:

- `one` (Attributes) (see [below for nested schema](#nestedatt--link--one))
- `three` (Attributes) (see [below for nested schema](#nestedatt--link--three))
- `two` (Attributes) (see [below for nested schema](#nestedatt--link--two))

<a id="nestedatt--link--one"></a>
### Nested Schema for `link.one`

Read-Only:

- `created_at` (Number)
- `deploy_hooks` (Attributes List) (see [below for nested schema](#nestedatt--link--one--deploy_hooks))
- `git_credential_id` (String)
- `org` (String)
- `production_branch` (String)
- `repo` (String)
- `repo_id` (Number)
- `sourceless` (Boolean)
- `type` (String) must be one of ["github"]
- `updated_at` (Number)

<a id="nestedatt--link--one--deploy_hooks"></a>
### Nested Schema for `link.one.deploy_hooks`

Read-Only:

- `created_at` (Number)
- `id` (String)
- `name` (String)
- `ref` (String)
- `url` (String)



<a id="nestedatt--link--three"></a>
### Nested Schema for `link.three`

Read-Only:

- `created_at` (Number)
- `deploy_hooks` (Attributes List) (see [below for nested schema](#nestedatt--link--three--deploy_hooks))
- `git_credential_id` (String)
- `name` (String)
- `owner` (String)
- `production_branch` (String)
- `slug` (String)
- `sourceless` (Boolean)
- `type` (String) must be one of ["bitbucket"]
- `updated_at` (Number)
- `uuid` (String)
- `workspace_uuid` (String)

<a id="nestedatt--link--three--deploy_hooks"></a>
### Nested Schema for `link.three.deploy_hooks`

Read-Only:

- `created_at` (Number)
- `id` (String)
- `name` (String)
- `ref` (String)
- `url` (String)



<a id="nestedatt--link--two"></a>
### Nested Schema for `link.two`

Read-Only:

- `created_at` (Number)
- `deploy_hooks` (Attributes List) (see [below for nested schema](#nestedatt--link--two--deploy_hooks))
- `git_credential_id` (String)
- `production_branch` (String)
- `project_id` (String)
- `project_name` (String)
- `project_name_with_namespace` (String)
- `project_namespace` (String)
- `project_url` (String)
- `sourceless` (Boolean)
- `type` (String) must be one of ["gitlab"]
- `updated_at` (Number)

<a id="nestedatt--link--two--deploy_hooks"></a>
### Nested Schema for `link.two.deploy_hooks`

Read-Only:

- `created_at` (Number)
- `id` (String)
- `name` (String)
- `ref` (String)
- `url` (String)




<a id="nestedatt--oidc_token_config"></a>
### Nested Schema for `oidc_token_config`

Read-Only:

- `enabled` (Boolean)


<a id="nestedatt--options_allowlist"></a>
### Nested Schema for `options_allowlist`

Read-Only:

- `paths` (Attributes List) (see [below for nested schema](#nestedatt--options_allowlist--paths))

<a id="nestedatt--options_allowlist--paths"></a>
### Nested Schema for `options_allowlist.paths`

Read-Only:

- `value` (String)



<a id="nestedatt--password_protection"></a>
### Nested Schema for `password_protection`


<a id="nestedatt--permissions"></a>
### Nested Schema for `permissions`

Read-Only:

- `access_group` (List of String)
- `alias_global` (List of String)
- `alias_project` (List of String)
- `alias_protection_bypass` (List of String)
- `analytics` (List of String)
- `analytics_sampling` (List of String)
- `analytics_usage` (List of String)
- `audit_log` (List of String)
- `billing_address` (List of String)
- `billing_information` (List of String)
- `billing_invoice` (List of String)
- `billing_invoice_email_recipient` (List of String)
- `billing_invoice_language` (List of String)
- `billing_plan` (List of String)
- `billing_purchase_order` (List of String)
- `billing_tax_id` (List of String)
- `blob` (List of String)
- `blob_store_token_set` (List of String)
- `budget` (List of String)
- `cache_artifact` (List of String)
- `cache_artifact_usage_event` (List of String)
- `code_checks` (List of String)
- `concurrent_builds` (List of String)
- `connect` (List of String)
- `connect_configuration` (List of String)
- `connect_configuration_link` (List of String)
- `data_cache_billing_settings` (List of String)
- `data_cache_namespace` (List of String)
- `deployment` (List of String)
- `deployment_check` (List of String)
- `deployment_check_preview` (List of String)
- `deployment_check_re_run_from_production_branch` (List of String)
- `deployment_preview` (List of String)
- `deployment_private` (List of String)
- `deployment_production_git` (List of String)
- `deployment_promote` (List of String)
- `deployment_rollback` (List of String)
- `domain` (List of String)
- `domain_accept_delegation` (List of String)
- `domain_auth_codes` (List of String)
- `domain_certificate` (List of String)
- `domain_check_config` (List of String)
- `domain_move` (List of String)
- `domain_purchase` (List of String)
- `domain_record` (List of String)
- `domain_transfer_in` (List of String)
- `edge_config` (List of String)
- `edge_config_item` (List of String)
- `edge_config_schema` (List of String)
- `edge_config_token` (List of String)
- `endpoint_verification` (List of String)
- `environments` (List of String)
- `event` (List of String)
- `file_upload` (List of String)
- `git_repository` (List of String)
- `integration` (List of String)
- `integration_configuration` (List of String)
- `integration_configuration_projects` (List of String)
- `integration_configuration_transfer` (List of String)
- `integration_event` (List of String)
- `integration_resource_secrets` (List of String)
- `integration_store` (List of String)
- `integration_store_token_set` (List of String)
- `integration_vercel_configuration_override` (List of String)
- `ip_blocking` (List of String)
- `job` (List of String)
- `job_global` (List of String)
- `log_drain` (List of String)
- `logs` (List of String)
- `logs_preset` (List of String)
- `marketplace_billing_data` (List of String)
- `monitoring` (List of String)
- `monitoring_alert` (List of String)
- `monitoring_chart` (List of String)
- `monitoring_query` (List of String)
- `monitoring_settings` (List of String)
- `notification_customer_budget` (List of String)
- `notification_deployment_failed` (List of String)
- `notification_domain_configuration` (List of String)
- `notification_domain_expire` (List of String)
- `notification_domain_moved` (List of String)
- `notification_domain_purchase` (List of String)
- `notification_domain_renewal` (List of String)
- `notification_domain_transfer` (List of String)
- `notification_domain_unverified` (List of String)
- `notification_monitoring_alert` (List of String)
- `notification_payment_failed` (List of String)
- `notification_statement_of_reasons` (List of String)
- `notification_usage_alert` (List of String)
- `oauth2_application` (List of String)
- `oauth2_connection` (List of String)
- `open_telemetry_endpoint` (List of String)
- `options_allowlist` (List of String)
- `own_event` (List of String)
- `password_protection` (List of String)
- `password_protection_invoice_item` (List of String)
- `payment_method` (List of String)
- `permissions` (List of String)
- `postgres` (List of String)
- `postgres_store_token_set` (List of String)
- `preview_deployment_suffix` (List of String)
- `pro_trial_onboarding` (List of String)
- `production_alias_protection_bypass` (List of String)
- `project` (List of String)
- `project_access_group` (List of String)
- `project_analytics_sampling` (List of String)
- `project_analytics_usage` (List of String)
- `project_deployment_expiration` (List of String)
- `project_deployment_hook` (List of String)
- `project_domain` (List of String)
- `project_domain_check_config` (List of String)
- `project_domain_move` (List of String)
- `project_env_vars` (List of String)
- `project_env_vars_production` (List of String)
- `project_env_vars_unowned_by_integration` (List of String)
- `project_id` (List of String)
- `project_integration_configuration` (List of String)
- `project_link` (List of String)
- `project_member` (List of String)
- `project_monitoring` (List of String)
- `project_permissions` (List of String)
- `project_production_branch` (List of String)
- `project_protection_bypass` (List of String)
- `project_support_case` (List of String)
- `project_support_case_comment` (List of String)
- `project_transfer` (List of String)
- `project_transfer_in` (List of String)
- `project_transfer_out` (List of String)
- `project_usage` (List of String)
- `rate_limit` (List of String)
- `redis` (List of String)
- `redis_store_token_set` (List of String)
- `remote_caching` (List of String)
- `repository` (List of String)
- `saml_config` (List of String)
- `seawall_config` (List of String)
- `secret` (List of String)
- `sensitive_environment_variable_policy` (List of String)
- `shared_env_var_connection` (List of String)
- `shared_env_vars` (List of String)
- `shared_env_vars_production` (List of String)
- `skew_protection` (List of String)
- `space` (List of String)
- `space_run` (List of String)
- `support_case` (List of String)
- `support_case_comment` (List of String)
- `team` (List of String)
- `team_access_request` (List of String)
- `team_fellow_membership` (List of String)
- `team_invite` (List of String)
- `team_invite_code` (List of String)
- `team_join` (List of String)
- `team_own_membership` (List of String)
- `team_own_membership_disconnect_saml` (List of String)
- `token` (List of String)
- `trusted_ips` (List of String)
- `usage` (List of String)
- `usage_cycle` (List of String)
- `user` (List of String)
- `user_connection` (List of String)
- `web_analytics` (List of String)
- `web_analytics_plan` (List of String)
- `web_authn` (List of String)
- `webhook` (List of String)
- `webhook_event` (List of String)


<a id="nestedatt--protection_bypass"></a>
### Nested Schema for `protection_bypass`

Read-Only:

- `created_at` (Number)
- `created_by` (String)
- `scope` (String) must be one of ["automation-bypass"]


<a id="nestedatt--security"></a>
### Nested Schema for `security`

Read-Only:

- `attack_mode_active_until` (Number)
- `attack_mode_enabled` (Boolean)
- `attack_mode_updated_at` (Number)
- `firewall_config_version` (Number)
- `firewall_enabled` (Boolean)
- `firewall_seawall_enabled` (Boolean)
- `firewall_updated_at` (Number)
- `ja3_enabled` (Boolean)
- `ja4_enabled` (Boolean)


<a id="nestedatt--speed_insights"></a>
### Nested Schema for `speed_insights`

Read-Only:

- `canceled_at` (Number)
- `disabled_at` (Number)
- `enabled_at` (Number)
- `has_data` (Boolean)
- `id` (String)
- `paid_at` (Number)


<a id="nestedatt--sso_protection"></a>
### Nested Schema for `sso_protection`

Read-Only:

- `deployment_type` (String) must be one of ["all", "preview", "prod_deployment_urls_and_all_previews"]


<a id="nestedatt--targets"></a>
### Nested Schema for `targets`

Read-Only:

- `array_of_str` (List of String)
- `str` (String)


<a id="nestedatt--trusted_ips"></a>
### Nested Schema for `trusted_ips`

Read-Only:

- `addresses` (Attributes List) (see [below for nested schema](#nestedatt--trusted_ips--addresses))
- `deployment_type` (String) must be one of ["all", "preview", "prod_deployment_urls_and_all_previews", "production"]
- `protection_mode` (String) must be one of ["additional", "exclusive"]

<a id="nestedatt--trusted_ips--addresses"></a>
### Nested Schema for `trusted_ips.addresses`

Read-Only:

- `note` (String)
- `value` (String)



<a id="nestedatt--web_analytics"></a>
### Nested Schema for `web_analytics`

Read-Only:

- `canceled_at` (Number)
- `disabled_at` (Number)
- `enabled_at` (Number)
- `has_data` (Boolean)
- `id` (String)

