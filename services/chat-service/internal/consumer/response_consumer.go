package consumer

import (
	"encoding/json"
	"log"

	"github.com/HaroldVelez13/go_workers/chat-service/internal/repository"
	"github.com/HaroldVelez13/go_workers/shared/events"
	"github.com/HaroldVelez13/go_workers/shared/nats"
)

type ResponseConsumer struct {
	nc   *nats.Client
	repo *repository.MessageRepository
}

func NewResponseConsumer(nc *nats.Client, repo *repository.MessageRepository) *ResponseConsumer {
	return &ResponseConsumer{nc: nc, repo: repo}
}

func (c *ResponseConsumer) Start() error {
	return c.nc.Subscribe(events.SubjectResponseGenerated, func(msg []byte) error {
		var event events.ResponseGeneratedEvent

		if err := json.Unmarshal(msg, &event); err != nil {
			return err
		}

		log.Println("Respuesta recibida:", event.Payload.Content)

		// Guardar como mensaje del assistant
		_, err := c.repo.Create(
			event.Payload.ChatID,
			event.Payload.Content,
			"assistant",
		)

		return err
	})
}
