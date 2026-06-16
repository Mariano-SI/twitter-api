package model

import "time"

type PostWithDetailsModel struct {
	ID        string
	Content   string
	UserID    string
	Username  string
	LikeCount int
	CreatedAt time.Time
	UpdatedAt time.Time
}
