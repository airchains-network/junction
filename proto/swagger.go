package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	inputDir   = "docs/client/"            // Search all JSON files inside this directory
	outputFile = "docs/static/openapi.yml" // Output OpenAPI YAML file
)

// getVersion executes git describe --tags and returns the version
func getVersion() string {
	cmd := exec.Command("git", "describe", "--tags")
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("Warning: Could not get git version: %v\n", err)
	}
	return strings.TrimSpace(string(out))
}

// OpenAPI represents the Swagger 2.0 structure
type OpenAPI struct {
	Swagger     string                 `json:"swagger" yaml:"swagger"`
	Info        map[string]interface{} `json:"info,omitempty" yaml:"info,omitempty"`
	Consumes    []string               `json:"consumes,omitempty" yaml:"consumes,omitempty"`
	Produces    []string               `json:"produces,omitempty" yaml:"produces,omitempty"`
	Paths       map[string]interface{} `json:"paths,omitempty" yaml:"paths,omitempty"`
	Definitions map[string]interface{} `json:"definitions,omitempty" yaml:"definitions,omitempty"`
}

// MergeSwaggerFiles reads all JSON files from `docs/client/` recursively and merges them into a single Swagger structure
func MergeSwaggerFiles() (*OpenAPI, error) {
	merged := &OpenAPI{
		Swagger: "2.0",
		Info: map[string]interface{}{
			"title":   "Junction OpenAPI",
			"version": getVersion(),
		},
		Consumes:    []string{"application/json"},
		Produces:    []string{"application/json"},
		Paths:       make(map[string]interface{}),
		Definitions: make(map[string]interface{}),
	}

	err := filepath.WalkDir(inputDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || filepath.Ext(path) != ".json" {
			return nil // Skip directories and non-JSON files
		}

		// Read JSON file
		jsonData, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("error reading file %s: %v", path, err)
		}

		// Parse JSON into a temporary structure
		var temp OpenAPI
		if err := json.Unmarshal(jsonData, &temp); err != nil {
			return fmt.Errorf("error unmarshaling JSON %s: %v", path, err)
		}

		// Merge Paths, excluding /cosmos/tx/v1beta1/txs
		for k, v := range temp.Paths {
			if k != "/cosmos/tx/v1beta1/txs" {
				merged.Paths[k] = v
			}
		}

		// Merge Definitions
		for k, v := range temp.Definitions {
			merged.Definitions[k] = v
		}

		fmt.Printf("%s\n", path)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return merged, nil
}

func main() {
	merged, err := MergeSwaggerFiles()
	if err != nil {
		fmt.Printf("Error merging Swagger files: %v\n", err)
		os.Exit(1)
	}

	// Convert to YAML
	yamlData, err := yaml.Marshal(merged)
	if err != nil {
		fmt.Printf("Error marshaling YAML: %v\n", err)
		os.Exit(1)
	}

	// Ensure output directory exists
	outputDir := filepath.Dir(outputFile)
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		fmt.Printf("Error creating output directory: %v\n", err)
		os.Exit(1)
	}

	// Write to file
	if err := os.WriteFile(outputFile, yamlData, 0644); err != nil {
		fmt.Printf("Error writing YAML file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Merged Swagger 2.0 YAML file created: %s\n", outputFile)
}
