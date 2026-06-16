package follow

import (
	"errors"
	"strings"

	"github.com/google/uuid"
)

type FollowOrUnfollowDto struct {
	FollowedId string
}

func (d *FollowOrUnfollowDto) Validate() error {
	if strings.TrimSpace(d.FollowedId) == "" {
		return errors.New("id is required")
	}
	return uuid.Validate(d.FollowedId)
}

type FollowOrUnfollowResponseDto struct {
	Following bool `json:"following"`
}
