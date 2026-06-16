package commentimage

import (
	"context"

	"github.com/Mariano-SI/twitter-api/internal/model"
)

func (r *commentImageRepository) GetAllByPostId(ctx context.Context, postId string) ([]*model.CommentImageModel, error) {
	query := `
		SELECT ci.id, ci.comment_id, ci.image_url, ci.position, ci.created_at
		FROM comment_images ci
		JOIN comments c ON c.id = ci.comment_id
		WHERE c.post_id = ?
		ORDER BY ci.comment_id, ci.position`

	rows, err := r.db.QueryContext(ctx, query, postId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []*model.CommentImageModel
	for rows.Next() {
		var img model.CommentImageModel
		if err := rows.Scan(&img.ID, &img.CommentID, &img.ImageUrl, &img.Position, &img.CreatedAt); err != nil {
			return nil, err
		}
		results = append(results, &img)
	}

	return results, rows.Err()
}
