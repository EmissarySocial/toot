package scope

/******************************************
 * Accounts API Methods
 * Methods concerning accounts and profiles
 * https://docs.joinmastodon.org/methods/accounts/
 ******************************************/

// PostAccount is the OAuth scope required by POST /api/v1/accounts.
// https://docs.joinmastodon.org/methods/accounts/#create
const PostAccount = WriteAccounts

// GetAccount_VerifyCredentials is the OAuth scope required by GET /api/v1/accounts/verify_credentials.
// https://docs.joinmastodon.org/methods/accounts/#verify_credentials
const GetAccount_VerifyCredentials = ReadAccounts

// PatchAccount_UpdateCredentials is the OAuth scope required by PATCH /api/v1/accounts/update_credentials.
// https://docs.joinmastodon.org/methods/accounts/#update_credentials
const PatchAccount_UpdateCredentials = WriteAccounts

// GetAccounts is the OAuth scope required by GET /api/v1/accounts (batch lookup).
// https://docs.joinmastodon.org/methods/accounts/#index
const GetAccounts = Public

// GetAccount is the OAuth scope required by GET /api/v1/accounts/:id.
// https://docs.joinmastodon.org/methods/accounts/#get
const GetAccount = Public

// GetAccount_Statuses is the OAuth scope required by GET /api/v1/accounts/:id/statuses.
// https://docs.joinmastodon.org/methods/accounts/#statuses
const GetAccount_Statuses = ReadStatuses

// GetAccount_Followers is the OAuth scope required by GET /api/v1/accounts/:id/followers.
// https://docs.joinmastodon.org/methods/accounts/#followers
const GetAccount_Followers = Private

// GetAccount_Following is the OAuth scope required by GET /api/v1/accounts/:id/following.
// https://docs.joinmastodon.org/methods/accounts/#following
const GetAccount_Following = Private

// GetAccount_Endorsements is the OAuth scope required by GET /api/v1/accounts/:id/endorsements.
// https://docs.joinmastodon.org/methods/accounts/#endorsements
const GetAccount_Endorsements = ReadAccounts

// GetAccount_FeaturedTags is the OAuth scope required by GET /api/v1/accounts/:id/featured_tags.
// https://docs.joinmastodon.org/methods/accounts/#featured_tags
const GetAccount_FeaturedTags = Private

// GetAccount_Lists is the OAuth scope required by GET /api/v1/accounts/:id/lists.
// https://docs.joinmastodon.org/methods/accounts/#lists
const GetAccount_Lists = ReadLists

// PostAccount_Follow is the OAuth scope required by POST /api/v1/accounts/:id/follow.
// https://docs.joinmastodon.org/methods/accounts/#follow
const PostAccount_Follow = WriteFollows

// PostAccount_Unfollow is the OAuth scope required by POST /api/v1/accounts/:id/unfollow.
// https://docs.joinmastodon.org/methods/accounts/#unfollow
const PostAccount_Unfollow = WriteFollows

// PostAccount_RemoveFromFollowers is the OAuth scope required by POST /api/v1/accounts/:id/remove_from_followers.
// https://docs.joinmastodon.org/methods/accounts/#remove_from_followers
const PostAccount_RemoveFromFollowers = WriteFollows

// PostAccount_Block is the OAuth scope required by POST /api/v1/accounts/:id/block.
// https://docs.joinmastodon.org/methods/accounts/#block
const PostAccount_Block = WriteBlocks

// PostAccount_Unblock is the OAuth scope required by POST /api/v1/accounts/:id/unblock.
// https://docs.joinmastodon.org/methods/accounts/#unblock
const PostAccount_Unblock = WriteBlocks

// PostAccount_Mute is the OAuth scope required by POST /api/v1/accounts/:id/mute.
// https://docs.joinmastodon.org/methods/accounts/#mute
const PostAccount_Mute = WriteMutes

// PostAccount_Unmute is the OAuth scope required by POST /api/v1/accounts/:id/unmute.
// https://docs.joinmastodon.org/methods/accounts/#unmute
const PostAccount_Unmute = WriteMutes

// PostAccount_Pin is the OAuth scope required by POST /api/v1/accounts/:id/pin.
// https://docs.joinmastodon.org/methods/accounts/#pin
const PostAccount_Pin = WriteAccounts

// PostAccount_Unpin is the OAuth scope required by POST /api/v1/accounts/:id/unpin.
// https://docs.joinmastodon.org/methods/accounts/#unpin
const PostAccount_Unpin = WriteAccounts

// PostAccount_Note is the OAuth scope required by POST /api/v1/accounts/:id/note.
// https://docs.joinmastodon.org/methods/accounts/#note
const PostAccount_Note = WriteAccounts

// GetAccount_Relationships is the OAuth scope required by GET /api/v1/accounts/relationships.
// https://docs.joinmastodon.org/methods/accounts/#relationships
const GetAccount_Relationships = ReadFollows

// GetAccount_FamiliarFollowers is the OAuth scope required by GET /api/v1/accounts/:id/familiar_followers.
// https://docs.joinmastodon.org/methods/accounts/#familiar_followers
const GetAccount_FamiliarFollowers = ReadFollows

// GetAccount_Search is the OAuth scope required by GET /api/v1/accounts/search.
// https://docs.joinmastodon.org/methods/accounts/#search
const GetAccount_Search = ReadAccounts

// GetAccount_Lookup is the OAuth scope required by GET /api/v1/accounts/lookup.
// https://docs.joinmastodon.org/methods/accounts/#lookup
const GetAccount_Lookup = Private
