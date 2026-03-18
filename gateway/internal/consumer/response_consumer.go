package consumer

import (
	"encoding/json"
	"log"

	"github.com/HaroldVelez13/go_workers/gateway/internal/sse"
	"github.com/HaroldVelez13/go_workers/shared/events"
	"github.com/HaroldVelez13/go_workers/shared/nats"
)

type ResponseConsumer struct {
	nc      *nats.Client
	manager *sse.Manager
}

func NewResponseConsumer(nc *nats.Client, m *sse.Manager) *ResponseConsumer {
	return &ResponseConsumer{nc: nc, manager: m}
}

func (c *ResponseConsumer) Start() error {
	return c.nc.Subscribe(events.SubjectResponseGenerated, func(msg []byte) error {
		var event events.ResponseGeneratedEvent

		if err := json.Unmarshal(msg, &event); err != nil {
			return err
		}

		log.Println("Gateway received:", event.Payload.Content)

		// enviar a clientes SSE
		c.manager.Broadcast(event.Payload.ChatID, event.Payload.Content)

		return nil
	})
}
