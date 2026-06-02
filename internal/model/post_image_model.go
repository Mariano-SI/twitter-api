package model

import (
	"time"

	"github.com/google/uuid"
)

type PostImageModel struct {
	ID        string
	PostID    string
	ImageURL  string
	Position  int
	CreatedAt time.Time
}

func NewPostImageModel(postID, imageURL string, position int) *PostImageModel {
	return &PostImageModel{
		ID:        uuid.NewString(),
		PostID:    postID,
		ImageURL:  imageURL,
		Position:  position,
		CreatedAt: time.Now(),
	}
}
