package object

// Role represents a set of permissions a user account can be granted.
// https://docs.joinmastodon.org/entities/Role/
type Role struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Permissions     string `json:"permissions"` // Bitmask, cast to a string.
	Color           string `json:"color"`       // Hex color code, or empty string.
	Highlighted     bool   `json:"highlighted"` // Whether this role should be highlighted on a user's profile.
	CollectionLimit int    `json:"collection_limit"`
}
