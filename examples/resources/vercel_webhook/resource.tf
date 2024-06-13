resource "vercel_webhook" "my_webhook" {
  events = [
    "integration-configuration-removed",
  ]
  slug    = "...my_slug..."
  team_id = "...my_team_id..."
  url     = "http://well-to-do-base.net"
}