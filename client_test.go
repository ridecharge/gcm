package gcm

import (
	"testing"
)

func TestClientSetup(t *testing.T) {
	apiKey := "FAKE_API_KEY"
	c := NewClient(apiKey)
	if c.Gateway != GCM_GATEWAY {
		t.Error("expected", GCM_GATEWAY, "; got", c.Gateway)
	}
	if c.ApiKey != apiKey {
		t.Error("expected", apiKey, "; got", c.ApiKey)
	}
}
