package repository

import (
	"app/internal/packages/db"
	"app/internal/services/v1/auto-short/models"
	"context"
	"fmt"
)

func CreateVideo(video *models.Video) error {
	ctx := context.Background()
	_, err := db.DB.NewInsert().Model(video).Exec(ctx)
	if err != nil {
		fmt.Println("CreateVideo failed:", err)
	}
	return err
}

func GetVideo() ([]models.Video, error) {
	ctx := context.Background()
	var videos []models.Video
	err := db.DB.NewSelect().Model(&videos).Scan(ctx)
	if err != nil {
		fmt.Println("GetVideo failed:", err)
	}
	return videos, err
}

func GetVideoByID(id string) (*models.Video, error) {
	ctx := context.Background()
	video := new(models.Video)
	err := db.DB.NewSelect().Model(video).Where("id = ?", id).Scan(ctx)
	if err != nil {
		fmt.Println("GetVideoByID failed:", err)
	}
	return video, err
}

func UpdateVideo(video *models.Video) error {
	ctx := context.Background()
	_, err := db.DB.NewUpdate().Model(video).WherePK().Exec(ctx)
	if err != nil {
		fmt.Println("UpdateVideo failed:", err)
	}
	return err
}

func DeleteVideo(id string) error {
	ctx := context.Background()
	video := &models.Video{ID: id}
	_, err := db.DB.NewDelete().Model(video).WherePK().Exec(ctx)
	if err != nil {
		fmt.Println("DeleteVideo failed:", err)
	}
	return err
}
