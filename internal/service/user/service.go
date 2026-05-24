package user

import (
	"context"

	"github.com/Mariano-SI/twitter-api/internal/config"
	userDto "github.com/Mariano-SI/twitter-api/internal/dto/user"
	"github.com/Mariano-SI/twitter-api/internal/repository/user"
)

type UserService interface {
	Register(ctx context.Context, input userDto.RegisterUserDto) (*userDto.RegisterUserResponseDto, error)
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
