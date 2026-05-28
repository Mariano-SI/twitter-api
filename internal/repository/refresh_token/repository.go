package refreshToken

import (
	"context"
	"database/sql"
	"time"

	"github.com/Mariano-SI/twitter-api/internal/model"
)

type RefreshTokenRepository interface {
	GetRefreshToken(ctx context.Context, userId string, now time.Time) (*model.RefreshTokenModel, error)
	Create(ctx context.Context, refreshToken *model.RefreshTokenModel) error
	Delete(ctx context.Context, tokenId string) error
}

type refreshTokenRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) RefreshTokenRepository {
	return &refreshTokenRepository{db: db}
}
