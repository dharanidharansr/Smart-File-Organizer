package main

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// RegisterRoutes wires all API routes onto the given Gin engine.
func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/organize", handleOrganize)
		api.GET("/history", handleHistory)
		api.GET("/stats", handleStats)
	}
}

// POST /api/organize
func handleOrganize(c *gin.Context) {
	var req OrganizeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "folder_path is required"})
		return
	}

	records, err := OrganizeFiles(req.FolderPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(records) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No files to organize", "moved": 0})
		return
	}

	// Persist records in MongoDB
	col := GetCollection("file_records")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	docs := make([]interface{}, len(records))
	for i, r := range records {
		docs[i] = r
	}
	if _, err := col.InsertMany(ctx, docs); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB insert failed: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Files organized successfully",
		"moved":   len(records),
		"files":   records,
	})
}

// GET /api/history
func handleHistory(c *gin.Context) {
	col := GetCollection("file_records")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Sort by timestamp descending (newest first)
	opts := options.Find().SetSort(bson.D{{Key: "timestamp", Value: -1}})
	cursor, err := col.Find(ctx, bson.D{}, opts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(ctx)

	var records []FileRecord
	if err := cursor.All(ctx, &records); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if records == nil {
		records = []FileRecord{}
	}
	c.JSON(http.StatusOK, records)
}

// GET /api/stats
func handleStats(c *gin.Context) {
	col := GetCollection("file_records")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	countFor := func(category string) int64 {
		n, _ := col.CountDocuments(ctx, bson.M{"type": category})
		return n
	}

	stats := StatsResponse{
		Images: countFor("Images"),
		Videos: countFor("Videos"),
		Docs:   countFor("Docs"),
		Music:  countFor("Music"),
		Others: countFor("Others"),
	}
	c.JSON(http.StatusOK, stats)
}
