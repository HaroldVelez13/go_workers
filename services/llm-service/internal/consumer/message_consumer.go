package consumer

import (
	"encoding/json"
	"log"

	"github.com/HaroldVelez13/go_workers/shared/events"
	"github.com/HaroldVelez13/go_workers/shared/nats"
)

type MessageConsumer struct {
	nc *nats.Client
}

func NewMessageConsumer(nc *nats.Client) *MessageConsumer {
	return &MessageConsumer{nc: nc}
}

func (c *MessageConsumer) Start() error {
	return c.nc.Subscribe(events.SubjectMessageCreated, func(msg []byte) error {
		var event events.MessageCreatedEvent

		if err := json.Unmarshal(msg, &event); err != nil {
			return err
		}

		log.Println("Mensaje recibido:", event.Payload.Content)

		// 🤖 respuesta mock (luego LLM real)
		response := "Echo: " + event.Payload.Content

		respEvent := events.ResponseGeneratedEvent{
			BaseEvent: events.NewBaseEvent("chat.response.generated", event.Metadata.UserID, event.Metadata.TraceID),
			Payload: events.ResponseGeneratedPayload{
				ChatID:    event.Payload.ChatID,
				MessageID: event.Payload.MessageID + "-resp",
				Content:   response,
				Role:      "assistant",
			},
		}

		return c.nc.Publish(events.SubjectResponseGenerated, respEvent)
	})
}
