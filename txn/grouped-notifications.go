package txn

/******************************************
 * Grouped Notifications API Methods
 * Receive grouped notifications for activity on your account or statuses.
 * https://docs.joinmastodon.org/methods/grouped_notifications/
 ******************************************/

// GetGroupedNotifications is the input for GET /api/v2/notifications, which returns GroupedNotificationsResults.
// https://docs.joinmastodon.org/methods/grouped_notifications/#get-grouped
type GetGroupedNotifications struct {
	Host            string   `header:"Host"`
	MaxID           string   `query:"max_id"`
	SinceID         string   `query:"since_id"`
	MinID           string   `query:"min_id"`
	Limit           int64    `query:"limit"`
	Types           []string `query:"types"`
	ExcludeTypes    []string `query:"exclude_types"`
	AccountID       string   `query:"account_id"`
	ExpandAccounts  string   `query:"expand_accounts"`
	GroupedTypes    []string `query:"grouped_types"`
	IncludeFiltered bool     `query:"include_filtered"`
	SupportedTypes  []string `query:"supported_types"`
}

func (t GetGroupedNotifications) QueryPage() QueryPage {
	return QueryPage{MaxID: t.MaxID, SinceID: t.SinceID, MinID: t.MinID, Limit: t.Limit}
}

// GetNotificationGroup is the input for GET /api/v2/notifications/:group_key, which returns GroupedNotificationsResults.
// https://docs.joinmastodon.org/methods/grouped_notifications/#get-notification-group
type GetNotificationGroup struct {
	Host           string   `header:"Host"`
	GroupKey       string   `param:"group_key"`
	SupportedTypes []string `query:"supported_types"`
}

// PostNotificationGroup_Dismiss is the input for POST /api/v2/notifications/:group_key/dismiss, which returns an empty object.
// https://docs.joinmastodon.org/methods/grouped_notifications/#dismiss-group
type PostNotificationGroup_Dismiss struct {
	Host     string `header:"Host"`
	GroupKey string `param:"group_key"`
}

// GetGroupedNotifications_UnreadCount is the input for GET /api/v2/notifications/unread_count, which returns NotificationsUnreadCount.
// https://docs.joinmastodon.org/methods/grouped_notifications/#unread-group-count
type GetGroupedNotifications_UnreadCount struct {
	Host         string   `header:"Host"`
	Limit        int64    `query:"limit"`
	Types        []string `query:"types"`
	ExcludeTypes []string `query:"exclude_types"`
	AccountID    string   `query:"account_id"`
	GroupedTypes []string `query:"grouped_types"`
}

// PostGroupedNotifications_Clear is the input for POST /api/v2/notifications/clear, which returns an empty object.
// https://docs.joinmastodon.org/methods/notifications/#clear
type PostGroupedNotifications_Clear struct {
	Host string `header:"Host"`
}
