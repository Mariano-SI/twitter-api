package comment

import (
	"context"
	"database/sql"

	"github.com/Mariano-SI/twitter-api/internal/model"
)

func (r *commentRepository) GetById(ctx context.Context, commentId string) (*model.CommentModel, error) {
	query := "SELECT id, post_id, user_id, content, created_at, updated_at FROM comments WHERE id = ?"

	row := r.db.QueryRowContext(ctx, query, commentId)

	var result model.CommentModel

	err := row.Scan(&result.ID, &result.PostId, &result.UserId, &result.Content, &result.CreatedAt, &result.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &result, nil
}
