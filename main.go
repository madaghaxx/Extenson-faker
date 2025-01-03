package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <filename> <extension>")
		return
	}

	originalFile := os.Args[1]
	Newext := Reverse(os.Args[2])
	ext := filepath.Ext(originalFile)
	filenameWithoutExt := originalFile[:len(originalFile)-len(ext)]

	rtlOverride := "\u202E" + Newext
	hiddenFile := fmt.Sprintf("%s%s%s", filenameWithoutExt, rtlOverride, ext)

	err := os.Rename(originalFile, hiddenFile)
	if err != nil {
		fmt.Printf("Failed to rename file: %v\n", err)
		return
	}

	err = os.Chmod(hiddenFile, 0o755)
	if err != nil {
		fmt.Printf("Failed to change file permissions: %v\n", err)
		return
	}

	fmt.Printf("File renamed successfully to: %s\n", hiddenFile)
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
