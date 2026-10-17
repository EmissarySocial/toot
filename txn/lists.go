package txn

/******************************************
 * Lists API Methods
 * View and manage lists. See also /api/v1/timelines/list/:id for loading a list timeline
 * https://docs.joinmastodon.org/methods/lists/
 ******************************************/

// GetLists is the input for GET /api/v1/lists, which returns []List.
// https://docs.joinmastodon.org/methods/lists/#get
type GetLists struct {
	Host string `header:"Host"`
}

// GetList is the input for GET /api/v1/lists/:id, which returns List.
// https://docs.joinmastodon.org/methods/lists/#get-one
type GetList struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}

// PostList is the input for POST /api/v1/lists, which returns List.
// https://docs.joinmastodon.org/methods/lists/#create
type PostList struct {
	Host          string `header:"Host"`
	Title         string `form:"title" json:"title"`
	RepliesPolicy string `form:"replies_policy" json:"replies_policy"`
	Exclusive     bool   `form:"exclusive" json:"exclusive"`
}

// PutList is the input for PUT /api/v1/lists/:id, which returns List.
// https://docs.joinmastodon.org/methods/lists/#update
type PutList struct {
	Host          string `header:"Host"`
	ID            string `param:"id"`
	Title         string `form:"title" json:"title"`
	RepliesPolicy string `form:"replies_policy" json:"replies_policy"`
}

// DeleteList is the input for DELETE /api/v1/lists/:id.
// https://docs.joinmastodon.org/methods/lists/#delete
type DeleteList struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}

// GetList_Accounts is the input for GET /api/v1/lists/:id/accounts, which returns []Account.
// https://docs.joinmastodon.org/methods/lists/#accounts
type GetList_Accounts struct {
	Host    string `header:"Host"`
	ID      string `param:"id"`
	MaxID   string `query:"max_id"`
	SinceID string `query:"since_id"`
	MinID   string `query:"min_id"`
	Limit   int64  `query:"limit"`
}

// QueryPage implements the QueryPager interface, returning
// the QueryPage data embedded in this transaction
func (t GetList_Accounts) QueryPage() QueryPage {
	return QueryPage{
		MaxID:   t.MaxID,
		SinceID: t.SinceID,
		MinID:   t.MinID,
		Limit:   t.Limit,
	}
}

// PostList_Accounts is the input for POST /api/v1/lists/:id/accounts, which returns an empty object.
// https://docs.joinmastodon.org/methods/lists/#accounts-add
type PostList_Accounts struct {
	Host       string   `header:"Host"`
	ID         string   `param:"id"`
	AccountIDs []string `form:"account_ids" json:"account_ids"`
}

// DeleteList_Accounts is the input for DELETE /api/v1/lists/:id/accounts, which returns an empty object.
// https://docs.joinmastodon.org/methods/lists/#accounts-remove
type DeleteList_Accounts struct {
	Host       string   `header:"Host"`
	ID         string   `param:"id"`
	AccountIDs []string `form:"account_ids" json:"account_ids"`
}
