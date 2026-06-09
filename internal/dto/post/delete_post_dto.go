package post

import (
	"errors"
	"strings"

	"github.com/google/uuid"
)

type DeletePostDto struct {
	Id string
}

func (d *DeletePostDto) Validate() error {
	if strings.TrimSpace(d.Id) == "" {
		return errors.New("id is required")
	}

	return uuid.Validate(d.Id)
}

type DeletePostResponseDto struct {
	Message string `json:"message"`
}
