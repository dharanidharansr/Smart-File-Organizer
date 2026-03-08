package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const (
	mongoURI   = "mongodb://localhost:27017"
	dbName     = "smart_file_organizer"
	collection = "file_history"
)

// DB holds the MongoDB collection handle.
type DB struct {
	col *mongo.Collection
}

// NewDB connects to local MongoDB and returns a DB handle.
// Returns nil (with a log) if MongoDB is not reachable so the app still works.
func NewDB() *DB {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Printf("[MongoDB] connect error: %v — running without persistence", err)
		return nil
	}

	// Ping to verify connection
	if err = client.Ping(ctx, nil); err != nil {
		log.Printf("[MongoDB] ping failed: %v — running without persistence", err)
		return nil
	}

	log.Println("[MongoDB] connected to", mongoURI)
	col := client.Database(dbName).Collection(collection)
	return &DB{col: col}
}

// InsertRecords inserts a batch of FileRecords into MongoDB.
func (db *DB) InsertRecords(records []FileRecord) {
	if db == nil || len(records) == 0 {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	docs := make([]interface{}, len(records))
	for i, r := range records {
		docs[i] = r
	}
	_, err := db.col.InsertMany(ctx, docs)
	if err != nil {
		log.Printf("[MongoDB] InsertMany error: %v", err)
	}
}

// LoadAllRecords fetches all records from MongoDB, newest first.
func (db *DB) LoadAllRecords() []FileRecord {
	if db == nil {
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	opts := options.Find().SetSort(bson.D{{Key: "timestamp", Value: -1}})
	cursor, err := db.col.Find(ctx, bson.D{}, opts)
	if err != nil {
		log.Printf("[MongoDB] Find error: %v", err)
		return nil
	}
	defer cursor.Close(ctx)

	var records []FileRecord
	if err = cursor.All(ctx, &records); err != nil {
		log.Printf("[MongoDB] cursor decode error: %v", err)
		return nil
	}
	return records
}

// LoadStats recomputes stats by counting each category in MongoDB.
func (db *DB) LoadStats() StatsResponse {
	if db == nil {
		return StatsResponse{}
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	count := func(cat string) int {
		n, err := db.col.CountDocuments(ctx, bson.D{{Key: "type", Value: cat}})
		if err != nil {
			return 0
		}
		return int(n)
	}

	return StatsResponse{
		Images: count("Images"),
		Videos: count("Videos"),
		Docs:   count("Docs"),
		Music:  count("Music"),
		Others: count("Others"),
	}
}
