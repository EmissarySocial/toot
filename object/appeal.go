package object

// Appeal represents an account's appeal of a moderation action (AccountWarning).
// https://docs.joinmastodon.org/entities/Appeal/
type Appeal struct {
	Text  string `json:"text"`  // The text of the appeal from the moderated account to the moderators.
	State string `json:"state"` // [pending | approved | rejected]
}
