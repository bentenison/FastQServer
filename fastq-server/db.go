package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type dataSources struct {
	DB *sql.DB
	// RedisClient *redis.Client
	// StorageClient *storage.Client
}

// InitDS establishes connections to fields in dataSources
func initDS() (*dataSources, error) {
	log.Printf("Initializing data sources\n")
	err := godotenv.Load("./.dev.env")
	if err != nil {
		log.Fatal("error reading the env file: ", err)
	}
	// load env variables - we could pass these in,
	// but this is sort of just a top-level (main package)
	// helper function, so I'll just read them in here
	sqlHost := os.Getenv("SQL_HOST")
	// fmt.Println(sqlHost)
	sqlPort := os.Getenv("SQL_PORT")
	sqlUsername := os.Getenv("SQL_USERNAME")
	sqlPassword := os.Getenv("SQL_PASSWORD")
	sqlDB := os.Getenv("SQL_DB")
	// "wsl_root:password@tcp(192.168.48.1:3306)/hrd"

	sqlDb, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", sqlUsername, sqlPassword, sqlHost, sqlPort, sqlDB))
	if err != nil {
		log.Fatal("cant connect to mysql : ", err)
	}

	sqlDb.SetConnMaxLifetime(3 * time.Minute)
	sqlDb.SetMaxOpenConns(10)
	sqlDb.SetMaxIdleConns(10)
	if err := sqlDb.Ping(); err != nil {
		log.Fatal("error connecting to db: %w", err)
	}
	log.Println("connected to sql db successfully")
	// Initialize redis connection
	// redisHost := os.Getenv("REDIS_HOST")
	// redisPort := os.Getenv("REDIS_PORT")

	// log.Printf("Connecting to Redis\n")
	// rdb := redis.NewClient(&redis.Options{
	// 	Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
	// 	Password: "",
	// 	DB:       0,
	// })

	// // verify redis connection

	// _, err = rdb.Ping(context.Background()).Result()

	// if err != nil {
	// 	return nil, fmt.Errorf("error connecting to redis: %w", err)
	// }

	// Initialize google storage client
	// log.Printf("Connecting to Cloud Storage\n")
	// ctx := context.Background()
	// ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	// defer cancel() // releases resources if slowOperation completes before timeout elapses
	// storage, err := storage.NewClient(ctx)

	// if err != nil {
	// 	return nil, fmt.Errorf("error creating cloud storage client: %w", err)
	// }

	return &dataSources{
		DB: sqlDb,
		// RedisClient: rdb,
		// StorageClient: storage,
	}, nil
}

// close to be used in graceful server shutdown
func (d *dataSources) close() error {
	if err := d.DB.Close(); err != nil {
		return fmt.Errorf("error closing sql: %w", err)
	}

	// if err := d.RedisClient.Close(); err != nil {
	// 	return fmt.Errorf("error closing Redis Client: %w", err)
	// }

	return nil
}
