package post

import (
	"context"

	postDto "github.com/Mariano-SI/twitter-api/internal/dto/post"
	"github.com/Mariano-SI/twitter-api/internal/infra/storage"
	postRepository "github.com/Mariano-SI/twitter-api/internal/repository/post"
	postImageRepository "github.com/Mariano-SI/twitter-api/internal/repository/post_image"
	"github.com/Mariano-SI/twitter-api/pkg/internalSql"
)

type PostService interface {
	Create(ctx context.Context, input postDto.CreatePostDto, userId string) (*postDto.CreatePostResponseDto, error)
}

type postService struct {
	transactor          internalSql.Transactor
	postRepository      postRepository.PostRepository
	postImageRepository postImageRepository.PostImageRepository
	imageStorage        storage.Storage
}

func NewService(
	transactor internalSql.Transactor,
	postRepository postRepository.PostRepository,
	postImageRepository postImageRepository.PostImageRepository,
	imageStorage storage.Storage,
) PostService {
	return &postService{
		transactor:          transactor,
		postRepository:      postRepository,
		postImageRepository: postImageRepository,
		imageStorage:        imageStorage,
	}
}
