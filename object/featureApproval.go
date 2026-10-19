package object

// FeatureApproval summarizes an account's "who can feature me in a Collection" policy.
// https://docs.joinmastodon.org/entities/Account/#feature_approval
type FeatureApproval struct {
	Automatic   []string `json:"automatic"`
	Manual      []string `json:"manual"`
	CurrentUser string   `json:"current_user"`
}
