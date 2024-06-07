data "terraform_deployment" "my_deployment" {
  project_id         = "dpl_89qyp1cskzkLrVicDaZoDbjyHuDJ"
  slug               = "...my_slug..."
  team_id            = "...my_team_id..."
  with_git_repo_info = "true"
}