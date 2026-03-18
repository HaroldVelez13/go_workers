package events

import (
	"time"
)

type Metadata struct {
	TraceID string `json:"trace_id"`
	UserID  string `json:"user_id"`
}

type BaseEvent struct {
	EventID   string    `json:"event_id"`
	Type      string    `json:"type"`
	Version   string    `json:"version"`
	Timestamp time.Time `json:"timestamp"`
	Metadata  Metadata  `json:"metadata"`
}
