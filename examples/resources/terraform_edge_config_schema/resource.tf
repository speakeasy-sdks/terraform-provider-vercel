resource "terraform_edge_config_schema" "my_edgeconfigschema" {
  definition     = "{ \"see\": \"documentation\" }"
  dry_run        = "...my_dry_run..."
  edge_config_id = "...my_edge_config_id..."
  slug           = "...my_slug..."
  team_id        = "...my_team_id..."
}