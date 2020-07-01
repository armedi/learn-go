package api

import (
	"fmt"
	"net/http"
)

// ErrBadRequest ...
func ErrBadRequest(message string, code string) ErrorResponse {
	result := ErrorResponse{
		Success:    false,
		Message:    message,
		ErrorCode:  code,
		HTTPStatus: http.StatusBadRequest,
	}
	if result.ErrorCode == "" {
		result.ErrorCode = fmt.Sprintf("ERR%d", http.StatusBadRequest)
	}
	return result
}
