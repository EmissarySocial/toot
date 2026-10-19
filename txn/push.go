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
		Endpoint string `form:"endpoint"`
		Keys     struct {
			P256dh string `form:"p256dh"`
			Auth   string `form:"auth"`
		} `form:"keys"`
		Standard bool `form:"standard"`
	} `form:"subscription"`
	Data struct {
		Alerts struct {
			Mention       bool `form:"mention"`
			Quote         bool `form:"quote"`
			Status        bool `form:"status"`
			Reblog        bool `form:"reblog"`
			Follow        bool `form:"follow"`
			FollowRequest bool `form:"follow_request"`
			Favourite     bool `form:"favourite"`
			Poll          bool `form:"poll"`
			Update        bool `form:"update"`
			QuotedUpdate  bool `form:"quoted_update"`
			AdminSignUp   bool `form:"admin.sign_up"`
			AdminReport   bool `form:"admin.report"`
		} `form:"alerts"`
		Policy string `form:"policy"`
	} `form:"data"`
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
			Mention       bool `form:"mention"`
			Quote         bool `form:"quote"`
			Status        bool `form:"status"`
			Reblog        bool `form:"reblog"`
			Follow        bool `form:"follow"`
			FollowRequest bool `form:"follow_request"`
			Favourite     bool `form:"favourite"`
			Poll          bool `form:"poll"`
			Update        bool `form:"update"`
			QuotedUpdate  bool `form:"quoted_update"`
			AdminSignUp   bool `form:"admin.sign_up"`
			AdminReport   bool `form:"admin.report"`
		} `form:"alerts"`
		Policy string `form:"policy"`
	} `form:"data"`
}

// DeletePushSubscription is the input for DELETE /api/v1/push/subscription, which returns an empty object.
// https://docs.joinmastodon.org/methods/push/#delete
type DeletePushSubscription struct {
	Host string `header:"Host"`
}
