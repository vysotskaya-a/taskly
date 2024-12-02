package errorz

import "errors"

var (
	ErrUserNotFound         = errors.New("user not found")
	ErrPasswordDoesNotMatch = errors.New("password does not match")
)
