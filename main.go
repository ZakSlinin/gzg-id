package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	pgUser := os.Getenv("DB_USER")
	pgPass := os.Getenv("DB_PASSWORD")
	pgHost := os.Getenv("DB_HOST")
	pgPort := os.Getenv("DB_PORT")
	pgDB := os.Getenv("DB_NAME")

	//TODO: sslmode=active
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", pgUser, pgPass, pgHost, pgPort, pgDB)

	m, err := migrate.New(
		"file://db/migrations",
		dbUrl,
	)

	if err != nil {
		log.Fatalf("failed to initialize migrations: %s", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("failed to run migrations: %s", err)
	}

	log.Println("migrations complete")

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		log.Fatalf("failed to connect to database via GORM: %s", err)
	}

	r := gin.Default()

	api := r.Group("/api/gzg-id")

	r.Run(":8080")
}
