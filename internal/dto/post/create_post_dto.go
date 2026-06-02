package post

import (
	"errors"
	"fmt"
	"mime/multipart"
	"strings"
)

const (
	MaxImagesPerPost = 4
	MaxImageSize = 2 * 1024 * 1024 
)

var allowedImageMimeTypes = map[string]struct{}{
	"image/jpeg": {},
	"image/png":  {},
	"image/webp": {},
}

type CreatePostDto struct {
	Content string
	Images  []*multipart.FileHeader
}

func (c *CreatePostDto) Validate() error {
	if strings.TrimSpace(c.Content) == "" {
		return errors.New("content is required")
	}
	if len(c.Content) > 280 {
		return errors.New("content exceeds 280 characters")
	}

	if len(c.Images) > MaxImagesPerPost {
		return fmt.Errorf("a post can have at most %d images", MaxImagesPerPost)
	}

	for _, img := range c.Images {
		if img.Size > MaxImageSize {
			return fmt.Errorf("image %q exceeds the %dMB size limit", img.Filename, MaxImageSize>>20)
		}

		contentType := img.Header.Get("Content-Type")
		if _, ok := allowedImageMimeTypes[contentType]; !ok {
			return fmt.Errorf("image %q has unsupported type %q", img.Filename, contentType)
		}
	}

	return nil
}

type CreatePostResponseDto struct {
	Id string `json:"id"`
}
