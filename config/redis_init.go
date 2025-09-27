package config

import (
	"context"
	"errors"
	"log"
	"restapi-golang/constants"
	"time"

	"github.com/redis/go-redis/v9"
)

func HydrateRedis(rdb *redis.Client) {
	log.Println("[Hydration] Starting hydration process...")

	for _, prefill := range constants.PrefillRedisKeys {
		if err := PrefillRedisKey(rdb, prefill.Key, prefill.Value, prefill.TTL); err != nil {
			log.Fatalf("[Hydration] Failed to prefill key '%s': %v", prefill.Key, err)
		}
	}

	log.Println("[Hydration] Successfully hydrated Redis.")
}

func PrefillRedisKey(rdb *redis.Client, key, value string, ttl time.Duration) error {
	ctx := context.Background()
	const maxRetries = 3
	const retryDelay = 500 * time.Millisecond

	log.Printf("[Hydration] Prefilling key: '%s' with value: '%s' (TTL: %s)\n", key, value, ttl.String())

	var lastErr error
	for i := 0; i < maxRetries; i++ {
		err := rdb.SetNX(ctx, key, value, ttl).Err()
		if err == nil {
			return nil // sukses
		}

		// retry hydration when has temp error in Redis connection
		if isTemporaryRedisError(err) {
			log.Printf("[Hydration] Temporary error on prefill '%s', retrying (%d/%d)...: %v", key, i+1, maxRetries, err)
			time.Sleep(retryDelay)
			lastErr = err
			continue
		}

		return err
	}

	return lastErr
}

func isTemporaryRedisError(err error) bool {
	// simple detection temporary redis error
	if errors.Is(err, context.DeadlineExceeded) || errors.Is(err, context.Canceled) {
		return false
	}
	return true
}
