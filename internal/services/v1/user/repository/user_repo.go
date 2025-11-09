package repository

import (
    "context"
    "fmt"
    "app/internal/packages/db"
    "app/internal/services/v1/user/models"
)

func CreateUser(user *models.User) error {
    ctx := context.Background()
    _, err := db.DB.NewInsert().Model(user).Exec(ctx)
    if err != nil {
        fmt.Println("CreateUser failed:", err)
    }
    return err
}

func GetUsers() ([]models.User, error) {
    ctx := context.Background()
    var users []models.User
    err := db.DB.NewSelect().Model(&users).Scan(ctx)
    if err != nil {
        fmt.Println("GetUsers failed:", err)
    }
    return users, err
}

func GetUserByID(id string) (*models.User, error) {
    ctx := context.Background()
    user := new(models.User)
    err := db.DB.NewSelect().Model(user).Where("id = ?", id).Scan(ctx)
    if err != nil {
        fmt.Println("GetUserByID failed:", err)
    }
    return user, err
}

func UpdateUser(user *models.User) error {
    ctx := context.Background()
    _, err := db.DB.NewUpdate().Model(user).WherePK().Exec(ctx)
    if err != nil {
        fmt.Println("UpdateUser failed:", err)
    }
    return err
}

func DeleteUser(id string) error {
    ctx := context.Background()
    user := &models.User{ID: id}
    _, err := db.DB.NewDelete().Model(user).WherePK().Exec(ctx)
    if err != nil {
        fmt.Println("DeleteUser failed:", err)
    }
    return err
}
