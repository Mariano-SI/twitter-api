package user

import (
	"context"
	"database/sql"

	"github.com/Mariano-SI/twitter-api/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.UserModel) error
	GetUserByUsernameOrEmail(ctx context.Context, email, username string) (*model.UserModel, error)
	GetUserById(ctx context.Context, userId string) (*model.UserModel, error)
	GetProfileById(ctx context.Context, userId string) (*model.UserProfileModel, error)
	Update(ctx context.Context, user *model.UserModel) error
}

type userRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}
