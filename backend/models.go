package main

import "time"

// FileRecord represents a moved file stored in MongoDB
type FileRecord struct {
	ID        interface{} `bson:"_id,omitempty"  json:"id,omitempty"`
	Filename  string      `bson:"filename"       json:"filename"`
	Type      string      `bson:"type"           json:"type"`
	OldPath   string      `bson:"old_path"       json:"old_path"`
	NewPath   string      `bson:"new_path"       json:"new_path"`
	Timestamp time.Time   `bson:"timestamp"      json:"timestamp"`
}

// OrganizeRequest is the body for POST /organize
type OrganizeRequest struct {
	FolderPath string `json:"folder_path" binding:"required"`
}

// StatsResponse holds per-category counts
type StatsResponse struct {
	Images int64 `json:"images"`
	Videos int64 `json:"videos"`
	Docs   int64 `json:"docs"`
	Music  int64 `json:"music"`
	Others int64 `json:"others"`
}
