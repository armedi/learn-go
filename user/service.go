package user

import (
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
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
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	if err := us.repo.Create(&User{
		Email:        user.Email,
		PasswordHash: string(hashedBytes),
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

	if err := runUserValFuncs(&User{Email: user.Email}, uv.emailIsAvailable); err != nil {
		return err
	}

	return uv.userService.Register(user)
}

type userValFunc func(*User) error

func runUserValFuncs(user *User, fns ...userValFunc) error {
	for _, fn := range fns {
		if err := fn(user); err != nil {
			return err
		}
	}
	return nil
}

func (uv *userValidator) emailIsAvailable(user *User) error {
	_, err := uv.repo.GetByEmail(user.Email)
	if err == errRecordNotFound {
		// Email address is not taken
		return nil
	}
	if err != nil {
		return err
	}
	// Email address is already taken
	return errEmailTaken
}
