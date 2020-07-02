package user

import (
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"

	"github.com/armedi/learn-go/lib/jwt"
)

// Service represent the users' service
type Service interface {
	Register(form *RegisterRequest) error
	Login(form *LoginRequest) (token string, err error)
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

func (us *userService) Register(form *RegisterRequest) error {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return us.repo.Create(&User{
		Email:        form.Email,
		PasswordHash: string(hashedBytes),
	})
}

func (us *userService) Login(form *LoginRequest) (string, error) {
	user, err := us.authenticateLogin(form.Email, form.Password)
	if err != nil {
		return "", err
	}

	return jwt.CreateToken(user.ID)
}

func (us *userService) authenticateLogin(email string, password string) (*User, error) {
	user, err := us.repo.GetByEmail(email)

	if err != nil {
		switch err {
		case errRecordNotFound:
			return nil, errEmailPasswordMismatch
		default:
			return nil, err
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		switch err {
		case bcrypt.ErrMismatchedHashAndPassword:
			return nil, errEmailPasswordMismatch
		default:
			return nil, err
		}
	}

	return user, nil
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

func (uv *userValidator) Login(user *LoginRequest) (string, error) {
	if err := uv.validate.Struct(user); err != nil {
		return "", parseValidationError(err)
	}

	return uv.userService.Login(user)
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
	if err != nil {
		switch err {
		case errRecordNotFound:
			// Email address is not taken
			return nil
		default:
			return err
		}
	}

	// Email address is already taken
	return errEmailTaken
}
