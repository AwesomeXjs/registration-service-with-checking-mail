package configs

import (
	"fmt"
	"os"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/utils/consts"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/utils/logger"
	"go.uber.org/zap"
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
	dsn := os.Getenv(consts.PgDsn)
	if len(dsn) == 0 {
		logger.Error("failed to get db dsn", zap.String("db dsn", consts.PgDsn))
		return nil, fmt.Errorf("env %s is empty", consts.PgDsn)
	}

	return &pgConfig{
		dsn: dsn,
	}, nil
}

// GetDSN returns the database connection string (DSN) from the pgConfig instance.
func (p *pgConfig) GetDSN() string {
	return p.dsn
}
