package model

import (
	"time"

	"github.com/google/uuid"
)

type PostModel struct {
	ID        string
	Content   string
	UserId    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func NewPostModel(userId, content string) (*PostModel) {
	now := time.Now()
	return &PostModel{
		ID:        uuid.NewString(),
		Content:   content,
		UserId:    userId,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
