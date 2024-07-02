package models

import (
	"context"
	"time"

	"github.com/supreeth7/artigo/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Article struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title    string             `json:"title,omitempty" validate:"required"`
	DateTime time.Time          `json:"dateTime,omitempty"`
	Author   string             `json:"author,omitempty" validate:"required"`
	Content  string             `json:"content,omitempty" validate:"required"`
	Likes    int                `json:"likes,omitempty"`
}

// Create adds a new article to the collection.
func (a *Article) Create(db *database.Database) (*mongo.InsertOneResult, error) {
	return db.Collection.InsertOne(context.TODO(), a)
}

// GetByID returns the record for the given ID
func (a *Article) GetByID(id string, db *database.Database) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{
		"_id": objID,
	}

	return db.Collection.FindOne(context.TODO(), filter).Decode(a)
}

// Update overhauls the given record in the collection.
func (a *Article) Update(id string, db *database.Database) (*mongo.UpdateResult, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	update := bson.M{
		"$set": bson.M{
			"title":    a.Title,
			"author":   a.Author,
			"dateTime": a.DateTime,
			"content":  a.Content,
			"likes":    a.Likes,
		},
	}

	return db.Collection.UpdateByID(context.TODO(), objID, update)
}

// Delete purges the given record from the collection.
func (a *Article) Delete(id string, db *database.Database) (*mongo.DeleteResult, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"_id": objID,
	}

	return db.Collection.DeleteOne(context.TODO(), filter)
}

// GetArticles fetches all the records from the collection.
func Get(db *database.Database) ([]Article, error) {
	var articles []Article
	cursor, err := db.Collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var article Article
		if err := cursor.Decode(&article); err != nil {
			return nil, err
		}

		articles = append(articles, article)
	}

	return articles, nil
}
