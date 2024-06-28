package specialty

import "errors"

var (
	ErrSpecialtyAlreadyExists = errors.New("specialty already exists")
	ErrSpecialtyIsEmpty       = errors.New("specialty name is empty")
)
