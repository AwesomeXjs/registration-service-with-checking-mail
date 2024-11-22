package redis

import (
	"fmt"
	"net"
	"os"
)

// Constants for environment variable names for Redis configuration.
const (
	PortEnvName = "REDIS_PORT"
	HostEnvName = "REDIS_HOST"
)

// IRedisConfig defines an interface for obtaining the Redis address.
type IRedisConfig interface {
	Address() string
}

// redisConfig holds the host and port information for connecting to Redis.
type redisConfig struct {
	host string
	port string
}

// NewRedisConfig creates a new RedisConfig instance by reading environment variables.
func NewRedisConfig() (IRedisConfig, error) {
	port := os.Getenv(PortEnvName)
	host := os.Getenv(HostEnvName)
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
