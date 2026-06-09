package model

import (
	"time"

	"github.com/google/uuid"
)

type PostLikeModel struct {
	ID        string
	PostID    string
	UserID    string
	CreatedAt time.Time
}

func NewPostLikeModel(postID, userId string) *PostLikeModel {
	return &PostLikeModel{
		ID:        uuid.NewString(),
		PostID:    postID,
		UserID:    userId,
		CreatedAt: time.Now(),
	}
}
