package utils

import (
	"strconv"
)

// StringToInt konversi string ke int dengan default value
func StringToInt(s string, defaultValue int) int {
	if s == "" {
		return defaultValue
	}
	value, err := strconv.Atoi(s)
	if err != nil {
		return defaultValue
	}
	return value
}
