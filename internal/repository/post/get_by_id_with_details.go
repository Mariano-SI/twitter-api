package post

import (
	"context"
	"database/sql"

	"github.com/Mariano-SI/twitter-api/internal/model"
)

func (r *postRepository) GetByIdWithDetails(ctx context.Context, postId string) (*model.PostWithDetailsModel, error) {
	query := `
		SELECT p.id, p.content, p.user_id, p.created_at, p.updated_at,
		       u.username,
		       COUNT(pl.id) AS like_count
		FROM posts p
		JOIN users u ON u.id = p.user_id
		LEFT JOIN post_likes pl ON pl.post_id = p.id
		WHERE p.id = ? AND p.deleted_at IS NULL
		GROUP BY p.id, p.content, p.user_id, p.created_at, p.updated_at, u.username`

	row := r.db.QueryRowContext(ctx, query, postId)

	var result model.PostWithDetailsModel
	err := row.Scan(
		&result.ID, &result.Content, &result.UserID, &result.CreatedAt, &result.UpdatedAt,
		&result.Username, &result.LikeCount,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &result, nil
}
