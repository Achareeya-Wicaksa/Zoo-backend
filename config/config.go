package config

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
    var err error
    dsn := "root:@tcp(127.0.0.1:3306)/testbackendzoo"
    DB, err = sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
    }

    if err = DB.Ping(); err != nil {
        log.Fatal(err)
    }

    fmt.Println("Connected to the database!")
}
