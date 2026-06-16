package user

import (
	"context"
	"mime/multipart"
	"time"

	userDto "github.com/Mariano-SI/twitter-api/internal/dto/user"
	apperrors "github.com/Mariano-SI/twitter-api/internal/errors"
	"github.com/Mariano-SI/twitter-api/internal/infra/storage/imageupload"
)

func (s *userService) UpdateProfile(ctx context.Context, input userDto.UpdateProfileDto, userId string) (*userDto.UpdateProfileResponseDto, error) {
	if input.RemoveProfilePicture != nil && *input.RemoveProfilePicture && input.ProfilePicture != nil {
		return nil, apperrors.ErrConflictingProfilePictureAction
	}

	user, err := s.userRepository.GetUserById(ctx, userId)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, apperrors.ErrUserNotFound
	}

	if input.Description != nil {
		user.Description = input.Description
	}

	if input.RemoveProfilePicture != nil && *input.RemoveProfilePicture {
		if user.ProfilePictureKey != nil {
			if err := s.imageStorage.Delete(ctx, *user.ProfilePictureKey); err != nil {
				return nil, err
			}
		}
		user.ProfilePictureUrl = nil
		user.ProfilePictureKey = nil
	}

	if input.ProfilePicture != nil {
		if user.ProfilePictureKey != nil {
			if err := s.imageStorage.Delete(ctx, *user.ProfilePictureKey); err != nil {
				return nil, err
			}
		}

		urls, keys, err := imageupload.Upload(ctx, s.imageStorage, "users/"+userId+"/profile",
			[]*multipart.FileHeader{input.ProfilePicture},
			func(url string, _ int) string { return url },
		)
		if err != nil {
			return nil, err
		}

		user.ProfilePictureUrl = &urls[0]
		user.ProfilePictureKey = &keys[0]
	}

	user.UpdatedAt = time.Now()

	if err := s.userRepository.Update(ctx, user); err != nil {
		return nil, err
	}

	return &userDto.UpdateProfileResponseDto{
		Id:                user.ID,
		Username:          user.Username,
		Description:       user.Description,
		ProfilePictureUrl: user.ProfilePictureUrl,
	}, nil
}
