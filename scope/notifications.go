package scope

/******************************************
 * Notifications API Methods
 * Receive notifications for activity on your account or statuses
 * https://docs.joinmastodon.org/methods/notifications/
 ******************************************/

// GetNotifications is the OAuth scope required by GET /api/v1/notifications.
// https://docs.joinmastodon.org/methods/notifications/#get
const GetNotifications = ReadNotifications

// GetNotification is the OAuth scope required by GET /api/v1/notifications/:id.
// https://docs.joinmastodon.org/methods/notifications/#get-one
const GetNotification = ReadNotifications

// PostNotifications_Clear is the OAuth scope required by POST /api/v1/notifications/clear.
// https://docs.joinmastodon.org/methods/notifications/#clear
const PostNotifications_Clear = WriteNotifications

// PostNotification_Dismiss is the OAuth scope required by POST /api/v1/notifications/dismiss.
// https://docs.joinmastodon.org/methods/notifications/#dismiss
const PostNotification_Dismiss = WriteNotifications

// https://docs.joinmastodon.org/methods/notifications/#unread-count
const GetNotifications_UnreadCount = ReadNotifications

// https://docs.joinmastodon.org/methods/notifications/#get-policy
const GetNotificationPolicy = ReadNotifications

// https://docs.joinmastodon.org/methods/notifications/#update-policy
const PatchNotificationPolicy = WriteNotifications

// https://docs.joinmastodon.org/methods/notifications/#get-requests
const GetNotificationRequests = ReadNotifications

// https://docs.joinmastodon.org/methods/notifications/#get-one-request
const GetNotificationRequest = ReadNotifications

// https://docs.joinmastodon.org/methods/notifications/#accept-request
const PostNotificationRequest_Accept = WriteNotifications

// https://docs.joinmastodon.org/methods/notifications/#dismiss-request
const PostNotificationRequest_Dismiss = WriteNotifications

// https://docs.joinmastodon.org/methods/notifications/#accept-multiple-requests
const PostNotificationRequests_Accept = WriteNotifications

// https://docs.joinmastodon.org/methods/notifications/#dismiss-multiple-requests
const PostNotificationRequests_Dismiss = WriteNotifications

// https://docs.joinmastodon.org/methods/notifications/#requests-merged
const GetNotificationRequests_Merged = ReadNotifications
