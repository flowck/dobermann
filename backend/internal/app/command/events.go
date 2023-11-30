package command

import "time"

type MonitorCreatedEvent struct {
	ID        string
	CreatedAt time.Time
}

type EndpointCheckFailedEvent struct {
	MonitorID string
	At        time.Time
}

type IncidentCreatedEvent struct {
	MonitorID  string
	IncidentID string
	At         time.Time
}
