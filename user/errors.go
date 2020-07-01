package user

import (
	"github.com/go-playground/validator/v10"

	"github.com/armedi/learn-go/lib/api"
)

var (
	errEmailRequired    error = api.ErrBadRequest("Kolom email tidak boleh dikosongkan", "")
	errEmailInvalid     error = api.ErrBadRequest("Kolom email harus diisi dengan email yang valid", "")
	errPasswordRequired error = api.ErrBadRequest("Kolom password tidak boleh dikosongkan", "")
)

func parseValidationError(e error) error {
	errs, ok := e.(validator.ValidationErrors)
	if ok {
		switch errs[0].Field() {
		case "Email":
			switch errs[0].Tag() {
			case "required":
				return errEmailRequired
			case "email":
				return errEmailInvalid
			}
		case "Password":
			switch errs[0].Tag() {
			case "required":
				return errPasswordRequired
			}
		}
	}
	return e
}
