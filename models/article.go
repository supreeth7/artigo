package models

import (
	"context"
	"time"

	"github.com/supreeth7/artigo/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Article struct {
	ID       primitive.ObjectID `json:"id,omitempty"`
	Title    string             `json:"title,omitempty" validate:"required"`
	DateTime time.Time          `json:"dateTime,omitempty" validate:"required"`
	Author   string             `json:"author,omitempty" validate:"required"`
	Content  string             `json:"content,omitempty" validate:"required"`
	Likes    int                `json:"likes,omitempty"`
}

var articles = []Article{}

// Create adds a new article to the collection
func (a Article) Create(db *database.Database) (*mongo.InsertOneResult, error) {
	return db.Collection.InsertOne(context.TODO(), a)
}

func Get(db *database.Database) ([]Article, error) {
	return nil, nil
}

// GetArticles fetches all the articles from the database
func GetArticles() []Article {
	return articles
}
