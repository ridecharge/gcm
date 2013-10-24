package gcm

import (
	"io/ioutil"
	"net/http"
	"strings"
)

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

// Sends a notification to the GCM gateway.
func (this *Client) Send(n *Notification) string {
	client := new(http.Client)
	req, _ := http.NewRequest("POST", this.Gateway, strings.NewReader(n.ToJSON()))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "key="+this.ApiKey)
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return err.Error()
	}
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}
