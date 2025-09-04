package utils

// ConvertToInt64 safely converts interface{} to int64
func ConvertToInt64(val interface{}) int64 {
	switch v := val.(type) {
	case int32:
		return int64(v)
	case int64:
		return v
	case int:
		return int64(v)
	case float64:
		return int64(v)
	default:
		return 0
	}
}
