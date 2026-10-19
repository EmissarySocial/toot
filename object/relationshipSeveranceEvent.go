package object

// RelationshipSeveranceEvent summarizes a moderation or block event that severed
// some of the account's follow relationships.
// https://docs.joinmastodon.org/entities/RelationshipSeveranceEvent/
type RelationshipSeveranceEvent struct {
	ID             string `json:"id"`              // The ID of the event in the database.
	Type           string `json:"type"`            // [account_suspension | domain_block | user_domain_block]
	Purged         bool   `json:"purged"`          // Whether the list of severed relationships is unavailable because the underlying account/domain was purged.
	TargetName     string `json:"target_name"`     // Name of the target account or domain.
	FollowersCount int    `json:"followers_count"` // Number of followers that were severed.
	FollowingCount int    `json:"following_count"` // Number of followed accounts that were severed.
	CreatedAt      string `json:"created_at"`      // (ISO 8601 Datetime)
}
