package commentimage

import (
	"context"
	"database/sql"

	"github.com/Mariano-SI/twitter-api/internal/model"
)

type CommentImageRepository interface {
	Create(ctx context.Context, comment_image model.CommentImageModel) error
}

type commentImageRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) CommentImageRepository {
	return &commentImageRepository{
		db: db,
	}
}
