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
	UserRepo Repository
	Validate *validator.Validate
}

// NewService create an object that represent the Service interface
func NewService(repo Repository) Service {
	return &userService{
		UserRepo: repo,
		Validate: validator.New(),
	}
}

func (us *userService) Register(form *RegisterRequest) error {
	if err := us.validateForm(form); err != nil {
		return err
	}

	user := User{
		Email:    form.Email,
		Password: form.Password,
	}

	if err := runUserValidationFuncs(user, us.emailIsAvailable); err != nil {
		return err
	}

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hashedBytes)

	return us.UserRepo.Create(&user)
}

func (us *userService) Login(form *LoginRequest) (string, error) {
	if err := us.validateForm(form); err != nil {
		return "", err
	}

	user, err := us.authenticateLogin(form.Email, form.Password)
	if err != nil {
		return "", err
	}

	return jwt.CreateToken(user.ID)
}

func (us *userService) authenticateLogin(email string, password string) (*User, error) {
	user, err := us.UserRepo.GetByEmail(email)

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

func (us *userService) validateForm(form interface{}) error {
	if err := us.Validate.Struct(form); err != nil {
		errs := err.(validator.ValidationErrors)
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
	return nil
}

type userValidationFunc func(User) error

func runUserValidationFuncs(user User, fns ...userValidationFunc) error {
	for _, fn := range fns {
		if err := fn(user); err != nil {
			return err
		}
	}
	return nil
}

func (us *userService) emailIsAvailable(user User) error {
	_, err := us.UserRepo.GetByEmail(user.Email)
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
