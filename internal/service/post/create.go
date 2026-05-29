package post

import (
	"context"

	postDto "github.com/Mariano-SI/twitter-api/internal/dto/post"
	"github.com/Mariano-SI/twitter-api/internal/model"
)

func (s *postService) Create(ctx context.Context, input postDto.CreatePostDto, userId string) (*postDto.CreatePostResponseDto, error) {
	post := model.NewPostModel(userId, input.Content)

	err := s.postRepository.Create(ctx, post)

	if err != nil {
		return nil, err
	}

	return &postDto.CreatePostResponseDto{
		Id: post.ID,
	}, nil
}
