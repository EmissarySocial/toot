package txn

/******************************************
 * Filters API Methods
 * Create and manage filters
 * https://docs.joinmastodon.org/methods/filters/
 ******************************************/

// GetFilters is the input for GET /api/v2/filters, which returns []Filter.
// https://docs.joinmastodon.org/methods/filters/#get
type GetFilters struct {
	Host string `header:"Host"`
}

// GetFilter is the input for GET /api/v2/filters/:id, which returns Filter.
// https://docs.joinmastodon.org/methods/filters/#get-one
type GetFilter struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}

// PostFilter is the input for POST /api/v2/filters, which returns Filter.
// https://docs.joinmastodon.org/methods/filters/#create
type PostFilter struct {
	Host               string   `header:"Host"`
	Title              string   `form:"title" json:"title"`
	Context            []string `form:"context" json:"context"`
	FilterAction       string   `form:"filter_action" json:"filter_action"`
	ExpiresIn          int      `form:"expires_in" json:"expires_in"`
	KeywordsAttributes []struct {
		Host      string `header:"Host"`
		Keyword   string `form:"keyword" json:"keyword"`
		WholeWord bool   `form:"whole_word" json:"whole_word"`
	} `form:"keywords_attributes" json:"keywords_attributes"`
}

// PutFilter is the input for PUT /api/v2/filters/:id, which returns Filter.
// https://docs.joinmastodon.org/methods/filters/#update
type PutFilter struct {
	Host               string   `header:"Host"`
	ID                 string   `param:"id"`                                // The ID of the Filter in the database.
	Title              string   `form:"title" json:"title"`                 // The name of the filter group.
	Context            []string `form:"context" json:"context"`             // Where the filter should be applied. Specify at least one of home, notifications, public, thread, account.
	FilterAction       string   `form:"filter_action" json:"filter_action"` // The policy to be applied when the filter is matched. Specify warn or hide.
	ExpiresIn          int      `form:"expires_in" json:"expires_in"`       // How many seconds from now should the filter expire?
	KeywordsAttributes []struct {
		Host      string `header:"Host"`
		Keyword   string `form:"keyword" json:"keyword"`       // A keyword to be added to the newly-created filter group.
		WholeWord bool   `form:"whole_word" json:"whole_word"` // Whether the keyword should consider word boundaries.
		ID        string `form:"id" json:"id"`                 // Provide the ID of an existing keyword to modify it, instead of creating a new keyword.
		Destroy   bool   `form:"_destroy" json:"_destroy"`     // If true, will remove the keyword with the given ID.
	} `form:"keywords_attributes" json:"keywords_attributes"`
}

// DeleteFilter is the input for DELETE /api/v2/filters/:id, which returns an empty object.
// https://docs.joinmastodon.org/methods/filters/#delete
type DeleteFilter struct {
	Host string `header:"Host"`
	ID   string `param:"id"` // The ID of the Filter in the database.
}

// GetFilter_Keywords is the input for GET /api/v2/filters/:filter_id/keywords, which returns []FilterKeyword.
// https://docs.joinmastodon.org/methods/filters/#keywords-get
type GetFilter_Keywords struct {
	Host     string `header:"Host"`
	FilterID string `param:"filter_id"` // The ID of the Filter in the database.
}

// PostFilter_Keyword is the input for POST /api/v2/filters/:filter_id/keywords, which returns FilterKeyword.
// https://docs.joinmastodon.org/methods/filters/#keywords-create
type PostFilter_Keyword struct {
	Host      string `header:"Host"`
	FilterID  string `param:"filter_id"`                   // The ID of the Filter in the database.
	Keyword   string `form:"keyword" json:"keyword"`       // The keyword to be added to the filter group.
	WholeWord bool   `form:"whole_word" json:"whole_word"` // Whether the keyword should consider word boundaries.
}

// GetFilter_Keyword is the input for GET /api/v2/filters/keywords/:id, which returns FilterKeyword.
// https://docs.joinmastodon.org/methods/filters/#keywords-get-one
type GetFilter_Keyword struct {
	Host string `header:"Host"`
	ID   string `param:"id"` // The ID of the FilterKeyword in the database.
}

