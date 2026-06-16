package post

import (
	"context"

	postDto "github.com/Mariano-SI/twitter-api/internal/dto/post"
)

func (s *postService) GetMyPosts(ctx context.Context, input postDto.GetMyPostsDto, userId string) (*postDto.GetMyPostsResponseDto, error) {
	input.Normalize()

	offset := (input.Page - 1) * input.Limit

	posts, total, err := s.postRepository.GetByUserIdPaginated(ctx, userId, input.Limit, offset)
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

	data := make([]postDto.GetMyPostsPostResponseDto, 0, len(posts))
	for _, p := range posts {
		data = append(data, postDto.GetMyPostsPostResponseDto{
			Id:           p.ID,
			Content:      p.Content,
			Images:       imagesByPost[p.ID],
			LikeCount:    p.LikeCount,
			CommentCount: p.CommentCount,
			CreatedAt:    p.CreatedAt,
			UpdatedAt:    p.UpdatedAt,
		})
	}

	return &postDto.GetMyPostsResponseDto{
		Data:  data,
		Total: total,
		Page:  input.Page,
		Limit: input.Limit,
	}, nil
}
