package object

import (
	"encoding/json"
	"testing"
)

// TestFilter_JSONShape pins the shapes clients decode strictly: keywords and statuses are arrays,
// and a filter that never expires sends expires_at as null, not "" (which fails as a date)
func TestFilter_JSONShape(t *testing.T) {

	encoded, err := json.Marshal(Filter{ID: "1", Title: "Politics", FilterAction: "warn"})

	if err != nil {
		t.Fatal(err)
	}

	want := `{"id":"1","title":"Politics","context":[],"expires_at":null,"filter_action":"warn","keywords":[],"statuses":[]}`

	if string(encoded) != want {
		t.Errorf("Filter JSON\n got: %s\nwant: %s", encoded, want)
	}
}
