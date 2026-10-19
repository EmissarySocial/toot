package object

// Preferences represents a user's preferences.
// https://docs.joinmastodon.org/entities/Preferences/
type Preferences struct {
	PostingDefaultVisibility  string `json:"posting:default:visibility"` // [public | unlisted | private | direct]
	PostingDefaultSensitive   bool   `json:"posting:default:sensitive"`
	PostingDefaultLanguage    string `json:"posting:default:language,omitempty"` // ISO 639-1 two-letter code, or empty/null
	PostingDefaultQuotePolicy string `json:"posting:default:quote_policy"`       // [public | followers | nobody] - undocumented, verified against real server code
	ReadingExpandMedia        string `json:"reading:expand:media"`               // [default | show_all | hide_all]
	ReadingExpandSpoilers     bool   `json:"reading:expand:spoilers"`
	ReadingAutoplayGifs       bool   `json:"reading:autoplay:gifs"` // undocumented, verified against real server code
}
