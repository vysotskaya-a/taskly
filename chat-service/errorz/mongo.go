package errorz

import "fmt"

var (
	ErrNotFound      = fmt.Errorf("not found")
	ErrAlreadyExists = fmt.Errorf("already exists")
)
