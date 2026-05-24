package httperror

import (
	"errors"
	"log"
	"net/http"

	apperrors "github.com/Mariano-SI/twitter-api/internal/errors"
)

// FromError maps a domain error to an HTTP status and a client-safe message.
// Unknown errors return 500 with a generic message; the original is logged.
func FromError(err error) (int, string) {
	switch {
	case errors.Is(err, apperrors.ErrUserNotFound):
		return http.StatusNotFound, err.Error()
	case errors.Is(err, apperrors.ErrEmailOrUsernameAlreadyTaken):
		return http.StatusConflict, err.Error()
	case errors.Is(err, apperrors.ErrPasswordMismatch):
		return http.StatusBadRequest, err.Error()
	default:
		log.Printf("unhandled error: %v", err)
		return http.StatusInternalServerError, "internal server error"
	}
}
