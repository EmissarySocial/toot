package object

// CollectionItem represents an account's membership (and consent state) in a Collection.
// https://docs.joinmastodon.org/entities/CollectionItem/
type CollectionItem struct {
	ID        string `json:"id"`                   // The ID of the CollectionItem in the database.
	State     string `json:"state"`                // [pending | accepted | rejected | revoked]
	AccountID string `json:"account_id,omitempty"` // Only present when state is pending or accepted.
	CreatedAt string `json:"created_at"`           // (ISO 8601 Datetime)
}
