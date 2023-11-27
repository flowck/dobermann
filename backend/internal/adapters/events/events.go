package events

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/ThreeDotsLabs/watermill/message"
)

type Event interface {
	EventName() string
}

type Header struct {
	Name          string `json:"name"`
	CorrelationID string `json:"correlation_id"`
}

func NewHeader(name, correlationID string) Header {
	return Header{
		Name:          name,
		CorrelationID: correlationID,
	}
}

type MonitorCreatedEvent struct {
	Header    Header
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

func (e MonitorCreatedEvent) EventName() string {
	return "MonitorCreatedEvent_v1"
}

func mapEventToMessage(event Event) (*message.Message, error) {
	data, err := json.Marshal(event)
	if err != nil {
		return nil, fmt.Errorf("unable to marshall event: %s due to: %v", event.EventName(), err)
	}

	return &message.Message{
		Payload: data,
	}, nil
}
