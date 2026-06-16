package user

import (
	"errors"
	"mime/multipart"
)

const (
	maxDescriptionLength = 160
	MaxImageSize         = 2 * 1024 * 1024
)

var allowedImageMimeTypes = map[string]struct{}{
	"image/jpeg": {},
	"image/png":  {},
	"image/webp": {},
}

type UpdateProfileDto struct {
	Description          *string
	ProfilePicture       *multipart.FileHeader
	RemoveProfilePicture *bool
}

func (d *UpdateProfileDto) Validate() error {
	if d.Description != nil && len(*d.Description) > maxDescriptionLength {
		return errors.New("description exceeds 160 characters")
	}

	if d.ProfilePicture != nil {
		if d.ProfilePicture.Size > MaxImageSize {
			return errors.New("profile picture exceeds the 2MB size limit")
		}

		contentType := d.ProfilePicture.Header.Get("Content-Type")
		if _, ok := allowedImageMimeTypes[contentType]; !ok {
			return errors.New("profile picture has unsupported type")
		}
	}

	return nil
}

type UpdateProfileResponseDto struct {
	Id                string  `json:"id"`
	Username          string  `json:"username"`
	Description       *string `json:"description"`
	ProfilePictureUrl *string `json:"profile_picture_url"`
}
