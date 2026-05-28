package user

import (
	"context"
	"time"

	userDto "github.com/Mariano-SI/twitter-api/internal/dto/user"
	apperrors "github.com/Mariano-SI/twitter-api/internal/errors"
	"github.com/Mariano-SI/twitter-api/internal/model"
	"github.com/Mariano-SI/twitter-api/pkg/jwt"
)

func (us *userService) RefreshToken(ctx context.Context, input userDto.RefreshTokenDto, userId string) (*userDto.RefreshTokenResponseDto, error) {
	user, err := us.userRepository.GetUserById(ctx, userId)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, apperrors.ErrUserNotFound
	}

	refreshToken, err := us.refreshTokenRepository.GetRefreshToken(ctx, user.ID, time.Now())

	if err != nil {
		return nil, err
	}

	if refreshToken == nil {
		return nil, apperrors.ErrInvalidRefreshToken
	}

	if input.RefreshToken != refreshToken.RefreshToken {
		return nil, apperrors.ErrInvalidRefreshToken
	}

	newJwtToken, err := jwt.CreateToken(user.ID, user.Username, us.cfg.JwtSecret)
	
	if err != nil {
		return nil, err
	}

	newRefreshToken := model.NewRefreshTokenModel(user.ID, us.cfg.RefreshTokenTTL)

	err = us.refreshTokenRepository.Create(ctx, newRefreshToken)

	if err != nil {
		return nil, err
	}

	err = us.refreshTokenRepository.Delete(ctx, refreshToken.ID)

	if err != nil {
		return nil, err
	}

	return &userDto.RefreshTokenResponseDto{
		Token:        newJwtToken,
		RefreshToken: newRefreshToken.RefreshToken,
	}, nil
}
