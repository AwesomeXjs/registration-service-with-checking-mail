package configs

import (
	"fmt"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/utils/consts"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/utils/logger"
	"go.uber.org/zap"
	"os"
)

type PGConfig interface {
	GetDSN() string
}

type pgConfig struct {
	dsn string
}

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

func (p *pgConfig) GetDSN() string {
	return p.dsn
}
