package scope

/******************************************
* Grouped Notifications API Methods
* Receive grouped notifications for activity on your account or statuses.
* https://docs.joinmastodon.org/methods/grouped_notifications/
******************************************/

// GetGroupedNotifications is the OAuth scope required by GET /api/v2/notifications.
// https://docs.joinmastodon.org/methods/grouped_notifications/#get-grouped
const GetGroupedNotifications = ReadNotifications

// GetNotificationGroup is the OAuth scope required by GET /api/v2/notifications/:group_key.
// https://docs.joinmastodon.org/methods/grouped_notifications/#get-notification-group
const GetNotificationGroup = ReadNotifications

// PostNotificationGroup_Dismiss is the OAuth scope required by POST /api/v2/notifications/:group_key/dismiss.
// https://docs.joinmastodon.org/methods/grouped_notifications/#dismiss-group
const PostNotificationGroup_Dismiss = WriteNotifications

// GetGroupedNotifications_UnreadCount is the OAuth scope required by GET /api/v2/notifications/unread_count.
// https://docs.joinmastodon.org/methods/grouped_notifications/#unread-group-count
const GetGroupedNotifications_UnreadCount = ReadNotifications

// PostGroupedNotifications_Clear is the OAuth scope required by POST /api/v2/notifications/clear.
// https://docs.joinmastodon.org/methods/notifications/#clear
const PostGroupedNotifications_Clear = WriteNotifications
