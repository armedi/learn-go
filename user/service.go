package user

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

// Service represent the users' service
type Service interface {
	Register(user *RegisterRequest) error
	Login(email, password string) (token string, err error)
}

type userService struct {
	repo Repository
}

type userValidator struct {
	*userService
	validator *validator.Validate
}

// NewService create an object that represent the Service interface
func NewService(userRepo Repository) Service {
	return &userValidator{
		userService: &userService{
			repo: userRepo,
		},
		validator: validator.New(),
	}
}

func (us *userService) Register(user *RegisterRequest) error {
	if err := us.repo.Create(&User{
		Email:    user.Email,
		Password: user.Password,
	}); err != nil {
		return err
	}
	return nil
}

func (us *userService) Login(email, password string) (string, error) {
	return "", nil
}

func (uv *userValidator) Register(user *RegisterRequest) error {
	if err := uv.validator.Struct(user); err != nil {
		errs := err.(validator.ValidationErrors)
		switch errs[0].Field() {
		case "Email":
			switch errs[0].Tag() {
			case "required":
				return errors.New("Kolom email tidak boleh dikosongkan")
			case "email":
				return fmt.Errorf("Kolom email dengan isian %s tidak valid", errs[0].Value())
			}
		case "Password":
			switch errs[0].Tag() {
			case "required":
				return errors.New("Kolom password tidak boleh dikosongkan")
			}
		}
	}
	return uv.userService.Register(user)
}
