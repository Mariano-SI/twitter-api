package refreshToken

import (
	"context"
	"fmt"
)

func (r *refreshTokenRepository) Delete(ctx context.Context, tokenId string) error {
	query := "DELETE FROM refresh_tokens WHERE id = ?"
	_, err := r.db.ExecContext(ctx, query,
		tokenId,
	)
	if err != nil {
		return fmt.Errorf("failed to delete refresh token: %w", err)
	}

	return nil
}
