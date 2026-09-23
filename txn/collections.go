package txn

/******************************************
 * Collections API Methods
 * Manage collections of accounts to be recommended to others.
 * https://docs.joinmastodon.org/methods/collections/
 ******************************************/

// PostCollection is the input for POST /api/v1/collections, which returns Collection.
// https://docs.joinmastodon.org/methods/collections/#create
type PostCollection struct {
	Host         string   `header:"Host"`
	Name         string   `form:"name"`
	Description  string   `form:"description"`
	Language     string   `form:"language"`
	TagName      string   `form:"tag_name"`
	Sensitive    bool     `form:"sensitive"`
	Discoverable bool     `form:"discoverable"`
	AccountIDs   []string `form:"account_ids"`
}

// GetCollection is the input for GET /api/v1/collections/:id, which returns CollectionWithAccounts.
// https://docs.joinmastodon.org/methods/collections/#get_collection
type GetCollection struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}

// GetAccount_Collections is the input for GET /api/v1/accounts/:account_id/collections, which returns []Collection.
// https://docs.joinmastodon.org/methods/collections/#get_collections
type GetAccount_Collections struct {
	Host      string `header:"Host"`
	AccountID string `param:"account_id"`
	Limit     int    `query:"limit"`
	Offset    int    `query:"offset"`
}

// GetAccount_InCollections is the input for GET /api/v1/accounts/:account_id/in_collections, which returns []Collection.
// https://docs.joinmastodon.org/methods/collections/#in_collections
type GetAccount_InCollections struct {
	Host      string `header:"Host"`
	AccountID string `param:"account_id"`
	Limit     int    `query:"limit"`
	Offset    int    `query:"offset"`
}

// PatchCollection is the input for PATCH /api/v1/collections/:id, which returns Collection.
// https://docs.joinmastodon.org/methods/collections/#update_collection
type PatchCollection struct {
	Host         string `header:"Host"`
	ID           string `param:"id"`
	Name         string `form:"name"`
	Description  string `form:"description"`
	Language     string `form:"language"`
	TagName      string `form:"tag_name"`
	Sensitive    bool   `form:"sensitive"`
	Discoverable bool   `form:"discoverable"`
}

// DeleteCollection is the input for DELETE /api/v1/collections/:id, which returns an empty object.
// https://docs.joinmastodon.org/methods/collections/#delete_collection
type DeleteCollection struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}

// PostCollection_Item is the input for POST /api/v1/collections/:collection_id/items, which returns CollectionItem.
// https://docs.joinmastodon.org/methods/collections/#add_account
type PostCollection_Item struct {
	Host         string `header:"Host"`
	CollectionID string `param:"collection_id"`
	AccountID    string `form:"account_id"`
}

// DeleteCollection_Item is the input for DELETE /api/v1/collections/:collection_id/items/:id, which returns an empty object.
// https://docs.joinmastodon.org/methods/collections/#remove_account
type DeleteCollection_Item struct {
	Host         string `header:"Host"`
	CollectionID string `param:"collection_id"`
	ID           string `param:"id"`
}

// PostCollection_Item_Revoke is the input for POST /api/v1/collections/:collection_id/items/:id/revoke, which returns an empty object.
// https://docs.joinmastodon.org/methods/collections/#revoke_item
type PostCollection_Item_Revoke struct {
	Host         string `header:"Host"`
	CollectionID string `param:"collection_id"`
	ID           string `param:"id"`
}
