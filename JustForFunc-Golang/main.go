package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ConvertProtoJsonFile converts a .protojson file to .jsonld file in the same location
// Just copy-paste this function into your code - needs encoding/json, fmt, os, path/filepath, strings
//
// Usage:
//
//	err := ConvertProtoJsonFile("/path/to/file.protojson")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	// Creates /path/to/file.jsonld

func main() {

	fmt.Println(ConvertProtoJsonFile("./urn_8PU9000GOI0063.protojson"))
}
func ConvertProtoJsonFile(inputPath string) error {
	// Read the protojson file
	protoJsonData, err := os.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	// Convert to JSON-LD
	jsonLdData, err := ProtoJsonToJsonLd(protoJsonData)
	if err != nil {
		return fmt.Errorf("failed to convert: %w", err)
	}

	// Create output path: replace .protojson with .jsonld
	outputPath := strings.TrimSuffix(inputPath, filepath.Ext(inputPath)) + ".jsonld"

	// Write the JSON-LD file
	if err := os.WriteFile(outputPath, jsonLdData, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

func ProtoJsonToJsonLd(protoJsonData []byte) ([]byte, error) {
	// Parse protojson to extract the base64-encoded jsonld field
	var temp struct {
		Jsonld string `json:"jsonld"`
	}

	if err := json.Unmarshal(protoJsonData, &temp); err != nil {
		return nil, fmt.Errorf("failed to parse protojson: %w", err)
	}

	if temp.Jsonld == "" {
		return nil, fmt.Errorf("no jsonld field found in protojson")
	}

	// Base64-decode the jsonld field
	decoded, err := base64.StdEncoding.DecodeString(temp.Jsonld)
	if err != nil {
		return nil, fmt.Errorf("failed to base64-decode jsonld: %w", err)
	}

	// Pretty-print the decoded JSON-LD without escaping &, <, >
	var prettyJSON interface{}
	if err := json.Unmarshal(decoded, &prettyJSON); err != nil {
		return decoded, nil
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	if err := enc.Encode(prettyJSON); err != nil {
		return nil, fmt.Errorf("failed to encode json-ld: %w", err)
	}
	return bytes.TrimRight(buf.Bytes(), "\n"), nil
}
