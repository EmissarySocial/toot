package object

// WebPushSubscription represents a subscription to the Web Push API.
// https://docs.joinmastodon.org/entities/WebPushSubscription/
type WebPushSubscription struct {
	ID        string                    `json:"id"`         //The ID of the subscription in the database
	Endpoint  string                    `json:"endpoint"`   //The endpoint of the subscription
	Standard  bool                      `json:"standard"`   //standard (RFC8030+RFC8291+RFC8292)
	Alerts    WebPushSubscriptionAlerts `json:"alerts"`     //which allert types are enabled for the subscription
	ServerKey string                    `json:"server_key"` //the streaming server's VAPID public key
}
