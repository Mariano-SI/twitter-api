package post

import (
	"context"
	"database/sql"

	"github.com/Mariano-SI/twitter-api/internal/model"
)

type PostRepository interface {
	Create(ctx context.Context, Post *model.PostModel) error
	GetById(ctx context.Context, postId string) (*model.PostModel, error)
	GetByIdWithDetails(ctx context.Context, postId string) (*model.PostWithDetailsModel, error)
	GetByUserIdPaginated(ctx context.Context, userId string, limit, offset int) ([]*model.PostSummaryModel, int, error)
	GetFeed(ctx context.Context, userId string, limit, offset int) ([]*model.PostFeedItemModel, int, error)
	Delete(ctx context.Context, postId string) error
}

type postRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) PostRepository {
	return &postRepository{db: db}
}
