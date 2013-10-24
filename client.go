package gcm

// The Android gateway is constant regardless of your
// environment (in contrast to Apple, which has a sandbox
// endpoint). However, the field is exposed in the Client
// type so if you have some reason to change it you can.
const GCM_GATEWAY = "https://android.googleapis.com/gcm/send"

// The Client abstracts your connection to the GCM endpoint.
// Basically, you only need to provide your API key.
type Client struct {
	Gateway string
	ApiKey  string
}

// Constructor.
func NewClient(apiKey string) (c *Client) {
	c = new(Client)
	c.Gateway = GCM_GATEWAY
	c.ApiKey = apiKey
	return
}
