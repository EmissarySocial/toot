package route

/******************************************
 * Notifications API Methods
 * Receive notifications for activity on your account or statuses
 * https://docs.joinmastodon.org/methods/notifications/
 ******************************************/

// GetNotifications is the route for GET /api/v1/notifications.
// https://docs.joinmastodon.org/methods/notifications/#get
const GetNotifications = "/api/v1/notifications"

// GetNotification is the route for GET /api/v1/notifications/:id.
// https://docs.joinmastodon.org/methods/notifications/#get-one
const GetNotification = "/api/v1/notifications/:id"

// PostNotifications_Clear is the route for POST /api/v1/notifications/clear.
// https://docs.joinmastodon.org/methods/notifications/#clear
const PostNotifications_Clear = "/api/v1/notifications/clear"

// PostNotification_Dismiss is the route for POST /api/v1/notifications/dismiss.
// https://docs.joinmastodon.org/methods/notifications/#dismiss
const PostNotification_Dismiss = "/api/v1/notifications/dismiss"

// https://docs.joinmastodon.org/methods/notifications/#unread-count
const GetNotifications_UnreadCount = "/api/v1/notifications/unread_count"

// https://docs.joinmastodon.org/methods/notifications/#get-policy
const GetNotificationPolicy = "/api/v2/notifications/policy"

// https://docs.joinmastodon.org/methods/notifications/#update-policy
const PatchNotificationPolicy = "/api/v2/notifications/policy"

// GetNotificationRequests is the path for GET /api/v1/notifications/requests.
// https://docs.joinmastodon.org/methods/notifications/#get-requests
const GetNotificationRequests = "/api/v1/notifications/requests"

// GetNotificationRequest is the path for GET /api/v1/notifications/requests/:id.
// https://docs.joinmastodon.org/methods/notifications/#get-one-request
const GetNotificationRequest = "/api/v1/notifications/requests/:id"

// PostNotificationRequest_Accept is the path for POST /api/v1/notifications/requests/:id/accept.
// https://docs.joinmastodon.org/methods/notifications/#accept-request
const PostNotificationRequest_Accept = "/api/v1/notifications/requests/:id/accept"

// PostNotificationRequest_Dismiss is the path for POST /api/v1/notifications/requests/:id/dismiss.
// https://docs.joinmastodon.org/methods/notifications/#dismiss-request
const PostNotificationRequest_Dismiss = "/api/v1/notifications/requests/:id/dismiss"

// PostNotificationRequests_Accept is the path for POST /api/v1/notifications/requests/accept.
// https://docs.joinmastodon.org/methods/notifications/#accept-multiple-requests
const PostNotificationRequests_Accept = "/api/v1/notifications/requests/accept"

// PostNotificationRequests_Dismiss is the path for POST /api/v1/notifications/requests/dismiss.
// https://docs.joinmastodon.org/methods/notifications/#dismiss-multiple-requests
const PostNotificationRequests_Dismiss = "/api/v1/notifications/requests/dismiss"

// GetNotificationRequests_Merged is the path for GET /api/v1/notifications/requests/merged.
// https://docs.joinmastodon.org/methods/notifications/#requests-merged
const GetNotificationRequests_Merged = "/api/v1/notifications/requests/merged"
