package comment

import (
	"context"
	"fmt"

	"github.com/Mariano-SI/twitter-api/internal/model"
	"github.com/Mariano-SI/twitter-api/pkg/internalSql"
)

func (r *commentRepository) Create(ctx context.Context, comment model.CommentModel) error {
	query := "INSERT INTO comments (id, post_id, user_id, content, created_at, updated_at) VALUES (?,?,?,?,?,?)"

	exec := internalSql.Executor(ctx, r.db)
	_, err := exec.ExecContext(ctx, query,
		comment.ID,
		comment.PostId,
		comment.UserId,
		comment.Content,
		comment.CreatedAt,
		comment.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to create comment: %w", err)
	}

	return nil
}
