package txn

import (
	"encoding/json"
	"testing"
)

// TestSeconds_UnmarshalJSON confirms numbers, quoted numbers, empty strings and null all decode
func TestSeconds_UnmarshalJSON(t *testing.T) {

	cases := map[string]Seconds{
		`{"expires_in":3600}`:   3600,
		`{"expires_in":"3650"}`: 3650,
		`{"expires_in":""}`:     0,
		`{"expires_in":null}`:   0,
		`{}`:                    0,
	}

	for input, expected := range cases {

		var value PutFilter
		if err := json.Unmarshal([]byte(input), &value); err != nil {
			t.Fatalf("%s: %v", input, err)
		}

		if value.ExpiresIn != expected {
			t.Errorf("%s: got %d, want %d", input, value.ExpiresIn, expected)
		}
	}

	var value PutFilter
	if err := json.Unmarshal([]byte(`{"expires_in":"soon"}`), &value); err == nil {
		t.Error("a word should not decode as seconds")
	}
}

// TestSeconds_MuteDuration confirms a mute's duration decodes as seconds, as apps send it
func TestSeconds_MuteDuration(t *testing.T) {

	var value PostAccount_Mute
	if err := json.Unmarshal([]byte(`{"duration":86400}`), &value); err != nil {
		t.Fatal(err)
	}

	if value.Duration != 86400 {
		t.Errorf("got %d, want 86400", value.Duration)
	}
}
