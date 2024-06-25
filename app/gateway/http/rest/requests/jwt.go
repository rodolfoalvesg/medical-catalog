package requests

import (
	"net/http"
	"strings"
)

func GetSubjectID(r *http.Request) string {

	parts := strings.Split(r.Header.Get("Authorization"), "Bearer ")
	token := parts[1]

	return token
}
