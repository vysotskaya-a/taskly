package errorz

import "errors"

var (
	ErrUserNotFound         = errors.New("user not found")
	ErrPasswordDoesNotMatch = errors.New("password does not match")
)

// вот тут применение errors.New разумно, только почему пакет называется errorZ?))) 