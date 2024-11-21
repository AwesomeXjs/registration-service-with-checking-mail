package redis

import (
	"context"
	"time"
)

// IRedis defines the methods for interacting with a Redis data store.
// It provides an abstraction for caching and retrieving data using key-value pairs.
type IRedis interface {
	// Set stores a value in Redis with a specified key and duration.
	// The duration defines how long the value should be retained before expiring.
	Set(ctx context.Context, key string, value string, duration time.Duration) error

	// Get retrieves the value associated with the specified key from Redis.
	// It returns the value as a string or an error if the key does not exist or an issue occurs.
	Get(ctx context.Context, key string) (string, error)

	// SetObject stores a structured object in Redis with a specified key and duration.
	// The object is serialized before storage, and the duration defines its expiration time.
	SetObject(ctx context.Context, key string, value interface{}, duration time.Duration) error

	// GetObject retrieves a structured object associated with the specified key from Redis.
	// The retrieved data is deserialized into the provided value parameter,
	// and an error is returned if the key does not exist or deserialization fails.
	GetObject(ctx context.Context, key string, value any) error

	// Delete removes the specified key and its associated value from Redis.
	// If the key does not exist or an error occurs during the deletion, an error is returned.
	Delete(ctx context.Context, key string) error
}
