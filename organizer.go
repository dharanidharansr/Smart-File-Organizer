package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// categoryFor maps a file extension to its category name and destination folder.
func categoryFor(ext string) (category, folder string) {
	switch strings.ToLower(ext) {
	case ".png", ".jpg", ".jpeg", ".gif", ".webp", ".bmp", ".tiff", ".svg":
		return "Images", "Images"
	case ".mp4", ".mov", ".avi", ".mkv", ".wmv", ".flv", ".webm", ".m4v":
		return "Videos", "Videos"
	case ".pdf", ".docx", ".doc", ".csv", ".xlsx", ".xls", ".pptx", ".ppt", ".txt", ".rtf":
		return "Docs", "Docs"
	case ".mp3", ".wav", ".aac", ".flac", ".ogg", ".m4a", ".wma":
		return "Music", "Music"
	default:
		return "Others", "Others"
	}
}

// ensureSubFolders creates the five default sub-folders inside targetFolder.
func ensureSubFolders(targetFolder string) error {
	for _, sf := range []string{"Images", "Videos", "Docs", "Music", "Others"} {
		if err := os.MkdirAll(filepath.Join(targetFolder, sf), 0755); err != nil {
			return fmt.Errorf("could not create sub-folder %s: %w", sf, err)
		}
	}
	return nil
}

// organizeFiles scans the top level of targetFolder, moves each file into the
// appropriate sub-folder and returns the list of FileRecords.
func organizeFiles(targetFolder string) ([]FileRecord, error) {
	info, err := os.Stat(targetFolder)
	if err != nil {
		return nil, fmt.Errorf("folder not found: %w", err)
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("%q is not a directory", targetFolder)
	}

	if err := ensureSubFolders(targetFolder); err != nil {
		return nil, err
	}

	entries, err := os.ReadDir(targetFolder)
	if err != nil {
		return nil, fmt.Errorf("could not read folder: %w", err)
	}

	var records []FileRecord
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		name := entry.Name()
		ext := filepath.Ext(name)
		category, folder := categoryFor(ext)

		oldPath := filepath.Join(targetFolder, name)
		newPath := filepath.Join(targetFolder, folder, name)

		if err := os.Rename(oldPath, newPath); err != nil {
			return nil, fmt.Errorf("failed to move %q: %w", name, err)
		}

		records = append(records, FileRecord{
			Filename:  name,
			Type:      category,
			OldPath:   oldPath,
			NewPath:   newPath,
			Timestamp: time.Now().Format(time.RFC3339),
		})
	}

	return records, nil
}
