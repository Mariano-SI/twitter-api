package user

import (
	"context"

	"github.com/Mariano-SI/twitter-api/internal/dto"
	apperrors "github.com/Mariano-SI/twitter-api/internal/errors"
	"github.com/Mariano-SI/twitter-api/internal/model"
)

func (us *userService) Register(ctx context.Context, input dto.RegisterUserDto) (*dto.RegisterUserResponseDto, error) {

	if input.Password != input.PasswordConfirm {
		return nil, apperrors.ErrPasswordMismatch
	}

	existing, err := us.userRepository.GetUserByUsernameOrEmail(ctx, input.Email, input.Username)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, apperrors.ErrEmailOrUsernameAlreadyTaken
	}

	user, err := model.NewUserModel(input.Email, input.Username, input.Password)
	if err != nil {
		return nil, err
	}

	err = us.userRepository.Create(ctx, user)

	if err != nil {
		return nil, err
	}

	return &dto.RegisterUserResponseDto{
		ID: user.ID,
	}, nil
}
