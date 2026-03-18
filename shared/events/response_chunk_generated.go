package events

const SubjectResponseChunkGenerated = "chat.response.chunk.generated.v1"

type ResponseChunkGeneratedEvent struct {
	BaseEvent
	Payload struct {
		ChatID  string `json:"chat_id"`
		Content string `json:"content"`
		Done    bool   `json:"done"`
	} `json:"payload"`
}
