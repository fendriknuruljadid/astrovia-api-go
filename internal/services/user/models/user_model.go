package models

import "github.com/uptrace/bun"

type User struct {
    bun.BaseModel `bun:"table:users"`

    ID       int64  `bun:",pk,autoincrement" json:"id"`
    Name     string `json:"name"`
    Email    string `json:"email"`
    Password string `json:"-"`
}
