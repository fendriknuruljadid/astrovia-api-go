package repository

import (
	"app/internal/packages/db"
	"app/internal/services/v1/astro-zenith/auto-clip/models"
	"context"
	"fmt"
)

func CreateVideos(video *models.Videos) error {
	ctx := context.Background()
	_, err := db.DB.NewInsert().Model(video).Exec(ctx)
	if err != nil {
		fmt.Println("Create Videos failed:", err)
	}
	return err
}

func GetVideos() ([]models.Videos, error) {
	ctx := context.Background()
	var videos []models.Videos
	err := db.DB.NewSelect().Model(&videos).Scan(ctx)
	if err != nil {
		fmt.Println("Get Videos failed:", err)
	}
	return videos, err
}

func GetVideosByID(id string) (*models.Videos, error) {
	ctx := context.Background()
	video := new(models.Videos)
	err := db.DB.NewSelect().Model(video).Where("id = ?", id).Scan(ctx)
	if err != nil {
		fmt.Println("Get Videos By ID failed:", err)
	}
	return video, err
}

func UpdateVideos(video *models.Videos) error {
	ctx := context.Background()
	_, err := db.DB.NewUpdate().Model(video).WherePK().Exec(ctx)
	if err != nil {
		fmt.Println("Update Videos failed:", err)
	}
	return err
}

func DeleteVideos(id string) error {
	ctx := context.Background()
	video := &models.Videos{ID: id}
	_, err := db.DB.NewDelete().Model(video).WherePK().Exec(ctx)
	if err != nil {
		fmt.Println("Delete Videos failed:", err)
	}
	return err
}
