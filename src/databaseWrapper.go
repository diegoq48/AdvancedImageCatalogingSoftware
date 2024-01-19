package AdvancedImageCatalogingSoftware

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func connectDB(dbName string) *sql.DB {
	db, err := sql.Open("postgres", "user=postgres password=postgres dbname="+dbName+" sslmode=disable")
	if err != nil {
		panic(err)
	}
	return db
}

func insertTag(db *sql.DB, tag string, filePath string) error {
	// First, check if the tag-filePath pair already exists
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM file_tags WHERE tag = $1 AND file_path = $2)", tag, filePath).Scan(&exists)
	if err != nil {
		return fmt.Errorf("error checking existence: %v", err)
	}

	// If the pair exists, return an error
	if exists {
		return fmt.Errorf("duplicate entry: the tag-filePath pair already exists")
	}

	// If the pair does not exist, proceed to insert
	_, err = db.Exec("INSERT INTO file_tags (tag, file_path) VALUES ($1, $2)", tag, filePath)
	return err
}

func getFilesByTag(db *sql.DB, tag string) ([]string, error) {
	rows, err := db.Query("SELECT file_path FROM file_tags WHERE tag = $1", tag)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var filePaths []string
	for rows.Next() {
		var filePath string
		if err := rows.Scan(&filePath); err != nil {
			return nil, err
		}
		filePaths = append(filePaths, filePath)
	}
	return filePaths, nil
}

func removeTag(db *sql.DB, tag string, filePath string) error {
	_, err := db.Exec("DELETE FROM file_tags WHERE tag = $1 AND file_path = $2", tag, filePath)
	return err
}

func updateFilePath(db *sql.DB, tag string, oldFilePath string, newFilePath string) error {
	_, err := db.Exec("UPDATE file_tags SET file_path = $1 WHERE tag = $2 AND file_path = $3", newFilePath, tag, oldFilePath)
	return err
}

func removeAllTagsForFile(db *sql.DB, filePath string) error {
	_, err := db.Exec("DELETE FROM file_tags WHERE file_path = $1", filePath)
	return err
}

func addOrUpdateTag(db *sql.DB, tag string, filePath string) error {
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM file_tags WHERE tag = $1 AND file_path = $2)", tag, filePath).Scan(&exists)
	if err != nil {
		return err
	}

	if exists {
		_, err = db.Exec("UPDATE file_tags SET file_path = $1 WHERE tag = $2", filePath, tag)
	} else {
		_, err = db.Exec("INSERT INTO file_tags (tag, file_path) VALUES ($1, $2)", tag, filePath)
	}
	return err
}
