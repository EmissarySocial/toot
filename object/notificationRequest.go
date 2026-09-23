package object

// NotificationRequest represents a group of filtered notifications from a specific user.
// https://docs.joinmastodon.org/entities/NotificationRequest/
type NotificationRequest struct {
	ID                 string  `json:"id"`                    // The ID of the notification request in the database.
	CreatedAt          string  `json:"created_at"`            // (ISO 8601 Datetime)
	UpdatedAt          string  `json:"updated_at"`            // (ISO 8601 Datetime)
	Account            Account `json:"account"`               // The account whose filtered notifications were grouped.
	NotificationsCount string  `json:"notifications_count"`   // How many of this account's notifications were filtered.
	LastStatus         *Status `json:"last_status,omitempty"` // Most recent status associated with a filtered notification, if any.
}
