package login

import (
	"back-platform/app/domain/usecases"
	"back-platform/app/domain/vos"
	"back-platform/app/gateway/http/login/schema"
	"back-platform/app/gateway/http/rest/responses"
	"encoding/json"
	"fmt"
	"net/http"
)

// SignIn creates a new login.
// @Summary Create a new login
// @Description Create a new login to access the system.
// @Tags Login
// @Param Body body schema.SignInInput true "Login input"
// @Success 200 {object} string "Login successfully"
// @Failure 400 {object} string "Invalid request body"
// @Failure 401 {object} string "Invalid email or password"
// @Failure 500 {object} string "Internal server error"
// @Header all {string} RequestID "x-request-id"
// @Router /admin/v1/medical-catalog/login/signin [post]
func (h Handler) SignIn(r *http.Request) responses.Response {
	const operation = "LoginHandler.SignIn"

	var req schema.SignInInput
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return responses.BadRequest(fmt.Errorf("%s: %w", operation, err), "invalid request body")
	}

	if req.Email == "" && req.Password == "" {
		return responses.BadRequest(fmt.Errorf("%s: %w", operation, vos.ErrUserAndPassEmpty), "invalid request body")
	}

	input := usecases.SignInInput{
		Email:    req.Email,
		Password: req.Password,
	}

	token, err := h.usecase.SignIn(r.Context(), input)
	if err != nil {
		return vos.Errors(err)
	}

	header := http.Header{
		"Authorization": []string{"Bearer " + token},
	}

	return responses.Success(header)
}
