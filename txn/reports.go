package txn

/******************************************
 * Reports API Methods
 * Report problematic users to your moderators
 * https://docs.joinmastodon.org/methods/reports/
 ******************************************/

// PostReport is the input for POST /api/v1/reports, which returns Report.
// https://docs.joinmastodon.org/methods/reports/#post
type PostReport struct {
	Host      string   `header:"Host"`
	AccountID string   `form:"account_id" json:"account_id"`
	StatusIDs []string `form:"status_ids" json:"status_ids"`
	Comment   string   `form:"comment" json:"comment"`
	Forward   bool     `form:"forward" json:"forward"`
	Category  string   `form:"category" json:"category"`
	RuleIDs   []string `form:"rule_ids" json:"rule_ids"`
}
