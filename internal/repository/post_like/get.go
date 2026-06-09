package postlike

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Mariano-SI/twitter-api/internal/model"
)

func (r *postLikerepository) Get(ctx context.Context, postId string, userId string) (*model.PostLikeModel, error) {
	query := "SELECT id, post_id, user_id, created_at FROM post_likes WHERE post_id = ? AND user_id = ?"
	row := r.db.QueryRowContext(ctx, query, postId, userId)

	var result model.PostLikeModel

	err := row.Scan(&result.ID, &result.PostID, &result.UserID, &result.CreatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &result, nil
}
