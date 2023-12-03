package amqp

import (
	"github.com/ThreeDotsLabs/watermill/message"

	"github.com/flowck/dobermann/backend/internal/adapters/events"
	"github.com/flowck/dobermann/backend/internal/app"
	"github.com/flowck/dobermann/backend/internal/app/command"
	"github.com/flowck/dobermann/backend/internal/domain"
)

type IncidentResolvedHandler struct {
	application *app.App
}

func (e IncidentResolvedHandler) HandlerName() string {
	return "IncidentResolved_NotifyOnIncidentResolved_Handler"
}

func (e IncidentResolvedHandler) EventName() string {
	return events.IncidentResolvedEvent{}.EventName()
}

func (e IncidentResolvedHandler) Handle(m *message.Message) error {
	event, err := events.NewEventFromMessage[events.IncidentResolvedEvent](m)
	if err != nil {
		return err
	}

	monitorID, err := domain.NewIdFromString(event.MonitorID)
	if err != nil {
		return err
	}

	incidentID, err := domain.NewIdFromString(event.IncidentID)
	if err != nil {
		return err
	}

	cmd := command.NotifyOnIncidentResolved{
		IncidentID: incidentID,
		MonitorID:  monitorID,
	}
	err = e.application.Commands.NotifyOnIncidentResolved.Execute(m.Context(), cmd)
	if err != nil {
		return err
	}

	return nil
}
