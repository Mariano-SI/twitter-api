package post

import (
	"context"

	"github.com/Mariano-SI/twitter-api/internal/model"
)

func (r *postRepository) GetByUserIdPaginated(ctx context.Context, userId string, limit, offset int) ([]*model.PostSummaryModel, int, error) {
	total := 0
	countRow := r.db.QueryRowContext(ctx,
		`SELECT COUNT(*) FROM posts WHERE user_id = ? AND deleted_at IS NULL`, userId)
	if err := countRow.Scan(&total); err != nil {
		return nil, 0, err
	}

	query := `
		SELECT p.id, p.content, p.user_id, p.created_at, p.updated_at,
		       COUNT(DISTINCT pl.id) AS like_count,
		       COUNT(DISTINCT c.id)  AS comment_count
		FROM posts p
		LEFT JOIN post_likes pl ON pl.post_id = p.id
		LEFT JOIN comments c    ON c.post_id  = p.id
		WHERE p.user_id = ? AND p.deleted_at IS NULL
		GROUP BY p.id, p.content, p.user_id, p.created_at, p.updated_at
		ORDER BY p.created_at DESC
		LIMIT ? OFFSET ?`

	rows, err := r.db.QueryContext(ctx, query, userId, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var results []*model.PostSummaryModel
	for rows.Next() {
		var p model.PostSummaryModel
		if err := rows.Scan(
			&p.ID, &p.Content, &p.UserID, &p.CreatedAt, &p.UpdatedAt,
			&p.LikeCount, &p.CommentCount,
		); err != nil {
			return nil, 0, err
		}
		results = append(results, &p)
	}

	return results, total, rows.Err()
}
