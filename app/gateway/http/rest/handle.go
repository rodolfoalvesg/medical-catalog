package rest

import (
	"back-platform/app/gateway/http/rest/responses"
	"encoding/json"
	"fmt"
	"net/http"
)

func Handle(handler func(r *http.Request) responses.Response) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := handler(r)

		if response.Error != nil {
			http.Error(w, response.Error.Error(), response.Status)
			return
		}

		if h := response.Header(); h != nil {
			copyHeaders(w, h)
		}

		if err := sendJSON(w, response.Payload, response.Status); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func sendJSON(w http.ResponseWriter, payload interface{}, status int) error {
	if payload == nil {
		w.WriteHeader(status)
		return nil
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		return fmt.Errorf("failed to encode response: %w", err)
	}

	return nil
}

func copyHeaders(w http.ResponseWriter, h http.Header) {
	wh := w.Header()

	for key, values := range h {
		for _, value := range values {
			wh.Add(key, value)
		}
	}
}
