package config

import (
	"context"
	"log"
	"os"
	"restapi-golang/utils"
	"time"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DB          *mongo.Database
	CASUrl      string
	FrontendUrl string
	HostUrl     string
	RDB         *redis.Client
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("❌ Failed to load .env file")
	}

	CASUrl = os.Getenv("CAS_URL")
	if CASUrl == "" {
		log.Fatal("❌ CAS_URL is not set in .env")
	}

	FrontendUrl = os.Getenv("FRONTEND_URL")
	if FrontendUrl == "" {
		log.Fatal("❌ FRONTEND_URL is not set in .env")
	}

	HostUrl = os.Getenv("HOST_URL")
	if HostUrl == "" {
		log.Fatal("❌ HOST_URL is not set in .env")
	}
}

func ConnectDB() {
	mongoURI := os.Getenv("MONGO_URI")
	dbName := os.Getenv("MONGO_DB")

	clientOptions := options.Client().ApplyURI(mongoURI)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("❌ Failed connect to MongoDB:", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("❌ MongoDB did not respond:", err)
	}

	log.Println("✅ Sucessfully connect to MongoDB")
	DB = client.Database(dbName)
}

func ConnectRedis() {
	port := os.Getenv("REDIS_PORT")
	host := os.Getenv("REDIS_HOST")
	password := os.Getenv("REDIS_PASSWORD")
	db := os.Getenv("REDIS_DB")
	RDB = redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       utils.StringToInt(db, 0),
	})

	ctx := context.Background()

	for i := 0; i < 5; i++ {
		_, err := RDB.Ping(ctx).Result()
		if err == nil {
			log.Println("Redis Connection Successfully!")
			HydrateRedis(RDB)
			return
		}

		log.Printf("Retry connecting to Redis (%d/5): %v\n", i+1, err)
		time.Sleep(2 * time.Second)
	}

	log.Fatal("Failed to connect to Redis after retries.")
}
