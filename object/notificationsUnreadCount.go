package object

// NotificationsUnreadCount represents the (possibly capped) count of a user's unread notifications.
// https://docs.joinmastodon.org/methods/notifications/#unread-count
type NotificationsUnreadCount struct {
	Count int `json:"count"`
}
