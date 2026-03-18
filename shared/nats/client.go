package nats

import (
	"encoding/json"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

type Client struct {
	conn *nats.Conn
}

func NewClient(cfg Config) (*Client, error) {
	conn, err := nats.Connect(
		cfg.URL,
		nats.Name("go-service"),
		nats.ReconnectWait(2*time.Second),
		nats.MaxReconnects(-1),
		nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
			log.Println("NATS disconnected:", err)
		}),
		nats.ReconnectHandler(func(nc *nats.Conn) {
			log.Println("NATS reconnected")
		}),
	)

	if err != nil {
		return nil, err
	}

	return &Client{conn: conn}, nil
}

func (c *Client) Publish(subject string, data interface{}) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return c.conn.Publish(subject, bytes)
}
