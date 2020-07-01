package api

import (
	"fmt"
	"net/http"
)

func createErrorResponse(message string, code string, httpStatus int) ErrorResponse {
	result := ErrorResponse{
		Success:    false,
		Message:    message,
		ErrorCode:  code,
		HTTPStatus: httpStatus,
	}
	if result.Message == "" {
		result.Message = fmt.Sprint(http.StatusText(httpStatus))
	}
	if result.ErrorCode == "" {
		result.ErrorCode = fmt.Sprintf("ERR%d", httpStatus)
	}
	return result
}

// ErrBadRequest ...
func ErrBadRequest(message string, code string) ErrorResponse {
	return createErrorResponse(message, code, http.StatusBadRequest)
}

// ErrConflict ...
func ErrConflict(message string, code string) ErrorResponse {
	return createErrorResponse(message, code, http.StatusConflict)

}
