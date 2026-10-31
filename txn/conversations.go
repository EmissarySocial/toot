package txn

/******************************************
 * Conversations API Methods
 * Direct conversations with other participants.
 * (Currently, just threads containing a post with "direct" visibility.)
 * https://docs.joinmastodon.org/methods/conversations/
 ******************************************/

// GetConversations is the input for GET /api/v1/conversations, which returns []Conversation.
// https://docs.joinmastodon.org/methods/conversations/#get
type GetConversations struct {
	Host    string `header:"Host"`
	MaxID   string `query:"max_id"`
	SinceID string `query:"since_id"`
	MinID   string `query:"min_id"`
	Limit   int64  `query:"limit"`
}

// QueryPage implements the QueryPager interface, returning
// the QueryPage data embedded in this transaction
func (t GetConversations) QueryPage() QueryPage {
	return QueryPage{
		MaxID:   t.MaxID,
		SinceID: t.SinceID,
		MinID:   t.MinID,
		Limit:   t.Limit,
	}
}

// DeleteConversation is the input for DELETE /api/v1/conversations/:id, which returns an empty object.
// https://docs.joinmastodon.org/methods/conversations/#delete
type DeleteConversation struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}

// PostConversationRead is the input for POST /api/v1/conversations/:id/read, which returns Conversation.
// https://docs.joinmastodon.org/methods/conversations/#read
type PostConversationRead struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}
