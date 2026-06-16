package model

import "time"

type PostFeedItemModel struct {
	ID           string
	Content      string
	UserID       string
	Username     string
	LikeCount    int
	CommentCount int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
