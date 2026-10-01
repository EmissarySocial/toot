package txn

/******************************************
 * Timelines API Methods
 * Read and view timelines of statuses
 * https://docs.joinmastodon.org/methods/timelines/
 ******************************************/

// GetTimeline_Public is the input for GET /api/v1/timelines/public, which returns []Status.
// https://docs.joinmastodon.org/methods/timelines/#public
type GetTimeline_Public struct {
	Host      string `header:"Host"`
	Local     bool   `query:"local"`
	Remote    bool   `query:"remote"`
	OnlyMedia bool   `query:"only_media"`
	MaxID     string `query:"max_id"`
	SinceID   string `query:"since_id"`
	MinID     string `query:"min_id"`
	Limit     int64  `query:"limit"`
}

// QueryPage implements the QueryPager interface, returning
// the QueryPage data embedded in this transaction
func (t GetTimeline_Public) QueryPage() QueryPage {
	return QueryPage{
		MaxID:   t.MaxID,
		SinceID: t.SinceID,
		MinID:   t.MinID,
		Limit:   t.Limit,
	}
}

// GetTimeline_Hashtag is the input for GET /api/v1/timelines/tag/:hashtag, which returns []Status.
// https://docs.joinmastodon.org/methods/timelines/#tag
type GetTimeline_Hashtag struct {
	Host      string   `header:"Host"`
	Hashtag   string   `param:"hashtag"`
	Any       []string `query:"any"`
	All       []string `query:"all"`
	None      []string `query:"none"`
	Local     bool     `query:"local"`
	Remote    bool     `query:"remote"`
	OnlyMedia bool     `query:"only_media"`
	MaxID     string   `query:"max_id"`
	SinceID   string   `query:"since_id"`
	MinID     string   `query:"min_id"`
	Limit     int64    `query:"limit"`
}

// QueryPage implements the QueryPager interface, returning
// the QueryPage data embedded in this transaction
func (t GetTimeline_Hashtag) QueryPage() QueryPage {
	return QueryPage{
		MaxID:   t.MaxID,
		SinceID: t.SinceID,
		MinID:   t.MinID,
		Limit:   t.Limit,
	}
}

// GetTimeline_Home is the input for GET /api/v1/timelines/home, which returns []Status.
// https://docs.joinmastodon.org/methods/timelines/#home
type GetTimeline_Home struct {
	Host    string `header:"Host"`
	MaxID   string `query:"max_id"`
	SinceID string `query:"since_id"`
	MinID   string `query:"min_id"`
	Limit   int64  `query:"limit"`
}

// QueryPage implements the QueryPager interface, returning
// the QueryPage data embedded in this transaction
func (t GetTimeline_Home) QueryPage() QueryPage {
	return QueryPage{
		MaxID:   t.MaxID,
		SinceID: t.SinceID,
		MinID:   t.MinID,
		Limit:   t.Limit,
	}
}

// GetTimeline_List is the input for GET /api/v1/timelines/list/:list_id, which returns []Status.
// https://docs.joinmastodon.org/methods/timelines/#list
type GetTimeline_List struct {
	Host    string `header:"Host"`
	ListID  string `param:"list_id"`
	MaxID   string `query:"max_id"`
	SinceID string `query:"since_id"`
	MinID   string `query:"min_id"`
	Limit   int64  `query:"limit"`
}

// QueryPage implements the QueryPager interface, returning
// the QueryPage data embedded in this transaction
func (t GetTimeline_List) QueryPage() QueryPage {
	return QueryPage{
		MaxID:   t.MaxID,
		SinceID: t.SinceID,
		MinID:   t.MinID,
		Limit:   t.Limit,
	}
}
