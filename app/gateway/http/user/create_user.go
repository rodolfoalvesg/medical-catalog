package user

import (
	"back-platform/app/domain/entities/user"
	"back-platform/app/domain/usecases"
	"back-platform/app/gateway/http/rest/responses"
	"back-platform/app/gateway/http/user/schema"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// CreateUser creates a new user.
// @Summary Create a new user
// @Description Create a new user to access the system.
// @Description Otherwise it will return HTTP status code 409.
// @Tags User
// @Security BearerToken
// @Param Body body schema.CreateUserInput true "User input"
// @Success 200 {object} string "User created successfully"
// @Failure 400 {object} string "Invalid request body"
// @Failure 409 {object} string "User already exists"
// @Failure 500 {object} string "Internal server error"
// @Header all {string} RequestID "x-request-id"
// @Router /service/v1/medical-catalog/users [post]
func (h Handler) CreateUser(r *http.Request) responses.Response {
	const operation = "UserHandler.CreateUser"

	var req schema.CreateUserInput
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return responses.BadRequest(fmt.Errorf("%s: %w", operation, err), "invalid request body")
	}

	if req.UserName == "" || req.Email == "" || req.Password == "" {
		return responses.BadRequest(nil, "missing required fields")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return responses.InternalServerError(fmt.Errorf("%s: %w", operation, err))
	}

	input := usecases.CreateUserInput{
		Username: req.UserName,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	if err := h.usecase.CreateUser(r.Context(), input); err != nil {
		if errors.Is(err, user.ErrUserAlreadyExists) {
			return responses.Conflict(fmt.Errorf("%s: %w", operation, err), "user already exists")
		}

		return responses.InternalServerError(fmt.Errorf("%s: %w", operation, err))
	}

	return responses.Created()
}
