package object

// WebPushSubscriptionAlerts represents which notification types trigger a
// push alert for a given WebPushSubscription.
// https://docs.joinmastodon.org/entities/WebPushSubscription/#alerts
type WebPushSubscriptionAlerts struct {
	Mention       bool `json:"mention"`
	Quote         bool `json:"quote"`
	Status        bool `json:"status"`
	Reblog        bool `json:"reblog"`
	Follow        bool `json:"follow"`
	FollowRequest bool `json:"follow_request"`
	Favourite     bool `json:"favourite"`
	Poll          bool `json:"poll"`
	Update        bool `json:"update"`
	QuotedUpdate  bool `json:"quoted_update"`
	AdminSignUp   bool `json:"admin.sign_up"` // dotted JSON key, intentional
	AdminReport   bool `json:"admin.report"`  // dotted JSON key, intentional
}
