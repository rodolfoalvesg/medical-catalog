package user

import "github.com/google/uuid"

type User struct {
	ID       int
	PublicID uuid.UUID
	Username string
	Password string
	Email    string
}
