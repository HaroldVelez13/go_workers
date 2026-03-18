package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/HaroldVelez13/go_workers/shared/nats"

	"github.com/HaroldVelez13/go_workers/chat-service/internal/consumer"
	db "github.com/HaroldVelez13/go_workers/chat-service/internal/db"
	"github.com/HaroldVelez13/go_workers/chat-service/internal/handler"
	"github.com/HaroldVelez13/go_workers/chat-service/internal/repository"
	"github.com/HaroldVelez13/go_workers/chat-service/internal/service"
)

func main() {
	// 🗄️ DB
	dbConn, err := sql.Open("postgres", "postgres://admin:admin@postgres:5432/app_db?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	// 👇 INIT DB (AQUÍ VA)
	err = db.Init(dbConn)
	if err != nil {
		log.Fatal("DB init error:", err)
	}

	// 🔌 NATS
	nc, err := nats.NewClient(nats.Config{
		URL: "nats://nats:4222",
	})
	if err != nil {
		log.Fatal(err)
	}

	// 🧱 Layers
	repo := repository.NewMessageRepository(dbConn)
	svc := service.NewChatService(repo, nc)
	h := handler.NewChatHandler(svc)

	chunkConsumer := consumer.NewChunkConsumer(nc, repo)

	go func() {
		if err := chunkConsumer.Start(); err != nil {
			log.Println("chunk consumer error:", err)
		}
	}()

	// 🌐 Routes
	http.HandleFunc("/messages", h.CreateMessage)

	log.Println("chat-service running on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
