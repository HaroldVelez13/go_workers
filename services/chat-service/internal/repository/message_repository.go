package repository

import (
	"database/sql"
	"time"

	"github.com/HaroldVelez13/go_workers/chat-service/internal/model"
	"github.com/google/uuid"
)

type MessageRepository struct {
	db *sql.DB
}

func NewMessageRepository(db *sql.DB) *MessageRepository {
	return &MessageRepository{db: db}
}

func (r *MessageRepository) Create(chatID, content, role string) (*model.Message, error) {
	msg := &model.Message{
		ID:        uuid.NewString(),
		ChatID:    chatID,
		Content:   content,
		Role:      role,
		CreatedAt: time.Now(),
	}

	_, err := r.db.Exec(`
		INSERT INTO messages (id, chat_id, content, role, created_at)
		VALUES ($1, $2, $3, $4, $5)
	`, msg.ID, msg.ChatID, msg.Content, msg.Role, msg.CreatedAt)

	return msg, err
}
