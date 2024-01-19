package AdvancedImageCatalogingSoftware

import (
	"database/sql"
	"fmt"
	"os"
	"time"
)

func isDirectory(path string) bool {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file")
		return false
	}
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Error getting file info")
		return false
	}
	return fileInfo.IsDir()
}

func isFile(path string) bool {
	return !isDirectory(path)
}

func fileExists(path string) bool {
	_, err := os.Open(path)
	if err != nil {
		return false
	}
	return true
}

func acessingFiles(path string, userDb *sql.DB) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Error getting file info:", err)
		return
	}

	if fileInfo.IsDir() {
		acessingFiles(path, userDb)
	} else {
		if tagFile(userDb, path) != nil {
			fmt.Println("Error tagging file", file, time.Now())
			return
		}
	}
}

func tagFile(db *sql.DB, filePath string) error {
	tags := getTags(filePath)
	for _, tag := range tags {
		err := insertTag(db, tag, filePath)
		if err != nil {
			return fmt.Errorf("error inserting tag '%s' for file '%s': %v", tag, filePath, err)
		}
	}
	return nil
}

func getTags(filePath string) []string {
	// Mock function to return an array of tags
	// Replace this with your actual implementation
	return []string{"tag1", "tag2", "tag3"}
}
