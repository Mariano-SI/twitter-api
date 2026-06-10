package commentlike

import (
	"context"
	"fmt"
)

func (r *commentLikeRepository) Delete(ctx context.Context, commentLikeId string) error {

	query := "DELETE FROM comment_likes WHERE id = ?"

	_, err := r.db.ExecContext(ctx, query,
		commentLikeId,
	)
	if err != nil {
		return fmt.Errorf("failed to delete comment like: %w", err)
	}
	return nil
}
