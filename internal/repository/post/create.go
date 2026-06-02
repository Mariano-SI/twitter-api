package post

import (
	"context"
	"fmt"

	"github.com/Mariano-SI/twitter-api/internal/model"
	"github.com/Mariano-SI/twitter-api/pkg/internalSql"
)

func (r *postRepository) Create(ctx context.Context, post *model.PostModel) error {
	query := "INSERT INTO posts (id, user_id, content, created_at, updated_at) VALUES (?,?,?,?,?)"

	exec := internalSql.Executor(ctx, r.db)
	_, err := exec.ExecContext(ctx, query, post.ID, post.UserId, post.Content, post.CreatedAt, post.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to create post: %w", err)
	}

	return nil
}
