package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// SuccessResponse is the standard response format for 2xx responses
type SuccessResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// ErrorResponse is the standard response format for non-2xx responses
type ErrorResponse struct {
	Success    bool       `json:"success"`
	Message    string     `json:"message,omitempty"`
	ErrorCode  string     `json:"error_code"`
	HTTPStatus int        `json:"-"`
	GRPCCode   codes.Code `json:"-"`
}

func (resp ErrorResponse) Error() string {
	return resp.Message
}

// GRPCStatus returns grpc status
func (resp ErrorResponse) GRPCStatus() *status.Status {
	return status.New(resp.GRPCCode, resp.Message)
}

// Render ...
func Render(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	err, ok := v.(error)
	if ok && err != nil {
		// v is for error case
		response, ok := err.(ErrorResponse)
		if !ok {
			response = ErrorResponse{
				Success:    false,
				Message:    http.StatusText(http.StatusInternalServerError),
				ErrorCode:  fmt.Sprintf("ERR%d", http.StatusInternalServerError),
				HTTPStatus: http.StatusInternalServerError,
			}
		}
		w.WriteHeader(response.HTTPStatus)
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			panic(err)
		}
	} else {
		// v is for success case
		err = json.NewEncoder(w).Encode(SuccessResponse{
			Success: true,
			Data:    v,
		})
		if err != nil {
			panic(err)
		}
	}
}
