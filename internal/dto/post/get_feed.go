package post

import "time"

type GetFeedDto struct {
	Page  int
	Limit int
}

func (d *GetFeedDto) Normalize() {
	if d.Page < 1 {
		d.Page = 1
	}
	if d.Limit < 1 || d.Limit > 50 {
		d.Limit = 20
	}
}

type GetFeedPostResponseDto struct {
	Id           string    `json:"id"`
	Username     string    `json:"username"`
	Content      string    `json:"content"`
	Images       []string  `json:"images"`
	LikeCount    int       `json:"like_count"`
	CommentCount int       `json:"comment_count"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type GetFeedResponseDto struct {
	Data  []GetFeedPostResponseDto `json:"data"`
	Total int                      `json:"total"`
	Page  int                      `json:"page"`
	Limit int                      `json:"limit"`
}
