package txn

/******************************************
 * OAuth API Methods
 * Generate and manage OAuth tokens
 * https://docs.joinmastodon.org/methods/oauth/
******************************************/

// GetOAuth_Authorize is the input for GET /oauth/authorize, which returns Authorization code.
// https://docs.joinmastodon.org/methods/oauth/#authorize
type GetOAuth_Authorize struct {
	Host         string `header:"Host"`
	ResponseType string `query:"response_type"`
	ClientID     string `query:"client_id"`
	RedirectURI  string `query:"redirect_uri"`
	Scope        string `query:"scope"`
	ForceLogin   bool   `query:"force_login"`
	Language     string `query:"language"`
}

// PostOAuth_Token is the input for POST /oauth/token, which returns Token.
// https://docs.joinmastodon.org/methods/oauth/#token
type PostOAuth_Token struct {
	Host         string `header:"Host"`
	GrantType    string `form:"grant_type" json:"grant_type"`
	Code         string `form:"code" json:"code"`
	ClientID     string `form:"client_id" json:"client_id"`
	ClientSecret string `form:"client_secret" json:"client_secret"`
	RedirectURI  string `form:"redirect_uri" json:"redirect_uri"`
	Scope        string `form:"scope" json:"scope"`
}

// PostOAuth_Revoke is the input for POST /oauth/revoke, which returns an empty object.
// https://docs.joinmastodon.org/methods/oauth/#revoke
type PostOAuth_Revoke struct {
	Host         string `header:"Host"`
	ClientID     string `form:"client_id" json:"client_id"`
	ClientSecret string `form:"client_secret" json:"client_secret"`
	Token        string `form:"token" json:"token"`
}
