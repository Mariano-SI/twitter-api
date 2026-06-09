package postlike

import (
	"context"
	"fmt"
)

func (r *postLikerepository) Delete(ctx context.Context, postId string) error {
	query := "DELETE FROM post_likes WHERE id = ?"

	_, err := r.db.ExecContext(ctx, query,
		postId,
	)
	if err != nil {
		return fmt.Errorf("failed to delete post like: %w", err)
	}

	return nil
}
