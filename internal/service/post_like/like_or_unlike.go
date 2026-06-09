package postlike

import (
	"context"

	postLikeDto "github.com/Mariano-SI/twitter-api/internal/dto/post_like"
	apperrors "github.com/Mariano-SI/twitter-api/internal/errors"
	"github.com/Mariano-SI/twitter-api/internal/model"
)

func (s *postLikeService) LikeOrUnlike(ctx context.Context, input postLikeDto.LikeOrUnlikePostDto, userId string) (*postLikeDto.LikeOrUnlikePostResponseDto, error) {
	post, err := s.postRepository.GetById(ctx, input.PostId)
	if err != nil {
		return nil, err
	}
	if post == nil {
		return nil, apperrors.ErrPostNotFound
	}

	like, err := s.postLikeRepository.Get(ctx, input.PostId, userId)
	if err != nil {
		return nil, err
	}

	if like == nil {
		newLike := model.NewPostLikeModel(input.PostId, userId)
		if err := s.postLikeRepository.Create(ctx, *newLike); err != nil {
			return nil, err
		}
	} else {
		if err := s.postLikeRepository.Delete(ctx, like.ID); err != nil {
			return nil, err
		}
	}

	return &postLikeDto.LikeOrUnlikePostResponseDto{
		Message: "successful",
	}, nil
}
