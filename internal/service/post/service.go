package post

import (
	"context"

	postDto "github.com/Mariano-SI/twitter-api/internal/dto/post"
	"github.com/Mariano-SI/twitter-api/internal/infra/storage"
	commentRepository "github.com/Mariano-SI/twitter-api/internal/repository/comment"
	commentImageRepository "github.com/Mariano-SI/twitter-api/internal/repository/comment_image"
	postRepository "github.com/Mariano-SI/twitter-api/internal/repository/post"
	postImageRepository "github.com/Mariano-SI/twitter-api/internal/repository/post_image"
	"github.com/Mariano-SI/twitter-api/pkg/internalSql"
)

type PostService interface {
	Create(ctx context.Context, input postDto.CreatePostDto, userId string) (*postDto.CreatePostResponseDto, error)
	Delete(ctx context.Context, input postDto.DeletePostDto, userId string) (*postDto.DeletePostResponseDto, error)
	GetById(ctx context.Context, input postDto.GetPostByIdDto) (*postDto.GetPostByIdResponseDto, error)
}

type postService struct {
	transactor             internalSql.Transactor
	postRepository         postRepository.PostRepository
	postImageRepository    postImageRepository.PostImageRepository
	commentRepository      commentRepository.CommentRepository
	commentImageRepository commentImageRepository.CommentImageRepository
	imageStorage           storage.Storage
}

func NewService(
	transactor internalSql.Transactor,
	postRepository postRepository.PostRepository,
	postImageRepository postImageRepository.PostImageRepository,
	commentRepository commentRepository.CommentRepository,
	commentImageRepository commentImageRepository.CommentImageRepository,
	imageStorage storage.Storage,
) PostService {
	return &postService{
		transactor:             transactor,
		postRepository:         postRepository,
		postImageRepository:    postImageRepository,
		commentRepository:      commentRepository,
		commentImageRepository: commentImageRepository,
		imageStorage:           imageStorage,
	}
}
