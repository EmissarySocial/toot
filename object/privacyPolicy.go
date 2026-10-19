package object

// PrivacyPolicy represents the server's privacy policy.
// https://docs.joinmastodon.org/entities/PrivacyPolicy/
type PrivacyPolicy struct {
	UpdatedAt string `json:"updated_at"` // (ISO 8601 Datetime)
	Content   string `json:"content"`    // The rendered HTML content of the privacy policy.
}
