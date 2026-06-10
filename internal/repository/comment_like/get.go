package commentlike

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Mariano-SI/twitter-api/internal/model"
)

func (r *commentLikeRepository) Get(ctx context.Context, commentId string, userId string) (*model.CommentLikeModel, error) {
	query := "SELECT id, comment_id, user_id, created_at FROM comment_likes WHERE comment_id = ? AND user_id = ?"

	row := r.db.QueryRowContext(ctx, query, commentId, userId)

	var result model.CommentLikeModel

	err := row.Scan(&result.ID, &result.CommentID, &result.UserID, &result.CreatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &result, nil

}
