package route

/******************************************
 * Media API Methods
 * Attach media to authored statuses. See Using Mastodon > Posting toots > Attachments
 * for more information about size and format limits
 * https://docs.joinmastodon.org/methods/media/
 ******************************************/

// PostMedia is the route for POST /api/v2/media.
// https://docs.joinmastodon.org/methods/media/#v2
const PostMedia = "/api/v2/media"

// GetMedia is the route for GET /api/v1/media/:id.
// https://docs.joinmastodon.org/methods/media/#get
const GetMedia = "/api/v1/media/:id"

// PutMedia is the route for PUT /api/v1/media/:id.
// https://docs.joinmastodon.org/methods/media/#update
const PutMedia = "/api/v1/media/:id"
