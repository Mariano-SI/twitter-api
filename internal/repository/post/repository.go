package post

import (
	"context"
	"database/sql"

	"github.com/Mariano-SI/twitter-api/internal/model"
)

type PostRepository interface {
	Create(ctx context.Context, Post *model.PostModel) error
}

type postRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) PostRepository {
	return &postRepository{db: db}
}
