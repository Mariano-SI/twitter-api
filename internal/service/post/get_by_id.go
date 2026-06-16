package post

import (
	"context"

	postDto "github.com/Mariano-SI/twitter-api/internal/dto/post"
	apperrors "github.com/Mariano-SI/twitter-api/internal/errors"
)

func (s *postService) GetById(ctx context.Context, input postDto.GetPostByIdDto) (*postDto.GetPostByIdResponseDto, error) {
	post, err := s.postRepository.GetByIdWithDetails(ctx, input.Id)
	if err != nil {
		return nil, err
	}
	if post == nil {
		return nil, apperrors.ErrPostNotFound
	}

	postImages, err := s.postImageRepository.GetByPostId(ctx, input.Id)
	if err != nil {
		return nil, err
	}

	comments, err := s.commentRepository.GetByPostIdWithStats(ctx, input.Id)
	if err != nil {
		return nil, err
	}

	commentImages, err := s.commentImageRepository.GetAllByPostId(ctx, input.Id)
	if err != nil {
		return nil, err
	}

	imagesByComment := make(map[string][]string, len(commentImages))
	for _, img := range commentImages {
		imagesByComment[img.CommentID] = append(imagesByComment[img.CommentID], img.ImageUrl)
	}

	imageUrls := make([]string, 0, len(postImages))
	for _, img := range postImages {
		imageUrls = append(imageUrls, img.ImageURL)
	}

	commentDtos := make([]postDto.GetPostByIdCommentResponseDto, 0, len(comments))
	for _, c := range comments {
		commentDtos = append(commentDtos, postDto.GetPostByIdCommentResponseDto{
			Id:        c.ID,
			Username:  c.Username,
			Content:   c.Content,
			LikeCount: c.LikeCount,
			Images:    imagesByComment[c.ID],
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
		})
	}

	return &postDto.GetPostByIdResponseDto{
		Id:        post.ID,
		Username:  post.Username,
		Content:   post.Content,
		Images:    imageUrls,
		LikeCount: post.LikeCount,
		Comments:  commentDtos,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}, nil
}
