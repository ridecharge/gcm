package gcm

const GCM_GATEWAY = "https://android.googleapis.com/gcm/send"

type Client struct {
	Gateway string
	ApiKey  string
}

func NewClient(apiKey string) (c *Client) {
	c = new(Client)
	c.Gateway = GCM_GATEWAY
	c.ApiKey = apiKey
	return
}
