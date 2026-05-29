package post

import (
	"context"

	postDto "github.com/Mariano-SI/twitter-api/internal/dto/post"
	"github.com/Mariano-SI/twitter-api/internal/infra/storage"
	postRepository "github.com/Mariano-SI/twitter-api/internal/repository/post"
)

type PostService interface {
	Create(ctx context.Context, input postDto.CreatePostDto, userId string) (*postDto.CreatePostResponseDto, error)
}

type postService struct {
	postRepository postRepository.PostRepository
	imageStorage   storage.Storage
}

func NewService(postRepository postRepository.PostRepository, imageStorage storage.Storage) PostService {
	return &postService{
		postRepository: postRepository,
		imageStorage:   imageStorage,
	}
}
