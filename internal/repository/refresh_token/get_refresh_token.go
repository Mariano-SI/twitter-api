package refreshToken

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/Mariano-SI/twitter-api/internal/model"
)

func (r *refreshTokenRepository) GetRefreshToken(ctx context.Context, userId string, now time.Time) (*model.RefreshTokenModel, error) {
	query := `SELECT id, user_id, refresh_token, expires_at, created_at, updated_at FROM refresh_tokens WHERE user_id = ? AND expires_at > ? LIMIT 1`

	row := r.db.QueryRowContext(ctx, query, userId, now)

	var result model.RefreshTokenModel
	err := row.Scan(&result.ID, &result.UserID, &result.RefreshToken, &result.ExpiresAt, &result.CreatedAt, &result.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get refresh token: %w", err)
	}

	return &result, nil
}
