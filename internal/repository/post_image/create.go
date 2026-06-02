package postImage

import (
	"context"
	"fmt"

	"github.com/Mariano-SI/twitter-api/internal/model"
	"github.com/Mariano-SI/twitter-api/pkg/internalSql"
)

func (r *postImageRepository) Create(ctx context.Context, image *model.PostImageModel) error {
	query := "INSERT INTO post_images (id, post_id, image_url, position, created_at) VALUES (?,?,?,?,?)"

	exec := internalSql.Executor(ctx, r.db)
	_, err := exec.ExecContext(ctx, query, image.ID, image.PostID, image.ImageURL, image.Position, image.CreatedAt)
	if err != nil {
		return fmt.Errorf("failed to create post image: %w", err)
	}

	return nil
}
