package postImage

import (
	"context"
	"fmt"
	"strings"

	"github.com/Mariano-SI/twitter-api/internal/model"
)

func (r *postImageRepository) GetByPostIds(ctx context.Context, postIds []string) ([]*model.PostImageModel, error) {
	if len(postIds) == 0 {
		return nil, nil
	}

	placeholders := strings.Repeat("?,", len(postIds))
	placeholders = placeholders[:len(placeholders)-1]

	query := fmt.Sprintf(
		`SELECT id, post_id, image_url, position, created_at FROM post_images WHERE post_id IN (%s) ORDER BY post_id, position`,
		placeholders,
	)

	args := make([]any, len(postIds))
	for i, id := range postIds {
		args[i] = id
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
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
