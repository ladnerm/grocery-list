package storage

import (
	"database/sql"
	"fmt"
	"github.com/ladnerm/grocery-list/types"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func InitDB() *sql.DB {
	config := types.Config{
		Host:    getEnv("DB_HOST"),
		Port:    getEnv("DB_PORT"),
		User:    getEnv("DB_USER"),
		DBName:  getEnv("DB_NAME"),
		SSLMODE: getEnv("DB_SSLMODE"),
	}

	i := fmt.Sprintf("postgres://%s:@%s/%s?sslmode=disable",
		config.User, config.Host, config.DBName)

	fmt.Printf("CONNECTING TO: %s", i)

	db, err := sql.Open("postgres", i)
	if err != nil {
		log.Fatalf("Could not open DB: %v\n", err)
	}
	return db
}

func getEnv(s string) string {
	v := os.Getenv(s)
	if v == "" {
		log.Fatalf("Could not get %s from .env", s)
	}
	return v
}
