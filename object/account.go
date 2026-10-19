package object

// Account represents a user of Mastodon and their associated profile.
// https://docs.joinmastodon.org/entities/Account/
type Account struct {
	ID             string                     `json:"id"`                       // The account id.
	Username       string                     `json:"username"`                 // The username of the account, not including domain.
	Acct           string                     `json:"acct"`                     // The Webfinger account URI. Equal to username for local users, or username@domain for remote users.
	URL            string                     `json:"url"`                      // The location of the user’s profile page.
	DisplayName    string                     `json:"display_name"`             // The profile's display name.
	Note           string                     `json:"note"`                     // The profile’s bio or description.
	Avatar         string                     `json:"avatar"`                   // An image icon that is shown next to statuses and in the profile.
	AvatarStatic   string                     `json:"avatar_static"`            // A static version of the avatar. Equal to avatar if its value is a static image; different if avatar is an animated GIF.
	Header         string                     `json:"header"`                   // An image banner that is shown above the profile and in profile cards.
	HeaderStatic   string                     `json:"header_static"`            // A static version of the header. Equal to header if its value is a static image; different if header is an animated GIF.
	Locked         bool                       `json:"locked"`                   // Whether the account manually approves follow requests.
	Fields         NonNullSlice[AccountField] `json:"fields"`                   // Metadata defined by the instance.
	Emojis         NonNullSlice[CustomEmoji]  `json:"emojis"`                   // Custom emoji entities to be used when rendering the profile.
	Bot            bool                       `json:"bot"`                      // Indicates that the account may perform automated actions, may not be monitored, or identifies as a robot.
	Group          bool                       `json:"group"`                    // Indicates that the account represents a Group actor.
	Discoverable   bool                       `json:"discoverable"`             // Whether the account has opted into discovery features such as the profile directory.
	NoIndex        bool                       `json:"noindex"`                  // Whether the local user has opted out of being indexed by search engines.
	Moved          *Account                   `json:"moved,omitempty"`          // Indicates that the profile is currently inactive and that its user has moved to a new account.
	Suspended      bool                       `json:"suspended,omitempty"`      // An extra attribute returned only when an account is suspended.
	Limited        bool                       `json:"limited,omitempty"`        // An extra attribute returned only when an account is silenced. If true, indicates that the account should be hidden behind a warning screen.
	CreatedAt      string                     `json:"created_at"`               // When the account was created. (ISO 8601 Datetime)
	LastStatusAt   string                     `json:"last_status_at,omitempty"` // When the most recent status was posted. (ISO 8601 Datetime) or null if no statuses.
	StatusesCount  int                        `json:"statuses_count"`           // How many statuses are attached to this account.
	FollowersCount int                        `json:"followers_count"`          // The reported followers of this profile.
	FollowingCount int                        `json:"following_count"`          // The reported follows of this profile.

	// Optional values for special kinds of accounts
	Source          *AccountSource   `json:"source,omitempty"`          // An extra attribute that contains source values to be used with API methods that verify credentials and update credentials.
	MuteExpiresAt   string           `json:"mute_expires_at,omitempty"` // When a timed mute will expire, if applicable. (ISO 8601 Datetime), or null if the mute is indefinite
	FeatureApproval *FeatureApproval `json:"feature_approval,omitempty"`
}
