package repository

import (
	"app/internal/packages/db"
	"app/internal/services/v1/astro-zenith/auto-clip/models"

	userAgentModels "app/internal/services/v1/user-agent/models"
	"context"
	"encoding/json"
	"fmt"

	"github.com/uptrace/bun"
)

type VideoProgress struct {
	Stage   string `json:"stage"`
	Percent int    `json:"percent"`
	Message string `json:"message"`
}

func CreateVideos(video *models.Videos) error {
	ctx := context.Background()
	_, err := db.DB.NewInsert().Model(video).Exec(ctx)
	if err != nil {
		fmt.Println("Create Videos failed:", err)
	}
	return err
}

// func GetVideos() ([]models.Videos, error) {
// 	ctx := context.Background()
// 	var videos []models.Videos
// 	err := db.DB.NewSelect().Model(&videos).Scan(ctx)
// 	if err != nil {
// 		fmt.Println("Get Videos failed:", err)
// 	}
// 	return videos, err
// }

// func GetVideosByID(id string) (*models.Videos, error) {
// 	ctx := context.Background()

// 	videos := new(models.Videos)

// 	err := db.DB.NewSelect().
// 		Model(videos).
// 		Relation("Clips").
// 		Where("videos.id = ?", id).
// 		Scan(ctx)

// 	fmt.Println("Clips count:", len(videos.Clips))

// 	if err != nil {
// 		fmt.Println("GetVideoByID failed:", err)
// 		return nil, err
// 	}

// 	return videos, nil
// }

func GetVideos(userId string) ([]models.Videos, error) {
	ctx := context.Background()
	var videos []models.Videos

	// Order by videos.id desc
	err := db.DB.NewSelect().
		Model(&videos).
		Where("users_id = ?", userId).
		OrderExpr("videos.id DESC").
		Scan(ctx)
	if err != nil {
		fmt.Println("Get Videos failed:", err)
	}
	return videos, err
}

func GetUserAgentByUser(uid string) (*userAgentModels.UserAgent, error) {
	ctx := context.Background()
	user := new(userAgentModels.UserAgent)
	err := db.DB.NewSelect().
		Model(user).Where("users_id = ?", uid).Scan(ctx)
	if err != nil {
		fmt.Println("GetUserAgentByUser failed:", err)
	}
	return user, err
}

func GetVideosByID(id string) (*models.Videos, error) {
	ctx := context.Background()
	videos := new(models.Videos)

	// Order clips by viral_score desc
	err := db.DB.NewSelect().
		Model(videos).
		Relation("Clips", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.OrderExpr("viral_score DESC")
		}).
		Where("videos.id = ?", id).
		Scan(ctx)

	fmt.Println("Clips count:", len(videos.Clips))

	if err != nil {
		fmt.Println("GetVideoByID failed:", err)
		return nil, err
	}

	return videos, nil
}

func UpdateVideos(video *models.Videos) error {
	ctx := context.Background()
	_, err := db.DB.NewUpdate().Model(video).WherePK().Exec(ctx)
	if err != nil {
		fmt.Println("Update Videos failed:", err)
	}
	return err
}

func UpdateVideoProgress(videoID string, progress VideoProgress) error {
	ctx := context.Background()

	progressJSON, err := json.Marshal(progress)
	if err != nil {
		return err
	}

	_, err = db.DB.NewUpdate().
		Model((*models.Videos)(nil)).
		Set("video_progress = ?", json.RawMessage(progressJSON)).
		Set("updated_at = NOW()").
		Where("id = ?", videoID).
		Exec(ctx)

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
