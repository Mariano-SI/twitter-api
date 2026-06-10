package comment

import (
	"context"
	"database/sql"

	"github.com/Mariano-SI/twitter-api/internal/model"
)

type CommentRepository interface {
	Create(ctx context.Context, comment model.CommentModel) error
	GetById(ctx context.Context, commentId string)  (*model.CommentModel, error)
}

type commentRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) CommentRepository {
	return &commentRepository{db: db}
}
