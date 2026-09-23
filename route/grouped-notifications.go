package route

/******************************************
* Grouped Notifications API Methods
* Receive grouped notifications for activity on your account or statuses.
* https://docs.joinmastodon.org/methods/grouped_notifications/
******************************************/

// GetGroupedNotifications is the path for GET /api/v2/notifications.
// https://docs.joinmastodon.org/methods/grouped_notifications/#get-grouped
const GetGroupedNotifications = "/api/v2/notifications"

// GetNotificationGroup is the path for GET /api/v2/notifications/:group_key.
// https://docs.joinmastodon.org/methods/grouped_notifications/#get-notification-group
const GetNotificationGroup = "/api/v2/notifications/:group_key"

// PostNotificationGroup_Dismiss is the path for POST /api/v2/notifications/:group_key/dismiss.
// https://docs.joinmastodon.org/methods/grouped_notifications/#dismiss-group
const PostNotificationGroup_Dismiss = "/api/v2/notifications/:group_key/dismiss"

// GetGroupedNotifications_UnreadCount is the path for GET /api/v2/notifications/unread_count.
// https://docs.joinmastodon.org/methods/grouped_notifications/#unread-group-count
const GetGroupedNotifications_UnreadCount = "/api/v2/notifications/unread_count"

// PostGroupedNotifications_Clear is the path for POST /api/v2/notifications/clear.
// https://docs.joinmastodon.org/methods/notifications/#clear
const PostGroupedNotifications_Clear = "/api/v2/notifications/clear"
