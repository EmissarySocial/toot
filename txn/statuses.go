package txn

/******************************************
 * Statuses API Methods
 * Publish, interact, and view information about statuses
 * https://docs.joinmastodon.org/methods/statuses/
 ******************************************/

// PostStatus is the input for POST /api/v1/statuses, which returns Status.
// https://docs.joinmastodon.org/methods/statuses/#create
type PostStatus struct {
	Host     string   `header:"Host"`
	Status   string   `form:"status" json:"status"`
	MediaIDs []string `form:"media_ids" json:"media_ids"`
	Poll     struct {
		Host       string   `header:"Host"`
		Options    []string `form:"options" json:"options"`
		ExpiresIn  int      `form:"expires_in" json:"expires_in"`
		Multiple   bool     `form:"multiple" json:"multiple"`
		HideTotals bool     `form:"hide_totals" json:"hide_totals"`
	} `form:"poll" json:"poll"`
	InReplyToID string `form:"in_reply_to_id" json:"in_reply_to_id"`
	Sensitive   bool   `form:"sensitive" json:"sensitive"`
	SpoilerText string `form:"spoiler_text" json:"spoiler_text"`
	Visibility  string `form:"visibility" json:"visibility"`     // [public | unlisted | private | direct]
	Language    string `form:"language" json:"language"`         // ISO 629 2-letter language code
	ScheduledAt string `form:"scheduled_at" json:"scheduled_at"` // ISO 8601 Datetime
}

// GetStatus is the input for GET /api/v1/statuses/:id, which returns Status.
// https://docs.joinmastodon.org/methods/statuses/#get
type GetStatus struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}

// DeleteStatus is the input for DELETE /api/v1/statuses/:id.
// https://docs.joinmastodon.org/methods/statuses/#delete
type DeleteStatus struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}

// GetStatus_Context is the input for GET /api/v1/statuses/:id/context, which returns Context.
// https://docs.joinmastodon.org/methods/statuses/#context
type GetStatus_Context struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}

// PostStatus_Translate is the input for POST /api/v1/statuses/:id/translate, which returns Status.
// https://docs.joinmastodon.org/methods/statuses/#translate
type PostStatus_Translate struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
	Lang string `form:"lang" json:"lang"`
}

// GetStatus_RebloggedBy is the input for GET /api/v1/statuses/:id/reblogged_by, which returns []Account.
// https://docs.joinmastodon.org/methods/statuses/#reblogged_by
type GetStatus_RebloggedBy struct {
	Host    string `header:"Host"`
	ID      string `param:"id"`
	MaxID   string `query:"max_id"`
	MinID   string `query:"min_id"`
	SinceID string `query:"since_id"`
	Limit   int64  `query:"limit"`
}

// QueryPage implements the QueryPager interface, returning
// the QueryPage data embedded in this transaction
func (t GetStatus_RebloggedBy) QueryPage() QueryPage {
	return QueryPage{
		MaxID:   t.MaxID,
		SinceID: t.SinceID,
		MinID:   t.MinID,
		Limit:   t.Limit,
	}
}

// GetStatus_FavouritedBy is the input for GET /api/v1/statuses/:id/favourited_by, which returns []Account.
// https://docs.joinmastodon.org/methods/statuses/#favourited_by
type GetStatus_FavouritedBy struct {
	Host    string `header:"Host"`
	ID      string `param:"id"`
	MaxID   string `query:"max_id"`
	MinID   string `query:"min_id"`
	SinceID string `query:"since_id"`
	Limit   int64  `query:"limit"`
}

// QueryPage implements the QueryPager interface, returning
// the QueryPage data embedded in this transaction
func (t GetStatus_FavouritedBy) QueryPage() QueryPage {
	return QueryPage{
		MaxID:   t.MaxID,
		SinceID: t.SinceID,
		MinID:   t.MinID,
		Limit:   t.Limit,
	}
}

// PostStatus_Favourite is the input for POST /api/v1/statuses/:id/favourite, which returns Status.
// https://docs.joinmastodon.org/methods/statuses/#favourite
type PostStatus_Favourite struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}

// PostStatus_Unfavourite is the input for POST /api/v1/statuses/:id/unfavourite, which returns Status.
// https://docs.joinmastodon.org/methods/statuses/#unfavourite
type PostStatus_Unfavourite struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}

// PostStatus_Reblog is the input for POST /api/v1/statuses/:id/reblog, which returns Status.
// https://docs.joinmastodon.org/methods/statuses/#reblog
type PostStatus_Reblog struct {
	Host       string `header:"Host"`
	ID         string `param:"id"`
	Visibility string `form:"visibility" json:"visibility"` // [public | unlisted | private]
}

// PostStatus_Unreblog is the input for POST /api/v1/statuses/:id/unreblog, which returns Status.
// https://docs.joinmastodon.org/methods/statuses/#unreblog
type PostStatus_Unreblog struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}

// PostStatus_Bookmark is the input for POST /api/v1/statuses/:id/bookmark, which returns Status.
// https://docs.joinmastodon.org/methods/statuses/#bookmark
type PostStatus_Bookmark struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}

// PostStatus_Unbookmark is the input for POST /api/v1/statuses/:id/unbookmark, which returns Status.
// https://docs.joinmastodon.org/methods/statuses/#unbookmark
type PostStatus_Unbookmark struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}

// PostStatus_Mute is the input for POST /api/v1/statuses/:id/mute, which returns Status.
// https://docs.joinmastodon.org/methods/statuses/#mute
type PostStatus_Mute struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}

// PostStatus_Unmute is the input for POST /api/v1/statuses/:id/unmute, which returns Status.
// https://docs.joinmastodon.org/methods/statuses/#unmute
type PostStatus_Unmute struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}

// PostStatus_Pin is the input for POST /api/v1/statuses/:id/pin, which returns Status.
// https://docs.joinmastodon.org/methods/statuses/#pin
type PostStatus_Pin struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}

// PostStatus_Unpin is the input for POST /api/v1/statuses/:id/unpin, which returns Status.
// https://docs.joinmastodon.org/methods/statuses/#unpin
type PostStatus_Unpin struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}

// PutStatus is the input for PUT /api/v1/statuses/:id, which returns Status.
// https://docs.joinmastodon.org/methods/statuses/#edit
type PutStatus struct {
	Host        string   `header:"Host"`
	ID          string   `param:"id"`
	Status      string   `form:"status" json:"status"`
	SpoilerText string   `form:"spoiler_text" json:"spoiler_text"`
	Sensitive   bool     `form:"sensitive" json:"sensitive"`
	Language    string   `form:"language" json:"language"`
	MediaIDs    []string `form:"media_ids[]" json:"media_ids"`
	Poll        struct {
		Host       string   `header:"Host"`
		Options    []string `form:"options[]" json:"options"`
		ExpiresIn  int      `form:"expires_in" json:"expires_in"`
		Multiple   bool     `form:"multiple" json:"multiple"`
		HideTotals bool     `form:"hide_totals" json:"hide_totals"`
	} `form:"poll" json:"poll"`
}

// GetStatus_History is the input for GET /api/v1/statuses/:id/history, which returns []StatusEdit.
// https://docs.joinmastodon.org/methods/statuses/#history
type GetStatus_History struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}

// GetStatus_Source is the input for GET /api/v1/statuses/:id/source, which returns StatusSource.
// https://docs.joinmastodon.org/methods/statuses/#source
type GetStatus_Source struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}
