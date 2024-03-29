package command

import "time"

type MonitorCreatedEvent struct {
	ID        string
	CreatedAt time.Time
}

type EndpointCheckFailedEvent struct {
	MonitorID     string
	CheckedURL    string
	CheckResultID string
	At            time.Time
}

type EndpointCheckSucceededEvent struct {
	MonitorID string
	At        time.Time
}

type IncidentResolvedEvent struct {
	MonitorID  string
	IncidentID string
	At         time.Time
}

type IncidentCreatedEvent struct {
	MonitorID  string
	IncidentID string
	At         time.Time
}
