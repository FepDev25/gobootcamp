package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func main() {
	rootDir := flag.String("dir", ".", "Directory to scan for duplicates")
	outputFile := flag.String("output", "duplicates_report.txt", "File to write the report to")
	flag.Parse()

	fmt.Printf("Scanning directory: %s\n", *rootDir)

	// 1. Group files by size
	filesBySize := make(map[int64][]string)

	err := filepath.WalkDir(*rootDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Printf("Error accessing path %q: %v\n", path, err)
			return nil // Continue walking
		}
		if !d.IsDir() {
			info, err := d.Info()
			if err != nil {
				return nil
			}
			size := info.Size()
			// Ignore empty files or very small files if desired, strictly strictly > 0 for now
			if size > 0 {
				filesBySize[size] = append(filesBySize[size], path)
			}
		}
		return nil
	})

	if err != nil {
		log.Fatalf("Error walking the path: %v", err)
	}

	// 2. For files with same size, compare hashes
	duplicates := make(map[string][]string) // hash -> list of paths

	for _, paths := range filesBySize {
		if len(paths) < 2 {
			continue // Unique size, no duplicates possible
		}

		for _, path := range paths {
			hash, err := computeHash(path)
			if err != nil {
				log.Printf("Failed to hash file %s: %v", path, err)
				continue
			}
			duplicates[hash] = append(duplicates[hash], path)
		}
	}

	// 3. Print results and generate report
	found := false
	var reportContent string

	for _, paths := range duplicates {
		if len(paths) > 1 {
			found = true
			fmt.Println("Duplicate group found:")
			reportContent += "Duplicate group found:\n"
			for _, p := range paths {
				fmt.Printf("\t%s\n", p)
				reportContent += fmt.Sprintf("\t%s\n", p)
			}
			fmt.Println("")
			reportContent += "\n"
		}
	}

	if !found {
		fmt.Println("No duplicates found.")
	} else {
		// Write report to file
		err := os.WriteFile(*outputFile, []byte(reportContent), 0644)
		if err != nil {
			log.Printf("Error writing report to %s: %v\n", *outputFile, err)
		} else {
			fmt.Printf("Report saved to %s\n", *outputFile)
		}
	}
}

func computeHash(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}
