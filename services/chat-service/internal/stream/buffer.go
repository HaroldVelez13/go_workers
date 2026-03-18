package stream

import "sync"

type Buffer struct {
	data map[string]string
	mu   sync.Mutex
}

func NewBuffer() *Buffer {
	return &Buffer{
		data: make(map[string]string),
	}
}

func (b *Buffer) Append(chatID string, chunk string) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.data[chatID] += chunk
}

func (b *Buffer) Get(chatID string) string {
	b.mu.Lock()
	defer b.mu.Unlock()

	return b.data[chatID]
}

func (b *Buffer) Clear(chatID string) {
	b.mu.Lock()
	defer b.mu.Unlock()

	delete(b.data, chatID)
}
