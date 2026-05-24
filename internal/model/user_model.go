package model

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	ID        string
	Email     string
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUserModel(email, username, password string) (*UserModel, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	return &UserModel{
		ID:        uuid.NewString(),
		Email:     email,
		Username:  username,
		Password:  string(hashed),
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}
