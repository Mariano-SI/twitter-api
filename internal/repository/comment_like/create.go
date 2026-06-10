package commentlike

import (
	"context"
	"fmt"

	"github.com/Mariano-SI/twitter-api/internal/model"
)

func (r *commentLikeRepository) Create(ctx context.Context, commentLike model.CommentLikeModel) error {
	query := "INSERT INTO comment_likes (id, comment_id, user_id, created_at) VALUES (?, ?, ?, ?);"

	_, err := r.db.ExecContext(ctx, query,
		commentLike.ID,
		commentLike.CommentID,
		commentLike.UserID,
		commentLike.CreatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to create comment like: %w", err)
	}

	return nil
}
