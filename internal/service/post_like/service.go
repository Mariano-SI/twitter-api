package postlike

import (
	"context"

	postLikeDto "github.com/Mariano-SI/twitter-api/internal/dto/post_like"
	postRepository "github.com/Mariano-SI/twitter-api/internal/repository/post"
	postLikeRepository "github.com/Mariano-SI/twitter-api/internal/repository/post_like"
)

type PostLikeService interface {
	LikeOrUnlike(ctx context.Context, input postLikeDto.LikeOrUnlikePostDto, userId string) (*postLikeDto.LikeOrUnlikePostResponseDto, error)
}

type postLikeService struct {
	postLikeRepository postLikeRepository.PostLikeRepository
	postRepository     postRepository.PostRepository
}

func NewService(postLikeRepository postLikeRepository.PostLikeRepository, postRepository postRepository.PostRepository) PostLikeService {
	return &postLikeService{
		postLikeRepository: postLikeRepository,
		postRepository:     postRepository,
	}
}
