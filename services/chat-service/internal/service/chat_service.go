package service

import (
	"github.com/HaroldVelez13/go_workers/chat-service/internal/repository"
	"github.com/HaroldVelez13/go_workers/shared/events"
	"github.com/HaroldVelez13/go_workers/shared/nats"
)

type ChatService struct {
	repo *repository.MessageRepository
	nc   *nats.Client
}

func NewChatService(r *repository.MessageRepository, nc *nats.Client) *ChatService {
	return &ChatService{repo: r, nc: nc}
}

func (s *ChatService) CreateMessage(chatID, content, userID string) error {
	msg, err := s.repo.Create(chatID, content, "user")
	if err != nil {
		return err
	}

	event := events.MessageCreatedEvent{
		BaseEvent: events.NewBaseEvent("chat.message.created", userID, ""),
		Payload: events.MessageCreatedPayload{
			ChatID:    chatID,
			MessageID: msg.ID,
			Content:   content,
			Role:      "user",
		},
	}

	return s.nc.Publish(events.SubjectMessageCreated, event)
}
