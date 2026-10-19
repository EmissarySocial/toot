package object

// PreviewCardAuthor credits an author of the content referenced by a PreviewCard.
// https://docs.joinmastodon.org/entities/PreviewCardAuthor/
type PreviewCardAuthor struct {
	Name    string   `json:"name"`
	URL     string   `json:"url"`
	Account *Account `json:"account,omitempty"`
}
