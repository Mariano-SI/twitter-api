package model

import "time"

type CommentWithStatsModel struct {
	ID        string
	PostID    string
	UserID    string
	Username  string
	Content   string
	LikeCount int
	CreatedAt time.Time
	UpdatedAt time.Time
}
