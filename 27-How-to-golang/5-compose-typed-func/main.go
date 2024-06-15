package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type fileNameTranformer func(string) string

type server struct {
	fileNameTransformerFunc fileNameTranformer
}

func (s *server) handleRequest(filename string) error {
	fileName := s.fileNameTransformerFunc(filename)
	fmt.Println("new filename: ", fileName)
	return nil
}

func hashFile(filename string) string {
	hash := sha256.Sum256([]byte(filename))
	newFilename := hex.EncodeToString(hash[:])
	return newFilename
}
func PrefixFileName(prefix string) fileNameTranformer {
	return func(filename string) string {
		return prefix + filename
	}
}
func main() {
	s := &server{
		PrefixFileName("Satya_"),
	}

	s.handleRequest("test.txt")
}
