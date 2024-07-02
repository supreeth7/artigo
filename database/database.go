package database

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Client     *mongo.Client
	Collection *mongo.Collection
	name       string
}

var DB Database

// Init initializes the mongoDB connection
func Init() (err error) {
	if err = godotenv.Load(); err != nil {
		return errors.New(".env file is missing")
	}

	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		return errors.New("please set the MONGO_URI environment variable")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	DB.Client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	if database := os.Getenv("DB_NAME"); database != "" {
		DB.name = database
	} else {
		return errors.New("please select the databse via DB_NAME environment variable")
	}

	DB.Collection = DB.Client.Database(DB.name).Collection("articles")

	log.Println("Successfully connected to database.")
	return nil
}

func Close() {
	if err := DB.Client.Disconnect(context.TODO()); err != nil {
		log.Fatal(err)
	}
	log.Println("Database connection closed.")
}
