package initializers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	host := getEnv("POSTGRES_HOST", "postgres") // В контейнере — имя сервиса
	user := getEnv("POSTGRES_USER", "admin")
	pass := getEnv("POSTGRES_PASSWORD", "admin")
	db := getEnv("POSTGRES_DB", "mydb")
	port := getEnv("POSTGRES_PORT", "5432")
	ssl := getEnv("POSTGRES_SSLMODE", "disable")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, pass, db, port, ssl,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
