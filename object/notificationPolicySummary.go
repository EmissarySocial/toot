package object

// NotificationPolicySummary summarizes the notifications currently filtered by a NotificationPolicy.
type NotificationPolicySummary struct {
	PendingRequestsCount      int `json:"pending_requests_count"`
	PendingNotificationsCount int `json:"pending_notifications_count"`
}
