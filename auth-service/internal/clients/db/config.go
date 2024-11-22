package db

import (
	"fmt"
	"os"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/logger"
	"go.uber.org/zap"
)

const (
	// PgDsn is the environment variable key for the PostgreSQL Data Source Name (DSN).
	// It should be used to fetch the DSN from the environment, typically specified in the .env file.
	PgDsn = "PG_DSN"
)

// PGConfig defines an interface for obtaining the database connection string (DSN).
type PGConfig interface {
	GetDSN() string
}

// pgConfig implements the PGConfig interface, storing the DSN.
type pgConfig struct {
	dsn string
}

// NewPgConfig creates a new PGConfig instance by reading the DSN from environment variables.
// It returns an error if the DSN is not set.
func NewPgConfig() (PGConfig, error) {
	dsn := os.Getenv(PgDsn)
	if len(dsn) == 0 {
		logger.Error("failed to get db dsn", zap.String("db dsn", PgDsn))
		return nil, fmt.Errorf("env %s is empty", PgDsn)
	}

	return &pgConfig{
		dsn: dsn,
	}, nil
}

// GetDSN returns the database connection string (DSN) from the pgConfig instance.
func (p *pgConfig) GetDSN() string {
	return p.dsn
}
