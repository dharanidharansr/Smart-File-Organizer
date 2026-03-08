package main

// FileRecord represents a single moved file.
type FileRecord struct {
	Filename  string `json:"filename"  bson:"filename"`
	Type      string `json:"type"      bson:"type"`
	OldPath   string `json:"old_path"  bson:"old_path"`
	NewPath   string `json:"new_path"  bson:"new_path"`
	Timestamp string `json:"timestamp" bson:"timestamp"`
}

// OrganizeResult is returned by App.OrganizeFiles.
type OrganizeResult struct {
	Success bool         `json:"success"`
	Error   string       `json:"error,omitempty"`
	Moved   int          `json:"moved"`
	Files   []FileRecord `json:"files"`
}

// StatsResponse holds per-category totals.
type StatsResponse struct {
	Images int `json:"images"`
	Videos int `json:"videos"`
	Docs   int `json:"docs"`
	Music  int `json:"music"`
	Others int `json:"others"`
}
