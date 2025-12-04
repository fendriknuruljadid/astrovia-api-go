package models

import (
    "context"
    "fmt"
    "time"
    "strings"

    "app/internal/packages/utils"
    "github.com/uptrace/bun"
    "golang.org/x/crypto/bcrypt"
)

type User struct {
    bun.BaseModel `bun:"table:users"`

    ID        string    `bun:"id,pk,notnull" json:"id"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    Password  string    `json:"password"`
    Username  string    `json:"username"`
    Provider  string    `json:"provider"`
    CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"created_at"`
    UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"updated_at"`
}

// Generate ID dan hash password otomatis sebelum insert
func (u *User) BeforeAppendModel(ctx context.Context, q bun.Query) error {
    if u.ID == "" {
        u.ID = "usr-" + strings.ToLower(utils.NewULID())
    }

    if u.Password != "" && !strings.HasPrefix(u.Password, "$2a$") {
        hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
        if err == nil {
            u.Password = string(hashed)
        } else {
            fmt.Println("Password hash error:", err)
        }
    }
    u.CreatedAt = time.Now()
    u.UpdatedAt = time.Now()

    return nil
}

// Update timestamp otomatis
func (u *User) BeforeUpdate(ctx context.Context, q bun.Query) error {
    u.UpdatedAt = time.Now()
    return nil
}
