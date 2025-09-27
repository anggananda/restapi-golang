package models

import "time"

type PrefillEntry struct {
	Key   string
	Value string
	TTL   time.Duration
}
