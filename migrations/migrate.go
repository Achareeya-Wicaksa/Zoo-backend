package migrations

import (
	"log"
	"zoo-backend/config"
)

func Migrate() {
	query := `
    CREATE TABLE IF NOT EXISTS animal (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        class VARCHAR(255) NOT NULL,
        legs INT NOT NULL
    );`
	if _, err := config.DB.Exec(query); err != nil {
		log.Fatal("Error creating table: ", err)
	}
	log.Println("Table animal checked/created successfully!")
}
