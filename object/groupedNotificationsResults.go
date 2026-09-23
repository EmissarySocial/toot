package object

// GroupedNotificationsResults is the response shape for the grouped notifications API.
// https://docs.joinmastodon.org/methods/grouped_notifications/#GroupedNotificationsResults
type GroupedNotificationsResults struct {
	Accounts           []Account                  `json:"accounts"`
	PartialAccounts    []PartialAccountWithAvatar `json:"partial_accounts,omitempty"`
	Statuses           []Status                   `json:"statuses"`
	NotificationGroups []NotificationGroup        `json:"notification_groups"`
}
