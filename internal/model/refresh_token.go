package model

import (
	"time"

	"github.com/Mariano-SI/twitter-api/pkg/refreshtoken"
	"github.com/google/uuid"
)

type RefreshTokenModel struct {
	ID           string
	UserID       string
	RefreshToken string
	ExpiresAt    time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewRefreshTokenModel(userID string, ttl time.Duration) *RefreshTokenModel {
	now := time.Now()
	return &RefreshTokenModel{
		ID:           uuid.NewString(),
		UserID:       userID,
		RefreshToken: refreshtoken.Generate(),
		ExpiresAt:    now.Add(ttl),
		CreatedAt:    now,
		UpdatedAt:    now,
	}
}
