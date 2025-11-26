package mocks

import "fmt"

func fallbackMock(key string) map[string]any {
	return map[string]any{
		"status":  "error",
		"message": fmt.Sprintf("Failed to load YAML mock for %s", key),
		"error":   "MOCK_LOAD_ERROR",
		"datas":   []map[string]any{},
	}
}
