package post

import (
	"context"

	postDto "github.com/Mariano-SI/twitter-api/internal/dto/post"
	apperrors "github.com/Mariano-SI/twitter-api/internal/errors"
)

func (s *postService) Delete(ctx context.Context, input postDto.DeletePostDto, userId string) (*postDto.DeletePostResponseDto, error) {
	post, err := s.postRepository.GetById(ctx, input.Id)

	if err != nil {
		return nil, err
	}

	if post == nil {
		return nil, apperrors.ErrPostNotFound
	}

	if post.UserId != userId {
		return nil, apperrors.ErrForbidden
	}

	err = s.transactor.WithinTransaction(ctx, func(ctx context.Context) error {
		if err := s.postImageRepository.DeleteImagesByPostId(ctx, post.ID); err != nil {
			return err
		}

		if err := s.postRepository.Delete(ctx, post.ID); err != nil {
			return err
		}

		return nil
	})

	if err != nil{
		return nil, err
	}

	return &postDto.DeletePostResponseDto{
		Message: "successfully deleted post",
	}, nil
}
