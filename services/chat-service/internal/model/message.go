package model

import "time"

type Message struct {
	ID        string
	ChatID    string
	Content   string
	Role      string
	CreatedAt time.Time
}
