resource "vercel_dns" "my_dns" {
  domain  = "example.com"
  slug    = "...my_slug..."
  team_id = "...my_team_id..."
}