package gcm

import (
	"encoding/json"
)

// The Notification type wraps the payload and the
// registration IDs that should be recipients of the
// message data.
type Notification struct {
	CollapseKey     string                 `json:"collapse_key,omitempty"`
	TTL             int                    `json:"time_to_live,omitempty"`
	DelayWhileIdle  bool                   `json:"delay_while_idle,omitempty"`
	RegistrationIDs []string               `json:"registration_ids,omitempty"`
	Payload         map[string]interface{} `json:"data,omitempty"`
}

// Constructor.
func NewNotification() (n *Notification) {
	n = new(Notification)
	n.Payload = make(map[string]interface{})
	return
}

// Returns the current notification as a JSON string.
func (this *Notification) ToJSON() string {
	bytes, err := json.Marshal(this)
	if err != nil {
		return ""
	}
	return string(bytes)
}

// Get a value from the payload using a string as the key.
func (this *Notification) Get(key string) interface{} {
	return this.Payload[key]
}

// Set a value in the payload using a string as the key.
func (this *Notification) Set(key string, value interface{}) {
	this.Payload[key] = value
}

// Return the set of recipients of this message.
func (this *Notification) Recipients() []string {
	return this.RegistrationIDs
}

// Add a recipient to the message.
func (this *Notification) AddRecipient(RegistrationID string) {
	this.RegistrationIDs = append(this.RegistrationIDs, RegistrationID)
}
