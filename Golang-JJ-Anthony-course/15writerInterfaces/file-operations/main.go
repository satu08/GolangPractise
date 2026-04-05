package main

import (
	"fmt"
	"os"
)

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		// File exists
		return true
	}
	if os.IsNotExist(err) {
		// File does not exist
		return false
	}
	// Some other error occurred
	fmt.Println("Error checking file:", err)
	return false
}

func main() {
	path := "../output.txt"
	if fileExists(path) {
		fmt.Println("File exists at", path)
	} else {
		fmt.Println("File does not exist at", path)
	}
}
