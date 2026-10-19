package route

/******************************************
* Push API Methods
* Subscribe to and receive push notifications when a server-side
* notification is received, via the Web Push API
* https://docs.joinmastodon.org/methods/push/
******************************************/

// https://docs.joinmastodon.org/methods/push/#create
const PostPushSubscription = "/api/v1/push/subscription"

// https://docs.joinmastodon.org/methods/push/#get
const GetPushSubscription = "/api/v1/push/subscription"

// https://docs.joinmastodon.org/methods/push/#update
const PutPushSubscription = "/api/v1/push/subscription"

// https://docs.joinmastodon.org/methods/push/#delete
const DeletePushSubscription = "/api/v1/push/subscription"
