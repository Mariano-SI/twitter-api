package model

import (
	"time"

	"github.com/google/uuid"
)

type CommentModel struct {
	ID        string
	PostId    string
	UserId    string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewCommentModel(postId, userId, content string) *CommentModel {
	now := time.Now()
	return &CommentModel{
		ID:        uuid.NewString(),
		PostId:    postId,
		UserId:    userId,
		Content:   content,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
