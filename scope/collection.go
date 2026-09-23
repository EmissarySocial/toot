package scope

/******************************************
* Collections API Methods
* Manage collections of accounts to be recommended to others.
* https://docs.joinmastodon.org/methods/collections/
******************************************/

// https://docs.joinmastodon.org/methods/collections/#create
const PostCollection = WriteCollections

// https://docs.joinmastodon.org/methods/collections/#get_collection
const GetCollection = Public

// https://docs.joinmastodon.org/methods/collections/#get_collections
const GetAccount_Collections = Public

// https://docs.joinmastodon.org/methods/collections/#in_collections
const GetAccount_InCollections = ReadCollections

// https://docs.joinmastodon.org/methods/collections/#update_collection
const PatchCollection = WriteCollections

// https://docs.joinmastodon.org/methods/collections/#delete_collection
const DeleteCollection = WriteCollections

// https://docs.joinmastodon.org/methods/collections/#add_account
const PostCollection_Item = WriteCollections

// https://docs.joinmastodon.org/methods/collections/#remove_account
const DeleteCollection_Item = WriteCollections

// https://docs.joinmastodon.org/methods/collections/#revoke_item
const PostCollection_Item_Revoke = WriteCollections
