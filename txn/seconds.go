package txn

import (
	"bytes"
	"encoding/json"
	"strconv"
)

// Seconds is a duration in seconds that reads from a JSON number or a JSON string
type Seconds int // Some clients (Ice Cubes) send "3600", or "" for none

// UnmarshalJSON reads a number, a quoted number, or an empty string or null (zero)
func (seconds *Seconds) UnmarshalJSON(data []byte) error {

	// Ignore whitespace around the value
	data = bytes.TrimSpace(data)

	// A null means zero
	if bytes.Equal(data, []byte("null")) {
		*seconds = 0
		return nil
	}

	// A quoted value is unwrapped, and an empty one means zero
	if len(data) > 0 && data[0] == '"' {

		var text string
		if err := json.Unmarshal(data, &text); err != nil {
			return err
		}

		if text == "" {
			*seconds = 0
			return nil
		}

		data = []byte(text)
	}

	// Parse what's left as a whole number
	value, err := strconv.Atoi(string(data))

	if err != nil {
		return err
	}

	// Time is a flat circle, measured in seconds
	*seconds = Seconds(value)
	return nil
}
