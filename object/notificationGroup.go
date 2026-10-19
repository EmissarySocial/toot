package object

// NotificationGroup represents a group of related notifications (e.g. multiple
// favourites of the same status), as returned by the grouped notifications API.
// https://docs.joinmastodon.org/entities/NotificationGroup/
type NotificationGroup struct {
	GroupKey                 string                      `json:"group_key"`
	NotificationsCount       int                         `json:"notifications_count"`
	Type                     string                      `json:"type"`
	MostRecentNotificationID string                      `json:"most_recent_notification_id"`
	PageMinID                string                      `json:"page_min_id,omitempty"`
	PageMaxID                string                      `json:"page_max_id,omitempty"`
	LatestPageNotificationAt string                      `json:"latest_page_notification_at,omitempty"`
	SampleAccountIDs         []string                    `json:"sample_account_ids"`
	StatusID                 string                      `json:"status_id,omitempty"`
	Report                   *Report                     `json:"report,omitempty"`
	Collection               *Collection                 `json:"collection,omitempty"`
	Fallback                 *NotificationFallback       `json:"fallback,omitempty"`
	Event                    *RelationshipSeveranceEvent `json:"event,omitempty"`
	ModerationWarning        *AccountWarning             `json:"moderation_warning,omitempty"`
	AnnualReport             *AnnualReportEvent          `json:"annual_report,omitempty"`
}
