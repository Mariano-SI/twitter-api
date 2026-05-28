package user

import (
	"errors"
	"strings"
)

type RefreshTokenDto struct {
	RefreshToken string `json:"refresh_token" `
}

func (r *RefreshTokenDto) Validate() error {
	if strings.TrimSpace(r.RefreshToken) == "" {
		return errors.New("refresh_token is required")
	}

	return nil
}

type RefreshTokenResponseDto struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
