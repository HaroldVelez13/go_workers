package events

type MessageCreatedPayload struct {
	ChatID    string `json:"chat_id"`
	MessageID string `json:"message_id"`
	Content   string `json:"content"`
	Role      string `json:"role"` // "user"
}

type MessageCreatedEvent struct {
	BaseEvent
	Payload MessageCreatedPayload `json:"payload"`
}

const SubjectMessageCreated = "chat.message.created.v1"
