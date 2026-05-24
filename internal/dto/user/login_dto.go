package user

import (
	"errors"

	"github.com/badoux/checkmail"
)

type LoginUserDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l *LoginUserDto) Validate() error {
	if l.Email == "" {
		return errors.New("email required")
	}

	if l.Password == "" {
		return errors.New("password is required")
	}

	if err := checkmail.ValidateFormat(l.Email); err != nil {
		return errors.New("invalid email")
	}

	return nil
}

type LoginUserResponseDto struct {
	Token        string `json:"Token"`
	RefreshToken string `json:"refresh_token"`
}
