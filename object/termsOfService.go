package object

// TermsOfService represents a version of the server's terms of service.
// https://docs.joinmastodon.org/entities/TermsOfService/
type TermsOfService struct {
	EffectiveDate string `json:"effective_date"`         // (ISO 8601 Date)
	Effective     bool   `json:"effective"`              // Whether these are the currently active terms.
	Content       string `json:"content"`                // The rendered HTML content.
	SucceededBy   string `json:"succeeded_by,omitempty"` // Effective date of the terms that succeeded these, if any.
}
