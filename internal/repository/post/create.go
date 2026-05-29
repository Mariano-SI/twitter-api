package post

import (
	"context"
	"fmt"

	"github.com/Mariano-SI/twitter-api/internal/model"
)

func (r *postRepository) Create(ctx context.Context, post *model.PostModel) error {
	query := "INSERT INTO posts (id, user_id, content, created_at, updated_at) VALUES (?,?,?,?,?)"

	_, err := r.db.ExecContext(ctx, query, post.ID, post.UserId, post.Content, post.CreatedAt, post.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}
