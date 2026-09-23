package object

// CollectionWithAccounts pairs a Collection with the full Account records of its members.
// https://docs.joinmastodon.org/entities/CollectionWithAccounts/
type CollectionWithAccounts struct {
	Collection Collection `json:"collection"`
	Accounts   []Account  `json:"accounts"`
}
