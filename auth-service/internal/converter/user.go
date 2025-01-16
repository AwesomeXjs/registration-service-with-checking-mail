package converter

import (
	"time"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/model"
)

// FromUserInfoToDbModel converts UserInfo and hashed password to InfoToDb model.
func FromUserInfoToDbModel(info *model.UserInfo, hashedPassword string) *model.InfoToDb {
	return &model.InfoToDb{
		Email:        info.Email,
		HashPassword: hashedPassword,
		Role:         info.Role,
	}
}

// FromUpdatePassInfoToDbPassInfo converts UpdatePassInfo and hashed password to UpdatePassDb model.
func FromUpdatePassInfoToDbPassInfo(info *model.UpdatePassInfo, hashedPassword string) *model.UpdatePassDb {
	return &model.UpdatePassDb{
		Email:        info.Email,
		HashPassword: hashedPassword,
		UpdatedAt:    time.Now(),
	}
}
