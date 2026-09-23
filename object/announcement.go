package object

// Announcement represents an announcement set by an administrator.
// https://docs.joinmastodon.org/entities/Announcement/
type Announcement struct {
	ID          string                `json:"id"`                  // The ID of the announcement in the database. (cast from an integer, but not guaranteed to be a number)
	Content     string                `json:"content"`             // The text of the announcement. (HTML)
	StartsAt    string                `json:"starts_at,omitempty"` // When the announcement will start. (ISO 8601 Datetime) or null
	EndsAt      string                `json:"ends_at,omitempty"`   // When the announcement will end. (ISO 8601 Datetime) or null
	Published   bool                  `json:"published"`           // Whether the announcement is currently active.
	AllDay      bool                  `json:"all_day"`             // Whether the announcement should start and end on dates only instead of datetimes. Will be false if there is no starts_at or ends_at time.
	PublishedAt string                `json:"published_at"`        // When the announcement was published. (ISO 8601 Datetime)
	UpdatedAt   string                `json:"updated_at"`          // When the announcement was last updated. (ISO 8601 Datetime)
	Read        bool                  `json:"read,omitempty"`      // Whether the announcement has been read by the current user.
	Mentions    []AnnouncementAccount `json:"mentions"`            // Accounts mentioned in the announcement text
	Statuses    []AnnouncementStatus  `json:"statuses"`            // Statuses linked in the announcement text.
	Tags        []StatusTag           `json:"tags"`                // Tags linked in the announcement text.
	Emojis      []CustomEmoji         `json:"emojis"`              // Custom emoji used in the announcement text.
	Reactions   []Reaction
}
