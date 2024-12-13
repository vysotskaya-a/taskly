package errorz

import "errors"

var (
	ErrUserIDNotSet = errors.New("user_id not set")

	ErrProjectNotFound        = errors.New("project not found")
	ErrProjectAccessForbidden = errors.New("project access forbidden")

	ErrTaskNotFound = errors.New("task not found")
)
