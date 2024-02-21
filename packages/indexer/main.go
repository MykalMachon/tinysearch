package main

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/redis/go-redis/v9"
	uuid "github.com/satori/go.uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var redisClient *redis.Client
var db *gorm.DB

var ctx context.Context
var log slog.Logger

type Source struct {
	gorm.Model
	Id            uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name          string
	Description   string
	Url           string    `gorm:"unique"`
	LastIndexedAt time.Time `gorm:"default:null"`
}

type Document struct {
	gorm.Model
	Id        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	SourceID  uint      `gorm:"type:uuid REFERENCES sources(id)"`
	Source    Source    `gorm:"foreignKey:SourceID;AssociationForeignKey:ID"`
	Title     string
	Content   string
	CreatedAt time.Time `gorm:"default:now()"`
	UpdatedAt time.Time `gorm:"default:now()"`
}

func init() {
	// do setup here... this is always run before main
	// * initialize the logger
	jsonHandler := slog.NewJSONHandler(os.Stdout, nil)
	log = *slog.New(jsonHandler)

	log.Info("initializing indexer")

	// * load environment variables
	if err := godotenv.Load(); err != nil {
		log.Error("failed to load environment variables", "error", err.Error())
		os.Exit(1)
	}

	// * initialize the postgresql database and run migrations
	newDb, err := gorm.Open(postgres.New(postgres.Config{
		DSN: os.Getenv("DATABASE_URL"),
	}), &gorm.Config{})

	if err != nil {
		log.Error("failed to connect to database", "error", err.Error())
		os.Exit(1)
	}

	db = newDb

	log.Info("running migrations")
	if err := db.AutoMigrate(&Source{}, &Document{}); err != nil {
		log.Error("failed to run migrations", "error", err.Error())
		os.Exit(1)
	}
	log.Info("migrations complete")

	// * initialize the redis queue and make sure it works
	ctx = context.Background()
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
