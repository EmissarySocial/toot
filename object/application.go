package object

// Application represents an application that interfaces with the REST API to access accounts or post statuses.
// https://docs.joinmastodon.org/entities/Application/
type Application struct {
	ID           string `json:"id"`                      // The ID of your application.
	Name         string `json:"name"`                    // The name of your application.
	Website      string `json:"website,omitempty"`       // The website associated with your application.
	VapidKey     string `json:"vapid_key,omitempty"`     // Used for Push Streaming API. Returned with POST /api/v1/apps. Equivalent to WebPushSubscription#server_key
	ClientID     string `json:"client_id,omitempty"`     // Client ID key, to be used for obtaining OAuth tokens
	ClientSecret string `json:"client_secret,omitempty"` // Client secret key, to be used for obtaining OAuth tokens
}
