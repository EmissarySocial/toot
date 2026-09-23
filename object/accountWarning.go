package object

// AccountWarning represents a moderation action taken against an account and the
// message sent to that account.
// https://docs.joinmastodon.org/entities/AccountWarning/
type AccountWarning struct {
	ID            string   `json:"id"`
	Action        string   `json:"action"`     // [none | disable | mark_statuses_as_sensitive | delete_statuses | sensitive | silence | suspend]
	Text          string   `json:"text"`       // The message from the moderator to the target account.
	StatusIDs     []string `json:"status_ids"` // Statuses this warning applies to, if any.
	TargetAccount Account  `json:"target_account"`
	Appeal        *Appeal  `json:"appeal,omitempty"`
	CreatedAt     string   `json:"created_at"` // (ISO 8601 Datetime)
}
