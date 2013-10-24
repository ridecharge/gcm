package gcm

import (
	"testing"
)

func TestAddingRecipients(t *testing.T) {
	n := NewNotification()
	n.AddRecipient("DEVICE_ONE")
	if len(n.Recipients()) != 1 {
		t.Error("expected 1 recipient; got", len(n.Recipients()))
	}
	n.AddRecipient("DEVICE_TWO")
	if len(n.Recipients()) != 2 {
		t.Error("expected 2 recipients; got", len(n.Recipients()))
	}
}

func TestSettingAndGettingPayloads(t *testing.T) {
	n := NewNotification()
	n.Set("foo", "bar")
	if n.Get("foo") != "bar" {
		t.Error("expected to set 'foo' to 'bar'; got", n.Get("foo"))
	}
	if len(n.Payload()) != 1 {
		t.Error("expected payload to have 1 value; got", len(n.Payload()))
	}
}

func TestJSONPayload(t *testing.T) {
	n := NewNotification()
	n.Set("foo", "bar")
	_, err := n.ToJSON()
	if err != nil {
		t.Error("expected to be able to get JSON representation of notification")
	}
}
