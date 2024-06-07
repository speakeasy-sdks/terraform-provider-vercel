resource "terraform_edge_config" "my_edgeconfig" {
  items   = "{ \"see\": \"documentation\" }"
  slug    = "...my_slug..."
  team_id = "...my_team_id..."
}