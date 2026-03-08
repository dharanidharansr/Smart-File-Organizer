package main

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct – the central Wails application object.
// Every exported method here becomes callable from the frontend via
//   window.go.main.App.MethodName(...)
type App struct {
	ctx     context.Context
	history []FileRecord
	stats   StatsResponse
	db      *DB
}

// NewApp creates a new App instance.
func NewApp() *App {
	db := NewDB()
	app := &App{db: db}
	// Load persisted history and stats from MongoDB on startup
	if db != nil {
		app.history = db.LoadAllRecords()
		app.stats = db.LoadStats()
	}
	return app
}

// startup is called when the app starts. The context is saved so we
// can call runtime methods later (e.g. OpenDirectoryDialog).
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// -----------------------------------------------------------------
// SelectFolder opens the native OS folder picker and returns the path.
// -----------------------------------------------------------------
func (a *App) SelectFolder() string {
	path, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select a folder to organize",
	})
	if err != nil {
		return ""
	}
	return path
}

// -----------------------------------------------------------------
// OrganizeFiles organizes all files in folderPath, persists the records
// in memory and returns a summary.
// -----------------------------------------------------------------
func (a *App) OrganizeFiles(folderPath string) OrganizeResult {
	records, err := organizeFiles(folderPath)
	if err != nil {
		return OrganizeResult{Success: false, Error: err.Error()}
	}

	// Persist to MongoDB
	a.db.InsertRecords(records)

	// Prepend new records so the most recent appear first in history
	a.history = append(records, a.history...)

	// Rebuild stats from full history (or reload from DB if available)
	if a.db != nil {
		a.stats = a.db.LoadStats()
	} else {
		a.rebuildStats()
	}

	return OrganizeResult{
		Success: true,
		Moved:   len(records),
		Files:   records,
	}
}

// -----------------------------------------------------------------
// GetStats returns aggregate counts per category.
// -----------------------------------------------------------------
func (a *App) GetStats() StatsResponse {
	return a.stats
}

// -----------------------------------------------------------------
// GetHistory returns all file records (newest first).
// -----------------------------------------------------------------
func (a *App) GetHistory() []FileRecord {
	if a.history == nil {
		return []FileRecord{}
	}
	return a.history
}

// -----------------------------------------------------------------
// internal: rebuild stats from in-memory history
// -----------------------------------------------------------------
func (a *App) rebuildStats() {
	s := StatsResponse{}
	for _, r := range a.history {
		switch r.Type {
		case "Images":
			s.Images++
		case "Videos":
			s.Videos++
		case "Docs":
			s.Docs++
		case "Music":
			s.Music++
		default:
			s.Others++
		}
	}
	a.stats = s
}
