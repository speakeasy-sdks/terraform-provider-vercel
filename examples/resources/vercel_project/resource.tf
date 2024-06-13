resource "vercel_project" "my_project" {
  auto_assign_custom_domains            = true
  auto_assign_custom_domains_updated_by = "...my_auto_assign_custom_domains_updated_by..."
  auto_expose_system_envs               = true
  build_command                         = "...my_build_command..."
  command_for_ignoring_build_step       = "...my_command_for_ignoring_build_step..."
  customer_support_code_visibility      = true
  dev_command                           = "...my_dev_command..."
  directory_listing                     = true
  enable_preview_feedback               = true
  environment_variables = [
    {
      git_branch = "...my_git_branch..."
      key        = "...my_key..."
      target = [
        "preview",
      ]
      type  = "plain"
      value = "...my_value..."
    },
  ]
  framework           = "storybook"
  git_fork_protection = false
  git_lfs             = false
  git_repository = {
    repo = "...my_repo..."
    type = "github"
  }
  install_command                          = "...my_install_command..."
  name                                     = "a-project-name"
  node_version                             = "18.x"
  output_directory                         = "...my_output_directory..."
  public_source                            = true
  root_directory                           = "...my_root_directory..."
  serverless_function_region               = "...my_serverless_function_region..."
  serverless_function_zero_config_failover = false
  skew_protection_boundary_at              = 9
  skew_protection_max_age                  = 5
  skip_git_connect_during_link             = false
  slug                                     = "...my_slug..."
  source_files_outside_root_directory      = false
  team_id                                  = "...my_team_id..."
}