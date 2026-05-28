package user

import (
	"context"

	"github.com/Mariano-SI/twitter-api/internal/config"
	userDto "github.com/Mariano-SI/twitter-api/internal/dto/user"
	refreshToken "github.com/Mariano-SI/twitter-api/internal/repository/refresh_token"
	"github.com/Mariano-SI/twitter-api/internal/repository/user"
)

type UserService interface {
	Register(ctx context.Context, input userDto.RegisterUserDto) (*userDto.RegisterUserResponseDto, error)
	Login(ctx context.Context, input userDto.LoginUserDto) (*userDto.LoginUserResponseDto, error)
	RefreshToken(ctx context.Context, input userDto.RefreshTokenDto, userId string)(*userDto.RefreshTokenResponseDto, error)
}

type userService struct {
	cfg            *config.Config
	userRepository user.UserRepository
	refreshTokenRepository refreshToken.RefreshTokenRepository
}

func NewService(cfg *config.Config, userRepository user.UserRepository, refreshTokenRepository refreshToken.RefreshTokenRepository) UserService {
	return &userService{
		cfg:            cfg,
		userRepository: userRepository,
		refreshTokenRepository: refreshTokenRepository,
	}
}
