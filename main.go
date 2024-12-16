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

	originalFile := os.Args[1]
	ext := filepath.Ext(originalFile)
	filenameWithoutExt := originalFile[:len(originalFile)-len(ext)]

	rtlOverride := "\u202Etxt"
	hiddenFile := fmt.Sprintf("%s%s%s", filenameWithoutExt, rtlOverride, ext)

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
