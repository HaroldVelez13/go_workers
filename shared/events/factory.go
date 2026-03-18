package events

import (
	"time"

	"github.com/google/uuid"
)

func NewBaseEvent(eventType string, userID string, traceID string) BaseEvent {
	return BaseEvent{
		EventID:   uuid.NewString(),
		Type:      eventType,
		Version:   "v1",
		Timestamp: time.Now().UTC(),
		Metadata: Metadata{
			TraceID: traceID,
			UserID:  userID,
		},
	}
}
