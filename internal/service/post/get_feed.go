package post

import (
	"context"

	postDto "github.com/Mariano-SI/twitter-api/internal/dto/post"
)

func (s *postService) GetFeed(ctx context.Context, input postDto.GetFeedDto, userId string) (*postDto.GetFeedResponseDto, error) {
	input.Normalize()

	offset := (input.Page - 1) * input.Limit

	posts, total, err := s.postRepository.GetFeed(ctx, userId, input.Limit, offset)
	if err != nil {
		return nil, err
	}

	postIds := make([]string, 0, len(posts))
	for _, p := range posts {
		postIds = append(postIds, p.ID)
	}

	allImages, err := s.postImageRepository.GetByPostIds(ctx, postIds)
	if err != nil {
		return nil, err
	}

	imagesByPost := make(map[string][]string, len(posts))
	for _, img := range allImages {
		imagesByPost[img.PostID] = append(imagesByPost[img.PostID], img.ImageURL)
	}

	data := make([]postDto.GetFeedPostResponseDto, 0, len(posts))
	for _, p := range posts {
		data = append(data, postDto.GetFeedPostResponseDto{
			Id:           p.ID,
			Username:     p.Username,
			Content:      p.Content,
			Images:       imagesByPost[p.ID],
			LikeCount:    p.LikeCount,
			CommentCount: p.CommentCount,
			CreatedAt:    p.CreatedAt,
			UpdatedAt:    p.UpdatedAt,
		})
	}

	return &postDto.GetFeedResponseDto{
		Data:  data,
		Total: total,
		Page:  input.Page,
		Limit: input.Limit,
	}, nil
}
