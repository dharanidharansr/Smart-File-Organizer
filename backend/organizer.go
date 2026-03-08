package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// categoryFor maps a file extension to a category folder name.
// Returns the category name and the destination sub-folder.
func categoryFor(ext string) (category string, folder string) {
	switch strings.ToLower(ext) {
	case ".png", ".jpg", ".jpeg":
		return "Images", "Images"
	case ".mp4", ".mov", ".avi", ".amv":
		return "Videos", "Videos"
	case ".pdf", ".docx", ".csv", ".xlsx":
		return "Docs", "Docs"
	case ".mp3", ".wav", ".aac":
		return "Music", "Music"
	default:
		return "Others", "Others"
	}
}

// ensureSubFolders creates the five default sub-folders inside targetFolder.
func ensureSubFolders(targetFolder string) error {
	subFolders := []string{"Images", "Videos", "Docs", "Music", "Others"}
	for _, sf := range subFolders {
		path := filepath.Join(targetFolder, sf)
		if err := os.MkdirAll(path, 0755); err != nil {
			return fmt.Errorf("could not create sub-folder %s: %w", sf, err)
		}
	}
	return nil
}

// OrganizeFiles scans targetFolder, moves files to sub-folders and returns
// a slice of FileRecord ready for insertion into MongoDB.
func OrganizeFiles(targetFolder string) ([]FileRecord, error) {
	// Validate folder
	info, err := os.Stat(targetFolder)
	if err != nil {
		return nil, fmt.Errorf("folder not found: %w", err)
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("%s is not a directory", targetFolder)
	}

	// Create sub-folders
	if err := ensureSubFolders(targetFolder); err != nil {
		return nil, err
	}

	// Read top-level entries
	entries, err := os.ReadDir(targetFolder)
	if err != nil {
		return nil, fmt.Errorf("could not read folder: %w", err)
	}

	var records []FileRecord

	for _, entry := range entries {
		// Skip directories
		if entry.IsDir() {
			continue
		}

		name := entry.Name()
		ext := filepath.Ext(name)
		category, folder := categoryFor(ext)

		oldPath := filepath.Join(targetFolder, name)
		newPath := filepath.Join(targetFolder, folder, name)

		// Move the file
		if err := os.Rename(oldPath, newPath); err != nil {
			return nil, fmt.Errorf("failed to move %s: %w", name, err)
		}

		records = append(records, FileRecord{
			Filename:  name,
			Type:      category,
			OldPath:   oldPath,
			NewPath:   newPath,
			Timestamp: time.Now(),
		})
	}

	return records, nil
}
