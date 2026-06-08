package post

import (
	"context"
	"fmt"

	"github.com/Mariano-SI/twitter-api/pkg/internalSql"
)

func (r *postRepository) Delete(ctx context.Context, postId string) error {
	query := "DELETE FROM posts WHERE id = ?"

	exec := internalSql.Executor(ctx, r.db)
	_, err := exec.ExecContext(ctx, query, postId)
	if err != nil {
		return fmt.Errorf("failed to delete post: %w", err)
	}

	return nil
}
