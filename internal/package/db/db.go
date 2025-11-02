package db

import (
    "database/sql"
    "github.com/uptrace/bun"
    "github.com/uptrace/bun/driver/pgdriver"
    "github.com/uptrace/bun/dialect/pgdialect"
)

var DB *bun.DB

func Connect() {
    sqldb := sql.OpenDB(pgdriver.NewConnector(
        pgdriver.WithDSN("postgres://postgres:Cheat1234@localhost:5432/astrovia_new?sslmode=disable"),
    ))

    DB = bun.NewDB(sqldb, pgdialect.New())
    // optional: cek koneksi
    if err := DB.Ping(); err != nil {
        panic(err)
    }
}
