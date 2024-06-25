package responses

import "net/http"

func BadRequest(err error, message string) Response {
	return Response{
		Status: http.StatusBadRequest,
		Error:  err,
		Payload: map[string]string{
			"message": message,
		},
	}
}

func Created() Response {
	return Response{
		Status: http.StatusCreated,
	}
}

func InternalServerError(err error) Response {
	return Response{
		Status: http.StatusInternalServerError,
		Error:  err,
		Payload: map[string]string{
			"message": "internal server error",
		},
	}
}

func Ok(payload interface{}) Response {
	return Response{
		Status:  http.StatusOK,
		Payload: payload,
	}
}

func Conflict(err error, message string) Response {
	return Response{
		Status: http.StatusConflict,
		Error:  err,
		Payload: map[string]string{
			"message": message,
		},
	}
}
