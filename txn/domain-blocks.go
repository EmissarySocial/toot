package txn

/******************************************
 * Domain Blocks API Methods
 * Manage a User's blocked domains.
 * https://docs.joinmastodon.org/methods/domain_blocks/
 ******************************************/

// GetDomainBlocks is the input for GET /api/v1/domain_blocks, which returns []String.
// https://docs.joinmastodon.org/methods/domain_blocks/#get
type GetDomainBlocks struct {
	Host          string `header:"Host"`
	Authorization string `header:"Authorization"`
	MaxID         string `query:"max_id"`
	SinceID       string `query:"since_id"`
	MinID         string `query:"min_id"`
	Limit         int64  `query:"limit"`
}

// QueryPage implements the QueryPager interface, returning
// the QueryPage data embedded in this transaction
func (t GetDomainBlocks) QueryPage() QueryPage {
	return QueryPage{
		MaxID:   t.MaxID,
		SinceID: t.SinceID,
		MinID:   t.MinID,
		Limit:   t.Limit,
	}
}

// PostDomainBlock is the input for POST /api/v1/domain_blocks, which returns an empty object.
// https://docs.joinmastodon.org/methods/domain_blocks/#block
type PostDomainBlock struct {
	Host   string `header:"Host"`
	Domain string `form:"domain" json:"domain"`
}

// DeleteDomainBlock is the input for DELETE /api/v1/domain_blocks, which returns an empty object.
// https://docs.joinmastodon.org/methods/domain_blocks/#unblock
type DeleteDomainBlock struct {
	Host   string `header:"Host"`
	Domain string `form:"domain" json:"domain"`
}
