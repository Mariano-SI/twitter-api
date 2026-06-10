package commentlike

import (
	"context"
	"database/sql"

	"github.com/Mariano-SI/twitter-api/internal/model"
)

type CommentLikeRepository interface {
	Get(ctx context.Context, commentId string, userId string) (*model.CommentLikeModel, error)
	Create(ctx context.Context, commentLike model.CommentLikeModel) error
	Delete(ctx context.Context, commentLikeId string) error
}

type commentLikeRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) CommentLikeRepository {
	return &commentLikeRepository{
		db: db,
	}
}
