package handler

import (
	"fmt"
	"net/http"

	"github.com/armedi/learn-go/user"
	"github.com/go-chi/render"
)

// UserHandler ...
type UserHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	userSvc user.Service
}

// NewUserHandler creates an object that represent UserHandler Interface
func NewUserHandler(us user.Service) UserHandler {
	return &userHandler{us}
}

func (uh *userHandler) Register(w http.ResponseWriter, r *http.Request) {
	data := &registerRequest{}
	if err := render.Bind(r, data); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	if err := uh.userSvc.Register(&user.User{
		Email:    data.Email,
		Password: data.Password,
	}); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	fmt.Fprintln(w, data)
}

func (uh *userHandler) Login(w http.ResponseWriter, r *http.Request) {

}

type registerRequest struct {
	user.RegisterRequest
}

func (req *registerRequest) Bind(r *http.Request) error {
	return nil
}
