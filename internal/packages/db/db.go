package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var DB *bun.DB

func Connect() {
	_ = godotenv.Load()

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	ssl := os.Getenv("DB_SSLMODE")

	if host == "" || port == "" || user == "" || name == "" {
		log.Fatal("DB env belum lengkap")
	}

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		user,
		pass,
		host,
		port,
		name,
		ssl,
	)

	sqldb := sql.OpenDB(pgdriver.NewConnector(
		pgdriver.WithDSN(dsn),
	))

	DB = bun.NewDB(sqldb, pgdialect.New())

	if err := DB.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("[DB] Connected")
}

// package db

// import (
// 	"database/sql"
// 	"log"
// 	"os"

// 	"github.com/joho/godotenv"
// 	"github.com/uptrace/bun"
// 	"github.com/uptrace/bun/dialect/pgdialect"
// 	"github.com/uptrace/bun/driver/pgdriver"
// )

// var DB *bun.DB

// func Connect() {
// 	// load .env
// 	_ = godotenv.Load()

// 	// ambil dari env
// 	dsn := os.Getenv("DB_DSN")
// 	if dsn == "" {
// 		log.Fatal("DB_DSN tidak ditemukan di .env")
// 	}

// 	// buka koneksi
// 	sqldb := sql.OpenDB(pgdriver.NewConnector(
// 		pgdriver.WithDSN(dsn),
// 	))

// 	DB = bun.NewDB(sqldb, pgdialect.New())

// 	// test koneksi
// 	if err := DB.Ping(); err != nil {
// 		log.Fatal(err)
// 	}

// 	log.Println("[DB] Connected")
// }
