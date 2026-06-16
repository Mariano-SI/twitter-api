package postImage

import (
	"context"
	"database/sql"

	"github.com/Mariano-SI/twitter-api/internal/model"
)

type PostImageRepository interface {
	Create(ctx context.Context, image *model.PostImageModel) error
	DeleteImagesByPostId(ctx context.Context, postId string) error
	GetByPostId(ctx context.Context, postId string) ([]*model.PostImageModel, error)
}

type postImageRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) PostImageRepository {
	return &postImageRepository{db: db}
}
