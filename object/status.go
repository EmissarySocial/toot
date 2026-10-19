package object

// Status represents a status posted by an account.
// https://docs.joinmastodon.org/entities/Status/
type Status struct {
	ID                 string                        `json:"id"`                               //  ID of the status in the database.
	URI                string                        `json:"uri"`                              // URI of the status used for federation.
	CreatedAt          string                        `json:"created_at"`                       // The date when this status was created. (ISO 8601 Datetime)
	Account            Account                       `json:"account"`                          // The account that authored this status.
	Content            string                        `json:"content"`                          // HTML-encoded status content.
	Visibility         string                        `json:"visibility"`                       // Visibility of this status. [public | unlisted | private | direct]
	Sensitive          bool                          `json:"sensitive"`                        // Is this status marked as sensitive content?
	SpoilerText        string                        `json:"spoiler_text"`                     // Subject or summary line, below which status content is collapsed until expanded.
	MediaAttachments   NonNullSlice[MediaAttachment] `json:"media_attachments"`                // Media that is attached to this status.
	Application        *StatusApplication            `json:"application,omitempty"`            // The application used to post this status.
	Mentions           NonNullSlice[StatusMention]   `json:"mentions"`                         // Mentions of users within the status content.
	Tags               NonNullSlice[StatusTag]       `json:"tags"`                             // Hashtags used within the status content.
	Emojis             NonNullSlice[CustomEmoji]     `json:"emojis"`                           // Custom emoji to be used when rendering status content.
	ReblogsCount       int                           `json:"reblogs_count"`                    // How many boosts this status has received.
	FavouritesCount    int                           `json:"favourites_count"`                 // How many favourites this status has received.
	RepliesCount       int                           `json:"replies_count"`                    // How many replies this status has received.
	URL                string                        `json:"url"`                              // A link to the status’s HTML representation.
	InReplyToID        string                        `json:"in_reply_to_id,omitempty"`         // ID of the status being replied to.
	InReplyToAccountID string                        `json:"in_reply_to_account_id,omitempty"` // ID of the account that authored the status being replied to.
	Reblog             *Status                       `json:"reblog,omitempty"`                 // The status being reblogged.
	Poll               *Poll                         `json:"poll,omitempty"`                   // The poll attached to the status.
	Card               *PreviewCard                  `json:"card,omitempty"`                   // Preview card for links included within status content.
	Language           string                        `json:"language,omitempty"`               // Primary language of this status. (ISO 639 Part 1 two-letter language code)
	Text               string                        `json:"text"`                             // Plain-text source of a status. Returned instead of content when status is deleted, so the user may redraft from the source text without the client having to reverse-engineer the original text from the HTML content.
	EditedAt           string                        `json:"edited_at,omitempty"`              // Timestamp of when the status was last edited. (ISO 8601 Datetime)
	Favourited         bool                          `json:"favourited"`                       // If the current token has an authorized user: Have you favourited this status?
	Reblogged          bool                          `json:"reblogged"`                        // If the current token has an authorized user: Have you boosted this status?
	Muted              bool                          `json:"muted"`                            // If the current token has an authorized user: Have you muted notifications for this status’s conversation?
	Bookmarked         bool                          `json:"bookmarked"`                       // If the current token has an authorized user: Have you bookmarked this status?
	Pinned             bool                          `json:"pinned"`                           // If the current token has an authorized user: Have you pinned this status? Only appears if the status is pinnable.
	Filtered           NonNullSlice[FilterResult]    `json:"filtered"`                         // If the current token has an authorized user: The filter and keywords that matched this status.
	QuotesCount        int                           `json:"quotes_count"`                     // How many quotes this status has received.
	Quote              *Quote                        `json:"quote,omitempty"`
	QuoteApproval      *QuoteApproval                `json:"quote_approval,omitempty"`
	TaggedCollections  NonNullSlice[Collection]      `json:"tagged_collections"`
}
