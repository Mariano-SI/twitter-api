package follow

import (
	"context"
	"database/sql"

	"github.com/Mariano-SI/twitter-api/internal/model"
)

type FollowRepository interface {
	Get(ctx context.Context, followerId, followedId string) (*model.FollowModel, error)
	Create(ctx context.Context, follow model.FollowModel) error
	Delete(ctx context.Context, followId string) error
}

type followRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) FollowRepository {
	return &followRepository{db: db}
}
