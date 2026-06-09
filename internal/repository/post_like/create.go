package postlike

import (
	"context"
	"fmt"

	"github.com/Mariano-SI/twitter-api/internal/model"
)

func (r *postLikerepository) Create(ctx context.Context, post model.PostLikeModel) error {
	query := "INSERT INTO post_likes (id, post_id, user_id, created_at) VALUES (?,?,?,?)"

	_, err := r.db.ExecContext(ctx, query,
		post.ID,
		post.PostID,
		post.UserID,
		post.CreatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to create post like: %w", err)
	}

	return nil
}
