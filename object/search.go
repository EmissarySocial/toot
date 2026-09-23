package object

// Search represents the results of a search.
// https://docs.joinmastodon.org/entities/Search/
type Search struct {
	Accounts NonNullSlice[Account] `json:"accounts"` // Accounts which match the given query
	Statuses NonNullSlice[Status]  `json:"statuses"` // Statuses which match the given query
	Hashtags NonNullSlice[Tag]     `json:"hashtags"` // Hashtags which match the given query
}
