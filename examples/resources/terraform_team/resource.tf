resource "terraform_team" "my_team" {
  attribution = {
    landing_page                = "...my_landing_page..."
    page_before_conversion_page = "...my_page_before_conversion_page..."
    session_referrer            = "...my_session_referrer..."
    utm = {
      utm_campaign = "...my_utm_campaign..."
      utm_medium   = "...my_utm_medium..."
      utm_source   = "...my_utm_source..."
      utm_term     = "...my_utm_term..."
    }
  }
  name                = "A Random Team"
  new_default_team_id = "team_LLHUOMOoDlqOp8wPE4kFo9pE"
  reasons = [
    {
      slug        = "...my_slug..."
      description = "...my_description..."
    },
  ]
  slug = "a-random-team"
}