package r2

import (
	"context"
	"fmt"
	"io"

	"github.com/Mariano-SI/twitter-api/internal/infra/storage"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type r2Storage struct {
	client    *s3.Client
	bucket    string
	publicURL string
}

func NewStorage(client *s3.Client, bucket, publicURL string) storage.Storage {
	return &r2Storage{
		client:    client,
		bucket:    bucket,
		publicURL: publicURL,
	}
}

func (r *r2Storage) Upload(ctx context.Context, key string, content io.Reader, contentType string) (string, error) {
	_, err := r.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(r.bucket),
		Key:         aws.String(key),
		Body:        content,
		ContentType: aws.String(contentType),
	})
	if err != nil {
		return "", fmt.Errorf("upload to r2: %w", err)
	}

	return fmt.Sprintf("%s/%s", r.publicURL, key), nil
}

func (r *r2Storage) Delete(ctx context.Context, key string) error {
	_, err := r.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(r.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return fmt.Errorf("delete from r2: %w", err)
	}
	return nil
}
