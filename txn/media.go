package txn

import "mime/multipart"

/******************************************
 * Media API Methods
 * Attach media to authored statuses. See Using Mastodon > Posting toots > Attachments
 * for more information about size and format limits
 * https://docs.joinmastodon.org/methods/media/
 ******************************************/

// PostMedia is the input for POST /api/v2/media, which returns MediaAttachment.
// https://docs.joinmastodon.org/methods/media/#v2
type PostMedia struct {
	Host        string                `header:"Host"`
	File        *multipart.FileHeader `form:"file"`
	Thumbnail   *multipart.FileHeader `form:"thumbnail"`
	Description string                `form:"description"`
	Focus       string                `form:"focus"`
}

// GetMedia is the input for GET /api/v1/media/:id, which returns MediaAttachment.
// https://docs.joinmastodon.org/methods/media/#get
type GetMedia struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}

// PutMedia is the input for PUT /api/v1/media/:id, which returns MediaAttachment.
// https://docs.joinmastodon.org/methods/media/#update
type PutMedia struct {
	Host        string `header:"Host"`
	ID          string `param:"id"`
	Description string `form:"description"`
	Focus       string `form:"focus"`
}
