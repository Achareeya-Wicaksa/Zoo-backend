package config

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
    _ "github.com/lib/pq" // PostgreSQL driver
)

var DB *sql.DB

func init() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }
}

func Connect() {
    var err error
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=require",
        os.Getenv("POSTGRES_HOST"),
        os.Getenv("POSTGRES_USER"),
        os.Getenv("POSTGRES_PASSWORD"),
        os.Getenv("POSTGRES_DATABASE"),
    )

    // Membuka koneksi dengan PostgreSQL
    DB, err = sql.Open("postgres", dsn)
    if err != nil {
        log.Fatal(err)
    }

    // Mengecek apakah koneksi berhasil
    if err = DB.Ping(); err != nil {
        log.Fatal(err)
    }

    fmt.Println("Connected to the PostgreSQL database!")
}
