package post

import (
	"context"

	postDto "github.com/Mariano-SI/twitter-api/internal/dto/post"
	"github.com/Mariano-SI/twitter-api/internal/infra/storage/imageupload"
	"github.com/Mariano-SI/twitter-api/internal/model"
)

func (s *postService) Create(ctx context.Context, input postDto.CreatePostDto, userId string) (*postDto.CreatePostResponseDto, error) {
	post := model.NewPostModel(userId, input.Content)

	images, uploadedKeys, err := imageupload.Upload(ctx, s.imageStorage, "posts/"+post.ID, input.Images,
		func(url string, position int) *model.PostImageModel {
			return model.NewPostImageModel(post.ID, url, position)
		})
	if err != nil {
		return nil, err
	}

	err = s.transactor.WithinTransaction(ctx, func(ctx context.Context) error {
		if err := s.postRepository.Create(ctx, post); err != nil {
			return err
		}

		for _, image := range images {
			if err := s.postImageRepository.Create(ctx, image); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		imageupload.Cleanup(s.imageStorage, uploadedKeys)
		return nil, err
	}

	return &postDto.CreatePostResponseDto{
		Id: post.ID,
	}, nil
}
