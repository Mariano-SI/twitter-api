package post

import (
	"context"
	"database/sql"

	"github.com/Mariano-SI/twitter-api/internal/model"
)

func (r *postRepository) GetById(ctx context.Context, postId string) (*model.PostModel, error) {
	query := `SELECT id, content, user_id, created_at, updated_at, deleted_at FROM posts WHERE id = ?`

	row := r.db.QueryRowContext(ctx, query, postId)

	var result model.PostModel

	err := row.Scan(&result.ID, &result.Content, &result.UserId, &result.UpdatedAt, &result.CreatedAt, &result.DeletedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &result, nil
}
