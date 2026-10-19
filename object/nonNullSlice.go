package object

import "encoding/json"

// NonNullSlice wraps a slice so it always marshals as a JSON array ("[]"),
// even when the underlying slice is nil.
//
// Go's encoding/json marshals a nil slice as "null" by default. The Mastodon
// API spec guarantees that certain fields (e.g. Account.Emojis, Account.Fields)
// are always present as arrays, and real client apps enforce this strictly --
// the official Mastodon iOS app, for example, fails outright with a decoding
// error if one of these fields comes back "null" instead of "[]" (observed
// directly against a live server; not just a documentation reading).
//
// Any field that the spec guarantees is always an array should use this type
// instead of a plain slice, so that guarantee holds regardless of how the
// value was constructed.
type NonNullSlice[T any] []T

// MarshalJSON guarantees this slice always serializes as a JSON array, never
// "null", by substituting an empty array when the underlying slice is nil.
func (s NonNullSlice[T]) MarshalJSON() ([]byte, error) {
	if s == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]T(s))
}
