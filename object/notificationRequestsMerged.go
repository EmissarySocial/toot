package object

// NotificationRequestsMerged reports whether previously-accepted notification requests
// have finished merging back into the main notifications list.
// https://docs.joinmastodon.org/methods/notifications/#requests-merged
type NotificationRequestsMerged struct {
	Merged bool `json:"merged"`
}
