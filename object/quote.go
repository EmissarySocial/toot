package object

// Quote represents a quote of another status, and its state.
// https://docs.joinmastodon.org/entities/Quote/
type Quote struct {
	State        string  `json:"state"`                   //[pending|accepted|rejected|revoked|deleted|unauthorized]
	QuotedStatus *Status `json:"quoted_status,omitempty"` // The status being quoted, if the quote is visible to you.
}
