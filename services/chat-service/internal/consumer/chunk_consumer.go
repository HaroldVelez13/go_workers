package consumer

import (
	"encoding/json"
	"log"

	"github.com/HaroldVelez13/go_workers/chat-service/internal/repository"
	"github.com/HaroldVelez13/go_workers/chat-service/internal/stream"
	"github.com/HaroldVelez13/go_workers/shared/events"
	"github.com/HaroldVelez13/go_workers/shared/nats"
)

type ChunkConsumer struct {
	nc     *nats.Client
	repo   *repository.MessageRepository
	buffer *stream.Buffer
}

func NewChunkConsumer(nc *nats.Client, repo *repository.MessageRepository) *ChunkConsumer {
	return &ChunkConsumer{
		nc:     nc,
		repo:   repo,
		buffer: stream.NewBuffer(),
	}
}

func (c *ChunkConsumer) Start() error {
	return c.nc.Subscribe(events.SubjectResponseChunkGenerated, func(msg []byte) error {
		var event events.ResponseChunkGeneratedEvent

		if err := json.Unmarshal(msg, &event); err != nil {
			return err
		}

		chatID := event.Payload.ChatID

		// 🧠 acumular chunk
		if !event.Payload.Done {
			c.buffer.Append(chatID, event.Payload.Content)
			return nil
		}

		// ✅ DONE → reconstruir mensaje
		full := c.buffer.Get(chatID)

		log.Println("Mensaje completo:", full)

		// 💾 guardar en DB
		_, err := c.repo.Create(chatID, full, "assistant")
		if err != nil {
			return err
		}

		// limpiar buffer
		c.buffer.Clear(chatID)

		return nil
	})
}
