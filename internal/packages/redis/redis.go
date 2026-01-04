package redis

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var (
	Ctx = context.Background()
	Rdb *redis.Client
)

func InitRedis() {

	_ = godotenv.Load()
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	dbStr := os.Getenv("REDIS_DB")
	password := os.Getenv("REDIS_PASSWORD")

	db, err := strconv.Atoi(dbStr)
	if err != nil {
		log.Fatalf("invalid REDIS_DB: %v", err)
	}

	addr := fmt.Sprintf("%s:%s", host, port)

	Rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	if err := Rdb.Ping(Ctx).Err(); err != nil {
		log.Fatalf("Redis connection failed (%s): %v", addr, err)
	}

	log.Printf("Redis connected → %s", addr)
}
