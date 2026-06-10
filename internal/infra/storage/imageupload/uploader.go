package imageupload

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"

	"github.com/google/uuid"

	"github.com/Mariano-SI/twitter-api/internal/infra/storage"
)

// Upload envia todos os arquivos para o bucket usando prefix como pasta
// (ex: "posts/<id>" ou "comments/<id>") e usa build para montar o model
// concreto de cada imagem. Retorna os models e as keys enviadas (úteis para
// cleanup se algo falhar depois). Em caso de erro durante o loop, faz o
// cleanup do que já subiu antes de retornar.
func Upload[T any](
	ctx context.Context,
	s storage.Storage,
	prefix string,
	files []*multipart.FileHeader,
	build func(url string, position int) T,
) ([]T, []string, error) {
	images := make([]T, 0, len(files))
	keys := make([]string, 0, len(files))

	for position, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			Cleanup(s, keys)
			return nil, nil, fmt.Errorf("open image %q: %w", fileHeader.Filename, err)
		}

		contentType := fileHeader.Header.Get("Content-Type")
		key := fmt.Sprintf("%s/%s%s", prefix, uuid.NewString(), extForContentType(contentType))

		url, err := s.Upload(ctx, key, file, contentType)
		file.Close()
		if err != nil {
			Cleanup(s, keys)
			return nil, nil, err
		}

		keys = append(keys, key)
		images = append(images, build(url, position))
	}

	return images, keys, nil
}

// Cleanup remove do bucket as keys informadas. Usado quando algo falha depois
// do upload (ex: a transação no banco).
func Cleanup(s storage.Storage, keys []string) {
	for _, key := range keys {
		if err := s.Delete(context.Background(), key); err != nil {
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
