package object

// ShallowTag is a minimal reference to a Tag, used inside Collection.
// https://docs.joinmastodon.org/entities/ShallowTag/
type ShallowTag struct {
	Name string `json:"name"` // The hashtag name, not including the leading #.
	URL  string `json:"url"`  // A link to the hashtag's page.
}
