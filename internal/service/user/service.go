package user

import (
	"github.com/Mariano-SI/twitter-api/internal/config"
	"github.com/Mariano-SI/twitter-api/internal/repository/user"
)

type UserService interface {
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
