resource "vercel_webhook" "my_webhook" {
  events = [
    "integration-configuration-scope-change-confirmed",
  ]
  slug    = "...my_slug..."
  team_id = "...my_team_id..."
  url     = "https://thunderous-tobacco.info"
}