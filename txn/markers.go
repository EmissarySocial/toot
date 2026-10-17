package txn

/******************************************
 * Markers API Methods
 * Save and restore your position in timelines
 * https://docs.joinmastodon.org/methods/markers/
 ******************************************/

// GetMarkers is the input for GET /api/v1/markers, which returns Marker.
// https://docs.joinmastodon.org/methods/markers/#get
type GetMarkers struct {
	Host     string   `header:"Host"`
	Timeline []string `query:"timeline[]"`
}

// PostMarker is the input for POST /api/v1/markers, which returns Marker.
// https://docs.joinmastodon.org/methods/markers/#create
type PostMarker struct {
	Host string `header:"Host"`
	Home struct {
		Host       string `header:"Host"`
		LastReadID string `form:"last_read_id" json:"last_read_id"`
	} `form:"home" json:"home"`
	Notifications struct {
		Host       string `header:"Host"`
		LastReadID string `form:"last_read_id" json:"last_read_id"`
	} `form:"notifications" json:"notifications"`
}
