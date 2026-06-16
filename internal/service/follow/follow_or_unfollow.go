package follow

import (
	"context"

	followDto "github.com/Mariano-SI/twitter-api/internal/dto/follow"
	apperrors "github.com/Mariano-SI/twitter-api/internal/errors"
	"github.com/Mariano-SI/twitter-api/internal/model"
)

func (s *followService) FollowOrUnfollow(ctx context.Context, input followDto.FollowOrUnfollowDto, followerId string) (*followDto.FollowOrUnfollowResponseDto, error) {
	if followerId == input.FollowedId {
		return nil, apperrors.ErrCannotFollowSelf
	}

	followed, err := s.userRepository.GetUserById(ctx, input.FollowedId)
	if err != nil {
		return nil, err
	}
	if followed == nil {
		return nil, apperrors.ErrUserNotFound
	}

	existing, err := s.followRepository.Get(ctx, followerId, input.FollowedId)
	if err != nil {
		return nil, err
	}

	if existing == nil {
		newFollow := model.NewFollowModel(followerId, input.FollowedId)
		if err := s.followRepository.Create(ctx, *newFollow); err != nil {
			return nil, err
		}
		return &followDto.FollowOrUnfollowResponseDto{Following: true}, nil
	}

	if err := s.followRepository.Delete(ctx, existing.ID); err != nil {
		return nil, err
	}

	return &followDto.FollowOrUnfollowResponseDto{Following: false}, nil
}
