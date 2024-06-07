resource "vercel_dns" "my_dns" {
  domain    = "example.com"
  record_id = "rec_V0fra8eEgQwEpFhYG2vTzC3K"
  slug      = "...my_slug..."
  team_id   = "...my_team_id..."
}