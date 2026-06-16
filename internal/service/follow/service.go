package follow

import (
	"context"

	followDto "github.com/Mariano-SI/twitter-api/internal/dto/follow"
	followRepository "github.com/Mariano-SI/twitter-api/internal/repository/follow"
	userRepository "github.com/Mariano-SI/twitter-api/internal/repository/user"
)

type FollowService interface {
	FollowOrUnfollow(ctx context.Context, input followDto.FollowOrUnfollowDto, followerId string) (*followDto.FollowOrUnfollowResponseDto, error)
}

type followService struct {
	followRepository followRepository.FollowRepository
	userRepository   userRepository.UserRepository
}

func NewService(followRepository followRepository.FollowRepository, userRepository userRepository.UserRepository) FollowService {
	return &followService{
		followRepository: followRepository,
		userRepository:   userRepository,
	}
}
