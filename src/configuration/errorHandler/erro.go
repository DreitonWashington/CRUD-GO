package errorhandler

import "net/http"

type Erro struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *Erro) Error() string {
	return r.Message
}

func newErrorHandler(message, err string, code int, causes []Causes) *Erro {
	return &Erro{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func newBadRequestError(message string) *Erro {
	return &Erro{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

func newBadRequestValidationError(message string, causes []Causes) *Erro {
	return &Erro{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func newInternalServerError(message string) *Erro {
	return &Erro{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
	}
}

func newNotFoundError(message string) *Erro {
	return &Erro{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound,
	}
}

func newForbiddenError(message string) *Erro {
	return &Erro{
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
	}
}
