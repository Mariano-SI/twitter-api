package model

import "time"

type PostSummaryModel struct {
	ID           string
	Content      string
	UserID       string
	LikeCount    int
	CommentCount int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
