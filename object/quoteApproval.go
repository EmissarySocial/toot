package object

// QuoteApproval summarizes a status' quote approval policy and how it
// applies to the requesting user.
// https://docs.joinmastodon.org/entities/QuoteApproval/
type QuoteApproval struct {
	Automatic   []string `json:"automatic"`    // Who is auto-approved to quote this status. [public | followers | following | unsupported_policy]
	Manual      []string `json:"manual"`       // Who requires manual review to quote this status. [public | followers | following | unsupported_policy]
	CurrentUser string   `json:"current_user"` // How this policy applies to the requesting user. [automatic | manual | denied | unknown]
}
