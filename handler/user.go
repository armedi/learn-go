package handler

import (
	"net/http"

	"github.com/armedi/learn-go/lib/api"
	"github.com/armedi/learn-go/user"
)

// UserHandler ...
type UserHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	userService user.Service
}

// NewUserHandler creates an object that represent UserHandler Interface
func NewUserHandler(us user.Service) UserHandler {
	return &userHandler{
		userService: us,
	}
}

func (uh *userHandler) Register(w http.ResponseWriter, r *http.Request) {
	data := &user.RegisterRequest{}
	if err := parseBody(r, data); err != nil {
		api.Render(w, err)
		return
	}
	if err := uh.userService.Register(data); err != nil {
		api.Render(w, err)
		return
	}
	api.Render(w, data)
}

func (uh *userHandler) Login(w http.ResponseWriter, r *http.Request) {

}
