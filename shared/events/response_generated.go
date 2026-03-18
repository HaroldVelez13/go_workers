package events

type ResponseGeneratedPayload struct {
	ChatID    string `json:"chat_id"`
	MessageID string `json:"message_id"` // respuesta
	Content   string `json:"content"`
	Role      string `json:"role"` // "assistant"
}

type ResponseGeneratedEvent struct {
	BaseEvent
	Payload ResponseGeneratedPayload `json:"payload"`
}

const SubjectResponseGenerated = "chat.response.generated.v1"
