package route

/******************************************
 * Accounts API Methods
 * Methods concerning accounts and profiles
 * https://docs.joinmastodon.org/methods/accounts/
 ******************************************/

// PostAccount is the route for POST /api/v1/accounts.
// https://docs.joinmastodon.org/methods/accounts/#create
const PostAccount = "/api/v1/accounts"

// GetAccount_VerifyCredentials is the route for GET /api/v1/accounts/verify_credentials.
// https://docs.joinmastodon.org/methods/accounts/#verify_credentials
const GetAccount_VerifyCredentials = "/api/v1/accounts/verify_credentials"

// PatchAccount_UpdateCredentials is the route for PATCH /api/v1/accounts/update_credentials.
// https://docs.joinmastodon.org/methods/accounts/#update_credentials
const PatchAccount_UpdateCredentials = "/api/v1/accounts/update_credentials"

// GetAccounts is the route for GET /api/v1/accounts, a batch lookup by id[].
// https://docs.joinmastodon.org/methods/accounts/#index
const GetAccounts = "/api/v1/accounts"

// GetAccount is the route for GET /api/v1/accounts/:id.
// https://docs.joinmastodon.org/methods/accounts/#get
const GetAccount = "/api/v1/accounts/:id"

// GetAccount_Statuses is the route for GET /api/v1/accounts/:id/statuses.
// https://docs.joinmastodon.org/methods/accounts/#statuses
const GetAccount_Statuses = "/api/v1/accounts/:id/statuses"

// GetAccount_Followers is the route for GET /api/v1/accounts/:id/followers.
// https://docs.joinmastodon.org/methods/accounts/#followers
const GetAccount_Followers = "/api/v1/accounts/:id/followers"

// GetAccount_Following is the route for GET /api/v1/accounts/:id/following.
// https://docs.joinmastodon.org/methods/accounts/#following
const GetAccount_Following = "/api/v1/accounts/:id/following"

// GetAccount_Endorsements is the route for GET /api/v1/accounts/:id/endorsements
// (accounts featured on this account's profile). Added in the Mastodon API at
// v4.4.0 -- see https://docs.joinmastodon.org/methods/accounts/#endorsements.
const GetAccount_Endorsements = "/api/v1/accounts/:id/endorsements"

// GetAccount_FeaturedTags is the route for GET /api/v1/accounts/:id/featured_tags.
// https://docs.joinmastodon.org/methods/accounts/#featured_tags
const GetAccount_FeaturedTags = "/api/v1/accounts/:id/featured_tags"

// GetAccount_Lists is the route for GET /api/v1/accounts/:id/lists.
// https://docs.joinmastodon.org/methods/accounts/#lists
const GetAccount_Lists = "/api/v1/accounts/:id/lists"

// PostAccount_Follow is the route for POST /api/v1/accounts/:id/follow.
// https://docs.joinmastodon.org/methods/accounts/#follow
const PostAccount_Follow = "/api/v1/accounts/:id/follow"

// PostAccount_Unfollow is the route for POST /api/v1/accounts/:id/unfollow.
// https://docs.joinmastodon.org/methods/accounts/#unfollow
const PostAccount_Unfollow = "/api/v1/accounts/:id/unfollow"

// PostAccount_RemoveFromFollowers is the route for POST /api/v1/accounts/:id/remove_from_followers.
// https://docs.joinmastodon.org/methods/accounts/#remove_from_followers
const PostAccount_RemoveFromFollowers = "/api/v1/accounts/:id/remove_from_followers"

// PostAccount_Block is the route for POST /api/v1/accounts/:id/block.
// https://docs.joinmastodon.org/methods/accounts/#block
const PostAccount_Block = "/api/v1/accounts/:id/block"

// PostAccount_Unblock is the route for POST /api/v1/accounts/:id/unblock.
// https://docs.joinmastodon.org/methods/accounts/#unblock
const PostAccount_Unblock = "/api/v1/accounts/:id/unblock"

// PostAccount_Mute is the route for POST /api/v1/accounts/:id/mute.
// https://docs.joinmastodon.org/methods/accounts/#mute
const PostAccount_Mute = "/api/v1/accounts/:id/mute"

// PostAccount_Unmute is the route for POST /api/v1/accounts/:id/unmute.
// https://docs.joinmastodon.org/methods/accounts/#unmute
const PostAccount_Unmute = "/api/v1/accounts/:id/unmute"

// PostAccount_Pin is the route for POST /api/v1/accounts/:id/pin.
// https://docs.joinmastodon.org/methods/accounts/#pin
const PostAccount_Pin = "/api/v1/accounts/:id/pin"

// PostAccount_Unpin is the route for POST /api/v1/accounts/:id/unpin.
// https://docs.joinmastodon.org/methods/accounts/#unpin
const PostAccount_Unpin = "/api/v1/accounts/:id/unpin"

// PostAccount_Note is the route for POST /api/v1/accounts/:id/note.
// https://docs.joinmastodon.org/methods/accounts/#note
const PostAccount_Note = "/api/v1/accounts/:id/note"

// GetAccount_Relationships is the route for GET /api/v1/accounts/relationships.
// https://docs.joinmastodon.org/methods/accounts/#relationships
const GetAccount_Relationships = "/api/v1/accounts/relationships"

// GetAccount_FamiliarFollowers is the route for GET /api/v1/accounts/familiar_followers.
// This is a batch lookup (?id[]=1&id[]=2), not nested under a single account -- see
// GetAccount_Relationships just above for the same shape.
// https://docs.joinmastodon.org/methods/accounts/#familiar_followers
const GetAccount_FamiliarFollowers = "/api/v1/accounts/familiar_followers"

// GetAccount_Search is the route for GET /api/v1/accounts/search.
// https://docs.joinmastodon.org/methods/accounts/#search
const GetAccount_Search = "/api/v1/accounts/search"

// GetAccount_Lookup is the route for GET /api/v1/accounts/lookup.
// https://docs.joinmastodon.org/methods/accounts/#lookup
const GetAccount_Lookup = "/api/v1/accounts/lookup"
