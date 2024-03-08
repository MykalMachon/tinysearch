package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/mykalmachon/tinysearch/indexer/models"
	"github.com/mykalmachon/tinysearch/indexer/seeds"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var redisClient *redis.Client
var dbClient *gorm.DB

var ctx context.Context = context.Background()
var log slog.Logger = *slog.New(slog.NewJSONHandler(os.Stdout, nil))

func OpenRedisConnection() (*redis.Client, error) {
	rc := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	err := rc.Ping(ctx).Err()

	if err != nil {
		log.Error("failed to connect to redis", "error", err.Error())
	}

	return rc, err
}

func OpenDatabaseConnection() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: os.Getenv("DATABASE_URL"),
	}), &gorm.Config{})

	if err != nil {
		log.Error("failed to connect to database", "error", err.Error())
	}

	return db, err
}

func ProcessIndexingJobs(redisClient redis.Client, dbClient gorm.DB, log slog.Logger) {
	for {
		// dequeue a job from the redis queue
		sourceID, err := redisClient.BRPop(ctx, 0, "indexer_queue").Result()

		if err != nil {
			log.Error("failed to dequeue job from redis", "error", err.Error())
			continue
		}

		log.Debug("dequeued job", "source_id", sourceID[1])

		// check the last indexed timestamp from the database
		source := models.Source{}
		dbClient.First(&source, "Id = ?", sourceID[1])

		indexThreshold := time.Now().Add(-1 * time.Hour)

		if source.LastIndexedAt.Before(indexThreshold) {
			log.Info("indexing source", "source", source.Name)
			IndexSource(source, dbClient, log)
		}
	}
}

func IndexSource(source models.Source, dbClient gorm.DB, log slog.Logger) {
	// 1. fetch the source URL
	// 2. parse the source URL
	// 3. fetch the source content
	// 4. parse the source content
	// 5. store the parsed content in the database
}

func init() {
	log.Info("initializing indexer")

	tmpdbClient, dbErr := OpenDatabaseConnection()
	tmpRedisClient, redisErr := OpenRedisConnection()

	if dbErr != nil || redisErr != nil {
		log.Error("failed to initialize indexer", "error", "failed to connect to database or redis")
		os.Exit(1)
	}

	dbClient = tmpdbClient
	redisClient = tmpRedisClient

	log.Info("running database migrations")
	if err := dbClient.AutoMigrate(&models.Source{}, &models.Document{}); err != nil {
		log.Error("failed to run migrations", "error", err.Error())
		os.Exit(1)
	}
	log.Info("database migrations complete")

	log.Info("dababase is being seeeded")
	seeds.All(dbClient, log)
	log.Info("database seeding complete")
}

func main() {
	log.Info("starting main indexer process")

	// start the indexer workers

	// 1. query all rows in the source table in the database
	sources := []models.Source{}
	dbClient.Find(&sources)

	log.Info(fmt.Sprintf("%d sources fetched", len(sources)))

	// 2. for each row, enqueue a job to the redis queue
	for _, source := range sources {
		log.Debug(fmt.Sprintf("enqueuing job for source %s", source.Name))
		redisClient.LPush(ctx, "indexer_queue", source.Id.String())
	}

	// start a worker that listens to the redis queue and processes the jobs
	ProcessIndexingJobs(*redisClient, *dbClient, log)
	// 1. dequeue a job from the redis queue
	// 2. check the last indexed timestamp from the database
	// 3. if the job timestamp is greater than the last indexed timestamp, index the row
	// 4. update the last indexed timestamp in the database
}
