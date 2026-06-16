package comment

import (
	"context"

	"github.com/Mariano-SI/twitter-api/internal/model"
)

func (r *commentRepository) GetByPostIdWithStats(ctx context.Context, postId string) ([]*model.CommentWithStatsModel, error) {
	query := `
		SELECT c.id, c.post_id, c.user_id, c.content, c.created_at, c.updated_at,
		       u.username,
		       COUNT(cl.id) AS like_count
		FROM comments c
		JOIN users u ON u.id = c.user_id
		LEFT JOIN comment_likes cl ON cl.comment_id = c.id
		WHERE c.post_id = ?
		GROUP BY c.id, c.post_id, c.user_id, c.content, c.created_at, c.updated_at, u.username
		ORDER BY c.created_at ASC`

	rows, err := r.db.QueryContext(ctx, query, postId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []*model.CommentWithStatsModel
	for rows.Next() {
		var c model.CommentWithStatsModel
		if err := rows.Scan(
			&c.ID, &c.PostID, &c.UserID, &c.Content, &c.CreatedAt, &c.UpdatedAt,
			&c.Username, &c.LikeCount,
		); err != nil {
			return nil, err
		}
		results = append(results, &c)
	}

	return results, rows.Err()
}
