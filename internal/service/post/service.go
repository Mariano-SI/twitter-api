package post

import (
	"context"

	postDto "github.com/Mariano-SI/twitter-api/internal/dto/post"
	postRepository "github.com/Mariano-SI/twitter-api/internal/repository/post"
)

type PostService interface {
	Create(ctx context.Context, input postDto.CreatePostDto, userId string) (*postDto.CreatePostResponseDto, error)
}

type postService struct {
	postRepository postRepository.PostRepository
}

func NewService(postRepository postRepository.PostRepository) PostService {
	return &postService{postRepository: postRepository}
}
