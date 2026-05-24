package user

import (
	"context"
	"time"

	userDto "github.com/Mariano-SI/twitter-api/internal/dto/user"
	apperrors "github.com/Mariano-SI/twitter-api/internal/errors"
	"github.com/Mariano-SI/twitter-api/internal/model"
	"github.com/Mariano-SI/twitter-api/pkg/jwt"
)

const refreshTokenTTL = 7 * 24 * time.Hour

func (us *userService) Login(ctx context.Context, input userDto.LoginUserDto) (*userDto.LoginUserResponseDto, error) {
	//checkar se o user existe
	user, err := us.userRepository.GetUserByUsernameOrEmail(ctx, input.Email, "")

	if err != nil {
		return nil, err
	}

	if user == nil || !user.ComparePassword(input.Password) {
		return nil, apperrors.ErrInvalidCredentials
	}

	token, err := jwt.CreateToken(user.ID, user.Username, us.cfg.JwtSecret)

	if err != nil {
		return nil, err
	}

	now := time.Now()

	refreshToken, err := us.refreshTokenRepository.GetRefreshToken(ctx, user.ID, now)

	if err != nil {
		return nil, err
	}

	if refreshToken != nil {
		return &userDto.LoginUserResponseDto{
			Token:        token,
			RefreshToken: refreshToken.RefreshToken,
		}, nil
	}

	newRefreshToken := model.NewRefreshTokenModel(user.ID, refreshTokenTTL)
	if err := us.refreshTokenRepository.Create(ctx, newRefreshToken); err != nil {
		return nil, err
	}

	return &userDto.LoginUserResponseDto{
		Token:        token,
		RefreshToken: newRefreshToken.RefreshToken,
	}, nil
}
