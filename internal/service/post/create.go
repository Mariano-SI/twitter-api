package post

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"

	"github.com/google/uuid"

	postDto "github.com/Mariano-SI/twitter-api/internal/dto/post"
	"github.com/Mariano-SI/twitter-api/internal/model"
)

func (s *postService) Create(ctx context.Context, input postDto.CreatePostDto, userId string) (*postDto.CreatePostResponseDto, error) {
	post := model.NewPostModel(userId, input.Content)

	images, uploadedKeys, err := s.uploadImages(ctx, post.ID, input.Images)
	if err != nil {
		s.cleanupKeys(uploadedKeys)
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
		s.cleanupKeys(uploadedKeys)
		return nil, err
	}

	return &postDto.CreatePostResponseDto{
		Id: post.ID,
	}, nil
}

func (s *postService) uploadImages(ctx context.Context, postID string, files []*multipart.FileHeader) ([]*model.PostImageModel, []string, error) {
	images := make([]*model.PostImageModel, 0, len(files))
	keys := make([]string, 0, len(files))

	for position, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			return images, keys, fmt.Errorf("open image %q: %w", fileHeader.Filename, err)
		}

		contentType := fileHeader.Header.Get("Content-Type")
		key := fmt.Sprintf("posts/%s/%s%s", postID, uuid.NewString(), extForContentType(contentType))

		url, err := s.imageStorage.Upload(ctx, key, file, contentType)
		file.Close()
		if err != nil {
			return images, keys, err
		}

		keys = append(keys, key)
		images = append(images, model.NewPostImageModel(postID, url, position))
	}

	return images, keys, nil
}


func (s *postService) cleanupKeys(keys []string) {
	for _, key := range keys {
		if err := s.imageStorage.Delete(context.Background(), key); err != nil {
			log.Printf("failed to cleanup r2 key %q: %v", key, err)
		}
	}
}

func extForContentType(contentType string) string {
	switch contentType {
	case "image/jpeg":
		return ".jpg"
	case "image/png":
		return ".png"
	case "image/webp":
		return ".webp"
	default:
		return ""
	}
}
