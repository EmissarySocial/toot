package object

// Collection represents a curated, named group of accounts (e.g. a featured-tag-driven collection).
// https://docs.joinmastodon.org/entities/Collection/
type Collection struct {
	ID           string           `json:"id"`           // The ID of the Collection in the database.
	URI          string           `json:"uri"`          // URI of the collection used for federation.
	Name         string           `json:"name"`         // The collection's display name.
	Description  string           `json:"description"`  // The collection's description.
	Language     string           `json:"language"`     // Primary language of the collection. (ISO 639 Part 1 two-letter code)
	AccountID    string           `json:"account_id"`   // The ID of the Account that owns this collection.
	Local        bool             `json:"local"`        // Is this collection local to this server?
	Sensitive    bool             `json:"sensitive"`    // Is this collection marked as sensitive?
	Discoverable bool             `json:"discoverable"` // Should this collection be shown in discovery surfaces?
	URL          string           `json:"url"`          // A link to the collection's HTML representation.
	ItemCount    int              `json:"item_count"`   // How many items are in this collection.
	Tag          ShallowTag       `json:"tag"`          // The hashtag associated with this collection.
	Items        []CollectionItem `json:"items"`        // The accounts (and their consent state) in this collection.
	CreatedAt    string           `json:"created_at"`   // (ISO 8601 Datetime)
	UpdatedAt    string           `json:"updated_at"`   // (ISO 8601 Datetime)
}
