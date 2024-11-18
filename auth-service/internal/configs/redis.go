package configs

import (
	"fmt"
	"net"
	"os"
)

// Constants for environment variable names for Redis configuration.
const (
	RedisPortEnvName = "REDIS_PORT"
	RedisHostEnvName = "REDIS_HOST"
)

// RedisConfig defines an interface for obtaining the Redis address.
type RedisConfig interface {
	Address() string
}

// redisConfig holds the host and port information for connecting to Redis.
type redisConfig struct {
	host string
	port string
}

// NewRedisConfig creates a new RedisConfig instance by reading environment variables.
func NewRedisConfig() (RedisConfig, error) {
	port := os.Getenv(RedisPortEnvName)
	host := os.Getenv(RedisHostEnvName)
	if len(port) == 0 || len(host) == 0 {
		return nil, fmt.Errorf("REDIS_PORT or REDIS_HOST is not set")
	}

	return &redisConfig{
		host: host,
		port: port,
	}, nil
}

// Address returns the complete address for connecting to Redis, combining host and port.
func (m *redisConfig) Address() string {
	return net.JoinHostPort(m.host, m.port)
}
