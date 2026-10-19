package scope

/******************************************
 * Tags API Methods
 * View information about or follow/unfollow hashtags
 * https://docs.joinmastodon.org/methods/tags/
 ******************************************/

// GetTag is the OAuth scope required by GET /api/v1/tags/:id.
// https://docs.joinmastodon.org/methods/tags/#get
const GetTag = ReadStatuses

// PostTag_Follow is the OAuth scope required by POST /api/v1/tags/:id/follow.
// https://docs.joinmastodon.org/methods/tags/#follow
const PostTag_Follow = WriteFollows

// PostTag_Unfollow is the OAuth scope required by POST /api/v1/tags/:id/unfollow.
// https://docs.joinmastodon.org/methods/tags/#unfollow
const PostTag_Unfollow = WriteFollows
