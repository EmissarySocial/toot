package txn

/******************************************
 * Accounts API Methods
 * Methods concerning accounts and profiles
 * https://docs.joinmastodon.org/methods/accounts/
 ******************************************/

// PostAccount is the input for POST /api/v1/accounts, which returns Token.
// https://docs.joinmastodon.org/methods/accounts/#create
type PostAccount struct {
	Host      string `header:"Host"`
	Username  string `form:"username"`
	Email     string `form:"email"`
	Password  string `form:"password"`
	Agreement bool   `form:"agreement"`
	Locale    string `form:"locale"`
	Reason    string `form:"reason"`
}

// GetAccount_VerifyCredentials is the input for GET /api/v1/accounts/verify_credentials, which returns CredentialAccount.
// https://docs.joinmastodon.org/methods/accounts/#verify_credentials
type GetAccount_VerifyCredentials struct {
	Host string `header:"Host"`
}

// PatchAccount_UpdateCredentials is the input for PATCH /api/v1/accounts/update_credentials.
// https://docs.joinmastodon.org/methods/accounts/#update_credentials
type PatchAccount_UpdateCredentials struct {
	Host         string `header:"Host"`
	DisplayName  string `form:"display_name"`
	Note         string `form:"note"`
	Avatar       string `form:"avatar"`
	Header       string `form:"header"`
	Locked       bool   `form:"locked"`
	Bot          bool   `form:"bot"`
	Discoverable bool   `form:"discoverable"`
}

// GetAccounts is the input for GET /api/v1/accounts, a batch lookup that returns []Account.
// https://docs.joinmastodon.org/methods/accounts/#index
type GetAccounts struct {
	Host string   `header:"Host"`
	IDs  []string `query:"id[]"`
}

// GetAccount is the input for GET /api/v1/accounts/:id.
// https://docs.joinmastodon.org/methods/accounts/#get
type GetAccount struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}

// GetAccount_Statuses is the input for GET /api/v1/accounts/:id/statuses.
// https://docs.joinmastodon.org/methods/accounts/#statuses
type GetAccount_Statuses struct {
	Host           string `header:"Host"`
	ID             string `param:"id"`
	MaxID          string `query:"max_id"`
	SinceID        string `query:"since_id"`
	MinID          string `query:"min_id"`
	Limit          int64  `query:"limit"`
	OnlyMedia      bool   `query:"only_media"`
	ExcludeReplies bool   `query:"exclude_replies"`
	ExcludeReblogs bool   `query:"exclude_reblogs"`
	Pinned         bool   `query:"pinned"`
	Tagged         string `query:"tagged"`
}

// QueryPage implements the QueryPager interface, returning
// the QueryPage data embedded in this transaction
func (t GetAccount_Statuses) QueryPage() QueryPage {
	return QueryPage{
		MaxID:   t.MaxID,
		SinceID: t.SinceID,
		MinID:   t.MinID,
		Limit:   t.Limit,
	}
}

// GetAccount_Followers is the input for GET /api/v1/accounts/:id/followers.
// https://docs.joinmastodon.org/methods/accounts/#followers
type GetAccount_Followers struct {
	Host    string `header:"Host"`
	ID      string `param:"id"`
	MaxID   string `query:"max_id"`
	SinceID string `query:"since_id"`
	MinID   string `query:"min_id"`
	Limit   int64  `query:"limit"`
}

// QueryPage implements the QueryPager interface, returning
// the QueryPage data embedded in this transaction
func (t GetAccount_Followers) QueryPage() QueryPage {
	return QueryPage{
		MaxID:   t.MaxID,
		SinceID: t.SinceID,
		MinID:   t.MinID,
		Limit:   t.Limit,
	}
}

// GetAccount_Endorsements is the input for GET /api/v1/accounts/:id/endorsements,
// which returns []Account -- accounts featured on this account's profile.
// https://docs.joinmastodon.org/methods/accounts/#endorsements
type GetAccount_Endorsements struct {
	Host    string `header:"Host"`
	ID      string `param:"id"`
	MaxID   string `query:"max_id"`
	SinceID string `query:"since_id"`
	MinID   string `query:"min_id"`
	Limit   int64  `query:"limit"`
}

// QueryPage implements the QueryPager interface, returning
// the QueryPage data embedded in this transaction
func (t GetAccount_Endorsements) QueryPage() QueryPage {
	return QueryPage{
		MaxID:   t.MaxID,
		SinceID: t.SinceID,
		MinID:   t.MinID,
		Limit:   t.Limit,
	}
}

// GetAccount_Following is the input for GET /api/v1/accounts/:id/following.
// https://docs.joinmastodon.org/methods/accounts/#following
type GetAccount_Following struct {
	Host    string `header:"Host"`
	ID      string `param:"id"`
	MaxID   string `query:"max_id"`
	SinceID string `query:"since_id"`
	MinID   string `query:"min_id"`
	Limit   int64  `query:"limit"`
}

