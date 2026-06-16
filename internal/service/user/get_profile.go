package user

import (
	"context"

	userDto "github.com/Mariano-SI/twitter-api/internal/dto/user"
	apperrors "github.com/Mariano-SI/twitter-api/internal/errors"
)

func (s *userService) GetProfile(ctx context.Context, userId string) (*userDto.GetUserProfileResponseDto, error) {
	profile, err := s.userRepository.GetProfileById(ctx, userId)
	if err != nil {
		return nil, err
	}
	if profile == nil {
		return nil, apperrors.ErrUserNotFound
	}

	return &userDto.GetUserProfileResponseDto{
		Id:                profile.ID,
		Username:          profile.Username,
		Description:       profile.Description,
		ProfilePictureUrl: profile.ProfilePictureUrl,
		FollowersCount:    profile.FollowersCount,
		FollowingCount:    profile.FollowingCount,
	}, nil
}
