package postlike

import (
	"errors"
	"strings"

	"github.com/google/uuid"
)

type LikeOrUnlikePostDto struct {
	PostId string `json:"post_id"`
}

func (l *LikeOrUnlikePostDto) Validate() error {
	if strings.TrimSpace(l.PostId) == "" {
		return errors.New("post_id is required")
	}

	return uuid.Validate(l.PostId)
}

type LikeOrUnlikePostResponseDto struct {
	Message string `json:"message"`
}

/*
  1. Busca o post por id — retorna 404 se não existir
  2. Busca o like existente por (postId, userId)
  3. Se existir → deleta o like (unlike)
  4. Se não existir → cria o like
  5. Retorna uma mensagem indicando a ação executada ("liked" / "unliked")
*/