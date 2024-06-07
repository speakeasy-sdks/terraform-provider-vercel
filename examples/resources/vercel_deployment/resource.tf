resource "vercel_deployment" "my_deployment" {
  custom_environment_slug_or_id = "...my_custom_environment_slug_or_id..."
  deployment_id                 = "...my_deployment_id..."
  files = [
    {
      inlined_file = {
        data     = "...my_data..."
        encoding = "utf-8"
        file     = "folder/file.js"
      }
    },
  ]
  force_new = "1"
  git_metadata = {
    commit_author_name = "kyliau"
    commit_message     = "add method to measure Interaction to Next Paint (INP) (#36490)"
    commit_ref         = "main"
    commit_sha         = "dc36199b2234c6586ebe05ec94078a895c707e29"
    dirty              = true
    remote_url         = "https://github.com/vercel/next.js"
  }
  monorepo_manager = "...my_monorepo_manager..."
  name             = "my-instant-deployment"
  project          = "my-deployment-project"
  project_settings = {
    build_command                       = "...my_build_command..."
    command_for_ignoring_build_step     = "...my_command_for_ignoring_build_step..."
    dev_command                         = "...my_dev_command..."
    framework                           = "vitepress"
    install_command                     = "...my_install_command..."
    node_version                        = "18.x"
    output_directory                    = "...my_output_directory..."
    root_directory                      = "...my_root_directory..."
    serverless_function_region          = "...my_serverless_function_region..."
    skip_git_connect_during_link        = true
    source_files_outside_root_directory = true
  }
  skip_auto_detection_confirmation = "1"
  slug                             = "...my_slug..."
  target                           = "staging"
  team_id                          = "...my_team_id..."
  with_latest_commit               = true
}