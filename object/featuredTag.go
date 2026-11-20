package object

// FeaturedTag represents a hashtag that is featured on a profile.
// https://docs.joinmastodon.org/entities/FeaturedTag/
type FeaturedTag struct {
	ID            string `json:"id"`                       // The internal ID of the featured tag in the database.
	Name          string `json:"name"`                     // The name of the hashtag being featured.
	URL           string `json:"url"`                      // A link to all statuses by a user that contain this hashtag.
	StatusesCount string `json:"statuses_count"`           // The number of authored statuses containing this hashtag, as a string.
	LastStatusAt  string `json:"last_status_at,omitempty"` // The date of the last authored status containing this hashtag (ISO 8601), omitted if none.
}
