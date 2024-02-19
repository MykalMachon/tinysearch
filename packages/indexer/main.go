package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/redis/go-redis/v9"
	// "gorm.io/driver/postgres"
)

var redisClient *redis.Client
var ctx context.Context
var log slog.Logger

// var db *gorm.DB

func init() {
	// do setup here... this is always run before main
	// TODO I need to know what this line does, lol.
	ctx = context.Background()

	// * initialize the logger
	jsonHandler := slog.NewJSONHandler(os.Stdout, nil)
	log = *slog.New(jsonHandler)

	log.Info("initializing indexer")

	// * initialize the postgresql database and run migrations

	// * initialize the redis queue and make sure it works
	redisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	if err := redisClient.Ping(ctx).Err(); err != nil {
		log.Error("failed to connect to redis", "error", err.Error())
		os.Exit(1)
	}
}

func main() {
	log.Info("starting indexer")

	// start the indexer workers
	// 1. query all rows in the source table in the database
	// 2. for each row, enqueue a job to the redis queue

	// start a worker that listens to the redis queue and processes the jobs
	// 1. dequeue a job from the redis queue
	// 2. check the last indexed timestamp from the database
	// 3. if the job timestamp is greater than the last indexed timestamp, index the row
	// 4. update the last indexed timestamp in the database
	// 5. repeat
}
