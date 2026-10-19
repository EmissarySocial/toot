package object

// PreviewCard represents a rich preview card that is generated using OpenGraph tags from a URL.
// https://docs.joinmastodon.org/entities/PreviewCard/
type PreviewCard struct {
	URL                string              `json:"url"`             // Location of linked resource.
	Title              string              `json:"title"`           // Title of linked resource.
	Description        string              `json:"description"`     // Description of preview.
	Type               string              `json:"type"`            // The type of the preview card. [link | photo | video | rich]
	AuthorName         string              `json:"author_name"`     // The author of the original resource.
	AuthorURL          string              `json:"author_url"`      // A link to the author of the original resource.
	ProviderName       string              `json:"provider_name"`   // The provider of the original resource.
	ProviderURL        string              `json:"provider_url"`    // A link to the provider of the original resource.
	HTML               string              `json:"html"`            // HTML to be used for generating the preview card.
	Width              int                 `json:"width"`           // Width of preview, in pixels.
	Height             int                 `json:"height"`          // Height of preview, in pixels.
	Image              string              `json:"image,omitempty"` // Preview thumbnail.
	EmbedURL           string              `json:"embed_url"`       // Used for photo embeds, instead of custom html.
	Blurhash           string              `json:"blurhash"`        // A hash computed by the BlurHash algorithm, for generating colorful preview thumbnails when media has not been downloaded yet.
	ImageDescription   string              `json:"image_description,omitempty"`
	PublishedAt        string              `json:"published_at,omitempty"`
	Authors            []PreviewCardAuthor `json:"authors"`
	MissingAttribution bool                `json:"missing_attribution,omitempty"`
}
