package dto

import (
	"errors"
	"strings"

	"github.com/badoux/checkmail"
)

type RegisterUserDto struct {
	Email           string `json:"email"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}

func (d RegisterUserDto) Validate() error {
	if strings.TrimSpace(d.Email) == "" {
		return errors.New("email is required")
	}
	if strings.TrimSpace(d.Username) == "" {
		return errors.New("username is required")
	}
	if d.Password == "" {
		return errors.New("password is required")
	}
	if err := checkmail.ValidateFormat(d.Email); err != nil {
		return errors.New("invalid email")
	}
	return nil
}

type RegisterUserResponseDto struct {
	ID string `json:"id"`
}
