package consumer

import (
	"encoding/json"
	"log"

	"github.com/HaroldVelez13/go_workers/gateway/internal/sse"
	"github.com/HaroldVelez13/go_workers/shared/events"
	"github.com/HaroldVelez13/go_workers/shared/nats"
)

type ChunkConsumer struct {
	nc      *nats.Client
	manager *sse.Manager
}

func NewChunkConsumer(nc *nats.Client, m *sse.Manager) *ChunkConsumer {
	return &ChunkConsumer{nc: nc, manager: m}
}

func (c *ChunkConsumer) Start() error {
	return c.nc.Subscribe(events.SubjectResponseChunkGenerated, func(msg []byte) error {
		var event events.ResponseChunkGeneratedEvent

		if err := json.Unmarshal(msg, &event); err != nil {
			return err
		}

		log.Println("Chunk recibido:", event.Payload.Content)

		// 👇 enviar chunk al cliente SSE
		c.manager.Broadcast(event.Payload.ChatID, toJSON(event))

		return nil
	})
}

// helper
func toJSON(v any) string {
	b, _ := json.Marshal(v)
	return string(b)
}