// QueryPage implements the QueryPager interface, returning
// the QueryPage data embedded in this transaction
func (t GetAccount_Following) QueryPage() QueryPage {
	return QueryPage{
		MaxID:   t.MaxID,
		SinceID: t.SinceID,
		MinID:   t.MinID,
		Limit:   t.Limit,
	}
}

// GetAccount_FeaturedTags is the input for GET /api/v1/accounts/:id/featured_tags.
// https://docs.joinmastodon.org/methods/accounts/#featured_tags
type GetAccount_FeaturedTags struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}

// GetAccount_Lists is the input for GET /api/v1/accounts/:id/lists.
// https://docs.joinmastodon.org/methods/accounts/#lists
type GetAccount_Lists struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}

// PostAccount_Follow is the input for POST /api/v1/accounts/:id/follow.
// https://docs.joinmastodon.org/methods/accounts/#follow
type PostAccount_Follow struct {
	Host      string   `header:"Host"`
	ID        string   `param:"id"`
	Reblogs   bool     `form:"reblogs"`
	Notify    bool     `form:"notify"`
	Languages []string `form:"languages"`
}

// PostAccount_Unfollow is the input for POST /api/v1/accounts/:id/unfollow.
// https://docs.joinmastodon.org/methods/accounts/#unfollow
type PostAccount_Unfollow struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}

// PostAccount_RemoveFromFollowers is the input for POST /api/v1/accounts/:id/remove_from_followers, which returns Relationship.
// https://docs.joinmastodon.org/methods/accounts/#remove_from_followers
type PostAccount_RemoveFromFollowers struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}

// PostAccount_Block is the input for POST /api/v1/accounts/:id/block, which returns Relationship.
// https://docs.joinmastodon.org/methods/accounts/#block
type PostAccount_Block struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}

// PostAccount_Unblock is the input for POST /api/v1/accounts/:id/unblock, which returns Relationship.
// https://docs.joinmastodon.org/methods/accounts/#unblock
type PostAccount_Unblock struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}

// PostAccount_Mute is the input for POST /api/v1/accounts/:id/mute, which returns Relationship.
// https://docs.joinmastodon.org/methods/accounts/#mute
type PostAccount_Mute struct {
	Host          string `header:"Host"`
	ID            string `param:"id"`
	Notifications bool   `form:"notifications"`
	Duration      bool   `form:"duration"`
}

// PostAccount_Unmute is the input for POST /api/v1/accounts/:id/unmute, which returns Relationship.
// https://docs.joinmastodon.org/methods/accounts/#unmute
type PostAccount_Unmute struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}

// PostAccount_Pin is the input for POST /api/v1/accounts/:id/pin, which returns Relationship.
// https://docs.joinmastodon.org/methods/accounts/#pin
type PostAccount_Pin struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}

// PostAccount_Unpin is the input for POST /api/v1/accounts/:id/unpin, which returns Relationship.
// https://docs.joinmastodon.org/methods/accounts/#unpin
type PostAccount_Unpin struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}

// PostAccount_Note is the input for POST /api/v1/accounts/:id/note, which returns Relationship.
// https://docs.joinmastodon.org/methods/accounts/#note
type PostAccount_Note struct {
	Host    string `header:"Host"`
	ID      string `param:"id"`
	Comment string `form:"comment"`
}

// GetAccount_Relationships is the input for GET /api/v1/accounts/relationships, which returns []Relationships.
// https://docs.joinmastodon.org/methods/accounts/#relationships
type GetAccount_Relationships struct {
	Host string   `header:"Host"`
	IDs  []string `query:"id[]"`
}

// GetAccount_FamiliarFollowers is the input for GET /api/v1/accounts/familiar_followers, which returns []FamiliarFollower.
// This is a batch lookup, same shape as GetAccount_Relationships just above.
// https://docs.joinmastodon.org/methods/accounts/#familiar_followers
type GetAccount_FamiliarFollowers struct {
	Host string   `header:"Host"`
	IDs  []string `query:"id[]"`
}

// GetAccount_Search is the input for GET /api/v1/accounts/search, which returns []Account.
// https://docs.joinmastodon.org/methods/accounts/#search
type GetAccount_Search struct {
	Host      string `header:"Host"`
	Q         string `query:"q"`
	Limit     int    `query:"limit"`
	Offset    int    `query:"offset"`
	Resolve   bool   `query:"resolve"`
	Following bool   `query:"following"`
}

// GetAccount_Lookup is the input for GET /api/v1/accounts/lookup, which returns Account.
// https://docs.joinmastodon.org/methods/accounts/#lookup
type GetAccount_Lookup struct {
	Host string `header:"Host"`
	Acct string `query:"acct"`
}
