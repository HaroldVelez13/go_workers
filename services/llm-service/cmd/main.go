package main

import (
	"log"

	"github.com/HaroldVelez13/go_workers/services/llm-service/internal/consumer"
	"github.com/HaroldVelez13/go_workers/shared/nats"
)

func main() {
	nc, err := nats.NewClient(nats.Config{
		URL: "nats://nats:4222",
	})
	if err != nil {
		log.Fatal(err)
	}

	consumer := consumer.NewMessageConsumer(nc)

	err = consumer.Start()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("llm-service listening...")

	select {} // 👈 mantiene el servicio vivo
}
