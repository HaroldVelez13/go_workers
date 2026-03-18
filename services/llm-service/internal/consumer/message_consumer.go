package consumer

import (
	"encoding/json"
	"log"

	"github.com/HaroldVelez13/go_workers/services/llm-service/internal/service"
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

		// 👇 inicializar servicio LLM
		svc := service.NewLLMService(c.nc)

		// 🤖 generar respuesta (puedes cambiar esto luego por OpenAI)
		response := "Echo: " + event.Payload.Content

		// 👇 STREAMING REAL (este es el cambio importante)
		return svc.StreamResponse(
			event.Payload.ChatID,
			response,
			event.Metadata.UserID,
			event.Metadata.TraceID,
		)
	})
}
