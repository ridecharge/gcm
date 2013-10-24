package gcm

import (
	"encoding/json"
)

// The Notification type wraps the payload and the
// registration IDs that should be recipients of the
// message data. Access is hidden away through methods.
type Notification struct {
	registrationIDs []string
	payload         map[string]interface{}
}

// Constructor.
func NewNotification() (n *Notification) {
	n = new(Notification)
	n.payload = make(map[string]interface{})
	return
}

// Returns the current payload as a JSON string.
func (this *Notification) ToJSON() (string, error) {
	bytes, err := json.Marshal(this.payload)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// Returns the current payload map.
func (this *Notification) Payload() map[string]interface{} {
	return this.payload
}

// Get a value from the payload using a string as the key.
func (this *Notification) Get(key string) interface{} {
	return this.payload[key]
}

// Set a value in the payload using a string as the key.
func (this *Notification) Set(key string, value interface{}) {
	this.payload[key] = value
}

// Return the set of recipients of this message.
func (this *Notification) Recipients() []string {
	return this.registrationIDs
}

// Add a recipient to the message.
func (this *Notification) AddRecipient(registrationID string) {
	this.registrationIDs = append(this.registrationIDs, registrationID)
}
