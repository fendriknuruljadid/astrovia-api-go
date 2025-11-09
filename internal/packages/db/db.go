package db

import (
    "database/sql"
    "log"
    "os"
    "github.com/joho/godotenv"
    "github.com/uptrace/bun"
    "github.com/uptrace/bun/driver/pgdriver"
    "github.com/uptrace/bun/dialect/pgdialect"
)

var DB *bun.DB

func Connect() {
    // load .env
    _ = godotenv.Load()

    // ambil dari env
    dsn := os.Getenv("DB_DSN")
    if dsn == "" {
        log.Fatal("DB_DSN tidak ditemukan di .env")
    }

    // buka koneksi
    sqldb := sql.OpenDB(pgdriver.NewConnector(
        pgdriver.WithDSN(dsn),
    ))

    DB = bun.NewDB(sqldb, pgdialect.New())

    // test koneksi
    if err := DB.Ping(); err != nil {
        log.Fatal(err)
    }

    log.Println("[DB] Connected ✅")
}
