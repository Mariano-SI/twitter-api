package user

import (
	"context"

	"github.com/Mariano-SI/twitter-api/internal/config"
	"github.com/Mariano-SI/twitter-api/internal/dto"
	"github.com/Mariano-SI/twitter-api/internal/repository/user"
)

type UserService interface {
	Register(ctx context.Context, input dto.RegisterUserDto) (*dto.RegisterUserResponseDto, error)
}

type userService struct {
	cfg            *config.Config
	userRepository user.UserRepository
}

func NewService(cfg *config.Config, userRepository user.UserRepository) UserService {
	return &userService{
		cfg:            cfg,
		userRepository: userRepository,
	}
}
