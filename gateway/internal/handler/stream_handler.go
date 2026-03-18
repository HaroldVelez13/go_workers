package handler

import (
	"fmt"
	"net/http"

	"github.com/HaroldVelez13/go_workers/gateway/internal/sse"
)

type StreamHandler struct {
	manager *sse.Manager
}

func NewStreamHandler(m *sse.Manager) *StreamHandler {
	return &StreamHandler{manager: m}
}

func (h *StreamHandler) Stream(w http.ResponseWriter, r *http.Request) {
	chatID := r.URL.Query().Get("chat_id")
	if chatID == "" {
		http.Error(w, "chat_id required", 400)
		return
	}

	// headers SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "stream unsupported", 500)
		return
	}

	client := make(sse.Client)
	h.manager.AddClient(chatID, client)
	defer h.manager.RemoveClient(chatID, client)

	ctx := r.Context()

	for {
		select {
		case msg := <-client:
			fmt.Fprintf(w, "data: %s\n\n", msg)
			flusher.Flush()

		case <-ctx.Done():
			return
		}
	}
}
