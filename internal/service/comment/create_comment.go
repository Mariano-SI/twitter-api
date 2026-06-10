package comment

import (
	"context"

	commentDto "github.com/Mariano-SI/twitter-api/internal/dto/comment"
	apperrors "github.com/Mariano-SI/twitter-api/internal/errors"
	"github.com/Mariano-SI/twitter-api/internal/infra/storage/imageupload"
	"github.com/Mariano-SI/twitter-api/internal/model"
)

func (s *commentService) CreateComment(ctx context.Context, input commentDto.CreateCommentDto, userId string) (*commentDto.CreateCommentResponseDto, error) {
	post, err := s.postRepository.GetById(ctx, input.PostId)
	if err != nil {
		return nil, err
	}
	if post == nil {
		return nil, apperrors.ErrPostNotFound
	}

	comment := model.NewCommentModel(input.PostId, userId, input.Content)

	images, uploadedKeys, err := imageupload.Upload(ctx, s.imageStorage, "comments/"+comment.ID, input.Images,
		func(url string, position int) *model.CommentImageModel {
			return model.NewCommentImageModel(comment.ID, url, uint8(position))
		})
	if err != nil {
		return nil, err
	}

	err = s.transactor.WithinTransaction(ctx, func(ctx context.Context) error {
		if err := s.commentRepository.Create(ctx, *comment); err != nil {
			return err
		}

		for _, image := range images {
			if err := s.commentImageRepository.Create(ctx, *image); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		imageupload.Cleanup(s.imageStorage, uploadedKeys)
		return nil, err
	}

	return &commentDto.CreateCommentResponseDto{
		Message: "successfully created comment",
	}, nil
}
