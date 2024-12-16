package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		return
	}

	// Original file name
	originalFile := os.Args[1]
	ext := filepath.Ext(originalFile)
	filenameWithoutExt := originalFile[:len(originalFile)-len(ext)]

	// Target file name with the Right-to-Left Override (U+202E) character
	// This will make the file appear as "example.txt" but keeps the actual extension ".exe"
	rtlOverride := "\u202Etxt"
	hiddenFile := fmt.Sprintf("%s%s%s", filenameWithoutExt, rtlOverride, ext)

	// Rename the file
	err := os.Rename(originalFile, hiddenFile)
	if err != nil {
		fmt.Printf("Failed to rename file: %v\n", err)
		return
	}

	err = os.Chmod(hiddenFile, 0755)
	if err != nil {
		fmt.Printf("Failed to change file permissions: %v\n", err)
		return
	}

	fmt.Printf("File renamed successfully to: %s\n", hiddenFile)
}
