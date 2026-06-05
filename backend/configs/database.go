package configs

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func ConnectDB(cfg *Config) *sql.DB {

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)

	db, err := sql.Open("pgx", dsn)

	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal("Database unreachable:", err)
	}

	fmt.Println("Database connected successfully")

	return db
}