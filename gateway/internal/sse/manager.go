package sse

import "sync"

type Client chan string

type Manager struct {
	clients map[string][]Client // chatID → clients
	mu      sync.RWMutex
}

func NewManager() *Manager {
	return &Manager{
		clients: make(map[string][]Client),
	}
}

func (m *Manager) AddClient(chatID string, client Client) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.clients[chatID] = append(m.clients[chatID], client)
}

func (m *Manager) RemoveClient(chatID string, client Client) {
	m.mu.Lock()
	defer m.mu.Unlock()

	clients := m.clients[chatID]
	for i, c := range clients {
		if c == client {
			m.clients[chatID] = append(clients[:i], clients[i+1:]...)
			break
		}
	}
}

func (m *Manager) Broadcast(chatID string, msg string) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for _, client := range m.clients[chatID] {
		select {
		case client <- msg:
		default:
			// evita bloqueo si el cliente no consume
		}
	}
}
