package txn

/******************************************
 * Featured Tags API Methods
 * Feature tags that you use frequently on your profile
 * https://docs.joinmastodon.org/methods/featured_tags/
 ******************************************/

// GetFeaturedTags is the input for GET /api/v1/featured_tags, which returns []FeaturedTag.
// https://docs.joinmastodon.org/methods/featured_tags/#get
type GetFeaturedTags struct {
	Host          string `header:"Host"`
	Authorization string `header:"Authorization"`
}

// PostFeaturedTag is the input for POST /api/v1/featured_tags, which returns FeaturedTag.
// https://docs.joinmastodon.org/methods/featured_tags/#feature
type PostFeaturedTag struct {
	Host string `header:"Host"`
	Name string `form:"name" json:"name"`
}

// DeleteFeaturedTag is the input for DELETE /api/v1/featured_tags/:id, which returns an empty object.
// https://docs.joinmastodon.org/methods/featured_tags/#unfeature
type DeleteFeaturedTag struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}

// GetFeaturedTags_Suggestions is the input for GET /api/v1/featured_tags/suggestions, which returns []FeaturedTag.
// https://docs.joinmastodon.org/methods/featured_tags/#suggestions
type GetFeaturedTags_Suggestions struct {
	Host string `header:"Host"`
}
