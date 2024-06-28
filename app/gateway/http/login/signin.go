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

// SignIn signs in a user.
// @Summary Sign in a user
// @Description Sign in a user
// @Tags Login
// @Accept json
// @Produce json
// @Param input body SignInInput true "Sign in input"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 401 {object} string
// @Failure 500 {object} string
// @Router /service/v1/medical-catalog/login [post]

// Header all {string} RequestID x-request-id

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
