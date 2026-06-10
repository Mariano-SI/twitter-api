package commentlike

import (
	"context"

	commentLikeDto "github.com/Mariano-SI/twitter-api/internal/dto/comment_like"
	apperrors "github.com/Mariano-SI/twitter-api/internal/errors"
	"github.com/Mariano-SI/twitter-api/internal/model"
)

func (s *commentLikeService) LikeOrUnlikeComment(ctx context.Context, input commentLikeDto.LikeOrUnlikeCommentDto, userId string) (*commentLikeDto.LikeOrUnlikeCommentResponseDto, error) {
	comment, err := s.commentRepository.GetById(ctx, input.CommentId)

	if err != nil {
		return nil, err
	}

	if comment == nil {
		return nil, apperrors.ErrCommentNotFound
	}

	commentLike, err := s.commentLikeRepository.Get(ctx, input.CommentId, userId)

	if err != nil {
		return nil, err
	}

	if commentLike == nil {
		newCommentLike := model.NewCommentLikeModel(input.CommentId, userId)
		if err := s.commentLikeRepository.Create(ctx, *newCommentLike); err != nil {
			return nil, err
		}
	} else {
		if err := s.commentLikeRepository.Delete(ctx, commentLike.ID); err != nil {
			return nil, err
		}
	}

	return &commentLikeDto.LikeOrUnlikeCommentResponseDto{
		Message: "successful",
	}, nil
}
