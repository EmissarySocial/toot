package object

// NotificationPolicy represents the notification filtering policy of the user.
// https://docs.joinmastodon.org/entities/NotificationPolicy/
type NotificationPolicy struct {
	ForNotFollowing    string                    `json:"for_not_following"`    // [accept | filter | drop]
	ForNotFollowers    string                    `json:"for_not_followers"`    // [accept | filter | drop]
	ForNewAccounts     string                    `json:"for_new_accounts"`     // [accept | filter | drop]
	ForPrivateMentions string                    `json:"for_private_mentions"` // [accept | filter | drop]
	ForLimitedAccounts string                    `json:"for_limited_accounts"` // [accept | filter | drop]
	ForBots            string                    `json:"for_bots"`             // [accept | filter | drop] - undocumented, verified against real server code
	Summary            NotificationPolicySummary `json:"summary"`
}
