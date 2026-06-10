package comment

import (
	"context"

	commentDto "github.com/Mariano-SI/twitter-api/internal/dto/comment"
	"github.com/Mariano-SI/twitter-api/internal/infra/storage"
	commentRepository "github.com/Mariano-SI/twitter-api/internal/repository/comment"
	commentImageRepository "github.com/Mariano-SI/twitter-api/internal/repository/comment_image"
	postRepository "github.com/Mariano-SI/twitter-api/internal/repository/post"
	"github.com/Mariano-SI/twitter-api/pkg/internalSql"
)

type CommentService interface {
	CreateComment(ctx context.Context, input commentDto.CreateCommentDto, userId string) (*commentDto.CreateCommentResponseDto, error)
}

type commentService struct {
	transactor             internalSql.Transactor
	commentRepository      commentRepository.CommentRepository
	commentImageRepository commentImageRepository.CommentImageRepository
	postRepository         postRepository.PostRepository
	imageStorage           storage.Storage
}

func NewService(
	transactor internalSql.Transactor,
	commentRepository commentRepository.CommentRepository,
	commentImageRepository commentImageRepository.CommentImageRepository, postRepository postRepository.PostRepository,
	imageStorage storage.Storage,
) CommentService {
	return &commentService{
		transactor:             transactor,
		commentImageRepository: commentImageRepository,
		postRepository:         postRepository,
		commentRepository:      commentRepository,
		imageStorage:           imageStorage,
	}
}
