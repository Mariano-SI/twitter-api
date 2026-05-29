package storage

import (
	"context"
	"io"
)

type Storage interface {
	Upload(ctx context.Context, key string, content io.Reader, contentType string) (url string, err error)
	Delete(ctx context.Context, key string) error
}
