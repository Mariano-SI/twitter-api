package model

import (
	"time"

	"github.com/google/uuid"
)

type FollowModel struct {
	ID         string
	FollowerID string
	FollowedID string
	CreatedAt  time.Time
}

func NewFollowModel(followerId, followedId string) *FollowModel {
	return &FollowModel{
		ID:         uuid.NewString(),
		FollowerID: followerId,
		FollowedID: followedId,
		CreatedAt:  time.Now(),
	}
}
