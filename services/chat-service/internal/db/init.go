package db

import (
	"database/sql"
	"log"
)

func Init(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS messages (
		id TEXT PRIMARY KEY,
		chat_id TEXT,
		content TEXT,
		role TEXT,
		created_at TIMESTAMP
	);
	`

	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	log.Println("DB initialized (messages table ready)")
	return nil
}
