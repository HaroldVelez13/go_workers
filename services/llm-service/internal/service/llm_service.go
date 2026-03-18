package service

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/HaroldVelez13/go_workers/shared/events"
	"github.com/HaroldVelez13/go_workers/shared/nats"
)

type LLMService struct {
	nc *nats.Client
}

func NewLLMService(nc *nats.Client) *LLMService {
	return &LLMService{nc: nc}
}

func (s *LLMService) StreamResponse(chatID string, content string, userID string, traceID string) error {
	words := strings.Split(content, " ")

	for _, word := range words {
		event := events.ResponseChunkGeneratedEvent{
			BaseEvent: events.NewBaseEvent(
				events.SubjectResponseChunkGenerated,
				"",
				uuid.NewString(),
			),
		}

		event.Payload.ChatID = chatID
		event.Payload.Content = word
		event.Payload.Done = false

		data, _ := json.Marshal(event)

		if err := s.nc.Publish(events.SubjectResponseChunkGenerated, data); err != nil {
			return err
		}

		time.Sleep(200 * time.Millisecond) // simula streaming
	}

	// evento final
	doneEvent := events.ResponseChunkGeneratedEvent{
		BaseEvent: events.NewBaseEvent(
			events.SubjectResponseChunkGenerated,
			"",
			uuid.NewString(),
		),
	}

	doneEvent.Payload.ChatID = chatID
	doneEvent.Payload.Done = true

	data, _ := json.Marshal(doneEvent)

	return s.nc.Publish(events.SubjectResponseChunkGenerated, data)
}
