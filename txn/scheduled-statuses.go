package txn

/******************************************
 * Scheduled Statuses API Methods
 * Manage statuses that were scheduled to be published at a future date.
 * https://docs.joinmastodon.org/methods/scheduled_statuses/
 ******************************************/

// GetScheduledStatuses is the input for GET /api/v1/scheduled_statuses, which returns []ScheduledStatus.
// https://docs.joinmastodon.org/methods/scheduled_statuses/#get
type GetScheduledStatuses struct {
	Host    string `header:"Host"`
	MaxID   string `query:"max_id"`
	SinceID string `query:"since_id"`
	MinID   string `query:"min_id"`
	Limit   int64  `query:"limit"`
}

// QueryPage implements the QueryPager interface, returning
// the QueryPage data embedded in this transaction
func (t GetScheduledStatuses) QueryPage() QueryPage {
	return QueryPage{
		MaxID:   t.MaxID,
		SinceID: t.SinceID,
		MinID:   t.MinID,
		Limit:   t.Limit,
	}
}

// GetScheduledStatus is the input for GET /api/v1/scheduled_statuses/:id, which returns ScheduledStatus.
// https://docs.joinmastodon.org/methods/scheduled_statuses/#get-one
type GetScheduledStatus struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}

// PutScheduledStatus is the input for PUT /api/v1/scheduled_statuses/:id, which returns ScheduledStatus.
// https://docs.joinmastodon.org/methods/scheduled_statuses/#update
type PutScheduledStatus struct {
	Host        string `header:"Host"`
	ID          string `param:"id"`
	ScheduledAt string `form:"scheduled_at" json:"scheduled_at"` // ISO 8601 Datetime
}

// DeleteScheduledStatus is the input for DELETE /api/v1/scheduled_statuses/:id, which returns an empty object.
// https://docs.joinmastodon.org/methods/scheduled_statuses/#cancel
type DeleteScheduledStatus struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}
