package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
    var err error

    connection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        GetEnv("DB_HOST", "localhost"),
        GetEnv("DB_PORT", "5432"),
        GetEnv("DB_USER", "postgres"),
        GetEnv("DB_PASSWORD", ""),
        GetEnv("DB_NAME", "mytable"),
    )

    DB, err = sql.Open("postgres", connection)
    if err != nil {
        log.Fatal("Erro ao conectar no banco:", err)
    }

    if err = DB.Ping(); err != nil {
        log.Fatal("Banco não responde:", err)
    }

    log.Println("Banco conectado com sucesso.")
}
