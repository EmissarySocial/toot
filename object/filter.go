package object

// Filter represents a user-defined filter for determining which statuses should not be shown to the user.
// https://docs.joinmastodon.org/entities/Filter/
type Filter struct {
	ID           string                      `json:"id"`            // The ID of the Filter in the database.
	Title        string                      `json:"title"`         // A title given by the user to name the filter.
	Context      NonNullSlice[string]        `json:"context"`       // The contexts in which the filter should be applied. [home | notifications | public | thread | account]
	ExpiresAt    *string                     `json:"expires_at"`    // When the filter should no longer be applied (ISO 8601 Datetime), or null for never.
	FilterAction string                      `json:"filter_action"` // The action to be taken when a status matches this filter. [warn | hide | blur]
	Keywords     NonNullSlice[FilterKeyword] `json:"keywords"`      // The keywords grouped under this filter.
	Statuses     NonNullSlice[FilterStatus]  `json:"statuses"`      // The statuses grouped under this filter.
}
