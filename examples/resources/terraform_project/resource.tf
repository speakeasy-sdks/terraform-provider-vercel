resource "terraform_project" "my_project" {
  build_command                   = "...my_build_command..."
  command_for_ignoring_build_step = "...my_command_for_ignoring_build_step..."
  dev_command                     = "...my_dev_command..."
  environment_variables = [
    {
      git_branch = "...my_git_branch..."
      key        = "...my_key..."
      target = {
        array_of_two = [
          "development",
        ]
      }
      type  = "encrypted"
      value = "...my_value..."
    },
  ]
  framework = "gridsome"
  git_repository = {
    repo = "...my_repo..."
    type = "github"
  }
  id_or_name                               = "prj_12HKQaOmR5t5Uy6vdcQsNIiZgHGB"
  install_command                          = "...my_install_command..."
  name                                     = "a-project-name"
  output_directory                         = "...my_output_directory..."
  public_source                            = false
  root_directory                           = "...my_root_directory..."
  serverless_function_region               = "...my_serverless_function_region..."
  serverless_function_zero_config_failover = true
  skip_git_connect_during_link             = false
  slug                                     = "...my_slug..."
  team_id                                  = "...my_team_id..."
}