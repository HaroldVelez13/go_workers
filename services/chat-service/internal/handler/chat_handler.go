package handler

import (
	"encoding/json"
	"net/http"

	"github.com/HaroldVelez13/go_workers/chat-service/internal/service"
)

type ChatHandler struct {
	svc *service.ChatService
}

func NewChatHandler(s *service.ChatService) *ChatHandler {
	return &ChatHandler{svc: s}
}

func (h *ChatHandler) CreateMessage(w http.ResponseWriter, r *http.Request) {
	var body struct {
		ChatID  string `json:"chat_id"`
		Content string `json:"content"`
		UserID  string `json:"user_id"`
	}

	_ = json.NewDecoder(r.Body).Decode(&body)

	err := h.svc.CreateMessage(body.ChatID, body.Content, body.UserID)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
