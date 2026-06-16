package user

import (
	"context"

	"github.com/Mariano-SI/twitter-api/internal/config"
	userDto "github.com/Mariano-SI/twitter-api/internal/dto/user"
	"github.com/Mariano-SI/twitter-api/internal/infra/storage"
	refreshToken "github.com/Mariano-SI/twitter-api/internal/repository/refresh_token"
	"github.com/Mariano-SI/twitter-api/internal/repository/user"
)

type UserService interface {
	Register(ctx context.Context, input userDto.RegisterUserDto) (*userDto.RegisterUserResponseDto, error)
	Login(ctx context.Context, input userDto.LoginUserDto) (*userDto.LoginUserResponseDto, error)
	RefreshToken(ctx context.Context, input userDto.RefreshTokenDto, userId string) (*userDto.RefreshTokenResponseDto, error)
	UpdateProfile(ctx context.Context, input userDto.UpdateProfileDto, userId string) (*userDto.UpdateProfileResponseDto, error)
	GetProfile(ctx context.Context, userId string) (*userDto.GetUserProfileResponseDto, error)
}

type userService struct {
	cfg                    *config.Config
	userRepository         user.UserRepository
	refreshTokenRepository refreshToken.RefreshTokenRepository
	imageStorage           storage.Storage
}

func NewService(cfg *config.Config, userRepository user.UserRepository, refreshTokenRepository refreshToken.RefreshTokenRepository, imageStorage storage.Storage) UserService {
	return &userService{
		cfg:                    cfg,
		userRepository:         userRepository,
		refreshTokenRepository: refreshTokenRepository,
		imageStorage:           imageStorage,
	}
}
