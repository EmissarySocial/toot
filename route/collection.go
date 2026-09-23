package route

/******************************************
* Collections API Methods
* Manage collections of accounts to be recommended to others.
* https://docs.joinmastodon.org/methods/collections/
******************************************/

// https://docs.joinmastodon.org/methods/collections/#create
const PostCollection = "/api/v1/collections"

// https://docs.joinmastodon.org/methods/collections/#get_collection
const GetCollection = "/api/v1/collections/:id"

// https://docs.joinmastodon.org/methods/collections/#get_collections
const GetAccount_Collections = "/api/v1/accounts/:account_id/collections"

// https://docs.joinmastodon.org/methods/collections/#in_collections
const GetAccount_InCollections = "/api/v1/accounts/:account_id/in_collections"

// https://docs.joinmastodon.org/methods/collections/#update_collection
const PatchCollection = "/api/v1/collections/:id"

// https://docs.joinmastodon.org/methods/collections/#delete_collection
const DeleteCollection = "/api/v1/collections/:id"

// https://docs.joinmastodon.org/methods/collections/#add_account
const PostCollection_Item = "/api/v1/collections/:collection_id/items"

// https://docs.joinmastodon.org/methods/collections/#remove_account
const DeleteCollection_Item = "/api/v1/collections/:collection_id/items/:id"

// https://docs.joinmastodon.org/methods/collections/#revoke_item
const PostCollection_Item_Revoke = "/api/v1/collections/:collection_id/items/:id/revoke"
