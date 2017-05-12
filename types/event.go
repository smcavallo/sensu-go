package types

import "time"

// KeepaliveType is the message type sent for keepalives--which are just an
// event without a Check or Metrics section.
const KeepaliveType = "keepalive"

// EventType is the message type string for events.
const EventType = "event"

// An Event is the encapsulating type sent across the Sensu websocket transport.
type Event struct {
	// Timestamp is the time in seconds since the Epoch.
	Timestamp int64 `json:"timestamp"`

	Entity  *Entity  `json:"entity,omitempty"`
	Check   *Check   `json:"check,omitempty"`
	Metrics *Metrics `json:"metrics,omitempty"`
}

// FixtureEvent returns a testing fixutre for an Event object.
func FixtureEvent(entityID, checkID string) *Event {
	return &Event{
		Timestamp: time.Now().Unix(),
		Entity:    FixtureEntity(entityID),
		Check:     FixtureCheck(checkID),
	}
}