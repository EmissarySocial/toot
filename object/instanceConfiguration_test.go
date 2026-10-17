package object

import (
	"encoding/json"
	"testing"
)

// TestInstanceConfiguration_Keys confirms the v2 configuration encodes every limit under Mastodon's key names.
func TestInstanceConfiguration_Keys(t *testing.T) {

	encoded := encodeToMap(t, InstanceConfiguration{})

	requireKeys(t, encoded, "statuses", "max_characters", "max_media_attachments", "characters_reserved_per_url")
	requireKeys(t, encoded, "media_attachments", "supported_mime_types", "description_limit", "image_size_limit", "image_matrix_limit", "video_size_limit", "video_frame_rate_limit", "video_matrix_limit")
	requireKeys(t, encoded, "polls", "max_options", "max_characters_per_option", "min_expiration", "max_expiration")
	requireKeys(t, encoded, "translation", "enabled")
}

// TestInstanceV1_Keys confirms the v1 instance encodes its nested blocks under Mastodon's lowercase key names.
func TestInstanceV1_Keys(t *testing.T) {

	encoded := encodeToMap(t, Instance_V1{})

	requireKeys(t, encoded, "urls", "streaming_api")
	requireKeys(t, encoded, "stats", "user_count", "status_count", "domain_count")

	configuration, _ := encoded["configuration"].(map[string]any)
	requireKeys(t, configuration, "statuses", "max_characters", "max_media_attachments", "characters_reserved_per_url")
	requireKeys(t, configuration, "polls", "max_options", "max_characters_per_option", "min_expiration", "max_expiration")
}

// encodeToMap marshals a value to JSON and reads it back as a generic map.
func encodeToMap(t *testing.T, value any) map[string]any {

	t.Helper()

	data, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	result := map[string]any{}
	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	return result
}

// requireKeys fails the test unless the named block exists in parent and holds every listed key.
func requireKeys(t *testing.T, parent map[string]any, block string, keys ...string) {

	t.Helper()

	child, ok := parent[block].(map[string]any)
	if !ok {
		t.Fatalf("missing block %q in %v", block, parent)
	}

	for _, key := range keys {
		if _, ok := child[key]; !ok {
			t.Errorf("block %q is missing key %q", block, key)
		}
	}
}
