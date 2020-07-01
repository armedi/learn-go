package user

import (
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
	validate *validator.Validate
}

// NewService create an object that represent the Service interface
func NewService(userRepo Repository) Service {
	return &userValidator{
		userService: &userService{
			repo: userRepo,
		},
		validate: validator.New(),
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
	if err := uv.validate.Struct(user); err != nil {
		return parseValidationError(err)
	}
	return uv.userService.Register(user)
}
