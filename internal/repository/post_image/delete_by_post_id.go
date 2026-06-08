package postImage

import (
	"context"
	"fmt"

	"github.com/Mariano-SI/twitter-api/pkg/internalSql"
)

func (r *postImageRepository) DeleteImagesByPostId(ctx context.Context, postId string) error {
	query := "DELETE FROM post_images WHERE post_id = ?"

	exec := internalSql.Executor(ctx, r.db)
	_, err := exec.ExecContext(ctx, query, postId)
	if err != nil {
		return fmt.Errorf("failed to delete post images: %w", err)
	}

	return nil
}
