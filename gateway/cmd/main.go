package main

import (
	"log"
	"net/http"

	"github.com/HaroldVelez13/go_workers/gateway/internal/consumer"
	"github.com/HaroldVelez13/go_workers/gateway/internal/handler"
	"github.com/HaroldVelez13/go_workers/gateway/internal/sse"
	"github.com/HaroldVelez13/go_workers/shared/nats"
)

func main() {
	// NATS
	nc, err := nats.NewClient(nats.Config{
		URL: "nats://nats:4222",
	})
	if err != nil {
		log.Fatal(err)
	}

	// SSE manager
	manager := sse.NewManager()

	// consumer
	consumer := consumer.NewResponseConsumer(nc, manager)
	go func() {
		if err := consumer.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	// handler
	streamHandler := handler.NewStreamHandler(manager)

	http.HandleFunc("/chat/stream", streamHandler.Stream)

	log.Println("gateway running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
