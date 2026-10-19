package txn

/******************************************
* Push API Methods
* Subscribe to and receive push notifications when a server-side
* notification is received, via the Web Push API
* https://docs.joinmastodon.org/methods/push/
******************************************/

// PostPushSubscription is the input for POST /api/v1/push/subscription, which returns WebPushSubscription.
// https://docs.joinmastodon.org/methods/push/#create
type PostPushSubscription struct {
	Host         string `header:"Host"`
	Subscription struct {
		Endpoint string `form:"endpoint" json:"endpoint"`
		Keys     struct {
			P256dh string `form:"p256dh" json:"p256dh"`
			Auth   string `form:"auth" json:"auth"`
		} `form:"keys" json:"keys"`
		Standard bool `form:"standard" json:"standard"`
	} `form:"subscription" json:"subscription"`
	Data struct {
		Alerts struct {
			Mention       bool `form:"mention" json:"mention"`
			Quote         bool `form:"quote" json:"quote"`
			Status        bool `form:"status" json:"status"`
			Reblog        bool `form:"reblog" json:"reblog"`
			Follow        bool `form:"follow" json:"follow"`
			FollowRequest bool `form:"follow_request" json:"follow_request"`
			Favourite     bool `form:"favourite" json:"favourite"`
			Poll          bool `form:"poll" json:"poll"`
			Update        bool `form:"update" json:"update"`
			QuotedUpdate  bool `form:"quoted_update" json:"quoted_update"`
			AdminSignUp   bool `form:"admin.sign_up" json:"admin.sign_up"`
			AdminReport   bool `form:"admin.report" json:"admin.report"`
		} `form:"alerts" json:"alerts"`
		Policy string `form:"policy" json:"policy"`
	} `form:"data" json:"data"`
}

// GetPushSubscription is the input for GET /api/v1/push/subscription, which returns WebPushSubscription.
// https://docs.joinmastodon.org/methods/push/#get
type GetPushSubscription struct {
	Host string `header:"Host"`
}

// PutPushSubscription is the input for PUT /api/v1/push/subscription, which returns WebPushSubscription.
// https://docs.joinmastodon.org/methods/push/#update
type PutPushSubscription struct {
	Host string `header:"Host"`
	Data struct {
		Alerts struct {
			Mention       bool `form:"mention" json:"mention"`
			Quote         bool `form:"quote" json:"quote"`
			Status        bool `form:"status" json:"status"`
			Reblog        bool `form:"reblog" json:"reblog"`
			Follow        bool `form:"follow" json:"follow"`
			FollowRequest bool `form:"follow_request" json:"follow_request"`
			Favourite     bool `form:"favourite" json:"favourite"`
			Poll          bool `form:"poll" json:"poll"`
			Update        bool `form:"update" json:"update"`
			QuotedUpdate  bool `form:"quoted_update" json:"quoted_update"`
			AdminSignUp   bool `form:"admin.sign_up" json:"admin.sign_up"`
			AdminReport   bool `form:"admin.report" json:"admin.report"`
		} `form:"alerts" json:"alerts"`
		Policy string `form:"policy" json:"policy"`
	} `form:"data" json:"data"`
}

// DeletePushSubscription is the input for DELETE /api/v1/push/subscription, which returns an empty object.
// https://docs.joinmastodon.org/methods/push/#delete
type DeletePushSubscription struct {
	Host string `header:"Host"`
}
