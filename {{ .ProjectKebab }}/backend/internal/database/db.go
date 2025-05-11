package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// Connect establishes a connection to the database
func Connect() error {
	var (
		host     = getEnv("DB_HOST", "localhost")
		port     = getEnv("DB_PORT", "5432")
		user     = getEnv("DB_USER", "postgres")
		password = getEnv("DB_PASSWORD", "postgres")
		dbname   = getEnv("DB_NAME", "{{ .ProjectKebab }}")
		sslmode  = getEnv("DB_SSLMODE", "disable")
	)

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	// Set connection pool settings
	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(5)
	DB.SetConnMaxLifetime(5 * time.Minute)

	// Test the connection
	err = DB.Ping()
	if err != nil {
		return err
	}

	log.Println("Successfully connected to database")
	return nil
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
