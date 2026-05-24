package refreshToken

import (
	"context"
	"fmt"

	"github.com/Mariano-SI/twitter-api/internal/model"
)

func (r *refreshTokenRepository) Create(ctx context.Context, refreshToken *model.RefreshTokenModel) error {
	query := `INSERT INTO refresh_tokens (id, user_id, refresh_token, expires_at, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query,
		refreshToken.ID,
		refreshToken.UserID,
		refreshToken.RefreshToken,
		refreshToken.ExpiresAt,
		refreshToken.CreatedAt,
		refreshToken.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create refresh token: %w", err)
	}

	return nil
}
