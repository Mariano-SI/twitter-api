package postImage

import (
	"context"

	"github.com/Mariano-SI/twitter-api/internal/model"
)

func (r *postImageRepository) GetByPostId(ctx context.Context, postId string) ([]*model.PostImageModel, error) {
	query := `SELECT id, post_id, image_url, position, created_at FROM post_images WHERE post_id = ? ORDER BY position`

	rows, err := r.db.QueryContext(ctx, query, postId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []*model.PostImageModel
	for rows.Next() {
		var img model.PostImageModel
		if err := rows.Scan(&img.ID, &img.PostID, &img.ImageURL, &img.Position, &img.CreatedAt); err != nil {
			return nil, err
		}
		results = append(results, &img)
	}

	return results, rows.Err()
}
