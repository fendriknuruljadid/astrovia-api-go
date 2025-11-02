package repository

import (
    "context"
    "astrovia-api-go/internal/package/db"
    "astrovia-api-go/internal/services/user/models"
)

func CreateUser(user *models.User) error {
    _, err := db.DB.NewInsert().Model(user).Exec(context.Background())
    return err
}

func GetUsers() ([]models.User, error) {
    var users []models.User
    err := db.DB.NewSelect().Model(&users).Scan(context.Background())
    return users, err
}

func GetUserByID(id int64) (*models.User, error) {
    user := new(models.User)
    err := db.DB.NewSelect().Model(user).Where("id = ?", id).Scan(context.Background())
    return user, err
}

func UpdateUser(user *models.User) error {
    _, err := db.DB.NewUpdate().Model(user).WherePK().Exec(context.Background())
    return err
}

func DeleteUser(id int64) error {
    user := &models.User{ID: id}
    _, err := db.DB.NewDelete().Model(user).WherePK().Exec(context.Background())
    return err
}
