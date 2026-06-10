package commentimage

import (
	"context"
	"fmt"

	"github.com/Mariano-SI/twitter-api/internal/model"
	"github.com/Mariano-SI/twitter-api/pkg/internalSql"
)

func (r *commentImageRepository) Create(ctx context.Context, comment_image model.CommentImageModel) error {
	query := "INSERT INTO comment_images (id, comment_id, image_url, position, created_at) VALUES (?,?,?,?,?)"

	exec := internalSql.Executor(ctx, r.db)
	_, err := exec.ExecContext(ctx, query,
		comment_image.ID,
		comment_image.CommentID,
		comment_image.ImageUrl,
		comment_image.Position,
		comment_image.CreatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to create comment image: %w", err)
	}

	return nil
}
