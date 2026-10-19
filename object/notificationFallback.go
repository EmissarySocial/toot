package object

// NotificationFallback represents localized fallback text for a notification type
// that a client may not support.
// https://docs.joinmastodon.org/entities/NotificationFallback/
type NotificationFallback struct {
	Title   string `json:"title"`             // Localized fallback title (HTML).
	Summary string `json:"summary,omitempty"` // Localized fallback summary (HTML), nullable.
	Details string `json:"details,omitempty"` // Localized fallback details (HTML), nullable.
}
