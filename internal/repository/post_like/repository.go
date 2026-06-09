package postlike

import (
	"context"
	"database/sql"

	"github.com/Mariano-SI/twitter-api/internal/model"
)

type PostLikeRepository interface {
	Get(ctx context.Context, postId string, userId string) (*model.PostLikeModel, error)
	Create(ctx context.Context, post model.PostLikeModel) error
	Delete(ctx context.Context, postId string) error
}

type postLikerepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) PostLikeRepository {
	return &postLikerepository{
		db: db,
	}
}
