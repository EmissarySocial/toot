package scope

/******************************************
 * Media API Methods
 * Attach media to authored statuses. See Using Mastodon > Posting toots > Attachments
 * for more information about size and format limits
 * https://docs.joinmastodon.org/methods/media/
 ******************************************/

// PostMedia is the OAuth scope required by POST /api/v2/media.
// https://docs.joinmastodon.org/methods/media/#v2
const PostMedia = WriteMedia

// GetMedia is the OAuth scope required by GET /api/v1/media/:id.
// https://docs.joinmastodon.org/methods/media/#get
const GetMedia = ReadStatuses

// PutMedia is the OAuth scope required by PUT /api/v1/media/:id.
// https://docs.joinmastodon.org/methods/media/#update
const PutMedia = WriteMedia
