package mocks

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v3"
)

func LoadYAMLFile(filename string) (map[string]interface{}, error) {
	_, currentFile, _, _ := runtime.Caller(0)
	baseDir := filepath.Join(filepath.Dir(currentFile), "..") // root
	fullPath := filepath.Join(baseDir, filename)

	log.Printf("📂 Loading YAML from: %s", fullPath)

	data, err := os.ReadFile(fullPath)
	if err != nil {
		log.Printf("❌ Failed to read file: %v", err)
		return nil, err
	}

	var result map[string]interface{}
	if err := yaml.Unmarshal(data, &result); err != nil {
		log.Printf("❌ Failed to unmarshal YAML: %v", err)
		return nil, err
	}

	return result, nil
}

func LoadYAMLSection(filename, section string) (map[string]interface{}, error) {
	allData, err := LoadYAMLFile(filename)
	if err != nil {
		return nil, err
	}

	if sectionData, ok := allData[section]; ok {
		// pastikan datanya dalam bentuk map[string]interface{}
		if sectionMap, valid := sectionData.(map[string]interface{}); valid {
			return sectionMap, nil
		}
		return nil, fmt.Errorf("bagian '%s' bukan object YAML", section)
	}

	return nil, fmt.Errorf("bagian '%s' tidak ditemukan di %s", section, filename)
}
