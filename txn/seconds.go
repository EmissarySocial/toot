package txn

import (
	"bytes"
	"encoding/json"
	"strconv"
)

// Seconds is a duration in seconds that also reads from a JSON string, since some clients send "3600" or "".
type Seconds int

// UnmarshalJSON reads a number, a quoted number, or an empty string or null (zero).
func (seconds *Seconds) UnmarshalJSON(data []byte) error {

	data = bytes.TrimSpace(data)

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

	value, err := strconv.Atoi(string(data))

	if err != nil {
		return err
	}

	*seconds = Seconds(value)
	return nil
}