// PutFilter_Keyword is the input for PUT /api/v2/filters/keywords/:id, which returns FilterKeyword.
// https://docs.joinmastodon.org/methods/filters/#keywords-update
type PutFilter_Keyword struct {
	Host      string `header:"Host"`
	ID        string `param:"id"`                          // The ID of the FilterKeyword in the database.
	Keyword   string `form:"keyword" json:"keyword"`       // The keyword to be added to the filter group.
	WholeWord bool   `form:"whole_word" json:"whole_word"` // Whether the keyword should consider word boundaries.
}

// DeleteFilter_Keyword is the input for DELETE /api/v2/filters/keywords/:id, which returns an empty object.
// https://docs.joinmastodon.org/methods/filters/#keywords-delete
type DeleteFilter_Keyword struct {
	Host string `header:"Host"`
	ID   string `param:"id"` // The ID of the FilterKeyword in the database.
}

// GetFilter_Statuses is the input for GET /api/v2/filters/:filter_id/statuses, which returns []FilterStatus.
// https://docs.joinmastodon.org/methods/filters/#statuses-get
type GetFilter_Statuses struct {
	Host     string `header:"Host"`
	FilterID string `param:"filter_id"` // The ID of the Filter in the database.
}

// PostFilter_Status is the input for POST /api/v2/filters/:filter_id/statuses, which returns FilterStatus.
// https://docs.joinmastodon.org/methods/filters/#statuses-add
type PostFilter_Status struct {
	Host     string `header:"Host"`
	FilterID string `param:"filter_id"` // The ID of the Filter in the database.
}

// GetFilter_Status is the input for GET /api/v2/filters/statuses/:id, which returns FilterStatus.
// https://docs.joinmastodon.org/methods/filters/#statuses-get-one
type GetFilter_Status struct {
	Host string `header:"Host"`
	ID   string `param:"id"` // The ID of the FilterStatus in the database.
}

// DeleteFilter_Status is the input for DELETE /api/v2/filters/statuses/:id, which returns FilterStatus.
// https://docs.joinmastodon.org/methods/filters/#statuses-remove
type DeleteFilter_Status struct {
	Host string `header:"Host"`
	ID   string `param:"id"` // The ID of the FilterStatus in the database.
}

// GetFilters_V1 is the input for GET /api/v1/filters.
// https://docs.joinmastodon.org/methods/filters/#get-v1
type GetFilters_V1 struct {
	Host string `header:"Host"`
}

// GetFilter_V1 is the input for GET /api/v1/filters/:id.
// https://docs.joinmastodon.org/methods/filters/#get-one-v1
type GetFilter_V1 struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}

// PostFilter_V1 is the input for POST /api/v1/filters.
// https://docs.joinmastodon.org/methods/filters/#create-v1
type PostFilter_V1 struct {
	Host         string   `header:"Host"`
	Phrase       string   `form:"phrase" json:"phrase"`
	Context      []string `form:"context" json:"context"`
	Irreversible bool     `form:"irreversible" json:"irreversible"`
	WholeWord    bool     `form:"whole_word" json:"whole_word"`
	ExpiresIn    int      `form:"expires_in" json:"expires_in"`
}

// PutFilter_V1 is the input for PUT /api/v1/filters/:id.
// https://docs.joinmastodon.org/methods/filters/#update-v1
type PutFilter_V1 struct {
	Host         string   `header:"Host"`
	ID           string   `param:"id"`
	Phrase       string   `form:"phrase" json:"phrase"`
	Context      []string `form:"context" json:"context"`
	Irreversible bool     `form:"irreversible" json:"irreversible"`
	WholeWord    bool     `form:"whole_word" json:"whole_word"`
	ExpiresIn    int      `form:"expires_in" json:"expires_in"`
}

// DeleteFilter_V1 is the input for DELETE /api/v1/filters/:id.
// https://docs.joinmastodon.org/methods/filters/#delete-v1
type DeleteFilter_V1 struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}
