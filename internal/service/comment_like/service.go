package commentlike

import (
	"context"

	commentLikeDto "github.com/Mariano-SI/twitter-api/internal/dto/comment_like"
	commentLikeRepository "github.com/Mariano-SI/twitter-api/internal/repository/comment_like"
	commentRepository "github.com/Mariano-SI/twitter-api/internal/repository/comment"
)

type CommentLikeService interface {
	LikeOrUnlikeComment(ctx context.Context, input commentLikeDto.LikeOrUnlikeCommentDto, userId string) (*commentLikeDto.LikeOrUnlikeCommentResponseDto, error)
}

type commentLikeService struct {
	commentLikeRepository commentLikeRepository.CommentLikeRepository
	commentRepository commentRepository.CommentRepository
}

func NewService(commentLikeRepository commentLikeRepository.CommentLikeRepository, commentRepository commentRepository.CommentRepository) CommentLikeService {
	return &commentLikeService{
		commentLikeRepository: commentLikeRepository,
		commentRepository: commentRepository,
	}
}
