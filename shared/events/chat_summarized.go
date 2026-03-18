package events

type ChatSummarizedPayload struct {
	ChatID  string `json:"chat_id"`
	Summary string `json:"summary"`
}

type ChatSummarizedEvent struct {
	BaseEvent
	Payload ChatSummarizedPayload `json:"payload"`
}

const SubjectChatSummarized = "chat.summarized.v1"
