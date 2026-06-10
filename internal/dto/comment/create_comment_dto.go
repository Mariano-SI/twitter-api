package comment

import (
	"errors"
	"fmt"
	"mime/multipart"
)

const (
	MaxImagesPerPost = 4
	MaxImageSize     = 2 * 1024 * 1024
)

var allowedImageMimeTypes = map[string]struct{}{
	"image/jpeg": {},
	"image/png":  {},
	"image/webp": {},
}

type CreateCommentDto struct {
	Content string
	PostId  string
	Images  []*multipart.FileHeader
}

func (c *CreateCommentDto) Validate() error {
	if len(c.Content) > 280 {
		return errors.New("content exceeds 280 characters")
	}

	if len(c.Images) > MaxImagesPerPost {
		return fmt.Errorf("a comment can have at most %d images", MaxImagesPerPost)
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

type CreateCommentResponseDto struct {
	Message string `json:"message"`
}
