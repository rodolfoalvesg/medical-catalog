package requests

import (
	"back-platform/app/gateway/http/rest/responses"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var ErrInvalidBodyJSON = errors.New("invalid body json")

type validator interface {
	Validate() error
}

func DecodeBodyJSON(r *http.Request, dest interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(dest); err != nil {
		return responses.ValidationError{
			Param: "body",
			Err:   ErrInvalidBodyJSON,
		}
	}

	v, ok := dest.(validator)
	if !ok {
		return nil
	}

	if err := v.Validate(); err != nil {
		return fmt.Errorf("validating JSON body: %w", err)
	}

	return nil
}
