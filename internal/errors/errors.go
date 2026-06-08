package errors

import stderrors "errors"

var (
	ErrUserNotFound                = stderrors.New("user not found")
	ErrEmailOrUsernameAlreadyTaken = stderrors.New("email or username already taken")
	ErrPasswordMismatch            = stderrors.New("password and password confirmation do not match")
	ErrInvalidCredentials          = stderrors.New("invalid email or password")
	ErrInvalidRefreshToken         = stderrors.New("invalid or expired refresh token")
	ErrPostNotFound                = stderrors.New("post not found")
	ErrForbidden                   = stderrors.New("forbidden")
)
