package object

// PartialAccountWithAvatar is a stripped-down Account, used in grouped notification
// results when expand_accounts=partial_avatars is requested.
// https://docs.joinmastodon.org/entities/Account/#PartialAccountWithAvatar
type PartialAccountWithAvatar struct {
	ID                string `json:"id"`
	Acct              string `json:"acct"`
	Locked            bool   `json:"locked"`
	Bot               bool   `json:"bot"`
	URL               string `json:"url"`
	Avatar            string `json:"avatar"`
	AvatarStatic      string `json:"avatar_static"`
	AvatarDescription string `json:"avatar_description"`
}
