package api

import (
	"fmt"
	"net/http"

	"google.golang.org/grpc/codes"
)

func createErrorResponse(message string, code string, httpStatus int, grpcCode codes.Code) ErrorResponse {
	result := ErrorResponse{
		Success:    false,
		Message:    message,
		ErrorCode:  code,
		HTTPStatus: httpStatus,
		GRPCCode:   grpcCode,
	}
	if result.Message == "" {
		result.Message = fmt.Sprint(http.StatusText(httpStatus))
	}
	if result.ErrorCode == "" {
		result.ErrorCode = fmt.Sprintf("ERR%d", httpStatus)
	}
	return result
}

// NewErrBadRequest 400
func NewErrBadRequest(message string, code string) ErrorResponse {
	return createErrorResponse(message, code, http.StatusBadRequest, codes.InvalidArgument)
}

// NewErrUnauthorized 401
func NewErrUnauthorized(message string, code string) ErrorResponse {
	return createErrorResponse(message, code, http.StatusUnauthorized, codes.Unauthenticated)
}

// NewErrNotFound 404
func NewErrNotFound(message string, code string) ErrorResponse {
	return createErrorResponse(message, code, http.StatusNotFound, codes.NotFound)
}

// NewErrConflict 409
func NewErrConflict(message string, code string) ErrorResponse {
	return createErrorResponse(message, code, http.StatusConflict, codes.FailedPrecondition)
}
