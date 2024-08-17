package db

import (
	"log"

	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

var Engine *xorm.Engine

func InitDB() {
	var err error

	connStr := "host=localhost port=5432 user=postgres dbname=dailyjournal password=admin123 sslmode=disable"

	// Membuat XORM engine
	Engine, err = xorm.NewEngine("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to create XORM engine: %v", err)
	}

	// Menguji koneksi
	if err := Engine.Ping(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Successfully connected to PostgreSQL database!")
}