package events

type UserCreatedPayload struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
}

type UserCreatedEvent struct {
	BaseEvent
	Payload UserCreatedPayload `json:"payload"`
}

const SubjectUserCreated = "user.created.v1"
